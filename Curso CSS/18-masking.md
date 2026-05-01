# Módulo 18 - Masking

En este módulo aprenderás sobre CSS Masking, una técnica avanzada que permite crear efectos visuales complejos ocultando partes de elementos usando formas, gradientes y imágenes.

## Introducción a CSS Masking

### ¿Qué es CSS Masking?

CSS Masking permite definir qué partes de un elemento son visibles y cuáles son transparentes, creando efectos visuales sofisticados como formas irregulares, transiciones suaves, y composiciones complejas.

```css
/* Máscara básica con imagen */
.element {
  mask-image: url('mask.png');
}

/* Máscara con gradiente */
.element {
  mask-image: linear-gradient(45deg, black 50%, transparent 50%);
}
```

## Mask Image

### Usando imágenes como máscaras

```css
/* Máscara con imagen PNG/SVG */
.element {
  mask-image: url('circle-mask.png');
  mask-size: cover;
  mask-repeat: no-repeat;
  mask-position: center;
}
```

### Propiedades de máscara

```css
.element {
  mask-image: url('mask.svg');
  mask-size: 100px 100px;
  mask-repeat: no-repeat;
  mask-position: center center;
  mask-origin: border-box;
  mask-clip: border-box;
}
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Masking - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: 'Arial', sans-serif;
      background: #1a1a1a;
      color: white;
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
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1, #96ceb4);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    .demo-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 30px;
      margin-bottom: 50px;
    }

    .demo-item {
      position: relative;
      height: 250px;
      border-radius: 15px;
      overflow: hidden;
      box-shadow: 0 10px 30px rgba(0,0,0,0.3);
    }

    .demo-item img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }

    /* Máscaras aplicadas */
    .circle-mask {
      -webkit-mask-image: radial-gradient(circle at center, black 50%, transparent 50%);
      mask-image: radial-gradient(circle at center, black 50%, transparent 50%);
    }

    .diamond-mask {
      -webkit-mask-image: linear-gradient(45deg, transparent 40%, black 40%, black 60%, transparent 60%);
      mask-image: linear-gradient(45deg, transparent 40%, black 40%, black 60%, transparent 60%);
    }

    .heart-mask {
      -webkit-mask-image: radial-gradient(circle at 25% 25%, black 20%, transparent 20%),
                          radial-gradient(circle at 75% 25%, black 20%, transparent 20%),
                          linear-gradient(to bottom, black 40%, transparent 40%);
      mask-image: radial-gradient(circle at 25% 25%, black 20%, transparent 20%),
                  radial-gradient(circle at 75% 25%, black 20%, transparent 20%),
                  linear-gradient(to bottom, black 40%, transparent 40%);
      -webkit-mask-composite: intersect;
      mask-composite: intersect;
    }

    .star-mask {
      -webkit-mask-image:
        polygon(50% 0%, 61% 35%, 98% 35%, 68% 57%, 79% 91%, 50% 70%, 21% 91%, 32% 57%, 2% 35%, 39% 35%);
      mask-image:
        polygon(50% 0%, 61% 35%, 98% 35%, 68% 57%, 79% 91%, 50% 70%, 21% 91%, 32% 57%, 2% 35%, 39% 35%);
    }

    .wave-mask {
      -webkit-mask-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cpath d='M0 50 Q25 0 50 50 T100 50 L100 100 L0 100 Z' fill='black'/%3e%3c/svg%3e");
      mask-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cpath d='M0 50 Q25 0 50 50 T100 50 L100 100 L0 100 Z' fill='black'/%3e%3c/svg%3e");
    }

    .demo-info {
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
      background: linear-gradient(transparent, rgba(0,0,0,0.8));
      color: white;
      padding: 20px;
      transform: translateY(100%);
      transition: transform 0.3s ease;
    }

    .demo-item:hover .demo-info {
      transform: translateY(0);
    }

    .demo-info h3 {
      margin-bottom: 5px;
      font-size: 1.2rem;
    }

    .demo-info p {
      opacity: 0.9;
      font-size: 0.9rem;
    }

    .interactive-demo {
      background: rgba(255,255,255,0.05);
      border-radius: 15px;
      padding: 30px;
      margin-top: 50px;
    }

    .interactive-demo h2 {
      text-align: center;
      margin-bottom: 30px;
      font-size: 2rem;
    }

    .mask-controls {
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
      color: #4ecdc4;
    }

    .control-group select,
    .control-group input {
      padding: 10px;
      border: 1px solid rgba(255,255,255,0.2);
      border-radius: 8px;
      background: rgba(255,255,255,0.1);
      color: white;
      font-size: 14px;
    }

    .control-group select:focus,
    .control-group input:focus {
      outline: none;
      border-color: #4ecdc4;
      box-shadow: 0 0 0 2px rgba(78, 205, 196, 0.3);
    }

    .preview-container {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 300px;
      background: rgba(255,255,255,0.05);
      border-radius: 10px;
      margin-top: 30px;
      position: relative;
      overflow: hidden;
    }

    .preview-image {
      width: 250px;
      height: 250px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1, #96ceb4);
      border-radius: 10px;
      transition: all 0.3s ease;
      -webkit-mask-size: cover;
      mask-size: cover;
      -webkit-mask-repeat: no-repeat;
      mask-repeat: no-repeat;
      -webkit-mask-position: center;
      mask-position: center;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>CSS Masking</h1>
      <p>Efectos visuales avanzados con máscaras</p>
    </header>

    <div class="demo-grid">
      <div class="demo-item circle-mask">
        <img src="https://picsum.photos/300/250?random=1" alt="Circle Mask">
        <div class="demo-info">
          <h3>Máscara Circular</h3>
          <p>Gradiente radial para crear formas circulares perfectas</p>
        </div>
      </div>

      <div class="demo-item diamond-mask">
        <img src="https://picsum.photos/300/250?random=2" alt="Diamond Mask">
        <div class="demo-info">
          <h3>Máscara Diamante</h3>
          <p>Gradiente lineal diagonal para formas geométricas</p>
        </div>
      </div>

      <div class="demo-item heart-mask">
        <img src="https://picsum.photos/300/250?random=3" alt="Heart Mask">
        <div class="demo-info">
          <h3>Máscara Corazón</h3>
          <p>Múltiples gradientes combinados para formas complejas</p>
        </div>
      </div>

      <div class="demo-item star-mask">
        <img src="https://picsum.photos/300/250?random=4" alt="Star Mask">
        <div class="demo-info">
          <h3>Máscara Estrella</h3>
          <p>Función polygon() para formas personalizadas precisas</p>
        </div>
      </div>

      <div class="demo-item wave-mask">
        <img src="https://picsum.photos/300/250?random=5" alt="Wave Mask">
        <div class="demo-info">
          <h3>Máscara Ola</h3>
          <p>SVG inline como máscara para curvas suaves</p>
        </div>
      </div>
    </div>

    <div class="interactive-demo">
      <h2>Demo Interactivo</h2>

      <div class="mask-controls">
        <div class="control-group">
          <label for="maskType">Tipo de Máscara:</label>
          <select id="maskType">
            <option value="circle">Circular</option>
            <option value="diamond">Diamante</option>
            <option value="star">Estrella</option>
            <option value="heart">Corazón</option>
            <option value="wave">Ola</option>
            <option value="custom">Personalizada</option>
          </select>
        </div>

        <div class="control-group">
          <label for="maskSize">Tamaño (%):</label>
          <input type="range" id="maskSize" min="50" max="150" value="100" step="5">
        </div>

        <div class="control-group">
          <label for="maskPosition">Posición:</label>
          <select id="maskPosition">
            <option value="center">Centro</option>
            <option value="top">Arriba</option>
            <option value="bottom">Abajo</option>
            <option value="left">Izquierda</option>
            <option value="right">Derecha</option>
          </select>
        </div>
      </div>

      <div class="preview-container">
        <div class="preview-image" id="previewImage"></div>
      </div>
    </div>
  </div>

  <script>
    const previewImage = document.getElementById('previewImage');
    const maskType = document.getElementById('maskType');
    const maskSize = document.getElementById('maskSize');
    const maskPosition = document.getElementById('maskPosition');

    const masks = {
      circle: 'radial-gradient(circle at center, black 50%, transparent 50%)',
      diamond: 'linear-gradient(45deg, transparent 40%, black 40%, black 60%, transparent 60%)',
      star: 'polygon(50% 0%, 61% 35%, 98% 35%, 68% 57%, 79% 91%, 50% 70%, 21% 91%, 32% 57%, 2% 35%, 39% 35%)',
      heart: 'radial-gradient(circle at 25% 25%, black 20%, transparent 20%), radial-gradient(circle at 75% 25%, black 20%, transparent 20%), linear-gradient(to bottom, black 40%, transparent 40%)',
      wave: 'url("data:image/svg+xml,%3csvg width=\'100\' height=\'100\' xmlns=\'http://www.w3.org/2000/svg\'%3e%3cpath d=\'M0 50 Q25 0 50 50 T100 50 L100 100 L0 100 Z\' fill=\'black\'/%3e%3c/svg%3e")',
      custom: 'linear-gradient(45deg, black 30%, transparent 30%, transparent 70%, black 70%)'
    };

    function updateMask() {
      const type = maskType.value;
      const size = maskSize.value;
      const position = maskPosition.value;

      previewImage.style.webkitMaskImage = masks[type];
      previewImage.style.maskImage = masks[type];
      previewImage.style.webkitMaskSize = `${size}%`;
      previewImage.style.maskSize = `${size}%`;
      previewImage.style.webkitMaskPosition = position;
      previewImage.style.maskPosition = position;
    }

    maskType.addEventListener('change', updateMask);
    maskSize.addEventListener('input', updateMask);
    maskPosition.addEventListener('change', updateMask);

    // Inicializar
    updateMask();
  </script>
</body>
</html>
```

## Gradientes como Máscaras

### Máscaras con gradientes lineales

```css
/* Máscara lineal */
.element {
  mask-image: linear-gradient(to right, transparent 0%, black 50%, transparent 100%);
}

/* Máscara radial */
.element {
  mask-image: radial-gradient(circle, black 30%, transparent 70%);
}

/* Máscara cónica */
.element {
  mask-image: conic-gradient(from 0deg, black 0deg, transparent 180deg, black 360deg);
}
```

## Mask Composite

### Composición de múltiples máscaras

```css
.element {
  mask-image:
    radial-gradient(circle at 30% 30%, black 20%, transparent 20%),
    radial-gradient(circle at 70% 30%, black 20%, transparent 20%),
    linear-gradient(to bottom, transparent 30%, black 30%, black 70%, transparent 70%);
  mask-composite: intersect;
}
```

## Mask Position y Size

### Controlando el tamaño y posición

```css
.element {
  mask-image: url('mask.png');
  mask-size: cover;        /* cover, contain, auto, o valores específicos */
  mask-position: center;   /* center, top, bottom, left, right, o valores específicos */
  mask-repeat: no-repeat;  /* repeat, no-repeat, repeat-x, repeat-y */
}
```

## Ejemplo con Animaciones

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Masking - Animaciones</title>
  <style>
    body {
      margin: 0;
      padding: 20px;
      font-family: 'Arial', sans-serif;
      background: #0a0a0a;
      color: white;
    }

    .container {
      max-width: 1000px;
      margin: 0 auto;
    }

    .animated-masks {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
      gap: 30px;
      margin-bottom: 50px;
    }

    .mask-animation {
      position: relative;
      height: 250px;
      border-radius: 15px;
      overflow: hidden;
      cursor: pointer;
    }

    .mask-animation img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      transition: filter 0.3s ease;
    }

    .mask-animation:hover img {
      filter: brightness(1.1) contrast(1.1);
    }

    /* Animaciones de máscara */
    .reveal-animation {
      -webkit-mask-image: linear-gradient(to right, transparent 0%, black 100%);
      mask-image: linear-gradient(to right, transparent 0%, black 100%);
      -webkit-mask-size: 200% 100%;
      mask-size: 200% 100%;
      -webkit-mask-position: -100% 0;
      mask-position: -100% 0;
      transition: -webkit-mask-position 0.8s ease, mask-position 0.8s ease;
    }

    .reveal-animation:hover {
      -webkit-mask-position: 0 0;
      mask-position: 0 0;
    }

    .circular-reveal {
      -webkit-mask-image: radial-gradient(circle, transparent 30%, black 30%);
      mask-image: radial-gradient(circle, transparent 30%, black 30%);
      -webkit-mask-size: 0% 0%;
      mask-size: 0% 0%;
      -webkit-mask-position: center;
      mask-position: center;
      transition: -webkit-mask-size 0.6s ease;
    }

    .circular-reveal:hover {
      -webkit-mask-size: 300% 300%;
      mask-size: 300% 300%;
    }

    .wipe-animation {
      -webkit-mask-image: linear-gradient(to bottom, transparent 0%, black 100%);
      mask-image: linear-gradient(to bottom, transparent 0%, black 100%);
      -webkit-mask-size: 100% 0%;
      mask-size: 100% 0%;
      -webkit-mask-position: center top;
      mask-position: center top;
      transition: -webkit-mask-size 0.5s ease;
    }

    .wipe-animation:hover {
      -webkit-mask-size: 100% 100%;
      mask-size: 100% 100%;
    }

    .shape-morph {
      -webkit-mask-image: polygon(50% 0%, 100% 50%, 50% 100%, 0% 50%);
      mask-image: polygon(50% 0%, 100% 50%, 50% 100%, 0% 50%);
      transition: -webkit-mask-image 0.6s ease, mask-image 0.6s ease;
    }

    .shape-morph:hover {
      -webkit-mask-image: polygon(50% 20%, 80% 80%, 20% 80%);
      mask-image: polygon(50% 20%, 80% 80%, 20% 80%);
    }

    .rotating-mask {
      -webkit-mask-image: conic-gradient(from 0deg, transparent 0deg, black 90deg, transparent 180deg, black 270deg, transparent 360deg);
      mask-image: conic-gradient(from 0deg, transparent 0deg, black 90deg, transparent 180deg, black 270deg, transparent 360deg);
      animation: rotate-mask 3s linear infinite;
    }

    @keyframes rotate-mask {
      0% { -webkit-mask-image: conic-gradient(from 0deg, transparent 0deg, black 90deg, transparent 180deg, black 270deg, transparent 360deg); }
      25% { -webkit-mask-image: conic-gradient(from 90deg, transparent 0deg, black 90deg, transparent 180deg, black 270deg, transparent 360deg); }
      50% { -webkit-mask-image: conic-gradient(from 180deg, transparent 0deg, black 90deg, transparent 180deg, black 270deg, transparent 360deg); }
      75% { -webkit-mask-image: conic-gradient(from 270deg, transparent 0deg, black 90deg, transparent 180deg, black 270deg, transparent 360deg); }
      100% { -webkit-mask-image: conic-gradient(from 360deg, transparent 0deg, black 90deg, transparent 180deg, black 270deg, transparent 360deg); }
    }

    .mask-info {
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
      background: linear-gradient(transparent, rgba(0,0,0,0.9));
      color: white;
      padding: 15px;
      transform: translateY(100%);
      transition: transform 0.3s ease;
    }

    .mask-animation:hover .mask-info {
      transform: translateY(0);
    }

    .mask-info h3 {
      margin-bottom: 5px;
      font-size: 1.1rem;
    }

    .mask-info p {
      opacity: 0.8;
      font-size: 0.9rem;
    }

    .advanced-demo {
      background: rgba(255,255,255,0.05);
      border-radius: 15px;
      padding: 30px;
      margin-top: 50px;
    }

    .advanced-demo h2 {
      text-align: center;
      margin-bottom: 30px;
      font-size: 2rem;
      color: #4ecdc4;
    }

    .text-mask-demo {
      text-align: center;
      margin-bottom: 40px;
    }

    .text-mask {
      font-size: 4rem;
      font-weight: bold;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1, #96ceb4);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
      -webkit-mask-image: linear-gradient(to right, transparent 0%, black 50%, transparent 100%);
      mask-image: linear-gradient(to right, transparent 0%, black 50%, transparent 100%);
      -webkit-mask-size: 200% 100%;
      mask-size: 200% 100%;
      -webkit-mask-position: -100% 0;
      mask-position: -100% 0;
      animation: text-reveal 2s ease-in-out infinite alternate;
    }

    @keyframes text-reveal {
      0% {
        -webkit-mask-position: -100% 0;
        mask-position: -100% 0;
      }
      100% {
        -webkit-mask-position: 100% 0;
        mask-position: 100% 0;
      }
    }

    .composite-demo {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 300px;
      background: linear-gradient(45deg, #1a1a1a, #2d2d2d);
      border-radius: 15px;
      position: relative;
    }

    .composite-element {
      width: 200px;
      height: 200px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      border-radius: 10px;
      -webkit-mask-image:
        radial-gradient(circle at 30% 30%, black 25%, transparent 25%),
        radial-gradient(circle at 70% 70%, black 25%, transparent 25%),
        linear-gradient(to right, transparent 40%, black 40%, black 60%, transparent 60%);
      mask-image:
        radial-gradient(circle at 30% 30%, black 25%, transparent 25%),
        radial-gradient(circle at 70% 70%, black 25%, transparent 25%),
        linear-gradient(to right, transparent 40%, black 40%, black 60%, transparent 60%);
      -webkit-mask-composite: intersect;
      mask-composite: intersect;
      animation: composite-rotate 4s ease-in-out infinite;
    }

    @keyframes composite-rotate {
      0%, 100% { transform: rotate(0deg); }
      25% { transform: rotate(90deg); }
      50% { transform: rotate(180deg); }
      75% { transform: rotate(270deg); }
    }
  </style>
</head>
<body>
  <div class="container">
    <h1 style="text-align: center; margin-bottom: 40px; font-size: 2.5rem;">
      CSS Masking - Animaciones
    </h1>

    <div class="animated-masks">
      <div class="mask-animation reveal-animation">
        <img src="https://picsum.photos/250/250?random=10" alt="Reveal">
        <div class="mask-info">
          <h3>Revelación Lateral</h3>
          <p>La máscara se desliza desde la izquierda revelando la imagen</p>
        </div>
      </div>

      <div class="mask-animation circular-reveal">
        <img src="https://picsum.photos/250/250?random=11" alt="Circular Reveal">
        <div class="mask-info">
          <h3>Revelación Circular</h3>
          <p>Un círculo expansivo revela la imagen desde el centro</p>
        </div>
      </div>

      <div class="mask-animation wipe-animation">
        <img src="https://picsum.photos/250/250?random=12" alt="Wipe">
        <div class="mask-info">
          <h3>Barrido Vertical</h3>
          <p>La máscara se expande desde arriba hacia abajo</p>
        </div>
      </div>

      <div class="mask-animation shape-morph">
        <img src="https://picsum.photos/250/250?random=13" alt="Shape Morph">
        <div class="mask-info">
          <h3>Morfing de Forma</h3>
          <p>La forma de la máscara cambia al hacer hover</p>
        </div>
      </div>

      <div class="mask-animation rotating-mask">
        <img src="https://picsum.photos/250/250?random=14" alt="Rotating">
        <div class="mask-info">
          <h3>Máscara Rotatoria</h3>
          <p>La máscara gira continuamente creando un efecto dinámico</p>
        </div>
      </div>
    </div>

    <div class="advanced-demo">
      <h2>Ejemplos Avanzados</h2>

      <div class="text-mask-demo">
        <div class="text-mask">MASKING</div>
        <p style="margin-top: 20px; opacity: 0.8;">Texto con máscara animada</p>
      </div>

      <div class="composite-demo">
        <div class="composite-element"></div>
      </div>
      <p style="text-align: center; margin-top: 20px; opacity: 0.8;">
        Máscara compuesta con múltiples gradientes
      </p>
    </div>
  </div>
</body>
</html>
```

## SVG Masks

### Usando SVG para máscaras complejas

```html
<svg width="0" height="0">
  <defs>
    <mask id="complexMask">
      <rect width="100%" height="100%" fill="white"/>
      <circle cx="50%" cy="50%" r="30%" fill="black"/>
      <rect x="20%" y="20%" width="60%" height="60%" fill="black"/>
    </mask>
  </defs>
</svg>

<style>
  .element {
    mask: url(#complexMask);
  }
</style>
```

## Casos de Uso Prácticos

### 1. Avatares con formas personalizadas

```css
.avatar {
  width: 100px;
  height: 100px;
  mask-image: url('hexagon-mask.svg');
  mask-size: cover;
}
```

### 2. Efectos de transición

```css
.page-transition {
  mask-image: radial-gradient(circle, black 0%, transparent 100%);
  mask-size: 0% 0%;
  transition: mask-size 0.5s ease;
}

.page-transition.active {
  mask-size: 300% 300%;
}
```

### 3. Texto con efectos

```css
.text-effect {
  background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  mask-image: linear-gradient(to right, transparent 0%, black 100%);
}
```

## Performance Considerations

### Optimizaciones

```css
/* Usar mask-image en lugar de múltiples elementos */
.efficient-mask {
  mask-image: url('mask.svg');
  /* Mejor que múltiples elementos posicionados */
}

/* Evitar máscaras complejas en elementos animados */
.simple-mask {
  mask-image: circle(50%);
  /* Mejor performance que gradientes complejos */
}
```

## Compatibilidad

### Fallbacks para navegadores antiguos

```css
.element {
  /* Fallback básico */
  border-radius: 50%;
  overflow: hidden;

  /* Máscara moderna */
  mask-image: url('circle-mask.svg');
}

@supports not (mask-image: url('test.svg')) {
  .mask-fallback {
    /* Estilos alternativos */
    clip-path: circle(50%);
  }
}
```

## Resumen

CSS Masking ofrece control creativo sobre la visibilidad de elementos:

- ✅ **Mask Image**: Imágenes, gradientes y SVG como máscaras
- ✅ **Composición**: Combinar múltiples máscaras con mask-composite
- ✅ **Animaciones**: Transiciones y animaciones fluidas
- ✅ **Performance**: Optimizado para uso moderno
- ✅ **Compatibilidad**: Excelente soporte con prefijos -webkit-
- ✅ **Creatividad**: Efectos visuales ilimitados

Las máscaras CSS son ideales para crear interfaces modernas, transiciones elegantes, y efectos visuales que van más allá de las formas tradicionales basadas en cajas rectangulares.