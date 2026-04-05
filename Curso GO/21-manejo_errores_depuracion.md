# Manejo de Errores y Depuración en Go

Go toma una filosofía diferente: errores como valores, no excepciones. Esto hace el código más explícito y controlable.

## 1. Modelo de errores de Go

### Error interface

```go
// Cualquier tipo que implemente Error() es un error
type error interface {
    Error() string
}

// Casi todos devuelven (resultado, error)
func Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

func main() {
    resultado, err := Dividir(10, 2)
    if err != nil {
        // Algo salió mal
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("Resultado:", resultado)
}
```

### Verificar siempre errores

```go
// ❌ Mal: ignorar errores
datos, _ := os.ReadFile("archivo.txt")

// ✅ Bien: verificar siempre
datos, err := os.ReadFile("archivo.txt")
if err != nil {
    // Manejar error aquí
    fmt.Println("Error:", err)
    return
}
```

---

## 2. Crear errores personalizados

### errors.New() - simple

```go
import "errors"

func ValidarEdad(edad int) error {
    if edad < 0 {
        return errors.New("edad no puede ser negativa")
    }
    return nil
}

func main() {
    err := ValidarEdad(-5)
    if err != nil {
        fmt.Println(err.Error())  // "edad no puede ser negativa"
    }
}
```

### fmt.Errorf - con variables

```go
import "fmt"

func Conectar(host string, puerto int) error {
    return fmt.Errorf("no se puede conectar a %s:%d: conexión rechazada", host, puerto)
}

func main() {
    err := Conectar("localhost", 5432)
    fmt.Println(err)
    // "no se puede conectar a localhost:5432: conexión rechazada"
}
```

### Tipo error personalizado

```go
// Crear tipo error reutilizable
type ErrorValidacion struct {
    Campo   string
    Motivo  string
}

func (e ErrorValidacion) Error() string {
    return fmt.Sprintf("error en %s: %s", e.Campo, e.Motivo)
}

func ValidarEmail(email string) error {
    if !strings.Contains(email, "@") {
        return ErrorValidacion{
            Campo: "email",
            Motivo: "no contiene @",
        }
    }
    return nil
}

func main() {
    err := ValidarEmail("invalido")
    if err != nil {
        fmt.Println(err)  // "error en email: no contiene @"
    }
}
```

### ErrorInterface avanzado

```go
// Errores con más información
type ErrorConTexto struct {
    Operacion string
    Archivo   string
    Err       error
}

func (e *ErrorConTexto) Error() string {
    return fmt.Sprintf("%s %s: %v", e.Operacion, e.Archivo, e.Err)
}

func Leer(archivo string) error {
    _, err := os.Open(archivo)
    if err != nil {
        return &ErrorConTexto{
            Operacion: "leer",
            Archivo:   archivo,
            Err:       err,
        }
    }
    return nil
}

func main() {
    err := Leer("no_existe.txt")
    if err != nil {
        fmt.Println(err)
        // "leer no_existe.txt: open no_existe.txt: no such file or directory"
    }
}
```

---

## 3. Técnicas: Wrapping errors

Go 1.13+ permite "wrap" errores para mantener contexto.

```go
import (
    "errors"
    "fmt"
    "os"
)

func ProcesarArchivo(archivo string) error {
    datos, err := os.ReadFile(archivo)
    if err != nil {
        // Wrap error: preserva original + agrega contexto
        return fmt.Errorf("procesando %s: %w", archivo, err)
    }
    
    if len(datos) == 0 {
        return fmt.Errorf("archivo vacío")
    }
    
    return nil
}

func main() {
    err := ProcesarArchivo("no_existe.txt")
    if err != nil {
        fmt.Println(err)
        // "procesando no_existe.txt: open no_existe.txt: no such file or directory"
        
        // Buscar causa original
        var osErr *os.PathError
        if errors.As(err, &osErr) {
            fmt.Println("Error original es PathError")
        }
        
        // Verificar tipo de error
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("Archivo no existe")
        }
    }
}
```

**`%w` vs `%v`**:
- `%w`: Wrap (mantiene cadena de origen)
- `%v`: Simplemente convierte a string

---

## 4. Defer, Panic, Recover

### Defer - ejecutar al final

```go
func LeerArchivo(archivo string) error {
    f, err := os.Open(archivo)
    if err != nil {
        return err
    }
    defer f.Close()  // Se ejecuta al salir de la función
    
    // Leer archivo...
    return nil
}

// Múltiples defers: LIFO (stack)
func main() {
    defer fmt.Println("3. Segundo")
    defer fmt.Println("2. Primero")
    fmt.Println("1. Ahora")
}
// Output:
// 1. Ahora
// 2. Primero
// 3. Segundo
```

### Panic - error crítico

```go
// Panic es para errores REALMENTE serios
func DividirCrítico(a, b float64) float64 {
    if b == 0 {
        panic("división por cero es prohibida!")  // ⚠️ Mata el programa
    }
    return a / b
}

func main() {
    x := DividirCrítico(10, 2)
    fmt.Println(x)
}
```

**Cuándo usar panic**:
- ✅ Error de programación (programming)
- ✅ Condición imposible
- ❌ ~~Entrada inválida~~
- ❌ ~~Fallo de red~~
- ❌ ~~Archivo no existe~~

### Recover - atrapar panic

```go
func DividirSeguro(a, b float64) (resultado float64, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    
    if b == 0 {
        panic("división por cero")
    }
    
    return a / b, nil
}

func main() {
    resultado, err := DividirSeguro(10, 0)
    if err != nil {
        fmt.Println(err)  // "panic: división por cero"
    }
}
```

---

## 5. Patrones comunes

### Early return

```go
// ✅ Mejor legibilidad con early returns
func ValidarUsuario(u Usuario) error {
    if u.Nombre == "" {
        return errors.New("nombre vacío")
    }
    
    if u.Edad < 0 {
        return errors.New("edad no válida")
    }
    
    if !strings.Contains(u.Email, "@") {
        return errors.New("email inválido")
    }
    
    return nil  // Todo OK
}
```

### Encadenar errores

```go
func GuardarBD(usuario Usuario) error {
    // Validar
    if err := ValidarUsuario(usuario); err != nil {
        return fmt.Errorf("validación fallida: %w", err)
    }
    
    // Guardar en BD
    if err := baseDatos.Insertar(usuario); err != nil {
        return fmt.Errorf("no se puede guardar usuario: %w", err)
    }
    
    return nil
}

func main() {
    u := Usuario{"", 25, "test@email.com"}
    if err := GuardarBD(u); err != nil {
        fmt.Println(err)
        // "validación fallida: nombre vacío"
    }
}
```

### Try pattern (aunque Go no tiene try)

Simular try-catch de otros lenguajes:

```go
// Simulación: Opcional
type Result struct {
    Value interface{}
    Err   error
}

func Try(f func() (interface{}, error)) *Result {
    v, err := f()
    return &Result{v, err}
}

func (r *Result) Then(f func(interface{}) (interface{}, error)) *Result {
    if r.Err != nil {
        return r
    }
    v, err := f(r.Value)
    return &Result{v, err}
}

// Uso (cadena de operaciones)
resultado := Try(func() (interface{}, error) {
    return os.ReadFile("archivo.txt")
}).Then(func(v interface{}) (interface{}, error) {
    datos := v.([]byte)
    // Procesar...
    return datos, nil
})
```

---

## 6. Depuración con print statements

### Printf debugging

```go
import "fmt"

func ProcesarDatos(arr []int) int {
    fmt.Printf("DEBUG: entrada arr = %v\n", arr)
    
    suma := 0
    for i, v := range arr {
        fmt.Printf("DEBUG: i=%d, v=%d, suma parcial=%d\n", i, v, suma)
        suma += v
    }
    
    fmt.Printf("DEBUG: resultado final = %d\n", suma)
    return suma
}

func main() {
    resultado := ProcesarDatos([]int{1, 2, 3})
    fmt.Println("Resultado:", resultado)
}
```

### Log package

Mejor que fmt para logs en producción:

```go
import (
    "log"
    "os"
)

func main() {
    // Logger a archivo
    archivo, _ := os.Create("app.log")
    logger := log.New(archivo, "INFO: ", log.LstdFlags)
    
    logger.Println("Aplicación iniciada")
    
    // Logs con nivel
    loggerError := log.New(archivo, "ERROR: ", log.LstdFlags)
    loggerError.Println("Algo salió mal")
}
```

### Structured logging (recomendado)

```go
import (
    "encoding/json"
    "fmt"
)

// Log estructurado
type LogEntry struct {
    Nivel     string      `json:"nivel"`
    Mensaje   string      `json:"mensaje"`
    Contexto  interface{} `json:"contexto,omitempty"`
}

func Log(nivel, mensaje string, contexto interface{}) {
    entry := LogEntry{nivel, mensaje, contexto}
    jsonBytes, _ := json.Marshal(entry)
    fmt.Println(string(jsonBytes))
}

func main() {
    Log("INFO", "Usuario conectado", map[string]string{
        "usuario": "alice",
        "hora":    "10:30",
    })
    // {"nivel":"INFO","mensaje":"Usuario conectado","contexto":{"usuario":"alice","hora":"10:30"}}
}
```

---

## 7. Debugging avanzado

### Stack trace

```go
import (
    "fmt"
    "runtime"
)

func MostrarStack() {
    pc := make([]uintptr, 15)
    n := runtime.Callers(1, pc)
    
    for _, p := range pc[:n] {
        fn := runtime.FuncForPC(p)
        file, line := fn.FileLine(p)
        fmt.Printf("%s - %s:%d\n", fn.Name(), file, line)
    }
}

func FuncionC() {
    MostrarStack()
}

func FuncionB() {
    FuncionC()
}

func FuncionA() {
    FuncionB()
}

func main() {
    FuncionA()
    // Muestra: main → FuncionA → FuncionB → FuncionC → MostrarStack
}
```

### Debug helper

```go
import (
    "fmt"
    "path/filepath"
    "runtime"
)

// Helper para debugging
func DEBUG(v interface{}) {
    pc, file, line, _ := runtime.Caller(1)
    fn := runtime.FuncForPC(pc)
    
    filename := filepath.Base(file)
    fmt.Printf("[%s:%d %s] %T = %v\n", filename, line, fn.Name(), v, v)
}

func main() {
    x := 42
    DEBUG(x)
    // [main.go:18 main.main] int = 42
}
```

---

## 8. Pruebas unitarias (Testing)

Go tiene testing integrado:

```go
// archivo: math_test.go
// Comando: go test

import "testing"

func Sumar(a, b int) int {
    return a + b
}

func TestSumar(t *testing.T) {
    casos := []struct {
        a, b, esperado int
    }{
        {2, 3, 5},
        {-1, 1, 0},
        {0, 0, 0},
    }
    
    for _, c := range casos {
        resultado := Sumar(c.a, c.b)
        if resultado != c.esperado {
            t.Errorf("Sumar(%d, %d) = %d, esperado %d",
                c.a, c.b, resultado, c.esperado)
        }
    }
}

// Subtest
func TestSumarSubtests(t *testing.T) {
    t.Run("positivos", func(t *testing.T) {
        if Sumar(2, 3) != 5 {
            t.Fail()
        }
    })
    
    t.Run("negativos", func(t *testing.T) {
        if Sumar(-2, -3) != -5 {
            t.Fail()
        }
    })
}

// Benchmark
func BenchmarkSumar(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sumar(10, 20)
    }
}
```

Ejecutar:
```bash
go test          # Correr tests
go test -v       # Verbose
go test -run TestSumar  # Test específico
go test -bench .  # Benchmarks
```

---

## 9. Checklist: Manejo de errores

### ✅ Bien

```go
// 1. Siempre devolver error
func ProcesarDatos(entrada string) ([]int, error) {
    if entrada == "" {
        return nil, errors.New("entrada vacía")
    }
    // ...
    return resultado, nil
}

// 2. Verificar errores inmediatamente
datos, err := os.ReadFile("archivo.txt")
if err != nil {
    return fmt.Errorf("leer archivo: %w", err)
}

// 3. Usar defer para limpiar
func main() {
    archivo, _ := os.Open("file.txt")
    defer archivo.Close()
}

// 4. Wrap errors con contexto
return fmt.Errorf("operación fallida: %w", err)

// 5. Crear errores descriptivos
return fmt.Errorf("usuario con ID %d no encontrado", id)

// 6. Logging estructurado
log.Printf("usuario=%s, acción=%s, resultado=%s", user, action, result)
```

### ❌ Mal

```go
// 1. Ignorar errores
datos, _ := os.ReadFile("archivo.txt")  // ❌

// 2. Usar panic para error normal
if usuario == nil {
    panic("usuario no encontrado")  // ❌
}

// 3. Mensaje genérico
return errors.New("error")  // ❌

// 4. No hacer defer
f := os.Open("file.txt")
// ... y nunca cerrar

// 5. Printf en producción
fmt.Println("Usuario:", usuario)  // ❌ Mejor usar log

// 6. Debugging dejado en código
fmt.Println("DEBUG: x =", x)  // ❌
```

---

## 10. Herramientas de debugging

### go run -gcflags

```bash
# Debugging con nivel de optimización
go run -gcflags=-N -gcflags=-l programa.go
```

### Delve debugger

```bash
# Instalar
go install github.com/go-delve/delve/cmd/dlv@latest

# Debuggear
dlv debug

# Breakpoints
(dlv) break main.main
(dlv) continue
(dlv) print x
(dlv) next
(dlv) step
```

### Profiling

```bash
import _ "net/http/pprof"

func main() {
    go http.ListenAndServe("localhost:6060", nil)
    
    // Luego visitar:
    // http://localhost:6060/debug/pprof
    // http://localhost:6060/debug/pprof/heap
}
```

---

## 11. Conclusión - Errores son tu responsabilidad

Go tiene un mensaje claro: **Los errores son valores, no excepciones.**

Esto es filosoficamente diferente a Java/Python/JavaScript.

### Por qué Go eligió este camino

**Antes (Excepciones)**:
```java
try {
    data = leerArchivo()
    usuario = parsearJSON(data)
    conexion = conectarBD()
    resultado = consultar(conexion)
} catch (IOException e) {
    // ¿Qué pasó? ¿Dónde? ¿Qué hago?
} catch (JSONException e) {
    // ...
}
```

Caos. No sabes dónde falló. Incógnita profunda.

**Después (Go)**:
```go
data, err := leerArchivo()
if err != nil {
    return fmt.Errorf("leer archivo: %w", err)
}

usuario, err := parsearJSON(data)
if err != nil {
    return fmt.Errorf("parsear JSON: %w", err)
}

conexion, err := conectarBD()
if err != nil {
    return fmt.Errorf("conectar BD: %w", err)
}

resultado, err := consultar(conexion)
if err != nil {
    return fmt.Errorf("consultar: %w", err)
}
```

**Es verbose, ¿verdad?** Sí. Pero tu stack trace es **crystal clear**.

### La verdad sobre Go y errores

Sí, escribes `if err != nil` 1000 veces.

Pero:
1. **Sabes exactamente dónde puede fallar**
2. **Manejas cada caso específicamente**
3. **No hay sorpresas con excepciones inesperadas**
4. **El código es explícito y debuggeable**

### Comparación: Robustez de Errores

| Aspecto | Go | Java | Python |
|---------|-----|------|--------|
| Errores esperados | Explícitos | try/catch | try/except |
| Stack traces | Claros | Muy profundos | Variables |
| Recovery | defer/recover | finally | finally |
| Causa root | Si wrappeas | Difícil rastrear | Difícil rastrear |
| Performance | O(1) | O(n) throw | O(n) raise |

### El cambio mental

Si vienes de excepciones:

**Viejo (Python)**:
```python
try:
    resultado = operacion_peligrosa()
    usar(resultado)
except Exception as e:
    log.error(f"Failed: {e}")
```

**Nuevo (Go)**:
```go
resultado, err := operacionPeligrosa()
if err != nil {
    return fmt.Errorf("operacion peligrosa: %w", err)
}
usar(resultado)
```

Go **força** que manejes errores. Es incómodo al principio. Luego es tu superpoder.

### Patrón final: El Wrap idiomático

```go
// Nivel bajo
func BuscarDB(id int) (*Usuario, error) {
    resultados, err_sql := db.Query("SELECT * WHERE id = ?", id)
    if err_sql != nil {
        return nil, err_sql  // ⚠️ Puede perder contexto
    }
    // ...
}

// Nivel medio
func ObtenerUsuario(id int) (*Usuario, error) {
    usuario, err := BuscarDB(id)
    if err != nil {
        return nil, fmt.Errorf("buscar en DB: %w", err)
    }
    return usuario, nil
}

// Nivel alto
func Handler(w http.ResponseWriter, r *http.Request) {
    usuario, err := ObtenerUsuario(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    // ...
}
```

Cada nivel agrega contexto. El usuario final ve:
```
buscar en DB: sql: no rows in result set
```

Claro, específico, actionable.

### Debugging mental

**Si tu programa falla**:

1. Lee el stack trace (Go da buenos)
2. Busca el `if err != nil` donde falló
3. Mira el `fmt.Errorf` para contexto
4. Usa `errors.Is()` / `errors.As()` para tipo específico
5. Arregla el root cause

**No hay magia oculta. Solo errores claros.**

### La verdad sobre `if err != nil`

Sí, escribirlo 1000 veces es tedioso.

Pero:
- Es **explícito**: No hay errores silenciosos
- Es **seguro**: No hay panic inesperado
- Es **debuggeable**: Stack traces claros
- Es **mantenible**: Futuro yo entiende qué puede fallar
- Es **Go**: Asi se hace. Acéptalo. Domínalo.

### Filosofía final

> "Explicit is better than implicit" - Go Proverbs

Los errores **no deben ser silenciosos**. Go força que pienses en ellos.

Es incómodo. Es correcto.

**Domina manejo de errores en Go y podrás escribir sistemas fiables en cualquier lenguaje.**
