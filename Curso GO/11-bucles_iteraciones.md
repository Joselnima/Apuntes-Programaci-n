# Bucles e Iteraciones en Go - Guía Completa

Los bucles (loops) te permiten repetir código una y otra vez. Son fundamentales: sin ellos, tendríamos que escribir el mismo código mil veces.

---

## Parte I: For - El único bucle de Go

Go tiene **solo una palabra clave para bucles: `for`**. Eso es diferente a otros lenguajes como Python (que tiene `while`) o JavaScript (que tiene `while` y `for`).

La filosofía de Go es: **una forma clara de hacer algo, no diez formas confusas**.

### For clásico (como en C)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
// Output: 0 1 2 3 4
```

**Desglose**:
- `i := 0`: Inicialización (antes de empezar)
- `i < 5`: Condición (continuar mientras sea verdadera)
- `i++`: Incremento (después de cada iteración)

**Flujo**:
1. Inicializa `i = 0`
2. ¿Es `i < 5`? Sí → ejecutar cuerpo
3. Incrementar `i` a 1
4. ¿Es `i < 5`? Sí → ejecutar cuerpo
5. ... repite hasta que `i = 5` y `i < 5` es falso

### For simple (como while)

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
// Output: 0 1 2 3 4
```

Cuando omites la inicialización e incremento, es como un `while` en otros lenguajes.

### Bucle infinito

```go
for {
    fmt.Println("Esto se repite para siempre...")
    break  // Necesitamos break para salir
}
```

Sin condición, el bucle nunca terminará (a menos que uses `break`).

### Break y Continue

```go
// break: sale del bucle completamente
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Sale del for en i = 5
    }
    fmt.Println(i)
}
// Output: 0 1 2 3 4
```

```go
// continue: salta a la siguiente iteración
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // Salta números pares
    }
    fmt.Println(i)
}
// Output: 1 3 5 7 9
```

### Bucles anidados

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("(%d,%d) ", i, j)
    }
    fmt.Println()
}
// Output:
// (0,0) (0,1) (0,2)
// (1,0) (1,1) (1,2)
// (2,0) (2,1) (2,2)
```

---

## Parte II: For Range - Iterar sobre colecciones

La forma más común e idiomática en Go para iterar sobre arrays, slices, maps, strings, etc.

### Range sobre slice

```go
numeros := []int{10, 20, 30, 40}

for i, valor := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", i, valor)
}
// Output:
// Índice: 0, Valor: 10
// Índice: 1, Valor: 20
// Índice: 2, Valor: 30
// Índice: 3, Valor: 40
```

`range` te da:
- El índice actual
- El valor en ese índice

**Si solo necesitas el índice**:
```go
for i := range numeros {
    fmt.Println(i)
}
// Output: 0 1 2 3
```

**Si solo necesitas el valor**:
```go
for _, valor := range numeros {
    fmt.Println(valor)
}
// Output: 10 20 30 40
```

El guion bajo `_` significa "ignoro este valor".

### Range sobre array

```go
estudiantes := [3]string{"Alice", "Bob", "Charlie"}

for i, nombre := range estudiantes {
    fmt.Printf("%d: %s\n", i, nombre)
}
```

Funciona igual que con slices.

### Range sobre map

```go
edades := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Charlie": 35,
}

for nombre, edad := range edades {
    fmt.Printf("%s: %d años\n", nombre, edad)
}
// Output (orden aleatorio):
// Alice: 30 años
// Bob: 25 años
// Charlie: 35 años
```

⚠️ **Importante**: El orden de iteración en un map es **aleatorio**. Si necesitas orden, primero ordena las claves.

```go
// Si necesitas orden específico
nombres := make([]string, 0, len(edades))
for nombre := range edades {
    nombres = append(nombres, nombre)
}
sort.Strings(nombres)
for _, nombre := range nombres {
    fmt.Println(nombre, edades[nombre])
}
```

### Range sobre string

```go
texto := "Hola"

for i, char := range texto {
    fmt.Printf("%d: %c (%d)\n", i, char, char)
}
// Output:
// 0: H (72)
// 1: o (111)
// 2: l (108)
// 3: a (97)
```

Cuando haces range sobre string, obtienes **runes** (caracteres Unicode), no bytes.

```go
texto := "Hola"
bytes := []byte(texto)

for i, b := range bytes {
    fmt.Printf("%d: %c (%d)\n", i, b, b)
}
```

### Range sobre string Unicode

```go
texto := "Hel日本世"

for i, char := range texto {
    fmt.Printf("%d: %c\n", i, char)
}
// Output:
// 0: H
// 1: e
// 2: l
// 3: 日
// 6: 本
// 9: 世
```

Nota que los índices saltan (0, 1, 2, 3, 6, 9). Eso es porque caracteres Unicode pueden ocupar múltiples bytes.

### Range sobre channel (avanzado)

```go
canal := make(chan int, 3)
canal <- 1
canal <- 2
canal <- 3
close(canal)

for valor := range canal {
    fmt.Println(valor)
}
// Output: 1 2 3
```

---

## Parte III: Patrones Comunes

### Patrón: Buscar en slice

```go
numeros := []int{10, 20, 30, 40, 50}
objetivo := 30

encontrado := false
indice := -1

for i, v := range numeros {
    if v == objetivo {
        encontrado = true
        indice = i
        break
    }
}

if encontrado {
    fmt.Printf("Encontrado en índice %d\n", indice)
} else {
    fmt.Println("No encontrado")
}
```

### Patrón: Accumulator (acumular valores)

```go
numeros := []int{1, 2, 3, 4, 5}
suma := 0

for _, n := range numeros {
    suma += n
}

fmt.Println(suma)  // 15
```

```go
// Concatenar strings
palabras := []string{"Hola", " ", "mundo", "!"}
resultado := ""

for _, palabra := range palabras {
    resultado += palabra
}

fmt.Println(resultado)  // "Hola mundo!"
```

⚠️ **Cuidado**: Concatenar strings en loops es ineficiente. Usa `strings.Builder`:

```go
import "strings"

palabras := []string{"Hola", " ", "mundo", "!"}
var sb strings.Builder

for _, palabra := range palabras {
    sb.WriteString(palabra)
}

fmt.Println(sb.String())  // "Hola mundo!"
```

### Patrón: Transformar/Mapear

```go
numeros := []int{1, 2, 3, 4, 5}
cuadrados := make([]int, 0, len(numeros))

for _, n := range numeros {
    cuadrados = append(cuadrados, n*n)
}

fmt.Println(cuadrados)  // [1 4 9 16 25]
```

### Patrón: Filtrar

```go
numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
pares := []int{}

for _, n := range numeros {
    if n%2 == 0 {
        pares = append(pares, n)
    }
}

fmt.Println(pares)  // [2 4 6 8 10]
```

### Patrón: Contar frecuencias

```go
texto := "hola"
frecuencias := make(map[rune]int)

for _, char := range texto {
    frecuencias[char]++
}

fmt.Println(frecuencias)
// Output: map[97:1 104:1 108:1 111:1]  (a:1, h:1, l:1, o:1)
```

---

## Parte IV: Ejemplos Integrados

### Ejemplo 1: Tabla de multiplicar

```go
tabla := 5

fmt.Printf("Tabla del %d:\n", tabla)
for i := 1; i <= 10; i++ {
    resultado := tabla * i
    fmt.Printf("%d x %d = %d\n", tabla, i, resultado)
}
// Output:
// Tabla del 5:
// 5 x 1 = 5
// 5 x 2 = 10
// ... hasta 5 x 10 = 50
```

### Ejemplo 2: Procesador de archivo línea por línea

```go
import (
    "bufio"
    "fmt"
    "os"
)

func ProcesarArchivo(nombreArchivo string) error {
    archivo, err := os.Open(nombreArchivo)
    if err != nil {
        return err
    }
    defer archivo.Close()
    
    scanner := bufio.NewScanner(archivo)
    numeroLinea := 0
    
    for scanner.Scan() {
        numeroLinea++
        linea := scanner.Text()
        fmt.Printf("Línea %d: %s\n", numeroLinea, linea)
    }
    
    return scanner.Err()
}
```

### Ejemplo 3: Búsqueda en matriz 2D

```go
matriz := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

objetivo := 5
encontrado := false

for i, fila := range matriz {
    for j, valor := range fila {
        if valor == objetivo {
            fmt.Printf("Encontrado en posición [%d][%d]\n", i, j)
            encontrado = true
            break
        }
    }
    if encontrado {
        break
    }
}
```

### Ejemplo 4: Histograma

```go
numeros := []int{3, 7, 5, 8, 2, 9, 4}

for _, cantidad := range numeros {
    // Dibujar asteriscos
    for i := 0; i < cantidad; i++ {
        fmt.Print("*")
    }
    fmt.Println()
}
// Output:
// ***
// *******
// *****
// ********
// **
// *********
// ****
```

### Ejemplo 5: Procesador de datos con validación

```go
type Estudiante struct {
    Nombre   string
    Notas    []float64
}

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

func ProcesarEstudiantes(estudiantes []Estudiante) {
    for _, e := range estudiantes {
        if len(e.Notas) == 0 {
            fmt.Printf("%s: Sin notas\n", e.Nombre)
            continue
        }
        
        promedio := CalcularPromedio(e.Notas)
        var estado string
        
        if promedio >= 90 {
            estado = "Excelente"
        } else if promedio >= 80 {
            estado = "Muy bien"
        } else if promedio >= 70 {
            estado = "Bien"
        } else {
            estado = "Insuficiente"
        }
        
        fmt.Printf("%s: %.2f (%s)\n", e.Nombre, promedio, estado)
    }
}

func main() {
    estudiantes := []Estudiante{
        {"Alice", []float64{85, 90, 92}},
        {"Bob", []float64{78, 72, 80}},
        {"Charlie", []float64{95, 98, 92}},
        {"Diana", []float64{}},
    }
    
    ProcesarEstudiantes(estudiantes)
}
```

---

## Mejores Prácticas

### ✅ Bien

```go
// 1. Range para iterar (idiomático)
for i, valor := range slice {
    // ...
}

// 2. Usar _ si ignoras algo
for _, valor := range slice {  // ignoro el índice
    // ...
}

// 3. Break cuando encuentres lo que buscas
for i, v := range numeros {
    if v == objetivo {
        break
    }
}

// 4. Nombrar bien las variables
for _, estudiante := range estudiantes {
    procesar(estudiante)
}

// 5. Acumular con strings.Builder para strings
var sb strings.Builder
for _, palabra := range palabras {
    sb.WriteString(palabra)
}

// 6. Evitar modificar slice durante iteración
for i, v := range original {
    // Copia a otro slice si necesitas cambiar
    nuevos = append(nuevos, algo(v))
}
```

### ❌ Mal

```go
// 1. Iteración manual antigua
for i := 0; i < len(slice); i++ {
    v := slice[i]
    // ❌ Usar range en su lugar
}

// 2. Concatenación de strings en loop
resultado := ""
for _, palabra := range palabras {
    resultado += palabra  // ❌ Ineficiente
}

// 3. Modificar slice mientras iteras
for i := range slice {
    if condicion {
        slice = append(slice, nuevoItem)  // ❌ Comportamiento impredecible
    }
}

// 4. Nesting excesivo
for i := range a {
    for j := range b {
        for k := range c {
            // ❌ 3 niveles es mucho
        }
    }
}

// 5. Olvidar que map es aleatorio
for k := range map {
    // ❌ Orden no determinista
}
```

---

## Tabla de Comparación

| Tipo | Uso |
|------|-----|
| `for i := 0; i < n; i++` | Cuando necesitas el índice exacto |
| `for i < n {}` | Bucle while condicional |
| `for {}` | Bucle infinito (con break) |
| `for i, v := range slice` | Iterar sobre slice (lo común) |
| `for _, v := range slice` | Iterar ignorando índice |
| `for k, v := range map` | Iterar sobre map |
| `for _, char := range string` | Iterar caracteres |

---

## Conclusión

Los bucles en Go:
- **Una sola palabra clave**: `for` para todo
- **Range es lo común**: `for i, v := range colección`
- **Break/continue**: para control fino
- **Anidación limitada**: Si necesitas 3+ niveles, redeseña

Domina `for` y podrás procesar datos eficientemente.
