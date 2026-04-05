# Tipo de dato: string

## Definición
En Go, `string` es una secuencia inmutable de bytes codificados en UTF-8. Cada `string` representa texto y puede contener caracteres ASCII, Unicode y símbolos.

## Para qué sirve
- Representar texto legible: nombres, mensajes, etiquetas, rutas y contenido de archivos.
- Guardar datos que no deben modificarse directamente.
- Trabajar con protocolos, URLs, JSON y comunicación entre sistemas.

## Cuándo se usa
- Cuando necesitas almacenar o mostrar texto.
- En entradas y salidas de consola.
- En concatenación de mensajes.
- Al procesar cadenas con funciones de la biblioteca estándar como `len`, `strings.Contains`, `strings.Split`, `strings.TrimSpace`, etc.

## Características importantes
- Los `string` son inmutables: no se puede cambiar un carácter dentro de una cadena existente.
- Acceder a un `string` con índice devuelve un `byte`, no un `rune`.
- Para manejar caracteres Unicode correctamente, se usan `rune` y `for range`.

## Funciones útiles
- `len(s)` : devuelve la longitud en bytes de `s`.
- `s1 + s2` : concatena cadenas.
- `strings.Contains(s, substr)` : verifica si `substr` está en `s`.
- `strings.HasPrefix(s, prefijo)` : comprueba el prefijo.
- `strings.HasSuffix(s, sufijo)` : comprueba el sufijo.
- `strings.Index(s, substr)` : busca la primera posición de `substr`.
- `strings.LastIndex(s, substr)` : busca la última posición de `substr`.
- `strings.Count(s, substr)` : cuenta ocurrencias de `substr`.
- `strings.Split(s, sep)` : divide la cadena por un separador.
- `strings.SplitN(s, sep, n)` : divide la cadena en hasta `n` partes.
- `strings.Fields(s)` : divide la cadena por espacios y saltos de línea.
- `strings.FieldsFunc(s, f)` : divide según una función personalizada.
- `strings.Join(parts, sep)` : une partes con un separador.
- `strings.TrimSpace(s)` : elimina espacios en los extremos.
- `strings.Trim(s, cutset)` : elimina caracteres en los extremos.
- `strings.TrimLeft(s, cutset)` : elimina caracteres al inicio.
- `strings.TrimRight(s, cutset)` : elimina caracteres al final.
- `strings.TrimPrefix(s, prefijo)` : elimina un prefijo.
- `strings.TrimSuffix(s, sufijo)` : elimina un sufijo.
- `strings.Replace(s, old, new, n)` : reemplaza subcadenas.
- `strings.ReplaceAll(s, old, new)` : reemplaza todas las ocurrencias.
- `strings.Repeat(s, n)` : repite la cadena `n` veces.
- `strings.ToUpper(s)` : convierte a mayúsculas.
- `strings.ToLower(s)` : convierte a minúsculas.
- `strings.Title(s)` : convierte la primera letra de cada palabra en mayúscula (deprecated en Go 1.20, usa `cases.Title` para Unicode completo).
- `strings.EqualFold(s1, s2)` : compara ignorando mayúsculas/minúsculas.
- `strings.Map(f, s)` : transforma cada carácter con una función.
- `strings.NewReplacer(old1, new1, old2, new2, ...)` : reemplaza múltiples pares.
- `strings.Cut(s, sep)` : divide en dos partes antes y después del separador.
- `strings.ContainsAny(s, chars)` : comprueba si alguno de `chars` aparece en `s`.
- `strings.SplitAfter(s, sep)` : divide la cadena incluyendo el separador.
- `strings.SplitAfterN(s, sep, n)` : divide en hasta `n` partes incluyendo el separador.
- `strings.TrimFunc(s, f)` : elimina caracteres en los extremos según `f`.
- `strings.TrimLeftFunc(s, f)` : elimina caracteres al inicio según `f`.
- `strings.TrimRightFunc(s, f)` : elimina caracteres al final según `f`.
- `fmt.Sprintf(format, args...)` : crea cadenas formateadas.
- `strconv.Itoa(n)` : convierte entero a string.
- `strconv.Atoi(s)` : convierte string a entero.

## Ejemplos de funciones
```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    texto := "  Hola, Mundo! "

    // Longitud en bytes del string
    fmt.Println("len:", len(texto))

    // Eliminar espacios en los extremos
    fmt.Println("TrimSpace:", strings.TrimSpace(texto))

    // Búsquedas simples
    fmt.Println("Contains 'Mundo':", strings.Contains(texto, "Mundo"))
    fmt.Println("HasPrefix '  Ho':", strings.HasPrefix(texto, "  Ho"))
    fmt.Println("HasSuffix '! ':", strings.HasSuffix(texto, "! "))

    indice := strings.Index(texto, "Mundo")
    fmt.Println("Index de 'Mundo':", indice)
    fmt.Println("LastIndex de 'o':", strings.LastIndex(texto, "o"))
    fmt.Println("Count de 'o':", strings.Count(texto, "o"))

    partes := strings.Split(strings.TrimSpace(texto), ", ")
    fmt.Println("Split:", partes)
    fmt.Println("Fields:", strings.Fields("uno dos   tres"))

    frase := strings.Join([]string{"Go", "es", "genial"}, " ")
    fmt.Println("Join:", frase)

    reemplazo := strings.ReplaceAll(frase, "genial", "rápido")
    fmt.Println("ReplaceAll:", reemplazo)

    mayus := strings.ToUpper(frase)
    minus := strings.ToLower(frase)
    fmt.Println("ToUpper:", mayus)
    fmt.Println("ToLower:", minus)
    fmt.Println("Repeat:", strings.Repeat("Go! ", 2))

    textoCorto := "Go"
    fmt.Println("EqualFold 'go' == 'Go':", strings.EqualFold(textoCorto, "go"))

    textoMap := strings.Map(func(r rune) rune {
        if r >= 'a' && r <= 'z' {
            return r - 32
        }
        return r
    }, "abc")
    fmt.Println("Map:", textoMap)

    parte1, parte2, ok := strings.Cut("nombre=Ana", "=")
    fmt.Println("Cut:", parte1, parte2, ok)

    fmt.Println("ContainsAny 'aeiou':", strings.ContainsAny(texto, "aeiou"))
    fmt.Println("SplitAfter:", strings.SplitAfter("a,b,c", ","))
    fmt.Println("SplitAfterN:", strings.SplitAfterN("a,b,c", ",", 2))

    trimmedCustom := strings.TrimFunc("...Hola...", func(r rune) bool {
        return r == '.'
    })
    fmt.Println("TrimFunc:", trimmedCustom)

    replacer := strings.NewReplacer("Hola", "Adiós", "Mundo", "Gophers")
    fmt.Println("NewReplacer:", replacer.Replace("Hola, Mundo!"))

    textoNum := "42"
    numero, _ := strconv.Atoi(textoNum)
    fmt.Println("Atoi:", numero)
    fmt.Println("Itoa:", strconv.Itoa(numero))

    nombre := "Ana"
    edad := 28
    mensaje := fmt.Sprintf("Hola %s, tienes %d años", nombre, edad)
    fmt.Println("Sprintf:", mensaje)
}
```

## Ejemplo
```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    saludo := "Hola, mundo"
    nombre := "Ana"

    mensaje := saludo + ", " + nombre + "!"
    fmt.Println(mensaje)
    fmt.Println("Longitud en bytes:", len(mensaje))
    fmt.Println("Contiene 'mundo'?:", strings.Contains(mensaje, "mundo"))

    partes := strings.Split(mensaje, ", ")
    fmt.Println(partes)
}
```

## Buenas prácticas
- Usa `string` para texto y `[]byte` para datos binarios.
- Evita modificar cadenas directamente; crea nuevas cadenas si necesitas un resultado distinto.
- Para iterar por caracteres Unicode, usa `for range`.
- Convierte `string` a `[]byte` cuando trabajes con I/O o cifrado.
