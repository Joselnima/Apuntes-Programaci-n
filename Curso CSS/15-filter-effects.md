# Módulo 15 - Filter Effects

En este módulo aprenderás sobre CSS Filter Effects, una poderosa herramienta que permite aplicar efectos visuales avanzados a elementos HTML como desenfoque, brillo, contraste, saturación y muchos más.

## Introducción a CSS Filters

### ¿Qué son los CSS Filters?

Los filtros CSS permiten aplicar efectos gráficos a elementos sin necesidad de software de edición de imágenes. Son procesados por el navegador y pueden ser animados.

```css
/* Aplicando un filtro de desenfoque */
.element {
  filter: blur(5px);
}

/* Múltiples filtros */
.element {
  filter: blur(5px) brightness(1.2) contrast(1.5);
}
```

## Filtros Básicos

### Blur (Desenfoque)

```css
.blur-light {
  filter: blur(2px);
}

.blur-medium {
  filter: blur(5px);
}

.blur-heavy {
  filter: blur(10px);
}
```

### Brightness (Brillo)

```css
.bright-darker {
  filter: brightness(0.5);
}

.bright-normal {
  filter: brightness(1);
}

.bright-lighter {
  filter: brightness(1.5);
}
```

### Contrast (Contraste)

```css
.contrast-low {
  filter: contrast(0.5);
}

.contrast-normal {
  filter: contrast(1);
}

.contrast-high {
  filter: contrast(2);
}
```

### Grayscale (Escala de grises)

```css
.grayscale-full {
  filter: grayscale(1);
}

.grayscale-half {
  filter: grayscale(0.5);
}

.grayscale-none {
  filter: grayscale(0);
}
```

### Sepia

```css
.sepia-light {
  filter: sepia(0.3);
}

.sepia-medium {
  filter: sepia(0.7);
}

.sepia-full {
  filter: sepia(1);
}
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Filters - Básico</title>
  <style>
    .container {
      max-width: 1200px;
      margin: 0 auto;
      padding: 20px;
    }

    .filter-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
      gap: 20px;
      margin-bottom: 40px;
    }

    .filter-item {
      position: relative;
      overflow: hidden;
      border-radius: 10px;
      box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    }

    .filter-item img {
      width: 100%;
      height: 200px;
      object-fit: cover;
      transition: filter 0.3s ease;
    }

    .filter-item:hover img {
      filter: none;
    }

    .filter-overlay {
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
      background: linear-gradient(transparent, rgba(0,0,0,0.7));
      color: white;
      padding: 20px;
      transform: translateY(100%);
      transition: transform 0.3s ease;
    }

    .filter-item:hover .filter-overlay {
      transform: translateY(0);
    }

    /* Filtros aplicados */
    .blur img { filter: blur(3px); }
    .brightness img { filter: brightness(1.3); }
    .contrast img { filter: contrast(1.5); }
    .grayscale img { filter: grayscale(1); }
    .sepia img { filter: sepia(0.8); }
  </style>
</head>
<body>
  <div class="container">
    <h1>CSS Filter Effects - Básicos</h1>

    <div class="filter-grid">
      <div class="filter-item blur">
        <img src="https://picsum.photos/300/200?random=1" alt="Blur">
        <div class="filter-overlay">
          <h3>Blur (3px)</h3>
          <p>Efecto de desenfoque suave</p>
        </div>
      </div>

      <div class="filter-item brightness">
        <img src="https://picsum.photos/300/200?random=2" alt="Brightness">
        <div class="filter-overlay">
          <h3>Brightness (1.3)</h3>
          <p>Aumento del brillo</p>
        </div>
      </div>

      <div class="filter-item contrast">
        <img src="https://picsum.photos/300/200?random=3" alt="Contrast">
        <div class="filter-overlay">
          <h3>Contrast (1.5)</h3>
          <p>Mayor contraste</p>
        </div>
      </div>

      <div class="filter-item grayscale">
        <img src="https://picsum.photos/300/200?random=4" alt="Grayscale">
        <div class="filter-overlay">
          <h3>Grayscale (100%)</h3>
          <p>Escala de grises completa</p>
        </div>
      </div>

      <div class="filter-item sepia">
        <img src="https://picsum.photos/300/200?random=5" alt="Sepia">
        <div class="filter-overlay">
          <h3>Sepia (80%)</h3>
          <p>Tono sepia vintage</p>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
```

## Filtros Avanzados

### Saturate (Saturación)

```css
.saturate-low {
  filter: saturate(0.5);
}

.saturate-high {
  filter: saturate(2);
}
```

### Hue-rotate (Rotación de tono)

```css
.hue-red {
  filter: hue-rotate(0deg);
}

.hue-blue {
  filter: hue-rotate(240deg);
}

.hue-green {
  filter: hue-rotate(120deg);
}
```

### Invert (Invertir colores)

```css
.invert-half {
  filter: invert(0.5);
}

.invert-full {
  filter: invert(1);
}
```

### Opacity

```css
.opacity-half {
  filter: opacity(0.5);
}

.opacity-quarter {
  filter: opacity(0.25);
}
```

### Drop-shadow (Sombra)

```css
.shadow {
  filter: drop-shadow(5px 5px 10px rgba(0,0,0,0.5));
}

.shadow-colored {
  filter: drop-shadow(0 0 10px rgba(255,0,0,0.8));
}
```

## Combinando Múltiples Filtros

### Efectos complejos

```css
.vintage-effect {
  filter: sepia(0.3) contrast(1.2) brightness(1.1) saturate(1.3);
}

.dramatic-effect {
  filter: contrast(1.5) brightness(0.8) saturate(1.2) hue-rotate(5deg);
}

.cartoon-effect {
  filter: contrast(2) saturate(1.5) brightness(1.2) hue-rotate(-10deg);
}
```

## Ejemplo con Filtros Combinados

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Filters - Combinados</title>
  <style>
    .container {
      max-width: 1000px;
      margin: 0 auto;
      padding: 20px;
    }

    .effect-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 30px;
    }

    .effect-item {
      position: relative;
      border-radius: 15px;
      overflow: hidden;
      box-shadow: 0 8px 25px rgba(0,0,0,0.15);
      transition: transform 0.3s ease;
    }

    .effect-item:hover {
      transform: translateY(-5px);
    }

    .effect-item img {
      width: 100%;
      height: 250px;
      object-fit: cover;
      transition: filter 0.3s ease;
    }

    .effect-info {
      padding: 20px;
      background: white;
    }

    .effect-info h3 {
      margin: 0 0 10px 0;
      color: #333;
    }

    .effect-info p {
      margin: 0;
      color: #666;
      font-size: 14px;
    }

    /* Efectos aplicados */
    .vintage img {
      filter: sepia(0.4) contrast(1.3) brightness(1.1) saturate(1.2);
    }

    .dramatic img {
      filter: contrast(1.8) brightness(0.7) saturate(1.4) hue-rotate(-5deg);
    }

    .cyberpunk img {
      filter: hue-rotate(180deg) saturate(2) brightness(1.2) contrast(1.4);
    }

    .warm img {
      filter: sepia(0.2) saturate(1.3) brightness(1.1) contrast(1.1);
    }

    .cool img {
      filter: hue-rotate(180deg) saturate(0.8) brightness(1.1) contrast(1.1);
    }

    .noir img {
      filter: grayscale(1) contrast(1.5) brightness(0.8);
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>CSS Filter Effects - Combinados</h1>

    <div class="effect-grid">
      <div class="effect-item vintage">
        <img src="https://picsum.photos/300/250?random=6" alt="Vintage">
        <div class="effect-info">
          <h3>Efecto Vintage</h3>
          <p>sepia(0.4) contrast(1.3) brightness(1.1) saturate(1.2)</p>
        </div>
      </div>

      <div class="effect-item dramatic">
        <img src="https://picsum.photos/300/250?random=7" alt="Dramatic">
        <div class="effect-info">
          <h3>Efecto Dramático</h3>
          <p>contrast(1.8) brightness(0.7) saturate(1.4) hue-rotate(-5deg)</p>
        </div>
      </div>

      <div class="effect-item cyberpunk">
        <img src="https://picsum.photos/300/250?random=8" alt="Cyberpunk">
        <div class="effect-info">
          <h3>Efecto Cyberpunk</h3>
          <p>hue-rotate(180deg) saturate(2) brightness(1.2) contrast(1.4)</p>
        </div>
      </div>

      <div class="effect-item warm">
        <img src="https://picsum.photos/300/250?random=9" alt="Warm">
        <div class="effect-info">
          <h3>Tonos Cálidos</h3>
          <p>sepia(0.2) saturate(1.3) brightness(1.1) contrast(1.1)</p>
        </div>
      </div>

      <div class="effect-item cool">
        <img src="https://picsum.photos/300/250?random=10" alt="Cool">
        <div class="effect-info">
          <h3>Tonos Fríos</h3>
          <p>hue-rotate(180deg) saturate(0.8) brightness(1.1) contrast(1.1)</p>
        </div>
      </div>

      <div class="effect-item noir">
        <img src="https://picsum.photos/300/250?random=11" alt="Noir">
        <div class="effect-info">
          <h3>Estilo Noir</h3>
          <p>grayscale(1) contrast(1.5) brightness(0.8)</p>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
```

## Filtros con Animaciones

### Transiciones suaves

```css
.element {
  filter: blur(0px) grayscale(0);
  transition: filter 0.5s ease;
}

.element:hover {
  filter: blur(2px) grayscale(1);
}
```

### Animaciones complejas

```css
@keyframes filterPulse {
  0%, 100% {
    filter: brightness(1) saturate(1);
  }
  50% {
    filter: brightness(1.3) saturate(1.5);
  }
}

.pulsing-element {
  animation: filterPulse 2s infinite;
}
```

## Filtros en SVG

### Usando filtros SVG personalizados

```html
<svg width="0" height="0">
  <defs>
    <filter id="customBlur">
      <feGaussianBlur stdDeviation="3"/>
    </filter>
  </defs>
</svg>

<style>
  .custom-filter {
    filter: url(#customBlur);
  }
</style>
```

## Performance y Mejores Prácticas

### Optimizaciones

```css
/* Usar transform3d para forzar aceleración por hardware */
.element {
  transform: translate3d(0, 0, 0);
  filter: blur(5px);
}

/* Evitar filtros en elementos que cambian frecuentemente */
.static-element {
  filter: blur(2px);
}

/* Usar will-change para animaciones */
.animated-element {
  will-change: filter;
}
```

## Casos de Uso Prácticos

### 1. Galerías de imágenes
```css
.gallery-item {
  filter: grayscale(1);
  transition: filter 0.3s ease;
}

.gallery-item:hover {
  filter: grayscale(0);
}
```

### 2. Estados de carga
```css
.loading {
  filter: blur(1px) opacity(0.7);
  pointer-events: none;
}
```

### 3. Efectos de foco
```css
.focused {
  filter: brightness(1.1) contrast(1.1) saturate(1.2);
}
```

### 4. Modo oscuro
```css
.dark-mode {
  filter: invert(1) hue-rotate(180deg);
}
```

## Compatibilidad

### Fallbacks para navegadores antiguos

```css
.element {
  /* Fallback */
  opacity: 0.8;
  /* Filtro moderno */
  filter: opacity(0.8);
}

@supports not (filter: blur(1px)) {
  .fallback-element {
    /* Estilos alternativos */
  }
}
```

## Ejemplo Interactivo Completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Filters - Interactivo</title>
  <style>
    .container {
      max-width: 800px;
      margin: 0 auto;
      padding: 20px;
    }

    .interactive-demo {
      position: relative;
      margin: 40px 0;
    }

    .demo-image {
      width: 100%;
      max-width: 500px;
      height: 300px;
      object-fit: cover;
      border-radius: 10px;
      box-shadow: 0 5px 15px rgba(0,0,0,0.2);
      transition: filter 0.3s ease;
    }

    .controls {
      margin-top: 20px;
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
      gap: 15px;
    }

    .control-group {
      display: flex;
      flex-direction: column;
    }

    .control-group label {
      margin-bottom: 5px;
      font-weight: bold;
      color: #333;
    }

    .control-group input {
      padding: 8px;
      border: 1px solid #ddd;
      border-radius: 5px;
    }

    .preset-buttons {
      margin-top: 20px;
      display: flex;
      flex-wrap: wrap;
      gap: 10px;
    }

    .preset-btn {
      padding: 8px 16px;
      border: 1px solid #007bff;
      background: white;
      color: #007bff;
      border-radius: 5px;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    .preset-btn:hover {
      background: #007bff;
      color: white;
    }

    .preset-btn.active {
      background: #007bff;
      color: white;
    }

    .filter-code {
      margin-top: 20px;
      padding: 15px;
      background: #f8f9fa;
      border-radius: 5px;
      font-family: monospace;
      border: 1px solid #e9ecef;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>CSS Filter Effects - Interactivo</h1>

    <div class="interactive-demo">
      <img id="demoImage" src="https://picsum.photos/500/300?random=12" alt="Demo" class="demo-image">

      <div class="controls">
        <div class="control-group">
          <label for="blur">Blur (px):</label>
          <input type="range" id="blur" min="0" max="20" value="0" step="1">
        </div>

        <div class="control-group">
          <label for="brightness">Brightness:</label>
          <input type="range" id="brightness" min="0" max="3" value="1" step="0.1">
        </div>

        <div class="control-group">
          <label for="contrast">Contrast:</label>
          <input type="range" id="contrast" min="0" max="3" value="1" step="0.1">
        </div>

        <div class="control-group">
          <label for="saturate">Saturate:</label>
          <input type="range" id="saturate" min="0" max="3" value="1" step="0.1">
        </div>

        <div class="control-group">
          <label for="grayscale">Grayscale:</label>
          <input type="range" id="grayscale" min="0" max="1" value="0" step="0.1">
        </div>

        <div class="control-group">
          <label for="sepia">Sepia:</label>
          <input type="range" id="sepia" min="0" max="1" value="0" step="0.1">
        </div>

        <div class="control-group">
          <label for="hueRotate">Hue Rotate (deg):</label>
          <input type="range" id="hueRotate" min="0" max="360" value="0" step="1">
        </div>

        <div class="control-group">
          <label for="invert">Invert:</label>
          <input type="range" id="invert" min="0" max="1" value="0" step="0.1">
        </div>
      </div>

      <div class="preset-buttons">
        <button class="preset-btn" data-preset="none">Sin Filtro</button>
        <button class="preset-btn" data-preset="vintage">Vintage</button>
        <button class="preset-btn" data-preset="dramatic">Dramático</button>
        <button class="preset-btn" data-preset="cyberpunk">Cyberpunk</button>
        <button class="preset-btn" data-preset="noir">Noir</button>
        <button class="preset-btn" data-preset="warm">Cálido</button>
        <button class="preset-btn" data-preset="cool">Frío</button>
      </div>

      <div class="filter-code">
        <strong>CSS aplicado:</strong><br>
        <code id="filterCode">filter: none;</code>
      </div>
    </div>
  </div>

  <script>
    const image = document.getElementById('demoImage');
    const controls = document.querySelectorAll('input[type="range"]');
    const presetButtons = document.querySelectorAll('.preset-btn');
    const filterCode = document.getElementById('filterCode');

    const presets = {
      none: { blur: 0, brightness: 1, contrast: 1, saturate: 1, grayscale: 0, sepia: 0, hueRotate: 0, invert: 0 },
      vintage: { blur: 0, brightness: 1.1, contrast: 1.2, saturate: 1.2, grayscale: 0, sepia: 0.3, hueRotate: 0, invert: 0 },
      dramatic: { blur: 0, brightness: 0.8, contrast: 1.5, saturate: 1.4, grayscale: 0, sepia: 0, hueRotate: 0, invert: 0 },
      cyberpunk: { blur: 0, brightness: 1.2, contrast: 1.4, saturate: 2, grayscale: 0, sepia: 0, hueRotate: 180, invert: 0 },
      noir: { blur: 0, brightness: 0.8, contrast: 1.5, saturate: 0, grayscale: 1, sepia: 0, hueRotate: 0, invert: 0 },
      warm: { blur: 0, brightness: 1.1, contrast: 1.1, saturate: 1.3, grayscale: 0, sepia: 0.2, hueRotate: 0, invert: 0 },
      cool: { blur: 0, brightness: 1.1, contrast: 1.1, saturate: 0.8, grayscale: 0, sepia: 0, hueRotate: 180, invert: 0 }
    };

    function updateFilter() {
      const blur = document.getElementById('blur').value;
      const brightness = document.getElementById('brightness').value;
      const contrast = document.getElementById('contrast').value;
      const saturate = document.getElementById('saturate').value;
      const grayscale = document.getElementById('grayscale').value;
      const sepia = document.getElementById('sepia').value;
      const hueRotate = document.getElementById('hueRotate').value;
      const invert = document.getElementById('invert').value;

      const filterValue = `blur(${blur}px) brightness(${brightness}) contrast(${contrast}) saturate(${saturate}) grayscale(${grayscale}) sepia(${sepia}) hue-rotate(${hueRotate}deg) invert(${invert})`;

      image.style.filter = filterValue;
      filterCode.textContent = `filter: ${filterValue};`;
    }

    controls.forEach(control => {
      control.addEventListener('input', updateFilter);
    });

    presetButtons.forEach(button => {
      button.addEventListener('click', () => {
        const preset = button.dataset.preset;
        const values = presets[preset];

        Object.keys(values).forEach(key => {
          document.getElementById(key).value = values[key];
        });

        presetButtons.forEach(btn => btn.classList.remove('active'));
        button.classList.add('active');

        updateFilter();
      });
    });

    // Inicializar
    updateFilter();
  </script>
</body>
</html>
```

## Resumen

CSS Filter Effects ofrecen un poder increíble para manipular imágenes y elementos visuales:

- ✅ **Filtros básicos**: blur, brightness, contrast, grayscale, sepia
- ✅ **Filtros avanzados**: saturate, hue-rotate, invert, opacity, drop-shadow
- ✅ **Combinaciones**: Crear efectos complejos combinando múltiples filtros
- ✅ **Animaciones**: Transiciones y animaciones suaves entre filtros
- ✅ **Performance**: Optimizado para uso moderno con aceleración por hardware
- ✅ **Compatibilidad**: Excelente soporte en navegadores modernos

Los filtros CSS son esenciales para crear interfaces modernas, galerías interactivas, efectos hover, y experiencias visuales ricas sin necesidad de procesamiento de imágenes en el servidor.