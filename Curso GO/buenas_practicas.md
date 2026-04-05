# Buenas Prácticas en Go - Guía Completa

## Introducción: El código vive más tiempo que crees

Escribir código que funciona es fácil. Escribir código que otros puedan **leer, entender y mantener** es difícil.

**Verdad incómoda**: Pasarás más tiempo leyendo código que escribiéndolo.
- **10%** del tiempo: escribir código nuevo
- **90%** del tiempo: leer, debuggear, modificar código existente

Por eso las "buenas prácticas" no son fanatismo. Son **respeto por el costo real de mantenimiento**.

Go tiene una filosofía clara al respecto:
> "Make it easy to read, not easy to write."

Las buenas prácticas son simplemente **escribir código como si alguien (probablemente tú en 6 meses) lo tuviera que entender sin contexto**.

---

## Parte I: Convenciones de Nombres

El nombre de una variable/función/paquete es lo primero que leen otros programadores.

### 1. Nombres claros y descriptivos

```go
// ❌ Mal: Abreviaturas crípticas
func cal(p, t float64) float64 {
    return p + (p * t)
}

u := struct {
    n string
    a int
}{"Alice", 30}

// ✅ Bien: Nombres auto-explicativos
func CalcularPrecioConImpuesto(precio, tasaImpuesto float64) float64 {
    return precio + (precio * tasaImpuesto)
}

usuario := struct {
    nombre string
    edad   int
}{"Alice", 30}
```

**Regla**: Si necesitas comentario para explicar qué es `x`, renómbralo.

### 2. Convención pública/privada (Mayúsculas)

```go
// En Go, mayúscula inicial = exportado (pública en otros lenguajes)
// minúscula inicial = privado (local al paquete)

type Persona struct {
    Nombre string      // ✅ Exportado
    Email  string      // ✅ Exportado
    edad   int         // ❌ Privado (no recomendado mezclar)
}

func (p *Persona) CalcularEdad() int {  // ✅ Exportado
    return time.Now().Year() - p.NacimientoAno
}

func (p *Persona) validarEmail() bool {  // Privado, ok para helpers
    return strings.Contains(p.Email, "@")
}
```

**Regla**: Si es parte de tu API pública, mayúscula. Si es helper interno, minúscula.

### 3. Nombres cortos vs descriptivos (contexto importa)

```go
// ✅ Variables loop pueden ser cortas (alcance es 1-2 líneas)
for i := 0; i < len(items); i++ {
    procesar(items[i])
}

// ❌ Variables de función deben ser claras
func p(u []map[string]interface{}) error {  // ¿Qué es p? ¿u?
    // ...
}

// ✅ Explícito
func ProcesarUsuarios(usuarios []map[string]interface{}) error {
    // ...
}

// ✅ En loops, una letra es OK
for _, usuario := range usuarios {
    procesar(usuario)
}
```

**Regla**: Alcance pequeño = nombre corto. Alcance grande = nombre largo.

### 4. Interfaces: Nombres con "-er"

```go
// ✅ Interfaces pequeñas, verbos como sufijo
type Writer interface {
    Write(p []byte) (n int, err error)
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// ❌ Nombres genéricos
type FileIO interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
    Close() error
}
```

**Regla**: Interfaces pequeñas y cohesivas. Nombra con "-er" si es acción.

### 5. Paquetes: nombres simples, no jerárquicos

```go
// ✅ Bien
package user
package account
package payment

// ❌ Redundante (el path ya dice la jerarquía)
package users      // import "myapp/users" ya es claro
package usermodels // No necesitas "models", usa package names simples

// ✅ Dentro del paquete
type User struct { }      // Aún es myapp.User cuando importas
func NewUser() *User { }  // Aún es myapp.NewUser()
```

**Regla**: Nombres de paquete simples, minúsculas, sin underscores.

---

## Parte II: Organización del Código

### 1. Estructura típica de paquete

```
myapp/
├── main.go                 # Punto de entrada
├── user/
│   ├── user.go            # Tipos e interfases
│   ├── user_test.go       # Tests
│   ├── service.go         # Lógica de negocio
│   └── service_test.go
├── account/
│   ├── account.go
│   ├── account_test.go
│   ├── service.go
│   └── service_test.go
├── storage/
│   ├── postgres.go        # Implementación DB
│   ├── postgres_test.go
│   └── storage.go         # Interfaz
└── go.mod
```

**Patrón**:
- `tipos.go` - Define structs/interfaces
- `service.go` - Lógica de negocio
- Implementaciones concretas en archivos separados
- `*_test.go` - Tests junto al código

### 2. Orden dentro de un archivo

```go
package user

import (
    "net/http"         // stdlib
    
    "github.com/pkg"   // terceros
    
    "myapp/storage"    // local
)

// Constantes
const (
    AdminRole = "admin"
    UserRole  = "user"
)

// Tipos
type User struct {
    ID    int
    Name  string
}

// Métodos de User
func (u *User) Validate() error {
    // ...
}

// Funciones públicas
func NewUser(name string) *User {
    // ...
}

func GetUser(id int) (*User, error) {
    // ...
}

// Funciones privadas
func validateEmail(email string) bool {
    // ...
}
```

**Patrón**: Imports → Consts → Types → Methods → Public funcs → Private funcs

### 3. Inicializa structs con constructores

```go
// ❌ Mal: Permitir constructor cero + validación después
type Config struct {
    Host string
    Port int
}

c := Config{}  // Inválido
if c.Port == 0 {
    c.Port = 8080
}

// ✅ Bien: Constructor que garantiza estado válido
type Config struct {
    host string    // privado
    port int
}

func NewConfig(host string, port int) (*Config, error) {
    if port < 1 || port > 65535 {
        return nil, fmt.Errorf("port inválido: %d", port)
    }
    return &Config{host, port}, nil
}

c, err := NewConfig("localhost", 8080)
// Garantizado que c es válido
```

---

## Parte III: Manejo de Errores (Repaso)

### 1. Siempre manejar errores

```go
// ❌ Malo
data, _ := ioutil.ReadFile("file.txt")

// ✅ Bien
data, err := ioutil.ReadFile("file.txt")
if err != nil {
    return fmt.Errorf("leer archivo: %w", err)
}
```

### 2. Wrap errores con contexto

```go
// ❌ Pierde información
if err != nil {
    return err
}

// ✅ Agrega contexto
if err != nil {
    return fmt.Errorf("procesar usuarios: %w", err)
}
```

### 3. Crear errores descriptivos

```go
// ❌ Genérico
if edad < 0 {
    return errors.New("error")
}

// ✅ Específico
if edad < 0 {
    return fmt.Errorf("edad no válida: %d (esperado >= 0)", edad)
}

// ✅ Con variables
if saldo < monto {
    return fmt.Errorf("saldo insuficiente: tienes %.2f, necesitas %.2f", 
        saldo, monto)
}
```

---

## Parte IV: Testing (Pruebas)

### 1. Tests co-ubicados con código

```go
// user.go
package user

func ValidarEmail(email string) bool {
    return strings.Contains(email, "@")
}

// user_test.go (en el mismo paquete)
package user

import "testing"

func TestValidarEmail(t *testing.T) {
    casos := []struct {
        email    string
        esperado bool
    }{
        {"alice@example.com", true},
        {"bob", false},
        {"", false},
    }
    
    for _, c := range casos {
        resultado := ValidarEmail(c.email)
        if resultado != c.esperado {
            t.Errorf("ValidarEmail(%q) = %v, esperado %v",
                c.email, resultado, c.esperado)
        }
    }
}
```

Ejecuta con: `go test ./user`

### 2. Tabla de casos de prueba

```go
func TestCalcularDescuento(t *testing.T) {
    tests := []struct {
        name     string
        precio   float64
        porcentaje int
        esperado float64
    }{
        {"sin descuento", 100, 0, 100},
        {"25% off", 100, 25, 75},
        {"50% off", 100, 50, 50},
        {"precio cero", 0, 50, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resultado := CalcularDescuento(tt.precio, tt.porcentaje)
            if resultado != tt.esperado {
                t.Errorf("got %.2f, want %.2f", resultado, tt.esperado)
            }
        })
    }
}
```

Ejecuta específico: `go test -run TestCalcularDescuento`

### 3. Benchmarks para performance

```go
func BenchmarkValidarEmail(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ValidarEmail("alice@example.com")
    }
}
```

Ejecuta: `go test -bench=. -benchmem`

### 4. Mocks y interfaces

```go
// storage.go
type Storage interface {
    Save(usuario *User) error
    Get(id int) (*User, error)
}

// service.go
type UserService struct {
    storage Storage
}

func NewUserService(s Storage) *UserService {
    return &UserService{storage: s}
}

func (s *UserService) CreateUser(nombre string) (*User, error) {
    u := &User{Name: nombre}
    return u, s.storage.Save(u)
}

// service_test.go
type MockStorage struct {
    saved []*User
}

func (m *MockStorage) Save(u *User) error {
    m.saved = append(m.saved, u)
    return nil
}

func (m *MockStorage) Get(id int) (*User, error) {
    return nil, errors.New("no encontrado")
}

func TestCreateUser(t *testing.T) {
    mock := &MockStorage{}
    service := NewUserService(mock)
    
    _, err := service.CreateUser("Alice")
    if err != nil {
        t.Fatalf("CreateUser failed: %v", err)
    }
    
    if len(mock.saved) != 1 {
        t.Errorf("esperaba 1 usuario salvado, got %d", len(mock.saved))
    }
}
```

---

## Parte V: Documentación y Comentarios

### 1. Documenta funciones públicas

```go
// ✅ Bien: Pequeña explicación, qué hace, qué retorna
// ValidarEmail verifica que el email tenga formato válido.
// Retorna true si es válido (contiene @), false en otros casos.
func ValidarEmail(email string) bool {
    return strings.Contains(email, "@")
}

// ✅ Para tipos
// User representa un usuario del sistema.
type User struct {
    ID    int       // ID único
    Name  string    // Nombre completo
    Email string    // Email de contacto
}

// ❌ Malo: Obvio
// Esta función suma dos números
func Sumar(a, b int) int {
    return a + b
}
```

**Regla**: Explica el **por qué**, no el **qué** (el código ya dice qué).

### 2. Código auto-documentado

```go
// ❌ Malo: Comentario críptico
if usuario.edad >= 18 && usuario.activo && usuario.verificado {
    // usuario puede votar
}

// ✅ Bien: Nombre claro
if PuedeVotar(usuario) {
    // ...
}

func PuedeVotar(u *User) bool {
    return u.Edad >= 18 && u.Activo && u.Verificado
}
```

### 3. Ejemplos en documentación

```go
// ParseID convierte un string a ID.
// Retorna error si el formato no es válido (ej: "user_123").
// 
// Ejemplo:
//     id, err := ParseID("user_123")
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(id)  // Output: 123
func ParseID(s string) (int, error) {
    // ...
}
```

Ejecuta ejemplos: `go test ./... -v`

---

## Parte VI: Performance y Eficiencia

### 1. Evita copias innecesarias

```go
// ❌ Copia el slice
func Procesar(items []Item) {
    copia := items  // Esto copia el header (3 campos)
}

// ✅ Pasa por referencia
func Procesar(items *[]Item) {
    // O mejor aún, no copies nada
}

// ❌ Copia la estructura
func ObtenerNombre(persona Persona) string {
    return persona.Nombre
}

// ✅ Pasa puntero si es grande
func ObtenerNombre(persona *Persona) string {
    return persona.Nombre
}
```

### 2. Pre-aloca slices si sabes tamaño

```go
// ❌ Append sin capacidad (re-aloca múltiples veces)
func GenerarNumeros(n int) []int {
    var numeros []int
    for i := 0; i < n; i++ {
        numeros = append(numeros, i)
    }
    return numeros
}

// ✅ Pre-aloca capacidad
func GenerarNumeros(n int) []int {
    numeros := make([]int, 0, n)  // capacidad n
    for i := 0; i < n; i++ {
        numeros = append(numeros, i)
    }
    return numeros
}
```

### 3. Usa strings.Builder para concatenación

```go
// ❌ Crea múltiples strings (ineficiente)
resultado := ""
for _, palabra := range palabras {
    resultado += palabra + " "
}

// ✅ Usa Builder
var sb strings.Builder
for _, palabra := range palabras {
    sb.WriteString(palabra)
    sb.WriteString(" ")
}
resultado := sb.String()
```

### 4. El operador := vs =

```go
// ✅ Usa := en el scope mínimo necesario
if datos, err := LeerArchivo(); err != nil {
    fmt.Println("Error:", err)
} else {
    Procesar(datos)
}
// datos no existe acá (scope limitado)

// ❌ Variable que vive más que la necesita
var datos []byte
var err error
if datos, err = LeerArchivo(); err != nil {
    fmt.Println("Error:", err)
}
Procesar(datos)
```

---

## Parte VII: Seguridad Básica

### 1. Valida entrada siempre

```go
// ❌ Confianza ciega
func DescuentoPorcentaje(p float64) float64 {
    return p * 0.9
}

desc := DescuentoPorcentaje(-100)  // ¿Descuento negativo?

// ✅ Valida entrada
func CalcularConDescuento(precio float64, porcentaje int) (float64, error) {
    if precio < 0 {
        return 0, fmt.Errorf("precio negativo: %.2f", precio)
    }
    if porcentaje < 0 || porcentaje > 100 {
        return 0, fmt.Errorf("porcentaje fuera de rango: %d", porcentaje)
    }
    return precio * (1 - float64(porcentaje)/100), nil
}
```

### 2. No loguees datos sensibles

```go
// ❌ Malo: puede exponer información sensible
func Login(username, password string) error {
    log.Printf("Login intento: usuario=%s, password=%s", username, password)
    // ...
}

// ✅ Bien: solo lo necesario
func Login(username, password string) error {
    log.Printf("Login intento: usuario=%s", username)
    // Nunca logs passwords
    // ...
}
```

### 3. Cierra recursos

```go
// ❌ Memory leak
func ReadConfig(path string) []byte {
    f, _ := os.Open(path)
    // 'f' nunca se cierra
    return ioutil.ReadAll(f)
}

// ✅ Usa defer
func ReadConfig(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()  // Garantizado cerrar
    return ioutil.ReadAll(f)
}
```

---

## Parte VIII: Estilo de Código

### 1. Usa gofmt

```bash
# Formatea automáticamente según estándares Go
gofmt -w .

# O integrado en VS Code: Format on Save
```

Go tiene ONE style, forzado por herramientas. No hay 10 debates sobre espacios/tabs.

### 2. Interface saturation (interfaces pequeñas)

```go
// ❌ Interfaz monolítica
type DataStore interface {
    Create(item *Item) error
    Read(id int) (*Item, error)
    Update(item *Item) error
    Delete(id int) error
    List() ([]*Item, error)
    Close() error
}

// ✅ Interfaces cohesivas pequeñas
type Writer interface {
    Write(item *Item) error
}

type Reader interface {
    Read(id int) (*Item, error)
}

type Closer interface {
    Close() error
}
```

### 3. Evita godoc anti-patterns

```go
// ❌ Malo
func F(a int) int {
    return a * 2
}

// ✅ Bien
// Double returns twice the input integer.
//
//     Double(5) returns 10
func Double(n int) int {
    return n * 2
}
```

### 4. Visibilidad explícita

```go
// ❌ Mezclar público/privado sin razón
type Database struct {
    Host     string      // público
    username string      // privado
    db       *sql.DB     // privado
}

// ✅ Consistente
type Database struct {
    host string          // Privado por defecto (constructor controla)
    db    *sql.DB
}

func NewDatabase(host string) (*Database, error) {
    // Constructor valida
    d := &Database{host: host}
    // ... inicializa
    return d, nil
}
```

---

## Parte IX: Patrones Comunes

### 1. Constructor pattern

```go
type Logger struct {
    file   *os.File
    level  LogLevel
    format string
}

func NewLogger(path string, level LogLevel) (*Logger, error) {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }
    return &Logger{
        file:   f,
        level:  level,
        format: "[%s] %s",
    }, nil
}

func (l *Logger) Close() error {
    return l.file.Close()
}
```

### 2. Builder pattern

```go
type QueryBuilder struct {
    table  string
    where  []string
    order  string
    limit  int
}

func (qb *QueryBuilder) Table(t string) *QueryBuilder {
    qb.table = t
    return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
    qb.where = append(qb.where, condition)
    return qb
}

func (qb *QueryBuilder) Build() string {
    // ...
}

// Uso
query := (&QueryBuilder{}).
    Table("users").
    Where("age > 18").
    Where("active = true").
    Build()
```

### 3. Dependency Injection

```go
// ✅ Inyecta dependencias, no las crees
type UserService struct {
    storage Storage
    mail    MailService
}

func NewUserService(s Storage, m MailService) *UserService {
    return &UserService{storage: s, mail: m}
}

func (us *UserService) RegisterUser(email string) error {
    // Usa storage y mail inyectados
    user := &User{Email: email}
    if err := us.storage.Save(user); err != nil {
        return err
    }
    return us.mail.SendWelcome(email)
}
```

---

## Parte X: Anti-patrones a evitar

```go
// ❌ 1. Ignorar errores
data, _ := ioutil.ReadFile("file")

// ❌ 2. Panic en código de producción
if x == nil {
    panic("x no puede ser nil")  // Usa error en su lugar
}

// ❌ 3. Globales mutables
var database *Database  // Evita mutación global

// ❌ 4. Omitir validación
func TransferMoney(from, to *Account, amount float64) {
    // Sin validar amounts, accounts, etc
    from.Balance -= amount
    to.Balance += amount
}

// ❌ 5. Código duplicado
// Si repites código 3 veces, EXTRAE función

// ❌ 6. Comentarios desactualizados
// Carga usuarios de BD
// (pero la función ahora es de otro lugar)
func GetUsers() []*User { /**/ }

// ❌ 7. Nesting profundo
for {
    for {
        for {
            // Muy profundo, refactoriza
        }
    }
}
```

---

## Parte XI: Herramientas Go

### 1. gofmt - Formateador

```bash
gofmt -w archivo.go      # Formatea en lugar
go fmt ./...             # Todos los archivos
```

### 2. golint / golangci-lint - Linter

```bash
golangci-lint run ./...
```

Detecta:
- Código inútil
- Variables no usadas
- Nombres no idiomáticos
- Potenciales bugs

### 3. go vet - Analizador estático

```bash
go vet ./...
```

Detecta:
- Llamadas a funciones con número inválido de argumentos
- Construcciones sospechosas

### 4. godoc - Documentación

```bash
godoc -http=:6060      # Servidor local de docs
```

Genera documentación a partir de comentarios.

---

## Resumen: Checklist de Buenas Prácticas

- [ ] **Nombres**: Claros, descriptivos, convenciones Go
- [ ] **Errores**: Siempre manejar, wrap con contexto
- [ ] **Tests**: Coverage > 80%, tabla de casos
- [ ] **Documentación**: Funciones públicas documentadas
- [ ] **Performance**: No optimize prematuramente, pero evita obvias
- [ ] **Seguridad**: Valida entrada, no logs sensitivos
- [ ] **Organización**: Paquetes pequeños, responsabilidad única
- [ ] **Recursos**: Cierra archivos, usa defer
- [ ] **Estilo**: Usa gofmt, sigue convenciones Go
- [ ] **Interfaces**: Pequeñas, cohesivas, bien nombradas

---

## Conclusión: El camino a código profesional

Las buenas prácticas no son dogma. Son:
- **Consenso de lo que funciona**
- **Respeto por quienes mantienen el código**
- **Protección contra bugs comunes**
- **Escalabilidad y mantenibilidad**

Go fuerza muchas de estas cosas (formato, errores explícitos). Eso es su fortaleza.

**Domina estas prácticas y escribirás código del que te orgulle, 6 meses después.**
