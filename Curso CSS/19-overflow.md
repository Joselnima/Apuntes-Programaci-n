# Módulo 19 - Overflow

En este módulo aprenderás sobre CSS Overflow, una propiedad fundamental que controla cómo se comporta el contenido que excede los límites de su contenedor, incluyendo técnicas avanzadas como scrollbars personalizados y overflow-x/y.

## Introducción a CSS Overflow

### ¿Qué es CSS Overflow?

CSS Overflow controla qué sucede cuando el contenido de un elemento es demasiado grande para caber en su área designada. Puede ser visible, oculto, o hacer que aparezca una barra de desplazamiento.

```css
/* Contenido visible fuera del contenedor */
.element {
  overflow: visible;
}

/* Contenido oculto */
.element {
  overflow: hidden;
}

/* Barra de desplazamiento automática */
.element {
  overflow: auto;
}

/* Siempre mostrar barra de desplazamiento */
.element {
  overflow: scroll;
}
```

## Propiedades Básicas de Overflow

### Overflow

```css
.container {
  width: 300px;
  height: 200px;
  border: 2px solid #333;
  overflow: auto; /* visible | hidden | scroll | auto */
}
```

### Overflow-x y Overflow-y

```css
/* Control independiente de ejes */
.element {
  overflow-x: auto;  /* scroll horizontal */
  overflow-y: hidden; /* ocultar vertical */
}
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Overflow - Básico</title>
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

    .demo-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
      gap: 30px;
      margin-bottom: 50px;
    }

    .demo-item {
      background: white;
      border-radius: 12px;
      padding: 25px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
      border: 1px solid #e0e0e0;
    }

    .demo-item h3 {
      margin-bottom: 15px;
      color: #333;
      font-size: 1.3rem;
    }

    .demo-item p {
      margin-bottom: 20px;
      color: #666;
      line-height: 1.6;
    }

    .overflow-container {
      width: 100%;
      height: 150px;
      border: 2px solid #ddd;
      border-radius: 8px;
      padding: 15px;
      background: #fafafa;
      font-family: monospace;
      font-size: 14px;
      line-height: 1.5;
      margin-bottom: 15px;
    }

    /* Diferentes tipos de overflow */
    .visible-overflow {
      overflow: visible;
    }

    .hidden-overflow {
      overflow: hidden;
    }

    .scroll-overflow {
      overflow: scroll;
    }

    .auto-overflow {
      overflow: auto;
    }

    .x-scroll {
      overflow-x: auto;
      overflow-y: hidden;
      white-space: nowrap;
    }

    .y-scroll {
      overflow-x: hidden;
      overflow-y: auto;
    }

    .content {
      background: linear-gradient(45deg, #667eea 0%, #764ba2 100%);
      color: white;
      padding: 20px;
      border-radius: 8px;
      margin-bottom: 10px;
    }

    .long-content {
      width: 200%;
      background: repeating-linear-gradient(
        45deg,
        #ff6b6b,
        #ff6b6b 20px,
        #4ecdc4 20px,
        #4ecdc4 40px
      );
      height: 100px;
      border-radius: 8px;
      margin-bottom: 10px;
    }

    .tall-content {
      height: 300px;
      background: repeating-linear-gradient(
        to bottom,
        #45b7d1,
        #45b7d1 30px,
        #96ceb4 30px,
        #96ceb4 60px
      );
      border-radius: 8px;
    }

    .code-example {
      background: #2d3748;
      color: #e2e8f0;
      padding: 15px;
      border-radius: 6px;
      font-family: 'Monaco', 'Menlo', monospace;
      font-size: 13px;
      margin-top: 15px;
    }

    .comparison-section {
      background: white;
      border-radius: 12px;
      padding: 30px;
      margin-top: 50px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
    }

    .comparison-section h2 {
      text-align: center;
      margin-bottom: 30px;
      color: #333;
      font-size: 2rem;
    }

    .comparison-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
      gap: 20px;
    }

    .comparison-item {
      border: 1px solid #e0e0e0;
      border-radius: 8px;
      overflow: hidden;
    }

    .comparison-header {
      background: #667eea;
      color: white;
      padding: 15px;
      text-align: center;
      font-weight: bold;
    }

    .comparison-content {
      padding: 20px;
      height: 200px;
    }

    .comparison-content p {
      margin-bottom: 15px;
      font-size: 14px;
    }

    .visible-demo {
      overflow: visible;
      position: relative;
    }

    .hidden-demo {
      overflow: hidden;
    }

    .scroll-demo {
      overflow: scroll;
    }

    .auto-demo {
      overflow: auto;
    }

    .demo-box {
      width: 120px;
      height: 80px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      border-radius: 8px;
      margin: 10px;
      display: inline-block;
      color: white;
      text-align: center;
      line-height: 80px;
      font-weight: bold;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>CSS Overflow</h1>
      <p>Controlando el contenido que excede los límites del contenedor</p>
    </header>

    <div class="demo-grid">
      <div class="demo-item">
        <h3>Overflow: Visible</h3>
        <p>El contenido se muestra fuera del contenedor. Este es el valor por defecto.</p>
        <div class="overflow-container visible-overflow">
          <div class="content">Contenido normal</div>
          <div class="demo-box">1</div>
          <div class="demo-box">2</div>
          <div class="demo-box">3</div>
          <div class="demo-box">4</div>
        </div>
        <div class="code-example">
overflow: visible;
        </div>
      </div>

      <div class="demo-item">
        <h3>Overflow: Hidden</h3>
        <p>El contenido que excede se oculta completamente.</p>
        <div class="overflow-container hidden-overflow">
          <div class="content">Contenido normal</div>
          <div class="demo-box">1</div>
          <div class="demo-box">2</div>
          <div class="demo-box">3</div>
          <div class="demo-box">4</div>
        </div>
        <div class="code-example">
overflow: hidden;
        </div>
      </div>

      <div class="demo-item">
        <h3>Overflow: Scroll</h3>
        <p>Siempre muestra barras de desplazamiento, incluso si no son necesarias.</p>
        <div class="overflow-container scroll-overflow">
          <div class="content">Contenido normal</div>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
        </div>
        <div class="code-example">
overflow: scroll;
        </div>
      </div>

      <div class="demo-item">
        <h3>Overflow: Auto</h3>
        <p>Muestra barras de desplazamiento solo cuando son necesarias.</p>
        <div class="overflow-container auto-overflow">
          <div class="content">Contenido normal</div>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
        </div>
        <div class="code-example">
overflow: auto;
        </div>
      </div>

      <div class="demo-item">
        <h3>Overflow-x: Auto</h3>
        <p>Control específico del desplazamiento horizontal.</p>
        <div class="overflow-container x-scroll">
          <div class="long-content"></div>
        </div>
        <div class="code-example">
overflow-x: auto;<br>
overflow-y: hidden;
        </div>
      </div>

      <div class="demo-item">
        <h3>Overflow-y: Auto</h3>
        <p>Control específico del desplazamiento vertical.</p>
        <div class="overflow-container y-scroll">
          <div class="tall-content"></div>
        </div>
        <div class="code-example">
overflow-x: hidden;<br>
overflow-y: auto;
        </div>
      </div>
    </div>

    <div class="comparison-section">
      <h2>Comparación Visual</h2>

      <div class="comparison-grid">
        <div class="comparison-item">
          <div class="comparison-header">Visible</div>
          <div class="comparison-content visible-demo">
            <p>Contenido que puede extenderse fuera del contenedor.</p>
            <div class="demo-box" style="position: relative; left: 50px;">Box</div>
          </div>
        </div>

        <div class="comparison-item">
          <div class="comparison-header">Hidden</div>
          <div class="comparison-content hidden-demo">
            <p>Contenido cortado en los bordes.</p>
            <div class="demo-box" style="position: relative; left: 50px;">Box</div>
          </div>
        </div>

        <div class="comparison-item">
          <div class="comparison-header">Scroll</div>
          <div class="comparison-content scroll-demo">
            <p>Barras de desplazamiento siempre visibles.</p>
            <p>Texto adicional para demostrar el scroll.</p>
          </div>
        </div>

        <div class="comparison-item">
          <div class="comparison-header">Auto</div>
          <div class="comparison-content auto-demo">
            <p>Barras solo cuando son necesarias.</p>
            <p>Este contenido es lo suficientemente largo para mostrar una barra de desplazamiento vertical.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
```

## Scrollbars Personalizados

### CSS Scrollbars (Webkit)

```css
/* Personalizar scrollbar en Webkit */
.custom-scrollbar::-webkit-scrollbar {
  width: 12px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 6px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: linear-gradient(45deg, #667eea, #764ba2);
  border-radius: 6px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(45deg, #5a67d8, #6b46c1);
}
```

### Firefox Scrollbars

```css
/* Personalizar scrollbar en Firefox */
.custom-scrollbar {
  scrollbar-width: thin;
  scrollbar-color: #667eea #f1f1f1;
}
```

## Ejemplo con Scrollbars Personalizados

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Overflow - Scrollbars</title>
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
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    .scrollbar-showcase {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 30px;
      margin-bottom: 50px;
    }

    .scrollbar-demo {
      background: rgba(255,255,255,0.05);
      border-radius: 15px;
      padding: 25px;
      border: 1px solid rgba(255,255,255,0.1);
    }

    .scrollbar-demo h3 {
      margin-bottom: 20px;
      color: #4ecdc4;
      font-size: 1.3rem;
    }

    .scroll-content {
      height: 200px;
      padding: 15px;
      background: rgba(255,255,255,0.02);
      border-radius: 8px;
      border: 1px solid rgba(255,255,255,0.05);
      overflow-y: auto;
      font-size: 14px;
      line-height: 1.6;
    }

    /* Scrollbar personalizado moderno */
    .modern-scroll::-webkit-scrollbar {
      width: 8px;
    }

    .modern-scroll::-webkit-scrollbar-track {
      background: rgba(255,255,255,0.1);
      border-radius: 4px;
    }

    .modern-scroll::-webkit-scrollbar-thumb {
      background: linear-gradient(45deg, #667eea, #764ba2);
      border-radius: 4px;
      transition: background 0.3s ease;
    }

    .modern-scroll::-webkit-scrollbar-thumb:hover {
      background: linear-gradient(45deg, #5a67d8, #6b46c1);
    }

    /* Scrollbar minimalista */
    .minimal-scroll::-webkit-scrollbar {
      width: 6px;
    }

    .minimal-scroll::-webkit-scrollbar-track {
      background: transparent;
    }

    .minimal-scroll::-webkit-scrollbar-thumb {
      background: rgba(255,255,255,0.3);
      border-radius: 3px;
    }

    .minimal-scroll::-webkit-scrollbar-thumb:hover {
      background: rgba(255,255,255,0.5);
    }

    /* Scrollbar colorido */
    .colorful-scroll::-webkit-scrollbar {
      width: 12px;
    }

    .colorful-scroll::-webkit-scrollbar-track {
      background: #2d3748;
      border-radius: 6px;
    }

    .colorful-scroll::-webkit-scrollbar-thumb {
      background: linear-gradient(to bottom, #ff6b6b, #4ecdc4, #45b7d1);
      border-radius: 6px;
      border: 2px solid #2d3748;
    }

    /* Scrollbar con patrón */
    .pattern-scroll::-webkit-scrollbar {
      width: 16px;
    }

    .pattern-scroll::-webkit-scrollbar-track {
      background: #1a202c;
      border-radius: 8px;
    }

    .pattern-scroll::-webkit-scrollbar-thumb {
      background: repeating-linear-gradient(
        45deg,
        #667eea,
        #667eea 4px,
        #764ba2 4px,
        #764ba2 8px
      );
      border-radius: 8px;
      border: 2px solid #1a202c;
    }

    /* Firefox scrollbar */
    .firefox-scroll {
      scrollbar-width: thin;
      scrollbar-color: #667eea #2d3748;
    }

    .code-display {
      background: #2d3748;
      border-radius: 8px;
      padding: 20px;
      margin-top: 15px;
      font-family: 'Monaco', 'Menlo', monospace;
      font-size: 13px;
      color: #e2e8f0;
      border: 1px solid #4a5568;
    }

    .interactive-demo {
      background: rgba(255,255,255,0.05);
      border-radius: 15px;
      padding: 30px;
      margin-top: 50px;
      border: 1px solid rgba(255,255,255,0.1);
    }

    .interactive-demo h2 {
      text-align: center;
      margin-bottom: 30px;
      font-size: 2rem;
      color: #4ecdc4;
    }

    .controls {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
      gap: 20px;
      margin-bottom: 30px;
    }

    .control-group {
      display: flex;
      flex-direction: column;
    }

    .control-group label {
      margin-bottom: 8px;
      font-weight: bold;
      color: #e0e0e0;
    }

    .control-group input,
    .control-group select {
      padding: 10px;
      border: 1px solid rgba(255,255,255,0.2);
      border-radius: 6px;
      background: rgba(255,255,255,0.05);
      color: #e0e0e0;
      font-size: 14px;
    }

    .control-group input:focus,
    .control-group select:focus {
      outline: none;
      border-color: #4ecdc4;
      box-shadow: 0 0 0 2px rgba(78, 205, 196, 0.3);
    }

    .preview-container {
      height: 250px;
      background: rgba(255,255,255,0.02);
      border-radius: 10px;
      border: 1px solid rgba(255,255,255,0.05);
      padding: 20px;
      overflow-y: auto;
      position: relative;
    }

    .preview-content {
      color: #e0e0e0;
      line-height: 1.6;
    }

    .custom-scroll {
      scrollbar-width: thin;
      scrollbar-color: #667eea #2d3748;
    }

    .custom-scroll::-webkit-scrollbar {
      width: 8px;
    }

    .custom-scroll::-webkit-scrollbar-track {
      background: #2d3748;
      border-radius: 4px;
    }

    .custom-scroll::-webkit-scrollbar-thumb {
      background: #667eea;
      border-radius: 4px;
      transition: background 0.3s ease;
    }

    .custom-scroll::-webkit-scrollbar-thumb:hover {
      background: #5a67d8;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>Scrollbars Personalizados</h1>
      <p>Diseña barras de desplazamiento únicas y modernas</p>
    </header>

    <div class="scrollbar-showcase">
      <div class="scrollbar-demo">
        <h3>Scrollbar Moderno</h3>
        <div class="scroll-content modern-scroll">
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
          <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
          <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>
        </div>
        <div class="code-display">
::-webkit-scrollbar { width: 8px; }<br>
::-webkit-scrollbar-track { background: rgba(255,255,255,0.1); }<br>
::-webkit-scrollbar-thumb { background: linear-gradient(45deg, #667eea, #764ba2); }
        </div>
      </div>

      <div class="scrollbar-demo">
        <h3>Scrollbar Minimalista</h3>
        <div class="scroll-content minimal-scroll">
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
          <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
          <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>
        </div>
        <div class="code-display">
::-webkit-scrollbar { width: 6px; }<br>
::-webkit-scrollbar-thumb { background: rgba(255,255,255,0.3); }
        </div>
      </div>

      <div class="scrollbar-demo">
        <h3>Scrollbar Colorido</h3>
        <div class="scroll-content colorful-scroll">
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
          <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
          <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>
        </div>
        <div class="code-display">
::-webkit-scrollbar-thumb { background: linear-gradient(to bottom, #ff6b6b, #4ecdc4, #45b7d1); }
        </div>
      </div>

      <div class="scrollbar-demo">
        <h3>Scrollbar con Patrón</h3>
        <div class="scroll-content pattern-scroll">
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
          <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
          <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>
        </div>
        <div class="code-display">
::-webkit-scrollbar-thumb { background: repeating-linear-gradient(45deg, #667eea, #667eea 4px, #764ba2 4px, #764ba2 8px); }
        </div>
      </div>
    </div>

    <div class="interactive-demo">
      <h2>Demo Interactivo</h2>

      <div class="controls">
        <div class="control-group">
          <label for="scrollbarWidth">Ancho (px):</label>
          <input type="range" id="scrollbarWidth" min="4" max="20" value="8">
        </div>

        <div class="control-group">
          <label for="trackColor">Color del Track:</label>
          <input type="color" id="trackColor" value="#2d3748">
        </div>

        <div class="control-group">
          <label for="thumbColor">Color del Thumb:</label>
          <input type="color" id="thumbColor" value="#667eea">
        </div>

        <div class="control-group">
          <label for="borderRadius">Radio del Borde (px):</label>
          <input type="range" id="borderRadius" min="0" max="10" value="4">
        </div>
      </div>

      <div class="preview-container custom-scroll" id="previewContainer">
        <div class="preview-content">
          <h3>Contenido de Ejemplo</h3>
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

          <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>

          <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>

          <p>Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.</p>

          <p>Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.</p>

          <p>Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur?</p>
        </div>
      </div>
    </div>
  </div>

  <script>
    const scrollbarWidth = document.getElementById('scrollbarWidth');
    const trackColor = document.getElementById('trackColor');
    const thumbColor = document.getElementById('thumbColor');
    const borderRadius = document.getElementById('borderRadius');
    const previewContainer = document.getElementById('previewContainer');

    function updateScrollbar() {
      const width = scrollbarWidth.value;
      const track = trackColor.value;
      const thumb = thumbColor.value;
      const radius = borderRadius.value;

      const style = document.getElementById('dynamic-styles') || document.createElement('style');
      style.id = 'dynamic-styles';
      style.textContent = `
        .custom-scroll::-webkit-scrollbar {
          width: ${width}px;
        }
        .custom-scroll::-webkit-scrollbar-track {
          background: ${track};
          border-radius: ${radius}px;
        }
        .custom-scroll::-webkit-scrollbar-thumb {
          background: ${thumb};
          border-radius: ${radius}px;
        }
        .custom-scroll::-webkit-scrollbar-thumb:hover {
          background: ${thumb}dd;
        }
      `;

      if (!document.head.contains(style)) {
        document.head.appendChild(style);
      }
    }

    scrollbarWidth.addEventListener('input', updateScrollbar);
    trackColor.addEventListener('input', updateScrollbar);
    thumbColor.addEventListener('input', updateScrollbar);
    borderRadius.addEventListener('input', updateScrollbar);

    // Inicializar
    updateScrollbar();
  </script>
</body>
</html>
```

## Overflow en Contextos Avanzados

### Text Overflow

```css
/* Texto que se corta con elipsis */
.text-truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* Múltiples líneas */
.text-multiline {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
```

### Scroll Snap

```css
/* Scroll con puntos de anclaje */
.scroll-snap-container {
  overflow-x: auto;
  scroll-snap-type: x mandatory;
}

.scroll-snap-item {
  scroll-snap-align: start;
  flex: 0 0 300px;
}
```

## Casos de Uso Prácticos

### 1. Contenedores responsivos

```css
.responsive-container {
  max-height: 400px;
  overflow-y: auto;
  overflow-x: hidden;
}
```

### 2. Galerías de imágenes

```css
.image-gallery {
  display: flex;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  gap: 10px;
}

.gallery-image {
  flex: 0 0 300px;
  scroll-snap-align: start;
}
```

### 3. Chat interfaces

```css
.chat-messages {
  height: 300px;
  overflow-y: auto;
  scroll-behavior: smooth;
}
```

## Performance y Mejores Prácticas

### Optimizaciones

```css
/* Usar transform3d para forzar aceleración por hardware */
.scroll-container {
  transform: translate3d(0, 0, 0);
  overflow: auto;
}

/* Evitar overflow en elementos con animaciones complejas */
.static-content {
  overflow: hidden;
  /* Mejor performance */
}
```

## Compatibilidad

### Fallbacks

```css
.element {
  /* Fallback básico */
  overflow: auto;

  /* Scrollbars personalizados para Webkit */
  &::-webkit-scrollbar {
    width: 8px;
  }

  /* Firefox */
  scrollbar-width: thin;
  scrollbar-color: #667eea #f1f1f1;
}
```

## Resumen

CSS Overflow es fundamental para controlar el contenido excedente:

- ✅ **Overflow básico**: visible, hidden, scroll, auto
- ✅ **Control por eje**: overflow-x y overflow-y independientes
- ✅ **Scrollbars personalizados**: Diseño único para barras de desplazamiento
- ✅ **Text overflow**: Elipsis y truncado de texto
- ✅ **Scroll snap**: Navegación con puntos de anclaje
- ✅ **Performance**: Optimizaciones para scroll suave

El overflow es esencial para crear interfaces scrollables, galerías, chats, y cualquier contenido que necesite control sobre cómo se maneja el espacio limitado.