# Biblioteca `math` en Go

La biblioteca `math` proporciona funciones matemáticas básicas y avanzadas para operaciones numéricas.

## Funciones principales

### Raíz y potencia
- `math.Sqrt(x)` : raíz cuadrada de `x`.
- `math.Sqrt(x)` : raíz cuadrada de `x`.
- `math.Cbrt(x)` : raíz cúbica de `x`.
- `math.Pow(x, y)` : `x` elevado a la potencia `y`.
- `math.Pow10(n)` : 10 elevado a la potencia `n`.
- `math.Exp(x)` : e elevado a la potencia `x`.
- `math.Exp2(x)` : 2 elevado a la potencia `x`.
- `math.Log(x)` : logaritmo natural de `x`.
- `math.Log10(x)` : logaritmo base 10 de `x`.
- `math.Log2(x)` : logaritmo base 2 de `x`.
- `math.Log1p(x)` : logaritmo natural de `1 + x`.

### Trigonometría
- `math.Sin(x)` : seno de `x` (radianes).
- `math.Cos(x)` : coseno de `x` (radianes).
- `math.Tan(x)` : tangente de `x` (radianes).
- `math.Asin(x)` : arcoseno de `x`.
- `math.Acos(x)` : arcocoseno de `x`.
- `math.Atan(x)` : arcotangente de `x`.
- `math.Atan2(y, x)` : arcotangente de `y/x` con consideración de cuadrante.
- `math.Sinh(x)` : seno hiperbólico de `x`.
- `math.Cosh(x)` : coseno hiperbólico de `x`.
- `math.Tanh(x)` : tangente hiperbólica de `x`.
- `math.Asinh(x)` : arcoseno hiperbólico de `x`.
- `math.Acosh(x)` : arcocoseno hiperbólico de `x`.
- `math.Atanh(x)` : arcotangente hiperbólica de `x`.

### Valor absoluto y redondeo
- `math.Abs(x)` : valor absoluto de `x`.
- `math.Floor(x)` : redondea hacia abajo.
- `math.Ceil(x)` : redondea hacia arriba.
- `math.Trunc(x)` : elimina la parte decimal.
- `math.Round(x)` : redondea al entero más cercano.
- `math.RoundToEven(x)` : redondea al entero más cercano (par si es intermedio).

### Máximo y mínimo
- `math.Max(x, y)` : máximo entre `x` e `y`.
- `math.Min(x, y)` : mínimo entre `x` e `y`.
- `math.Dim(x, y)` : diferencia positiva (máximo de `x - y` y 0).

### Mod y residuo
- `math.Mod(x, y)` : residuo de `x / y`.
- `math.Remainder(x, y)` : residuo IEEE 754 de `x / y`.
- `math.Fmod(x, y)` : residuo con signo de `x / y`.

### Manipulación de bits y valores especiales
- `math.Copysign(x, y)` : copia el signo de `y` a `x`.
- `math.Signbit(x)` : determina si `x` es negativo.
- `math.Ldexp(frac, exp)` : calcula `frac * 2^exp`.
- `math.Frexp(x)` : descompone `x` en fracción y exponente.
- `math.Modf(x)` : separa la parte entera y decimal de `x`.

### Valores especiales
- `math.NaN()` : retorna "Not a Number".
- `math.Inf(sign)` : retorna infinito con signo.
- `math.IsNaN(x)` : verifica si `x` es NaN.
- `math.IsInf(x, sign)` : verifica si `x` es infinito.
- `math.IsFinite(x)` : verifica si `x` es finito.

### Constantes
- `math.Pi` : π (3.14159...).
- `math.E` : e (2.71828...).
- `math.Phi` : número áureo (1.61803...).
- `math.Sqrt2` : √2 (1.41421...).
- `math.SqrtE` : √e (1.64872...).
- `math.SqrtPi` : √π (1.77245...).
- `math.SqrtPhi` : √φ (1.27201...).
- `math.Ln2` : ln(2) (0.69314...).
- `math.Log2E` : log₂(e) (1.44269...).
- `math.Ln10` : ln(10) (2.30258...).
- `math.Log10E` : log₁₀(e) (0.43429...).
- `math.MaxFloat64` : valor máximo de `float64`.
- `math.MinFloat64` : valor mínimo positivo de `float64`.
- `math.MaxInt` : valor máximo de `int`.
- `math.MinInt` : valor mínimo de `int`.

## Ejemplos

### Raíz cuadrada y potencia
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Raíz cuadrada
    fmt.Println("√16:", math.Sqrt(16))
    fmt.Println("√2:", math.Sqrt(2))
    
    // Raíz cúbica
    fmt.Println("∛8:", math.Cbrt(8))
    
    // Potencia
    fmt.Println("2^8:", math.Pow(2, 8))
    fmt.Println("10^3:", math.Pow10(3))
    fmt.Println("e^2:", math.Exp(2))
    fmt.Println("2^4:", math.Exp2(4))
}
```

### Logaritmos
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Logaritmos
    fmt.Println("ln(10):", math.Log(10))
    fmt.Println("log₁₀(100):", math.Log10(100))
    fmt.Println("log₂(8):", math.Log2(8))
    fmt.Println("ln(1 + 0.5):", math.Log1p(0.5))
}
```

### Trigonometría básica
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    angulo := math.Pi / 4 // 45 grados
    
    fmt.Println("sin(45°):", math.Sin(angulo))
    fmt.Println("cos(45°):", math.Cos(angulo))
    fmt.Println("tan(45°):", math.Tan(angulo))
    
    // Funciones inversas
    fmt.Println("asin(0.7):", math.Asin(0.7))
    fmt.Println("acos(0.7):", math.Acos(0.7))
    fmt.Println("atan(1):", math.Atan(1))
    fmt.Println("atan2(1, 1):", math.Atan2(1, 1))
}
```

### Trigonometría hiperbólica
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    x := 0.5
    
    fmt.Println("sinh(0.5):", math.Sinh(x))
    fmt.Println("cosh(0.5):", math.Cosh(x))
    fmt.Println("tanh(0.5):", math.Tanh(x))
    
    fmt.Println("asinh(0.5):", math.Asinh(x))
    fmt.Println("acosh(1.5):", math.Acosh(1.5))
    fmt.Println("atanh(0.5):", math.Atanh(0.5))
}
```

### Redondeo y valor absoluto
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    x := 3.7
    y := -4.2
    
    fmt.Println("Floor(3.7):", math.Floor(x))
    fmt.Println("Ceil(3.7):", math.Ceil(x))
    fmt.Println("Trunc(3.7):", math.Trunc(x))
    fmt.Println("Round(3.7):", math.Round(x))
    
    fmt.Println("Abs(-4.2):", math.Abs(y))
}
```

### Máximo, mínimo y diferencia
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    a, b := 10.5, 7.3
    
    fmt.Println("Max(10.5, 7.3):", math.Max(a, b))
    fmt.Println("Min(10.5, 7.3):", math.Min(a, b))
    fmt.Println("Dim(10.5, 7.3):", math.Dim(a, b))
}
```

### Operaciones con módulo
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    x, y := 10.5, 3.0
    
    fmt.Println("Mod(10.5, 3):", math.Mod(x, y))
    fmt.Println("Remainder(10.5, 3):", math.Remainder(x, y))
    fmt.Println("Fmod(10.5, 3):", math.Fmod(x, y))
}
```

### Separar partes entera y decimal
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    x := 3.75
    
    // Modf retorna la parte entera e.g. (3.0, 0.75)
    entero, decimal := math.Modf(x)
    fmt.Println("Modf(3.75):", entero, "+", decimal)
    
    // Frexp descompone en fracción y exponente
    frac, exp := math.Frexp(8.0)
    fmt.Println("Frexp(8.0):", frac, "* 2^", exp)
    
    // Ldexp reconstruye
    resultado := math.Ldexp(frac, exp)
    fmt.Println("Ldexp(0.5, 4):", resultado)
}
```

### Valores especiales
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    nan := math.NaN()
    inf := math.Inf(1)
    
    fmt.Println("NaN:", nan)
    fmt.Println("Infinito positivo:", inf)
    fmt.Println("Infinito negativo:", math.Inf(-1))
    
    fmt.Println("¿Es NaN?", math.IsNaN(nan))
    fmt.Println("¿Es Inf?", math.IsInf(inf, 1))
    fmt.Println("¿Es finito 5.0?", math.IsFinite(5.0))
}
```

### Constantes matemáticas
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("π:", math.Pi)
    fmt.Println("e:", math.E)
    fmt.Println("√2:", math.Sqrt2)
    fmt.Println("√π:", math.SqrtPi)
    fmt.Println("Número áureo:", math.Phi)
    fmt.Println("ln(2):", math.Ln2)
    fmt.Println("ln(10):", math.Ln10)
    fmt.Println("log₂(e):", math.Log2E)
    fmt.Println("log₁₀(e):", math.Log10E)
}
```

## Buenas prácticas
- Usa `math.IsNaN()` e `math.IsInf()` para validar resultados antes de usar valores especiales.
- Convierte grados a radianes usando `grados * math.Pi / 180` para las funciones trigonométricas.
- Usa `math.Abs()` para comparar números con precisión limitada.
- Prefiere `math.Remainder()` sobre `math.Mod()` para operaciones matemáticas precisas.
