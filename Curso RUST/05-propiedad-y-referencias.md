# 05 - Propiedad, Préstamo y Referencias

Este es uno de los conceptos más importantes de Rust. ¡Dominálo y entenderás por qué Rust es tan seguro!

## El Sistema de Propiedad (Ownership)

El "ownership" es un conjunto de reglas que gestiona la memoria automáticamente sin recolector de basura:

### Las Tres Reglas

1. **Cada valor tiene un propietario** (owner)
2. **Solo puede haber un propietario al tiempo**
3. **Cuando el propietario se destruye, el valor se libera**

Veamos con un ejemplo:

```rust
fn main() {
    let s1 = String::from("Hola");  // s1 es el propietario
    let s2 = s1;  // ¿Qué pasa aquí?
    
    // println!("{}", s1);  // ❌ Error: s1 ya no es dueño
    println!("{}", s2);  // ✅ s2 es el nuevo dueño
}
```

Cuando haces `let s2 = s1`, el ownership se **transfiere** de s1 a s2. s1 queda inválido.

### ¿Por Qué?

Imagina dos personas con las llaves de la misma casa:
- Si ambas pueden cambiar cosas, ¡caos!
- Rust evita esto: solo **una** persona tiene las llaves a la vez

### Tipos de Copia

Algunos tipos no siguen estas reglas. Los **tipos simples** se copian automáticamente:

```rust
fn main() {
    let x = 5;      // i32 se copia
    let y = x;      // y obtiene una copia, x sigue siendo válido
    
    println!("x = {}, y = {}", x, y);  // ✅ Ambos funcionan
}
```

Tipos que se copian: `i32`, `f64`, `bool`, `char`, etc.
Tipos que se mueven: `String`, `Vec`, etc.

## Funciones y Propiedad

### Transferencia de Propiedad

```rust
fn tomar_propiedad(s: String) {
    println!("Recibí: {}", s);
}  // s se destruye aquí

fn main() {
    let s = String::from("Hola");
    tomar_propiedad(s);
    
    // println!("{}", s);  // ❌ Error: s ya no existe
}
```

### Devolver Propiedad

```rust
fn dar_propiedad() -> String {
    String::from("Mío")
}

fn main() {
    let s = dar_propiedad();  // s obtiene la propiedad
    println!("{}", s);  // ✅ Funciona
}
```

Esto es tedioso si necesitas usar la variable después de llamar una función. ¡Para eso existen las referencias!

## Referencias y Préstamo

Una **referencia** te permite usar un valor sin tomar su propiedad. Es como prestarle algo a alguien; después lo devuelve:

```rust
fn tomar_referencia(s: &String) {  // s es una referencia
    println!("Recibí: {}", s);
}

fn main() {
    let s = String::from("Hola");
    tomar_referencia(&s);  // &s crea una referencia
    
    println!("{}", s);  // ✅ s sigue siendo mío
}
```

El `&` significa "referencia" (pedir prestado).

### Referencias Mutables

Por defecto, las referencias son **inmutables** (no puedes cambiar lo que apuntan). Si necesitas modificar, usa `&mut`:

```rust
fn modificar(s: &mut String) {
    s.push_str(" Mundo!");
}

fn main() {
    let mut s = String::from("Hola");  // Variable mutable
    modificar(&mut s);  // Referencia mutable
    
    println!("{}", s);  // "Hola Mundo!"
}
```

### Reglas de Préstamo

Rust enforce dos reglas:

1. **O tienes muchas referencias inmutables, O una referencia mutable** (no ambas)
2. **Las referencias no pueden vivir más que lo que apuntan**

Ejemplo de violación de la primera regla:

```rust
fn main() {
    let mut s = String::from("Hola");
    
    let r1 = &s;      // ✅ Referencia inmutable 1
    let r2 = &s;      // ✅ Referencia inmutable 2
    let r3 = &mut s;  // ❌ Error: no puedes tener mutable + inmutables
}
```

Solución: usa referencias mutables solo cuando las necesites:

```rust
fn main() {
    let mut s = String::from("Hola");
    
    let r1 = &s;
    let r2 = &s;
    println!("{}, {}", r1, r2);  // Úsalas aquí
    
    let r3 = &mut s;  // ✅ Ahora puedes porque r1 y r2 no se usan más
    r3.push_str(" Mundo!");
    println!("{}", r3);
}
```

## Slices (Rebanadas)

Un slice es una referencia a parte de una colección:

### String Slices

```rust
fn main() {
    let s = String::from("Hola Mundo");
    
    let hola = &s[0..4];      // "Hola" (índices 0, 1, 2, 3)
    let mundo = &s[5..10];    // "Mundo"
    let todo = &s[..];        // Toda la cadena
    
    println!("{}", hola);  // "Hola"
}
```

Más ejemplos:

```rust
fn main() {
    let s = String::from("Hola");
    
    let inicio = &s[..2];     // "Ho"
    let final_ = &s[2..];     // "la"
}
```

### Array Slices

```rust
fn main() {
    let numeros = [1, 2, 3, 4, 5];
    
    let slice = &numeros[1..3];  // [2, 3]
    
    println!("{:?}", slice);  // [2, 3]
}
```

### Función Usando Slices

Una función útil que encuentra la primera palabra:

```rust
fn primera_palabra(s: &str) -> &str {
    let bytes = s.as_bytes();
    
    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    
    &s[..]
}

fn main() {
    let s = String::from("Hola Mundo");
    let palabra = primera_palabra(&s);
    println!("{}", palabra);  // "Hola"
}
```

## Resumen Visual

```rust
// MOVIMIENTO (transferencia de ownership)
let s1 = String::from("Hola");
let s2 = s1;  // s1 ya no es válido

// COPIA (tipos simples)
let x = 5;
let y = x;  // x sigue siendo válido

// REFERENCIA INMUTABLE (pedir prestado sin cambiar)
let s = String::from("Hola");
let r = &s;  // &s es una referencia

// REFERENCIA MUTABLE (pedir prestado para cambiar)
let mut s = String::from("Hola");
let r = &mut s;  // &mut s es una referencia mutable
r.push_str(" Mundo");

// SLICE (referencia a parte de una colección)
let s = String::from("Hola");
let parte = &s[0..2];  // "Ho"
```

## Tabla Comparativa

| Concepto | Símbolo | Cambios | Múltiples |
|----------|---------|---------|-----------|
| Propiedad | `let x = ...` | ✅ Sí (mut) | ❌ No |
| Ref. Inmutable | `&x` | ❌ No | ✅ Sí |
| Ref. Mutable | `&mut x` | ✅ Sí | ❌ No |

## Siguiente Paso

Ahora que entiendes ownership, aprenderás a organizar datos con **structs** (estructuras) y **enums** (enumeraciones), que son formas poderosas de agrupar información.

---

**Consejo**: Si no entiendes ownership en la primera lectura, ¡no te preocupes! Es el concepto más desafiante. Practica con ejemplos hasta que te sienta natural.
