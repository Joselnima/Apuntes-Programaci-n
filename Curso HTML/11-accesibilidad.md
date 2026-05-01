# Módulo 11 - Accesibilidad web

La accesibilidad permite que más personas usen tu sitio, especialmente quienes tienen dificultades visuales, motoras o cognitivas.

## Texto alternativo en imágenes

```html
<img src="logo.png" alt="Logotipo de la empresa">
```

- `alt` describe la imagen.
- si la imagen es decorativa, usa `alt=""`.

## Etiquetas `label`

```html
<label for="email">Correo</label>
<input id="email" type="email" name="email">
```

Esto ayuda a lectores de pantalla y mejora la usabilidad.

## Uso de roles y ARIA

```html
<nav aria-label="Menú principal">
  <a href="#home">Inicio</a>
</nav>
```

- `aria-label` describe secciones para tecnología asistiva.
- usa `role` solo cuando una etiqueta semántica no existe.

## Enlaces claros

```html
<a href="/contacto">Contacta al equipo</a>
```

Evita enlaces como "haz clic aquí". El texto debe explicar a dónde lleva.

## Teclado y foco

Asegúrate de que el usuario pueda navegar con el teclado.

```css
button:focus,
a:focus {
  outline: 2px solid #007bff;
}
```

## Ejemplo real: formulario accesible

```html
<form>
  <div>
    <label for="nombre">Nombre completo</label>
    <input id="nombre" name="nombre" type="text" required>
  </div>

  <div>
    <label for="mensaje">Mensaje</label>
    <textarea id="mensaje" name="mensaje"></textarea>
  </div>

  <button type="submit">Enviar</button>
</form>
```

## Buenas prácticas

- usa `alt` en todas las imágenes.
- siempre asocia `label` con `input`.
- usa encabezados en orden lógico (`h1`, `h2`, `h3`).
- prueba tu página con el teclado.
