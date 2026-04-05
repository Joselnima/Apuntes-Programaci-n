# ORM y Query Builders en Go - Profesional

> **Elegir la herramienta correcta para acceder a bases de datos sin caídas ni problemas en producción**

---

## ¿Por qué existen los ORMs?

Imagina que escribes SQL crudo:

```go
// ❌ SQL crudo (frágil)
query := "SELECT * FROM usuarios WHERE id = " + strconv.Itoa(id)
rows, err := db.Query(query)
```

**Problemas:**
- 🔴 **SQL Injection**: Si `id` viene del usuario, ¡es un aguero de seguridad!
- 🔴 **Type-unsafe**: `rows.Scan()` requiere saber tipos exactos
- 🔴 **Code duplication**: Escribir parsing manual en muchos lugares
- 🔴 **Cambios en BD**: Si tabla cambia, todo falla

Un **ORM** o **Query Builder** resuelve esto:
- ✅ **Seguridad**: Parámetros preparados automáticamente
- ✅ **Type-safe**: Compilación verifica tipos
- ✅ **DRY**: Una definición de tabla, múltiples operaciones
- ✅ **Abstracción**: Mismo código para PostgreSQL, MySQL, etc.

---

## Parte I: SQL Puro vs ORM

### ❌ SQL Puro (Bajo nivel, propenso a errores)

```go
func obtenerUsuario(db *sql.DB, id int) (*Usuario, error) {
    // Vulnerable a SQL injection
    query := fmt.Sprintf("SELECT id, nombre, email FROM usuarios WHERE id = %d", id)
    
    row := db.QueryRow(query)
    
    var u Usuario
    if err := row.Scan(&u.ID, &u.Nombre, &u.Email); err != nil {
        return nil, err
    }
    
    return &u, nil
}
```

**Problemas:**
- SQL injection
- Manual row parsing
- No reusable
- Fácil olvidar campos al cambiar tabla

### ✅ Con Query Builder (GORM)

```go
func obtenerUsuario(db *gorm.DB, id int) (*Usuario, error) {
    var u Usuario
    if err := db.First(&u, id).Error; err != nil {
        return nil, err
    }
    return &u, nil
}
```

**Ventajas:**
- Sin SQL injection
- Automático parsing de tipos
- Reusable y mantenible
- Cambios en tabla se reflejan automáticamente

---

## Parte II: GORM (El ORM más popular en Go)

### ¿Por qué GORM?

GORM es el estándar de facto en Go profesional:
- ✅ **Maduro**: Millones de líneas en producción
- ✅ **Completo**: CRUD, hooks, validación, migraciones
- ✅ **Compatible**: PostgreSQL, MySQL, SQLite, SQL Server
- ✅ **Type-safe**: Compilación verifica tipos
- ✅ **Active**: Mantenimiento constante

### Instalación

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### Definir Modelo

```go
type Usuario struct {
    ID        uint           `gorm:"primaryKey"`
    Nombre    string         `gorm:"index"`
    Email     string         `gorm:"uniqueIndex"`
    Edad      int
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Nombre de tabla (por defecto: usuarios)
func (Usuario) TableName() string {
    return "usuarios"
}
```

### Conectar a BD

```go
import "gorm.io/driver/postgres"
import "gorm.io/gorm"

// PostgreSQL
dsn := "host=localhost port=5432 user=postgres password=pass dbname=miapp sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// MySQL
dsn := "user:pass@tcp(localhost:3306)/dbname?charset=utf8mb4"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// SQLite
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
```

### CRUD - Create

```go
// Crear un usuario
usuario := Usuario{Nombre: "Juan", Email: "juan@example.com", Edad: 30}

result := db.Create(&usuario)
if result.Error != nil {
    log.Fatal(result.Error)
}

fmt.Println("Usuario creado con ID:", usuario.ID)
```

### CRUD - Read

```go
// Obtener por ID
var usuario Usuario
db.First(&usuario, 1)  // ID = 1

// Obtener con WHERE
var usuarios []Usuario
db.Where("edad > ?", 25).Find(&usuarios)

// Obtener un registro
var usuario Usuario
db.Where("email = ?", "juan@example.com").First(&usuario)

// Obtener todos
var usuarios []Usuario
db.Find(&usuarios)

// Con limit y offset
db.Limit(10).Offset(0).Find(&usuarios)
```

### CRUD - Update

```go
// Actualizar un campo
db.Model(&Usuario{}).Where("id = ?", 1).Update("nombre", "Carlos")

// Actualizar múltiples campos
db.Model(&usuario).Updates(Usuario{Nombre: "Carlos", Edad: 31})

// O actualizar la estructura
usuario.Nombre = "Carlos"
db.Save(&usuario)
```

### CRUD - Delete

```go
// Borrar por ID
db.Delete(&Usuario{}, 1)

// Borrar con condición
db.Where("edad < ?", 18).Delete(&Usuario{})

// Borrar todo un registro
usuario := Usuario{ID: 1}
db.Delete(&usuario)
```

### Query Avanzado

```go
// Búsqueda con múltiples condiciones
var usuarios []Usuario
db.Where("edad > ? AND nombre LIKE ?", 25, "%Juan%").
    Order("edad DESC").
    Limit(10).
    Find(&usuarios)

// Relaciones (JOINS)
type Pedido struct {
    ID        uint
    UsuarioID uint
    Usuario   Usuario  // Relación
    Monto     float64
}

var pedidos []Pedido
db.Preload("Usuario").Find(&pedidos)  // Carga usuarios asociados

// Subqueries
var usuarios []Usuario
db.Where("id IN (?)", db.Select("usuario_id").From("pedidos")).Find(&usuarios)
```

### Migraciones (estructura de BD)

```go
type Usuario struct {
    ID    uint   `gorm:"primaryKey"`
    Nombre string
    Email string `gorm:"uniqueIndex"`
}

func main() {
    db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    // ✅ Auto-migración (crea tabla si no existe)
    db.AutoMigrate(&Usuario{})
    
    // O migración manual
    db.Migrator().CreateTable(&Usuario{})
    db.Migrator().AddIndex(&Usuario{}, "email")
}
```

### Transacciones

```go
// Transacción manual
tx := db.BeginTx(ctx, &sql.TxOptions{})

if err := tx.Create(&usuario1).Error; err != nil {
    tx.Rollback()
    return err
}

if err := tx.Create(&usuario2).Error; err != nil {
    tx.Rollback()
    return err
}

tx.Commit()  // Todo éxito

// O con callback
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&usuario1).Error; err != nil {
        return err
    }
    return tx.Create(&usuario2).Error
})
```

### ✅ Patrón profesional - Repository

```go
type UsuarioRepository struct {
    db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
    return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) Crear(ctx context.Context, usuario *Usuario) error {
    return r.db.WithContext(ctx).Create(usuario).Error
}

func (r *UsuarioRepository) ObtenerPorID(ctx context.Context, id uint) (*Usuario, error) {
    var usuario Usuario
    if err := r.db.WithContext(ctx).First(&usuario, id).Error; err != nil {
        return nil, err
    }
    return &usuario, nil
}

func (r *UsuarioRepository) ObtenerPorEmail(ctx context.Context, email string) (*Usuario, error) {
    var usuario Usuario
    if err := r.db.WithContext(ctx).Where("email = ?", email).First(&usuario).Error; err != nil {
        return nil, err
    }
    return &usuario, nil
}

func (r *UsuarioRepository) Actualizar(ctx context.Context, usuario *Usuario) error {
    return r.db.WithContext(ctx).Save(usuario).Error
}

func (r *UsuarioRepository) Borrar(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(&Usuario{}, id).Error
}

// Uso
func main() {
    db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    repo := NewUsuarioRepository(db)
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    usuario, err := repo.ObtenerPorEmail(ctx, "juan@example.com")
    if err != nil {
        log.Fatal(err)
    }
}
```

---

## Parte III: sqlc (Type-Safe, Recomendado)

### ¿Por qué sqlc?

sqlc es **no mágico**: tú escribes SQL, sqlc genera código Go type-safe.

**Ventajas sobre GORM:**
- ✅ **100% SQL puro**: Tienes control total
- ✅ **Type-safe**: Errores de tipo en compilación
- ✅ **Performance**: Sin reflexión como GORM
- ✅ **Mantenible**: SQL separado de Go
- ✅ **Auditable**: Lo que ves es lo que ejecuta

**Desventaja:**
- ❌ Más boilerplate inicial

### Instalar y setup

```bash
# Instalar sqlc
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# Crear config
# sqlc.yaml
version: "2"
sql:
  - engine: "postgres"
    queries: "queries/"
    schema: "schema/"
    out: "db"
```

### Definir schema

```sql
-- schema/schema.sql
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    edad INT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### Escribir queries

```sql
-- queries/usuarios.sql

-- name: CreateUsuario :one
INSERT INTO usuarios (nombre, email, edad)
VALUES ($1, $2, $3)
RETURNING id, nombre, email, edad, created_at;

-- name: GetUsuario :one
SELECT id, nombre, email, edad, created_at
FROM usuarios
WHERE id = $1;

-- name: ListUsuarios :many
SELECT id, nombre, email, edad, created_at
FROM usuarios
ORDER BY id;

-- name: UpdateUsuario :exec
UPDATE usuarios
SET nombre = $1, email = $2, edad = $3
WHERE id = $4;

-- name: DeleteUsuario :exec
DELETE FROM usuarios
WHERE id = $1;
```

### Generar código

```bash
sqlc generate
```

Genera `db/queries.sql.go` con funciones type-safe.

### Usar código generado

```go
import (
    "context"
    "github.com/jackc/pgx/v5"
    "myapp/db"
)

func main() {
    conn, _ := pgx.Connect(context.Background(), dsn)
    queries := db.New(conn)
    
    // Create
    usuario, _ := queries.CreateUsuario(context.Background(), db.CreateUsuarioParams{
        Nombre: "Juan",
        Email:  "juan@example.com",
        Edad:   30,
    })
    
    // Read
    u, _ := queries.GetUsuario(context.Background(), usuario.ID)
    fmt.Println(u.Nombre)
    
    // List
    usuarios, _ := queries.ListUsuarios(context.Background())
    for _, u := range usuarios {
        fmt.Println(u.Nombre)
    }
    
    // Update
    queries.UpdateUsuario(context.Background(), db.UpdateUsuarioParams{
        ID:     usuario.ID,
        Nombre: "Carlos",
    })
    
    // Delete
    queries.DeleteUsuario(context.Background(), usuario.ID)
}
```

### Patrón profesional con sqlc

```go
type UsuarioStore struct {
    queries *db.Queries
}

func NewUsuarioStore(conn *pgx.Conn) *UsuarioStore {
    return &UsuarioStore{queries: db.New(conn)}
}

func (s *UsuarioStore) Crear(ctx context.Context, usuario *Usuario) error {
    _, err := s.queries.CreateUsuario(ctx, db.CreateUsuarioParams{
        Nombre: usuario.Nombre,
        Email:  usuario.Email,
        Edad:   int32(usuario.Edad),
    })
    return err
}

func (s *UsuarioStore) ObtenerPorID(ctx context.Context, id int32) (*Usuario, error) {
    u, err := s.queries.GetUsuario(ctx, id)
    if err != nil {
        return nil, err
    }
    return &Usuario{
        ID:     int(u.ID),
        Nombre: u.Nombre,
        Email:  u.Email,
        Edad:   int(u.Edad.Int32),
    }, nil
}

// Uso
func main() {
    conn, _ := pgx.Connect(context.Background(), dsn)
    store := NewUsuarioStore(conn)
    
    usuario, _ := store.ObtenerPorID(context.Background(), 1)
    fmt.Println(usuario.Nombre)
}
```

---

## Parte IV: ent (Entity Framework para Go)

### ¿Qué es ent?

ent es un entity framework moderno, similar a GORM pero con enfoque diferente:

**Ventajas:**
- ✅ **Code-first**: Define schema en Go
- ✅ **Graph queries**: Navegar relaciones fácilmente
- ✅ **Type-safe**: Compilación strict
- ✅ **Hooks y middleware**

### Instalación

```bash
go get entgo.io/ent/cmd/ent
ent new User
```

### Definir schema

```go
// ent/schema/user.go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type User struct {
    ent.Schema
}

func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id"),
        field.String("nombre").NotEmpty(),
        field.String("email").Unique(),
        field.Int("edad"),
    }
}
```

### CRUD con ent

```go
import "myapp/ent"

// Create
usuario, _ := client.User.Create().
    SetNombre("Juan").
    SetEmail("juan@example.com").
    SetEdad(30).
    Save(ctx)

// Read
usuario, _ := client.User.Get(ctx, 1)

// Update
usuario, _ := usuario.Update().
    SetNombre("Carlos").
    Save(ctx)

// Delete
client.User.DeleteOneID(1).Exec(ctx)

// Query lista
usuarios, _ := client.User.Query().
    Where(user.AgedGT(25)).
    All(ctx)
```

---

## Parte V: Cuándo Usar Cada Uno

### GORM - Mejor para:
- ✅ **Prototipado rápido**
- ✅ **Cambios frecuentes de schema**
- ✅ **Equipo pequeño**
- ✅ **SQL complejo con ORM**
- 📍 Proyectos medianos y MVPs

### sqlc - Mejor para:
- ✅ **SQL optimizado crítico**
- ✅ **Control total de queries**
- ✅ **Performance importa**
- ✅ **Equipos grandes**
- 📍 Producción, startups serias

### ent - Mejor para:
- ✅ **Grandes aplicaciones**
- ✅ **Mucha navegación de relaciones**
- ✅ **Schema code-first**
- ✅ **Equipos estructurados**
- 📍 Arquitectura limpia

### SQL Puro - SOLO si:
- ⚠️ **Casos EXCEPCIONALES**
- ⚠️ **SQL muy especializado**
- ⚠️ **Puedes garantizar que es seguro**

---

## Parte VI: Anti-patrones

### ❌ SQL Injection

```go
// ❌ Malo
query := "SELECT * FROM usuarios WHERE id = " + id

// ✅ Bien: GORM
db.Where("id = ?", id).First(&usuario)

// ✅ Bien: sqlc (parameterizado)
// SQL: WHERE id = $1
```

### ❌ Ignorar errores de BD

```go
// ❌ Malo
db.Where("id = ?", id).First(&usuario)
utilizar(usuario)  // ¿Qué si no existe?

// ✅ Bien
if err := db.Where("id = ?", id).First(&usuario).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return errors.New("usuario no encontrado")
    }
    return err
}
```

### ❌ Blasphemy, usar global DB

```go
// ❌ Malo: Global
var DB *gorm.DB

// ✅ Bien: Inyectar
func NewUserHandler(db *gorm.DB) *UserHandler {
    return &UserHandler{db: db}
}
```

### ❌ Olvidar context

```go
// ❌ Malo
db.First(&usuario).Error

// ✅ Bien (GORM)
db.WithContext(ctx).First(&usuario).Error

// ✅ Bien (sqlc)
queries.GetUsuario(ctx, id)
```

---

## Parte VII: Mejores Prácticas

### 1. Siempre usar context

```go
// ✅ Respeta timeouts
func (r *UsuarioRepository) Crear(ctx context.Context, u *Usuario) error {
    return r.db.WithContext(ctx).Create(u).Error
}
```

### 2. Usar transactions para operaciones críticas

```go
// ✅ ACID guaranteado
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&usuario1).Error; err != nil {
        return err
    }
    return tx.Create(&usuario2).Error
})
```

### 3. Repository pattern

```go
// ✅ Abstracción limpia
type UserRepository interface {
    Crear(ctx context.Context, u *Usuario) error
    ObtenerPorID(ctx context.Context, id int) (*Usuario, error)
}
```

### 4. Migraciones versionadas (sqlc)

```sql
-- migrations/001_create_usuarios.sql
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL
);
```

### 5. Índices para queries frecuentes

```go
// GORM
type Usuario struct {
    ID    uint   `gorm:"primaryKey"`
    Email string `gorm:"index"`  // ← Índice para búsqueda rápida
}
```

---

## Resumen: Jerarquía de Recomendación

```
Proyecto        │  Recomendado
────────────────┼─────────────
MVP/Prototipo   │  GORM
Producción      │  sqlc (mejor performance)
Grande/Complejo │  ent
Especial        │  SQL puro (con cuidado)
```

---

## Conclusión

**En Go profesional, elige entre GORM, sqlc o ent**:

- 🏆 **sqlc para performance crítico**: Type-safe, SQL puro, rápido
- 🏆 **Gorm para desarrollo rápido**: Completo, fácil, MÁS adoptado
- 🏆 **ent para grandes aplicaciones**: Arquitectura limpia, escalable

**Nunca SQL puro** a menos que realmente no haya alternativa.

**Go tiene ecosistema maduro de ORMs.Úsalos.**
