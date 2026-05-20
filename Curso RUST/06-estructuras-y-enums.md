# 06 - Estructuras (Structs) y Enumeraciones (Enums)

## Estructuras (Structs)

Las estructuras permiten agrupar múltiples valores relacionados en un solo tipo personalizado.

### Definir y Usar Structs

```rust
// Definir una estructura
struct Usuario {
    nombre: String,
    edad: u32,
    correo: String,
}

fn main() {
    // Crear una instancia
    let usuario = Usuario {
        nombre: String::from("Ana"),
        edad: 30,
        correo: String::from("ana@example.com"),
    };
    
    // Acceder a los campos
    println!("Nombre: {}", usuario.nombre);
    println!("Edad: {}", usuario.edad);
}
```

### Variables Mutables en Structs

Para cambiar campos, la struct debe ser mutable:

```rust
struct Usuario {
    nombre: String,
    edad: u32,
}

fn main() {
    let mut usuario = Usuario {
        nombre: String::from("Carlos"),
        edad: 25,
    };
    
    usuario.edad = 26;  // Cambiar edad
    println!("Nueva edad: {}", usuario.edad);
}
```

### Función que Retorna Struct

```rust
fn crear_usuario(nombre: String, edad: u32) -> Usuario {
    Usuario {
        nombre,  // Abreviatura: nombre: nombre,
        edad,    // Abreviatura: edad: edad,
    }
}

fn main() {
    let usuario = crear_usuario(String::from("Diana"), 28);
    println!("Usuario: {} años {}", usuario.nombre, usuario.edad);
}
```

### Actualización de Structs

```rust
struct Usuario {
    nombre: String,
    edad: u32,
    correo: String,
}

fn main() {
    let usuario1 = Usuario {
        nombre: String::from("Ana"),
        edad: 30,
        correo: String::from("ana@example.com"),
    };
    
    // Crear usuario2 a partir de usuario1, cambiando solo edad
    let usuario2 = Usuario {
        edad: 31,
        ..usuario1  // El resto de campos vienen de usuario1
    };
    
    println!("Usuario2: {} años", usuario2.edad);
}
```

## Tuple Structs

Structs sin nombres para los campos:

```rust
struct Color(i32, i32, i32);
struct Punto(f64, f64);

fn main() {
    let negro = Color(0, 0, 0);
    let origen = Punto(0.0, 0.0);
    
    println!("Rojo: {}", negro.0);  // Acceder por índice
    println!("X: {}", origen.0);
}
```

## Unit Structs

Structs sin campos (útiles para marcadores):

```rust
struct Marcador;

fn main() {
    let _m = Marcador;
}
```

## Métodos

Los métodos son funciones asociadas a un tipo. Se definen con `impl`:

```rust
struct Rectangulo {
    ancho: u32,
    alto: u32,
}

impl Rectangulo {
    // Método que toma &self (no mutable)
    fn area(&self) -> u32 {
        self.ancho * self.alto
    }
    
    // Método que toma &mut self (mutable)
    fn cambiar_ancho(&mut self, nuevo_ancho: u32) {
        self.ancho = nuevo_ancho;
    }
    
    // Método que toma self (consume el objeto)
    fn mostrar_info(self) {
        println!("Rect: {}x{}", self.ancho, self.alto);
    }
}

fn main() {
    let mut rect = Rectangulo { ancho: 10, alto: 5 };
    
    println!("Área: {}", rect.area());  // 50
    
    rect.cambiar_ancho(20);
    println!("Nueva área: {}", rect.area());  // 100
    
    rect.mostrar_info();  // rect se destruye aquí
}
```

### Funciones Asociadas

Las funciones sin `self` se llaman "funciones asociadas" o constructores:

```rust
impl Rectangulo {
    fn cuadrado(lado: u32) -> Rectangulo {
        Rectangulo {
            ancho: lado,
            alto: lado,
        }
    }
}

fn main() {
    let cuadrado = Rectangulo::cuadrado(5);
    println!("Área: {}", cuadrado.area());  // 25
}
```

Se llaman con `::` (Type::function).

## Derivar Traits Automáticamente

```rust
#[derive(Debug)]
struct Punto {
    x: i32,
    y: i32,
}

fn main() {
    let p = Punto { x: 5, y: 10 };
    println!("{:?}", p);  // Punto { x: 5, y: 10 }
    println!("{:#?}", p); // Formato bonito con saltos de línea
}
```

## Enumeraciones (Enums)

Las enums permiten definir tipos con un conjunto de variantes posibles:

### Enum Básico

```rust
enum Dirección {
    Arriba,
    Abajo,
    Izquierda,
    Derecha,
}

fn mover_personaje(dir: Dirección) {
    match dir {
        Dirección::Arriba => println!("Moviendo arriba"),
        Dirección::Abajo => println!("Moviendo abajo"),
        Dirección::Izquierda => println!("Moviendo izquierda"),
        Dirección::Derecha => println!("Moviendo derecha"),
    }
}

fn main() {
    let dir = Dirección::Arriba;
    mover_personaje(dir);
}
```

### Enum con Datos

Las variantes pueden contener datos:

```rust
enum Resultado {
    Exito(String),
    Error(String),
}

fn procesar() -> Resultado {
    Resultado::Exito(String::from("Operación completada"))
}

fn main() {
    match procesar() {
        Resultado::Exito(msg) => println!("Éxito: {}", msg),
        Resultado::Error(msg) => println!("Error: {}", msg),
    }
}
```

### Option: El Enum de "Posiblemente Nulo"

Rust no tiene `null`. En su lugar, usa `Option`:

```rust
enum Option<T> {
    Some(T),
    None,
}

fn buscar_número(números: &[i32], objetivo: i32) -> Option<usize> {
    for (i, &num) in números.iter().enumerate() {
        if num == objetivo {
            return Some(i);
        }
    }
    None
}

fn main() {
    let números = [1, 2, 3, 4, 5];
    
    match buscar_número(&números, 3) {
        Some(índice) => println!("Encontrado en índice {}", índice),
        None => println!("No encontrado"),
    }
}
```

### Result: El Enum de "Éxito o Error"

Para operaciones que pueden fallar:

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}

fn dividir(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err(String::from("No se puede dividir por cero"))
    } else {
        Ok(a / b)
    }
}

fn main() {
    match dividir(10.0, 2.0) {
        Ok(resultado) => println!("Resultado: {}", resultado),
        Err(error) => println!("Error: {}", error),
    }
}
```

## Pattern Matching (if let)

A veces solo te importa una variante:

```rust
fn main() {
    let valor = Some(3);
    
    // Forma completa
    match valor {
        Some(3) => println!("Es tres"),
        _ => println!("Es otra cosa"),
    }
    
    // Forma abreviada con if let
    if let Some(3) = valor {
        println!("Es tres");
    }
}
```

### Option con if let

```rust
fn main() {
    let números = [1, 2, 3];
    let búsqueda = algunos_números(&números, 2);
    
    if let Some(índice) = búsqueda {
        println!("Encontrado en {}", índice);
    }
}
```

### while let

```rust
fn main() {
    let mut valores = vec![1, 2, 3];
    
    while let Some(valor) = valores.pop() {
        println!("Valor: {}", valor);  // Imprime 3, 2, 1
    }
}
```

## Métodos en Enums

```rust
enum Mensaje {
    Escribir(String),
    Cambiar { r: i32, g: i32, b: i32 },
    Mover { x: i32, y: i32 },
    Salir,
}

impl Mensaje {
    fn procesar(&self) {
        match self {
            Mensaje::Escribir(texto) => println!("Texto: {}", texto),
            Mensaje::Cambiar { r, g, b } => println!("Color: {},{},{}", r, g, b),
            Mensaje::Mover { x, y } => println!("Posición: {},{}",x, y),
            Mensaje::Salir => println!("Adiós"),
        }
    }
}

fn main() {
    let msg = Mensaje::Cambiar { r: 255, g: 0, b: 0 };
    msg.procesar();
}
```

## Resumen

| Concepto | Uso |
|----------|-----|
| Struct | Agrupar datos relacionados |
| Método | Función en struct con `&self` |
| Función Asociada | Función en struct sin `self` (constructores) |
| Enum | Tipo con variantes posibles |
| Option | Para "posiblemente valor" |
| Result | Para "éxito o error" |

## Siguiente Paso

Ahora tienes structs y enums para organizar datos. El próximo capítulo profundiza en **pattern matching**, que es poderoso con estos tipos.

---

**Consejo**: Las structs y enums son la base de la mayoría de programas Rust. Domínalas bien.
