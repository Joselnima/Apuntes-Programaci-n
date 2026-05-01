# Módulo 03 - Métodos Útiles de Tipos de Datos Primitivos

JavaScript proporciona cientos de métodos nativos integrados directamente en los tipos de datos primitivos a través de sus correspondientes "Objetos Envoltorio" (Wrapper Objects como `String`, `Number`, `Math`). A continuación, exploraremos de forma exhaustiva sus capacidades.

## 1. Métodos Exahustivos de Strings

Las cadenas de texto tienen métodos para limpiar, buscar, reemplazar y manipular. **Todos estos métodos devuelven un nuevo string**, nunca modifican el original porque los primitivos en JS son inmutables.

```javascript
const texto = "   JavaScript es asombroso!   ";

// --- LIMPIEZA Y FORMATO ---
console.log(texto.trim());      // "JavaScript es asombroso!"
console.log(texto.trimStart()); // "JavaScript es asombroso!   "
console.log(texto.trimEnd());   // "   JavaScript es asombroso!"
console.log(texto.toLowerCase()); // "   javascript es asombroso!   "
console.log(texto.toUpperCase()); // "   JAVASCRIPT ES ASOMBROSO!   "

// padStart() y padEnd() (ES8) - Relleno
console.log("5".padStart(3, "0")); // "005"
console.log("A".padEnd(4, "."));   // "A..."

// --- EXTRACCIÓN ---
const limpio = "JavaScript";
console.log(limpio.slice(0, 4));    // "Java" (Admite índices negativos ej. slice(-6))
console.log(limpio.substring(4, 10)); // "Script" (Similar a slice pero sin negativos)

// Extracción de un solo carácter
console.log(limpio.charAt(0));       // "J"
console.log(limpio.charCodeAt(0));   // 74 (Código numérico UTF-16)
console.log(limpio.codePointAt(0));  // 74 (Código Unicode completo)

// --- REEMPLAZO Y DIVISIÓN ---
const frase = "gato, perro, gato";
console.log(frase.replace("gato", "ave"));    // "ave, perro, gato" (Solo el primero)
console.log(frase.replaceAll("gato", "ave")); // "ave, perro, ave" (Todos)

console.log(frase.split(", ")); // ["gato", "perro", "gato"]
console.log("Hola".repeat(3));  // "HolaHolaHola"
console.log("Hola".concat(" ", "Mundo")); // "Hola Mundo"

// --- BÚSQUEDA ---
console.log(frase.indexOf("perro"));     // 6
console.log(frase.lastIndexOf("gato"));  // 13
console.log(frase.includes("perro"));    // true
console.log(frase.startsWith("gato"));   // true
console.log(frase.endsWith("ave"));      // false

// Búsqueda con Expresiones Regulares (RegEx)
console.log(frase.search(/perro/));      // 6
console.log(frase.match(/gato/g));       // ["gato", "gato"]
// matchAll() devuelve un iterador con información detallada de cada coincidencia
```

## 2. Métodos de Numbers y Floats

El objeto estático `Number` y el prototipo de los números nos dan el control sobre precisión y conversiones.

```javascript
let num = 125.789;

// --- FORMATEO (Devuelven Strings) ---
console.log(num.toFixed(2));      // "125.79" (Fija la cantidad de decimales y redondea)
console.log(num.toPrecision(4));  // "125.8" (Fija la cantidad TOTAL de dígitos)
console.log(num.toExponential()); // "1.25789e+2" (Notación científica)

// Cambio de base (Binario, Hexadecimal, etc.)
let byte = 255;
console.log(byte.toString(2));  // "11111111" (Binario)
console.log(byte.toString(16)); // "ff" (Hexadecimal)

// --- COMPROBACIONES Y CONVERSIÓN (Objeto Number) ---
console.log(Number.isInteger(125));       // true
console.log(Number.isFinite(125));        // true
console.log(Number.isSafeInteger(99999)); // true (Verifica si no pasa el límite de Float64)

// NaN (Not A Number) es un tipo numérico de error
console.log(Number.isNaN(NaN));           // true
console.log(Number.isNaN(0 / 0));         // true

// Parseo estático
console.log(Number.parseInt("10.5"));     // 10
console.log(Number.parseFloat("10.5"));   // 10.5
```

## 3. El Poderoso Objeto `Math`

`Math` contiene un arsenal matemático. No se usa con `new`, todos sus métodos son estáticos.

```javascript
// Constantes
console.log(Math.PI); // 3.141592653589793
console.log(Math.E);  // 2.718281828459045 (Número de Euler)

// Redondeo
console.log(Math.round(4.5)); // 5 (Hacia el más cercano)
console.log(Math.floor(4.9)); // 4 (Hacia abajo siempre)
console.log(Math.ceil(4.1));  // 5 (Hacia arriba siempre)
console.log(Math.trunc(4.9)); // 4 (Corta la parte decimal sin redondear)

// Operaciones
console.log(Math.pow(2, 3));  // 8 (2 al cubo, equivalente a 2 ** 3)
console.log(Math.sqrt(16));   // 4 (Raíz cuadrada)
console.log(Math.cbrt(27));   // 3 (Raíz cúbica)
console.log(Math.abs(-50));   // 50 (Valor absoluto)
console.log(Math.sign(-50));  // -1 (Signo: -1 negativo, 1 positivo, 0 cero, -0)

// Máximos y Mínimos
console.log(Math.max(10, 50, -2, 100)); // 100
console.log(Math.min(10, 50, -2, 100)); // -2

// Aleatoriedad
console.log(Math.random()); // Entre 0.000... y 0.999...
// Truco para random entre 1 y 10:
console.log(Math.floor(Math.random() * 10) + 1); 

// Trigonometría (Trabaja en Radianes)
console.log(Math.sin(0)); // 0
console.log(Math.cos(Math.PI)); // -1
console.log(Math.tan(0)); // 0
// También incluye acos, asin, atan, cosh, sinh, tanh, etc.

// Logaritmos
console.log(Math.log(Math.E)); // 1 (Logaritmo natural)
console.log(Math.log10(100));  // 2 (Logaritmo base 10)
```

## 4. Funciones de Fechas (`Date`)

El objeto `Date` contiene la fecha exacta con milisegundos desde el 1 de Enero de 1970 (Epoch).

```javascript
const hoy = new Date(); 

// --- GETTERS (Lectura) ---
console.log(hoy.getFullYear()); // 2026 (Siempre usar este, no getYear())
console.log(hoy.getMonth());    // Mes (0 = Enero a 11 = Diciembre)
console.log(hoy.getDate());     // Día del mes (1 - 31)
console.log(hoy.getDay());      // Día de la semana (0 = Domingo, 1 = Lunes...)
console.log(hoy.getHours());    // Horas (0 - 23)
console.log(hoy.getMinutes());  // Minutos (0 - 59)
console.log(hoy.getSeconds());  // Segundos (0 - 59)
console.log(hoy.getMilliseconds()); // Milisegundos (0 - 999)

// Tiempo UNIX
console.log(hoy.getTime()); // Milisegundos transcurridos desde 1970
console.log(Date.now());    // Igual que getTime(), pero método estático directo (más rápido)

// --- SETTERS (Modificación) ---
let cita = new Date();
cita.setFullYear(2030);
cita.setMonth(11); // Diciembre
cita.setDate(31);
cita.setHours(23, 59, 59); // Se pueden pasar minutos y segundos al mismo tiempo
console.log(cita); // 31 de Dic de 2030 a las 23:59:59

// --- FORMATEO (Conversión a String) ---
console.log(hoy.toDateString());       // "Fri May 01 2026"
console.log(hoy.toTimeString());       // "10:30:00 GMT-0500"
console.log(hoy.toISOString());        // "2026-05-01T15:30:00.000Z" (Estándar global)
console.log(hoy.toLocaleDateString()); // "1/5/2026" (Formateo según tu país)
console.log(hoy.toLocaleTimeString()); // "10:30:00 AM" (Formato hora según tu país)
```

## 5. Booleans y Datos Binarios

Los booleanos no tienen métodos complejos (`toString()` y `valueOf()`), pero su conversión es vital:

```javascript
// Forzar conversión a boolean
console.log(Boolean("Texto")); // true
console.log(!! "Texto");       // true (Doble negación)

// Falsy values:
// 0, "", null, undefined, false, NaN
```

Los datos pueden manipularse a nivel de bits (1s y 0s):

```javascript
let a = 5; // En binario: 0101
let b = 3; // En binario: 0011

// Operadores Bitwise
console.log(a & b); // 1 (AND: 0001)
console.log(a | b); // 7 (OR:  0111)
console.log(a ^ b); // 6 (XOR: 0110)
console.log(~a);    // -6 (NOT)
console.log(a << 1); // 10 (Desplazamiento izquierda: multiplica por 2)
console.log(a >> 1); // 2 (Desplazamiento derecha: divide por 2 y trunca)
```
