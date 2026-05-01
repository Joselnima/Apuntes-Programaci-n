# Módulo 29 - CSS Nesting

En este módulo aprenderás sobre CSS Nesting, una nueva característica nativa de CSS que permite anidar selectores de manera similar a Sass/SCSS, pero con sintaxis nativa y mejores prestaciones.

## Introducción a CSS Nesting

### ¿Qué es CSS Nesting?

CSS Nesting es una característica nativa de CSS que permite escribir selectores anidados sin necesidad de preprocesadores. Simplifica la escritura y lectura de CSS complejo.

```css
/* ❌ Sin nesting */
.card {
  background: white;
}
.card .title {
  font-size: 1.5rem;
}
.card .title:hover {
  color: blue;
}

/* ✅ Con CSS Nesting */
.card {
  background: white;

  .title {
    font-size: 1.5rem;

    &:hover {
      color: blue;
    }
  }
}
```

## Sintaxis Básica

### Anidamiento simple

```css
.parent {
  color: blue;

  .child {
    color: red;
  }
}

/* Equivalente a: */
.parent { color: blue; }
.parent .child { color: red; }
```

### Selectores anidados

```css
.article {
  font-size: 1rem;

  h2 {
    font-size: 1.5em;
  }

  p {
    margin-bottom: 1em;

    a {
      color: blue;

      &:hover {
        color: darkblue;
      }
    }
  }
}
```

### El símbolo `&` (ampersand)

```css
.button {
  background: blue;
  color: white;

  /* & representa el selector padre */
  &:hover {
    background: darkblue;
  }

  &:focus {
    outline: 2px solid yellow;
  }

  /* Modificadores */
  &--primary {
    background: green;
  }

  &--large {
    padding: 1rem 2rem;
  }
}
```

## Casos de Uso Comunes

### Componentes con variantes

```css
.card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);

  /* Estados */
  &:hover {
    box-shadow: 0 4px 16px rgba(0,0,0,0.2);
  }

  /* Partes del componente */
  .card__header {
    padding: 1rem;
    border-bottom: 1px solid #eee;
  }

  .card__body {
    padding: 1rem;
  }

  .card__footer {
    padding: 1rem;
    border-top: 1px solid #eee;
    background: #f9f9f9;
  }

  /* Variantes */
  &--featured {
    border: 2px solid gold;

    .card__header {
      background: gold;
      color: white;
    }
  }

  &--compact {
    .card__body {
      padding: 0.5rem;
    }
  }
}
```

### Navegación y menús

```css
.nav {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;

  .nav__item {
    margin: 0 1rem;

    .nav__link {
      color: #333;
      text-decoration: none;
      padding: 0.5rem 1rem;
      border-radius: 4px;
      transition: background-color 0.2s;

      &:hover {
        background-color: #f0f0f0;
      }

      &:focus {
        outline: 2px solid #007bff;
        outline-offset: 2px;
      }

      /* Estado activo */
      &.is-active {
        background-color: #007bff;
        color: white;
      }
    }

    /* Submenús */
    &.has-submenu {
      position: relative;

      &:hover .nav__submenu {
        display: block;
      }

      .nav__submenu {
        display: none;
        position: absolute;
        top: 100%;
        left: 0;
        background: white;
        border: 1px solid #ddd;
        border-radius: 4px;
        box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        min-width: 200px;

        .nav__submenu-item {
          .nav__submenu-link {
            display: block;
            padding: 0.75rem 1rem;
            color: #333;
            text-decoration: none;

            &:hover {
              background-color: #f8f8f8;
            }
          }
        }
      }
    }
  }
}
```

## Técnicas Avanzadas

### Anidamiento múltiple

```css
.form {
  .form__group {
    margin-bottom: 1rem;

    .form__label {
      display: block;
      margin-bottom: 0.5rem;
      font-weight: 500;
    }

    .form__input {
      width: 100%;
      padding: 0.75rem;
      border: 1px solid #ddd;
      border-radius: 4px;

      &:focus {
        outline: none;
        border-color: #007bff;
        box-shadow: 0 0 0 3px rgba(0,123,255,0.1);
      }

      &:invalid {
        border-color: #dc3545;

        &:focus {
          box-shadow: 0 0 0 3px rgba(220,53,69,0.1);
        }
      }

      /* Placeholder */
      &::placeholder {
        color: #999;
      }
    }

    /* Estados de validación */
    &.is-valid .form__input {
      border-color: #28a745;
    }

    &.is-invalid .form__input {
      border-color: #dc3545;
    }
  }

  .form__actions {
    display: flex;
    gap: 1rem;
    margin-top: 2rem;

    .form__button {
      padding: 0.75rem 1.5rem;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      font-size: 1rem;
      transition: background-color 0.2s;

      &--primary {
        background-color: #007bff;
        color: white;

        &:hover {
          background-color: #0056b3;
        }
      }

      &--secondary {
        background-color: #6c757d;
        color: white;

        &:hover {
          background-color: #545b62;
        }
      }

      &:disabled {
        opacity: 0.6;
        cursor: not-allowed;
      }
    }
  }
}
```

### Media queries anidadas

```css
.card {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;

  @media (min-width: 768px) {
    grid-template-columns: 1fr 1fr;

    .card__image {
      grid-column: 1 / -1;
    }
  }

  @media (min-width: 1024px) {
    grid-template-columns: 1fr 2fr;

    .card__sidebar {
      background: #f8f9fa;
      padding: 1rem;
    }
  }
}
```

### Container queries con nesting

```css
.card {
  container-type: inline-size;

  .card__content {
    display: flex;
    flex-direction: column;
  }

  @container (min-width: 400px) {
    .card__content {
      flex-direction: row;

      .card__image {
        flex: 0 0 200px;
      }

      .card__text {
        flex: 1;
        margin-left: 1rem;
      }
    }
  }
}
```

## Selectores Complejos

### Combinadores y pseudo-clases

```css
.article {
  /* Descendiente */
  p {
    margin-bottom: 1em;
  }

  /* Hijo directo */
  > .highlight {
    background: yellow;
  }

  /* Hermano adyacente */
  h2 + p {
    margin-top: 0;
  }

  /* Pseudo-clases */
  &:first-child {
    margin-top: 0;
  }

  &:last-child {
    margin-bottom: 0;
  }

  /* Pseudo-elementos */
  &::before {
    content: "📖 ";
  }

  /* Atributos */
  [data-featured] {
    border: 2px solid gold;
  }
}
```

### Selectores avanzados

```css
.gallery {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;

  /* nth-child */
  .gallery__item {
    &:nth-child(3n + 1) {
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
    }

    &:nth-child(3n + 2) {
      background: linear-gradient(45deg, #667eea, #764ba2);
    }

    &:nth-child(3n + 3) {
      background: linear-gradient(45deg, #f093fb, #f5576c);
    }

    /* Hover en elementos hermanos */
    &:hover ~ .gallery__item {
      opacity: 0.7;
    }
  }
}
```

## Mejores Prácticas

### Evitar anidamiento excesivo

```css
/* ❌ Demasiados niveles de anidamiento */
.header {
  .nav {
    .nav__list {
      .nav__item {
        .nav__link {
          &:hover {
            .nav__icon {
              transform: rotate(45deg);
            }
          }
        }
      }
    }
  }
}

/* ✅ Anidamiento razonable */
.header {
  .nav {
    .nav__link {
      &:hover .nav__icon {
        transform: rotate(45deg);
      }
    }
  }
}
```

### Organizar por componentes

```css
/* Componente Button */
.button {
  /* Estilos base */
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;

  /* Estados */
  &:hover {
    transform: translateY(-1px);
  }

  &:focus {
    outline: 2px solid #007bff;
    outline-offset: 2px;
  }

  /* Variantes */
  &--primary {
    background: #007bff;
    color: white;
  }

  &--secondary {
    background: #6c757d;
    color: white;
  }

  /* Tamaños */
  &--small {
    padding: 0.5rem 1rem;
    font-size: 0.875rem;
  }

  &--large {
    padding: 1rem 2rem;
    font-size: 1.125rem;
  }
}
```

## Compatibilidad y Fallbacks

### Detectar soporte

```css
/* Verificar soporte de nesting */
@supports (selector(&)) {
  .component {
    /* Usar nesting */
    .child {
      color: blue;
    }
  }
}

/* Fallback sin nesting */
@supports not (selector(&)) {
  .component .child {
    color: blue;
  }
}
```

### PostCSS como fallback

```javascript
// postcss.config.js
module.exports = {
  plugins: [
    require('postcss-nesting'), // Polyfill para nesting
    require('autoprefixer'),
    require('cssnano')
  ]
}
```

## Ejemplos Prácticos

### Sistema de diseño con nesting

```css
/* Variables */
:root {
  --color-primary: #007bff;
  --color-secondary: #6c757d;
  --spacing-unit: 1rem;
  --border-radius: 4px;
}

/* Componente Card */
.card {
  background: white;
  border-radius: var(--border-radius);
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  overflow: hidden;

  .card__header {
    padding: calc(var(--spacing-unit) * 1.5);
    background: #f8f9fa;
    border-bottom: 1px solid #dee2e6;

    .card__title {
      margin: 0;
      font-size: 1.25rem;
      color: #333;
    }

    .card__subtitle {
      margin: 0.25rem 0 0 0;
      font-size: 0.875rem;
      color: #6c757d;
    }
  }

  .card__body {
    padding: calc(var(--spacing-unit) * 1.5);

    .card__text {
      margin-bottom: calc(var(--spacing-unit) * 0.75);
      line-height: 1.6;

      &:last-child {
        margin-bottom: 0;
      }
    }
  }

  .card__footer {
    padding: calc(var(--spacing-unit) * 1.5);
    background: #f8f9fa;
    border-top: 1px solid #dee2e6;

    .card__actions {
      display: flex;
      justify-content: flex-end;
      gap: calc(var(--spacing-unit) * 0.5);

      .card__button {
        padding: 0.5rem 1rem;
        border: 1px solid var(--color-secondary);
        background: white;
        color: var(--color-secondary);
        border-radius: var(--border-radius);
        cursor: pointer;
        transition: all 0.2s;

        &:hover {
          background: var(--color-secondary);
          color: white;
        }

        &--primary {
          border-color: var(--color-primary);
          color: var(--color-primary);

          &:hover {
            background: var(--color-primary);
            color: white;
          }
        }
      }
    }
  }

  /* Estados */
  &:hover {
    box-shadow: 0 4px 16px rgba(0,0,0,0.15);
    transform: translateY(-2px);
  }

  /* Variantes */
  &--elevated {
    box-shadow: 0 8px 32px rgba(0,0,0,0.2);
  }

  &--bordered {
    border: 1px solid #dee2e6;
  }
}
```

### Layout responsive con nesting

```css
.layout {
  display: grid;
  min-height: 100vh;

  .layout__sidebar {
    background: #2c3e50;
    color: white;
    padding: 2rem;

    .layout__nav {
      .layout__nav-item {
        margin-bottom: 1rem;

        .layout__nav-link {
          color: #bdc3c7;
          text-decoration: none;
          display: block;
          padding: 0.75rem 1rem;
          border-radius: 4px;
          transition: all 0.2s;

          &:hover {
            background: rgba(255,255,255,0.1);
            color: white;
          }

          &.is-active {
            background: #3498db;
            color: white;
          }
        }
      }
    }
  }

  .layout__main {
    padding: 2rem;
    background: #ecf0f1;

    .layout__content {
      background: white;
      border-radius: 8px;
      padding: 2rem;
      box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    }
  }

  /* Responsive */
  @media (max-width: 768px) {
    grid-template-columns: 1fr;

    .layout__sidebar {
      order: 2;

      .layout__nav {
        display: flex;
        flex-wrap: wrap;
        gap: 0.5rem;

        .layout__nav-item {
          margin-bottom: 0;

          .layout__nav-link {
            padding: 0.5rem 1rem;
            font-size: 0.875rem;
          }
        }
      }
    }

    .layout__main {
      order: 1;

      .layout__content {
        padding: 1rem;
      }
    }
  }
}
```

## Debugging y Desarrollo

### Herramientas de desarrollo

```css
/* Mostrar estructura de nesting en desarrollo */
.component {
  /* Nivel 1 */
  background: white;

  .child {
    /* Nivel 2 */
    color: blue;

    .grandchild {
      /* Nivel 3 */
      font-size: 0.875rem;
    }
  }
}
```

### Linting con Stylelint

```javascript
// .stylelintrc.js
module.exports = {
  extends: ['stylelint-config-standard'],
  rules: {
    'max-nesting-depth': 3,
    'selector-max-compound-selectors': 3,
    'selector-no-qualifying-type': null, // Permitir anidamiento
  }
}
```

## Resumen

CSS Nesting revoluciona la escritura de CSS:

- ✅ **Sintaxis más limpia y legible**
- ✅ **Menos repetición de selectores**
- ✅ **Mejor organización del código**
- ✅ **Transición natural desde preprocesadores**
- ✅ **Mejor mantenibilidad**

CSS Nesting hace que escribir CSS complejo sea tan natural como escribir HTML anidado, mejorando significativamente la experiencia de desarrollo.