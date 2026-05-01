# Módulo 03 - Colores, unidades y valores

En este módulo aprenderás todo sobre colores en CSS, las diferentes unidades de medida disponibles y cómo trabajar con valores numéricos y de texto.

## Sistema de colores en CSS

### 1. Nombres de colores
CSS incluye 147 nombres de colores predefinidos:

```css
h1 { color: red; }
p { color: blue; }
div { background-color: yellow; }
```

**Colores básicos:**
- `red`, `blue`, `green`, `yellow`, `purple`, `orange`
- `black`, `white`, `gray`, `silver`
- `maroon`, `navy`, `olive`, `lime`, `aqua`, `fuchsia`

### 2. Colores hexadecimales (#RRGGBB)
Sistema de 16.7 millones de colores:

```css
/* Formato completo */
h1 { color: #ff0000; } /* Rojo */
p { color: #00ff00; }   /* Verde */
div { color: #0000ff; } /* Azul */

/* Formato abreviado (cuando dígitos se repiten) */
h1 { color: #f00; }     /* Equivalente a #ff0000 */
p { color: #0f0; }       /* Equivalente a #00ff00 */
div { color: #00f; }     /* Equivalente a #0000ff */
```

### 3. Colores RGB/RGBA
Sistema aditivo con valores de 0-255:

```css
/* RGB - sin transparencia */
h1 { color: rgb(255, 0, 0); }     /* Rojo */
p { color: rgb(0, 255, 0); }       /* Verde */
div { color: rgb(0, 0, 255); }     /* Azul */
span { color: rgb(128, 128, 128); } /* Gris */

/* RGBA - con transparencia (alpha) */
.overlay {
  background-color: rgba(0, 0, 0, 0.5); /* Negro semi-transparente */
}
.modal {
  background-color: rgba(255, 255, 255, 0.9); /* Blanco casi transparente */
}
```

### 4. Colores HSL/HSLA
Sistema basado en matiz, saturación y luminosidad:

```css
/* HSL - Hue, Saturation, Lightness */
h1 { color: hsl(0, 100%, 50%); }     /* Rojo */
p { color: hsl(120, 100%, 50%); }     /* Verde */
div { color: hsl(240, 100%, 50%); }   /* Azul */
span { color: hsl(60, 100%, 50%); }   /* Amarillo */

/* HSLA - con transparencia */
.overlay {
  background-color: hsla(0, 100%, 0%, 0.7); /* Negro semi-transparente */
}
.accent {
  background-color: hsla(45, 100%, 50%, 0.8); /* Amarillo semi-transparente */
}
```

**Explicación HSL:**
- **Hue (Matiz)**: 0-360° (0°=rojo, 120°=verde, 240°=azul)
- **Saturation (Saturación)**: 0-100% (0%=gris, 100%=color puro)
- **Lightness (Luminosidad)**: 0-100% (0%=negro, 100%=blanco, 50%=color normal)

## Unidades de medida en CSS

### Unidades absolutas

```css
/* Píxeles - unidad más común */
div { width: 300px; height: 200px; }

/* Puntos - usado en impresión */
.print { font-size: 12pt; }

/* Pulgadas */
.ruler { width: 6in; }

/* Centímetros */
.card { width: 10cm; }

/* Milímetros */
.border { border-width: 5mm; }
```

### Unidades relativas al viewport

```css
/* 1% del ancho del viewport */
.full-width { width: 100vw; }

/* 1% del alto del viewport */
.full-height { height: 100vh; }

/* El menor entre vw o vh */
.small-screen { font-size: 5vmin; }

/* El mayor entre vw o vh */
.large-screen { font-size: 5vmax; }
```

### Unidades relativas a la fuente

```css
/* Relativo al elemento padre */
.child { font-size: 1.2em; }  /* 20% más grande que el padre */
.small { font-size: 0.8em; }  /* 20% más pequeño que el padre */

/* Relativo al elemento root (<html>) */
html { font-size: 16px; }
.title { font-size: 2rem; }    /* 32px */
.subtitle { font-size: 1.5rem; } /* 24px */
```

### Unidades relativas al contenedor

```css
/* Porcentaje del elemento padre */
.column { width: 50%; }
.full { width: 100%; }

/* Relativo al ancho del contenedor */
.flexible { width: 75%; padding: 5%; }
```

## Funciones de color modernas

### 1. Color con transparencia
```css
/* Usando alpha en diferentes formatos */
.overlay {
  background-color: #00000080; /* Hex con alpha */
  background-color: rgb(0 0 0 / 0.5); /* RGB moderno con / */
  background-color: hsl(0 0% 0% / 50%); /* HSL moderno */
}
```

### 2. Colores CSS modernos
```css
/* Colores con mejor legibilidad */
.accent { color: lch(50% 100 120); } /* LCH color space */
.srgb { color: color(srgb 1 0.5 0); } /* Color spaces modernos */
```

## Propiedades que aceptan valores de color

```css
.element {
  /* Texto */
  color: #333;
  
  /* Fondos */
  background-color: rgba(255, 255, 255, 0.9);
  
  /* Bordes */
  border-color: hsl(200, 50%, 40%);
  border-top-color: #ff0000;
  
  /* Sombras */
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
  
  /* Gradientes (veremos en módulos posteriores) */
  background: linear-gradient(to right, #ff0000, #0000ff);
}
```

## Sistema de valores numéricos

### Números enteros y decimales
```css
div {
  z-index: 10;        /* Entero */
  opacity: 0.8;       /* Decimal */
  font-weight: 400;   /* Entero */
  line-height: 1.5;   /* Decimal */
}
```

### Números con unidades
```css
.dimensions {
  width: 300px;
  height: 200px;
  margin: 20px;
  padding: 15px;
  border-width: 2px;
  font-size: 16px;
  line-height: 1.5;
}
```

### Valores calculados
```css
.calculated {
  width: calc(100% - 40px);    /* 100% menos 40px */
  height: calc(100vh - 80px);  /* 100vh menos 80px */
  font-size: calc(1rem + 2px); /* 1rem más 2px */
}
```

## Valores de texto y palabras clave

### Palabras clave predefinidas
```css
.position {
  position: static;    /* Por defecto */
  position: relative;  /* Relativo al flujo normal */
  position: absolute;  /* Fuera del flujo */
  position: fixed;     /* Fijo en viewport */
  position: sticky;    /* Pegajoso */
}

.display {
  display: block;      /* Elemento de bloque */
  display: inline;     /* Elemento en línea */
  display: inline-block; /* Ambos */
  display: flex;       /* Contenedor flex */
  display: grid;       /* Contenedor grid */
  display: none;       /* Oculto */
}

.float {
  float: left;         /* Flotar a la izquierda */
  float: right;        /* Flotar a la derecha */
  float: none;         /* Sin flotar */
}
```

### Cadenas de texto
```css
.content::before {
  content: "→ ";       /* Texto literal */
  content: attr(data-icon); /* Valor de atributo */
  content: "Capítulo " counter(capitulo); /* Contador */
}
```

## Funciones CSS útiles

### Funciones matemáticas
```css
.element {
  width: min(500px, 80vw);     /* El menor valor */
  height: max(200px, 50vh);    /* El mayor valor */
  padding: clamp(10px, 5vw, 50px); /* Entre min y max */
}
```

### Funciones de color
```css
.colors {
  /* Mezclar colores */
  background-color: color-mix(in srgb, red 50%, blue 50%);
  
  /* Contraste automático */
  color: color-contrast(wheat vs tan, sienna, #d2691e, olive);
  
  /* Relativo a otro color */
  border-color: color-mix(in srgb, canvasText 20%, transparent);
}
```

## Propiedades shorthand vs longhand

### Shorthand (abreviado)
```css
.margin {
  margin: 10px 20px 15px 25px; /* top right bottom left */
  margin: 10px 20px 15px;       /* top right/left bottom */
  margin: 10px 20px;            /* top/bottom right/left */
  margin: 10px;                 /* todos los lados */
}

.padding {
  padding: 15px 20px; /* vertical horizontal */
}

.border {
  border: 2px solid #333; /* width style color */
}

.font {
  font: italic bold 16px/1.5 Arial, sans-serif; /* style weight size/line-height family */
}

.background {
  background: #f0f0f0 url('pattern.png') no-repeat center top; /* color image repeat position */
}
```

### Longhand (extendido)
```css
.margin {
  margin-top: 10px;
  margin-right: 20px;
  margin-bottom: 15px;
  margin-left: 25px;
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Colores y Unidades CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>Sistema de Colores CSS</h1>
      <p>Demostración de colores, unidades y valores</p>
    </header>

    <main class="main">
      <section class="color-section">
        <h2>Colores por Nombre</h2>
        <div class="color-box red">Rojo</div>
        <div class="color-box blue">Azul</div>
        <div class="color-box green">Verde</div>
      </section>

      <section class="color-section">
        <h2>Colores Hexadecimales</h2>
        <div class="color-box hex1">#FF6B6B</div>
        <div class="color-box hex2">#4ECDC4</div>
        <div class="color-box hex3">#45B7D1</div>
      </section>

      <section class="color-section">
        <h2>Colores RGB/RGBA</h2>
        <div class="color-box rgb1">RGB sólido</div>
        <div class="color-box rgb2">RGBA transparente</div>
      </section>

      <section class="units-section">
        <h2>Unidades de Medida</h2>
        <div class="unit-demo">
          <div class="box px">300px</div>
          <div class="box percent">50%</div>
          <div class="box vw">80vw</div>
          <div class="box vh">60vh</div>
        </div>
      </section>

      <section class="calc-section">
        <h2>Valores Calculados</h2>
        <div class="calc-demo">
          <div class="calc-box">Ancho calculado</div>
          <div class="calc-box">Responsive</div>
        </div>
      </section>
    </main>
  </div>
</body>
</html>
```

```css
/* Reset y base */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  line-height: 1.6;
  color: #333;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

/* Header */
.header {
  text-align: center;
  margin-bottom: 40px;
  color: white;
}

.header h1 {
  font-size: calc(2rem + 1vw);
  margin-bottom: 10px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
}

.header p {
  font-size: 1.2rem;
  opacity: 0.9;
}

/* Secciones */
.color-section, .units-section, .calc-section {
  background: rgba(255, 255, 255, 0.95);
  margin-bottom: 30px;
  padding: 25px;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.1);
}

.color-section h2, .units-section h2, .calc-section h2 {
  margin-bottom: 20px;
  color: #2c3e50;
  border-bottom: 3px solid #3498db;
  padding-bottom: 10px;
}

/* Color boxes */
.color-box {
  display: inline-block;
  width: calc(30% - 10px);
  margin: 5px;
  padding: 20px;
  color: white;
  text-align: center;
  border-radius: 8px;
  font-weight: bold;
  box-shadow: 0 4px 15px rgba(0,0,0,0.2);
  transition: transform 0.3s ease;
}

.color-box:hover {
  transform: translateY(-5px);
}

/* Colores específicos */
.red { background-color: red; }
.blue { background-color: blue; }
.green { background-color: green; }
.hex1 { background-color: #FF6B6B; }
.hex2 { background-color: #4ECDC4; }
.hex3 { background-color: #45B7D1; }
.rgb1 { background-color: rgb(155, 89, 182); }
.rgb2 { background-color: rgba(155, 89, 182, 0.7); }

/* Units demo */
.unit-demo {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
}

.box {
  padding: 20px;
  background: linear-gradient(45deg, #ff9a9e, #fecfef);
  color: white;
  border-radius: 8px;
  text-align: center;
  font-weight: bold;
  box-shadow: 0 4px 15px rgba(0,0,0,0.2);
  transition: all 0.3s ease;
}

.box:hover {
  transform: scale(1.05);
}

.px { width: 300px; }
.percent { width: 50%; min-width: 200px; }
.vw { width: 80vw; max-width: 400px; }
.vh { height: 60vh; width: 200px; }

/* Calc demo */
.calc-demo {
  display: flex;
  gap: 20px;
  justify-content: center;
}

.calc-box {
  width: calc(100% / 3 - 20px);
  min-width: 200px;
  height: calc(100px + 2vh);
  background: linear-gradient(45deg, #a8edea, #fed6e3);
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  font-weight: bold;
  color: #333;
  box-shadow: 0 4px 15px rgba(0,0,0,0.2);
}

/* Responsive */
@media (max-width: 768px) {
  .color-box {
    width: calc(45% - 10px);
  }
  
  .unit-demo {
    flex-direction: column;
    align-items: center;
  }
  
  .box {
    width: 250px;
  }
  
  .calc-box {
    width: calc(90vw - 40px);
  }
}
```

## Resumen

En este módulo aprendiste:

- ✅ **Colores**: Nombres, hexadecimal, RGB/RGBA, HSL/HSLA
- ✅ **Unidades**: Absolutas (px, pt), relativas (em, rem, %), viewport (vw, vh)
- ✅ **Valores**: Números, texto, palabras clave, funciones
- ✅ **Shorthand vs Longhand**: Propiedades abreviadas
- ✅ **Funciones modernas**: calc(), min(), max(), clamp(), color-mix()

Los colores y unidades son fundamentales en CSS. Practica combinándolos para crear diseños atractivos y responsivos.