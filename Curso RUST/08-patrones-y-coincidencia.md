# 09 - Patrones y Pattern Matching

El pattern matching es una característica poderosa de Rust que permite descomponer y analizar valores.

## El Operador match

Ya lo vimos, pero profundicemos:

```rust
fn main() {
    let número = 3;
    
    match número {
        1 => println!("Uno"),
        2 => println!("Dos"),
        3 => println!("Tres"),
        _ => println!("Otro"),  // _ captura todos los demás casos
    }
}
```

El `_` es un pattern que "ignora" el valor.

### Patterns con Range

```rust
fn main() {
    let número = 5;
    
    match número {
        1..=3 => println!("De 1 a 3"),
        4..=6 => println!("De 4 a 6"),
        _ => println!("Mayor a 6"),
    }
}
```

### Patterns Destructuring en match

```rust
#[derive(Debug)]
enum Color {
    Rgb(i32, i32, i32),
    Hsv(i32, i32, i32),
}

fn main() {
    let color = Color::Rgb(255, 0, 128);
    
    match color {
        Color::Rgb(r, g, b) => {
            println!("Rojo: {}, Verde: {}, Azul: {}", r, g, b);
        }
        Color::Hsv(h, s, v) => {
            println!("Matiz: {}, Saturación: {}, Valor: {}", h, s, v);
        }
    }
}
```

### Guard (Condiciones en Patterns)

```rust
fn main() {
    let número = 4;
    
    match número {
        x if x < 5 => println!("Menor que 5"),
        x if x == 5 => println!("Es 5"),
        _ => println!("Mayor que 5"),
    }
}
```

### Pattern con OR

```rust
fn main() {
    let número = 1;
    
    match número {
        1 | 2 => println!("Uno o Dos"),
        3 | 4 => println!("Tres o Cuatro"),
        _ => println!("Otro"),
    }
}
```

## Destructuring en Asignaciones

### Destructuring de Tuplas

```rust
fn main() {
    let tupla = (5, 6.4, "Hola");
    
    let (x, y, z) = tupla;
    println!("x: {}, y: {}, z: {}", x, y, z);
    
    // Ignorar algunos valores
    let (x, _, z) = tupla;
}
```

### Destructuring de Structs

```rust
struct Usuario {
    nombre: String,
    edad: u32,
}

fn main() {
    let usuario = Usuario {
        nombre: String::from("Ana"),
        edad: 30,
    };
    
    let Usuario { nombre, edad } = usuario;
    println!("Nombre: {}, Edad: {}", nombre, edad);
    
    // Con renombre
    let Usuario { nombre: n, edad: a } = usuario;
}
```

### Destructuring Anidado

```rust
enum Mensaje {
    Cambiar { x: i32, y: i32 },
}

fn main() {
    let msg = Mensaje::Cambiar { x: 10, y: 20 };
    
    match msg {
        Mensaje::Cambiar { x: a, y: b } => {
            println!("Posición: {}, {}", a, b);
        }
    }
}
```

## Patrones de Ignorado

```rust
fn main() {
    let (x, _, z) = (1, 2, 3);  // Ignora el medio
    
    let _y = 5;  // Ignora esta variable (no da warning)
    
    let s = Some(String::from("Hola"));
    if let Some(_) = s {  // Ignora el valor
        println!("Hay algo");
    }
    
    match (1, 2, 3) {
        (1, _, _) => println!("Empieza con 1"),
        _ => println!("Otro patrón"),
    }
}
```

## Patrones Ref

Obtener referencias sin ownership:

```rust
fn main() {
    let punto = (1, 2);
    
    match punto {
        (ref x, ref y) => {
            println!("Punto: {}, {}", x, y);
            // x e y son referencias
        }
    }
}
```

## Patrones Binding

```rust
fn main() {
    enum Mensaje {
        Escribir(String),
    }
    
    let msg = Mensaje::Escribir(String::from("Hola"));
    
    match msg {
        Mensaje::Escribir(text) => {
            println!("Escribir: {}", text);  // text tiene el String
        }
    }
}
```

## Ejemplo Completo: Parser Sencillo

```rust
enum Token {
    Número(i32),
    Operador(char),
    Fin,
}

fn evaluar_token(token: Token) {
    match token {
        Token::Número(n) => println!("Número: {}", n),
        Token::Operador('+') => println!("Suma"),
        Token::Operador('-') => println!("Resta"),
        Token::Operador(op) => println!("Operador: {}", op),
        Token::Fin => println!("Fin del programa"),
    }
}

fn main() {
    let tokens = vec![
        Token::Número(5),
        Token::Operador('+'),
        Token::Número(3),
        Token::Fin,
    ];
    
    for token in tokens {
        evaluar_token(token);
    }
}
```

## Resumen de Patterns

| Pattern | Ejemplo | Descripción |
|---------|---------|-------------|
| Literal | `1`, `"hola"` | Valor exacto |
| Range | `1..=5` | Rango de valores |
| Destructuring | `(x, y)` | Separar estructura |
| Wildcard | `_` | Ignorar valor |
| Guard | `x if x > 5` | Condición adicional |
| OR | `1 \| 2` | Múltiples opciones |
| Ref | `ref x` | Obtener referencia |
| Bind | `x @ 1..=5` | Capturar y vincular |

## Siguiente Paso

Ahora dominas los patrones. El próximo capítulo introduce **genéricos y traits**, dos conceptos que hacen a Rust extremadamente poderoso.

---

**Consejo**: El pattern matching es expresivo. Úsalo frecuentemente para hacer tu código más limpio y seguro.
