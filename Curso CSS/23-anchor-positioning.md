# Módulo 23 - Anchor Positioning

En este módulo aprenderás sobre Anchor Positioning, una nueva característica de CSS que permite posicionar elementos relativos a otros elementos (anchors) de manera declarativa, sin necesidad de JavaScript para cálculos de posición.

## Introducción a Anchor Positioning

### ¿Qué es Anchor Positioning?

Anchor Positioning es una especificación CSS que permite posicionar elementos de manera relativa a otros elementos llamados "anchors". Esto simplifica tareas complejas como tooltips, menús desplegables, popovers, y elementos flotantes.

```css
/* Elemento ancla */
.anchor-element {
  anchor-name: --my-anchor;
}

/* Elemento posicionado */
.positioned-element {
  position: absolute;
  position-anchor: --my-anchor;
  top: anchor(bottom);
  left: anchor(center);
}
```

## Conceptos Básicos

### Anchor Name y Position Anchor

```css
/* Definir un ancla */
.button {
  anchor-name: --button-anchor;
}

/* Referenciar el ancla */
.tooltip {
  position: absolute;
  position-anchor: --button-anchor;
}
```

### Funciones de Anclaje

```css
/* Posicionamiento usando funciones de ancla */
.tooltip {
  position-anchor: --button-anchor;
  top: anchor(bottom);
  left: anchor(start);
  right: anchor(end);
  bottom: anchor(top);
  /* También: anchor(center) */
}
```

## Ejemplo Básico Interactivo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Anchor Positioning - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      min-height: 100vh;
      padding: 40px;
      color: #333;
    }

    .container {
      max-width: 1200px;
      margin: 0 auto;
    }

    .header {
      text-align: center;
      margin-bottom: 50px;
    }

    .header h1 {
      font-size: 3rem;
      color: white;
      margin-bottom: 10px;
      text-shadow: 0 2px 4px rgba(0,0,0,0.3);
    }

    .header p {
      font-size: 1.2rem;
      color: rgba(255,255,255,0.9);
    }

    .demo-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 30px;
      margin-bottom: 50px;
    }

    .demo-card {
      background: rgba(255,255,255,0.95);
      border-radius: 16px;
      padding: 30px;
      box-shadow: 0 8px 32px rgba(0,0,0,0.1);
      backdrop-filter: blur(10px);
      border: 1px solid rgba(255,255,255,0.2);
    }

    .demo-card h3 {
      margin-bottom: 20px;
      color: #333;
      font-size: 1.5rem;
      text-align: center;
    }

    /* Demo 1: Tooltip básico */
    .tooltip-demo {
      position: relative;
      text-align: center;
    }

    .tooltip-button {
      padding: 12px 24px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      color: white;
      border: none;
      border-radius: 8px;
      font-size: 16px;
      cursor: pointer;
      transition: transform 0.2s ease;
      anchor-name: --tooltip-button;
    }

    .tooltip-button:hover {
      transform: translateY(-2px);
    }

    .tooltip {
      position: absolute;
      background: #333;
      color: white;
      padding: 12px 16px;
      border-radius: 6px;
      font-size: 14px;
      opacity: 0;
      pointer-events: none;
      transition: opacity 0.2s ease;
      position-anchor: --tooltip-button;
      top: anchor(bottom);
      left: anchor(center);
      translate: -50% 8px;
      max-width: 200px;
    }

    .tooltip::before {
      content: '';
      position: absolute;
      top: -6px;
      left: 50%;
      transform: translateX(-50%);
      width: 0;
      height: 0;
      border-left: 6px solid transparent;
      border-right: 6px solid transparent;
      border-bottom: 6px solid #333;
    }

    .tooltip-button:hover + .tooltip,
    .tooltip-button:focus + .tooltip {
      opacity: 1;
    }

    /* Demo 2: Menú desplegable */
    .dropdown-demo {
      position: relative;
    }

    .dropdown-button {
      padding: 12px 24px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      color: white;
      border: none;
      border-radius: 8px;
      font-size: 16px;
      cursor: pointer;
      transition: transform 0.2s ease;
      anchor-name: --dropdown-button;
    }

    .dropdown-button:hover {
      transform: translateY(-2px);
    }

    .dropdown-menu {
      position: absolute;
      background: white;
      border-radius: 8px;
      box-shadow: 0 8px 32px rgba(0,0,0,0.2);
      min-width: 200px;
      opacity: 0;
      pointer-events: none;
      transition: opacity 0.2s ease, transform 0.2s ease;
      position-anchor: --dropdown-button;
      top: anchor(bottom);
      left: anchor(start);
      translate: 0 8px;
      border: 1px solid rgba(0,0,0,0.1);
    }

    .dropdown-menu.show {
      opacity: 1;
      pointer-events: auto;
      transform: translateY(0);
    }

    .dropdown-item {
      padding: 12px 16px;
      border-bottom: 1px solid rgba(0,0,0,0.05);
      cursor: pointer;
      transition: background-color 0.2s ease;
    }

    .dropdown-item:hover {
      background: rgba(102, 126, 234, 0.1);
    }

    .dropdown-item:last-child {
      border-bottom: none;
    }

    /* Demo 3: Popover */
    .popover-demo {
      position: relative;
      height: 200px;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .popover-trigger {
      padding: 12px 24px;
      background: linear-gradient(45deg, #96ceb4, #45b7d1);
      color: white;
      border: none;
      border-radius: 8px;
      font-size: 16px;
      cursor: pointer;
      transition: transform 0.2s ease;
      anchor-name: --popover-trigger;
    }

    .popover-trigger:hover {
      transform: translateY(-2px);
    }

    .popover {
      position: absolute;
      background: white;
      border-radius: 12px;
      padding: 24px;
      box-shadow: 0 12px 48px rgba(0,0,0,0.25);
      max-width: 300px;
      opacity: 0;
      pointer-events: none;
      transition: opacity 0.3s ease, transform 0.3s ease;
      position-anchor: --popover-trigger;
      top: anchor(center);
      left: anchor(end);
      translate: 16px -50%;
      border: 1px solid rgba(0,0,0,0.1);
    }

    .popover.show {
      opacity: 1;
      pointer-events: auto;
      transform: translateX(0) translateY(-50%);
    }

    .popover::before {
      content: '';
      position: absolute;
      top: 50%;
      left: -8px;
      transform: translateY(-50%);
      width: 0;
      height: 0;
      border-top: 8px solid transparent;
      border-bottom: 8px solid transparent;
      border-right: 8px solid white;
    }

    .popover h4 {
      margin-bottom: 12px;
      color: #333;
    }

    .popover p {
      color: #666;
      line-height: 1.6;
      margin-bottom: 16px;
    }

    .popover-actions {
      display: flex;
      gap: 12px;
    }

    .popover-btn {
      padding: 8px 16px;
      border: none;
      border-radius: 6px;
      font-size: 14px;
      cursor: pointer;
      transition: background-color 0.2s ease;
    }

    .popover-btn.primary {
      background: #667eea;
      color: white;
    }

    .popover-btn.secondary {
      background: #f8f9fa;
      color: #333;
      border: 1px solid #dee2e6;
    }

    /* Demo 4: Floating Action Button */
    .fab-demo {
      position: relative;
      height: 300px;
      background: rgba(255,255,255,0.1);
      border-radius: 12px;
      display: flex;
      align-items: flex-end;
      justify-content: flex-end;
      padding: 20px;
    }

    .fab {
      width: 56px;
      height: 56px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      border-radius: 50%;
      border: none;
      color: white;
      font-size: 24px;
      cursor: pointer;
      transition: transform 0.2s ease, box-shadow 0.2s ease;
      box-shadow: 0 4px 16px rgba(255, 107, 107, 0.3);
      anchor-name: --fab;
    }

    .fab:hover {
      transform: scale(1.1);
      box-shadow: 0 6px 24px rgba(255, 107, 107, 0.4);
    }

    .fab-menu {
      position: absolute;
      display: flex;
      flex-direction: column;
      gap: 12px;
      opacity: 0;
      pointer-events: none;
      transition: opacity 0.3s ease;
      position-anchor: --fab;
      bottom: anchor(top);
      right: anchor(center);
      translate: 50% -12px;
    }

    .fab-menu.show {
      opacity: 1;
      pointer-events: auto;
    }

    .fab-option {
      width: 48px;
      height: 48px;
      background: white;
      border-radius: 50%;
      border: none;
      color: #333;
      font-size: 20px;
      cursor: pointer;
      transition: transform 0.2s ease, box-shadow 0.2s ease;
      box-shadow: 0 2px 8px rgba(0,0,0,0.2);
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .fab-option:hover {
      transform: scale(1.1);
      box-shadow: 0 4px 16px rgba(0,0,0,0.3);
    }

    .fab-option:nth-child(1) { background: #667eea; color: white; }
    .fab-option:nth-child(2) { background: #764ba2; color: white; }
    .fab-option:nth-child(3) { background: #4ecdc4; color: white; }

    .controls {
      background: rgba(255,255,255,0.95);
      border-radius: 12px;
      padding: 20px;
      margin-top: 30px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
    }

    .controls h3 {
      margin-bottom: 15px;
      color: #333;
    }

    .control-buttons {
      display: flex;
      gap: 15px;
      flex-wrap: wrap;
    }

    .control-btn {
      padding: 10px 20px;
      background: #667eea;
      color: white;
      border: none;
      border-radius: 6px;
      cursor: pointer;
      font-size: 14px;
      transition: background-color 0.2s ease;
    }

    .control-btn:hover {
      background: #5a67d8;
    }

    .control-btn.active {
      background: #4ecdc4;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>Anchor Positioning</h1>
      <p>Posicionamiento relativo a elementos ancla</p>
    </header>

    <div class="demo-grid">
      <!-- Tooltip Demo -->
      <div class="demo-card">
        <h3>Tooltip</h3>
        <div class="tooltip-demo">
          <button class="tooltip-button" id="tooltipBtn">Hover me</button>
          <div class="tooltip" id="tooltip">
            Este tooltip se posiciona automáticamente relativo al botón
          </div>
        </div>
      </div>

      <!-- Dropdown Demo -->
      <div class="demo-card">
        <h3>Menú Desplegable</h3>
        <div class="dropdown-demo">
          <button class="dropdown-button" id="dropdownBtn">Menú</button>
          <div class="dropdown-menu" id="dropdownMenu">
            <div class="dropdown-item">Opción 1</div>
            <div class="dropdown-item">Opción 2</div>
            <div class="dropdown-item">Opción 3</div>
            <div class="dropdown-item">Opción 4</div>
          </div>
        </div>
      </div>

      <!-- Popover Demo -->
      <div class="demo-card">
        <h3>Popover</h3>
        <div class="popover-demo">
          <button class="popover-trigger" id="popoverBtn">Mostrar Info</button>
          <div class="popover" id="popover">
            <h4>Información Importante</h4>
            <p>Este popover se posiciona automáticamente al lado del botón trigger usando anchor positioning.</p>
            <div class="popover-actions">
              <button class="popover-btn primary">Aceptar</button>
              <button class="popover-btn secondary">Cancelar</button>
            </div>
          </div>
        </div>
      </div>

      <!-- FAB Demo -->
      <div class="demo-card">
        <h3>FAB Menu</h3>
        <div class="fab-demo">
          <button class="fab" id="fabBtn">+</button>
          <div class="fab-menu" id="fabMenu">
            <button class="fab-option">📝</button>
            <button class="fab-option">📷</button>
            <button class="fab-option">🎵</button>
          </div>
        </div>
      </div>
    </div>

    <div class="controls">
      <h3>Controles de Demo</h3>
      <div class="control-buttons">
        <button class="control-btn active" onclick="toggleTooltip()">Toggle Tooltip</button>
        <button class="control-btn" onclick="toggleDropdown()">Toggle Dropdown</button>
        <button class="control-btn" onclick="togglePopover()">Toggle Popover</button>
        <button class="control-btn" onclick="toggleFAB()">Toggle FAB Menu</button>
        <button class="control-btn" onclick="resetAll()">Reset All</button>
      </div>
    </div>
  </div>

  <script>
    // Control functions
    function toggleTooltip() {
      const tooltip = document.getElementById('tooltip');
      const btn = document.getElementById('tooltipBtn');

      if (tooltip.style.opacity === '1') {
        tooltip.style.opacity = '0';
        btn.classList.remove('active');
      } else {
        tooltip.style.opacity = '1';
        btn.classList.add('active');
      }
    }

    function toggleDropdown() {
      const menu = document.getElementById('dropdownMenu');
      menu.classList.toggle('show');
    }

    function togglePopover() {
      const popover = document.getElementById('popover');
      popover.classList.toggle('show');
    }

    function toggleFAB() {
      const menu = document.getElementById('fabMenu');
      menu.classList.toggle('show');
    }

    function resetAll() {
      document.getElementById('tooltip').style.opacity = '0';
      document.getElementById('dropdownMenu').classList.remove('show');
      document.getElementById('popover').classList.remove('show');
      document.getElementById('fabMenu').classList.remove('show');

      document.querySelectorAll('.active').forEach(el => el.classList.remove('active'));
    }

    // Event listeners
    document.getElementById('tooltipBtn').addEventListener('mouseenter', () => {
      document.getElementById('tooltip').style.opacity = '1';
    });

    document.getElementById('tooltipBtn').addEventListener('mouseleave', () => {
      document.getElementById('tooltip').style.opacity = '0';
    });

    document.getElementById('dropdownBtn').addEventListener('click', toggleDropdown);

    document.getElementById('popoverBtn').addEventListener('click', togglePopover);

    document.getElementById('fabBtn').addEventListener('click', toggleFAB);

    // Close menus when clicking outside
    document.addEventListener('click', (e) => {
      if (!e.target.closest('.dropdown-demo')) {
        document.getElementById('dropdownMenu').classList.remove('show');
      }
      if (!e.target.closest('.popover-demo')) {
        document.getElementById('popover').classList.remove('show');
      }
      if (!e.target.closest('.fab-demo')) {
        document.getElementById('fabMenu').classList.remove('show');
      }
    });
  </script>
</body>
</html>
```

## Funciones de Anclaje Avanzadas

### Posicionamiento con inset-area

```css
/* Posicionamiento automático inteligente */
.tooltip {
  position: absolute;
  position-anchor: --my-anchor;

  /* inset-area elige automáticamente la mejor posición */
  inset-area: top; /* top, bottom, left, right */
  /* o combinaciones: top left, bottom right, etc. */
}

/* Con fallback manual */
.tooltip {
  position: absolute;
  position-anchor: --my-anchor;

  inset-area: top;

  /* Fallback si inset-area no es soportado */
  top: anchor(bottom);
  left: anchor(center);
  translate: -50% 8px;
}
```

### Anclas múltiples

```css
/* Un elemento puede tener múltiples anclas */
.element {
  position: absolute;
  position-anchor: --anchor-1, --anchor-2;

  /* Elegir ancla dinámicamente */
  top: anchor(--anchor-1 bottom);
  left: anchor(--anchor-2 center);
}
```

## Ejemplo Avanzado: Sistema de Tooltips Inteligente

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Anchor Positioning - Sistema Inteligente</title>
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
      padding: 40px;
      line-height: 1.6;
    }

    .container {
      max-width: 1400px;
      margin: 0 auto;
    }

    .header {
      text-align: center;
      margin-bottom: 60px;
    }

    .header h1 {
      font-size: 3.5rem;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #667eea, #764ba2);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
      margin-bottom: 20px;
    }

    .header p {
      font-size: 1.3rem;
      color: #a0aec0;
      max-width: 600px;
      margin: 0 auto;
    }

    .demo-layout {
      display: grid;
      grid-template-columns: 1fr 300px;
      gap: 40px;
      margin-bottom: 60px;
    }

    .content-area {
      background: rgba(255,255,255,0.05);
      border-radius: 16px;
      padding: 40px;
      border: 1px solid rgba(255,255,255,0.1);
    }

    .content-area h2 {
      margin-bottom: 30px;
      color: #e0e0e0;
      font-size: 2rem;
    }

    .content-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
      gap: 20px;
      margin-bottom: 40px;
    }

    .content-button {
      padding: 16px 24px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      color: white;
      border: none;
      border-radius: 12px;
      font-size: 16px;
      cursor: pointer;
      transition: all 0.3s ease;
      position: relative;
      anchor-name: var(--anchor-name);
    }

    .content-button:hover {
      transform: translateY(-3px);
      box-shadow: 0 8px 32px rgba(102, 126, 234, 0.3);
    }

    .content-button:nth-child(1) { --anchor-name: --btn-1; }
    .content-button:nth-child(2) { --anchor-name: --btn-2; }
    .content-button:nth-child(3) { --anchor-name: --btn-3; }
    .content-button:nth-child(4) { --anchor-name: --btn-4; }
    .content-button:nth-child(5) { --anchor-name: --btn-5; }
    .content-button:nth-child(6) { --anchor-name: --btn-6; }
    .content-button:nth-child(7) { --anchor-name: --btn-7; }
    .content-button:nth-child(8) { --anchor-name: --btn-8; }

    .smart-tooltip {
      position: absolute;
      background: #2d3748;
      color: white;
      padding: 16px 20px;
      border-radius: 8px;
      font-size: 14px;
      opacity: 0;
      pointer-events: none;
      transition: all 0.3s ease;
      box-shadow: 0 8px 32px rgba(0,0,0,0.3);
      max-width: 250px;
      z-index: 1000;
      position-anchor: var(--current-anchor);
      inset-area: top;
      margin: 8px;
    }

    .smart-tooltip.show {
      opacity: 1;
      pointer-events: auto;
    }

    .smart-tooltip::before {
      content: '';
      position: absolute;
      width: 0;
      height: 0;
      border: 6px solid transparent;
    }

    /* Posicionamiento dinámico basado en inset-area */
    .smart-tooltip[inset-area="top"]::before {
      bottom: -12px;
      left: 50%;
      transform: translateX(-50%);
      border-top: 6px solid #2d3748;
      border-bottom: none;
    }

    .smart-tooltip[inset-area="bottom"]::before {
      top: -12px;
      left: 50%;
      transform: translateX(-50%);
      border-bottom: 6px solid #2d3748;
      border-top: none;
    }

    .smart-tooltip[inset-area="left"]::before {
      right: -12px;
      top: 50%;
      transform: translateY(-50%);
      border-left: 6px solid #2d3748;
      border-right: none;
    }

    .smart-tooltip[inset-area="right"]::before {
      left: -12px;
      top: 50%;
      transform: translateY(-50%);
      border-right: 6px solid #2d3748;
      border-left: none;
    }

    .tooltip-content h4 {
      margin-bottom: 8px;
      color: #4ecdc4;
      font-size: 16px;
    }

    .tooltip-content p {
      margin-bottom: 12px;
      color: #a0aec0;
      font-size: 13px;
    }

    .tooltip-actions {
      display: flex;
      gap: 8px;
    }

    .tooltip-btn {
      padding: 6px 12px;
      border: none;
      border-radius: 4px;
      font-size: 12px;
      cursor: pointer;
      transition: background-color 0.2s ease;
    }

    .tooltip-btn.primary {
      background: #4ecdc4;
      color: white;
    }

    .tooltip-btn.secondary {
      background: rgba(255,255,255,0.1);
      color: #a0aec0;
      border: 1px solid rgba(255,255,255,0.2);
    }

    .controls-panel {
      background: rgba(255,255,255,0.05);
      border-radius: 16px;
      padding: 30px;
      border: 1px solid rgba(255,255,255,0.1);
      position: sticky;
      top: 40px;
      height: fit-content;
    }

    .controls-panel h3 {
      margin-bottom: 20px;
      color: #e0e0e0;
      font-size: 1.5rem;
    }

    .control-group {
      margin-bottom: 25px;
    }

    .control-group label {
      display: block;
      margin-bottom: 8px;
      color: #a0aec0;
      font-size: 14px;
      font-weight: 500;
    }

    .control-select {
      width: 100%;
      padding: 10px 12px;
      background: rgba(255,255,255,0.05);
      border: 1px solid rgba(255,255,255,0.2);
      border-radius: 6px;
      color: #e0e0e0;
      font-size: 14px;
    }

    .control-select:focus {
      outline: none;
      border-color: #4ecdc4;
    }

    .control-buttons {
      display: flex;
      flex-direction: column;
      gap: 10px;
    }

    .control-btn {
      padding: 12px 16px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      color: white;
      border: none;
      border-radius: 8px;
      cursor: pointer;
      font-size: 14px;
      transition: all 0.2s ease;
      text-align: left;
    }

    .control-btn:hover {
      transform: translateY(-1px);
      box-shadow: 0 4px 16px rgba(102, 126, 234, 0.3);
    }

    .control-btn.active {
      background: linear-gradient(45deg, #4ecdc4, #44a08d);
    }

    .stats-panel {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 20px;
      margin-top: 20px;
      border: 1px solid rgba(255,255,255,0.1);
    }

    .stats-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
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

    .floating-demo {
      margin-top: 60px;
      background: rgba(255,255,255,0.05);
      border-radius: 16px;
      padding: 40px;
      border: 1px solid rgba(255,255,255,0.1);
    }

    .floating-demo h2 {
      margin-bottom: 30px;
      color: #e0e0e0;
      font-size: 2rem;
    }

    .floating-container {
      position: relative;
      height: 400px;
      background: rgba(255,255,255,0.02);
      border-radius: 12px;
      border: 1px solid rgba(255,255,255,0.05);
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .floating-trigger {
      padding: 16px 32px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      color: white;
      border: none;
      border-radius: 12px;
      font-size: 18px;
      cursor: pointer;
      transition: all 0.3s ease;
      anchor-name: --floating-trigger;
    }

    .floating-trigger:hover {
      transform: scale(1.05);
      box-shadow: 0 8px 32px rgba(255, 107, 107, 0.3);
    }

    .floating-panel {
      position: absolute;
      background: rgba(255,255,255,0.95);
      backdrop-filter: blur(20px);
      border-radius: 16px;
      padding: 30px;
      box-shadow: 0 16px 64px rgba(0,0,0,0.3);
      max-width: 400px;
      opacity: 0;
      pointer-events: none;
      transition: all 0.4s ease;
      position-anchor: --floating-trigger;
      inset-area: top;
      margin: 16px;
      border: 1px solid rgba(255,255,255,0.2);
    }

    .floating-panel.show {
      opacity: 1;
      pointer-events: auto;
      transform: translateY(0);
    }

    .floating-panel h3 {
      margin-bottom: 16px;
      color: #333;
      font-size: 1.5rem;
    }

    .floating-panel p {
      color: #666;
      line-height: 1.6;
      margin-bottom: 20px;
    }

    .panel-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
      gap: 15px;
      margin-bottom: 20px;
    }

    .panel-item {
      text-align: center;
      padding: 12px;
      background: rgba(102, 126, 234, 0.1);
      border-radius: 8px;
      border: 1px solid rgba(102, 126, 234, 0.2);
    }

    .panel-item span {
      display: block;
      font-size: 20px;
      margin-bottom: 4px;
    }

    .panel-item small {
      color: #667eea;
      font-size: 12px;
    }

    .panel-actions {
      display: flex;
      gap: 12px;
      justify-content: flex-end;
    }

    .panel-btn {
      padding: 10px 20px;
      border: none;
      border-radius: 6px;
      font-size: 14px;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    .panel-btn.primary {
      background: #667eea;
      color: white;
    }

    .panel-btn.secondary {
      background: #f8f9fa;
      color: #333;
      border: 1px solid #dee2e6;
    }

    .panel-btn:hover {
      transform: translateY(-1px);
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>Sistema de Tooltips Inteligente</h1>
      <p>Anchor Positioning con posicionamiento automático inteligente</p>
    </header>

    <div class="demo-layout">
      <div class="content-area">
        <h2>Contenido Interactivo</h2>
        <div class="content-grid">
          <button class="content-button">Botón 1</button>
          <button class="content-button">Botón 2</button>
          <button class="content-button">Botón 3</button>
          <button class="content-button">Botón 4</button>
          <button class="content-button">Botón 5</button>
          <button class="content-button">Botón 6</button>
          <button class="content-button">Botón 7</button>
          <button class="content-button">Botón 8</button>
        </div>

        <p>Haz hover sobre los botones para ver los tooltips inteligentes que se posicionan automáticamente según el espacio disponible.</p>
      </div>

      <div class="controls-panel">
        <h3>Configuración</h3>

        <div class="control-group">
          <label for="positionMode">Modo de Posicionamiento:</label>
          <select class="control-select" id="positionMode">
            <option value="auto">Automático</option>
            <option value="top">Siempre Arriba</option>
            <option value="bottom">Siempre Abajo</option>
            <option value="left">Siempre Izquierda</option>
            <option value="right">Siempre Derecha</option>
          </select>
        </div>

        <div class="control-group">
          <label for="animationType">Tipo de Animación:</label>
          <select class="control-select" id="animationType">
            <option value="fade">Fade</option>
            <option value="slide">Slide</option>
            <option value="scale">Scale</option>
            <option value="bounce">Bounce</option>
          </select>
        </div>

        <div class="control-buttons">
          <button class="control-btn active" onclick="showAllTooltips()">Mostrar Todos</button>
          <button class="control-btn" onclick="hideAllTooltips()">Ocultar Todos</button>
          <button class="control-btn" onclick="randomizePositions()">Posiciones Aleatorias</button>
          <button class="control-btn" onclick="resetPositions()">Reset</button>
        </div>

        <div class="stats-panel">
          <div class="stats-grid">
            <div class="stat-item">
              <span class="stat-value" id="activeTooltips">0</span>
              <span class="stat-label">Tooltips Activos</span>
            </div>
            <div class="stat-item">
              <span class="stat-value" id="totalInteractions">0</span>
              <span class="stat-label">Interacciones</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="floating-demo">
      <h2>Panel Flotante Inteligente</h2>
      <div class="floating-container">
        <button class="floating-trigger" id="floatingTrigger">Mostrar Panel</button>

        <div class="floating-panel" id="floatingPanel">
          <h3>Panel de Información</h3>
          <p>Este panel se posiciona automáticamente usando Anchor Positioning, eligiendo la mejor ubicación según el espacio disponible en el viewport.</p>

          <div class="panel-grid">
            <div class="panel-item">
              <span>📊</span>
              <small>Analytics</small>
            </div>
            <div class="panel-item">
              <span>⚙️</span>
              <small>Settings</small>
            </div>
            <div class="panel-item">
              <span>👥</span>
              <small>Users</small>
            </div>
            <div class="panel-item">
              <span>📈</span>
              <small>Reports</small>
            </div>
          </div>

          <div class="panel-actions">
            <button class="panel-btn secondary">Cancelar</button>
            <button class="panel-btn primary">Aceptar</button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Tooltip template -->
  <div class="smart-tooltip" id="tooltipTemplate">
    <div class="tooltip-content">
      <h4 id="tooltipTitle">Título del Tooltip</h4>
      <p id="tooltipText">Contenido del tooltip que se muestra al hacer hover sobre el elemento.</p>
      <div class="tooltip-actions">
        <button class="tooltip-btn primary" onclick="tooltipAction('primary')">Acción</button>
        <button class="tooltip-btn secondary" onclick="tooltipAction('secondary')">Más Info</button>
      </div>
    </div>
  </div>

  <script>
    class SmartTooltipSystem {
      constructor() {
        this.activeTooltips = new Set();
        this.totalInteractions = 0;
        this.currentTooltip = null;
        this.positionMode = 'auto';
        this.animationType = 'fade';

        this.init();
      }

      init() {
        // Crear tooltips para cada botón
        document.querySelectorAll('.content-button').forEach((button, index) => {
          this.createTooltipForButton(button, index);
        });

        // Configurar controles
        this.setupControls();

        // Configurar panel flotante
        this.setupFloatingPanel();

        // Actualizar stats iniciales
        this.updateStats();
      }

      createTooltipForButton(button, index) {
        const tooltip = document.getElementById('tooltipTemplate').cloneNode(true);
        tooltip.id = `tooltip-${index}`;
        tooltip.style.positionAnchor = `var(--btn-${index + 1})`;

        // Personalizar contenido
        const titles = ['Información', 'Configuración', 'Ayuda', 'Estadísticas', 'Perfil', 'Notificaciones', 'Ajustes', 'Soporte'];
        const contents = [
          'Información detallada sobre esta funcionalidad.',
          'Configura las opciones según tus necesidades.',
          'Obtén ayuda y soporte técnico.',
          'Visualiza estadísticas y métricas.',
          'Gestiona tu perfil de usuario.',
          'Revisa tus notificaciones pendientes.',
          'Ajusta la configuración del sistema.',
          'Contacta con nuestro equipo de soporte.'
        ];

        tooltip.querySelector('#tooltipTitle').textContent = titles[index];
        tooltip.querySelector('#tooltipText').textContent = contents[index];

        document.body.appendChild(tooltip);

        // Event listeners
        button.addEventListener('mouseenter', () => this.showTooltip(tooltip, button));
        button.addEventListener('mouseleave', () => this.hideTooltip(tooltip));
      }

      showTooltip(tooltip, button) {
        if (this.currentTooltip) {
          this.hideTooltip(this.currentTooltip);
        }

        // Aplicar modo de posicionamiento
        this.applyPositioningMode(tooltip);

        // Aplicar animación
        this.applyAnimation(tooltip);

        tooltip.classList.add('show');
        this.activeTooltips.add(tooltip);
        this.currentTooltip = tooltip;
        this.totalInteractions++;

        this.updateStats();
      }

      hideTooltip(tooltip) {
        tooltip.classList.remove('show');
        this.activeTooltips.delete(tooltip);
        if (this.currentTooltip === tooltip) {
          this.currentTooltip = null;
        }
        this.updateStats();
      }

      applyPositioningMode(tooltip) {
        switch (this.positionMode) {
          case 'auto':
            tooltip.style.insetArea = 'top';
            break;
          case 'top':
            tooltip.style.insetArea = 'top';
            break;
          case 'bottom':
            tooltip.style.insetArea = 'bottom';
            break;
          case 'left':
            tooltip.style.insetArea = 'left';
            break;
          case 'right':
            tooltip.style.insetArea = 'right';
            break;
        }
      }

      applyAnimation(tooltip) {
        // Resetear animaciones previas
        tooltip.style.animation = '';

        // Aplicar nueva animación
        setTimeout(() => {
          switch (this.animationType) {
            case 'fade':
              tooltip.style.animation = 'fadeIn 0.3s ease';
              break;
            case 'slide':
              tooltip.style.animation = 'slideIn 0.3s ease';
              break;
            case 'scale':
              tooltip.style.animation = 'scaleIn 0.3s ease';
              break;
            case 'bounce':
              tooltip.style.animation = 'bounceIn 0.4s ease';
              break;
          }
        }, 10);
      }

      setupControls() {
        document.getElementById('positionMode').addEventListener('change', (e) => {
          this.positionMode = e.target.value;
          if (this.currentTooltip) {
            this.applyPositioningMode(this.currentTooltip);
          }
        });

        document.getElementById('animationType').addEventListener('change', (e) => {
          this.animationType = e.target.value;
        });
      }

      setupFloatingPanel() {
        const trigger = document.getElementById('floatingTrigger');
        const panel = document.getElementById('floatingPanel');

        trigger.addEventListener('click', () => {
          panel.classList.toggle('show');
        });

        // Cerrar al hacer click fuera
        document.addEventListener('click', (e) => {
          if (!trigger.contains(e.target) && !panel.contains(e.target)) {
            panel.classList.remove('show');
          }
        });
      }

      updateStats() {
        document.getElementById('activeTooltips').textContent = this.activeTooltips.size;
        document.getElementById('totalInteractions').textContent = this.totalInteractions;
      }
    }

    // Funciones globales para controles
    function showAllTooltips() {
      document.querySelectorAll('.smart-tooltip').forEach(tooltip => {
        tooltip.classList.add('show');
      });
    }

    function hideAllTooltips() {
      document.querySelectorAll('.smart-tooltip').forEach(tooltip => {
        tooltip.classList.remove('show');
      });
    }

    function randomizePositions() {
      const modes = ['auto', 'top', 'bottom', 'left', 'right'];
      document.getElementById('positionMode').value = modes[Math.floor(Math.random() * modes.length)];
      document.getElementById('positionMode').dispatchEvent(new Event('change'));
    }

    function resetPositions() {
      document.getElementById('positionMode').value = 'auto';
      document.getElementById('animationType').value = 'fade';
      document.getElementById('positionMode').dispatchEvent(new Event('change'));
      hideAllTooltips();
    }

    function tooltipAction(type) {
      alert(`Acción ${type} ejecutada`);
    }

    // Agregar estilos de animación
    const style = document.createElement('style');
    style.textContent = `
      @keyframes fadeIn {
        from { opacity: 0; }
        to { opacity: 1; }
      }

      @keyframes slideIn {
        from { opacity: 0; transform: translateY(-10px); }
        to { opacity: 1; transform: translateY(0); }
      }

      @keyframes scaleIn {
        from { opacity: 0; transform: scale(0.9); }
        to { opacity: 1; transform: scale(1); }
      }

      @keyframes bounceIn {
        0% { opacity: 0; transform: scale(0.3); }
        50% { opacity: 1; transform: scale(1.05); }
        70% { transform: scale(0.9); }
        100% { opacity: 1; transform: scale(1); }
      }
    `;
    document.head.appendChild(style);

    // Inicializar sistema
    new SmartTooltipSystem();
  </script>
</body>
</html>
```

## Técnicas Avanzadas

### Anclas condicionales

```css
/* Posicionamiento condicional basado en espacio disponible */
.tooltip {
  position: absolute;
  position-anchor: --my-anchor;

  /* Intentar posicionar arriba primero */
  inset-area: top;

  /* Si no hay espacio arriba, usar abajo */
  inset-area: bottom;
}

/* Con JavaScript para lógica más compleja */
.tooltip {
  --preferred-position: top;
  --fallback-position: bottom;
}

.tooltip[data-position="top"] {
  inset-area: top;
}

.tooltip[data-position="bottom"] {
  inset-area: bottom;
}
```

### Anclas dinámicas

```javascript
// Cambiar ancla dinámicamente
function updateTooltipAnchor(tooltip, newAnchor) {
  tooltip.style.setProperty('--current-anchor', newAnchor);
  tooltip.style.positionAnchor = 'var(--current-anchor)';
}
```

## Compatibilidad y Fallbacks

### Detectar soporte

```javascript
// Detectar soporte de Anchor Positioning
function supportsAnchorPositioning() {
  return CSS.supports('position-anchor: --test');
}

// Fallback para navegadores sin soporte
if (!supportsAnchorPositioning()) {
  // Usar posicionamiento tradicional con JavaScript
  function positionTooltip(tooltip, anchor) {
    const anchorRect = anchor.getBoundingClientRect();
    const tooltipRect = tooltip.getBoundingClientRect();
    const viewport = {
      width: window.innerWidth,
      height: window.innerHeight
    };

    // Lógica de posicionamiento manual
    let top = anchorRect.top - tooltipRect.height - 10;
    let left = anchorRect.left + (anchorRect.width / 2) - (tooltipRect.width / 2);

    // Ajustar si se sale del viewport
    if (top < 0) {
      top = anchorRect.bottom + 10;
    }

    if (left < 0) {
      left = 10;
    } else if (left + tooltipRect.width > viewport.width) {
      left = viewport.width - tooltipRect.width - 10;
    }

    tooltip.style.top = top + 'px';
    tooltip.style.left = left + 'px';
  }
}
```

## Casos de Uso Prácticos

### 1. Tooltips inteligentes

```css
.smart-tooltip {
  position: absolute;
  position-anchor: --tooltip-anchor;
  inset-area: top;
  margin: 8px;
}
```

### 2. Menús contextuales

```css
.context-menu {
  position: absolute;
  position-anchor: --context-anchor;
  inset-area: bottom left;
}
```

### 3. Popovers responsive

```css
.popover {
  position: absolute;
  position-anchor: --popover-anchor;
  inset-area: right;
}

@media (max-width: 768px) {
  .popover {
    inset-area: bottom;
  }
}
```

### 4. Elementos flotantes

```css
.floating-element {
  position: absolute;
  position-anchor: --float-anchor;
  top: anchor(center);
  left: anchor(end);
  translate: 16px -50%;
}
```

## Performance Considerations

### Optimizaciones

```css
/* Evitar repaints innecesarios */
.tooltip {
  will-change: transform;
  contain: layout style paint;
}

/* Usar transform en lugar de top/left cuando sea posible */
.tooltip {
  position: absolute;
  position-anchor: --anchor;
  inset-area: top;
  translate: 0 8px; /* Mejor que margin-top */
}
```

## Resumen

Anchor Positioning revoluciona el posicionamiento web:

- ✅ **Declarativo**: Posicionamiento sin JavaScript complejo
- ✅ **Automático**: `inset-area` elige la mejor posición
- ✅ **Flexible**: Múltiples anclas y funciones de posicionamiento
- ✅ **Responsive**: Se adapta automáticamente al viewport
- ✅ **Performante**: Optimizado para animaciones y transiciones

Anchor Positioning simplifica tareas que antes requerían JavaScript complejo, permitiendo crear interfaces más robustas y mantenibles con menos código.