# Módulo 08 - Multimedia y elementos interactivos

En este módulo veremos los elementos HTML que permiten audio, video, contenido embebido y acciones interactivas.

## Audio en HTML

```html
<audio controls>
  <source src="cancion.mp3" type="audio/mpeg">
  <source src="cancion.ogg" type="audio/ogg">
  Tu navegador no soporta audio.
</audio>
```

## Video en HTML

```html
<video controls width="640">
  <source src="video.mp4" type="video/mp4">
  <source src="video.webm" type="video/webm">
  Tu navegador no soporta video.
</video>
```

## Imágenes avanzadas con `picture`

```html
<picture>
  <source srcset="imagen.webp" type="image/webp">
  <img src="imagen.jpg" alt="Ejemplo de imagen" loading="lazy">
</picture>
```

## Contenido embebido

### Iframe

```html
<iframe src="https://www.ejemplo.com" title="Contenido embebido" width="600" height="400"></iframe>
```

### Embed

```html
<embed src="documento.pdf" type="application/pdf" width="600" height="400">
```

### Object

```html
<object data="documento.pdf" type="application/pdf" width="600" height="400">
  <p>Tu navegador no soporta PDF embebido.</p>
</object>
```

## Elementos interactivos nativos

### Details y summary

```html
<details>
  <summary>Pregunta frecuente</summary>
  <p>Respuesta de la pregunta.</p>
</details>
```

### Progreso y métrica

```html
<progress value="70" max="100">70%</progress>
<meter value="0.6">60%</meter>
```

### Dialog

```html
<dialog open>
  <p>Este es un cuadro de diálogo.</p>
</dialog>
```

## Ejemplo real

```html
<section>
  <h2>Video de introducción</h2>
  <video controls width="560">
    <source src="intro.mp4" type="video/mp4">
  </video>
</section>

<section>
  <h2>Preguntas frecuentes</h2>
  <details>
    <summary>¿Qué es HTML?</summary>
    <p>HTML es el lenguaje que estructura las páginas web.</p>
  </details>
</section>
```

## Resumen

HTML incluye muchos elementos que no son solo texto. Con audio, video y controles interactivos puedes ofrecer experiencias más ricas sin añadir JavaScript.
