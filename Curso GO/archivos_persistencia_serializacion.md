# Archivos, Persistencia y Serialización en Go

Trabajar con archivos es fundamental: leer datos, guardar información, persistir estado. Go proporciona herramientas poderosas en el paquete `os` e `io`.

## 1. Operaciones básicas con archivos

### Leer un archivo completo

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Leer archivo completo (simple)
    contenido, err := os.ReadFile("archivo.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println(string(contenido))
}
```

### Leer línea por línea

```go
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    archivo, err := os.Open("archivo.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer archivo.Close()  // Cerrar al terminar
    
    scanner := bufio.NewScanner(archivo)
    for scanner.Scan() {
        linea := scanner.Text()
        fmt.Println(linea)
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Println("Error de lectura:", err)
    }
}
```

### Escribir en archivo

```go
import (
    "fmt"
    "os"
)

func main() {
    // Escribir contenido
    contenido := "Hola, mundo!\n"
    
    err := os.WriteFile("salida.txt", []byte(contenido), 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("Archivo escrito")
}
```

### Append (agregar al final)

```go
import (
    "os"
)

func main() {
    archivo, err := os.OpenFile(
        "log.txt",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0644,
    )
    if err != nil {
        panic(err)
    }
    defer archivo.Close()
    
    archivo.WriteString("Nueva línea en log\n")
}
```

**Flags**:
- `O_RDONLY`: Solo lectura
- `O_WRONLY`: Solo escritura
- `O_RDWR`: Lectura y escritura
- `O_CREATE`: Crear si no existe
- `O_APPEND`: Append al final
- `O_TRUNCATE`: Limpiar contenido

### Permisos (0644, 0755)

```
0644 = rw-r--r--
  |   |  |  |
  |   |  |  +---- otros (4 = leer)
  |   |  +------- grupo (4 = leer)
  |   +---------- propietario (6 = leer+escribir)
  
0755 = rwxr-xr-x
  |   |  |  |
  |   |  |  +---- otros (5 = leer+ejecutar)
  |   |  +------- grupo (5 = leer+ejecutar)
  |   +---------- propietario (7 = todo)
```

---

## 2. Trabajar con directorios

### Listar archivos

```go
import (
    "fmt"
    "os"
)

func main() {
    entries, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    for _, e := range entries {
        if e.IsDir() {
            fmt.Println("[DIR]", e.Name())
        } else {
            fmt.Println("[FILE]", e.Name())
        }
    }
}
```

### Crear directorio

```go
os.Mkdir("mi_carpeta", 0755)      // Solo un nivel
os.MkdirAll("a/b/c", 0755)        // Todos los niveles
```

### Eliminar archivo/directorio

```go
os.Remove("archivo.txt")      // Eliminar archivo
os.RemoveAll("carpeta")       // Eliminar carpeta y contenido
```

### Información de archivo

```go
info, err := os.Stat("archivo.txt")
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println("Nombre:", info.Name())
fmt.Println("Tamaño:", info.Size(), "bytes")
fmt.Println("Modificado:", info.ModTime())
fmt.Println("Es directorio:", info.IsDir())
```

---

## 3. Lectura y escritura eficiente

### Usar bufio para mejor rendimiento

```go
import (
    "bufio"
    "os"
)

func main() {
    archivo, _ := os.Open("archivo.txt")
    defer archivo.Close()
    
    // Reader bufferizado
    reader := bufio.NewReader(archivo)
    
    // Leer byte por byte
    b, _ := reader.ReadByte()
    
    // Leer hasta un delimitador
    linea, _ := reader.ReadString('\n')
    
    // Leer línea completa
    linea, _ := reader.ReadLine()  // sin salto
}
```

### Escritura bufferizada

```go
import (
    "bufio"
    "os"
)

func main() {
    archivo, _ := os.Create("salida.txt")
    defer archivo.Close()
    
    writer := bufio.NewWriter(archivo)
    
    writer.WriteString("Línea 1\n")
    writer.WriteString("Línea 2\n")
    
    writer.Flush()  // ⚠️ Vaciar buffer a disco
}
```

**Importante**: `Flush()` asegura que todo se escriba.

---

## 4. Serialización: JSON

JSON es el formato más común para datos estructurados.

### Struct a JSON

```go
import (
    "encoding/json"
    "fmt"
)

type Persona struct {
    Nombre string `json:"nombre"`
    Edad   int    `json:"edad"`
}

func main() {
    p := Persona{"Alice", 30}
    
    // Struct → JSON
    jsonBytes, _ := json.Marshal(p)
    fmt.Println(string(jsonBytes))
    // {"nombre":"Alice","edad":30}
    
    // Con indentación
    jsonIndent, _ := json.MarshalIndent(p, "", "  ")
    fmt.Println(string(jsonIndent))
}
```

### JSON a Struct

```go
import (
    "encoding/json"
    "fmt"
)

func main() {
    jsonStr := `{"nombre":"Bob","edad":25}`
    
    var p Persona
    err := json.Unmarshal([]byte(jsonStr), &p)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println(p.Nombre, p.Edad)  // Bob 25
}
```

### Tags JSON

```go
type Usuario struct {
    ID       int    `json:"id"`
    Nombre   string `json:"nombre"`
    Email    string `json:"email,omitempty"`  // omitempty: no incluir si vacío
    Activo   bool   `json:"-"`                // -: nunca incluir
    Interno  string                           // sin tag: se incluye como "Interno"
}
```

### Leer/escribir JSON de archivos

```go
import (
    "encoding/json"
    "os"
)

func GuardarJSON(archivo string, data interface{}) error {
    jsonBytes, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(archivo, jsonBytes, 0644)
}

func CargarJSON(archivo string, data interface{}) error {
    jsonBytes, err := os.ReadFile(archivo)
    if err != nil {
        return err
    }
    return json.Unmarshal(jsonBytes, data)
}

func main() {
    p := Persona{"Alice", 30}
    GuardarJSON("persona.json", p)
    
    var p2 Persona
    CargarJSON("persona.json", &p2)
    fmt.Println(p2)
}
```

---

## 5. Serialización: CSV

Para datos tabulares.

### Leer CSV

```go
import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    archivo, _ := os.Open("datos.csv")
    defer archivo.Close()
    
    reader := csv.NewReader(archivo)
    registros, _ := reader.ReadAll()
    
    for _, registro := range registros {
        fmt.Println(registro)
        // registro es []string
    }
}
```

### Escribir CSV

```go
import (
    "encoding/csv"
    "os"
)

func main() {
    archivo, _ := os.Create("salida.csv")
    defer archivo.Close()
    
    writer := csv.NewWriter(archivo)
    defer writer.Flush()
    
    // Escribir encabezados
    writer.Write([]string{"ID", "Nombre", "Edad"})
    
    // Escribir datos
    writer.Write([]string{"1", "Alice", "30"})
    writer.Write([]string{"2", "Bob", "25"})
}
```

### Usar con structs

```go
import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type Persona struct {
    Nombre string
    Edad   int
}

func GuardarCSV(archivo string, personas []Persona) error {
    w, _ := os.Create(archivo)
    defer w.Close()
    
    writer := csv.NewWriter(w)
    defer writer.Flush()
    
    for _, p := range personas {
        writer.Write([]string{
            p.Nombre,
            strconv.Itoa(p.Edad),
        })
    }
    return nil
}
```

---

## 6. Serialización: XML

Menos común que JSON, pero a veces necesario.

```go
import (
    "encoding/xml"
    "fmt"
)

type Persona struct {
    Nombre string `xml:"nombre"`
    Edad   int    `xml:"edad"`
}

func main() {
    p := Persona{"Alice", 30}
    
    // Struct → XML
    xmlBytes, _ := xml.MarshalIndent(p, "", "  ")
    fmt.Println(string(xmlBytes))
    
    // XML → Struct
    xmlStr := `<Persona><nombre>Bob</nombre><edad>25</edad></Persona>`
    var p2 Persona
    xml.Unmarshal([]byte(xmlStr), &p2)
    fmt.Println(p2)
}
```

---

## 7. Persistencia: Base de datos simple con JSON

```go
import (
    "encoding/json"
    "os"
    "sync"
)

// Base de datos simple en archivo
type DB struct {
    archivo string
    mu      sync.Mutex
    datos   map[string]interface{}
}

func AbrirDB(archivo string) *DB {
    db := &DB{
        archivo: archivo,
        datos:   make(map[string]interface{}),
    }
    
    // Cargar datos existentes
    if jsonBytes, err := os.ReadFile(archivo); err == nil {
        json.Unmarshal(jsonBytes, &db.datos)
    }
    
    return db
}

func (db *DB) Guardar(clave string, valor interface{}) error {
    db.mu.Lock()
    defer db.mu.Unlock()
    
    db.datos[clave] = valor
    
    // Guardar a disco
    jsonBytes, _ := json.MarshalIndent(db.datos, "", "  ")
    return os.WriteFile(db.archivo, jsonBytes, 0644)
}

func (db *DB) Obtener(clave string) interface{} {
    db.mu.Lock()
    defer db.mu.Unlock()
    
    return db.datos[clave]
}

func main() {
    db := AbrirDB("datos.json")
    
    db.Guardar("usuario1", map[string]string{
        "nombre": "Alice",
        "email":  "alice@example.com",
    })
    
    fmt.Println(db.Obtener("usuario1"))
}
```

---

## 8. Persistencia: Binary (formato eficiente)

Para datos que necesitan compresión o acceso rápido.

```go
import (
    "encoding/gob"
    "os"
)

type Cuenta struct {
    Titular string
    Saldo   float64
}

func GuardarBinary(archivo string, data interface{}) error {
    f, _ := os.Create(archivo)
    defer f.Close()
    
    encoder := gob.NewEncoder(f)
    return encoder.Encode(data)
}

func CargarBinary(archivo string, data interface{}) error {
    f, _ := os.Open(archivo)
    defer f.Close()
    
    decoder := gob.NewDecoder(f)
    return decoder.Decode(data)
}

func main() {
    cuenta := Cuenta{"Alice", 1000}
    GuardarBinary("cuenta.bin", cuenta)
    
    var cuenta2 Cuenta
    CargarBinary("cuenta.bin", &cuenta2)
    fmt.Println(cuenta2)
}
```

---

## 9. Copiar archivos

```go
import (
    "io"
    "os"
)

func CopiarArchivo(origen, destino string) error {
    src, _ := os.Open(origen)
    defer src.Close()
    
    dst, _ := os.Create(destino)
    defer dst.Close()
    
    // Copiar contenido
    _, err := io.Copy(dst, src)
    return err
}

func main() {
    CopiarArchivo("archivo.txt", "copia.txt")
}
```

---

## 10. Trabajar con paths

```go
import (
    "os"
    "path/filepath"
)

func main() {
    // Construir path portátil
    ruta := filepath.Join("carpeta", "subcarpeta", "archivo.txt")
    
    // Obtener directorio
    dir := filepath.Dir(ruta)
    
    // Obtener nombre
    nombre := filepath.Base(ruta)
    
    // Extensión
    ext := filepath.Ext(ruta)
    
    // Path absoluto
    abs, _ := filepath.Abs("./archivo.txt")
}
```

---

## 11. Tabla de formatos comunes

| Formato | Legible | Eficiente | Casos de uso |
|---------|---------|-----------|--------------|
| JSON | ✅ Sí | Medio | APIs, config |
| CSV | ✅ Sí | Medio | Datos tabulares |
| XML | ✅ Sí | Bajo | SOAP, legacy |
| Binary | ❌ No | ✅ Alto | Bases datos |
| Text | ✅ Sí | Medio | Logs |

---

## 12. Mejores prácticas

### ✅ Bien
```go
// Siempre cerrar archivos
archivo, err := os.Open("archivo.txt")
if err != nil {
    return err
}
defer archivo.Close()

// Usar bufio para lectura/escritura eficiente
reader := bufio.NewReader(archivo)
writer := bufio.NewWriter(archivo)

// Verificar errores
if err != nil {
    fmt.Println("Error:", err)
    return err
}

// JSON con tags
type Persona struct {
    Nombre string `json:"nombre"`
    Edad   int    `json:"edad"`
}

// Separar lógica de I/O
func GuardarDatos(archivo string, datos interface{}) error {
    jsonBytes, _ := json.Marshal(datos)
    return os.WriteFile(archivo, jsonBytes, 0644)
}
```

### ❌ Mal
```go
// No cerrar archivos → memory leak
archivo, _ := os.Open("archivo.txt")
contenido, _ := os.ReadFile("archivo.txt")

// Ignorar errores
contents, _ := os.ReadFile("archivo.txt")

// Leer archivo entero si es muy grande
contenido, _ := os.ReadFile("archivo-1GB.bin")

// JSON sin estructura
data := map[string]interface{}{}  // Menos tipo-seguro
```

---

## 12. Conclusión - Persistencia en el Mundo Real

Trabajar con archivos y persistencia en Go es una tarea diaria. Hemos aprendido:

### Los tres pilares de la persistencia

1. **Lectura/Escritura**: `os.ReadFile`, `os.WriteFile`, bufio
2. **Serialización**: JSON, CSV, XML, Binary
3. **Bases de datos**: Desde archivos simples hasta SQL

### Guía rápida: ¿Qué usar?

| Caso | Solución |
|------|----------|
| Configuración simple | JSON en archivo |
| Datos estructurados | Struct → JSON |
| Datos tabulares | CSV |
| Cache rápido | Binary con gob |
| Datos complejos | Base de datos SQL |
| Logs | Append a archivo de texto |

### Antipatrones a evitar

```go
// ❌ Ignorar errores
datos, _ := os.ReadFile("file.txt")

// ❌ No cerrar archivos (memory leak)
f := os.Open("file.txt")
contenido, _ := ioutil.ReadAll(f)

// ❌ Concatenar strings en loop
for _, item := range items {
    contenido += item  // Ineficiente
}

// ❌ Cargar archivo gigante en memoria
todo := ioutil.ReadFile("1GB_file.bin")
```

### Patrones correctos

```go
// ✅ Siempre cerrar con defer
f, err := os.Open("file.txt")
if err != nil {
    return err
}
defer f.Close()

// ✅ Usar readers/writers
reader := bufio.NewReader(f)
line, _ := reader.ReadString('\n')

// ✅ JSON para datos
type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

// ✅ Builders para eficiencia
var buf strings.Builder
for _, item := range items {
    buf.WriteString(item)
}
resultado := buf.String()
```

### Flujo típico: Read → Parse → Transform → Write

```go
// 1. Leer JSON
data, _ := os.ReadFile("input.json")

// 2. Parsear
var entrada []Persona
json.Unmarshal(data, &entrada)

// 3. Transformar
salida := procesarPersonas(entrada)

// 4. Escribir
jsonSalida, _ := json.MarshalIndent(salida, "", "  ")
os.WriteFile("output.json", jsonSalida, 0644)
```

### Performance: Cómo saber qué es rápido

- **O(1)**: Acceso a byte específico de small file
- **O(n)**: Leer archivo completo
- **O(n log n)**: Parsear JSON
- **O(n)**: Serializar a JSON
- **O(n)**: Escribir archivo

Si tu archivo es > 100MB, considera streaming en lugar de cargas completas.

### La filosofía de Go para I/O

**Go hace I/O simple pero explícito**:
- Check de errores (no excepciones)
- Readers/Writers comunes (io.Reader, io.Writer)
- Buffering cuando necesitas performance
- Defer para limpieza

Con esta filosofía, podrás manejar persistencia desde chicos POCs hasta sistemas enormes de producción.

**Go I/O te enseña a ser respetuoso con la memoria y los recursos. Aprende bien.**
