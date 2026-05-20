# 16 - Entrada/Salida y Trabajo con Archivos

Trabajar con el sistema de archivos y entrada/salida es común en Rust.

## Leer Líneas desde Entrada Estándar

```rust
use std::io;

fn main() {
    println!("¿Cuál es tu nombre?");
    
    let mut entrada = String::new();
    io::stdin().read_line(&mut entrada)
        .expect("Fallo al leer línea");
    
    // entrada contiene el newline al final
    let nombre = entrada.trim();
    println!("Hola, {}!", nombre);
}
```

### Leer Entrada Entera

```rust
use std::io::{self, Read};

fn main() {
    let mut contenido = String::new();
    io::stdin().read_to_string(&mut contenido)
        .expect("Error al leer");
    
    println!("Leí:\n{}", contenido);
}
```

## Leer Archivos

```rust
use std::fs;

fn main() {
    let contenido = fs::read_to_string("archivo.txt")
        .expect("No se pudo leer el archivo");
    
    println!("Contenido:\n{}", contenido);
}
```

### Manejo de Errores Mejor

```rust
use std::fs;
use std::io;

fn leer_archivo(ruta: &str) -> io::Result<String> {
    fs::read_to_string(ruta)
}

fn main() {
    match leer_archivo("archivo.txt") {
        Ok(contenido) => println!("{}", contenido),
        Err(e) => println!("Error: {}", e),
    }
}
```

### Leer Línea por Línea

```rust
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let archivo = File::open("datos.txt")
        .expect("No se puede abrir el archivo");
    
    let lector = BufReader::new(archivo);
    
    for línea in lector.lines() {
        match línea {
            Ok(contenido) => println!("{}", contenido),
            Err(e) => println!("Error: {}", e),
        }
    }
}
```

## Escribir Archivos

### Crear y Escribir

```rust
use std::fs;

fn main() {
    let contenido = "Hola, mundo!";
    
    fs::write("salida.txt", contenido)
        .expect("No se pudo escribir el archivo");
}
```

### Agregar Contenido

```rust
use std::fs::OpenOptions;
use std::io::Write;

fn main() {
    let mut archivo = OpenOptions::new()
        .append(true)
        .open("logs.txt")
        .expect("No se puede abrir");
    
    writeln!(archivo, "Nueva línea de log")
        .expect("No se puede escribir");
}
```

## Trabajar con Rutas

```rust
use std::path::Path;

fn main() {
    let ruta = Path::new("datos/archivo.txt");
    
    println!("¿Existe? {}", ruta.exists());
    println!("Nombre: {:?}", ruta.file_name());
    println!("Extensión: {:?}", ruta.extension());
}
```

## Directorios

```rust
use std::fs;

fn main() {
    // Crear directorio
    fs::create_dir_all("output/resultados")
        .expect("No se pudo crear el directorio");
    
    // Listar contenido
    for entrada in fs::read_dir(".").unwrap() {
        let entrada = entrada.unwrap();
        let ruta = entrada.path();
        println!("{:?}", ruta);
    }
}
```

## Metadatos

```rust
use std::fs;

fn main() {
    let metadatos = fs::metadata("archivo.txt")
        .expect("Error al obtener metadatos");
    
    println!("¿Es archivo? {}", metadatos.is_file());
    println!("¿Es directorio? {}", metadatos.is_dir());
    println!("Tamaño: {} bytes", metadatos.len());
}
```

## Serialización JSON

Primero, añade a `Cargo.toml`:

```toml
[dependencies]
serde_json = "1.0"
```

### Escribir JSON

```rust
use serde_json::json;
use std::fs;

fn main() {
    let datos = json!({
        "nombre": "Ana",
        "edad": 30,
        "ciudad": "Madrid"
    });
    
    fs::write("datos.json", datos.to_string())
        .expect("Error al escribir");
}
```

### Leer JSON

```rust
use serde_json::json;
use std::fs;

fn main() {
    let contenido = fs::read_to_string("datos.json")
        .expect("No se pudo leer");
    
    let datos: serde_json::Value = serde_json::from_str(&contenido)
        .expect("JSON inválido");
    
    println!("Nombre: {}", datos["nombre"]);
    println!("Edad: {}", datos["edad"]);
}
```

## Ejemplo Práctico: Logger Simple

```rust
use std::fs::OpenOptions;
use std::io::Write;
use std::time::SystemTime;

struct Logger {
    archivo: String,
}

impl Logger {
    fn nuevo(archivo: &str) -> Self {
        Logger {
            archivo: archivo.to_string(),
        }
    }
    
    fn log(&self, mensaje: &str) {
        let timestamp = format!("{:?}", SystemTime::now());
        let linea = format!("[{}] {}\n", timestamp, mensaje);
        
        let mut f = OpenOptions::new()
            .create(true)
            .append(true)
            .open(&self.archivo)
            .expect("Error al abrir log");
        
        f.write_all(linea.as_bytes())
            .expect("Error al escribir log");
    }
}

fn main() {
    let logger = Logger::nuevo("app.log");
    
    logger.log("Aplicación iniciada");
    logger.log("Procesando datos");
    logger.log("Aplicación finalizada");
}
```

## Procesamiento de CSV

Con la crate `csv`:

```toml
[dependencies]
csv = "1.1"
```

```rust
use std::fs::File;
use csv::Reader;

fn main() {
    let archivo = File::open("datos.csv")
        .expect("No se puede abrir");
    
    let mut lector = Reader::from_reader(archivo);
    
    for resultado in lector.records() {
        let registro = resultado.expect("Error al leer");
        println!("{:?}", registro);
    }
}
```

## Siguiente Paso

Ya sabes interactuar con el sistema. El próximo capítulo enseña **buenas prácticas** para escribir código profesional en Rust.

---

**Consejo**: Siempre maneja errores de I/O. Los archivos pueden no existir, la red falla, etc.
