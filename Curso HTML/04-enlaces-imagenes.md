# Módulo 04 - Enlaces, imágenes y multimedia

En este módulo aprenderás a conectar páginas, mostrar imágenes y usar elementos multimedia.

## Enlaces con `<a>`

```html
<a href="https://www.ejemplo.com">Visita este sitio</a>
```

### Tipos de enlaces

- Enlace externo:
  ```html
  <a href="https://www.ejemplo.com" target="_blank" rel="noreferrer noopener">Sitio externo</a>
  ```
- Enlace interno:
  ```html
  <a href="./pagina.html">Página interna</a>
  ```
- Enlace de correo:
  ```html
  <a href="mailto:correo@ejemplo.com">Enviar correo</a>
  ```
- Enlace de teléfono:
  ```html
  <a href="tel:+34123456789">Llamar</a>
  ```
- Enlace de descarga:
  ```html
  <a href="documento.pdf" download>Descargar PDF</a>
  ```

## Imágenes con `<img>`

```html
<img src="logo.png" alt="Logotipo de la marca">
```

### Buenas prácticas con imágenes

- `alt` describe la imagen.
- `loading="lazy"` mejora el rendimiento.
- evita imágenes sin `alt` si aportan información.

## Figura y pie de foto

```html
<figure>
  <img src="producto.jpg" alt="Camiseta azul">
  <figcaption>Camiseta azul con diseño moderno.</figcaption>
</figure>
```

## Imágenes responsivas con `picture`

```html
<picture>
  <source srcset="imagen.webp" type="image/webp">
  <img src="imagen.jpg" alt="Ejemplo de imagen">
</picture>
```

## Audio y video

### Audio

```html
<audio controls>
  <source src="audio.mp3" type="audio/mpeg">
  Tu navegador no soporta audio.
</audio>
```

### Video

```html
<video controls width="640">
  <source src="video.mp4" type="video/mp4">
  Tu navegador no soporta video.
</video>
```

## Integración con `iframe` y `embed`

### Iframe

```html
<iframe src="https://www.ejemplo.com" title="Sitio embebido" width="600" height="400"></iframe>
```

### Embed

```html
<embed src="documento.pdf" type="application/pdf" width="600" height="400">
```

## Ejemplo real: tarjeta con enlace e imagen

```html
<article>
  <a href="producto.html">
    <figure>
      <img src="producto.jpg" alt="Camiseta azul">
      <figcaption>Camiseta azul, algodón 100%.</figcaption>
    </figure>
    <h2>Camiseta azul</h2>
  </a>
</article>
```

## Resumen

Enlaces y multimedia son parte esencial de cualquier página web. Usa etiquetas correctas y siempre añade texto descriptivo para accesibilidad.
