# Módulo 07 - Bordes y sombras

En este módulo aprenderás a crear bordes atractivos, sombras realistas y efectos visuales avanzados que harán que tus diseños destaquen.

## Propiedades de borde básicas

### Ancho del borde (border-width)

```css
.element {
  border-width: thin;      /* Delgado */
  border-width: medium;    /* Medio (default) */
  border-width: thick;     /* Grueso */
  border-width: 2px;       /* Valor específico */
  border-width: 0.5em;     /* Relativo al tamaño de fuente */
}
```

### Estilo del borde (border-style)

```css
.element {
  border-style: none;      /* Sin borde */
  border-style: solid;     /* Línea continua */
  border-style: dotted;    /* Puntos */
  border-style: dashed;     /* Líneas discontinuas */
  border-style: double;     /* Doble línea */
  border-style: groove;     /* Efecto 3D hundido */
  border-style: ridge;      /* Efecto 3D elevado */
  border-style: inset;      /* Efecto 3D interno */
  border-style: outset;     /* Efecto 3D externo */
}
```

### Color del borde (border-color)

```css
.element {
  border-color: #000;                    /* Hex */
  border-color: rgb(0, 0, 0);           /* RGB */
  border-color: rgba(0, 0, 0, 0.5);     /* RGBA */
  border-color: hsl(0, 0%, 0%);         /* HSL */
  border-color: currentColor;           /* Color del texto */
  border-color: transparent;             /* Transparente */
}
```

## Propiedad shorthand border

```css
.element {
  /* Orden: width style color */
  border: 2px solid #000;
  border: 1px dashed red;
  border: 3px double rgba(0,0,0,0.3);
  border: none; /* Sin borde */
}
```

## Bordes individuales

```css
.element {
  border-top: 2px solid red;
  border-right: 3px dashed blue;
  border-bottom: 4px dotted green;
  border-left: 5px double purple;
  
  /* Propiedades individuales */
  border-top-width: 2px;
  border-top-style: solid;
  border-top-color: red;
}
```

## Bordes redondeados (border-radius)

```css
.element {
  border-radius: 5px;           /* Todos los lados */
  border-radius: 10px 20px;      /* Arriba-izq/abajo-der | Arriba-der/abajo-izq */
  border-radius: 10px 20px 30px; /* Arriba-izq | Arriba-der/abajo-izq | Abajo-der */
  border-radius: 10px 20px 30px 40px; /* Arriba-izq | Arriba-der | Abajo-der | Abajo-izq */
  
  /* Específico por esquina */
  border-top-left-radius: 10px;
  border-top-right-radius: 20px;
  border-bottom-right-radius: 30px;
  border-bottom-left-radius: 40px;
  
  /* Formas especiales */
  border-radius: 50%;            /* Círculo */
  border-radius: 50px / 20px;    /* Elipse (radio-x / radio-y) */
}
```

## Sombras de caja (box-shadow)

### Sintaxis básica

```css
.element {
  box-shadow: offset-x offset-y blur-radius spread-radius color inset;
  
  /* Ejemplos */
  box-shadow: 5px 5px;                          /* Sombra simple */
  box-shadow: 5px 5px 10px;                     /* Con desenfoque */
  box-shadow: 5px 5px 10px 2px rgba(0,0,0,0.3); /* Con extensión */
  box-shadow: inset 5px 5px 10px rgba(0,0,0,0.3); /* Interior */
}
```

### Sombras realistas

```css
/* Sombra sutil */
.card {
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

/* Sombra media */
.card {
  box-shadow: 0 4px 8px rgba(0,0,0,0.15);
}

/* Sombra fuerte */
.card {
  box-shadow: 0 8px 16px rgba(0,0,0,0.2);
}

/* Sombra flotante */
.floating {
  box-shadow: 0 10px 30px rgba(0,0,0,0.3);
}
```

### Múltiples sombras

```css
.element {
  box-shadow: 
    0 1px 3px rgba(0,0,0,0.12),
    0 1px 2px rgba(0,0,0,0.24);
    
  /* Efecto de profundidad */
  box-shadow: 
    0 1px 2px rgba(0,0,0,0.1),
    0 2px 4px rgba(0,0,0,0.1),
    0 4px 8px rgba(0,0,0,0.1);
}
```

## Sombras de texto (text-shadow)

```css
.element {
  text-shadow: offset-x offset-y blur-radius color;
  
  /* Ejemplos */
  text-shadow: 1px 1px 2px rgba(0,0,0,0.5);     /* Sombra sutil */
  text-shadow: 2px 2px 4px #000;                 /* Sombra negra */
  text-shadow: 0 0 10px rgba(255,0,0,0.8);      /* Resplandor */
  text-shadow: 1px 1px 0 #000, -1px -1px 0 #000; /* Borde */
}
```

### Efectos de texto avanzados

```css
/* Texto 3D */
.text-3d {
  text-shadow: 
    1px 1px 0 #ccc,
    2px 2px 0 #c9c9c9,
    3px 3px 0 #bbb,
    4px 4px 0 #b9b9b9,
    5px 5px 0 #aaa,
    6px 6px 1px rgba(0,0,0,0.1),
    0 0 5px rgba(0,0,0,0.1);
}

/* Texto con resplandor */
.glow {
  text-shadow: 0 0 5px #fff, 0 0 10px #fff, 0 0 15px #fff, 0 0 20px #ff00de;
}

/* Texto fuego */
.fire {
  text-shadow: 
    0 0 5px #fff,
    0 0 10px #fff,
    0 0 15px #fff,
    0 0 20px #ff00de,
    0 0 35px #ff00de,
    0 0 40px #ff00de;
}
```

## Bordes con imágenes

### Border-image-source

```css
.element {
  border-image-source: url('border-pattern.png');
  border-image-source: linear-gradient(45deg, #ff0000, #0000ff);
}
```

### Border-image-slice

```css
.element {
  border-image-slice: 30;        /* Un valor para todos */
  border-image-slice: 30 35;     /* Vertical | Horizontal */
  border-image-slice: 30 35 40;  /* Arriba | Horizontal | Abajo */
  border-image-slice: 30 35 40 45; /* Arriba | Derecha | Abajo | Izquierda */
  
  /* Con fill */
  border-image-slice: 30 fill;   /* Rellena el centro */
}
```

### Border-image-width

```css
.element {
  border-image-width: 10px;      /* Ancho fijo */
  border-image-width: 1;         /* Multiplicador del border-width */
  border-image-width: auto;      /* Automático */
}
```

### Border-image-outset

```css
.element {
  border-image-outset: 10px;     /* Expande el borde */
  border-image-outset: 1 2;      /* Vertical | Horizontal */
}
```

### Border-image-repeat

```css
.element {
  border-image-repeat: stretch;   /* Estira (default) */
  border-image-repeat: repeat;    /* Repite */
  border-image-repeat: round;     /* Ajusta para llenar */
  border-image-repeat: space;     /* Espacia */
}
```

### Propiedad shorthand

```css
.element {
  border-image: url('border.png') 30 / 10px / 5px round;
  /* source | slice / width / outset | repeat */
}
```

## Efectos avanzados

### Bordes animados

```css
.animated-border {
  position: relative;
  padding: 20px;
  background: white;
}

.animated-border::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, #ff0000, #00ff00, #0000ff, #ff0000);
  background-size: 400% 400%;
  z-index: -1;
  animation: borderRotate 3s linear infinite;
}

@keyframes borderRotate {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}
```

### Sombras dinámicas

```css
.dynamic-shadow {
  transition: box-shadow 0.3s ease;
}

.dynamic-shadow:hover {
  box-shadow: 
    0 20px 40px rgba(0,0,0,0.1),
    0 10px 20px rgba(0,0,0,0.1),
    0 5px 10px rgba(0,0,0,0.1);
}
```

### Efectos de profundidad

```css
.depth-effect {
  box-shadow: 
    0 1px 3px rgba(0,0,0,0.12),
    0 1px 2px rgba(0,0,0,0.24);
  transition: all 0.3s cubic-bezier(.25,.8,.25,1);
}

.depth-effect:hover {
  box-shadow: 
    0 14px 28px rgba(0,0,0,0.25),
    0 10px 10px rgba(0,0,0,0.22);
  transform: translateY(-2px);
}
```

## Técnicas de diseño

### Tarjetas con sombras

```css
.card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  transition: box-shadow 0.3s ease;
}

.card:hover {
  box-shadow: 0 8px 25px rgba(0,0,0,0.15);
}
```

### Botones con efectos

```css
.btn {
  padding: 12px 24px;
  background: #3498db;
  color: white;
  border: none;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.2);
  transition: all 0.3s ease;
}

.btn:hover {
  box-shadow: 0 4px 12px rgba(52,152,219,0.4);
  transform: translateY(-2px);
}

.btn:active {
  box-shadow: 0 1px 2px rgba(0,0,0,0.2);
  transform: translateY(0);
}
```

### Formularios con bordes

```css
.input {
  padding: 12px;
  border: 2px solid #ddd;
  border-radius: 4px;
  transition: border-color 0.3s ease;
}

.input:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52,152,219,0.1);
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Bordes y Sombras CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="container">
    <header>
      <h1>Bordes y Sombras Maestros</h1>
      <p>Domina los efectos visuales en CSS</p>
    </header>

    <main>
      <section class="border-styles">
        <h2>Estilos de Borde</h2>
        <div class="border-examples">
          <div class="example solid">Solid</div>
          <div class="example dotted">Dotted</div>
          <div class="example dashed">Dashed</div>
          <div class="example double">Double</div>
          <div class="example groove">Groove</div>
          <div class="example ridge">Ridge</div>
        </div>
      </section>

      <section class="border-radius">
        <h2>Bordes Redondeados</h2>
        <div class="radius-examples">
          <div class="example small-radius">Pequeño</div>
          <div class="example medium-radius">Medio</div>
          <div class="example large-radius">Grande</div>
          <div class="example circle">Círculo</div>
          <div class="example pill">Píldora</div>
        </div>
      </section>

      <section class="box-shadows">
        <h2>Sombras de Caja</h2>
        <div class="shadow-examples">
          <div class="example subtle">Sutil</div>
          <div class="example medium">Media</div>
          <div class="example strong">Fuerte</div>
          <div class="example floating">Flotante</div>
          <div class="example inset">Interior</div>
        </div>
      </section>

      <section class="text-shadows">
        <h2>Sombras de Texto</h2>
        <div class="text-examples">
          <h3 class="subtle-text">Sombra sutil</h3>
          <h3 class="glow-text">Texto brillante</h3>
          <h3 class="fire-text">Texto fuego</h3>
          <h3 class="outline-text">Texto con borde</h3>
        </div>
      </section>

      <section class="cards">
        <h2>Tarjetas con Efectos</h2>
        <div class="card-grid">
          <div class="card">
            <h3>Tarjeta Básica</h3>
            <p>Sombra sutil para contenido normal</p>
          </div>
          <div class="card hover-card">
            <h3>Tarjeta Interactiva</h3>
            <p>Efectos al pasar el mouse</p>
          </div>
          <div class="card depth-card">
            <h3>Tarjeta 3D</h3>
            <p>Efectos de profundidad</p>
          </div>
        </div>
      </section>

      <section class="buttons">
        <h2>Botones con Sombras</h2>
        <div class="button-examples">
          <button class="btn primary">Botón Primario</button>
          <button class="btn secondary">Botón Secundario</button>
          <button class="btn danger">Botón Peligro</button>
          <button class="btn animated">Botón Animado</button>
        </div>
      </section>

      <section class="forms">
        <h2>Formularios con Bordes</h2>
        <form class="form-example">
          <div class="form-group">
            <label for="name">Nombre</label>
            <input type="text" id="name" class="input" placeholder="Tu nombre">
          </div>
          <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" class="input" placeholder="tu@email.com">
          </div>
          <button type="submit" class="btn submit">Enviar</button>
        </form>
      </section>

      <section class="advanced-effects">
        <h2>Efectos Avanzados</h2>
        <div class="effect-examples">
          <div class="animated-border">
            <p>Borde animado con gradiente</p>
          </div>
          <div class="glass-effect">
            <p>Efecto vidrio (glassmorphism)</p>
          </div>
          <div class="neon-effect">
            <p>Efecto neón</p>
          </div>
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
  --danger: #e74c3c;
  --warning: #f39c12;
  --dark: #2c3e50;
  --light: #ecf0f1;
  --shadow: rgba(0,0,0,0.1);
  --shadow-medium: rgba(0,0,0,0.15);
  --shadow-strong: rgba(0,0,0,0.2);
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

/* Header */
header {
  text-align: center;
  padding: 60px 20px;
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: white;
}

header h1 {
  font-size: 3rem;
  margin-bottom: 10px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
}

header p {
  font-size: 1.2rem;
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

/* Estilos de borde */
.border-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.example {
  height: 120px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 1.1rem;
  background: white;
  color: var(--dark);
}

.solid { border: 3px solid var(--primary); }
.dotted { border: 3px dotted var(--secondary); }
.dashed { border: 3px dashed var(--danger); }
.double { border: 3px double var(--warning); }
.groove { border: 3px groove var(--primary); }
.ridge { border: 3px ridge var(--secondary); }

/* Bordes redondeados */
.radius-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 20px;
}

.small-radius { border: 3px solid var(--primary); border-radius: 5px; }
.medium-radius { border: 3px solid var(--secondary); border-radius: 15px; }
.large-radius { border: 3px solid var(--danger); border-radius: 25px; }
.circle { border: 3px solid var(--warning); border-radius: 50%; width: 120px; height: 120px; }
.pill { border: 3px solid var(--primary); border-radius: 60px; padding: 10px 20px; }

/* Sombras de caja */
.shadow-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 20px;
}

.subtle { box-shadow: 0 2px 4px var(--shadow); }
.medium { box-shadow: 0 4px 8px var(--shadow-medium); }
.strong { box-shadow: 0 8px 16px var(--shadow-strong); }
.floating { box-shadow: 0 10px 30px rgba(0,0,0,0.3); }
.inset { box-shadow: inset 0 2px 4px var(--shadow); background: #f8f9fa; }

/* Sombras de texto */
.text-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  text-align: center;
}

.subtle-text { text-shadow: 1px 1px 2px rgba(0,0,0,0.3); }
.glow-text { text-shadow: 0 0 10px rgba(52,152,219,0.8); color: var(--primary); }
.fire-text { 
  text-shadow: 0 0 5px #fff, 0 0 10px #fff, 0 0 15px #fff, 0 0 20px #ff00de;
  color: #ff00de;
}
.outline-text { 
  text-shadow: 1px 1px 0 #000, -1px -1px 0 #000, 1px -1px 0 #000, -1px 1px 0 #000;
  color: white;
}

/* Tarjetas */
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
}

.card {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 8px var(--shadow);
  transition: all 0.3s ease;
}

.card h3 {
  margin-bottom: 15px;
  color: var(--dark);
}

.card p {
  color: #666;
}

.hover-card:hover {
  box-shadow: 0 8px 25px var(--shadow-medium);
  transform: translateY(-5px);
}

.depth-card {
  box-shadow: 
    0 1px 3px rgba(0,0,0,0.12),
    0 1px 2px rgba(0,0,0,0.24);
}

.depth-card:hover {
  box-shadow: 
    0 14px 28px rgba(0,0,0,0.25),
    0 10px 10px rgba(0,0,0,0.22);
  transform: translateY(-2px);
}

/* Botones */
.button-examples {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
}

.btn {
  padding: 12px 24px;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0,0,0,0.2);
}

.primary { background: var(--primary); color: white; }
.secondary { background: var(--secondary); color: white; }
.danger { background: var(--danger); color: white; }
.submit { background: var(--primary); color: white; }

.btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0,0,0,0.3);
}

.btn:active {
  transform: translateY(0);
  box-shadow: 0 1px 2px rgba(0,0,0,0.2);
}

.primary:hover { box-shadow: 0 4px 12px rgba(52,152,219,0.4); }
.secondary:hover { box-shadow: 0 4px 12px rgba(46,204,113,0.4); }
.danger:hover { box-shadow: 0 4px 12px rgba(231,76,60,0.4); }

.animated {
  background: linear-gradient(45deg, var(--primary), var(--secondary));
  background-size: 200% 200%;
  animation: gradientShift 2s ease infinite;
}

@keyframes gradientShift {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* Formularios */
.form-example {
  max-width: 500px;
  margin: 0 auto;
  background: white;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 4px 15px var(--shadow);
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: 600;
  color: var(--dark);
}

.input {
  width: 100%;
  padding: 12px;
  border: 2px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.input:focus {
  outline: none;
  border-color: var(--primary);
  box-shadow: 0 0 0 3px rgba(52,152,219,0.1);
}

/* Efectos avanzados */
.effect-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
}

.animated-border {
  position: relative;
  padding: 30px;
  background: white;
  border-radius: 8px;
  text-align: center;
}

.animated-border::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(45deg, #ff0000, #00ff00, #0000ff, #ff0000);
  background-size: 400% 400%;
  border-radius: 10px;
  z-index: -1;
  animation: borderRotate 3s linear infinite;
}

@keyframes borderRotate {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.glass-effect {
  padding: 30px;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  text-align: center;
  box-shadow: 0 8px 32px rgba(0,0,0,0.1);
}

.neon-effect {
  padding: 30px;
  background: #000;
  border: 2px solid #00ff00;
  border-radius: 8px;
  text-align: center;
  color: #00ff00;
  text-shadow: 0 0 5px #00ff00, 0 0 10px #00ff00, 0 0 15px #00ff00;
  box-shadow: 0 0 5px #00ff00, 0 0 10px #00ff00, 0 0 15px #00ff00;
}

/* Responsive */
@media (max-width: 768px) {
  section {
    padding: 40px 15px;
  }
  
  .border-examples, .radius-examples, .shadow-examples, 
  .text-examples, .card-grid, .effect-examples {
    grid-template-columns: 1fr;
  }
  
  .button-examples {
    flex-direction: column;
    align-items: center;
  }
  
  header h1 {
    font-size: 2rem;
  }
  
  h2 {
    font-size: 2rem;
  }
}
```

## Resumen

Los bordes y sombras son herramientas poderosas para el diseño web:

- ✅ **Bordes**: Ancho, estilo, color y formas redondeadas
- ✅ **Sombras de caja**: Efectos de profundidad y realismo
- ✅ **Sombras de texto**: Efectos visuales en tipografía
- ✅ **Bordes con imágenes**: Diseños personalizados
- ✅ **Efectos avanzados**: Animaciones, glassmorphism, neón

Experimenta con combinaciones para crear interfaces modernas y atractivas. Las sombras pueden agregar profundidad y los bordes pueden definir jerarquías visuales.