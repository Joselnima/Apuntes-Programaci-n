# Módulo 26 - Expresiones Regulares (RegEx)

Las expresiones regulares son patrones utilizados para encontrar, validar o reemplazar combinaciones de caracteres dentro de cadenas de texto. Son universales en casi todos los lenguajes de programación.

## Creación de una RegEx

En JavaScript existen dos formas de crear una Expresión Regular:

```javascript
// 1. Sintaxis literal (Recomendada y más rápida)
const regex1 = /hola/i;

// 2. Usando el constructor RegExp (Útil si el patrón es una variable dinámica)
const patronStr = "hola";
const regex2 = new RegExp(patronStr, "i");
```

Las letras al final (`i`, `g`, etc.) son **banderas (flags)**:
- `i` (Insensitive): Ignora mayúsculas y minúsculas.
- `g` (Global): Busca todas las coincidencias, no solo la primera.
- `m` (Multiline): Trata el texto como múltiples líneas.

## Métodos de RegEx

El objeto RegExp tiene métodos propios para evaluar strings:

```javascript
const regexEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

// .test(): Devuelve true o false si el string cumple el patrón
console.log(regexEmail.test("usuario@email.com")); // true
console.log(regexEmail.test("usuario.com"));       // false

// .exec(): Devuelve un array con detalles de la coincidencia, o null si no hay
const extraeNumero = /\d+/; // \d busca dígitos
console.log(extraeNumero.exec("El precio es 50 dolares")); // ["50", index: 13, input: ...]
```

## Métodos de String que usan RegEx

```javascript
const frase = "El gato, el perro y el GATO se llevan bien";

// .match(): Devuelve un array con las coincidencias
console.log(frase.match(/gato/gi)); // ["gato", "GATO"]

// .replace() / .replaceAll()
console.log(frase.replace(/gato/gi, "pájaro")); // "El pájaro, el perro y el pájaro se llevan bien"

// .search(): Similar a indexOf pero con RegEx (devuelve la posición)
console.log(frase.search(/perro/)); // 12
```

## Sintaxis Básica de Patrones

- `\d` : Busca cualquier dígito (0-9).
- `\w` : Busca caracteres alfanuméricos (letras, números, guiones bajos).
- `\s` : Busca espacios en blanco (espacio, tab, salto de línea).
- `.` : Busca CUALQUIER carácter (excepto saltos de línea).
- `^` : Inicio del string (ej. `/^Hola/` solo coincide si "Hola" está al principio).
- `$` : Fin del string (ej. `/Mundo$/` solo coincide si termina en "Mundo").
- `[abc]` : Busca 'a', 'b' o 'c'.
- `+` : Una o más repeticiones (ej. `\d+` = "5", "50", "500").
- `*` : Cero o más repeticiones.
- `?` : Cero o una repetición (Opcional).

### Ejemplo práctico: Validar contraseña
```javascript
// La contraseña debe tener: Al menos 8 caracteres, 1 mayúscula, 1 número
const regexClave = /^(?=.*[A-Z])(?=.*\d).{8,}$/;
console.log(regexClave.test("hola123")); // false (sin mayúscula)
console.log(regexClave.test("Hola1234")); // true
```
