# Módulo 05 - Texto y tipografía

En este módulo aprenderás a controlar completamente la apariencia del texto en tus páginas web, desde fuentes hasta espaciado y efectos visuales.

## Propiedades básicas de fuente

### Familia de fuente (font-family)

```css
body {
  font-family: Arial, sans-serif; /* Fuente principal con fallback */
}

h1 {
  font-family: 'Georgia', serif; /* Fuente serif para títulos */
}

code {
  font-family: 'Courier New', monospace; /* Fuente monoespaciada */
}
```

**Familias genéricas:**
- `serif`: Con serifs (ej: Times New Roman, Georgia)
- `sans-serif`: Sin serifs (ej: Arial, Helvetica)
- `monospace`: Ancho fijo (ej: Courier, Consolas)
- `cursive`: Manuscrita (ej: Comic Sans)
- `fantasy`: Decorativa (ej: Impact)

### Tamaño de fuente (font-size)

```css
.element {
  font-size: 16px;    /* Tamaño absoluto */
  font-size: 1.2em;   /* Relativo al elemento padre */
  font-size: 1.5rem;  /* Relativo al root */
  font-size: 120%;    /* Porcentaje */
  font-size: larger;  /* Palabra clave */
  font-size: 2vw;     /* Relativo al viewport */
}
```

**Palabras clave:**
- `xx-small`, `x-small`, `small`, `medium`, `large`, `x-large`, `xx-large`
- `larger`, `smaller`

### Peso de fuente (font-weight)

```css
.text {
  font-weight: normal;    /* 400 */
  font-weight: bold;      /* 700 */
  font-weight: lighter;   /* Más ligero que el heredado */
  font-weight: bolder;    /* Más pesado que el heredado */
  font-weight: 100;       /* Thin */
  font-weight: 200;       /* Extra Light */
  font-weight: 300;       /* Light */
  font-weight: 400;       /* Normal */
  font-weight: 500;       /* Medium */
  font-weight: 600;       /* Semi Bold */
  font-weight: 700;       /* Bold */
  font-weight: 800;       /* Extra Bold */
  font-weight: 900;       /* Black */
}
```

### Estilo de fuente (font-style)

```css
.text {
  font-style: normal;   /* Normal */
  font-style: italic;   /* Cursiva */
  font-style: oblique;  /* Oblícua (forzada) */
}
```

### Variante de fuente (font-variant)

```css
.text {
  font-variant: normal;      /* Normal */
  font-variant: small-caps;  /* Versalitas */
}
```

## Propiedad shorthand font

```css
/* Orden: style variant weight size/line-height family */
.element {
  font: italic small-caps bold 16px/1.5 Arial, sans-serif;
  font: 24px Georgia, serif;
  font: bold 14px/1.2 'Helvetica Neue', sans-serif;
}
```

## Propiedades de texto

### Color del texto (color)

```css
.text {
  color: #333;                    /* Hex */
  color: rgb(51, 51, 51);         /* RGB */
  color: rgba(51, 51, 51, 0.8);   /* RGBA */
  color: hsl(0, 0%, 20%);         /* HSL */
  color: currentColor;            /* Color heredado */
}
```

### Alineación de texto (text-align)

```css
.text {
  text-align: left;      /* Izquierda (default) */
  text-align: right;     /* Derecha */
  text-align: center;    /* Centro */
  text-align: justify;   /* Justificado */
}
```

### Decoración de texto (text-decoration)

```css
.link {
  text-decoration: none;           /* Sin decoración */
  text-decoration: underline;      /* Subrayado */
  text-decoration: overline;       /* Sobrerayado */
  text-decoration: line-through;   /* Tachado */
}

.fancy {
  text-decoration: underline wavy red; /* Estilo moderno */
}
```

### Transformación de texto (text-transform)

```css
.text {
  text-transform: none;        /* Sin transformación */
  text-transform: uppercase;   /* MAYÚSCULAS */
  text-transform: lowercase;   /* minúsculas */
  text-transform: capitalize;  /* Primera letra mayúscula */
}
```

### Sombra de texto (text-shadow)

```css
.text {
  text-shadow: 2px 2px 4px rgba(0,0,0,0.5);    /* horizontal vertical blur color */
  text-shadow: 0 0 10px #ff0000;                /* Efecto neón */
  text-shadow: 1px 1px 0 #000, -1px -1px 0 #000; /* Borde */
}
```

## Propiedades de espaciado

### Alto de línea (line-height)

```css
.text {
  line-height: 1.5;      /* Multiplicador (recomendado) */
  line-height: 24px;     /* Tamaño absoluto */
  line-height: 150%;     /* Porcentaje */
  line-height: normal;   /* Default (~1.2) */
}
```

### Espaciado entre letras (letter-spacing)

```css
.text {
  letter-spacing: normal;   /* Default */
  letter-spacing: 2px;      /* Más espacio */
  letter-spacing: -1px;     /* Menos espacio */
}
```

### Espaciado entre palabras (word-spacing)

```css
.text {
  word-spacing: normal;   /* Default */
  word-spacing: 5px;      /* Más espacio */
  word-spacing: -2px;     /* Menos espacio */
}
```

## Propiedades avanzadas

### Dirección del texto (direction)

```css
.arabic {
  direction: rtl;  /* Right to Left */
}

.english {
  direction: ltr;  /* Left to Right (default) */
}
```

### Orientación vertical (writing-mode)

```css
.vertical {
  writing-mode: vertical-rl;  /* Vertical right-to-left */
  writing-mode: vertical-lr;  /* Vertical left-to-right */
}
```

### Sangría de primera línea (text-indent)

```css
.paragraph {
  text-indent: 20px;     /* Sangría fija */
  text-indent: 2em;      /* Sangría relativa */
  text-indent: -10px;    /* Sangría negativa */
}
```

### Desbordamiento de texto (text-overflow)

```css
.truncate {
  white-space: nowrap;       /* Evita saltos de línea */
  overflow: hidden;          /* Oculta desbordamiento */
  text-overflow: ellipsis;   /* Añade "..." */
  text-overflow: clip;       /* Corta el texto */
}
```

### Múltiples líneas con ellipsis

```css
.multi-line {
  display: -webkit-box;
  -webkit-line-clamp: 3;     /* Número de líneas */
  -webkit-box-orient: vertical;
  overflow: hidden;
}
```

## Fuentes web (@font-face)

### Importar fuentes personalizadas

```css
@font-face {
  font-family: 'MiFuente';
  src: url('fuente.woff2') format('woff2'),
       url('fuente.woff') format('woff');
  font-weight: normal;
  font-style: normal;
}

.element {
  font-family: 'MiFuente', sans-serif;
}
```

### Fuentes desde Google Fonts

```html
<link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap" rel="stylesheet">
```

```css
.element {
  font-family: 'Roboto', sans-serif;
}
```

## Propiedades de lista

### Tipo de marcador (list-style-type)

```css
ul {
  list-style-type: disc;      /* • */
  list-style-type: circle;    /* ◦ */
  list-style-type: square;    /* ■ */
  list-style-type: none;      /* Sin marcador */
}

ol {
  list-style-type: decimal;   /* 1, 2, 3 */
  list-style-type: lower-alpha; /* a, b, c */
  list-style-type: upper-roman; /* I, II, III */
}
```

### Posición del marcador (list-style-position)

```css
ul {
  list-style-position: outside; /* Fuera del flujo (default) */
  list-style-position: inside;  /* Dentro del flujo */
}
```

### Imagen como marcador (list-style-image)

```css
ul {
  list-style-image: url('checkmark.png');
}
```

### Shorthand list-style

```css
ul {
  list-style: disc outside url('bullet.png');
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Tipografía CSS</title>
  <link rel="stylesheet" href="styles.css">
  <link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;700&family=Roboto:wght@300;400;500;700&display=swap" rel="stylesheet">
</head>
<body>
  <div class="container">
    <header>
      <h1>Maestría en Tipografía CSS</h1>
      <p class="subtitle">Domina el arte del texto en la web</p>
    </header>

    <main>
      <section class="typography-basics">
        <h2>Familias de Fuente</h2>
        
        <div class="font-examples">
          <p class="serif">Esta es una fuente serif. Times New Roman es un ejemplo clásico.</p>
          <p class="sans-serif">Esta es una fuente sans-serif. Arial y Helvetica son muy comunes.</p>
          <p class="monospace">Esta es una fuente monospace. Perfecta para código.</p>
        </div>
      </section>

      <section class="text-properties">
        <h2>Propiedades de Texto</h2>
        
        <p class="normal">Texto normal sin modificaciones.</p>
        <p class="uppercase">Texto en mayúsculas.</p>
        <p class="lowercase">Texto en minúsculas.</p>
        <p class="capitalize">Texto con capitalización.</p>
        
        <p class="underline">Texto subrayado.</p>
        <p class="strikethrough">Texto tachado.</p>
        <p class="shadow">Texto con sombra.</p>
      </section>

      <section class="spacing">
        <h2>Espaciado y Alineación</h2>
        
        <p class="left">Texto alineado a la izquierda. Este es el comportamiento por defecto en la mayoría de los idiomas.</p>
        
        <p class="center">Texto centrado. Comúnmente usado en títulos y encabezados.</p>
        
        <p class="right">Texto alineado a la derecha. Menos común pero útil en algunos diseños.</p>
        
        <p class="justify">Texto justificado. Ambas márgenes están alineadas. Esto crea un aspecto más formal y periódico, pero puede crear espacios irregulares entre palabras.</p>
      </section>

      <section class="advanced">
        <h2>Técnicas Avanzadas</h2>
        
        <div class="text-effects">
          <p class="gradient-text">Texto con gradiente</p>
          <p class="outline-text">Texto con contorno</p>
          <p class="emboss">Texto en relieve</p>
        </div>
        
        <div class="spacing-examples">
          <p class="tight">Texto con espaciado ajustado.</p>
          <p class="loose">Texto con espaciado amplio.</p>
          <p class="indented">Texto con sangría de primera línea. Esta técnica se usa comúnmente en párrafos largos para mejorar la legibilidad y crear una jerarquía visual clara.</p>
        </div>
      </section>

      <section class="lists">
        <h2>Listas Estilizadas</h2>
        
        <h3>Lista desordenada</h3>
        <ul class="custom-bullets">
          <li>Elemento de lista personalizado</li>
          <li>Otro elemento con estilo único</li>
          <li>Más elementos para demostrar</li>
        </ul>
        
        <h3>Lista ordenada</h3>
        <ol class="roman-numerals">
          <li>Primer elemento</li>
          <li>Segundo elemento</li>
          <li>Tercer elemento</li>
        </ol>
      </section>

      <section class="responsive-text">
        <h2>Texto Responsivo</h2>
        <p class="responsive">Este texto cambia de tamaño según el ancho de la pantalla. Prueba redimensionando la ventana del navegador.</p>
      </section>
    </main>
  </div>
</body>
</html>
```

```css
/* Importar fuentes locales */
@font-face {
  font-family: 'CustomFont';
  src: url('custom-font.woff2') format('woff2');
  font-weight: normal;
  font-style: normal;
}

/* Reset y variables */
:root {
  --primary-color: #2c3e50;
  --secondary-color: #3498db;
  --text-color: #333;
  --background: #f8f9fa;
  --serif: 'Playfair Display', serif;
  --sans: 'Roboto', sans-serif;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: var(--sans);
  line-height: 1.6;
  color: var(--text-color);
  background: var(--background);
  font-size: 16px;
}

.container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

/* Header */
header {
  text-align: center;
  margin-bottom: 50px;
  padding: 40px 20px;
  background: linear-gradient(135deg, var(--primary-color), var(--secondary-color));
  color: white;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.1);
}

h1 {
  font-family: var(--serif);
  font-size: clamp(2rem, 5vw, 3.5rem);
  font-weight: 700;
  margin-bottom: 10px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
}

.subtitle {
  font-size: 1.2rem;
  opacity: 0.9;
  font-weight: 300;
}

/* Secciones */
section {
  margin-bottom: 50px;
  padding: 30px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
}

h2 {
  font-family: var(--serif);
  font-size: 2rem;
  color: var(--primary-color);
  margin-bottom: 25px;
  border-bottom: 3px solid var(--secondary-color);
  padding-bottom: 10px;
}

h3 {
  font-size: 1.5rem;
  color: var(--primary-color);
  margin-bottom: 15px;
}

/* Familias de fuente */
.font-examples p {
  margin-bottom: 20px;
  padding: 15px;
  border-radius: 6px;
}

.serif {
  font-family: var(--serif);
  background: #f8f8f8;
  border-left: 4px solid #e74c3c;
}

.sans-serif {
  font-family: var(--sans);
  background: #f0f8ff;
  border-left: 4px solid #3498db;
}

.monospace {
  font-family: 'Courier New', monospace;
  background: #f0fff0;
  border-left: 4px solid #27ae60;
}

/* Propiedades de texto */
.text-properties p {
  margin-bottom: 15px;
  padding: 10px;
  border-radius: 4px;
}

.normal {
  background: #f8f8f8;
}

.uppercase {
  text-transform: uppercase;
  background: #fff5cc;
  font-weight: bold;
}

.lowercase {
  text-transform: lowercase;
  background: #e8f5e8;
}

.capitalize {
  text-transform: capitalize;
  background: #e3f2fd;
}

.underline {
  text-decoration: underline wavy var(--secondary-color);
  background: #f3e5f5;
}

.strikethrough {
  text-decoration: line-through;
  background: #ffebee;
}

.shadow {
  text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
  background: #f5f5f5;
  font-weight: bold;
}

/* Alineación */
.left {
  text-align: left;
  background: #e8f5e8;
}

.center {
  text-align: center;
  background: #fff3e0;
}

.right {
  text-align: right;
  background: #fce4ec;
}

.justify {
  text-align: justify;
  background: #f3e5f5;
  text-justify: inter-word;
}

/* Efectos avanzados */
.text-effects {
  margin-bottom: 30px;
}

.gradient-text {
  background: linear-gradient(45deg, #ff6b6b, #4ecdc4, #45b7d1);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  font-size: 2rem;
  font-weight: bold;
  text-align: center;
}

.outline-text {
  -webkit-text-stroke: 2px #3498db;
  -webkit-text-fill-color: white;
  font-size: 2rem;
  font-weight: bold;
  text-align: center;
  background: var(--primary-color);
  padding: 10px;
  border-radius: 8px;
}

.emboss {
  text-shadow: 1px 1px 0 #ccc, -1px -1px 0 #333;
  font-size: 1.5rem;
  font-weight: bold;
  color: #888;
}

/* Espaciado */
.spacing-examples p {
  margin-bottom: 20px;
  padding: 15px;
  border-radius: 6px;
}

.tight {
  letter-spacing: -1px;
  word-spacing: -2px;
  background: #fff8e1;
}

.loose {
  letter-spacing: 2px;
  word-spacing: 5px;
  background: #f1f8e9;
}

.indented {
  text-indent: 2em;
  background: #e0f2f1;
  line-height: 1.8;
}

/* Listas */
.custom-bullets {
  list-style: none;
  padding-left: 0;
}

.custom-bullets li {
  position: relative;
  padding-left: 30px;
  margin-bottom: 10px;
  padding: 10px 0 10px 35px;
}

.custom-bullets li::before {
  content: '✓';
  position: absolute;
  left: 0;
  top: 10px;
  color: #27ae60;
  font-weight: bold;
  font-size: 1.2em;
}

.roman-numerals {
  list-style-type: upper-roman;
  padding-left: 20px;
}

.roman-numerals li {
  margin-bottom: 8px;
  font-weight: 500;
}

/* Texto responsivo */
.responsive {
  font-size: clamp(1rem, 4vw, 2rem);
  background: linear-gradient(45deg, #667eea, #764ba2);
  color: white;
  padding: 20px;
  border-radius: 8px;
  text-align: center;
  font-weight: 500;
}

/* Animaciones sutiles */
section {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

section:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0,0,0,0.15);
}

/* Responsive */
@media (max-width: 768px) {
  .container {
    padding: 10px;
  }
  
  section {
    padding: 20px;
  }
  
  .font-examples p {
    font-size: 0.9rem;
  }
  
  .text-effects p {
    font-size: 1.5rem;
  }
}
```

## Resumen

La tipografía es crucial para una buena experiencia de usuario:

- ✅ **Fuentes**: Familia, tamaño, peso y estilo
- ✅ **Texto**: Color, alineación, transformación, decoración
- ✅ **Espaciado**: Líneas, letras y palabras
- ✅ **Efectos**: Sombras, gradientes, contornos
- ✅ **Listas**: Marcadores personalizados
- ✅ **Responsive**: Texto que se adapta

Una buena tipografía mejora la legibilidad, la jerarquía visual y la experiencia general del usuario.