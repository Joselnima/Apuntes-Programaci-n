# 09 - Colecciones: Vec, String, Slice, HashMap, HashSet

Hasta ahora has trabajado con tipos pequeños. Ahora aprenderás a almacenar múltiples valores de formas diferentes.

## Vectores (Vec<T>)

Un vector es una colección dinámica que puede crecer o encogerse.

### Crear Vectores

```rust
fn main() {
    // Forma 1: Con macro vec!
    let números = vec![1, 2, 3];
    
    // Forma 2: Vacío con tipo explícito
    let mut lista: Vec<i32> = Vec::new();
    
    // Forma 3: Con capacidad inicial (optimiza memory)
    let mut con_capacidad = Vec::with_capacity(10);
    
    // Forma 4: Repetir valor
    let repetido = vec![0; 5];  // [0, 0, 0, 0, 0]
    
    println!("{:?}", números);  // [1, 2, 3]
}
```

### Insertar y Eliminar

#### `push(valor)` - Agregar al Final
- **Qué hace**: Añade un elemento al final del vector
- **Cuándo usarlo**: Cuando construyes un vector y agregas elementos secuencialmente
- **Complejidad**: O(1) amortizado

```rust
let mut v = vec![1, 2, 3];
v.push(4);
println!("{:?}", v);  // [1, 2, 3, 4]
```

#### `pop()` - Eliminar el Último
- **Qué hace**: Extrae y devuelve el último elemento (como una pila/stack)
- **Devuelve**: `Option<T>` - `Some(valor)` si hay elemento, `None` si está vacío
- **Cuándo usarlo**: Procesamiento de pilas, deshacer acciones (undo), últimas operaciones

```rust
let mut v = vec![1, 2, 3];
if let Some(último) = v.pop() {
    println!("Extraído: {}", último);  // 3
    println!("Vector ahora: {:?}", v);  // [1, 2]
}
```

#### `insert(índice, valor)` - Insertar en Posición
- **Qué hace**: Inserta un elemento en un índice específico, desplazando los siguientes
- **Complejidad**: O(n) - lento para vectores grandes en índices bajos
- **Cuándo usarlo**: Cuando necesitas insertar en medio, pero hazlo con cuidado (es costoso)

```rust
let mut v = vec![1, 2, 3];
v.insert(1, 10);  // Inserta 10 en posición 1
println!("{:?}", v);  // [1, 10, 2, 3]
```

#### `remove(índice)` - Eliminar en Posición
- **Qué hace**: Elimina y devuelve el elemento en un índice específico
- **Complejidad**: O(n) - lento para vectores grandes
- **Devuelve**: El valor eliminado
- **Cuándo usarlo**: Cuando sabes exactamente qué elemento eliminar por posición

```rust
let mut v = vec![1, 2, 3];
let eliminado = v.remove(1);
println!("Eliminado: {}", eliminado);  // 2
println!("Vector: {:?}", v);  // [1, 3]
```

#### Resumen de Operaciones Stack:
```rust
let mut stack = Vec::new();
stack.push(1);
stack.push(2);
stack.push(3);
assert_eq!(stack.pop(), Some(3));  // LIFO: Last In, First Out
```

### Acceder a Elementos

```rust
fn main() {
    let números = vec![1, 2, 3, 4, 5];
    
    // Indexación (panics si está fuera)
    println!("{}", números[0]);  // 1
    
    // get() - devuelve Option (más seguro)
    match números.get(2) {
        Some(valor) => println!("Índice 2: {}", valor),
        None => println!("Índice no existe"),
    }
    
    // Primero y último
    if let Some(primero) = números.first() {
        println!("Primero: {}", primero);
    }
    
    if let Some(último) = números.last() {
        println!("Último: {}", último);
    }
}
```

### Búsqueda y Ordenamiento

#### `contains(&elemento)` - Verificar Existencia
- **Qué hace**: Devuelve `true` si el elemento existe en el vector
- **Complejidad**: O(n) - busca linealmente
- **Cuándo usarlo**: Cuando necesitas verificar si algo existe (búsqueda simple)

```rust
let números = vec![1, 2, 3, 4, 5];
if números.contains(&3) {
    println!("El vector contiene 3");
}
```

#### `position()` - Encontrar Índice
- **Qué hace**: Devuelve el índice del primer elemento que cumple una condición
- **Devuelve**: `Option<usize>`
- **Cuándo usarlo**: Cuando necesitas saber dónde está un elemento

```rust
let números = vec![10, 20, 30, 40];
if let Some(pos) = números.iter().position(|&x| x == 30) {
    println!("30 está en índice {}", pos);  // índice 2
}
```

#### `sort()` - Ordenar Ascendente
- **Qué hace**: Ordena el vector in-place de menor a mayor
- **Complejidad**: O(n log n) - promedio
- **Modifica**: El vector original (no crea uno nuevo)
- **Cuándo usarlo**: Cuando necesitas datos ordenados de forma predeterminada

```rust
let mut números = vec![5, 2, 8, 1, 9];
números.sort();
println!("{:?}", números);  // [1, 2, 5, 8, 9]
```

#### `sort_by()` - Ordenar Personalizado
- **Qué hace**: Ordena usando una función comparadora personalizada
- **Cuándo usarlo**: Cuando necesitas lógica de ordenamiento especial

```rust
let mut números = vec![5, 2, 8, 1, 9];

// Descendente
números.sort_by(|a, b| b.cmp(a));
println!("{:?}", números);  // [9, 8, 5, 2, 1]

// Por valor absoluto
let mut con_negativos = vec![-5, 3, -1, 8];
con_negativos.sort_by_key(|x| x.abs());
println!("{:?}", con_negativos);  // [-1, 3, -5, 8]
```

#### `reverse()` - Invertir Orden
- **Qué hace**: Invierte el orden de todos los elementos
- **Complejidad**: O(n)
- **Cuándo usarlo**: Para procesar en orden inverso o deshacer un orden

```rust
let mut números = vec![1, 2, 3, 4];
números.reverse();
println!("{:?}", números);  // [4, 3, 2, 1]
```

#### `binary_search(&elemento)` - Búsqueda Rápida
- **Qué hace**: Busca eficientemente en un vector **ya ordenado**
- **Complejidad**: O(log n) - mucho más rápido que contains()
- **Devuelve**: `Result<índice, índice_donde_debería_ir>`
- **Cuándo usarlo**: Cuando el vector está ordenado y necesitas búsqueda frecuente
- **IMPORTANTE**: ¡Requiere vector ordenado!

```rust
let números = vec![1, 3, 5, 7, 9];  // Debe estar ordenado

match números.binary_search(&5) {
    Ok(pos) => println!("Encontrado en índice {}", pos),
    Err(pos) => println!("No encontrado, iría en índice {}", pos),
}
```

### Deduplicación y Filtrado

#### `dedup()` - Eliminar Duplicados Consecutivos
- **Qué hace**: Elimina elementos duplicados consecutivos (requiere estar ordenado primero)
- **Complejidad**: O(n)
- **Cuándo usarlo**: Después de ordenar, cuando solo quieres valores únicos consecutivos
- **IMPORTANTE**: Solo funciona en duplicados **consecutivos**

```rust
let mut números = vec![1, 1, 2, 2, 3, 3, 3, 4];
números.dedup();
println!("{:?}", números);  // [1, 2, 3, 4]

// Esto NO funciona correctamente si no está ordenado primero:
let mut desordenado = vec![1, 2, 1, 2, 1];
desordenado.dedup();
println!("{:?}", desordenado);  // [1, 2, 1, 2, 1] - no cambia!
```

#### `retain(predicado)` - Mantener Solo Elementos que Cumplen Condición
- **Qué hace**: Mantiene solo elementos que cumplen la condición del closure, elimina el resto
- **Complejidad**: O(n)
- **Modifica**: El vector original
- **Cuándo usarlo**: Filtrado in-place, cuando quieres descartar elementos
- **Mejor que**: Crear un nuevo vector con filter() cuando la memoria es crítica

```rust
let mut números = vec![1, 2, 3, 4, 5, 6, 7];

// Mantener solo pares
números.retain(|&x| x % 2 == 0);
println!("{:?}", números);  // [2, 4, 6]

// Mantener solo mayores a 5
let mut valores = vec![1, 5, 3, 8, 2, 9];
valores.retain(|&x| x > 5);
println!("{:?}", valores);  // [8, 9]
```

#### `drain(rango)` - Extraer Rango y Limpiar
- **Qué hace**: Elimina un rango de elementos y devuelve un iterador sobre ellos
- **Cuándo usarlo**: Cuando quieres procesar y eliminar un rango al mismo tiempo
- **Eficiente**: No crea copia, solo traslada

```rust
let mut v = vec![1, 2, 3, 4, 5];

// Extraer elementos del índice 1 al 3
for valor in v.drain(1..3) {
    println!("Extraído: {}", valor);  // Imprime 2, 3
}

println!("Vector después: {:?}", v);  // [1, 4, 5]
```

### Iterar sobre Vectores

```rust
fn main() {
    let números = vec![1, 2, 3, 4, 5];
    
    // Lectura (referencias)
    for num in &números {
        println!("{}", num);
    }
    
    // Con índice
    for (i, num) in números.iter().enumerate() {
        println!("Posición {}: {}", i, num);
    }
    
    // Modificar (referencias mutables)
    let mut números = números;
    for num in &mut números {
        *num *= 2;
    }
    println!("{:?}", números);  // [2, 4, 6, 8, 10]
    
    // Tomar propiedad
    let números = vec!["a", "b", "c"];
    for palabra in números {  // Se consume el vector
        println!("{}", palabra);
    }
}
```

### Métodos Informacionales

```rust
fn main() {
    let v = vec![1, 2, 3, 4, 5];
    
    println!("Longitud: {}", v.len());              // 5
    println!("¿Vacío? {}", v.is_empty());           // false
    println!("Capacidad: {}", v.capacity());        // >= 5
    
    let v2: Vec<i32> = Vec::new();
    println!("Vector vacío: {}", v2.is_empty());    // true
}
```

### Clonación y Conversión

```rust
fn main() {
    let v1 = vec![1, 2, 3];
    
    // Clonar
    let v2 = v1.clone();
    println!("Copia: {:?}", v2);
    
    // De array a vector
    let array = [1, 2, 3];
    let v3 = array.to_vec();
    
    // De slice a vector
    let slice: &[i32] = &[1, 2, 3];
    let v4 = slice.to_vec();
}
```

## Strings y &str

Los strings en Rust son complejos pero poderosos.

### Tipos de String

```rust
fn main() {
    // &str - String slice, inmutable, referencia fija
    let s1: &str = "Hola";  // Literal
    
    // String - String dinámico, mutable, en el heap
    let s2: String = String::from("Hola");
    
    // Convertir
    let s3 = s1.to_string();
    let s4: String = "Hola".to_owned();
    
    // String vacío
    let mut s5 = String::new();
    s5.push_str("Contenido");
}
```

### Construcción de Strings

```rust
fn main() {
    // String::from
    let s1 = String::from("Hola");
    
    // .to_string()
    let s2 = "Mundo".to_string();
    
    // format!
    let nombre = "Ana";
    let s3 = format!("Hola, {}", nombre);
    
    // Repetir
    let s4 = "ab".repeat(3);  // "ababab"
}
```

### Concatenación

```rust
fn main() {
    // push_str() - agrega &str
    let mut s = String::from("Hola");
    s.push_str(" Mundo");
    println!("{}", s);  // "Hola Mundo"
    
    // push() - agrega char
    s.push('!');
    println!("{}", s);  // "Hola Mundo!"
    
    // + operator
    let s1 = String::from("Hola");
    let s2 = String::from("Mundo");
    let s3 = s1 + " " + &s2;  // s1 se consume
    
    // format! - más flexible
    let resultado = format!("{} {}", "Hola", "Mundo");
}
```

### Insertar y Eliminar

```rust
fn main() {
    // insert() - insertar char en posición
    let mut s = String::from("Hola");
    s.insert(1, 'X');
    println!("{}", s);  // "HXola"
    
    // insert_str() - insertar string
    s.insert_str(1, "YY");
    println!("{}", s);  // "HYYXola"
    
    // remove() - eliminar char en posición
    let c = s.remove(1);
    println!("{}", s);  // "HYXola"
    
    // pop() - eliminar último char
    if let Some(c) = s.pop() {
        println!("Eliminado: {}", c);
    }
}
```

### Búsqueda y Slicing

#### `contains(&substring)` - Verificar Contenido
- **Qué hace**: Devuelve `true` si el string contiene el substring
- **Complejidad**: O(n*m) - donde n es longitud del string, m es la del substring
- **Cuándo usarlo**: Búsquedas simples de presencia

```rust
let s = "¡Hola, Mundo!";
if s.contains("Mundo") {
    println!("Contiene 'Mundo'");  // Se imprime
}
```

#### `starts_with()` y `ends_with()` - Verificar Inicio/Final
- **Qué hace**: Comprueba si el string empieza o termina con el patrón
- **Cuándo usarlo**: Validación de formatos, extensiones de archivos, etc.

```rust
let s = "archivo.txt";
println!("{}", s.starts_with("archivo"));  // true
println!("{}", s.ends_with(".txt"));       // true

// Validar protocolo
let url = "https://example.com";
if url.starts_with("https://") {
    println!("Conexión segura");
}
```

#### `find()` y `rfind()` - Encontrar Posición
- **Qué hace**: Busca un substring y devuelve su posición (índice)
- **Devuelve**: `Option<usize>`
- **find()**: Busca desde el inicio
- **rfind()**: Busca desde el final (últimaocurrencia)
- **Cuándo usarlo**: Cuando necesitas la posición exacta para extraer partes

```rust
let s = "contacto@correo.com";

// Encontrar @
if let Some(pos) = s.find('@') {
    let usuario = &s[..pos];
    let dominio = &s[pos+1..];
    println!("Usuario: {}", usuario);  // "contacto"
    println!("Dominio: {}", dominio);  // "correo.com"
}

// Última coma
let s2 = "a,b,c,d";
if let Some(pos) = s2.rfind(',') {
    println!("Última coma en posición {}", pos);  // 5
}
```

#### Slicing de String
- **Qué hace**: Extrae una porción del string usando `&s[inicio..fin]`
- **Cuándo usarlo**: Cuando necesitas una parte específica
- **CUIDADO**: Rust funciona por bytes, no por caracteres. Con emojis puede fallar

```rust
let s = "Hola";
let slice = &s[0..2];  // "Ho"
let slice2 = &s[..];    // Todo el string

// Esto puede fallar con caracteres especiales:
// let s2 = "Hola 😊";
// let parte = &s2[0..5];  // ❌ Panic! El emoji ocupa 4 bytes
```

### Transformación de Caso

#### `to_uppercase()` y `to_lowercase()` - Cambiar Mayúsculas
- **Qué hace**: Convierte todo a mayúsculas o minúsculas
- **Devuelve**: Nuevo string (no modifica el original)
- **Cuándo usarlo**: Comparaciones case-insensitive, formateo, normalizando entrada

```rust
let s = "Hola Mundo";

println!("{}", s.to_uppercase());    // "HOLA MUNDO"
println!("{}", s.to_lowercase());    // "hola mundo"

// Comparación case-insensitive
let entrada = "RUST";
if entrada.to_lowercase() == "rust" {
    println!("Usuario escribió 'rust'");
}
```

### Dividir Strings

#### `split(separador)` - Dividir por Delimitador
- **Qué hace**: Divide el string en partes usando un delimitador
- **Devuelve**: Iterador de substrings
- **Cuándo usarlo**: Parsing de CSV, líneas de comando, configuración

```rust
let datos = "Nombre,Edad,Ciudad";
for campo in datos.split(',') {
    println!("Campo: {}", campo);
}
// Output: Nombre, Edad, Ciudad

// Con trim para quitar espacios
let lista = "uno, dos, tres";
for item in lista.split(',') {
    println!("- {}", item.trim());
}
```

#### `split_whitespace()` - Dividir por Espacios
- **Qué hace**: Divide por espacios en blanco (espacios, tabulaciones, newlines)
- **Cuándo usarlo**: Procesar líneas de comandos, palabras

```rust
let frase = "Rust es fantástico";
for palabra in frase.split_whitespace() {
    println!("Palabra: {}", palabra);
}
// Output: Rust, es, fantástico
```

#### `lines()` - Dividir por Saltos de Línea
- **Qué hace**: Divide el string por `\n`, quitándolos automáticamente
- **Cuándo usarlo**: Procesar archivos de texto, logs línea por línea

```rust
let texto = "línea 1\nlínea 2\nlínea 3";
for (num, línea) in texto.lines().enumerate() {
    println!("Línea {}: {}", num + 1, línea);
}
```

#### `split_at()` - Dividir en Posición Específica
- **Qué hace**: Divide el string en dos partes en una posición exacta
- **Devuelve**: Tupla `(&str, &str)`
- **Cuándo usarlo**: Cuando sabes exactamente dónde dividir

```rust
let s = "Hola";
let (parte1, parte2) = s.split_at(2);
println!("'{}' | '{}'", parte1, parte2);  // 'Ho' | 'la'
```

### Limpieza de Strings

#### `trim()` - Eliminar Espacios
- **Qué hace**: Quita espacios en blanco (espacios, tabs, newlines) del inicio y final
- **Cuándo usarlo**: Limpiar entrada de usuario, archivos, APIs

```rust
let s = "  hola mundo  ";
println!("'{}'", s.trim());        // 'hola mundo'
println!("'{}'", s.trim_start());  // 'hola mundo  '
println!("'{}'", s.trim_end());    // '  hola mundo'
```

#### `trim_matches()` - Eliminar Caracteres Específicos
- **Qué hace**: Quita caracteres específicos del inicio y final
- **Cuándo usarlo**: Remover comillas, caracteres especiales, etc.

```rust
let s = "xxxHolaxxx";
println!("{}", s.trim_matches('x'));  // "Hola"

let s2 = "\"Hola\"";
println!("{}", s2.trim_matches('"'));  // "Hola"

// Con predicado
let s3 = "123abc456";
println!("{}", s3.trim_matches(|c: char| c.is_numeric()));  // "abc"
```

### Reemplazo

#### `replace(de, por)` - Reemplazar Todos
- **Qué hace**: Reemplaza **todas** las ocurrencias
- **Devuelve**: Nuevo string (no modifica original)
- **Complejidad**: O(n*m)
- **Cuándo usarlo**: Sustituciones simples

```rust
let s = "gato gato gato";
println!("{}", s.replace("gato", "perro"));  // "perro perro perro"

// Templating sencillo
let saludo = "Hola NOMBRE, bienvenido a CIUDAD";
let mensaje = saludo
    .replace("NOMBRE", "Ana")
    .replace("CIUDAD", "Madrid");
println!("{}", mensaje);
```

#### `replacen(de, por, contador)` - Reemplazar Primeros N
- **Qué hace**: Reemplaza solo las primeras n ocurrencias
- **Cuándo usarlo**: Cuando necesitas controlar cuántos reemplazos

```rust
let s = "el el el";
println!("{}", s.replacen("el", "un", 1));  // "un el el"
println!("{}", s.replacen("el", "un", 2));  // "un un el"
```

### Iteración sobre Caracteres

```rust
fn main() {
    let s = "Hola";
    
    // chars() - iterador de caracteres
    for c in s.chars() {
        println!("Carácter: {}", c);
    }
    
    // bytes() - iterador de bytes
    for b in s.bytes() {
        println!("Byte: {}", b);
    }
    
    // Acceso individual
    let c = s.chars().nth(1);  // Option
    println!("Segundo carácter: {:?}", c);
}
```

### Información de String

```rust
fn main() {
    let s = "¡Hola!";
    
    println!("Bytes: {}", s.len());           // 7 (¡ toma 2 bytes)
    println!("Caracteres: {}", s.chars().count());  // 6
    println!("¿Vacío? {}", s.is_empty());    // false
    println!("¿ASCII? {}", s.is_ascii());     // false
}
```

## Slices (&[T])

Un slice es una referencia a una porción de una colección.

### Crear Slices

```rust
fn main() {
    let números = vec![1, 2, 3, 4, 5];
    
    // Diferentes formas de slicing
    let todo: &[i32] = &números[..];      // Todos
    let inicio: &[i32] = &números[..3];   // Primeros 3
    let final_: &[i32] = &números[2..];   // Desde índice 2
    let rango: &[i32] = &números[1..4];   // Índices 1, 2, 3
    
    println!("Todo: {:?}", todo);
    println!("Primeros 3: {:?}", inicio);
    println!("Desde 2: {:?}", final_);
    println!("Rango 1-4: {:?}", rango);
}
```

### Métodos Útiles de Slices

```rust
fn main() {
    let números = [1, 2, 3, 4, 5];
    let slice = &números[1..4];
    
    // Información
    println!("Longitud: {}", slice.len());
    println!("¿Vacío? {}", slice.is_empty());
    
    // Acceso
    println!("Primero: {:?}", slice.first());
    println!("Último: {:?}", slice.last());
    println!("En índice 1: {:?}", slice.get(1));
    
    // Búsqueda
    println!("¿Contiene 3? {}", slice.contains(&3));
    
    // Ordenamiento (si es mutable)
    let mut slice_mut = [3, 1, 2];
    slice_mut.sort();
    println!("Ordenado: {:?}", slice_mut);
}
```

### Iteración sobre Slices

```rust
fn main() {
    let slice = &[10, 20, 30, 40];
    
    // Lectura
    for &num in slice {
        println!("{}", num);
    }
    
    // Con índice
    for (i, &num) in slice.iter().enumerate() {
        println!("Índice {}: {}", i, num);
    }
    
    // Ventanas
    for ventana in slice.windows(2) {
        println!("Ventana: {:?}", ventana);  // [10, 20], [20, 30], ...
    }
    
    // Chunks
    for chunk in slice.chunks(2) {
        println!("Chunk: {:?}", chunk);  // [10, 20], [30, 40]
    }
}
```

## HashMap<K, V>

Un mapa de clave-valor.

### Crear HashMap

```rust
use std::collections::HashMap;

fn main() {
    // Nuevo vacío
    let mut puntuaciones = HashMap::new();
    
    // Desde pares
    let pares = vec![("Rojo", 10), ("Azul", 20)];
    let colores: HashMap<&str, i32> = pares.into_iter().collect();
    
    println!("Colores: {:?}", colores);
}
```

### Métodos Esenciales

#### `insert(clave, valor)` - Añadir o Actualizar
- **Qué hace**: Inserta una clave-valor. Si la clave existe, actualiza el valor
- **Devuelve**: `Option` con el valor anterior (None si no existía)
- **Cuándo usarlo**: Agregar datos al mapa

```rust
use std::collections::HashMap;

let mut edades = HashMap::new();
edades.insert("Ana", 28);
edades.insert("Bob", 35);

// Si la clave ya existe, actualiza
if let Some(edad_anterior) = edades.insert("Ana", 29) {
    println!("Edad anterior de Ana: {}", edad_anterior);  // 28
}
```

#### `get(&clave)` - Obtener Valor
- **Qué hace**: Busca una clave y devuelve una referencia al valor
- **Devuelve**: `Option<&V>` - None si no existe
- **Complejidad**: O(1) promedio
- **Cuándo usarlo**: Consultar valores existentes sin modificar
- **IMPORTANTE**: Devuelve Option, debes manejarlo

```rust
let mut datos = HashMap::new();
datos.insert("Python", 10);
datos.insert("Rust", 8);

// Forma segura: con match
match datos.get("Rust") {
    Some(puntos) => println!("Rust: {} puntos", puntos),
    None => println!("Lenguaje no registrado"),
}

// Forma rápida: con unwrap_or
let puntos_python = datos.get("Python").unwrap_or(&0);
println!("Python: {}", puntos_python);
```

#### `remove(&clave)` - Eliminar y Extraer
- **Qué hace**: Elimina una clave-valor del mapa y devuelve el valor
- **Devuelve**: `Option<V>` - el valor eliminado o None
- **Cuándo usarlo**: Cuando necesitas remover y procesar el valor

```rust
let mut carrito = HashMap::new();
carrito.insert("manzanas", 3);
carrito.insert("naranjas", 2);

if let Some(cantidad) = carrito.remove("manzanas") {
    println!("Removidas {} manzanas", cantidad);  // 3
}

println!("Carrito: {:?}", carrito);  // Solo {naranjas: 2}
```

#### `contains_key(&clave)` - Verificar Existencia
- **Qué hace**: Devuelve `true` si la clave existe
- **Complejidad**: O(1) promedio
- **Cuándo usarlo**: Validación rápida antes de get/insert

```rust
let mut inventario = HashMap::new();
inventario.insert("laptop", 5);

if !inventario.contains_key("mouse") {
    inventario.insert("mouse", 20);
}

println!("{:?}", inventario);
```

### Iteración

#### `keys()` - Obtener Todas las Claves
- **Qué hace**: Devuelve iterador sobre todas las claves
- **Cuándo usarlo**: Listar todas las claves, búsquedas, validación

```rust
let mut frutas = HashMap::new();
frutas.insert("manzana", 5);
frutas.insert("plátano", 3);
frutas.insert("naranja", 7);

println!("Frutas disponibles:");
for fruta in frutas.keys() {
    println!("- {}", fruta);
}
```

#### `values()` - Obtener Todos los Valores
- **Qué hace**: Devuelve iterador sobre todos los valores
- **Cuándo usarlo**: Cálculos, estadísticas, búsquedas

```rust
let mut precios = HashMap::new();
precios.insert("café", 2.50);
precios.insert("té", 1.80);
precios.insert("jugo", 2.00);

let total: f64 = precios.values().sum();
println!("Valor total del inventario: {}", total);

let precio_máximo = precios.values().max_by(|a, b| a.partial_cmp(b).unwrap());
println!("Artículo más caro: ${:?}", precio_máximo);
```

#### `iter()` - Iterar Pares Clave-Valor
- **Qué hace**: Devuelve tuplas de referencias `(&K, &V)`
- **Cuándo usarlo**: Cuando necesitas ambos, clave y valor

```rust
let mut tareas = HashMap::new();
tareas.insert("tarea1", "completada");
tareas.insert("tarea2", "pendiente");
tareas.insert("tarea3", "completada");

for (id, estado) in tareas.iter() {
    println!("{}: {}", id, estado);
}
```

#### `iter_mut()` - Modificar Valores During Iteración
- **Qué hace**: Devuelve referencias mutables para modificar valores
- **Cuándo usarlo**: Cuando necesitas actualizar valores basándose en lógica

```rust
let mut puntuaciones = HashMap::new();
puntuaciones.insert("nivel1", 10);
puntuaciones.insert("nivel2", 20);

// Doblar todas las puntuaciones
for puntos in puntuaciones.values_mut() {
    *puntos *= 2;
}

println!("{:?}", puntuaciones);  // {level1: 20, level2: 40}
```

### Operaciones Avanzadas

#### `entry()` - Acceso Eficiente
- **Qué hace**: Acceso optimizado para operaciones condicionadas (check + insert/update)
- **Evita**: Dos búsquedas (una con get, otra con insert)
- **Cuándo usarlo**: Cuando necesitas "obtener si existe, sino insertar"

```rust
use std::collections::HashMap;

let mut contadores = HashMap::new();

// Sin entry (ineficiente: 2 búsquedas)
if !contadores.contains_key("rust") {
    contadores.insert("rust", 0);
}
*contadores.get_mut("rust").unwrap() += 1;

// Con entry (eficiente: 1 búsqueda)
*contadores.entry("python").or_insert(0) += 1;

println!("{:?}", contadores);
```

#### `or_insert()` - Insertar si No Existe
- **Qué hace**: Si la clave no existe, inserta el valor. Devuelve referencia mutable
- **Cuándo usarlo**: Inicialización con valores por defecto

```rust
let mut usuarios = HashMap::new();

// Obtener o crear
let contador = usuarios.entry("Alice").or_insert(0);
*contador += 5;

let contador = usuarios.entry("Alice").or_insert(0);
*contador += 3;

println!("{:?}", usuarios);  // {Alice: 8}
```

#### `or_insert_with()` - Insertar Condicionalmente
- **Qué hace**: Como or_insert, pero calcula el valor con una función
- **Cuándo usarlo**: Cuando insertar requiere cálculo costoso

```rust
let mut caché = HashMap::new();

// Solo calcula si la clave no existe
let valor = caché.entry("clave1").or_insert_with(|| {
    println!("Calculando...");
    5 * 2  // Cálculo costoso
});

println!("Valor: {}", valor);

// Segunda vez: No entra al bloque
let valor2 = caché.entry("clave1").or_insert_with(|| {
    println!("Esto NO se ejecuta");
    999
});
```

## HashSet<T>

Un conjunto de valores únicos sin claves.

### Crear HashSet

```rust
use std::collections::HashSet;

fn main() {
    // Vacío
    let mut conjunto = HashSet::new();
    
    // Desde iterador
    let nums = vec![1, 2, 3, 2, 1];
    let único: HashSet<i32> = nums.into_iter().collect();
    
    println!("Único: {:?}", único);  // {1, 2, 3}
}
```

### Operaciones

```rust
use std::collections::HashSet;

fn main() {
    let mut conjunto = HashSet::new();
    
    // Insertar
    conjunto.insert(1);
    conjunto.insert(2);
    conjunto.insert(3);
    
    // ¿Contiene?
    if conjunto.contains(&2) {
        println!("Contiene 2");
    }
    
    // Eliminar
    conjunto.remove(&2);
    println!("{:?}", conjunto);  // {1, 3}
}
```

### Operaciones de Conjuntos

```rust
use std::collections::HashSet;

fn main() {
    let a: HashSet<i32> = [1, 2, 3].iter().copied().collect();
    let b: HashSet<i32> = [2, 3, 4].iter().copied().collect();
    
    // Unión
    let unión: HashSet<_> = a.union(&b).copied().collect();
    println!("Unión: {:?}", unión);  // {1, 2, 3, 4}
    
    // Intersección
    let intersección: HashSet<_> = a.intersection(&b).copied().collect();
    println!("Intersección: {:?}", intersección);  // {2, 3}
    
    // Diferencia
    let diferencia: HashSet<_> = a.difference(&b).copied().collect();
    println!("Diferencia: {:?}", diferencia);  // {1}
    
    // Diferencia simétrica
    let simetría: HashSet<_> = a.symmetric_difference(&b).copied().collect();
    println!("Diferencia simétrica: {:?}", simetría);  // {1, 4}
}
```

## BTreeMap y BTreeSet

Versiones ordenadas de HashMap y HashSet.

```rust
use std::collections::BTreeMap;

fn main() {
    let mut mapa = BTreeMap::new();
    mapa.insert("b", 2);
    mapa.insert("a", 1);
    mapa.insert("c", 3);
    
    // Mantiene orden de claves
    for (k, v) in &mapa {
        println!("{}: {}", k, v);  // a:1, b:2, c:3
    }
}
```

## Comparación de Colecciones

| Colección | Acceso | Inserción | Búsqueda | Ordenado |
|-----------|--------|-----------|----------|----------|
| Vec<T> | O(1) | O(n) | O(n) | No |
| HashMap<K,V> | O(1) | O(1) | O(1) | No |
| HashSet<T> | - | O(1) | O(1) | No |
| BTreeMap<K,V> | O(log n) | O(log n) | O(log n) | Sí |
| BTreeSet<T> | - | O(log n) | O(log n) | Sí |

## Ejemplo Práctico 1: Contador de Palabras

```rust
use std::collections::HashMap;

fn contar_palabras(texto: &str) -> HashMap<&str, u32> {
    let mut contador = HashMap::new();
    
    for palabra in texto.split_whitespace() {
        let palabra_limpia = palabra.trim_matches(|c: char| !c.is_alphanumeric());
        let cuenta = contador.entry(palabra_limpia).or_insert(0);
        *cuenta += 1;
    }
    
    contador
}

fn main() {
    let texto = "el el la la la gato";
    let resultado = contar_palabras(texto);
    
    // Ordenar por frecuencia
    let mut pares: Vec<_> = resultado.into_iter().collect();
    pares.sort_by(|a, b| b.1.cmp(&a.1));
    
    for (palabra, cuenta) in pares {
        println!("{}: {}", palabra, cuenta);
    }
}
```

Output:
```
la: 3
el: 2
gato: 1
```

## Ejemplo Práctico 2: Procesamiento de Datos

```rust
fn main() {
    // Datos de temperaturas
    let mut temperaturas = vec![15.2, 18.5, 22.3, 19.8, 16.1, 21.5];
    
    // Información
    println!("Cantidad de mediciones: {}", temperaturas.len());
    println!("Mínima: {}", temperaturas.iter().fold(f64::INFINITY, |a, &b| a.min(b)));
    println!("Máxima: {}", temperaturas.iter().fold(f64::NEG_INFINITY, |a, &b| a.max(b)));
    
    // Promedio
    let promedio: f64 = temperaturas.iter().sum::<f64>() / temperaturas.len() as f64;
    println!("Promedio: {:.1}", promedio);
    
    // Filtrar valores altos
    let altas: Vec<_> = temperaturas.iter().filter(|&&t| t > 20.0).copied().collect();
    println!("Temperaturas altas: {:?}", altas);
    
    // Ordenar
    temperaturas.sort_by(|a, b| a.partial_cmp(b).unwrap());
    println!("Ordenadas: {:?}", temperaturas);
}
```

## Ejemplo Práctico 3: Procesamiento de Strings

```rust
fn main() {
    let log = "ERROR: conexión fallida\nWARN: timeout\nERROR: recurso no encontrado";
    
    // Contar líneas
    println!("Líneas totales: {}", log.lines().count());
    
    // Filtrar errores
    let errores: Vec<_> = log.lines()
        .filter(|line| line.starts_with("ERROR"))
        .collect();
    
    println!("Errores encontrados: {}", errores.len());
    for error in errores {
        println!("- {}", error);
    }
    
    // Extraer mensajes
    let mensajes: Vec<_> = log.lines()
        .map(|line| line.split(": ").nth(1).unwrap_or("desconocido"))
        .collect();
    
    println!("Mensajes: {:?}", mensajes);
}
```

## Rendimiento: Cuándo Usar Cada Colección

```rust
fn main() {
    // Vec: cuando necesitas índice rápido
    let v = vec![1, 2, 3, 4, 5];
    println!("Acceso rápido: {}", v[0]);  // O(1)
    
    // HashMap: cuando necesitas búsqueda por clave
    let mut mapa = std::collections::HashMap::new();
    mapa.insert("nombre", "Ana");
    println!("Búsqueda rápida: {:?}", mapa.get("nombre"));  // O(1)
    
    // HashSet: cuando solo importan valores únicos
    let mut conjunto = std::collections::HashSet::new();
    conjunto.insert(1);
    println!("¿Único? {}", conjunto.contains(&1));  // O(1)
    
    // BTreeMap: cuando necesitas ordenamiento
    let mut btree = std::collections::BTreeMap::new();
    btree.insert("a", 1);
    btree.insert("b", 2);
    for (k, v) in &btree {
        println!("{}: {}", k, v);  // Siempre ordenado
    }
}
```

## Siguiente Paso

Ahora dominas todas las colecciones principales de Rust. El próximo capítulo enseña cómo **manejar errores** de forma segura con `Result`.

---

**Consejo**: Las colecciones son fundamentales. Domina sus métodos - son lo que usarás a diario en Rust. Experimenta en el [Rust Playground](https://play.rust-lang.org/) con diferentes casos de uso.
