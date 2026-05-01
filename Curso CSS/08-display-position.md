# Módulo 08 - Display y position

En este módulo aprenderás los conceptos fundamentales de flujo de documento, posicionamiento y display que controlan cómo se comportan los elementos en la página web.

## El modelo de caja y flujo normal

### Flujo normal del documento

```css
/* Los elementos siguen el flujo normal por defecto */
.element {
  /* display: block; */  /* Valor por defecto para div, p, h1, etc. */
}

/* Elementos inline */
span, a, strong {
  /* display: inline; */ /* Valor por defecto */
}
```

### Comportamiento de elementos block vs inline

```css
/* Elementos block */
div, p, h1, section {
  display: block;        /* Ocupan todo el ancho disponible */
  width: 100%;          /* Por defecto */
  margin: auto;         /* Centrado horizontal */
}

/* Elementos inline */
span, a, em, strong {
  display: inline;      /* Solo ocupan el espacio necesario */
  width: auto;          /* No se puede cambiar */
  margin: 0;            /* Solo left/right */
}
```

## Propiedad display

### Valores principales

```css
.element {
  display: block;        /* Elemento de bloque */
  display: inline;       /* Elemento en línea */
  display: inline-block; /* Mezcla de ambos */
  display: none;         /* Oculta el elemento */
  display: flex;         /* Contenedor flexbox */
  display: grid;         /* Contenedor grid */
  display: table;        /* Comportamiento de tabla */
}
```

### Inline-block

```css
.inline-block-element {
  display: inline-block;
  width: 200px;         /* Ahora sí funciona */
  height: 100px;        /* Ahora sí funciona */
  margin: 10px;         /* Funciona en todas direcciones */
  vertical-align: top;  /* Alineación vertical */
}
```

### Display: none vs visibility: hidden

```css
.hidden-with-display {
  display: none;        /* Elimina completamente del flujo */
}

.hidden-with-visibility {
  visibility: hidden;   /* Mantiene el espacio, solo oculta visualmente */
}
```

## Posicionamiento (position)

### Static (posicionamiento por defecto)

```css
.element {
  position: static;     /* Valor por defecto */
  /* No responde a top, right, bottom, left */
}
```

### Relative

```css
.relative-element {
  position: relative;
  top: 20px;           /* Se mueve desde su posición original */
  left: 30px;
  /* Mantiene su espacio original en el flujo */
}
```

### Absolute

```css
.absolute-element {
  position: absolute;
  top: 50px;
  right: 20px;
  /* Se posiciona relativo al primer ancestro posicionado */
  /* Sale completamente del flujo normal */
}
```

### Fixed

```css
.fixed-element {
  position: fixed;
  top: 0;
  left: 0;
  /* Se posiciona relativo a la ventana del navegador */
  /* Permanece fijo durante el scroll */
}
```

### Sticky

```css
.sticky-element {
  position: sticky;
  top: 0;
  /* Se comporta como relative hasta que llega al punto especificado */
  /* Luego se comporta como fixed */
}
```

## Z-index y contexto de apilamiento

### Z-index básico

```css
.layer-1 {
  position: relative;
  z-index: 1;
}

.layer-2 {
  position: relative;
  z-index: 2;    /* Aparece encima de layer-1 */
}

.layer-3 {
  position: absolute;
  z-index: 999;  /* Aparece encima de todos */
}
```

### Contexto de apilamiento

```css
.parent {
  position: relative;
  z-index: 1;    /* Crea un nuevo contexto de apilamiento */
}

.child {
  position: relative;
  z-index: 100;  /* Solo afecta dentro del contexto del parent */
}
```

## Técnicas de centrado

### Centrado horizontal

```css
/* Método 1: Auto margins */
.centered-block {
  width: 300px;
  margin: 0 auto;      /* Solo funciona en elementos block */
}

/* Método 2: Text-align */
.centered-inline {
  text-align: center;
}

/* Método 3: Flexbox */
.flex-center {
  display: flex;
  justify-content: center;
}
```

### Centrado vertical

```css
/* Método 1: Line-height (para texto) */
.single-line-text {
  height: 100px;
  line-height: 100px;   /* Altura = line-height */
}

/* Método 2: Flexbox */
.flex-center-vertical {
  display: flex;
  align-items: center;
  height: 200px;
}

/* Método 3: Absolute positioning */
.absolute-center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
```

### Centrado perfecto (horizontal + vertical)

```css
/* Método 1: Flexbox */
.perfect-center {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 300px;
}

/* Método 2: Grid */
.grid-center {
  display: grid;
  place-items: center;
  height: 300px;
}

/* Método 3: Absolute + Transform */
.absolute-perfect-center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
```

## Float y clear

### Float básico

```css
.float-left {
  float: left;
  width: 200px;
}

.float-right {
  float: right;
  width: 200px;
}
```

### Clear

```css
.clearfix::after {
  content: '';
  display: table;
  clear: both;
}

.clear-both {
  clear: both;    /* No permite floats a izquierda ni derecha */
}

.clear-left {
  clear: left;    /* No permite floats a la izquierda */
}

.clear-right {
  clear: right;   /* No permite floats a la derecha */
}
```

## Técnicas de layout tradicionales

### Layout de 2 columnas con float

```css
.container {
  width: 100%;
  overflow: hidden; /* Contiene los floats */
}

.sidebar {
  float: left;
  width: 30%;
}

.content {
  float: right;
  width: 65%;
}
```

### Layout de 3 columnas

```css
.header, .footer {
  clear: both;
}

.left-sidebar {
  float: left;
  width: 20%;
}

.content {
  float: left;
  width: 60%;
}

.right-sidebar {
  float: right;
  width: 20%;
}
```

## Modern CSS Layout

### Flexbox para layouts simples

```css
.flex-layout {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  flex: 0 0 250px;  /* No crece, no se encoge, ancho fijo */
}

.content {
  flex: 1;           /* Crece para ocupar espacio restante */
}
```

### CSS Grid para layouts complejos

```css
.grid-layout {
  display: grid;
  grid-template-columns: 250px 1fr 200px;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
  grid-template-areas: 
    "header header header"
    "sidebar content aside"
    "footer footer footer";
}

.header { grid-area: header; }
.sidebar { grid-area: sidebar; }
.content { grid-area: content; }
.aside { grid-area: aside; }
.footer { grid-area: footer; }
```

## Técnicas de posicionamiento avanzado

### Posicionamiento absoluto dentro de contenedor

```css
.relative-container {
  position: relative;
  height: 300px;
  background: #f0f0f0;
}

.absolute-child {
  position: absolute;
  top: 20px;
  right: 20px;
  width: 100px;
  height: 100px;
  background: #3498db;
}
```

### Elementos superpuestos

```css
.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  z-index: 1000;
}

.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 1001;
  background: white;
  padding: 20px;
  border-radius: 8px;
}
```

### Navegación sticky

```css
.sticky-nav {
  position: sticky;
  top: 0;
  background: white;
  z-index: 100;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
```

## Responsive positioning

### Cambiar posicionamiento según pantalla

```css
.responsive-element {
  position: static;  /* Por defecto */
}

@media (min-width: 768px) {
  .responsive-element {
    position: fixed;
    top: 20px;
    right: 20px;
  }
}
```

### Layout responsive con flexbox

```css
.responsive-layout {
  display: flex;
  flex-direction: column;  /* Columnas en móvil */
}

@media (min-width: 768px) {
  .responsive-layout {
    flex-direction: row;   /* Filas en desktop */
  }
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Display y Position CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <nav class="sticky-nav">
    <div class="nav-container">
      <h1>Display & Position</h1>
      <ul>
        <li><a href="#display">Display</a></li>
        <li><a href="#position">Position</a></li>
        <li><a href="#centering">Centrado</a></li>
        <li><a href="#layouts">Layouts</a></li>
      </ul>
    </div>
  </nav>

  <main class="main-content">
    <section id="display" class="section">
      <h2>Propiedad Display</h2>
      
      <div class="display-examples">
        <div class="example">
          <h3>Block</h3>
          <div class="block-element">Elemento block</div>
          <div class="block-element">Otro elemento block</div>
        </div>
        
        <div class="example">
          <h3>Inline</h3>
          <p>Este es un párrafo con <span class="inline-element">elemento inline</span> dentro de él.</p>
        </div>
        
        <div class="example">
          <h3>Inline-block</h3>
          <div class="inline-block-container">
            <div class="inline-block-element">Uno</div>
            <div class="inline-block-element">Dos</div>
            <div class="inline-block-element">Tres</div>
          </div>
        </div>
      </div>
    </section>

    <section id="position" class="section">
      <h2>Posicionamiento</h2>
      
      <div class="position-examples">
        <div class="position-container">
          <div class="static-box">Static</div>
          <div class="relative-box">Relative (+20px, +20px)</div>
          <div class="absolute-box">Absolute</div>
        </div>
        
        <div class="fixed-example">
          <p>Scroll para ver el elemento fixed...</p>
          <div class="fixed-box">Fixed</div>
        </div>
        
        <div class="sticky-example">
          <div class="sticky-content">
            <p>Contenido que hace scroll...</p>
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
            <div class="sticky-box">Sticky</div>
            <p>Más contenido después del elemento sticky...</p>
          </div>
        </div>
      </div>
    </section>

    <section id="centering" class="section">
      <h2>Técnicas de Centrado</h2>
      
      <div class="centering-examples">
        <div class="centering-method">
          <h3>Margin Auto</h3>
          <div class="centered-block">Centrado con margin: 0 auto</div>
        </div>
        
        <div class="centering-method">
          <h3>Flexbox</h3>
          <div class="flex-center">Centrado con Flexbox</div>
        </div>
        
        <div class="centering-method">
          <h3>Absolute + Transform</h3>
          <div class="absolute-center-container">
            <div class="absolute-centered">Centrado absoluto</div>
          </div>
        </div>
        
        <div class="centering-method">
          <h3>Grid</h3>
          <div class="grid-center">Centrado con Grid</div>
        </div>
      </div>
    </section>

    <section id="layouts" class="section">
      <h2>Layouts Prácticos</h2>
      
      <div class="layout-examples">
        <div class="layout-example">
          <h3>Float Layout</h3>
          <div class="float-layout">
            <div class="sidebar">Sidebar</div>
            <div class="main-content">Contenido principal</div>
          </div>
        </div>
        
        <div class="layout-example">
          <h3>Flexbox Layout</h3>
          <div class="flex-layout">
            <div class="sidebar">Sidebar</div>
            <div class="main-content">Contenido principal</div>
          </div>
        </div>
        
        <div class="layout-example">
          <h3>Grid Layout</h3>
          <div class="grid-layout">
            <div class="sidebar">Sidebar</div>
            <div class="main-content">Contenido principal</div>
          </div>
        </div>
      </div>
    </section>

    <section class="modal-section">
      <h2>Modal con Posicionamiento</h2>
      <button class="open-modal-btn">Abrir Modal</button>
      
      <div class="modal-overlay" id="modal">
        <div class="modal-content">
          <h3>Modal</h3>
          <p>Este modal usa posicionamiento absoluto para centrarse en la pantalla.</p>
          <button class="close-modal-btn">Cerrar</button>
        </div>
      </div>
    </section>
  </main>

  <footer class="footer">
    <p>&copy; 2024 Display & Position CSS</p>
  </footer>

  <script src="script.js"></script>
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
  --gray: #95a5a6;
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

/* Navegación sticky */
.sticky-nav {
  position: sticky;
  top: 0;
  background: white;
  z-index: 1000;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  padding: 1rem 0;
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-container h1 {
  color: var(--primary);
}

nav ul {
  display: flex;
  list-style: none;
  gap: 2rem;
}

nav a {
  text-decoration: none;
  color: var(--dark);
  font-weight: 500;
  transition: color 0.3s ease;
}

nav a:hover {
  color: var(--primary);
}

/* Secciones */
.section {
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

/* Display examples */
.display-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 40px;
  margin-bottom: 40px;
}

.example {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.example h3 {
  margin-bottom: 15px;
  color: var(--primary);
}

.block-element {
  background: var(--primary);
  color: white;
  padding: 10px;
  margin-bottom: 10px;
  border-radius: 4px;
}

.inline-element {
  background: var(--secondary);
  color: white;
  padding: 5px 10px;
  border-radius: 4px;
}

.inline-block-container {
  background: #f8f9fa;
  padding: 10px;
  border-radius: 4px;
}

.inline-block-element {
  display: inline-block;
  background: var(--accent);
  color: white;
  padding: 10px 15px;
  margin: 5px;
  border-radius: 4px;
}

/* Position examples */
.position-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 40px;
}

.position-container {
  position: relative;
  height: 250px;
  background: #f8f9fa;
  border: 2px dashed var(--gray);
  border-radius: 8px;
  padding: 20px;
}

.static-box {
  background: var(--primary);
  color: white;
  padding: 10px;
  border-radius: 4px;
  margin-bottom: 10px;
}

.relative-box {
  position: relative;
  top: 20px;
  left: 20px;
  background: var(--secondary);
  color: white;
  padding: 10px;
  border-radius: 4px;
}

.absolute-box {
  position: absolute;
  top: 20px;
  right: 20px;
  background: var(--accent);
  color: white;
  padding: 10px;
  border-radius: 4px;
}

.fixed-example {
  height: 200px;
  background: #f8f9fa;
  border-radius: 8px;
  padding: 20px;
  overflow: auto;
  position: relative;
}

.fixed-box {
  position: fixed;
  top: 20px;
  right: 20px;
  background: var(--primary);
  color: white;
  padding: 10px;
  border-radius: 4px;
  z-index: 100;
}

.sticky-example {
  height: 300px;
  background: #f8f9fa;
  border-radius: 8px;
  padding: 20px;
  overflow: auto;
}

.sticky-content {
  height: 500px;
}

.sticky-box {
  position: sticky;
  top: 10px;
  background: var(--secondary);
  color: white;
  padding: 15px;
  border-radius: 4px;
  margin: 10px 0;
}

/* Centering examples */
.centering-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
}

.centering-method {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  height: 150px;
}

.centering-method h3 {
  margin-bottom: 15px;
  color: var(--primary);
  font-size: 1.1rem;
}

.centered-block {
  width: 200px;
  background: var(--primary);
  color: white;
  padding: 15px;
  border-radius: 4px;
  margin: 0 auto;
}

.flex-center {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80px;
  background: var(--secondary);
  color: white;
  border-radius: 4px;
}

.absolute-center-container {
  position: relative;
  height: 80px;
  background: #f8f9fa;
  border-radius: 4px;
}

.absolute-centered {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: var(--accent);
  color: white;
  padding: 10px 15px;
  border-radius: 4px;
}

.grid-center {
  display: grid;
  place-items: center;
  height: 80px;
  background: var(--primary);
  color: white;
  border-radius: 4px;
}

/* Layout examples */
.layout-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 40px;
}

.layout-example {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.layout-example h3 {
  margin-bottom: 20px;
  color: var(--primary);
}

/* Float layout */
.float-layout {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 4px;
  overflow: hidden;
}

.float-layout .sidebar {
  float: left;
  width: 30%;
  background: var(--secondary);
  color: white;
  padding: 15px;
  border-radius: 4px;
  margin-right: 5%;
}

.float-layout .main-content {
  float: left;
  width: 65%;
  background: var(--primary);
  color: white;
  padding: 15px;
  border-radius: 4px;
}

/* Flex layout */
.flex-layout {
  display: flex;
  background: #f8f9fa;
  padding: 20px;
  border-radius: 4px;
  gap: 20px;
}

.flex-layout .sidebar {
  flex: 0 0 200px;
  background: var(--secondary);
  color: white;
  padding: 15px;
  border-radius: 4px;
}

.flex-layout .main-content {
  flex: 1;
  background: var(--primary);
  color: white;
  padding: 15px;
  border-radius: 4px;
}

/* Grid layout */
.grid-layout {
  display: grid;
  grid-template-columns: 200px 1fr;
  gap: 20px;
  background: #f8f9fa;
  padding: 20px;
  border-radius: 4px;
}

.grid-layout .sidebar {
  background: var(--secondary);
  color: white;
  padding: 15px;
  border-radius: 4px;
}

.grid-layout .main-content {
  background: var(--primary);
  color: white;
  padding: 15px;
  border-radius: 4px;
}

/* Modal */
.modal-section {
  text-align: center;
  padding: 40px 20px;
  background: white;
  margin: 40px 0;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.open-modal-btn {
  background: var(--primary);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.3s ease;
}

.open-modal-btn:hover {
  background: #2980b9;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: none;
  z-index: 1000;
}

.modal-content {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: white;
  padding: 30px;
  border-radius: 8px;
  max-width: 400px;
  width: 90%;
  text-align: center;
  box-shadow: 0 10px 30px rgba(0,0,0,0.3);
}

.close-modal-btn {
  background: var(--accent);
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 20px;
}

/* Footer */
.footer {
  background: var(--dark);
  color: white;
  text-align: center;
  padding: 20px;
}

/* Responsive */
@media (max-width: 768px) {
  .nav-container {
    flex-direction: column;
    gap: 1rem;
  }
  
  nav ul {
    flex-wrap: wrap;
    justify-content: center;
  }
  
  .section {
    padding: 40px 15px;
  }
  
  .display-examples, .position-examples, 
  .centering-examples, .layout-examples {
    grid-template-columns: 1fr;
  }
  
  .float-layout .sidebar,
  .float-layout .main-content {
    float: none;
    width: 100%;
    margin-bottom: 10px;
  }
  
  .flex-layout {
    flex-direction: column;
  }
  
  .grid-layout {
    grid-template-columns: 1fr;
  }
}
```

```javascript
// script.js
document.addEventListener('DOMContentLoaded', function() {
  const openModalBtn = document.querySelector('.open-modal-btn');
  const closeModalBtn = document.querySelector('.close-modal-btn');
  const modal = document.getElementById('modal');
  
  openModalBtn.addEventListener('click', function() {
    modal.style.display = 'block';
  });
  
  closeModalBtn.addEventListener('click', function() {
    modal.style.display = 'none';
  });
  
  // Cerrar modal al hacer click fuera
  modal.addEventListener('click', function(e) {
    if (e.target === modal) {
      modal.style.display = 'none';
    }
  });
});
```

## Resumen

Display y position son conceptos fundamentales en CSS:

- ✅ **Display**: Controla el flujo del documento (block, inline, flex, grid)
- ✅ **Position**: Controla el posicionamiento (static, relative, absolute, fixed, sticky)
- ✅ **Centrado**: Múltiples técnicas para centrar elementos
- ✅ **Z-index**: Control del orden de apilamiento
- ✅ **Layouts**: De floats tradicionales a modernas técnicas CSS

Comprender estos conceptos es esencial para crear layouts complejos y controlar el comportamiento de los elementos en la página.