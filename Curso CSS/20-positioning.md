# Módulo 20 - Positioning

En este módulo aprenderás sobre CSS Positioning, una de las propiedades más importantes y poderosas de CSS que controla cómo se posicionan los elementos en el layout, incluyendo static, relative, absolute, fixed, sticky y técnicas avanzadas.

## Introducción a CSS Positioning

### ¿Qué es CSS Positioning?

CSS Positioning determina cómo se posiciona un elemento en el documento. Controla si un elemento se posiciona de manera normal en el flujo del documento o se saca del flujo normal para posicionarse de manera especial.

```css
/* Posicionamiento estático (por defecto) */
.element {
  position: static;
}

/* Posicionamiento relativo */
.element {
  position: relative;
  top: 10px;
  left: 20px;
}

/* Posicionamiento absoluto */
.element {
  position: absolute;
  top: 50px;
  left: 100px;
}

/* Posicionamiento fijo */
.element {
  position: fixed;
  top: 0;
  right: 0;
}

/* Posicionamiento sticky */
.element {
  position: sticky;
  top: 20px;
}
```

## Tipos de Positioning

### Static (Estático)

```css
.static-element {
  position: static; /* Valor por defecto */
  /* No responde a top, right, bottom, left */
}
```

El posicionamiento estático es el comportamiento normal de los elementos en el flujo del documento.

### Relative (Relativo)

```css
.relative-element {
  position: relative;
  top: 20px;
  left: 30px;
}
```

Los elementos relativamente posicionados se desplazan desde su posición normal, pero mantienen su espacio original en el flujo del documento.

### Absolute (Absoluto)

```css
.absolute-element {
  position: absolute;
  top: 50px;
  right: 20px;
}
```

Los elementos absolutamente posicionados se sacan completamente del flujo normal y se posicionan respecto a su ancestro posicionado más cercano.

### Fixed (Fijo)

```css
.fixed-element {
  position: fixed;
  top: 10px;
  left: 10px;
}
```

Los elementos fijos se posicionan respecto a la ventana del navegador y permanecen en la misma posición incluso al hacer scroll.

### Sticky (Pegajoso)

```css
.sticky-element {
  position: sticky;
  top: 0;
}
```

Los elementos sticky combinan relative y fixed: se comportan como relative hasta que alcanzan una posición específica, entonces se "pegan" como fixed.

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Positioning - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: 'Arial', sans-serif;
      background: #f5f5f7;
      color: #333;
      padding: 20px;
      line-height: 1.6;
    }

    .container {
      max-width: 1200px;
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

    .positioning-container {
      position: relative;
      height: 300px;
      background: linear-gradient(45deg, #f8f9fa, #e9ecef);
      border: 2px dashed #dee2e6;
      border-radius: 8px;
      margin-bottom: 20px;
      overflow: hidden;
    }

    .box {
      width: 80px;
      height: 80px;
      border-radius: 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-weight: bold;
      color: white;
      box-shadow: 0 4px 8px rgba(0,0,0,0.2);
      position: relative;
    }

    /* Static positioning */
    .static-box {
      position: static;
      background: linear-gradient(45deg, #6c757d, #495057);
    }

    /* Relative positioning */
    .relative-box {
      position: relative;
      top: 20px;
      left: 30px;
      background: linear-gradient(45deg, #28a745, #20c997);
    }

    /* Absolute positioning */
    .absolute-box {
      position: absolute;
      top: 50px;
      right: 50px;
      background: linear-gradient(45deg, #dc3545, #fd7e14);
    }

    /* Fixed positioning */
    .fixed-box {
      position: fixed;
      top: 20px;
      right: 20px;
      background: linear-gradient(45deg, #007bff, #6610f2);
      z-index: 1000;
    }

    /* Sticky positioning */
    .sticky-box {
      position: sticky;
      top: 20px;
      background: linear-gradient(45deg, #e83e8c, #6f42c1);
    }

    .code-example {
      background: #2d3748;
      color: #e2e8f0;
      padding: 15px;
      border-radius: 6px;
      font-family: 'Monaco', 'Menlo', monospace;
      font-size: 14px;
      margin-top: 15px;
    }

    .explanation {
      margin-top: 15px;
      padding: 15px;
      background: #f8f9fa;
      border-radius: 6px;
      border-left: 4px solid #667eea;
    }

    .comparison-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 20px;
      margin-top: 30px;
    }

    .comparison-item {
      background: #f8f9fa;
      border-radius: 8px;
      padding: 20px;
      border: 1px solid #e0e0e0;
    }

    .comparison-item h3 {
      margin-bottom: 15px;
      color: #333;
      font-size: 1.2rem;
    }

    .comparison-item .type {
      display: inline-block;
      padding: 4px 8px;
      background: #667eea;
      color: white;
      border-radius: 4px;
      font-size: 12px;
      font-weight: bold;
      margin-bottom: 10px;
    }

    .comparison-item p {
      color: #666;
      font-size: 14px;
      line-height: 1.5;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>CSS Positioning</h1>
      <p>Controlando la posición de los elementos en el layout</p>
    </header>

    <div class="demo-section">
      <h2>Static Positioning</h2>
      <div class="positioning-container">
        <div class="box static-box">Static</div>
        <p>Los elementos static siguen el flujo normal del documento y no responden a propiedades de posicionamiento como top, left, etc.</p>
      </div>
      <div class="code-example">
position: static; /* Valor por defecto */
      </div>
      <div class="explanation">
        <strong>Características:</strong> Comportamiento normal, no se ve afectado por top/right/bottom/left, mantiene su posición en el flujo del documento.
      </div>
    </div>

    <div class="demo-section">
      <h2>Relative Positioning</h2>
      <div class="positioning-container">
        <div class="box relative-box">Relative</div>
        <p>Los elementos relative se desplazan desde su posición normal pero mantienen su espacio original.</p>
      </div>
      <div class="code-example">
position: relative;<br>
top: 20px;<br>
left: 30px;
      </div>
      <div class="explanation">
        <strong>Características:</strong> Se desplaza desde su posición original, mantiene el espacio ocupado, afecta el posicionamiento de elementos hermanos.
      </div>
    </div>

    <div class="demo-section">
      <h2>Absolute Positioning</h2>
      <div class="positioning-container">
        <div class="box absolute-box">Absolute</div>
        <p>Los elementos absolute se sacan del flujo normal y se posicionan respecto a su ancestro posicionado más cercano.</p>
      </div>
      <div class="code-example">
position: absolute;<br>
top: 50px;<br>
right: 50px;
      </div>
      <div class="explanation">
        <strong>Características:</strong> Se saca del flujo normal, no ocupa espacio, se posiciona respecto al primer ancestro posicionado (no static).
      </div>
    </div>

    <div class="demo-section">
      <h2>Fixed Positioning</h2>
      <p>Los elementos fixed se posicionan respecto a la ventana del navegador.</p>
      <div class="code-example">
position: fixed;<br>
top: 20px;<br>
right: 20px;
      </div>
      <div class="explanation">
        <strong>Características:</strong> Se posiciona respecto a la ventana del navegador, permanece fijo durante el scroll, se saca del flujo normal.
      </div>
    </div>

    <div class="demo-section">
      <h2>Sticky Positioning</h2>
      <div style="height: 200px; overflow-y: auto; background: #f8f9fa; border-radius: 8px; padding: 20px; margin-bottom: 20px;">
        <div class="box sticky-box">Sticky</div>
        <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
        <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      </div>
      <div class="code-example">
position: sticky;<br>
top: 20px;
      </div>
      <div class="explanation">
        <strong>Características:</strong> Comportamiento híbrido entre relative y fixed. Se comporta como relative hasta alcanzar la posición especificada, entonces "se pega" como fixed.
      </div>
    </div>

    <div class="comparison-grid">
      <div class="comparison-item">
        <span class="type">STATIC</span>
        <h3>Flujo Normal</h3>
        <p>Los elementos siguen el flujo normal del documento. Es el comportamiento por defecto de todos los elementos.</p>
      </div>

      <div class="comparison-item">
        <span class="type">RELATIVE</span>
        <h3>Desplazamiento Relativo</h3>
        <p>Se desplaza desde su posición normal pero mantiene su espacio original en el layout.</p>
      </div>

      <div class="comparison-item">
        <span class="type">ABSOLUTE</span>
        <h3>Posicionamiento Absoluto</h3>
        <p>Se posiciona respecto a su ancestro posicionado más cercano, fuera del flujo normal.</p>
      </div>

      <div class="comparison-item">
        <span class="type">FIXED</span>
        <h3>Posicionamiento Fijo</h3>
        <p>Se posiciona respecto a la ventana del navegador y permanece fijo durante el scroll.</p>
      </div>

      <div class="comparison-item">
        <span class="type">STICKY</span>
        <h3>Posicionamiento Pegajoso</h3>
        <p>Comportamiento híbrido: relative hasta una posición específica, luego fixed.</p>
      </div>
    </div>
  </div>

  <!-- Fixed element demo -->
  <div class="box fixed-box">Fixed</div>
</body>
</html>
```

## Z-Index y Contextos de Apilamiento

### Controlando el orden de apilamiento

```css
.layer-1 {
  position: relative;
  z-index: 1;
}

.layer-2 {
  position: relative;
  z-index: 2;
}

.layer-3 {
  position: absolute;
  z-index: 10;
}
```

### Contextos de apilamiento

```css
.stacking-context {
  position: relative;
  z-index: 1; /* Crea un nuevo contexto de apilamiento */
}

.child-element {
  position: relative;
  z-index: 999; /* Solo afecta dentro de su contexto */
}
```

## Técnicas Avanzadas de Positioning

### Centrado absoluto

```css
.centered-absolute {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
```

### Sticky navigation

```css
.sticky-nav {
  position: sticky;
  top: 0;
  background: white;
  z-index: 100;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}
```

### Modal overlays

```css
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  z-index: 1000;
}

.modal-content {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: white;
  z-index: 1001;
}
```

## Ejemplo Complejo: Layout con Positioning

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Positioning - Layout Complejo</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: #1a1a1a;
      color: #e0e0e0;
      line-height: 1.6;
    }

    .app {
      position: relative;
      min-height: 100vh;
    }

    .sidebar {
      position: fixed;
      top: 0;
      left: 0;
      width: 280px;
      height: 100vh;
      background: linear-gradient(180deg, #2d3748 0%, #1a202c 100%);
      padding: 30px 20px;
      box-shadow: 2px 0 10px rgba(0,0,0,0.3);
      z-index: 200;
      overflow-y: auto;
    }

    .sidebar-header {
      margin-bottom: 30px;
    }

    .sidebar-title {
      font-size: 1.5rem;
      font-weight: bold;
      color: #4ecdc4;
      margin-bottom: 10px;
    }

    .sidebar-subtitle {
      color: #a0aec0;
      font-size: 0.9rem;
    }

    .nav-item {
      display: block;
      color: #e2e8f0;
      text-decoration: none;
      padding: 12px 16px;
      margin-bottom: 8px;
      border-radius: 8px;
      transition: all 0.2s ease;
      position: relative;
    }

    .nav-item:hover {
      background: rgba(78, 205, 196, 0.1);
      color: #4ecdc4;
      transform: translateX(5px);
    }

    .nav-item.active {
      background: rgba(78, 205, 196, 0.2);
      color: #4ecdc4;
      border-left: 3px solid #4ecdc4;
    }

    .main-content {
      margin-left: 280px;
      padding: 30px;
      min-height: 100vh;
    }

    .header {
      position: sticky;
      top: 0;
      background: rgba(26, 26, 26, 0.95);
      backdrop-filter: blur(10px);
      padding: 20px 0;
      margin: -30px -30px 30px -30px;
      border-bottom: 1px solid #2d3748;
      z-index: 100;
    }

    .header-content {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .page-title {
      font-size: 2rem;
      font-weight: bold;
      color: #e0e0e0;
    }

    .header-actions {
      display: flex;
      gap: 15px;
    }

    .btn {
      padding: 10px 20px;
      border: none;
      border-radius: 6px;
      cursor: pointer;
      font-weight: 500;
      transition: all 0.2s ease;
    }

    .btn-primary {
      background: linear-gradient(45deg, #667eea, #764ba2);
      color: white;
    }

    .btn-primary:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
    }

    .content-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
      gap: 25px;
      margin-bottom: 50px;
    }

    .card {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 25px;
      border: 1px solid rgba(255,255,255,0.1);
      position: relative;
      overflow: hidden;
      transition: all 0.3s ease;
    }

    .card:hover {
      transform: translateY(-5px);
      box-shadow: 0 10px 30px rgba(0,0,0,0.3);
      border-color: rgba(78, 205, 196, 0.3);
    }

    .card-icon {
      width: 50px;
      height: 50px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      border-radius: 10px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 24px;
      margin-bottom: 20px;
    }

    .card-title {
      font-size: 1.3rem;
      font-weight: bold;
      margin-bottom: 15px;
      color: #e0e0e0;
    }

    .card-text {
      color: #a0aec0;
      line-height: 1.6;
    }

    .floating-action-btn {
      position: fixed;
      bottom: 30px;
      right: 30px;
      width: 60px;
      height: 60px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      border: none;
      border-radius: 50%;
      color: white;
      font-size: 24px;
      cursor: pointer;
      box-shadow: 0 4px 20px rgba(255, 107, 107, 0.4);
      z-index: 300;
      transition: all 0.3s ease;
    }

    .floating-action-btn:hover {
      transform: scale(1.1);
      box-shadow: 0 6px 25px rgba(255, 107, 107, 0.6);
    }

    .modal-overlay {
      position: fixed;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: rgba(0,0,0,0.8);
      display: none;
      z-index: 1000;
      backdrop-filter: blur(5px);
    }

    .modal-content {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      background: #2d3748;
      padding: 30px;
      border-radius: 12px;
      max-width: 500px;
      width: 90%;
      box-shadow: 0 20px 40px rgba(0,0,0,0.5);
    }

    .modal-header {
      margin-bottom: 20px;
    }

    .modal-title {
      font-size: 1.5rem;
      color: #e0e0e0;
      margin-bottom: 10px;
    }

    .modal-close {
      position: absolute;
      top: 15px;
      right: 15px;
      background: none;
      border: none;
      color: #a0aec0;
      font-size: 24px;
      cursor: pointer;
      padding: 5px;
      border-radius: 50%;
      transition: all 0.2s ease;
    }

    .modal-close:hover {
      background: rgba(255,255,255,0.1);
      color: #e0e0e0;
    }

    .modal-body p {
      color: #a0aec0;
      line-height: 1.6;
      margin-bottom: 20px;
    }

    .modal-actions {
      display: flex;
      justify-content: flex-end;
      gap: 10px;
    }

    .tooltip {
      position: relative;
      display: inline-block;
    }

    .tooltip-text {
      position: absolute;
      bottom: 125%;
      left: 50%;
      transform: translateX(-50%);
      background: #2d3748;
      color: #e0e0e0;
      padding: 8px 12px;
      border-radius: 6px;
      font-size: 14px;
      white-space: nowrap;
      opacity: 0;
      visibility: hidden;
      transition: all 0.3s ease;
      z-index: 400;
    }

    .tooltip-text::after {
      content: '';
      position: absolute;
      top: 100%;
      left: 50%;
      transform: translateX(-50%);
      border: 5px solid transparent;
      border-top-color: #2d3748;
    }

    .tooltip:hover .tooltip-text {
      opacity: 1;
      visibility: visible;
    }

    @media (max-width: 768px) {
      .sidebar {
        transform: translateX(-100%);
        transition: transform 0.3s ease;
      }

      .sidebar.open {
        transform: translateX(0);
      }

      .main-content {
        margin-left: 0;
      }

      .menu-toggle {
        display: block;
        position: fixed;
        top: 20px;
        left: 20px;
        z-index: 250;
        background: #2d3748;
        border: none;
        color: #e0e0e0;
        padding: 10px;
        border-radius: 6px;
        cursor: pointer;
      }
    }

    .menu-toggle {
      display: none;
    }
  </style>
</head>
<body>
  <div class="app">
    <aside class="sidebar">
      <div class="sidebar-header">
        <h1 class="sidebar-title">Dashboard</h1>
        <p class="sidebar-subtitle">Panel de control</p>
      </div>

      <nav>
        <a href="#" class="nav-item active">🏠 Inicio</a>
        <a href="#" class="nav-item">📊 Analytics</a>
        <a href="#" class="nav-item">👥 Usuarios</a>
        <a href="#" class="nav-item">⚙️ Configuración</a>
        <a href="#" class="nav-item">📝 Reportes</a>
        <a href="#" class="nav-item">🔔 Notificaciones</a>
      </nav>
    </aside>

    <main class="main-content">
      <header class="header">
        <div class="header-content">
          <h1 class="page-title">Panel de Control</h1>
          <div class="header-actions">
            <button class="btn btn-primary tooltip">
              Exportar
              <span class="tooltip-text">Exportar datos a CSV</span>
            </button>
            <button class="btn btn-primary" onclick="openModal()">Nuevo</button>
          </div>
        </div>
      </header>

      <div class="content-grid">
        <div class="card">
          <div class="card-icon">📈</div>
          <h3 class="card-title">Estadísticas</h3>
          <p class="card-text">Visualiza el rendimiento de tu aplicación con gráficos interactivos y métricas en tiempo real.</p>
        </div>

        <div class="card">
          <div class="card-icon">👤</div>
          <h3 class="card-title">Usuarios</h3>
          <p class="card-text">Gestiona los usuarios de tu plataforma con herramientas avanzadas de administración.</p>
        </div>

        <div class="card">
          <div class="card-icon">💰</div>
          <h3 class="card-title">Ingresos</h3>
          <p class="card-text">Monitorea los ingresos y analiza tendencias financieras con reportes detallados.</p>
        </div>

        <div class="card">
          <div class="card-icon">🔧</div>
          <h3 class="card-title">Configuración</h3>
          <p class="card-text">Personaliza la configuración de tu aplicación según tus necesidades específicas.</p>
        </div>
      </div>
    </main>

    <button class="floating-action-btn" onclick="openModal()">+</button>

    <div class="modal-overlay" id="modal">
      <div class="modal-content">
        <button class="modal-close" onclick="closeModal()">×</button>
        <div class="modal-header">
          <h2 class="modal-title">Crear Nuevo Elemento</h2>
        </div>
        <div class="modal-body">
          <p>Esta es una ventana modal que demuestra el uso de posicionamiento absoluto para centrar contenido sobre otros elementos.</p>
          <p>El overlay utiliza <code>position: fixed</code> para cubrir toda la pantalla, mientras que el contenido modal se centra con <code>position: absolute</code> y <code>transform: translate()</code>.</p>
        </div>
        <div class="modal-actions">
          <button class="btn btn-primary" onclick="closeModal()">Cerrar</button>
        </div>
      </div>
    </div>
  </div>

  <script>
    function openModal() {
      document.getElementById('modal').style.display = 'block';
    }

    function closeModal() {
      document.getElementById('modal').style.display = 'none';
    }

    // Cerrar modal al hacer click fuera
    document.getElementById('modal').addEventListener('click', function(e) {
      if (e.target === this) {
        closeModal();
      }
    });
  </script>
</body>
</html>
```

## Casos de Uso Prácticos

### 1. Tooltips

```css
.tooltip {
  position: relative;
  display: inline-block;
}

.tooltip-text {
  position: absolute;
  bottom: 125%;
  left: 50%;
  transform: translateX(-50%);
  visibility: hidden;
  opacity: 0;
}

.tooltip:hover .tooltip-text {
  visibility: visible;
  opacity: 1;
}
```

### 2. Dropdown menus

```css
.dropdown {
  position: relative;
}

.dropdown-content {
  position: absolute;
  top: 100%;
  left: 0;
  display: none;
}

.dropdown:hover .dropdown-content {
  display: block;
}
```

### 3. Image overlays

```css
.image-container {
  position: relative;
}

.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.image-container:hover .image-overlay {
  opacity: 1;
}
```

## Mejores Prácticas

### 1. Evitar sobreposicionamiento innecesario

```css
/* ✅ Bueno: usar el posicionamiento correcto */
.sticky-header {
  position: sticky;
  top: 0;
}

/* ❌ Evitar: absolute cuando no es necesario */
.unnecessary-absolute {
  position: absolute;
  top: 0;
  left: 0;
}
```

### 2. Gestionar z-index correctamente

```css
/* Sistema de z-index consistente */
.modal { z-index: 1000; }
.dropdown { z-index: 100; }
.tooltip { z-index: 200; }
.sticky-nav { z-index: 50; }
```

### 3. Considerar performance

```css
/* Usar transform para animaciones en lugar de cambiar top/left */
.animated-element {
  position: absolute;
  transition: transform 0.3s ease;
}

.animated-element.move {
  transform: translateX(100px);
}
```

## Resumen

CSS Positioning es fundamental para crear layouts complejos:

- ✅ **Static**: Comportamiento normal del flujo del documento
- ✅ **Relative**: Desplazamiento desde posición original
- ✅ **Absolute**: Posicionamiento respecto a ancestro posicionado
- ✅ **Fixed**: Posicionamiento fijo respecto a la ventana
- ✅ **Sticky**: Comportamiento híbrido relative/fixed
- ✅ **Z-index**: Control del orden de apilamiento
- ✅ **Contextos de apilamiento**: Gestión de capas complejas

El posicionamiento CSS es esencial para crear interfaces modernas, modales, tooltips, menús desplegables, y cualquier elemento que necesite salir del flujo normal del documento.