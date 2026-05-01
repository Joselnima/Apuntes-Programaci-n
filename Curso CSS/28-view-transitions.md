# Módulo 31 - View Transitions

En este módulo aprenderás sobre View Transitions API, una característica moderna de CSS que permite crear transiciones suaves entre diferentes estados o páginas de una aplicación web, mejorando significativamente la experiencia del usuario.

## Introducción a View Transitions

### ¿Qué son las View Transitions?

View Transitions permiten animar cambios entre diferentes vistas o estados de una aplicación, creando transiciones suaves similares a las de aplicaciones nativas.

```javascript
// Transición básica
document.startViewTransition(() => {
  // Cambiar el DOM
  updateContent();
});
```

```css
/* Animar la transición */
::view-transition-old(root) {
  animation: fade-out 0.5s ease-in-out;
}

::view-transition-new(root) {
  animation: fade-in 0.5s ease-in-out;
}
```

## Pseudo-elementos de View Transition

### Elementos principales

```css
/* Raíz de la transición */
::view-transition {
  /* Contenedor de toda la transición */
}

/* Estado anterior */
::view-transition-old(root) {
  animation: slide-out 0.6s ease-in-out;
}

/* Estado nuevo */
::view-transition-new(root) {
  animation: slide-in 0.6s ease-in-out;
}

/* Imagen de fondo durante la transición */
::view-transition-image-pair(root) {
  /* Estilos para la composición */
}

/* Capas individuales */
::view-transition-group(root) {
  animation: scale 0.6s ease-in-out;
}
```

### Transiciones por defecto

```css
/* Transición cross-fade por defecto */
::view-transition-old(root) {
  animation: 0.6s ease-in-out both cross-fade;
}

::view-transition-new(root) {
  animation: 0.6s ease-in-out both cross-fade;
}
```

## View Transition Types

### Transiciones nombradas

```javascript
// Transición con nombre específico
document.startViewTransition(() => {
  updateContent();
}, {
  types: ['slide']
});
```

```css
/* Transición nombrada */
::view-transition-old(slide) {
  animation: slide-out-left 0.5s ease-in-out;
}

::view-transition-new(slide) {
  animation: slide-in-right 0.5s ease-in-out;
}
```

### Múltiples tipos

```javascript
document.startViewTransition(() => {
  updateContent();
}, {
  types: ['slide', 'fade']
});
```

## Ejemplo Básico Interactivo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>View Transitions - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: #f5f5f7;
      color: #333;
      padding: 20px;
    }

    .container {
      max-width: 800px;
      margin: 0 auto;
    }

    .header {
      text-align: center;
      margin-bottom: 40px;
    }

    .header h1 {
      font-size: 2.5rem;
      margin-bottom: 10px;
      color: #333;
    }

    .header p {
      font-size: 1.1rem;
      color: #666;
    }

    .demo-section {
      background: white;
      border-radius: 12px;
      padding: 30px;
      margin-bottom: 30px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
      border: 1px solid #e0e0e0;
    }

    .demo-section h2 {
      margin-bottom: 20px;
      color: #333;
      font-size: 1.8rem;
      border-bottom: 2px solid #667eea;
      padding-bottom: 10px;
    }

    .content-card {
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
      padding: 40px;
      border-radius: 12px;
      text-align: center;
      margin-bottom: 30px;
      transition: all 0.3s ease;
    }

    .content-card:hover {
      transform: translateY(-5px);
      box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
    }

    .content-card h3 {
      font-size: 2rem;
      margin-bottom: 15px;
    }

    .content-card p {
      font-size: 1.1rem;
      opacity: 0.9;
      line-height: 1.6;
    }

    .controls {
      display: flex;
      gap: 15px;
      justify-content: center;
      flex-wrap: wrap;
      margin-bottom: 30px;
    }

    .control-btn {
      padding: 12px 24px;
      background: #667eea;
      color: white;
      border: none;
      border-radius: 8px;
      font-size: 16px;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    .control-btn:hover {
      background: #5a67d8;
      transform: translateY(-2px);
    }

    .control-btn:active {
      transform: translateY(0);
    }

    .transition-types {
      display: flex;
      gap: 10px;
      justify-content: center;
      flex-wrap: wrap;
      margin-bottom: 20px;
    }

    .type-btn {
      padding: 8px 16px;
      background: #f8f9fa;
      color: #333;
      border: 1px solid #dee2e6;
      border-radius: 6px;
      font-size: 14px;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    .type-btn:hover {
      background: #e9ecef;
    }

    .type-btn.active {
      background: #667eea;
      color: white;
      border-color: #667eea;
    }

    .status {
      text-align: center;
      padding: 15px;
      background: #f8f9fa;
      border-radius: 8px;
      border: 1px solid #e0e0e0;
      font-family: monospace;
      font-size: 14px;
    }

    /* View Transition Styles */
    ::view-transition-old(root) {
      animation: fade-out 0.6s ease-in-out;
    }

    ::view-transition-new(root) {
      animation: fade-in 0.6s ease-in-out;
    }

    @keyframes fade-out {
      to {
        opacity: 0;
      }
    }

    @keyframes fade-in {
      from {
        opacity: 0;
      }
      to {
        opacity: 1;
      }
    }

    /* Transición slide */
    ::view-transition-old(slide) {
      animation: slide-out-left 0.5s ease-in-out;
    }

    ::view-transition-new(slide) {
      animation: slide-in-right 0.5s ease-in-out;
    }

    @keyframes slide-out-left {
      to {
        transform: translateX(-100px);
        opacity: 0;
      }
    }

    @keyframes slide-in-right {
      from {
        transform: translateX(100px);
        opacity: 0;
      }
      to {
        transform: translateX(0);
        opacity: 1;
      }
    }

    /* Transición scale */
    ::view-transition-old(scale) {
      animation: scale-out 0.5s ease-in-out;
    }

    ::view-transition-new(scale) {
      animation: scale-in 0.5s ease-in-out;
    }

    @keyframes scale-out {
      to {
        transform: scale(0.8);
        opacity: 0;
      }
    }

    @keyframes scale-in {
      from {
        transform: scale(1.2);
        opacity: 0;
      }
      to {
        transform: scale(1);
        opacity: 1;
      }
    }

    /* Transición rotate */
    ::view-transition-old(rotate) {
      animation: rotate-out 0.6s ease-in-out;
    }

    ::view-transition-new(rotate) {
      animation: rotate-in 0.6s ease-in-out;
    }

    @keyframes rotate-out {
      to {
        transform: rotate(-10deg) scale(0.9);
        opacity: 0;
      }
    }

    @keyframes rotate-in {
      from {
        transform: rotate(10deg) scale(1.1);
        opacity: 0;
      }
      to {
        transform: rotate(0deg) scale(1);
        opacity: 1;
      }
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>View Transitions API</h1>
      <p>Transiciones suaves entre estados</p>
    </header>

    <div class="demo-section">
      <h2>Contenido Dinámico</h2>

      <div class="content-card" id="dynamicContent">
        <h3>Bienvenido</h3>
        <p>Este contenido cambiará con transiciones suaves cuando hagas clic en los botones de abajo.</p>
      </div>

      <div class="controls">
        <button class="control-btn" onclick="changeContent('welcome')">Bienvenido</button>
        <button class="control-btn" onclick="changeContent('about')">Sobre mí</button>
        <button class="control-btn" onclick="changeContent('projects')">Proyectos</button>
        <button class="control-btn" onclick="changeContent('contact')">Contacto</button>
      </div>

      <div class="transition-types">
        <button class="type-btn active" onclick="setTransitionType('')">Default</button>
        <button class="type-btn" onclick="setTransitionType('slide')">Slide</button>
        <button class="type-btn" onclick="setTransitionType('scale')">Scale</button>
        <button class="type-btn" onclick="setTransitionType('rotate')">Rotate</button>
      </div>

      <div class="status" id="status">
        Estado: Listo para transiciones
      </div>
    </div>
  </div>

  <script>
    let currentTransitionType = '';

    const contents = {
      welcome: {
        title: 'Bienvenido',
        text: 'Este contenido cambiará con transiciones suaves cuando hagas clic en los botones de abajo.'
      },
      about: {
        title: 'Sobre mí',
        text: 'Soy un desarrollador web apasionado por crear experiencias digitales excepcionales. Me especializo en tecnologías frontend modernas.'
      },
      projects: {
        title: 'Mis Proyectos',
        text: 'He trabajado en diversos proyectos incluyendo aplicaciones web, sitios e-commerce, dashboards interactivos y mucho más.'
      },
      contact: {
        title: 'Contacto',
        text: '¿Interesado en trabajar juntos? No dudes en contactarme. Estoy disponible para proyectos freelance y oportunidades laborales.'
      }
    };

    function setTransitionType(type) {
      currentTransitionType = type;

      // Update active button
      document.querySelectorAll('.type-btn').forEach(btn => {
        btn.classList.remove('active');
      });
      event.target.classList.add('active');

      updateStatus(`Tipo de transición: ${type || 'Default'}`);
    }

    function changeContent(type) {
      if (!document.startViewTransition) {
        updateStatus('❌ View Transitions no soportadas en este navegador');
        // Fallback sin transición
        updateContent(type);
        return;
      }

      updateStatus('🎬 Iniciando transición...');

      const transition = document.startViewTransition(() => {
        updateContent(type);
      }, {
        types: currentTransitionType ? [currentTransitionType] : []
      });

      transition.finished.then(() => {
        updateStatus('✅ Transición completada');
      }).catch(() => {
        updateStatus('❌ Error en la transición');
      });
    }

    function updateContent(type) {
      const content = contents[type];
      const card = document.getElementById('dynamicContent');

      card.innerHTML = `
        <h3>${content.title}</h3>
        <p>${content.text}</p>
      `;
    }

    function updateStatus(message) {
      document.getElementById('status').textContent = message;
    }

    // Inicializar
    updateStatus('Estado: Listo para transiciones');
  </script>
</body>
</html>
```

## Técnicas Avanzadas

### Transiciones con elementos específicos

```css
/* Transiciones para elementos específicos */
::view-transition-old(header) {
  animation: slide-up 0.5s ease-in-out;
}

::view-transition-new(header) {
  animation: slide-down 0.5s ease-in-out;
}

::view-transition-old(main) {
  animation: fade-out 0.4s ease-in-out;
}

::view-transition-new(main) {
  animation: fade-in 0.4s ease-in-out;
}
```

### Usando `view-transition-name`

```html
<header style="view-transition-name: header;">
  <h1>Mi Sitio Web</h1>
</header>

<main style="view-transition-name: main;">
  <p>Contenido principal</p>
</main>
```

```css
::view-transition-old(header) {
  animation: header-out 0.5s ease-in-out;
}

::view-transition-new(header) {
  animation: header-in 0.5s ease-in-out;
}
```

## SPA Navigation con View Transitions

### React Router con View Transitions

```javascript
import { useNavigate } from 'react-router-dom';

function Navigation() {
  const navigate = useNavigate();

  const handleNavigation = (path) => {
    if (!document.startViewTransition) {
      navigate(path);
      return;
    }

    document.startViewTransition(() => {
      navigate(path);
    });
  };

  return (
    <nav>
      <button onClick={() => handleNavigation('/')}>Home</button>
      <button onClick={() => handleNavigation('/about')}>About</button>
      <button onClick={() => handleNavigation('/projects')}>Projects</button>
    </nav>
  );
}
```

### Vue.js con View Transitions

```javascript
// router/index.js
import { createRouter } from 'vue-router'

const router = createRouter({
  // ... configuración del router
})

// main.js
import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

app.mixin({
  beforeRouteUpdate(to, from, next) {
    if (!document.startViewTransition) {
      next();
      return;
    }

    document.startViewTransition(() => {
      next();
    });
  }
});

app.use(router);
app.mount('#app');
```

## Ejemplo Complejo: Navegación SPA

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>SPA con View Transitions</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: #0a0a0a;
      color: #e0e0e0;
      min-height: 100vh;
    }

    .app {
      min-height: 100vh;
      display: flex;
      flex-direction: column;
    }

    .navbar {
      background: rgba(255,255,255,0.05);
      backdrop-filter: blur(10px);
      border-bottom: 1px solid rgba(255,255,255,0.1);
      padding: 1rem 2rem;
      position: sticky;
      top: 0;
      z-index: 100;
      view-transition-name: navbar;
    }

    .nav-container {
      max-width: 1200px;
      margin: 0 auto;
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .nav-brand {
      font-size: 1.5rem;
      font-weight: bold;
      background: linear-gradient(45deg, #667eea, #764ba2);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    .nav-links {
      display: flex;
      gap: 2rem;
      list-style: none;
    }

    .nav-link {
      color: #a0aec0;
      text-decoration: none;
      padding: 0.5rem 1rem;
      border-radius: 6px;
      transition: all 0.2s ease;
      cursor: pointer;
    }

    .nav-link:hover {
      color: #e0e0e0;
      background: rgba(255,255,255,0.1);
    }

    .nav-link.active {
      color: #667eea;
      background: rgba(102, 126, 234, 0.1);
    }

    .main-content {
      flex: 1;
      padding: 2rem;
      max-width: 1200px;
      margin: 0 auto;
      width: 100%;
      view-transition-name: main;
    }

    .page {
      background: rgba(255,255,255,0.02);
      border-radius: 12px;
      padding: 3rem;
      border: 1px solid rgba(255,255,255,0.05);
      min-height: 60vh;
    }

    .page h1 {
      font-size: 3rem;
      margin-bottom: 1rem;
      background: linear-gradient(45deg, #667eea, #764ba2);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    .page p {
      font-size: 1.1rem;
      line-height: 1.8;
      color: #a0aec0;
      margin-bottom: 2rem;
    }

    .page-content {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 2rem;
      margin-top: 3rem;
    }

    .content-card {
      background: rgba(255,255,255,0.05);
      border-radius: 8px;
      padding: 2rem;
      border: 1px solid rgba(255,255,255,0.1);
    }

    .content-card h3 {
      color: #e0e0e0;
      margin-bottom: 1rem;
    }

    .content-card p {
      color: #a0aec0;
      line-height: 1.6;
    }

    .loading {
      display: flex;
      align-items: center;
      justify-content: center;
      min-height: 60vh;
      font-size: 1.2rem;
      color: #667eea;
    }

    /* View Transitions */
    ::view-transition-old(root) {
      animation: fade-out 0.4s ease-in-out;
    }

    ::view-transition-new(root) {
      animation: fade-in 0.4s ease-in-out;
    }

    ::view-transition-old(navbar) {
      animation: slide-down 0.3s ease-in-out;
    }

    ::view-transition-new(navbar) {
      animation: slide-up 0.3s ease-in-out;
    }

    ::view-transition-old(main) {
      animation: scale-out 0.4s ease-in-out;
    }

    ::view-transition-new(main) {
      animation: scale-in 0.4s ease-in-out;
    }

    @keyframes fade-out {
      to { opacity: 0; }
    }

    @keyframes fade-in {
      from { opacity: 0; }
      to { opacity: 1; }
    }

    @keyframes slide-down {
      to { transform: translateY(-20px); opacity: 0; }
    }

    @keyframes slide-up {
      from { transform: translateY(-20px); opacity: 0; }
      to { transform: translateY(0); opacity: 1; }
    }

    @keyframes scale-out {
      to { transform: scale(0.95); opacity: 0; }
    }

    @keyframes scale-in {
      from { transform: scale(1.05); opacity: 0; }
      to { transform: scale(1); opacity: 1; }
    }

    /* Transiciones específicas por página */
    .page-home::view-transition-old(main) {
      animation: slide-out-left 0.5s ease-in-out;
    }

    .page-about::view-transition-new(main) {
      animation: slide-in-right 0.5s ease-in-out;
    }

    .page-projects::view-transition-old(main) {
      animation: rotate-out 0.5s ease-in-out;
    }

    .page-projects::view-transition-new(main) {
      animation: rotate-in 0.5s ease-in-out;
    }

    @keyframes slide-out-left {
      to { transform: translateX(-50px); opacity: 0; }
    }

    @keyframes slide-in-right {
      from { transform: translateX(50px); opacity: 0; }
      to { transform: translateX(0); opacity: 1; }
    }

    @keyframes rotate-out {
      to { transform: rotate(-5deg) scale(0.95); opacity: 0; }
    }

    @keyframes rotate-in {
      from { transform: rotate(5deg) scale(1.05); opacity: 0; }
      to { transform: rotate(0deg) scale(1); opacity: 1; }
    }

    @media (max-width: 768px) {
      .nav-links {
        display: none;
      }

      .main-content {
        padding: 1rem;
      }

      .page {
        padding: 2rem 1.5rem;
      }

      .page h1 {
        font-size: 2rem;
      }
    }
  </style>
</head>
<body>
  <div class="app">
    <nav class="navbar">
      <div class="nav-container">
        <div class="nav-brand">Mi Portfolio</div>
        <ul class="nav-links">
          <li><a class="nav-link" onclick="navigateTo('home')">Inicio</a></li>
          <li><a class="nav-link" onclick="navigateTo('about')">Sobre mí</a></li>
          <li><a class="nav-link" onclick="navigateTo('projects')">Proyectos</a></li>
          <li><a class="nav-link" onclick="navigateTo('contact')">Contacto</a></li>
        </ul>
      </div>
    </nav>

    <main class="main-content">
      <div id="page-content" class="page">
        <!-- Contenido dinámico -->
      </div>
    </main>
  </div>

  <script>
    const pages = {
      home: {
        title: 'Bienvenido a mi Portfolio',
        content: `
          <p>Soy un desarrollador web apasionado por crear experiencias digitales excepcionales. Especializado en tecnologías modernas de frontend y backend.</p>

          <div class="page-content">
            <div class="content-card">
              <h3>🚀 Desarrollo Web</h3>
              <p>Creo aplicaciones web modernas con React, Vue.js y Node.js, enfocándome en performance y experiencia de usuario.</p>
            </div>
            <div class="content-card">
              <h3>🎨 Diseño UI/UX</h3>
              <p>Diseño interfaces intuitivas y atractivas, combinando principios de diseño con las últimas tendencias tecnológicas.</p>
            </div>
            <div class="content-card">
              <h3>⚡ Optimización</h3>
              <p>Me aseguro de que cada proyecto sea rápido, accesible y funcione perfectamente en todos los dispositivos.</p>
            </div>
          </div>
        `
      },
      about: {
        title: 'Sobre mí',
        content: `
          <p>Con más de 5 años de experiencia en desarrollo web, he trabajado en proyectos de diversos tamaños y complejidades.</p>

          <div class="page-content">
            <div class="content-card">
              <h3>💼 Experiencia</h3>
              <p>Desarrollo full-stack con experiencia en startups y empresas Fortune 500. He liderado equipos y mentoreado desarrolladores junior.</p>
            </div>
            <div class="content-card">
              <h3>🎯 Enfoque</h3>
              <p>Creo en el desarrollo sostenible, escribiendo código mantenible y escalable. Siempre busco las mejores prácticas y tecnologías emergentes.</p>
            </div>
            <div class="content-card">
              <h3>🌱 Aprendizaje Continuo</h3>
              <p>La tecnología evoluciona rápidamente. Me mantengo actualizado con las últimas tendencias y mejores prácticas de la industria.</p>
            </div>
          </div>
        `
      },
      projects: {
        title: 'Mis Proyectos',
        content: `
          <p>Aquí algunos de los proyectos en los que he trabajado recientemente. Cada uno representa un desafío único y una oportunidad de aprendizaje.</p>

          <div class="page-content">
            <div class="content-card">
              <h3>🛒 E-commerce Platform</h3>
              <p>Plataforma completa de comercio electrónico con React, Node.js y Stripe. Incluye panel de administración y sistema de inventario.</p>
            </div>
            <div class="content-card">
              <h3>📊 Dashboard Analytics</h3>
              <p>Dashboard interactivo con D3.js para visualización de datos en tiempo real. Procesamiento de big data y exportación de reportes.</p>
            </div>
            <div class="content-card">
              <h3>🎮 Web Game</h3>
              <p>Juego web multijugador desarrollado con WebSockets y Canvas API. Compatible con móviles y con sistema de rankings.</p>
            </div>
          </div>
        `
      },
      contact: {
        title: 'Hablemos',
        content: `
          <p>¿Tienes un proyecto en mente? ¿Quieres colaborar? Estoy disponible para discutir oportunidades y cómo podemos trabajar juntos.</p>

          <div class="page-content">
            <div class="content-card">
              <h3>📧 Email</h3>
              <p>hola@miportfolio.com</p>
            </div>
            <div class="content-card">
              <h3>💼 LinkedIn</h3>
              <p>linkedin.com/in/miportfolio</p>
            </div>
            <div class="content-card">
              <h3>🐙 GitHub</h3>
              <p>github.com/miportfolio</p>
            </div>
          </div>
        `
      }
    };

    let currentPage = 'home';

    function navigateTo(pageId) {
      if (pageId === currentPage) return;

      if (!document.startViewTransition) {
        // Fallback sin transiciones
        updatePage(pageId);
        return;
      }

      // Actualizar navegación
      document.querySelectorAll('.nav-link').forEach(link => {
        link.classList.remove('active');
      });
      event.target.classList.add('active');

      // Iniciar transición
      const transition = document.startViewTransition(() => {
        updatePage(pageId);
      });

      transition.finished.then(() => {
        currentPage = pageId;
      });
    }

    function updatePage(pageId) {
      const pageData = pages[pageId];
      const contentDiv = document.getElementById('page-content');

      contentDiv.className = `page page-${pageId}`;
      contentDiv.innerHTML = `
        <h1>${pageData.title}</h1>
        ${pageData.content}
      `;
    }

    // Inicializar
    updatePage('home');
    document.querySelector('.nav-link').classList.add('active');
  </script>
</body>
</html>
```

## Técnicas Avanzadas

### Transiciones condicionales

```javascript
function navigateWithTransition(to, condition) {
  const transitionType = condition ? 'slide' : 'fade';

  document.startViewTransition(() => {
    navigate(to);
  }, {
    types: [transitionType]
  });
}
```

### Transiciones con datos

```javascript
// Transición basada en el tipo de cambio
function transitionBasedOnChange(oldData, newData) {
  let transitionType = 'fade';

  if (oldData.category !== newData.category) {
    transitionType = 'slide';
  } else if (oldData.priority !== newData.priority) {
    transitionType = 'scale';
  }

  document.startViewTransition(() => {
    updateUI(newData);
  }, {
    types: [transitionType]
  });
}
```

### Performance y optimización

```css
/* Evitar repaints durante transiciones */
::view-transition-old(root),
::view-transition-new(root) {
  /* Usar solo transform y opacity */
  transform: translateZ(0); /* Force hardware acceleration */
}

/* Optimizar para móviles */
@media (max-width: 768px) {
  ::view-transition-old(root) {
    animation-duration: 0.3s; /* Más rápido en móviles */
  }

  ::view-transition-new(root) {
    animation-duration: 0.3s;
  }
}
```

## Compatibilidad y Fallbacks

### Detectar soporte

```javascript
function supportsViewTransitions() {
  return typeof document.startViewTransition === 'function';
}

// Fallback con animaciones CSS
function transitionWithFallback(updateFunction) {
  if (supportsViewTransitions()) {
    document.startViewTransition(updateFunction);
  } else {
    // Fallback con clases CSS
    document.body.classList.add('transitioning');
    updateFunction();

    // Remover clase después de la animación
    setTimeout(() => {
      document.body.classList.remove('transitioning');
    }, 300);
  }
}
```

### Polyfill para navegadores antiguos

```html
<!-- Polyfill para View Transitions -->
<script src="https://unpkg.com/@oddbird/css-anchor-positioning@latest/dist/css-anchor-positioning-polyfill.js"></script>
<script>
  // Inicializar polyfill si es necesario
  if (!supportsViewTransitions()) {
    console.log('Usando polyfill para View Transitions');
  }
</script>
```

## Casos de Uso Prácticos

### 1. Cambio de tema

```javascript
function toggleTheme() {
  document.startViewTransition(() => {
    document.body.classList.toggle('dark-theme');
  });
}
```

### 2. Filtros y búsqueda

```javascript
function applyFilter(filterType) {
  document.startViewTransition(() => {
    updateListItems(filterType);
  }, {
    types: ['filter']
  });
}
```

### 3. Modal y overlays

```javascript
function showModal() {
  document.startViewTransition(() => {
    modal.classList.add('visible');
  }, {
    types: ['modal']
  });
}
```

## Resumen

View Transitions API revoluciona las transiciones web:

- ✅ **Transiciones nativas**: Sin JavaScript complejo para animaciones
- ✅ **Performance óptima**: Animaciones hardware-accelerated
- ✅ **Experiencia nativa**: Similar a aplicaciones móviles
- ✅ **Flexibilidad total**: Control completo sobre las animaciones
- ✅ **Progressive enhancement**: Funciona sin soporte nativo

View Transitions permite crear aplicaciones web con transiciones suaves y profesionales, elevando significativamente la experiencia del usuario.