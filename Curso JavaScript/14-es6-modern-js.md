# Módulo 10 - JavaScript moderno (ES6+)

ES6 y versiones posteriores introdujeron muchas mejoras que hacen el código más claro y poderoso.

## `let` y `const`

- `let` para valores cambiantes.
- `const` para valores que no cambian.

## Arrow functions

```javascript
const sumar = (a, b) => a + b;
```

## Template literals

```javascript
const nombre = 'Luis';
const saludo = `Hola ${nombre}, bienvenido.`;
```

## Destructuring

```javascript
const persona = { nombre: 'Ana', edad: 25 };
const { nombre, edad } = persona;

const numeros = [1, 2, 3];
const [primero, segundo] = numeros;
```

## Spread y rest

```javascript
const valores = [1, 2, 3];
const copia = [...valores];

function sumar(...numeros) {
  return numeros.reduce((total, n) => total + n, 0);
}
```

## Parámetros por defecto

```javascript
function crearSaludo(nombre = 'amigo') {
  return `Hola ${nombre}`;
}
```

## Objetos mejorados

```javascript
const edad = 30;
const usuario = {
  nombre: 'Luis',
  edad,
  saludar() {
    console.log('Hola');
  }
};
```

## `const` con arrays y objetos

```javascript
const lista = [1, 2, 3];
lista.push(4); // funciona
```

## Novedades de ES7 (ECMAScript 2016)

Aunque ES7 fue una actualización pequeña, introdujo dos características muy útiles:

### `Array.prototype.includes()`
Antes usábamos `indexOf()` para saber si un elemento estaba en un array. Ahora es mucho más directo:
```javascript
const frutas = ['manzana', 'pera', 'uva'];
console.log(frutas.includes('pera')); // true
console.log(frutas.includes('sandía')); // false
```

### Operador de Exponenciación (`**`)
Una alternativa más corta a `Math.pow()`:
```javascript
console.log(2 ** 3); // 8 (2 elevado a 3)
console.log(5 ** 2); // 25
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>ES6</title>
</head>
<body>
  <h1>JavaScript moderno</h1>
  <script>
    const nombre = 'Clara';
    const edad = 27;

    const persona = { nombre, edad };
    const mensaje = `La persona se llama ${persona.nombre} y tiene ${persona.edad} años.`;

    console.log(mensaje);

    const numeros = [1, 2, 3, 4];
    const dobles = numeros.map(n => n * 2);
    console.log(dobles);
  </script>
</body>
</html>
```

## Ejercicio

1. Usa destructuring para extraer propiedades de un objeto.
2. Crea una función con parámetros rest.
3. Combina dos arrays usando spread.

## Resumen

ES6 modernizó JavaScript con sintaxis más simple y herramientas más potentes. Es la base del código moderno.
