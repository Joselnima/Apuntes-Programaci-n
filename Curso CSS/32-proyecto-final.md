# Módulo 17 - Proyecto Final: Sitio Web Completo

En este módulo final construiremos un sitio web completo que integre todos los conceptos aprendidos en el curso. Crearemos un portfolio personal moderno con múltiples secciones, diseño responsive, animaciones, y optimización de performance.

## Descripción del Proyecto

### Objetivos del proyecto

- ✅ **Aplicar todos los conceptos CSS**: Desde básico hasta avanzado
- ✅ **Diseño responsive**: Móvil-first con breakpoints estratégicos
- ✅ **Animaciones y transiciones**: Interacciones fluidas
- ✅ **Arquitectura CSS**: BEM + ITCSS para organización
- ✅ **Optimización**: Performance y Core Web Vitals
- ✅ **Accesibilidad**: WCAG 2.1 compliance
- ✅ **Modern CSS**: Grid, Flexbox, Custom Properties

### Estructura del sitio

```
portfolio/
├── index.html (Home/Landing)
├── about.html (Sobre mí)
├── projects.html (Proyectos)
├── blog.html (Blog/Artículos)
├── contact.html (Contacto)
├── assets/
│   ├── css/
│   │   ├── main.css (CSS principal)
│   │   ├── critical.css (CSS crítico)
│   │   └── components/ (Componentes modulares)
│   ├── js/
│   │   ├── main.js (JavaScript principal)
│   │   └── components/ (Componentes JS)
│   └── images/ (Imágenes optimizadas)
└── README.md
```

## Configuración del Proyecto

### Estructura de archivos CSS (ITCSS)

```
css/
├── 01-settings/
│   ├── _colors.scss
│   ├── _typography.scss
│   └── _spacing.scss
├── 02-tools/
│   ├── _mixins.scss
│   └── _functions.scss
├── 03-generic/
│   ├── _reset.scss
│   ├── _normalize.scss
│   └── _base.scss
├── 04-elements/
│   ├── _headings.scss
│   ├── _links.scss
│   └── _images.scss
├── 05-objects/
│   ├── _container.scss
│   ├── _grid.scss
│   └── _media.scss
├── 06-components/
│   ├── _button.scss
│   ├── _card.scss
│   ├── _modal.scss
│   ├── _navigation.scss
│   └── _hero.scss
├── 07-utilities/
│   ├── _spacing.scss
│   ├── _colors.scss
│   ├── _typography.scss
│   └── _visibility.scss
└── main.scss
```

### Variables y configuración

```scss
// 01-settings/_colors.scss
:root {
  // Colores principales
  --color-primary: #6366f1;
  --color-primary-dark: #4f46e5;
  --color-primary-light: #a5b4fc;
  
  // Colores secundarios
  --color-secondary: #06b6d4;
  --color-secondary-dark: #0891b2;
  --color-secondary-light: #67e8f9;
  
  // Colores neutros
  --color-gray-50: #f9fafb;
  --color-gray-100: #f3f4f6;
  --color-gray-200: #e5e7eb;
  --color-gray-300: #d1d5db;
  --color-gray-400: #9ca3af;
  --color-gray-500: #6b7280;
  --color-gray-600: #4b5563;
  --color-gray-700: #374151;
  --color-gray-800: #1f2937;
  --color-gray-900: #111827;
  
  // Estados
  --color-success: #10b981;
  --color-warning: #f59e0b;
  --color-error: #ef4444;
  
  // Tema oscuro
  --color-dark-bg: #0f172a;
  --color-dark-surface: #1e293b;
  --color-dark-text: #f1f5f9;
}
```

```scss
// 01-settings/_typography.scss
:root {
  // Fuentes
  --font-family-sans: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  --font-family-mono: 'JetBrains Mono', 'Fira Code', monospace;
  
  // Tamaños
  --font-size-xs: 0.75rem;
  --font-size-sm: 0.875rem;
  --font-size-base: 1rem;
  --font-size-lg: 1.125rem;
  --font-size-xl: 1.25rem;
  --font-size-2xl: 1.5rem;
  --font-size-3xl: 1.875rem;
  --font-size-4xl: 2.25rem;
  --font-size-5xl: 3rem;
  
  // Pesos
  --font-weight-light: 300;
  --font-weight-normal: 400;
  --font-weight-medium: 500;
  --font-weight-semibold: 600;
  --font-weight-bold: 700;
  
  // Líneas
  --line-height-tight: 1.25;
  --line-height-normal: 1.5;
  --line-height-relaxed: 1.625;
}
```

```scss
// 01-settings/_spacing.scss
:root {
  // Espaciado
  --spacing-1: 0.25rem;
  --spacing-2: 0.5rem;
  --spacing-3: 0.75rem;
  --spacing-4: 1rem;
  --spacing-5: 1.25rem;
  --spacing-6: 1.5rem;
  --spacing-8: 2rem;
  --spacing-10: 2.5rem;
  --spacing-12: 3rem;
  --spacing-16: 4rem;
  --spacing-20: 5rem;
  --spacing-24: 6rem;
  
  // Breakpoints
  --breakpoint-sm: 640px;
  --breakpoint-md: 768px;
  --breakpoint-lg: 1024px;
  --breakpoint-xl: 1280px;
  --breakpoint-2xl: 1536px;
}
```

## Componentes Base

### Botón reutilizable

```scss
// 06-components/_button.scss
.c-button {
  --button-padding-x: var(--spacing-4);
  --button-padding-y: var(--spacing-2);
  --button-border-radius: 0.375rem;
  --button-font-size: var(--font-size-sm);
  --button-font-weight: var(--font-weight-medium);
  
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: var(--spacing-2);
  padding: var(--button-padding-y) var(--button-padding-x);
  font-family: var(--font-family-sans);
  font-size: var(--button-font-size);
  font-weight: var(--button-font-weight);
  line-height: var(--line-height-tight);
  text-decoration: none;
  border: 1px solid transparent;
  border-radius: var(--button-border-radius);
  cursor: pointer;
  transition: all 0.2s ease-in-out;
  user-select: none;
  
  &:focus {
    outline: 2px solid var(--color-primary);
    outline-offset: 2px;
  }
  
  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
  
  // Variantes
  &--primary {
    background: var(--color-primary);
    color: white;
    
    &:hover:not(:disabled) {
      background: var(--color-primary-dark);
      transform: translateY(-1px);
      box-shadow: 0 4px 12px rgba(99, 102, 241, 0.4);
    }
  }
  
  &--secondary {
    background: white;
    color: var(--color-gray-700);
    border-color: var(--color-gray-300);
    
    &:hover:not(:disabled) {
      background: var(--color-gray-50);
      border-color: var(--color-gray-400);
    }
  }
  
  &--outline {
    background: transparent;
    color: var(--color-primary);
    border-color: var(--color-primary);
    
    &:hover:not(:disabled) {
      background: var(--color-primary);
      color: white;
    }
  }
  
  // Tamaños
  &--sm {
    --button-padding-x: var(--spacing-3);
    --button-padding-y: var(--spacing-1);
    --button-font-size: var(--font-size-xs);
  }
  
  &--lg {
    --button-padding-x: var(--spacing-6);
    --button-padding-y: var(--spacing-3);
    --button-font-size: var(--font-size-base);
  }
  
  // Estados
  &--loading {
    position: relative;
    color: transparent !important;
    
    &::after {
      content: '';
      position: absolute;
      width: 1rem;
      height: 1rem;
      border: 2px solid currentColor;
      border-right-color: transparent;
      border-radius: 50%;
      animation: spin 1s linear infinite;
    }
  }
}

@keyframes spin {
  to { transform: rotate(360deg); }
}
```

### Sistema de grid

```scss
// 05-objects/_grid.scss
.o-grid {
  display: grid;
  gap: var(--spacing-6);
  
  &--1-col {
    grid-template-columns: 1fr;
  }
  
  &--2-col {
    grid-template-columns: repeat(2, 1fr);
  }
  
  &--3-col {
    grid-template-columns: repeat(3, 1fr);
  }
  
  &--4-col {
    grid-template-columns: repeat(4, 1fr);
  }
  
  // Responsive
  @media (max-width: 768px) {
    &--2-col,
    &--3-col,
    &--4-col {
      grid-template-columns: 1fr;
    }
  }
  
  @media (min-width: 769px) and (max-width: 1024px) {
    &--3-col,
    &--4-col {
      grid-template-columns: repeat(2, 1fr);
    }
  }
}
```

### Tarjetas de proyecto

```scss
// 06-components/_card.scss
.c-card {
  background: white;
  border-radius: 0.75rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1), 0 1px 2px rgba(0, 0, 0, 0.06);
  overflow: hidden;
  transition: all 0.3s ease;
  
  &:hover {
    transform: translateY(-4px);
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15), 0 4px 10px rgba(0, 0, 0, 0.1);
  }
  
  &__image {
    width: 100%;
    height: 200px;
    object-fit: cover;
    transition: transform 0.3s ease;
    
    .c-card:hover & {
      transform: scale(1.05);
    }
  }
  
  &__content {
    padding: var(--spacing-6);
  }
  
  &__title {
    font-size: var(--font-size-xl);
    font-weight: var(--font-weight-semibold);
    color: var(--color-gray-900);
    margin-bottom: var(--spacing-2);
  }
  
  &__description {
    color: var(--color-gray-600);
    line-height: var(--line-height-relaxed);
    margin-bottom: var(--spacing-4);
  }
  
  &__tags {
    display: flex;
    flex-wrap: wrap;
    gap: var(--spacing-2);
    margin-bottom: var(--spacing-4);
  }
  
  &__tag {
    background: var(--color-gray-100);
    color: var(--color-gray-700);
    padding: var(--spacing-1) var(--spacing-2);
    border-radius: 0.25rem;
    font-size: var(--font-size-xs);
    font-weight: var(--font-weight-medium);
  }
  
  &__actions {
    display: flex;
    gap: var(--spacing-3);
  }
}
```

## Layout Principal

### Header con navegación

```scss
// 06-components/_navigation.scss
.c-navigation {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid var(--color-gray-200);
  z-index: 1000;
  transition: all 0.3s ease;
  
  &__container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 var(--spacing-4);
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 4rem;
  }
  
  &__logo {
    font-size: var(--font-size-xl);
    font-weight: var(--font-weight-bold);
    color: var(--color-primary);
    text-decoration: none;
  }
  
  &__menu {
    display: flex;
    list-style: none;
    gap: var(--spacing-8);
    margin: 0;
    padding: 0;
    
    @media (max-width: 768px) {
      display: none;
    }
  }
  
  &__link {
    color: var(--color-gray-700);
    text-decoration: none;
    font-weight: var(--font-weight-medium);
    transition: color 0.2s ease;
    
    &:hover,
    &--active {
      color: var(--color-primary);
    }
  }
  
  &__mobile-toggle {
    display: none;
    background: none;
    border: none;
    cursor: pointer;
    padding: var(--spacing-2);
    
    @media (max-width: 768px) {
      display: block;
    }
  }
  
  // Mobile menu
  &__mobile-menu {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: white;
    border-bottom: 1px solid var(--color-gray-200);
    display: none;
    
    &.is-open {
      display: block;
    }
    
    @media (min-width: 769px) {
      display: none !important;
    }
  }
  
  &__mobile-list {
    list-style: none;
    margin: 0;
    padding: var(--spacing-4);
  }
  
  &__mobile-link {
    display: block;
    padding: var(--spacing-3);
    color: var(--color-gray-700);
    text-decoration: none;
    border-radius: 0.375rem;
    transition: background-color 0.2s ease;
    
    &:hover {
      background: var(--color-gray-50);
    }
  }
}
```

### Hero section

```scss
// 06-components/_hero.scss
.c-hero {
  min-height: 100vh;
  display: flex;
  align-items: center;
  background: linear-gradient(135deg, var(--color-primary-light) 0%, var(--color-secondary-light) 100%);
  position: relative;
  overflow: hidden;
  
  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: url('hero-pattern.svg') repeat;
    opacity: 0.1;
  }
  
  &__container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 var(--spacing-4);
    position: relative;
    z-index: 1;
  }
  
  &__content {
    max-width: 600px;
    
    @media (max-width: 768px) {
      text-align: center;
    }
  }
  
  &__title {
    font-size: clamp(var(--font-size-4xl), 8vw, var(--font-size-5xl));
    font-weight: var(--font-weight-bold);
    color: var(--color-gray-900);
    margin-bottom: var(--spacing-4);
    line-height: var(--line-height-tight);
  }
  
  &__subtitle {
    font-size: var(--font-size-xl);
    color: var(--color-gray-700);
    margin-bottom: var(--spacing-8);
    line-height: var(--line-height-relaxed);
  }
  
  &__actions {
    display: flex;
    gap: var(--spacing-4);
    flex-wrap: wrap;
    
    @media (max-width: 768px) {
      justify-content: center;
    }
  }
  
  &__image {
    position: absolute;
    right: 0;
    top: 50%;
    transform: translateY(-50%);
    width: 400px;
    height: 400px;
    border-radius: 50%;
    background: white;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
    display: flex;
    align-items: center;
    justify-content: center;
    
    @media (max-width: 1024px) {
      display: none;
    }
    
    img {
      width: 80%;
      height: 80%;
      object-fit: cover;
      border-radius: 50%;
    }
  }
}
```

## Páginas del Sitio

### Home (index.html)

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Juan Pérez - Desarrollador Web</title>
  <meta name="description" content="Portfolio de Juan Pérez, desarrollador web full-stack especializado en React, Node.js y diseño UX/UI.">
  
  <!-- Preconnect para performance -->
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  
  <!-- Fonts -->
  <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap" rel="stylesheet">
  
  <!-- CSS crítico inline -->
  <style>
    body{margin:0;font-family:'Inter',sans-serif;background:#fff;color:#111827}
    .c-navigation{position:fixed;top:0;left:0;right:0;background:rgba(255,255,255,.95);backdrop-filter:blur(10px);border-bottom:1px solid #e5e7eb;z-index:1000}
    .c-navigation__container{max-width:1200px;margin:0 auto;padding:0 1rem;display:flex;align-items:center;justify-content:space-between;height:4rem}
    .c-navigation__logo{font-size:1.25rem;font-weight:700;color:#6366f1;text-decoration:none}
    .c-hero{min-height:100vh;display:flex;align-items:center;background:linear-gradient(135deg,#a5b4fc 0%,#67e8f9 100%)}
    .c-hero__container{max-width:1200px;margin:0 auto;padding:0 1rem;position:relative;z-index:1}
    .c-hero__title{font-size:clamp(2.25rem,8vw,3rem);font-weight:700;color:#111827;margin-bottom:1rem;line-height:1.25}
    .c-hero__subtitle{font-size:1.25rem;color:#374151;margin-bottom:2rem;line-height:1.625}
    .c-button{display:inline-flex;align-items:center;justify-content:center;padding:.5rem 1rem;border:1px solid transparent;border-radius:.375rem;text-decoration:none;cursor:pointer;font-weight:500;color:#fff;background:#6366f1}
  </style>
  
  <!-- CSS restante -->
  <link rel="stylesheet" href="assets/css/main.css">
</head>
<body>
  <!-- Navigation -->
  <nav class="c-navigation">
    <div class="c-navigation__container">
      <a href="#" class="c-navigation__logo">Juan Pérez</a>
      <ul class="c-navigation__menu">
        <li><a href="#about" class="c-navigation__link">Sobre mí</a></li>
        <li><a href="#projects" class="c-navigation__link">Proyectos</a></li>
        <li><a href="#blog" class="c-navigation__link">Blog</a></li>
        <li><a href="#contact" class="c-navigation__link">Contacto</a></li>
      </ul>
      <button class="c-navigation__mobile-toggle" aria-label="Toggle menu">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M4 6h16M4 12h16M4 18h16"></path>
        </svg>
      </button>
    </div>
    <div class="c-navigation__mobile-menu">
      <ul class="c-navigation__mobile-list">
        <li><a href="#about" class="c-navigation__mobile-link">Sobre mí</a></li>
        <li><a href="#projects" class="c-navigation__mobile-link">Proyectos</a></li>
        <li><a href="#blog" class="c-navigation__mobile-link">Blog</a></li>
        <li><a href="#contact" class="c-navigation__mobile-link">Contacto</a></li>
      </ul>
    </div>
  </nav>

  <!-- Hero Section -->
  <section class="c-hero">
    <div class="c-hero__container">
      <div class="c-hero__content">
        <h1 class="c-hero__title">
          Hola, soy Juan Pérez
          <span class="c-hero__title-accent">Desarrollador Web</span>
        </h1>
        <p class="c-hero__subtitle">
          Creo experiencias web excepcionales combinando diseño moderno, 
          código limpio y las mejores prácticas de desarrollo.
        </p>
        <div class="c-hero__actions">
          <a href="#projects" class="c-button c-button--primary c-button--lg">
            Ver mis trabajos
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M7 17L17 7M17 7H7M17 7V17"></path>
            </svg>
          </a>
          <a href="#contact" class="c-button c-button--outline c-button--lg">
            Contactar
          </a>
        </div>
      </div>
      <div class="c-hero__image">
        <img src="assets/images/profile.jpg" alt="Juan Pérez" loading="eager">
      </div>
    </div>
  </section>

  <!-- About Section -->
  <section id="about" class="section">
    <div class="container">
      <div class="section__header">
        <h2 class="section__title">Sobre mí</h2>
        <p class="section__subtitle">Conoce mi trayectoria y experiencia</p>
      </div>
      
      <div class="o-grid o-grid--2-col">
        <div class="about__content">
          <h3>Mi historia</h3>
          <p>
            Soy un desarrollador web apasionado con más de 5 años de experiencia 
            creando aplicaciones web modernas y escalables. Me especializo en 
            tecnologías frontend como React, Vue.js y CSS avanzado.
          </p>
          <p>
            Creo en el poder del diseño centrado en el usuario y en escribir 
            código que sea mantenible y escalable. Cuando no estoy programando, 
            me gusta contribuir a proyectos open source y compartir conocimientos 
            con la comunidad.
          </p>
          
          <div class="skills">
            <h4>Habilidades principales</h4>
            <div class="skills__grid">
              <div class="skill">
                <span class="skill__name">JavaScript</span>
                <div class="skill__bar">
                  <div class="skill__progress" style="width: 90%"></div>
                </div>
              </div>
              <div class="skill">
                <span class="skill__name">React</span>
                <div class="skill__bar">
                  <div class="skill__progress" style="width: 85%"></div>
                </div>
              </div>
              <div class="skill">
                <span class="skill__name">CSS/SCSS</span>
                <div class="skill__bar">
                  <div class="skill__progress" style="width: 95%"></div>
                </div>
              </div>
              <div class="skill">
                <span class="skill__name">Node.js</span>
                <div class="skill__bar">
                  <div class="skill__progress" style="width: 80%"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="about__image">
          <img src="assets/images/about.jpg" alt="Juan trabajando" loading="lazy">
        </div>
      </div>
    </div>
  </section>

  <!-- Projects Section -->
  <section id="projects" class="section section--gray">
    <div class="container">
      <div class="section__header">
        <h2 class="section__title">Proyectos destacados</h2>
        <p class="section__subtitle">Algunos de mis trabajos más recientes</p>
      </div>
      
      <div class="o-grid o-grid--3-col">
        <article class="c-card">
          <img src="assets/images/project1.jpg" alt="Proyecto 1" class="c-card__image" loading="lazy">
          <div class="c-card__content">
            <h3 class="c-card__title">E-commerce Platform</h3>
            <p class="c-card__description">
              Plataforma de comercio electrónico completa con React, Node.js y Stripe.
            </p>
            <div class="c-card__tags">
              <span class="c-card__tag">React</span>
              <span class="c-card__tag">Node.js</span>
              <span class="c-card__tag">MongoDB</span>
            </div>
            <div class="c-card__actions">
              <a href="#" class="c-button c-button--sm">Ver proyecto</a>
              <a href="#" class="c-button c-button--outline c-button--sm">Código</a>
            </div>
          </div>
        </article>
        
        <article class="c-card">
          <img src="assets/images/project2.jpg" alt="Proyecto 2" class="c-card__image" loading="lazy">
          <div class="c-card__content">
            <h3 class="c-card__title">Task Management App</h3>
            <p class="c-card__description">
              Aplicación de gestión de tareas con drag & drop y sincronización en tiempo real.
            </p>
            <div class="c-card__tags">
              <span class="c-card__tag">Vue.js</span>
              <span class="c-card__tag">Firebase</span>
              <span class="c-card__tag">PWA</span>
            </div>
            <div class="c-card__actions">
              <a href="#" class="c-button c-button--sm">Ver proyecto</a>
              <a href="#" class="c-button c-button--outline c-button--sm">Código</a>
            </div>
          </div>
        </article>
        
        <article class="c-card">
          <img src="assets/images/project3.jpg" alt="Proyecto 3" class="c-card__image" loading="lazy">
          <div class="c-card__content">
            <h3 class="c-card__title">Design System</h3>
            <p class="c-card__description">
              Sistema de diseño completo con componentes reutilizables y documentación.
            </p>
            <div class="c-card__tags">
              <span class="c-card__tag">CSS</span>
              <span class="c-card__tag">Storybook</span>
              <span class="c-card__tag">Figma</span>
            </div>
            <div class="c-card__actions">
              <a href="#" class="c-button c-button--sm">Ver proyecto</a>
              <a href="#" class="c-button c-button--outline c-button--sm">Código</a>
            </div>
          </div>
        </article>
      </div>
      
      <div class="text-center" style="margin-top: 3rem;">
        <a href="projects.html" class="c-button c-button--primary c-button--lg">
          Ver todos los proyectos
        </a>
      </div>
    </div>
  </section>

  <!-- Contact Section -->
  <section id="contact" class="section">
    <div class="container">
      <div class="section__header">
        <h2 class="section__title">Hablemos</h2>
        <p class="section__subtitle">¿Tienes un proyecto en mente? ¡Conversemos!</p>
      </div>
      
      <div class="o-grid o-grid--2-col">
        <div class="contact__info">
          <h3>Información de contacto</h3>
          <div class="contact__item">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path>
              <polyline points="22,6 12,13 2,6"></polyline>
            </svg>
            <div>
              <strong>Email</strong>
              <p>juan@example.com</p>
            </div>
          </div>
          
          <div class="contact__item">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
              <circle cx="12" cy="10" r="3"></circle>
            </svg>
            <div>
              <strong>Ubicación</strong>
              <p>Madrid, España</p>
            </div>
          </div>
        </div>
        
        <form class="contact__form">
          <div class="form-group">
            <label for="name">Nombre</label>
            <input type="text" id="name" name="name" required>
          </div>
          
          <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" name="email" required>
          </div>
          
          <div class="form-group">
            <label for="message">Mensaje</label>
            <textarea id="message" name="message" rows="5" required></textarea>
          </div>
          
          <button type="submit" class="c-button c-button--primary c-button--lg">
            Enviar mensaje
          </button>
        </form>
      </div>
    </div>
  </section>

  <!-- Footer -->
  <footer class="footer">
    <div class="container">
      <div class="footer__content">
        <div class="footer__brand">
          <h3>Juan Pérez</h3>
          <p>Desarrollador web apasionado por crear experiencias digitales excepcionales.</p>
        </div>
        
        <div class="footer__links">
          <h4>Enlaces</h4>
          <ul>
            <li><a href="#about">Sobre mí</a></li>
            <li><a href="#projects">Proyectos</a></li>
            <li><a href="#blog">Blog</a></li>
            <li><a href="#contact">Contacto</a></li>
          </ul>
        </div>
        
        <div class="footer__social">
          <h4>Sígueme</h4>
          <div class="social-links">
            <a href="#" aria-label="GitHub">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
            </a>
            <a href="#" aria-label="LinkedIn">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
              </svg>
            </a>
            <a href="#" aria-label="Twitter">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
                <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
              </svg>
            </a>
          </div>
        </div>
      </div>
      
      <div class="footer__bottom">
        <p>&copy; 2024 Juan Pérez. Todos los derechos reservados.</p>
      </div>
    </div>
  </footer>

  <!-- Scripts -->
  <script src="assets/js/main.js" defer></script>
</body>
</html>
```

### JavaScript Interactivo

```javascript
// assets/js/main.js
document.addEventListener('DOMContentLoaded', function() {
  // Navigation mobile toggle
  const mobileToggle = document.querySelector('.c-navigation__mobile-toggle');
  const mobileMenu = document.querySelector('.c-navigation__mobile-menu');
  
  if (mobileToggle && mobileMenu) {
    mobileToggle.addEventListener('click', function() {
      mobileMenu.classList.toggle('is-open');
    });
  }
  
  // Smooth scrolling
  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
      e.preventDefault();
      const target = document.querySelector(this.getAttribute('href'));
      if (target) {
        target.scrollIntoView({
          behavior: 'smooth',
          block: 'start'
        });
      }
    });
  });
  
  // Intersection Observer for animations
  const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px'
  };
  
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('animate-in');
      }
    });
  }, observerOptions);
  
  // Observe elements for animation
  document.querySelectorAll('.c-card, .section').forEach(el => {
    observer.observe(el);
  });
  
  // Form handling
  const contactForm = document.querySelector('.contact__form');
  if (contactForm) {
    contactForm.addEventListener('submit', function(e) {
      e.preventDefault();
      
      const submitBtn = this.querySelector('button[type="submit"]');
      const originalText = submitBtn.textContent;
      
      // Show loading state
      submitBtn.classList.add('c-button--loading');
      submitBtn.textContent = 'Enviando...';
      
      // Simulate form submission
      setTimeout(() => {
        submitBtn.classList.remove('c-button--loading');
        submitBtn.textContent = '¡Mensaje enviado!';
        submitBtn.classList.add('c-button--success');
        
        // Reset form
        this.reset();
        
        // Reset button after 3 seconds
        setTimeout(() => {
          submitBtn.textContent = originalText;
          submitBtn.classList.remove('c-button--success');
        }, 3000);
      }, 2000);
    });
  }
  
  // Theme toggle (bonus feature)
  const themeToggle = document.createElement('button');
  themeToggle.innerHTML = '🌓';
  themeToggle.className = 'theme-toggle';
  themeToggle.setAttribute('aria-label', 'Toggle theme');
  
  document.body.appendChild(themeToggle);
  
  themeToggle.addEventListener('click', function() {
    document.documentElement.classList.toggle('dark-theme');
    localStorage.setItem('theme', 
      document.documentElement.classList.contains('dark-theme') ? 'dark' : 'light'
    );
  });
  
  // Load saved theme
  if (localStorage.getItem('theme') === 'dark') {
    document.documentElement.classList.add('dark-theme');
  }
});
```

### Tema oscuro

```scss
// Dark theme variables
.dark-theme {
  --color-gray-50: #1f2937;
  --color-gray-100: #374151;
  --color-gray-200: #4b5563;
  --color-gray-700: #d1d5db;
  --color-gray-800: #e5e7eb;
  --color-gray-900: #f9fafb;
  
  --color-dark-bg: #0f172a;
  --color-dark-surface: #1e293b;
  --color-dark-text: #f1f5f9;
  
  body {
    background: var(--color-dark-bg);
    color: var(--color-dark-text);
  }
  
  .c-card {
    background: var(--color-dark-surface);
    color: var(--color-dark-text);
  }
  
  .c-navigation {
    background: rgba(30, 41, 59, 0.95);
    border-bottom-color: var(--color-gray-200);
  }
  
  .c-button--outline {
    color: var(--color-primary-light);
    border-color: var(--color-primary-light);
  }
}
```

## Optimización y Performance

### CSS Crítico

```css
/* assets/css/critical.css - CSS crítico inline */
body{margin:0;font-family:'Inter',sans-serif}
.c-navigation{position:fixed;top:0;left:0;right:0;background:rgba(255,255,255,.95);z-index:1000}
.c-hero{min-height:100vh;display:flex;align-items:center;background:linear-gradient(135deg,#a5b4fc 0%,#67e8f9 100%)}
/* ... más CSS crítico ... */
```

### Build process

```javascript
// build.js
const fs = require('fs');
const path = require('path');
const postcss = require('postcss');
const cssnano = require('cssnano');
const autoprefixer = require('autoprefixer');

async function buildCSS() {
  const input = fs.readFileSync('css/main.scss', 'utf8');
  
  const result = await postcss([
    autoprefixer,
    cssnano({
      preset: 'advanced'
    })
  ]).process(input, {
    from: 'css/main.scss',
    to: 'assets/css/main.css'
  });
  
  fs.writeFileSync('assets/css/main.css', result.css);
  
  console.log(`CSS built: ${result.css.length} bytes`);
}

buildCSS();
```

### Optimización de imágenes

```html
<!-- Imágenes responsivas -->
<picture>
  <source srcset="assets/images/hero.avif" type="image/avif">
  <source srcset="assets/images/hero.webp" type="image/webp">
  <img src="assets/images/hero.jpg" 
       alt="Hero image" 
       loading="lazy"
       decoding="async">
</picture>
```

## Testing y QA

### Lighthouse audit

```bash
# Ejecutar Lighthouse
lighthouse http://localhost:3000 --output html --output-path ./reports/lighthouse.html

# Solo performance
lighthouse http://localhost:3000 --only-categories performance
```

### CSS testing

```javascript
// Verificar que los componentes se rendericen correctamente
describe('Button Component', () => {
  it('should render primary button', () => {
    const button = document.createElement('button');
    button.className = 'c-button c-button--primary';
    button.textContent = 'Click me';
    
    document.body.appendChild(button);
    
    expect(button).toBeVisible();
    expect(button).toHaveStyle({
      backgroundColor: 'var(--color-primary)'
    });
  });
});
```

## Despliegue

### Netlify

```toml
# netlify.toml
[build]
  publish = "."
  command = "npm run build"

[build.environment]
  NODE_VERSION = "18"

[[redirects]]
  from = "/api/*"
  to = "https://api.example.com/:splat"
  status = 200

[[headers]]
  for = "/assets/css/*.css"
  [headers.values]
    Cache-Control = "public, max-age=31536000"

[[headers]]
  for = "/assets/js/*.js"
  [headers.values]
    Cache-Control = "public, max-age=31536000"

[[headers]]
  for = "/assets/images/*"
  [headers.values]
    Cache-Control = "public, max-age=31536000"
```

### GitHub Actions

```yaml
# .github/workflows/deploy.yml
name: Deploy to Netlify
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Node.js
        uses: actions/setup-node@v3
        with:
          node-version: '18'
          cache: 'npm'
      
      - name: Install dependencies
        run: npm ci
      
      - name: Build
        run: npm run build
      
      - name: Deploy to Netlify
        uses: nwtgck/actions-netlify@v2
        with:
          publish-dir: './dist'
          production-branch: main
          github-token: ${{ secrets.GITHUB_TOKEN }}
          deploy-message: "Deploy from GitHub Actions"
          enable-pull-request-comment: true
          enable-commit-comment: true
          overwrites-pull-request-comment: true
        env:
          NETLIFY_AUTH_TOKEN: ${{ secrets.NETLIFY_AUTH_TOKEN }}
          NETLIFY_SITE_ID: ${{ secrets.NETLIFY_SITE_ID }}
        timeout-minutes: 1
```

## Documentación

### README.md

```markdown
# Portfolio - Juan Pérez

Sitio web portfolio personal construido con tecnologías modernas.

## 🚀 Características

- ✅ Diseño responsive (móvil-first)
- ✅ Performance optimizada (Core Web Vitals)
- ✅ Arquitectura CSS escalable (ITCSS + BEM)
- ✅ Componentes reutilizables
- ✅ Tema oscuro/claro
- ✅ Animaciones fluidas
- ✅ Accesibilidad WCAG 2.1

## 🛠️ Tecnologías

- **HTML5** - Semántico y accesible
- **CSS3** - Variables, Grid, Flexbox, Animaciones
- **JavaScript** - ES6+, Intersection Observer
- **Sass/SCSS** - Preprocesador CSS
- **PostCSS** - Autoprefixer, CSSNano
- **Netlify** - Despliegue y hosting

## 📦 Instalación

```bash
# Clonar repositorio
git clone https://github.com/juanperez/portfolio.git
cd portfolio

# Instalar dependencias
npm install

# Desarrollar
npm run dev

# Build para producción
npm run build

# Preview producción
npm run preview
```

## 🎨 Arquitectura CSS

```
css/
├── 01-settings/     # Variables y configuración
├── 02-tools/        # Mixins y funciones
├── 03-generic/      # Resets y base
├── 04-elements/     # Estilos de elementos HTML
├── 05-objects/      # Layout y estructuras
├── 06-components/   # Componentes UI
└── 07-utilities/    # Utilidades y helpers
```

## 📱 Navegadores soportados

- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## 📈 Performance

- **Lighthouse Score**: 95+ en todas las métricas
- **First Contentful Paint**: < 1.5s
- **Largest Contentful Paint**: < 2.5s
- **Cumulative Layout Shift**: < 0.1
- **Bundle size**: < 50KB gzipped

## 🤝 Contribuir

1. Fork el proyecto
2. Crea tu rama (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## 📞 Contacto

Juan Pérez - [@juanperez](https://twitter.com/juanperez) - juan@example.com

Link del proyecto: [https://juanperez.dev](https://juanperez.dev)
```

## Conclusión

Este proyecto final demuestra la aplicación práctica de todos los conceptos aprendidos:

- **CSS Avanzado**: Variables, funciones, mixins, Grid, Flexbox
- **Arquitectura**: ITCSS, BEM, componentes modulares
- **Performance**: Optimización, minificación, CSS crítico
- **Responsive**: Móvil-first, breakpoints estratégicos
- **Accesibilidad**: Semántica, ARIA, navegación por teclado
- **Modern CSS**: Custom Properties, containment, will-change
- **Herramientas**: Sass, PostCSS, Autoprefixer, CSSNano
- **Despliegue**: CDN, caching, compression

El resultado es un sitio web moderno, performante y mantenible que sirve como excelente ejemplo de las mejores prácticas en desarrollo CSS profesional.