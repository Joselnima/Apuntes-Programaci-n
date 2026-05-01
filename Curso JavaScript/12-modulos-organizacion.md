# Módulo 13 - Módulos y organización

Organizar tu código en archivos y módulos hace que tus proyectos sean más claros y fáciles de mantener.

## ¿Por qué dividir el código?

- Facilita la lectura.
- Reduce errores.
- Permite reutilizar funciones.
- Mejora el trabajo en equipo.

## Módulos en JavaScript moderno

Un módulo es un archivo que exporta funciones, objetos o valores, y otro archivo puede importarlos.

### Exportar

```javascript
// saludar.js
export function saludar(nombre) {
  return `Hola ${nombre}`;
}
```

### Importar

```javascript
// app.js
import { saludar } from './saludar.js';
console.log(saludar('Ana'));
```

## Export default

```javascript
// calculadora.js
export default function sumar(a, b) {
  return a + b;
}

// app.js
import sumar from './calculadora.js';
```

## Importaciones agrupadas

```javascript
// utilidades.js
export function sumar(a, b) {
  return a + b;
}
export function restar(a, b) {
  return a - b;
}

// app.js
import { sumar, restar } from './utilidades.js';
```

## Organización de carpetas

Un ejemplo sencillo:

```
src/
  index.html
  js/
    app.js
    saludar.js
    calculadora.js
```

## Módulos y el navegador

Para usar módulos en el navegador, agrega `type="module"` en el script:

```html
<script type="module" src="js/app.js"></script>
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Módulos JS</title>
</head>
<body>
  <h1>Módulos en JavaScript</h1>
  <script type="module">
    import { saludar } from './saludar.js';
    console.log(saludar('Carlos'));
  </script>
</body>
</html>
```

## Ejercicio

1. Crea un archivo `math.js` con una función `multiplicar`.
2. Importa esa función en otro archivo y úsala.
3. Organiza un pequeño proyecto con carpetas `css`, `js` e `index.html`.

## Resumen

Los módulos ayudan a mantener tu código ordenado. En proyectos grandes, son indispensables.
