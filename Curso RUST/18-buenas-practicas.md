# 18 - Buenas Prácticas en Rust

Este capítulo resume patrones y estrategias para escribir código Rust profesional.

## 1. Naming Conventions (Convenciones de Nombres)

```rust
// Variables: snake_case
let mi_variable = 5;

// Funciones: snake_case
fn calcular_promedio() {}

// Tipos: PascalCase
struct MiEstructura {}
enum MiEnum {}

// Constantes: SCREAMING_SNAKE_CASE
const MAX_CONEXIONES: u32 = 100;

// Métodos privados: _snake_case
fn _funcion_interna() {}

// Generics: Letras mayúsculas simples
fn procesar<T>(valor: T) {}
```

## 2. Organización del Código

### Estructura de Carpetas

```
proyecto/
├── src/
│   ├── main.rs          # Punto de entrada
│   ├── lib.rs           # API pública
│   ├── utils/           # Utilidades generales
│   │   ├── mod.rs
│   │   ├── logger.rs
│   │   └── validators.rs
│   ├── models/          # Estructuras de datos
│   │   ├── mod.rs
│   │   ├── usuario.rs
│   │   └── producto.rs
│   └── handlers/        # Lógica de negocio
│       ├── mod.rs
│       └── procesamiento.rs
├── tests/               # Tests de integración
├── examples/            # Ejemplos de uso
├── Cargo.toml
└── README.md
```

### Organizar Módulos

```rust
// lib.rs
pub mod utils;
pub mod models;
pub mod handlers;

// Exportar APIs principales
pub use models::Usuario;
pub use handlers::procesar;
```

## 3. Error Handling

### Nunca Ignores Errores

```rust
// ❌ Malo
let _ = archivo.read_to_string(&mut contenido);

// ✅ Bueno
match archivo.read_to_string(&mut contenido) {
    Ok(n) => println!("Leí {} bytes", n),
    Err(e) => eprintln!("Error: {}", e),
}
```

### Propaga Errores Apropiadamente

```rust
// ❌ Malo: usar unwrap() en código de producción
let datos = archivo.read_to_string(&mut s).unwrap();

// ✅ Bueno: propagar al caller
fn leer_config(ruta: &str) -> Result<String, std::io::Error> {
    std::fs::read_to_string(ruta)
}

// ✅ Muy bueno: con contexto
fn leer_config(ruta: &str) -> Result<String, Box<dyn std::error::Error>> {
    let contenido = std::fs::read_to_string(ruta)
        .map_err(|e| format!("No se pudo leer {}: {}", ruta, e))?;
    Ok(contenido)
}
```

### Crear Tipos de Error Personalizados

```rust
use std::fmt;

#[derive(Debug)]
pub enum AppError {
    IoError(std::io::Error),
    ParseError(String),
    ValidationError(String),
}

impl fmt::Display for AppError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            AppError::IoError(e) => write!(f, "Error de I/O: {}", e),
            AppError::ParseError(msg) => write!(f, "Parse error: {}", msg),
            AppError::ValidationError(msg) => write!(f, "Validación falló: {}", msg),
        }
    }
}

impl std::error::Error for AppError {}

pub type Result<T> = std::result::Result<T, AppError>;
```

## 4. Immutability First

```rust
// ❌ Usas mut cuando no lo necesitas
let mut contador = 0;
contador += 1;
println!("{}", contador);

// ✅ Mejor: sin mut
let contador = 1;
println!("{}", contador);

// ✅ Mejor aún: si necesitas cambiar
let contador = 0;
let contador = contador + 1;  // Shadowing
```

## 5. Comentarios Documentados

```rust
/// Calcula el factorial de un número.
///
/// # Panics
/// Panics si el número es mayor a 20 (para evitar overflow).
///
/// # Examples
///
/// ```
/// assert_eq!(factorial(5), 120);
/// ```
pub fn factorial(n: u32) -> u32 {
    assert!(n <= 20, "Factorial muy grande");
    match n {
        0 | 1 => 1,
        _ => n * factorial(n - 1),
    }
}

/// Modifica el valor internamente.
///
/// # Arguments
///
/// * `valor` - El nuevo valor
pub fn establecer_valor(&mut self, valor: i32) {
    self.valor = valor;
}
```

## 6. Type Safety

```rust
// ❌ Tipo genérico, menos claro
fn procesar(valor: i32) -> i32 {
    valor * 2
}

// ✅ Mejor: tipos específicos con newtype pattern
#[derive(Debug, Clone, Copy)]
pub struct UserId(u32);

#[derive(Debug, Clone)]
pub struct UserName(String);

fn buscar_usuario(id: UserId) -> Option<UserName> {
    // Ahora es imposible mezclar ID con nombre
    None
}
```

## 7. RAII (Resource Acquisition Is Initialization)

Usa `Drop` para limpiar recursos automáticamente:

```rust
struct Archivo {
    manejador: std::fs::File,
}

impl Drop for Archivo {
    fn drop(&mut self) {
        // La lógica de limpieza ocurre automáticamente
        println!("Archivo cerrado");
    }
}

fn main() {
    let _archivo = Archivo {
        manejador: std::fs::File::create("temp.txt").unwrap(),
    };
    // Se limpia automáticamente al salir de scope
}
```

## 8. Testing

Escribe tests al lado del código:

```rust
pub fn sumar(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_sumar_positivos() {
        assert_eq!(sumar(2, 2), 4);
    }
    
    #[test]
    fn test_sumar_negativos() {
        assert_eq!(sumar(-1, -1), -2);
    }
}
```

## 9. Use Performance Wisely

```rust
// ❌ Copia innecesaria
fn procesar(s: String) -> String {
    format!("Procesado: {}", s)
}

// ✅ Mejor: usa referencias
fn procesar(s: &str) -> String {
    format!("Procesado: {}", s)
}

// ✅ Para grandes estructuras
fn procesar(datos: &[u8]) {
    // Evita copiar
}
```

## 10. Logging

Usa `log` en librerías, `env_logger` en binarios:

```toml
[dependencies]
log = "0.4"
env_logger = "0.11"
```

```rust
use log::{debug, info, warn, error};

fn procesar_usuario(id: u32) {
    debug!("Buscando usuario {}", id);
    info!("Usuario encontrado");
    warn!("Acceso a datos sensibles");
    error!("Falló la verificación");
}

fn main() {
    env_logger::init();
    procesar_usuario(123);
}
```

Ejecuta con:
```bash
RUST_LOG=debug cargo run
```

## 11. Validación de Entrada

```rust
pub struct Email(String);

impl Email {
    pub fn nuevo(email: &str) -> Result<Self, String> {
        if email.contains('@') && email.len() > 5 {
            Ok(Email(email.to_string()))
        } else {
            Err("Email inválido".to_string())
        }
    }
}

fn main() {
    match Email::nuevo("usuario@example.com") {
        Ok(email) => println!("Email válido: {}", email.0),
        Err(e) => println!("Error: {}", e),
    }
}
```

## 12. Documentation y Examples

```bash
# Generar documentación
cargo doc --open

# Ejecutar tests de documentación
cargo test --doc
```

## 13. Use External Crates Wisely

```toml
[dependencies]
serde = { version = "1.0", features = ["derive"] }
tokio = { version = "1", features = ["full"] }
anyhow = "1.0"  # Para errores rápidos
```

## 14. Clippy (Linter)

```bash
cargo clippy
```

Sigue sus sugerencias para código más idiomático.

## 15. Code Formatting

```bash
# Formatea automáticamente
cargo fmt

# Verifica formato
cargo fmt -- --check
```

Configura en `rustfmt.toml`:

```toml
max_width = 100
tab_spaces = 4
edition = "2024"
```

## 16. Benchmarking

```rust
#![feature(test)]
extern crate test;

fn fibonacci(n: u32) -> u32 {
    match n {
        0 | 1 => 1,
        _ => fibonacci(n - 1) + fibonacci(n - 2),
    }
}

#[cfg(test)]
mod benches {
    use super::*;
    use test::Bencher;
    
    #[bench]
    fn bench_fib(b: &mut Bencher) {
        b.iter(|| fibonacci(20));
    }
}
```

## 17. Concurrency Best Practices

```rust
use std::sync::{Arc, Mutex};
use std::thread;

// ✅ Bien: Arc + Mutex para compartir datos
let datos = Arc::new(Mutex::new(0));

for _ in 0..4 {
    let datos_clone = Arc::clone(&datos);
    thread::spawn(move || {
        let mut d = datos_clone.lock().unwrap();
        *d += 1;
    });
}
```

## 18. Security Considerations

```rust
// ✅ Valida entrada siempre
fn procesar_comando(cmd: &str) -> Result<(), String> {
    if cmd.contains("delete") && cmd.contains("--force") {
        Err("Comando peligroso bloqueado".to_string())
    } else {
        Ok(())
    }
}

// ✅ Usa tipos seguros
pub struct SecureString(String);

impl Drop for SecureString {
    fn drop(&mut self) {
        // Limpiar datos sensibles de memoria
        self.0.clear();
    }
}
```

## Checklist Final

Antes de publicar tu código, verifica:

- [ ] ✅ No hay `unwrap()` innecesarios en código de producción
- [ ] ✅ Todos los errores son manejados
- [ ] ✅ Los nombres siguen las convenciones Rust
- [ ] ✅ Hay tests para funcionalidad crítica
- [ ] ✅ La documentación es clara
- [ ] ✅ Pasó `cargo clippy` sin warnings
- [ ] ✅ El código está formateado (`cargo fmt`)
- [ ] ✅ Los datos sensibles se limpian
- [ ] ✅ La entrada se valida
- [ ] ✅ Los recursos se limpian automáticamente

## Recursos Adicionales

- [Rust API Guidelines](https://rust-lang.github.io/api-guidelines/)
- [Rust by Example](https://doc.rust-lang.org/rust-by-example/)
- [The Rustonomicon](https://doc.rust-lang.org/nomicon/) (para unsafe)
- [Concurrency Patterns](https://doc.rust-lang.org/book/ch16-00-concurrency.html)

## Conclusión

Rust es un lenguaje poderoso que te obliga a escribir código seguro y eficiente. Estas buenas prácticas te ayudarán a escribir código que no solo compila, sino que es mantenible, eficiente y seguro.

### Los Tres Pilares

1. **Seguridad**: El compilador te ayuda
2. **Velocidad**: Sin overhead de garbage collection
3. **Concurrencia**: Sin race conditions

Bienvenido a la comunidad Rust. ¡Ahora eres un Rustacean! 🦀

---

**Última recomendación**: Lee el código de otros. La comunidad Rust es excelente. Explora crates populares en [crates.io](https://crates.io/) y aprende de sus patrones.
