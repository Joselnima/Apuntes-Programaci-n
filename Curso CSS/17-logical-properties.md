# Módulo 17 - Logical Properties

En este módulo aprenderás sobre CSS Logical Properties, una característica avanzada que permite crear layouts que se adaptan automáticamente a diferentes direcciones de escritura y orientaciones de texto.

## Introducción a Logical Properties

### ¿Qué son las Logical Properties?

Las propiedades lógicas CSS reemplazan las propiedades físicas tradicionales (top, right, bottom, left) con propiedades que se basan en el flujo lógico del contenido, adaptándose automáticamente a diferentes idiomas y direcciones de escritura.

```css
/* Propiedades físicas tradicionales */
.element {
  margin-top: 20px;
  margin-right: 15px;
  margin-bottom: 20px;
  margin-left: 15px;
}

/* Propiedades lógicas equivalentes */
.element {
  margin-block-start: 20px;
  margin-inline-end: 15px;
  margin-block-end: 20px;
  margin-inline-start: 15px;
}
```

## Block vs Inline Dimensions

### Entendiendo las direcciones lógicas

```css
/* Block direction: dirección del flujo de bloques (vertical en idiomas LTR) */
.element {
  margin-block-start: 10px;  /* Equivalente a margin-top en LTR */
  margin-block-end: 10px;    /* Equivalente a margin-bottom en LTR */
  padding-block-start: 15px; /* Equivalente a padding-top en LTR */
  padding-block-end: 15px;   /* Equivalente a padding-bottom en LTR */
}

/* Inline direction: dirección del flujo inline (horizontal en idiomas LTR) */
.element {
  margin-inline-start: 20px; /* Equivalente a margin-left en LTR */
  margin-inline-end: 20px;   /* Equivalente a margin-right en LTR */
  padding-inline-start: 25px;/* Equivalente a padding-left en LTR */
  padding-inline-end: 25px;  /* Equivalente a padding-right en LTR */
}
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es" dir="ltr">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Logical Properties - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: 'Arial', sans-serif;
      background: #f5f5f7;
      color: #333;
      padding: 20px;
    }

    .container {
      max-width: 1000px;
      margin: 0 auto;
      background: white;
      border-radius: 12px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
      overflow: hidden;
    }

    .header {
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
      padding-block-start: 30px;
      padding-block-end: 30px;
      padding-inline-start: 40px;
      padding-inline-end: 40px;
    }

    .header h1 {
      font-size: 2rem;
      margin-block-end: 10px;
    }

    .header p {
      opacity: 0.9;
      font-size: 1.1rem;
    }

    .content {
      padding-block-start: 40px;
      padding-block-end: 40px;
      padding-inline-start: 40px;
      padding-inline-end: 40px;
    }

    .demo-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 30px;
      margin-block-end: 40px;
    }

    .demo-card {
      background: #f8f9fa;
      border-radius: 8px;
      padding-block-start: 25px;
      padding-block-end: 25px;
      padding-inline-start: 25px;
      padding-inline-end: 25px;
      border-inline-start: 4px solid #667eea;
    }

    .demo-card h3 {
      color: #667eea;
      margin-block-end: 15px;
      font-size: 1.2rem;
    }

    .demo-card p {
      line-height: 1.6;
      color: #666;
    }

    .code-example {
      background: #2d3748;
      color: #e2e8f0;
      padding-block-start: 20px;
      padding-block-end: 20px;
      padding-inline-start: 20px;
      padding-inline-end: 20px;
      border-radius: 6px;
      font-family: 'Monaco', 'Menlo', monospace;
      font-size: 14px;
      margin-block-start: 15px;
      overflow-x: auto;
    }

    .direction-toggle {
      text-align: center;
      padding-block-start: 30px;
      padding-block-end: 30px;
      border-block-start: 1px solid #e0e0e0;
    }

    .toggle-btn {
      background: #667eea;
      color: white;
      border: none;
      padding-block-start: 12px;
      padding-block-end: 12px;
      padding-inline-start: 24px;
      padding-inline-end: 24px;
      border-radius: 6px;
      cursor: pointer;
      font-size: 16px;
      transition: background-color 0.2s ease;
    }

    .toggle-btn:hover {
      background: #5a67d8;
    }

    .comparison {
      margin-block-start: 40px;
      padding-block-start: 30px;
      border-block-start: 1px solid #e0e0e0;
    }

    .comparison h2 {
      text-align: center;
      margin-block-end: 30px;
      color: #333;
    }

    .comparison-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 20px;
    }

    .comparison-item {
      background: #f8f9fa;
      padding-block-start: 20px;
      padding-block-end: 20px;
      padding-inline-start: 20px;
      padding-inline-end: 20px;
      border-radius: 8px;
    }

    .comparison-item h4 {
      margin-block-end: 10px;
      color: #667eea;
    }

    .comparison-item .physical {
      background: #fed7d7;
      border-inline-start: 3px solid #e53e3e;
    }

    .comparison-item .logical {
      background: #c6f6d5;
      border-inline-start: 3px solid #38a169;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>CSS Logical Properties</h1>
      <p>Propiedades que se adaptan a la dirección del texto</p>
    </header>

    <div class="content">
      <div class="demo-grid">
        <div class="demo-card">
          <h3>Block Direction</h3>
          <p>Se refiere a la dirección en la que fluyen los bloques de contenido. En idiomas de izquierda a derecha (LTR), esto es vertical.</p>
          <div class="code-example">
margin-block-start: 20px; /* top en LTR */<br>
margin-block-end: 20px;   /* bottom en LTR */
          </div>
        </div>

        <div class="demo-card">
          <h3>Inline Direction</h3>
          <p>Se refiere a la dirección en la que fluye el contenido dentro de una línea. En idiomas LTR, esto es horizontal.</p>
          <div class="code-example">
margin-inline-start: 15px; /* left en LTR */<br>
margin-inline-end: 15px;   /* right en LTR */
          </div>
        </div>
      </div>

      <div class="direction-toggle">
        <button class="toggle-btn" onclick="toggleDirection()">Cambiar Dirección (RTL/LTR)</button>
      </div>

      <div class="comparison">
        <h2>Comparación: Propiedades Físicas vs Lógicas</h2>

        <div class="comparison-grid">
          <div class="comparison-item physical">
            <h4>Propiedades Físicas</h4>
            <p>No se adaptan a diferentes direcciones de escritura. En RTL, margin-left sigue siendo margin-left.</p>
            <div class="code-example">
margin-top: 20px;<br>
margin-right: 15px;<br>
margin-bottom: 20px;<br>
margin-left: 15px;
            </div>
          </div>

          <div class="comparison-item logical">
            <h4>Propiedades Lógicas</h4>
            <p>Se adaptan automáticamente. En RTL, margin-inline-start se convierte en margin-right.</p>
            <div class="code-example">
margin-block-start: 20px;<br>
margin-inline-end: 15px;<br>
margin-block-end: 20px;<br>
margin-inline-start: 15px;
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <script>
    function toggleDirection() {
      const html = document.documentElement;
      const currentDir = html.getAttribute('dir');
      const newDir = currentDir === 'ltr' ? 'rtl' : 'ltr';
      html.setAttribute('dir', newDir);

      // Actualizar el botón
      const button = document.querySelector('.toggle-btn');
      button.textContent = `Cambiar Dirección (${newDir.toUpperCase()})`;
    }
  </script>
</body>
</html>
```

## Propiedades Lógicas Comunes

### Margin, Padding y Border

```css
/* Margin */
.element {
  margin-block-start: 10px;   /* margin-top */
  margin-block-end: 10px;     /* margin-bottom */
  margin-inline-start: 15px;  /* margin-left */
  margin-inline-end: 15px;    /* margin-right */
}

/* Padding */
.element {
  padding-block-start: 20px;  /* padding-top */
  padding-block-end: 20px;    /* padding-bottom */
  padding-inline-start: 25px; /* padding-left */
  padding-inline-end: 25px;   /* padding-right */
}

/* Border */
.element {
  border-block-start: 2px solid #333;  /* border-top */
  border-block-end: 2px solid #333;    /* border-bottom */
  border-inline-start: 2px solid #333; /* border-left */
  border-inline-end: 2px solid #333;   /* border-right */
}
```

### Position

```css
/* Position offsets */
.element {
  position: absolute;
  inset-block-start: 20px;   /* top */
  inset-block-end: 20px;     /* bottom */
  inset-inline-start: 15px;  /* left */
  inset-inline-end: 15px;    /* right */
}

/* Shorthand */
.element {
  inset: 20px 15px 20px 15px; /* block-start, inline-end, block-end, inline-start */
}
```

### Width y Height

```css
/* Dimensiones lógicas */
.element {
  block-size: 200px;   /* height en LTR */
  inline-size: 300px;  /* width en LTR */
}

/* Min/Max */
.element {
  min-block-size: 100px;   /* min-height */
  max-block-size: 500px;   /* max-height */
  min-inline-size: 200px;  /* min-width */
  max-inline-size: 800px;  /* max-width */
}
```

## Ejemplo con Layout Complejo

```html
<!DOCTYPE html>
<html lang="es" dir="ltr">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Logical Properties - Layout</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: #f8f9fa;
      color: #333;
    }

    .app {
      display: flex;
      flex-direction: column;
      min-block-size: 100vh;
    }

    .header {
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
      padding-block-start: 20px;
      padding-block-end: 20px;
      padding-inline-start: 24px;
      padding-inline-end: 24px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      box-shadow: 0 2px 10px rgba(0,0,0,0.1);
    }

    .logo {
      font-size: 1.5rem;
      font-weight: bold;
    }

    .nav {
      display: flex;
      gap: 20px;
    }

    .nav-link {
      color: white;
      text-decoration: none;
      padding-block-start: 8px;
      padding-block-end: 8px;
      padding-inline-start: 12px;
      padding-inline-end: 12px;
      border-radius: 6px;
      transition: background-color 0.2s ease;
    }

    .nav-link:hover {
      background: rgba(255,255,255,0.1);
    }

    .main {
      display: flex;
      flex: 1;
    }

    .sidebar {
      background: white;
      inline-size: 280px;
      padding-block-start: 30px;
      padding-block-end: 30px;
      padding-inline-start: 24px;
      padding-inline-end: 24px;
      border-inline-end: 1px solid #e0e0e0;
      box-shadow: inset -1px 0 0 rgba(0,0,0,0.1);
    }

    .sidebar-title {
      font-size: 1.1rem;
      font-weight: 600;
      margin-block-end: 20px;
      color: #333;
    }

    .sidebar-item {
      display: block;
      color: #666;
      text-decoration: none;
      padding-block-start: 10px;
      padding-block-end: 10px;
      padding-inline-start: 12px;
      padding-inline-end: 12px;
      border-radius: 6px;
      margin-block-end: 5px;
      transition: all 0.2s ease;
    }

    .sidebar-item:hover {
      background: #f0f0f0;
      color: #333;
    }

    .sidebar-item.active {
      background: #667eea;
      color: white;
    }

    .content {
      flex: 1;
      padding-block-start: 40px;
      padding-block-end: 40px;
      padding-inline-start: 40px;
      padding-inline-end: 40px;
      background: white;
    }

    .content-header {
      margin-block-end: 30px;
    }

    .content-title {
      font-size: 2rem;
      margin-block-end: 10px;
      color: #333;
    }

    .content-subtitle {
      color: #666;
      font-size: 1.1rem;
    }

    .card-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 24px;
    }

    .card {
      background: #f8f9fa;
      border-radius: 12px;
      padding-block-start: 24px;
      padding-block-end: 24px;
      padding-inline-start: 24px;
      padding-inline-end: 24px;
      border: 1px solid #e0e0e0;
      transition: box-shadow 0.2s ease;
    }

    .card:hover {
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
    }

    .card-title {
      font-size: 1.2rem;
      margin-block-end: 12px;
      color: #333;
    }

    .card-text {
      color: #666;
      line-height: 1.6;
    }

    .footer {
      background: #2d3748;
      color: white;
      padding-block-start: 30px;
      padding-block-end: 30px;
      padding-inline-start: 40px;
      padding-inline-end: 40px;
      text-align: center;
    }

    .direction-controls {
      position: fixed;
      inset-block-end: 20px;
      inset-inline-end: 20px;
      background: #667eea;
      color: white;
      padding-block-start: 12px;
      padding-block-end: 12px;
      padding-inline-start: 16px;
      padding-inline-end: 16px;
      border-radius: 8px;
      border: none;
      cursor: pointer;
      box-shadow: 0 4px 12px rgba(0,0,0,0.3);
    }

    /* Responsive */
    @media (max-inline-size: 768px) {
      .main {
        flex-direction: column;
      }

      .sidebar {
        inline-size: 100%;
        border-inline-end: none;
        border-block-end: 1px solid #e0e0e0;
      }

      .nav {
        display: none;
      }
    }
  </style>
</head>
<body>
  <div class="app">
    <header class="header">
      <div class="logo">Mi App</div>
      <nav class="nav">
        <a href="#" class="nav-link">Inicio</a>
        <a href="#" class="nav-link">Productos</a>
        <a href="#" class="nav-link">Contacto</a>
      </nav>
    </header>

    <div class="main">
      <aside class="sidebar">
        <h2 class="sidebar-title">Menú</h2>
        <a href="#" class="sidebar-item active">Dashboard</a>
        <a href="#" class="sidebar-item">Usuarios</a>
        <a href="#" class="sidebar-item">Configuración</a>
        <a href="#" class="sidebar-item">Reportes</a>
      </aside>

      <main class="content">
        <header class="content-header">
          <h1 class="content-title">Dashboard</h1>
          <p class="content-subtitle">Bienvenido a tu panel de control</p>
        </header>

        <div class="card-grid">
          <div class="card">
            <h3 class="card-title">Estadísticas</h3>
            <p class="card-text">Visualiza tus métricas principales con gráficos interactivos y análisis en tiempo real.</p>
          </div>

          <div class="card">
            <h3 class="card-title">Usuarios</h3>
            <p class="card-text">Gestiona los usuarios de tu aplicación con herramientas avanzadas de administración.</p>
          </div>

          <div class="card">
            <h3 class="card-title">Configuración</h3>
            <p class="card-text">Personaliza la configuración de tu aplicación según tus necesidades específicas.</p>
          </div>

          <div class="card">
            <h3 class="card-title">Reportes</h3>
            <p class="card-text">Genera reportes detallados con datos exportables en múltiples formatos.</p>
          </div>
        </div>
      </main>
    </div>

    <footer class="footer">
      <p>&copy; 2024 Mi App. Todos los derechos reservados.</p>
    </footer>

    <button class="direction-controls" onclick="toggleDirection()">
      Cambiar Dirección
    </button>
  </div>

  <script>
    function toggleDirection() {
      const html = document.documentElement;
      const currentDir = html.getAttribute('dir');
      const newDir = currentDir === 'ltr' ? 'rtl' : 'ltr';
      html.setAttribute('dir', newDir);
    }
  </script>
</body>
</html>
```

## Border Radius Lógico

```css
/* Border radius lógico */
.element {
  border-start-start-radius: 10px;  /* top-left en LTR */
  border-start-end-radius: 10px;    /* top-right en LTR */
  border-end-start-radius: 10px;    /* bottom-left en LTR */
  border-end-end-radius: 10px;      /* bottom-right en LTR */
}
```

## Text Align Lógico

```css
/* Alineación de texto lógica */
.text-start {
  text-align: start;  /* left en LTR, right en RTL */
}

.text-end {
  text-align: end;    /* right en LTR, left en RTL */
}
```

## Float Lógico

```css
/* Float lógico */
.float-start {
  float: inline-start;  /* left en LTR, right en RTL */
}

.float-end {
  float: inline-end;    /* right en LTR, left en RTL */
}
```

## Casos de Uso Prácticos

### 1. Componentes reutilizables

```css
/* Componente de tarjeta que funciona en cualquier dirección */
.card {
  padding-block-start: 20px;
  padding-block-end: 20px;
  padding-inline-start: 24px;
  padding-inline-end: 24px;
  border-inline-start: 4px solid #667eea;
  margin-inline-end: 16px;
}
```

### 2. Navegación responsive

```css
/* Navegación que se adapta */
.navbar {
  display: flex;
  justify-content: space-between;
  padding-inline-start: 20px;
  padding-inline-end: 20px;
}

.nav-menu {
  margin-inline-start: auto;
}
```

### 3. Layouts de grid flexibles

```css
/* Grid que respeta la dirección */
.grid-container {
  display: grid;
  grid-template-columns: 1fr 3fr;
  gap: 24px;
  padding-inline-start: 20px;
  padding-inline-end: 20px;
}
```

## Compatibilidad y Fallbacks

### Verificando soporte

```css
@supports (margin-inline-start: 10px) {
  .logical-supported {
    margin-inline-start: 20px;
    margin-inline-end: 20px;
  }
}

@supports not (margin-inline-start: 10px) {
  .logical-fallback {
    margin-left: 20px;
    margin-right: 20px;
  }
}
```

### PostCSS para compatibilidad

```javascript
// Usando postcss-logical para convertir propiedades lógicas
// a propiedades físicas con fallbacks automáticos
```

## Mejores Prácticas

### 1. Usar propiedades lógicas desde el inicio

```css
/* ✅ Bueno: usar propiedades lógicas */
.component {
  margin-inline-start: 16px;
  padding-block-end: 20px;
}

/* ❌ Evitar: mezclar propiedades físicas y lógicas */
.component {
  margin-left: 16px;      /* físico */
  padding-block-end: 20px; /* lógico */
}
```

### 2. Considerar el contexto

```css
/* Para componentes que siempre deben estar en un lado específico */
.left-aligned {
  margin-inline-start: 0; /* Siempre al inicio de la línea */
}

.right-aligned {
  margin-inline-end: 0; /* Siempre al final de la línea */
}
```

### 3. Testing en múltiples direcciones

```css
/* Asegurarse de que funciona en RTL */
[dir="rtl"] .component {
  /* Ajustes específicos para RTL si son necesarios */
}
```

## Resumen

CSS Logical Properties revolucionan el diseño web internacional:

- ✅ **Block vs Inline**: Dimensiones basadas en el flujo del contenido
- ✅ **Adaptabilidad**: Se ajustan automáticamente a LTR/RTL
- ✅ **Compatibilidad**: Excelente soporte en navegadores modernos
- ✅ **Mantenibilidad**: Código más limpio y semántico
- ✅ **Internacionalización**: Ideal para sitios multiidioma
- ✅ **Futuro-proof**: Preparado para futuras direcciones de escritura

Las Logical Properties son esenciales para crear interfaces que funcionen perfectamente en cualquier idioma y dirección de escritura, desde árabe y hebreo hasta idiomas verticales como el japonés.