# 14 - Módulos y Paquetes

Cuando tus proyectos crecen, necesitas organizar el código. Los módulos y paquetes lo hacen posible.

## Módulos

Un módulo es un contenedor de código que agrupa funcionalidad relacionada.

### Declarar Módulos

```rust
mod sonidos {
    pub fn grito() {
        println!("¡AAHHH!");
    }
    
    fn susurro() {  // Privado
        println!("psst");
    }
}

fn main() {
    sonidos::grito();     // ✅ Funciona
    // sonidos::susurro();  // ❌ Error: privado
}
```

- `pub` hace visible públicamente
- Sin `pub` es privado (solo visible dentro del módulo)

### Módulos Anidados

```rust
mod sonidos {
    pub mod vocales {
        pub fn vocal_a() {
            println!("aaa");
        }
    }
    
    pub mod consonantes {
        pub fn consonante_t() {
            println!("ttt");
        }
    }
}

fn main() {
    sonidos::vocales::vocal_a();
    sonidos::consonantes::consonante_t();
}
```

### Super (Acceso a Padre)

```rust
fn nivel_raíz() {
    println!("Estoy en raíz");
}

mod sonidos {
    pub fn llamar_raíz() {
        super::nivel_raíz();  // Llama a función del módulo padre
    }
}

fn main() {
    sonidos::llamar_raíz();
}
```

## Separar en Archivos

Para proyectos grandes, separa módulos en archivos:

Estructura:
```
mi_proyecto/
├── src/
│   ├── main.rs
│   ├── sonidos.rs
│   └── animales.rs
```

### main.rs

```rust
mod sonidos;
mod animales;

fn main() {
    sonidos::grito();
    animales::perro::ladrar();
}
```

### sonidos.rs

```rust
pub fn grito() {
    println!("¡AAHHH!");
}
```

### animales.rs

```rust
pub mod perro {
    pub fn ladrar() {
        println!("¡Guau!");
    }
}
```

## Use (Importaciones)

`use` trae items al scope actual:

```rust
mod sonidos {
    pub fn grito() {
        println!("¡AAHHH!");
    }
}

use sonidos::grito;

fn main() {
    grito();  // Sin necesidad del prefijo
}
```

### Importar Múltiples Items

```rust
use sonidos::{grito, susurro};

// O con glob
use sonidos::*;
```

### Renombrar Imports

```rust
use sonidos::grito as gritar;

fn main() {
    gritar();
}
```

### Re-exportar

```rust
pub use sonidos::grito;  // Lo hace público de este módulo
```

## Paquetes (Crates)

Un paquete es una colección de crates (librerías o ejecutables).

### Estructura de Paquete

```
mi_paquete/
├── Cargo.toml
└── src/
    ├── main.rs      # Crate ejecutable
    ├── lib.rs       # Crate librería
    ├── bin/
    │   └── herramienta.rs
    └── modulos/
        ├── mod.rs
        └── utilidades.rs
```

### Cargo.toml para Múltiples Binarios

```toml
[package]
name = "mi_paquete"
version = "0.1.0"

[[bin]]
name = "principal"
path = "src/main.rs"

[[bin]]
name = "herramienta"
path = "src/bin/herramienta.rs"
```

### Crear Librería

```bash
cargo new --lib mi_librería
```

Esto crea `lib.rs` en lugar de `main.rs`.

### lib.rs (Punto de Entrada de Librería)

```rust
pub mod utilidades {
    pub fn sumar(a: i32, b: i32) -> i32 {
        a + b
    }
}

pub fn función_pública() {
    println!("Público");
}
```

### Usar tu Librería

```rust
use mi_librería::utilidades::sumar;

fn main() {
    println!("{}", sumar(5, 3));  // 8
}
```

## Rutas Públicas

Por defecto, solo `main` es público. Para hacer tu librería útil, exporta lo importante:

```rust
// lib.rs
pub mod utilidades;
pub use utilidades::sumar;  // Re-exporta

pub fn función_principal() {
    // ...
}
```

Así los usuarios pueden hacer:

```rust
use mi_librería::sumar;  // Importación directa
```

## Ejemplo Práctico: Calculadora

Estructura:
```
calculadora/
├── src/
│   ├── main.rs
│   ├── lib.rs
│   └── operaciones.rs
```

### lib.rs

```rust
pub mod operaciones;

pub use operaciones::{sumar, restar, multiplicar, dividir};
```

### operaciones.rs

```rust
pub fn sumar(a: f64, b: f64) -> f64 {
    a + b
}

pub fn restar(a: f64, b: f64) -> f64 {
    a - b
}

pub fn multiplicar(a: f64, b: f64) -> f64 {
    a * b
}

pub fn dividir(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err("No se puede dividir por cero".to_string())
    } else {
        Ok(a / b)
    }
}
```

### main.rs

```rust
use calculadora::{sumar, restar, multiplicar, dividir};

fn main() {
    println!("5 + 3 = {}", sumar(5.0, 3.0));
    println!("5 - 3 = {}", restar(5.0, 3.0));
    println!("5 × 3 = {}", multiplicar(5.0, 3.0));
    
    match dividir(10.0, 2.0) {
        Ok(resultado) => println!("10 ÷ 2 = {}", resultado),
        Err(e) => println!("Error: {}", e),
    }
}
```

## Reglas de Privacidad

- Los items son **privados por defecto**
- Solo el módulo padre puede ver items privados de un submódulo
- `pub` los hace públicos
- `pub use` re-exporta para que el padre vea

## Siguiente Paso

Organizaste tu código. El próximo capítulo enseña **pruebas automatizadas** para asegurar calidad.

---

**Consejo**: Piensa en la API que quieres exponer. Usa `pub` estratégicamente. El mejor código es aquel que se entiende a primer vistazo.
