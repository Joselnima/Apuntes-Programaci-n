# Módulo 15 - Optimización y Performance CSS

En este módulo aprenderás técnicas avanzadas de optimización CSS, estrategias para mejorar el rendimiento de renderizado, reducción de tamaño de archivos, y mejores prácticas para CSS en producción.

## Introducción a la Optimización CSS

### ¿Por qué optimizar CSS?

```css
/* ❌ CSS no optimizado */
body { font-family: Arial, sans-serif; margin: 0; padding: 0; }
.container { width: 100%; max-width: 1200px; margin: 0 auto; padding: 20px; }
.button { background: red; color: white; padding: 10px 20px; border: none; border-radius: 5px; }

/* ✅ CSS optimizado */
body{font-family:Arial,sans-serif;margin:0;padding:0}
.container{width:100%;max-width:1200px;margin:0 auto;padding:20px}
.button{background:red;color:#fff;padding:10px 20px;border:0;border-radius:5px}
```

### Métricas de rendimiento

- **First Contentful Paint (FCP)**: Tiempo hasta primer contenido
- **Largest Contentful Paint (LCP)**: Tiempo hasta elemento más grande
- **Cumulative Layout Shift (CLS)**: Cambio de layout acumulado
- **Time to Interactive (TTI)**: Tiempo hasta interactividad

## Minificación y Compresión

### Minificación manual

```css
/* Original */
.header {
  background-color: #ffffff;
  border: 1px solid #cccccc;
  margin-bottom: 20px;
  padding: 10px 15px;
}

/* Minificado */
.header{background-color:#fff;border:1px solid #ccc;margin-bottom:20px;padding:10px 15px}
```

### Herramientas de minificación

```bash
# CSSNano
npm install cssnano --save-dev

# PostCSS con CSSNano
postcss input.css -o output.css --use cssnano

# Opciones avanzadas
postcss input.css -o output.css --use cssnano --cssnano.preset advanced
```

### Configuración PostCSS

```javascript
// postcss.config.js
module.exports = {
  plugins: [
    require('cssnano')({
      preset: ['default', {
        discardComments: {
          removeAll: true,
        },
        normalizeWhitespace: false
      }]
    })
  ]
}
```

## Eliminación de CSS no utilizado (Purging)

### PurgeCSS

```bash
# Instalar PurgeCSS
npm install @fullhuman/postcss-purgecss --save-dev

# Configurar con PostCSS
postcss input.css -o output.css --use @fullhuman/postcss-purgecss
```

### Configuración PurgeCSS

```javascript
// postcss.config.js
module.exports = {
  plugins: [
    require('@fullhuman/postcss-purgecss')({
      content: [
        './src/**/*.html',
        './src/**/*.js',
        './src/**/*.jsx',
        './src/**/*.ts',
        './src/**/*.tsx',
        './src/**/*.vue',
      ],
      defaultExtractor: content => content.match(/[\w-/:]+(?<!:)/g) || [],
      safelist: [
        'active',
        'disabled',
        'focus',
        'hover',
        /^nav-/,
        /^btn-/,
      ]
    }),
    require('cssnano')
  ]
}
```

### Ejemplo de purgado

```html
<!-- index.html -->
<div class="container">
  <button class="btn btn-primary">Click me</button>
  <button class="btn btn-secondary">Secondary</button>
</div>
```

```css
/* input.css */
.container { max-width: 1200px; margin: 0 auto; }
.btn { padding: 10px 20px; border: none; border-radius: 4px; }
.btn-primary { background: blue; color: white; }
.btn-secondary { background: gray; color: white; }
.btn-danger { background: red; color: white; } /* No usado - será removido */
.card { background: white; padding: 20px; } /* No usado - será removido */
```

```css
/* output.css - Solo CSS usado */
.container{max-width:1200px;margin:0 auto}
.btn{padding:10px 20px;border:0;border-radius:4px}
.btn-primary{background:#00f;color:#fff}
.btn-secondary{background:#808080;color:#fff}
```

## Optimización del CSSOM (CSS Object Model)

### Reducir profundidad de selectores

```css
/* ❌ Selectores profundos */
.header .nav .menu .item .link { color: red; }

/* ✅ Selectores planos */
.c-link { color: red; }
.c-link--active { color: blue; }
```

### Evitar selectores universales

```css
/* ❌ Mal */
* { box-sizing: border-box; }
div * { margin: 0; }

/* ✅ Mejor */
html { box-sizing: border-box; }
*, *::before, *::after { box-sizing: inherit; }
```

### Optimización de especificidad

```css
/* ❌ Alta especificidad */
#main .content .article .title { font-size: 24px; }

/* ✅ Baja especificidad */
.c-article-title { font-size: 24px; }
```

## Optimización de Renderizado

### Contenido Above the Fold

```html
<!-- CSS crítico inline -->
<style>
  .hero { background: #007bff; color: white; padding: 2rem; text-align: center; }
  .hero h1 { font-size: 3rem; margin: 0; }
  .hero p { font-size: 1.2rem; margin: 1rem 0; }
  .btn { background: white; color: #007bff; padding: 1rem 2rem; text-decoration: none; border-radius: 4px; }
</style>

<div class="hero">
  <h1>Bienvenido</h1>
  <p>Contenido importante above the fold</p>
  <a href="#" class="btn">Comenzar</a>
</div>

<!-- Resto del CSS cargado después -->
<link rel="stylesheet" href="styles.css">
```

### Critical CSS

```javascript
// Extraer CSS crítico
const critical = require('critical');

critical.generate({
  inline: true,
  base: 'dist/',
  src: 'index.html',
  dest: 'index.html',
  minify: true,
  width: 1300,
  height: 900
});
```

### CSS no bloqueante

```html
<!-- CSS no crítico con media queries -->
<link rel="stylesheet" href="print.css" media="print">
<link rel="stylesheet" href="mobile.css" media="screen and (max-width: 768px)">
<link rel="stylesheet" href="desktop.css" media="screen and (min-width: 769px)">

<!-- CSS asíncrono -->
<link rel="preload" href="styles.css" as="style" onload="this.onload=null;this.rel='stylesheet'">
<noscript><link rel="stylesheet" href="styles.css"></noscript>
```

## Optimización de Imágenes y Assets

### Imágenes responsivas

```css
/* Imágenes fluidas */
img {
  max-width: 100%;
  height: auto;
}

/* Background images responsivas */
.hero {
  background-image: url('hero-small.jpg');
}

@media (min-width: 768px) {
  .hero {
    background-image: url('hero-medium.jpg');
  }
}

@media (min-width: 1200px) {
  .hero {
    background-image: url('hero-large.jpg');
  }
}
```

### WebP con fallbacks

```html
<picture>
  <source srcset="image.webp" type="image/webp">
  <img src="image.jpg" alt="Descripción">
</picture>
```

```css
/* CSS para WebP */
.webp .hero {
  background-image: url('hero.webp');
}

.no-webp .hero {
  background-image: url('hero.jpg');
}
```

### Optimización de fuentes

```css
/* Font loading optimizado */
@font-face {
  font-family: 'Inter';
  font-style: normal;
  font-weight: 400;
  font-display: swap; /* Mejora FCP */
  src: url('inter-regular.woff2') format('woff2'),
       url('inter-regular.woff') format('woff');
}

@font-face {
  font-family: 'Inter';
  font-style: normal;
  font-weight: 700;
  font-display: swap;
  src: url('inter-bold.woff2') format('woff2'),
       url('inter-bold.woff') format('woff');
}
```

## Optimización de Animaciones

### Usar transform y opacity

```css
/* ✅ Animaciones optimizadas */
.element {
  transform: translateX(0);
  opacity: 1;
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.element:hover {
  transform: translateX(10px);
  opacity: 0.8;
}

/* ❌ Animaciones no optimizadas */
.element {
  width: 100px;
  height: 100px;
  margin-left: 0;
  transition: width 0.3s ease, height 0.3s ease, margin-left 0.3s ease;
}

.element:hover {
  width: 120px;
  height: 120px;
  margin-left: 20px;
}
```

### will-change property

```css
/* Indicar propiedades que cambiarán */
.element {
  will-change: transform, opacity;
}

/* Reset después de la animación */
.element:not(:hover) {
  will-change: auto;
}
```

### Containment

```css
/* Aislar subárboles para mejor performance */
.expensive-operation {
  contain: layout style paint;
}
```

## Optimización de Layout

### Evitar Layout Thrashing

```javascript
// ❌ Mal - Múltiples lecturas/escrituras
elements.forEach(el => {
  const height = el.offsetHeight; // Lectura
  el.style.height = height + 'px'; // Escritura
  const width = el.offsetWidth; // Lectura
  el.style.width = width + 'px'; // Escritura
});

// ✅ Mejor - Agrupar lecturas y escrituras
elements.forEach(el => {
  const height = el.offsetHeight; // Lectura
  const width = el.offsetWidth; // Lectura
  // Escrituras agrupadas
  el.style.height = height + 'px';
  el.style.width = width + 'px';
});
```

### CSS Containment

```css
/* Contener layout, style, paint */
.component {
  contain: layout style paint;
}

/* Contener solo layout */
.component {
  contain: layout;
}

/* Contener paint */
.component {
  contain: paint;
}
```

## Optimización de Selectores

### Análisis de eficiencia de selectores

```css
/* ❌ Selectores ineficientes */
div > p:first-child { color: red; }
.header .nav .menu .item .link:hover { background: blue; }

/* ✅ Selectores eficientes */
.c-link { color: red; }
.c-link:hover { background: blue; }
```

### Reglas de eficiencia

1. **ID selectors**: Más rápidos (`#id`)
2. **Class selectors**: Rápidos (`.class`)
3. **Type selectors**: Moderados (`div`, `p`)
4. **Universal selectors**: Lentos (`*`)
5. **Attribute selectors**: Lentos (`[attr="value"]`)
6. **Pseudo-selectors**: Lentos (`:first-child`, `:nth-child`)

## Herramientas de Optimización

### CSS Stats

```javascript
// Analizar CSS
const cssstats = require('cssstats');

cssstats('styles.css').then(stats => {
  console.log('Reglas:', stats.rules.total);
  console.log('Selectores:', stats.selectors.total);
  console.log('Declarations:', stats.declarations.total);
  console.log('Tamaño gzipped:', stats.gzipSize);
});
```

### Lighthouse

```bash
# Ejecutar Lighthouse
lighthouse https://example.com --output html --output-path ./report.html

# Solo métricas de performance
lighthouse https://example.com --only-categories performance
```

### WebPageTest

```bash
# Prueba de velocidad
webpagetest test https://example.com --location 'Dulles:Chrome' --runs 3
```

## Optimización para Core Web Vitals

### First Contentful Paint (FCP)

```html
<!-- CSS crítico inline -->
<style>
  body { margin: 0; font-family: Arial, sans-serif; }
  .hero { background: #007bff; color: white; padding: 2rem; text-align: center; }
  .hero h1 { font-size: 3rem; margin: 0; }
</style>

<div class="hero">
  <h1>Contenido visible inmediatamente</h1>
</div>
```

### Largest Contentful Paint (LCP)

```css
/* Optimizar elementos LCP */
.hero-image {
  width: 100%;
  height: 400px;
  background: #f0f0f0;
  /* Placeholder mientras carga */
}

.hero-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
```

### Cumulative Layout Shift (CLS)

```css
/* Evitar CLS */
.card {
  min-height: 200px; /* Altura mínima */
  aspect-ratio: 16 / 9; /* Relación de aspecto */
}

.image-container {
  position: relative;
  width: 100%;
  padding-bottom: 56.25%; /* 16:9 aspect ratio */
}

.image-container img {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}
```

## Optimización de CSS Grid y Flexbox

### Grid implícito vs explícito

```css
/* ❌ Grid implícito ineficiente */
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

/* ✅ Grid explícito optimizado */
.grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1rem;
}

@media (max-width: 768px) {
  .grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 480px) {
  .grid {
    grid-template-columns: 1fr;
  }
}
```

### Flexbox optimizado

```css
/* ✅ Flexbox eficiente */
.flex-container {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.flex-item {
  flex: 1 1 200px; /* flex-grow, flex-shrink, flex-basis */
  min-width: 0; /* Evitar overflow */
}
```

## Optimización de Animaciones CSS

### Hardware acceleration

```css
/* Forzar aceleración por hardware */
.animated-element {
  transform: translateZ(0);
  backface-visibility: hidden;
  perspective: 1000px;
}
```

### Animaciones eficientes

```css
/* Usar transform y opacity */
@keyframes slideIn {
  from {
    transform: translateX(-100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.slide-in {
  animation: slideIn 0.3s ease-out;
}
```

## Monitoreo y Métricas

### CSS en producción

```javascript
// Monitorear CSS con Performance API
const observer = new PerformanceObserver((list) => {
  for (const entry of list.getEntries()) {
    if (entry.name.includes('.css')) {
      console.log('CSS loaded:', entry.name, entry.duration + 'ms');
    }
  }
});

observer.observe({ entryTypes: ['resource'] });
```

### Budgets de performance

```javascript
// package.json
{
  "bundlesize": [
    {
      "path": "./dist/css/*.css",
      "maxSize": "50 kB"
    }
  ]
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Performance Master</title>
  
  <!-- CSS crítico inline -->
  <style>
    body{margin:0;font-family:Arial,sans-serif;background:#f8f9fa}
    .hero{background:#007bff;color:#fff;padding:3rem;text-align:center;min-height:60vh;display:flex;flex-direction:column;justify-content:center}
    .hero h1{font-size:3rem;margin:0}
    .hero p{font-size:1.2rem;margin:1rem 0}
    .btn{background:#fff;color:#007bff;padding:1rem 2rem;text-decoration:none;border-radius:4px;display:inline-block;margin-top:1rem}
  </style>
  
  <!-- Preload CSS no crítico -->
  <link rel="preload" href="styles.css" as="style" onload="this.onload=null;this.rel='stylesheet'">
  <noscript><link rel="stylesheet" href="styles.css"></noscript>
</head>
<body>
  <div class="hero">
    <h1>Optimización CSS</h1>
    <p>Performance y mejores prácticas</p>
    <a href="#content" class="btn">Explorar</a>
  </div>

  <main id="content">
    <section class="section">
      <div class="container">
        <h2>Técnicas de Optimización</h2>
        <div class="grid">
          <div class="card">
            <h3>Minificación</h3>
            <p>Reducir tamaño eliminando espacios y comentarios.</p>
          </div>
          <div class="card">
            <h3>Purgado</h3>
            <p>Eliminar CSS no utilizado con herramientas como PurgeCSS.</p>
          </div>
          <div class="card">
            <h3>Compresión</h3>
            <p>Gzip y Brotli para reducir transferencia de datos.</p>
          </div>
          <div class="card">
            <h3>Critical CSS</h3>
            <p>Cargar CSS crítico inline para mejor FCP.</p>
          </div>
        </div>
      </section>
    </main>

  <script>
    // Lazy load imágenes
    const images = document.querySelectorAll('img[data-src]');
    const imageObserver = new IntersectionObserver((entries, observer) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          const img = entry.target;
          img.src = img.dataset.src;
          img.classList.remove('lazy');
          observer.unobserve(img);
        }
      });
    });

    images.forEach(img => imageObserver.observe(img));
  </script>
</body>
</html>
```

```css
/* styles.css - Optimizado */
:root{--primary:#007bff;--secondary:#6c757d;--light:#f8f9fa;--dark:#343a40;--shadow:0 2px 8px rgba(0,0,0,.1)}

.container{max-width:1200px;margin:0 auto;padding:0 1rem}
.section{padding:3rem 0}
.grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(250px,1fr));gap:1rem}
.card{background:#fff;border-radius:8px;box-shadow:var(--shadow);padding:1.5rem;overflow:hidden;contain:layout style paint}
.card h3{margin:0 0 1rem 0;color:var(--dark)}
.card p{margin:0;color:#666;line-height:1.6}

/* Animaciones optimizadas */
@keyframes fadeIn{from{opacity:0;transform:translateY(20px)}to{opacity:1;transform:translateY(0)}}
.fade-in{animation:fadeIn .6s ease-out}

/* Responsive */
@media(max-width:768px){
  .grid{grid-template-columns:1fr}
  .hero{padding:2rem 1rem}
  .hero h1{font-size:2rem}
}

/* Lazy loading */
img.lazy{opacity:0;transition:opacity .3s}
img.lazy.loaded{opacity:1}
```

```javascript
// build.js - Proceso de optimización
const fs = require('fs');
const postcss = require('postcss');
const cssnano = require('cssnano');
const purgecss = require('@fullhuman/postcss-purgecss');

async function optimizeCSS() {
  const inputCSS = fs.readFileSync('src/styles.css', 'utf8');
  
  const result = await postcss([
    purgecss({
      content: ['src/**/*.html', 'src/**/*.js'],
      safelist: ['active', 'fade-in', /^btn/]
    }),
    cssnano({
      preset: 'advanced'
    })
  ]).process(inputCSS, { from: 'src/styles.css', to: 'dist/styles.css' });
  
  fs.writeFileSync('dist/styles.css', result.css);
  
  console.log('CSS optimizado:');
  console.log(`Tamaño original: ${inputCSS.length} bytes`);
  console.log(`Tamaño optimizado: ${result.css.length} bytes`);
  console.log(`Reducción: ${((inputCSS.length - result.css.length) / inputCSS.length * 100).toFixed(1)}%`);
}

optimizeCSS();
```

## Resumen

La optimización CSS es crucial para el rendimiento web:

- **Minificación**: Reducir tamaño eliminando caracteres innecesarios
- **Purgado**: Eliminar CSS no utilizado
- **Critical CSS**: Cargar CSS crítico primero
- **Compresión**: Gzip/Brotli para transferencia
- **Selectores eficientes**: Reducir profundidad y complejidad
- **Animaciones optimizadas**: Usar transform y opacity
- **Containment**: Aislar componentes para mejor performance

Herramientas como CSSNano, PurgeCSS, Critical, y Lighthouse son esenciales para mantener CSS optimizado. Monitorea Core Web Vitals y establece budgets de performance.