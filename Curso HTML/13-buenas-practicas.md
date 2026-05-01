# Módulo 13 - Buenas prácticas

Estas son recomendaciones para escribir HTML limpio, mantenible y profesional.

## Indentación y legibilidad

```html
<section>
  <h2>Título</h2>
  <p>Texto claro.</p>
</section>
```

Usa sangrías para que el código sea fácil de leer.

## Nombres claros para clases y archivos

- `.header`
- `.product-card`
- `index.html`
- `styles.css`

Evita nombres como `.box` cuando no describen su función.

## Comentarios útiles

```html
<!-- Sección de productos destacados -->
<section>
  ...
</section>
```

Usa comentarios solo cuando aclaren el propósito.

## Validación y limpieza

- revisa tu HTML con el validador de W3C.
- corrige etiquetas abiertas o atributos faltantes.
- elimina código muerto o duplicado.

## Accesibilidad y semántica

- usa `alt` en todas las imágenes.
- asocia siempre `label` con `input`.
- usa etiquetas semánticas como `main`, `article`, `nav`.

## Rendimiento

- optimiza imágenes para web.
- usa solo el CSS que necesitas.
- evita cargar scripts innecesarios.

## Ejemplo de estructura buena

```html
<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Mi página</title>
    <link rel="stylesheet" href="styles.css">
  </head>
  <body>
    <header>
      <h1>Mi sitio</h1>
    </header>

    <main>
      <section>
        <h2>Bienvenida</h2>
        <p>Texto introductorio.</p>
      </section>
    </main>

    <footer>
      <p>© 2026</p>
    </footer>
  </body>
</html>
```

## Resumen

HTML bien escrito es la base de páginas profesionales. Sigue estas prácticas para que tu código sea más fácil de mantener, de entender y de escalar.