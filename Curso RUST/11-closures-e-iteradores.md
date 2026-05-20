# 11 - Closures e Iteradores

Dos características funcionales que hacen a Rust muy expresivo.

## Closures (Cierres)

Un closure es una función anónima que puede capturar variables de su entorno.

### Closures Básicos

```rust
fn main() {
    let x = 5;
    
    // Closure que captura x
    let agregar_x = |número| número + x;
    
    println!("{}", agregar_x(3));  // 8
}
```

El closure se define entre `| |`. Puede acceder a variables externas.

### Tipos de Captura

Rust infiere automáticamente cómo capturar:

```rust
fn main() {
    let s = String::from("Hola");
    
    // Captura referencia (lectura)
    let usa = || println!("{}", s);
    usa();
    
    println!("Puedo usar s: {}", s);  // Aún válido
}
```

Para captura mutable:

```rust
fn main() {
    let mut x = 5;
    
    let mut incrementar = || {
        x += 1;
    };
    
    incrementar();
    println!("{}", x);  // 6
}
```

Para tomar propiedad:

```rust
fn main() {
    let s = String::from("Hola");
    
    let tomar = move || {
        println!("{}", s);
    };
    
    tomar();
    // println!("{}", s);  // ❌ Error: s se movió
}
```

### Closures como Parámetros

```rust
fn aplicar_closure<F>(x: i32, f: F) -> i32
where
    F: Fn(i32) -> i32,
{
    f(x)
}

fn main() {
    let resultado = aplicar_closure(5, |x| x * 2);
    println!("{}", resultado);  // 10
}
```

Traits para closures:
- `Fn`: Captura referencia (puede usar múltiples veces)
- `FnMut`: Captura referencia mutable
- `FnOnce`: Toma propiedad (una sola vez)

### Retornar Closures

```rust
fn hacer_multiplicador(n: i32) -> Box<dyn Fn(i32) -> i32> {
    Box::new(move |x| x * n)
}

fn main() {
    let multiplicar_por_2 = hacer_multiplicador(2);
    println!("{}", multiplicar_por_2(5));  // 10
}
```

## Iteradores

Los iteradores permiten procesar elementos secuencialmente.

### Crear Iteradores

```rust
fn main() {
    let números = vec![1, 2, 3];
    
    // iter() devuelve referencias
    for num in números.iter() {
        println!("{}", num);
    }
    
    // into_iter() toma propiedad
    for num in números {
        println!("{}", num);
    }
    
    // iter_mut() da referencias mutables
    let mut números = vec![1, 2, 3];
    for num in &mut números {
        *num *= 2;
    }
}
```

### Métodos de Iterador

#### map

Transforma cada elemento:

```rust
fn main() {
    let números = vec![1, 2, 3];
    
    let duplicados: Vec<i32> = números
        .iter()
        .map(|x| x * 2)
        .collect();
    
    println!("{:?}", duplicados);  // [2, 4, 6]
}
```

#### filter

Selecciona elementos:

```rust
fn main() {
    let números = vec![1, 2, 3, 4, 5];
    
    let pares: Vec<i32> = números
        .iter()
        .filter(|x| x % 2 == 0)
        .copied()
        .collect();
    
    println!("{:?}", pares);  // [2, 4]
}
```

#### fold

Acumula valores:

```rust
fn main() {
    let números = vec![1, 2, 3, 4];
    
    let suma = números
        .iter()
        .fold(0, |acum, x| acum + x);
    
    println!("{}", suma);  // 10
}
```

#### find

Encuentra el primero que cumple condición:

```rust
fn main() {
    let números = vec![1, 2, 3, 4, 5];
    
    if let Some(par) = números.iter().find(|x| *x % 2 == 0) {
        println!("Primer par: {}", par);  // 2
    }
}
```

#### any / all

Comprueba condiciones:

```rust
fn main() {
    let números = vec![1, 2, 3, 4, 5];
    
    let hay_par = números.iter().any(|x| x % 2 == 0);  // true
    let todos_positivos = números.iter().all(|x| x > &0);  // true
}
```

### Cadenas de Iteradores

Combina múltiples operaciones:

```rust
fn main() {
    let números = vec![1, 2, 3, 4, 5];
    
    let resultado: Vec<i32> = números
        .iter()
        .filter(|x| x % 2 == 0)  // Filtra pares
        .map(|x| x * 2)          // Duplica
        .collect();
    
    println!("{:?}", resultado);  // [4, 8]
}
```

### Crear Iteradores Personalizados

```rust
struct Contador {
    desde: i32,
    hasta: i32,
}

impl Iterator for Contador {
    type Item = i32;
    
    fn next(&mut self) -> Option<i32> {
        if self.desde <= self.hasta {
            let res = self.desde;
            self.desde += 1;
            Some(res)
        } else {
            None
        }
    }
}

fn main() {
    let contador = Contador { desde: 1, hasta: 3 };
    for num in contador {
        println!("{}", num);  // 1, 2, 3
    }
}
```

## Ejemplo Completo: Procesar Datos

```rust
struct Persona {
    nombre: String,
    edad: u32,
}

fn main() {
    let personas = vec![
        Persona { nombre: "Ana".to_string(), edad: 28 },
        Persona { nombre: "Carlos".to_string(), edad: 35 },
        Persona { nombre: "Diana".to_string(), edad: 22 },
    ];
    
    // Filtra mayores de 25, obtiene nombres, los ordena
    let adultos: Vec<_> = personas
        .iter()
        .filter(|p| p.edad > 25)
        .map(|p| p.nombre.clone())
        .collect();
    
    println!("Adultos: {:?}", adultos);  // ["Ana", "Carlos"]
}
```

## Performance: Iteradores vs Loops

Los iteradores se optimizan a nivel de compilación:

```rust
// Ambos generan el mismo código máquina
let suma1: i32 = (1..=5).sum();

let mut suma2 = 0;
for i in 1..=5 {
    suma2 += i;
}
```

**Recomendación**: Usa iteradores cuando sea posible - son idiomáticos en Rust y eficientes.

## Resumen

| Método | Qué hace |
|--------|----------|
| `map` | Transforma cada elemento |
| `filter` | Filtra elementos |
| `fold` | Acumula valores |
| `find` | Encuentra elemento |
| `any` | ¿Hay alguno que cumpla? |
| `all` | ¿Todos cumplen? |
| `collect` | Recolecta en colección |

## Siguiente Paso

Ahora sabes programar con estilo funcional. El próximo capítulo enseña **smart pointers**, que son referencias inteligentes para casos especiales.

---

**Consejo**: Los iteradores son idiomáticos en Rust. Aprende a usar `.map()` y `.filter()` bien, transformarán tu código.
