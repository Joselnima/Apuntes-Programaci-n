# Estructura Básica de un Archivo Go

> **La anatomía de todo programa Go: desde el primer carácter hasta la ejecución**

---

## ¿Por qué Go tiene una estructura específica?

Cuando abres un archivo `.go`, ves un **patrón muy específico**. No es capricho: es diseño intencional.

Go **obliga** una estructura clara porque:
1. Reduce ambigüedad
2. Hace visible qué dependen de qué
3. Facilita encontrar cosas rápidamente
4. Evita problemas de inicialización

Un archivo Go tiene **siempre** este orden:
1. **Package declaration**
2. **Import statements**
3. **Constantes y variables**
4. **Funciones y tipos**

Veamos cada parte.

---

## Parte I: Package Declaration

### ¿Qué es un package?

Un **package** es un grupo de código Go compilado junto. Todo archivo debe pertenencer a un package.

### package main

```go
package main
```

- `main` es el package especial: **el punto de entrada** de tu programa
- **Solo puede haber un `func main()` en todo tu programa**
- Sin `package main`, Go no sabe dónde comenzar

### Otros packages

```go
package usuarios    // Package privado (local a tu proyecto)
package math        // Package de stdlib de Go
```

**Regla importante:**
- Un **archivo** pertenece a un package
- Una **carpeta** puede tener múltiples archivos del mismo package
- Todos los `.go` en una carpeta **deben tener el mismo package**

### Ejemplo de estructura

```
proyecto/
├── main.go          (package main)
├── config.go        (package main)
├── usuarios/
│   ├── usuario.go   (package usuarios)
│   └── service.go   (package usuarios)
└── productos/
    ├── producto.go  (package productos)
    └── service.go   (package productos)
```

---

## Parte II: Importaciones (Import Statements)

### Importación simple

```go
package main

import "fmt"       // Import un paquete standard
```

### Múltiples importaciones

```go
// Forma 1: Múltiples import
import "fmt"
import "strings"
import "strconv"

// Forma 2: Agrupadas (recomendado)
import (
    "fmt"
    "strings"
    "strconv"
)
```

**✅ Siempre usa la forma agrupada para múltiples imports.**

### Orden de imports (convención Go)

```go
// ✅ Orden correcto:
import (
    // Stdlib primero (alfabético)
    "fmt"
    "math"
    "net/http"
    "strings"

    // Espacios en blanco entre stdlib y terceros

    // Terceros (github, external packages)
    "github.com/sirupsen/logrus"
    "github.com/user/mypackage"

    // Espacios en blanco entre terceros y local

    // Local (tu proyecto)
    "myproject/config"
    "myproject/usuarios"
)
```

### Alias de import

```go
import (
    f "fmt"              // Alias f
    "math/rand"
)

func main() {
    f.Println("Usando alias")   // Acceder con alias
}
```

### Blank import (para efectos secundarios)

```go
import (
    _ "github.com/lib/pq"   // Solo ejecuta init()
)
```

Se usa cuando necesitas que un package se inicialice pero no lo usas directamente (ej: drivers de BD).

### ❌ Errores comunes

```go
// ❌ Importar paquete que no usas (ERROR DE COMPILACIÓN)
import "fmt"
func main() {
    // fmt nunca se usa
}

// ✅ Usar lo que importas
import "fmt"
func main() {
    fmt.Println("Hola")
}

// ❌ No importar lo que necesitas
func main() {
    fmt.Println("Hola")    // ERROR: fmt no importado
}

// ✅ Importar necesarios
import "fmt"
func main() {
    fmt.Println("Hola")
}
```

---

## Parte III: Package-Level Variables y Constantes

Después de imports, declaras variables/constantes a nivel de paquete:

```go
package main

import "fmt"

// Variables a nivel de paquete
var (
    NombreApp = "MiApp"
    Versión = "1.0.0"
)

// Constantes
const MaxUsuarios = 1000

// Inicialización (si necesitas lógica)
var base datos.DB

func init() {
    // Se ejecuta ANTES de main()
    base, _ = abrirBaseDatos()
}

func main() {
    // ...
}
```

**Orden de ejecución:**
1. Variables se declaran
2. `func init()` se ejecuta
3. `func main()` se ejecuta

---

## Parte IV: Función main()

### Anatomía de main()

```go
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```

**Características de main():**
- **Sin parámetros** (no puede recibir argumentos en la firma)
- **Sin retorno**
- **Se ejecuta automáticamente** cuando corres el programa
- **Obligatoria en package main**

### Acceder a argumentos de línea de comandos

Para recibir argumentos cuando ejecutas el programa:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // os.Args[0] es el nombre del programa
    // os.Args[1:] son los argumentos
    
    if len(os.Args) < 2 {
        fmt.Println("Uso: programa nombre")
        return
    }

    nombre := os.Args[1]
    fmt.Printf("Hola, %s!\n", nombre)
}
```

**Ejecución:**
```bash
go run main.go Juan
# Output: Hola, Juan!
```

### Función init()

```go
package main

import "fmt"

func init() {
    fmt.Println("Se ejecuta PRIMERO")
}

func main() {
    fmt.Println("Se ejecuta SEGUNDO")
}
```

**Output:**
```
Se ejecuta PRIMERO
Se ejecuta SEGUNDO
```

Usa `init()` para:
- Inicializar variables globales
- Validar configuración
- Conectar a BD
- Registrar handlers

---

## Parte V: Comentarios

### Comentario simple (una línea)

```go
// Este es un comentario de una línea
x := 5  // Comentario al lado del código
```

### Comentario multilínea

```go
/*
Este es un comentario
de múltiples líneas.
Útil para explicaciones largas.
*/
```

### ✅ Documentación (Godoc)

Los comentarios **directamente antes** de funciones/tipos son documentación:

```go
// Sum añade dos números y retorna el resultado.
// 
// Ejemplo:
//     resultado := Sum(5, 3)
//     fmt.Println(resultado)  // 8
func Sum(a, b int) int {
    return a + b
}

// User representa un usuario del sistema.
type User struct {
    ID   int       // ID único
    Name string    // Nombre completo
    Email string   // Email de contacto
}
```

**Genera documentación local:**
```bash
go doc ./...       # Muestra documentación en terminal
godoc -http=:6060 # Servidor web en localhost:6060
```

### ❌ Comentarios malos

```go
// ❌ Obvio del código
x := 5  // Asignar 5 a x

// ❌ Desactualizado
y := 10  // Carga usuarios de la BD (falso, no hace eso)

// ❌ Ruido
// ---------------------
// SECCIÓN IMPORTANTE
// ---------------------

// ✅ Explica POR QUÉ, no QUÉ
maxGoroutines := 1000  // Limitar para evitar exhaustión de recursos
```

---

## Parte VI: Importando Paquetes Locales

### Estructura del Proyecto

```
miproyecto/
├── go.mod           # Manifiesto del módulo
├── main.go          # Punto de entrada
├── config/
│   └── config.go    # package config
└── usuarios/
    └── usuario.go   # package usuarios
```

### go.mod

```
module miproyecto

go 1.21
```

### Importar paquete local

```go
// main.go
package main

import (
    "fmt"
    "miproyecto/config"  // Paquete local
    "miproyecto/usuarios"
)

func main() {
    cfg := config.Cargar()
    fmt.Println(cfg)
}
```

```go
// config/config.go
package config

func Cargar() string {
    return "Configuración cargada"
}
```

**Importar siempre con el nombre del **módulo** del go.mod, no la ruta del filesystem.**

---

## Parte VII: Orden Típico de un Archivo .go

```go
// 1. Package declaration
package main

// 2. Importaciones (agrupadas, alfabéticas)
import (
    "fmt"
    "strings"
    
    "github.com/alguna/lib"
    
    "myapp/config"
)

// 3. Constantes
const (
    MaxUsuarios = 1000
    DefaultPort = 8080
)

// 4. Variables a nivel de paquete
var (
    logger Logger
    db     *sql.DB
)

// 5. Tipos y estructuras
type Usuario struct {
    ID    int
    Name  string
}

// 6. init() si necesitas inicialización
func init() {
    // Cargar BD, configuración, etc
}

// 7. Funciones públicas (PascalCase)
func CrearUsuario(name string) *Usuario {
    return &Usuario{Name: name}
}

// 8. Funciones privadas (camelCase)
func validarNombre(name string) bool {
    return len(name) > 0
}

// 9. main() (si es package main)
func main() {
    // Punto de entrada
}
```

---

## Parte VIII: Ejecutar un Programa Go

### Con go run (desarrollo)

```bash
go run main.go              # Ejecutar directamente
go run main.go Juan         # Con argumentos
```

### Con go build (producción)

```bash
go build -o miapp main.go   # Compilar a binario
./miapp                     # Ejecutar binario
./miapp Juan                # Con argumentos
```

### Con go install (global)

```bash
go install .                # Instala en $GOPATH/bin
miapp                       # Ejecutar desde cualquier lado
```

---

## Parte IX: Ejemplo Completo - Calculadora Simple

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

const (
    errorArgumentos = "Uso: calculadora número1 operador número2"
)

// Realiza operación aritmética
func Operar(a float64, op string, b float64) (float64, error) {
    switch op {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, fmt.Errorf("división por cero")
        }
        return a / b, nil
    default:
        return 0, fmt.Errorf("operador desconocido: %s", op)
    }
}

func main() {
    // Validar argumentos
    if len(os.Args) != 4 {
        fmt.Println(errorArgumentos)
        return
    }

    // Parsear argumentos
    a, err1 := strconv.ParseFloat(os.Args[1], 64)
    op := os.Args[2]
    b, err2 := strconv.ParseFloat(os.Args[3], 64)

    if err1 != nil || err2 != nil {
        fmt.Println("Error: argumentos no son números")
        return
    }

    // Usar función
    resultado, err := Operar(a, op, b)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Mostrar resultado
    fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, resultado)
}
```

**Ejecución:**
```bash
go run main.go 10 + 5
# Output: 10.00 + 5.00 = 15.00

go run main.go 20 / 4
# Output: 20.00 / 4.00 = 5.00

go run main.go 10 / 0
# Output: Error: división por cero
```

---

## Parte X: Anti-patrones

### ❌ Importar sin usar

```go
import (
    "fmt"
    _ "unused"    // ERROR en compilación
)
```

### ❌ main sin package main

```go
// ❌ Esto no compila
package usuarios

func main() {
    // Main solo en package main
}
```

### ❌ Circular imports

```go
// ❌ a.go
package a
import "b"

// ❌ b.go
package b
import "a"   // ERROR: import circular
```

### ❌ Comentarios obsoletos

```go
// ❌ Comentario desactuali zado
// Conectar a BD MySQL
// (pero ahora usamos PostgreSQL)
db := conectarPostgres()
```

---

## Resumen: Estructura de un .go

1. **Package** - Declara a qué paquete pertenece
2. **Imports** - Qué dependendencias necesita
3. **Constantes/Variables** - Datos a nivel de paquete
4. **Tipos** - Structs, interfaces
5. **init()** - Inicialización (opcional)
6. **Funciones** - Lógica del programa
7. **main()** - Punto de entrada (si es package main)

**Siguiendo este patrón, todo programa Go es predecible y fácil de leer.**
