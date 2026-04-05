# Tipos de Datos en Go - Guía Completa

Entender los tipos de datos es lo más fundamental en la programación. Go es un lenguaje fuertemente tipado, lo que significa que cada variable tiene un tipo definido en tiempo de compilación. Esto nos da seguridad: el compilador verifica que usemos los tipos correctamente.

---

## Parte I: Tipos Básicos

Los tipos básicos son los bloques de construcción fundamentales. Si dominas estos, puedes construir cualquier cosa.

### 1. Bool - Valores de Verdad

```go
var activo bool = true
var deshabilitado bool = false
```

**¿Qué es?**
Un `bool` es el tipo más simple: simplemente puede ser `true` (verdadero) o `false` (falso). No hay nada intermedio.

**¿Cuándo lo usas?**
- Banderas y estados: `usuarioActivo`, `esPrimo`, `esVálido`
- Condiciones: `if x > 5`
- Iteraciones: `while (condición)`

**Operadores lógicos**:
```go
a && b   // AND (y)
a || b   // OR (o)
!a       // NOT (no)
```

**Ejemplo práctico**:
```go
func EsMayorDeEdad(edad int) bool {
    return edad >= 18
}

func main() {
    esMayorEdad := EsMayorDeEdad(25)
    if esMayorEdad {
        fmt.Println("Puede votar")
    }
}
```

---

### 2. Tipos Numéricos Enteros (Integers)

Los enteros son números sin decimales. Go tiene varias opciones porque diferentes situaciones necesitan diferentes tamaños.

#### Enteros con signo: `int`, `int8`, `int16`, `int32`, `int64`

```go
var contador int = 42
var temperatura int8 = -50    // -128 a 127
var año int32 = 2026
var distancia int64 = 9999999999
```

**Tamaños y rangos**:

| Tipo | Rango | Uso |
|------|-------|-----|
| `int8` | -128 a 127 | Datos muy pequeños |
| `int16` | -32,768 a 32,767 | Temperaturas, ángulos |
| `int32` | -2M a 2M aprox | Universal para muchos casos |
| `int64` | -9 billones a 9 billones | Muy grandes números, timestamps |
| `int` | Depende de SO (32 o 64 bits) | Estándar, mejor opción |

**¿Cuál usar?**
- **`int`**: El 95% de las veces. Go elige automáticamente 32 o 64 bits según tu máquina.
- **`int64`**: Cuando necesites números gigantes o timestamp en milisegundos
- **`int32`, `int8`, `int16`**: Cuando necesites compatibilidad externa o ahorrar memoria en arrays gigantes

```go
func main() {
    var contador int = 0
    var temperatura int8 = 25
    var años int64 = 10000000000  // 10 mil millones
}
```

#### Enteros sin signo: `uint`, `uint8`, `uint16`, `uint32`, `uint64`

```go
var edad uint = 25          // Nunca negativa
var r uint8 = 255          // Ese es el máximo
var pixeles uint32 = 1920
```

**¿Cuándo usar?**
- **`uint8` (byte)**: Para bytes individuales, valores 0-255
- **`uint`**: Para conteos, índices, longitudes
- Generalmente: no necesitas unsigned a menos que explícitamente haya valores negativos prohibidos

```go
// ✅ Bien: edad nunca es negativa
func ValidarEdad(edad uint) {
    if edad > 120 {
        fmt.Println("Edad no realista")
    }
}

// ❌ No sería malo, pero int también funciona
func ValidarEdad(edad int) {
    if edad < 0 || edad > 120 {
        fmt.Println("Inválido")
    }
}
```

---

### 3. Tipos de Coma Flotante (Decimales)

Para números con parte fraccionaria.

```go
var pi float64 = 3.14159
var temperatura float32 = 98.6
```

**Comparación**:

| Tipo | Precisión | Uso |
|------|-----------|-----|
| `float32` | ~7 dígitos | Gráficos, cuando ahorro de memoria importa |
| `float64` | ~15 dígitos | Estándar: ciencia, finanzas, cualquier precisión |

**Ejemplo práctico**:
```go
func CalcularPromedio(calificaciones []float64) float64 {
    suma := 0.0
    for _, cal := range calificaciones {
        suma += cal
    }
    return suma / float64(len(calificaciones))
}

func main() {
    califs := []float64{85.5, 90.0, 78.5, 92.0}
    promedio := CalcularPromedio(califs)
    fmt.Printf("Promedio: %.2f\n", promedio)  // 86.50
}
```

**Cuidado con floats**:
```go
// Los floats no son exactos
x := 0.1 + 0.2
y := 0.3
fmt.Println(x == y)  // false! (porque x = 0.30000000000000004)

// Usar comparación aproximada
func AproximadoIgual(a, b, tolerancia float64) bool {
    diferencia := a - b
    if diferencia < 0 {
        diferencia = -diferencia
    }
    return diferencia < tolerancia
}
```

---

### 4. Rune y Byte - Caracteres Individuales

Go maneja texto de forma especial. Necesitas entender `byte` y `rune`.

#### `byte` - Un carácter ASCII (0-255)

```go
var caracter byte = 'A'  // 65
var numero byte = 48     // '0'
```

Un `byte` es un número 0-255 que representa un carácter en ASCII simple.

#### `rune` - Un carácter Unicode

```go
var letra rune = 'A'      // 65 (mismo que byte para ASCII)
var emoji rune = '😀'      // 128512 (Unicode)
var chino rune = '中'      // 20013 (Unicode)
```

Un `rune` es un carácter Unicode completo (puede ser emoji, caracteres asiáticos, etc.).

**Diferencia fundamental**:

```go
texto := "Hola"
fmt.Println(len(texto))           // 4 (4 bytes)

textoUnicode := "こんにちは"      // Japonés
fmt.Println(len(textoUnicode))    // 15 bytes (!!! no 5 caracteres)

// Convertir a runes para contar caracteres reales
runes := []rune(textoUnicode)
fmt.Println(len(runes))           // 5 caracteres
```

**Ejemplo práctico**:
```go
func ContarCaracteres(texto string) int {
    return len([]rune(texto))  // Convertir a runes primero
}

func main() {
    fmt.Println(ContarCaracteres("Hola"))        // 4
    fmt.Println(ContarCaracteres("Hel日本"))     // 6
    fmt.Println(ContarCaracteres("你好"))        // 2
}
```

---

### 5. String - Texto

```go
var nombre string = "Alice"
var mensaje string = "Hola, mundo!"
var vacio string = ""
```

**¿Qué es?**
Un `string` es una secuencia inmutable de bytes UTF-8 que representa texto. **Inmutable** significa que no puedes cambiarla después de crearla.

**Características**:
- Inmutable: `s[0] = 'x'` ❌ ILLEGAL
- Codificación UTF-8 automática
- Comparables: puedes usar `==`, `<`, `>`

**Operaciones comunes**:

```go
import "strings"

s := "Hola"

// Concatenación
resultado := s + " Mundo"      // "Hola Mundo"

// Longitud
fmt.Println(len(s))            // 4

// Acceso a un byte específico
fmt.Println(s[0])              // 72 (byte 'H')

// Buscar substring
strings.Contains(s, "ol")      // true

// Convertir a minúsculas/mayúsculas
strings.ToLower(s)             // "hola"
strings.ToUpper(s)             // "HOLA"

// Dividir
partes := strings.Split("a,b,c", ",")  // ["a", "b", "c"]

// Unir
resultado := strings.Join([]string{"a", "b", "c"}, "-")  // "a-b-c"

// Remover espacios
strings.TrimSpace("  hola  ")  // "hola"

// Reemplazar
strings.ReplaceAll("aaa", "a", "b")  // "bbb"
```

**Conversiones**:
```go
// String a byte slice
bytes := []byte("Hola")

// Byte slice a string
texto := string([]byte{72, 111, 108, 97})  // "Hola"

// String a runes
runes := []rune("Hola")  // ['H', 'o', 'l', 'a']
```

---

### 6. Números Complejos (Avanzado)

Para matemática compleja; rara vez necesarios.

```go
var c complex128 = 3 + 4i
var c2 complex64 = 1 + 2i
```

```go
c := 3 + 4i
fmt.Println(real(c))   // 3
fmt.Println(imag(c))   // 4
fmt.Println(abs(c))    // 5 (√(3² + 4²))
```

**Cuándo**: Procesamiento de señales, transformadas de Fourier.

---

### 7. Tipos Especiales

#### `uintptr` - Direcciones de memoria

Para operaciones de bajo nivel con memoria (unsafe):

```go
var dirección uintptr
// Rara vez necesitas esto
```

---

## Parte II: Tipos Compuestos

Ahora que entiendes los tipos básicos, podemos combinarlos para crear estructuras más complejas.

### 1. Array - Colección Fija

```go
var notas [5]float64
notas[0] = 85.5
notas[1] = 90.0

// O con inicialización
estudiantes := [3]string{"Alice", "Bob", "Charlie"}
```

**¿Qué es?**
Un array es una colección de elementos **del mismo tipo** con **tamaño fijo** decidido en tiempo de compilación.

**Características**:
- Tamaño fijo: `[5]int` siempre tiene 5 elementos
- Índice 0: el primero siempre es índice 0
- Tipo homogéneo: todos los elementos deben ser del mismo tipo
- Valor vs referencia: se pasan por valor (copian datos)

**Ejemplo**:
```go
func CalcularPromedio(calificaciones [5]float64) float64 {
    suma := 0.0
    for _, cal := range calificaciones {
        suma += cal
    }
    return suma / 5
}

func main() {
    mis_califs := [5]float64{85.5, 90.0, 78.5, 92.0, 88.5}
    promedio := CalcularPromedio(mis_califs)
    fmt.Printf("Promedio: %.2f\n", promedio)
}
```

---

### 2. Slice - Colección Dinámica

```go
var numeros []int              // Slice vacío
numeros = append(numeros, 1, 2, 3)

calificaciones := []float64{85.5, 90.0, 78.5}
```

**¿Qué es?**
Un slice es una **vista dinámica** de un array. Puedes cambiar su tamaño, pero internamente apunta a un array subyacente.

**Slice vs Array**:

| Aspecto | Array | Slice |
|---------|-------|-------|
| Tamaño | Fijo | Dinámico |
| Sintaxis | `[5]int` | `[]int` |
| Pase | Copia datos | Pasa referencia |
| `append` | ❌ No | ✅ Sí |
| Uso común | 5% | 95% de los casos |

**Operaciones**:
```go
s := []int{1, 2, 3, 4, 5}

// Acceso
fmt.Println(s[0])              // 1
fmt.Println(s[2])              // 3

// Length (tamaño actual)
fmt.Println(len(s))            // 5

// Capacity (espacio reservado)
fmt.Println(cap(s))            // 5 o más

// Subslice (slice de slice)
fmt.Println(s[1:3])            // [2, 3] (desde índice 1 hasta 3 exclusive)
fmt.Println(s[:2])             // [1, 2]
fmt.Println(s[2:])             // [3, 4, 5]

// Agregar elementos
s = append(s, 6, 7)            // [1, 2, 3, 4, 5, 6, 7]

// Crear slice con capacidad reservada
s2 := make([]int, 0, 10)       // longitud 0, capacidad 10
```

**Ejemplo práctico**:
```go
func FiltrarPares(numeros []int) []int {
    var resultado []int
    for _, n := range numeros {
        if n%2 == 0 {
            resultado = append(resultado, n)
        }
    }
    return resultado
}

func main() {
    entrada := []int{1, 2, 3, 4, 5, 6}
    pares := FiltrarPares(entrada)
    fmt.Println(pares)  // [2, 4, 6]
}
```

---

### 3. Map - Diccionario / Hash Table

```go
edades := map[string]int{
    "Alice": 30,
    "Bob":   25,
}

// Acceso
fmt.Println(edades["Alice"])   // 30

// Agregar
edades["Charlie"] = 35

// Verificar existencia
edad, existe := edades["Diana"]
if existe {
    fmt.Println(edad)
} else {
    fmt.Println("No encontrado")
}

// Eliminar
delete(edades, "Bob")
```

**¿Qué es?**
Un map es una tabla hash que almacena pares clave-valor. Es extremadamente rápido para búsquedas.

**Características**:
- Claves únicas: cada clave aparece una sola vez
- Búsqueda O(1): acceso instantáneo
- Dinámico: puedes agregar/eliminar en cualquier momento
- Desordenado: el orden de iteración es aleatorio

**Casos de uso**:
```go
// Contar frecuencias
func ContarPalabras(texto string) map[string]int {
    palabras := strings.Fields(texto)
    conteos := make(map[string]int)
    for _, palabra := range palabras {
        conteos[palabra]++
    }
    return conteos
}

// Índice de usuarios
usuarios := make(map[int]string)  // ID → Nombre
usuarios[1] = "Alice"
usuarios[2] = "Bob"

// Configuración
config := map[string]string{
    "host": "localhost",
    "puerto": "5432",
}
```

---

### 4. Struct - Objeto Con Campos

```go
type Persona struct {
    Nombre string
    Edad   int
    Email  string
}

p := Persona{"Alice", 30, "alice@example.com"}
// O con nombres
p2 := Persona{
    Nombre: "Bob",
    Edad:   25,
    Email:  "bob@example.com",
}
```

**¿Qué es?**
Un struct agrupa múltiples campos de diferentes tipos bajo un solo nombre. Es como un "objeto" que contiene datos relacionados.

**Ejemplo de dominio real**:
```go
type Cuenta struct {
    Titular    string
    Saldo      float64
    NumeroCuenta string
    Activa     bool
}

type Banco struct {
    Nombre     string
    Cuentas    []Cuenta
    Capital    float64
}

func main() {
    c1 := Cuenta{
        Titular: "Alice",
        Saldo: 1500.00,
        NumeroCuenta: "001-234",
        Activa: true,
    }
    
    c2 := Cuenta{
        Titular: "Bob",
        Saldo: 2500.00,
        NumeroCuenta: "001-235",
        Activa: true,
    }
    
    banco := Banco{
        Nombre: "MiBanco",
        Cuentas: []Cuenta{c1, c2},
        Capital: 1000000,
    }
    
    fmt.Printf("Banco: %s\n", banco.Nombre)
    fmt.Printf("Cuenta 1: %+v\n", banco.Cuentas[0])
}
```

---

### 5. Pointer - Referencia a Memoria

```go
x := 42
px := &x              // px apunta a x

fmt.Println(*px)      // 42 (dereferenciar)
*px = 100             // Cambiar el valor

fmt.Println(x)        // 100 (cambió!)
```

**¿Qué es?**
Un pointer (puntero) es una dirección de memoria. Usando `&` obtenemos dirección, usando `*` accedemos al valor.

**Cuándo usar**:
- Pasar por referencia en funciones
- Evitar copiar estructuras grandes
- Implementar métodos que mutan el receptor

**Ejemplo práctico**:
```go
// Función que necesita modificar el original
func Depositar(c *Cuenta, monto float64) {
    if monto > 0 {
        c.Saldo += monto
    }
}

// Función que NO necesita modificar
func ObtenerSaldo(c Cuenta) float64 {
    return c.Saldo
}

func main() {
    cuenta := Cuenta{
        Titular: "Alice",
        Saldo: 1000.00,
    }
    
    Depositar(&cuenta, 500)
    fmt.Println(cuenta.Saldo)  // 1500.00
}
```

---

### 6. Function Type - Funciones como Valores

```go
// Definir un tipo function
type Operacion func(int, int) int

func Sumar(a, b int) int {
    return a + b
}

func Restar(a, b int) int {
    return a - b
}

func Aplicar(op Operacion, x, y int) int {
    return op(x, y)
}

func main() {
    var suma Operacion = Sumar
    resultado := Aplicar(suma, 5, 3)
    fmt.Println(resultado)  // 8
}
```

---

### 7. Interface - Contrato de Métodos

```go
type Escritor interface {
    Escribir(texto string) error
}

// Cualquier tipo que implemente Escribir() satisface Escritor
```

Una interface define qué métodos debe tener un tipo. Es un contrato.

**Ejemplo completo**:
```go
type Dispositivo interface {
    Encender()
    Apagar()
    Mostrar(texto string)
}

type Pantalla struct {
    Modelo string
}

func (p *Pantalla) Encender() {
    fmt.Println("Pantalla encendida")
}

func (p *Pantalla) Apagar() {
    fmt.Println("Pantalla apagada")
}

func (p *Pantalla) Mostrar(texto string) {
    fmt.Println("Mostrando:", texto)
}

type Parlante struct {
    Marca string
}

func (pa *Parlante) Encender() {
    fmt.Println("Parlante encendido")
}

func (pa *Parlante) Apagar() {
    fmt.Println("Parlante apagado")
}

func (pa *Parlante) Mostrar(texto string) {
    fmt.Println("Sonando:", texto)
}

func main() {
    // Ambos satisfacen Dispositivo
    dispositivos := []Dispositivo{
        &Pantalla{Modelo: "Samsung"},
        &Parlante{Marca: "Sony"},
    }
    
    for _, d := range dispositivos {
        d.Encender()
        d.Mostrar("Hola")
        d.Apagar()
    }
}
```

---

### 8. Channel - Comunicación Concurrente

```go
c := make(chan int)
go func() {
    c <- 42  // Enviar valor
}()
valor := <-c  // Recibir valor
```

Channels permiten goroutines (threads ligeros) comunicarse de forma segura.

```go
import (
    "fmt"
    "time"
)

func main() {
    respuestas := make(chan string)
    
    // Goroutine 1
    go func() {
        time.Sleep(1 * time.Second)
        respuestas <- "Respuesta 1"
    }()
    
    // Goroutine 2
    go func() {
        time.Sleep(2 * time.Second)
        respuestas <- "Respuesta 2"
    }()
    
    // Esperar respuestas
    fmt.Println(<-respuestas)  // Respuesta 1
    fmt.Println(<-respuestas)  // Respuesta 2
}
```

---

## Parte III: Tipos Especiales

### 1. `time.Time` - Fecha y Hora (⭐ IMPORTANTE)

```go
import "time"

ahora := time.Now()
fmt.Println(ahora)  // 2026-04-04 14:30:45.123456789 +0000 UTC

// Crear fecha específica
fecha := time.Date(2026, 12, 25, 15, 30, 0, 0, time.UTC)
fmt.Println(fecha)
```

**Operaciones comunes**:
```go
import "time"

ahora := time.Now()

// Componentes
ahora.Year()       // 2026
ahora.Month()      // April (4)
ahora.Day()        // 4
ahora.Hour()       // 14
ahora.Minute()     // 30
ahora.Second()     // 45

// Agregar/restar tiempo
mañana := ahora.AddDate(0, 0, 1)              // +1 día
dentro_una_semana := ahora.Add(7 * time.Day) // +7 días
en_dos_horas := ahora.Add(2 * time.Hour)     // +2 horas

// Diferencia
diferencia := mañana.Sub(ahora)
fmt.Println(diferencia)                      // 24h0m0s

// Formato
ahora.Format("2006-01-02")        // "2026-04-04"
ahora.Format("2006-01-02 15:04") // "2026-04-04 14:30"

// Parse (string a Time)
fecha, _ := time.Parse("2006-01-02", "2026-04-04")

// Comparar
if mañana.After(ahora) {
    fmt.Println("mañana es después de ahora")
}
```

**Durations** (intervalos de tiempo):
```go
d := 5 * time.Second
d2 := 2 * time.Minute
d3 := 1 * time.Hour

// Convertir
segundos := int64(d.Seconds())       // 5
milisegundos := d.Milliseconds()     // 5000
```

**Ejemplo práctico - Registrar tiempo de ejecución**:
```go
func main() {
    inicio := time.Now()
    
    // Hacer algo
    time.Sleep(2 * time.Second)
    
    duracion := time.Since(inicio)
    fmt.Printf("Tomó: %v\n", duracion)  // Tomó: 2.000123456s
}
```

---

### 2. `error` - Manejo de Errores

```go
import "errors"

func Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

func main() {
    resultado, err := Dividir(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Resultado:", resultado)
}
```

---

### 3. `interface{}` - Tipo Universal

```go
var x interface{} = "cadena"
var y interface{} = 42
var z interface{} = []int{1, 2, 3}
```

`interface{}` significa "cualquier tipo". Se usa cuando necesitas flexibilidad extrema.

```go
func Imprimir(valor interface{}) {
    switch v := valor.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Entero:", v)
    case []int:
        fmt.Println("Slice:", v)
    default:
        fmt.Println("Desconocido:", v)
    }
}

func main() {
    Imprimir("Hola")
    Imprimir(42)
    Imprimir([]int{1, 2, 3})
}
```

---

## Parte IV: Ejercicios Prácticos Integrados

### Ejercicio 1: Sistema de Biblioteca

```go
package main

import (
    "fmt"
    "time"
)

type Libro struct {
    Titulo    string
    Autor     string
    Publicado int
    Disponible bool
}

type Biblioteca struct {
    Nombre string
    Libros []Libro
    Abierta bool
}

func (b *Biblioteca) AgregarLibro(libro Libro) {
    b.Libros = append(b.Libros, libro)
}

func (b *Biblioteca) BuscarPorAutor(autor string) []Libro {
    var resultados []Libro
    for _, libro := range b.Libros {
        if libro.Autor == autor {
            resultados = append(resultados, libro)
        }
    }
    return resultados
}

func (b *Biblioteca) Prestar(titulo string) error {
    for i, libro := range b.Libros {
        if libro.Titulo == titulo {
            if !libro.Disponible {
                return fmt.Errorf("libro '%s' no disponible", titulo)
            }
            b.Libros[i].Disponible = false
            return nil
        }
    }
    return fmt.Errorf("libro '%s' no encontrado", titulo)
}

func (b *Biblioteca) Devolver(titulo string) error {
    for i, libro := range b.Libros {
        if libro.Titulo == titulo {
            b.Libros[i].Disponible = true
            return nil
        }
    }
    return fmt.Errorf("libro '%s' no encontrado", titulo)
}

func main() {
    biblioteca := Biblioteca{
        Nombre: "Biblioteca Central",
        Abierta: true,
    }
    
    biblioteca.AgregarLibro(Libro{
        Titulo: "1984",
        Autor: "George Orwell",
        Publicado: 1949,
        Disponible: true,
    })
    
    biblioteca.AgregarLibro(Libro{
        Titulo: "Fundación",
        Autor: "Isaac Asimov",
        Publicado: 1951,
        Disponible: true,
    })
    
    biblioteca.AgregarLibro(Libro{
        Titulo: "Yo, Robot",
        Autor: "Isaac Asimov",
        Publicado: 1950,
        Disponible: true,
    })
    
    fmt.Println("Libros de Isaac Asimov:")
    librosAsimov := biblioteca.BuscarPorAutor("Isaac Asimov")
    for _, libro := range librosAsimov {
        fmt.Printf("- %s (%d)\n", libro.Titulo, libro.Publicado)
    }
    
    fmt.Println("\nPrestando '1984'...")
    if err := biblioteca.Prestar("1984"); err == nil {
        fmt.Println("Préstamo exitoso")
    }
    
    fmt.Println("Intentando prestar '1984' nuevamente...")
    if err := biblioteca.Prestar("1984"); err != nil {
        fmt.Println("Error:", err)
    }
    
    fmt.Println("\nDevolviendo '1984'...")
    biblioteca.Devolver("1984")
    fmt.Println("Devolución completada")
}
```

---

## Resumen de Tipos

| Tipo | Propósito | Mutable | Ejemplo |
|------|-----------|---------|---------|
| `bool` | Verdadero/falso | - | `true` |
| `int`, `float64` | Números | - | `42`, `3.14` |
| `string` | Texto | ❌ No | `"Hola"` |
| `[5]int` | Array fijo | ✅ Sí | Tamaño conocido |
| `[]int` | Slice dinámico | ✅ Sí | Tamaño flexible |
| `map[string]int` | Tabla hash | ✅ Sí | Búsqueda O(1) |
| `struct` | Datos agrupados | ✅ Sí | Objetos |
| `*T` | Referencia | ✅ Sí | Puntero a T |
| `func(int)int` | Función | - | Callable |
| `interface{}` | Cualquier tipo | - | Flexible |
| `time.Time` | Fecha/hora | - | Dates |
| `error` | Manejo errores | - | Errores |

---

## Conclusión

Los tipos en Go son:
- **Explícitos**: El compilador verifica todo
- **Seguros**: Errores de tipo se detectan antes de ejecutar
- **Claros**: Código self-documenting

Domina estos tipos y podrás construir cualquier programa en Go.