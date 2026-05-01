# Módulo 14 - Proyecto práctico

En este módulo vas a aplicar todo lo aprendido con una página HTML completa y bien estructurada.

## Objetivo

Construir una página que incluya:

- estructura semántica (`header`, `main`, `section`, `footer`)
- navegación con enlaces internos
- secciones de contenido claras
- formulario de contacto
- multimedia y detalles interactivos

## Estructura HTML del proyecto

```html
<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Página de presentación del proyecto HTML">
    <title>Proyecto HTML</title>
    <link rel="icon" href="favicon.ico">
  </head>
  <body>
    <header>
      <h1>Mi proyecto HTML</h1>
      <nav>
        <a href="#inicio">Inicio</a>
        <a href="#servicios">Servicios</a>
        <a href="#contacto">Contacto</a>
      </nav>
    </header>

    <main>
      <section id="inicio">
        <h2>Bienvenido</h2>
        <p>Presentación breve sobre el proyecto y su propósito.</p>
      </section>

      <section id="servicios">
        <h2>Servicios</h2>
        <article>
          <h3>Servicio 1</h3>
          <p>Descripción del servicio.</p>
        </article>
        <article>
          <h3>Servicio 2</h3>
          <p>Descripción del servicio.</p>
        </article>
      </section>

      <section id="multimedia">
        <h2>Multimedia</h2>
        <figure>
          <img src="imagen.jpg" alt="Descripción de la imagen">
          <figcaption>Ejemplo de imagen con pie de foto.</figcaption>
        </figure>
        <video controls>
          <source src="video.mp4" type="video/mp4">
        </video>
      </section>

      <section id="faq">
        <h2>Preguntas frecuentes</h2>
        <details>
          <summary>¿Qué ofrece este proyecto?</summary>
          <p>Ofrece una estructura HTML completa y accesible.</p>
        </details>
      </section>

      <section id="contacto">
        <h2>Contacto</h2>
        <form action="/contacto" method="post">
          <label for="nombre">Nombre</label>
          <input id="nombre" name="nombre" type="text" required>

          <label for="correo">Correo</label>
          <input id="correo" name="correo" type="email" required>

          <button type="submit">Enviar</button>
        </form>
      </section>
    </main>

    <footer>
      <p>© 2026 Proyecto HTML</p>
    </footer>
  </body>
</html>
```

## Ejercicio del proyecto

- añade una lista de características.
- agrega una tabla de precios o comparación.
- usa `figure` y `figcaption` para imágenes.
- incluye un video o audio.
- pon `details` para preguntas frecuentes.

## Resumen

Este proyecto debe ser una página completa y clara. Usa solo HTML bien estructurado; la apariencia se puede mejorar después en otro curso.
