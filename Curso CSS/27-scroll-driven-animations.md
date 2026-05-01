# Módulo 30 - Scroll-driven Animations

En este módulo aprenderás sobre Scroll-driven Animations, una poderosa API de CSS que permite crear animaciones basadas en el progreso del scroll, creando experiencias interactivas y dinámicas.

## Introducción a Scroll-driven Animations

### ¿Qué son las Scroll-driven Animations?

Las Scroll-driven Animations permiten animar elementos basándose en la posición del scroll, creando efectos parallax, reveals progresivos, y animaciones que responden al movimiento del usuario.

```css
@keyframes slide-in {
  from {
    opacity: 0;
    transform: translateX(-100px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.scroll-element {
  animation: slide-in linear;
  animation-timeline: scroll();
}
```

## Scroll Timelines

### `scroll()` timeline

```css
.element {
  animation: slide-in linear;
  animation-timeline: scroll();
}

/* Especificar contenedor de scroll */
.element {
  animation: slide-in linear;
  animation-timeline: scroll(nearest);
}

/* Scroll horizontal */
.element {
  animation: slide-in linear;
  animation-timeline: scroll(x);
}
```

### Parámetros de `scroll()`

```css
/* Contenedor específico */
.scroll-container {
  overflow: auto;
  height: 400px;
}

.animated-element {
  animation: slide-in linear;
  animation-timeline: scroll(.scroll-container);
}

/* Dirección específica */
.horizontal-scroll {
  animation-timeline: scroll(x);
}

.vertical-scroll {
  animation-timeline: scroll(y);
}

/* Ámbito */
.block-scroll {
  animation-timeline: scroll(nearest block);
}

.inline-scroll {
  animation-timeline: scroll(nearest inline);
}
```

## View Timelines

### `view()` timeline

```css
.element {
  animation: fade-in linear;
  animation-timeline: view();
}

/* Especificar umbrales */
.element {
  animation: fade-in linear;
  animation-timeline: view(50% 75%);
}

/* Contenedor específico */
.element {
  animation: fade-in linear;
  animation-timeline: view(.container);
}
```

### Parámetros de `view()`

```css
/* Umbrales de entrada/salida */
.element {
  animation-timeline: view(0% 100%);  /* Del 0% al 100% */
  animation-timeline: view(25% 75%);  /* Del 25% al 75% */
}

/* Ejes específicos */
.element {
  animation-timeline: view(block 50%);   /* Eje vertical */
  animation-timeline: view(inline 50%);  /* Eje horizontal */
}

/* Combinaciones */
.element {
  animation-timeline: view(block 0% 100%);
}
```

## Ejemplos Prácticos

### Reveal Animation Básica

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Scroll-driven Animations - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      line-height: 1.6;
      color: #333;
    }

    .hero {
      height: 100vh;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      text-align: center;
    }

    .hero h1 {
      font-size: 4rem;
      margin-bottom: 1rem;
      animation: fade-in-up linear;
      animation-timeline: view();
    }

    .hero p {
      font-size: 1.5rem;
      opacity: 0.9;
      animation: fade-in linear;
      animation-timeline: view();
      animation-delay: 0.2s;
    }

    @keyframes fade-in-up {
      from {
        opacity: 0;
        transform: translateY(50px);
      }
      to {
        opacity: 1;
        transform: translateY(0);
      }
    }

    @keyframes fade-in {
      from { opacity: 0; }
      to { opacity: 1; }
    }

    .content {
      padding: 5rem 2rem;
      max-width: 1200px;
      margin: 0 auto;
    }

    .section {
      margin-bottom: 5rem;
      opacity: 0;
      animation: slide-in linear;
      animation-timeline: view();
    }

    .section:nth-child(odd) {
      animation-direction: reverse;
    }

    @keyframes slide-in {
      from {
        opacity: 0;
        transform: translateX(-100px);
      }
      to {
        opacity: 1;
        transform: translateX(0);
      }
    }

    .card {
      background: white;
      border-radius: 12px;
      padding: 2rem;
      margin-bottom: 2rem;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
      animation: scale-in linear;
      animation-timeline: view();
    }

    @keyframes scale-in {
      from {
        opacity: 0;
        transform: scale(0.8);
      }
      to {
        opacity: 1;
        transform: scale(1);
      }
    }

    .progress-bar {
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      height: 4px;
      background: rgba(255,255,255,0.2);
      z-index: 1000;
    }

    .progress-fill {
      height: 100%;
      background: linear-gradient(90deg, #667eea, #764ba2);
      width: 0%;
      animation: progress linear;
      animation-timeline: scroll();
    }

    @keyframes progress {
      from { width: 0%; }
      to { width: 100%; }
    }

    .spacer {
      height: 200vh;
      background: linear-gradient(to bottom,
        #667eea 0%,
        #764ba2 25%,
        #f093fb 50%,
        #f5576c 75%,
        #4ecdc4 100%);
    }
  </style>
</head>
<body>
  <div class="progress-bar">
    <div class="progress-fill"></div>
  </div>

  <section class="hero">
    <div>
      <h1>Scroll-driven Animations</h1>
      <p>Animaciones que responden al scroll</p>
    </div>
  </section>

  <div class="content">
    <section class="section">
      <h2>Sección 1</h2>
      <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
    </section>

    <section class="section">
      <h2>Sección 2</h2>
      <div class="card">
        <h3>Card 1</h3>
        <p>Contenido de la primera card con animación de escala.</p>
      </div>
      <div class="card">
        <h3>Card 2</h3>
        <p>Contenido de la segunda card con animación de escala.</p>
      </div>
    </section>

    <section class="section">
      <h2>Sección 3</h2>
      <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
    </section>
  </div>

  <div class="spacer"></div>
</body>
</html>
```

## Técnicas Avanzadas

### Parallax Scrolling

```css
.parallax-container {
  height: 200vh;
  overflow: hidden;
  position: relative;
}

.parallax-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 300%;
  background: url('mountain.jpg') center/cover;
  animation: parallax linear;
  animation-timeline: scroll(nearest);
}

.parallax-content {
  position: relative;
  z-index: 2;
  padding: 5rem 2rem;
  color: white;
}

@keyframes parallax {
  from {
    transform: translateY(0);
  }
  to {
    transform: translateY(-100%);
  }
}
```

### Sticky Animations

```css
.sticky-section {
  height: 200vh;
  position: relative;
}

.sticky-element {
  position: sticky;
  top: 50%;
  transform: translateY(-50%);
  animation: rotate-on-scroll linear;
  animation-timeline: scroll(nearest);
}

@keyframes rotate-on-scroll {
  from {
    transform: translateY(-50%) rotate(0deg);
  }
  to {
    transform: translateY(-50%) rotate(360deg);
  }
}
```

### Multiple Scroll Ranges

```css
.multi-range-element {
  animation:
    phase-1 linear 0s 0s,
    phase-2 linear 0s 25%,
    phase-3 linear 0s 50%,
    phase-4 linear 0s 75%;
  animation-timeline: scroll();
}

@keyframes phase-1 {
  from { background: red; }
  to { background: orange; }
}

@keyframes phase-2 {
  from { background: orange; }
  to { background: yellow; }
}

@keyframes phase-3 {
  from { background: yellow; }
  to { background: green; }
}

@keyframes phase-4 {
  from { background: green; }
  to { background: blue; }
}
```

## Scroll Timeline Ranges

### Control de rangos de animación

```css
/* Animar solo en la primera mitad del scroll */
.element {
  animation: slide-in linear;
  animation-timeline: scroll();
  animation-range: 0% 50%;
}

/* Animar en rangos específicos */
.element {
  animation: grow linear;
  animation-timeline: scroll();
  animation-range: 25% 75%;
}

/* Múltiples rangos */
.element {
  animation:
    fade-in linear 0s 0s,
    grow linear 0s 50%;
  animation-timeline: scroll();
  animation-range: 0% 50%, 50% 100%;
}
```

### Rangos con view timelines

```css
/* Animar cuando el elemento está parcialmente visible */
.element {
  animation: reveal linear;
  animation-timeline: view();
  animation-range: cover 0% cover 100%;
}

/* Animar desde que entra hasta que sale completamente */
.element {
  animation: parallax linear;
  animation-timeline: view();
  animation-range: contain 0% contain 100%;
}
```

## Ejemplo Complejo: Portfolio Interactivo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Portfolio con Scroll Animations</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      line-height: 1.6;
      color: #333;
      background: #0a0a0a;
      color: #e0e0e0;
    }

    .hero {
      height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      position: relative;
      overflow: hidden;
    }

    .hero-bg {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background: url('hero-pattern.svg') repeat;
      opacity: 0.1;
      animation: bg-move linear;
      animation-timeline: scroll();
    }

    @keyframes bg-move {
      from { transform: translateY(0); }
      to { transform: translateY(-50px); }
    }

    .hero-content {
      text-align: center;
      color: white;
      z-index: 2;
      position: relative;
    }

    .hero h1 {
      font-size: clamp(2rem, 8vw, 6rem);
      margin-bottom: 1rem;
      animation: title-reveal linear;
      animation-timeline: view();
    }

    .hero p {
      font-size: clamp(1rem, 4vw, 1.5rem);
      opacity: 0.9;
      animation: subtitle-reveal linear;
      animation-timeline: view();
      animation-delay: 0.2s;
    }

    @keyframes title-reveal {
      from {
        opacity: 0;
        transform: translateY(50px) scale(0.9);
      }
      to {
        opacity: 1;
        transform: translateY(0) scale(1);
      }
    }

    @keyframes subtitle-reveal {
      from {
        opacity: 0;
        transform: translateY(30px);
      }
      to {
        opacity: 1;
        transform: translateY(0);
      }
    }

    .projects {
      padding: 5rem 2rem;
      max-width: 1200px;
      margin: 0 auto;
    }

    .project {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 4rem;
      margin-bottom: 8rem;
      align-items: center;
      opacity: 0;
      animation: project-reveal linear;
      animation-timeline: view();
    }

    .project:nth-child(even) {
      direction: rtl;
    }

    .project:nth-child(even) .project-content {
      text-align: right;
    }

    @keyframes project-reveal {
      from {
        opacity: 0;
        transform: translateY(100px);
      }
      to {
        opacity: 1;
        transform: translateY(0);
      }
    }

    .project-image {
      width: 100%;
      height: 400px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      border-radius: 12px;
      animation: image-float linear;
      animation-timeline: view();
    }

    @keyframes image-float {
      from {
        transform: translateY(50px) scale(0.9);
      }
      to {
        transform: translateY(0) scale(1);
      }
    }

    .project-content h3 {
      font-size: 2.5rem;
      margin-bottom: 1rem;
      background: linear-gradient(45deg, #667eea, #764ba2);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    .project-content p {
      font-size: 1.1rem;
      line-height: 1.8;
      margin-bottom: 2rem;
      color: #a0aec0;
    }

    .project-tags {
      display: flex;
      gap: 1rem;
      flex-wrap: wrap;
      margin-bottom: 2rem;
    }

    .tag {
      background: rgba(102, 126, 234, 0.1);
      color: #667eea;
      padding: 0.5rem 1rem;
      border-radius: 20px;
      font-size: 0.875rem;
      border: 1px solid rgba(102, 126, 234, 0.2);
    }

    .project-link {
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      color: #667eea;
      text-decoration: none;
      font-weight: 500;
      transition: color 0.2s;
    }

    .project-link:hover {
      color: #764ba2;
    }

    .skills {
      padding: 5rem 2rem;
      background: rgba(255,255,255,0.02);
    }

    .skills-container {
      max-width: 1200px;
      margin: 0 auto;
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
      gap: 2rem;
    }

    .skill-card {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 2rem;
      border: 1px solid rgba(255,255,255,0.1);
      animation: skill-reveal linear;
      animation-timeline: view();
    }

    @keyframes skill-reveal {
      from {
        opacity: 0;
        transform: translateY(50px) rotateX(10deg);
      }
      to {
        opacity: 1;
        transform: translateY(0) rotateX(0deg);
      }
    }

    .skill-icon {
      width: 60px;
      height: 60px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      border-radius: 12px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 2rem;
      margin-bottom: 1rem;
    }

    .skill-name {
      font-size: 1.25rem;
      font-weight: 600;
      margin-bottom: 0.5rem;
      color: #e0e0e0;
    }

    .skill-description {
      color: #a0aec0;
      line-height: 1.6;
    }

    .progress-ring {
      width: 100px;
      height: 100px;
      margin: 1rem auto;
      animation: progress-fill linear;
      animation-timeline: view();
    }

    @keyframes progress-fill {
      from {
        --progress: 0;
      }
      to {
        --progress: 85;
      }
    }

    .contact {
      padding: 5rem 2rem;
      text-align: center;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
    }

    .contact h2 {
      font-size: 3rem;
      margin-bottom: 1rem;
    }

    .contact p {
      font-size: 1.2rem;
      opacity: 0.9;
      margin-bottom: 3rem;
    }

    .contact-form {
      max-width: 600px;
      margin: 0 auto;
      display: grid;
      gap: 1.5rem;
    }

    .form-group {
      text-align: left;
    }

    .form-group label {
      display: block;
      margin-bottom: 0.5rem;
      font-weight: 500;
    }

    .form-group input,
    .form-group textarea {
      width: 100%;
      padding: 1rem;
      border: none;
      border-radius: 8px;
      background: rgba(255,255,255,0.1);
      color: white;
      font-size: 1rem;
      backdrop-filter: blur(10px);
    }

    .form-group textarea {
      resize: vertical;
      min-height: 150px;
    }

    .submit-btn {
      padding: 1rem 2rem;
      background: white;
      color: #667eea;
      border: none;
      border-radius: 8px;
      font-size: 1.1rem;
      font-weight: 600;
      cursor: pointer;
      transition: transform 0.2s;
    }

    .submit-btn:hover {
      transform: translateY(-2px);
    }

    .scroll-indicator {
      position: fixed;
      bottom: 2rem;
      right: 2rem;
      width: 50px;
      height: 50px;
      background: rgba(102, 126, 234, 0.9);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      cursor: pointer;
      animation: bounce linear;
      animation-timeline: scroll();
      animation-range: 0% 10%;
    }

    @keyframes bounce {
      from {
        transform: translateY(0);
      }
      to {
        transform: translateY(-10px);
      }
    }

    @media (max-width: 768px) {
      .project {
        grid-template-columns: 1fr;
        gap: 2rem;
      }

      .project:nth-child(even) {
        direction: ltr;
      }

      .project:nth-child(even) .project-content {
        text-align: left;
      }
    }
  </style>
</head>
<body>
  <section class="hero">
    <div class="hero-bg"></div>
    <div class="hero-content">
      <h1>Portfolio Interactivo</h1>
      <p>Desarrollador Web Full Stack</p>
    </div>
  </section>

  <section class="projects">
    <div class="project">
      <div class="project-image"></div>
      <div class="project-content">
        <h3>E-commerce Platform</h3>
        <p>Plataforma de comercio electrónico completa con React, Node.js y MongoDB. Incluye autenticación, pagos integrados y panel de administración.</p>
        <div class="project-tags">
          <span class="tag">React</span>
          <span class="tag">Node.js</span>
          <span class="tag">MongoDB</span>
          <span class="tag">Stripe</span>
        </div>
        <a href="#" class="project-link">Ver proyecto →</a>
      </div>
    </div>

    <div class="project">
      <div class="project-content">
        <h3>Task Management App</h3>
        <p>Aplicación de gestión de tareas con Vue.js y Firebase. Características incluyen colaboración en tiempo real, drag & drop, y notificaciones push.</p>
        <div class="project-tags">
          <span class="tag">Vue.js</span>
          <span class="tag">Firebase</span>
          <span class="tag">PWA</span>
          <span class="tag">WebSockets</span>
        </div>
        <a href="#" class="project-link">Ver proyecto →</a>
      </div>
      <div class="project-image"></div>
    </div>

    <div class="project">
      <div class="project-image"></div>
      <div class="project-content">
        <h3>Data Visualization Dashboard</h3>
        <p>Dashboard interactivo para visualización de datos con D3.js, React y Express. Incluye gráficos en tiempo real y exportación de reportes.</p>
        <div class="project-tags">
          <span class="tag">D3.js</span>
          <span class="tag">React</span>
          <span class="tag">Express</span>
          <span class="tag">PostgreSQL</span>
        </div>
        <a href="#" class="project-link">Ver proyecto →</a>
      </div>
    </div>
  </section>

  <section class="skills">
    <div class="skills-container">
      <div class="skill-card">
        <div class="skill-icon">⚛️</div>
        <h3 class="skill-name">React</h3>
        <p class="skill-description">Framework JavaScript para interfaces de usuario. Desarrollo de aplicaciones web modernas y escalables.</p>
      </div>

      <div class="skill-card">
        <div class="skill-icon">🟢</div>
        <h3 class="skill-name">Node.js</h3>
        <p class="skill-description">Runtime JavaScript del lado del servidor. APIs REST, GraphQL y aplicaciones en tiempo real.</p>
      </div>

      <div class="skill-card">
        <div class="skill-icon">🎨</div>
        <h3 class="skill-name">CSS</h3>
        <p class="skill-description">Diseño web moderno con CSS Grid, Flexbox, animaciones y responsive design.</p>
      </div>

      <div class="skill-card">
        <div class="skill-icon">🗄️</div>
        <h3 class="skill-name">Databases</h3>
        <p class="skill-description">MongoDB, PostgreSQL y Redis. Diseño de esquemas, consultas optimizadas y caching.</p>
      </div>
    </div>
  </section>

  <section class="contact">
    <h2>¿Hablamos?</h2>
    <p>Estoy disponible para proyectos freelance y oportunidades laborales</p>

    <form class="contact-form">
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
        <textarea id="message" name="message" required></textarea>
      </div>

      <button type="submit" class="submit-btn">Enviar mensaje</button>
    </form>
  </section>

  <div class="scroll-indicator" onclick="window.scrollTo({top: 0, behavior: 'smooth'})">
    ↑
  </div>
</body>
</html>
```

## Técnicas Avanzadas

### Scroll Velocity Detection

```css
/* Animación basada en velocidad de scroll */
.fast-scroll-element {
  animation: fast-effect linear;
  animation-timeline: scroll();
  animation-duration: 0.1s; /* Duración corta = alta sensibilidad */
}

.slow-scroll-element {
  animation: slow-effect linear;
  animation-timeline: scroll();
  animation-duration: 2s; /* Duración larga = baja sensibilidad */
}
```

### Intersection con Scroll Animations

```css
/* Combinar con Intersection Observer */
.scroll-triggered {
  animation: reveal linear;
  animation-timeline: view();
  animation-fill-mode: both;
}

/* Animación solo cuando es visible */
@keyframes reveal {
  from {
    opacity: 0;
    transform: translateY(50px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
```

### Performance Optimization

```css
/* Usar transform y opacity para mejores performance */
.animated-element {
  animation: performant-animation linear;
  animation-timeline: scroll();
  will-change: transform, opacity;
}

@keyframes performant-animation {
  from {
    transform: translateY(100px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* Evitar animaciones costosas */
.avoid-this {
  /* ❌ */
  animation: width-change linear;
  animation-timeline: scroll();
}

@keyframes width-change {
  from { width: 100px; }
  to { width: 200px; } /* Layout thrashing */
}
```

## Compatibilidad y Fallbacks

### Detectar soporte

```javascript
// Detectar soporte de Scroll-driven Animations
function supportsScrollDrivenAnimations() {
  return CSS.supports('animation-timeline: scroll()');
}

// Fallback con Intersection Observer
if (!supportsScrollDrivenAnimations()) {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('animate');
      }
    });
  });

  document.querySelectorAll('.scroll-element').forEach(el => {
    observer.observe(el);
  });
}
```

### Progressive Enhancement

```css
/* Fallback básico */
.scroll-element {
  opacity: 0;
  transform: translateY(50px);
  transition: opacity 0.6s ease, transform 0.6s ease;
}

/* Enhancement con scroll animations */
@supports (animation-timeline: scroll()) {
  .scroll-element {
    animation: slide-in linear;
    animation-timeline: scroll();
    transition: none; /* Remover transition */
  }
}
```

## Casos de Uso Avanzados

### Storytelling Scroll

```css
.story-section {
  height: 100vh;
  position: relative;
  overflow: hidden;
}

.story-element {
  position: absolute;
  animation: story-progression linear;
  animation-timeline: scroll(nearest);
}

@keyframes story-progression {
  0% {
    opacity: 0;
    transform: scale(0.5) rotate(-10deg);
  }
  25% {
    opacity: 1;
    transform: scale(1) rotate(0deg);
  }
  75% {
    opacity: 1;
    transform: scale(1.1) rotate(5deg);
  }
  100% {
    opacity: 0;
    transform: scale(1.2) rotate(10deg);
  }
}
```

### Infinite Scroll Animation

```css
.infinite-scroll {
  animation: infinite-move linear infinite;
  animation-timeline: scroll();
}

@keyframes infinite-move {
  from { transform: translateX(0); }
  to { transform: translateX(-100px); }
}
```

## Resumen

Scroll-driven Animations transforman la experiencia web:

- ✅ **Animaciones naturales**: Respondiendo al comportamiento del usuario
- ✅ **Performance óptima**: Animaciones hardware-accelerated
- ✅ **Experiencias inmersivas**: Storytelling y navegación fluida
- ✅ **Progressive enhancement**: Funciona sin JavaScript
- ✅ **Compatibilidad futura**: API moderna y robusta

Las Scroll-driven Animations abren nuevas posibilidades para crear interfaces web más dinámicas y atractivas, mejorando significativamente la experiencia del usuario.