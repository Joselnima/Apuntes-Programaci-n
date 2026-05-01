# Módulo 09 - Flexbox

En este módulo aprenderás Flexbox (Flexible Box Layout), el sistema de layout moderno que revolucionó el diseño web. Flexbox es perfecto para layouts unidimensionales.

## Conceptos básicos de Flexbox

### Contenedor flex (flex container)

```css
.flex-container {
  display: flex;        /* Activa flexbox */
  /* display: inline-flex; */ /* Versión inline */
}
```

### Elementos flex (flex items)

```css
/* Los hijos directos del contenedor flex se convierten automáticamente en flex items */
.flex-item {
  /* Propiedades aplicables a los items */
}
```

## Propiedades del contenedor flex

### Dirección del flujo principal (flex-direction)

```css
.container {
  flex-direction: row;            /* Izquierda a derecha (default) */
  flex-direction: row-reverse;    /* Derecha a izquierda */
  flex-direction: column;         /* Arriba a abajo */
  flex-direction: column-reverse; /* Abajo a arriba */
}
```

### Envoltura (flex-wrap)

```css
.container {
  flex-wrap: nowrap;     /* No envolver (default) */
  flex-wrap: wrap;       /* Envolver a la siguiente línea */
  flex-wrap: wrap-reverse; /* Envolver en dirección opuesta */
}
```

### Dirección y envoltura combinadas (flex-flow)

```css
.container {
  flex-flow: row wrap;         /* flex-direction + flex-wrap */
  flex-flow: column nowrap;    /* Dirección columna, sin envoltura */
  flex-flow: row-reverse wrap; /* Dirección reversa, con envoltura */
}
```

## Alineación principal (justify-content)

```css
.container {
  justify-content: flex-start;     /* Inicio del eje principal (default) */
  justify-content: flex-end;       /* Fin del eje principal */
  justify-content: center;         /* Centro */
  justify-content: space-between;  /* Espacio entre items */
  justify-content: space-around;   /* Espacio alrededor de cada item */
  justify-content: space-evenly;   /* Espacio uniforme entre todos */
}
```

## Alineación transversal (align-items)

```css
.container {
  align-items: stretch;      /* Estirar para llenar (default) */
  align-items: flex-start;   /* Inicio del eje transversal */
  align-items: flex-end;     /* Fin del eje transversal */
  align-items: center;       /* Centro */
  align-items: baseline;     /* Línea base del texto */
}
```

## Alineación de múltiples líneas (align-content)

```css
.container {
  flex-wrap: wrap; /* Necesario para que funcione */
  align-content: stretch;        /* Estirar líneas (default) */
  align-content: flex-start;     /* Inicio del contenedor */
  align-content: flex-end;       /* Fin del contenedor */
  align-content: center;         /* Centro */
  align-content: space-between;  /* Espacio entre líneas */
  align-content: space-around;   /* Espacio alrededor de líneas */
}
```

## Propiedades de los elementos flex

### Orden (order)

```css
.item {
  order: 0;  /* Valor por defecto, orden natural */
}

.first {
  order: -1; /* Aparece primero */
}

.last {
  order: 1;  /* Aparece último */
}
```

### Crecimiento (flex-grow)

```css
.item {
  flex-grow: 0;  /* No crece (default) */
}

.grow {
  flex-grow: 1;  /* Crece para ocupar espacio disponible */
}

.grow-more {
  flex-grow: 2;  /* Crece el doble que flex-grow: 1 */
}
```

### Encogimiento (flex-shrink)

```css
.item {
  flex-shrink: 1;  /* Se encoge si es necesario (default) */
}

.no-shrink {
  flex-shrink: 0;  /* No se encoge */
}
```

### Base (flex-basis)

```css
.item {
  flex-basis: auto;    /* Tamaño basado en contenido (default) */
  flex-basis: 200px;   /* Tamaño fijo */
  flex-basis: 50%;     /* Porcentaje */
}
```

### Propiedad shorthand flex

```css
.item {
  flex: 0 1 auto;      /* flex-grow flex-shrink flex-basis */
  flex: 1;             /* flex-grow: 1, otros valores por defecto */
  flex: 2 0 200px;     /* Crece 2x, no se encoge, base 200px */
  flex: none;          /* Equivalente a 0 0 auto */
}
```

### Alineación individual (align-self)

```css
.item {
  align-self: auto;        /* Hereda de align-items (default) */
  align-self: flex-start;  /* Inicio del eje transversal */
  align-self: flex-end;    /* Fin del eje transversal */
  align-self: center;      /* Centro */
  align-self: stretch;     /* Estirar */
  align-self: baseline;    /* Línea base */
}
```

## Patrones comunes de Flexbox

### Navegación horizontal

```css
.nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: #333;
}

.nav-links {
  display: flex;
  list-style: none;
  gap: 2rem;
}

.nav-link {
  color: white;
  text-decoration: none;
}
```

### Tarjetas en grid

```css
.card-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.card {
  flex: 1 1 300px;  /* Crece, se encoge, base 300px */
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 8px;
}
```

### Layout de página básico

```css
.page {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.header {
  /* Header fijo arriba */
}

.main {
  flex: 1;  /* Ocupa el espacio restante */
  display: flex;
}

.sidebar {
  flex: 0 0 250px;
  background: #f5f5f5;
}

.content {
  flex: 1;
  padding: 2rem;
}

.footer {
  /* Footer fijo abajo */
}
```

### Centrado perfecto

```css
.center-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.centered-content {
  text-align: center;
}
```

## Técnicas avanzadas

### Flexbox holy grail layout

```css
.holy-grail {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.header, .footer {
  flex: 0 0 auto;
}

.main {
  flex: 1;
  display: flex;
}

.content {
  flex: 1;
  order: 2;
}

.nav {
  flex: 0 0 200px;
  order: 1;
}

.ads {
  flex: 0 0 200px;
  order: 3;
}
```

### Responsive flexbox

```css
.responsive-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

@media (min-width: 768px) {
  .responsive-container {
    flex-direction: row;
  }
  
  .sidebar {
    flex: 0 0 250px;
  }
  
  .content {
    flex: 1;
  }
}
```

### Flexbox con imágenes

```css
.image-gallery {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.gallery-item {
  flex: 1 1 200px;
  height: 200px;
  overflow: hidden;
  border-radius: 8px;
}

.gallery-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
```

## Comparación con otros métodos

### Ventajas de Flexbox sobre floats

```css
/* Con floats (problemático) */
.float-layout {
  overflow: hidden; /* Necesario para contener floats */
}

.float-layout::after {
  content: '';
  display: table;
  clear: both;
}

.sidebar {
  float: left;
  width: 30%;
}

.content {
  float: right;
  width: 65%;
}

/* Con Flexbox (simple) */
.flex-layout {
  display: flex;
  gap: 20px;
}

.sidebar {
  flex: 0 0 30%;
}

.content {
  flex: 1;
}
```

### Flexbox vs CSS Grid

```css
/* Flexbox: Unidimensional */
.flex-container {
  display: flex;
  justify-content: space-between; /* Eje principal */
  align-items: center;           /* Eje transversal */
}

/* Grid: Bidimensional */
.grid-container {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr; /* Control completo de filas y columnas */
  grid-template-rows: auto 1fr auto;
  gap: 20px;
}
```

## Debugging de Flexbox

### Herramientas del navegador

```css
.debug-flex {
  display: flex;
  /* En DevTools, inspecciona el contenedor para ver: */
  /* - Eje principal y transversal */
  /* - Líneas de grid flex */
  /* - Propiedades de cada item */
}
```

### Clases de debug

```css
.debug-item {
  border: 1px solid red;
  padding: 10px;
  background: rgba(255, 0, 0, 0.1);
}

.debug-item::before {
  content: attr(class);
  position: absolute;
  top: -20px;
  background: yellow;
  padding: 2px 4px;
  font-size: 12px;
}
```

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Flexbox CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>Maestría en Flexbox</h1>
      <nav class="nav">
        <a href="#justify">Justify</a>
        <a href="#align">Align</a>
        <a href="#flex">Flex</a>
        <a href="#patterns">Patrones</a>
      </nav>
    </header>

    <main class="main">
      <section id="justify" class="section">
        <h2>Justify Content</h2>
        <div class="justify-examples">
          <div class="example">
            <h3>flex-start</h3>
            <div class="flex-container justify-start">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
            </div>
          </div>
          
          <div class="example">
            <h3>center</h3>
            <div class="flex-container justify-center">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
            </div>
          </div>
          
          <div class="example">
            <h3>space-between</h3>
            <div class="flex-container justify-between">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
            </div>
          </div>
          
          <div class="example">
            <h3>space-around</h3>
            <div class="flex-container justify-around">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
            </div>
          </div>
        </div>
      </section>

      <section id="align" class="section">
        <h2>Align Items</h2>
        <div class="align-examples">
          <div class="example">
            <h3>flex-start</h3>
            <div class="flex-container align-start">
              <div class="item tall">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
            </div>
          </div>
          
          <div class="example">
            <h3>center</h3>
            <div class="flex-container align-center">
              <div class="item tall">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
            </div>
          </div>
          
          <div class="example">
            <h3>stretch</h3>
            <div class="flex-container align-stretch">
              <div class="item">1</div>
              <div class="item">2</div>
              <div class="item">3</div>
            </div>
          </div>
        </div>
      </section>

      <section id="flex" class="section">
        <h2>Propiedad Flex</h2>
        <div class="flex-examples">
          <div class="example">
            <h3>flex-grow</h3>
            <div class="flex-container">
              <div class="item grow-1">1 (grow: 1)</div>
              <div class="item grow-2">2 (grow: 2)</div>
              <div class="item">3 (grow: 0)</div>
            </div>
          </div>
          
          <div class="example">
            <h3>flex-basis</h3>
            <div class="flex-container">
              <div class="item basis-100">100px</div>
              <div class="item basis-200">200px</div>
              <div class="item basis-auto">auto</div>
            </div>
          </div>
          
          <div class="example">
            <h3>order</h3>
            <div class="flex-container">
              <div class="item order-3">1 (order: 3)</div>
              <div class="item order-1">2 (order: 1)</div>
              <div class="item order-2">3 (order: 2)</div>
            </div>
          </div>
        </div>
      </section>

      <section id="patterns" class="section">
        <h2>Patrones Comunes</h2>
        <div class="pattern-examples">
          <div class="pattern">
            <h3>Navegación</h3>
            <nav class="nav-example">
              <div class="logo">Logo</div>
              <ul class="nav-links">
                <li><a href="#">Inicio</a></li>
                <li><a href="#">Acerca</a></li>
                <li><a href="#">Contacto</a></li>
              </ul>
              <button class="btn">Login</button>
            </nav>
          </div>
          
          <div class="pattern">
            <h3>Card Grid</h3>
            <div class="card-grid">
              <div class="card">
                <h4>Card 1</h4>
                <p>Contenido de la tarjeta</p>
              </div>
              <div class="card">
                <h4>Card 2</h4>
                <p>Más contenido aquí</p>
              </div>
              <div class="card">
                <h4>Card 3</h4>
                <p>Tercera tarjeta</p>
              </div>
            </div>
          </div>
          
          <div class="pattern">
            <h3>Layout de Página</h3>
            <div class="page-layout">
              <header class="page-header">Header</header>
              <div class="page-main">
                <aside class="page-sidebar">Sidebar</aside>
                <main class="page-content">Contenido principal</main>
              </div>
              <footer class="page-footer">Footer</footer>
            </div>
          </div>
          
          <div class="pattern">
            <h3>Galería de Imágenes</h3>
            <div class="image-gallery">
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300" alt="Imagen 1">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=300" alt="Imagen 2">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=300" alt="Imagen 3">
              </div>
              <div class="gallery-item">
                <img src="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=300" alt="Imagen 4">
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>

    <footer class="footer">
      <p>&copy; 2024 Flexbox Master Course</p>
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
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.header {
  background: white;
  padding: 1rem 2rem;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header h1 {
  color: var(--primary);
}

.nav {
  display: flex;
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
  flex: 1;
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

/* Ejemplos de justify */
.justify-examples, .align-examples, .flex-examples {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
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

.flex-container {
  display: flex;
  height: 120px;
  background: #f8f9fa;
  border-radius: 4px;
  padding: 10px;
  gap: 10px;
}

.item {
  background: var(--primary);
  color: white;
  padding: 10px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  min-width: 40px;
}

/* Justify content variations */
.justify-start { justify-content: flex-start; }
.justify-center { justify-content: center; }
.justify-between { justify-content: space-between; }
.justify-around { justify-content: space-around; }

/* Align items variations */
.align-start { align-items: flex-start; }
.align-center { align-items: center; }
.align-stretch { align-items: stretch; }

.tall {
  height: 80px;
}

/* Flex property examples */
.grow-1 { flex-grow: 1; }
.grow-2 { flex-grow: 2; }

.basis-100 { flex-basis: 100px; }
.basis-200 { flex-basis: 200px; }
.basis-auto { flex-basis: auto; }

.order-1 { order: 1; }
.order-2 { order: 2; }
.order-3 { order: 3; }

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

/* Navegación */
.nav-example {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--dark);
  color: white;
  padding: 1rem;
  border-radius: 4px;
}

.nav-links {
  display: flex;
  list-style: none;
  gap: 1.5rem;
}

.nav-links a {
  color: white;
  text-decoration: none;
}

.btn {
  background: var(--secondary);
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
}

/* Card grid */
.card-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.card {
  flex: 1 1 200px;
  background: #f8f9fa;
  padding: 1rem;
  border-radius: 4px;
  border: 1px solid #ddd;
}

.card h4 {
  margin-bottom: 0.5rem;
  color: var(--primary);
}

/* Page layout */
.page-layout {
  display: flex;
  flex-direction: column;
  height: 300px;
  border: 1px solid #ddd;
  border-radius: 4px;
  overflow: hidden;
}

.page-header {
  background: var(--primary);
  color: white;
  padding: 0.5rem;
  text-align: center;
}

.page-main {
  flex: 1;
  display: flex;
}

.page-sidebar {
  background: var(--secondary);
  color: white;
  padding: 1rem;
  width: 150px;
}

.page-content {
  flex: 1;
  background: #f8f9fa;
  padding: 1rem;
}

.page-footer {
  background: var(--dark);
  color: white;
  padding: 0.5rem;
  text-align: center;
}

/* Image gallery */
.image-gallery {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.gallery-item {
  flex: 1 1 150px;
  height: 150px;
  overflow: hidden;
  border-radius: 4px;
}

.gallery-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Responsive */
@media (max-width: 768px) {
  .header {
    flex-direction: column;
    gap: 1rem;
  }
  
  .nav {
    flex-wrap: wrap;
    justify-content: center;
  }
  
  .justify-examples, .align-examples, .flex-examples, .pattern-examples {
    grid-template-columns: 1fr;
  }
  
  .nav-example {
    flex-direction: column;
    gap: 1rem;
  }
  
  .nav-links {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .page-main {
    flex-direction: column;
  }
  
  .page-sidebar {
    width: auto;
  }
}
```

## Resumen

Flexbox es el sistema de layout moderno para CSS:

- ✅ **Contenedor flex**: `display: flex` activa el modo flexbox
- ✅ **Dirección**: `flex-direction` controla el flujo principal
- ✅ **Alineación**: `justify-content` (eje principal), `align-items` (eje transversal)
- ✅ **Flexibilidad**: `flex-grow`, `flex-shrink`, `flex-basis`
- ✅ **Orden**: `order` permite reordenar elementos visualmente

Flexbox es perfecto para layouts unidimensionales y es mucho más simple que los métodos tradicionales con floats.