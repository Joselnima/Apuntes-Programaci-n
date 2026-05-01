# Módulo 09 - Metadatos y estructura avanzada

Este módulo explica el contenido del `<head>`, metadatos importantes y otros elementos avanzados del documento HTML.

## Metadatos esenciales

```html
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="description" content="Curso completo de HTML">
<meta name="author" content="Tu Nombre">
<title>Curso de HTML</title>
```

## Otras etiquetas en `<head>`

```html
<link rel="icon" href="favicon.ico">
<link rel="canonical" href="https://www.ejemplo.com/pagina">
<meta name="robots" content="index,follow">
```

## Base y enlaces

```html
<base href="https://www.ejemplo.com/">
```

- `base` define la URL base para enlaces relativos.

## Open Graph y redes sociales

```html
<meta property="og:title" content="Curso de HTML">
<meta property="og:description" content="Aprende HTML con ejemplos claros.">
<meta property="og:image" content="https://www.ejemplo.com/imagen.jpg">
```

## Etiquetas de script y noscript

```html
<script src="app.js" defer></script>
<noscript>Necesitas JavaScript para ver este contenido.</noscript>
```

- `defer` descarga el script y lo ejecuta después del parseo.
- `noscript` muestra contenido cuando JavaScript está deshabilitado.

## Elementos avanzados de HTML

```html
<template>
  <p>Contenido reutilizable</p>
</template>
<slot></slot>
```

- `<template>` contiene HTML que no se renderiza inmediatamente.
- `<slot>` se usa en componentes web para marcar donde va contenido insertado.

## Atributos globales útiles

- `lang="es"`
- `dir="ltr"`
- `hidden`
- `tabindex="0"`
- `data-` atributos personalizados

## Ejemplo real

```html
<html lang="es">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Aprende HTML paso a paso.">
    <link rel="icon" href="favicon.ico">
    <title>Curso HTML</title>
  </head>
  <body>
    <header>
      <h1>Mi curso de HTML</h1>
    </header>
  </body>
</html>
```

## Resumen

El `<head>` contiene información del documento que no se muestra directamente. Usa metadatos correctos para mejorar la compatibilidad, la indexación y la experiencia del usuario.
