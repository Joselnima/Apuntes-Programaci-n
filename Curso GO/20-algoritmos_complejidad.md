# Algoritmos y Complejidad en Go

Entender complejidad (Big O) es crítico para escribir código eficiente. Un algoritmo "correcto" pero ineficiente es un problema.

## 1. Notación Big O

Mide cómo escala el tiempo/espacio con el tamaño de entrada (n).

### Órdenes comunes

```
O(1)        Constante       Hash lookup
O(log n)    Logarítmica     Binary search
O(n)        Lineal          Buscar un elemento
O(n log n)  Lineal-log      Merge sort, Quick sort
O(n²)       Cuadrática      Bubble sort, Insertion sort
O(n³)       Cúbica          3 loops anidados
O(2ⁿ)       Exponencial     Recursión sin memoización
O(n!)       Factorial       Permutaciones
```

### Visualización

```
O(1)
|          ___
|         |
|_________|

O(log n)
|        /
|       /
|      /
|___/

O(n)
|     /
|    /
|   /
|__/

O(n²)
|        /
|       /
|      /
|     /
|____/

O(2ⁿ)
|              /
|            /
|          /
|       /
|_____/
```

---

## 2. Búsqueda lineal vs Búsqueda binaria

### Búsqueda Lineal - O(n)

```go
// Peor caso: recorrer todo el array
func BusquedaLineal(arr []int, objetivo int) int {
    for i, v := range arr {
        if v == objetivo {
            return i
        }
    }
    return -1
}

// n = 1,000,000 → hasta 1,000,000 comparaciones
```

### Búsqueda Binaria - O(log n)

```go
// Array DEBE estar ordenado
// Divide y conquista
func BusquedaBinaria(arr []int, objetivo int) int {
    izq, der := 0, len(arr)-1
    
    for izq <= der {
        mid := (izq + der) / 2
        
        if arr[mid] == objetivo {
            return mid
        } else if arr[mid] < objetivo {
            izq = mid + 1  // Buscar mitad derecha
        } else {
            der = mid - 1  // Buscar mitad izquierda
        }
    }
    
    return -1
}

// n = 1,000,000 → máximo 20 comparaciones (log₂(1M) ≈ 20)
```

**Comparación**:
- Array 1,000 elementos: Linear = 1000 ops, Binary = 10 ops
- Array 1,000,000 elementos: Linear = 1M ops, Binary = 20 ops

---

## 3. Ordenamiento

### Bubble Sort - O(n²)

```go
func BubbleSort(arr []int) {
    n := len(arr)
    
    // Externo: n pasadas
    for i := 0; i < n; i++ {
        // Interno: n comparaciones cada pasada
        for j := 0; j < n-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// Caso peor: array inverso → n² comparaciones
// No recomendado para datos grandes
```

### Quick Sort - O(n log n) promedio

```go
func QuickSort(arr []int, baja, alta int) {
    if baja < alta {
        // Partir array
        p := Particionar(arr, baja, alta)
        
        // Ordenar mitades recursivamente
        QuickSort(arr, baja, p-1)
        QuickSort(arr, p+1, alta)
    }
}

func Particionar(arr []int, baja, alta int) int {
    pivote := arr[alta]
    i := baja - 1
    
    for j := baja; j < alta; j++ {
        if arr[j] < pivote {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    
    arr[i+1], arr[alta] = arr[alta], arr[i+1]
    return i + 1
}

// Mucho más rápido que Bubble Sort en práctica
// Go usa Quick Sort + Insertion Sort en sort.Ints()
```

### Merge Sort - O(n log n) garantizado

```go
func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    // Dividir
    mid := len(arr) / 2
    izq := MergeSort(arr[:mid])
    der := MergeSort(arr[mid:])
    
    // Conquistar (mezclar)
    return Mezclar(izq, der)
}

func Mezclar(izq, der []int) []int {
    resultado := make([]int, 0, len(izq)+len(der))
    i, j := 0, 0
    
    for i < len(izq) && j < len(der) {
        if izq[i] <= der[j] {
            resultado = append(resultado, izq[i])
            i++
        } else {
            resultado = append(resultado, der[j])
            j++
        }
    }
    
    resultado = append(resultado, izq[i:]...)
    resultado = append(resultado, der[j:]...)
    return resultado
}

// Estable, predecible, pero usa más memoria O(n)
```

### Usar sort de Go

```go
import (
    "fmt"
    "sort"
)

func main() {
    arr := []int{5, 2, 8, 1, 9}
    sort.Ints(arr)
    fmt.Println(arr)  // [1 2 5 8 9]
    
    str := []string{"c", "a", "b"}
    sort.Strings(str)
    fmt.Println(str)  // [a b c]
    
    // Ordenar custom
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]  // Descendente
    })
}
```

| Algoritmo | Tiempo | Espacio | Estable | Uso |
|-----------|--------|---------|---------|-----|
| Bubble | O(n²) | O(1) | ✅ | Educativo |
| Quick | O(n log n) | O(log n) | ❌ | General |
| Merge | O(n log n) | O(n) | ✅ | Datos grandes |
| Insertion | O(n²) | O(1) | ✅ | Arrays pequeños |

---

## 4. Estructuras de datos y complejidad

```go
// Array/Slice
Acceso:   O(1)    arr[5]
Buscar:   O(n)    hayar valor
Insertar: O(n)    insertar en medio
Eliminar: O(n)    eliminar del medio

// Map
Acceso:   O(1)    map[clave]
Insertar: O(1)    map[clave] = valor
Buscar:   O(1)    buscar en mapa
Eliminar: O(1)    delete(map, clave)

// Linked List
Acceso:   O(n)    necesita recorrer
Buscar:   O(n)    recorrer lista
Insertar: O(1)    si tenemos posición
Eliminar: O(1)    si tenemos nodo

// Binary Search Tree
Acceso:   O(log n) promedio, O(n) peor
Buscar:   O(log n)
Insertar: O(log n)
Eliminar: O(log n)

// Hash Table (Map)
Acceso:   O(1)
Buscar:   O(1)
Insertar: O(1)
Eliminar: O(1)
```

---

## 5. Técnicas de optimización

### Memoización (Caching)

Guardar resultados para evitar recálculos.

```go
// Fibonacci sin memoización - O(2ⁿ) ❌
func Fib(n int) int {
    if n <= 1 {
        return n
    }
    return Fib(n-1) + Fib(n-2)
    // Fib(40) = millones de llamadas!
}

// Fibonacci con memoización - O(n) ✅
func FibMemo(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    
    if val, existe := memo[n]; existe {
        return val  // Retornar resultado cachado
    }
    
    resultado := FibMemo(n-1, memo) + FibMemo(n-2, memo)
    memo[n] = resultado
    return resultado
}

func main() {
    memo := make(map[int]int)
    fmt.Println(FibMemo(40, memo))  // Instantáneo
}
```

### Programación Dinámica

Construir solución de abajo hacia arriba.

```go
// Fibonacci iterativo - O(n), O(1) espacio
func FibDP(n int) int {
    if n <= 1 {
        return n
    }
    
    prev, curr := 0, 1
    for i := 2; i <= n; i++ {
        prev, curr = curr, prev+curr
    }
    return curr
}
```

### Divide y Conquista

Partir problema en subproblemas independientes.

```go
// Binary Search es divide y conquista
// Merge Sort es divide y conquista
// Quick Sort es divide y conquista

// Contar inversiones en array (ejemplar)
func ContarInversiones(arr []int) int {
    if len(arr) <= 1 {
        return 0
    }
    
    mid := len(arr) / 2
    izq := arr[:mid]
    der := arr[mid:]
    
    // Contar en mitades + durante merge
    inversiones := ContarInversiones(izq) + 
                   ContarInversiones(der)
    
    // Contar inversiones entre mitades
    i, j := 0, 0
    for i < len(izq) && j < len(der) {
        if izq[i] > der[j] {
            inversiones += len(izq) - i
            j++
        } else {
            i++
        }
    }
    
    return inversiones
}
```

---

## 6. Análisis práctico

### Ejemplo: Duplicados

```go
// Solución 1: Fuerza bruta - O(n²)
func TieneDuplicadosFB(arr []int) bool {
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] == arr[j] {
                return true
            }
        }
    }
    return false
}

// Solución 2: Map - O(n)
func TieneDuplicadosMap(arr []int) bool {
    visto := make(map[int]bool)
    for _, v := range arr {
        if visto[v] {
            return true
        }
        visto[v] = true
    }
    return false
}

// Array 1,000,000 elementos
// Fuerza bruta:  1 billón de comparaciones
// Map:            1 millón de operaciones
```

### Ejemplo: Encontrar pares que suman X

```go
// Solución 1: Fuerza bruta - O(n²)
func ParesSumanFB(arr []int, suma int) [][]int {
    pares := [][]int{}
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i]+arr[j] == suma {
                pares = append(pares, []int{arr[i], arr[j]})
            }
        }
    }
    return pares
}

// Solución 2: Hash map - O(n)
func ParesSumanMap(arr []int, suma int) [][]int {
    pares := [][]int{}
    visto := make(map[int]bool)
    
    for _, v := range arr {
        complemento := suma - v
        if visto[complemento] {
            pares = append(pares, []int{v, complemento})
        }
        visto[v] = true
    }
    return pares
}
```

---

## 7. Recursión vs Iteración

### Recursión: Elegante, pero cuidado con stack

```go
// Elegant pero O(n) stack space
func SumaRecursiva(arr []int, idx int) int {
    if idx >= len(arr) {
        return 0
    }
    return arr[idx] + SumaRecursiva(arr, idx+1)
}

// Stack overflow si n es gigante
```

### Iteración: Más eficiente

```go
// O(1) stack space
func SumaIterativa(arr []int) int {
    suma := 0
    for _, v := range arr {
        suma += v
    }
    return suma
}
```

---

## 8. Profiling en Go

Medir rendimiento real.

```go
import (
    "fmt"
    "time"
)

func Benchmark(nombre string, f func()) {
    inicio := time.Now()
    f()
    duracion := time.Since(inicio)
    fmt.Printf("%s: %v\n", nombre, duracion)
}

func main() {
    arr := make([]int, 1000000)
    for i := range arr {
        arr[i] = i
    }
    
    Benchmark("Búsqueda Lineal", func() {
        BusquedaLineal(arr, 999999)
    })
    
    Benchmark("Búsqueda Binaria", func() {
        BusquedaBinaria(arr, 999999)
    })
}
```

---

## 9. Problemas clásicos

### Fibonacci
```
O(2ⁿ) recursión pura
O(n) con memoización
O(n) iterativo
O(1) solo últimos dos valores
```

### Factorial
```go
func Factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * Factorial(n-1)  // O(n) time, O(n) stack
}
```

### Towers of Hanoi
```
3 discos: 7 movimientos
4 discos: 15 movimientos
n discos: 2ⁿ - 1 movimientos → O(2ⁿ)
```

---

## 10. Consejos prácticos

### ✅ Bien
```go
// Elegir estructura de datos correcta
usuarios := make(map[string]Usuario)  // Búsqueda O(1)

// Evitar loops anidados cuando sea posible
for _, x := range []int{...} {
    if m[x] {  // O(1) lookup
        // ...
    }
}

// Usar sort.Slice para lógica compleja
sort.Slice(datos, func(i, j int) bool {
    return datos[i].Prioridad > datos[j].Prioridad
})

// Medir antes de optimizar
// "Premature optimization is the root of all evil" - Knuth
```

### ❌ Mal
```go
// Triple loop anidado = O(n³) ❌
for i := range arr {
    for j := range arr {
        for k := range arr {
            // ...
        }
    }
}

// Búsqueda lineal en loop = O(n²) ❌
for _, usuario := range usuarios {
    if indiceUsuario(usuarios, usuario.ID) > 0 {
        // ...
    }
}

// Crear copia innecesaria = O(n) extra
resultado := arr[:]  // mejor que append([]int{}, arr...)
```

---

## 12. Conclusión - El Camino Hacia la Eficiencia

Entender complejidad y algoritmos no es luxuria, es **necesario**.

### La verdad incómoda

Escrbir código que funciona es fácil. Escribir código **que funciona rápido** requiere pensamiento.

Un algoritmo O(n²) que funciona en 0.1 segundos con n=1000, podría tomar **16 minutos** con n=1,000,000.

Eso es la diferencia entre un sistema que escala y uno que colapsa.

### Regla 80/20 de optimización

- **80%** del tiempo, un algoritmo O(n) es suficiente
- **19%** del tiempo, necesitas O(log n) o O(n log n)
- **1%** del tiempo, necesitas algo ingeniero/especial

No optimices prematuramente. Pero cuando necesites, **sabe dónde cambiar**.

### Decisiones clave para datos grandes

| Tamaño | Estructura | Algoritmo |
|--------|-----------|-----------|
| < 100 items | Array/Slice | O(n) search OK |
| 100-10k | Map | O(1) lookup |
| 10k-1M | Indexed DB | Binary search |
| 1M+ | Database | Indexed queries |
| Streaming | Pipes | Process as comes |

### La jerarquía de prioridades

```
1. Funciona correctamente
2. Es legible (otro puede entender)
3. Es lo suficientemente rápido (measure!)
4. Es elegante/eficiente
```

Muchos programadores voltean (3) y (4). Big mistake.

### Ejemplo real: búsqueda de usuario

**Versión 1** (Naïve): Loop lineal O(n)
```go
func BuscarUsuario(id int) *Usuario {
    for _, u := range usuarios {
        if u.ID == id {
            return &u
        }
    }
    return nil
}
 // Con 1M usuarios: hasta 1M comparaciones
```

**Versión 2** (Perezoso): Map O(1)
```go
var usuariosMap = make(map[int]*Usuario)

func BuscarUsuario(id int) *Usuario {
    return usuariosMap[id]
}
// Con 1M usuarios: una operación hash
```

**Diferencia**: **1M operaciones vs 1 operación**. Eso es 1,000,000x más rápido.

### Tools para medir

```bash
# Benchmark en Go
go test -bench=. -benchmem

# Profiling
go run -cpuprofile=cpu.prof programa.go
go tool pprof cpu.prof

# Tracing
import _ "net/http/pprof"
# Va a http://localhost:6060/debug/pprof
```

### Conceptos finales

**Big O es conversación mental**: "¿Cómo escala esto?" No necesitas ser exacto, necesitas saber orden de magnitud.

**Leer código es más importante que escribirlo**: Alguien mantendrá tu código. Haz que sea claro, luego rápido.

**La mejor optimización es simple**: SQL con índice, Map en lugar de Array, Algoritmo diferente. No micro-optimizaciones tontas.

### La jorn del programador con complejidad

**Semana 1**: "¿Qué es Big O?"

**Mes 1**: Aprendes O(n²) es malo.

**Mes 3**: Entiendes por qué Binary Search es 2^16x más rápido.

**Año 1**: Automáticamente elegis estructura de datos correcta.

**Después**: Casi nunca piensas en ello porque diseñas bien desde el inicio.

**Go te enseña a respetar la computadora. Aprende esto en Go, aplícalo en cualquier lenguaje.**
