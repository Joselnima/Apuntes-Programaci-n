# Módulo 01 - Introducción a JavaScript

JavaScript es el lenguaje que hace que las páginas web sean interactivas. Con JavaScript podemos cambiar el contenido, responder a acciones del usuario y crear experiencias dinámicas.

## ¿Qué es JavaScript?

- Es un lenguaje de programación que se ejecuta en el navegador.
- Permite modificar elementos de la página sin recargar.
- Se usa en el frontend, y también en el backend con Node.js.

## ¿Por qué aprender JavaScript?

- Casi todas las páginas web lo usan.
- Es el lenguaje más importante para el desarrollo web.
- Te permite crear aplicaciones reales, juegos y herramientas.

## Primer ejemplo

Crea un archivo llamado `index.html` y pega esto:

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Mi primer JavaScript</title>
</head>
<body>
  <h1>Hola desde JavaScript</h1>
  <button id="boton">Haz clic</button>

  <script>
    const boton = document.getElementById('boton');
    boton.addEventListener('click', () => {
      alert('¡Has hecho clic!');
    });
  </script>
</body>
</html>
```

### ¿Qué hace este código?

- Crea un botón en la página.
- Usa `document.getElementById` para encontrarlo.
- Añade un evento `click` que muestra un mensaje.

## Cómo ejecutar JavaScript

- Abre el archivo `index.html` en tu navegador.
- Cada vez que hagas clic en el botón, verás la alerta.

### Herramientas recomendadas

- Editor de código: Visual Studio Code, Atom o Sublime Text.
- Navegador: Chrome, Edge, Firefox o Safari.
- Consola del navegador: presiona `F12` o `Ctrl+Shift+I`.

## ¿Qué veremos en este curso?

- Variables y tipos
- Condicionales y bucles
- Funciones y objetos
- Trabajo con el DOM
- Asincronía y `async/await`
- Características modernas de ECMAScript 8

## Ejercicio

1. Cambia el texto del botón a "Presióname".
2. Modifica la alerta para que muestre "¡Bienvenido a JavaScript!".
3. Añade un segundo botón que cambie el color del fondo.

## Resumen

JavaScript es el lenguaje que da vida a las páginas web. En el siguiente módulo comenzaremos con las variables y los tipos de datos, que son la base de cualquier programa.
