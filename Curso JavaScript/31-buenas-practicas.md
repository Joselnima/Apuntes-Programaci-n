# Módulo 15 - Buenas prácticas y estilo

Escribir código claro y ordenado hace que tus programas sean más fáciles de entender y mantener.

## Nombres descriptivos

```javascript
let totalProductos = 10;
const nombreUsuario = 'Ana';
```

- Prefiere `totalProductos` a `tp`.
- Usa nombres que expliquen el propósito.

## Evita comentarios innecesarios

Un código bien escrito no necesita demasiados comentarios.

```javascript
// Malo
let x = 5; // asigna 5 a x

// Bueno
let contador = 5;
```

## Usa funciones pequeñas

Cada función debe hacer una sola tarea.

```javascript
function validarEmail(email) {
  return email.includes('@');
}
```

## Evita la repetición (DRY)

No repitas código. Si necesitas hacer lo mismo varias veces, ponlo en una función.

## Organización clara

- Separa lógica, datos y presentación.
- Usa módulos cuando el proyecto crece.
- Mantén archivos cortos y coherentes.

## Formato y estilo

- Usa indentación de 2 o 4 espacios.
- Añade espacios alrededor de operadores.
- Usa `camelCase` para variables y funciones.

## Comentarios útiles

Comenta el por qué, no el qué.

```javascript
// Calcula el precio final con IVA
function calcularPrecioConIVA(precio) {
  return precio * 1.21;
}
```

## Versionado y control de cambios

Usa Git para guardar tu trabajo paso a paso. Haz commits claros.

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Buenas prácticas</title>
</head>
<body>
  <h1>Buenas prácticas</h1>
  <script>
    function esNumeroPar(numero) {
      return numero % 2 === 0;
    }

    const numeros = [1, 2, 3, 4];
    const pares = numeros.filter(esNumeroPar);

    console.log('Números pares:', pares);
  </script>
</body>
</html>
```

## Ejercicio

1. Reescribe un código antiguo usando nombres claros.
2. Extrae una parte repetida en una función.
3. Decide cuándo usar `const` y cuándo usar `let`.

## Resumen

La buena práctica no es opcional. Un código limpio te ahorra tiempo hoy y mañana.
