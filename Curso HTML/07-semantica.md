# Módulo 07 - Etiquetas semánticas

La semántica en HTML consiste en elegir la etiqueta correcta según el significado del contenido.

## Estructura principal

```html
<header>...</header>
<nav>...</nav>
<main>...</main>
<section>...</section>
<article>...</article>
<aside>...</aside>
<footer>...</footer>
```

### ¿Por qué es importante?

- mejora la accesibilidad
- ayuda a motores de búsqueda a entender la página
- facilita el mantenimiento
- hace el HTML más legible

## Más etiquetas semánticas

```html
<figure>...
  <figcaption>...</figcaption>
</figure>
<address>...</address>
<details>...
  <summary>...</summary>
</details>
<time datetime="2026-04-30">30 de abril de 2026</time>
<mark>...</mark>
<address>...</address>
```

### Ejemplo real: artículo con imagen y fecha

```html
<article>
  <header>
    <h1>Guía de HTML</h1>
    <p>Publicado el <time datetime="2026-04-30">30 de abril de 2026</time></p>
  </header>

  <figure>
    <img src="html-logo.png" alt="Logo de HTML">
    <figcaption>Logo oficial de HTML</figcaption>
  </figure>

  <section>
    <h2>¿Qué es HTML?</h2>
    <p>HTML es un lenguaje de marcado que estructura las páginas web.</p>
  </section>

  <aside>
    <p>Dato curioso: HTML5 introduce muchas etiquetas para contenido multimedia.</p>
  </aside>

  <footer>
    <p>Autor: Escuela web</p>
  </footer>
</article>
```

## Elementos de texto con significado

- `<strong>` importante
- `<em>` énfasis
- `<small>` texto secundario
- `<abbr>` abreviatura
- `<cite>` fuente o referencia
- `<blockquote>` cita larga
- `<q>` cita corta
- `<code>` código en línea
- `<pre>` texto preformateado

## Elementos interactivos

```html
<details>
  <summary>Mostrar más</summary>
  <p>Contenido oculto que se revela al usuario.</p>
</details>
```

## Buenas prácticas

- usa la etiqueta más específica disponible.
- no uses `<div>` cuando hay una alternativa semántica.
- organiza secciones con `section`, `article` y `aside`.
- incluye `header` y `footer` cuando sea necesario.
