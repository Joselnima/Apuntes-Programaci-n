# Módulo 12 - Responsive Design

En este módulo aprenderás a crear diseños que se adaptan a diferentes tamaños de pantalla. El responsive design es fundamental para la web moderna donde los usuarios acceden desde dispositivos de todo tipo.

## Conceptos básicos del responsive design

### Viewport meta tag

```html
<!-- Esencial para responsive design -->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<!-- width=device-width: ancho del dispositivo -->
<!-- initial-scale=1.0: zoom inicial 100% -->
```

### Media queries

```css
/* Sintaxis básica */
@media media-type and (condition) {
  /* Reglas CSS */
}

/* Ejemplos comunes */
@media screen and (max-width: 768px) {
  /* Móvil */
}

@media screen and (min-width: 769px) and (max-width: 1024px) {
  /* Tablet */
}

@media screen and (min-width: 1025px) {
  /* Desktop */
}
```

## Breakpoints comunes

### Breakpoints estándar

```css
/* Extra small devices (phones, < 576px) */
@media (max-width: 575.98px) { }

/* Small devices (phones, 576px and up) */
@media (min-width: 576px) and (max-width: 767.98px) { }

/* Medium devices (tablets, 768px and up) */
@media (min-width: 768px) and (max-width: 991.98px) { }

/* Large devices (desktops, 992px and up) */
@media (min-width: 992px) and (max-width: 1199.98px) { }

/* Extra large devices (large desktops, 1200px and up) */
@media (min-width: 1200px) { }
```

### Breakpoints móviles primero

```css
/* Móvil primero (recomendado) */
.container {
  width: 100%;
  padding: 15px;
}

/* Tablet */
@media (min-width: 768px) {
  .container {
    width: 750px;
    margin: 0 auto;
  }
}

/* Desktop */
@media (min-width: 992px) {
  .container {
    width: 970px;
  }
}

/* Large desktop */
@media (min-width: 1200px) {
  .container {
    width: 1170px;
  }
}
```

## Unidades responsivas

### Unidades relativas

```css
/* Viewport units */
.element {
  width: 50vw;    /* 50% del ancho del viewport */
  height: 100vh;  /* 100% del alto del viewport */
  font-size: 5vw; /* 5% del ancho del viewport */
  padding: 2vh;   /* 2% del alto del viewport */
}

/* REM units (recomendado) */
html {
  font-size: 16px; /* Base */
}

.element {
  font-size: 1rem;    /* 16px */
  padding: 1.5rem;    /* 24px */
  margin: 2rem;      /* 32px */
}

/* EM units */
.parent {
  font-size: 1.2em; /* Relativo al padre */
}

.child {
  font-size: 1.5em; /* 1.5 * 1.2em del abuelo */
}
```

### Funciones calc()

```css
.element {
  width: calc(100% - 20px);        /* 100% menos 20px */
  height: calc(100vh - 80px);      /* Viewport height menos header */
  font-size: calc(1rem + 0.5vw);   /* Rem + viewport unit */
  margin: calc(2rem + 10px);       /* Combinación */
}
```

## Layouts responsivos

### Flexbox responsive

```css
.flex-container {
  display: flex;
  flex-direction: column; /* Móvil: columna */
  gap: 1rem;
}

@media (min-width: 768px) {
  .flex-container {
    flex-direction: row;  /* Desktop: fila */
  }
  
  .flex-item {
    flex: 1; /* Items iguales */
  }
}
```

### CSS Grid responsive

```css
.grid-container {
  display: grid;
  grid-template-columns: 1fr; /* Móvil: 1 columna */
  gap: 1rem;
}

@media (min-width: 768px) {
  .grid-container {
    grid-template-columns: repeat(2, 1fr); /* Tablet: 2 columnas */
  }
}

@media (min-width: 1024px) {
  .grid-container {
    grid-template-columns: repeat(3, 1fr); /* Desktop: 3 columnas */
  }
}
```

### Grid con auto-fit

```css
.auto-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}
/* Se adapta automáticamente al espacio disponible */
```

## Imágenes responsivas

### Imágenes fluidas

```css
.responsive-image {
  max-width: 100%;    /* Nunca más grande que el contenedor */
  height: auto;       /* Mantener proporción */
}
```

### Picture element

```html
<picture>
  <!-- WebP para navegadores modernos -->
  <source srcset="image-1200.webp" media="(min-width: 1200px)">
  <source srcset="image-768.webp" media="(min-width: 768px)">
  <source srcset="image-480.webp" media="(min-width: 480px)">
  
  <!-- Fallback para navegadores antiguos -->
  <img src="image-480.jpg" alt="Imagen responsiva">
</picture>
```

### Srcset y sizes

```html
<img src="small.jpg" 
     srcset="small.jpg 480w,
             medium.jpg 768w,
             large.jpg 1200w"
     sizes="(max-width: 480px) 100vw,
            (max-width: 768px) 50vw,
            33vw"
     alt="Imagen con srcset">
```

## Tipografía responsiva

### Font-size fluido

```css
/* Usando viewport units */
.responsive-text {
  font-size: clamp(1rem, 2.5vw, 2rem);
  /* Mínimo 1rem, ideal 2.5vw, máximo 2rem */
}

/* Escala modular */
:root {
  --font-size-base: clamp(1rem, 1.5vw, 1.25rem);
  --font-size-lg: calc(var(--font-size-base) * 1.25);
  --font-size-xl: calc(var(--font-size-base) * 1.5);
}

h1 { font-size: var(--font-size-xl); }
h2 { font-size: var(--font-size-lg); }
p { font-size: var(--font-size-base); }
```

### Line-height responsivo

```css
.responsive-line-height {
  line-height: clamp(1.4, 1.8, 2);
  /* Se adapta según el tamaño de fuente */
}
```

## Navegación responsiva

### Menú hamburguesa

```html
<nav class="navbar">
  <div class="logo">Logo</div>
  <button class="menu-toggle" aria-label="Toggle menu">
    <span></span>
    <span></span>
    <span></span>
  </button>
  <ul class="nav-menu">
    <li><a href="#">Inicio</a></li>
    <li><a href="#">Acerca</a></li>
    <li><a href="#">Contacto</a></li>
  </ul>
</nav>
```

```css
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
}

.nav-menu {
  display: flex;
  list-style: none;
  gap: 2rem;
}

.menu-toggle {
  display: none;
  flex-direction: column;
  background: none;
  border: none;
  cursor: pointer;
}

.menu-toggle span {
  width: 25px;
  height: 3px;
  background: #333;
  margin: 3px 0;
  transition: 0.3s;
}

@media (max-width: 768px) {
  .menu-toggle {
    display: flex;
  }
  
  .nav-menu {
    position: fixed;
    top: 100%;
    left: 0;
    width: 100%;
    height: 0;
    flex-direction: column;
    background: white;
    overflow: hidden;
    transition: height 0.3s ease;
  }
  
  .nav-menu.open {
    height: 200px;
  }
  
  .nav-menu li {
    padding: 1rem;
    border-bottom: 1px solid #eee;
  }
}
```

## Componentes responsivos

### Tarjetas adaptativas

```css
.card {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 8px;
}

.card-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 4px;
}

@media (min-width: 768px) {
  .card {
    grid-template-columns: 200px 1fr;
    align-items: start;
  }
  
  .card-image {
    height: 150px;
  }
}
```

### Formularios responsivos

```css
.form-container {
  display: grid;
  gap: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

@media (min-width: 768px) {
  .form-container {
    grid-template-columns: 1fr 1fr;
  }
  
  .form-group.full-width {
    grid-column: 1 / -1;
  }
}
```

## Técnicas avanzadas

### Container queries (CSS Containment)

```css
/* Próximamente en CSS */
@container (min-width: 400px) {
  .card {
    display: grid;
    grid-template-columns: 150px 1fr;
  }
}
```

### Aspect ratio

```css
.aspect-ratio-box {
  aspect-ratio: 16 / 9; /* Relación 16:9 */
  background: #f0f0f0;
  /* Funciona sin especificar height */
}
```

### CSS locks

```css
/* Tamaño de fuente que no escala más allá de ciertos límites */
.locked-text {
  font-size: clamp(1rem, 2vw + 1rem, 2rem);
}
```

## Testing y debugging

### Herramientas de desarrollo

```css
/* Simular diferentes dispositivos en DevTools */
/* - Toggle device toolbar */
/* - Responsive design mode */
/* - Network throttling */
/* - Touch simulation */
```

### Media query debugging

```css
/* Añadir indicadores visuales */
@media (max-width: 768px) {
  body::before {
    content: "Móvil";
    position: fixed;
    top: 0;
    left: 0;
    background: red;
    color: white;
    padding: 5px;
    z-index: 9999;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  body::before {
    content: "Tablet";
    background: orange;
  }
}

@media (min-width: 1025px) {
  body::before {
    content: "Desktop";
    background: green;
  }
}
```

## Mejores prácticas

### Performance

```css
/* Evitar layout shifts */
.responsive-image {
  aspect-ratio: 16 / 9; /* Define proporción antes de cargar */
}

/* Lazy loading */
<img loading="lazy" src="image.jpg" alt="Lazy loaded image">
```

### Accesibilidad

```css
/* No depender solo de hover */
@media (hover: hover) {
  .hover-only {
    /* Solo para dispositivos con hover */
  }
}

/* Respetar preferencias de movimiento */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    transition-duration: 0.01ms !important;
  }
}
```

### SEO y performance

```css
/* Critical CSS primero */
<link rel="preload" href="critical.css" as="style" onload="this.onload=null;this.rel='stylesheet'">
<noscript><link rel="stylesheet" href="critical.css"></noscript>

/* Cargar CSS condicionalmente */
<link rel="stylesheet" media="screen and (min-width: 768px)" href="desktop.css">
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Responsive Design CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="debug-indicator"></div>
  
  <nav class="navbar">
    <div class="logo">RD Master</div>
    <button class="menu-toggle" aria-label="Toggle menu">
      <span></span>
      <span></span>
      <span></span>
    </button>
    <ul class="nav-menu">
      <li><a href="#typography">Tipografía</a></li>
      <li><a href="#layouts">Layouts</a></li>
      <li><a href="#images">Imágenes</a></li>
      <li><a href="#components">Componentes</a></li>
    </ul>
  </nav>

  <header class="hero">
    <div class="hero-content">
      <h1>Responsive Design</h1>
      <p>Diseños que se adaptan a todos los dispositivos</p>
      <button class="cta-button">Comenzar</button>
    </div>
  </header>

  <main class="main">
    <section id="typography" class="section">
      <h2>Tipografía Fluida</h2>
      <div class="typography-examples">
        <p class="fluid-text">Este texto cambia de tamaño según la pantalla</p>
        <p class="locked-text">Texto con límites (clamp)</p>
        <div class="text-scale">
          <h3>H3 fluido</h3>
          <h4>H4 fluido</h4>
          <p>Párrafo con line-height responsivo</p>
        </div>
      </div>
    </section>

    <section id="layouts" class="section">
      <h2>Layouts Responsivos</h2>
      <div class="layout-examples">
        <div class="flex-layout">
          <div class="layout-item">Flex Item 1</div>
          <div class="layout-item">Flex Item 2</div>
          <div class="layout-item">Flex Item 3</div>
        </div>
        
        <div class="grid-layout">
          <div class="grid-item">Grid Item 1</div>
          <div class="grid-item">Grid Item 2</div>
          <div class="grid-item">Grid Item 3</div>
          <div class="grid-item">Grid Item 4</div>
        </div>
        
        <div class="auto-grid">
          <div class="auto-item">Auto 1</div>
          <div class="auto-item">Auto 2</div>
          <div class="auto-item">Auto 3</div>
          <div class="auto-item">Auto 4</div>
          <div class="auto-item">Auto 5</div>
          <div class="auto-item">Auto 6</div>
        </div>
      </div>
    </section>

    <section id="images" class="section">
      <h2>Imágenes Responsivas</h2>
      <div class="image-examples">
        <img class="responsive-image" src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=800" alt="Responsive image">
        
        <picture class="picture-example">
          <source srcset="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=400" media="(max-width: 480px)">
          <source srcset="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=800" media="(max-width: 768px)">
          <img src="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=1200" alt="Picture element">
        </picture>
        
        <div class="aspect-ratio-box">
          <p>Caja con aspect ratio 16:9</p>
        </div>
      </div>
    </section>

    <section id="components" class="section">
      <h2>Componentes Responsivos</h2>
      <div class="component-examples">
        <div class="responsive-card">
          <img src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300" alt="Card image">
          <div class="card-content">
            <h3>Tarjeta Responsiva</h3>
            <p>Esta tarjeta cambia su layout en diferentes tamaños de pantalla.</p>
            <button class="card-button">Leer más</button>
          </div>
        </div>
        
        <form class="responsive-form">
          <div class="form-group">
            <label for="name">Nombre</label>
            <input type="text" id="name" placeholder="Tu nombre">
          </div>
          <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" placeholder="tu@email.com">
          </div>
          <div class="form-group full-width">
            <label for="message">Mensaje</label>
            <textarea id="message" placeholder="Tu mensaje"></textarea>
          </div>
          <button type="submit" class="submit-button">Enviar</button>
        </form>
      </div>
    </section>
  </main>

  <footer class="footer">
    <p>&copy; 2024 Responsive Design Master</p>
  </footer>

  <script src="script.js"></script>
</body>
</html>
```

```css
/* Variables CSS */
:root {
  --primary: #3498db;
  --secondary: #2ecc71;
  --accent: #e74c3c;
  --dark: #2c3e50;
  --light: #ecf0f1;
  --gray: #95a5a6;
  
  /* Espaciado fluido */
  --space-sm: clamp(0.5rem, 2vw, 1rem);
  --space-md: clamp(1rem, 4vw, 2rem);
  --space-lg: clamp(1.5rem, 6vw, 3rem);
  
  /* Tipografía fluida */
  --font-size-base: clamp(1rem, 2.5vw, 1.25rem);
  --font-size-lg: calc(var(--font-size-base) * 1.25);
  --font-size-xl: calc(var(--font-size-base) * 1.5);
  --font-size-xxl: calc(var(--font-size-base) * 2);
}

/* Reset */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  line-height: 1.6;
  color: var(--dark);
  background: var(--light);
  font-size: var(--font-size-base);
}

/* Debug indicator */
.debug-indicator::before {
  content: "Desktop";
  position: fixed;
  top: 10px;
  right: 10px;
  background: green;
  color: white;
  padding: 5px 10px;
  border-radius: 4px;
  font-size: 12px;
  z-index: 9999;
}

@media (max-width: 768px) {
  .debug-indicator::before {
    content: "Móvil";
    background: red;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .debug-indicator::before {
    content: "Tablet";
    background: orange;
  }
}

/* Navbar */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--space-md);
  background: white;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.logo {
  font-size: var(--font-size-lg);
  font-weight: bold;
  color: var(--primary);
}

.nav-menu {
  display: flex;
  list-style: none;
  gap: 2rem;
}

.nav-menu a {
  text-decoration: none;
  color: var(--dark);
  font-weight: 500;
  transition: color 0.3s ease;
}

.nav-menu a:hover {
  color: var(--primary);
}

.menu-toggle {
  display: none;
  flex-direction: column;
  background: none;
  border: none;
  cursor: pointer;
  padding: 5px;
}

.menu-toggle span {
  width: 25px;
  height: 3px;
  background: var(--dark);
  margin: 3px 0;
  transition: 0.3s;
}

.menu-toggle.active span:nth-child(1) {
  transform: rotate(-45deg) translate(-5px, 5px);
}

.menu-toggle.active span:nth-child(2) {
  opacity: 0;
}

.menu-toggle.active span:nth-child(3) {
  transform: rotate(45deg) translate(-5px, -5px);
}

/* Hero */
.hero {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: white;
  padding: var(--space-lg);
  text-align: center;
  min-height: 50vh;
  display: flex;
  align-items: center;
}

.hero-content h1 {
  font-size: var(--font-size-xxl);
  margin-bottom: var(--space-sm);
}

.hero-content p {
  font-size: var(--font-size-lg);
  margin-bottom: var(--space-md);
  opacity: 0.9;
}

.cta-button {
  padding: 12px 24px;
  background: white;
  color: var(--primary);
  border: none;
  border-radius: 6px;
  font-size: var(--font-size-base);
  cursor: pointer;
  transition: all 0.3s ease;
}

.cta-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0,0,0,0.2);
}

/* Secciones */
.section {
  padding: var(--space-lg);
  max-width: 1200px;
  margin: 0 auto;
}

.section h2 {
  font-size: var(--font-size-xl);
  text-align: center;
  margin-bottom: var(--space-lg);
  color: var(--dark);
  position: relative;
}

.section h2::after {
  content: '';
  position: absolute;
  bottom: -10px;
  left: 50%;
  transform: translateX(-50%);
  width: 80px;
  height: 4px;
  background: linear-gradient(to right, var(--primary), var(--secondary));
  border-radius: 2px;
}

/* Typography */
.typography-examples {
  display: grid;
  gap: var(--space-md);
}

.fluid-text {
  font-size: clamp(1.2rem, 4vw, 2rem);
  color: var(--primary);
  text-align: center;
}

.locked-text {
  font-size: clamp(1rem, 2vw + 1rem, 2rem);
  color: var(--secondary);
  text-align: center;
}

.text-scale h3 {
  font-size: var(--font-size-lg);
  margin-bottom: var(--space-sm);
}

.text-scale h4 {
  font-size: calc(var(--font-size-base) * 1.1);
  margin-bottom: var(--space-sm);
}

.text-scale p {
  line-height: clamp(1.4, 1.8, 2);
}

/* Layouts */
.layout-examples {
  display: grid;
  gap: var(--space-lg);
}

.flex-layout {
  display: flex;
  flex-direction: column;
  gap: var(--space-sm);
  padding: var(--space-md);
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.layout-item {
  padding: var(--space-sm);
  background: var(--primary);
  color: white;
  border-radius: 4px;
  text-align: center;
}

.grid-layout {
  display: grid;
  grid-template-columns: 1fr;
  gap: var(--space-sm);
  padding: var(--space-md);
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.grid-item {
  padding: var(--space-sm);
  background: var(--secondary);
  color: white;
  border-radius: 4px;
  text-align: center;
}

.auto-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: var(--space-sm);
  padding: var(--space-md);
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.auto-item {
  padding: var(--space-sm);
  background: var(--accent);
  color: white;
  border-radius: 4px;
  text-align: center;
}

/* Images */
.image-examples {
  display: grid;
  gap: var(--space-md);
}

.responsive-image {
  width: 100%;
  height: auto;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.picture-example img {
  width: 100%;
  height: auto;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.aspect-ratio-box {
  aspect-ratio: 16 / 9;
  background: linear-gradient(45deg, var(--primary), var(--secondary));
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  font-weight: bold;
}

/* Components */
.component-examples {
  display: grid;
  gap: var(--space-lg);
}

.responsive-card {
  display: grid;
  grid-template-columns: 1fr;
  gap: var(--space-md);
  padding: var(--space-md);
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.responsive-card img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 4px;
}

.card-content h3 {
  margin-bottom: var(--space-sm);
  color: var(--primary);
}

.card-button {
  padding: 8px 16px;
  background: var(--primary);
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.card-button:hover {
  background: #2980b9;
}

.responsive-form {
  display: grid;
  gap: var(--space-md);
  padding: var(--space-md);
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 5px;
  font-weight: 600;
  color: var(--dark);
}

.form-group input,
.form-group textarea {
  padding: 12px;
  border: 2px solid #ddd;
  border-radius: 4px;
  font-size: var(--font-size-base);
  transition: border-color 0.3s ease;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: var(--primary);
}

.form-group textarea {
  resize: vertical;
  min-height: 100px;
}

.submit-button {
  padding: 12px;
  background: var(--secondary);
  color: white;
  border: none;
  border-radius: 4px;
  font-size: var(--font-size-base);
  cursor: pointer;
  transition: background 0.3s ease;
}

.submit-button:hover {
  background: #27ae60;
}

/* Footer */
.footer {
  background: var(--dark);
  color: white;
  text-align: center;
  padding: var(--space-md);
}

/* Media Queries */
@media (max-width: 768px) {
  .navbar {
    position: relative;
  }
  
  .menu-toggle {
    display: flex;
  }
  
  .nav-menu {
    position: absolute;
    top: 100%;
    left: 0;
    width: 100%;
    background: white;
    flex-direction: column;
    padding: var(--space-md);
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    transform: translateY(-100%);
    opacity: 0;
    visibility: hidden;
    transition: all 0.3s ease;
  }
  
  .nav-menu.active {
    transform: translateY(0);
    opacity: 1;
    visibility: visible;
  }
  
  .nav-menu li {
    padding: var(--space-sm) 0;
    border-bottom: 1px solid #eee;
  }
  
  .hero {
    padding: var(--space-md);
    min-height: 40vh;
  }
  
  .hero-content h1 {
    font-size: var(--font-size-xl);
  }
  
  .flex-layout {
    flex-direction: column;
  }
  
  .grid-layout {
    grid-template-columns: 1fr;
  }
  
  .auto-grid {
    grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  }
  
  .responsive-card {
    grid-template-columns: 1fr;
  }
  
  .responsive-form {
    grid-template-columns: 1fr;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .flex-layout {
    flex-direction: row;
    flex-wrap: wrap;
  }
  
  .grid-layout {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .responsive-card {
    grid-template-columns: 200px 1fr;
  }
  
  .responsive-form {
    grid-template-columns: 1fr 1fr;
  }
  
  .form-group.full-width {
    grid-column: 1 / -1;
  }
}

@media (min-width: 1025px) {
  .flex-layout {
    flex-direction: row;
  }
  
  .grid-layout {
    grid-template-columns: repeat(3, 1fr);
  }
  
  .responsive-card {
    grid-template-columns: 250px 1fr;
  }
  
  .responsive-form {
    grid-template-columns: 1fr 1fr;
  }
}

/* Accesibilidad */
@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    animation-duration: 0.01ms !important;
    transition-duration: 0.01ms !important;
  }
}
```

```javascript
// script.js
document.addEventListener('DOMContentLoaded', function() {
  // Toggle menú móvil
  const menuToggle = document.querySelector('.menu-toggle');
  const navMenu = document.querySelector('.nav-menu');
  
  menuToggle.addEventListener('click', function() {
    navMenu.classList.toggle('active');
    menuToggle.classList.toggle('active');
  });
  
  // Cerrar menú al hacer click en un enlace
  const navLinks = document.querySelectorAll('.nav-menu a');
  navLinks.forEach(link => {
    link.addEventListener('click', function() {
      navMenu.classList.remove('active');
      menuToggle.classList.remove('active');
    });
  });
  
  // Smooth scroll para anclas
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
  
  // Mostrar tamaño de ventana en tiempo real (para debugging)
  function updateWindowSize() {
    const width = window.innerWidth;
    const height = window.innerHeight;
    console.log(`Viewport: ${width}x${height}`);
  }
  
  window.addEventListener('resize', updateWindowSize);
  updateWindowSize(); // Llamada inicial
});
```

## Resumen

El responsive design es esencial para la web moderna:

- ✅ **Media queries**: `@media` para diferentes tamaños de pantalla
- ✅ **Unidades fluidas**: `vw`, `vh`, `rem`, `clamp()` para escalabilidad
- ✅ **Imágenes responsivas**: `srcset`, `<picture>`, `aspect-ratio`
- ✅ **Layouts adaptativos**: Flexbox y Grid responsive
- ✅ **Navegación móvil**: Menús hamburguesa y navegación adaptativa

El responsive design no es opcional, es obligatorio. Todos los sitios modernos deben funcionar perfectamente en todos los dispositivos.