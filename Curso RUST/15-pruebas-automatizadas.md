# 15 - Pruebas Automatizadas

Las pruebas aseguran que tu código funciona como esperas. Rust facilita escribirlas.

## Pruebas Básicas

### Función de Prueba

```rust
fn suma(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_suma() {
        assert_eq!(suma(2, 2), 4);
    }
}
```

- `#[cfg(test)]` compila solo en tests
- `#[test]` marca una función como prueba
- `assert_eq!()` verifica igualdad

### Macros de Aserción

```rust
#[test]
fn pruebas_basicas() {
    // Verdadero
    assert!(true);
    assert_eq!(2 + 2, 4);
    assert_ne!(2 + 2, 5);
    
    // Con mensaje
    assert!(false, "Esto falló porque: {}", "razón");
}
```

### Pruebas que Deben Fallar

```rust
#[test]
#[should_panic]
fn esto_debe_fallar() {
    panic!("¡Panicó!");
}

#[test]
#[should_panic(expected = "mensaje específico")]
fn panic_especifico() {
    panic!("mensaje específico");
}
```

### Pruebas con Result

```rust
#[test]
fn prueba_con_result() -> Result<(), String> {
    if 2 + 2 == 4 {
        Ok(())
    } else {
        Err(String::from("Math is broken"))
    }
}
```

## Estructura de Pruebas

### Arrange, Act, Assert

```rust
fn agregar_dos(a: i32) -> i32 {
    a + 2
}

#[test]
fn test_agregar_dos() {
    // Arrange (Preparar)
    let entrada = 2;
    
    // Act (Actuar)
    let resultado = agregar_dos(entrada);
    
    // Assert (Verificar)
    assert_eq!(resultado, 4);
}
```

## Ejecutar Pruebas

```bash
# Ejecutar todas
cargo test

# Con salida detallada
cargo test -- --nocapture

# Pruebas específicas
cargo test test_suma

# Un solo thread (para debugging)
cargo test -- --test-threads=1

# Mostrar pruebas que pasaron
cargo test -- --nocapture --test-threads=1
```

## Tests de Integración

Para archivos en la carpeta `tests/`:

Estructura:
```
proyecto/
├── src/
│   ├── lib.rs
│   └── main.rs
└── tests/
    └── integracion.rs
```

### tests/integracion.rs

```rust
use proyecto::sumar;

#[test]
fn test_integracion() {
    assert_eq!(sumar(3, 2), 5);
}
```

Se ejecutan con:
```bash
cargo test --test integracion
```

## Ejemplo Completo: Calculadora con Pruebas

```rust
// lib.rs
pub struct Calculadora {
    resultado: f64,
}

impl Calculadora {
    pub fn nueva() -> Self {
        Calculadora { resultado: 0.0 }
    }
    
    pub fn sumar(&mut self, valor: f64) -> f64 {
        self.resultado += valor;
        self.resultado
    }
    
    pub fn restar(&mut self, valor: f64) -> f64 {
        self.resultado -= valor;
        self.resultado
    }
    
    pub fn obtener(&self) -> f64 {
        self.resultado
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_nueva() {
        let calc = Calculadora::nueva();
        assert_eq!(calc.obtener(), 0.0);
    }
    
    #[test]
    fn test_sumar() {
        let mut calc = Calculadora::nueva();
        calc.sumar(5.0);
        assert_eq!(calc.obtener(), 5.0);
    }
    
    #[test]
    fn test_secuencia() {
        let mut calc = Calculadora::nueva();
        calc.sumar(10.0);
        calc.restar(5.0);
        assert_eq!(calc.obtener(), 5.0);
    }
}
```

Ejecuta con `cargo test`.

## Tests Documentados

Los ejemplos en comentarios de documentación también son pruebas:

```rust
/// Suma dos números
/// 
/// # Ejemplos
/// 
/// ```
/// let resultado = mi_librería::sumar(2, 2);
/// assert_eq!(resultado, 4);
/// ```
pub fn sumar(a: i32, b: i32) -> i32 {
    a + b
}
```

Ejecútalos con:
```bash
cargo test --doc
```

## Benchmarks (Avanzado)

Mide performance:

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
    fn bench_fibonacci(b: &mut Bencher) {
        b.iter(|| fibonacci(20));
    }
}
```

Ejecuta con:
```bash
cargo bench
```

## Mocking y Test Fixtures

Usa variables de prueba:

```rust
#[test]
fn test_con_fixture() {
    // Fixture
    let entrada = vec![1, 2, 3];
    
    // Test
    let suma: i32 = entrada.iter().sum();
    
    // Assert
    assert_eq!(suma, 6);
}
```

Para mocking de traits:

```rust
trait Servicio {
    fn obtener_datos(&self) -> String;
}

struct MockServicio;

impl Servicio for MockServicio {
    fn obtener_datos(&self) -> String {
        "datos de prueba".to_string()
    }
}
```

## Siguiente Paso

Las pruebas aseguran confianza en tu código. El próximo capítulo enseña **entrada/salida y trabajo con archivos**.

---

**Consejo**: Escribe pruebas mientras desarrollas, no después. Hará tu código más limpio desde el inicio.
