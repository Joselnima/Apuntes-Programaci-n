# 12 - Smart Pointers

Los smart pointers son referencias que tienen comportamiento adicional, como gestión automática de memoria.

## Box<T>

`Box` es el smart pointer más simple: asigna valores en el heap en lugar del stack.

### Uso Básico

```rust
fn main() {
    let b = Box::new(5);
    println!("b = {}", b);  // 5
}
```

Cuando `b` sale de scope, la memoria se libera automáticamente.

### Por Qué Usar Box

**1. Valores recursivos:**

```rust
enum Lista {
    Cons(i32, Box<Lista>),
    Vacia,
}

fn main() {
    let lista = Lista::Cons(1, Box::new(
        Lista::Cons(2, Box::new(
            Lista::Vacia
        ))
    ));
}
```

Sin `Box`, el tamaño sería infinito.

**2. Trait objects (polimorfismo):**

```rust
trait Animal {
    fn sonido(&self);
}

fn main() {
    let animales: Vec<Box<dyn Animal>> = vec![
        Box::new(Perro),
        Box::new(Gato),
    ];
}
```

## Rc<T> (Reference Counting)

`Rc` permite múltiples propietarios del mismo dato:

```rust
use std::rc::Rc;

fn main() {
    let valor = Rc::new(String::from("Hola"));
    
    let a = valor.clone();  // Incrementa contador
    let b = valor.clone();  // Incrementa contador
    
    println!("a: {}", a);
    println!("b: {}", b);
    
    // Cuando sale de scope, el contador decrece
}
```

`clone()` en `Rc` no copia los datos, solo incrementa el contador.

### Obtener el Contador

```rust
use std::rc::Rc;

fn main() {
    let valor = Rc::new(5);
    let contador_inicial = Rc::strong_count(&valor);  // 1
    
    let _a = valor.clone();
    let contador_con_uno = Rc::strong_count(&valor);  // 2
    
    println!("Conteo: {}", contador_con_uno);
}
```

## RefCell<T>

`RefCell` permite mutabilidad interior: cambiar valores sin ser mut.

```rust
use std::cell::RefCell;

fn main() {
    let valor = RefCell::new(5);
    
    // Acceso mutable sin que valor sea mut
    *valor.borrow_mut() = 10;
    
    println!("{}", *valor.borrow());  // 10
}
```

**Cuidado**: Los préstamos se verifican en runtime, no en compilación. Pueden causar panic.

### Rc + RefCell

Combinan múltiples propietarios + mutabilidad interior:

```rust
use std::rc::Rc;
use std::cell::RefCell;

struct Nodo {
    valor: i32,
    siguiente: Option<Rc<RefCell<Nodo>>>,
}

fn main() {
    let nodo = Rc::new(RefCell::new(Nodo {
        valor: 1,
        siguiente: None,
    }));
    
    // Múltiples referencias
    let copia = nodo.clone();
    
    // Modificar a través de cualquier referencia
    nodo.borrow_mut().valor = 2;
    println!("{}", copia.borrow().valor);  // 2
}
```

## Arc<T> (Atomic Reference Counting)

Como `Rc` pero thread-safe (lo veremos en concurrencia).

## Mutex<T>

Protege datos para acceso seguro en threads:

```rust
use std::sync::Mutex;

fn main() {
    let número = Mutex::new(5);
    
    {
        let mut num = número.lock().unwrap();
        *num = 6;
    }  // El lock se libera aquí
    
    println!("Número: {}", número.lock().unwrap());  // 6
}
```

## Comparación de Smart Pointers

| Pointer | Múltiples | Thread-safe | Mutabilidad |
|---------|-----------|------------|-------------|
| Box | No | No | Con &mut |
| Rc | Sí | No | Con RefCell |
| Arc | Sí | Sí | Con Mutex |

## Ejemplo Práctico: Árbol Compartido

```rust
use std::rc::Rc;
use std::cell::RefCell;

struct Nodo {
    valor: i32,
    izquierda: Option<Rc<RefCell<Nodo>>>,
    derecha: Option<Rc<RefCell<Nodo>>>,
}

impl Nodo {
    fn nuevo(valor: i32) -> Rc<RefCell<Nodo>> {
        Rc::new(RefCell::new(Nodo {
            valor,
            izquierda: None,
            derecha: None,
        }))
    }
}

fn main() {
    let raíz = Nodo::nuevo(1);
    let izq = Nodo::nuevo(2);
    
    raíz.borrow_mut().izquierda = Some(izq.clone());
    
    println!("Raíz: {}", raíz.borrow().valor);
}
```

## Deref Trait

Los smart pointers implementan `Deref`, permitiendo usar `*` para acceder:

```rust
use std::ops::Deref;

fn main() {
    let b = Box::new(5);
    println!("{}", *b);  // Deref automático
    
    let nombre = String::from("Rust");
    println!("{}", &nombre[0..2]);  // String implementa Deref a str
}
```

## Drop Trait

Controla qué pasa cuando un valor se sale de scope:

```rust
struct Persona {
    nombre: String,
}

impl Drop for Persona {
    fn drop(&mut self) {
        println!("{} se va", self.nombre);
    }
}

fn main() {
    let p = Persona { 
        nombre: String::from("Ana") 
    };
    // Al salir de scope: "Ana se va"
}
```

## Siguiente Paso

Ahora dominas gestión de memoria avanzada. El próximo capítulo enseña **concurrencia segura en Rust**.

---

**Consejo**: Los smart pointers parecen complejos, pero son herramientas específicas. Usa `Box` por defecto, recurre a otros cuando los necesites.
