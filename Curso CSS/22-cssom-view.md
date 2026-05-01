# Módulo 22 - CSSOM View

En este módulo aprenderás sobre CSSOM View API, una interfaz de programación que permite acceder y manipular la geometría visual de los elementos en el viewport, incluyendo métodos para obtener posiciones, dimensiones, y controlar el scroll.

## Introducción a CSSOM View

### ¿Qué es CSSOM View?

CSSOM View es una API de JavaScript que proporciona métodos y propiedades para acceder a información geométrica de elementos DOM, como posiciones, dimensiones, y estado del viewport. Es parte de las especificaciones de CSS Object Model.

```javascript
// Obtener el rectángulo delimitador de un elemento
const rect = element.getBoundingClientRect();

// Obtener dimensiones del viewport
const viewportWidth = window.innerWidth;
const viewportHeight = window.innerHeight;

// Controlar el scroll
window.scrollTo(0, 500);
```

## Métodos Principales

### getBoundingClientRect()

```javascript
const element = document.getElementById('myElement');
const rect = element.getBoundingClientRect();

console.log('Posición X:', rect.left);
console.log('Posición Y:', rect.top);
console.log('Ancho:', rect.width);
console.log('Alto:', rect.height);
console.log('Derecha:', rect.right);
console.log('Abajo:', rect.bottom);
```

### Propiedades del Viewport

```javascript
// Dimensiones del viewport
const viewportWidth = window.innerWidth;
const viewportHeight = window.innerHeight;

// Dimensiones incluyendo scrollbars
const outerWidth = window.outerWidth;
const outerHeight = window.outerHeight;

// Posición del scroll
const scrollX = window.scrollX || window.pageXOffset;
const scrollY = window.scrollY || window.pageYOffset;
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSSOM View - Básico</title>
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
    }

    .container {
      max-width: 1000px;
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

    .element-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
      gap: 20px;
      margin-bottom: 30px;
    }

    .demo-element {
      height: 120px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      border-radius: 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-weight: bold;
      cursor: pointer;
      transition: transform 0.2s ease;
      position: relative;
    }

    .demo-element:hover {
      transform: translateY(-5px);
    }

    .element-info {
      background: rgba(0,0,0,0.8);
      color: white;
      padding: 15px;
      border-radius: 6px;
      font-size: 12px;
      position: absolute;
      top: -10px;
      left: -10px;
      opacity: 0;
      transition: opacity 0.2s ease;
      pointer-events: none;
      z-index: 10;
    }

    .demo-element:hover .element-info {
      opacity: 1;
    }

    .viewport-info {
      background: #f8f9fa;
      border-radius: 8px;
      padding: 20px;
      margin-bottom: 20px;
      border: 1px solid #e0e0e0;
    }

    .info-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
      gap: 15px;
    }

    .info-item {
      text-align: center;
    }

    .info-label {
      font-size: 12px;
      color: #666;
      margin-bottom: 5px;
      text-transform: uppercase;
      letter-spacing: 0.5px;
    }

    .info-value {
      font-size: 18px;
      font-weight: bold;
      color: #333;
    }

    .scroll-controls {
      display: flex;
      gap: 15px;
      margin-bottom: 20px;
      flex-wrap: wrap;
    }

    .scroll-btn {
      padding: 10px 20px;
      background: #667eea;
      color: white;
      border: none;
      border-radius: 6px;
      cursor: pointer;
      font-size: 14px;
      transition: background-color 0.2s ease;
    }

    .scroll-btn:hover {
      background: #5a67d8;
    }

    .scroll-btn:disabled {
      background: #ccc;
      cursor: not-allowed;
    }

    .long-content {
      height: 1000px;
      background: linear-gradient(to bottom,
        #ff6b6b 0%,
        #4ecdc4 25%,
        #45b7d1 50%,
        #96ceb4 75%,
        #667eea 100%);
      border-radius: 8px;
      margin: 20px 0;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-size: 24px;
      font-weight: bold;
    }

    .intersection-demo {
      margin-top: 50px;
      background: white;
      border-radius: 12px;
      padding: 30px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
    }

    .intersection-demo h2 {
      margin-bottom: 20px;
      color: #333;
      font-size: 1.8rem;
    }

    .intersection-target {
      height: 100px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      border-radius: 8px;
      margin: 20px 0;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-weight: bold;
      transition: opacity 0.3s ease;
    }

    .intersection-status {
      padding: 15px;
      background: #f8f9fa;
      border-radius: 6px;
      border: 1px solid #e0e0e0;
      font-family: monospace;
      font-size: 14px;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>CSSOM View API</h1>
      <p>Accediendo a la geometría visual de elementos</p>
    </header>

    <div class="demo-section">
      <h2>getBoundingClientRect()</h2>

      <div class="element-grid">
        <div class="demo-element" id="element1">
          Elemento 1
          <div class="element-info" id="info1">
            Cargando...
          </div>
        </div>

        <div class="demo-element" id="element2">
          Elemento 2
          <div class="element-info" id="info2">
            Cargando...
          </div>
        </div>

        <div class="demo-element" id="element3">
          Elemento 3
          <div class="element-info" id="info3">
            Cargando...
          </div>
        </div>

        <div class="demo-element" id="element4">
          Elemento 4
          <div class="element-info" id="info4">
            Cargando...
          </div>
        </div>
      </div>

      <div class="viewport-info">
        <div class="info-grid">
          <div class="info-item">
            <div class="info-label">Viewport Width</div>
            <div class="info-value" id="viewportWidth">--</div>
          </div>
          <div class="info-item">
            <div class="info-label">Viewport Height</div>
            <div class="info-value" id="viewportHeight">--</div>
          </div>
          <div class="info-item">
            <div class="info-label">Scroll X</div>
            <div class="info-value" id="scrollX">--</div>
          </div>
          <div class="info-item">
            <div class="info-label">Scroll Y</div>
            <div class="info-value" id="scrollY">--</div>
          </div>
        </div>
      </div>

      <div class="scroll-controls">
        <button class="scroll-btn" onclick="scrollToTop()">Ir Arriba</button>
        <button class="scroll-btn" onclick="scrollToBottom()">Ir Abajo</button>
        <button class="scroll-btn" onclick="scrollToElement('element3')">Ir a Elemento 3</button>
        <button class="scroll-btn" onclick="smoothScroll()">Scroll Suave</button>
      </div>
    </div>

    <div class="long-content">
      Contenido Largo - Scroll para ver más
    </div>

    <div class="intersection-demo">
      <h2>Intersection Observer</h2>

      <div class="intersection-target" id="intersectionTarget">
        Elemento a Observar
      </div>

      <div class="intersection-status" id="intersectionStatus">
        Estado: --
      </div>
    </div>
  </div>

  <script>
    // Función para actualizar información de elementos
    function updateElementInfo() {
      for (let i = 1; i <= 4; i++) {
        const element = document.getElementById(`element${i}`);
        const info = document.getElementById(`info${i}`);

        if (element && info) {
          const rect = element.getBoundingClientRect();
          info.innerHTML = `
            X: ${Math.round(rect.left)}<br>
            Y: ${Math.round(rect.top)}<br>
            W: ${Math.round(rect.width)}<br>
            H: ${Math.round(rect.height)}
          `;
        }
      }
    }

    // Función para actualizar información del viewport
    function updateViewportInfo() {
      const viewportWidth = document.getElementById('viewportWidth');
      const viewportHeight = document.getElementById('viewportHeight');
      const scrollX = document.getElementById('scrollX');
      const scrollY = document.getElementById('scrollY');

      viewportWidth.textContent = window.innerWidth;
      viewportHeight.textContent = window.innerHeight;
      scrollX.textContent = Math.round(window.scrollX || window.pageXOffset);
      scrollY.textContent = Math.round(window.scrollY || window.pageYOffset);
    }

    // Funciones de scroll
    function scrollToTop() {
      window.scrollTo({ top: 0, behavior: 'smooth' });
    }

    function scrollToBottom() {
      window.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' });
    }

    function scrollToElement(elementId) {
      const element = document.getElementById(elementId);
      if (element) {
        element.scrollIntoView({ behavior: 'smooth', block: 'center' });
      }
    }

    function smoothScroll() {
      window.scrollTo({
        top: window.innerHeight,
        behavior: 'smooth'
      });
    }

    // Intersection Observer
    const intersectionTarget = document.getElementById('intersectionTarget');
    const intersectionStatus = document.getElementById('intersectionStatus');

    const observer = new IntersectionObserver((entries) => {
      entries.forEach(entry => {
        const isIntersecting = entry.isIntersecting;
        const intersectionRatio = Math.round(entry.intersectionRatio * 100);

        intersectionStatus.innerHTML = `
          Estado: ${isIntersecting ? 'Visible' : 'No visible'}<br>
          Ratio: ${intersectionRatio}%<br>
          Rect: ${Math.round(entry.boundingClientRect.top)}, ${Math.round(entry.boundingClientRect.left)}
        `;

        intersectionTarget.style.opacity = isIntersecting ? '1' : '0.5';
      });
    }, {
      threshold: [0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0]
    });

    observer.observe(intersectionTarget);

    // Event listeners
    window.addEventListener('scroll', () => {
      updateElementInfo();
      updateViewportInfo();
    });

    window.addEventListener('resize', () => {
      updateElementInfo();
      updateViewportInfo();
    });

    // Inicializar
    updateElementInfo();
    updateViewportInfo();
  </script>
</body>
</html>
```

## Métodos de Scroll

### scrollTo(), scrollBy(), scroll()

```javascript
// Scroll a posición específica
window.scrollTo(0, 500);

// Scroll con opciones
window.scrollTo({
  top: 500,
  left: 0,
  behavior: 'smooth'
});

// Scroll relativo
window.scrollBy(0, 100);

// Scroll a elemento
element.scrollIntoView();
element.scrollIntoView({ behavior: 'smooth', block: 'center' });
```

### Control de scroll programático

```javascript
// Scroll horizontal y vertical
window.scrollTo({
  top: 1000,
  left: 500,
  behavior: 'smooth'
});

// Scroll solo vertical
window.scrollTo({
  top: 1000,
  behavior: 'auto'
});
```

## Intersection Observer API

### Detectando visibilidad de elementos

```javascript
const observer = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      // Elemento visible
      entry.target.classList.add('visible');
    } else {
      // Elemento no visible
      entry.target.classList.remove('visible');
    }
  });
}, {
  threshold: 0.5, // 50% del elemento visible
  rootMargin: '10px' // Margen alrededor del viewport
});

observer.observe(document.querySelector('.elemento-a-observar'));
```

## Ejemplo Avanzado: Infinite Scroll

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSSOM View - Infinite Scroll</title>
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
    }

    .app {
      max-width: 800px;
      margin: 0 auto;
      padding: 20px;
    }

    .header {
      text-align: center;
      margin-bottom: 40px;
    }

    .header h1 {
      font-size: 2.5rem;
      margin-bottom: 10px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    .controls {
      display: flex;
      justify-content: center;
      gap: 20px;
      margin-bottom: 30px;
      flex-wrap: wrap;
    }

    .control-group {
      display: flex;
      flex-direction: column;
      align-items: center;
    }

    .control-group label {
      margin-bottom: 8px;
      font-size: 14px;
      color: #a0aec0;
    }

    .control-group select,
    .control-group input {
      padding: 8px 12px;
      border: 1px solid rgba(255,255,255,0.2);
      border-radius: 6px;
      background: rgba(255,255,255,0.05);
      color: #e0e0e0;
      font-size: 14px;
    }

    .control-group select:focus,
    .control-group input:focus {
      outline: none;
      border-color: #4ecdc4;
    }

    .content-container {
      position: relative;
    }

    .content-list {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
      gap: 20px;
      margin-bottom: 40px;
    }

    .content-item {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 20px;
      border: 1px solid rgba(255,255,255,0.1);
      transition: all 0.3s ease;
      opacity: 0;
      transform: translateY(20px);
    }

    .content-item.visible {
      opacity: 1;
      transform: translateY(0);
    }

    .content-item:hover {
      transform: translateY(-5px);
      border-color: rgba(78, 205, 196, 0.3);
    }

    .item-header {
      display: flex;
      align-items: center;
      margin-bottom: 15px;
    }

    .item-avatar {
      width: 40px;
      height: 40px;
      border-radius: 50%;
      background: linear-gradient(45deg, #667eea, #764ba2);
      margin-right: 15px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-weight: bold;
    }

    .item-title {
      font-size: 1.1rem;
      font-weight: bold;
      color: #e0e0e0;
    }

    .item-content {
      color: #a0aec0;
      line-height: 1.6;
      margin-bottom: 15px;
    }

    .item-meta {
      display: flex;
      justify-content: space-between;
      font-size: 12px;
      color: #718096;
    }

    .loading-indicator {
      text-align: center;
      padding: 40px;
      color: #a0aec0;
    }

    .loading-spinner {
      width: 40px;
      height: 40px;
      border: 3px solid rgba(78, 205, 196, 0.3);
      border-top: 3px solid #4ecdc4;
      border-radius: 50%;
      animation: spin 1s linear infinite;
      margin: 0 auto 20px;
    }

    @keyframes spin {
      0% { transform: rotate(0deg); }
      100% { transform: rotate(360deg); }
    }

    .end-message {
      text-align: center;
      padding: 40px;
      color: #a0aec0;
      font-style: italic;
    }

    .scroll-to-top {
      position: fixed;
      bottom: 30px;
      right: 30px;
      width: 50px;
      height: 50px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      border: none;
      border-radius: 50%;
      color: white;
      font-size: 20px;
      cursor: pointer;
      box-shadow: 0 4px 20px rgba(0,0,0,0.3);
      opacity: 0;
      transform: translateY(20px);
      transition: all 0.3s ease;
      z-index: 1000;
    }

    .scroll-to-top.visible {
      opacity: 1;
      transform: translateY(0);
    }

    .stats {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 20px;
      margin-bottom: 30px;
      border: 1px solid rgba(255,255,255,0.1);
    }

    .stats-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
      gap: 15px;
    }

    .stat-item {
      text-align: center;
    }

    .stat-value {
      font-size: 24px;
      font-weight: bold;
      color: #4ecdc4;
      display: block;
    }

    .stat-label {
      font-size: 12px;
      color: #a0aec0;
      text-transform: uppercase;
      letter-spacing: 0.5px;
    }
  </style>
</head>
<body>
  <div class="app">
    <header class="header">
      <h1>Infinite Scroll Demo</h1>
      <p>Demostración de CSSOM View con scroll infinito</p>
    </header>

    <div class="stats">
      <div class="stats-grid">
        <div class="stat-item">
          <span class="stat-value" id="totalItems">0</span>
          <span class="stat-label">Items</span>
        </div>
        <div class="stat-item">
          <span class="stat-value" id="visibleItems">0</span>
          <span class="stat-label">Visibles</span>
        </div>
        <div class="stat-item">
          <span class="stat-value" id="scrollPosition">0</span>
          <span class="stat-label">Scroll Y</span>
        </div>
        <div class="stat-item">
          <span class="stat-value" id="viewportHeight">--</span>
          <span class="stat-label">Viewport</span>
        </div>
      </div>
    </div>

    <div class="controls">
      <div class="control-group">
        <label for="itemsPerLoad">Items por carga:</label>
        <select id="itemsPerLoad">
          <option value="5">5</option>
          <option value="10" selected>10</option>
          <option value="20">20</option>
          <option value="50">50</option>
        </select>
      </div>

      <div class="control-group">
        <label for="loadDelay">Delay de carga (ms):</label>
        <input type="number" id="loadDelay" value="500" min="0" max="2000" step="100">
      </div>

      <div class="control-group">
        <label for="intersectionThreshold">Threshold:</label>
        <input type="number" id="intersectionThreshold" value="0.1" min="0" max="1" step="0.1">
      </div>
    </div>

    <div class="content-container">
      <div class="content-list" id="contentList"></div>

      <div class="loading-indicator" id="loadingIndicator">
        <div class="loading-spinner"></div>
        <p>Cargando más contenido...</p>
      </div>

      <div class="end-message" id="endMessage" style="display: none;">
        <p>¡Has llegado al final! No hay más contenido para cargar.</p>
      </div>
    </div>

    <button class="scroll-to-top" id="scrollToTop">↑</button>
  </div>

  <script>
    class InfiniteScrollDemo {
      constructor() {
        this.contentList = document.getElementById('contentList');
        this.loadingIndicator = document.getElementById('loadingIndicator');
        this.endMessage = document.getElementById('endMessage');
        this.scrollToTopBtn = document.getElementById('scrollToTop');

        this.itemsPerLoad = parseInt(document.getElementById('itemsPerLoad').value);
        this.loadDelay = parseInt(document.getElementById('loadDelay').value);
        this.intersectionThreshold = parseFloat(document.getElementById('intersectionThreshold').value);

        this.totalItems = 0;
        this.isLoading = false;
        this.hasMoreContent = true;
        this.maxItems = 200; // Límite para demo

        this.init();
      }

      init() {
        // Crear Intersection Observer
        this.observer = new IntersectionObserver((entries) => {
          entries.forEach(entry => {
            if (entry.isIntersecting && !this.isLoading && this.hasMoreContent) {
              this.loadMoreItems();
            }
          });
        }, {
          threshold: this.intersectionThreshold
        });

        // Crear elemento sentinel
        this.sentinel = document.createElement('div');
        this.sentinel.style.height = '10px';
        this.contentList.appendChild(this.sentinel);
        this.observer.observe(this.sentinel);

        // Event listeners para controles
        this.setupControls();

        // Scroll to top button
        this.setupScrollToTop();

        // Carga inicial
        this.loadMoreItems();

        // Actualizar stats
        this.updateStats();
      }

      setupControls() {
        document.getElementById('itemsPerLoad').addEventListener('change', (e) => {
          this.itemsPerLoad = parseInt(e.target.value);
        });

        document.getElementById('loadDelay').addEventListener('change', (e) => {
          this.loadDelay = parseInt(e.target.value);
        });

        document.getElementById('intersectionThreshold').addEventListener('input', (e) => {
          this.intersectionThreshold = parseFloat(e.target.value);
          this.observer.threshold = this.intersectionThreshold;
        });
      }

      setupScrollToTop() {
        window.addEventListener('scroll', () => {
          const scrolled = window.scrollY;
          const scrollThreshold = 300;

          if (scrolled > scrollThreshold) {
            this.scrollToTopBtn.classList.add('visible');
          } else {
            this.scrollToTopBtn.classList.remove('visible');
          }

          this.updateStats();
        });

        this.scrollToTopBtn.addEventListener('click', () => {
          window.scrollTo({
            top: 0,
            behavior: 'smooth'
          });
        });
      }

      async loadMoreItems() {
        if (this.isLoading || !this.hasMoreContent) return;

        this.isLoading = true;
        this.loadingIndicator.style.display = 'block';

        // Simular delay de carga
        await new Promise(resolve => setTimeout(resolve, this.loadDelay));

        const newItems = [];
        for (let i = 0; i < this.itemsPerLoad; i++) {
          if (this.totalItems >= this.maxItems) {
            this.hasMoreContent = false;
            break;
          }

          const item = this.createItem(this.totalItems + 1);
          newItems.push(item);
          this.totalItems++;
        }

        // Agregar items al DOM
        newItems.forEach(item => {
          this.contentList.insertBefore(item, this.sentinel);
        });

        // Animar items visibles
        setTimeout(() => {
          const visibleItems = this.getVisibleItems();
          visibleItems.forEach(item => {
            item.classList.add('visible');
          });
        }, 100);

        this.isLoading = false;
        this.loadingIndicator.style.display = 'none';

        if (!this.hasMoreContent) {
          this.endMessage.style.display = 'block';
          this.observer.unobserve(this.sentinel);
        }

        this.updateStats();
      }

      createItem(index) {
        const item = document.createElement('div');
        item.className = 'content-item';

        const colors = ['#ff6b6b', '#4ecdc4', '#45b7d1', '#96ceb4', '#667eea', '#764ba2'];
        const randomColor = colors[Math.floor(Math.random() * colors.length)];

        item.innerHTML = `
          <div class="item-header">
            <div class="item-avatar" style="background: linear-gradient(45deg, ${randomColor}, ${this.adjustColor(randomColor, -20)})">
              ${index}
            </div>
            <div class="item-title">Item ${index}</div>
          </div>
          <div class="item-content">
            Este es el elemento número ${index} cargado dinámicamente. Los elementos se cargan automáticamente cuando el usuario hace scroll hacia abajo y el elemento sentinel entra en el viewport.
          </div>
          <div class="item-meta">
            <span>Hace ${Math.floor(Math.random() * 24)} horas</span>
            <span>${Math.floor(Math.random() * 100)} vistas</span>
          </div>
        `;

        return item;
      }

      adjustColor(color, amount) {
        // Función simple para ajustar el brillo del color
        const usePound = color[0] === '#';
        const col = usePound ? color.slice(1) : color;

        const num = parseInt(col, 16);
        let r = (num >> 16) + amount;
        let g = (num >> 8 & 0x00FF) + amount;
        let b = (num & 0x0000FF) + amount;

        r = r > 255 ? 255 : r < 0 ? 0 : r;
        g = g > 255 ? 255 : g < 0 ? 0 : g;
        b = b > 255 ? 255 : b < 0 ? 0 : b;

        return (usePound ? '#' : '') + (r << 16 | g << 8 | b).toString(16);
      }

      getVisibleItems() {
        const items = Array.from(document.querySelectorAll('.content-item'));
        return items.filter(item => {
          const rect = item.getBoundingClientRect();
          return rect.top < window.innerHeight && rect.bottom > 0;
        });
      }

      updateStats() {
        document.getElementById('totalItems').textContent = this.totalItems;
        document.getElementById('visibleItems').textContent = this.getVisibleItems().length;
        document.getElementById('scrollPosition').textContent = Math.round(window.scrollY);
        document.getElementById('viewportHeight').textContent = window.innerHeight;
      }
    }

    // Inicializar demo
    new InfiniteScrollDemo();
  </script>
</body>
</html>
```

## Resize Observer

### Detectando cambios de tamaño

```javascript
const resizeObserver = new ResizeObserver(entries => {
  entries.forEach(entry => {
    const { width, height } = entry.contentRect;
    console.log('Nuevo tamaño:', width, height);

    // Ajustar layout según el nuevo tamaño
    if (width < 768) {
      entry.target.classList.add('mobile');
    } else {
      entry.target.classList.remove('mobile');
    }
  });
});

resizeObserver.observe(document.querySelector('.responsive-element'));
```

## Técnicas Avanzadas

### Lazy Loading con Intersection Observer

```javascript
const lazyImages = document.querySelectorAll('img[data-src]');

const imageObserver = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      const img = entry.target;
      img.src = img.dataset.src;
      img.classList.remove('lazy');
      imageObserver.unobserve(img);
    }
  });
});

lazyImages.forEach(img => imageObserver.observe(img));
```

### Scroll-based Animations

```javascript
const animatedElements = document.querySelectorAll('.animate-on-scroll');

const scrollObserver = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      entry.target.style.animationPlayState = 'running';
    }
  });
}, {
  threshold: 0.1
});

animatedElements.forEach(el => {
  el.style.animationPlayState = 'paused';
  scrollObserver.observe(el);
});
```

## Casos de Uso Prácticos

### 1. Sticky Navigation

```javascript
const nav = document.querySelector('.nav');
const navTop = nav.offsetTop;

window.addEventListener('scroll', () => {
  if (window.scrollY >= navTop) {
    nav.classList.add('sticky');
  } else {
    nav.classList.remove('sticky');
  }
});
```

### 2. Parallax Scrolling

```javascript
window.addEventListener('scroll', () => {
  const scrolled = window.scrollY;
  const parallaxElements = document.querySelectorAll('.parallax');

  parallaxElements.forEach(element => {
    const speed = element.dataset.speed || 0.5;
    element.style.transform = `translateY(${scrolled * speed}px)`;
  });
});
```

### 3. Infinite Scroll

```javascript
let loading = false;

window.addEventListener('scroll', () => {
  if (window.innerHeight + window.scrollY >= document.body.offsetHeight - 1000 && !loading) {
    loading = true;
    loadMoreContent().then(() => {
      loading = false;
    });
  }
});
```

## Performance Considerations

### Optimizaciones

```javascript
// Throttling para eventos de scroll
let scrollTimeout;
window.addEventListener('scroll', () => {
  if (!scrollTimeout) {
    scrollTimeout = setTimeout(() => {
      // Código de scroll aquí
      scrollTimeout = null;
    }, 16); // ~60fps
  }
});

// Usar passive listeners
window.addEventListener('scroll', handleScroll, { passive: true });
```

## Compatibilidad

### Fallbacks

```javascript
// Fallback para getBoundingClientRect en navegadores antiguos
function getBoundingRect(element) {
  if (element.getBoundingClientRect) {
    return element.getBoundingClientRect();
  }

  // Fallback manual
  const rect = {
    left: 0,
    top: 0,
    width: element.offsetWidth,
    height: element.offsetHeight
  };

  let el = element;
  while (el) {
    rect.left += el.offsetLeft;
    rect.top += el.offsetTop;
    el = el.offsetParent;
  }

  rect.right = rect.left + rect.width;
  rect.bottom = rect.top + rect.height;

  return rect;
}
```

## Resumen

CSSOM View API proporciona control preciso sobre la geometría visual:

- ✅ **getBoundingClientRect()**: Obtener posición y dimensiones de elementos
- ✅ **Viewport properties**: Acceder a dimensiones y posición del scroll
- ✅ **Scroll methods**: Control programático del desplazamiento
- ✅ **Intersection Observer**: Detectar visibilidad de elementos eficientemente
- ✅ **Resize Observer**: Monitorear cambios de tamaño
- ✅ **Performance**: Técnicas optimizadas para scroll y animaciones

CSSOM View es esencial para crear interfaces interactivas modernas, lazy loading, infinite scroll, y cualquier funcionalidad que requiera conocimiento preciso de la geometría de elementos en el viewport.