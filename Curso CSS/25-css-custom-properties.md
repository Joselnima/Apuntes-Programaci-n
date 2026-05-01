# Módulo 28 - CSS Custom Properties Avanzado

En este módulo profundizaremos en CSS Custom Properties (variables CSS nativas), explorando técnicas avanzadas, patrones de diseño, y casos de uso complejos que van más allá de las variables básicas.

## Introducción a Custom Properties Avanzadas

### Más allá de las variables básicas

```css
/* Variables básicas */
:root {
  --color-primary: #3498db;
  --spacing-unit: 1rem;
}

/* Variables avanzadas con lógica */
:root {
  --color-primary-hue: 204;
  --color-primary-saturation: 70%;
  --color-primary-lightness: 53%;

  --color-primary: hsl(var(--color-primary-hue), var(--color-primary-saturation), var(--color-primary-lightness));
  --color-primary-dark: hsl(var(--color-primary-hue), var(--color-primary-saturation), calc(var(--color-primary-lightness) - 10%));
  --color-primary-light: hsl(var(--color-primary-hue), var(--color-primary-saturation), calc(var(--color-primary-lightness) + 10%));
}
```

## Variables Dinámicas y Contextuales

### Variables basadas en contexto

```css
/* Variables que cambian según el contexto */
:root {
  --theme-color: #3498db;
  --theme-bg: #ffffff;
}

[data-theme="dark"] {
  --theme-color: #74b9ff;
  --theme-bg: #2d3436;
}

[data-theme="neon"] {
  --theme-color: #00ff88;
  --theme-bg: #0a0a0a;
}

/* Uso */
.button {
  background: var(--theme-color);
  color: var(--theme-bg);
}
```

### Variables responsive

```css
:root {
  --text-size: 1rem;
  --spacing: 1rem;
}

@media (min-width: 768px) {
  :root {
    --text-size: 1.125rem;
    --spacing: 1.25rem;
  }
}

@media (min-width: 1024px) {
  :root {
    --text-size: 1.25rem;
    --spacing: 1.5rem;
  }
}

.element {
  font-size: var(--text-size);
  padding: var(--spacing);
}
```

## Funciones Avanzadas con Custom Properties

### Cálculos matemáticos complejos

```css
:root {
  --base-size: 16;
  --scale-ratio: 1.25;

  /* Escala tipográfica */
  --text-xs: calc(var(--base-size) * 1px);
  --text-sm: calc(var(--text-xs) * var(--scale-ratio));
  --text-base: calc(var(--text-sm) * var(--scale-ratio));
  --text-lg: calc(var(--text-base) * var(--scale-ratio));
  --text-xl: calc(var(--text-lg) * var(--scale-ratio));
  --text-2xl: calc(var(--text-xl) * var(--scale-ratio));
  --text-3xl: calc(var(--text-2xl) * var(--scale-ratio));
}
```

### Variables condicionales con @supports

```css
:root {
  --supports-grid: 0;
}

@supports (display: grid) {
  :root {
    --supports-grid: 1;
  }
}

.grid-container {
  display: flex;
}

@supports (display: grid) {
  .grid-container {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  }
}
```

## Sistema de Diseño con Custom Properties

### Design Tokens

```css
:root {
  /* Colores */
  --color-brand-primary: #6366f1;
  --color-brand-secondary: #06b6d4;
  --color-text-primary: #1f2937;
  --color-text-secondary: #6b7280;
  --color-bg-primary: #ffffff;
  --color-bg-secondary: #f9fafb;

  /* Espaciado */
  --space-1: 0.25rem;
  --space-2: 0.5rem;
  --space-3: 0.75rem;
  --space-4: 1rem;
  --space-5: 1.25rem;
  --space-6: 1.5rem;
  --space-8: 2rem;
  --space-10: 2.5rem;
  --space-12: 3rem;

  /* Tipografía */
  --font-family-sans: 'Inter', system-ui, sans-serif;
  --font-family-mono: 'JetBrains Mono', monospace;

  --font-size-xs: 0.75rem;
  --font-size-sm: 0.875rem;
  --font-size-base: 1rem;
  --font-size-lg: 1.125rem;
  --font-size-xl: 1.25rem;
  --font-size-2xl: 1.5rem;

  /* Bordes */
  --border-radius-sm: 0.125rem;
  --border-radius-md: 0.375rem;
  --border-radius-lg: 0.5rem;
  --border-radius-xl: 0.75rem;
  --border-radius-full: 9999px;

  /* Sombras */
  --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.05);
  --shadow-md: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
  --shadow-lg: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);
}
```

### Componentes con design tokens

```css
.button {
  font-family: var(--font-family-sans);
  font-size: var(--font-size-base);
  padding: var(--space-3) var(--space-4);
  border-radius: var(--border-radius-md);
  border: 1px solid transparent;
  background-color: var(--color-brand-primary);
  color: white;
  box-shadow: var(--shadow-sm);
  transition: all 0.2s ease;
}

.button:hover {
  background-color: color-mix(in srgb, var(--color-brand-primary), black 10%);
  box-shadow: var(--shadow-md);
}

.button--secondary {
  background-color: var(--color-bg-primary);
  color: var(--color-brand-primary);
  border-color: var(--color-brand-primary);
}

.button--large {
  font-size: var(--font-size-lg);
  padding: var(--space-4) var(--space-6);
}
```

## Variables Computadas y Lógicas

### Variables con lógica condicional

```css
:root {
  --is-dark: 0;
  --is-compact: 0;
}

@media (prefers-color-scheme: dark) {
  :root {
    --is-dark: 1;
  }
}

@media (max-width: 640px) {
  :root {
    --is-compact: 1;
  }
}

/* Variables computadas */
:root {
  --bg-color: hsl(0, 0%, calc(100% - (var(--is-dark) * 85%)));
  --text-color: hsl(0, 0%, calc(10% + (var(--is-dark) * 80%)));
  --padding: calc(var(--space-4) * (1 + var(--is-compact)));
}
```

### Sistema de temas dinámico

```css
:root {
  /* Tema base */
  --hue-primary: 240;
  --hue-secondary: 200;
  --saturation: 70%;
  --lightness-bg: 98%;
  --lightness-text: 15%;

  /* Generar colores */
  --color-primary: hsl(var(--hue-primary), var(--saturation), 50%);
  --color-secondary: hsl(var(--hue-secondary), var(--saturation), 50%);
  --color-bg: hsl(0, 0%, var(--lightness-bg));
  --color-text: hsl(0, 0%, var(--lightness-text));
}

/* Tema oscuro */
[data-theme="dark"] {
  --lightness-bg: 10%;
  --lightness-text: 90%;
}

/* Tema cálido */
[data-theme="warm"] {
  --hue-primary: 25;
  --hue-secondary: 45;
}

/* Tema frío */
[data-theme="cool"] {
  --hue-primary: 200;
  --hue-secondary: 240;
}
```

## Animaciones con Custom Properties

### Variables animables

```css
:root {
  --animation-duration: 0.3s;
  --animation-easing: ease-out;
  --scale-hover: 1.05;
  --translate-y-hover: -2px;
}

.card {
  transition:
    transform var(--animation-duration) var(--animation-easing),
    box-shadow var(--animation-duration) var(--animation-easing);
}

.card:hover {
  transform: scale(var(--scale-hover)) translateY(var(--translate-y-hover));
}
```

### Animaciones dinámicas

```css
@keyframes pulse {
  0%, 100% {
    transform: scale(var(--scale-base, 1));
  }
  50% {
    transform: scale(var(--scale-pulse, 1.1));
  }
}

.element {
  --scale-base: 1;
  --scale-pulse: 1.2;
  animation: pulse 2s infinite;
}

.element[data-intense] {
  --scale-pulse: 1.5;
}
```

## Custom Properties en JavaScript

### Manipulación desde JavaScript

```javascript
// Obtener valor de custom property
const rootStyles = getComputedStyle(document.documentElement);
const primaryColor = rootStyles.getPropertyValue('--color-primary');

// Establecer valor
document.documentElement.style.setProperty('--color-primary', '#ff6b6b');

// Cambiar tema dinámicamente
function setTheme(theme) {
  const themes = {
    light: {
      '--bg-color': '#ffffff',
      '--text-color': '#333333',
      '--accent-color': '#007bff'
    },
    dark: {
      '--bg-color': '#1a1a1a',
      '--text-color': '#ffffff',
      '--accent-color': '#4dabf7'
    }
  };

  Object.entries(themes[theme]).forEach(([property, value]) => {
    document.documentElement.style.setProperty(property, value);
  });
}
```

### Sistema de configuración dinámico

```javascript
class ThemeManager {
  constructor() {
    this.root = document.documentElement;
    this.themes = new Map();
  }

  defineTheme(name, properties) {
    this.themes.set(name, properties);
  }

  applyTheme(name) {
    const theme = this.themes.get(name);
    if (!theme) return;

    Object.entries(theme).forEach(([property, value]) => {
      this.root.style.setProperty(property, value);
    });
  }

  getThemeValue(property) {
    return getComputedStyle(this.root).getPropertyValue(property);
  }
}

// Uso
const themeManager = new ThemeManager();

themeManager.defineTheme('ocean', {
  '--primary-hue': '200',
  '--primary-saturation': '80%',
  '--primary-lightness': '50%'
});

themeManager.applyTheme('ocean');
```

## Técnicas Avanzadas

### Variables con unidades dinámicas

```css
:root {
  --viewport-unit: 1vw;
  --container-max-width: 1200px;
  --container-padding: calc(var(--space-4) + var(--viewport-unit) * 2);

  /* Responsive typography */
  --fluid-font-size: clamp(1rem, 2vw + 0.5rem, 2rem);
}

/* Container con padding fluido */
.container {
  max-width: var(--container-max-width);
  padding: 0 var(--container-padding);
  margin: 0 auto;
}
```

### Variables para layouts complejos

```css
:root {
  /* Grid system */
  --grid-columns: 12;
  --grid-gutter: var(--space-4);
  --grid-margin: var(--space-4);

  /* Generar clases de grid dinámicamente */
  --col-1: calc(100% / var(--grid-columns));
  --col-2: calc(var(--col-1) * 2);
  --col-3: calc(var(--col-1) * 3);
  /* ... */
}

.grid {
  display: grid;
  grid-template-columns: repeat(var(--grid-columns), 1fr);
  gap: var(--grid-gutter);
  margin: var(--grid-margin);
}

.col-6 {
  grid-column: span calc(var(--grid-columns) / 2);
}
```

## Debugging y Desarrollo

### Herramientas para debugging

```css
/* Mostrar valores de variables en desarrollo */
:root {
  --debug: 'Valor actual: ' var(--color-primary);
}

/* Usar en pseudo-elementos para debugging */
.debug::before {
  content: var(--debug);
  position: fixed;
  top: 0;
  left: 0;
  background: red;
  color: white;
  padding: 5px;
  font-size: 12px;
  z-index: 9999;
}
```

### Validación de variables

```css
/* Variables con valores por defecto robustos */
.button {
  background: var(--button-bg, var(--color-primary, #007bff));
  color: var(--button-color, var(--color-text-on-primary, white));
  padding: var(--button-padding, var(--space-3) var(--space-4));
}
```

## Casos de Uso Avanzados

### Sistema de componentes escalable

```css
/* Componente base */
.c-button {
  --button-bg: var(--color-primary);
  --button-color: white;
  --button-padding: var(--space-3) var(--space-4);
  --button-radius: var(--border-radius-md);

  background: var(--button-bg);
  color: var(--button-color);
  padding: var(--button-padding);
  border-radius: var(--button-radius);
  border: none;
  cursor: pointer;
}

/* Variantes usando custom properties */
.c-button--secondary {
  --button-bg: var(--color-secondary);
}

.c-button--outline {
  --button-bg: transparent;
  --button-color: var(--color-primary);
  border: 1px solid var(--button-bg);
}

.c-button--large {
  --button-padding: var(--space-4) var(--space-6);
}
```

### Tema con múltiples variantes

```css
:root {
  /* Tema base */
  --theme-name: 'default';
  --theme-colors: 'blue';
  --theme-mode: 'light';

  /* Colores dinámicos */
  --color-primary: var(--color-blue-500);
  --color-secondary: var(--color-blue-600);
}

/* Tema rojo */
[data-theme-colors="red"] {
  --theme-colors: 'red';
  --color-primary: var(--color-red-500);
  --color-secondary: var(--color-red-600);
}

/* Modo oscuro */
[data-theme-mode="dark"] {
  --theme-mode: 'dark';
  --color-bg: var(--color-gray-900);
  --color-text: var(--color-gray-100);
}
```

## Performance y Mejores Prácticas

### Optimizaciones

```css
/* Evitar variables en animaciones críticas */
.element {
  /* ❌ Mal: variable en propiedad animada */
  transform: translateX(var(--offset));

  /* ✅ Bueno: calcular en CSS */
  --offset: 100px;
  transform: translateX(var(--offset));
}

/* Usar variables para valores estáticos */
.element {
  /* ✅ Bueno para valores que no cambian durante animaciones */
  color: var(--text-color);
  background: var(--bg-color);
}
```

### Estrategias de carga

```css
/* Variables críticas inline */
<style>
  :root {
    --color-critical: #007bff;
    --font-critical: 'Arial', sans-serif;
  }
</style>

/* Variables no críticas en archivo separado */
@import 'variables.css';
```

## Resumen

CSS Custom Properties avanzadas permiten:

- ✅ **Sistemas de diseño dinámicos**
- ✅ **Temas contextuales y responsive**
- ✅ **Cálculos matemáticos complejos**
- ✅ **Integración perfecta con JavaScript**
- ✅ **Componentes reutilizables y escalables**
- ✅ **Debugging y desarrollo eficientes**

Las custom properties transforman CSS de un lenguaje estático a uno dinámico y programable, abriendo posibilidades ilimitadas para el desarrollo frontend moderno.