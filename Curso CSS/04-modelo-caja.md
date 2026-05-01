# Módulo 04 - Modelo de caja (Box Model)

En este módulo aprenderás el concepto fundamental del modelo de caja en CSS, que explica cómo se dimensionan y espacian todos los elementos HTML.

## ¿Qué es el Box Model?

Todos los elementos HTML se representan como **rectángulos** (cajas) en el navegador. El modelo de caja describe cómo se calcula el espacio total que ocupa cada elemento.

### Componentes del Box Model

```
┌─────────────────────────────────────┐ ← Margin (margen exterior)
│  ┌─────────────────────────────────┐  │
│  │        Border (borde)          │  │
│  │  ┌─────────────────────────────┐ │  │
│  │  │      Padding (relleno)     │ │  │
│  │  │  ┌─────────────────────────┐ │ │  │
│  │  │  │    Content (contenido) │ │ │  │
│  │  │  │                         │ │ │  │
│  │  │  └─────────────────────────┘ │ │  │
│  │  └─────────────────────────────┘ │ │  │
│  └─────────────────────────────────┘  │
└─────────────────────────────────────┘
```

### Propiedades de cada capa

#### 1. Content (Contenido)
- **width/height**: Dimensiones del contenido
- Es el área donde se muestra el texto, imágenes, etc.

#### 2. Padding (Relleno)
- Espacio entre el contenido y el borde
- **padding**: Aplica a todos los lados
- **padding-top/right/bottom/left**: Lados individuales

#### 3. Border (Borde)
- Línea que rodea el padding
- **border-width**: Grosor del borde
- **border-style**: Estilo (solid, dashed, dotted, etc.)
- **border-color**: Color del borde

#### 4. Margin (Margen)
- Espacio exterior, entre elementos
- **margin**: Aplica a todos los lados
- **margin-top/right/bottom/left**: Lados individuales

## Propiedades del Box Model

### Dimensiones del contenido

```css
.box {
  width: 300px;   /* Ancho del contenido */
  height: 200px;  /* Alto del contenido */
  
  /* Dimensiones mínimas y máximas */
  min-width: 200px;
  max-width: 500px;
  min-height: 150px;
  max-height: 400px;
}
```

### Padding (relleno interior)

```css
.element {
  /* Padding para todos los lados */
  padding: 20px;
  
  /* Padding para lados específicos */
  padding-top: 10px;
  padding-right: 15px;
  padding-bottom: 10px;
  padding-left: 15px;
  
  /* Shorthand: top right bottom left */
  padding: 10px 15px 10px 15px;
  
  /* Shorthand: top/bottom right/left */
  padding: 10px 15px;
}
```

### Border (bordes)

```css
.border-example {
  /* Propiedades individuales */
  border-width: 2px;
  border-style: solid;
  border-color: #333;
  
  /* Shorthand completo */
  border: 2px solid #333;
  
  /* Bordes individuales */
  border-top: 1px dashed red;
  border-right: 2px solid blue;
  border-bottom: 1px dotted green;
  border-left: 2px double purple;
  
  /* Border radius para esquinas redondeadas */
  border-radius: 8px;
  border-radius: 50%; /* Círculo */
}
```

### Margin (margen exterior)

```css
.spacing {
  /* Margin para todos los lados */
  margin: 20px;
  
  /* Margin para lados específicos */
  margin-top: 10px;
  margin-right: 15px;
  margin-bottom: 10px;
  margin-left: 15px;
  
  /* Shorthand igual que padding */
  margin: 10px 15px; /* vertical horizontal */
  margin: 10px 15px 20px; /* top right/left bottom */
  margin: 10px 15px 20px 25px; /* top right bottom left */
  
  /* Margin especial */
  margin: 0 auto; /* Centrar horizontalmente */
}
```

## Box-sizing: El gran cambio

### Content-box (por defecto)
```css
.element {
  box-sizing: content-box; /* Valor por defecto */
  width: 300px;
  padding: 20px;
  border: 2px solid black;
  /* Ancho total = 300px + 40px + 4px = 344px */
}
```

### Border-box (recomendado)
```css
.element {
  box-sizing: border-box;
  width: 300px;
  padding: 20px;
  border: 2px solid black;
  /* Ancho total = 300px (padding y border incluidos) */
}
```

**Reset recomendado:**
```css
* {
  box-sizing: border-box;
}
```

## Propiedades avanzadas del Box Model

### Outline (contorno)
Similar al border pero no ocupa espacio:
```css
.button {
  outline: 2px solid blue;
  outline-offset: 5px; /* Espacio entre outline y border */
}
```

### Box-shadow (sombras)
```css
.card {
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1); /* horizontal vertical blur color */
  box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.1); /* sombra interior */
}
```

## Colapso de márgenes

Cuando dos elementos verticales tienen márgenes, estos se **colapsan**:

```html
<div class="top">Elemento superior</div>
<div class="bottom">Elemento inferior</div>
```

```css
.top {
  margin-bottom: 20px;
}

.bottom {
  margin-top: 30px;
}

/* El margen entre elementos será de 30px (el mayor), no 50px */
```

**Excepciones al colapso:**
- Elementos con `float`
- Elementos con `position: absolute`
- Elementos flex/grid
- Padding o border entre ellos

## Centrado de elementos

### Centrado horizontal
```css
.center-h {
  margin: 0 auto; /* Para elementos con ancho definido */
  text-align: center; /* Para texto/inline elements */
}
```

### Centrado vertical
```css
.center-v {
  /* Método 1: Flexbox */
  display: flex;
  align-items: center;
  justify-content: center;
  
  /* Método 2: Grid */
  display: grid;
  place-items: center;
  
  /* Método 3: Position absolute */
  position: relative;
}
.center-v > * {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
```

## Overflow (desbordamiento)

```css
.container {
  width: 300px;
  height: 200px;
  overflow: visible;   /* Muestra contenido que se sale (default) */
  overflow: hidden;    /* Oculta contenido que se sale */
  overflow: scroll;    /* Siempre muestra barras de scroll */
  overflow: auto;      /* Muestra scroll solo cuando es necesario */
  
  /* Control individual */
  overflow-x: auto;
  overflow-y: hidden;
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Box Model CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="container">
    <h1>Modelo de Caja en CSS</h1>
    
    <div class="box-model-demo">
      <div class="box content-box">
        <h3>Content Box</h3>
        <p>Ancho total: width + padding + border</p>
      </div>
      
      <div class="box border-box">
        <h3>Border Box</h3>
        <p>Ancho total: width (incluye padding y border)</p>
      </div>
    </div>

    <div class="spacing-examples">
      <h2>Ejemplos de Espaciado</h2>
      
      <div class="padding-example">
        <h3>Padding</h3>
        <p>Espacio interior entre contenido y borde</p>
      </div>
      
      <div class="margin-example">
        <h3>Margin</h3>
        <p>Espacio exterior entre elementos</p>
      </div>
      
      <div class="border-example">
        <h3>Border</h3>
        <p>Línea que rodea el elemento</p>
      </div>
    </div>

    <div class="advanced-examples">
      <h2>Ejemplos Avanzados</h2>
      
      <div class="shadow-box">
        <h3>Box Shadow</h3>
        <p>Sombras para dar profundidad</p>
      </div>
      
      <div class="rounded-box">
        <h3>Border Radius</h3>
        <p>Esquinas redondeadas</p>
      </div>
      
      <div class="outline-box">
        <h3>Outline</h3>
        <p>Contorno exterior</p>
      </div>
    </div>

    <div class="overflow-demo">
      <h2>Overflow</h2>
      <div class="overflow-container">
        <p>Este contenido es muy largo y se sale del contenedor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
      </div>
    </div>
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
  background: #f8f9fa;
  padding: 20px;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
}

h1, h2, h3 {
  color: #2c3e50;
  margin-bottom: 20px;
}

h1 {
  text-align: center;
  font-size: 2.5rem;
  margin-bottom: 40px;
}

/* Demo del Box Model */
.box-model-demo {
  display: flex;
  gap: 40px;
  margin-bottom: 60px;
  justify-content: center;
}

.box {
  width: 300px;
  min-height: 150px;
  padding: 20px;
  border: 3px solid #3498db;
  background: white;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
}

.content-box {
  box-sizing: content-box;
}

.border-box {
  box-sizing: border-box;
}

.box h3 {
  color: #e74c3c;
  margin-bottom: 10px;
}

/* Explicación visual del box model */
.box::before {
  content: '';
  position: absolute;
  top: -25px;
  left: 10px;
  background: #f39c12;
  color: white;
  padding: 2px 8px;
  border-radius: 3px;
  font-size: 12px;
  font-weight: bold;
}

.content-box::before {
  content: 'CONTENT-BOX';
}

.border-box::before {
  content: 'BORDER-BOX';
}

/* Ejemplos de espaciado */
.spacing-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
  margin-bottom: 60px;
}

.padding-example, .margin-example, .border-example {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
}

.padding-example {
  padding: 40px;
  background: linear-gradient(45deg, #f093fb 0%, #f5576c 100%);
  color: white;
}

.margin-example {
  margin: 30px;
  background: linear-gradient(45deg, #4facfe 0%, #00f2fe 100%);
  color: white;
}

.border-example {
  border: 8px solid;
  border-image: linear-gradient(45deg, #ff9a9e, #fecfef) 1;
  background: white;
}

/* Ejemplos avanzados */
.advanced-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
  margin-bottom: 60px;
}

.shadow-box {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.2), 0 1px 8px rgba(0,0,0,0.1);
}

.rounded-box {
  background: linear-gradient(45deg, #a8edea, #fed6e3);
  padding: 30px;
  border-radius: 50px 20px;
  color: #333;
}

.outline-box {
  background: white;
  padding: 30px;
  border-radius: 8px;
  outline: 3px solid #e74c3c;
  outline-offset: 5px;
}

/* Demo de overflow */
.overflow-demo {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
}

.overflow-container {
  width: 100%;
  height: 120px;
  padding: 15px;
  border: 2px solid #bdc3c7;
  border-radius: 4px;
  overflow-y: auto;
  background: #ecf0f1;
}

.overflow-container p {
  margin: 0;
  line-height: 1.4;
}

/* Responsive */
@media (max-width: 768px) {
  .box-model-demo {
    flex-direction: column;
    align-items: center;
  }
  
  .box {
    width: 100%;
    max-width: 400px;
  }
  
  .spacing-examples, .advanced-examples {
    grid-template-columns: 1fr;
  }
}

/* Animaciones sutiles */
.box, .padding-example, .margin-example, .border-example, 
.shadow-box, .rounded-box, .outline-box, .overflow-demo {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.box:hover, .padding-example:hover, .margin-example:hover, 
.border-example:hover, .shadow-box:hover, .rounded-box:hover, 
.outline-box:hover, .overflow-demo:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0,0,0,0.15);
}
```

## Resumen

El modelo de caja es fundamental en CSS:

- ✅ **Content**: Área del contenido real
- ✅ **Padding**: Espacio interior
- ✅ **Border**: Línea alrededor del padding
- ✅ **Margin**: Espacio exterior entre elementos
- ✅ **Box-sizing**: Controla cómo se calcula el ancho total
- ✅ **Overflow**: Maneja contenido que se sale
- ✅ **Box-shadow**: Añade profundidad visual

Dominar el box model te permite controlar completamente el layout y espaciado de tus diseños web.