# HTTP y Servidores Web en Go

> **Cómo construir servidores web escalables y APIs REST con Go**

---

## ¿Por qué Go es perfecto para servidores web?

Desde el principio, Go fue diseñado pensando en **internet**. La razón:

1. **Concurrencia nativa**: Goroutines hacen fácil manejar 10,000 conexiones simultáneas
2. **Bajo overhead**: Un servidor web en Go consume muy pocos recursos
3. **Estándar "net/http"**: Incluido en la stdlib, robusto y probado
4. **Compilación estática**: Un binario, sin dependencias externas

**Dato curioso**: Google, Kubernetes, Docker, y Hertz se escribieron en Go precisamente por esto.

---

## Parte I: Servidor HTTP Mínimo

### El servidor más simple posible

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Manejar ruta raíz
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "¡Hola, mundo!")
    })

    // Escuchar en puerto 8080
    fmt.Println("Servidor en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Ejecución:**
```bash
go run main.go
# Servidor en http://localhost:8080
```

Abre navegador y ve http://localhost:8080 → verás "¡Hola, mundo!"

**¿Qué pasó?**
1. `HandleFunc` registra una función que maneja peticiones a esa ruta
2. La función recibe `ResponseWriter` (para responder) y `Request` (los datos del cliente)
3. `ListenAndServe` inicia el servidor en puerto 8080

---

## Parte II: Objeto Request (Recibir Datos)

El `*http.Request` contiene toda la información del cliente:

### Métodos HTTP

```go
func miHandler(w http.ResponseWriter, r *http.Request) {
    // Verificar método
    if r.Method == "GET" {
        fmt.Fprintf(w, "Recibí GET")
    } else if r.Method == "POST" {
        fmt.Fprintf(w, "Recibí POST")
    }
}
```

### Parámetros URL (query parameters)

```go
// URL: http://localhost:8080/buscar?nombre=Juan&edad=30

func buscar(w http.ResponseWriter, r *http.Request) {
    // Obtener parámetros
    nombre := r.URL.Query().Get("nombre")
    edad := r.URL.Query().Get("edad")
    
    fmt.Fprintf(w, "Nombre: %s, Edad: %s", nombre, edad)
}

func main() {
    http.HandleFunc("/buscar", buscar)
    http.ListenAndServe(":8080", nil)
}
```

**Output:** "Nombre: Juan, Edad: 30"

### Headers

```go
func miHandler(w http.ResponseWriter, r *http.Request) {
    // Leer header
    userAgent := r.Header.Get("User-Agent")
    fmt.Fprintf(w, "Tu navegador: %s", userAgent)
}
```

### Body (para POST, PUT)

```go
import (
    "encoding/json"
    "io"
)

type Usuario struct {
    Nombre string `json:"nombre"`
    Email  string `json:"email"`
}

func crearUsuario(w http.ResponseWriter, r *http.Request) {
    // Leer body
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()

    // Parsear JSON
    var usuario Usuario
    json.Unmarshal(body, &usuario)

    fmt.Fprintf(w, "Creado: %s", usuario.Nombre)
}

func main() {
    http.HandleFunc("/usuarios", crearUsuario)
    http.ListenAndServe(":8080", nil)
}
```

**Test con curl:**
```bash
curl -X POST http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan","email":"juan@example.com"}'
# Output: Creado: Juan
```

---

## Parte III: Objeto ResponseWriter (Enviar Respuesta)

### Función básica

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Escribir respuesta
    fmt.Fprintf(w, "Texto simple")
}
```

### Enviar JSON

```go
func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
    // Preparar datos
    usuario := map[string]string{
        "nombre": "Juan",
        "email":  "juan@example.com",
    }

    // Establecer header
    w.Header().Set("Content-Type", "application/json")

    // Escribir JSON
    json.NewEncoder(w).Encode(usuario)
}
```

### Estados HTTP (status codes)

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 200 OK (defecto)
    fmt.Fprintf(w, "Ok")
    
    // 404 Not Found
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "No encontrado")
    
    // 500 Error interno
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error en servidor")
}
```

```go
// Uso práctico
func obtenerProducto(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    if id == "" {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ID requerido")
        return
    }
    
    // Buscar producto...
    if !encontrado {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Producto no encontrado")
        return
    }
    
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Producto encontrado")
}
```

### Headers de respuesta

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Agregar headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom", "Mi valor")
    w.Header().Set("Access-Control-Allow-Origin", "*")  // CORS
    
    fmt.Fprintf(w, `{"status":"ok"}`)
}
```

---

## Parte IV: Enrutamiento (Routing)

### Rutas básicas

```go
func main() {
    http.HandleFunc("/", inicio)
    http.HandleFunc("/usuarios", usuarios)
    http.HandleFunc("/productos", productos)
    
    http.ListenAndServe(":8080", nil)
}

func inicio(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Página de inicio")
}

func usuarios(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Lista de usuarios")
}

func productos(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Lista de productos")
}
```

### Rutas con parámetros (opción manual)

Go estándar NO soporta `/usuarios/{id}` directamente. Tienes que extraer manualmente:

```go
func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
    // URL: /usuario/5
    id := r.URL.Path[len("/usuario/"):]  // Extrae "5"
    
    fmt.Fprintf(w, "Usuario con ID: %s", id)
}

func main() {
    http.HandleFunc("/usuario/", obtenerUsuario)
    http.ListenAndServe(":8080", nil)
}
```

### ✅ Mejor: Usar router (middleware)

Para proyectos reales, usa un router como `gorilla/mux`:

```bash
go get github.com/gorilla/mux
```

```go
import "github.com/gorilla/mux"

func main() {
    r := mux.NewRouter()
    
    // Rutas con variables
    r.HandleFunc("/usuario/{id}", obtenerUsuario).Methods("GET")
    r.HandleFunc("/usuarios", crearUsuario).Methods("POST")
    
    http.ListenAndServe(":8080", r)
}

func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "Usuario: %s", id)
}
```

---

## Parte V: Ejemplo Práctico - API REST de Usuarios

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "sync"
)

// Estructura de Usuario
type Usuario struct {
    ID    int    `json:"id"`
    Nombre string `json:"nombre"`
    Email  string `json:"email"`
}

// Almacenamiento en memoria (para ejemplo)
var (
    usuarios = make(map[int]*Usuario)
    mu       sync.RWMutex
    nextID   = 1
)

// GET /usuarios - Listar todos
func listarUsuarios(w http.ResponseWriter, r *http.Request) {
    mu.RLock()
    defer mu.RUnlock()

    w.Header().Set("Content-Type", "application/json")
    
    lista := make([]*Usuario, 0, len(usuarios))
    for _, u := range usuarios {
        lista = append(lista, u)
    }
    
    json.NewEncoder(w).Encode(lista)
}

// GET /usuario/{id} - Obtener uno
func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    idInt, err := strconv.Atoi(id)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ID inválido")
        return
    }

    mu.RLock()
    usuario, existe := usuarios[idInt]
    mu.RUnlock()

    if !existe {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Usuario no encontrado")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usuario)
}

// POST /usuario - Crear
func crearUsuario(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var u Usuario
    json.NewDecoder(r.Body).Decode(&u)

    mu.Lock()
    u.ID = nextID
    usuarios[nextID] = &u
    nextID++
    mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(u)
}

// PUT /usuario/{id} - Actualizar
func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    idInt, err := strconv.Atoi(id)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    var u Usuario
    json.NewDecoder(r.Body).Decode(&u)

    mu.Lock()
    usuario, existe := usuarios[idInt]
    if existe {
        usuario.Nombre = u.Nombre
        usuario.Email = u.Email
    }
    mu.Unlock()

    if !existe {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usuario)
}

// DELETE /usuario/{id} - Borrar
func borrarUsuario(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    idInt, err := strconv.Atoi(id)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    mu.Lock()
    _, existe := usuarios[idInt]
    if existe {
        delete(usuarios, idInt)
    }
    mu.Unlock()

    if !existe {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func main() {
    // Rutas
    http.HandleFunc("/usuarios", listarUsuarios)
    http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            obtenerUsuario(w, r)
        case "POST":
            crearUsuario(w, r)
        case "PUT":
            actualizarUsuario(w, r)
        case "DELETE":
            borrarUsuario(w, r)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    })

    // Health check
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "OK")
    })

    fmt.Println("Servidor en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Tests con curl:**

```bash
# Crear usuario
curl -X POST http://localhost:8080/usuario \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan","email":"juan@example.com"}'
# Output: {"id":1,"nombre":"Juan","email":"juan@example.com"}

# Listar todos
curl http://localhost:8080/usuarios
# Output: [{"id":1,"nombre":"Juan","email":"juan@example.com"}]

# Obtener uno
curl http://localhost:8080/usuario?id=1
# Output: {"id":1,"nombre":"Juan","email":"juan@example.com"}

# Actualizar
curl -X PUT http://localhost:8080/usuario?id=1 \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Jupiter","email":"jupiter@example.com"}'

# Borrar
curl -X DELETE http://localhost:8080/usuario?id=1
```

---

## Parte VI: Middleware

El **middleware** es código que se ejecuta antes/después de tus handlers.

### Middleware simple

```go
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("[%s] %s %s\n", time.Now().String(), r.Method, r.URL)
        next(w, r)
    }
}

func main() {
    http.HandleFunc("/usuarios", loggerMiddleware(listarUsuarios))
    http.ListenAndServe(":8080", nil)
}
```

**Output:**
```
[2025-01-10 10:30:45.123456789 +0000 UTC] GET /usuarios
```

### Middleware para CORS

```go
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next(w, r)
    }
}
```

---

## Parte VII: Manejo de Errores

### Función de utilidad para respuestas de error

```go
type ErrorResponse struct {
    Message string `json:"message"`
    Code    int    `json:"code"`
}

func errorJSON(w http.ResponseWriter, message string, code int) {
    w.WriteHeader(code)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ErrorResponse{
        Message: message,
        Code:    code,
    })
}

// Uso
func obtenerProducto(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    if id == "" {
        errorJSON(w, "ID requerido", http.StatusBadRequest)
        return
    }

    // Lógica...
}
```

---

## Parte VIII: Anti-patrones

### ❌ No manejar errores

```go
// ❌ Ignorar errores de parsing
var u Usuario
json.NewDecoder(r.Body).Decode(&u)  // ¿Qué si falla?

// ✅ Manejar
if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "JSON inválido")
    return
}
```

### ❌ Bloquear lectura sin deadline

```go
// ❌ Cliente puede conectar y no enviar nada, bloquea forever
http.ListenAndServe(":8080", nil)

// ✅ Con timeout
server := &http.Server{
    Addr:         ":8080",
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 5 * time.Second,
}
server.ListenAndServe()
```

### ❌ Datos de usuario sin validación

```go
// ❌ Confiar ciegamente en input
usuario := Usuario{
    Nombre: r.URL.Query().Get("nombre"),
    Email:  r.URL.Query().Get("email"),
}

// ✅ Validar
if len(usuario.Nombre) == 0 {
    errorJSON(w, "Nombre requerido", http.StatusBadRequest)
    return
}
if !strings.Contains(usuario.Email, "@") {
    errorJSON(w, "Email inválido", http.StatusBadRequest)
    return
}
```

---

## Resumen: HTTP en Go

| Concepto | Uso |
|----------|-----|
| `http.HandleFunc()` | Registrar handler para una ruta |
| `*http.Request` | Recibir datos (método, params, body) |
| `http.ResponseWriter` | Enviar respuesta (status, headers, body) |
| `json.NewEncoder()` | Enviar JSON |
| `json.NewDecoder()` | Recibir JSON |
| Middleware | Procesar antes/después de handlers |
| Status codes | Comunicar resultado (200, 404, 500) |

**Go es excelente para servidores porque es simple, robusto y escalable.**
