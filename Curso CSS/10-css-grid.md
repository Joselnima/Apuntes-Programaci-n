# Módulo 10 - CSS Grid

En este módulo aprenderás CSS Grid Layout, el sistema de layout bidimensional más poderoso de CSS. Grid es perfecto para crear layouts complejos de dos dimensiones.

## Conceptos básicos de CSS Grid

### Contenedor grid (grid container)

```css
.grid-container {
  display: grid;        /* Activa CSS Grid */
  /* display: inline-grid; */ /* Versión inline */
}
```

### Elementos grid (grid items)

```css
/* Los hijos directos del contenedor grid se convierten automáticamente en grid items */
.grid-item {
  /* Propiedades aplicables a los items */
}
```

## Definiendo la estructura del grid

### Columnas (grid-template-columns)

```css
.container {
  grid-template-columns: 100px 200px 100px;  /* 3 columnas fijas */
  grid-template-columns: 1fr 2fr 1fr;        /* 3 columnas fraccionarias */
  grid-template-columns: repeat(3, 1fr);     /* 3 columnas iguales */
  grid-template-columns: 200px auto 200px;   /* Mixtas */
  grid-template-columns: minmax(200px, 1fr) 1fr; /* Con minmax */
}
```

### Filas (grid-template-rows)

```css
.container {
  grid-template-rows: 100px 200px 100px;    /* 3 filas fijas */
  grid-template-rows: auto auto auto;        /* Altura automática */
  grid-template-rows: 1fr 2fr 1fr;          /* Fraccionarias */
  grid-template-rows: repeat(3, 100px);     /* Repetición */
}
```

### Áreas de grid (grid-template-areas)

```css
.container {
  grid-template-areas: 
    "header header header"
    "sidebar content aside"
    "footer footer footer";
  
  /* Áreas nombradas */
  grid-template-areas: 
    "nav nav nav"
    "main main sidebar"
    ". . .";  /* Punto = área vacía */
}
```

### Template shorthand (grid-template)

```css
.container {
  grid-template:
    "header header header" 100px
    "sidebar content aside" 1fr
    "footer footer footer" 100px
    / 200px 1fr 200px;  /* Columnas después de / */
}
```

## Controlando el flujo y gaps

### Gap entre celdas (gap)

```css
.container {
  gap: 20px;              /* Gap uniforme */
  gap: 10px 20px;         /* Row gap | Column gap */
  row-gap: 10px;          /* Solo filas */
  column-gap: 20px;       /* Solo columnas */
}
```

### Flujo automático (grid-auto-flow)

```css
.container {
  grid-auto-flow: row;         /* Llenar por filas (default) */
  grid-auto-flow: column;      /* Llenar por columnas */
  grid-auto-flow: dense;       /* Llenar huecos */
}
```

### Tamaño automático de filas/columnas

```css
.container {
  grid-auto-rows: 100px;       /* Altura automática de filas */
  grid-auto-columns: 200px;    /* Ancho automático de columnas */
  grid-auto-rows: minmax(100px, auto); /* Con minmax */
}
```

## Posicionamiento de grid items

### Líneas de grid

```css
.item {
  grid-column-start: 1;    /* Línea de columna inicial */
  grid-column-end: 3;      /* Línea de columna final */
  grid-row-start: 1;       /* Línea de fila inicial */
  grid-row-end: 2;         /* Línea de fila final */
}
```

### Shorthand para líneas

```css
.item {
  grid-column: 1 / 3;      /* start / end */
  grid-row: 1 / 2;         /* start / end */
  grid-column: 1 / span 2; /* start / span count */
  grid-row: 2 / span 3;    /* row / span count */
}
```

### Posicionamiento por áreas nombradas

```css
.header {
  grid-area: header;       /* Nombre del área */
}

.sidebar {
  grid-area: sidebar;
}

.content {
  grid-area: content;
}
```

## Alineación en CSS Grid

### Alineación del contenedor

```css
.container {
  justify-items: start;      /* Alineación horizontal de items */
  align-items: center;       /* Alineación vertical de items */
  place-items: center center; /* justify-items + align-items */
  
  justify-content: space-between; /* Alineación del grid completo */
  align-content: center;          /* Alineación vertical del grid */
  place-content: center center;    /* justify-content + align-content */
}
```

### Alineación individual de items

```css
.item {
  justify-self: center;     /* Alineación horizontal individual */
  align-self: end;          /* Alineación vertical individual */
  place-self: center end;   /* justify-self + align-self */
}
```

## Unidades y funciones de Grid

### Fracción (fr)

```css
.container {
  grid-template-columns: 1fr 2fr 1fr;  /* 1+2+1 = 4 partes */
  /* Primera columna: 1/4, segunda: 2/4, tercera: 1/4 */
}
```

### Minmax()

```css
.container {
  grid-template-columns: minmax(200px, 1fr) 1fr;
  /* Mínimo 200px, máximo 1fr */
}
```

### Fit-content()

```css
.container {
  grid-template-columns: fit-content(300px) 1fr;
  /* Ajusta al contenido, máximo 300px */
}
```

### Repeat() con auto-fit/auto-fill

```css
.container {
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  /* Crea tantas columnas como quepan */
  grid-template-columns: repeat(auto-fill, 200px);
  /* Llena el espacio disponible */
}
```

## Patrones comunes de CSS Grid

### Grid de 12 columnas (Bootstrap-like)

```css
.grid-12 {
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  gap: 20px;
}

.col-1 { grid-column: span 1; }
.col-2 { grid-column: span 2; }
.col-3 { grid-column: span 3; }
.col-4 { grid-column: span 4; }
.col-6 { grid-column: span 6; }
.col-12 { grid-column: span 12; }
```

### Layout de página completo

```css
.page {
  display: grid;
  grid-template:
    "nav nav nav" 60px
    "sidebar main main" 1fr
    "sidebar footer footer" 60px
    / 200px 1fr 1fr;
  min-height: 100vh;
}

.nav { grid-area: nav; }
.sidebar { grid-area: sidebar; }
.main { grid-area: main; }
.footer { grid-area: footer; }
```

### Galería de imágenes responsiva

```css
.gallery {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.gallery-item {
  aspect-ratio: 1; /* Cuadradas */
  overflow: hidden;
}

.gallery-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
```

### Grid de tarjetas

```css
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
}

.card {
  display: grid;
  grid-template-rows: auto 1fr auto;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 15px rgba(0,0,0,0.1);
}

.card-header { /* Ocupa primera fila */ }
.card-content { /* Ocupa fila media (1fr) */ }
.card-footer { /* Ocupa última fila */ }
```

## Técnicas avanzadas

### Grid implícito vs explícito

```css
.container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);  /* Grid explícito: 3 columnas */
  grid-template-rows: 100px 100px;       /* Grid explícito: 2 filas */
  /* Los items que excedan se colocan en grid implícito */
  grid-auto-rows: 50px;                  /* Altura de filas implícitas */
}
```

### Grid con áreas nombradas complejas

```css
.dashboard {
  display: grid;
  grid-template-areas: 
    "header header header header"
    "nav main main aside"
    "nav main main aside"
    "footer footer footer footer";
  grid-template-rows: 60px 1fr 1fr 60px;
  grid-template-columns: 200px 1fr 1fr 250px;
  gap: 20px;
}
```

### Grid responsive con media queries

```css
.grid-responsive {
  display: grid;
  grid-template-columns: 1fr;
  gap: 20px;
}

@media (min-width: 768px) {
  .grid-responsive {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1200px) {
  .grid-responsive {
    grid-template-columns: repeat(3, 1fr);
  }
}
```

### Subgrids (CSS Grid Level 2)

```css
.parent {
  display: grid;
  grid-template-columns: 200px 1fr;
}

.child {
  display: grid;
  grid-template-columns: subgrid;  /* Hereda las líneas del padre */
  grid-column: 2;                  /* Ocupa la segunda columna del padre */
}
```

## Comparación Grid vs Flexbox

### Cuándo usar Grid

```css
/* Grid: Layouts bidimensionales complejos */
.dashboard {
  display: grid;
  grid-template-areas: 
    "header header header"
    "sidebar content aside"
    "footer footer footer";
}

/* Grid: Alineación perfecta en 2D */
.photo-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}
```

### Cuándo usar Flexbox

```css
/* Flexbox: Layouts unidimensionales */
.navigation {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* Flexbox: Contenido dinámico */
.card {
  display: flex;
  flex-direction: column;
}
```

## Debugging de CSS Grid

### Herramientas de desarrollo

```css
.debug-grid {
  display: grid;
  /* En DevTools, inspecciona para ver: */
  /* - Líneas de grid numeradas */
  /* - Áreas nombradas */
  /* - Posicionamiento de items */
  /* - Gaps y alineación */
}
```

### Visualización de grid

```css
.debug-grid {
  background: 
    linear-gradient(to right, transparent 99%, rgba(255,0,0,0.2) 100%),
    linear-gradient(to bottom, transparent 99%, rgba(255,0,0,0.2) 100%);
  background-size: var(--column-width) var(--row-height);
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Grid Layout</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>CSS Grid Mastery</h1>
      <nav class="nav">
        <a href="#basics">Básicos</a>
        <a href="#positioning">Posicionamiento</a>
        <a href="#patterns">Patrones</a>
        <a href="#advanced">Avanzado</a>
      </nav>
    </header>

    <main class="main">
      <section id="basics" class="section">
        <h2>Conceptos Básicos</h2>
        
        <div class="grid-examples">
          <div class="example">
            <h3>Grid Simple</h3>
            <div class="simple-grid">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
              <div class="item">4</div>
              <div class="item">5</div>
              <div class="item">6</div>
            </div>
          </div>
          
          <div class="example">
            <h3>Fracciones (fr)</h3>
            <div class="fr-grid">
              <div class="item">1fr</div>
              <div class="item">2fr</div>
              <div class="item">1fr</div>
            </div>
          </div>
          
          <div class="example">
            <h3>Repeat y Auto-fit</h3>
            <div class="auto-grid">
              <div class="item">A</div>
              <div class="item">B</div>
              <div class="item">C</div>
              <div class="item">D</div>
              <div class="item">E</div>
              <div class="item">F</div>
              <div class="item">G</div>
              <div class="item">H</div>
            </div>
          </div>
        </div>
      </section>

      <section id="positioning" class="section">
        <h2>Posicionamiento</h2>
        
        <div class="positioning-examples">
          <div class="example">
            <h3>Líneas de Grid</h3>
            <div class="line-grid">
              <div class="item span-2">Span 2 columnas</div>
              <div class="item">Normal</div>
              <div class="item">Normal</div>
              <div class="item span-2-rows">Span 2 filas</div>
              <div class="item">Normal</div>
            </div>
          </div>
          
          <div class="example">
            <h3>Áreas Nombradas</h3>
            <div class="named-areas">
              <div class="header">Header</div>
              <div class="sidebar">Sidebar</div>
              <div class="content">Content</div>
              <div class="footer">Footer</div>
            </div>
          </div>
          
          <div class="example">
            <h3>Alineación</h3>
            <div class="alignment-grid">
              <div class="item start">Start</div>
              <div class="item center">Center</div>
              <div class="item end">End</div>
            </div>
          </div>
        </div>
      </section>

      <section id="patterns" class="section">
        <h2>Patrones Comunes</h2>
        
        <div class="pattern-examples">
          <div class="pattern">
            <h3>Layout de Página</h3>
            <div class="page-layout">
              <header class="page-header">Header</header>
              <nav class="page-nav">Navigation</nav>
              <main class="page-main">Main Content</main>
              <aside class="page-aside">Aside</aside>
              <footer class="page-footer">Footer</footer>
            </div>
          </div>
          
          <div class="pattern">
            <h3>Galería de Imágenes</h3>
            <div class="image-gallery">
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300" alt="1">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=300" alt="2">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300" alt="3">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=300" alt="4">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300" alt="5">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=300" alt="6">
              </div>
            </div>
          </div>
          
          <div class="pattern">
            <h3>Grid de Tarjetas</h3>
            <div class="card-grid">
              <div class="card">
                <div class="card-header">
                  <h4>Tarjeta 1</h4>
                </div>
                <div class="card-content">
                  <p>Contenido de la primera tarjeta con texto de ejemplo.</p>
                </div>
                <div class="card-footer">
                  <button>Leer más</button>
                </div>
              </div>
              <div class="card">
                <div class="card-header">
                  <h4>Tarjeta 2</h4>
                </div>
                <div class="card-content">
                  <p>Más contenido aquí para mostrar el layout de la tarjeta.</p>
                </div>
                <div class="card-footer">
                  <button>Leer más</button>
                </div>
              </div>
              <div class="card">
                <div class="card-header">
                  <h4>Tarjeta 3</h4>
                </div>
                <div class="card-content">
                  <p>Contenido adicional para completar el ejemplo de grid.</p>
                </div>
                <div class="card-footer">
                  <button>Leer más</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section id="advanced" class="section">
        <h2>Técnicas Avanzadas</h2>
        
        <div class="advanced-examples">
          <div class="example">
            <h3>Grid Implícito</h3>
            <div class="implicit-grid">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
              <div class="item">4</div>
              <div class="item">5</div>
              <div class="item">6</div>
              <div class="item">7</div>
              <div class="item">8</div>
              <div class="item">9</div>
            </div>
          </div>
          
          <div class="example">
            <h3>Minmax y Fit-content</h3>
            <div class="minmax-grid">
              <div class="item">Contenido corto</div>
              <div class="item">Contenido más largo que debería ocupar más espacio</div>
              <div class="item">Medio</div>
            </div>
          </div>
          
          <div class="example">
            <h3>Grid Responsivo</h3>
            <div class="responsive-grid">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
              <div class="item">4</div>
              <div class="item">5</div>
              <div class="item">6</div>
            </div>
          </div>
        </div>
      </section>
    </main>

    <footer class="footer">
      <p>&copy; 2024 CSS Grid Master Course</p>
    </footer>
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

/* Layout principal */
.container {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
}

.header {
  background: white;
  padding: 1rem 2rem;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  display: grid;
  grid-template-columns: 1fr auto;
  align-items: center;
  gap: 2rem;
}

.header h1 {
  color: var(--primary);
}

.nav {
  display: grid;
  grid-template-columns: repeat(4, auto);
  gap: 2rem;
}

.nav a {
  text-decoration: none;
  color: var(--dark);
  font-weight: 500;
  transition: color 0.3s ease;
}

.nav a:hover {
  color: var(--primary);
}

.main {
  padding: 2rem;
}

.footer {
  background: var(--dark);
  color: white;
  text-align: center;
  padding: 1rem;
}

/* Secciones */
.section {
  margin-bottom: 4rem;
}

.section h2 {
  text-align: center;
  font-size: 2.5rem;
  margin-bottom: 2rem;
  color: var(--dark);
  position: relative;
}

.section h2::after {
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

/* Grid examples */
.grid-examples, .positioning-examples, .advanced-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 2rem;
  margin-bottom: 2rem;
}

.example {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.example h3 {
  margin-bottom: 1rem;
  color: var(--primary);
  font-size: 1.2rem;
}

/* Items básicos */
.item {
  background: var(--primary);
  color: white;
  padding: 1rem;
  border-radius: 4px;
  display: grid;
  place-items: center;
  font-weight: bold;
}

/* Grid simple */
.simple-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(2, 100px);
  gap: 10px;
}

/* Fracciones */
.fr-grid {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  grid-template-rows: 100px;
  gap: 10px;
}

/* Auto-fit */
.auto-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(80px, 1fr));
  gap: 10px;
}

/* Posicionamiento */
.line-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(3, 80px);
  gap: 10px;
}

.span-2 {
  grid-column: span 2;
  background: var(--secondary) !important;
}

.span-2-rows {
  grid-row: span 2;
  background: var(--accent) !important;
}

/* Áreas nombradas */
.named-areas {
  display: grid;
  grid-template-areas: 
    "header header"
    "sidebar content"
    "footer footer";
  grid-template-rows: 60px 1fr 60px;
  grid-template-columns: 150px 1fr;
  gap: 10px;
  height: 250px;
}

.named-areas .header {
  grid-area: header;
  background: var(--primary);
}

.named-areas .sidebar {
  grid-area: sidebar;
  background: var(--secondary);
}

.named-areas .content {
  grid-area: content;
  background: var(--accent);
}

.named-areas .footer {
  grid-area: footer;
  background: var(--dark);
}

/* Alineación */
.alignment-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: 120px;
  gap: 10px;
  align-items: end;
}

.start { justify-self: start; }
.center { justify-self: center; }
.end { justify-self: end; }

/* Patrones */
.pattern-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 2rem;
}

.pattern {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.pattern h3 {
  margin-bottom: 1rem;
  color: var(--primary);
}

/* Page layout */
.page-layout {
  display: grid;
  grid-template-areas: 
    "header header header"
    "nav main aside"
    "footer footer footer";
  grid-template-rows: 50px 1fr 50px;
  grid-template-columns: 120px 1fr 150px;
  gap: 10px;
  height: 300px;
  font-size: 0.9rem;
}

.page-header { grid-area: header; background: var(--primary); }
.page-nav { grid-area: nav; background: var(--secondary); }
.page-main { grid-area: main; background: var(--accent); }
.page-aside { grid-area: aside; background: var(--dark); color: white; }
.page-footer { grid-area: footer; background: var(--gray); }

/* Image gallery */
.image-gallery {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 15px;
}

.gallery-item {
  aspect-ratio: 1;
  overflow: hidden;
  border-radius: 8px;
}

.gallery-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Card grid */
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.card {
  display: grid;
  grid-template-rows: auto 1fr auto;
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.card-header {
  background: var(--light);
  padding: 1rem;
  border-bottom: 1px solid #ddd;
}

.card-content {
  padding: 1rem;
}

.card-footer {
  background: var(--light);
  padding: 1rem;
  border-top: 1px solid #ddd;
  text-align: center;
}

.card button {
  background: var(--primary);
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
}

/* Avanzado */
.implicit-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(2, 80px);
  grid-auto-rows: 60px;
  gap: 10px;
}

.minmax-grid {
  display: grid;
  grid-template-columns: fit-content(200px) 1fr fit-content(150px);
  gap: 10px;
}

.responsive-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 10px;
}

/* Responsive */
@media (max-width: 768px) {
  .header {
    grid-template-columns: 1fr;
    text-align: center;
  }
  
  .nav {
    grid-template-columns: repeat(2, 1fr);
    gap: 1rem;
  }
  
  .grid-examples, .positioning-examples, .advanced-examples, .pattern-examples {
    grid-template-columns: 1fr;
  }
  
  .named-areas {
    grid-template-areas: 
      "header"
      "sidebar"
      "content"
      "footer";
    grid-template-columns: 1fr;
    grid-template-rows: 60px auto 1fr 60px;
  }
  
  .page-layout {
    grid-template-areas: 
      "header"
      "nav"
      "main"
      "aside"
      "footer";
    grid-template-columns: 1fr;
    grid-template-rows: 50px 80px 1fr 80px 50px;
  }
  
  .card-grid {
    grid-template-columns: 1fr;
  }
}
```

## Resumen

CSS Grid es el sistema de layout bidimensional más poderoso:

- ✅ **Grid explícito**: `grid-template-columns` y `grid-template-rows`
- ✅ **Áreas nombradas**: `grid-template-areas` para layouts complejos
- ✅ **Posicionamiento**: `grid-column` y `grid-row` para items individuales
- ✅ **Unidades flexibles**: `fr`, `minmax()`, `fit-content()`
- ✅ **Auto-placement**: `repeat(auto-fit, ...)` para responsividad

Grid es perfecto para layouts bidimensionales complejos, mientras que Flexbox es mejor para layouts unidimensionales.