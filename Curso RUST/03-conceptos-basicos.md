# 03 - Conceptos Básicos: Variables, Tipos y Operaciones

## Variables y Mutabilidad

En Rust, las variables son **inmutables por defecto**. Esto significa que una vez que asignas un valor, no puedes cambiarlo.

### Variables Inmutables

```rust
fn main() {
    let x = 5;
    println!("El valor de x es: {}", x);
    
    // x = 6;  // ❌ Error! x es inmutable
}
```

### Variables Mutables

Si necesitas cambiar el valor, declara la variable como `mut`:

```rust
fn main() {
    let mut x = 5;
    println!("El valor de x es: {}", x);
    
    x = 6;  // ✅ Funciona
    println!("El valor de x ahora es: {}", x);
}
```

Output:
```
El valor de x es: 5
El valor de x ahora es: 6
```

### ¿Por Qué Inmutable por Defecto?

Rust hace esto por **seguridad**. Las variables inmutables previenen errores accidentales donde cambias un valor sin darte cuenta.

**Regla de oro**: Usa `let` para valores que no cambiarán, y `let mut` cuando necesites modificar.

## Constantes

Las constantes son **siempre inmutables** y deben tener el tipo explícito:

```rust
const VELOCIDAD_LUZ: u32 = 299792458;

fn main() {
    println!("Velocidad de la luz: {} m/s", VELOCIDAD_LUZ);
}
```

Diferencias con variables:
- Usas `const` en lugar de `let`
- Debes especificar el tipo explícitamente (`: u32`)
- Solo aceptan valores constantes (nada calculado en tiempo de ejecución)
- Se usan para valores que nunca cambian

## Tipos de Datos Básicos

Rust tiene cuatro categorías principales de tipos escalares (valores únicos):

### 1. Enteros (Integers)

Números sin punto decimal. Hay dos variantes:

**Sin signo** (solo positivos):
```rust
let u8_var: u8 = 255;     // 0 a 255
let u16_var: u16 = 65535;
let u32_var: u32 = 4294967295;
let u64_var: u64 = 18446744073709551615;
let u128_var: u128 = 340282366920938463463374607431768211455;
let usize_var: usize = 999;  // Tamaño del pointer (32 o 64 bits)
```

**Con signo** (positivos y negativos):
```rust
let i8_var: i8 = -128;      // -128 a 127
let i16_var: i16 = -32768;
let i32_var: i32 = -2147483648;
let i64_var: i64 = -9223372036854775808;
let i128_var: i128 = -170141183460469231731687303715884105728;
let isize_var: isize = -999;  // Tamaño del pointer con signo
```

**Tipo por defecto**: `i32`

Ejemplos prácticos:

```rust
fn main() {
    let x = 5;           // Tipo inferido: i32
    let y: i64 = 5;      // Especificado: i64
    let z = 98_222;      // Separador visual (equivale a 98222)
    let hex = 0xff;      // Hexadecimal
    let octal = 0o77;    // Octal
    let binario = 0b1111_0000;  // Binario
}
```

### 2. Decimales (Floating Point)

Números con punto decimal:

```rust
fn main() {
    let x = 2.0;          // f64 (por defecto)
    let y: f32 = 3.0;     // f32 explícito
}
```

Tipos:
- `f32` - 32 bits, precisión simple
- `f64` - 64 bits, precisión doble (DEFAULT)

Úsalos para cálculos matemáticos:

```rust
fn main() {
    let pi = 3.14159;
    let diametro = 10.0;
    let circunferencia = pi * diametro;
    println!("Circunferencia: {}", circunferencia);
}
```

### 3. Booleanos

Valores verdadero/falso:

```rust
fn main() {
    let verdadero = true;
    let falso = false;
    
    let resultado: bool = 5 > 3;  // true
    println!("¿5 es mayor que 3? {}", resultado);
}
```

El tipo es `bool` y solo tiene dos valores: `true` y `false`.

### 4. Caracteres

Un carácter único (Unicode):

```rust
fn main() {
    let c = 'z';
    let emoji = '😻';
    let numero = '3';
    
    println!("Carácter: {}, Emoji: {}, Número: {}", c, emoji, numero);
}
```

**Importante**: Usa comillas simples `'` para caracteres (`char`), no comillas dobles `"` (eso es para strings).

## Operaciones Aritméticas

```rust
fn main() {
    // Suma
    let suma = 5 + 6;
    println!("5 + 6 = {}", suma);  // 11
    
    // Resta
    let resta = 95.5 - 4.3;
    println!("95.5 - 4.3 = {}", resta);  // 91.2
    
    // Multiplicación
    let multiplicacion = 4 * 30;
    println!("4 * 30 = {}", multiplicacion);  // 120
    
    // División
    let division = 56.7 / 32.2;
    println!("56.7 / 32.2 = {}", division);  // 1.7608...
    
    // Módulo (resto)
    let modulo = 43 % 5;
    println!("43 % 5 = {}", modulo);  // 3
}
```

### División Entera

```rust
fn main() {
    let x = 10 / 3;
    println!("{}", x);  // 3 (no 3.333...)
}
```

Cuando divides enteros, el resultado es entero.

## Operaciones Comparación y Lógica

```rust
fn main() {
    // Comparación
    println!("5 == 5: {}", 5 == 5);      // true
    println!("5 != 6: {}", 5 != 6);      // true
    println!("5 < 6: {}", 5 < 6);        // true
    println!("5 <= 5: {}", 5 <= 5);      // true
    println!("5 > 4: {}", 5 > 4);        // true
    println!("5 >= 5: {}", 5 >= 5);      // true
    
    // Lógica
    println!("true && false: {}", true && false);  // false
    println!("true || false: {}", true || false);  // true
    println!("!true: {}", !true);                  // false
}
```

## Shadowing (Enmascaramiento)

Puedes declarar una nueva variable con el mismo nombre. La nueva "oculta" a la anterior:

```rust
fn main() {
    let x = 5;
    println!("x = {}", x);  // 5
    
    let x = x + 1;  // Nueva variable x
    println!("x = {}", x);  // 6
    
    let x = x * 2;
    println!("x = {}", x);  // 12
}
```

Esto es útil cuando quieres cambiar el tipo:

```rust
fn main() {
    let espacios = "   ";
    let espacios = espacios.len();  // Cambió de string a número
    println!("Espacios: {}", espacios);  // 3
}
```

## Type Casting (Conversión de Tipos)

```rust
fn main() {
    let x: i32 = 5;
    let y: i64 = x as i64;
    
    let flotante: f32 = 3.14;
    let entero: i32 = flotante as i32;  // 3
    
    println!("y = {}", y);  // 5
    println!("entero = {}", entero);  // 3
}
```

Usa `as` para convertir entre tipos.

## Resumen

| Concepto | Ejemplo |
|----------|---------|
| Variable inmutable | `let x = 5;` |
| Variable mutable | `let mut x = 5;` |
| Constante | `const MAX: u32 = 100;` |
| Entero | `let x: i32 = 42;` |
| Decimal | `let x: f64 = 3.14;` |
| Booleano | `let x: bool = true;` |
| Carácter | `let x: char = 'A';` |

## Siguiente Paso

Ahora que entiendes variables y tipos, aprenderás a crear funciones y controlar el flujo de tu programa con condicionales y bucles.

---

**Consejo**: Si ves un error de tipo, Rust te lo dirá claramente. Lee el mensaje de error, que Rust es muy amable con sus mensajes.
