# Funciones en Go - Guía Completa

## Introducción: ¿Por qué existen las funciones?

Imagina que escribes un programa que necesita calcular el promedio de notas 50 veces. Podrías copiar/pegar el mismo código 50 veces, pero sería desastroso:
- **Código duplicado**: Si encuentras un bug, debes arreglarlo en 50 lugares
- **Mantenimiento imposible**: Cambiar la lógica es una pesadilla
- **Inlegible**: El código se vuelve un desorden

Las funciones son la solución: **un bloque de código reutilizable que realiza una tarea específica**.

En Go, las funciones son **"ciudadanos de primera clase"**, lo que significa que puedes:
- ✅ Pasarlas como argumentos a otras funciones
- ✅ Devolverlas como valores
- ✅ Almacenarlas en variables
- ✅ Usarlas en estructuras de datos (slices, maps, etc.)

Esto hace que Go sea excepcionally flexible para programación funcional.

**En esta guía aprenderás**:
1. Cómo definir funciones básicas
2. Parámetros (simples, variables, punteros)
3. Devoluciones (uno, múltiples, nombrados)
4. Funciones anónimas y closures
5. Funciones como argumentos (callbacks)
6. Defer, panic, recover control avanzado
7. Recursión y patrones funcionales

---

## Una metáfora: Funciones como máquinas

Piensa en una función como una máquina:
- **Entrada**: Parámetros (ingredientes)
- **Proceso**: Cuerpo de la función
- **Salida**: Valor de retorno (producto)

```
Parámetros              Cuerpo                Retorno
    ↓                     ↓                      ↓
   5, 3   →   [Máquina Sumar]   →   8
```

Una buena función:
- Realiza **una tarea específica** (responsabilidad única)
- Tiene **entradas claras** (parámetros con nombres descriptivos)
- Tiene **una salida clara** (retorna algo útil)
- Es **predecible** (mismo input siempre da mismo output)

---

A continuación aprenderemos cómo construir estas "máquinas" en Go.

## Conceptos Clave

Una **función** es un bloque de código reutilizable que realiza una tarea específica. En Go, las funciones son "ciudadanos de primera clase", lo que significa que puedes pasarlas como argumentos y devolverlas como valores.

## 1. Función básica

```go
// Sintaxis:
// func nombre(parametros) tipoRetorno {
//     cuerpo
//     return valor
// }

package main

import "fmt"

// Función simple: sumar dos números
func Sumar(a int, b int) int {
    return a + b
}

func main() {
    resultado := Sumar(5, 3)
    fmt.Println(resultado)  // 8
}
```

### Características
- **func**: palabra clave para definir función
- **Nombre**: Comienza con mayúscula si es pública (exportada)
- **Parámetros**: Entre paréntesis con tipo
- **Tipo retorno**: Después de paréntesis
- **return**: Devuelve valor

---

## 2. Parámetros

### Parámetros del mismo tipo

```go
// Forma larga (explicita)
func Restar(a int, b int) int {
    return a - b
}

// Forma corta (Go permite agrupar tipos)
func Restar(a, b int) int {
    return a - b
}

// Ambos hacen lo mismo, usa la forma corta
```

### Parámetros variables (variadic)

```go
func Suma(numeros ...int) int {
    total := 0
    for _, n := range numeros {
        total += n
    }
    return total
}

func main() {
    fmt.Println(Suma(1, 2, 3))           // 6
    fmt.Println(Suma(1, 2, 3, 4, 5))     // 15
    
    // También puedes pasar un slice
    nums := []int{10, 20, 30}
    fmt.Println(Suma(nums...))  // 60 (desempaqueta el slice)
}
```

**Nota**: El parámetro variadic debe ser el último

### Parámetros por referencia (punteros)

```go
// Pasar por VALUE (copia)
func IncrementarCopia(x int) {
    x = x + 1
    // No afecta el original
}

// Pasar por REFERENCIA (puntero)
func IncrementarReferencia(x *int) {
    *x = *x + 1
    // Modifica el original
}

func main() {
    a := 5
    IncrementarCopia(a)
    fmt.Println(a)  // 5 (sin cambios)
    
    b := 5
    IncrementarReferencia(&b)
    fmt.Println(b)  // 6 (modificado)
}
```

---

## 3. Valores de retorno

### Un valor

```go
func Dividir(a, b int) int {
    return a / b
}

resultado := Dividir(10, 2)  // 5
```

### Múltiples valores

```go
// Función que devuelve dos valores
func DividirConResto(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    cociente, resto := DividirConResto(10, 3)
    fmt.Println(cociente, resto)  // 3 1
    
    // Si ignoras un valor, usa _
    cociente, _ := DividirConResto(10, 3)
    fmt.Println(cociente)  // 3
}
```

### Retorno nombrado

```go
// Los valores de retorno pueden tener nombres
func DividirNombrado(a, b int) (cociente int, resto int) {
    cociente = a / b
    resto = a % b
    return  // return vacío devuelve los valores nombrados
}

func main() {
    c, r := DividirNombrado(10, 3)
    fmt.Println(c, r)  // 3 1
}
```

**Ventaja**: Documenta qué devuelves, puedes hacer return sin argumentos

### Patrón error/valor

```go
// Go usa mucho este patrón
func Buscar(slice []int, valor int) (indice int, encontrado bool) {
    for i, v := range slice {
        if v == valor {
            return i, true
        }
    }
    return -1, false
}

func main() {
    nums := []int{10, 20, 30, 40}
    idx, ok := Buscar(nums, 20)
    if ok {
        fmt.Println("Encontrado en índice:", idx)
    } else {
        fmt.Println("No encontrado")
    }
}
```

---

## 4. Funciones anónimas y closures

### Función anónima

```go
func main() {
    // Función sin nombre (anónima)
    suma := func(a, b int) int {
        return a + b
    }
    
    fmt.Println(suma(5, 3))  // 8
    
    // O ejecutarla inmediatamente
    resultado := func(x int) int {
        return x * 2
    }(10)
    
    fmt.Println(resultado)  // 20
}
```

### Closures (capturar variables del contexto)

```go
func Contador() func() int {
    count := 0
    
    // Esta función "recuerda" count
    return func() int {
        count++
        return count
    }
}

func main() {
    c1 := Contador()
    c2 := Contador()
    
    fmt.Println(c1())  // 1
    fmt.Println(c1())  // 2
    fmt.Println(c1())  // 3
    
    fmt.Println(c2())  // 1 (su propio count)
    fmt.Println(c2())  // 2
}
```

Cada llamada a `Contador()` crea su propia variable `count`.

### Closures modificables

```go
func Generador(inicio int) func() int {
    valor := inicio
    return func() int {
        valor += 2
        return valor
    }
}

func main() {
    gen := Generador(0)
    fmt.Println(gen())  // 2
    fmt.Println(gen())  // 4
    fmt.Println(gen())  // 6
}
```

---

## 5. Funciones como argumentos

### Pasar función como parámetro

```go
// Función que recibe otra función
func Aplicar(f func(int) int, valor int) int {
    return f(valor)
}

func main() {
    duplicar := func(x int) int {
        return x * 2
    }
    
    resultado := Aplicar(duplicar, 5)
    fmt.Println(resultado)  // 10
}
```

### Callbacks (función ejecutada después)

```go
func Procesar(datos []int, callback func(int)) {
    for _, d := range datos {
        callback(d)
    }
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    
    Procesar(numeros, func(n int) {
        fmt.Println("Procesando:", n)
    })
}
```

### Map/Filter pattern

```go
// Map: aplicar función a cada elemento
func Map(slice []int, f func(int) int) []int {
    resultado := make([]int, len(slice))
    for i, v := range slice {
        resultado[i] = f(v)
    }
    return resultado
}

// Filter: mantener elementos que pasan test
func Filter(slice []int, f func(int) bool) []int {
    var resultado []int
    for _, v := range slice {
        if f(v) {
            resultado = append(resultado, v)
        }
    }
    return resultado
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    
    // Duplicar cada número
    duplicados := Map(nums, func(x int) int {
        return x * 2
    })
    fmt.Println(duplicados)  // [2 4 6 8 10]
    
    // Mantener solo pares
    pares := Filter(nums, func(x int) bool {
        return x%2 == 0
    })
    fmt.Println(pares)  // [2 4]
}
```

---

## 6. Defer - Ejecutar al finalizar

`defer` ejecuta código cuando la función termina, útil para limpieza.

```go
import "fmt"

func main() {
    fmt.Println("1. Inicio")
    
    defer func() {
        fmt.Println("3. Limpieza (deferred)")
    }()
    
    fmt.Println("2. Proceso")
}

// Output:
// 1. Inicio
// 2. Proceso
// 3. Limpieza (deferred)
```

### Múltiples defer (LIFO)

```go
func main() {
    defer fmt.Println("1. Primero")
    defer fmt.Println("2. Segundo")
    defer fmt.Println("3. Tercero")
}

// Output:
// 3. Tercero
// 2. Segundo
// 1. Primero

// Se ejecutan en orden inverso (LIFO)
```

### Casos prácticos

```go
import "fmt"

// Cerrar archivo automáticamente
func Leer(archivo string) {
    f := AbrirArchivo(archivo)
    defer f.Close()  // Se ejecutará al final
    
    // Usar f
    f.Leer()
}

// Liberar recurso
func ConLock(m *sync.Mutex) {
    m.Lock()
    defer m.Unlock()  // Desbloquea al final
    
    // Usar recurso protegido
}

// Recuperar panic
func Seguro() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Error recuperado:", r)
        }
    }()
    
    // Código que puede fallar
}
```

---

## 7. Panic y recover

### Panic - Crash controlado

```go
func Dividir(a, b int) int {
    if b == 0 {
        panic("división por cero")
    }
    return a / b
}

func main() {
    resultado := Dividir(10, 2)
    fmt.Println(resultado)  // 5
    
    // resultado := Dividir(10, 0)  // PANIC: división por cero
}
```

### Recover - Capturar panic

```go
func Seguro(f func()) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Problema atrapado:", r)
        }
    }()
    
    f()  // Ejecutar función
}

func main() {
    Seguro(func() {
        panic("algo salió mal")
    })
    
    fmt.Println("El programa continúa")
}

// Output:
// Problema atrapado: algo salió mal
// El programa continúa
```

---

## 8. Funciones recursivas

```go
// Factorial: 5! = 5 * 4 * 3 * 2 * 1 = 120
func Factorial(n int) int {
    if n <= 1 {
        return 1  // Caso base
    }
    return n * Factorial(n-1)  // Caso recursivo
}

func main() {
    fmt.Println(Factorial(5))  // 120
}
```

### Recursión con memoización

```go
// Fibonacci: 0, 1, 1, 2, 3, 5, 8, 13...
func Fibonacci(n int, memo map[int]int) int {
    if n <= 1 {
        return n
    }
    
    if valor, existe := memo[n]; existe {
        return valor  // Ya calculado
    }
    
    resultado := Fibonacci(n-1, memo) + Fibonacci(n-2, memo)
    memo[n] = resultado
    return resultado
}

func main() {
    memo := make(map[int]int)
    fmt.Println(Fibonacci(10, memo))  // 55 (muy rápido)
}
```

---

## 9. Funciones de orden superior

```go
// Función que devuelve una función
func Multiplicador(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    doble := Multiplicador(2)
    triple := Multiplicador(3)
    
    fmt.Println(doble(5))   // 10
    fmt.Println(triple(5))  // 15
}
```

### Composición de funciones

```go
// Compose: f(g(x))
func Compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

func main() {
    sumar2 := func(x int) int { return x + 2 }
    multiplicar3 := func(x int) int { return x * 3 }
    
    h := Compose(sumar2, multiplicar3)
    
    fmt.Println(h(5))  // sumar2(multiplicar3(5)) = sumar2(15) = 17
}
```

---

## 10. Inline functions vs Named functions

```go
// ✅ Usa named functions cuando:
// - La lógica se reutiliza
// - Es compleja o documentable
// - Necesitas recursión
func Procesar(datos []int) {
    resultado := make([]int, 0)
    for _, d := range datos {
        if EsValido(d) {
            resultado = append(resultado, d*2)
        }
    }
}

func EsValido(n int) bool {
    return n > 0
}

// ✅ Usa inline functions cuando:
// - Solo la usas una vez
// - Es simple (1-2 líneas)
// - Necesitas closure

items := []int{1, -2, 3}
Filter(items, func(x int) bool {
    return x > 0
})
```

---

## 11. Tabla de tiempos de ejecución

| Operación | Tiempo | Ejemplo |
|-----------|--------|---------|
| Llamada función | O(1) | `f()` |
| Recursión profunda | O(n) | Fibonacci sin memo |
| Con memoización | O(n) | Fibonacci con memo |
| Función variadic | O(n) | `f(...args)` |
| Defer | O(1) | `defer cleanup()` |

---

## 12. Mejores prácticas

### ✅ Bien
```go
// Nombres claros
func CalcularPrecioTotal(precio, impuesto float64) float64 {
    return precio + (precio * impuesto)
}

// Retorno nombrado documenta
func BuscarUsuario(id int) (usuario *Usuario, encontrado bool) {
    // ...
    return usuario, true
}

// Usar defer para limpieza
func LeerArchivo(archivo string) ([]byte, error) {
    f, err := os.Open(archivo)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    // ...
}

// Parámetros claros
func Saludar(nombre string) string {
    return "Hola, " + nombre
}
```

### ❌ Mal
```go
// Nombres confusos
func calc(p, i float64) float64 {
    return p + (p * i)
}

// Sin documentar retornos
func get(id int) (*Usuario, bool) {
    // ¿Qué retorna?
}

// No limpiar recursos
func leer(archivo string) []byte {
    f, _ := os.Open(archivo)
    // f nunca se cierra (memory leak)
}

// Demasiados parámetros
func guardar(a, b, c, d, e, f string) {
    // Difícil de usar
}
```

---

## 13. Conclusión - El Poder de las Funciones

Funciones en Go:
- **Simples**: Sintaxis clara y explícil
- **Flexibles**: Soportan closures, parámetros variadic, retornos múltiples
- **Seguras**: defer y recover para manejo robusto
- **Funcionales**: Funciones como valores abren mundo de posibilidades

### Resumen de patrones

| Patrón | Caso de Uso |
|--------|------------|
| Función simple | La mayoría de casos |
| Parámetros variadic | Número variable de args |
| Múltiples retornos | Retornar resultado + error |
| Closures | Capturar estado del contexto |
| Callbacks | Ejecutar lógica externa |
| Defer | Garantizar limpieza |
| Panic/Recover | Manejo de errores críticos |
| Recursión | Problemas dividibles |

### Reglas de oro

1. **Una función, una responsabilidad**: Si describe la función y requiere "y" en la descripción, probablemente hace demasiado.

2. **Nombres claros**: Un buen nombre de función es casi documentación.

3. **Parámetros simples**: Si necesitas 5+ parámetros, usa un struct.

4. **Retorna errores**: Usa el patrón `(resultado, error)`.

5. **Usa defer**: Para cerrar archivos, desbloquear, etc.

6. **Evita side effects**: Una función pura es más testeable.

### El viaje de un programador con funciones

**Principiante**: Escribe funciones básicas que hacen tareas simples.

**Intermedio**: Descubre closures, callbacks, orden superior. ¡Ahhhhh!

**Avanzado**: Combina todo esto para crear abstracciones elegantes y reutilizables.

Go te permite llegar a "avanzado" sin complejidad innecesaria. No hay decoradores esotéricos, no hay métodos privados vs públicos complicados. Solo funciones claras y poderosas.

**Domina funciones y dominarás Go.**
- **Claras**: Nombres descriptivos (convención)

Con funciones bien diseñadas, tu código será modular, testeable y mantenible.
