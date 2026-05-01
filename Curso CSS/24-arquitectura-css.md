# Módulo 14 - Arquitectura CSS

En este módulo aprenderás sobre arquitecturas CSS, metodologías y mejores prácticas para organizar y mantener código CSS a gran escala. Cubriremos BEM, SMACSS, OOCSS y otras estrategias profesionales.

## Introducción a la Arquitectura CSS

### ¿Por qué necesitamos arquitectura CSS?

```css
/* ❌ Código CSS sin estructura */
.header .nav .item .link { color: red; }
.footer .nav .item .link { color: red; }
.sidebar .nav .item .link { color: red; }

/* ✅ Código CSS con arquitectura */
.c-link { color: red; }
.c-link--active { color: blue; }
```

### Beneficios de una buena arquitectura

- **Mantenibilidad**: Fácil de modificar y actualizar
- **Escalabilidad**: Crece con el proyecto
- **Reutilización**: Componentes reutilizables
- **Colaboración**: Múltiples desarrolladores trabajando eficientemente
- **Performance**: CSS optimizado y sin redundancias

## Metodología BEM (Block Element Modifier)

### Conceptos básicos de BEM

```css
/* Block: Componente independiente */
.block { }

/* Element: Parte del block */
.block__element { }

/* Modifier: Variación del block o element */
.block--modifier { }
.block__element--modifier { }
```

### Ejemplo práctico de BEM

```html
<!-- HTML con clases BEM -->
<header class="header">
  <nav class="nav">
    <ul class="nav__list">
      <li class="nav__item">
        <a href="#" class="nav__link nav__link--active">Inicio</a>
      </li>
      <li class="nav__item">
        <a href="#" class="nav__link">Acerca</a>
      </li>
      <li class="nav__item">
        <a href="#" class="nav__link nav__link--disabled">Contacto</a>
      </li>
    </ul>
  </nav>
  <button class="button button--primary">Login</button>
</header>
```

```css
/* CSS con BEM */
.header {
  background: #f8f9fa;
  padding: 1rem;
}

.nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav__list {
  display: flex;
  list-style: none;
  gap: 2rem;
}

.nav__item {
  /* Estilos base del item */
}

.nav__link {
  color: #333;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.nav__link:hover {
  background: #e9ecef;
}

.nav__link--active {
  background: #007bff;
  color: white;
}

.nav__link--disabled {
  color: #6c757d;
  cursor: not-allowed;
}

.nav__link--disabled:hover {
  background: transparent;
}

.button {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.button--primary {
  background: #007bff;
  color: white;
}

.button--primary:hover {
  background: #0056b3;
}
```

### Ventajas de BEM

- **Claridad**: Nombres descriptivos que explican la función
- **Predecibilidad**: Estructura consistente y fácil de entender
- **Modularidad**: Componentes independientes
- **Escalabilidad**: Fácil agregar nuevas variaciones
- **Mantenimiento**: Cambios localizados

### Casos de uso avanzados

```css
/* Componente Card con BEM */
.card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  overflow: hidden;
}

.card__header {
  padding: 1.5rem;
  border-bottom: 1px solid #e9ecef;
}

.card__title {
  margin: 0;
  font-size: 1.25rem;
  color: #333;
}

.card__subtitle {
  margin: 0.5rem 0 0 0;
  font-size: 0.875rem;
  color: #6c757d;
}

.card__body {
  padding: 1.5rem;
}

.card__footer {
  padding: 1rem 1.5rem;
  background: #f8f9fa;
  border-top: 1px solid #e9ecef;
}

.card--featured {
  border: 2px solid #007bff;
  box-shadow: 0 4px 16px rgba(0,123,255,0.2);
}

.card--compact .card__body {
  padding: 1rem;
}

.card--compact .card__header,
.card--compact .card__footer {
  padding: 0.75rem 1rem;
}
```

## SMACSS (Scalable and Modular Architecture for CSS)

### Categorías en SMACSS

```css
/* 1. Base - Estilos por defecto */
body, h1, h2, h3, p, a { /* ... */ }

/* 2. Layout - Estilos de layout */
.l-header { /* ... */ }
.l-sidebar { /* ... */ }
.l-main { /* ... */ }

/* 3. Module - Componentes reutilizables */
.button { /* ... */ }
.card { /* ... */ }

/* 4. State - Estados de los componentes */
.is-active { /* ... */ }
.is-hidden { /* ... */ }

/* 5. Theme - Temas y skins */
.theme-dark { /* ... */ }
```

### Ejemplo completo SMACSS

```css
/* base.css */
* {
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', sans-serif;
  line-height: 1.6;
  color: #333;
  background: #fff;
}

h1, h2, h3, h4, h5, h6 {
  margin: 0 0 1rem 0;
  font-weight: 600;
}

a {
  color: #007bff;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}
```

```css
/* layout.css */
.l-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.l-header {
  background: #fff;
  border-bottom: 1px solid #e9ecef;
  padding: 1rem 0;
}

.l-sidebar {
  background: #f8f9fa;
  padding: 2rem;
}

.l-main {
  padding: 2rem;
}

.l-footer {
  background: #343a40;
  color: white;
  padding: 2rem 0;
  text-align: center;
}
```

```css
/* modules.css */
.button {
  display: inline-block;
  padding: 0.75rem 1.5rem;
  border: 1px solid transparent;
  border-radius: 4px;
  text-align: center;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.button-primary {
  background: #007bff;
  border-color: #007bff;
  color: white;
}

.button-secondary {
  background: #6c757d;
  border-color: #6c757d;
  color: white;
}

.card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  overflow: hidden;
}

.card-header {
  padding: 1.5rem;
  border-bottom: 1px solid #e9ecef;
}

.card-body {
  padding: 1.5rem;
}

.card-footer {
  padding: 1rem 1.5rem;
  background: #f8f9fa;
  border-top: 1px solid #e9ecef;
}
```

```css
/* states.css */
.is-hidden {
  display: none !important;
}

.is-visible {
  display: block !important;
}

.is-active {
  background: #007bff !important;
  color: white !important;
}

.is-disabled {
  opacity: 0.6;
  pointer-events: none;
}

.is-loading {
  position: relative;
}

.is-loading::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 20px;
  height: 20px;
  margin: -10px 0 0 -10px;
  border: 2px solid #f3f3f3;
  border-top: 2px solid #007bff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
```

```css
/* theme.css */
.theme-dark {
  --bg-color: #2c3e50;
  --text-color: #ecf0f1;
  --border-color: #34495e;
}

.theme-dark body {
  background: var(--bg-color);
  color: var(--text-color);
}

.theme-dark .card {
  background: #34495e;
  color: var(--text-color);
}

.theme-dark .button-primary {
  background: #3498db;
  border-color: #3498db;
}
```

## OOCSS (Object Oriented CSS)

### Principios de OOCSS

```css
/* Separar estructura de apariencia */
.skin {
  border: 1px solid;
  background: white;
}

.skin-dark {
  border-color: #333;
  background: #333;
  color: white;
}

.size {
  width: 100px;
  height: 100px;
}

.size-large {
  width: 200px;
  height: 200px;
}

/* Combinar clases */
<div class="skin skin-dark size size-large"></div>
```

### Ejemplo práctico OOCSS

```css
/* Objetos base */
.media {
  overflow: hidden;
  margin-bottom: 1rem;
}

.media__img {
  float: left;
  margin-right: 1rem;
}

.media__body {
  overflow: hidden;
}

/* Extensiones */
.media--small .media__img {
  width: 64px;
  height: 64px;
}

.media--large .media__img {
  width: 128px;
  height: 128px;
}

/* Skins */
.media--bordered {
  border: 1px solid #e9ecef;
  border-radius: 4px;
  padding: 1rem;
}

.media--shadowed {
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}
```

## ITCSS (Inverted Triangle CSS)

### Arquitectura ITCSS

```css
/* 1. Settings - Variables */
@import "settings.colors";
@import "settings.typography";

/* 2. Tools - Mixins y funciones */
@import "tools.mixins";
@import "tools.functions";

/* 3. Generic - Resets y normalize */
@import "generic.reset";
@import "generic.normalize";

/* 4. Elements - Estilos base de elementos HTML */
@import "elements.headings";
@import "elements.links";

/* 5. Objects - Objetos OOCSS */
@import "objects.media";
@import "objects.flag";

/* 6. Components - Componentes específicos */
@import "components.button";
@import "components.card";

/* 7. Utilities - Utilidades y helpers */
@import "utilities.spacing";
@import "utilities.colors";
```

### Implementación ITCSS

```css
/* settings/_colors.scss */
$color-primary: #007bff;
$color-secondary: #6c757d;
$color-success: #28a745;
$color-danger: #dc3545;
$color-warning: #ffc107;

/* settings/_typography.scss */
$font-family-base: 'Segoe UI', sans-serif;
$font-size-base: 1rem;
$line-height-base: 1.6;

/* tools/_mixins.scss */
@mixin button-variant($color) {
  background: $color;
  border: 1px solid $color;
  color: white;
  
  &:hover {
    background: darken($color, 10%);
    border-color: darken($color, 10%);
  }
}

/* generic/_reset.scss */
*,
*::before,
*::after {
  box-sizing: border-box;
}

* {
  margin: 0;
}

/* elements/_headings.scss */
h1, h2, h3, h4, h5, h6 {
  margin-bottom: 0.5rem;
  font-weight: 600;
  line-height: 1.3;
}

h1 { font-size: 2.5rem; }
h2 { font-size: 2rem; }
h3 { font-size: 1.75rem; }
h4 { font-size: 1.5rem; }
h5 { font-size: 1.25rem; }
h6 { font-size: 1rem; }

/* objects/_media.scss */
.o-media {
  display: flex;
  align-items: flex-start;
}

.o-media__img {
  margin-right: 1rem;
}

.o-media__body {
  flex: 1;
}

/* components/_button.scss */
.c-button {
  display: inline-block;
  padding: 0.75rem 1.5rem;
  border: 1px solid transparent;
  border-radius: 4px;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.c-button--primary {
  @include button-variant($color-primary);
}

.c-button--secondary {
  @include button-variant($color-secondary);
}

/* utilities/_spacing.scss */
.u-m-0 { margin: 0 !important; }
.u-m-1 { margin: 0.25rem !important; }
.u-m-2 { margin: 0.5rem !important; }
.u-m-3 { margin: 1rem !important; }
.u-m-4 { margin: 1.5rem !important; }
.u-m-5 { margin: 3rem !important; }

.u-p-0 { padding: 0 !important; }
.u-p-1 { padding: 0.25rem !important; }
.u-p-2 { padding: 0.5rem !important; }
.u-p-3 { padding: 1rem !important; }
.u-p-4 { padding: 1.5rem !important; }
.u-p-5 { padding: 3rem !important; }
```

## Atomic Design

### Conceptos de Atomic Design

```css
/* Atoms - Elementos más pequeños */
.a-button {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
}

.a-input {
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

/* Molecules - Grupos de atoms */
.m-form-field {
  margin-bottom: 1rem;
}

.m-form-field label {
  display: block;
  margin-bottom: 0.5rem;
}

/* Organisms - Grupos complejos de molecules */
.o-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: #f8f9fa;
}

.o-navigation {
  display: flex;
  gap: 1rem;
}

/* Templates - Instancias de organisms */
.t-dashboard {
  display: grid;
  grid-template-columns: 1fr 3fr;
  gap: 2rem;
  min-height: 100vh;
}

/* Pages - Templates con contenido real */
.p-home {
  /* Estilos específicos de página */
}
```

## CSS Modules

### Conceptos básicos

```css
/* button.module.css */
.button {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.primary {
  background: #007bff;
  color: white;
}

.secondary {
  background: #6c757d;
  color: white;
}
```

```javascript
// React con CSS Modules
import styles from './button.module.css';

function Button({ variant = 'primary', children }) {
  return (
    <button className={styles[variant]}>
      {children}
    </button>
  );
}
```

## Estrategias de organización de archivos

### Estructura de carpetas recomendada

```
css/
├── abstracts/
│   ├── _variables.scss
│   ├── _mixins.scss
│   ├── _functions.scss
│   └── _placeholders.scss
├── base/
│   ├── _reset.scss
│   ├── _typography.scss
│   └── _base.scss
├── components/
│   ├── _buttons.scss
│   ├── _cards.scss
│   ├── _forms.scss
│   ├── _modals.scss
│   └── _navigation.scss
├── layout/
│   ├── _header.scss
│   ├── _footer.scss
│   ├── _sidebar.scss
│   └── _grid.scss
├── pages/
│   ├── _home.scss
│   ├── _about.scss
│   └── _contact.scss
├── themes/
│   ├── _default.scss
│   └── _dark.scss
└── utilities/
    ├── _spacing.scss
    ├── _colors.scss
    ├── _typography.scss
    └── _visibility.scss
```

### Archivo principal

```scss
// main.scss
@import 'abstracts/variables';
@import 'abstracts/mixins';
@import 'abstracts/functions';

@import 'base/reset';
@import 'base/typography';
@import 'base/base';

@import 'layout/header';
@import 'layout/footer';
@import 'layout/sidebar';
@import 'layout/grid';

@import 'components/buttons';
@import 'components/cards';
@import 'components/forms';
@import 'components/modals';
@import 'components/navigation';

@import 'pages/home';
@import 'pages/about';
@import 'pages/contact';

@import 'themes/default';

@import 'utilities/spacing';
@import 'utilities/colors';
@import 'utilities/typography';
@import 'utilities/visibility';
```

## Mejores prácticas

### Convenciones de nomenclatura

```css
/* ✅ Prefijos descriptivos */
.c-card { /* component */ }
.l-header { /* layout */ }
.u-text-center { /* utility */ }
.js-modal { /* javascript hook */ }

/* ✅ Nombres en inglés */
.button-primary { /* ✅ */ }
.boton-primario { /* ❌ */ }

/* ✅ Nombres kebab-case para CSS */
.nav-bar { /* ✅ */ }
.navbar { /* ❌ */ }

/* ✅ Evitar IDs para estilos */
#sidebar { /* ❌ */ }
.sidebar { /* ✅ */ }
```

### Especificidad controlada

```css
/* ❌ Alta especificidad */
.header .nav .item .link.active { /* ... */ }

/* ✅ Baja especificidad */
.c-link { /* ... */ }
.c-link--active { /* ... */ }
```

### Documentación

```css
/**
 * Componente Button
 * 
 * @param {string} variant - Tipo de botón (primary, secondary, danger)
 * @param {string} size - Tamaño del botón (small, medium, large)
 * @param {boolean} disabled - Estado deshabilitado
 * 
 * @example
 * <button class="c-button c-button--primary c-button--large">Click me</button>
 */
.c-button {
  /* ... */
}
```

## Herramientas y automatización

### Linters y formatters

```json
// .stylelintrc
{
  "extends": ["stylelint-config-standard"],
  "rules": {
    "selector-class-pattern": "^[a-z][a-zA-Z0-9]*(-[a-z][a-zA-Z0-9]*)*$",
    "max-nesting-depth": 3,
    "selector-max-specificity": "0,3,0"
  }
}
```

### CSS Custom Properties (Variables CSS)

```css
:root {
  --color-primary: #007bff;
  --color-secondary: #6c757d;
  --spacing-small: 0.5rem;
  --spacing-medium: 1rem;
  --spacing-large: 2rem;
}

.theme-dark {
  --color-primary: #3498db;
  --color-secondary: #95a5a6;
}

.button {
  background: var(--color-primary);
  padding: var(--spacing-medium);
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Architecture Master</title>
  <link rel="stylesheet" href="dist/css/main.css">
</head>
<body>
  <!-- Layout con SMACSS -->
  <div class="l-container">
    <header class="l-header">
      <h1>Arquitectura CSS</h1>
      <nav class="c-navigation">
        <ul class="c-navigation__list">
          <li class="c-navigation__item">
            <a href="#" class="c-navigation__link is-active">Inicio</a>
          </li>
          <li class="c-navigation__item">
            <a href="#" class="c-navigation__link">BEM</a>
          </li>
          <li class="c-navigation__item">
            <a href="#" class="c-navigation__link">SMACSS</a>
          </li>
        </ul>
      </nav>
    </header>

    <main class="l-main">
      <!-- Componentes con BEM -->
      <section class="section">
        <h2>Ejemplos de Arquitectura</h2>
        
        <div class="grid">
          <div class="card card--featured">
            <div class="card__header">
              <h3 class="card__title">BEM Methodology</h3>
              <p class="card__subtitle">Block Element Modifier</p>
            </div>
            <div class="card__body">
              <p>Metodología para nombrar clases CSS de forma consistente y escalable.</p>
              <button class="button button--primary">Leer más</button>
            </div>
          </div>

          <div class="card">
            <div class="card__header">
              <h3 class="card__title">SMACSS</h3>
              <p class="card__subtitle">Scalable and Modular Architecture</p>
            </div>
            <div class="card__body">
              <p>Arquitectura que categoriza CSS en 5 tipos: Base, Layout, Module, State, Theme.</p>
              <button class="button button--secondary">Explorar</button>
            </div>
          </div>

          <div class="card">
            <div class="card__header">
              <h3 class="card__title">ITCSS</h3>
              <p class="card__subtitle">Inverted Triangle CSS</p>
            </div>
            <div class="card__body">
              <p>Triángulo invertido que organiza CSS desde genérico hasta específico.</p>
              <button class="button button--success">Descubrir</button>
            </div>
          </div>
        </div>
      </section>

      <!-- Media Object con OOCSS -->
      <section class="section">
        <h2>Media Objects</h2>
        <div class="media media--bordered">
          <img src="avatar.jpg" alt="Avatar" class="media__img">
          <div class="media__body">
            <h4>John Doe</h4>
            <p>Desarrollador frontend apasionado por la arquitectura CSS y las mejores prácticas.</p>
          </div>
        </div>
      </section>
    </main>

    <footer class="l-footer">
      <p>&copy; 2024 CSS Architecture Course</p>
    </footer>
  </div>

  <!-- Modal con estado -->
  <div class="modal is-hidden" id="modal">
    <div class="modal__overlay"></div>
    <div class="modal__content">
      <h3>Modal Title</h3>
      <p>Contenido del modal con arquitectura CSS.</p>
      <button class="button button--primary">Cerrar</button>
    </div>
  </div>
</body>
</html>
```

```css
/* main.css - Arquitectura completa */
:root {
  --color-primary: #007bff;
  --color-secondary: #6c757d;
  --color-success: #28a745;
  --color-light: #f8f9fa;
  --color-dark: #343a40;
  --border-radius: 4px;
  --box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

/* Base */
* {
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', sans-serif;
  line-height: 1.6;
  color: #333;
  background: var(--color-light);
}

/* Layout */
.l-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.l-header {
  background: white;
  border-bottom: 1px solid #e9ecef;
  padding: 1rem 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.l-main {
  padding: 2rem 0;
}

.l-footer {
  background: var(--color-dark);
  color: white;
  padding: 2rem 0;
  text-align: center;
}

/* Components - BEM */
.button {
  display: inline-block;
  padding: 0.75rem 1.5rem;
  border: 1px solid transparent;
  border-radius: var(--border-radius);
  text-decoration: none;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.button--primary {
  background: var(--color-primary);
  border-color: var(--color-primary);
  color: white;
}

.button--primary:hover {
  background: #0056b3;
  border-color: #0056b3;
}

.button--secondary {
  background: var(--color-secondary);
  border-color: var(--color-secondary);
  color: white;
}

.button--success {
  background: var(--color-success);
  border-color: var(--color-success);
  color: white;
}

.card {
  background: white;
  border-radius: 8px;
  box-shadow: var(--box-shadow);
  overflow: hidden;
  margin-bottom: 1rem;
}

.card--featured {
  border: 2px solid var(--color-primary);
  box-shadow: 0 4px 16px rgba(0,123,255,0.2);
}

.card__header {
  padding: 1.5rem;
  border-bottom: 1px solid #e9ecef;
}

.card__title {
  margin: 0;
  font-size: 1.25rem;
  color: #333;
}

.card__subtitle {
  margin: 0.5rem 0 0 0;
  font-size: 0.875rem;
  color: #6c757d;
}

.card__body {
  padding: 1.5rem;
}

.c-navigation {
  display: flex;
}

.c-navigation__list {
  display: flex;
  list-style: none;
  gap: 2rem;
}

.c-navigation__item {
  /* Estilos base */
}

.c-navigation__link {
  color: #333;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: var(--border-radius);
  transition: all 0.3s ease;
}

.c-navigation__link:hover {
  background: #e9ecef;
}

/* Objects - OOCSS */
.media {
  display: flex;
  align-items: flex-start;
  margin-bottom: 1rem;
}

.media--bordered {
  border: 1px solid #e9ecef;
  border-radius: var(--border-radius);
  padding: 1rem;
}

.media__img {
  margin-right: 1rem;
  width: 64px;
  height: 64px;
  border-radius: 50%;
}

.media__body {
  flex: 1;
}

/* Utilities */
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
}

.section {
  margin-bottom: 3rem;
}

/* States */
.is-hidden {
  display: none !important;
}

.is-active {
  background: var(--color-primary) !important;
  color: white !important;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1000;
}

.modal__overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.5);
}

.modal__content {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: white;
  padding: 2rem;
  border-radius: var(--border-radius);
  box-shadow: var(--box-shadow);
  max-width: 500px;
  width: 90%;
}
```

## Resumen

La arquitectura CSS es fundamental para proyectos escalables:

- **BEM**: Nomenclatura clara y consistente
- **SMACSS**: Categorización lógica del CSS
- **OOCSS**: Separación de estructura y apariencia
- **ITCSS**: Organización triangular del código
- **Atomic Design**: Componentes desde átomos hasta páginas

Implementar una arquitectura sólida desde el inicio salva tiempo y dolores de cabeza a largo plazo. Elige la metodología que mejor se adapte a tu equipo y proyecto.