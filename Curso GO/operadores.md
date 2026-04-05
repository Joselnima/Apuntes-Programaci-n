# Operadores en Go

Go tiene operadores para realizar operaciones matemáticas, comparaciones y lógica. A continuación se describen los más usados y ejemplos de cuándo emplearlos.

## Operadores aritméticos

- `+` : suma
- `-` : resta
- `*` : multiplicación
- `/` : división
- `%` : módulo (resto de la división)

Para qué sirven:
- calcular valores numéricos
- actualizar contadores
- operaciones con tiempo, distancias, cantidades, etc.

Cuándo se usan:
- `a + b` para sumar números enteros o flotantes.
- `x - y` para restar valores o realizar diferencias.
- `p * q` para multiplicar cantidades.
- `total / n` para dividir valores cuando se conoce que `n` no es cero.
- `n % 2` para obtener el resto y determinar paridad.

Ejemplo:

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3
    fmt.Println(a + b) // 13
    fmt.Println(a - b) // 7
    fmt.Println(a * b) // 30
    fmt.Println(a / b) // 3
    fmt.Println(a % b) // 1
}
```

## Operadores de asignación

- `=` : asignar valor
- `+=` : suma y asigna
- `-=` : resta y asigna
- `*=` : multiplica y asigna
- `/=` : divide y asigna
- `%=` : módulo y asigna

Para qué sirven:
- actualizar variables sin escribir la variable completa de nuevo.
- simplificar expresiones acumulativas.

Cuándo se usan:
- `count += 1` en contadores.
- `total *= 2` para duplicar valores.
- `x %= 10` para conservar el resto.

Ejemplo:

```go
x := 5
x += 3 // x = 8
x *= 2 // x = 16
```

## Operadores relacionales (comparación)

- `==` : igual a
- `!=` : distinto de
- `<` : menor que
- `<=` : menor o igual que
- `>` : mayor que
- `>=` : mayor o igual que

Para qué sirven:
- comparar valores
- controlar flujo con condicionales
- verificar estados y condiciones de error

Cuándo se usan:
- `if edad >= 18 { ... }` para chequear mayoría de edad.
- `if a == b { ... }` para comparar igualdad.
- `if saldo != 0 { ... }` para detectar valores no nulos.

Ejemplo:

```go
if score >= 90 {
    fmt.Println("Excelente")
}
```

## Operadores lógicos

- `&&` : AND lógico
- `||` : OR lógico
- `!` : NOT lógico

Para qué sirven:
- combinar múltiples condiciones
- construir filtros y validaciones complejas
- invertir resultados booleanos

Cuándo se usan:
- `if x > 0 && x < 10 { ... }` para verificar un rango.
- `if isAdmin || isOwner { ... }` para permisos.
- `if !ok { ... }` para ejecutar cuando la condición es falsa.

Ejemplo:

```go
if edad >= 18 && tieneID {
    fmt.Println("Acceso permitido")
}
```

## Operadores bit a bit

- `&` : AND bit a bit
- `|` : OR bit a bit
- `^` : XOR bit a bit
- `&^` : AND NOT bit a bit
- `<<` : desplazamiento a la izquierda
- `>>` : desplazamiento a la derecha

Para qué sirven:
- manipular datos en nivel de bits
- optimizar máscaras y banderas
- trabajar con valores binarios y permisos

Cuándo se usan:
- `mask & value` para extraer bits.
- `flags | feature` para activar banderas.
- `x << 1` para multiplicar por 2 en enteros.

Ejemplo:

```go
const flagRead = 1 << 0
const flagWrite = 1 << 1
flags := flagRead | flagWrite
if flags&flagWrite != 0 {
    fmt.Println("Escribible")
}
```

## Operadores de otras categorías

- `++` : incremento en 1
- `--` : decremento en 1
- `:=` : declaración y asignación corta

Para qué sirven:
- `++` y `--` para ajustar contadores de forma simple.
- `:=` para declarar variables nuevas con tipo inferido.

Cuándo se usan:
- `i++` en bucles.
- `j--` cuando se reduce un índice.
- `x := 42` para declarar `x` automáticamente.

Ejemplo:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## Resumen rápido

- Aritméticos: `+`, `-`, `*`, `/`, `%`
- Asignación: `=`, `+=`, `-=`, `*=`, `/=`, `%=`
- Comparación: `==`, `!=`, `<`, `<=`, `>`, `>=`
- Lógicos: `&&`, `||`, `!`
- Bit a bit: `&`, `|`, `^`, `&^`, `<<`, `>>`
- Otros: `++`, `--`, `:=`

## Consejos prácticos

- Usa `&&` y `||` para condiciones compuestas.
- Evita dividir entre cero con `/`.
- Prefiere `int` y `float64` en cálculos generales.
- Usa operadores bit a bit solo cuando trabajes con banderas, permisos o datos binarios.
