# 08 - Manejo de Errores

Rust no usa excepciones. En su lugar, usa el tipo `Result` para errores recuperables. Los errores no recuperables usan `panic!`.

## Errores No Recuperables: panic!

Un panic es un error del cual el programa no se puede recuperar:

```rust
fn main() {
    panic!("¡Error crítico!");
}
```

Output:
```
thread 'main' panicked at '¡Error crítico!'
```

### Cuando Ocurren Panics

Algunos panics son implícitos:

```rust
fn main() {
    let números = vec![1, 2, 3];
    
    println!("{}", números[10]);  // ❌ Panic: index out of bounds
}
```

### Para Debugging

```rust
fn main() {
    let valor = Some(5);
    
    valor.unwrap();  // Devuelve 5 si Some, panic si None
    
    valor.expect("Debería tener un valor");  // Panic con mensaje personalizado
}
```

## Errores Recuperables: Result

Para errores que tu código podría manejar:

```rust
enum Result<T, E> {
    Ok(T),      // Éxito con valor
    Err(E),     // Error
}
```

### Ejemplo Básico

```rust
use std::fs;

fn main() {
    let contenido = fs::read_to_string("archivo.txt");
    
    match contenido {
        Ok(datos) => println!("Contenido: {}", datos),
        Err(error) => println!("Error: {}", error),
    }
}
```

### Propagar Errores

En lugar de manejar cada error, puedes propagarlo al caller:

```rust
use std::fs;

fn leer_archivo() -> Result<String, std::io::Error> {
    fs::read_to_string("archivo.txt")
}

fn main() {
    match leer_archivo() {
        Ok(contenido) => println!("Éxito: {}", contenido),
        Err(e) => println!("Error: {}", e),
    }
}
```

### El Operador `?` (El Quesiónador)

Es una forma elegante de propagar errores:

```rust
use std::fs;

fn leer_archivo() -> Result<String, std::io::Error> {
    let contenido = fs::read_to_string("archivo.txt")?;
    Ok(contenido)
}

fn main() {
    match leer_archivo() {
        Ok(contenido) => println!("Éxito: {}", contenido),
        Err(e) => println!("Error: {}", e),
    }
}
```

El `?` devuelve el error si ocurre, o continúa con el valor si es Ok.

### Métodos Útiles de Result

```rust
fn main() {
    let resultado: Result<i32, &str> = Ok(5);
    
    // map: transforma el valor Ok
    let duplicado = resultado.map(|x| x * 2);  // Ok(10)
    
    // map_err: transforma el error
    let resultado2: Result<i32, String> = 
        Err("algo salió mal").map_err(|e| e.to_string());
    
    // unwrap_or: valor por defecto si error
    let valor = Err("error").unwrap_or(0);  // 0
    
    // is_ok / is_err
    if resultado.is_ok() {
        println!("Es Ok");
    }
}
```

### if let con Result

```rust
fn main() {
    let resultado: Result<i32, &str> = Ok(5);
    
    if let Ok(valor) = resultado {
        println!("Valor: {}", valor);
    }
}
```

## Crear Tipos de Error Personalizados

```rust
use std::fmt;

#[derive(Debug)]
enum ErrorPersonalizado {
    NoEncontrado,
    PermisoDenegado,
    FormatoInvalido,
}

impl fmt::Display for ErrorPersonalizado {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Self::NoEncontrado => write!(f, "Recurso no encontrado"),
            Self::PermisoDenegado => write!(f, "Permiso denegado"),
            Self::FormatoInvalido => write!(f, "Formato inválido"),
        }
    }
}

impl std::error::Error for ErrorPersonalizado {}

fn buscar_archivo(nombre: &str) -> Result<String, ErrorPersonalizado> {
    if nombre.is_empty() {
        return Err(ErrorPersonalizado::NoEncontrado);
    }
    Ok(format!("Contenido de {}", nombre))
}

fn main() {
    match buscar_archivo("datos.txt") {
        Ok(contenido) => println!("Éxito: {}", contenido),
        Err(e) => println!("Error: {}", e),
    }
}
```

## Patrón: Option vs Result

| Situación | Tipo | Uso |
|-----------|------|-----|
| Valor existe o no | `Option<T>` | Búsquedas, valores opcionales |
| Operación éxito/error | `Result<T, E>` | I/O, parsing, validación |

## Ejemplo Completo: Parseador

```rust
fn parsear_número(s: &str) -> Result<i32, String> {
    s.trim()
        .parse::<i32>()
        .map_err(|_| format!("'{}' no es un número válido", s))
}

fn main() {
    let números = vec!["42", "abc", "100"];
    
    for num_str in números {
        match parsear_número(num_str) {
            Ok(n) => println!("Número: {}", n),
            Err(e) => println!("Error: {}", e),
        }
    }
}
```

Output:
```
Número: 42
Error: 'abc' no es un número válido
Número: 100
```

## Resumen: Estrategia de Errores

1. **Usa `panic!()` solo para:**
   - Bugs en tu código
   - Situaciones irrecuperables
   - Tests

2. **Usa `Result` para:**
   - Errores del usuario (entrada inválida)
   - Operaciones que pueden fallar (I/O, red)
   - Código de librería

3. **Propaga errores con `?` para:**
   - Mantener el código limpio
   - Permitir que el caller maneje el error

## Siguiente Paso

Ahora sabes manejar errores. El próximo capítulo enseña **genéricos y traits**, que permite escribir código flexible y reutilizable.

---

**Consejo**: El manejo de errores es importante. No ignores los `Result`, mánejalielos adecuadamente.
