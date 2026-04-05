# Control de Flujo en Go - Guía Completa

El control de flujo es lo que permite que tus programas tomen decisiones. Sin él, tu programa sería una secuencia linear aburrida. Con control de flujo, puedes responder a diferentes situaciones.

---

## Parte I: Sentencia if/else

La sentencia `if` es la forma más básica de tomar una decisión: "si X es verdadero, haz A, si no haz B".

### If simple

```go
edad := 18

if edad >= 18 {
    fmt.Println("Puedes votar")
}
```

**¿Qué pasa aquí?**
1. Evaluamos la condición `edad >= 18`
2. Si es `true`, ejecutamos el bloque de código dentro de `{}`
3. Si es `false`, saltamos ese bloque

### If/else

```go
edad := 16

if edad >= 18 {
    fmt.Println("Puedes votar")
} else {
    fmt.Println("Eres muy joven para votar")
}
```

Ahora tenemos dos caminos: uno si es verdad, otro si es falso.

### If/else if/else

Para múltiples condiciones:

```go
calificacion := 75

if calificacion >= 90 {
    fmt.Println("Excelente")
} else if calificacion >= 80 {
    fmt.Println("Muy bien")
} else if calificacion >= 70 {
    fmt.Println("Bien")
} else if calificacion >= 60 {
    fmt.Println("Satisfecho")
} else {
    fmt.Println("Insuficiente")
}
```

Go evalúa las condiciones de arriba a abajo. Cuando encuentra una verdadera, ejecuta ese bloque y **ignora el resto**.

**Importante**: El último `else` es nuestro "plan de respaldo" si nada más coincide.

### If con inicialización

```go
// Declarar e initializar variable DENTRO del if
if x := 10; x > 5 {
    fmt.Println(x)  // ✅ Puedo usar x aquí
}

// fmt.Println(x)  // ❌ ILLEGAL: x no existe fuera del if
```

Esto es útil cuando la variable solo se necesita en ese if. Mantiene el scope limpio.

**Ejemplo práctico**:
```go
// Validar entrada de usuario
if entrada, err := LeerArchivo("datos.txt"); err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Datos:", entrada)
}
```

### Condiciones complejas

Puedes combinar condiciones con `&&` (y), `||` (o), `!` (no):

```go
edad := 25
tieneLicencia := true
estaSobrio := true

// AND: ambas deben ser verdaderas
if tieneLicencia && estaSobrio {
    fmt.Println("Puedes conducir")
}

// OR: al menos una debe ser verdadera
if edad < 18 || !tieneLicencia {
    fmt.Println("No puedes conducir")
}

// NOT
if !tieneLicencia {
    fmt.Println("Obtén una licencia")
}
```

**Orden de evaluación**:
- Go evalúa de izquierda a derecha
- `&&` detiene si encuentra `false` (cortocircuito)
- `||` detiene si encuentra `true` (cortocircuito)

```go
// Cortocircuito en acción
if x > 0 && dividir(10, x) {  // Si x <= 0, nunca llama dividir()
    // ...
}
```

---

## Parte II: Switch

Cuando tienes **muchas** opciones para una sola variable, `switch` es más limpio que múltiples `if/else if`.

### Switch básico

```go
dia := 3

switch dia {
case 1:
    fmt.Println("Lunes")
case 2:
    fmt.Println("Martes")
case 3:
    fmt.Println("Miércoles")
case 4:
    fmt.Println("Jueves")
case 5:
    fmt.Println("Viernes")
default:
    fmt.Println("Fin de semana")
}
// Output: Miércoles
```

**¿Cómo funciona?**
1. Evalúa `dia`
2. Compara con cada `case`
3. Ejecuta el que coincida
4. Si ninguno coincide, ejecuta `default`

**Importante**: En Go, después de ejecutar un case, el flujo sale del switch (NO hay "fall-through" como en C/Java).

### Fall-through

Si quieres ejecutar múltiples cases, usa `fallthrough`:

```go
tipo := "auto"

switch tipo {
case "auto":
    fmt.Println("Tiene 4 ruedas")
    fallthrough
case "camión":
    fmt.Println("Tiene gasolina")
    fallthrough
case "bicicleta":
    fmt.Println("Es un vehículo")
default:
    fmt.Println("No conozco este tipo")
}
// Output:
// Tiene 4 ruedas
// Tiene gasolina
// Es un vehículo
```

### Switch con múltiples valores en un case

```go
mes := 12

switch mes {
case 12, 1, 2:
    fmt.Println("Invierno")
case 3, 4, 5:
    fmt.Println("Primavera")
case 6, 7, 8:
    fmt.Println("Verano")
case 9, 10, 11:
    fmt.Println("Otoño")
}
```

Un case puede tener múltiples valores separados por comas.

### Switch sin expresión (es como if/else if)

```go
edad := 25
tieneNegocio := true

switch {
case edad < 18:
    fmt.Println("Eres menor")
case edad >= 18 && edad < 65 && tieneNegocio:
    fmt.Println("Puedes iniciar un negocio")
case edad >= 65:
    fmt.Println("Edad de jubilación")
}
```

Cuando no hay expresión después de `switch`, cada `case` puede ser cualquier condición boolean.

### Type switch (avanzado)

Para determinar el tipo de una variable:

```go
func Procesar(valor interface{}) {
    switch v := valor.(type) {
    case string:
        fmt.Println("Es string:", v)
    case int:
        fmt.Println("Es int:", v)
    case []int:
        fmt.Println("Es slice de ints:", v)
    case float64:
        fmt.Println("Es float64:", v)
    default:
        fmt.Println("Tipo desconocido")
    }
}

func main() {
    Procesar("Hola")
    Procesar(42)
    Procesar([]int{1, 2, 3})
    Procesar(3.14)
}
```

---

## Parte III: Operadores Ternarios

**Nota importante**: Go **NO tiene operador ternario** como `condicion ? valor1 : valor2` (que existe en C, Java, JavaScript).

Pero hay **varias formas** de lograr lo mismo:

### Opción 1: If/else como Expresión (Recomendado)

Go permite usar `if/else` inline (aunque requiere inicializador):

```go
// Forma 1: Variable con if/else
edad := 25
categoria := ""

if edad >= 18 {
    categoria = "Adulto"
} else {
    categoria = "Menor"
}

// Mejor: Declarar e inicializar en la misma línea
var estado string
if edad >= 18 {
    estado = "Adulto"
} else {
    estado = "Menor"
}
fmt.Println(estado)  // Adulto
```

### Opción 2: Función Helper Ternaria

Crear una función que simule el comportamiento:

```go
// Función ternaria genérica (Go 1.18+ con Generics)
func Ternary[T any](condicion bool, siVerdad T, siFlaso T) T {
    if condicion {
        return siVerdad
    }
    return siFlaso
}

// Uso
edad := 25
categoria := Ternary(edad >= 18, "Adulto", "Menor")
fmt.Println(categoria)  // Adulto

// Otro ejemplo
puntos := 85
resultado := Ternary(puntos >= 80, "Aprobado", "Reprobado")
fmt.Println(resultado)  // Aprobado

// Con números
precio := 100.0
descuento := Ternary(precio > 50, precio*0.90, precio)
fmt.Println(descuento)  // 90
```

### Opción 3: Map (Para Valores Simples)

```go
// Para valores predefinidos
estado := map[bool]string{
    true:  "Adulto",
    false: "Menor",
}

edad := 25
fmt.Println(estado[edad >= 18])  // Adulto
```

### Opción 4: Switch Corto

```go
edad := 25

categoria := func() string {
    switch {
    case edad >= 18:
        return "Adulto"
    default:
        return "Menor"
    }
}()

fmt.Println(categoria)  // Adulto
```

### Comparación de Formas

```go
edad := 25

// ❌ No existe (otros lenguajes)
// categoria := edad >= 18 ? "Adulto" : "Menor"

// ✅ Opción 1: If/else directo (Simple y clara)
var categoria string
if edad >= 18 {
    categoria = "Adulto"
} else {
    categoria = "Menor"
}

// ✅ Opción 2: Función ternaria (Reutilizable)
categoria := Ternary(edad >= 18, "Adulto", "Menor")

// ✅ Opción 3: Map (Para valores predefinidos)
estados := map[bool]string{true: "Adulto", false: "Menor"}
categoria := estados[edad >= 18]

// ✅ Opción 4: Switch inline (Para lógica compleja)
categoria := func() string {
    switch {
    case edad >= 65:
        return "Jubilado"
    case edad >= 18:
        return "Adulto"
    default:
        return "Menor"
    }
}()
```

### Ejemplo Práctico: Sistema de Categorías

```go
type Producto struct {
    Precio float64
    Stock  int
}

func ObtenerCategoria(p Producto) string {
    // Opción 1: If/else claro
    if p.Precio > 100 {
        return "Premium"
    } else if p.Precio > 50 {
        return "Estándar"
    } else {
        return "Económico"
    }
}

func ObtenerDisponibilidad(p Producto) string {
    // Opción 2: Ternaria con función
    return Ternary(p.Stock > 10, "Disponible", "Pocas unidades")
}

func ObtenerDescuento(p Producto) float64 {
    // Opción 3: Ternaria con números
    return Ternary(p.Precio > 100, p.Precio*0.90, p.Precio)
}

func main() {
    productos := []Producto{
        {150, 20},
        {75, 5},
        {25, 0},
    }
    
    for i, p := range productos {
        categoria := ObtenerCategoria(p)
        disponibilidad := ObtenerDisponibilidad(p)
        descuento := ObtenerDescuento(p)
        
        fmt.Printf("Producto %d: %s - %s - Precio: $%.2f (desc: $%.2f)\n",
            i+1, categoria, disponibilidad, p.Precio, descuento)
    }
}

// Output:
// Producto 1: Premium - Disponible - Precio: $150.00 (desc: $135.00)
// Producto 2: Estándar - Pocas unidades - Precio: $75.00 (desc: $75.00)
// Producto 3: Económico - Pocas unidades - Precio: $25.00 (desc: $25.00)
```

### Ternarios Anidados (¡Evitar!)

```go
// ❌ Complicado y difícil de leer
val := Ternary(x > 10, 
    Ternary(y > 5, "Muy Alto", "Alto"), 
    Ternary(y > 5, "Medio", "Bajo"))

// ✅ Mejor: usar if/else if/else
var val string
if x > 10 {
    if y > 5 {
        val = "Muy Alto"
    } else {
        val = "Alto"
    }
} else {
    if y > 5 {
        val = "Medio"
    } else {
        val = "Bajo"
    }
}

// ✅ O mejor aún: switch
switch {
case x > 10 && y > 5:
    val = "Muy Alto"
case x > 10:
    val = "Alto"
case y > 5:
    val = "Medio"
default:
    val = "Bajo"
}
```

---

## Parte IV: Goto (⚠️ Evitar en general)

Go tiene `goto`, pero generalmente **NO deberías usarlo**. Hace el código confuso.

```go
i := 0

start:
    if i >= 5 {
        goto fin
    }
    fmt.Println(i)
    i++
    goto start

fin:
    fmt.Println("Terminado")
```

**Cuándo usar goto**: Casi nunca. Go tiene mejores formas (loops, funciones).

---

## Parte V: Ejemplos Integrados

### Ejemplo 1: Calculadora Interactiva

```go
import "fmt"

func CalcularPromedio(notas []float64) float64 {
    if len(notas) == 0 {
        return 0
    }
    suma := 0.0
    for _, nota := range notas {
        suma += nota
    }
    return suma / float64(len(notas))
}

func ObtenerGrado(promedio float64) string {
    if promedio >= 90 {
        return "A"
    } else if promedio >= 80 {
        return "B"
    } else if promedio >= 70 {
        return "C"
    } else if promedio >= 60 {
        return "D"
    } else {
        return "F"
    }
}

func main() {
    notas := []float64{85, 92, 78, 88, 95}
    promedio := CalcularPromedio(notas)
    grado := ObtenerGrado(promedio)
    
    fmt.Printf("Promedio: %.2f\n", promedio)
    fmt.Printf("Grado: %s\n", grado)
}
```

### Ejemplo 2: Validador de Datos

```go
type Usuario struct {
    Nombre   string
    Email    string
    Edad     int
}

// Función ternaria auxiliar
func Ternary[T any](condicion bool, siVerdad T, siFlaso T) T {
    if condicion {
        return siVerdad
    }
    return siFlaso
}

func ValidarUsuario(u Usuario) (bool, string) {
    // Validar nombre
    if len(u.Nombre) == 0 {
        return false, "Nombre no puede estar vacío"
    }
    
    if len(u.Nombre) < 3 {
        return false, "Nombre muy corto (mínimo 3 caracteres)"
    }
    
    // Validar email
    if !strings.Contains(u.Email, "@") {
        return false, "Email inválido (falta @)"
    }
    
    // Validar edad
    if u.Edad < 0 {
        return false, "Edad no puede ser negativa"
    }
    
    if u.Edad < 13 {
        return false, "Debe ser mayor de 13 años"
    }
    
    if u.Edad > 150 {
        return false, "Edad no realista"
    }
    
    // Todo OK
    return true, "Usuario válido"
}

func ObtenerCategoria(u Usuario) string {
    // Usando ternaria
    return Ternary(
        u.Edad >= 18,
        Ternary(u.Edad >= 65, "Adulto Mayor", "Adulto"),
        "Menor",
    )
}

func main() {
    usuarios := []Usuario{
        {"Alice", "alice@example.com", 25},
        {"Bo", "bob@example.com", 30},
        {"Charlie", "charliemail.com", 28},
        {"Diana", "diana@example.com", 70},
    }
    
    for _, u := range usuarios {
        if valido, mensaje := ValidarUsuario(u); valido {
            categoria := ObtenerCategoria(u)
            fmt.Printf("✅ %s (%s): %s\n", u.Nombre, categoria, mensaje)
        } else {
            fmt.Printf("❌ %s: %s\n", u.Nombre, mensaje)
        }
    }
}
```

### Ejemplo 3: Procesador de Pedidos

```go
type Pedido struct {
    ID       int
    Total    float64
    Estado   string
}

func ProcesarPedido(p Pedido) string {
    // Switch en estado
    switch p.Estado {
    case "pendiente":
        return "Procesando pedido..."
    case "procesando":
        if p.Total > 1000 {
            return "Requiere aprobación manual"
        }
        return "Avanzando a envío..."
    case "enviado":
        return "En tránsito"
    case "entregado":
        return "Pedido completado"
    case "cancelado":
        return "Pedido cancelado"
    default:
        return "Estado desconocido"
    }
}

func ObtenerColor(estado string) string {
    switch estado {
    case "pendiente", "procesando":
        return "🟡 Yellow"
    case "enviado":
        return "🔵 Blue"
    case "entregado":
        return "🟢 Green"
    case "cancelado":
        return "🔴 Red"
    default:
        return "⚪ Grey"
    }
}

func main() {
    pedidos := []Pedido{
        {1, 500, "pendiente"},
        {2, 1500, "procesando"},
        {3, 750, "enviado"},
        {4, 200, "entregado"},
    }
    
    for _, p := range pedidos {
        accion := ProcesarPedido(p)
        color := ObtenerColor(p.Estado)
        fmt.Printf("Pedido #%d: %s %s (Total: $%.2f)\n", p.ID, color, accion, p.Total)
    }
}
```

---

## Mejores Prácticas

### ✅ Bien

```go
// 1. If para una o dos condiciones
if x > 0 {
    hacer_algo()
}

// 2. Switch para muchos valores
switch dia {
case 1: fmt.Println("Lunes")
case 2: fmt.Println("Martes")
// ...
}

// 3. Validar temprano (early return)
if !valido {
    return err
}
// Resto del código con supuesto válido

// 4. Condiciones legibles
if edad >= 18 && tieneLicencia {
    permitir_conducir()
}

// 5. Evitar nesting profundo
// En lugar de múltiples if anidados, usa early returns
```

### ❌ Mal

```go
// 1. If/else if/else excesivo
if x == 1 {
} else if x == 2 {
} else if x == 3 {
} else if x == 4 {
}  // ❌ Usar switch

// 2. Condiciones complejas sin variables
if user.age >= 18 && user.hasLicense && user.paymentOk && notBlacklisted {
}  // ❌ Difícil de leer

// 3. Nesting profundo
if x > 0 {
    if y > 0 {
        if z > 0 {
            // ... 3 niveles 😫
        }
    }
}  // ❌ Usar early returns

// 4. Olvidar que Go termina switch tras case
switch x {
case 1:
    a()
    // ❌ Pensaste que también ejecutaría case 2
case 2:
    b()
}
```

---

## Resumen

| Estructura | Uso |
|-----------|-----|
| `if` | Una condición simple |
| `if/else` | Dos caminos |
| `if/else if/else` | Múltiples caminos (máx ~3-4) |
| `switch` | Muchos valores para una variable |
| `switch sin expr` | Múltiples condiciones complejas |

---

## Conclusión

El control de flujo es el corazón de la lógica de programación:
- **if**: Lo más común para decisiones simples
- **switch**: Cuando tienes muchas opciones
- **Mantén simple**: Código anidado profundo es difícil de leer
- **Early returns**: Mejor que if/else por todos lados

Domina if/switch y podrás escribir lógica clara y mantenible.
