# Módulo 05 - Operadores y expresiones

Los operadores son símbolos que permiten realizar cálculos, comparar valores y combinar expresiones.

## Operadores aritméticos

```javascript
let suma = 5 + 3;       // 8
let resta = 5 - 3;      // 2
let multiplicacion = 5 * 3; // 15
let division = 10 / 2;  // 5
let modulo = 10 % 3;    // 1
let potencia = 2 ** 3;  // 8
```

## Operadores de asignación

```javascript
let x = 10;
x += 5;  // x = x + 5 => 15
x -= 2;  // x = x - 2 => 13
x *= 2;  // x = x * 2 => 26
x /= 2;  // x = x / 2 => 13
```

## Operadores de comparación

```javascript
console.log(5 > 3);   // true
console.log(5 < 3);   // false
console.log(5 >= 5);  // true
console.log(5 <= 4);  // false
console.log(5 === '5'); // false
console.log(5 == '5');  // true
console.log(5 !== '5'); // true
```

### Diferencia entre `==` y `===`

- `==` compara valores después de convertirlos.
- `===` compara valor y tipo sin conversión.

**Usa siempre `===`** para evitar sorpresas.

## Operadores lógicos y Cortocircuito (Short-circuiting)

Los operadores lógicos no solo devuelven booleanos. En JavaScript, devuelven el valor de uno de los operandos originales.

```javascript
// Cortocircuito con OR (||)
// Devuelve el primer valor "truthy" que encuentre, o el último si todos son "falsy".
let nombreUsuario = "" || "Invitado"; // "Invitado"
let configuracion = "Oscuro" || "Claro"; // "Oscuro"

// Cortocircuito con AND (&&)
// Devuelve el primer valor "falsy", o el último si todos son "truthy".
let estaLogueado = true;
estaLogueado && console.log("Bienvenido de vuelta!"); // Se ejecuta el console.log
```

## Operadores Unarios

Son operadores que actúan sobre un solo operando.

```javascript
// typeof: Devuelve el tipo de dato como string
console.log(typeof 42); // "number"
console.log(typeof "Hola"); // "string"
console.log(typeof true); // "boolean"

// Unario de negación lógica (!)
console.log(!true); // false
console.log(!!"Hola"); // true (Doble negación convierte a booleano real)

// Incremento y decremento
let contador = 5;
contador++; // Sube a 6
contador--; // Baja a 5
```

## Operadores Relacionales (Instancias y Propiedades)

```javascript
// in: Comprueba si una propiedad existe en un objeto
const auto = { marca: "Ford" };
console.log("marca" in auto); // true

// instanceof: Comprueba si un objeto fue creado por una clase/constructor
const fecha = new Date();
console.log(fecha instanceof Date); // true
```

## Operador ternario

```javascript
let edad = 18;
let mensaje = edad >= 18 ? 'Eres mayor de edad' : 'Eres menor de edad';
```

## Concatenación de strings

```javascript
let nombre = 'Ana';
let saludo = 'Hola ' + nombre + '!';
```

## Template literals (plantillas de texto)

```javascript
let saludo2 = `Hola ${nombre}, tienes ${edad} años.`;
```

## Expresiones

Una expresión es cualquier fragmento de código que produce un valor.

```javascript
let resultado = (5 + 3) * 2;
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Operadores</title>
</head>
<body>
  <h1>Operadores y expresiones</h1>
  <script>
    let precio = 100;
    let descuento = 20;
    let total = precio - descuento;

    let mensaje = `Precio original: ${precio} €\nDescuento: ${descuento} €\nTotal: ${total} €`;
    console.log(mensaje);
  </script>
</body>
</html>
```

## Ejercicio

1. Calcula el área de un rectángulo con base 7 y altura 4.
2. Usa el operador ternario para mostrar si un número es par o impar.
3. Crea un mensaje con template literals que incluya tu nombre y tu edad.

## Resumen

Los operadores permiten hacer cálculos, comparar valores y tomar decisiones. Juntar expresiones es la base de cualquier programa.
