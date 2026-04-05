# Módulo 06 - HTML Desde Cero

> **Estructura de páginas web - Lo que Angular renderiza**

---

## ¿QUÉ es HTML?

**HTML = Lenguaje para describir estructura de páginas**

Es como un plano:
- 🏗️ Plano = HTML (estructura)
- 🎨 Pintura = CSS (diseño)
- ⚙️ Electricidad = JavaScript (interactividad)

```html
<!-- ❌ Sin HTML: Nada -->

<!-- ✅ Con HTML: Estructura -->
<h1>Título</h1>
<p>Párrafo</p>
<button>Botón</button>
```

---

## ¿PARA QUÉ?

### Problema 1: Mostrar información

```html
<!-- ❌ Sin HTML -->
Hola mundo  <!-- Solo texto plano, sin estructura -->

<!-- ✅ Con HTML -->
<h1>Hola mundo</h1>  <!-- Título importante -->
<p>Párrafo descriptivo</p>
```

### Problema 2: Entrada del usuario

```html
<!-- ❌ Sin HTML -->
Escribe aquí:  <!-- ¿Dónde escribe? -->

<!-- ✅ Con HTML -->
<label>Tu nombre:</label>
<input type="text" placeholder="Escribe aquí">
```

### Problema 3: Estructura semántica

```html
<!-- ❌ Confuso (todo es div) -->
<div>Menú</div>
<div>Contenido</div>
<div>Pie</div>

<!-- ✅ Claro (estructura clara) -->
<header>Menú</header>
<main>Contenido</main>
<footer>Pie</footer>
```

---

## ¿CÓMO funciona?

### 1. Estructura Básica

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Mi Página</title>
</head>
<body>
  <!-- Contenido aquí -->
  <h1>Hola</h1>
</body>
</html>
```

**¿Qué es cada parte?**
- `<!DOCTYPE html>` - Esto es HTML5 (moderno)
- `<html>` - Raíz del documento
- `<head>` - Metadatos (título, estilos)
- `<body>` - Lo visible en pantalla

### 2. Etiquetas Básicas

```html
<!-- Títulos (h1 a h6) -->
<h1>Título Principal</h1>
<h2>Subtítulo</h2>
<h3>Sub-subtítulo</h3>

<!-- Párrafos -->
<p>Un párrafo de texto normal.</p>

<!-- Enlaces -->
<a href="https://ejemplo.com">Ir a Ejemplo</a>

<!-- Imágenes -->
<img src="foto.jpg" alt="Descripción">

<!-- Lista no ordenada -->
<ul>
  <li>Item 1</li>
  <li>Item 2</li>
  <li>Item 3</li>
</ul>

<!-- Lista ordenada -->
<ol>
  <li>Primero</li>
  <li>Segundo</li>
  <li>Tercero</li>
</ol>
```

### 3. Formularios (MUY IMPORTANTE EN ANGULAR)

```html
<!-- Formulario simple -->
<form>
  <!-- Input de texto -->
  <label for="nombre">Nombre:</label>
  <input id="nombre" type="text" placeholder="Tu nombre">
  
  <!-- Input de email -->
  <label for="email">Email:</label>
  <input id="email" type="email" placeholder="tu@email.com">
  
  <!-- Input de contraseña -->
  <label for="password">Contraseña:</label>
  <input id="password" type="password">
  
  <!-- Textarea (texto largo) -->
  <label for="mensaje">Mensaje:</label>
  <textarea id="mensaje" rows="5"></textarea>
  
  <!-- Select (dropdown) -->
  <label for="pais">País:</label>
  <select id="pais">
    <option value="">Selecciona...</option>
    <option value="mx">México</option>
    <option value="ar">Argentina</option>
    <option value="pe">Perú</option>
  </select>
  
  <!-- Checkbox -->
  <label>
    <input type="checkbox"> Acepto términos
  </label>
  
  <!-- Radio -->
  <label>
    <input type="radio" name="genero"> Masculino
  </label>
  <label>
    <input type="radio" name="genero"> Femenino
  </label>
  
  <!-- Botones -->
  <button type="submit">Enviar</button>
  <button type="reset">Limpiar</button>
  <button type="button">Cancelar</button>
</form>
```

### 4. Estructura Semántica

```html
<!-- ❌ Viejo (todo div) -->
<div class="header">
  <div>Logo</div>
  <div>Menú</div>
</div>
<div class="content">
  <div class="article">...</div>
  <div class="sidebar">...</div>
</div>
<div class="footer">...</div>

<!-- ✅ Moderno (semántico) -->
<header>
  <figure>Logo</figure>
  <nav>Menú</nav>
</header>
<main>
  <article>Contenido</article>
  <aside>Sidebar</aside>
</main>
<footer>Pie de página</footer>
```

### 5. Atributos Comunes

```html
<!-- id: identificador único -->
<div id="main">Contenido principal</div>

<!-- class: para CSS -->
<div class="card">Tarjeta</div>

<!-- data-*: datos personalizados -->
<button data-product-id="123">Comprar</button>

<!-- aria-*: accesibilidad -->
<button aria-label="Menú">☰</button>

<!-- title: tooltip -->
<button title="Haz clic para guardar">Guardar</button>

<!-- disabled: deshabilitado -->
<button disabled>No se puede hacer clic</button>

<!-- required: obligatorio en formularios -->
<input type="email" required>
```

---

## En Angular

**Angular usa exactamente este HTML:**

```typescript
// app.component.ts
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <h1>{{ titulo }}</h1>
    <button (click)="saludar()">Haz clic</button>
    <p *ngIf="mostrar">Texto visible</p>
  `
})
export class AppComponent {
  titulo = 'Mi App Angular';
  mostrar = true;
  
  saludar() {
    alert('¡Hola!');
  }
}
```

**Lo verás como HTML normal en el navegador** ✅

---

## Tabla de Etiquetas Comunes

| Etiqueta | Para | Ejemplo |
|----------|------|---------|
| `<h1>-<h6>` | Títulos | `<h1>Mi título</h1>` |
| `<p>` | Párrafo | `<p>Texto</p>` |
| `<a>` | Enlace | `<a href="#">Link</a>` |
| `<img>` | Imagen | `<img src="">` |
| `<button>` | Botón | `<button>Click</button>` |
| `<input>` | Entrada | `<input type="text">` |
| `<textarea>` | Texto largo | `<textarea></textarea>` |
| `<select>` | Dropdown | `<select><option>...</option></select>` |
| `<ul>/<ol>` | Listas | `<ul><li>Item</li></ul>` |
| `<table>` | Tabla | `<table><tr><td>Celda</td></tr></table>` |
| `<header>` | Encabezado | `<header>Logo</header>` |
| `<footer>` | Pie | `<footer>Copyright</footer>` |
| `<nav>` | Navegación | `<nav>Menú</nav>` |
| `<main>` | Contenido | `<main>Contenido</main>` |

---

## Checklist HTML

- [ ] Entiendes estructura básica
- [ ] Sabes usar formularios
- [ ] Etiquetas semánticas claras
- [ ] Atributos correctos
- [ ] Nidificación correcta
- [ ] Accesibilidad (labels, alt)

---

## Próximo

Módulo 07: CSS desde cero (hacer que se vea bonito)
