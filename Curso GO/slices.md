# Tipo de dato: slice

## Definición
Un `slice` en Go es una vista dinámica sobre un array subyacente. Tiene longitud y capacidad variables, y permite manipular colecciones flexibles sin definir un tamaño fijo.

## Para qué sirve
- Almacenar listas de valores de tamaño variable.
- Trabajar con colecciones que crecen o se reducen en tiempo de ejecución.
- Pasar subconjuntos de arrays sin copiar todos los datos.

## Cuándo se usa
- En la mayoría de los casos cuando necesitas una colección de elementos.
- Cuando no conoces el tamaño exacto de los datos al compilar.
- Para construir arreglos dinámicos, buffers y listas de elementos.

## Características importantes
- Un `slice` contiene un puntero al array subyacente, una longitud y una capacidad.
- Al crecer con `append`, Go puede asignar un nuevo array si la capacidad se excede.
- Copiar un `slice` copia la estructura, pero no los datos subyacentes (a menos que se use `copy`).

## Funciones útiles
- `len(s)` : devuelve la longitud del slice.
- `cap(s)` : devuelve la capacidad del slice.
- `append(s, elems...)` : agrega elementos al slice.
- `append(s, otherSlice...)` : concatena con otro slice.
- `copy(dest, src)` : copia elementos de un slice a otro.
- `make([]T, len, cap)` : crea un slice con longitud y capacidad definidas.
- `s[i:j]` : extrae un sub-slice.
- `s[i:]` : crea un sub-slice desde `i` hasta el final.
- `s[:j]` : crea un sub-slice desde el inicio hasta `j`.
- `s[:0]` : vacía el slice manteniendo la capacidad.
- `s = s[:len(s)-1]` : elimina el último elemento.
- `s = append(s[:i], s[i+1:]...)` : elimina el elemento en la posición `i`.
- `append([]T{}, s...)` : copia un slice a uno nuevo.
- `sort.Ints(s)` : ordena un slice de enteros.
- `sort.Strings(s)` : ordena un slice de strings.
- `sort.Slice(s, func(i, j int) bool)` : ordena con comparación personalizada.
- `strings.Join(slice, sep)` : convierte un slice de strings en un string.
- `range` : recorre todos los elementos.
- `reflect.DeepEqual(a, b)` : compara slices completos.
- `sort.SearchInts(s, x)` : busca un valor en un slice ordenado.
- `sort.SearchStrings(s, x)` : busca un string en un slice ordenado.
- `append([]T{x}, s...)` : inserta un elemento al inicio.
- `s = s[:0]` : conserva la capacidad y vacía los datos.
- `s = append(s[:i], s[i+1:]...)` : elimina un elemento en la posición `i`.

## Ejemplos de funciones
```go
package main

import (
    "fmt"
    "reflect"
    "sort"
    "strings"
)

func main() {
    frutas := []string{"manzana", "pera", "uva"}
    fmt.Println("Slice:", frutas)
    fmt.Println("len:", len(frutas))
    fmt.Println("cap:", cap(frutas))

    frutas = append(frutas, "banana", "kiwi")
    fmt.Println("append:", frutas)
    fmt.Println("len después de append:", len(frutas))
    fmt.Println("cap después de append:", cap(frutas))

    otro := []string{"melón", "limón"}
    combinado := append(frutas, otro...)
    fmt.Println("concat slice:", combinado)

    sub := frutas[1:4]
    fmt.Println("sub-slice:", sub)

    copia := make([]string, len(frutas))
    copy(copia, frutas)
    fmt.Println("copy:", copia)

    copia2 := append([]string{}, frutas...)
    fmt.Println("copy con append:", copia2)

    // Insertar elemento al inicio
    frutas = append([]string{"fresa"}, frutas...)
    fmt.Println("Insertar al inicio:", frutas)

    // Eliminar elemento en la posición 2
    frutas = append(frutas[:2], frutas[3:]...)
    fmt.Println("eliminar índice 2:", frutas)

    // Vaciar el slice conservando la capacidad
    frutas = frutas[:0]
    fmt.Println("vaciar slice:", frutas, "capacidad:", cap(frutas))

    ordenados := []int{3, 1, 4, 2}
    sort.Ints(ordenados)
    fmt.Println("sort.Ints:", ordenados)

    palabras := []string{"alpha", "beta", "zeta"}
    sort.Strings(palabras)
    fmt.Println("sort.Strings:", palabras)

    idx := sort.SearchStrings(palabras, "beta")
    fmt.Println("SearchStrings 'beta':", idx)

    sort.Slice(palabras, func(i, j int) bool {
        return palabras[i] < palabras[j]
    })
    fmt.Println("sort.Slice:", palabras)

    fmt.Println("Join:", strings.Join(palabras, ", "))
    fmt.Println("Equal slices:", reflect.DeepEqual(copia, copia2))
}
```

## Ejemplo
```go
package main

import "fmt"

func main() {
    frutas := []string{"manzana", "pera", "uva"}
    fmt.Println(frutas)
    fmt.Println("Longitud:", len(frutas))
    fmt.Println("Capacidad:", cap(frutas))

    frutas = append(frutas, "banana")
    fmt.Println(frutas)

    sub := frutas[1:3]
    fmt.Println("Sub-slice:", sub)

    copia := make([]string, len(frutas))
    copy(copia, frutas)
    fmt.Println("Copia:", copia)
}
```

## Buenas prácticas
- Usa `append` para agregar elementos.
- Usa `make([]T, len, cap)` cuando quieras reservar capacidad.
- Ten cuidado: cambiar elementos en un sub-slice también puede cambiar el array original.
- Usa `copy` si necesitas duplicar los datos en un slice independiente.
