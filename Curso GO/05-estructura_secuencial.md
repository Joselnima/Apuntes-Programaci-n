# Estructura Secuencial en Go - Guía Completa

> **Cómo Go ejecuta código línea por línea y por qué esto es el fundamento de la programación**

---

## ¿Qué es Estructura Secuencial?

**Estructura secuencial** significa que tu programa ejecuta instrucciones **una después de otra, en orden**. Es el tipo más básico de estructura de control.

Imagina que quieres hacer un pastel:
1. Primero compras harina ← se ejecuta primero
2. Luego mezclas ingredientes ← se ejecuta segundo
3. Luego viertes en el molde ← se ejecuta tercero
4. Finalmente horneas ← se ejecuta al final

No puedes hornear antes de mezclar, ¿cierto? Go funciona exactamente así: línea por línea, en orden.

### La Magia del Orden

```go
package main

import "fmt"

func main() {
    x := 5      // Paso 1: Crear variable x con valor 5
    y := 10     // Paso 2: Crear variable y con valor 10
    suma := x + y   // Paso 3: Sumar x y y
    fmt.Println(suma)   // Paso 4: Imprimir el resultado (15)
}
```

**¿Qué pasa si cambias el orden?**

```go
// ❌ ESTO FALLA
fmt.Println(suma)   // Error: suma no existe aún
suma := x + y
x := 5
```

Go dice "Error: suma no definida". No puede usar una variable antes de crearla. Este es el **orden secuencial en acción**.

---

## Parte I: Variables y Asignaciones

### Declaración Básica

```go
var nombre string = "Juan"
var edad int = 30
var saldo float64 = 1500.50
```

**¿Qué pasó?**
1. Línea 1: Se crea variable `nombre` con valor "Juan"
2. Línea 2: Se crea variable `edad` con valor 30
3. Línea 3: Se crea variable `saldo` con valor 1500.50

### La Forma Corta (Más Común)

```go
nombre := "Juan"    // String inferido
edad := 30          // Int inferido
saldo := 1500.50    // Float64 inferido
```

Go **infiere el tipo** automáticamente. Es más rápido de escribir.

### Múltiples Variables

```go
// Forma 1: Una por una
x := 1
y := 2
z := 3

// Forma 2: Juntas
x, y, z := 1, 2, 3

// Forma 3: Bloque (para var)
var (
    nombre string = "Ana"
    edad int = 25
    ciudad string = "Madrid"
)
```

**¿Cuándo usar cada una?**
- **Una por una**: Cuando asignas en momentos diferentes
- **Juntas**: Cuando asignas al mismo tiempo (más legible)
- **Bloque**: Cuando tienes muchas variables de tipos diferentes

---

## Parte II: Operaciones Secuenciales

### Operaciones Aritméticas

```go
package main

import "fmt"

func main() {
    // Paso 1: Crear números
    a := 20
    b := 5

    // Paso 2: Operaciones (en orden)
    suma := a + b       // 25
    resta := a - b      // 15
    multiplicacion := a * b  // 100
    division := a / b   // 4
    modulo := a % b     // 0

    // Paso 3: Imprimir resultados
    fmt.Println("Suma:", suma)
    fmt.Println("Resta:", resta)
    fmt.Println("Multiplicación:", multiplicacion)
    fmt.Println("División:", division)
    fmt.Println("Módulo:", modulo)
}
```

**Salida:**
```
Suma: 25
Resta: 15
Multiplicación: 100
División: 4
Módulo: 0
```

**¿Qué pasó?**
1. Creamos `a` y `b`
2. Realizamos operaciones con `a` y `b`
3. Almacenamos resultados
4. Imprimimos cada uno

Si intentas imprimir antes de calcular:
```go
fmt.Println(suma)  // ❌ Error: suma no existe
suma := a + b      // ← definida aquí
```

### Operaciones con Strings

```go
package main

import "fmt"

func main() {
    // Paso 1: Crear strings
    nombre := "Juan"
    apellido := "Pérez"

    // Paso 2: Concatenación (con +)
    nombreCompleto := nombre + " " + apellido

    // Paso 3: Conversión de tipos
    edad := 30
    mensaje := "Mi nombre es " + nombreCompleto + " y tengo " + fmt.Sprint(edad) + " años"

    // Paso 4: Imprimir
    fmt.Println(mensaje)
}
```

**Salida:**
```
Mi nombre es Juan Pérez y tengo 30 años
```

**¿Por qué usamos `fmt.Sprint()`?**

En Go, **no puedes concatenar directamente strings con números**:
```go
// ❌ ESTO FALLA
resultado := "Edad: " + 30

// ✅ ESTO FUNCIONA
resultado := "Edad: " + fmt.Sprint(30)
```

Go es **type-safe**: respeta los tipos. Debes convertir números a strings.

---

## Parte III: Actualizaciones Secuenciales

### Modificar Variables

Una variable **puede cambiar su valor**, pero el tipo permanece:

```go
package main

import "fmt"

func main() {
    // Paso 1: Crear variable
    contador := 0
    fmt.Println("Inicial:", contador)  // 0

    // Paso 2: Modificar (actualización)
    contador = contador + 1
    fmt.Println("Después de +1:", contador)  // 1

    // Paso 3: Modificar otra vez
    contador = contador + 5
    fmt.Println("Después de +5:", contador)  // 6

    // Paso 4: Operador compuesto (más corto)
    contador += 1  // Equivalente a: contador = contador + 1
    fmt.Println("Después de +=1:", contador)  // 7
}
```

**Salida:**
```
Inicial: 0
Después de +1: 1
Después de +5: 6
Después de +=1: 7
```

### Operadores de Actualización

| Operador | Significado | Ejemplo |
|----------|-------------|---------|
| `+=` | Suma y asigna | `x += 5` → `x = x + 5` |
| `-=` | Resta y asigna | `x -= 3` → `x = x - 3` |
| `*=` | Multiplica y asigna | `x *= 2` → `x = x * 2` |
| `/=` | Divide y asigna | `x /= 4` → `x = x / 4` |
| `%=` | Módulo y asigna | `x %= 2` → `x = x % 2` |
| `++` | Incrementa en 1 | `x++` → `x = x + 1` |
| `--` | Decrementa en 1 | `x--` → `x = x - 1` |

**Nota especial**: ¡En Go, `++` y `--` son **sentencias**, no expresiones!

```go
x := 5

// ✅ CORRECTO
x++
fmt.Println(x)  // 6

// ❌ INCORRECTO
y := x++  // Error: ++ no retorna valor

// ✅ CORRECTO
y := x + 1  // Usa + 1 si necesitas el valor
```

---

## Parte IV: Entrada y Salida Secuencial

### Leyendo Input (Entrada del Usuario)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Paso 1: Solicitar nombre
    fmt.Print("¿Cuál es tu nombre? ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    nombre := strings.TrimSpace(scanner.Text())

    // Paso 2: Solicitar edad
    fmt.Print("¿Cuál es tu edad? ")
    scanner.Scan()
    edadStr := scanner.Text()
    edad, _ := strconv.Atoi(edadStr)  // Convierte string a int

    // Paso 3: Procesar (en orden)
    anoNacimiento := 2024 - age
    
    // Paso 4: Mostrar resultado
    fmt.Printf("Hola %s, naciste aproximadamente en %d\n", nombre, anoNacimiento)
}
```

**¿Qué pasa?**

1. **Línea 1**: Imprime pregunta
2. **Línea 2-3**: Lee respuesta del usuario
3. **Línea 4**: Guarda el nombre
4. **Línea 5-8**: Repite con edad
5. **Línea 9-10**: Calcula año de nacimiento
6. **Línea 11**: Muestra resultado

**El orden es CRÍTICO**: No puedes usar `nombre` o `edad` antes de que el usuario las ingrese.

### Ejemplo Completo: Calculadora Secuencial

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    // Paso 1: Obtener primer número
    fmt.Print("Ingresa el primer número: ")
    num1Str, _ := reader.ReadString('\n')
    num1, _ := strconv.ParseFloat(num1Str[:len(num1Str)-1], 64)

    // Paso 2: Obtener segundo número
    fmt.Print("Ingresa el segundo número: ")
    num2Str, _ := reader.ReadString('\n')
    num2, _ := strconv.ParseFloat(num2Str[:len(num2Str)-1], 64)

    // Paso 3: Realizar operaciones
    suma := num1 + num2
    resta := num1 - num2
    multiplicacion := num1 * num2
    division := num1 / num2

    // Paso 4: Mostrar todos los resultados
    fmt.Printf("\n%v + %v = %v\n", num1, num2, suma)
    fmt.Printf("%v - %v = %v\n", num1, num2, resta)
    fmt.Printf("%v × %v = %v\n", num1, num2, multiplicacion)
    fmt.Printf("%v ÷ %v = %v\n", num1, num2, division)
}
```

**Ejecución:**
```
Ingresa el primer número: 10
Ingresa el segundo número: 5

10 + 5 = 15
10 - 5 = 5
10 × 5 = 50
10 ÷ 5 = 2
```

---

## Parte V: Flujo de Datos Secuencial

### El Concepto de Pipeline (Tubería de Datos)

La estructura secuencial es como una **tubería**:

```
Entrada → Procesamiento 1 → Procesamiento 2 → Salida
```

**Ejemplo práctico:**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Paso 1: Entrada
    texto := "hola mundo"

    // Paso 2: Procesamiento 1 (convertir a mayúsculas)
    textoProcesado := strings.ToUpper(texto)
    fmt.Println("Paso 1:", textoProcesado)  // HOLA MUNDO

    // Paso 3: Procesamiento 2 (reemplazar espacios)
    textoProcesado = strings.ReplaceAll(textoProcesado, " ", "_")
    fmt.Println("Paso 2:", textoProcesado)  // HOLA_MUNDO

    // Paso 4: Procesamiento 3 (agregar prefijo)
    resultado := "TEXTO: " + textoProcesado
    fmt.Println("Paso 3:", resultado)  // TEXTO: HOLA_MUNDO

    // Paso 5: Salida final
    fmt.Println("\nResultado final:", resultado)
}
```

**Salida:**
```
Paso 1: HOLA MUNDO
Paso 2: HOLA_MUNDO
Paso 3: TEXTO: HOLA_MUNDO

Resultado final: TEXTO: HOLA_MUNDO
```

**¿Qué aprendemos?**

Cada paso depende del anterior. Si cambias el orden:
```go
// ❌ Orden incorrecto
textoProcesado = strings.ReplaceAll(textoProcesado, " ", "_")
textoProcesado = strings.ToUpper(texto)  // ¡Vuelve a convertir!
```

Obtendrías:
```
HOLA MUNDO  (los guiones se perdieron)
```

---

## Parte VI: Ejemplo Integrado - Sistema de Carrito de Compras

```go
package main

import (
    "fmt"
)

func main() {
    // ===== PASO 1: DEFINIR PRODUCTOS =====
    producto1 := "Laptop"
    precio1 := 1200.00
    cantidad1 := 2

    producto2 := "Mouse"
    precio2 := 25.50
    cantidad2 := 3

    // ===== PASO 2: CALCULAR SUBTOTALES =====
    subtotal1 := precio1 * float64(cantidad1)
    subtotal2 := precio2 * float64(cantidad2)

    fmt.Println("=== CARRITO DE COMPRAS ===")
    fmt.Printf("%s: %d × $%.2f = $%.2f\n", producto1, cantidad1, precio1, subtotal1)
    fmt.Printf("%s: %d × $%.2f = $%.2f\n", producto2, cantidad2, precio2, subtotal2)

    // ===== PASO 3: CALCULAR TOTAL =====
    subtotalBase := subtotal1 + subtotal2
    fmt.Printf("\nSubtotal: $%.2f\n", subtotalBase)

    // ===== PASO 4: APLICAR IMPUESTO (21%) =====
    impuesto := subtotalBase * 0.21
    fmt.Printf("Impuesto (21%%): $%.2f\n", impuesto)

    // ===== PASO 5: APLICAR DESCUENTO (si aplica) =====
    descuentoPorcentaje := 0.0
    if subtotalBase > 1000 {
        descuentoPorcentaje = 0.10  // 10% descuento
        fmt.Println("¡Descuento aplicado: 10%!")
    }
    descuento := subtotalBase * descuentoPorcentaje
    fmt.Printf("Descuento: -$%.2f\n", descuento)

    // ===== PASO 6: CALCULAR TOTAL FINAL =====
    totalFinal := (subtotalBase + impuesto) - descuento
    fmt.Printf("\n💰 TOTAL: $%.2f\n", totalFinal)

    // ===== PASO 7: MOSTRAR RESUMEN =====
    fmt.Println("\n=== RESUMEN ===")
    fmt.Printf("Productos: 2\n")
    fmt.Printf("Items: %d\n", cantidad1+cantidad2)
    fmt.Printf("Total a pagar: $%.2f\n", totalFinal)
}
```

**Salida:**
```
=== CARRITO DE COMPRAS ===
Laptop: 2 × $1200.00 = $2400.00
Mouse: 3 × $25.50 = $76.50

Subtotal: $2476.50
Impuesto (21%): $519.07
¡Descuento aplicado: 10%!
Descuento: -$247.65

💰 TOTAL: $2747.92

=== RESUMEN ===
Productos: 2
Items: 5
Total a pagar: $2747.92
```

---

## Parte VII: Mejores Prácticas en Estructura Secuencial

### ✅ BUENAS PRÁCTICAS

**1. Inicializa variables antes de usarlas**
```go
// ✅ Correcto
x := 0
x = x + 5
fmt.Println(x)

// ❌ Incorrecto
fmt.Println(x)  // Error: x no existe
x := 0
```

**2. Agrupa operaciones relacionadas**
```go
// ✅ Correcto: Agrupa lectura
fmt.Print("Nombre: ")
scanner.Scan()
nombre := scanner.Text()

fmt.Print("Edad: ")
scanner.Scan()
edad := scanner.Text()

// Luego procesa
```

**3. Usa nombres descriptivos**
```go
// ✅ Correcto
totalConImpuesto := subtotal + (subtotal * 0.21)

// ❌ Incorrecto
x := y + (y * 0.21)
```

**4. Comenta pasos importantes**
```go
// ✅ Bueno
// Paso 1: Leer datos
scanner.Scan()
edad := scanner.Text()

// Paso 2: Convertir a número
ageInt, _ := strconv.Atoi(edad)

// Paso 3: Validar
if ageInt < 0 {
    fmt.Println("Edad inválida")
}
```

### ❌ ANTI-PATRONES

**1. Variables mágicas sin contexto**
```go
// ❌ Malo: ¿Qué es 0.21?
total := subtotal * 0.21

// ✅ Bueno
impuestoPorcentaje := 0.21
total := subtotal * impuestoPorcentaje
```

**2. Cálculos sin guardar resultados intermedios**
```go
// ❌ Muy difícil de leer
fmt.Println(((a + b) * c - d) / e + f)

// ✅ Claro y fácil de debuggear
paso1 := a + b
paso2 := paso1 * c
paso3 := paso2 - d
paso4 := paso3 / e
resultado := paso4 + f
fmt.Println(resultado)
```

**3. No validar entrada**
```go
// ❌ Peligroso
edad, _ := strconv.Atoi(input)  // ¿Qué si input no es número?

// ✅ Seguro
edad, err := strconv.Atoi(input)
if err != nil {
    fmt.Println("Error: debe ser un número")
    return
}
```

---

## Tabla de Referencia: Estructura Secuencial

| Concepto | Ejemplo | Notas |
|----------|---------|-------|
| **Declaración** | `x := 5` | Crea variable, Go infiere tipo |
| **Asignación** | `x = 10` | Cambia valor (tipo permanece) |
| **Suma y asigna** | `x += 5` | Equivalente a `x = x + 5` |
| **Concatenación** | `s := "A" + "B"` | Solo strings, no mixed types |
| **Conversión** | `fmt.Sprint(5)` | Convierte número a string |
| **Lectura** | `scanner.Scan()` | Lee línea del usuario |
| **Cálculo** | `resultado := a + b` | Operación y almacena |
| **Impresión** | `fmt.Println()` | Muestra en consola |

---

## Conclusión: La Belleza de lo Secuencial

La estructura secuencial es el **fundamento de toda programación**. Parece simple, pero es poderosa:

- **Previsibilidad**: Sabes exactamente qué pasa y en qué orden
- **Debuggeo**: Fácil encontrar errores (ejecuta paso por paso)
- **Claridad**: El código se lee como instrucciones normales

En Go, se prioriza esta claridad. Prefiere código que se lee de arriba hacia abajo, sin saltos mentales. Como dijo Dijkstra:

> "Estructuración secuencial es la base de entender programas complejos"

Tu viaje como programador Go comienza aquí: domina la secuencialidad, y todo lo demás será más fácil.

**Próximo paso**: Una vez domines esto, aprenderás `control_flujo.md` para tomar decisiones, y `bucles_iteraciones.md` para repetir operaciones.

¡Vamos! 🚀
