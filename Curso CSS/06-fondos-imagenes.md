# Módulo 06 - Fondos e imágenes

En este módulo aprenderás a trabajar con fondos, imágenes de fondo, gradientes y técnicas avanzadas para crear diseños visuales atractivos.

## Propiedades básicas de fondo

### Color de fondo (background-color)

```css
.element {
  background-color: #f0f0f0;        /* Hex */
  background-color: rgb(240, 240, 240); /* RGB */
  background-color: rgba(240, 240, 240, 0.8); /* RGBA */
  background-color: hsl(0, 0%, 94%); /* HSL */
  background-color: transparent;    /* Transparente */
}
```

### Imagen de fondo (background-image)

```css
.element {
  background-image: url('imagen.jpg');
  background-image: url('imagen.webp'); /* Mejor compresión */
  background-image: none; /* Sin imagen */
}
```

### Repetición de fondo (background-repeat)

```css
.element {
  background-repeat: repeat;     /* Repite en ambas direcciones (default) */
  background-repeat: no-repeat;  /* No repite */
  background-repeat: repeat-x;   /* Solo horizontal */
  background-repeat: repeat-y;   /* Solo vertical */
  background-repeat: round;      /* Ajusta para llenar sin cortar */
  background-repeat: space;      /* Espacia uniformemente */
}
```

### Posición de fondo (background-position)

```css
.element {
  background-position: left top;        /* Palabras clave */
  background-position: center center;   /* Centro */
  background-position: right bottom;    /* Esquina inferior derecha */
  background-position: 50% 50%;         /* Porcentajes */
  background-position: 20px 30px;       /* Valores absolutos */
}
```

### Tamaño de fondo (background-size)

```css
.element {
  background-size: auto;         /* Tamaño original (default) */
  background-size: cover;        /* Cubre todo el elemento */
  background-size: contain;      /* Cabe dentro del elemento */
  background-size: 100% 100%;    /* Estira al 100% */
  background-size: 200px 150px;  /* Tamaño específico */
}
```

## Propiedad shorthand background

```css
.element {
  /* Orden: color image position/size repeat attachment origin clip */
  background: #f0f0f0 url('pattern.png') center/cover no-repeat fixed padding-box content-box;
  
  /* Ejemplos comunes */
  background: #fff url('hero.jpg') center/cover no-repeat;
  background: linear-gradient(to right, #ff0000, #0000ff);
  background: radial-gradient(circle, #ff0000, #0000ff);
}
```

## Propiedades avanzadas de fondo

### Adjunto de fondo (background-attachment)

```css
.element {
  background-attachment: scroll;  /* Se mueve con el contenido (default) */
  background-attachment: fixed;   /* Fijo en la ventana */
  background-attachment: local;   /* Se mueve con el elemento */
}
```

### Origen de fondo (background-origin)

```css
.element {
  background-origin: padding-box;  /* Desde el padding (default) */
  background-origin: border-box;   /* Desde el borde */
  background-origin: content-box;  /* Desde el contenido */
}
```

### Recorte de fondo (background-clip)

```css
.element {
  background-clip: border-box;   /* Hasta el borde (default) */
  background-clip: padding-box;  /* Hasta el padding */
  background-clip: content-box;  /* Hasta el contenido */
  background-clip: text;         /* Solo el texto (con -webkit-) */
}
```

## Múltiples fondos

```css
.element {
  background: 
    url('overlay.png') center/cover no-repeat,
    linear-gradient(45deg, rgba(0,0,0,0.3), rgba(0,0,0,0.7)),
    url('background.jpg') center/cover no-repeat;
  
  /* Cada fondo separado por coma */
}
```

## Gradientes

### Gradiente lineal (linear-gradient)

```css
.element {
  background: linear-gradient(direction, color1, color2, ...);
  
  /* Direcciones */
  background: linear-gradient(to right, #ff0000, #0000ff);
  background: linear-gradient(to bottom, #ff0000, #0000ff);
  background: linear-gradient(45deg, #ff0000, #0000ff);
  background: linear-gradient(to top right, #ff0000, #0000ff);
  
  /* Múltiples colores */
  background: linear-gradient(to right, red, orange, yellow, green);
  
  /* Con posiciones */
  background: linear-gradient(to right, red 0%, orange 25%, yellow 50%, green 100%);
}
```

### Gradiente radial (radial-gradient)

```css
.element {
  background: radial-gradient(shape size at position, color1, color2, ...);
  
  /* Formas básicas */
  background: radial-gradient(circle, red, blue);
  background: radial-gradient(ellipse, red, blue);
  
  /* Con tamaño */
  background: radial-gradient(circle 100px, red, blue);
  background: radial-gradient(ellipse 200px 100px, red, blue);
  
  /* Con posición */
  background: radial-gradient(circle at top, red, blue);
  background: radial-gradient(circle at 50px 50px, red, blue);
}
```

### Gradiente cónico (conic-gradient)

```css
.element {
  background: conic-gradient(from angle at position, color1, color2, ...);
  
  background: conic-gradient(red, yellow, green);
  background: conic-gradient(from 45deg, red, yellow, green);
  background: conic-gradient(from 45deg at center, red, yellow, green);
}
```

## Patrones y texturas

### Patrones CSS puros

```css
/* Cuadros */
.checkered {
  background: 
    linear-gradient(45deg, #ccc 25%, transparent 25%),
    linear-gradient(-45deg, #ccc 25%, transparent 25%),
    linear-gradient(45deg, transparent 75%, #ccc 75%),
    linear-gradient(-45deg, transparent 75%, #ccc 75%);
  background-size: 20px 20px;
  background-position: 0 0, 0 10px, 10px -10px, -10px 0px;
}

/* Rayas */
.striped {
  background: linear-gradient(90deg, #fff 50%, #000 50%);
  background-size: 20px 20px;
}

/* Polka dots */
.polka {
  background: radial-gradient(circle, #000 10%, transparent 10%);
  background-size: 20px 20px;
}
```

## Imágenes responsivas

### Imágenes de fondo responsivas

```css
.hero {
  background-image: url('hero-small.jpg');
  background-size: cover;
  background-position: center;
}

@media (min-width: 768px) {
  .hero {
    background-image: url('hero-medium.jpg');
  }
}

@media (min-width: 1200px) {
  .hero {
    background-image: url('hero-large.jpg');
  }
}
```

### Usando image-set para diferentes densidades

```css
.element {
  background-image: image-set(
    url('image-1x.jpg') 1x,
    url('image-2x.jpg') 2x,
    url('image-3x.jpg') 3x
  );
}
```

## Técnicas avanzadas

### Fondos con parallax

```css
.parallax {
  background-image: url('mountain.jpg');
  background-attachment: fixed;
  background-size: cover;
  background-position: center;
  height: 100vh;
}
```

### Fondos animados

```css
.animated-bg {
  background: linear-gradient(45deg, #ff0000, #0000ff, #00ff00);
  background-size: 300% 300%;
  animation: gradientShift 5s ease infinite;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}
```

### Efectos de vidrio (glassmorphism)

```css
.glass {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}
```

## Optimización de imágenes

### Formatos modernos

```css
/* WebP con fallback */
.element {
  background-image: url('image.webp');
}

@supports not (background-image: url('image.webp')) {
  .element {
    background-image: url('image.jpg');
  }
}
```

### Lazy loading para fondos

```javascript
// Usando Intersection Observer
const observer = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      entry.target.style.backgroundImage = `url(${entry.target.dataset.bg})`;
      observer.unobserve(entry.target);
    }
  });
});

document.querySelectorAll('.lazy-bg').forEach(el => {
  observer.observe(el);
});
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Fondos e Imágenes CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="container">
    <header class="hero">
      <h1>Maestría en Fondos CSS</h1>
      <p>Descubre el poder de los fondos en CSS</p>
    </header>

    <main>
      <section class="background-colors">
        <h2>Colores de Fondo</h2>
        <div class="color-samples">
          <div class="sample solid">Color sólido</div>
          <div class="sample transparent">Transparente</div>
          <div class="sample rgba">RGBA</div>
        </div>
      </section>

      <section class="background-images">
        <h2>Imágenes de Fondo</h2>
        <div class="image-examples">
          <div class="example no-repeat">No repeat</div>
          <div class="example repeat">Repeat</div>
          <div class="example cover">Cover</div>
          <div class="example contain">Contain</div>
        </div>
      </section>

      <section class="gradients">
        <h2>Gradientes</h2>
        <div class="gradient-examples">
          <div class="gradient linear">Lineal</div>
          <div class="gradient radial">Radial</div>
          <div class="gradient conic">Cónico</div>
          <div class="gradient multi">Múltiple</div>
        </div>
      </section>

      <section class="patterns">
        <h2>Patrones CSS</h2>
        <div class="pattern-examples">
          <div class="pattern stripes">Rayas</div>
          <div class="pattern checkered">Cuadros</div>
          <div class="pattern dots">Puntos</div>
        </div>
      </section>

      <section class="advanced">
        <h2>Técnicas Avanzadas</h2>
        <div class="advanced-examples">
          <div class="example parallax">Parallax</div>
          <div class="example glass">Glassmorphism</div>
          <div class="example animated">Animado</div>
        </div>
      </section>

      <section class="responsive">
        <h2>Imágenes Responsivas</h2>
        <div class="responsive-bg">
          <p>Esta imagen de fondo cambia según el tamaño de pantalla</p>
        </div>
      </section>
    </main>
  </div>
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
}

/* Hero section */
.hero {
  height: 60vh;
  background: linear-gradient(135deg, var(--primary), var(--secondary)),
              url('https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=1200') center/cover no-repeat;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: white;
  text-align: center;
  position: relative;
}

.hero::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4);
}

.hero h1, .hero p {
  position: relative;
  z-index: 1;
}

.hero h1 {
  font-size: clamp(2rem, 8vw, 4rem);
  margin-bottom: 1rem;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
}

.hero p {
  font-size: clamp(1rem, 4vw, 1.5rem);
  opacity: 0.9;
}

/* Secciones */
section {
  padding: 60px 20px;
  max-width: 1200px;
  margin: 0 auto;
}

h2 {
  text-align: center;
  font-size: 2.5rem;
  margin-bottom: 40px;
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

/* Colores de fondo */
.color-samples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.sample {
  height: 150px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  font-size: 1.2rem;
  box-shadow: 0 4px 15px rgba(0,0,0,0.1);
}

.solid { background-color: var(--primary); }
.transparent { background-color: transparent; border: 2px solid var(--primary); }
.rgba { background-color: rgba(52, 152, 219, 0.7); }

/* Imágenes de fondo */
.image-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.example {
  height: 200px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.7);
  background: url('https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=400') center/cover no-repeat;
}

.no-repeat { background-repeat: no-repeat; background-size: auto; }
.repeat { background-repeat: repeat; background-size: 50px 50px; }
.cover { background-size: cover; }
.contain { background-size: contain; background-color: var(--light); }

/* Gradientes */
.gradient-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.gradient {
  height: 150px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  text-shadow: 1px 1px 2px rgba(0,0,0,0.5);
}

.linear { background: linear-gradient(45deg, #ff6b6b, #4ecdc4); }
.radial { background: radial-gradient(circle, #667eea, #764ba2); }
.conic { background: conic-gradient(from 45deg, #ff9a9e, #fecfef, #a8edea); }
.multi { background: linear-gradient(45deg, transparent 30%, rgba(255,255,255,0.5) 50%, transparent 70%), linear-gradient(to right, #ff6b6b, #4ecdc4); }

/* Patrones */
.pattern-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.pattern {
  height: 150px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--dark);
  font-weight: bold;
}

.stripes {
  background: repeating-linear-gradient(45deg, #fff, #fff 10px, #000 10px, #000 20px);
}

.checkered {
  background: 
    linear-gradient(45deg, #000 25%, transparent 25%),
    linear-gradient(-45deg, #000 25%, transparent 25%),
    linear-gradient(45deg, transparent 75%, #000 75%),
    linear-gradient(-45deg, transparent 75%, #000 75%);
  background-size: 20px 20px;
  background-position: 0 0, 0 10px, 10px -10px, -10px 0px;
}

.dots {
  background: radial-gradient(circle, #000 3px, transparent 3px);
  background-size: 20px 20px;
}

/* Técnicas avanzadas */
.advanced-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.parallax {
  height: 300px;
  background: url('https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=600') center/cover no-repeat fixed;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 1.5rem;
  font-weight: bold;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.7);
}

.glass {
  height: 200px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--dark);
  font-weight: bold;
}

.animated {
  height: 200px;
  background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1, #667eea);
  background-size: 400% 400%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  animation: gradientShift 4s ease infinite;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* Imagen responsiva */
.responsive-bg {
  height: 300px;
  background: url('https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=400') center/cover no-repeat;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 1.2rem;
  font-weight: bold;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.7);
  border-radius: 8px;
}

@media (min-width: 768px) {
  .responsive-bg {
    background-image: url('https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=800');
  }
}

@media (min-width: 1200px) {
  .responsive-bg {
    background-image: url('https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=1200');
  }
}

/* Responsive */
@media (max-width: 768px) {
  section {
    padding: 40px 15px;
  }
  
  .color-samples, .image-examples, .gradient-examples, 
  .pattern-examples, .advanced-examples {
    grid-template-columns: 1fr;
  }
  
  .example, .sample, .gradient, .pattern {
    height: 120px;
  }
}
```

## Resumen

Los fondos son esenciales para el diseño web moderno:

- ✅ **Colores**: Sólidos, transparentes, RGBA
- ✅ **Imágenes**: Posición, tamaño, repetición
- ✅ **Gradientes**: Lineales, radiales, cónicos
- ✅ **Patrones**: CSS puro sin imágenes
- ✅ **Técnicas avanzadas**: Parallax, glassmorphism, animaciones
- ✅ **Responsive**: Imágenes que se adaptan

Los fondos pueden transformar completamente la apariencia de un sitio web. Experimenta con combinaciones para crear diseños únicos.