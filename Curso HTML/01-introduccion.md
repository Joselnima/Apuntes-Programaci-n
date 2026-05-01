# Módulo 01 - Introducción a HTML

HTML es el lenguaje que estructura el contenido de una página web. No es un lenguaje de programación: es un lenguaje de marcas.

## ¿Para qué sirve HTML?

- definir títulos y párrafos
- mostrar imágenes y enlaces
- crear listas y tablas
- construir formularios
- dar sentido a la información para navegadores y motores de búsqueda

## El primer documento HTML

```html
<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="UTF-8">
    <title>Mi primera página</title>
  </head>
  <body>
    <h1>Hola, mundo</h1>
    <p>Esta es mi primera página web con HTML.</p>
  </body>
</html>
```

### Qué hace cada parte

- `<!DOCTYPE html>`: indica que estamos usando HTML5.
- `<html lang="es">`: define el documento HTML y el idioma.
- `<head>`: contiene información oculta para el navegador.
- `<meta charset="UTF-8">`: asegura que los caracteres se muestren bien.
- `<title>`: texto que aparece en la pestaña del navegador.
- `<body>`: aquí va todo lo que se muestra en la página.

## Ejemplo real: una página de presentación

```html
<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sobre mí</title>
  </head>
  <body>
    <header>
      <h1>Mi Portafolio</h1>
      <p>Desarrollador web en formación.</p>
    </header>

    <section>
      <h2>Sobre mí</h2>
      <p>Me gusta crear páginas limpias y fáciles de usar.</p>
    </section>

    <footer>
      <p>Contacto: correo@ejemplo.com</p>
    </footer>
  </body>
</html>
```

## Resumen

HTML es el primer paso para construir páginas web. Aprenderlo bien te permite crear estructuras claras, accesibles y fáciles de ampliar con CSS y JavaScript.