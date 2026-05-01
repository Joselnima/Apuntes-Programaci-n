# Módulo 13 - Preprocesadores CSS

En este módulo aprenderás sobre preprocesadores CSS, herramientas que extienden las capacidades de CSS con variables, funciones, mixins y más. Nos enfocaremos en Sass/SCSS, el preprocesador más popular.

## Introducción a Sass/SCSS

### ¿Qué es un preprocesador CSS?

```scss
// Los preprocesadores toman código como este:
$primary-color: #3498db;
$font-size-base: 16px;

.button {
  background: $primary-color;
  font-size: $font-size-base;
  
  &:hover {
    background: darken($primary-color, 10%);
  }
}

// Y lo convierten en CSS estándar:
.button {
  background: #3498db;
  font-size: 16px;
}

.button:hover {
  background: #2980b9;
}
```

### Instalación y uso

```bash
# Instalar Sass
npm install -g sass

# Compilar archivo
sass input.scss output.css

# Compilar y vigilar cambios
sass --watch input.scss output.css

# Compilar directorio completo
sass --watch src/scss:dist/css
```

## Variables

### Declaración de variables

```scss
// Variables
$primary-color: #3498db;
$secondary-color: #2ecc71;
$font-stack: 'Helvetica Neue', Helvetica, Arial, sans-serif;
$border-radius: 4px;

// Uso
.button {
  background-color: $primary-color;
  color: white;
  font-family: $font-stack;
  border-radius: $border-radius;
}
```

### Scope de variables

```scss
$global-var: 16px; // Variable global

.container {
  $local-var: 20px; // Variable local
  font-size: $global-var;
  padding: $local-var;
  
  .nested {
    // $local-var está disponible aquí
    margin: $local-var;
  }
}

// $local-var no está disponible aquí
```

### Variables con !default

```scss
// En una librería o framework
$primary-color: #3498db !default;
$font-size: 16px !default;

// El usuario puede sobreescribir
$primary-color: #e74c3c; // Esto gana

// Uso
.button {
  background: $primary-color; // #e74c3c
}
```

## Anidamiento (Nesting)

### Anidamiento básico

```scss
.navbar {
  background: #333;
  
  .logo {
    float: left;
    font-size: 1.5em;
  }
  
  ul {
    float: right;
    
    li {
      display: inline-block;
      
      a {
        color: white;
        text-decoration: none;
        
        &:hover {
          color: #3498db;
        }
      }
    }
  }
}
```

### Compila a:

```css
.navbar {
  background: #333;
}

.navbar .logo {
  float: left;
  font-size: 1.5em;
}

.navbar ul {
  float: right;
}

.navbar ul li {
  display: inline-block;
}

.navbar ul li a {
  color: white;
  text-decoration: none;
}

.navbar ul li a:hover {
  color: #3498db;
}
```

### Anidamiento de propiedades

```scss
.button {
  font: {
    family: Arial, sans-serif;
    size: 1.2em;
    weight: bold;
  }
  
  border: {
    width: 2px;
    style: solid;
    color: #ccc;
    radius: 4px;
  }
}
```

## Partials y Imports

### Creando partials

```scss
// _variables.scss
$primary: #3498db;
$spacing: 1rem;

// _mixins.scss
@mixin button($color: $primary) {
  background: $color;
  // ...
}

// _buttons.scss
@import 'variables';
@import 'mixins';

.button {
  @include button();
}
```

### Import con @import

```scss
// main.scss
@import 'variables';
@import 'mixins';
@import 'buttons';
@import 'layout';

// También puedes importar archivos CSS normales
@import 'normalize.css';
```

### Import con @use (Sass moderno)

```scss
// _variables.scss
$primary: #3498db;

// main.scss
@use 'variables';

.button {
  background: variables.$primary;
}
```

## Mixins

### Mixins básicos

```scss
@mixin border-radius($radius: 5px) {
  -webkit-border-radius: $radius;
  -moz-border-radius: $radius;
  border-radius: $radius;
}

.button {
  @include border-radius(10px);
}

.card {
  @include border-radius(); // Usa valor por defecto
}
```

### Mixins con contenido

```scss
@mixin media-query($breakpoint) {
  @if $breakpoint == small {
    @media (max-width: 768px) { @content; }
  } @else if $breakpoint == medium {
    @media (min-width: 769px) and (max-width: 1024px) { @content; }
  } @else if $breakpoint == large {
    @media (min-width: 1025px) { @content; }
  }
}

.container {
  width: 100%;
  
  @include media-query(small) {
    padding: 10px;
  }
  
  @include media-query(medium) {
    width: 750px;
    margin: 0 auto;
  }
  
  @include media-query(large) {
    width: 1170px;
  }
}
```

### Mixins avanzados

```scss
@mixin flex-center {
  display: flex;
  justify-content: center;
  align-items: center;
}

@mixin button-variant($color, $hover-color: darken($color, 10%)) {
  background: $color;
  border: 1px solid $color;
  
  &:hover {
    background: $hover-color;
    border-color: $hover-color;
  }
}

.center-box {
  @include flex-center;
  width: 200px;
  height: 200px;
}

.primary-btn {
  @include button-variant(#3498db);
}

.danger-btn {
  @include button-variant(#e74c3c, #c0392b);
}
```

## Funciones

### Funciones personalizadas

```scss
@function px-to-rem($px, $base: 16px) {
  @return ($px / $base) * 1rem;
}

@function color-luminance($color) {
  $red: red($color);
  $green: green($color);
  $blue: blue($color);
  
  @return (0.299 * $red + 0.587 * $green + 0.114 * $blue) / 255;
}

@function contrast-color($color) {
  @if color-luminance($color) > 0.5 {
    @return #000;
  } @else {
    @return #fff;
  }
}

.button {
  font-size: px-to-rem(24px);
  background: #3498db;
  color: contrast-color(#3498db); // Retorna blanco
}
```

### Funciones integradas útiles

```scss
.element {
  // Colores
  background: lighten(#3498db, 20%);    // Más claro
  border: darken(#3498db, 10%);         // Más oscuro
  color: saturate(#3498db, 20%);        // Más saturado
  opacity: alpha(rgba(52, 152, 219, 0.5)); // Extraer alpha
  
  // Números
  width: round(13.7px);                 // 14px
  margin: ceil(13.1px);                 // 14px
  padding: floor(13.9px);               // 13px
  height: abs(-10px);                   // 10px
  
  // Strings
  content: quote("Hello World");        // "Hello World"
  font-family: unquote("Arial, sans-serif"); // Arial, sans-serif
  
  // Listas
  padding: nth(10px 20px 30px 40px, 2); // 20px
}
```

## Estructuras de control

### @if / @else

```scss
@mixin button-size($size) {
  @if $size == small {
    padding: 8px 16px;
    font-size: 0.875rem;
  } @else if $size == large {
    padding: 16px 32px;
    font-size: 1.25rem;
  } @else {
    padding: 12px 24px;
    font-size: 1rem;
  }
}

.small-btn {
  @include button-size(small);
}

.large-btn {
  @include button-size(large);
}
```

### @for

```scss
@for $i from 1 through 3 {
  .col-#{$i} {
    width: 100% / $i;
  }
}

// Genera:
// .col-1 { width: 100%; }
// .col-2 { width: 50%; }
// .col-3 { width: 33.333%; }
```

### @each

```scss
$colors: (
  primary: #3498db,
  secondary: #2ecc71,
  danger: #e74c3c,
  warning: #f39c12
);

@each $name, $color in $colors {
  .text-#{$name} {
    color: $color;
  }
  
  .bg-#{$name} {
    background: $color;
  }
}
```

### @while

```scss
$i: 1;
@while $i <= 5 {
  .mt-#{$i} {
    margin-top: #{$i}rem;
  }
  $i: $i + 1;
}
```

## Maps

### Trabajando con maps

```scss
$breakpoints: (
  small: 576px,
  medium: 768px,
  large: 992px,
  xlarge: 1200px
);

$colors: (
  primary: (
    base: #3498db,
    light: lighten(#3498db, 10%),
    dark: darken(#3498db, 10%)
  ),
  secondary: (
    base: #2ecc71,
    light: lighten(#2ecc71, 10%),
    dark: darken(#2ecc71, 10%)
  )
);

// Acceder a valores
.primary-color {
  color: map-get($colors, primary, base);
}

// Funciones útiles con maps
@function breakpoint($name) {
  @return map-get($breakpoints, $name);
}

@media (min-width: breakpoint(medium)) {
  .container {
    width: 750px;
  }
}
```

## Arquitectura y organización

### Patrón 7-1

```
scss/
├── abstracts/
│   ├── _variables.scss
│   ├── _functions.scss
│   ├── _mixins.scss
│   └── _placeholders.scss
├── base/
│   ├── _reset.scss
│   ├── _typography.scss
│   └── _base.scss
├── components/
│   ├── _buttons.scss
│   ├── _cards.scss
│   ├── _modals.scss
│   └── _forms.scss
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
└── main.scss
```

### Archivo main.scss

```scss
// Abstracts
@import 'abstracts/variables';
@import 'abstracts/functions';
@import 'abstracts/mixins';

// Base
@import 'base/reset';
@import 'base/typography';
@import 'base/base';

// Layout
@import 'layout/header';
@import 'layout/footer';
@import 'layout/sidebar';
@import 'layout/grid';

// Components
@import 'components/buttons';
@import 'components/cards';
@import 'components/modals';
@import 'components/forms';

// Pages
@import 'pages/home';
@import 'pages/about';
@import 'pages/contact';

// Themes
@import 'themes/default';
```

## Placeholders y extends

### Placeholders

```scss
%button-base {
  display: inline-block;
  padding: 12px 24px;
  text-decoration: none;
  border-radius: 4px;
  transition: all 0.3s ease;
}

%button-primary {
  @extend %button-base;
  background: $primary-color;
  color: white;
}

%button-secondary {
  @extend %button-base;
  background: $secondary-color;
  color: white;
}

.primary-btn {
  @extend %button-primary;
}

.secondary-btn {
  @extend %button-secondary;
}
```

## Optimización y mejores prácticas

### Compresión y sourcemaps

```bash
# Desarrollo (con sourcemaps)
sass --watch --style=expanded --sourcemap src/scss:dist/css

# Producción (comprimido)
sass --style=compressed src/scss/main.scss dist/css/main.css
```

### Mejores prácticas

```scss
// ✅ Bien: Variables descriptivas
$color-primary: #3498db;
$spacing-large: 2rem;

// ❌ Mal: Variables poco descriptivas
$c1: #3498db;
$s1: 2rem;

// ✅ Bien: Mixins reutilizables
@mixin button-variant($color) {
  background: $color;
  border: 1px solid $color;
  
  &:hover {
    background: darken($color, 10%);
  }
}

// ✅ Bien: Funciones puras
@function spacing($multiplier: 1) {
  @return $base-spacing * $multiplier;
}

// ✅ Bien: Anidamiento limitado (máximo 3-4 niveles)
.navbar {
  .nav-item {
    .nav-link {
      // No anidar más
    }
  }
}
```

## Integración con herramientas modernas

### Con Webpack

```javascript
// webpack.config.js
module.exports = {
  module: {
    rules: [
      {
        test: /\.scss$/,
        use: [
          'style-loader',
          'css-loader',
          'sass-loader'
        ]
      }
    ]
  }
};
```

### Con Gulp

```javascript
// gulpfile.js
const gulp = require('gulp');
const sass = require('gulp-sass')(require('sass'));
const autoprefixer = require('gulp-autoprefixer');

gulp.task('sass', function() {
  return gulp.src('src/scss/**/*.scss')
    .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
    .pipe(autoprefixer())
    .pipe(gulp.dest('dist/css'));
});

gulp.task('watch', function() {
  gulp.watch('src/scss/**/*.scss', gulp.series('sass'));
});
```

## Ejemplo práctico completo

```scss
// _variables.scss
$colors: (
  primary: #3498db,
  secondary: #2ecc71,
  danger: #e74c3c,
  warning: #f39c12,
  dark: #2c3e50,
  light: #ecf0f1
);

$spacing: (
  xs: 0.25rem,
  sm: 0.5rem,
  md: 1rem,
  lg: 1.5rem,
  xl: 2rem
);

$breakpoints: (
  sm: 576px,
  md: 768px,
  lg: 992px,
  xl: 1200px
);

$border-radius: 4px;
$box-shadow: 0 2px 4px rgba(0,0,0,0.1);

// _mixins.scss
@mixin flex-center {
  display: flex;
  justify-content: center;
  align-items: center;
}

@mixin button-variant($color, $hover-color: null) {
  $hover: if($hover-color, $hover-color, darken($color, 10%));
  
  background: $color;
  border: 2px solid $color;
  color: white;
  padding: map-get($spacing, md) map-get($spacing, lg);
  border-radius: $border-radius;
  text-decoration: none;
  display: inline-block;
  transition: all 0.3s ease;
  box-shadow: $box-shadow;
  
  &:hover {
    background: $hover;
    border-color: $hover;
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(0,0,0,0.15);
  }
}

@mixin media-query($breakpoint) {
  $value: map-get($breakpoints, $breakpoint);
  
  @if $value {
    @media (min-width: $value) {
      @content;
    }
  } @else {
    @warn "#{$breakpoint} no es un breakpoint válido. Usar: #{map-keys($breakpoints)}";
  }
}

// _functions.scss
@function color($name) {
  @return map-get($colors, $name);
}

@function spacing($size) {
  @return map-get($spacing, $size);
}

@function contrast-color($color) {
  $luminance: (red($color) * 0.299 + green($color) * 0.587 + blue($color) * 0.114) / 255;
  
  @if $luminance > 0.5 {
    @return #000;
  } @else {
    @return #fff;
  }
}

// main.scss
@import 'variables';
@import 'mixins';
@import 'functions';

// Base styles
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  line-height: 1.6;
  color: color(dark);
  background: color(light);
}

// Buttons
.btn {
  @include button-variant(color(primary));
  
  &-secondary {
    @include button-variant(color(secondary));
  }
  
  &-danger {
    @include button-variant(color(danger));
  }
}

// Cards
.card {
  background: white;
  border-radius: $border-radius;
  box-shadow: $box-shadow;
  padding: spacing(lg);
  margin-bottom: spacing(md);
  
  &-header {
    margin-bottom: spacing(md);
    
    h3 {
      color: color(primary);
    }
  }
  
  &-body {
    p {
      margin-bottom: spacing(sm);
    }
  }
}

// Layout
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 spacing(md);
  
  @include media-query(md) {
    padding: 0 spacing(lg);
  }
}

.grid {
  display: grid;
  gap: spacing(md);
  
  @include media-query(sm) {
    grid-template-columns: repeat(2, 1fr);
  }
  
  @include media-query(lg) {
    grid-template-columns: repeat(3, 1fr);
  }
}

// Utilities
.text-center {
  text-align: center;
}

.mb-md {
  margin-bottom: spacing(md);
}

.flex-center {
  @include flex-center;
}
```

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Sass/SCSS Preprocessor</title>
  <link rel="stylesheet" href="dist/css/main.css">
</head>
<body>
  <div class="container">
    <header>
      <h1>Sass/SCSS Master</h1>
      <p>Preprocesadores CSS avanzados</p>
    </header>

    <main>
      <section>
        <h2>Botones con Mixins</h2>
        <div class="button-examples">
          <a href="#" class="btn">Botón Primario</a>
          <a href="#" class="btn btn-secondary">Botón Secundario</a>
          <a href="#" class="btn btn-danger">Botón Danger</a>
        </div>
      </section>

      <section>
        <h2>Tarjetas con Variables</h2>
        <div class="grid">
          <div class="card">
            <div class="card-header">
              <h3>Tarjeta 1</h3>
            </div>
            <div class="card-body">
              <p>Contenido de la primera tarjeta usando variables Sass.</p>
            </div>
          </div>
          
          <div class="card">
            <div class="card-header">
              <h3>Tarjeta 2</h3>
            </div>
            <div class="card-body">
              <p>Más contenido con funciones y mixins aplicados.</p>
            </div>
          </div>
          
          <div class="card">
            <div class="card-header">
              <h3>Tarjeta 3</h3>
            </div>
            <div class="card-body">
              <p>Grid responsivo generado con media queries.</p>
            </div>
          </div>
        </div>
      </section>

      <section>
        <h2>Utilidades</h2>
        <div class="utilities">
          <div class="flex-center" style="height: 200px; background: #f8f9fa; border-radius: 8px;">
            <p class="text-center mb-md">Contenido centrado con mixin flex-center</p>
          </div>
        </div>
      </section>
    </main>

    <footer>
      <p>&copy; 2024 Sass/SCSS Preprocessor Course</p>
    </footer>
  </div>
</body>
</html>
```

## Resumen

Los preprocesadores CSS revolucionan el desarrollo frontend:

- ✅ **Variables**: Reutilización de valores y temas
- ✅ **Mixins**: Código reutilizable con parámetros
- ✅ **Funciones**: Lógica y cálculos avanzados
- ✅ **Anidamiento**: Estructura más legible y mantenible
- ✅ **Partials**: Organización modular del código
- ✅ **Bucles y condicionales**: Generación dinámica de CSS

Sass/SCSS es esencial para proyectos CSS grandes y complejos. Aprenderlo es obligatorio para desarrolladores profesionales.