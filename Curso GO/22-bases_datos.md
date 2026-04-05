# Bases de Datos en Go - Guía Completa

> **Cómo conectar, consultar y gestionar datos con PostgreSQL, MySQL, Oracle, SQL Server y bases de datos NoSQL**

---

## ¿Por Qué Necesitas Bases de Datos?

Imagina que tienes una tienda en línea. ¿Dónde guardas:
- Los datos de clientes (nombre, email, teléfono)
- Los productos (nombre, precio, stock)
- Los pedidos (qué compró, cuándo, cuánto pagó)

**No puedes guardar esto en memoria** (desaparece cuando cierras el programa). Necesitas **persistencia**: guardar datos de forma permanente.

Las **bases de datos** son software especializado para:
1. **Guardar datos** de forma segura
2. **Recuperar datos** rápidamente
3. **Relacionar datos** entre ellos
4. **Proteger datos** con contraseñas y permisos

---

## Landscape de Bases de Datos

### Tipos Principales

```
Bases de Datos
├─ Relacionales (SQL)
│  ├─ PostgreSQL (favorita de Go)
│  ├─ MySQL / MariaDB
│  ├─ Oracle (empresarial)
│  ├─ SQL Server (Microsoft)
│  └─ SQLite (pequeña, local)
│
└─ No Relacionales (NoSQL)
   ├─ Documentos: MongoDB
   ├─ Clave-Valor: Redis
   ├─ Columnar: Cassandra
   ├─ Búsqueda: Elasticsearch
   └─ Grafos: Neo4j
```

---

## PARTE I: CONCEPTOS FUNDAMENTALES

### SQL vs NoSQL

| Aspecto | SQL | NoSQL |
|---------|-----|-------|
| **Estructura** | Tablas con esquema fijo | Documentos/clave-valor flexibles |
| **Relaciones** | Claves foráneas | Documentos anidados |
| **Transacciones** | ACID garantizado | Eventual consistency |
| **Escalabilidad** | Vertical (mejor máquina) | Horizontal (más servidores) |
| **Complejidad** | Consultas complejas con JOINs | Menos JOINs, más redundancia |
| **Uso típico** | Datos estructurados, finanzas | Datos flexible, redes sociales |

### Conexión con Go

```
Tu Programa Go
     ↓
Driver (paquete específico)
     ↓
SQL Server (PostgreSQL, MySQL, etc.)
     ↓
Base de Datos
```

**El flujo siempre es:**
1. Conectar a la BD
2. Preparar consulta
3. Ejecutar
4. Procesar resultados
5. Cerrar conexión

---

## PARTE II: POSTGRESQL (La Favorita de Go)

PostgreSQL es la más popular en Go. Es open source, poderosa y confiable.

### Instalación del Driver

```bash
go get github.com/lib/pq
```

### Paso 1: Conexión Básica

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func main() {
    // Paso 1: Conectar a PostgreSQL
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        "localhost", "5432", "usuario", "contraseña", "mibasedatos")

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Paso 2: Verificar conexión
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("✅ Conectado a PostgreSQL")
}
```

**¿Qué significa cada parámetro?**
- `host`: Servidor (localhost = tu máquina)
- `port`: Puerto (5432 es el default de PostgreSQL)
- `user`: Usuario de PostgreSQL
- `password`: Contraseña
- `dbname`: Nombre de la base de datos
- `sslmode=disable`: Sin encriptación SSL (para desarrollo)

### Paso 2: Crear Tabla

```go
func main() {
    db := conectarBD()
    defer db.Close()

    // Crear tabla de usuarios
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS usuarios (
        id SERIAL PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE,
        edad INT,
        fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

    _, err := db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("✅ Tabla creada")
}

func conectarBD() *sql.DB {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        "localhost", "5432", "usuario", "contraseña", "mibasedatos")

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    return db
}
```

### Paso 3: Insertar Datos

```go
func InsertarUsuario(db *sql.DB, nombre string, email string, edad int) {
    insertSQL := `
    INSERT INTO usuarios (nombre, email, edad)
    VALUES ($1, $2, $3)
    RETURNING id;
    `

    var id int
    err := db.QueryRow(insertSQL, nombre, email, edad).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Usuario insertado con ID: %d\n", id)
}

// Uso:
InsertarUsuario(db, "Juan", "juan@example.com", 30)
InsertarUsuario(db, "Ana", "ana@example.com", 25)
```

**¿Qué es `$1, $2, $3`?**

Markers para valores seguramente escapados (previene SQL injection):
```go
// ❌ MALO: Vulnerable a SQL injection
query := fmt.Sprintf("INSERT INTO usuarios VALUES ('%s', '%s')", nombre, email)

// ✅ BUENO: Seguro
query := "INSERT INTO usuarios VALUES ($1, $2)"
db.Exec(query, nombre, email)
```

### Paso 4: Leer Datos

```go
type Usuario struct {
    ID       int
    Nombre   string
    Email    string
    Edad     int
}

func ObtenerTodosUsuarios(db *sql.DB) []Usuario {
    rows, err := db.Query("SELECT id, nombre, email, edad FROM usuarios ORDER BY id")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var usuarios []Usuario

    for rows.Next() {
        var u Usuario
        err := rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        if err != nil {
            log.Fatal(err)
        }
        usuarios = append(usuarios, u)
    }

    return usuarios
}

func ObtenerUsuarioPorID(db *sql.DB, id int) Usuario {
    var u Usuario
    err := db.QueryRow("SELECT id, nombre, email, edad FROM usuarios WHERE id = $1", id).
        Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
    if err != nil {
        log.Fatal(err)
    }
    return u
}

// Uso:
usuarios := ObtenerTodosUsuarios(db)
for _, u := range usuarios {
    fmt.Printf("ID: %d, Nombre: %s, Email: %s, Edad: %d\n", u.ID, u.Nombre, u.Email, u.Edad)
}

usuario := ObtenerUsuarioPorID(db, 1)
fmt.Println(usuario)
```

### Paso 5: Actualizar y Eliminar

```go
func ActualizarUsuario(db *sql.DB, id int, edad int) {
    updateSQL := "UPDATE usuarios SET edad = $1 WHERE id = $2"
    
    result, err := db.Exec(updateSQL, edad, id)
    if err != nil {
        log.Fatal(err)
    }

    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("Filas actualizadas: %d\n", rowsAffected)
}

func EliminarUsuario(db *sql.DB, id int) {
    deleteSQL := "DELETE FROM usuarios WHERE id = $1"
    
    result, err := db.Exec(deleteSQL, id)
    if err != nil {
        log.Fatal(err)
    }

    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("Filas eliminadas: %d\n", rowsAffected)
}

// Uso:
ActualizarUsuario(db, 1, 31)
EliminarUsuario(db, 2)
```

---

## PARTE III: OPERACIONES CRUD EN TODAS LAS BASES DE DATOS

### Estructura de Datos (Usada en todos los ejemplos)

```go
type Usuario struct {
    ID    int
    Nombre string
    Email string
    Edad  int
}
```

---

### CRUD POSTGRESQL

**1. CREATE TABLE:**
```sql
CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    edad INT,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**2. CREATE (Insertar):**
```go
func CreateUsuariosPostgreSQL(db *sql.DB, nombre, email string, edad int) int {
    var id int
    err := db.QueryRow(`
        INSERT INTO usuarios (nombre, email, edad) VALUES ($1, $2, $3) RETURNING id`, 
        nombre, email, edad).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }
    return id
}
```

**3. READ (Leer):**
```go
func ReadUsuariosPostgreSQL(db *sql.DB) []Usuario {
    rows, _ := db.Query(`SELECT id, nombre, email, edad FROM usuarios ORDER BY id`)
    defer rows.Close()
    
    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        usuarios = append(usuarios, u)
    }
    return usuarios
}
```

**4. UPDATE (Actualizar):**
```go
func UpdateUsuariosPostgreSQL(db *sql.DB, id int, edad int) {
    _, err := db.Exec(`UPDATE usuarios SET edad = $1 WHERE id = $2`, edad, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

**5. DELETE (Eliminar):**
```go
func DeleteUsuariosPostgreSQL(db *sql.DB, id int) {
    _, err := db.Exec(`DELETE FROM usuarios WHERE id = $1`, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### CRUD MYSQL

**1. CREATE TABLE:**
```sql
CREATE TABLE IF NOT EXISTS usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    edad INT,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**2. CREATE (Insertar):**
```go
func CreateUsuariosMySQL(db *sql.DB, nombre, email string, edad int) int64 {
    result, err := db.Exec(`
        INSERT INTO usuarios (nombre, email, edad) VALUES (?, ?, ?)`,
        nombre, email, edad)
    if err != nil {
        log.Fatal(err)
    }
    id, _ := result.LastInsertId()
    return id
}
```

**3. READ (Leer):**
```go
func ReadUsuariosMySQL(db *sql.DB) []Usuario {
    rows, _ := db.Query(`SELECT id, nombre, email, edad FROM usuarios ORDER BY id`)
    defer rows.Close()
    
    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        usuarios = append(usuarios, u)
    }
    return usuarios
}
```

**4. UPDATE (Actualizar):**
```go
func UpdateUsuariosMySQL(db *sql.DB, id int, edad int) {
    _, err := db.Exec(`UPDATE usuarios SET edad = ? WHERE id = ?`, edad, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

**5. DELETE (Eliminar):**
```go
func DeleteUsuariosMySQL(db *sql.DB, id int) {
    _, err := db.Exec(`DELETE FROM usuarios WHERE id = ?`, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### CRUD ORACLE

**1. CREATE TABLE:**
```sql
CREATE TABLE usuarios (
    id NUMBER PRIMARY KEY,
    nombre VARCHAR2(100) NOT NULL,
    email VARCHAR2(100) UNIQUE,
    edad NUMBER,
    fecha_registro TIMESTAMP DEFAULT SYSDATE
);

-- Crear secuencia para auto-incremento
CREATE SEQUENCE usuarios_seq START WITH 1 INCREMENT BY 1;
```

**2. CREATE (Insertar):**
```go
func CreateUsuariosOracle(db *sql.DB, nombre, email string, edad int) int {
    var id int
    err := db.QueryRow(`
        INSERT INTO usuarios (id, nombre, email, edad) 
        VALUES (usuarios_seq.NEXTVAL, :1, :2, :3) 
        RETURNING id INTO :4`,
        nombre, email, edad, &id).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }
    return id
}
```

**3. READ (Leer):**
```go
func ReadUsuariosOracle(db *sql.DB) []Usuario {
    rows, _ := db.Query(`SELECT id, nombre, email, edad FROM usuarios ORDER BY id`)
    defer rows.Close()
    
    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        usuarios = append(usuarios, u)
    }
    return usuarios
}
```

**4. UPDATE (Actualizar):**
```go
func UpdateUsuariosOracle(db *sql.DB, id int, edad int) {
    _, err := db.Exec(`UPDATE usuarios SET edad = :1 WHERE id = :2`, edad, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

**5. DELETE (Eliminar):**
```go
func DeleteUsuariosOracle(db *sql.DB, id int) {
    _, err := db.Exec(`DELETE FROM usuarios WHERE id = :1`, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### CRUD SQL SERVER

**1. CREATE TABLE:**
```sql
CREATE TABLE usuarios (
    id INT PRIMARY KEY IDENTITY(1,1),
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    edad INT,
    fecha_registro DATETIME DEFAULT GETDATE()
);
```

**2. CREATE (Insertar):**
```go
func CreateUsuariosSQLServer(db *sql.DB, nombre, email string, edad int) int {
    var id int
    err := db.QueryRow(`
        INSERT INTO usuarios (nombre, email, edad)
        OUTPUT INSERTED.id
        VALUES (?, ?, ?)`,
        nombre, email, edad).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }
    return id
}
```

**3. READ (Leer):**
```go
func ReadUsuariosSQLServer(db *sql.DB) []Usuario {
    rows, _ := db.Query(`SELECT id, nombre, email, edad FROM usuarios ORDER BY id`)
    defer rows.Close()
    
    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        usuarios = append(usuarios, u)
    }
    return usuarios
}
```

**4. UPDATE (Actualizar):**
```go
func UpdateUsuariosSQLServer(db *sql.DB, id int, edad int) {
    _, err := db.Exec(`UPDATE usuarios SET edad = ? WHERE id = ?`, edad, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

**5. DELETE (Eliminar):**
```go
func DeleteUsuariosSQLServer(db *sql.DB, id int) {
    _, err := db.Exec(`DELETE FROM usuarios WHERE id = ?`, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### CRUD SQLITE

**1. CREATE TABLE:**
```sql
CREATE TABLE IF NOT EXISTS usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    email TEXT UNIQUE,
    edad INTEGER,
    fecha_registro DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

**2. CREATE (Insertar):**
```go
func CreateUsuariosSQLite(db *sql.DB, nombre, email string, edad int) int64 {
    result, err := db.Exec(`
        INSERT INTO usuarios (nombre, email, edad) VALUES (?, ?, ?)`,
        nombre, email, edad)
    if err != nil {
        log.Fatal(err)
    }
    id, _ := result.LastInsertId()
    return id
}
```

**3. READ (Leer):**
```go
func ReadUsuariosSQLite(db *sql.DB) []Usuario {
    rows, _ := db.Query(`SELECT id, nombre, email, edad FROM usuarios ORDER BY id`)
    defer rows.Close()
    
    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        usuarios = append(usuarios, u)
    }
    return usuarios
}
```

**4. UPDATE (Actualizar):**
```go
func UpdateUsuariosSQLite(db *sql.DB, id int, edad int) {
    _, err := db.Exec(`UPDATE usuarios SET edad = ? WHERE id = ?`, edad, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

**5. DELETE (Eliminar):**
```go
func DeleteUsuariosSQLite(db *sql.DB, id int) {
    _, err := db.Exec(`DELETE FROM usuarios WHERE id = ?`, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### CRUD COCKROACHDB

**Nota:** CockroachDB es compatible con SQL de PostgreSQL.

**1. CREATE TABLE:**
```sql
CREATE TABLE usuarios (
    id INT PRIMARY KEY DEFAULT unique_rowid(),
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    edad INT,
    fecha_registro TIMESTAMP DEFAULT now()
);
```

**2. CREATE (Insertar):**
```go
func CreateUsuariosCockroachDB(db *sql.DB, nombre, email string, edad int) int {
    var id int
    err := db.QueryRow(`
        INSERT INTO usuarios (nombre, email, edad) VALUES ($1, $2, $3) RETURNING id`,
        nombre, email, edad).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }
    return id
}
```

**3. READ (Leer):**
```go
func ReadUsuariosCockroachDB(db *sql.DB) []Usuario {
    rows, _ := db.Query(`SELECT id, nombre, email, edad FROM usuarios ORDER BY id`)
    defer rows.Close()
    
    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        usuarios = append(usuarios, u)
    }
    return usuarios
}
```

**4. UPDATE (Actualizar):**
```go
func UpdateUsuariosCockroachDB(db *sql.DB, id int, edad int) {
    _, err := db.Exec(`UPDATE usuarios SET edad = $1 WHERE id = $2`, edad, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

**5. DELETE (Eliminar):**
```go
func DeleteUsuariosCockroachDB(db *sql.DB, id int) {
    _, err := db.Exec(`DELETE FROM usuarios WHERE id = $1`, id)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### CRUD MONGODB

**1. CREATE (Insertar - No hay tabla, es colección):**
```go
func CreateUsuariosMongoDB(client *mongo.Client, nombre, email string, edad int) string {
    col := client.Database("mibasedatos").Collection("usuarios")
    ctx := context.Background()
    
    usuario := bson.M{
        "nombre": nombre,
        "email":  email,
        "edad":   edad,
        "fecha_registro": time.Now(),
    }
    
    resultado, err := col.InsertOne(ctx, usuario)
    if err != nil {
        log.Fatal(err)
    }
    
    return resultado.InsertedID.(primitive.ObjectID).Hex()
}
```

**2. READ (Leer):**
```go
func ReadUsuariosMongoDB(client *mongo.Client) []bson.M {
    col := client.Database("mibasedatos").Collection("usuarios")
    ctx := context.Background()
    
    cursor, err := col.Find(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)
    
    var usuarios []bson.M
    err = cursor.All(ctx, &usuarios)
    if err != nil {
        log.Fatal(err)
    }
    
    return usuarios
}
```

**3. UPDATE (Actualizar):**
```go
func UpdateUsuariosMongoDB(client *mongo.Client, id string, edad int) {
    col := client.Database("mibasedatos").Collection("usuarios")
    ctx := context.Background()
    
    objID, _ := primitive.ObjectIDFromHex(id)
    
    _, err := col.UpdateOne(
        ctx,
        bson.M{"_id": objID},
        bson.M{"$set": bson.M{"edad": edad}},
    )
    if err != nil {
        log.Fatal(err)
    }
}
```

**4. DELETE (Eliminar):**
```go
func DeleteUsuariosMongoDB(client *mongo.Client, id string) {
    col := client.Database("mibasedatos").Collection("usuarios")
    ctx := context.Background()
    
    objID, _ := primitive.ObjectIDFromHex(id)
    
    _, err := col.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        log.Fatal(err)
    }
}
```

---

### CRUD REDIS

**Nota:** Redis no tiene "tablas". Se usa clave-valor. Adaptamos el concepto.

**1. CREATE (Guardar):**
```go
func CreateUsuariosRedis(client *redis.Client, id string, nombre, email string, edad int) {
    ctx := context.Background()
    
    // Guardar usuario como hash
    err := client.HSet(ctx, "usuario:"+id, map[string]interface{}{
        "nombre": nombre,
        "email":  email,
        "edad":   edad,
    }).Err()
    if err != nil {
        log.Fatal(err)
    }
    
    // También guardar en un set para listado
    client.SAdd(ctx, "usuarios:keys", id)
}
```

**2. READ (Leer):**
```go
func ReadUsuariosRedis(client *redis.Client) map[string]map[string]string {
    ctx := context.Background()
    
    // Obtener todas las claves
    keys := client.SMembers(ctx, "usuarios:keys").Val()
    
    usuarios := make(map[string]map[string]string)
    for _, key := range keys {
        userData, _ := client.HGetAll(ctx, "usuario:"+key).Result()
        usuarios[key] = userData
    }
    
    return usuarios
}
```

**3. UPDATE (Actualizar):**
```go
func UpdateUsuariosRedis(client *redis.Client, id string, edad int) {
    ctx := context.Background()
    
    err := client.HSet(ctx, "usuario:"+id, "edad", edad).Err()
    if err != nil {
        log.Fatal(err)
    }
}
```

**4. DELETE (Eliminar):**
```go
func DeleteUsuariosRedis(client *redis.Client, id string) {
    ctx := context.Background()
    
    client.Del(ctx, "usuario:"+id)
    client.SRem(ctx, "usuarios:keys", id)
}
```

---

### TABLA COMPARATIVA: SINTAXIS CRUD POR BD

| Operación | PostgreSQL | MySQL | Oracle | SQL Server | SQLite | CockroachDB | MongoDB |
|-----------|-----------|-------|--------|-----------|--------|------------|---------|
| **Placeholder** | `$1, $2` | `?` | `:1, :2` | `?` | `?` | `$1, $2` | N/A |
| **Auto-increment** | SERIAL | AUTO_INCREMENT | SEQUENCE | IDENTITY | AUTOINCREMENT | unique_rowid() | ObjectID |
| **INSERT retorna ID** | RETURNING id | LastInsertId() | RETURNING | OUTPUT | LastInsertId() | RETURNING | InsertedID |
| **LIMIT** | LIMIT n | LIMIT n | ROWNUM <= n | TOP n | LIMIT n | LIMIT n | limit() |
| **Transacciones** | Sí (ACID) | Sí (ACID) | Sí (ACID) | Sí (ACID) | Sí (ACID) | Sí (ACID) | No (eventual) |
| **Complejidad** | Media | Media | Alta | Media | Baja | Media | Baja |

---

## PARTE IV: CONEXIONES EN TODAS LAS BDs

### ✅ BUENAS PRÁCTICAS

**1. Siempre usar prepared statements**
```go
// ✅ Seguro
db.Query("SELECT * FROM usuarios WHERE id = ?", userInput)

// ❌ Vulnerable a SQL injection
db.Query(fmt.Sprintf("SELECT * FROM usuarios WHERE id = %s", userInput))
```

**2. Cerrar recursos siempre**
```go
// ✅ Correcto
rows, _ := db.Query("SELECT * FROM usuarios")
defer rows.Close()

// ❌ Memory leak
rows, _ := db.Query("SELECT * FROM usuarios")
// rows nunca se cierra
```

**3. Manejar errores**
```go
// ✅ Correcto
err := db.Ping()
if err != nil {
    log.Fatal("No se puede conectar:", err)
}

// ❌ Ignorar errores
db.Ping()  // ¿Y si falla?
```

**4. Pool de conexiones**
```go
// ✅ Bueno
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(time.Hour)
```

**5. Usar context para timeouts**
```go
// ✅ Prevenir queries infinitas
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, "SELECT * FROM usuarios")
```

### ❌ ANTI-PATRONES

**1. Crear nueva conexión cada vez**
```go
// ❌ Terrible rendimiento
for i := 0; i < 1000; i++ {
    db := sql.Open("postgres", dsn)  // Nueva conexión CADA ITERACIÓN
    // query...
    db.Close()
}

// ✅ Reutilizar conexión
db := sql.Open("postgres", dsn)
defer db.Close()
for i := 0; i < 1000; i++ {
    // query...
}
```

**2. Ignorar ExecContext**
```go
// ❌ Sin timeout
result, _ := db.Exec("UPDATE usuarios SET...")  // ¿Y si tarda 1 hora?

// ✅ Con timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
result, _ := db.ExecContext(ctx, "UPDATE usuarios SET...")
```

**3. Queries N+1**
```go
// ❌ Lento: N queries
usuarios := ObtenerTodosUsuarios(db)
for _, u := range usuarios {
    pedidos := ObtenerPedidosPorUsuario(db, u.ID)  // Query por cada usuario
}

// ✅ Rápido: Una query
pedidos := ObtenerTodosPedidosConUsuarios(db)  // JOIN
```

---

## PARTE XII: OPERACIONES AVANZADAS (SELECT, UPDATE, DELETE, Procedimientos, Vistas)

### SELECT Avanzado

#### 1. SELECT con WHERE y ORDER BY

```go
func ObtenerUsuariosActivos(db *sql.DB) []Usuario {
    query := `
    SELECT id, nombre, email, edad 
    FROM usuarios 
    WHERE edad >= 18 
    ORDER BY nombre ASC
    LIMIT 10;
    `
    
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        err := rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        if err != nil {
            log.Fatal(err)
        }
        usuarios = append(usuarios, u)
    }
    
    return usuarios
}
```

#### 2. SELECT con GROUP BY y COUNT

```go
func ObtenerEstadísticas(db *sql.DB) {
    query := `
    SELECT edad, COUNT(*) as cantidad
    FROM usuarios
    GROUP BY edad
    HAVING COUNT(*) > 1
    ORDER BY edad;
    `
    
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var edad int
        var cantidad int
        err := rows.Scan(&edad, &cantidad)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Edad: %d, Cantidad: %d usuarios\n", edad, cantidad)
    }
}
```

#### 3. SELECT con JOIN

```go
type Pedido struct {
    ID       int
    UsuarioID int
    Monto    float64
    Usuario  string
}

func ObtenerPedidosConClientes(db *sql.DB) []Pedido {
    query := `
    SELECT p.id, p.usuario_id, p.monto, u.nombre
    FROM pedidos p
    INNER JOIN usuarios u ON p.usuario_id = u.id
    ORDER BY p.id;
    `
    
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var pedidos []Pedido
    for rows.Next() {
        var p Pedido
        err := rows.Scan(&p.ID, &p.UsuarioID, &p.Monto, &p.Usuario)
        if err != nil {
            log.Fatal(err)
        }
        pedidos = append(pedidos, p)
    }
    
    return pedidos
}
```

#### 4. SELECT con Agregación (SUM, AVG, MAX, MIN)

```go
func ObtenerEstadísticasVentas(db *sql.DB) {
    query := `
    SELECT 
        COUNT(*) as total_pedidos,
        SUM(monto) as monto_total,
        AVG(monto) as promedio,
        MAX(monto) as maximo,
        MIN(monto) as minimo
    FROM pedidos;
    `
    
    var totalPedidos int
    var montoTotal float64
    var promedio float64
    var maximo float64
    var minimo float64
    
    err := db.QueryRow(query).Scan(&totalPedidos, &montoTotal, &promedio, &maximo, &minimo)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Total pedidos: %d\n", totalPedidos)
    fmt.Printf("Monto total: $%.2f\n", montoTotal)
    fmt.Printf("Promedio: $%.2f\n", promedio)
    fmt.Printf("Máximo: $%.2f\n", maximo)
    fmt.Printf("Mínimo: $%.2f\n", minimo)
}
```

---

### UPDATE Avanzado

#### 1. UPDATE Simple

```go
func ActualizarEmail(db *sql.DB, id int, nuevoEmail string) {
    query := `UPDATE usuarios SET email = $1 WHERE id = $2`
    
    result, err := db.Exec(query, nuevoEmail, id)
    if err != nil {
        log.Fatal(err)
    }
    
    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("Filas actualizadas: %d\n", rowsAffected)
}
```

#### 2. UPDATE con Condiciones Múltiples

```go
func IncementarEdadMayoresDeEdad(db *sql.DB) {
    query := `
    UPDATE usuarios 
    SET edad = edad + 1 
    WHERE edad >= 18
    `
    
    result, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
    
    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("Usuarios actualizados: %d\n", rowsAffected)
}
```

#### 3. UPDATE con CASE (Condicional)

```go
func ActualizarTarifa(db *sql.DB) {
    query := `
    UPDATE usuarios 
    SET tarifa = CASE 
        WHEN edad < 18 THEN 'junior'
        WHEN edad >= 18 AND edad < 65 THEN 'adulto'
        ELSE 'senior'
    END
    WHERE tarifa IS NULL
    `
    
    result, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
    
    rowsAffected, _ := result.RowsAffected()
    fmt.Printf("Tarifas actualizadas: %d\n", rowsAffected)
}
```

#### 4. UPDATE con JOIN

```go
func ActualizarStockDePedidos(db *sql.DB) {
    query := `
    UPDATE productos 
    SET stock = stock - 1
    WHERE id IN (
        SELECT producto_id 
        FROM pedidos 
        WHERE estado = 'completado'
    )
    `
    
    result, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Productos actualizados: %d\n", result.RowsAffected())
}
```

---

### DELETE Avanzado

#### 1. DELETE Simple

```go
func EliminarUsuarioPorID(db *sql.DB, id int) {
    query := `DELETE FROM usuarios WHERE id = $1`
    
    result, err := db.Exec(query, id)
    if err != nil {
        log.Fatal(err)
    }
    
    rowsDeleted, _ := result.RowsAffected()
    fmt.Printf("Usuarios eliminados: %d\n", rowsDeleted)
}
```

#### 2. DELETE con Condiciones

```go
func EliminarUsuariosInactivos(db *sql.DB, diasSinActividad int) {
    query := `
    DELETE FROM usuarios 
    WHERE ultima_actividad < NOW() - INTERVAL '1 day' * $1
    `
    
    result, err := db.Exec(query, diasSinActividad)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Usuarios inactivos eliminados: %d\n", result.RowsAffected())
}
```

#### 3. DELETE con Transacción (Seguro)

```go
func EliminarUsuarioYSusPedidos(db *sql.DB, usuarioID int) error {
    // Iniciar transacción
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Eliminar pedidos del usuario
    _, err = tx.Exec(`DELETE FROM pedidos WHERE usuario_id = $1`, usuarioID)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    // Eliminar usuario
    _, err = tx.Exec(`DELETE FROM usuarios WHERE id = $1`, usuarioID)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    // Confirmar transacción
    err = tx.Commit()
    if err != nil {
        return err
    }
    
    fmt.Println("Usuario y sus pedidos eliminados correctamente")
    return nil
}

// Uso:
err := EliminarUsuarioYSusPedidos(db, 5)
if err != nil {
    log.Fatal(err)
}
```

---

### PROCEDIMIENTOS ALMACENADOS EN TODAS LAS BASES DE DATOS

Un **procedimiento almacenado** es una serie de sentencias SQL compiladas y guardadas en la BD. Se ejecutan directamente en el servidor, ganando velocidad.

#### PostgreSQL

**Crear Procedimiento:**
```sql
CREATE OR REPLACE FUNCTION calcular_descuento(monto DECIMAL)
RETURNS DECIMAL AS $$
BEGIN
    IF monto > 1000 THEN
        RETURN monto * 0.90;
    ELSIF monto > 500 THEN
        RETURN monto * 0.95;
    ELSE
        RETURN monto;
    END IF;
END;
$$ LANGUAGE plpgsql;
```

**Llamar desde Go:**
```go
func CalcularDescuentoPostgreSQL(db *sql.DB, monto float64) float64 {
    var descuento float64
    
    // En PostgreSQL se ejecuta como SELECT
    err := db.QueryRow(`SELECT calcular_descuento($1)`, monto).
        Scan(&descuento)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Descuento: $%.2f\n", descuento)
    return descuento
}
```

---

#### MySQL

**Crear Procedimiento:**
```sql
DELIMITER $$

CREATE PROCEDURE CalcularDescuento(
    IN monto DECIMAL(10,2),
    OUT descuento DECIMAL(10,2)
)
BEGIN
    IF monto > 1000 THEN
        SET descuento = monto * 0.90;
    ELSEIF monto > 500 THEN
        SET descuento = monto * 0.95;
    ELSE
        SET descuento = monto;
    END IF;
END$$

DELIMITER ;
```

**Llamar desde Go:**
```go
func CalcularDescuentoMySQL(db *sql.DB, monto float64) float64 {
    var descuento float64
    
    // MySQL usa CALL y parámetro de salida
    err := db.QueryRow(
        `CALL CalcularDescuento(?, @descuento); SELECT @descuento;`,
        monto,
    ).Scan(&descuento)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Descuento MySQL: $%.2f\n", descuento)
    return descuento
}
```

---

#### Oracle Database

**Crear Procedimiento:**
```sql
CREATE OR REPLACE PROCEDURE calcular_descuento(
    p_monto IN NUMBER,
    p_descuento OUT NUMBER
)
IS
BEGIN
    IF p_monto > 1000 THEN
        p_descuento := p_monto * 0.90;
    ELSIF p_monto > 500 THEN
        p_descuento := p_monto * 0.95;
    ELSE
        p_descuento := p_monto;
    END IF;
END;
/
```

**Llamar desde Go:**
```go
func CalcularDescuentoOracle(db *sql.DB, monto float64) float64 {
    var descuento float64
    
    // Oracle usa BEGIN ... END anónimo
    err := db.QueryRow(`
        BEGIN
            calcular_descuento(?, ?);
        END;
    `, monto).Scan(&descuento)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Descuento Oracle: $%.2f\n", descuento)
    return descuento
}

// Alternativa: usando ExecContext
func CalcularDescuentoOracleAlt(db *sql.DB, monto float64) float64 {
    var descuento float64
    
    stmt, err := db.Prepare(`
        BEGIN
            calcular_descuento(:1, :2);
        END;
    `)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
    
    err = stmt.QueryRow(monto).Scan(&descuento)
    if err != nil {
        log.Fatal(err)
    }
    
    return descuento
}
```

---

#### SQL Server (Microsoft)

**Crear Procedimiento:**
```sql
CREATE PROCEDURE sp_CalcularDescuento
    @monto DECIMAL(10,2),
    @descuento DECIMAL(10,2) OUTPUT
AS
BEGIN
    IF @monto > 1000
        SET @descuento = @monto * 0.90
    ELSE IF @monto > 500
        SET @descuento = @monto * 0.95
    ELSE
        SET @descuento = @monto
END
GO
```

**Llamar desde Go:**
```go
func CalcularDescuentoSQLServer(db *sql.DB, monto float64) float64 {
    var descuento float64
    
    // SQL Server usa DECLARE para variables locales
    err := db.QueryRow(`
        DECLARE @descuento DECIMAL(10,2);
        EXEC sp_CalcularDescuento @monto = ?, @descuento = @descuento OUTPUT;
        SELECT @descuento;
    `, monto).Scan(&descuento)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Descuento SQL Server: $%.2f\n", descuento)
    return descuento
}
```

---

#### SQLite

**Nota:** SQLite **no soporta procedimientos almacenados** como tal. En su lugar, usamos funciones SQL personalizadas en Go.

**Alternativa: Función SQL en Go**
```go
func RegistrarFuncionesSQLite(db *sql.DB) error {
    // Definir función personalizada en SQLite desde Go
    conn, err := db.Conn(context.Background())
    if err != nil {
        return err
    }
    defer conn.Close()
    
    // SQLite permite crear funciones dinámicamente (en versiones recientes)
    // O simplemente crear procedimientos como funciones en Go
    return nil
}

// En su lugar, escribir lógica en Go:
func CalcularDescuentoSQLite(db *sql.DB, monto float64) float64 {
    var descuento float64
    
    if monto > 1000 {
        descuento = monto * 0.90
    } else if monto > 500 {
        descuento = monto * 0.95
    } else {
        descuento = monto
    }
    
    return descuento
}
```

---

#### MongoDB (Agregaciones - Similar a Procedimientos)

MongoDB usa **agregaciones** en lugar de procedimientos. Es una serie de etapas que transforman datos.

**Agregación (como procedimiento):**
```go
func ObtenerResumenVentasMongoDBPipeline(client *mongo.Client) {
    col := client.Database("tienda").Collection("pedidos")
    ctx := context.Background()
    
    // Pipeline es como un procedimiento: serie de operaciones
    pipeline := mongo.Pipeline{
        bson.D{
            bson.E{Key: "$match", Value: bson.D{
                bson.E{Key: "estado", Value: "completado"},
            }},
        },
        bson.D{
            bson.E{Key: "$group", Value: bson.D{
                bson.E{Key: "_id", Value: "$usuario_id"},
                bson.E{Key: "total", Value: bson.D{
                    bson.E{Key: "$sum", Value: "$monto"},
                }},
                bson.E{Key: "cantidad", Value: bson.D{
                    bson.E{Key: "$sum", Value: 1},
                }},
            }},
        },
        bson.D{
            bson.E{Key: "$sort", Value: bson.D{
                bson.E{Key: "total", Value: -1},
            }},
        },
    }
    
    cursor, err := col.Aggregate(ctx, pipeline)
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)
    
    var results []bson.M
    err = cursor.All(ctx, &results)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Resultados agregación:")
    for _, result := range results {
        fmt.Printf("Usuario %v: Total $%v, Pedidos: %v\n", 
            result["_id"], result["total"], result["cantidad"])
    }
}
```

---

### EJEMPLO COMPLETO: PROCEDIMIENTO EXISTENTE EN TODAS LAS BDs

Imagina que en TODAS las bases de datos existe este procedimiento:

**PostgreSQL:**
```sql
CREATE OR REPLACE FUNCTION obtener_top_clientes(limite INT)
RETURNS TABLE(id INT, nombre VARCHAR, total_gastos DECIMAL) AS $$
BEGIN
    RETURN QUERY
    SELECT u.id, u.nombre, COALESCE(SUM(p.monto), 0) as total
    FROM usuarios u
    LEFT JOIN pedidos p ON u.id = p.usuario_id
    GROUP BY u.id, u.nombre
    ORDER BY total DESC
    LIMIT limite;
END;
$$ LANGUAGE plpgsql;
```

**MySQL:**
```sql
DELIMITER $$

CREATE PROCEDURE sp_ObtenerTopClientes(IN p_limite INT)
BEGIN
    SELECT 
        u.id,
        u.nombre,
        COALESCE(SUM(p.monto), 0) as total_gastos
    FROM usuarios u
    LEFT JOIN pedidos p ON u.id = p.usuario_id
    GROUP BY u.id, u.nombre
    ORDER BY total_gastos DESC
    LIMIT p_limite;
END$$

DELIMITER ;
```

**Oracle:**
```sql
CREATE OR REPLACE PROCEDURE obtener_top_clientes(
    p_limite IN NUMBER,
    p_cursor OUT SYS_REFCURSOR
)
IS
BEGIN
    OPEN p_cursor FOR
    SELECT u.id, u.nombre, NVL(SUM(p.monto), 0) as total_gastos
    FROM usuarios u
    LEFT JOIN pedidos p ON u.id = p.usuario_id
    GROUP BY u.id, u.nombre
    ORDER BY total_gastos DESC
    FETCH FIRST p_limite ROWS ONLY;
END obtener_top_clientes;
/
```

**SQL Server:**
```sql
CREATE PROCEDURE sp_ObtenerTopClientes
    @limite INT
AS
BEGIN
    SELECT TOP(@limite)
        u.id,
        u.nombre,
        ISNULL(SUM(p.monto), 0) as total_gastos
    FROM usuarios u
    LEFT JOIN pedidos p ON u.id = p.usuario_id
    GROUP BY u.id, u.nombre
    ORDER BY total_gastos DESC
END
GO
```

---

### INTERFAZ UNIFICADA EN GO (Llamar cualquier BD)

**Crear una interfaz que funcione con todas:**

```go
type ClienteTop struct {
    ID        int
    Nombre    string
    TotalGastos float64
}

type BaseDatos interface {
    ObtenerTopClientes(limite int) []ClienteTop
}

// Implementación PostgreSQL
type PostgreSQLDB struct {
    db *sql.DB
}

func (p *PostgreSQLDB) ObtenerTopClientes(limite int) []ClienteTop {
    rows, err := p.db.Query("SELECT * FROM obtener_top_clientes($1)", limite)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var clientes []ClienteTop
    for rows.Next() {
        var c ClienteTop
        rows.Scan(&c.ID, &c.Nombre, &c.TotalGastos)
        clientes = append(clientes, c)
    }
    return clientes
}

// Implementación MySQL
type MySQLDB struct {
    db *sql.DB
}

func (m *MySQLDB) ObtenerTopClientes(limite int) []ClienteTop {
    rows, err := m.db.Query("CALL sp_ObtenerTopClientes(?)", limite)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var clientes []ClienteTop
    for rows.Next() {
        var c ClienteTop
        rows.Scan(&c.ID, &c.Nombre, &c.TotalGastos)
        clientes = append(clientes, c)
    }
    return clientes
}

// Implementación SQL Server
type SQLServerDB struct {
    db *sql.DB
}

func (s *SQLServerDB) ObtenerTopClientes(limite int) []ClienteTop {
    rows, err := s.db.Query("EXEC sp_ObtenerTopClientes @limite = ?", limite)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var clientes []ClienteTop
    for rows.Next() {
        var c ClienteTop
        rows.Scan(&c.ID, &c.Nombre, &c.TotalGastos)
        clientes = append(clientes, c)
    }
    return clientes
}

// USAR INDISTINTAMENTE
func MostrarTopClientes(bd BaseDatos) {
    clientes := bd.ObtenerTopClientes(10)
    for _, c := range clientes {
        fmt.Printf("%s - $%.2f\n", c.Nombre, c.TotalGastos)
    }
}

// Uso:
var bd BaseDatos
// Cambiar según necesidad:
// bd = &PostgreSQLDB{db}
// bd = &MySQLDB{db}
// bd = &SQLServerDB{db}

MostrarTopClientes(bd)
```

---

### TABLA: SINTAXIS DE PROCEDIMIENTOS POR BD

| BD | Crear | Llamar desde Go | Parámetros |
|---|---|---|---|
| **PostgreSQL** | `CREATE OR REPLACE FUNCTION` | `SELECT funcion($1)` | `$1, $2, $3` |
| **MySQL** | `CREATE PROCEDURE` | `CALL proc(?, ?)` | `?` |
| **Oracle** | `CREATE PROCEDURE` | `BEGIN proc(:1, :2); END;` | `:1, :2` |
| **SQL Server** | `CREATE PROCEDURE` | `EXEC sp_proc @p1=?` | `@nombre` |
| **SQLite** | No soporta | Lógica en Go | N/A |
| **MongoDB** | Agregaciones | `col.Aggregate(pipeline)` | Pipeline |

---

### MEJORES PRÁCTICAS CON PROCEDIMIENTOS

✅ **BUENAS PRÁCTICAS:**

```go
// 1. Usar parámetros, nunca concatenar
err := db.QueryRow(`CALL obtener_datos(?, ?)`, param1, param2).Scan()

// 2. Manejar errores
if err != nil {
    log.Printf("Error llamando procedimiento: %v", err)
    return err
}

// 3. Cerrar rows
rows, _ := db.Query(...)
defer rows.Close()

// 4. Context con timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
db.QueryContext(ctx, ...)
```

❌ **ANTI-PATRONES:**

```go
// ❌ Concatenar directamente
query := fmt.Sprintf("CALL proc('%s')", param)  // SQL Injection!

// ❌ Ignorar errores
db.Query(...)  // ¿Se ejecutó bien?

// ❌ Sin timeout
db.Query("SELECT * FROM tabla_gigante")  // ¿Si demora 1 hora?
```

---

### FUNCIONES ALMACENADAS

Las funciones son similares a procedimientos pero retornan un valor.

#### PostgreSQL: Función que Retorna Tabla

```sql
CREATE OR REPLACE FUNCTION obtener_usuarios_mayores(edad_minima INT)
RETURNS TABLE(id INT, nombre VARCHAR, email VARCHAR, edad INT) AS $$
BEGIN
    RETURN QUERY
    SELECT u.id, u.nombre, u.email, u.edad
    FROM usuarios u
    WHERE u.edad >= edad_minima
    ORDER BY u.edad DESC;
END;
$$ LANGUAGE plpgsql;
```

#### Ejecutar desde Go

```go
type Usuario struct {
    ID     int
    Nombre string
    Email  string
    Edad   int
}

func ObtenerUsuariosMayores(db *sql.DB, edadMinima int) []Usuario {
    rows, err := db.Query(`
        SELECT * FROM obtener_usuarios_mayores($1)
    `, edadMinima)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        err := rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.Edad)
        if err != nil {
            log.Fatal(err)
        }
        usuarios = append(usuarios, u)
    }
    
    return usuarios
}

// Uso:
usuarios := ObtenerUsuariosMayores(db, 18)
for _, u := range usuarios {
    fmt.Printf("%s - %d años\n", u.Nombre, u.Edad)
}
```

---

### VISTAS (VIEWS)

Una vista es una tabla "virtual" basada en una consulta. Útil para simplificar queries complejas.

#### PostgreSQL: Crear Vista

```sql
-- Crear vista que muestre usuarios con sus pedidos
CREATE VIEW usuarios_con_pedidos AS
SELECT 
    u.id,
    u.nombre,
    u.email,
    COUNT(p.id) as cantidad_pedidos,
    SUM(p.monto) as monto_total
FROM usuarios u
LEFT JOIN pedidos p ON u.id = p.usuario_id
GROUP BY u.id, u.nombre, u.email;
```

#### Consultar Vista desde Go

```go
type UsuarioConPedidos struct {
    ID              int
    Nombre          string
    Email           string
    CantidadPedidos int
    MontoTotal      float64
}

func ObtenerUsuariosConPedidos(db *sql.DB) []UsuarioConPedidos {
    // Se consulta como una tabla normal
    rows, err := db.Query(`
        SELECT id, nombre, email, cantidad_pedidos, monto_total
        FROM usuarios_con_pedidos
        WHERE cantidad_pedidos > 0
        ORDER BY monto_total DESC
    `)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var usuarios []UsuarioConPedidos
    for rows.Next() {
        var u UsuarioConPedidos
        err := rows.Scan(&u.ID, &u.Nombre, &u.Email, &u.CantidadPedidos, &u.MontoTotal)
        if err != nil {
            log.Fatal(err)
        }
        usuarios = append(usuarios, u)
    }
    
    return usuarios
}

// Uso:
usuarios := ObtenerUsuariosConPedidos(db)
for _, u := range usuarios {
    fmt.Printf("%s - %d pedidos - $%.2f total\n", u.Nombre, u.CantidadPedidos, u.MontoTotal)
}
```

#### Actualizar Vista

```sql
-- Para actualizar vista (reemplazar la antigua)
CREATE OR REPLACE VIEW usuarios_con_pedidos AS
SELECT 
    u.id,
    u.nombre,
    u.email,
    COUNT(p.id) as cantidad_pedidos,
    SUM(p.monto) as monto_total,
    AVG(p.monto) as promedio_pedidos
FROM usuarios u
LEFT JOIN pedidos p ON u.id = p.usuario_id
GROUP BY u.id, u.nombre, u.email;
```

---

### TRANSACCIONES (ACID)

Las transacciones garantizan que múltiples operaciones se ejecutan en conjunto (todo o nada).

#### Inserción Transaccional Compleja

```go
func CrearPedidoCompleto(db *sql.DB, usuarioID int, monto float64) error {
    // Iniciar transacción
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Validar usuario existe
    var existe bool
    err = tx.QueryRow(`SELECT 1 FROM usuarios WHERE id = $1`, usuarioID).Scan(&existe)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("usuario no existe")
    }
    
    // Validar saldo suficiente
    var saldo float64
    err = tx.QueryRow(`SELECT saldo FROM usuarios WHERE id = $1`, usuarioID).Scan(&saldo)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    if saldo < monto {
        tx.Rollback()
        return fmt.Errorf("saldo insuficiente")
    }
    
    // Insertar pedido
    var pedidoID int
    err = tx.QueryRow(`
        INSERT INTO pedidos (usuario_id, monto, fecha)
        VALUES ($1, $2, NOW())
        RETURNING id
    `, usuarioID, monto).Scan(&pedidoID)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    // Actualizar saldo del usuario
    _, err = tx.Exec(`UPDATE usuarios SET saldo = saldo - $1 WHERE id = $2`, monto, usuarioID)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    // Insertar log de transacción
    _, err = tx.Exec(`
        INSERT INTO logs (usuario_id, accion, monto)
        VALUES ($1, 'compra', $2)
    `, usuarioID, monto)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    // Confirmar transacción
    err = tx.Commit()
    if err != nil {
        return err
    }
    
    fmt.Printf("Pedido #%d creado exitosamente\n", pedidoID)
    return nil
}

// Uso:
err := CrearPedidoCompleto(db, 1, 500.00)
if err != nil {
    fmt.Println("Error:", err)
}
```

#### Transacción con Savepoint

```go
func TransaccionConSavepoint(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Paso 1
    _, err = tx.Exec(`INSERT INTO usuarios (nombre, email) VALUES ($1, $2)`, "Juan", "juan@ex.com")
    if err != nil {
        tx.Rollback()
        return err
    }
    
    // Savepoint (punto de restauración)
    _, err = tx.Exec(`SAVEPOINT sp1`)
    if err != nil {
        tx.Rollback()
        return err
    }
    
    // Paso 2 (puede fallar)
    _, err = tx.Exec(`INSERT INTO pedidos (usuario_id, monto) VALUES ($1, $2)`, 1, -100) // Error: monto negativo
    if err != nil {
        // Rollback solo al savepoint
        tx.Exec(`ROLLBACK TO sp1`)
        fmt.Println("Pedido falló, usuario creado pero sin pedido")
    }
    
    // Commit
    return tx.Commit()
}
```

---

### TRIGGERS

Los triggers son acciones automáticas que se ejecutan cuando ocurren eventos.

#### PostgreSQL: Crear Trigger

```sql
-- Crear tabla de auditoría
CREATE TABLE usuarios_auditoria (
    id SERIAL PRIMARY KEY,
    usuario_id INT,
    accion VARCHAR(50),
    fecha TIMESTAMP DEFAULT NOW()
);

-- Crear función que ejecuta el trigger
CREATE OR REPLACE FUNCTION registrar_cambio_usuario()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO usuarios_auditoria (usuario_id, accion)
    VALUES (NEW.id, 'actualizado');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Crear trigger
CREATE TRIGGER trigger_actualizar_usuario
AFTER UPDATE ON usuarios
FOR EACH ROW
EXECUTE FUNCTION registrar_cambio_usuario();
```

#### Verificar Auditoría desde Go

```go
func ObtenerAuditoriaUsuarios(db *sql.DB) {
    rows, err := db.Query(`
        SELECT usuario_id, accion, fecha
        FROM usuarios_auditoria
        ORDER BY fecha DESC
        LIMIT 10
    `)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var usuarioID int
        var accion string
        var fecha time.Time
        
        err := rows.Scan(&usuarioID, &accion, &fecha)
        if err != nil {
            log.Fatal(err)
        }
        
        fmt.Printf("Usuario %d - %s - %s\n", usuarioID, accion, fecha.Format("2006-01-02 15:04:05"))
    }
}
```

---

### ÍNDICES (Optimización)

Los índices aceleran búsquedas pero ralentizan inserciones.

#### Crear Índices

```sql
-- Índice simple
CREATE INDEX idx_usuarios_email ON usuarios(email);

-- Índice compuesto
CREATE INDEX idx_pedidos_usuario_fecha ON pedidos(usuario_id, fecha);

-- Índice único
CREATE UNIQUE INDEX idx_email_unico ON usuarios(email);

-- Índice partial (solo registros que cumplen condición)
CREATE INDEX idx_usuarios_activos ON usuarios(id) WHERE activo = true;
```

#### Desde Go: Verificar Índices

```go
func MostrarIndicesPostgreSQL(db *sql.DB) {
    query := `
    SELECT indexname FROM pg_indexes
    WHERE tablename = 'usuarios'
    `
    
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Índices de usuarios:")
    for rows.Next() {
        var indexName string
        err := rows.Scan(&indexName)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("  - %s\n", indexName)
    }
}
```

---

## TABLA COMPARATIVA: CUÁL BD USAR

| Caso de Uso | Mejor Opción | Razón |
|------------|--------------|-------|
| **Datos estructurados** | PostgreSQL | Robusto, características, open source |
| **Compatibilidad MySQL** | MySQL | Estándar web, hosting barato |
| **Empresa grande** | Oracle | Power, soporte, pero caro |
| **Windows/Empresa Microsoft** | SQL Server | Integración, .NET compatible |
| **Aplicación pequeña**| SQLite | Cero config, perfecto |
| **Datos flexibles JSON** | MongoDB | Esquema dinámico, escala fácil |
| **Cache/Sesiones** | Redis | Ultra rápido, en memoria |
| **Búsqueda full-text** | Elasticsearch | Indexación potente |
| **Relaciones complejas** | Neo4j | Grafos (recomendaciones, redes) |
| **Time series** | InfluxDB/Prometheus | Optimizado para métricas |

---

## CONCLUSIÓN: El Ecosistema de Datos

Go brilla porque tiene **drivers excelentes** para casi cualquier BD:

- **SQL Relacionales**: PostgreSQL, MySQL, Oracle, SQL Server, SQLite
- **NoSQL Documentos**: MongoDB, CouchDB
- **NoSQL Clave-Valor**: Redis, Memcached
- **NoSQL Búsqueda**: Elasticsearch, Algolia
- **NoSQL Grafos**: Neo4j
- **Time Series**: InfluxDB, Prometheus

**Tu viaje:**
1. Aprende `database/sql` (interfaz común)
2. Elige una BD según tu caso
3. Instala driver específico
4. Escribe código seguro y eficiente

**Regla de Oro**: 
> "Usa PostgreSQL a menos que tengas motivo para otra cosa"

PostgreSQL es la opción más segura, poderosa y versátil. Es la favorita de comunidad Go.

¡Vamos! 🚀
