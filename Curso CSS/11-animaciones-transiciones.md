# Módulo 11 - Animaciones y transiciones

En este módulo aprenderás a crear animaciones y transiciones suaves que harán que tus interfaces cobren vida. Las animaciones mejoran la experiencia del usuario y añaden profesionalismo.

## Transiciones (transitions)

### Propiedad básica transition

```css
.element {
  transition: property duration timing-function delay;
  
  /* Ejemplos */
  transition: background-color 0.3s ease;
  transition: all 0.5s ease-in-out;
  transition: transform 0.2s ease, opacity 0.3s ease;
}
```

### Propiedades individuales

```css
.element {
  transition-property: background-color, transform;  /* Propiedad(es) */
  transition-duration: 0.3s, 0.5s;                   /* Duración */
  transition-timing-function: ease, ease-in-out;     /* Función de tiempo */
  transition-delay: 0s, 0.2s;                       /* Retraso */
}
```

### Funciones de tiempo (timing functions)

```css
.element {
  transition-timing-function: ease;           /* Suave (default) */
  transition-timing-function: linear;         /* Lineal */
  transition-timing-function: ease-in;        /* Aceleración inicial */
  transition-timing-function: ease-out;       /* Desaceleración final */
  transition-timing-function: ease-in-out;    /* Aceleración y desaceleración */
  transition-timing-function: cubic-bezier(0.25, 0.46, 0.45, 0.94); /* Personalizada */
}
```

## Animaciones con @keyframes

### Sintaxis básica

```css
@keyframes nombre-animacion {
  from {
    /* Estado inicial */
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    /* Estado final */
    opacity: 1;
    transform: translateY(0);
  }
}

/* O con porcentajes */
@keyframes slide-in {
  0% {
    transform: translateX(-100%);
    opacity: 0;
  }
  50% {
    opacity: 0.5;
  }
  100% {
    transform: translateX(0);
    opacity: 1;
  }
}
```

### Aplicando animaciones

```css
.element {
  animation: name duration timing-function delay iteration-count direction fill-mode;
  
  /* Ejemplos */
  animation: slide-in 0.5s ease-out;
  animation: bounce 1s ease-in-out infinite;
  animation: fade-in 0.3s ease 0.2s 1 normal forwards;
}
```

### Propiedades individuales de animación

```css
.element {
  animation-name: slide-in;           /* Nombre de la animación */
  animation-duration: 0.5s;           /* Duración */
  animation-timing-function: ease;    /* Función de tiempo */
  animation-delay: 0.2s;              /* Retraso */
  animation-iteration-count: 1;       /* Número de repeticiones */
  animation-direction: normal;        /* Dirección */
  animation-fill-mode: forwards;       /* Modo de relleno */
  animation-play-state: running;       /* Estado de reproducción */
}
```

## Animaciones comunes

### Fade in/out

```css
@keyframes fade-in {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes fade-out {
  from { opacity: 1; }
  to { opacity: 0; }
}

.fade-in {
  animation: fade-in 0.5s ease;
}
```

### Slide animations

```css
@keyframes slide-in-left {
  from { transform: translateX(-100%); }
  to { transform: translateX(0); }
}

@keyframes slide-in-right {
  from { transform: translateX(100%); }
  to { transform: translateX(0); }
}

@keyframes slide-in-up {
  from { transform: translateY(100%); }
  to { transform: translateY(0); }
}

@keyframes slide-in-down {
  from { transform: translateY(-100%); }
  to { transform: translateY(0); }
}
```

### Scale animations

```css
@keyframes scale-in {
  from { transform: scale(0); }
  to { transform: scale(1); }
}

@keyframes scale-out {
  from { transform: scale(1); }
  to { transform: scale(0); }
}

@keyframes pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}
```

### Rotate animations

```css
@keyframes rotate-in {
  from { transform: rotate(-180deg); }
  to { transform: rotate(0deg); }
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.spinner {
  animation: spin 1s linear infinite;
}
```

### Bounce animation

```css
@keyframes bounce {
  0%, 20%, 53%, 80%, 100% {
    transform: translateY(0);
  }
  40%, 43% {
    transform: translateY(-30px);
  }
  70% {
    transform: translateY(-15px);
  }
  90% {
    transform: translateY(-4px);
  }
}

.bounce {
  animation: bounce 1s ease;
}
```

## Transforms para animaciones

### Transform 2D

```css
.element {
  transform: translate(x, y);     /* Mover */
  transform: rotate(angle);       /* Rotar */
  transform: scale(x, y);         /* Escalar */
  transform: skew(x-angle, y-angle); /* Inclinar */
  
  /* Combinadas */
  transform: translate(50px, 100px) rotate(45deg) scale(1.2);
}
```

### Transform 3D

```css
.element {
  transform: translate3d(x, y, z);    /* Mover en 3D */
  transform: rotateX(angle);         /* Rotar en X */
  transform: rotateY(angle);         /* Rotar en Y */
  transform: rotateZ(angle);         /* Rotar en Z */
  transform: scale3d(x, y, z);       /* Escalar en 3D */
  
  /* Perspective */
  transform: perspective(1000px) rotateY(45deg);
}
```

### Transform origin

```css
.element {
  transform-origin: center center;    /* Centro (default) */
  transform-origin: top left;         /* Esquina superior izquierda */
  transform-origin: 50px 100px;       /* Coordenadas específicas */
  transform-origin: 50% 100%;         /* Porcentajes */
}
```

## Performance y optimización

### Propiedades que no requieren repaint

```css
/* Buenas para animar (solo transform + opacity) */
.good-animation {
  transition: transform 0.3s ease, opacity 0.3s ease;
  /* Estas propiedades usan GPU */
}

/* Malas para animar (requieren repaint) */
.bad-animation {
  transition: width 0.3s ease, height 0.3s ease;
  /* Estas propiedades recalculan layout */
}
```

### Hardware acceleration

```css
.hardware-accelerated {
  transform: translateZ(0);  /* Fuerza GPU */
  backface-visibility: hidden; /* Evita flickering */
  perspective: 1000px;       /* Crea contexto 3D */
}
```

### Will-change

```css
.optimised-element {
  will-change: transform, opacity;  /* Indica qué animar */
}

/* Reset después de la animación */
.optimised-element:not(:hover) {
  will-change: auto;
}
```

## Animaciones complejas

### Loading spinner

```css
.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid #3498db;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
```

### Progress bar

```css
.progress-bar {
  width: 100%;
  height: 20px;
  background: #f0f0f0;
  border-radius: 10px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #3498db, #2ecc71);
  border-radius: 10px;
  animation: progress 2s ease-out;
}

@keyframes progress {
  from { width: 0%; }
  to { width: 75%; }
}
```

### Hover effects

```css
.card {
  transition: all 0.3s ease;
}

.card:hover {
  transform: translateY(-10px);
  box-shadow: 0 10px 30px rgba(0,0,0,0.2);
}

.button {
  transition: all 0.2s ease;
  position: relative;
  overflow: hidden;
}

.button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
  transition: left 0.5s ease;
}

.button:hover::before {
  left: 100%;
}
```

## Librerías de animación

### Animate.css

```html
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css">
```

```html
<div class="animate__animated animate__bounce">Contenido</div>
<div class="animate__animated animate__fadeInLeft">Contenido</div>
```

### Uso programático

```javascript
// Añadir clase de animación
element.classList.add('animate__animated', 'animate__bounce');

// Remover después de completar
element.addEventListener('animationend', () => {
  element.classList.remove('animate__animated', 'animate__bounce');
});
```

## Animaciones con JavaScript

### Control básico

```javascript
const element = document.querySelector('.animated');

function animate() {
  element.style.animation = 'slide-in 0.5s ease forwards';
}

// Control de animación
element.style.animationPlayState = 'paused';  // Pausar
element.style.animationPlayState = 'running'; // Reanudar
```

### Animaciones con Web Animations API

```javascript
const animation = element.animate([
  { transform: 'translateX(0)', opacity: 1 },
  { transform: 'translateX(100px)', opacity: 0 }
], {
  duration: 500,
  easing: 'ease-out',
  fill: 'forwards'
});

// Control
animation.pause();
animation.play();
animation.reverse();
animation.cancel();
```

## Accesibilidad en animaciones

### Respectar preferencias del usuario

```css
/* Respetar prefers-reduced-motion */
@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
```

### Animaciones opcionales

```css
/* Solo animar si no se prefiere movimiento reducido */
@media (prefers-reduced-motion: no-preference) {
  .animated-element {
    animation: complex-animation 2s ease;
  }
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Animaciones y Transiciones CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>Animaciones & Transiciones</h1>
      <p>Da vida a tus interfaces</p>
    </header>

    <main class="main">
      <section class="transitions-section">
        <h2>Transiciones</h2>
        <div class="transition-examples">
          <div class="transition-card">
            <h3>Hover suave</h3>
            <div class="hover-box">Pasa el mouse</div>
          </div>
          
          <div class="transition-card">
            <h3>Múltiples propiedades</h3>
            <div class="multi-transition">Transición compleja</div>
          </div>
          
          <div class="transition-card">
            <h3>Botón animado</h3>
            <button class="animated-button">Click me</button>
          </div>
        </div>
      </section>

      <section class="animations-section">
        <h2>Animaciones Keyframes</h2>
        <div class="animation-examples">
          <div class="animation-card">
            <h3>Fade In</h3>
            <div class="fade-in-box">Apareciendo...</div>
          </div>
          
          <div class="animation-card">
            <h3>Slide In</h3>
            <div class="slide-in-box">Deslizándose...</div>
          </div>
          
          <div class="animation-card">
            <h3>Bounce</h3>
            <div class="bounce-box">Rebotando...</div>
          </div>
          
          <div class="animation-card">
            <h3>Scale</h3>
            <div class="scale-box">Escalando...</div>
          </div>
        </div>
      </section>

      <section class="transforms-section">
        <h2>Transforms</h2>
        <div class="transform-examples">
          <div class="transform-card">
            <h3>Rotate</h3>
            <div class="rotate-box">Rotando</div>
          </div>
          
          <div class="transform-card">
            <h3>3D Transform</h3>
            <div class="transform-3d">3D</div>
          </div>
          
          <div class="transform-card">
            <h3>Origin</h3>
            <div class="origin-box">Origen personalizado</div>
          </div>
        </div>
      </section>

      <section class="performance-section">
        <h2>Performance</h2>
        <div class="performance-examples">
          <div class="performance-card">
            <h3>Hardware Accelerated</h3>
            <div class="hardware-box">GPU</div>
          </div>
          
          <div class="performance-card">
            <h3>Loading Spinner</h3>
            <div class="spinner"></div>
          </div>
          
          <div class="performance-card">
            <h3>Progress Bar</h3>
            <div class="progress-container">
              <div class="progress-bar">
                <div class="progress-fill"></div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section class="complex-section">
        <h2>Animaciones Complejas</h2>
        <div class="complex-examples">
          <div class="card-hover">
            <h3>Tarjeta Interactiva</h3>
            <p>Efectos hover avanzados</p>
          </div>
          
          <div class="staggered-list">
            <h3>Lista Escalada</h3>
            <ul>
              <li>Item 1</li>
              <li>Item 2</li>
              <li>Item 3</li>
              <li>Item 4</li>
            </ul>
          </div>
          
          <div class="morphing-shape">
            <h3>Forma Cambiante</h3>
            <div class="shape"></div>
          </div>
        </div>
      </section>
    </main>

    <footer class="footer">
      <p>&copy; 2024 Animations & Transitions Master</p>
    </footer>
  </div>

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
}

/* Layout */
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  text-align: center;
  margin-bottom: 40px;
}

.header h1 {
  font-size: 3rem;
  color: var(--primary);
  margin-bottom: 10px;
}

.header p {
  font-size: 1.2rem;
  opacity: 0.8;
}

/* Secciones */
section {
  margin-bottom: 60px;
}

h2 {
  font-size: 2.5rem;
  text-align: center;
  margin-bottom: 30px;
  color: var(--dark);
  position: relative;
}

h2::after {
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

/* Transiciones */
.transition-examples, .animation-examples, .transform-examples, 
.performance-examples, .complex-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
  margin-bottom: 40px;
}

.transition-card, .animation-card, .transform-card, 
.performance-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  text-align: center;
}

.transition-card h3, .animation-card h3, .transform-card h3,
.performance-card h3 {
  margin-bottom: 20px;
  color: var(--primary);
}

/* Hover box */
.hover-box {
  width: 150px;
  height: 150px;
  background: var(--primary);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  margin: 0 auto;
}

.hover-box:hover {
  background: var(--secondary);
  transform: scale(1.1);
  box-shadow: 0 10px 30px rgba(52, 152, 219, 0.3);
}

/* Multi transition */
.multi-transition {
  width: 150px;
  height: 150px;
  background: linear-gradient(45deg, var(--primary), var(--secondary));
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease, border-radius 0.5s ease;
  margin: 0 auto;
}

.multi-transition:hover {
  transform: rotate(5deg) scale(1.05);
  box-shadow: 0 15px 35px rgba(0,0,0,0.2);
  border-radius: 50%;
}

/* Animated button */
.animated-button {
  padding: 12px 24px;
  background: var(--primary);
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
}

.animated-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
  transition: left 0.5s ease;
}

.animated-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(52, 152, 219, 0.4);
}

.animated-button:hover::before {
  left: 100%;
}

.animated-button:active {
  transform: translateY(0);
}

/* Animaciones */
.fade-in-box {
  width: 150px;
  height: 150px;
  background: var(--secondary);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  animation: fadeIn 2s ease-in-out infinite;
  margin: 0 auto;
}

@keyframes fadeIn {
  0%, 100% { opacity: 0.3; }
  50% { opacity: 1; }
}

.slide-in-box {
  width: 150px;
  height: 150px;
  background: var(--accent);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  animation: slideIn 2s ease-in-out infinite;
  margin: 0 auto;
}

@keyframes slideIn {
  0%, 100% { transform: translateX(-20px); }
  50% { transform: translateX(20px); }
}

.bounce-box {
  width: 150px;
  height: 150px;
  background: #9b59b6;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  animation: bounce 2s ease-in-out infinite;
  margin: 0 auto;
}

@keyframes bounce {
  0%, 20%, 53%, 80%, 100% {
    transform: translateY(0);
  }
  40%, 43% {
    transform: translateY(-20px);
  }
  70% {
    transform: translateY(-10px);
  }
  90% {
    transform: translateY(-2px);
  }
}

.scale-box {
  width: 150px;
  height: 150px;
  background: #f39c12;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  animation: scalePulse 2s ease-in-out infinite;
  margin: 0 auto;
}

@keyframes scalePulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

/* Transforms */
.rotate-box {
  width: 150px;
  height: 150px;
  background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.3s ease;
  margin: 0 auto;
}

.rotate-box:hover {
  transform: rotate(15deg) scale(1.05);
}

.transform-3d {
  width: 150px;
  height: 150px;
  background: linear-gradient(45deg, #667eea, #764ba2);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.3s ease;
  margin: 0 auto;
}

.transform-3d:hover {
  transform: perspective(1000px) rotateY(15deg) rotateX(5deg);
}

.origin-box {
  width: 150px;
  height: 150px;
  background: linear-gradient(45deg, #f093fb, #f5576c);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.3s ease;
  transform-origin: bottom right;
  margin: 0 auto;
}

.origin-box:hover {
  transform: rotate(-15deg);
}

/* Performance */
.hardware-box {
  width: 150px;
  height: 150px;
  background: var(--primary);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.3s ease;
  transform: translateZ(0);
  margin: 0 auto;
}

.hardware-box:hover {
  transform: translateY(-10px);
}

.spinner {
  width: 60px;
  height: 60px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid var(--primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.progress-container {
  width: 200px;
  margin: 0 auto;
}

.progress-bar {
  width: 100%;
  height: 20px;
  background: #f0f0f0;
  border-radius: 10px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--primary), var(--secondary));
  border-radius: 10px;
  animation: progress 3s ease-out;
}

@keyframes progress {
  from { width: 0%; }
  to { width: 85%; }
}

/* Complejas */
.card-hover {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  cursor: pointer;
  transition: all 0.3s ease;
}

.card-hover:hover {
  transform: translateY(-10px);
  box-shadow: 0 15px 35px rgba(0,0,0,0.15);
}

.staggered-list {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.staggered-list ul {
  list-style: none;
  padding: 0;
}

.staggered-list li {
  padding: 10px 0;
  border-bottom: 1px solid #eee;
  opacity: 0;
  transform: translateX(-20px);
  animation: staggerIn 0.5s ease forwards;
}

.staggered-list li:nth-child(1) { animation-delay: 0.1s; }
.staggered-list li:nth-child(2) { animation-delay: 0.2s; }
.staggered-list li:nth-child(3) { animation-delay: 0.3s; }
.staggered-list li:nth-child(4) { animation-delay: 0.4s; }

@keyframes staggerIn {
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.morphing-shape {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  display: flex;
  justify-content: center;
}

.shape {
  width: 100px;
  height: 100px;
  background: var(--primary);
  animation: morph 4s ease-in-out infinite;
}

@keyframes morph {
  0%, 100% { border-radius: 50%; transform: rotate(0deg); }
  25% { border-radius: 0; transform: rotate(90deg); }
  50% { border-radius: 50% 0 50% 0; transform: rotate(180deg); }
  75% { border-radius: 0 50% 0 50%; transform: rotate(270deg); }
}

/* Footer */
.footer {
  text-align: center;
  margin-top: 60px;
  padding: 20px;
  background: var(--dark);
  color: white;
}

/* Responsive */
@media (max-width: 768px) {
  .transition-examples, .animation-examples, .transform-examples, 
  .performance-examples, .complex-examples {
    grid-template-columns: 1fr;
  }
  
  .header h1 {
    font-size: 2rem;
  }
  
  h2 {
    font-size: 2rem;
  }
}

/* Accesibilidad */
@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
```

```javascript
// script.js
document.addEventListener('DOMContentLoaded', function() {
  // Animación de entrada para las secciones
  const sections = document.querySelectorAll('section');
  
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.style.opacity = '1';
        entry.target.style.transform = 'translateY(0)';
      }
    });
  });
  
  sections.forEach(section => {
    section.style.opacity = '0';
    section.style.transform = 'translateY(30px)';
    section.style.transition = 'opacity 0.6s ease, transform 0.6s ease';
    observer.observe(section);
  });
  
  // Control de animaciones con botones
  const animatedButton = document.querySelector('.animated-button');
  let clickCount = 0;
  
  animatedButton.addEventListener('click', function() {
    clickCount++;
    this.textContent = `Clicked ${clickCount} times!`;
    
    // Animación extra al hacer click
    this.style.animation = 'none';
    setTimeout(() => {
      this.style.animation = 'bounce 0.5s ease';
    }, 10);
  });
});
```

## Resumen

Las animaciones y transiciones son esenciales para interfaces modernas:

- ✅ **Transiciones**: `transition` para cambios suaves entre estados
- ✅ **Animaciones**: `@keyframes` para secuencias complejas
- ✅ **Transforms**: `translate`, `rotate`, `scale` para movimiento
- ✅ **Performance**: Usar `transform` y `opacity` para animaciones fluidas
- ✅ **Accesibilidad**: Respetar `prefers-reduced-motion`

Las animaciones deben mejorar la UX, no distraer. Usa con moderación y propósito.