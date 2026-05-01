# Módulo 02 - Estructura de un documento HTML

Este módulo explica la estructura mínima de cualquier documento HTML y qué etiquetas son indispensables.

## La estructura básica

```html
<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="UTF-8">
    <title>Mi sitio web</title>
  </head>
  <body>
    <!-- Contenido visible -->
  </body>
</html>
```

### Elementos principales

- `<!DOCTYPE html>`: declara la versión de HTML.
- `<html>`: raíz del documento.
- `<head>`: metadatos, enlaces a estilos, scripts y título.
- `<body>`: contenido visible en la página.

## Etiquetas importantes dentro de `<head>`

```html
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="Curso completo de HTML">
  <title>Curso HTML</title>
  <link rel="stylesheet" href="styles.css">
</head>
```

- `meta charset`: define la codificación de caracteres.
- `meta viewport`: hace que la página sea responsiva en móviles.
- `meta description`: ayuda al SEO.
- `link rel="stylesheet"`: conecta CSS.

## Ejemplo real con estructura completa

```html
<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Página personal con HTML y CSS">
    <title>Mi portafolio</title>
    <link rel="stylesheet" href="styles.css">
  </head>
  <body>
    <header>
      <h1>Mi portafolio personal</h1>
    </header>

    <main>
      <section>
        <h2>Sobre mí</h2>
        <p>Bienvenido a mi sitio web.</p>
      </section>
    </main>

    <footer>
      <p>Derechos reservados © 2026</p>
    </footer>
  </body>
</html>
```

## Buenas prácticas

- usa siempre `lang="es"` si tu contenido está en español.
- coloca el `title` dentro de `<head>`.
- pon los estilos en archivos separados cuando el proyecto crece.
- escribe comentarios con `<!-- comentario -->` para recordarte la función de cada sección.
