# Context, Timeouts y Cancelación en Go

> **Cómo controlar el ciclo de vida de operaciones, evitar que se cuelguen y manejar cancelaciones de forma elegante**

---

## ¿Por qué existe Context?

Imagina un escenario real: tu servidor web recibe una solicitud del cliente, inicia una consulta a BD que tardará 30 segundos. Pero el cliente se aburre y **cierra la conexión después de 5 segundos**.

El servidor sigue gastando recursos en una consulta que **nadie más quiere**. Pior aún, si el cliente cierra la conexión y ves a una goroutine haciendo trabajo inútil, eso es un **memory leak**.

**Context** fue inventado exactamente para esto:
- 🎯 **Cancelar operaciones** cuando el cliente se va
- ⏱️ **Poner timeouts** a operaciones largas
- 🔗 **Pasar valores** entre funciones
- 📡 **Propagandel información** a lo largo de la pila de llamadas

Sin context, tu código es **frágil**: puede colgarse, puede perder recursos, es impredecible en producción.

Con context, tu código es **robusto**: cancela lo que no necesita, respeta deadlines, es profesional.

---

## Parte I: ¿Qué es Context?

### El objeto context.Context

Context es una **interfaz simple**:

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

**¿Qué significa?**
- `Deadline()` - ¿Cuándo debe terminar esta operación?
- `Done()` - Un canal que se cierra cuando debe cancelarse
- `Err()` - ¿Por qué se canceló? (context.Canceled, context.DeadlineExceeded)
- `Value()` - Acceso a datos contextuales

### Context es una cadena

```
context.Background()  ← Raíz (sin deadline)
    ↓
context.WithTimeout() ← Heredar con timeout
    ↓
context.WithCancel()  ← Heredar con cancelación manual
    ↓
(la función que recibe el context)
```

**Principio 1:** El contexto fluye como una cadena. El padre cancela, todos los hijos se cancelan.

### Context.Background()

El contexto "raíz", sin timeout ni cancelación:

```go
ctx := context.Background()

// Propiedades:
ctx.Deadline()    // (time.Time{}, false) - sin deadline
ctx.Err()         // nil - nunca está cancelado
<-ctx.Done()      // nunca se cierra
```

**Cuándo usarlo:**
- En `func main()`
- En servidores (punto de entrada de solicitudes)
- En tests

### Context.TODO()

Parecido a Background(), pero **comunica intención**:

```go
ctx := context.TODO()  // "Aún no sé qué contexto poner"
```

**Cuándo usarlo:**
- Cuando estás refactorizando
- Cuando no estás seguro del contexto adecuado

---

## Parte II: Timeouts (Plazos)

### Operación sin timeout (❌ Malo)

```go
func ConsultarBD() error {
    // Si la BD se cuelga, ¡tu servidor se cuelga!
    conn := abrirConexion()
    datos := consultar(conn)
    return procesarDatos(datos)
}

func main() {
    // ❌ Se cuelga indefinidamente si BD falla
    ConsultarBD()
}
```

**Problema:** Si la BD se cae, el servidor espera **para siempre**.

### Con Timeout (✅ Bien)

```go
func ConsultarBD(ctx context.Context) error {
    // Si tarda más de 5 segundos, se cancela
    conn, err := abrirConexion(ctx)
    if err != nil {
        return err
    }

    datos := consultar(ctx, conn)
    return procesarDatos(datos)
}

func main() {
    // Timeout de 5 segundos
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()  // Limpiar recursos

    if err := ConsultarBD(ctx); err != nil {
        fmt.Println("Error:", err)
    }
}
```

**¿Qué pasó?**
1. Creamos un contexto con timeout de 5 segundos
2. Después de 5 segundos, `ctx.Done()` se cierra automáticamente
3. Todas las operaciones que usan `ctx` notan que se canceló
4. `defer cancel()` limpia recursos

### Estructura perfecta con timeout

```go
func manejarSolicitud(w http.ResponseWriter, r *http.Request) {
    // ✅ PATRÓN: Crear contexto con timeout de solicitud
    ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
    defer cancel()

    // Pasar ctx a todas las operaciones
    usuario, err := obtenerUsuario(ctx, r.URL.Query().Get("id"))
    if err == context.DeadlineExceeded {
        w.WriteHeader(http.StatusRequestTimeout)
        fmt.Fprintf(w, "La solicitud tardó demasiado")
        return
    }

    // ... responder
}

func obtenerUsuario(ctx context.Context, id string) (*Usuario, error) {
    // Respetar el contexto en TODAS las operaciones
    datos, err := obtenerDeBD(ctx, id)
    if err != nil {
        return nil, err
    }
    return datos, nil
}
```

### Errores de timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

time.Sleep(2 * time.Second)  // Exceeds timeout

err := ctx.Err()
if err == context.DeadlineExceeded {
    fmt.Println("Se excedió el deadline")
}
if err == context.Canceled {
    fmt.Println("Se canceló")
}
```

---

## Parte III: Cancelación Manual

### WithCancel - cancelar manualmente

```go
func descargarDatos(ctx context.Context, url string) {
    // Hacer algo largo
    for i := 0; i < 100; i++ {
        select {
        case <-ctx.Done():
            fmt.Println("Cancelado:", ctx.Err())
            return
        default:
            // Procesar dato i
            time.Sleep(100 * time.Millisecond)
        }
    }
    fmt.Println("Descarga completada")
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go descargarDatos(ctx, "http://example.com")

    time.Sleep(500 * time.Millisecond)
    cancel()  // ← Cancelar la descarga

    time.Sleep(100 * time.Millisecond)
}
```

**Output:**
```
Cancelado: context canceled
```

### Patrón profesional: Cleanup con cancelación

```go
type Recurso struct {
    conn *Connection
}

func CrearRecurso(ctx context.Context) (*Recurso, error) {
    conn, err := abrirConexion()
    if err != nil {
        return nil, err
    }

    r := &Recurso{conn: conn}

    // Si contexto se cancela, cerrar conexión
    go func() {
        <-ctx.Done()
        r.Cerrar()  // Cleanup automático
    }()

    return r, nil
}

func (r *Recurso) Cerrar() {
    r.conn.Close()
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    recurso, _ := CrearRecurso(ctx)
    defer cancel()

    // Usar recurso...
    // Al salir de main, cancel() cierra recurso automáticamente
}
```

---

## Parte IV: Pasar Valores en Context

### context.WithValue() (usar con cuidado)

```go
// ❌ Evitar: keys string (pueden colisionar)
ctx := context.WithValue(context.Background(), "user_id", 123)
id := ctx.Value("user_id")  // ¿Qué tipo es?

// ✅ Bien: keys privadas (evita colisiones)
type userIDKey struct{}

ctx := context.WithValue(context.Background(), userIDKey{}, 123)
id := ctx.Value(userIDKey{}).(int)  // Type-safe
```

### Caso real: Pasar ID de solicitud

```go
type requestIDKey struct{}

func manejarSolicitud(w http.ResponseWriter, r *http.Request) {
    // Generar ID único para esta solicitud
    requestID := uuid.New().String()
    
    // Pasar a través del contexto
    ctx := context.WithValue(r.Context(), requestIDKey{}, requestID)
    ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    procesarDatos(ctx)
}

func procesarDatos(ctx context.Context) {
    requestID := ctx.Value(requestIDKey{}).(string)
    
    // Usar requestID para logging
    log.Printf("[%s] Procesando datos", requestID)
}
```

**✅ Ventaja:** El request ID se propaga automáticamente sin parámetros adicionales.

---

## Parte V: Patrón Profesional - HTTP Handler con Context

### ✅ Estructura recomendada

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

// Tipos de contexto (evita colisiones)
type requestIDKey struct{}
type userIDKey struct{}

// Handler que respeta timeout
func handleUsuario(w http.ResponseWriter, r *http.Request) {
    // 1. Crear contexto con timeout
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()

    // 2. Pasar datos si es necesario
    ctx = context.WithValue(ctx, requestIDKey{}, r.Header.Get("X-Request-ID"))

    // 3. Obtener datos (respeta contexto)
    usuario, err := obtenerUsuario(ctx, r.URL.Query().Get("id"))

    // 4. Manejar errores de timeout
    if err == context.DeadlineExceeded {
        w.WriteHeader(http.StatusRequestTimeout)
        fmt.Fprintf(w, "Solicitud expirada")
        return
    }

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error: %v", err)
        return
    }

    // 5. Responder
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"id":%d,"nombre":"%s"}`, usuario.ID, usuario.Nombre)
}

func obtenerUsuario(ctx context.Context, id string) (*Usuario, error) {
    // Respetar el contexto
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
    }

    // Simular queryla BD
    time.Sleep(1 * time.Second)

    // Verificar ANTES de retornar
    if err := ctx.Err(); err != nil {
        return nil, err
    }

    return &Usuario{ID: 1, Nombre: "Juan"}, nil
}

type Usuario struct {
    ID     int
    Nombre string
}

func main() {
    http.HandleFunc("/usuario", handleUsuario)
    http.ListenAndServe(":8080", nil)
}
```

---

## Parte VI: Anti-patrones

### ❌ No pasar contexto

```go
// ❌ Malo: Sin contexto
func consultarBD() ([]Usuario, error) {
    // Si la BD se cuelga, esto es indefinido
    conn := abrirConexion()
    return db.QueryContext(context.Background(), query)
}

// ✅ Bien: Con contexto
func consultarBD(ctx context.Context) ([]Usuario, error) {
    return db.QueryContext(ctx, query)
}
```

### ❌ Ignorar cancelación

```go
// ❌ Malo: Ignora cuando contexto se cancela
func procesarDatos(ctx context.Context) {
    for i := 0; i < 1000000; i++ {
        // No verifica si ctx se canceló
        procesarItem(i)
    }
}

// ✅ Bien: Respeta cancelación
func procesarDatos(ctx context.Context) {
    for i := 0; i < 1000000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            procesarItem(i)
        }
    }
}
```

### ❌ Crear contexto sin pasarlo

```go
// ❌ Malo: Cada función crea su propio timeout
func manejarSolicitud(w http.ResponseWriter, r *http.Request) {
    usuario, _ := obtenerUsuario()      // Timeout 5s
    orden, _ := obtenerOrden()          // Timeout 5s
    pago, _ := obtenerPago()            // Timeout 5s
    // Total: hasta 15s, pero cliente esperaba 10s
}

// ✅ Bien: Un deadline para toda la solicitud
func manejarSolicitud(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()

    usuario, _ := obtenerUsuario(ctx)   // Comparte deadline
    orden, _ := obtenerOrden(ctx)       // Comparte deadline
    pago, _ := obtenerPago(ctx)         // Comparte deadline
    // Total: máximo 10s
}
```

### ❌ Olvidar defer cancel()

```go
// ❌ Malo: Leak de goroutines
func manejarVarias() {
    for i := 0; i < 100; i++ {
        ctx, cancel := context.WithCancel(context.Background())
        operacion(ctx)
        // ¡Nunca llamas a cancel()!
    }
}

// ✅ Bien: Siempre defer cancel()
func manejarVarias() {
    for i := 0; i < 100; i++ {
        ctx, cancel := context.WithCancel(context.Background())
        operacion(ctx)
        cancel()  // O mejor: defer cancel() en una función
    }
}
```

---

## Parte VII: Timeout en diferentes operaciones

### BD con timeout

```go
func obtenerUsuario(ctx context.Context, id int) (*Usuario, error) {
    // ✅ Driver respeta el contexto
    row := db.QueryRowContext(ctx, "SELECT * FROM usuarios WHERE id = ?", id)
    
    var u Usuario
    if err := row.Scan(&u.ID, &u.Nombre); err != nil {
        return nil, err
    }
    return &u, nil
}
```

### HTTP client con timeout

```go
func consultarAPI(ctx context.Context, url string) ([]byte, error) {
    // ✅ El client respeta el contexto
    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}
```

### Goroutine con timeout

```go
func consultarGoroutine(ctx context.Context, resultado chan int) {
    for {
        select {
        case <-ctx.Done():
            return
        case <-time.After(1 * time.Second):
            // Trabajar
            resultado <- 42
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    resultado := make(chan int)
    go consultarGoroutine(ctx, resultado)

    for r := range resultado {
        fmt.Println(r)
    }
}
```

---

## Parte VIII: Mejores Prácticas

### 1. Pasar contexto como primer parámetro

```go
// ✅ Convención Go
func ConsultarBD(ctx context.Context, id int) (*Usuario, error) {
    // ...
}

// ❌ No hacerlo no seguir convención
func ConsultarBD(id int, ctx context.Context) (*Usuario, error) {
    // ...
}
```

### 2. Siempre respectar el contexto

```go
// ✅ Respeta deadlines
func processoLargo(ctx context.Context) error {
    for i := 0; i < 1000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Procesar
        }
    }
    return nil
}
```

### 3. Usar Values para request-scoped data

```go
// ✅ Bien: Pasar datos contextuales
type userKey struct{}
ctx := context.WithValue(context.Background(), userKey{}, user)

// ❌ Malo: Usar globals
var currentUser *User  // No escalable
```

### 4. Crear contexto con timeout en entrada

```go
// ✅ HTTP
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
    defer cancel()
    // ...
}

// ✅ gRPC
func (s *Server) GetUser(ctx context.Context, req *Request) (*Response, error) {
    // ctx ya tiene deadline del cliente
    // ...
}
```

### 5. Cleanup con context

```go
// ✅ Limpiar recursos automáticamente
func Connect(ctx context.Context) (*Connection, error) {
    conn := &Connection{}
    go func() {
        <-ctx.Done()
        conn.Close()
    }()
    return conn, nil
}
```

---

## Resumen: Context en Acción

| Caso | Patrón |
|------|--------|
| **Solicitud HTTP** | `context.WithTimeout(r.Context(), 30*s)` |
| **Operación BD** | `db.QueryContext(ctx, query)` |
| **HTTP Client** | `http.NewRequestWithContext(ctx, ...)` |
| **Cancelación manual** | `ctx, cancel := context.WithCancel()` ... `cancel()` |
| **Pasar datos** | `context.WithValue(ctx, key, value)` |
| **Cleanup** | `go func() { <-ctx.Done(); cleanup() }()` |

---

## Conclusión

**Context no es opcional en Go profesional**, es fundamental:

✅ **Evita operaciones colgadas**  
✅ **Respeta timeouts de cliente**  
✅ **Limpia recursos automáticamente**  
✅ **Propaga cancelación** a través de la pila  
✅ **Es estándar en la industria**

**Domina context y escribirás servidores robustos.**
