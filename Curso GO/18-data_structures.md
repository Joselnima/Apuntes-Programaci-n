# Estructuras de Datos en Go - Guía Completa

## ¿Qué es una estructura de datos?

Una **estructura de datos** es una forma especializada de organizar y almacenar datos en la memoria para permitir operaciones eficientes. Piensa en ello como un contenedor con una forma particular:

- **Array**: Como una fila de cajas numeradas, todas del mismo tamaño
- **Map**: Como un diccionario donde buscas por clave
- **Pila**: Como una pila de platos, donde sacas el de arriba
- **Cola**: Como una fila de personas, donde entra al final y sale del inicio

Cada estructura tiene **ventajas y desventajas**:
- ¿Cuán rápido es acceder a un elemento?
- ¿Cuán rápido es buscar algo?
- ¿Cuánta memoria ocupa?
- ¿Puedo modificarla fácilmente?

En Go, tenemos estructuras **integradas** (arrays, slices, maps) y podemos **crear las nuestras propias** (listas enlazadas, árboles, grafos).

## 1. Arrays - La estructura más básica

Un **array** es una colección de elementos **del mismo tipo** con **tamaño fijo**. Imagina una fila de casillas numeradas desde 0, todas del mismo tamaño.

### ¿Cómo funciona un array?

```
Array de 5 números:
┌────┬────┬────┬────┬────┐
│ 10 │ 20 │ 30 │ 40 │ 50 │
└────┴────┴────┴────┴────┘
  0    1    2    3    4
```

- **Índice**: Cada posición tiene un número único (0, 1, 2...)
- **Acceso directo O(1)**: Acceder a `arr[3]` es instantáneo. El programa sabe exactamente dónde está en memoria

### Características
- **Tamaño fijo**: Definido en compilación, no pueden crecer
- **Acceso O(1)**: Obtener un elemento por índice es muy rápido
- **Memoria contigua**: Todos los elementos están juntos en memoria
- **Homogéneo**: Todos los elementos son del mismo tipo

### Declaración y uso detallada
```go
package main

import "fmt"

func main() {
    // Forma 1: Declarar vacío (todos con valor por defecto: 0)
    var arr [5]int
    arr[0] = 10
    arr[1] = 20
    fmt.Println(arr)  // [10 20 0 0 0]
    
    // Forma 2: Inicializar con valores específicos
    numeros := [3]int{1, 2, 3}
    fmt.Println(numeros)  // [1 2 3]
    
    // Forma 3: Tamaño inferido (cuenta los elementos)
    datos := [...]string{"Go", "Rust", "Python"}
    fmt.Println(datos)  // [Go Rust Python]
    
    // Acceso a elementos
    fmt.Println("Primer elemento:", numeros[0])
    fmt.Println("Último elemento:", numeros[2])
    
    // Longitud
    fmt.Println("Longitud:", len(numeros))  // 3
    
    // Iteración 1: Por índice
    for i := 0; i < len(numeros); i++ {
        fmt.Println("Índice", i, "->", numeros[i])
    }
    
    // Iteración 2: Con "range" (recomendado, más simple)
    for i, v := range numeros {
        fmt.Println("Posición", i, "=", v)
    }
}
```

### ¿Por qué es rápido acceder por índice?

Cuando declares `arr [5]int`, GO reserva memoria para 5 números enteros consecutivamente. Si cada entero ocupa 8 bytes:

```
Dirección de memoria:
1000  1008  1016  1024  1032
[arr[0]] [arr[1]] [arr[2]] [arr[3]] [arr[4]]
```

Para obtener `arr[3]`, el programa sólo calcula: `dirección_base + 3*tamaño_elemento`

### Complejidad
- **Acceso**: O(1) - Acceder a cualquier elemento es instantáneo
- **Búsqueda**: O(n) - Necesitas revisar todos los elementos
- **Inserción**: O(n) - Necesitas mover elementos (además, arrays no cambian tamaño)
- **Eliminación**: O(n) - Necesitas mover elementos

### Limitaciones de arrays
```go
// ❌ Esto NO funciona en Go
arr := [5]int{1, 2, 3}
arr = append(arr, 4)  // ERROR: no puedo crecer

// ❌ El tamaño debe ser conocido en compilación
n := 5
arr := [n]int{}  // ERROR: n no es una constante
```

Para datos variables, usa **Slices** (siguiente sección).

---

## 2. Slices - Arrays dinámicos (lo que más usarás)

Un **slice** es como un "array flexible". Internamente usa un array, pero te permite cambiar el tamaño.

### ¿Cómo funciona un slice?

Un slice tiene tres componentes:
- **Puntero**: Dónde comienza en memoria
- **Longitud (len)**: Cuántos elementos tiene ahora
- **Capacidad (cap)**: Cuántos elementos *podrían* almacenarse

```
Slice {puntero, len=3, cap=5}
         ↓
Array subyacente: [10, 20, 30, ?, ?]
```

### Características principales
- **Tamaño dinámico**: Puedes crecer con `append`
- **Flexible**: Usar la mayoría de programas Go usan slices, no arrays
- **Paso por referencia**: Los cambios afectan todas las referencias
- **Capacidad automática**: Go reserva espacio extra para futuros elementos

### Declaración y uso detallada

```go
package main

import "fmt"

func main() {
    // Forma 1: Declaración vacía (nil)
    var s []int
    fmt.Println("Slice vacío:", s, "len:", len(s), "cap:", cap(s))
    
    // Forma 2: Crear con make (tamaño inicial)
    slice1 := make([]int, 5)        // 5 elementos, capacidad 5
    fmt.Println("Con make:", slice1)  // [0 0 0 0 0]
    
    // Forma 3: Crear con make (tamaño y capacidad)
    slice2 := make([]int, 3, 10)    // 3 elementos visible, pero 10 de capacidad
    fmt.Println("len:", len(slice2), "cap:", cap(slice2))
    
    // Forma 4: Literal (como array pero sin [número])
    slice3 := []int{1, 2, 3, 4, 5}
    fmt.Println("Literal:", slice3)
    
    // Forma 5: Sub-slice (corte de un array)
    arr := [5]int{10, 20, 30, 40, 50}
    subSlice := arr[1:4]            // Elementos en índices 1,2,3
    fmt.Println("Sub-slice:", subSlice)  // [20 30 40]
    
    // IMPORTANTE: Sub-slice comparte memoria con el array original
    subSlice[0] = 999
    fmt.Println("Array modificado:", arr)  // [10 999 30 40 50]
}
```

### append - Crecimiento dinámico

`append` es la función mágica que hace crecer el slice:

```go
package main

import "fmt"

func main() {
    s := make([]int, 0, 3)  // len=0, cap=3 (reserva espacio para 3)
    fmt.Println("Inicio: len:", len(s), "cap:", cap(s))
    
    s = append(s, 10)
    fmt.Println("Después de 1: len:", len(s), "cap:", cap(s))
    
    s = append(s, 20)
    fmt.Println("Después de 2: len:", len(s), "cap:", cap(s))
    
    s = append(s, 30)
    fmt.Println("Después de 3: len:", len(s), "cap:", cap(s))
    
    // Ahora la capacidad se agota, Go crea un array más grande
    s = append(s, 40)
    fmt.Println("Después de 4: len:", len(s), "cap:", cap(s))  // cap=6 (duplica)
    
    // Puedes agregar múltiples elementos de una vez
    s = append(s, 50, 60, 70)
    fmt.Println("Final:", s)
}
```

Cuando `len == cap` y haces `append`, Go crea un **nuevo array más grande** (usualmente el doble), copia todos los elementos, y actualiza el puntero.

### copy - Copiar datos entre slices

```go
package main

import "fmt"

func main() {
    original := []int{1, 2, 3, 4, 5}
    
    // Copiar a un nuevo slice
    copia := make([]int, len(original))
    copy(copia, original)
    
    // Modificar la copia NO afecta el original
    copia[0] = 999
    fmt.Println("Original:", original)  // [1 2 3 4 5]
    fmt.Println("Copia:", copia)        // [999 2 3 4 5]
}
```

### Operaciones útiles con el paquete "slices"

```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    s := []int{3, 1, 4, 1, 5, 9, 2, 6}
    
    // Ordenar
    slices.Sort(s)
    fmt.Println("Ordenado:", s)  // [1 1 2 3 4 5 6 9]
    
    // Buscar índice de un elemento
    idx := slices.Index(s, 4)
    fmt.Println("Índice de 4:", idx)  // 4
    
    // ¿Contiene un elemento?
    tiene := slices.Contains(s, 3)
    fmt.Println("¿Contiene 3?:", tiene)  // true
    
    // Clonar (copia profunda)
    clon := slices.Clone(s)
    clon[0] = 999
    fmt.Println("Original después:", s)  // Sin cambios
    
    // Invertir
    slices.Reverse(s)
    fmt.Println("Invertido:", s)
    
    // Comparar
    s2 := []int{1, 1, 2, 3, 4, 5, 6, 9}
    fmt.Println("¿Iguales?:", slices.Equal(s, s2))
}
```

### Insertar y eliminar elementos

Como los slices no tiene inserción/eliminación nativa, necesitas hacer pequeños trucos:

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 4, 5}
    
    // Insertar 3 en la posición 2
    indice := 2
    s = append(s[:indice], append([]int{3}, s[indice:]...)...)
    fmt.Println("Después de insertar:", s)  // [1 2 3 4 5]
    
    // Eliminar elemento en una posición
    s = append(s[:2], s[3:]...)  // Elimina índice 2
    fmt.Println("Después de eliminar:", s)  // [1 2 4 5]
}
```

### Complejidad
- **Acceso**: O(1) - Igual que arrays
- **Append** (amortizado): O(1) - La mayoría de veces es instantáneo
- **Búsqueda**: O(n) - Necesita revisar elementos
- **Inserción media**: O(n) - Requiere mover elementos
- **Eliminación**: O(n) - Requiere mover elementos

### ¿Cuándo usar slices?
✅ **Casi siempre** - Son más flexibles que arrays
❌ No uses cuando necesites tamaño fijo constantemente o memoria fija predecible

---

## 3. Maps - Búsqueda rápida por clave (diccionarios/hash tables)

Un **map** es como un diccionario: asocia una **clave** con un **valor**. En lugar de buscar por posición (índice), buscas por nombre.

### ¿Cómo funciona un map?

```
Map[string]int
┌─────────┬────────┐
│ "Ana"   │ 30     │
│ "Bob"   │ 25     │
│ "Carlos"│ 35     │
└─────────┴────────┘
```

Internamente usa una **tabla hash**: convierte la clave a una posición de memoria usando una función hash.

### Características
- **Búsqueda O(1)**: Acceso ultra rápido al valor si conoces la clave
- **Clave única**: No puede haber dos entradas con la misma clave
- **Sin orden garantizado**: Cuando iteras, el orden es aleatorio cada vez
- **Dinámico**: Crece automáticamente
- **Paso por referencia**: Los cambios se ven en todas partes

### Declaración y uso detallada

```go
package main

import "fmt"

func main() {
    // Forma 1: Declarar vacío (nil)
    var m map[string]int
    // ❌ No puedes usar sin inicializar
    // m["Ana"] = 30  // ERROR
    
    // Forma 2: Inicializar con make
    edades := make(map[string]int)
    edades["Ana"] = 30
    edades["Bob"] = 25
    edades["Carlos"] = 35
    
    fmt.Println("Map:", edades)
    fmt.Println("Edad de Ana:", edades["Ana"])  // 30
    
    // Forma 3: Literal (inicializar con valores)
    ciudades := map[string]string{
        "ES": "Madrid",
        "FR": "París",
        "IT": "Roma",
    }
    fmt.Println("Ciudades:", ciudades)
    
    // Acceso a valor inexistente
    fmt.Println("Edad de David:", edades["David"])  // 0 (valor por defecto)
    
    // Verificar si existe una clave
    edad, existe := edades["David"]
    fmt.Println("David:", edad, "existe:", existe)  // 0, false
    
    edad2, existe2 := edades["Ana"]
    fmt.Println("Ana:", edad2, "existe:", existe2)  // 30, true
}
```

### Modificar un map

```go
package main

import "fmt"

func main() {
    personas := make(map[string]int)
    
    // Agregar
    personas["Alice"] = 25
    personas["Bob"] = 30
    fmt.Println("Inicio:", personas)
    
    // Actualizar (la clave ya existe)
    personas["Alice"] = 26
    fmt.Println("Después de actualizar:", personas)
    
    // Eliminar una clave
    delete(personas, "Bob")
    fmt.Println("Después de eliminar:", personas)
    
    // ¿Cuántas entradas?
    fmt.Println("Cantidad de personas:", len(personas))  // 1
    
    // Limpiar completamente
    personas = make(map[string]int)  // "Reinicia" el map
    fmt.Println("Map limpio:", personas)  // map[]
}
```

### Iterar sobre un map

```go
package main

import "fmt"

func main() {
    calificaciones := map[string]int{
        "Math":    95,
        "Physics": 87,
        "Chemistry": 92,
    }
    
    // Iterar sobre clave y valor
    for materia, calificacion := range calificaciones {
        fmt.Println(materia, ":", calificacion)
    }
    // Nota: El orden es aleatorio cada ejecución
    
    // Solo claves
    for materia := range calificaciones {
        fmt.Println("Materia:", materia)
    }
    
    // Solo valores (raro, pero posible)
    for _, calificacion := range calificaciones {
        fmt.Println("Calificación:", calificacion)
    }
}
```

### Casos de uso reales

```go
package main

import "fmt"

func main() {
    // Caso 1: Contar ocurrencias de palabras
    texto := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
    conteo := make(map[string]int)
    
    for _, palabra := range texto {
        conteo[palabra]++
    }
    
    fmt.Println("Conteo:", conteo)
    // map[apple:3 banana:2 cherry:1]
    
    // Caso 2: Caché de datos
    cache := map[string]string{
        "key1": "value1",
        "key2": "value2",
    }
    
    if valor, existe := cache["key1"]; existe {
        fmt.Println("Encontrado en caché:", valor)
    }
    
    // Caso 3: Traducción de códigos
    codigoPais := map[string]string{
        "ES": "España",
        "FR": "Francia",
        "IT": "Italia",
        "DE": "Alemania",
    }
    
    codigo := "ES"
    pais, ok := codigoPais[codigo]
    if ok {
        fmt.Println(codigo, "->", pais)
    }
}
```

### Map de structs

```go
package main

import "fmt"

type Persona struct {
    Edad int
    Ciudad string
}

func main() {
    personas := map[string]Persona{
        "Alice": {Edad: 30, Ciudad: "Madrid"},
        "Bob": {Edad: 25, Ciudad: "Barcelona"},
    }
    
    fmt.Println(personas["Alice"])  // {30 Madrid}
    fmt.Println(personas["Alice"].Edad)  // 30
}
```

### Complejidad
- **Acceso**: O(1) promedio - Casi instantáneo
- **Inserción**: O(1) promedio - Muy rápido
- **Eliminación**: O(1) promedio - Muy rápido
- **Búsqueda**: O(1) promedio - Búsqueda directa por clave
- **Iteración**: O(n) - Necesita revisar todas las entradas

### ¿Cuándo usar maps?
✅ Cuando necesitas buscar por **clave**
✅ Contadores, cachés, asociaciones
✅ Diccionarios, traducciones
❌ Cuando necesitas **orden** (usa slices ordenados)
❌ Cuando necesitas duplicados (usa slices)

---

## 4. Structs - Agrupar datos relacionados

Una **struct** (estructura) agrupa datos relacionados de **diferentes tipos** bajo un mismo nombre. Es como crear tu propio tipo de dato.

### ¿Para qué sirven?

En lugar de:
```go
nombre := "Alice"
edad := 30
ciudad := "Madrid"
```

Haces:
```go
type Persona struct {
    Nombre string
    Edad int
    Ciudad string
}

p := Persona{Nombre: "Alice", Edad: 30, Ciudad: "Madrid"}
```

Mucho más organizado y claro.

### Características
- **Tipado estáticamente**: Cada campo tiene un tipo definido
- **Campos nombrados**: Los datos son accesibles por nombre, no por posición
- **Composición**: Puede contener slices, maps, otras structs
- **Métodos**: Puedes asociar funciones a structs

### Declaración y uso detallada

```go
package main

import "fmt"

// Definir un tipo de estructura
type Persona struct {
    Nombre string
    Edad   int
    Ciudad string
}

func main() {
    // Forma 1: Con nombres de campo (recomendado)
    p1 := Persona{
        Nombre: "Alice",
        Edad:   30,
        Ciudad: "Madrid",
    }
    fmt.Println(p1)  // {Alice 30 Madrid}
    
    // Forma 2: Por posición (fácil de cometer errores)
    p2 := Persona{"Bob", 25, "Barcelona"}
    
    // Forma 3: Parcial (otros campos son valor por defecto)
    p3 := Persona{Nombre: "Carlos"}
    fmt.Println(p3)  // {Carlos 0 }
    
    // Acceso a campos
    fmt.Println("Nombre:", p1.Nombre)
    fmt.Println("Edad:", p1.Edad)
    
    // Modificar campos
    p1.Edad = 31
    fmt.Println("Edad actualizada:", p1.Edad)
    
    // Comparar structs (deben tener los mismos valores)
    p4 := Persona{"Alice", 30, "Madrid"}
    fmt.Println("¿Iguales?:", p1 == p4)  // false (porque cambié edad)
}
```

### Structs anidadas (composición)

```go
package main

import "fmt"

type Dirección struct {
    Calle string
    Ciudad string
    Código string
}

type Persona struct {
    Nombre string
    Edad int
    Dirección Dirección  // Struct dentro de struct
}

type Empresa struct {
    Nombre string
    Director Persona
    Empleados []Persona  // Slice de structs
}

func main() {
    // Crear con anidamiento
    dir := Dirección{
        Calle: "Calle Principal 123",
        Ciudad: "Madrid",
        Código: "28001",
    }
    
    p := Persona{
        Nombre: "Alice",
        Edad: 30,
        Dirección: dir,
    }
    
    fmt.Println("Nombre:", p.Nombre)
    fmt.Println("Ciudad:", p.Dirección.Ciudad)  // Acceso anidado
    
    // Crear empresa
    emp := Empresa{
        Nombre: "TechCorp",
        Director: p,
        Empleados: []Persona{p},
    }
    
    fmt.Println("Director:", emp.Director.Nombre)
    fmt.Println("Empleados:", len(emp.Empleados))
}
```

### Métodos en structs (funciones asociadas)

Un **método** es una función asociada a un tipo. Un receptor indica a qué tipo pertenece.

```go
package main

import (
    "fmt"
    "math"
)

type Punto struct {
    X, Y float64
}

// Método con RECEPTOR POR VALOR (copia)
// No modifica el original
func (p Punto) Distancia() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Método con RECEPTOR POR PUNTERO
// Puede modificar el original
func (p *Punto) Trasladar(dx, dy float64) {
    p.X += dx
    p.Y += dy
}

// Otro método
func (p Punto) String() string {
    return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

func main() {
    punto := Punto{3, 4}
    
    // Llamar método (receptor por valor)
    distancia := punto.Distancia()
    fmt.Println("Distancia al origen:", distancia)  // 5
    
    // Llamar método (receptor por puntero)
    punto.Trasladar(1, 2)
    fmt.Println("Punto trasladado:", punto)  // (4, 6)
    
    // String() se llama automáticamente en fmt.Println
    fmt.Println("Punto:", punto)  // (4.00, 6.00)
}
```

### ¿Receptor por valor o puntero?

```go
// Usa VALOR cuando:
// - No necesitas modificar el struct
// - Es pequeño (pocos campos)
func (p Persona) SaludarString str {
    return "Hola, " + p.Nombre
}

// Usa PUNTERO cuando:
// - Necesitas modificar el struct
// - Es grande (muchos campos, copy sería caro)
// - Es más eficiente (no hay copia)
func (p *Persona) CumplirAños() {
    p.Edad++
}
```

### Structs vacías y de utilidad

```go
// Struct vacío (sin campos)
// Útil para serialización, markers
type Marca struct{}

// Usar como clave en map (ocupa 0 bytes)
m := map[string]Marca{}
m["key1"] = Marca{}

// Usar para canalización (channels)
ch := make(chan Marca)
ch <- Marca{}
```

### Complejidad
- **Acceso a campo**: O(1) - Directo
- **Crear instancia**: O(1) - Copia los datos
- **Métodos**: Depende de lo que hagan

### ¿Cuándo usar structs?
✅ Agrupar datos relacionados
✅ Cuando necesitas lógica asociada (métodos)
✅ Pasar múltiples valores juntos a funciones
❌ Cuando solo necesitas arrays (usa slices)

---

## 5. Interfaces - Contratos de métodos

Una **interface** define un **contrato**: qué métodos debe implementar un tipo. Es como decir: "si tu tipo implementa estos métodos, puedes usarse como interface X".

### ¿Para qué sirven?

Permiten **polimorfismo**: el mismo código puede trabajar con **diferentes tipos** que implementen los mismos métodos.

### Características
- **Implementación implícita**: No necesitas decir "implemento esta interface"
- **Polimorfismo**: El mismo código funciona con múltiples tipos
- **Flexibilidad**: Cambiar implementaciones sin modificar código
- `interface{}`: Acepta **cualquier tipo**

### Ejemplo del mundo real

```go
package main

import "fmt"

// Definir una interfaz
type Animal interface {
    Sonido() string
    Nombre() string
}

// Tipo Perro
type Perro struct {
    nombre string
}

func (p Perro) Sonido() string {
    return "¡Guau!"
}

func (p Perro) Nombre() string {
    return p.nombre
}

// Tipo Gato
type Gato struct {
    nombre string
}

func (g Gato) Sonido() string {
    return "¡Miau!"
}

func (g Gato) Nombre() string {
    return g.nombre
}

// Función que trabaja con CUALQUIER Animal
func HacerSonido(a Animal) {
    fmt.Println(a.Nombre(), "hace:", a.Sonido())
}

func main() {
    // Ambos tipos implementan Animal (sin decirlo explícitamente)
    perro := Perro{"Rex"}
    gato := Gato{"Félix"}
    
    // La misma función funciona con ambos
    HacerSonido(perro)  // Rex hace: ¡Guau!
    HacerSonido(gato)   // Félix hace: ¡Miau!
    
    // Puedes crear un slice de la interface
    animales := []Animal{perro, gato}
    for _, a := range animales {
        HacerSonido(a)
    }
}
```

### Interface vacía `interface{}`

Acepta **cualquier tipo**:

```go
package main

import "fmt"

func Imprimir(v interface{}) {
    fmt.Println("Valor:", v)
}

func main() {
    Imprimir(42)           // int
    Imprimir("Hola")       // string
    Imprimir(3.14)         // float64
    Imprimir([]int{1,2,3}) // slice
}
```

### Type assertion - Descubrir el tipo real

```go
package main

import "fmt"

func main() {
    var i interface{} = "Hola"
    
    // Type assertion 1: Con verificación
    s, ok := i.(string)
    if ok {
        fmt.Println("Es string:", s)
    }
    
    // Falla sin verificación
    //n := i.(int)  // ❌ PANIC si no es int
    
    // Verificar different tipos
    numero, ok := i.(int)
    if !ok {
        fmt.Println("No es int, es:", fmt.Sprintf("%T", i))
    }
}
```

### Type switch - Múltiples tipos

```go
package main

import "fmt"

func Procesar(v interface{}) {
    switch valor := v.(type) {
    case string:
        fmt.Println("String:", valor)
        fmt.Println("Longitud:", len(valor))
    case int:
        fmt.Println("Entero:", valor)
        fmt.Println("Cuadrado:", valor*valor)
    case []int:
        fmt.Println("Slice de ints:", valor)
        fmt.Println("Suma:", suma(valor))
    default:
        fmt.Println("Tipo desconocido:", fmt.Sprintf("%T", v))
    }
}

func suma(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    Procesar("Hola")      // String: Hola...
    Procesar(5)           // Entero: 5...
    Procesar([]int{1,2,3})  // Slice de ints: [1 2 3]...
    Procesar(3.14)        // Tipo desconocido: float64
}
```

### Interfaces útiles del estándar

Go proporciona interfaces comunes:

```go
// Reader - Leer datos
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer - Escribir datos
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Stringer - Convertir a string
type Stringer interface {
    String() string
}

// Error - Manejo de errores
type error interface {
    Error() string
}
```

Implementar `String()` hace que `fmt.Println` use tu versión:

```go
package main

import "fmt"

type Punto struct {
    X, Y float64
}

// Implementar Stringer
func (p Punto) String() string {
    return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

func main() {
    p := Punto{3, 4}
    fmt.Println(p)  // (3.0, 4.0)
}
```

### ¿Cuándo usar interfaces?
✅ Código que funciona con múltiples tipos
✅ Dependencias inyectables (pasar interfaces, no structs)
✅ Definir contratos entre módulos
❌ Cuando solo necesitas trabajar con un tipo específico (es overhead)

---

## 6. Punteros - Referencias a memoria

Un **puntero** almacena la **dirección de memoria** de una variable, no su valor. Es como tener la dirección de una casa en lugar de vivir en ella.

### ¿Por qué punteros?

1. **Passar por referencia**: Modificar variable original en función
2. **Memoria dinámica**: Crear structs grandes en montículo (heap)
3. **Estructuras recursivas**: Listas, árboles necesitan auto-referencias

### Características
- **`&` (dirección)**: Obtiene la dirección de una variable
- **`*` (desreferencia)**: Accede al valor apuntado
- **`nil`**: Puntero que no apunta a nada (equivalente a NULL)
- **Tipo `*T`**: Puntero a tipo T

### Concepto básico

```
Variable: x = 5
Dirección en memoria: 0x1234

Puntero: p := &x
p contiene: 0x1234
*p produce: 5
```

### Uso detallado

```go
package main

import "fmt"

func main() {
    // Variable original
    x := 10
    
    // Obtener dirección (crear puntero)
    p := &x
    
    fmt.Println("x:", x)              // 10
    fmt.Println("p (dirección):", p)  // 0xc00000...
    fmt.Println("*p (valor):", *p)    // 10
    
    // Modificar a través del puntero
    *p = 20
    fmt.Println("x ahora:", x)        // 20 (cambió!)
    
    // Comparar punteros
    p2 := &x
    fmt.Println("¿p == p2?:", p == p2)  // true (apuntan a lo mismo)
}
```

### Punteros en funciones

**Importante**: Go pasa por **valor por defecto**. Los cambios en la función NO afectan el original.

```go
package main

import "fmt"

// Función que INTENTA modificar (no funciona)
func IncrementarMal(x int) {
    x = x + 1  // Solo modifica la copia local
}

// Función que SÍ modifica (con puntero)
func IncrementarBien(x *int) {
    *x = *x + 1  // Modifica el original
}

func main() {
    contador := 5
    
    IncrementarMal(contador)
    fmt.Println("Después de mal:", contador)   // 5 (sin cambios)
    
    IncrementarBien(&contador)
    fmt.Println("Después de bien:", contador)  // 6 (cambió!)
}
```

### Punteros a structs

```go
package main

import "fmt"

type Persona struct {
    Nombre string
    Edad   int
}

func main() {
    // Forma 1: Crear y obtener puntero
    p := Persona{"Alice", 30}
    puntero := &p
    
    // Acceso a campos (ambas formas son equivalentes)
    fmt.Println("Nombre:", puntero.Nombre)  // Go desdeñarencia automáticamente
    fmt.Println("Nombre:", (*puntero).Nombre)  // Forma explícita (rara)
    
    // Modificar
    puntero.Edad = 31
    fmt.Println("Persona:", p)  // {Alice 31}
    
    // Forma 2: Crear directamente con puntero
    p2 := &Persona{"Bob", 25}  // p2 es tipo *Persona
    fmt.Println("Bob:", p2.Nombre)
    
    // Forma 3: Crear con new
    p3 := new(Persona)
    p3.Nombre = "Carlos"
    p3.Edad = 28
    fmt.Println("Carlos:", p3)
}
```

### Nil - Puntero vacío

```go
package main

import "fmt"

func main() {
    // Puntero sin inicializar (nil)
    var p *int
    
    fmt.Println("¿Es nil?:", p == nil)  // true
    
    // ❌ Esto causa PANIC si desreferencias nil
    // valor := *p  // PANIC!
    
    // Siempre verifica antes
    if p != nil {
        fmt.Println("Valor:", *p)
    } else {
        fmt.Println("Puntero es nil")
    }
    
    // Asignar un valor
    x := 10
    p = &x
    fmt.Println("Valor:", *p)  // 10
}
```

### Punteros vs Valores - Cuándo usar cada uno

```go
// Usar por VALOR cuando:
func (p Persona) Saludar() {  // pequeño, no modifica
    fmt.Println("Hola, soy", p.Nombre)
}

// Usar por PUNTERO cuando:
func (p *Persona) CumplirAños() {  // necesita modificar
    p.Edad++
}

func (p *PersonaGrande) Procesar() {  // struct muy grande (eficiencia)
    // ...
}
```

### Allocar memoria con make y new

```go
package main

import "fmt"

func main() {
    // new: Crea un puntero a un struct (inicializado a valores por defecto)
    p1 := new(int)
    *p1 = 42
    
    // make: Para types que necesitan inicialización interna (slices, maps, channels)
    s := make([]int, 5)  // No *[]int, solo []int
    m := make(map[string]int)
    
    fmt.Println("p1:", *p1)
    fmt.Println("s:", s)
    fmt.Println("m:", m)
}
```

### Punteros en slices

```go
package main

import "fmt"

type Persona struct {
    Nombre string
    Edad   int
}

func main() {
    // Slice de punteros a structs
    personas := []*Persona{
        &Persona{"Alice", 30},
        &Persona{"Bob", 25},
    }
    
    // Acceder y modificar
    personas[0].Edad = 31
    
    fmt.Println("Primera persona:", personas[0])
}
```

### ¿Por qué Go tiene punteros?

```go
// Sin punteros: cada función hace una copia (ineficiente)
func ProcesarGrande(persona Persona) {
    // hace copia de toda la persona
}

// Con punteros: solo la dirección (eficiente)
func ProcesarGrande(persona *Persona) {
    // solo copia la dirección (8 bytes)
}
```

### Complejidad
- **Crear puntero**: O(1)
- **Desreferencia**: O(1)
- **Copiar puntero**: O(1)

### Errores comunes

```go
❌ // Puntero sin inicializar
var p *int
fmt.Println(*p)  // PANIC

✅ // Inicializar
x := 5
p := &x
fmt.Println(*p)  // OK

❌ // Olvidar & cuando necesita puntero
func Modificar(p *int) {
    *p = 10
}
Modificar(contador)  // ERROR

✅ // Usar & para obtener dirección
Modificar(&contador)  // OK
```

---

## 7. Listas Enlazadas - Inserción eficiente al inicio

Una **lista enlazada** es una colección de **nodos** donde cada nodo tiene un **valor** y un **puntero al siguiente nodo**.

### ¿Por qué listas enlazadas?

Arrays/Slices:
- Acceso: O(1) ✅
- Inserción al inicio: O(n) ❌ (necesita mover todos)

Listas enlazadas:
- Acceso: O(n)
- Inserción al inicio: O(1) ✅

```
Array [1, 2, 3, 4] - Insertar 0 al inicio requiere mover todos
Array [0, 1, 2, 3, 4]

Lista enlazada:
1 -> 2 -> 3 -> 4 -> nil
0 -> 1 -> 2 -> 3 -> 4 -> nil (solo actualizar punteros)
```

### Estructura de nodo

```go
type Nodo struct {
    Valor int       // Los datos
    Siguiente *Nodo // Puntero al siguiente nodo
}
```

### Implementación completa

```go
package main

import "fmt"

type Nodo struct {
    Valor int
    Siguiente *Nodo
}

type ListaEnlazada struct {
    Cabeza *Nodo  // Primer elemento
}

// Insertar al INICIO
func (l *ListaEnlazada) InsertarInicio(valor int) {
    nuevoNodo := &Nodo{Valor: valor}
    nuevoNodo.Siguiente = l.Cabeza  // El nuevo apunta al anterior primero
    l.Cabeza = nuevoNodo             // El nuevo es el primero
}

// Insertar al FINAL
func (l *ListaEnlazada) InsertarFinal(valor int) {
    nuevoNodo := &Nodo{Valor: valor}
    
    // Si está vacía
    if l.Cabeza == nil {
        l.Cabeza = nuevoNodo
        return
    }
    
    // Buscar el último
    actual := l.Cabeza
    for actual.Siguiente != nil {
        actual = actual.Siguiente
    }
    
    // Conectar al final
    actual.Siguiente = nuevoNodo
}

// Buscar un valor
func (l *ListaEnlazada) Contiene(valor int) bool {
    actual := l.Cabeza
    for actual != nil {
        if actual.Valor == valor {
            return true
        }
        actual = actual.Siguiente
    }
    return false
}

// Eliminar primer elemento
func (l *ListaEnlazada) EliminarInicio() {
    if l.Cabeza != nil {
        l.Cabeza = l.Cabeza.Siguiente
    }
}

// Eliminar un valor específico
func (l *ListaEnlazada) Eliminar(valor int) {
    // Caso especial: eliminar el primero
    if l.Cabeza != nil && l.Cabeza.Valor == valor {
        l.Cabeza = l.Cabeza.Siguiente
        return
    }
    
    // Buscar y eliminar
    actual := l.Cabeza
    for actual != nil && actual.Siguiente != nil {
        if actual.Siguiente.Valor == valor {
            actual.Siguiente = actual.Siguiente.Siguiente
            return
        }
        actual = actual.Siguiente
    }
}

// Imprimir la lista
func (l *ListaEnlazada) Imprimir() {
    actual := l.Cabeza
    for actual != nil {
        fmt.Print(actual.Valor)
        if actual.Siguiente != nil {
            fmt.Print(" -> ")
        }
        actual = actual.Siguiente
    }
    fmt.Println(" -> nil")
}

// Longitud
func (l *ListaEnlazada) Longitud() int {
    count := 0
    actual := l.Cabeza
    for actual != nil {
        count++
        actual = actual.Siguiente
    }
    return count
}

func main() {
    lista := &ListaEnlazada{}
    
    // Agregar elementos
    lista.InsertarInicio(3)
    lista.InsertarInicio(2)
    lista.InsertarInicio(1)
    lista.InsertarFinal(4)
    lista.Imprimir()  // 1 -> 2 -> 3 -> 4 -> nil
    
    // Encontrar
    fmt.Println("¿Contiene 3?:", lista.Contiene(3))   // true
    fmt.Println("¿Contiene 10?:", lista.Contiene(10)) // false
    
    // Longitud
    fmt.Println("Longitud:", lista.Longitud())  // 4
    
    // Eliminar
    lista.Eliminar(2)
    lista.Imprimir()  // 1 -> 3 -> 4 -> nil
}
```

### Complejidad
- **Acceso**: O(n) - Necesita recorrer desde el inicio
- **Búsqueda**: O(n) - Buscar secuencialmente
- **Inserción al inicio**: O(1) - Solo actualizar punteros ✅
- **Inserción al final**: O(n) - Necesita encontrar el final
- **Eliminación**: O(n) - Necesita buscar

### ¿Cuándo usar listas enlazadas?
✅ Inserciones/eliminaciones frecuentes al inicio
✅ Cuando no necesitas acceso rápido por índice
❌ Cuando necesitas acceso aleatorio (usa slices)

---

## 8. Pilas (Stacks) - LIFO: Último Entra, Primero Sale

Una **pila** es como una pila de platos: agregas desde arriba (push) y sacas desde arriba (pop).

### Visualización

```
LIFO - Last In First Out

Pila: [1, 2, 3, 4]
        ↑
    (top/cima)

Push 5: [1, 2, 3, 4, 5]
Pop:    [1, 2, 3, 4]
Sacamos el 5 (el último que entramos)
```

### ¿Para qué se usan?

1. **Llamadas de funciones**: El stack de recursión
2. **Deshacer (Undo)**: Guardar acciones reversibles
3. **Expresiones matemáticas**: Evaluar `2 + 3 * 4`
4. **Parsear estructuras**: HTML, paréntesis

### Implementación con slice

```go
package main

import "fmt"

type Pila struct {
    items []int  // Almacenaremos elementos en un slice
}

// Agregar elemento al top
func (p *Pila) Push(valor int) {
    p.items = append(p.items, valor)
}

// Sacar del top
func (p *Pila) Pop() (int, bool) {
    if len(p.items) == 0 {
        return 0, false  // Pila vacía
    }
    // Último elemento
    ultimo := p.items[len(p.items)-1]
    // Eliminar el último
    p.items = p.items[:len(p.items)-1]
    return ultimo, true
}

// Ver el top sin sacarlo
func (p *Pila) Peek() (int, bool) {
    if len(p.items) == 0 {
        return 0, false
    }
    return p.items[len(p.items)-1], true
}

// ¿Está vacía?
func (p *Pila) Vacia() bool {
    return len(p.items) == 0
}

// Tamaño
func (p *Pila) Tamaño() int {
    return len(p.items)
}

func main() {
    pila := &Pila{}
    
    fmt.Println("¿Vacía?:", pila.Vacia())  // true
    
    // Agregar elementos
    pila.Push(10)
    pila.Push(20)
    pila.Push(30)
    pila.Push(40)
    
    fmt.Println("Tamaño:", pila.Tamaño())  // 4
    
    // Ver el top
    top, _ := pila.Peek()
    fmt.Println("Top:", top)  // 40
    
    // Sacar elementos (LIFO)
    for !pila.Vacia() {
        v, _ := pila.Pop()
        fmt.Println("Pop:", v)  // 40 30 20 10
    }
}
```

### Caso práctico: Verificar paréntesis balanceados

```go
package main

import "fmt"

type Pila struct {
    items []rune
}

func (p *Pila) Push(v rune) {
    p.items = append(p.items, v)
}

func (p *Pila) Pop() (rune, bool) {
    if len(p.items) == 0 {
        return 0, false
    }
    v := p.items[len(p.items)-1]
    p.items = p.items[:len(p.items)-1]
    return v, true
}

func EsBalanceado(expr string) bool {
    pila := &Pila{}
    
    for _, char := range expr {
        if char == '(' || char == '[' || char == '{' {
            pila.Push(char)
        } else if char == ')' || char == ']' || char == '}' {
            top, existe := pila.Pop()
            if !existe {
                return false  // Más cierres que aperturas
            }
            // Verificar coincidencia
            if (char == ')' && top != '(') ||
               (char == ']' && top != '[') ||
               (char == '}' && top != '{') {
                return false
            }
        }
    }
    
    return len(pila.items) == 0  // No sobran aperturas
}

func main() {
    fmt.Println(EsBalanceado("(a+b)"))           // true
    fmt.Println(EsBalanceado("(a+b]"))           // false
    fmt.Println(EsBalanceado("((a+b)"))          // false
    fmt.Println(EsBalanceado("(a+[b*c])"))       // true
}
```

### Complejidad
- **Push**: O(1) amortizado
- **Pop**: O(1)
- **Peek**: O(1)
- **Búsqueda**: O(n)

### ¿Cuándo usar pilas?
✅ Recursión y backtracking
✅ Undo/Redo
✅ Evaluación de expresiones
✅ Balanceo de paréntesis

---

## 9. Colas (Queues) - FIFO: Primero Entra, Primero Sale

Una **cola** es como una fila: agregas al final y sacas del inicio.

### Visualización

```
FIFO - First In First Out

Cola: [1, 2, 3, 4]
       ↑           ↑
     entra      sale
     (final)    (inicio)

Encolar 5: [1, 2, 3, 4, 5]
Desencolar: [2, 3, 4, 5]
Sale el 1 (el primero que entró)
```

### ¿Para qué se usan?

1. **Procesos**: Sistema operativo ejecuta de forma FIFO
2. **Impresoras**: Cola de trabajos
3. **Servers**: Procesando requests del usuario
4. **Búsqueda BFS**: Gráfos, niveles

### Implementación básica

```go
package main

import "fmt"

type Cola struct {
    items []int
}

// Encolar (agregar al final)
func (c *Cola) Encolar(valor int) {
    c.items = append(c.items, valor)
}

// Desencolar (sacar del inicio)
func (c *Cola) Desencolar() (int, bool) {
    if len(c.items) == 0 {
        return 0, false
    }
    primero := c.items[0]
    c.items = c.items[1:]  // Quitar el primero
    return primero, true
}

// Ver el primero sin sacar
func (c *Cola) Frente() (int, bool) {
    if len(c.items) == 0 {
        return 0, false
    }
    return c.items[0], true
}

// ¿Está vacía?
func (c *Cola) Vacia() bool {
    return len(c.items) == 0
}

// Tamaño
func (c *Cola) Tamaño() int {
    return len(c.items)
}

func main() {
    cola := &Cola{}
    
    // Agregar elementos
    cola.Encolar(10)
    cola.Encolar(20)
    cola.Encolar(30)
    cola.Encolar(40)
    
    fmt.Println("Tamaño:", cola.Tamaño())  // 4
    
    // Ver el frente
    frente, _ := cola.Frente()
    fmt.Println("Frente:", frente)  // 10
    
    // Desencolar elementos (FIFO)
    for !cola.Vacia() {
        v, _ := cola.Desencolar()
        fmt.Println("Desencolar:", v)  // 10 20 30 40
    }
}
```

### Cola eficiente con índices

El desencolar con `c.items[1:]` es O(n). Para mejor rendimiento:

```go
package main

import "fmt"

type ColaEficiente struct {
    items []int
    inicio int  // Índice del primer elemento
    fin    int  // Índice del último + 1
}

func (c *ColaEficiente) Encolar(valor int) {
    // Si alcanzamos capacidad, redimensionar
    if c.fin >= len(c.items) {
        if c.inicio > 0 {
            // Compactar
            copy(c.items, c.items[c.inicio:c.fin])
            c.fin -= c.inicio
            c.inicio = 0
        } else {
            // Crecer
            nuevos := make([]int, len(c.items)*2)
            copy(nuevos, c.items)
            c.items = nuevos
        }
    }
    c.items[c.fin] = valor
    c.fin++
}

func (c *ColaEficiente) Desencolar() (int, bool) {
    if c.inicio >= c.fin {
        return 0, false
    }
    valor := c.items[c.inicio]
    c.inicio++
    return valor, true
}

func main() {
    cola := &ColaEficiente{items: make([]int, 10)}
    
    cola.Encolar(1)
    cola.Encolar(2)
    cola.Encolar(3)
    
    v1, _ := cola.Desencolar()
    v2, _ := cola.Desencolar()
    
    fmt.Println("Sacados:", v1, v2)  // 1 2
}
```

### Complejidad
- **Encolar**: O(1) amortizado (con implementación eficiente)
- **Desencolar**: O(1) con índices (O(n) con slice simple)
- **Búsqueda**: O(n)

### ¿Cuándo usar colas?
✅ Procesamiento secuencial
✅ BFS (breadth-first search)
✅ Tareas en sistema operativo
✅ Simulación de procesos

---

## 10. Conjuntos (Sets) - Valores únicos sin orden

Un **conjunto** (set) almacena solo **valores únicos** sin duplicados. Es como un map pero solo nos interesan las claves.

### ¿Para qué se usan?

1. **Deduplicación**: Remover duplicados de una lista
2. **Pertenencia**: ¿Está este valor en el conjunto?
3. **Operaciones con conjuntos**: Unión, intersección, diferencia
4. **Contadores**: Números únicos vistos

### Visualización

```
Array con duplicados:
[1, 2, 3, 2, 4, 1, 5] → Conversión a Set → {1, 2, 3, 4, 5}

Conjunto:
┌─────┐
│  1  │
│  2  │  (sin duplicados, sin orden específico)
│  3  │
│  4  │
│  5  │
└─────┘
```

### Implementación con map

```go
package main

import (
    "fmt"
    "sort"
)

type Conjunto struct {
    elementos map[int]bool  // Usa map como estructura interna
}

// Crear conjunto vacío
func NuevoConjunto() *Conjunto {
    return &Conjunto{
        elementos: make(map[int]bool),
    }
}

// Agregar elemento
func (s *Conjunto) Agregar(valor int) {
    s.elementos[valor] = true  // El valor del map no importa
}

// Eliminar elemento
func (s *Conjunto) Eliminar(valor int) {
    delete(s.elementos, valor)
}

// ¿Contiene?
func (s *Conjunto) Contiene(valor int) bool {
    return s.elementos[valor]  // false si no existe
}

// Tamaño
func (s *Conjunto) Tamaño() int {
    return len(s.elementos)
}

// Obtener elementos como slice
func (s *Conjunto) Items() []int {
    items := make([]int, 0, len(s.elementos))
    for k := range s.elementos {
        items = append(items, k)
    }
    sort.Ints(items)  // Para orden predecible
    return items
}

// ¿Está vacío?
func (s *Conjunto) Vacio() bool {
    return len(s.elementos) == 0
}

// Limpiar
func (s *Conjunto) Limpiar() {
    s.elementos = make(map[int]bool)
}

func main() {
    conjunto := NuevoConjunto()
    
    // Agregar elementos
    conjunto.Agregar(1)
    conjunto.Agregar(2)
    conjunto.Agregar(3)
    conjunto.Agregar(2)  // Duplicado, se ignora
    
    fmt.Println("¿Contiene 2?:", conjunto.Contiene(2))   // true
    fmt.Println("¿Contiene 10?:", conjunto.Contiene(10)) // false
    fmt.Println("Tamaño:", conjunto.Tamaño())             // 3
    fmt.Println("Items:", conjunto.Items())               // [1 2 3]
    
    // Eliminar
    conjunto.Eliminar(2)
    fmt.Println("Después de eliminar:", conjunto.Items())  // [1 3]
}
```

### Operaciones matemáticas con conjuntos

```go
package main

import "fmt"

type Conjunto struct {
    elementos map[int]bool
}

func NuevoConjunto(vals ...int) *Conjunto {
    c := &Conjunto{elementos: make(map[int]bool)}
    for _, v := range vals {
        c.elementos[v] = true
    }
    return c
}

func (s *Conjunto) Agregar(valor int) {
    s.elementos[valor] = true
}

func (s *Conjunto) Contiene(valor int) bool {
    return s.elementos[valor]
}

// UNIÓN: Todos los elementos de A y B
func (s *Conjunto) Union(otro *Conjunto) *Conjunto {
    resultado := NuevoConjunto()
    for k := range s.elementos {
        resultado.Agregar(k)
    }
    for k := range otro.elementos {
        resultado.Agregar(k)
    }
    return resultado
}

// INTERSECCIÓN: Elementos en A y en B
func (s *Conjunto) Interseccion(otro *Conjunto) *Conjunto {
    resultado := NuevoConjunto()
    for k := range s.elementos {
        if otro.Contiene(k) {
            resultado.Agregar(k)
        }
    }
    return resultado
}

// DIFERENCIA: Elementos en A pero no en B
func (s *Conjunto) Diferencia(otro *Conjunto) *Conjunto {
    resultado := NuevoConjunto()
    for k := range s.elementos {
        if !otro.Contiene(k) {
            resultado.Agregar(k)
        }
    }
    return resultado
}

// ¿Es subconjunto? (todos los elementos de A están en B)
func (s *Conjunto) EsSubconjunto(otro *Conjunto) bool {
    for k := range s.elementos {
        if !otro.Contiene(k) {
            return false
        }
    }
    return true
}

func (s *Conjunto) Items() []int {
    items := make([]int, 0, len(s.elementos))
    for k := range s.elementos {
        items = append(items, k)
    }
    return items
}

func main() {
    A := NuevoConjunto(1, 2, 3, 4)
    B := NuevoConjunto(3, 4, 5, 6)
    
    fmt.Println("A:", A.Items())  // [1 2 3 4]
    fmt.Println("B:", B.Items())  // [3 4 5 6]
    
    // Unión
    union := A.Union(B)
    fmt.Println("A ∪ B:", union.Items())  // [1 2 3 4 5 6]
    
    // Intersección
    intersec := A.Interseccion(B)
    fmt.Println("A ∩ B:", intersec.Items())  // [3 4]
    
    // Diferencia
    diferencia := A.Diferencia(B)
    fmt.Println("A - B:", diferencia.Items())  // [1 2]
    
    // Subconjunto
    C := NuevoConjunto(3, 4)
    fmt.Println("¿C ⊆ A?:", C.EsSubconjunto(A))  // true
    fmt.Println("¿B ⊆ A?:", B.EsSubconjunto(A))  // false
}
```

### Caso práctico: Remover duplicados

```go
package main

import "fmt"

func RemoverDuplicados(nums []int) []int {
    conjunto := make(map[int]bool)
    
    // Agregar todos al conjunto
    for _, n := range nums {
        conjunto[n] = true
    }
    
    // Convertir a slice
    resultado := make([]int, 0, len(conjunto))
    for k := range conjunto {
        resultado = append(resultado, k)
    }
    
    return resultado
}

func main() {
    nums := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
    unicos := RemoverDuplicados(nums)
    fmt.Println("Únicos:", unicos)  // [1 2 3 4 5]
}
```

### Con strings

```go
package main

import "fmt"

type ConjuntoString struct {
    elementos map[string]bool
}

func (s *ConjuntoString) Agregar(valor string) {
    if s.elementos == nil {
        s.elementos = make(map[string]bool)
    }
    s.elementos[valor] = true
}

func (s *ConjuntoString) Contiene(valor string) bool {
    return s.elementos[valor]
}

func main() {
    lenguajes := &ConjuntoString{}
    lenguajes.Agregar("Go")
    lenguajes.Agregar("Python")
    lenguajes.Agregar("Go")  // Duplicado
    
    fmt.Println("¿Tiene Go?:", lenguajes.Contiene("Go"))
    fmt.Println("Tamaño:", len(lenguajes.elementos))  // 2
}
```

### Complejidad
- **Agregar**: O(1)
- **Eliminar**: O(1)
- **Contiene**: O(1)
- **Unión**: O(n + m)
- **Intersección**: O(n) o O(m) (el menor)
- **Diferencia**: O(n)

### ¿Cuándo usar conjuntos?
✅ Deduplicar datos
✅ Comprobar pertenencia rápida
✅ Operaciones matemáticas
✅ Cuando el orden no importa
❌ Cuando necesitas mantener el orden

---

## 11. Árboles (Trees) - Búsqueda logarítmica en datos ordenados

Un **árbol** es una estructura **jerárquica** donde un nodo puede tener múltiples hijos. Un **árbol binario** es especial: cada nodo tiene máximo 2 hijos.

### Visualización

```
Árbol binario de búsqueda:
               50
              /  \
            30    70
           / \    / \
          20 40  60 80

Propiedades:
- Izquierda < Padre < Derecha
- Búsqueda rápida: O(log n)
```

### ¿Para qué se usan?

1. **Búsqueda rápida y ordenada**: Mejor que arrays + maps
2. **Datos jerárquicos**: Sistemas de archivos, jerarquías
3. **Expresiones matemáticas**: Parsing de fórmulas
4. **Autocompletado**: Trie (árbol de búsqueda de strings)

### Árbol binario de búsqueda completo
### Árbol binario de búsqueda completo

```go
package main

import "fmt"

type NodoArbol struct {
    Valor int
    Izquierda *NodoArbol
    Derecha *NodoArbol
}

type ArbolBinario struct {
    Raiz *NodoArbol
}

// INSERTAR
func (a *ArbolBinario) Insertar(valor int) {
    a.Raiz = a.insertarRecursivo(a.Raiz, valor)
}

func (a *ArbolBinario) insertarRecursivo(nodo *NodoArbol, valor int) *NodoArbol {
    if nodo == nil {
        return &NodoArbol{Valor: valor}
    }
    
    if valor < nodo.Valor {
        nodo.Izquierda = a.insertarRecursivo(nodo.Izquierda, valor)
    } else if valor > nodo.Valor {
        nodo.Derecha = a.insertarRecursivo(nodo.Derecha, valor)
    }
    // Si es igual, no insertar (evitar duplicados)
    
    return nodo
}

// BUSCAR
func (a *ArbolBinario) Buscar(valor int) bool {
    return a.buscarRecursivo(a.Raiz, valor)
}

func (a *ArbolBinario) buscarRecursivo(nodo *NodoArbol, valor int) bool {
    if nodo == nil {
        return false
    }
    
    if valor == nodo.Valor {
        return true
    } else if valor < nodo.Valor {
        return a.buscarRecursivo(nodo.Izquierda, valor)
    } else {
        return a.buscarRecursivo(nodo.Derecha, valor)
    }
}

// ENCONTRAR MÍNIMO
func (a *ArbolBinario) FindMin() *int {
    if a.Raiz == nil {
        return nil
    }
    nodo := a.Raiz
    for nodo.Izquierda != nil {
        nodo = nodo.Izquierda
    }
    return &nodo.Valor
}

// ENCONTRAR MÁXIMO
func (a *ArbolBinario) FindMax() *int {
    if a.Raiz == nil {
        return nil
    }
    nodo := a.Raiz
    for nodo.Derecha != nil {
        nodo = nodo.Derecha
    }
    return &nodo.Valor
}

// RECORRIDOS (formas de visitar todos los nodos)

// Inorden: Izquierda => Raíz => Derecha (orden ascendente)
func (a *ArbolBinario) RecorridoInorden() {
    a.recorridoInordenRecursivo(a.Raiz)
    fmt.Println()
}

func (a *ArbolBinario) recorridoInordenRecursivo(nodo *NodoArbol) {
    if nodo == nil {
        return
    }
    a.recorridoInordenRecursivo(nodo.Izquierda)
    fmt.Print(nodo.Valor, " ")
    a.recorridoInordenRecursivo(nodo.Derecha)
}

// Preorden: Raíz => Izquierda => Derecha
func (a *ArbolBinario) RecorridoPreorden() {
    a.recorridoPreordenRecursivo(a.Raiz)
    fmt.Println()
}

func (a *ArbolBinario) recorridoPreordenRecursivo(nodo *NodoArbol) {
    if nodo == nil {
        return
    }
    fmt.Print(nodo.Valor, " ")
    a.recorridoPreordenRecursivo(nodo.Izquierda)
    a.recorridoPreordenRecursivo(nodo.Derecha)
}

// Postorden: Izquierda => Derecha => Raíz
func (a *ArbolBinario) RecorridoPostorden() {
    a.recorridoPostordenRecursivo(a.Raiz)
    fmt.Println()
}

func (a *ArbolBinario) recorridoPostordenRecursivo(nodo *NodoArbol) {
    if nodo == nil {
        return
    }
    a.recorridoPostordenRecursivo(nodo.Izquierda)
    a.recorridoPostordenRecursivo(nodo.Derecha)
    fmt.Print(nodo.Valor, " ")
}

// Altura del árbol
func (a *ArbolBinario) Altura() int {
    return a.alturaRecursiva(a.Raiz)
}

func (a *ArbolBinario) alturaRecursiva(nodo *NodoArbol) int {
    if nodo == nil {
        return 0
    }
    izquierda := a.alturaRecursiva(nodo.Izquierda)
    derecha := a.alturaRecursiva(nodo.Derecha)
    
    if izquierda > derecha {
        return izquierda + 1
    }
    return derecha + 1
}

// Contar nodos
func (a *ArbolBinario) Contar() int {
    return a.contarRecursivo(a.Raiz)
}

func (a *ArbolBinario) contarRecursivo(nodo *NodoArbol) int {
    if nodo == nil {
        return 0
    }
    return 1 + a.contarRecursivo(nodo.Izquierda) + a.contarRecursivo(nodo.Derecha)
}

func main() {
    arbol := &ArbolBinario{}
    
    // Insertar valores (crea un árbol balanceado)
    valores := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35}
    for _, v := range valores {
        arbol.Insertar(v)
    }
    
    fmt.Print("Inorden (ascendente): ")
    arbol.RecorridoInorden()  // 10 20 25 30 35 40 50 60 70 80
    
    fmt.Print("Preorden: ")
    arbol.RecorridoPreorden()
    
    fmt.Print("Postorden: ")
    arbol.RecorridoPostorden()
    
    // Búsqueda
    fmt.Println("\n¿Buscar 40?:", arbol.Buscar(40))      // true
    fmt.Println("¿Buscar 999?:", arbol.Buscar(999))    // false
    
    // Min y Max
    min := arbol.FindMin()
    max := arbol.FindMax()
    fmt.Println("Mínimo:", *min)  // 10
    fmt.Println("Máximo:", *max)  // 80
    
    // Propiedades
    fmt.Println("Altura:", arbol.Altura())      // 4
    fmt.Println("Total de nodos:", arbol.Contar())  // 10
}
```

### Tipos de recorrido explicados

```
Árbol:
       50
      /  \
    30    70
   / \
  20 40

Inorden (izq-raíz-der):     20 30 40 50 70
Preorden (raíz-izq-der):    50 30 20 40 70
Postorden (izq-der-raíz):   20 40 30 70 50
```

### Eliminación (más complejo)

```go
// Eliminar es más complicado porque:
// 1. Si es hoja (sin hijos): simplemente eliminar
// 2. Si tiene 1 hijo: reemplazar con ese hijo
// 3. Si tiene 2 hijos: encuentrar el sucesor (mínimo del subárbol derecho)

func (a *ArbolBinario) Eliminar(valor int) {
    a.Raiz = a.eliminarRecursivo(a.Raiz, valor)
}

func (a *ArbolBinario) eliminarRecursivo(nodo *NodoArbol, valor int) *NodoArbol {
    if nodo == nil {
        return nil
    }
    
    if valor < nodo.Valor {
        nodo.Izquierda = a.eliminarRecursivo(nodo.Izquierda, valor)
    } else if valor > nodo.Valor {
        nodo.Derecha = a.eliminarRecursivo(nodo.Derecha, valor)
    } else {
        // Encontrado, ahora eliminar
        
        // Caso 1: Sin hijos (hoja)
        if nodo.Izquierda == nil && nodo.Derecha == nil {
            return nil
        }
        
        // Caso 2: Un hijo
        if nodo.Izquierda == nil {
            return nodo.Derecha
        }
        if nodo.Derecha == nil {
            return nodo.Izquierda
        }
        
        // Caso 3: Dos hijos
        // Encontrar el más pequeño del subárbol derecho (sucesor)
        sucesor := a.Raiz
        for sucesor.Izquierda != nil {
            sucesor = sucesor.Izquierda
        }
        nodo.Valor = sucesor.Valor
        nodo.Derecha = a.eliminarRecursivo(nodo.Derecha, sucesor.Valor)
    }
    
    return nodo
}
```

### Complejidad (árbol balanceado)
- **Búsqueda**: O(log n) ✅
- **Inserción**: O(log n) ✅
- **Eliminación**: O(log n) ✅
- **Traversal (todos)**: O(n)

**Nota**: Si el árbol está desbalanceado (ejemplo: insertar 1,2,3,4,5), se vuelve una lista enlazada: O(n) ❌

### Árbol AVL o Red-Black

Para evitar desbalanceo, usa árboles **auto-balanceables**:
- **AVL**: Rebalancea tras cada inserción
- **Red-Black**: Rebalancea menos frecuentemente (más eficiente)

### ¿Cuándo usar árboles?
✅ Búsqueda rápida en datos ordenados
✅ Datos jerárquicos
✅ Autocompletado (Trie)
❌ Cuando necesitas acceso rápido por clave (usa map)
❌ Cuando los datos son muy pequeños (usa arrays)

---

## 12. Grafos - Redes complejas de conexiones

Un **grafo** es una colección de **nodos** (vértices) conectados por **aristas** (bordes). Es la estructura más flexible para representar relaciones complejas.

### Visualización

```
Grafo:
  0 --- 1
  |     |
  2 --- 3

Nodos: 0, 1, 2, 3
Aristas: 0-1, 0-2, 1-3, 2-3
```

### Representación: Lista de adyacencia

```go
Grafo {
  0: [1, 2]
  1: [0, 3]
  2: [0, 3]
  3: [1, 2]
}
```

### ¿Para qué se usan?

1. **Redes sociales**: Quién conecta con quién
2. **GPS/Mapas**: Ciudades y carreteras, encontrar rutas
3. **Recomendaciones**: Qué películas ver
4. **Análisis de dependencias**: Módulos que dependen de otros
5. **Juegos**: Movimiento en mapa

### Grafo con lista de adyacencia completo

```go
package main

import (
    "fmt"
    "sort"
)

type Grafo struct {
    adyacencia map[int][]int
    nodos      map[int]bool
}

func NuevoGrafo() *Grafo {
    return &Grafo{
        adyacencia: make(map[int][]int),
        nodos:      make(map[int]bool),
    }
}

// Agregar nodo
func (g *Grafo) AgregarNodo(nodo int) {
    g.nodos[nodo] = true
    if _, existe := g.adyacencia[nodo]; !existe {
        g.adyacencia[nodo] = []int{}
    }
}

// Agregar arista (conexión)
func (g *Grafo) AgregarArista(u, v int) {
    g.AgregarNodo(u)
    g.AgregarNodo(v)
    
    // Agregar v a la lista de u
    g.adyacencia[u] = append(g.adyacencia[u], v)
    
    // Para grafo no dirigido, también agregar u a v
    // g.adyacencia[v] = append(g.adyacencia[v], u)
}

// Obtener vecinos de un nodo
func (g *Grafo) Vecinos(nodo int) []int {
    if vecinos, existe := g.adyacencia[nodo]; existe {
        copia := make([]int, len(vecinos))
        copy(copia, vecinos)
        return copia
    }
    return []int{}
}

// ¿Existe arista entre u y v?
func (g *Grafo) ExisteArista(u, v int) bool {
    for _, vecino := range g.adyacencia[u] {
        if vecino == v {
            return true
        }
    }
    return false
}

// Cantidad de nodos
func (g *Grafo) CantidadNodos() int {
    return len(g.nodos)
}

// Cantidad de aristas
func (g *Grafo) CantidadAristas() int {
    count := 0
    for _, vecinos := range g.adyacencia {
        count += len(vecinos)
    }
    return count
}

// BFS (Búsqueda en amplitud) - Nivel por nivel
func (g *Grafo) BFS(inicio int) {
    visitados := make(map[int]bool)
    cola := []int{inicio}
    visitados[inicio] = true
    
    for len(cola) > 0 {
        nodo := cola[0]
        cola = cola[1:]  // Sacamos del inicio
        fmt.Print(nodo, " ")
        
        // Agregar vecinos no visitados
        for _, vecino := range g.adyacencia[nodo] {
            if !visitados[vecino] {
                visitados[vecino] = true
                cola = append(cola, vecino)
            }
        }
    }
    fmt.Println()
}

// DFS (Búsqueda en profundidad) - Profundo primero
func (g *Grafo) DFS(inicio int) {
    visitados := make(map[int]bool)
    g.dfsRecursivo(inicio, visitados)
    fmt.Println()
}

func (g *Grafo) dfsRecursivo(nodo int, visitados map[int]bool) {
    visitados[nodo] = true
    fmt.Print(nodo, " ")
    
    // Visitar todos los vecinos
    for _, vecino := range g.adyacencia[nodo] {
        if !visitados[vecino] {
            g.dfsRecursivo(vecino, visitados)
        }
    }
}

// Detectar ciclos
func (g *Grafo) TieneCiclo() bool {
    visitados := make(map[int]bool)
    
    for nodo := range g.nodos {
        if !visitados[nodo] {
            if g.tieneCicloRecursivo(nodo, visitados, -1) {
                return true
            }
        }
    }
    return false
}

func (g *Grafo) tieneCicloRecursivo(nodo int, visitados map[int]bool, padre int) bool {
    visitados[nodo] = true
    
    for _, vecino := range g.adyacencia[nodo] {
        if !visitados[vecino] {
            if g.tieneCicloRecursivo(vecino, visitados, nodo) {
                return true
            }
        } else if vecino != padre {
            // Visitado y no es el padre = ciclo
            return true
        }
    }
    return false
}

// Distancia más corta entre dos nodos (BFS)
func (g *Grafo) DistanciaCorta(inicio, fin int) int {
    if inicio == fin {
        return 0
    }
    
    visitados := make(map[int]bool)
    distancias := make(map[int]int)
    cola := []int{inicio}
    visitados[inicio] = true
    distancias[inicio] = 0
    
    for len(cola) > 0 {
        nodo := cola[0]
        cola = cola[1:]
        
        for _, vecino := range g.adyacencia[nodo] {
            if !visitados[vecino] {
                visitados[vecino] = true
                distancias[vecino] = distancias[nodo] + 1
                
                if vecino == fin {
                    return distancias[vecino]
                }
                
                cola = append(cola, vecino)
            }
        }
    }
    
    return -1  // No hay camino
}

// Camino completo entre dos nodos
func (g *Grafo) CaminoCorto(inicio, fin int) []int {
    visitados := make(map[int]bool)
    padre := make(map[int]int)
    cola := []int{inicio}
    visitados[inicio] = true
    padre[inicio] = -1
    
    for len(cola) > 0 {
        nodo := cola[0]
        cola = cola[1:]
        
        if nodo == fin {
            // Reconstruir camino
            camino := []int{}
            actual := fin
            for actual != -1 {
                camino = append([]int{actual}, camino...)
                actual = padre[actual]
            }
            return camino
        }
        
        for _, vecino := range g.adyacencia[nodo] {
            if !visitados[vecino] {
                visitados[vecino] = true
                padre[vecino] = nodo
                cola = append(cola, vecino)
            }
        }
    }
    
    return []int{}  // No hay camino
}

// Conéctados: ¿todos los nodos están conectados?
func (g *Grafo) EstaConectado() bool {
    if g.CantidadNodos() == 0 {
        return true
    }
    
    // Tomar un nodo arbitrario
    var inicio int
    for nodo := range g.nodos {
        inicio = nodo
        break
    }
    
    visitados := make(map[int]bool)
    g.dfsContador(inicio, visitados)
    
    return len(visitados) == len(g.nodos)
}

func (g *Grafo) dfsContador(nodo int, visitados map[int]bool) {
    visitados[nodo] = true
    for _, vecino := range g.adyacencia[nodo] {
        if !visitados[vecino] {
            g.dfsContador(vecino, visitados)
        }
    }
}

// Imprimir grafo
func (g *Grafo) Imprimir() {
    nodos := make([]int, 0, len(g.nodos))
    for n := range g.nodos {
        nodos = append(nodos, n)
    }
    sort.Ints(nodos)
    
    for _, nodo := range nodos {
        fmt.Printf("%d -> %v\n", nodo, g.adyacencia[nodo])
    }
}

func main() {
    g := NuevoGrafo()
    
    // Crear grafo
    g.AgregarArista(0, 1)
    g.AgregarArista(0, 2)
    g.AgregarArista(1, 3)
    g.AgregarArista(2, 3)
    g.AgregarArista(3, 4)
    
    fmt.Println("Grafo:")
    g.Imprimir()
    
    fmt.Print("\nBFS desde 0: ")
    g.BFS(0)
    
    fmt.Print("DFS desde 0: ")
    g.DFS(0)
    
    fmt.Println("¿Tiene ciclos?:", g.TieneCiclo())
    
    fmt.Println("¿Está conectado?:", g.EstaConectado())
    
    fmt.Println("Distancia 0 -> 4:", g.DistanciaCorta(0, 4))
    
    fmt.Println("Camino 0 -> 4:", g.CaminoCorto(0, 4))
    
    // Grafo con ciclo
    g2 := NuevoGrafo()
    g2.AgregarArista(0, 1)
    g2.AgregarArista(1, 2)
    g2.AgregarArista(2, 0)  // Crea ciclo
    fmt.Println("\n¿Grafo 2 tiene ciclos?:", g2.TieneCiclo())  // true
}
```

### Diferencias BFS vs DFS

```
       0
      / \
     1   2
     |   |
     3   4

BFS (por niveles):    0 1 2 3 4
DFS (profundo):       0 1 3 2 4

BFS = Cola   (FIFO)
DFS = Pila   (LIFO) o Recursión
```

### Complejidad
- **BFS**: O(V + E) - Visita cada nodo y arista una vez
- **DFS**: O(V + E) - Igual que BFS
- **Camino más corto**: O(V + E)
- **Detectar ciclos**: O(V + E)

Donde V = cantidad de nodos, E = cantidad de aristas

### ¿Cuándo usar grafos?
✅ Redes con conexiones complejas
✅ Encontrar rutas (GPS, recomendaciones)
✅ Detectar ciclos o dependencias
✅ Análisis de conectividad
❌ Datos simples lineales (usa arrays)

---

---

## 13. Tabla de Comparación Completa

| Estructura | Acceso | Búsqueda | Inserción | Eliminación | Memoria | Uso |
|-----------|--------|----------|-----------|-------------|---------|-----|
| Array | O(1) | O(n) | O(n) | O(n) | Fijo ✅ | Datos fijos |
| Slice | O(1) | O(n) | O(n) | O(n) | Dinámico ✅ | Datos dinámicos |
| Map | O(1) | O(1) ✅ | O(1) ✅ | O(1) ✅ | Dinámico ✅ | Clave-valor |
| Struct | - | - | - | - | Fijo | Agrupar datos |
| Lista enlazada | O(n) | O(n) | O(1)* | O(1)* | Dinámico ✅ | Inserción frecuente |
| Pila | - | O(n) | O(1)** | O(1) ✅ | Dinámico | LIFO, recursión |
| Cola | - | O(n) | O(1)** | O(1)*** | Dinámico | FIFO, BFS |
| Árbol binario | O(log n) ✅ | O(log n) ✅ | O(log n) ✅ | O(log n) ✅ | Dinámico ✅ | Búsqueda ordenada |
| Conjunto | - | O(1) ✅ | O(1) ✅ | O(1) ✅ | Dinámico ✅ | Valores únicos |
| Grafo | Varía | Varía | Varía | Varía | Dinámico ✅ | Redes complejas |

*Al inicio
**O(1) amortizado
***O(n) con slice simple, O(1) con índices

---

## 14. Guía de decisión práctica

### Problema: "Tengo datos y necesito almacenarlos"

```
¿Necesitas acceso por índice?
│
├─ SÍ ─→ ¿Tamaño variab?
│        ├─ SÍ ──→ SLICE (99% de casos)
│        └─ NO ──→ ARRAY (casos especiales)
│
└─ NO ─→ ¿Acceso por clave?
         ├─ SÍ ──→ MAP (búsqueda rápida)
         └─ NO ──→ STRUCT (agrupar datos relacionados)
```

### Problema: "Necesito datos ordenados y búsqueda rápida"

```
¿Datos muy grandes?
│
├─ SÍ ──→ ¿Necesitas insertar/eliminar frecuentemente en orden?
│         ├─ SÍ ──→ ÁRBOL (mantieneorden automáticamente)
│         └─ NO ──→ MAP + slice ordenado (sort)
│
└─ NO ──→ SLICE + sort package (más simple)
```

### Problema: "Procesamiento en orden específico"

```
¿Qué orden?
│
├─ Primero entrado = Primero salido ──→ COLA (FIFO)
│
└─ Último entrado = Primero salido ──→ PILA (LIFO)
```

### Problema: "Navegar redes complejas"

```
¿Relaciones complejas entre datos?
│
├─ SÍ, muchas conexiones ──→ GRAFO
│   ├─ ¿Buscar distancia? ──→ BFS
│   └─ ¿Explorar todo? ──→ DFS
│
└─ NO, relación padre-hijo ──→ ÁRBOL
```

---

## 15. Ejemplos del mundo real

### Ejemplo 1: Sistema de recomendaciones

```go
// Grafo de usuarios conectados
grafo := NuevoGrafo()  // Users como nodos

// Qué películas vieron amigos (BFS)
recomendaciones := obtenerRecomendaciones(usuarioActual, 2)  // Amigos de amigos
```

### Ejemplo 2: Sistema de archivos

```go
// Árbol de directorios
raiz := &Directorio{
    nombre: "/",
    subdirs: []*Directorio{...},  // Árbol
    archivos: []string{...},
}

// Buscar archivo
buscar(raiz, "/home/usuario/doc.txt")
```

### Ejemplo 3: Generador de sugerencias (autocompletado)

```go
// Trie (árbol especializado) para strings
type TrieNode struct {
    hijos map[rune]*TrieNode
    esFinPalabra bool
}

// Insertar palabras: "cat", "car", "dog"
// Buscar: "ca" → ["cat", "car"]
```

### Ejemplo 4: Sistema de caché

```go
// Map para acceso O(1) rápido
cache := map[string]interface{}{}
cache["usuario:123"] = usuario
cache["post:456"] = post

// Con TTL (tiempo de expiración)
cacheTTL := map[string]CacheEntry{
    "key": {valor: data, expiracion: time.Now().Add(5*time.Minute)},
}
```

### Ejemplo 5: Verificar paréntesis balanceados

```go
// Pila para tracking
pila := [...]rune{}
for _, char := range "foo(bar(baz))" {
    if char == '(' {
        pila.push('(')
    } else if char == ')' {
        pila.pop()  // Verifica coincidencia
    }
}
// Si pila está vacía al final, está balanceado
```

---

## 16. Mejores prácticas detalladas

### 1. Inicialización correcta

```go
// ❌ MALO
var m map[string]int
m["key"] = 1  // PANIC

// ✅ BIEN
m := make(map[string]int)
m["key"] = 1

// ❌ MALO
var s []int
s = append(s, 1)  // OK pero may ser ineficiente

// ✅ BIEN
s := make([]int, 0, 10)  // Preasigna capacidad
s = append(s, 1)
```

### 2. Usar punteros sabiamente

```go
// ❌ MALO - copiar struct grande
type GranEstructura struct {
    datos [10000]int
}

func Procesar(g GranEstructura) {  // Copia toda
    // ...
}

// ✅ BIEN - pasar puntero
func Procesar(g *GranEstructura) {  // Solo puntero (8 bytes)
    // ...
}
```

### 3. Evitar nil panics

```go
// ❌ MALO
var p *int
fmt.Println(*p)  // PANIC

// ✅ BIEN
var p *int
if p != nil {
    fmt.Println(*p)
}

// ✅ BIEN - usar default
var s []int  // nil slice es válido
for _, v := range s {  // itera 0 veces
    fmt.Println(v)
}
```

### 4. Mensajes de error claros

```go
// ❌ MALO
if !existe {
    return fmt.Errorf("error")
}

// ✅ BIEN
if !existe {
    return fmt.Errorf("usuario %d no encontrado en el sistema", userID)
}
```

### 5. Documentar decisiones

```go
// ✅ BIEN - explica por qué usas esta estructura
// Mapa para acceso O(1) a usuarios por ID
usuarios := make(map[string]*Usuario)

// ✅ BIEN - explica complejidad importante
// Tree binario de búsqueda mantiene datos ordenados con búsqueda O(log n)
arbol := &ArbolBinario{}
```

### 6. Rendimiento y profiling

```go
// Usa pprof para medir verdadero rendimiento
import _ "net/http/pprof"

go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()

// http://localhost:6060/debug/pprof
// No adivines, mide
```

### 7. Concurrencia segura

```go
// Slices no son thread-safe. Para goroutines usa:
// - channels
// - sync.Mutex
// - sync.Map (para maps)

type MiEstructura struct {
    mu       sync.Mutex
    slices   []int
}

func (m *MiEstructura) Agregar(valor int) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.slices = append(m.slices, valor)
}
```

---

## 17. Antipatrones a evitar

```go
// ❌ NO: Iterar y modificar simultáneamente
slice := []int{1,2,3,4,5}
for i, v := range slice {
    if v == 3 {
        slice = append(slice[:i], slice[i+1:]...)  // Evitar
    }
}

// ✅ SÍ: Crear nuevo slice
slice := []int{1,2,3,4,5}
resultado := []int{}
for _, v := range slice {
    if v != 3 {
        resultado = append(resultado, v)
    }
}

// ❌ NO: Usar reflection innecesariamente
//go (pseudo-código)
valor := reflect.ValueOf(datos)  // Lento

// ✅ SÍ: Usa types concretos
// Específica tu tipo correctamente

// ❌ NO: No preasignar capacidad
items := []Item{}
for i := 0; i < 1000000; i++ {
    items = append(items, Item{...})  // Realocaciones frecuentes
}

// ✅ SÍ: Preasignar
items := make([]Item, 0, 1000000)
for i := 0; i < 1000000; i++ {
    items = append(items, Item{...})
}
```

---

## 18. Resumen y recomendaciones finales

### Lo más importante

1. **Arrays y slices** resuelven 90% de problemas cotidianos
2. **Maps** son tu herramienta para búsqueda rápida por clave
3. **Structs** te ayudan a organizar código
4. **Punteros** permiten modificaciones y eficiencia
5. **Interfaces** hacen código flexible y testeable

### Flujo de decisión simple

```
Necesitas almacenar datos?
│
├─ Datos de tamaño variable → SLICE
├─ Búsqueda por clave → MAP
├─ Agrupar tipos diferentes → STRUCT
├─ Procesar en orden especial → PILA/COLA
├─ Búsqueda ordenada → ÁRBOL
├─ Red de conexiones → GRAFO
└─ Valores sin duplicados → CONJUNTO
```

### Antes de optimizar

1. **Escribe código claro** - Correctitud primero
2. **Mide rendimiento** - No adivines, usa `pprof`
3. **Optimiza específicamente** - Solo dónde duele
4. **Reutiliza estándar** - Go stdlib es muy bueno

### Debes conocer y practicar

- ✅ Arrays y slices (usarás diariamente)
- ✅ Maps (muy común)
- ✅ Structs e interfaces (esencial para Go)
- ✅ Punteros (fundamental)
- ✅ Árboles (para búsquedas avanzadas)
- ✅ Grafos (análisis de relaciones)
- ⚠️ Listas enlazadas (raro en Go, usa slices)

---

## 19. Conclusión

Las estructuras de datos son la base de la programación eficiente. En Go:

- Comienza simple: **slice, map, struct**
- Aprende cuándo cambiar: Mide, vea el cuello de botella
- Implementa lo necesario: Los árboles y grafos cuando los necesites
- Nunca prematuramente optimices: Claridad > rendimiento micro

**Recuerda**: *La mejor estructura de datos es la que tu equipo entiende y puede mantener.*

Go facilita esto con `interfaces` que permiten cambiar implementaciones sin modificar código cliente. Aprovecha eso.

¡Ahora tienes el fundamento completo para resolver casi cualquier problema de datos!
