# Biblioteca `math/rand` en Go

La biblioteca `math/rand` sirve para generar números pseudoaleatorios en Go.

## Funciones y tipos principales

- `rand.Seed(seed)` : inicializa el generador global con una semilla.
- `rand.Seed64(seed)` : inicializa el generador global con una semilla de 64 bits.
- `rand.Int()` : genera un entero no negativo aleatorio.
- `rand.Intn(n)` : genera un entero aleatorio en `[0, n)`.
- `rand.Int31()` : genera un entero de 31 bits no negativo.
- `rand.Int31n(n)` : genera un entero de 31 bits en `[0, n)`.
- `rand.Int63()` : genera un entero de 63 bits no negativo.
- `rand.Int63n(n)` : genera un entero de 63 bits en `[0, n)`.
- `rand.Uint32()` : genera un entero sin signo de 32 bits.
- `rand.Uint64()` : genera un entero sin signo de 64 bits.
- `rand.Float32()` : genera un `float32` en `[0.0, 1.0)`.
- `rand.Float64()` : genera un `float64` en `[0.0, 1.0)`.
- `rand.NormFloat64()` : genera un valor normal (distribución gaussiana, media 0, desviación 1).
- `rand.ExpFloat64()` : genera un valor exponencial (media 1).
- `rand.Perm(n)` : genera una permutación aleatoria de `0..n-1`.
- `rand.Shuffle(n, swap)` : mezcla una secuencia de tamaño `n`.
- `rand.Read(p)` : rellena el slice `p` con bytes pseudoaleatorios.
- `rand.NewSource(seed)` : crea una fuente determinista de pseudoaleatoriedad.
- `rand.NewTextSource(seed)` : crea una fuente de texto determinista (Go 1.20+).
- `rand.New(src)` : crea un generador independiente a partir de una fuente.

## Ejemplos

### Generar enteros y flotantes
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    fmt.Println("Int:", rand.Int())
    fmt.Println("Intn 0-9:", rand.Intn(10))
    fmt.Println("Int63:", rand.Int63())
    fmt.Println("Float32:", rand.Float32())
    fmt.Println("Float64:", rand.Float64())
}
```

### Mezclar y permutar
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    datos := []int{1, 2, 3, 4, 5}
    rand.Shuffle(len(datos), func(i, j int) {
        datos[i], datos[j] = datos[j], datos[i]
    })
    fmt.Println("Shuffle:", datos)

    perm := rand.Perm(5)
    fmt.Println("Perm:", perm)
}
```

### Generadores independientes y semilla fija
```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    src := rand.NewSource(42)
    r := rand.New(src)

    fmt.Println("Rand1:", r.Intn(100))
    fmt.Println("Rand2:", r.Intn(100))
}
```

### Distribuciones avanzadas
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    fmt.Println("NormFloat64:", rand.NormFloat64())
    fmt.Println("ExpFloat64:", rand.ExpFloat64())
}
```

## Buenas prácticas
- Usa `rand.Seed(time.Now().UnixNano())` para variar los valores en cada ejecución.
- Para resultados reproducibles en pruebas, usa generadores independientes con `rand.New(rand.NewSource(seed))`.
- No uses `math/rand` para criptografía; para eso está `crypto/rand`.
