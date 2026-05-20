# 10 - Genéricos, Traits y Tiempos de Vida

Estos tres conceptos son el corazón del sistema de tipos avanzado de Rust.

## Genéricos

Los genéricos permiten escribir código que funciona con muchos tipos diferentes:

### Funciones Genéricas

```rust
fn obtener_primero<T>(lista: &[T]) -> Option<&T> {
    if lista.is_empty() {
        None
    } else {
        Some(&lista[0])
    }
}

fn main() {
    let números = vec![1, 2, 3];
    let palabras = vec!["hola", "mundo"];
    
    println!("{:?}", obtener_primero(&números));  // Some(1)
    println!("{:?}", obtener_primero(&palabras)); // Some("hola")
}
```

`<T>` significa "cualquier tipo". Rust genera código para cada tipo que uses.

### Structs Genéricos

```rust
struct Caja<T> {
    contenido: T,
}

impl<T> Caja<T> {
    fn obtener_referencia(&self) -> &T {
        &self.contenido
    }
}

fn main() {
    let caja_número: Caja<i32> = Caja { contenido: 5 };
    let caja_string: Caja<String> = Caja { 
        contenido: String::from("Hola") 
    };
    
    println!("{}", caja_número.obtener_referencia());
}
```

### Múltiples Parámetros Genéricos

```rust
struct Par<T, U> {
    primero: T,
    segundo: U,
}

fn main() {
    let par = Par {
        primero: 5,
        segundo: String::from("Cinco"),
    };
}
```

## Traits: Interfaces Rust

Un trait define un conjunto de métodos que un tipo debe implementar:

### Definir un Trait

```rust
trait Animal {
    fn sonido(&self) -> String;
}

struct Perro;
struct Gato;

impl Animal for Perro {
    fn sonido(&self) -> String {
        String::from("¡Guau!")
    }
}

impl Animal for Gato {
    fn sonido(&self) -> String {
        String::from("¡Miau!")
    }
}

fn main() {
    let perro = Perro;
    let gato = Gato;
    
    println!("{}", perro.sonido());  // ¡Guau!
    println!("{}", gato.sonido());   // ¡Miau!
}
```

### Traits como Parámetros

```rust
trait Animal {
    fn sonido(&self) -> String;
}

// Función que acepta cualquier tipo que implemente Animal
fn hacer_ruido(animal: &dyn Animal) {
    println!("{}", animal.sonido());
}

struct Perro;
impl Animal for Perro {
    fn sonido(&self) -> String {
        String::from("¡Guau!")
    }
}

fn main() {
    let perro = Perro;
    hacer_ruido(&perro);
}
```

El `dyn` significa "dynamic dispatch" - el método se elige en tiempo de ejecución.

### Retornar Traits

```rust
trait Animal {
    fn sonido(&self) -> String;
}

struct Perro;
impl Animal for Perro {
    fn sonido(&self) -> String {
        String::from("¡Guau!")
    }
}

fn crear_animal() -> impl Animal {  // impl Trait
    Perro
}

fn main() {
    let animal = crear_animal();
    println!("{}", animal.sonido());
}
```

El `impl Animal` significa "devuelve cualquier tipo que implemente Animal".

### Métodos por Defecto

```rust
trait Animal {
    fn nombre(&self) -> &str;
    
    fn describir(&self) -> String {
        format!("Soy un {}", self.nombre())
    }
}

struct Perro;

impl Animal for Perro {
    fn nombre(&self) -> &str {
        "perro"
    }
}

fn main() {
    let perro = Perro;
    println!("{}", perro.describir());  // "Soy un perro"
}
```

### Traits Bounds

Limitar genéricos a tipos que implementan un trait:

```rust
trait Imprimible {
    fn mostrar(&self);
}

struct Número(i32);

impl Imprimible for Número {
    fn mostrar(&self) {
        println!("Número: {}", self.0);
    }
}

fn mostrar_algo<T: Imprimible>(cosa: &T) {
    cosa.mostrar();
}

fn main() {
    let num = Número(5);
    mostrar_algo(&num);
}
```

`<T: Imprimible>` significa "T debe implementar Imprimible".

### Múltiples Trait Bounds

```rust
fn hacer_algo<T: Clone + Display>(valor: T) {
    let copia = valor.clone();
    println!("{}", copia);
}
```

## Tiempos de Vida (Lifetimes)

Los lifetimes controlan cuánto tiempo viven las referencias:

### Problema sin Lifetimes

```rust
// ❌ No compila - ¿Cuál String se retorna?
fn obtener_string(a: &String, b: &String) -> &String {
    if a.len() > b.len() {
        a
    } else {
        b
    }
}
```

Rust no sabe si el resultado vive tanto como `a` o `b`.

### Solución: Anotar Lifetimes

```rust
// ✅ Compila - a y b tienen el mismo lifetime
fn obtener_string<'a>(a: &'a String, b: &'a String) -> &'a String {
    if a.len() > b.len() {
        a
    } else {
        b
    }
}

fn main() {
    let s1 = String::from("Hola");
    let s2 = String::from("Mundo");
    
    let resultado = obtener_string(&s1, &s2);
    println!("{}", resultado);
}
```

`'a` es un lifetime. Significa "el resultado vive tanto como los parámetros".

### Lifetimes en Structs

```rust
struct Citación<'a> {
    texto: &'a str,
    autor: &'a str,
}

fn main() {
    let cita_texto = "La vida es lo que pasa mientras haces planes";
    let cita_autor = "John Lennon";
    
    let cita = Citación {
        texto: cita_texto,
        autor: cita_autor,
    };
    
    println!("{} - {}", cita.texto, cita.autor);
}
```

### Reglas de Elision de Lifetimes

Rust infiere lifetimes en casos simples:

```rust
// Sin anotación explícita
fn obtener_primero(s: &String) -> &String {
    s
}

// Rust asume implícitamente:
fn obtener_primero<'a>(s: &'a String) -> &'a String {
    s
}
```

Si hay una referencia de entrada y una de salida, Rust asume que tienen el mismo lifetime.

## Combinando Todo: Ejemplo Completo

```rust
trait Descritor {
    fn descripción(&self) -> String;
}

struct Libro<'a> {
    título: &'a str,
    autor: &'a str,
}

impl<'a> Descritor for Libro<'a> {
    fn descripción(&self) -> String {
        format!("{} de {}", self.título, self.autor)
    }
}

fn mostrar_descripción<'a, T: Descritor>(cosa: &'a T) {
    println!("Descripción: {}", cosa.descripción());
}

fn main() {
    let libro = Libro {
        título: "1984",
        autor: "George Orwell",
    };
    
    mostrar_descripción(&libro);
}
```

## Resumen

| Concepto | Uso |
|----------|-----|
| Genéricos | Escribir para múltiples tipos |
| Traits | Definir interfaces/contratos |
| Trait Bounds | Limitar genéricos a tipos específicos |
| Lifetimes | Controlar validez de referencias |

## Siguiente Paso

Estos conceptos son avanzados pero fundamentales. Practica con ejemplos para dominarlos. El próximo capítulo enseña **closures e iteradores**, que son herramientas funcionales poderosas.

---

**Consejo**: Los lifetimes parecen complejos, pero con práctica se vuelven intuitivos. Recuerda: Rust está siendo paranoico por tu seguridad.
