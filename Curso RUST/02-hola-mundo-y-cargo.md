# 02 - Hola, Mundo! y Cargo

## Tu Primer Programa en Rust

Vamos a crear el programa más simple: "Hola, mundo!"

### Opción 1: Sin Cargo (Forma Manual)

Crea una carpeta y un archivo llamado `main.rs`:

```rust
fn main() {
    println!("¡Hola, mundo!");
}
```

Compila el código:

```bash
rustc main.rs
```

Ejecuta el programa resultante:

```bash
# En Windows
main.exe

# En macOS y Linux
./main
```

Verás en pantalla:
```
¡Hola, mundo!
```

### Opción 2: Con Cargo (Forma Recomendada)

Cargo es el gestor de proyectos de Rust. Te permite:
- Gestionar dependencias fácilmente
- Ejecutar tests
- Documentar código
- Publicar librerías

Crea un nuevo proyecto:

```bash
cargo new hola_mundo
```

Entra en la carpeta:

```bash
cd hola_mundo
```

Verás esta estructura:

```
hola_mundo/
├── Cargo.toml        # Configuración del proyecto
├── src/
│   └── main.rs       # Tu código
└── .gitignore        # Archivos ignorados por Git
```

### Entendiendo `Cargo.toml`

Este archivo es como el "corazón" de tu proyecto:

```toml
[package]
name = "hola_mundo"
version = "0.1.0"
edition = "2024"

[dependencies]
# Aquí van las librerías externas que necesites
```

- `name`: Nombre de tu proyecto
- `version`: Versión actual del proyecto
- `edition`: La edición de Rust que usas (2024 es la más nueva)
- `[dependencies]`: Donde especificas librerías externas

### Tu Primer Archivo main.rs

Por defecto, `cargo new` crea un archivo `main.rs` con:

```rust
fn main() {
    println!("Hello, world!");
}
```

Cámbialo por:

```rust
fn main() {
    println!("¡Hola, mundo!");
}
```

## Ejecutar tu Proyecto

### Modo Desarrollo

Para compilar y ejecutar durante el desarrollo:

```bash
cargo run
```

Output esperado:
```
   Compiling hola_mundo v0.1.0
    Finished dev [unoptimized + debuginfo] target(s) in 0.42s
     Running `target/debug/hola_mundo`
¡Hola, mundo!
```

La primera vez tarda más porque Cargo compila todo. Las siguientes veces es más rápido.

### Modo Release (Optimizado)

Para un programa listo para producción:

```bash
cargo build --release
```

Esto compila tu código con optimizaciones, lo que lo hace:
- Más rápido en ejecución
- Más grande en tamaño
- Tarda más en compilar

El binario compilado está en `target/release/`

## Archivos Generados

Cuando compiles tu código, Cargo crea:

```
hola_mundo/
├── target/              # Archivos compilados
│   ├── debug/           # Compilación de desarrollo
│   └── release/         # Compilación optimizada
├── Cargo.lock           # Archivo de control de versiones de dependencias
├── Cargo.toml
├── src/
│   └── main.rs
└── .gitignore
```

## Comandos Útiles de Cargo

| Comando | Descripción |
|---------|-------------|
| `cargo new` | Crear nuevo proyecto |
| `cargo build` | Compilar el proyecto |
| `cargo run` | Compilar y ejecutar |
| `cargo check` | Verificar si el código compila (sin crear ejecutable) |
| `cargo clean` | Eliminar la carpeta `target/` |
| `cargo doc` | Generar documentación |
| `cargo test` | Ejecutar tests |
| `cargo fmt` | Formatear automáticamente el código |
| `cargo clippy` | Verificar mejoras de código |

## Desglosando el Código

Analicemos línea por línea:

```rust
fn main() {
    println!("¡Hola, mundo!");
}
```

- `fn` - Palabra clave para declarar una función
- `main()` - Nombre especial: es el punto de entrada de todo programa Rust
- `{}` - Bloque que contiene el cuerpo de la función
- `println!()` - Macro que imprime texto en la pantalla (nota la `!`)
- `"¡Hola, mundo!"` - Texto entre comillas (String literal)
- `;` - Termina la declaración

## ¿Qué es una Macro?

En Rust, `println!` es una **macro**, no una función. Las macas se reconocen por el `!` al final.

Las macros son más poderosas que las funciones:
- Pueden tomar un número variable de argumentos
- Se expanden en tiempo de compilación
- Pueden generar código

Ahora mismo, solo necesitas saber que `println!()` imprime texto.

## Imprimir con Formato

`println!()` permite más que solo texto:

```rust
fn main() {
    let nombre = "Rust";
    let version = "1.90.0";
    
    println!("Bienvenido a {}", nombre);
    println!("Versión: {}", version);
    println!("{} {}", nombre, version);
}
```

Output:
```
Bienvenido a Rust
Versión: 1.90.0
Rust 1.90.0
```

Los `{}` son placeholders que se reemplazan con los valores siguientes.

## Estructuración de Código

En Rust, los proyectos con Cargo siguen una estructura estándar:

- **`main.rs`** - Archivo principal del ejecutable (aplicación)
- **`lib.rs`** - Archivo principal de una librería (código reutilizable)
- **Otros archivos en `src/`** - Módulos que puedes organizar

Aprenderás más sobre módulos en capítulos posteriores.

## Siguiente Paso

Ahora que sabes cómo ejecutar código Rust, vamos a aprender sobre variables, tipos de datos y las primeras operaciones. ¡Continuemos!

---

**Nota**: Si ves mensajes de error, revisa que:
1. Instalaste Rust correctamente
2. Estás en la carpeta correcta del proyecto
3. El archivo `main.rs` tiene la sintaxis correcta
