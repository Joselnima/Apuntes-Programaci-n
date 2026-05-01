# Módulo 04 - Control de flujo

El control de flujo permite decidir qué parte del código se ejecuta según condiciones y repetir acciones con bucles.

## Condicional `if`

```javascript
let edad = 18;

if (edad >= 18) {
  console.log('Eres mayor de edad');
}
```

## `if...else`

```javascript
let clima = 'lluvia';

if (clima === 'sol') {
  console.log('Haz ejercicio al aire libre');
} else {
  console.log('Lleva un paraguas');
}
```

## `else if`

```javascript
let nota = 7;

if (nota >= 9) {
  console.log('Excelente');
} else if (nota >= 7) {
  console.log('Bien');
} else {
  console.log('Necesitas mejorar');
}
```

## `switch`

```javascript
let dia = 'martes';

switch (dia) {
  case 'lunes':
    console.log('Primer día de la semana');
    break;
  case 'martes':
    console.log('Segundo día');
    break;
  default:
    console.log('No es lunes ni martes');
}
```

### El patrón `switch(true)`
Es una técnica avanzada que permite usar rangos o condiciones complejas dentro de un `switch`, algo que normalmente requeriría múltiples `if...else`.

```javascript
const puntuacion = 85;

switch (true) {
  case puntuacion >= 90:
    console.log("A - Sobresaliente");
    break;
  case puntuacion >= 80:
    console.log("B - Notable"); // Entrará aquí
    break;
  case puntuacion >= 70:
    console.log("C - Aprobado");
    break;
  default:
    console.log("F - Suspendido");
}
```

## Operador Ternario (Alternativa a `if...else`)

El operador ternario permite escribir un `if...else` en una sola línea. Es muy útil para asignar valores basados en una condición.

Su sintaxis es: `condición ? valor_si_verdadero : valor_si_falso`

```javascript
let edad = 20;

// Usando if...else clásico
let mensaje;
if (edad >= 18) {
  mensaje = 'Mayor de edad';
} else {
  mensaje = 'Menor de edad';
}

// Usando operador ternario (mucho más limpio)
let mensajeTernario = edad >= 18 ? 'Mayor de edad' : 'Menor de edad';

console.log(mensajeTernario); // "Mayor de edad"
```

## Mapeo con Objetos (Alternativa a `switch` o múltiples `if`)

A menudo, usar un objeto como diccionario (o mapa) es más eficiente y legible que usar un `switch` enorme o encadenar muchos `else if`.

```javascript
const rol = 'admin';

// Alternativa con Mapeo (Diccionario)
const permisos = {
  admin: 'Control total',
  editor: 'Puede editar artículos',
  viewer: 'Solo puede leer',
  default: 'Rol no válido'
};

// Si 'rol' existe en el objeto 'permisos', lo devuelve. Si no, devuelve el 'default'.
const permisoActual = permisos[rol] || permisos['default'];

console.log(permisoActual); // "Control total"
```
Esto evita tener que escribir bloques de `case` y `break`, y hace que añadir nuevos roles sea tan simple como agregar una línea al objeto.

## Bucles `for`

```javascript
for (let i = 0; i < 5; i++) {
  console.log(i);
}
```

## Bucle `while`

```javascript
let contador = 0;
while (contador < 5) {
  console.log(contador);
  contador++;
}
```

## Bucle `do...while`

```javascript
let numero = 0;
do {
  console.log(numero);
  numero++;
} while (numero < 5);
```

## `break` y `continue`

```javascript
for (let i = 0; i < 10; i++) {
  if (i === 5) break;
  if (i % 2 === 0) continue;
  console.log(i);
}
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Control de flujo</title>
</head>
<body>
  <h1>Control de flujo</h1>
  <script>
    const numero = 7;

    if (numero % 2 === 0) {
      console.log('El número es par');
    } else {
      console.log('El número es impar');
    }

    for (let i = 1; i <= 5; i++) {
      console.log(`Iteración ${i}`);
    }
  </script>
</body>
</html>
```

## Ejercicio

1. Haz un condicional que muestre si un número es positivo, negativo o cero.
2. Recorre un array con un bucle `for` y muestra cada elemento.
3. Usa `switch` para mostrar el nombre del día según un número del 1 al 7.

## Resumen

Las estructuras de control permiten tomar decisiones y repetir acciones. Son indispensables para programar cualquier lógica.
