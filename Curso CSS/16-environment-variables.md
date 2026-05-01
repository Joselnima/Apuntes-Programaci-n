# Módulo 16 - Environment Variables

En este módulo aprenderás sobre CSS Environment Variables (también conocidas como CSS Custom Properties for Environment Values), una característica avanzada que permite acceder a valores del entorno del navegador y del dispositivo.

## Introducción a Environment Variables

### ¿Qué son las Environment Variables?

Las variables de entorno CSS permiten acceder a información del entorno del navegador, como el área segura del viewport, la orientación del dispositivo, y otras propiedades del entorno de visualización.

```css
/* Accediendo al área segura del viewport */
.element {
  padding-top: env(safe-area-inset-top);
  padding-bottom: env(safe-area-inset-bottom);
  padding-left: env(safe-area-inset-left);
  padding-right: env(safe-area-inset-right);
}
```

## Safe Area Inset Variables

### Área segura para dispositivos con notch

```css
/* Variables de área segura */
.safe-area-top {
  padding-top: env(safe-area-inset-top);
}

.safe-area-bottom {
  padding-bottom: env(safe-area-inset-bottom);
}

.safe-area-left {
  padding-left: env(safe-area-inset-left);
}

.safe-area-right {
  padding-right: env(safe-area-inset-right);
}
```

### Aplicación completa

```css
.full-safe-area {
  padding-top: env(safe-area-inset-top);
  padding-bottom: env(safe-area-inset-bottom);
  padding-left: env(safe-area-inset-left);
  padding-right: env(safe-area-inset-right);
}
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Environment Variables - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: 'Arial', sans-serif;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      min-height: 100vh;
      color: white;
    }

    .container {
      max-width: 1200px;
      margin: 0 auto;
      padding: 20px;
      /* Área segura para dispositivos con notch */
      padding-top: max(20px, env(safe-area-inset-top));
      padding-bottom: max(20px, env(safe-area-inset-bottom));
      padding-left: max(20px, env(safe-area-inset-left));
      padding-right: max(20px, env(safe-area-inset-right));
    }

    .header {
      background: rgba(255, 255, 255, 0.1);
      backdrop-filter: blur(10px);
      border-radius: 15px;
      padding: 30px;
      margin-bottom: 30px;
      text-align: center;
      /* Área segura superior */
      padding-top: max(30px, env(safe-area-inset-top) + 10px);
    }

    .header h1 {
      font-size: 2.5rem;
      margin-bottom: 10px;
    }

    .header p {
      font-size: 1.1rem;
      opacity: 0.9;
    }

    .content-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 20px;
      margin-bottom: 30px;
    }

    .card {
      background: rgba(255, 255, 255, 0.1);
      backdrop-filter: blur(10px);
      border-radius: 15px;
      padding: 25px;
      border: 1px solid rgba(255, 255, 255, 0.2);
    }

    .card h3 {
      margin-bottom: 15px;
      font-size: 1.3rem;
    }

    .card p {
      line-height: 1.6;
      opacity: 0.9;
    }

    .safe-area-demo {
      background: rgba(255, 0, 0, 0.3);
      border: 2px solid red;
      padding: 10px;
      margin: 10px 0;
      border-radius: 5px;
      font-size: 0.9rem;
    }

    .footer {
      text-align: center;
      padding: 20px;
      background: rgba(0, 0, 0, 0.2);
      border-radius: 10px;
      /* Área segura inferior */
      margin-bottom: env(safe-area-inset-bottom);
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>Environment Variables CSS</h1>
      <p>Variables de entorno para diseño responsivo seguro</p>
    </header>

    <div class="content-grid">
      <div class="card">
        <h3>¿Qué son?</h3>
        <p>Las variables de entorno CSS permiten acceder a información del entorno del navegador, como áreas seguras de dispositivos con notch, barras de navegación, y otras propiedades del viewport.</p>
      </div>

      <div class="card">
        <h3>Safe Area Inset</h3>
        <p>Las variables <code>safe-area-inset-*</code> proporcionan el espacio necesario para evitar elementos del sistema como notches, barras de estado, y controles de navegación.</p>
      </div>

      <div class="card">
        <h3>Compatibilidad</h3>
        <p>Soportadas en Safari (iOS 11.2+), Chrome (69+), y otros navegadores modernos. En navegadores sin soporte, las variables se resuelven a 0.</p>
      </div>
    </div>

    <div class="safe-area-demo">
      <strong>Demo de Área Segura:</strong><br>
      Este contenedor respeta las áreas seguras del dispositivo. En un iPhone X o similar, el contenido no se superpone con el notch.
    </div>

    <footer class="footer">
      <p>Este footer también respeta el área segura inferior del dispositivo.</p>
    </footer>
  </div>
</body>
</html>
```

## Viewport Segment Variables

### Variables para segmentos del viewport

```css
/* Variables disponibles (todavía en desarrollo) */
.element {
  /* Altura del segmento de título */
  padding-top: env(titlebar-area-height);

  /* Área de la barra de título */
  margin-top: env(titlebar-area-height);

  /* Área de controles del teclado */
  padding-bottom: env(keyboard-inset-height);
}
```

## Fallbacks y Valores por Defecto

### Proporcionando fallbacks

```css
/* Con valor por defecto */
.element {
  padding-top: env(safe-area-inset-top, 20px);
  padding-bottom: env(safe-area-inset-bottom, 20px);
}

/* Usando max() para mejores fallbacks */
.element {
  padding-top: max(20px, env(safe-area-inset-top));
  padding-bottom: max(20px, env(safe-area-inset-bottom));
}
```

## Aplicaciones Prácticas

### 1. Headers fijos con área segura

```css
.fixed-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  padding-top: env(safe-area-inset-top);
  background: white;
  z-index: 1000;
}
```

### 2. Footers con controles

```css
.bottom-controls {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding-bottom: max(20px, env(safe-area-inset-bottom));
  background: white;
  border-top: 1px solid #eee;
}
```

### 3. Contenido full-screen

```css
.fullscreen-content {
  position: fixed;
  top: env(safe-area-inset-top);
  bottom: env(safe-area-inset-bottom);
  left: env(safe-area-inset-left);
  right: env(safe-area-inset-right);
  overflow: auto;
}
```

## Ejemplo Complejo: App Mobile

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Environment Variables - App Mobile</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: #f5f5f5;
      color: #333;
    }

    .app {
      position: fixed;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      display: flex;
      flex-direction: column;
      background: white;
    }

    .status-bar {
      height: env(safe-area-inset-top, 0px);
      background: #000;
      position: relative;
      z-index: 10;
    }

    .header {
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
      padding: 15px 20px;
      padding-top: max(15px, env(safe-area-inset-top));
      display: flex;
      align-items: center;
      justify-content: space-between;
      box-shadow: 0 2px 10px rgba(0,0,0,0.1);
      position: relative;
      z-index: 5;
    }

    .header h1 {
      font-size: 1.2rem;
      font-weight: 600;
    }

    .menu-btn {
      width: 30px;
      height: 30px;
      background: rgba(255,255,255,0.2);
      border: none;
      border-radius: 5px;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
    }

    .menu-btn::before {
      content: '☰';
      color: white;
      font-size: 16px;
    }

    .main-content {
      flex: 1;
      overflow-y: auto;
      padding: 20px;
      padding-bottom: calc(80px + max(20px, env(safe-area-inset-bottom)));
    }

    .content-card {
      background: white;
      border-radius: 12px;
      padding: 20px;
      margin-bottom: 20px;
      box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    }

    .content-card h3 {
      margin-bottom: 10px;
      color: #333;
    }

    .content-card p {
      line-height: 1.6;
      color: #666;
    }

    .bottom-nav {
      position: fixed;
      bottom: 0;
      left: 0;
      right: 0;
      background: white;
      border-top: 1px solid #e0e0e0;
      padding: 10px 20px;
      padding-bottom: max(10px, env(safe-area-inset-bottom));
      display: flex;
      justify-content: space-around;
      align-items: center;
      z-index: 10;
    }

    .nav-item {
      display: flex;
      flex-direction: column;
      align-items: center;
      padding: 5px;
      border-radius: 8px;
      transition: background-color 0.2s ease;
      cursor: pointer;
      min-width: 60px;
    }

    .nav-item:hover {
      background: #f0f0f0;
    }

    .nav-item.active {
      background: rgba(102, 126, 234, 0.1);
      color: #667eea;
    }

    .nav-icon {
      font-size: 20px;
      margin-bottom: 2px;
    }

    .nav-label {
      font-size: 10px;
      font-weight: 500;
    }

    .safe-area-indicator {
      position: fixed;
      top: 0;
      left: 0;
      right: 0;
      height: env(safe-area-inset-top, 0px);
      background: rgba(255, 0, 0, 0.3);
      pointer-events: none;
      z-index: 1000;
    }

    .safe-area-indicator::after {
      content: 'Área Segura';
      position: absolute;
      top: 2px;
      left: 10px;
      font-size: 10px;
      color: red;
      font-weight: bold;
    }
  </style>
</head>
<body>
  <div class="app">
    <div class="status-bar"></div>

    <header class="header">
      <button class="menu-btn"></button>
      <h1>Mi App</h1>
      <div style="width: 30px;"></div>
    </header>

    <main class="main-content">
      <div class="content-card">
        <h3>Environment Variables en Acción</h3>
        <p>Esta aplicación demuestra cómo usar CSS Environment Variables para crear interfaces que respetan las áreas seguras de los dispositivos modernos.</p>
      </div>

      <div class="content-card">
        <h3>Header Adaptable</h3>
        <p>El header se ajusta automáticamente al área segura superior, evitando superposiciones con notches o barras de estado.</p>
      </div>

      <div class="content-card">
        <h3>Navigation Seguro</h3>
        <p>La barra de navegación inferior respeta el área segura, asegurando que los controles sean accesibles en todos los dispositivos.</p>
      </div>

      <div class="content-card">
        <h3>Contenido Scrollable</h3>
        <p>El contenido principal puede hacer scroll sin interferir con los elementos fijos, manteniendo el padding adecuado para áreas seguras.</p>
      </div>
    </main>

    <nav class="bottom-nav">
      <div class="nav-item active">
        <div class="nav-icon">🏠</div>
        <div class="nav-label">Inicio</div>
      </div>
      <div class="nav-item">
        <div class="nav-icon">🔍</div>
        <div class="nav-label">Buscar</div>
      </div>
      <div class="nav-item">
        <div class="nav-icon">❤️</div>
        <div class="nav-label">Favoritos</div>
      </div>
      <div class="nav-item">
        <div class="nav-icon">👤</div>
        <div class="nav-label">Perfil</div>
      </div>
    </nav>

    <div class="safe-area-indicator"></div>
  </div>
</body>
</html>
```

## Variables Personalizadas con Environment

### Combinando custom properties con environment variables

```css
:root {
  --header-height: 60px;
  --safe-padding: max(20px, env(safe-area-inset-top));
}

.header {
  height: var(--header-height);
  padding-top: var(--safe-padding);
}
```

## Compatibilidad y Polyfills

### Verificando soporte

```css
@supports (padding: env(safe-area-inset-top)) {
  .safe-area-supported {
    padding-top: env(safe-area-inset-top);
  }
}

@supports not (padding: env(safe-area-inset-top)) {
  .safe-area-fallback {
    padding-top: 20px;
  }
}
```

### JavaScript fallback

```javascript
// Detectar soporte de environment variables
const supportsSafeArea = CSS.supports('padding: env(safe-area-inset-top)');

if (!supportsSafeArea) {
  // Aplicar fallbacks con JavaScript
  document.documentElement.style.setProperty('--safe-area-top', '20px');
}
```

## Mejores Prácticas

### 1. Siempre usar fallbacks

```css
.element {
  /* Fallback primero */
  padding: 20px;

  /* Environment variable */
  padding: env(safe-area-inset-top, 20px);
}
```

### 2. Usar max() para valores mínimos

```css
.element {
  padding-top: max(20px, env(safe-area-inset-top));
  padding-bottom: max(20px, env(safe-area-inset-bottom));
}
```

### 3. Considerar diferentes dispositivos

```css
/* Para diferentes tipos de dispositivos */
@media (max-width: 768px) {
  .mobile-safe {
    padding-top: max(15px, env(safe-area-inset-top));
  }
}

@media (min-width: 769px) {
  .desktop-safe {
    padding-top: 20px;
  }
}
```

## Casos de Uso Avanzados

### 1. Progressive Web Apps (PWAs)

```css
.pwa-layout {
  position: fixed;
  top: env(safe-area-inset-top);
  bottom: env(safe-area-inset-bottom);
  left: env(safe-area-inset-left);
  right: env(safe-area-inset-right);
}
```

### 2. Video players fullscreen

```css
.video-fullscreen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  padding-top: env(safe-area-inset-top);
  padding-bottom: env(safe-area-inset-bottom);
}
```

### 3. Modales y overlays

```css
.modal {
  position: fixed;
  top: env(safe-area-inset-top);
  bottom: env(safe-area-inset-bottom);
  left: env(safe-area-inset-left);
  right: env(safe-area-inset-right);
  background: rgba(0,0,0,0.8);
}
```

## Resumen

CSS Environment Variables proporcionan acceso a información crítica del entorno del navegador:

- ✅ **Safe Area Inset**: Variables para áreas seguras en dispositivos con notch
- ✅ **Fallbacks**: Siempre proporcionar valores por defecto con `env(var, fallback)`
- ✅ **max() function**: Combinar con funciones CSS para mejores resultados
- ✅ **Compatibilidad**: Excelente soporte en dispositivos móviles modernos
- ✅ **PWAs**: Esencial para Progressive Web Apps de calidad
- ✅ **Responsive Design**: Diseño responsivo que respeta las limitaciones del hardware

Las Environment Variables son cruciales para crear experiencias móviles nativas en la web, especialmente en la era de los dispositivos con pantallas complejas como iPhone X, Samsung Galaxy con notch, y otros dispositivos con áreas seguras.