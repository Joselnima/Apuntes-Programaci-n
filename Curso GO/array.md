# Tipo de dato: array - Guía Completa

> **Colecciones de tamaño fijo: cómo y cuándo usar arrays en Go**

---

## ¿Por qué existen los arrays?

Imagina que quieres almacenar las **calificaciones de 5 estudiantes** de una clase pequeña. Sabes exactamente que son 5, ni más ni menos. Podrías crear 5 variables:

```go
// ❌ Mucho código innecesario
nota1 := 85
nota2 := 90
nota3 := 78
nota4 := 92
nota5 := 88
```

Pero esto es **repetitivo, difícil de manejar y frágil**. ¿Y si después necesitas 20 calificaciones?

Los **arrays** son la solución: **una agrupación de elementos del mismo tipo con tamaño fijo**, sabido desde el inicio.

```go
// ✅ Limpio y manejable
calificaciones := [5]int{85, 90, 78, 92, 88}
```

**La diferencia clave:**
- **Array**: Tamaño **fijo** (conocido en tiempo de compilación)
- **Slice**: Tamaño **dinámico** (crece/decrece en runtime)

**¿Cuándo cada uno?**
| Situación | Usa |
|-----------|-----|
| Conoces el tamaño exacto (ej: 5 notas) | Array |
| No sabes cuántos elementos tendrás | Slice |
| Necesitas garantía de memoria fija | Array |
| Necesitas agregar/quitar elementos | Slice |

---

## Definición

En Go, un `array` es una colección de elementos del mismo tipo con un tamaño **fijo** conocido en tiempo de compilación. El tamaño forma parte del tipo.

**Ejemplo de tipo:** `[5]int` es un array de 5 enteros. `[5]int` y `[4]int` son **tipos completamente diferentes**.

## Características Fundamentales

1. **Tamaño es parte del tipo**: `[3]string` ≠ `[4]string`
2. **Tamaño fijo**: No cambias la longitud después de crear el array
3. **Copiable**: Array pequeño = copia por valor (eficiente)
4. **Compactación en memoria**: Los elementos están contiguos
5. **Acceso rápido**: O(1) para acceder a cualquier elemento por índice

**Comparación con Slice:**

```go
// Array: tamaño fijo, es el tipo
var numeros [5]int     // Siempre 5 elementos

// Slice: tamaño dinámico, vista de array
var items []int        // Puede crecer/decrecer
```

## Para qué sirve

- Almacenar colecciones con **longitud garantizada y conocida**
- Modelar **estructuras estáticas** (coordenadas de un triángulo, RGB de color)
- Buffers de **bajo nivel** donde la eficiencia importa
- Asegurar **memoria predecible** (sin reasignaciones)

## Cuándo se usa

- **Cuando conoces el tamaño exacto desde el diseño** (5 calificaciones, 3 dimensiones XYZ, 256 píxeles)
- En **algoritmos especializados** que necesitan buffer fijo
- Para **datos predefinidos** que nunca cambian de tamaño
- En **código de bajo nivel** o **librerías de performance crítica**

---

## Parte I: Declaración y Creación

### Forma 1: Declarar vacío y llenar

```go
// ❌ Declarar array sin inicializar (los valores son cero)
var numeros [5]int
numeros[0] = 10
numeros[1] = 20
// numeros = [5]int{10, 20, 0, 0, 0}

// ✅ Forma más clara (inicializador)
numeros := [5]int{10, 20, 30, 40, 50}
```

### Forma 2: Dejar que Go cuente (inferencia)

```go
// Go infiere el tamaño del array por los elementos
numeros := [...]int{10, 20, 30, 40, 50}
// Go entiende: [5]int

// Bueno para arrays pequeños, no es recomendable para producción
```

### Forma 3: Parcialmente inicializado

```go
// Array de 10 elementos, pero solo inicializa algunos
// El resto son valores cero automáticamente
calificaciones := [10]int{85, 90, 78}
// calificaciones = [10]int{85, 90, 78, 0, 0, 0, 0, 0, 0, 0}
```

### Forma 4: Inicializador con índices

```go
// Especificar índices explícitamente
vertices := [5]float64{
    0: 0.0,
    2: 5.5,
    4: 10.0,
}
// vertices = [5]float64{0.0, 0, 5.5, 0, 10.0}
```

**Comparación de formas:**

```go
// Forma larga: var con tipo e inicialización
var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

// Forma corta: inferencia con :=
arr2 := [5]int{1, 2, 3, 4, 5}

// Forma más corta: inferencia de tamaño
arr3 := [...]int{1, 2, 3, 4, 5}

// Usa la forma corta en la mayoría de casos
```

---

## Parte II: Acceso y Modificación

### Acceder a elementos (indexación)

```go
frutas := [3]string{"manzana", "plátano", "naranja"}

// Acceso por índice (comienza en 0)
fmt.Println(frutas[0])      // "manzana"
fmt.Println(frutas[1])      // "plátano"
fmt.Println(frutas[2])      // "naranja"
fmt.Println(frutas[len(frutas)-1])  // Último elemento: "naranja"

// ❌ Índice fuera de rango (PANIC)
// fmt.Println(frutas[3])  // panic: index out of range
```

**Recorrer con índice:**

```go
for i := 0; i < len(frutas); i++ {
    fmt.Printf("frutas[%d] = %s\n", i, frutas[i])
}
```

### Recorrer con range (más idiomático en Go)

```go
frutas := [3]string{"manzana", "plátano", "naranja"}

// range devuelve (índice, valor)
for i, fruta := range frutas {
    fmt.Printf("%d: %s\n", i, fruta)
}

// Si solo necesitas el valor:
for _, fruta := range frutas {
    fmt.Println(fruta)
}

// Si solo necesitas el índice:
for i := range frutas {
    fmt.Println(i)
}
```

### Modificar elemento

```go
numeros := [5]int{10, 20, 30, 40, 50}

// Cambiar elemento en índice 2
numeros[2] = 99
// numeros = [5]int{10, 20, 99, 40, 50}

// Modificar múltiples
numeros[0] = 5
numeros[4] = 200
// numeros = [5]int{5, 20, 99, 40, 200}
```

---

## Parte III: Restricción Fundamental - Tamaño Fijo

### ❌ No puedes agregar elementos (error común)

```go
numeros := [5]int{1, 2, 3}

// ❌ Esto NO COMPILA
numeros = append(numeros, 4)  // Error: no se puede append a array

// ✅ Solución: convertir a slice
numerosSlice := append(numeros[:], 4)
// numerosSlice es ahora un slice, no array
```

### ❌ No puedes eliminar elementos

```go
numeros := [5]int{10, 20, 30, 40, 50}

// ❌ No hay forma de eliminar manteniendo la estructura
// Opción 1: asignar cero
numeros[2] = 0  // [10 20 0 40 50] - pero sigue siendo 5 elementos

// ✅ Opción 2: crear slice sin ese elemento (los elementos que quedan)
sinIndice2 := append(numeros[:2], numeros[3:]...)
// sinIndice2 es ahora [10 20 40 50] (4 elementos)
```

**¿Se necesita tamaño dinámico?** → **USA SLICE EN SU LUGAR**

```go
// ✅ Mejor enfoque: usa slice desde el principio
numeros := []int{1, 2, 3}  // slice, no array
numeros = append(numeros, 4, 5)  // ✅ Funciona
```

---

## Parte IV: Copiar Arrays (Comportamiento Importante)

### Arrays se copian por valor

```go
original := [3]int{1, 2, 3}
copia := original  // Copia TODOS los elementos

copia[0] = 999
fmt.Println(original[0])  // 1 (sin cambios, son arrays separados)
fmt.Println(copia[0])     // 999
```

**Esto es diferente a slices:**

```go
// Slices NO se copian
original := []int{1, 2, 3}
vistaOtra := original  // Ambos apuntan al mismo array subyacente

vistaOtra[0] = 999
fmt.Println(original[0])  // 999 (¡cambió porque comparten datos!)
```

### Copiar arrays con copy()

```go
original := [5]int{1, 2, 3, 4, 5}
destino := [5]int{}

copy(destino[:], original[:])
// destino = [1, 2, 3, 4, 5]
```

**Por qué necesitas `[:]`:** Arrays no son directamente copiables con `copy()`. Tienes que convertirlos a slice primero.



## Parte V: Operaciones Comunes

### Información del array

```go
numeros := [5]int{10, 20, 30, 40, 50}

// Tamaño (siempre conocido)
fmt.Println(len(numeros))   // 5

// Capacidad (arrays = tamaño)
fmt.Println(cap(numeros))   // 5 (igual a len para arrays)
```

### Convertir array a slice

```go
numeros := [5]int{10, 20, 30, 40, 50}

// Array completo como slice
sliceTotal := numeros[:]
// sliceTotal es []int{10, 20, 30, 40, 50}

// Slice parcial
slicePartial := numeros[1:4]
// slicePartial es []int{20, 30, 40} (índices 1, 2, 3)

sliceDesde := numeros[2:]
// sliceDesde es []int{30, 40, 50}

sliceHasta := numeros[:3]
// sliceHasta es []int{10, 20, 30}
```

### Buscar elemento en array

```go
numeros := [5]int{10, 20, 30, 40, 50}

// ❌ Go NO tiene indexOf nativo para arrays
// Solución 1: buscar manualmente
func encontrar(arr [5]int, valor int) int {
    for i, v := range arr {
        if v == valor {
            return i
        }
    }
    return -1
}

indice := encontrar(numeros, 30)
if indice != -1 {
    fmt.Println("Encontrado en índice:", indice)  // 2
} else {
    fmt.Println("No encontrado")
}

// Solución 2: convertir a slice y usar helper
sliceNums := numeros[:]
indice2 := indexOf(sliceNums, 30)
```

### Verificar si existe elemento

```go
palabras := [4]string{"gato", "perro", "pajaro", "pez"}

func contiene(arr [4]string, valor string) bool {
    for _, v := range arr {
        if v == valor {
            return true
        }
    }
    return false}

fmt.Println(contiene(palabras, "gato"))     // true
fmt.Println(contiene(palabras, "elefante")) // false
```

---

## Parte VI: Ejemplo Práctico - Gestor de Inventario Simple

Supongamos que una tienda pequeña tiene exactamente **5 productos** siempre en stock.

```go
package main

import "fmt"

// Estructura para un producto
type Producto struct {
    nombre string
    precio float64
    stock  int
}

// Inventario fijo de 5 productos
var inventario [5]Producto

// Inicializar inventario
func inicializarInventario() {
    inventario = [5]Producto{
        {"Laptop", 999.99, 3},
        {"Mouse", 25.50, 15},
        {"Teclado", 75.00, 8},
        {"Monitor", 299.99, 5},
        {"Cable HDMI", 12.99, 20},
    }
}

// Obtener precio total de stock
func calcularValorTotal() float64 {
    total := 0.0
    for _, producto := range inventario {
        total += producto.precio * float64(producto.stock)
    }
    return total
}

// Listar todos los productos
func listarProductos() {
    fmt.Println("\n=== INVENTARIO ===")
    for i, p := range inventario {
        fmt.Printf("%d. %s - $%.2f (Stock: %d)\n", i+1, p.nombre, p.precio, p.stock)
    }
}

// Buscar producto por nombre
func buscarProducto(nombre string) (int, bool) {
    for i, p := range inventario {
        if p.nombre == nombre {
            return i, true
        }
    }
    return -1, false
}

// Realizar compra
func hacerCompra(nombreProducto string) bool {
    indice, encontrado := buscarProducto(nombreProducto)
    if !encontrado {
        fmt.Printf("Error: Producto '%s' no encontrado\n", nombreProducto)
        return false
    }

    if inventario[indice].stock <= 0 {
        fmt.Printf("Error: %s sin stock\n", nombreProducto)
        return false
    }

    inventario[indice].stock--
    fmt.Printf("Compra exitosa: %s. Stock restante: %d\n", 
        nombreProducto, inventario[indice].stock)
    return true
}

func main() {
    inicializarInventario()
    listarProductos()

    // Realizar compras
    hacerCompra("Laptop")
    hacerCompra("Mouse")
    hacerCompra("Mouse")
    hacerCompra("Smartphone") // No existe

    listarProductos()

    // Mostrar valor total del inventario
    valorTotal := calcularValorTotal()
    fmt.Printf("\nValor total del inventario: $%.2f\n", valorTotal)
}
```

**Output:**
```
=== INVENTARIO ===
1. Laptop - $999.99 (Stock: 3)
2. Mouse - $25.50 (Stock: 15)
3. Teclado - $75.00 (Stock: 8)
4. Monitor - $299.99 (Stock: 5)
5. Cable HDMI - $12.99 (Stock: 20)

Compra exitosa: Laptop. Stock restante: 2
Compra exitosa: Mouse. Stock restante: 14
Compra exitosa: Mouse. Stock restante: 13
Error: Producto 'Smartphone' no encontrado

=== INVENTARIO ===
1. Laptop - $999.99 (Stock: 2)
2. Mouse - $25.50 (Stock: 13)
3. Teclado - $75.00 (Stock: 8)
4. Monitor - $299.99 (Stock: 5)
5. Cable HDMI - $12.99 (Stock: 20)

Valor total del inventario: $4699.82
```

---

## Parte VII: Ordenamiento

### Ordenar array de números

```go
import "sort"

numeros := [5]int{50, 20, 40, 10, 30}

// Convertir a slice y ordenar
sort.Ints(numeros[:])
fmt.Println(numeros)  // [10 20 30 40 50]
```

### Ordenar array de strings

```go
palabras := [4]string{"zebra", "apple", "mango", "banana"}

sort.Strings(palabras[:])
fmt.Println(palabras)  // [apple banana mango zebra]
```

### Ordenamiento personalizado

```go
type Estudiante struct {
    nombre string
    edad   int
}

// Array de estructuras
estudiantes := [3]Estudiante{
    {"Alice", 22},
    {"Bob", 20},
    {"Carlos", 21},
}

// Ordenar por edad (usando slice)
sort.Slice(estudiantes[:], func(i, j int) bool {
    return estudiantes[i].edad < estudiantes[j].edad
})

for _, e := range estudiantes {
    fmt.Printf("%s: %d años\n", e.nombre, e.edad)
}
// Output:
// Bob: 20 años
// Carlos: 21 años
// Alice: 22 años
```

---

## Parte VIII: Pasar Arrays a Funciones

### ⚠️ Atención: Los arrays se copian completos

```go
func modificarArray(arr [5]int) {
    arr[0] = 999
    // Modifica la COPIA, no el original
}

numeros := [5]int{1, 2, 3, 4, 5}
modificarArray(numeros)
fmt.Println(numeros[0])  // 1 (sin cambios)
```

**Esto es ineficiente para arrays grandes.** Solución: pasar por puntero.

### ✅ Pasar por puntero (recomendado)

```go
func modificarArray(arr *[5]int) {
    arr[0] = 999
    // Modifica el original
}

numeros := [5]int{1, 2, 3, 4, 5}
modificarArray(&numeros)
fmt.Println(numeros[0])  // 999 (cambió)
```

### ✅ Pasar como slice (más idiomático en Go)

```go
func procesarDatos(datos []int) {
    for i, v := range datos {
        datos[i] = v * 2
    }
}

numeros := [5]int{1, 2, 3, 4, 5}
procesarDatos(numeros[:])  // Convertir a slice
fmt.Println(numeros)  // [2 4 6 8 10]
```

**Comparación:**

```go
// ❌ Copia todo el array (ineficiente si es grande)
func fn1(arr [1000]int) { }

// ✅ Solo copia puntero
func fn2(arr *[1000]int) { }

// ✅ Solo copia header del slice (recomendado)
func fn3(arr []int) { }
```

**Regla en Go:** Si necesitas modificar un array, usa slice. Es más idiomático.

---

## Parte IX: Anti-patrones - Errores Comunes

### ❌ Ignorar que los arrays tienen tamaño fijo

```go
// ❌ Código que falla
func agregarProducto(arr [5]string) {
    arr = append(arr, "Nuevo")  // COMPILE ERROR
}

// ✅ Usar slice en su lugar
func agregarProducto(arr []string) []string {
    return append(arr, "Nuevo")
}
```

### ❌ Usar array cuando necesitas dinámico

```go
// ❌ Problema: ¿Cuántos usuarios tenemos?
usuarios := [100]User{}

// ✅ Usar slice
var usuarios []User
usuarios = append(usuarios, user1)
usuarios = append(usuarios, user2)
```

### ❌ No convertir a slice para copy()

```go
arr1 := [5]int{1, 2, 3, 4, 5}
arr2 := [5]int{}

// ❌ No compila
copy(arr2, arr1)

// ✅ Convertir a slice
copy(arr2[:], arr1[:])
```

### ❌ Asumir que slice copia el array

```go
original := [5]int{1, 2, 3, 4, 5}
vista := original[:]  // Vista del mismo array

vista[0] = 999
fmt.Println(original[0])  // 999 (ambos apuntan al mismo lugar)
```

### ❌ Pasar array grande a función por valor

```go
// ❌ INEFICIENTE: copia 1MB en la stack
largeArray := [262144]float64{}  // ~2MB
procesarDatos(largeArray)

// ✅ Usa slice o puntero
procesarDatos(largeArray[:])
```

---

## Parte X: Funciones Útiles - Referencia Rápida

| Función | Uso | Ejemplo |
|---------|-----|---------|
| `len(arr)` | Tamaño del array | `len(numeros)` → 5 |
| `cap(arr)` | Capacidad (= tamaño para arrays) | `cap(numeros)` → 5 |
| `copy(dest, src)` | Copiar elementos | `copy(dest[:], src[:])` |
| `range` | Iterar sobre elementos | `for i, v := range arr` |
| `sort.Ints(arr[:])` | Ordenar números | `sort.Ints(nums[:])` |
| `sort.Strings(arr[:])` | Ordenar strings | `sort.Strings(palabras[:])` |
| `sort.Slice(arr[:], fn)` | Ordenamiento personalizado | `sort.Slice(arr[:], func(i,j)...)` |
| `arr[i:j]` | Crear slice | `arr[1:4]` → elementos 1, 2, 3 |
| `arr[:]` | Array como slice | `arr[:]` → slice del array completo |

---

## Parte XI: Ejemplo de Algoritmo - Búsqueda Binaria

Para buscar en un **array ordenado**, búsqueda binaria es eficiente:

```go
// Búsqueda binaria en array ordenado
func busquedaBinaria(arr [10]int, objetivo int) int {
    izquierda, derecha := 0, len(arr)-1

    for izquierda <= derecha {
        mitad := (izquierda + derecha) / 2

        if arr[mitad] == objetivo {
            return mitad  // Encontrado
        } else if arr[mitad] < objetivo {
            izquierda = mitad + 1  // Buscar en la derecha
        } else {
            derecha = mitad - 1    // Buscar en la izquierda
        }
    }

    return -1  // No encontrado
}

func main() {
    numeros := [10]int{2, 5, 8, 12, 16, 23, 38, 45, 56, 67}
    
    fmt.Println(busquedaBinaria(numeros, 23))   // 5
    fmt.Println(busquedaBinaria(numeros, 100))  // -1
}
```

---

## Parte XII: Mejores Prácticas

### 1. Preferir slices sobre arrays (en la mayoría de casos)

```go
// ❌ Arrays fixos son limitados
func procesarNumeros(arr [100]int) {
    // ¿Y si tengo 50 o 200?
}

// ✅ Slices son flexibles
func procesarNumeros(arr []int) {
    // Funciona con cualquier cantidad
}
```

### 2. Documentar el tamaño fijo

```go
// ✅ Si necesitas array fijo, documenta POR QUÉ
// ColorRGB representa un color con 3 canales (R, G, B)
type ColorRGB [3]uint8

color := ColorRGB{255, 128, 0}  // Naranja
```

### 3. Usar nombres descriptivos

```go
// ❌ Poco claro
arr := [5]int{10, 20, 30, 40, 50}

// ✅ Claro
calificaciones := [5]int{10, 20, 30, 40, 50}
```

### 4. Pasar como slice si vas a modificar

```go
// ✅ Mejor práctica
func actualizarPrecios(productos []Producto) {
    for i := range productos {
        productos[i].precio *= 1.1
    }
}
```

### 5. Valores cero son tus aliados

```go
// Arrays autom
aticamente inicializados a valores cero
puntuaciones := [10]int{}  // [0 0 0 0 0 0 0 0 0 0]
nombres := [5]string{}     // ["" "" "" "" ""]
activos := [3]bool{}       // [false false false]
```

---

## Resumen: Arrays vs Slices

| Aspecto | Array | Slice |
|---------|-------|-------|
| Tamaño | ✅ Fijo | ❌ Dinámico |
| Tipo | Tamaño es parte del tipo | Sin tamaño en el tipo |
| Copia | Por valor (todos los datos) | Por referencia (header) |
| `append()` | ❌ No soporta | ✅ Soporta |
| Cuándo usar | Datos estáticos, buffers fijos | Colecciones variables |
| Performance | Mejor para datos fijos | Mejor flexibilidad |

---

## Conclusión

**Arrays son poderosos cuando los necesitas**, pero para la mayoría de aplicaciones Go, **slices son la opción correcta**.

- **Usa arrays cuando:** Sabes exactamente cuántos elementos tendrás siempre
- **Usa slices cuando:** No estás seguro del tamaño o necesitas flexibilidad

La filosofía de Go favorece **slices por defecto, arrays solo cuando tenga sentido**.
