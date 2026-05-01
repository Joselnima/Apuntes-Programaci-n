# Módulo 16 - Frameworks CSS

En este módulo aprenderás sobre los frameworks CSS más populares, sus características, ventajas, desventajas, y cuándo usar cada uno. Cubriremos Bootstrap, Tailwind CSS, Bulma, Foundation y otros frameworks modernos.

## Introducción a Frameworks CSS

### ¿Qué es un framework CSS?

```html
<!-- Sin framework - CSS personalizado -->
<style>
  .btn { padding: 10px 20px; background: blue; color: white; border: none; border-radius: 4px; }
  .container { max-width: 1200px; margin: 0 auto; padding: 0 15px; }
  .row { display: flex; flex-wrap: wrap; margin: 0 -15px; }
  .col { flex: 1; padding: 0 15px; }
</style>

<!-- Con framework - clases predefinidas -->
<button class="btn btn-primary">Botón</button>
<div class="container">
  <div class="row">
    <div class="col-md-6">Columna 1</div>
    <div class="col-md-6">Columna 2</div>
  </div>
</div>
```

### Ventajas de usar frameworks

- **Rapidez de desarrollo**: Componentes listos para usar
- **Consistencia**: Diseño uniforme en toda la aplicación
- **Responsive**: Diseño móvil-first incluido
- **Comunidad**: Soporte y actualizaciones constantes
- **Documentación**: Guías completas y ejemplos
- **Accesibilidad**: Componentes accesibles incluidos

### Desventajas

- **Tamaño**: Incluye CSS que no se usa
- **Personalización limitada**: Difícil modificar componentes base
- **Aprendizaje**: Curva de aprendizaje inicial
- **Dependencia**: Código depende del framework
- **Bloat**: Estilos innecesarios

## Bootstrap

### Introducción a Bootstrap

Bootstrap es el framework CSS más popular, creado por Twitter. Ofrece un sistema de grid de 12 columnas, componentes preestilizados y utilidades CSS.

### Instalación

```html
<!-- CDN -->
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>

<!-- NPM -->
npm install bootstrap
```

```scss
// Con Sass
@import '~bootstrap/scss/bootstrap';
```

### Sistema de Grid

```html
<div class="container">
  <div class="row">
    <div class="col-sm-12 col-md-6 col-lg-4">
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">Card title</h5>
          <p class="card-text">Some quick example text.</p>
          <a href="#" class="btn btn-primary">Go somewhere</a>
        </div>
      </div>
    </div>
  </div>
</div>
```

### Componentes principales

```html
<!-- Botones -->
<button class="btn btn-primary">Primary</button>
<button class="btn btn-secondary">Secondary</button>
<button class="btn btn-success btn-lg">Large Success</button>

<!-- Formularios -->
<form>
  <div class="mb-3">
    <label for="email" class="form-label">Email</label>
    <input type="email" class="form-control" id="email">
  </div>
  <div class="mb-3">
    <label for="password" class="form-label">Password</label>
    <input type="password" class="form-control" id="password">
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
</form>

<!-- Navbar -->
<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <div class="container-fluid">
    <a class="navbar-brand" href="#">Navbar</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="navbar-nav">
          <a class="nav-link active" href="#">Home</a>
        </li>
      </ul>
    </div>
  </div>
</nav>
```

### Utilidades CSS

```html
<!-- Espaciado -->
<div class="m-3 p-3">Margin y padding</div>
<div class="mt-5 mb-3">Margin top y bottom</div>

<!-- Colores -->
<div class="text-primary">Texto primario</div>
<div class="bg-success text-white">Fondo success</div>

<!-- Display -->
<div class="d-none d-md-block">Oculto en móvil, visible en desktop</div>
<div class="d-flex justify-content-center align-items-center">Flexbox utilities</div>
```

### Personalización

```scss
// custom.scss
@import '~bootstrap/scss/functions';
@import '~bootstrap/scss/variables';

// Personalizar variables
$primary: #007bff;
$secondary: #6c757d;
$success: #28a745;

// Importar Bootstrap
@import '~bootstrap/scss/bootstrap';
```

### Ventajas de Bootstrap

- ✅ **Maduro y estable**: 10+ años de desarrollo
- ✅ **Gran comunidad**: Miles de temas y componentes
- ✅ **Completo**: Todo lo que necesitas incluido
- ✅ **Documentación excelente**: Guías detalladas
- ✅ **JavaScript incluido**: Componentes interactivos

### Desventajas

- ❌ **Pesado**: ~200KB minificado
- ❌ **Genérico**: Todos los sitios se ven similares
- ❌ **Overriding**: Difícil personalizar profundamente
- ❌ **Dependiente**: Actualizaciones pueden romper cosas

## Tailwind CSS

### Introducción a Tailwind

Tailwind CSS es un framework utility-first que proporciona clases de bajo nivel para construir diseños personalizados sin salir del HTML.

### Instalación

```bash
npm install -D tailwindcss
npx tailwindcss init
```

```css
/* tailwind.css */
@tailwind base;
@tailwind components;
@tailwind utilities;
```

### Filosofía Utility-First

```html
<!-- Bootstrap approach -->
<div class="card">
  <div class="card-body">
    <h5 class="card-title">Card title</h5>
    <p class="card-text">Some text</p>
    <a href="#" class="btn btn-primary">Button</a>
  </div>
</div>

<!-- Tailwind approach -->
<div class="max-w-sm mx-auto bg-white rounded-xl shadow-md overflow-hidden">
  <div class="p-6">
    <h2 class="text-2xl font-bold text-gray-900">Card title</h2>
    <p class="text-gray-600 mt-2">Some text here</p>
    <button class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
      Button
    </button>
  </div>
</div>
```

### Sistema de diseño

```html
<!-- Layout -->
<div class="container mx-auto px-4">
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <!-- Grid items -->
  </div>
</div>

<!-- Typography -->
<h1 class="text-4xl font-bold text-gray-900">Heading</h1>
<p class="text-lg text-gray-600 leading-relaxed">Paragraph text</p>

<!-- Colors -->
<div class="bg-blue-500 text-white p-4">Blue background</div>
<div class="bg-red-100 text-red-800 p-4">Red alert</div>

<!-- Spacing -->
<div class="m-4 p-8 space-y-4">
  <div>Item 1</div>
  <div>Item 2</div>
</div>
```

### Responsive Design

```html
<!-- Responsive utilities -->
<div class="w-full md:w-1/2 lg:w-1/3">
  <h1 class="text-2xl md:text-3xl lg:text-4xl">Responsive text</h1>
  <p class="hidden md:block">Visible on medium and up</p>
</div>
```

### Componentes personalizados

```css
/* tailwind.config.js */
module.exports = {
  theme: {
    extend: {
      colors: {
        'brand': '#007acc',
      },
      spacing: {
        '18': '4.5rem',
      }
    }
  }
}
```

```css
/* components.css */
@layer components {
  .btn {
    @apply px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors;
  }
  
  .card {
    @apply bg-white rounded-lg shadow-md p-6;
  }
}
```

### JIT Compiler

```javascript
// tailwind.config.js - JIT mode
module.exports = {
  mode: 'jit',
  purge: ['./src/**/*.{js,ts,jsx,tsx}'],
  // ...
}
```

### Ventajas de Tailwind

- ✅ **Utility-first**: Construye cualquier diseño
- ✅ **Ligero**: Solo incluye utilidades usadas (con purging)
- ✅ **Consistente**: Sistema de diseño coherente
- ✅ **Personalizable**: Fácil modificar y extender
- ✅ **Developer experience**: Rápido desarrollo

### Desventajas

- ❌ **Curva de aprendizaje**: Muchas clases que recordar
- ❌ **HTML verboso**: Mucho código en markup
- ❌ **Mantenimiento**: Cambios requieren modificar HTML
- ❌ **Inicial setup**: Configuración más compleja

## Bulma

### Introducción a Bulma

Bulma es un framework CSS moderno basado en Flexbox, con un enfoque en simplicidad y elegancia.

### Instalación

```html
<!-- CDN -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">

<!-- NPM -->
npm install bulma
```

```scss
// Con Sass
@import '~bulma/bulma';
```

### Sistema de columnas

```html
<div class="columns">
  <div class="column">
    <div class="notification is-primary">
      <p class="title">Column 1</p>
      <p>Content here</p>
    </div>
  </div>
  <div class="column">
    <div class="notification is-info">
      <p class="title">Column 2</p>
      <p>Content here</p>
    </div>
  </div>
  <div class="column">
    <div class="notification is-warning">
      <p class="title">Column 3</p>
      <p>Content here</p>
    </div>
  </div>
</div>
```

### Componentes

```html
<!-- Hero -->
<section class="hero is-primary">
  <div class="hero-body">
    <p class="title">Hero title</p>
    <p class="subtitle">Hero subtitle</p>
  </div>
</section>

<!-- Cards -->
<div class="card">
  <div class="card-content">
    <p class="title">Card title</p>
    <p class="subtitle">Card subtitle</p>
  </div>
  <footer class="card-footer">
    <a href="#" class="card-footer-item">Save</a>
    <a href="#" class="card-footer-item">Edit</a>
  </footer>
</div>

<!-- Buttons -->
<button class="button is-primary is-large">Large Primary</button>
<button class="button is-loading">Loading</button>
```

### Modificadores

```html
<!-- Colors -->
<button class="button is-primary">Primary</button>
<button class="button is-link">Link</button>
<button class="button is-info">Info</button>

<!-- Sizes -->
<button class="button is-small">Small</button>
<button class="button is-normal">Normal</button>
<button class="button is-large">Large</button>

<!-- States -->
<button class="button is-loading">Loading</button>
<button class="button is-disabled">Disabled</button>
```

### Variables Sass

```scss
// Custom variables
$primary: #00d1b2;
$info: #209cee;

// Import Bulma
@import '~bulma/bulma';
```

### Ventajas de Bulma

- ✅ **Simple**: API limpia y fácil de aprender
- ✅ **Flexbox**: Moderno y flexible
- ✅ **Ligero**: ~200KB pero modular
- ✅ **Elegante**: Diseño moderno por defecto
- ✅ **Sass**: Fácil personalización

### Desventajas

- ❌ **Menos componentes**: Menos componentes que Bootstrap
- ❌ **Comunidad pequeña**: Menos recursos disponibles
- ❌ **JavaScript**: No incluye JS (solo CSS)

## Foundation

### Introducción a Foundation

Foundation es un framework responsive avanzado creado por ZURB, enfocado en desarrollo profesional.

### Instalación

```html
<!-- CDN -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/foundation-sites@6.7.5/dist/css/foundation.min.css">
<script src="https://cdn.jsdelivr.net/npm/foundation-sites@6.7.5/dist/js/foundation.min.js"></script>

<!-- NPM -->
npm install foundation-sites
```

### Grid system avanzado

```html
<div class="grid-container">
  <div class="grid-x grid-margin-x">
    <div class="cell small-12 medium-6 large-4">
      <div class="card">
        <div class="card-section">
          <h4>Card 1</h4>
          <p>Content</p>
        </div>
      </div>
    </div>
  </div>
</div>
```

### Componentes avanzados

```html
<!-- Orbit slider -->
<div class="orbit" role="region" aria-label="Favorite Space Pictures" data-orbit>
  <div class="orbit-wrapper">
    <div class="orbit-controls">
      <button class="orbit-previous"><span class="show-for-sr">Previous Slide</span></button>
      <button class="orbit-next"><span class="show-for-sr">Next Slide</span></button>
    </div>
    <ul class="orbit-container">
      <li class="is-active orbit-slide">
        <figure class="orbit-figure">
          <img class="orbit-image" src="img1.jpg" alt="Space">
        </figure>
      </li>
    </ul>
  </div>
</div>

<!-- Menu -->
<ul class="menu">
  <li><a href="#">Item 1</a></li>
  <li><a href="#">Item 2</a></li>
  <li class="is-active"><a href="#">Active Item</a></li>
</ul>
```

### Utilidades avanzadas

```html
<!-- Visibility -->
<div class="show-for-small-only">Visible solo en móvil</div>
<div class="hide-for-medium-up">Oculto en tablet y desktop</div>

<!-- Float -->
<div class="float-left">Izquierda</div>
<div class="float-right">Derecha</div>

<!-- Flexbox -->
<div class="flex-container align-center">
  <div class="flex-child-auto">Auto</div>
  <div class="flex-child-shrink">Shrink</div>
</div>
```

## Materialize CSS

### Introducción a Materialize

Materialize implementa Material Design de Google, con componentes modernos y animaciones fluidas.

### Instalación

```html
<!-- CDN -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>

<!-- NPM -->
npm install materialize-css
```

### Componentes Material Design

```html
<!-- Navbar -->
<nav>
  <div class="nav-wrapper">
    <a href="#" class="brand-logo">Logo</a>
    <ul id="nav-mobile" class="right hide-on-med-and-down">
      <li><a href="#">Home</a></li>
      <li><a href="#">About</a></li>
    </ul>
  </div>
</nav>

<!-- Cards -->
<div class="row">
  <div class="col s12 m6">
    <div class="card">
      <div class="card-image">
        <img src="img.jpg">
        <span class="card-title">Card Title</span>
      </div>
      <div class="card-content">
        <p>Content here</p>
      </div>
      <div class="card-action">
        <a href="#">Action</a>
      </div>
    </div>
  </div>
</div>
```

### JavaScript components

```javascript
// Inicializar componentes
document.addEventListener('DOMContentLoaded', function() {
  // Modal
  var elems = document.querySelectorAll('.modal');
  var instances = M.Modal.init(elems);
  
  // Dropdown
  var elems = document.querySelectorAll('.dropdown-trigger');
  var instances = M.Dropdown.init(elems);
  
  // Carousel
  var elems = document.querySelectorAll('.carousel');
  var instances = M.Carousel.init(elems);
});
```

## Comparación de Frameworks

| Framework | Tamaño | Curva aprendizaje | Personalización | Comunidad |
|-----------|--------|-------------------|-----------------|-----------|
| Bootstrap | Grande | Baja | Media | Excelente |
| Tailwind | Pequeño* | Media | Alta | Excelente |
| Bulma | Mediano | Baja | Alta | Buena |
| Foundation | Grande | Media | Alta | Buena |
| Materialize | Mediano | Baja | Media | Buena |

*Con purging automático

## Cuándo usar cada framework

### Usa Bootstrap si:
- Quieres empezar rápido
- Necesitas muchos componentes
- Tienes un equipo grande
- El proyecto es corporativo/tradicional

### Usa Tailwind si:
- Quieres control total del diseño
- El proyecto es altamente personalizado
- Prefieres utility-first approach
- Te importa el tamaño del bundle

### Usa Bulma si:
- Quieres algo simple pero elegante
- Te gusta Flexbox
- No necesitas JavaScript incluido
- Proyecto pequeño/mediano

### Usa Foundation si:
- Necesitas componentes avanzados
- Proyecto enterprise
- Quieres máxima flexibilidad
- Tienes experiencia con frameworks

### Usa Materialize si:
- Sigues Material Design
- Necesitas animaciones fluidas
- Proyecto mobile-first
- Te gusta el estilo de Google

## Integración con herramientas modernas

### Con React

```jsx
// Bootstrap con React
import { Button, Card } from 'react-bootstrap';

function App() {
  return (
    <Card>
      <Card.Body>
        <Card.Title>Card Title</Card.Title>
        <Button variant="primary">Button</Button>
      </Card.Body>
    </Card>
  );
}

// Tailwind con React
function Button({ children, variant = 'primary' }) {
  const baseClasses = 'px-4 py-2 rounded font-medium';
  const variants = {
    primary: 'bg-blue-500 text-white hover:bg-blue-600',
    secondary: 'bg-gray-500 text-white hover:bg-gray-600'
  };
  
  return (
    <button className={`${baseClasses} ${variants[variant]}`}>
      {children}
    </button>
  );
}
```

### Con Vue.js

```vue
<template>
  <div class="container mx-auto">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div v-for="item in items" :key="item.id" class="bg-white p-4 rounded shadow">
        <h3 class="text-xl font-bold">{{ item.title }}</h3>
        <p class="text-gray-600">{{ item.description }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [...]
    }
  }
}
</script>
```

## Mejores prácticas

### 1. Elegir el framework adecuado

```javascript
// Evalúa tus necesidades
const requirements = {
  timeline: 'rápido', // rápido vs personalizado
  customization: 'media', // baja vs alta
  team: 'grande', // pequeño vs grande
  design: 'consistente' // consistente vs único
};
```

### 2. Optimizar el bundle

```javascript
// Tailwind - purging automático
module.exports = {
  purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
  // ...
}

// Bootstrap - importar solo lo necesario
import 'bootstrap/dist/css/bootstrap.min.css';
// O
import Button from 'bootstrap/js/dist/button';
```

### 3. Personalización consistente

```scss
// Variables globales
:root {
  --primary-color: #007bff;
  --secondary-color: #6c757d;
  --font-family: 'Inter', sans-serif;
}

// Tailwind config
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: 'var(--primary-color)',
        secondary: 'var(--secondary-color)',
      },
      fontFamily: {
        sans: 'var(--font-family)',
      }
    }
  }
}
```

### 4. Mantener actualizado

```bash
# Verificar versiones
npm outdated

# Actualizar
npm update bootstrap
# o
npm update tailwindcss
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Frameworks Comparison</title>
  
  <!-- Bootstrap -->
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
  
  <!-- Tailwind (CDN para demo) -->
  <script src="https://cdn.tailwindcss.com"></script>
  
  <!-- Bulma -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
</head>
<body>
  <!-- Bootstrap Example -->
  <section class="bg-light py-5">
    <div class="container">
      <h2 class="text-center mb-4">Bootstrap Components</h2>
      <div class="row">
        <div class="col-md-4">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">Bootstrap Card</h5>
              <p class="card-text">Framework completo con muchos componentes.</p>
              <button class="btn btn-primary">Bootstrap Button</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- Tailwind Example -->
  <section class="bg-gray-100 py-16">
    <div class="container mx-auto px-4">
      <h2 class="text-center text-3xl font-bold mb-8">Tailwind Components</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div class="bg-white p-6 rounded-lg shadow-md">
          <h3 class="text-xl font-semibold mb-2">Tailwind Card</h3>
          <p class="text-gray-600 mb-4">Utility-first framework altamente personalizable.</p>
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
            Tailwind Button
          </button>
        </div>
      </div>
    </div>
  </section>

  <!-- Bulma Example -->
  <section class="hero is-info">
    <div class="hero-body">
      <div class="container">
        <h2 class="title">Bulma Components</h2>
        <div class="columns">
          <div class="column">
            <div class="box">
              <h3 class="title is-4">Bulma Box</h3>
              <p>Framework moderno basado en Flexbox.</p>
              <button class="button is-primary">Bulma Button</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
```

## Resumen

Los frameworks CSS aceleran el desarrollo pero requieren elegir sabiamente:

- **Bootstrap**: Completo, maduro, ideal para proyectos rápidos
- **Tailwind**: Utility-first, ligero, máxima personalización
- **Bulma**: Simple, elegante, basado en Flexbox
- **Foundation**: Avanzado, profesional, máxima flexibilidad
- **Materialize**: Material Design, animaciones fluidas

Elige basado en:
- **Tamaño del proyecto**: Pequeño → Bulma, Grande → Bootstrap
- **Personalización**: Alta → Tailwind, Baja → Bootstrap
- **Experiencia del equipo**: Inexperto → Bootstrap, Experto → Tailwind
- **Diseño**: Consistente → Framework, Único → Tailwind

Recuerda: Los frameworks son herramientas, no soluciones universales. Aprende CSS puro primero, luego usa frameworks para ser más productivo.