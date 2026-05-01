# Módulo 02 - Variables y tipos de datos

Las variables son lugares donde guardamos información. En JavaScript, las variables almacenan datos como números, texto o valores verdaderos/falsos.

## Declarar variables

### `let`

```javascript
let nombre = 'María';
let edad = 25;
```

### `const`

```javascript
const PI = 3.1416;
```

### `var` (ya no se usa tanto)

```javascript
var saludo = 'Hola';
```

**Recomendación:** usa `let` y `const` en lugar de `var`.

## Tipos de datos Primitivos y sus Límites de Memoria

JavaScript tiene un comportamiento especial a nivel de memoria que debes comprender.

### Number (Float64 por defecto)

En JavaScript clásico, **todos** los números son del tipo `Number`. Bajo el capó, no existe diferencia técnica entre un número entero (`10`) y un decimal (`3.14`). Ambos se guardan utilizando el formato de coma flotante de doble precisión de 64 bits (**Float64** o *IEEE 754*).

```javascript
let numero = 10;
let decimal = 3.14;
```

> **Limitación:** El formato Float64 puede almacenar con precisión absoluta enteros hasta `9007199254740991` (`Number.MAX_SAFE_INTEGER`). Más allá de eso, se pierde precisión.

### BigInt (Enteros ultra grandes)

A partir de ES11, se introdujo `BigInt` para poder manejar números enteros mayores a los 64 bits clásicos sin perder precisión. Se denotan con una `n` al final.

```javascript
let muyGrande = 9007199254740991000000n;
```

### Cadena de texto (String)

```javascript
let palabra = 'JavaScript';
let frase = "Estoy aprendiendo JS";
```

### Boolean

```javascript
let esVerdadero = true;
let esFalso = false;
```

### Symbol

Un tipo de dato inmutable y único. Muy usado en el desarrollo avanzado para crear claves de objetos que jamás colisionarán con otras.

```javascript
const idUnico = Symbol('id');
```

### Null y Undefined

```javascript
let valorNulo = null; // Ausencia intencional de valor
let valorIndefinido;  // Declarado, pero sin asignar
```

## Arreglos Tipados (Control estricto de memoria: 8, 16, 32, 64 bits)

Con la llegada de WebGL para juegos 3D y WebAssembly, a veces JS necesita manejar la memoria RAM con extrema precisión, justo como lo hacen C o C++. Para ello existen los **Typed Arrays**. 

Estos arreglos obligan a almacenar únicamente el tipo de dato especificado en cada "casilla" de memoria:

- **`Int8Array`**: Enteros de 8 bits con signo (-128 a 127).
- **`Uint8Array`**: Enteros de 8 bits sin signo (0 a 255). Se usa para manejar Bytes de imágenes o archivos.
- **`Int16Array`**: Enteros de 16 bits.
- **`Int32Array`**: Enteros de 32 bits.
- **`Float32Array`**: Decimales de 32 bits (Comunes en cálculos gráficos de tarjeta de video).
- **`Float64Array`**: Decimales de 64 bits (idénticos al tipo Number por defecto).

```javascript
// Este array solo permite números del -128 al 127
const pixeles = new Int8Array(3);
pixeles[0] = 100;
pixeles[1] = -50;
pixeles[2] = 200; // ¡Cuidado! 200 no cabe en 8 bits con signo, se transformará en negativo.
```

## Tipos de Referencia (Objetos Estructurales)

A diferencia de los tipos primitivos que guardan un valor directo en memoria, los tipos de referencia apuntan a un espacio en memoria. JavaScript cuenta con varios objetos estructurales nativos.

### Objetos y Arrays Genéricos
```javascript
let persona = { nombre: 'Ana', edad: 30 };
let lista = [1, 2, 3];
```

### Date (Fechas)
Se utiliza para almacenar y manipular fechas, horas y temporizadores.
```javascript
let hoy = new Date(); // Guarda la fecha y hora actuales
```

### RegExp (Expresiones Regulares)
Estructura nativa para trabajar con patrones de búsqueda en cadenas de texto.
```javascript
let patron = /[A-Z]/g;
```

### Funciones
¡En JavaScript, las funciones también son un tipo de dato! Pueden guardarse en variables, pasarse como argumentos o retornarse desde otras funciones.
```javascript
let saludar = function() { console.log("Hola"); };
```

### Otros tipos avanzados
Existen muchas otras colecciones estructurales que estudiarás más adelante:
- **Map y Set**: Colecciones avanzadas para pares clave-valor y conjuntos únicos.
- **Error**: Objetos para el rastreo y manejo de excepciones.
- **Promise**: Estructuras nativas para controlar operaciones asíncronas futuras.

## Operaciones con variables

```javascript
let a = 5;
let b = 2;
let suma = a + b; // 7
let texto = 'Hola ' + 'mundo';
```

## Tipos dinámicos

JavaScript cambia el tipo de una variable según su valor.

```javascript
let valor = 'Hola';
valor = 10; // ahora es número
```

## Comprobando el tipo

```javascript
console.log(typeof 'Hola'); // string
console.log(typeof 123);    // number
console.log(typeof true);   // boolean
console.log(typeof null);   // object (es una excepción histórica)
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Variables y tipos</title>
</head>
<body>
  <h1>Variables y tipos</h1>
  <script>
    const nombre = 'Carlos';
    let edad = 28;
    let esEstudiante = true;

    console.log('Nombre:', nombre);
    console.log('Edad:', edad);
    console.log('¿Es estudiante?', esEstudiante);
  </script>
</body>
</html>
```

## Buenas prácticas

- Usa `const` cuando el valor no cambiará.
- Usa `let` para valores que sí cambian.
- El nombre de la variable debe describir lo que almacena.
- Evita nombres genéricos como `a`, `x` o `valor`.

## Ejercicio

1. Declara una variable `ciudad` con tu ciudad favorita.
2. Declara una variable `anio` con el año actual.
3. Declara una variable `tieneMascota` con valor `true` o `false`.
4. Muestra esos valores en la consola.

## Resumen

Las variables son el primer paso para escribir programas. Conocer los tipos de datos te ayuda a trabajar mejor con JavaScript.
