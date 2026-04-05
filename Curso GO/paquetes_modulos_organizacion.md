# Paquetes, Módulos y Organización de Proyecto en Go

> **Cómo estructurar proyectos grandes, gestionar dependencias y reutilizar código**

---

## ¿Por qué organización importa en Go?

Un programa pequeño cabe en un archivo. Pero cuando crece:

```
❌ Sin organización:
main.go  (5000 líneas)   ← Imposible de mantener

✅ Con organización:
main.go           (500 líneas)
usuarios/
  usuario.go      (300 líneas)
  service.go      (400 líneas)
productos/
  producto.go     (200 líneas)
  service.go      (350 líneas)
storage/
  postgres.go     (600 líneas)
  storage.go      (150 líneas)
```

Go **obliga explícitamente** a pensar en estructura. Sin paquetes, todo colisiona. Con ellos, el código es escalable.

---

## Parte I: Paquetes (Packages)

### ¿Qué es un paquete?

Un **paquete** es un directorio que contiene archivos `.go` que **comparten el mismo nombre de paquete**:

```
usuarios/
├── usuario.go        (package usuarios)
├── service.go        (package usuarios)
└── validator.go      (package usuarios)
```

**Todos son `package usuarios`** aunque estén en archivos diferentes.

### Creación de paquete

```
proyecto/
├── main.go
├── usuarios/
│   ├── usuario.go
│   └── service.go
```

**usuarios/usuario.go:**
```go
package usuarios    // ← Nombre del paquete

type Usuario struct {
    ID    int
    Nombre string
}

func NewUsuario(nombre string) *Usuario {
    return &Usuario{Nombre: nombre}
}
```

**usuarios/service.go:**
```go
package usuarios    // ← MISMO paquete que usuario.go

func ObtenerUsuario(id int) (*Usuario, error) {
    // Lógica
}
```

**main.go:**
```go
package main

import (
    "fmt"
    "proyecto/usuarios"  // Importar paquete
)

func main() {
    u := usuarios.NewUsuario("Juan")
    fmt.Println(u.Nombre)
}
```

### Visibilidad (exportado vs privado)

En Go, **la visibilidad depende de mayúsculas:**

```go
// ✅ Exportado (primera letra mayúscula)
type Usuario struct { }
func NewUsuario(nombre string) *Usuario { }
func ObtenerUsuarios() []*Usuario { }

// ❌ Privado (primera letra minúscula)
type usuarioInterno struct { }  // Solo dentro del paquete
func crearUsuario() { }         // Solo dentro del paquete
```

**En otro paquete:**
```go
import "proyecto/usuarios"

usuarios.NewUsuario("Juan")      // ✅ Funciona (exportado)
usuarios.crearUsuario()          // ❌ Error: no se puede acceder (privado)
usuarios.Usuario{}               // ✅ Funciona
usuarios.usuarioInterno{}        // ❌ Error: privado
```

### init() en paquetes

Cada paquete puede tener una función `init()` que se ejecuta automáticamente:

```go
// usuarios/usuario.go
package usuarios

var logger Logger

func init() {
    // Se ejecuta cuando se importa el paquete
    logger = crearLogger()
}

// Orden de ejecución:
// 1. Se ejecutan init() de todos los paquetes importados
// 2. Se ejecuta init() del paquete main
// 3. Se ejecuta main()
```

---

## Parte II: Módulos (go.mod)

### ¿Qué es el módulo?

Un **módulo** es un proyecto completo con todas sus dependencias. El archivo `go.mod` es el "manifiesto":

```
proyecto/
├── go.mod           ← Declara módulo y versiones
├── go.sum           ← Log de checksums (no edites)
├── main.go
├── usuarios/
│   └── usuario.go
└── productos/
    └── producto.go
```

### Crear un módulo

```bash
cd proyecto
go mod init miproyecto

# Crea go.mod:
# module miproyecto
# go 1.21
```

El nombre (`miproyecto`) debe ser:
- En GitHub: `github.com/usuario/miproyecto`
- Local: cualquier nombre que tenga sentido

### go.mod - versiones de dependencias

```
module miproyecto

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.9
)
```

- `require` = dependencias directas
- `go 1.21` = versión mínima de Go

### Agregar dependencia

```bash
# Descargar dependencia
go get github.com/gorilla/mux

# Actualizar todos
go get -u ./...

# Instalar versión específica
go get github.com/gorilla/mux@v1.8.0
```

### Limpiar dependencias sin usar

```bash
go mod tidy
```

---

## Parte III: Estructura Típica de Proyecto

### Pequeño proyecto

```
miaplicacion/
├── go.mod
├── go.sum
├── main.go              # Punto de entrada
├── config/
│   └── config.go        # Configuración
├── usuarios/
│   ├── usuario.go       # Tipos
│   └── service.go       # Lógica
└── storage/
    └── postgres.go      # Acceso a BD
```

### Proyecto mediano

```
empresa/
├── go.mod
├── main.go
├── cmd/              ← Ejecutables
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── pkg/              ← Código reutilizable
│   ├── usuarios/
│   │   ├── usuario.go
│   │   └── service.go
│   ├── productos/
│   │   ├── producto.go
│   │   └── service.go
│   └── storage/
│       ├── storage.go   # Interfaz
│       ├── postgres.go
│       └── mysql.go
├── internal/         ← Código privado del proyecto
│   ├── config/
│   └── middleware/
└── tests/
    └── integration_test.go
```

### Estructura estándar (Go way)

```
github.com/usuario/proyecto/
├── cmd/
│   └── myapp/
│       └── main.go              # El ejecutable principal
├── pkg/
│   ├── usuarios/
│   ├── productos/
│   └── storage/
├── internal/
│   ├── config/
│   └── middleware/
├── tests/
│   └── integration_test.go
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```

**Reglas:**
- `cmd/` = Ejecutables (main)
- `pkg/` = Código público (reutilizable por otros)
- `internal/` = Código privado (solo para este proyecto)
- `tests/` = Tests de integración

---

## Parte IV: Importar Paquetes Locales

### Estructura

```
proyecto/
├── go.mod          (module miproyecto)
├── main.go
└── usuarios/
    └── usuario.go
```

### Importar en main.go

```go
package main

import (
    "fmt"
    "miproyecto/usuarios"  // ← Nombre del módulo + ruta
)

func main() {
    u := usuarios.NewUsuario("Juan")
    fmt.Println(u.Nombre)
}
```

**Regla:** Importar siempre con el **nombre del módulo** (en `go.mod`), no la ruta del filesystem.

### Ciclos de importación (errores comunes)

```go
// ❌ CIRCULAR - Error de compilación
// main.go
import "proyecto/usuarios"

// usuarios/usuario.go
import "proyecto"  // ← ERROR: import circular
```

**Solución:** Crear paquete separado para tipos compartidos:

```
proyecto/
├── main.go
├── models/
│   └── usuario.go    (tipos, sin dependencias de otros)
├── usuarios/
│   └── service.go    (usa models)
└── productos/
    └── service.go    (usa models)
```

---

## Parte V: Paquetes de terceros (Third-party)

### Agregar dependencia

```bash
go get github.com/gorilla/mux
```

Auto-actualiza `go.mod`:
```
require (
    github.com/gorilla/mux v1.8.0
)
```

### Usar

```go
import "github.com/gorilla/mux"

func main() {
    r := mux.NewRouter()
}
```

### go.sum - no editar manualmente

`go.sum` contiene checksums de seguridad. Go lo maneja automáticamente:

```
github.com/gorilla/mux v1.8.0 h1:...
github.com/gorilla/mux v1.8.0/go.mod h1:...
```

---

## Parte VI: Ejemplo Práctico - API Estructurada

### Estructura del proyecto

```
empresa/
├── go.mod
├── main.go
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   └── middleware/
│       └── logger.go
└── pkg/
    ├── usuarios/
    │   ├── usuario.go
    │   └── service.go
    └── storage/
        ├── storage.go
        └── postgres.go
```

### go.mod

```
module empresa

go 1.21

require (
    github.com/lib/pq v1.10.9
)
```

### users/usuario.go

```go
package usuarios

type Usuario struct {
    ID    int    `json:"id"`
    Nombre string `json:"nombre"`
    Email  string `json:"email"`
}

func NewUsuario(nombre, email string) *Usuario {
    return &Usuario{
        Nombre: nombre,
        Email:  email,
    }
}

func (u *Usuario) EsValido() bool {
    return len(u.Nombre) > 0 && len(u.Email) > 0
}
```

### storage/storage.go (Interfaz)

```go
package storage

import "empresa/pkg/usuarios"

type Store interface {
    CrearUsuario(u *usuarios.Usuario) error
    ObtenerUsuario(id int) (*usuarios.Usuario, error)
    ListarUsuarios() ([]*usuarios.Usuario, error)
}
```

### storage/postgres.go (Implementación)

```go
package storage

import (
    "database/sql"
    "empresa/pkg/usuarios"

    _ "github.com/lib/pq"
)

type PostgresStore struct {
    db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return &PostgresStore{db: db}, nil
}

func (ps *PostgresStore) CrearUsuario(u *usuarios.Usuario) error {
    return ps.db.QueryRow(
        "INSERT INTO usuarios (nombre, email) VALUES ($1, $2) RETURNING id",
        u.Nombre, u.Email).Scan(&u.ID)
}

func (ps *PostgresStore) ObtenerUsuario(id int) (*usuarios.Usuario, error) {
    u := &usuarios.Usuario{}
    err := ps.db.QueryRow(
        "SELECT id, nombre, email FROM usuarios WHERE id = $1", id).
        Scan(&u.ID, &u.Nombre, &u.Email)
    if err != nil {
        return nil, err
    }
    return u, nil
}

func (ps *PostgresStore) ListarUsuarios() ([]*usuarios.Usuario, error) {
    rows, err := ps.db.Query("SELECT id, nombre, email FROM usuarios")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var usuarios []*usuarios.Usuario
    for rows.Next() {
        u := &usuarios.Usuario{}
        if err := rows.Scan(&u.ID, &u.Nombre, &u.Email); err != nil {
            return nil, err
        }
        usuarios = append(usuarios, u)
    }
    return usuarios, nil
}
```

### internal/config/config.go

```go
package config

type Config struct {
    DBHost string
    DBPort string
    DBUser string
    DBPass string
    Port   string
}

func LoadConfig() Config {
    return Config{
        DBHost: "localhost",
        DBPort: "5432",
        DBUser: "postgres",
        DBPass: "password",
        Port:   "8080",
    }
}
```

### main.go

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "empresa/internal/config"
    "empresa/pkg/usuarios"
    "empresa/pkg/storage"
)

func main() {
    // Cargar configuración
    cfg := config.LoadConfig()

    // Conectar a BD
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass)
    
    store, err := storage.NewPostgresStore(connStr)
    if err != nil {
        log.Fatal(err)
    }

    // Crear usuario de ejemplo
    u := usuarios.NewUsuario("Juan", "juan@example.com")
    if err := store.CrearUsuario(u); err != nil {
        log.Fatal(err)
    }

    // API
    http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
        usuarios, _ := store.ListarUsuarios()
        fmt.Fprintf(w, "Total: %d usuarios", len(usuarios))
    })

    fmt.Printf("Servidor en localhost:%s\n", cfg.Port)
    http.ListenAndServe(":"+cfg.Port, nil)
}
```

---

## Parte VII: Best Practices

### 1. Usa `internal/` para código privado

```go
// ✅ Bien: private al proyecto
internal/middleware/logger.go

// ❌ Evitar: todo en pkg
pkg/private/middleware/logger.go
```

### 2. Paquetes pequeños y cohesivos

```go
// ❌ Paquete gigante
usuarios/
└── usuario.go (2000 líneas de TODO)

// ✅ Separar responsabilidades
usuarios/
├── usuario.go      (tipos)
├── service.go      (lógica)
├── validator.go    (validación)
└── repository.go   (acceso a datos)
```

### 3. Una responsabilidad por paquete

```go
// ❌ Mezclar responsabilidades
paquete/
└── paquete.go  (usuarios, productos, órdenes)

// ✅ Separar
├── usuarios/usuario.go
├── productos/producto.go
└── ordenes/orden.go
```

### 4. Nombres explícitos de paquete

```go
// ❌ Confuso
utils/
utils/utils.go

// ✅ Claro
strings/
math/
storage/
```

### 5. Agrupa tipos relacionados

```go
// ✅ Bien: tipos en un paquete
usuarios/
├── usuario.go       (type Usuario)
├── rol.go          (type Rol)
└── permisos.go     (type Permiso)

// ❌ Peor: tipos en archivos separados
usuarios/usuario.go
roles/rol.go
permisos/permiso.go
```

---

## Parte VIII: Anti-patrones

### ❌ Circular imports

```go
// usuarios/usuario.go
package usuarios
import "empresa/productos"   // ← usuarios depende de productos

// productos/producto.go
package productos
import "empresa/usuarios"    // ← ERROR: circular
```

### ❌ Paquete `util` gigante

```go
utils/
├── utils.go       (1000 líneas de TODO)
├── strings.go
├── math.go
├── slice.go
```

**Solución:** distribuir en paquetes específicos:
```go
strings/
math/
slice/
```

### ❌ Guardar estado global mutable

```go
// ❌ Malo
var db *sql.DB    // Compartido globalmente

// ✅ Bien
func NewUserService(db *sql.DB) *UserService {
    return &UserService{db: db}
}
```

### ❌ Init() con lógica pesada

```go
// ❌ init() se ejecuta automáticamente, difícil de debuggear
func init() {
    conectarABaseDatos()
    descargarDatos()
    procesarMil archivos()
}

// ✅ Lógica en main()
func main() {
    db, err := conectar()
    // ... manejo de errores
}
```

---

## Resumen: Estructura de Go

| Concepto | Uso |
|----------|-----|
| **Paquete** | Agrupa código relacionado en un directorio |
| **Módulo** | Proyecto completo con `go.mod` |
| **Visibilidad** | MAYÚSCULA = exportado, minúscula = privado |
| **`cmd/`** | Ejecutables (main.go) |
| **`pkg/`** | Código público y reutilizable |
| **`internal/`** | Código privado del proyecto |
| **`go.mod`** | Declara módulo y dependencias |

**Un proyecto bien organizado es un proyecto mantenible.**
