# POO en Go - Programación Orientada a Objetos

## Introducción: El enfoque radical de Go para OOP

Go tiene un enfoque radical para orientación a objetos: **sin clases**.

Esto confunde a muchos programadores que vienen de Java, C++, o Python. Piensan: "¿Sin clases? ¿Cómo hago OOP?"

La respuesta es elegante: Go usa **structs y métodos** en lugar de clases. Es como OOP, pero más flexible y simple.

**¿Por qué Go rechaza la herencia?**
Según los creadores de Go, la herencia es sobrevalorada. Causa:
- Jerarquías profundas y complejas (la famosa "gorilla/jungle problem")
- Acoplamiento débil entre clases
- Dificultad para cambiar código sin romper todo

**El enfoque de Go**: Usa **composición** (combinar tipos) e **interfaces** (contratos) en lugar de herencia.

Resultado: Código más flexible, predecible y mantenible.

### Metáfora: Lego vs Cadena de herencia

**Herencia tradicional** (Java):
```
Animal → Mamífero → Perro → MiPerro
         (cadena lineal, inflexible)
```

Si necesitas un "Perro volador", necesitas redeseñar toda la jerarquía.

**Composición en Go** (Lego):
```
Perro = {Cabeza, Patas, Cuerpo, Alas}
Gato = {Cabeza, Patas, Cuerpo}
Pajaro = {Cabeza, Alas, Cuerpo}
```

Puedes combinar comportamientos libremente.

---

Go **no es un lenguaje POO tradicional** (sin clases, sin herencia), pero permite código orientado a objetos a través de **structs, métodos e interfaces**. Este es el "enfoque de Go" y es más flexible.

## 1. Structs - Objetos

Un struct es tu equivalente a una clase:

```go
// Definir un objeto
type Persona struct {
    Nombre string
    Edad   int
    Salario float64
}

func main() {
    // Crear instancia
    p := Persona{
        Nombre: "Alice",
        Edad: 30,
        Salario: 50000,
    }
    
    fmt.Println(p.Nombre)
    fmt.Println(p.Edad)
}
```

### Encapsulación con mayúsculas/minúsculas

```go
// Público (accesible desde otros paquetes)
type Persona struct {
    Nombre string  // Público (comienza con mayúscula)
    edad   int     // Privado (comienza con minúscula)
}

func main() {
    p := Persona{Nombre: "Alice"}
    fmt.Println(p.Nombre)  // ✅ Funciona
    fmt.Println(p.edad)    // ❌ Error: no exportado
}
```

**Convención de Go**: No hay `public`/`private`, usa mayúsculas.

### Constructores

```go
// No hay constructores automáticos, creamos funciones
type Usuario struct {
    ID   int
    Nombre string
    Email string
    activo bool  // privado
}

// Constructor (convención: usar "New" + NombreTipo)
func NewUsuario(id int, nombre, email string) *Usuario {
    return &Usuario{
        ID: id,
        Nombre: nombre,
        Email: email,
        activo: true,  // valor por defecto
    }
}

func main() {
    u := NewUsuario(1, "Bob", "bob@example.com")
    fmt.Println(u.Nombre)
}
```

---

## 2. Métodos - Funciones asociadas a objetos

Un **método** es una función con receptor (vinculada a un tipo).

```go
type Persona struct {
    Nombre string
    Edad   int
}

// Método con receptor por VALOR (no modifica original)
func (p Persona) Presentarse() string {
    return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
}

// Método con receptor por PUNTERO (modifica original)
func (p *Persona) CumplirAños() {
    p.Edad++
}

func main() {
    persona := Persona{"Alice", 30}
    
    fmt.Println(persona.Presentarse())  // Hola, soy Alice y tengo 30 años
    
    persona.CumplirAños()
    fmt.Println(persona.Edad)  // 31
}
```

### Receptor por valor vs puntero

```go
// VALOR: Copia la estructura
func (p PersonaV) VerEdad() int {
    return p.Edad
}

// PUNTERO: Accede directamente
func (p *PersonaP) VerEdad() int {
    return p.Edad
}

// Regla práctica:
// - Usa VALOR si no modificas
// - Usa PUNTERO si modificas
// - Usa PUNTERO si la struct es grande (evitar copia)
```

---

## 3. Interfaces - Contrato de métodos

Una **interface** define un conjunto de métodos que un tipo debe implementar.

```go
// Contrato: todo lo que es "Animal" debe tener estos métodos
type Animal interface {
    Hacer() string
    Dormir()
}

// Implementación 1
type Perro struct {
    Nombre string
}

func (p Perro) Hacer() string {
    return "¡Guau!"
}

func (p Perro) Dormir() {
    fmt.Println(p.Nombre, "está durmiendo")
}

// Implementación 2
type Gato struct {
    Nombre string
}

func (g Gato) Hacer() string {
    return "¡Miau!"
}

func (g Gato) Dormir() {
    fmt.Println(g.Nombre, "está durmiendo")
}

// Función que acepta Animal (cualquier tipo que lo implemente)
func PasarTiempo(a Animal) {
    fmt.Println("El animal", a.Hacer())
    a.Dormir()
}

func main() {
    perro := Perro{"Rex"}
    gato := Gato{"Félix"}
    
    PasarTiempo(perro)  // Funciona
    PasarTiempo(gato)   // Funciona también
}
```

**Ventaja**: Mismo código funciona con múltiples tipos (polimorfismo)

### Interfaces integradas

```go
// Go proporciona interfaces comunes

// io.Reader: leer datos
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer: escribir datos
type Writer interface {
    Write(p []byte) (n int, err error)
}

// fmt.Stringer: convertir a string
type Stringer interface {
    String() string
}

// Implementar Stringer
type Punto struct {
    X, Y float64
}

func (p Punto) String() string {
    return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

func main() {
    p := Punto{3, 4}
    fmt.Println(p)  // Llama String() automáticamente
}
```

---

## 4. Composición - Herencia "simulada"

Go NO tiene herencia de clases. En su lugar, usa **composición** (embedding):

```go
type Persona struct {
    Nombre string
    Edad   int
}

func (p Persona) Presentarse() string {
    return fmt.Sprintf("Soy %s", p.Nombre)
}

// Empleado CONTIENE Persona (composición)
type Empleado struct {
    Persona      // "Embedir" otro struct
    Puesto string
}

func main() {
    e := Empleado{
        Persona: Persona{"Alice", 30},
        Puesto: "Ingeniera",
    }
    
    // Acceso directo a campos de Persona (sin e.Persona.Nombre)
    fmt.Println(e.Nombre)  // Alice
    fmt.Println(e.Puesto)  // Ingeniera
    
    // Métodos de Persona también disponibles
    fmt.Println(e.Presentarse())  // Soy Alice
}
```

### Método sobrescrito en struct embebido

```go
type Persona struct {
    Nombre string
}

func (p Persona) Saludar() string {
    return "Hola, soy persona"
}

type Empleado struct {
    Persona
    Puesto string
}

func (e Empleado) Saludar() string {
    return "Hola, soy empleado"
}

func main() {
    e := Empleado{Persona{"Alice"}, "Ingeniera"}
    
    fmt.Println(e.Saludar())         // Hola, soy empleado
    fmt.Println(e.Persona.Saludar()) // Hola, soy persona (acceso explícito)
}
```

---

## 5. Getters y Setters

Para controlar el acceso a campos privados:

```go
type Cuenta struct {
    saldo float64  // Privado
}

// Getter: acceso de lectura
func (c *Cuenta) Saldo() float64 {
    return c.saldo
}

// Setter: acceso de escritura con validación
func (c *Cuenta) Depositar(cantidad float64) error {
    if cantidad <= 0 {
        return fmt.Errorf("cantidad debe ser positiva")
    }
    c.saldo += cantidad
    return nil
}

func (c *Cuenta) Retirar(cantidad float64) error {
    if cantidad > c.saldo {
        return fmt.Errorf("saldo insuficiente")
    }
    c.saldo -= cantidad
    return nil
}

func main() {
    cuenta := &Cuenta{saldo: 100}
    
    fmt.Println("Saldo:", cuenta.Saldo())  // 100
    
    cuenta.Depositar(50)
    fmt.Println("Saldo:", cuenta.Saldo())  // 150
    
    if err := cuenta.Retirar(200); err != nil {
        fmt.Println("Error:", err)  // Error: saldo insuficiente
    }
}
```

---

## 6. Polimorfismo con interfaces

```go
// Interface define el contrato
type Vehiculo interface {
    Conducir() string
    Parar() string
}

// Implementación 1
type Auto struct {
    Marca string
}

func (a Auto) Conducir() string {
    return a.Marca + " está conduciendo"
}

func (a Auto) Parar() string {
    return a.Marca + " se detuvo"
}

// Implementación 2
type Moto struct {
    Modelo string
}

func (m Moto) Conducir() string {
    return "Moto " + m.Modelo + " está conduciendo"
}

func (m Moto) Parar() string {
    return "Moto " + m.Modelo + " se detuvo"
}

// Función que funciona con CUALQUIER Vehiculo
func Viajar(v Vehiculo) {
    fmt.Println(v.Conducir())
    // ... hacer cosas
    fmt.Println(v.Parar())
}

func main() {
    auto := Auto{"Toyota"}
    moto := Moto{"Harley"}
    
    Viajar(auto)   // Funciona
    Viajar(moto)   // Funciona también
}
```

---

## 7. Type assertion - Descubrir tipo en runtime

```go
var i interface{} = "Hola"

// Type assertion simple
s := i.(string)
fmt.Println(s)  // Hola

// Type assertion con verificación
s, ok := i.(string)
if ok {
    fmt.Println("Es string:", s)
} else {
    fmt.Println("No es string")
}

// Fallar sin verificación
// n := i.(int)  // PANIC si no es int
```

### Type switch

```go
func Procesar(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Println("String de longitud:", len(v))
    case int:
        fmt.Println("Entero:", v)
    case []int:
        fmt.Println("Slice de ints:", v)
    default:
        fmt.Println("Tipo desconocido")
    }
}

func main() {
    Procesar("Hola")         // String de longitud: 5
    Procesar(42)             // Entero: 42
    Procesar([]int{1, 2, 3}) // Slice de ints: [1 2 3]
}
```

---

## 8. Ejemplo completo: Sistema bancario

```go
package main

import "fmt"

// Interface: contrato para cuentas
type Cuenta interface {
    Depositar(monto float64) error
    Retirar(monto float64) error
    ObtenerSaldo() float64
}

// Implementación: Cuenta ahorros
type CuentaAhorros struct {
    titular string
    saldo   float64
    interes float64
}

func NewCuentaAhorros(titular string) *CuentaAhorros {
    return &CuentaAhorros{
        titular: titular,
        saldo:   0,
        interes: 0.05,
    }
}

func (c *CuentaAhorros) Depositar(monto float64) error {
    if monto <= 0 {
        return fmt.Errorf("monto debe ser positivo")
    }
    c.saldo += monto
    return nil
}

func (c *CuentaAhorros) Retirar(monto float64) error {
    if monto > c.saldo {
        return fmt.Errorf("saldo insuficiente")
    }
    c.saldo -= monto
    return nil
}

func (c *CuentaAhorros) ObtenerSaldo() float64 {
    return c.saldo
}

// Implementación: Cuenta corriente
type CuentaCorriente struct {
    titular string
    saldo   float64
}

func NewCuentaCorriente(titular string) *CuentaCorriente {
    return &CuentaCorriente{
        titular: titular,
        saldo:   0,
    }
}

func (c *CuentaCorriente) Depositar(monto float64) error {
    if monto <= 0 {
        return fmt.Errorf("monto debe ser positivo")
    }
    c.saldo += monto
    return nil
}

func (c *CuentaCorriente) Retirar(monto float64) error {
    if monto > c.saldo {
        return fmt.Errorf("saldo insuficiente")
    }
    c.saldo -= monto
    return nil
}

func (c *CuentaCorriente) ObtenerSaldo() float64 {
    return c.saldo
}

// Banco: trabaja con cualquier tipo de Cuenta
type Banco struct {
    cuentas []Cuenta
}

func (b *Banco) AgregarCuenta(c Cuenta) {
    b.cuentas = append(b.cuentas, c)
}

func (b *Banco) ReporteSaldos() {
    for i, c := range b.cuentas {
        fmt.Printf("Cuenta %d: $%.2f\n", i+1, c.ObtenerSaldo())
    }
}

func main() {
    banco := &Banco{}
    
    ahorros := NewCuentaAhorros("Alice")
    corriente := NewCuentaCorriente("Bob")
    
    ahorros.Depositar(1000)
    corriente.Depositar(500)
    
    banco.AgregarCuenta(ahorros)
    banco.AgregarCuenta(corriente)
    
    banco.ReporteSaldos()
    // Cuenta 1: $1000.00
    // Cuenta 2: $500.00
}
```

---

## 9. Tabla resumen: Go vs POO tradicional

| Concepto | Go | POO Tradicional |
|----------|----|----|
| Clases | ❌ Usa structs | ✅ Clases nativas |
| Herencia | ❌ Composición | ✅ Herencia de clases |
| Encapsulación | ✅ Mayúsculas/minúsculas | ✅ public/private |
| Polimorfismo | ✅ Interfaces | ✅ Overriding, interfaces |
| Métodos | ✅ Receptor implícito | ✅ this/self |
| Constructores | Manual (New...) | ✅ __init__ |
| Abstracción | ✅ Interfaces | ✅ Clases abstractas |

---

## 10. Mejores prácticas

### ✅ Bien
```go
// Usar composición en lugar de herencia
type Empleado struct {
    Persona
    Salario float64
}

// Métodos claros
func (e *Empleado) CalcularBono() float64 {
    return e.Salario * 0.1
}

// Interfaces pequeñas y cohesivas
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Getter/setter solo si necesitas lógica
func (c *Cuenta) Depositar(monto float64) error {
    // validación
    c.saldo += monto
}
```

### ❌ Mal
```go
// Intentar simular herencia profunda
type A struct{}
type B struct { A }
type C struct { B }
type D struct { C }  // Jerarquía profunda - complicado

// Métodos poco claros
func (p *Persona) A() {}
func (p *Persona) B() {}
func (p *Persona) C() {}

// Interfaces muy grandes
type Todo interface {
    Read()
    Write()
    Delete()
    Update()
    List()
    // ...30 métodos masVer malo
}
```

---

## 11. Conclusión - El Camino del OOP en Go

Go's enfoque de POO es radicalmente diferente, pero es lo correcto.

### Los pilares de Go para OOP

1. **Structs**: Agrupan datos (campos)
2. **Métodos**: Agrupan comportamiento (funciones con receiver)
3. **Composición**: Combina structs para code reuse
4. **Interfaces**: Define comportamientos esperados
5. **Duck typing**: Si se ve como un pato y suena como un pato, es un pato

### Comparación: Go vs Java/Python

| Aspecto | Java | Go |
|---------|------|-----|
| Clases | ✅ Sí | ❌ No (structs) |
| Herencia | ✅ Múltiple | ❌ No, composición |
| Polimorfismo | ✅ Interfaces | ✅ Interfaces |
| Encapsulación | ✅ public/private | ✅ Mayús./minús. |
| Constructor | ✅ Explicit | ❌ Convención |
| Abstracción | ✅ Abstract classess | ✅ Interfaces |

### El cambio mental necesario

**Si vienes de Java/C++**:
- Olvida herencia. Usa composición.
- Olvida métodos privados. Usa parámetros pequeños.
- Olvida getters/setters. Usa métodos cuando haya lógica.

**Si vienes de Python**:
- Necesitarás tipos más explícitos (structs vs dicts)
- Las interfaces son más rígidas que duck typing puro
- El receiver es extraño al principio, luego es natural

### Patrones comunes en Go OOP

```go
// Patrón 1: Constructors
type Logger struct {
    file *os.File
    buf  *bufio.Writer
}

func NewLogger(path string) (*Logger, error) {
    f, err := os.Create(path)
    if err != nil { return nil, err }
    return &Logger{file: f, buf: bufio.NewWriter(f)}, nil
}

// Patrón 2: Builder pattern
type RequestBuilder struct {
    method string
    url    string
    body   string
}

func (b *RequestBuilder) SetMethod(m string) *RequestBuilder {
    b.method = m
    return b
}

r := RequestBuilder{}.
    SetMethod("GET").
    SetURL("https://example.com")

// Patrón 3: Composición, no herencia
type Reader interface {
    Read() string
}

type Logger struct {
    r Reader
}

func (l *Logger) Log() {
    fmt.Println(l.r.Read())
}

// Patrón 4: Option pattern
func NewServer(opts ...func(*Server)) *Server {
    s := &Server{port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}

s := NewServer(
    WithPort(9000),
    WithTLS(true),
)
```

### Por qué Go es especial

1. **Pragmatismo**: No sigue dogma POO
2. **Claridad**: Código explícito es mejor que implícito
3. **Flexibilidad**: Composición permite recomposición
4. **Performance**: Sin overhead de jerarquías complejas
5. **Simplicidad**: Un programador puede entender todo el código

### El viaje del POOP (Programmer) en Go

**Fase 1**: "¿Dónde está inheritance? ¿Cómo hago abstract classes?"

**Fase 2**: Vuelves a leer el código después de 6 meses. ¡Es fácil de entender!

**Fase 3**: Empiezas a usar composición en otros lenguajes también.

**Fase 4**: Dominas Go y escribes código elegante y mantenible.

### La Verdad sobre Go y OOP

Go **es POO**, pero:
- No heredita obsesivamente
- No abusa de polimorfismo
- Elije composición estratégicamente
- Prioriza legibilidad sobre teoría

**Go OOP es pragmático. Y eso es su fortaleza.**

Lee las interfaces pequeñas, compón estructuras simples, escribe código claro. Así se hace OOP en Go.
