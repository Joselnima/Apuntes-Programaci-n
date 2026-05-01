# Módulo 05 - Funciones

Las funciones son bloques de código que realizan una tarea. Puedes llamarlas varias veces y pasarles datos.

## Función básica

```javascript
function saludar() {
  console.log('Hola');
}

saludar();
```

## Función con parámetros

```javascript
function saludar(nombre) {
  console.log(`Hola ${nombre}`);
}

saludar('Lucía');
```

## Función con valor de retorno

```javascript
function sumar(a, b) {
  return a + b;
}

let resultado = sumar(3, 4);
console.log(resultado); // 7
```

## Función anónima

```javascript
const restar = function(a, b) {
  return a - b;
};

console.log(restar(10, 3));
```

## Arrow functions

```javascript
const multiplicar = (a, b) => a * b;
console.log(multiplicar(4, 5));
```

## Función con varios parámetros y valores por defecto

```javascript
function crearMensaje(nombre = 'amigo', mensaje = 'Hola') {
  return `${mensaje}, ${nombre}`;
}

console.log(crearMensaje('Marta'));
```

## Funciones de Orden Superior y Callbacks

En JavaScript, las funciones son "ciudadanos de primera clase". Pueden guardarse en variables y pasarse como argumentos a otras funciones (eso es un callback).

```javascript
function procesarDato(dato, callback) {
  let modificado = dato + 10;
  callback(modificado);
}

procesarDato(5, function(resultado) {
  console.log("El resultado es:", resultado); // 15
});
```

## Closures (Clausuras)

Un closure ocurre cuando una función interna "recuerda" y tiene acceso a las variables de la función externa, incluso después de que la externa haya terminado.

```javascript
function crearContador() {
  let cuenta = 0; // Esta variable está "protegida"
  
  return function() {
    cuenta++;
    return cuenta;
  };
}

const contador = crearContador();
console.log(contador()); // 1
console.log(contador()); // 2
```

## IIFE (Expresión de Función Invocada Inmediatamente)

Una función que se ejecuta en cuanto se define. Muy usada históricamente para evitar contaminar el scope global.

```javascript
(function() {
  const variablePrivada = "Secreto";
  console.log("Me ejecuto inmediatamente!");
})();
```

## El Contexto (`this`) y `call`, `apply`, `bind`

Cada función normal tiene su propio contexto `this`. Puedes alterarlo forzosamente usando:

- `call()`: Llama a la función pasándole el objeto que será `this` y parámetros sueltos.
- `apply()`: Igual, pero los parámetros se pasan como array.
- `bind()`: Devuelve una copia de la función con el `this` atado para siempre.

```javascript
function presentarse(saludo) {
  console.log(`${saludo}, soy ${this.nombre}`);
}

const persona = { nombre: 'Carlos' };

presentarse.call(persona, 'Hola');  // "Hola, soy Carlos"
presentarse.apply(persona, ['Hey']); // "Hey, soy Carlos"

const funcionAtada = presentarse.bind(persona);
funcionAtada('Buenas'); // "Buenas, soy Carlos"
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Funciones</title>
</head>
<body>
  <h1>Funciones en JavaScript</h1>
  <script>
    function calcularIVA(precio) {
      const IVA = 0.21;
      return precio + precio * IVA;
    }

    const precioFinal = calcularIVA(50);
    console.log(`Precio final con IVA: ${precioFinal} €`);
  </script>
</body>
</html>
```

## Ejercicio

1. Crea una función que calcule el área de un círculo.
2. Crea una función que reciba un nombre y devuelva un saludo personalizado.
3. Crea una arrow function que calcule el doble de un número.

## Resumen

Las funciones te permiten reutilizar código y organizar mejor tus programas. Cada función debe hacer una sola cosa.
