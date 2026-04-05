# Variables, Constantes y Tipos en Go - Guía Completa

> **De qué forma almacenar dados en memoria y transformarlos según tus necesidades**

---

## ¿Por qué existen las variables?

Imagina que escribes un programa que pide el nombre de un usuario:

```
¿Cuál es tu nombre?
> Juan

Hola Juan, bienvenido!
```

El nombre **"Juan" debe almacenarse en algún lugar** de la memoria para que luego puedas imprimirlo en el mensaje. Las **variables son exactamente eso: espacios en memoria donde guardas datos**.

Sin variables, cada vez que necesitarias el nombre tendrías que pedirlo de nuevo. Las variables son **la forma de Go de recordar información**.

---

## Parte I: Declaración de Variables - Lo Esencial

### Paso 1: Entender qué es "declarar" una variable

**Declarar = Crear una caja en memoria y decirle a Go qué tipo de cosas va a tener.**

```go
var nombre string = "Juan"
```

Esto significa:
- `var` = "Hey Go, crea una variable"
- `nombre` = El nombre de la caja (cómo la llamarás adelante)
- `string` = Tipo de cosa que va a guardar (solo texto)
- `"Juan"` = El valor que pones adentro al principio

**Es como si dijeras:** "Crea una caja llamada 'nombre' que solo puede guardar texto, y mete 'Juan' adentro"

---

### Forma 1: `var` con tipo explícito (La forma larga)

```go
var nombre string = "Juan"
var edad int = 30
var saldo float64 = 1500.50
```

Esta es **la más clara y explícita** porque dices exactamente qué tipo es.

**¿Cuándo la usas?** Cuando quieres que sea super claro, especialmente en código que otros van a leer.

---

### Forma 2: `var` dejando que Go adivine el tipo (Forma media)

```go
var nombre = "Juan"     // Go ve "Juan" y sabe que es string
var edad = 30           // Go ve 30 y sabe que es int  
var saldo = 1500.50     // Go ve 1500.50 y sabe que es float64
```

**¿Cómo lo adivina?** Go mira lo que hay del lado derecho del `=`:
- Si ves `"texto"` entre comillas → es `string`
- Si ves `123` sin punto decimal → es `int`
- Si ves `123.45` con punto → es `float64`
- Si ves `true` o `false` → es `bool`

**¿Cuándo la usas?** Es menos verboso que la forma 1, pero el tipo es obvio.

---

### Forma 3: `:=` - La forma CORTA (Esto es lo que usarás en 90% del código)

```go
nombre := "Juan"
edad := 30
saldo := 1500.50
```

**¿Qué significa `:=`?**
- Es una forma **abreviada** de declarar + asignar al mismo tiempo
- Go automáticamente **adivina el tipo** mirando el lado derecho
- Es más rápido de escribir: 1 paso en vez de 3

**Compara:**

```go
// Forma larga (3 pasos)
var nombre string = "Juan"

// Forma media (Go adivina, 2 pasos)
var nombre = "Juan"

// Forma corta (Go adivina, 1 paso) ← Lo que usarás siempre
nombre := "Juan"
```

---

### ⚠️ LA RESTRICCIÓN IMPORTANTE: `:=` SOLO EN FUNCIONES

**❌ Esto NO funciona:**
```go
package main

nombre := "Juan"   // ERROR❌ (fuera de función)

func main() {
    fmt.Println(nombre)
}
```

**✅ Esto SÍ funciona:**
```go
package main

func main() {
    nombre := "Juan"   // OK✅ (dentro de función)
    fmt.Println(nombre)
}
```

**¿POR QUÉ?** Go tiene dos "niveles":
1. **Nivel de paquete** (código que no está dentro de ninguna función) - Usa `var`
2. **Dentro de función** (código dentro de `func main()` o cualquier otra función) - Usa `:=`

**Piénsalo así:**
- `:=` es la forma "rápida" para código que solo existe **dentro de funciones**
- `var` es la forma "formal" para código de **nivel superior**

**Tabla para decidir:**

| Código donde está | Usar | Ejemplo |
|---|---|---|
| Dentro de `func main()` | `:=` | `nombre := "Juan"` |
| Dentro de cualquier función | `:=` | `edad := 30` |
| Fuera de funciones (nivel paquete) | `var` | `var nombre = "Juan"` |

---

### Forma 4: `:=` con múltiples variables (Práctico)

```go
// Declarar varias al mismo tiempo
nombre, edad := "Juan", 30

// También puedes hacerlo en líneas separadas
a := 1
b := 2
c := 3
```

---

### Forma 5: `var` mostrando variedad (Nivel paquete)

Cuando estás **fuera de funciones** y quieres declarar múltiples variables:

```go
var (
    nombre string = "Juan"
    edad int = 30
    ciudad string = "Madrid"
)
```

---

### REASIGNAR vs REDECLARAR (Esto confunde a muchos)

**REASIGNAR = Dar un nuevo valor a una variable que ya existe**

```go
nombre := "Juan"
nombre = "Carlos"   // ✅ Reasignar (usa solo =, no :=)
```

Esto está OK. Ya creaste la caja, ahora cambias lo que hay adentro.

---

**REDECLARAR = Intentar crear la misma variable DOS VECES**

```go
nombre := "Juan"
nombre := "Carlos"  // ❌ ERROR (nombre ya existe)
```

Esto es un ERROR. No puedes crear la misma caja dos veces.

---

### RESUMEN RÁPIDO DE LAS 3 FORMAS

```go
// FORMA 1: Larga y explícita (nivel paquete)
var nombre string = "Juan"

// FORMA 2: Corta con var (nivel paquete o dentro función)
var nombre = "Juan"

// FORMA 3: Ultracorta con := (SOLO dentro de función)
nombre := "Juan"
```

**Regla de oro:**
- En 90% de tu código dentro de funciones → usa `:=`
- Cuando necesites variables de nivel paquete → usa `var`
- Sé explícito si es confuso → agrega el tipo

---

## Parte II: Nombres de Variables - Reglas Simples

Go no es "libre" con los nombres. Hay convenciones que todos siguen.

---

### Regla 1: Usa `camelCase` (minúscula al inicio)

**camelCase** significa: primera palabra minúscula, siguientes palabras con mayúscula. Sin guiones bajos.

```go
// ✅ BIEN (camelCase)
nombreUsuario := "Juan"
edadMinima := 18
precioTotal := 99.99

// ❌ MAL (snake_case, no idiomático en Go)
nombre_usuario := "Juan"
edad_minima := 18
precio_total := 99.99
```

**¿Por qué?** Go fue diseñado así. Es lo que ves en todo código Go profesional.

---

### Regla 2: Mayúscula inicial = "Público" (puede usarse en otros archivos)

Si una variable empieza con **MAYÚSCULA**, es como si dijeras: "Esta variable es pública, otros paquetes pueden usarla"

```go
// DENTRO DE UN ARCHIVO Go

var NombreAplicacion = "MiApp"    // Mayúscula = otros paquetes pueden usarla
var versionInterna = "1.0.0"      // minúscula = solo este paquete
```

(No te preocupes por esto ahorita, es más avanzado. Por ahora, usa minúsculas.)

---

### Regla 3: Nombres cortos en loops, largos en variables que "viven"

```go
// ✅ OK: i en un loop es temporal, 1-2 líneas
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// ❌ MAL: x para algo que vivirá más
x := usuario.ObtenerEdad()    // ¿Qué es x?

// ✅ BIEN: nombre descriptivo
edadUsuario := usuario.ObtenerEdad()
```

**Regla práctica:**
- Variables en loops: OK nombres cortos (`i`, `j`, `k`)
- Variables que usarás después: nombres descriptivos

---

### Regla 4: Nombres que expliquen QUÉ es el dato

```go
// ❌ MALO (no dice qué es)
x := 100
p := usuario
t := 5

// ✅ BIEN (claro qué es)
maxProductosPorPagina := 100
propietarioProducto := usuario
tiempoEsperaSegundos := 5
```

---

## Parte III: Tipos Básicos (Los que SIEMPRE usarás)

Go tiene muchos tipos, pero debes entender primero los **básicos**.

---

### Tipo 1: `string` - Texto

```go
nombre := "Juan"
ciudad := "Madrid"
mensaje := "Hola, mundo"
```

**Eso es todo.** Un string es texto entre comillas dobles.

```go
// Para texto de múltiples líneas, usa backticks
poema := `Esto es
texto en
varias líneas`
```

---

### Tipo 2: `int` - Número entero

```go
edad := 30
cantidad := 100
saldo := -50
```

`int` significa "entero". Sin decimales. Solo Go decide si 32 o 64 bits (no te preocupes).

**Variantes que existen pero casi nunca usas:**
```go
var a int8 = 100       // De -128 a 127 (muy pequeño)
var b int16 = 1000     // Un poco mayor
var c int32 = 100000   // Más grande
var d int64 = 1000000  // Muy grande (si realmente lo necesitas)
```

**Regla:** Usa `int`. Es suficiente para 99% de casos.

---

### Tipo 3: `float64` - Número con decimales

```go
precio := 19.99
saldo := 1500.50
pi := 3.14159
```

`float64` es para números con punto decimal. Es el estándar en Go.

**También existe `float32`** (menos precisión) **pero usa `float64` por defecto.**

---

### Tipo 4: `bool` - Verdadero o Falso

```go
esActivo := true
esMayor := false
esVacio := (nombre == "")
```

`bool` tiene solo dos valores: `true` o `false`. Se usa en condiciones.

---

### Tipo 5: `byte` y `rune` - Un carácter (Raro, ignora por ahora)

```go
// byte es un número (0-255)
// rune es un carácter Unicode

letra := 'A'    // rune
numero := byte(65)   // byte (65 = 'A')
```

**Puedes ignorar esto por ahora.** Es avanzado.

---

### RESUMEN DE TIPOS BÁSICOS

| Tipo | Qué es | Ejemplo |
|------|--------|---------|
| `string` | Texto | `"Juan"` |
| `int` | Número entero | `30` |
| `float64` | Número con decimales | `3.14` |
| `bool` | Verdadero/Falso | `true` |

**Eso es lo que necesitas saber para empezar.**

---

## Parte IV: Zero Values (Valores por Defecto) - ¿Qué pasa si no inicializo?

**Si creas una variable pero NO le das un valor inicial, Go automáticamente le da un valor "por defecto".**

---

### Ejemplos

```go
var numero int              // ¿Cuál es el valor? 
fmt.Println(numero)         // Imprime: 0

var precio float64          // ¿Cuál es el valor?
fmt.Println(precio)         // Imprime: 0.0

var nombre string           // ¿Cuál es el valor?
fmt.Println(nombre)         // Imprime: (nada, strings vacío)

var activo bool             // ¿Cuál es el valor?
fmt.Println(activo)         // Imprime: false
```

---

### Tabla de Zero Values

| Tipo | Valor por defecto |
|------|---|
| `int`, `int8`, `int16`, etc. | `0` |
| `float32`, `float64` | `0.0` |
| `string` | `""` (texto vacío) |
| `bool` | `false` |
| `slice`, `map` | `nil` |
| `pointer` | `nil` |

---

### ¿POR QUÉ IMPORTA?

**Porque no es "undefined" o "basura"**, como en otros lenguajes.

En Go, **siempre sabes qué valor tiene una variable sin inicializar:** siempre es cero para su tipo.

```go
// ✅ Seguro: sabes que contador vale 0
var contador int
contador++  // Ahora es 1

// ❌ Contadora no inicializada sería basura en otros lenguajes
// Pero en Go, es 0 siempre
```

---

### REGLA PRÁCTICA: Inicializa siempre de todos modos

Aunque Go te da un valor por defecto, es **mejor inicializar explícitamente** para que sea claro:

```go
// ❌ Confuso: ¿fue inicializado a propósito o por defecto?
var contador int

// ✅ Claro: quiero empezar en 0
contador := 0
```

---

## Parte V: Constantes - Valores que NO Cambian

Una **constante es como una variable, pero que NO PUEDES CAMBIAR después de crearla.**

Imagina que tienes una constante llamada `Pi`:

```go
const Pi = 3.14159

// Más tarde en el código
Pi = 3.14160         // ❌ ERROR: No puedes cambiar una constante
```

---

### Cuándo usar const vs var

```go
// ❌ MAL: Valor fijo que nunca cambia, pero usa var
var MaxProductosPorPagina = 10

// ✅ BIEN: Valor fijo → const
const MaxProductosPorPagina = 10

// ✅ var para valores que SÍ cambian
var productoActualMostrado = 0
productoActualMostrado = 1  // Puede cambiar
```

**Regla simple:**
- ¿El valor **nunca cambia**? → `const`
- ¿El valor **puede cambiar**? → `var`

---

### Ejemplos de constantes

```go
const (
    NombreApp = "MiApp"
    Version = "1.0.0"
    AutorContacto = "soporte@miapp.com"
)

const (
    Lunes = 1
    Martes = 2
    Miercoles = 3
)

const PiMath = 3.14159265358979
```

---

### iota - Contador automático (Avanzado, puedes ignorar)

Para crear series de números automáticamente:

```go
const (
    Rojo = iota      // 0
    Verde            // 1
    Azul             // 2
    Amarillo         // 3
)

fmt.Println(Rojo)    // Imprime: 0
fmt.Println(Verde)   // Imprime: 1
```

`iota` es "perezoso" - Go incremente automáticamente.

**Puedes ignorar `iota` por ahora. Es un patrón avanzado.**

---

### RESUMEN VAR vs CONST

```go
const MaxUsuarios = 100        // No puede cambiar
var usuariosActuales = 0       // Puede cambiar (empieza en 0)
usuariosActuales++             // Ahora es 1
```

---

## Parte VI: Conversión de Tipos - Cambiar de un tipo a otro

**Go NO hace conversiones automáticas** entre tipos. Si necesitas cambiar de tipo, debes hacerlo explícitamente.

---

### ❌ Esto NO funciona

```go
var numero int = 5
var decimal float64 = numero    // ❌ ERROR: int no es float64 automáticamente
```

Go te dirá: "Hey, son tipos diferentes. Conviértelos explícitamente."

---

### ✅ Esto SÍ funciona: Conversión explícita

```go
var numero int = 5
var decimal float64 = float64(numero)   // ✅ Ahora sí
fmt.Println(decimal)                    // Imprime: 5.0
```

**Sintaxis:**
```go
nuevoTipo(variable)
```

Es simple: pones el tipo destino entre paréntesis, seguido de la variable.

---

### Ejemplos prácticos

#### Entero a Float

```go
edad := 25
edadDecimal := float64(edad)    // 25 → 25.0
fmt.Println(edadDecimal)        // 25.0
```

#### Float a Entero (⚠️ Pierde decimales)

```go
precio := 19.99
precioEntero := int(precio)
fmt.Println(precioEntero)       // 19 (perdió .99)
```

**⚠️ CUIDADO:** Cuando conviertes float a int, **desaparece la parte decimal.**

---

#### Número a String

```go
numero := 42
texto := strconv.Itoa(numero)   // int → string
fmt.Println(texto)              // "42"

// También puedes usar Sprintf
texto2 := fmt.Sprintf("%d", numero)
fmt.Println(texto2)             // "42"
```

---

#### String a Número (⚠️ Puede fallar)

```go
// String a entero
texto := "123"
numero, err := strconv.Atoi(texto)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println(numero)         // 123
}

// ¿Qué pasa si el string no es número?
texto2 := "abc"
numero2, err := strconv.Atoi(texto2)
if err != nil {
    fmt.Println("Error: no es número")  // Llegará aquí
}
```

**⚠️ IMPORTANTE:** Convertir string a número puede fallar, por eso necesitas `err`.

---

#### String a Float

```go
texto := "3.14"
pi, err := strconv.ParseFloat(texto, 64)
if err != nil {
    fmt.Println("Error")
} else {
    fmt.Println(pi)             // 3.14
}
```

---

### RESUMEN DE CONVERSIONES

| Conversión | Código |
|---|---|
| int → float64 | `float64(numero)` |
| float64 → int | `int(numero)` (pierde decimales) |
| int → string | `strconv.Itoa(numero)` |
| string → int | `strconv.Atoi(texto)` (puede fallar) |
| string → float64 | `strconv.ParseFloat(texto, 64)` (puede fallar) |
| float64 → string | `fmt.Sprintf("%.2f", numero)` |

---

## Parte VII: Operadores - El = y sus amigos

### El operador `=` (Asignación simple)

```go
x := 10
x = 20              // Cambia el valor de x a 20
```

Es lo que ya vimos: cambiar lo que está en la caja.

---

### Operadores "perezosos" (Suma, resta, etc. todo junto)

En lugar de escribir `x = x + 5`, puedes escribir `x += 5`:

```go
x := 10
x += 5              // x = x + 5  →  x ahora es 15
fmt.Println(x)      // 15

x -= 3              // x = x - 3  →  x ahora es 12
x *= 2              // x = x * 2  →  x ahora es 24
x /= 4              // x = x / 4  →  x ahora es 6
x %= 2              // x = x % 2  →  x ahora es 0 (resto de dividir)
```

**Tabla:**

| Operador | Significado | Ejemplo |
|---|---|---|
| `+=` | Suma | `x += 5` es `x = x + 5` |
| `-=` | Resta | `x -= 3` es `x = x - 3` |
| `*=` | Multiplica | `x *= 2` es `x = x * 2` |
| `/=` | Divide | `x /= 4` es `x = x / 4` |
| `%=` | Resto | `x %= 2` es `x = x % 2` |

**Esto es solo "azúcar sintáctico"** (forma más rápida de escribir lo mismo).

---

### Incremento y Decremento

```go
x := 5
x++                 // x = x + 1  →  x ahora es 6
fmt.Println(x)      // 6

x--                 // x = x - 1  →  x ahora es 5
fmt.Println(x)      // 5
```

**⚠️ SOLO NOTACIÓN POSTFIJA EN GO:**
```go
x++     // ✅ Bien
++x     // ❌ ERROR (Go no permite esto)
```

---

### Intercambiar valores (Múltiple asignación)

```go
a := 1
b := 2

// Intercambiar: a y b cambian de lugar
a, b = b, a

fmt.Println(a)      // 2
fmt.Println(b)      // 1
```

Es MUY fácil en Go intercambiar valores.

---

## Parte VIII: Ejemplo Práctico - Calculadora de Presupuesto

Imagina que quieres **calcular el total de una compra** con descuento e impuestos.

```go
package main

import "fmt"

func main() {
    // Precios de productos
    precioLaptop := 999.99
    precicoMouse := 25.50
    precioTeclado := 75.00
    
    // Cantidades
    cantidadLaptop := 1
    cantidadMouse := 2
    cantidadTeclado := 1
    
    // Calcular subtotal
    subtotal := 0.0
    subtotal += precioLaptop * float64(cantidadLaptop)
    subtotal += precioMouse * float64(cantidadMouse)
    subtotal += precioTeclado * float64(cantidadTeclado)
    
    fmt.Printf("Subtotal: $%.2f\n", subtotal)
    
    // Aplicar descuento del 10%
    descuentoPorcentaje := 10.0
    descuento := subtotal * descuentoPorcentaje / 100
    subTotalConDescuento := subtotal - descuento
    
    fmt.Printf("Descuento (10%%): -$%.2f\n", descuento)
    fmt.Printf("Después descuento: $%.2f\n", subTotalConDescuento)
    
    // Aplicar impuesto del 16%
    tasaImpuesto := 0.16
    impuesto := subTotalConDescuento * tasaImpuesto
    totalFinal := subTotalConDescuento + impuesto
    
    fmt.Printf("Impuesto (16%%): +$%.2f\n", impuesto)
    fmt.Printf("TOTAL A PAGAR: $%.2f\n", totalFinal)
}
```

**Output:**
```
Subtotal: $1200.49
Descuento (10%): -$120.05
Después descuento: $1080.44
Impuesto (16%): +$172.87
TOTAL A PAGAR: $1253.31
```

---

### ¿Qué pasó acá?

1. **Creamos variables** para cada precio y cantidad con `:=`
2. **Calculamos el subtotal** usando `+=` para ir sumando
3. **Aplicamos descuento**: restamos el 10% del subtotal
4. **Aplicamos impuesto**: sumamos el 16% al resultado
5. **Imprimimos** el total con formato `%.2f` (2 decimales)

---

## Parte IX: Errores Comunes - Qué NO hacer

### ❌ Error 1: Mezclar tipos sin conversión

```go
// MALO
numero := 5
resultado := numero + 2.5    // ❌ Error: int + float64

// BIEN
resultado := float64(numero) + 2.5   // ✅ Convertir primero
```

---

### ❌ Error 2: Redeclarar una variable

```go
// MALO
nombre := "Juan"
nombre := "Carlos"           // ❌ ERROR: nombre ya existe

// BIEN
nombre := "Juan"
nombre = "Carlos"            // ✅ Reasignar, no redeclarar
```

**Recuerda:** `:=` es para CREAR. `=` es para CAMBIAR.

---

### ❌ Error 3: Usar `:=` fuera de función

```go
// MALO
nombre := "Juan"             // ❌ ERROR (fuera de función)

func main() {
    fmt.Println(nombre)
}

// BIEN
var nombre = "Juan"          // ✅ Usar var a nivel paquete

func main() {
    fmt.Println(nombre)
}
```

---

### ❌ Error 4: Convertir string a número sin verificar errores

```go
// MALO (ingenuo)
cantidad := "abc"
numero, _ := strconv.Atoi(cantidad)  // Ignora el error
fmt.Println(numero)                   // Imprime: 0 (confuso)

// BIEN
cantidad := "abc"
numero, err := strconv.Atoi(cantidad)
if err != nil {
    fmt.Println("No es un número válido:", err)
} else {
    fmt.Println("El número es:", numero)
}
```

---

### ❌ Error 5: No usar constantes para valores fijos

```go
// MALO (número mágico)
if edad > 18 {               // ¿De dónde viene 18?
    fmt.Println("Eres mayor")
}

// BIEN
const EdadMayoria = 18
if edad > EdadMayoria {
    fmt.Println("Eres mayor")
}
```

---

## Parte X: Mejores Prácticas - Resumen

### 1. Usa `:=` dentro de funciones, `var` afuera

```go
// ✅ Dentro de función
func main() {
    nombre := "Juan"
}

// ✅ Fuera de función (nivel paquete)
var nombre = "Juan"
```

---

### 2. Nombres claros siempre

```go
// ❌ Confuso
n := 10
p := 99.99

// ✅ Claro
maxProductosPorPagina := 10
precioProducto := 99.99
```

---

### 3. Usa `const` para valores que no cambian

```go
// ❌
var MaxConexiones = 100

// ✅
const MaxConexiones = 100
```

---

### 4. Convierte tipos explícitamente

```go
// ❌
precio := 19
total := precio + 0.50

// ✅
total := float64(precio) + 0.50
```

---

### 5. Maneja errores en conversiones string → número

```go
// ✅
numero, err := strconv.Atoi("123")
if err != nil {
    fmt.Println("Error:", err)
}
```

---

### 6. Inicializa variables siempre

```go
// ❌ Ambiguo
var contador int

// ✅ Claro
contador := 0
```

---

## Resumen Final: ¿Qué debes recordar?

**Las 3 formas de crear variables:**

1. **Dentro de función → `:=`** (la forma rápida)
   ```go
   nombre := "Juan"
   ```

2. **Fuera de función → `var`** (la forma formal)
   ```go
   var nombre = "Juan"
   ```

3. **Valor que nunca cambia → `const`** (la forma permanente)
   ```go
   const MaxUsuarios = 100
   ```

**Los 4 tipos básicos:**
- `string` → Texto
- `int` → Número entero
- `float64` → Número con decimales
- `bool` → Verdadero/Falso

**Conversión de tipos:** Siempre explícita

```go
float64(numero)           // int → float64
strconv.Atoi(texto)       // string → int
fmt.Sprintf("%d", numero) // int → string
```

---

**Con esto, ya sabes lo suficiente para escribir Go.**

El resto son detalles que aprenderás mientras escribes código.
