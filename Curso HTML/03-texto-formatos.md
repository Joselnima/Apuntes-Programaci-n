# Módulo 03 - Texto y formato

En este módulo vamos a dominar las etiquetas de texto más comunes y su uso semántico.

## Encabezados

```html
<h1>Título principal</h1>
<h2>Subtítulo</h2>
<h3>Sección</h3>
```

- usa un solo `h1` por página.
- sigue el orden lógico: `h1`, `h2`, `h3`, `h4`, `h5`, `h6`.
- los encabezados ayudan a la lectura y al SEO.

## Párrafos y saltos de línea

```html
<p>Este es un párrafo.</p>
<p>Otro párrafo separado.</p>
<br>
```

- `<p>` crea párrafos separados.
- `<br>` inserta un salto de línea dentro de un mismo párrafo.

## Énfasis y texto fuerte

```html
<p>Texto <strong>importante</strong> y texto <em>enfatizado</em>.</p>
```

- `<strong>` señala que el texto es importante.
- `<em>` indica énfasis en la lectura.

## Texto semántico y de código

```html
<mark>Texto resaltado</mark>
<small>Texto pequeño</small>
<abbr title="HTML">HTML</abbr>
<time datetime="2026-04-30">30 de abril de 2026</time>
<code>const nombre = 'HTML';</code>
<pre>Texto con espacios preservados</pre>
```

- `<mark>` resalta texto relevante.
- `<abbr>` define abreviaturas.
- `<time>` representa fechas y horas.
- `<code>` y `<pre>` muestran código o texto preformateado.

## Citas y referencias

```html
<blockquote>
  <p>Este es un párrafo citado de otra fuente.</p>
</blockquote>
<p>Texto normal con una cita corta <q>entre comillas</q>.</p>
```

- `<blockquote>` es para citas largas.
- `<q>` es para citas breves en línea.

## Texto técnico y especial

```html
<dfn>HTML</dfn>
<kbd>Ctrl + C</kbd>
<samp>Output del programa</samp>
<var>x</var>
<sup>10</sup>
<sub>2</sub>
```

- `<dfn>` define un término.
- `<kbd>` indica teclas.
- `<samp>` muestra resultados de un programa.
- `<var>` representa variables.
- `<sup>` y `<sub>` crean superíndice y subíndice.

## Entidades HTML

```html
<p>Menor &lt; mayor &gt; y ampersand &amp; o doble comilla &quot;.</p>
```

Las entidades son necesarias para mostrar símbolos especiales.

## Ejemplo real

```html
<article>
  <h1>¿Qué es HTML?</h1>
  <p>HTML es un lenguaje de marcado que estructura contenido en la web.</p>
  <h2>Ventajas</h2>
  <ul>
    <li><strong>Fácil de aprender</strong></li>
    <li><em>Compatible con todos los navegadores</em></li>
    <li><abbr title="HyperText Markup Language">HTML</abbr> es el estándar de la web</li>
  </ul>
  <p>La fecha de publicación es <time datetime="2026-04-30">30 de abril de 2026</time>.</p>
</article>
```

## Resumen

Usa siempre la etiqueta más adecuada para cada tipo de texto. Esto ayuda a los usuarios y a las máquinas que analizan tu página.
