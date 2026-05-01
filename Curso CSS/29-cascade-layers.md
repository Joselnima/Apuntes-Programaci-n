# Módulo 32 - Cascade Layers

En este módulo aprenderás sobre Cascade Layers (Capas de Cascada), una característica revolucionaria de CSS que permite organizar y controlar la especificidad de los estilos de manera más predecible y mantenible.

## Introducción a Cascade Layers

### ¿Qué son las Cascade Layers?

Cascade Layers permiten agrupar reglas CSS en capas ordenadas, donde las reglas de capas posteriores tienen mayor especificidad que las de capas anteriores, independientemente de la especificidad de los selectores individuales.

```css
@layer base, components, utilities;

/* Capa base - menor prioridad */
@layer base {
  .button {
    padding: 1rem;
    background: blue;
  }
}

/* Capa components - prioridad media */
@layer components {
  .button {
    background: green; /* Gana sobre base */
  }
}

/* Capa utilities - mayor prioridad */
@layer utilities {
  .button {
    background: red; /* Gana sobre components */
  }
}
```

## Sintaxis Básica

### Declaración de capas

```css
/* Declarar capas al inicio */
@layer base, theme, components, utilities;

/* O definir capas en línea */
@layer base {
  /* Reglas CSS */
}

@layer theme {
  /* Reglas CSS */
}
```

### Capas anidadas

```css
@layer framework {
  @layer base {
    /* Reglas base del framework */
  }

  @layer theme {
    /* Tema del framework */
  }
}
```

## Orden de Prioridad

### Reglas de precedencia

1. **Capas posteriores** tienen mayor prioridad
2. **Especificidad del selector** dentro de la misma capa
3. **Orden de aparición** dentro de la misma capa

```css
@layer base, components;

/* 1. Capa base - menor prioridad */
@layer base {
  .button { background: blue; }    /* Especificidad: 0,1,0 */
  div.button { background: cyan; } /* Especificidad: 0,1,1 */
}

/* 2. Capa components - mayor prioridad */
@layer components {
  .button { background: green; }   /* Gana por capa */
}
```

## Ejemplo Interactivo Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Cascade Layers - Básico</title>
  <style>
    /* Declarar capas */
    @layer base, components, utilities;

    /* Capa base - reset y estilos base */
    @layer base {
      * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
      }

      body {
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
        background: #f8f9fa;
        color: #333;
        line-height: 1.6;
        padding: 20px;
      }

      .container {
        max-width: 800px;
        margin: 0 auto;
        background: white;
        border-radius: 12px;
        padding: 30px;
        box-shadow: 0 4px 20px rgba(0,0,0,0.1);
      }

      h1 {
        color: #333;
        margin-bottom: 20px;
        font-size: 2rem;
      }

      .button {
        display: inline-block;
        padding: 12px 24px;
        border: none;
        border-radius: 6px;
        text-decoration: none;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.2s ease;
      }
    }

    /* Capa components - componentes reutilizables */
    @layer components {
      .card {
        background: #f8f9fa;
        border: 1px solid #e9ecef;
        border-radius: 8px;
        padding: 20px;
        margin-bottom: 20px;
      }

      .button {
        background: #007bff;
        color: white;
      }

      .button:hover {
        background: #0056b3;
        transform: translateY(-1px);
      }

      .button.secondary {
        background: #6c757d;
      }

      .button.secondary:hover {
        background: #545b62;
      }
    }

    /* Capa utilities - utilidades de alto nivel */
    @layer utilities {
      .text-center {
        text-align: center;
      }

      .mb-3 {
        margin-bottom: 1rem;
      }

      .bg-primary {
        background: #007bff !important;
      }

      .text-white {
        color: white !important;
      }

      .rounded {
        border-radius: 50px;
      }
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Cascade Layers - Ejemplo Básico</h1>

    <div class="card">
      <h2>Entendiendo las Capas</h2>
      <p>Las reglas CSS están organizadas en capas con diferente prioridad:</p>
      <ul>
        <li><strong>Base:</strong> Estilos base y reset</li>
        <li><strong>Components:</strong> Componentes reutilizables</li>
        <li><strong>Utilities:</strong> Utilidades de alto nivel (mayor prioridad)</li>
      </ul>
    </div>

    <div class="card">
      <h2>Botones de Ejemplo</h2>
      <p>Haz clic en los botones para ver cómo las capas afectan los estilos:</p>

      <div style="display: flex; gap: 10px; flex-wrap: wrap; margin: 20px 0;">
        <button class="button">Botón Base</button>
        <button class="button secondary">Botón Secondary</button>
        <button class="button bg-primary text-white rounded">Botón Utilities</button>
      </div>

      <div style="background: #f8f9fa; padding: 15px; border-radius: 6px; border: 1px solid #e9ecef;">
        <strong>Nota:</strong> El último botón combina estilos de diferentes capas. La utilidad <code>bg-primary</code> gana porque está en la capa <code>utilities</code>.
      </div>
    </div>

    <div class="card text-center">
      <h2>Texto Centrado con Utility</h2>
      <p class="mb-3">Este texto está centrado usando la clase <code>text-center</code> de la capa utilities.</p>
      <button class="button">Acción</button>
    </div>
  </div>
</body>
</html>
```

## Arquitectura de CSS con Layers

### Patrón ITCSS con Layers

```css
@layer reset, base, objects, components, utilities;

/* 1. Reset - normalización */
@layer reset {
  *, *::before, *::after {
    box-sizing: border-box;
  }
}

/* 2. Base - elementos HTML base */
@layer base {
  body { font-family: system-ui; }
  h1, h2, h3 { line-height: 1.2; }
}

/* 3. Objects - patrones de diseño */
@layer objects {
  .container { max-width: 1200px; margin: 0 auto; }
  .grid { display: grid; gap: 1rem; }
}

/* 4. Components - componentes específicos */
@layer components {
  .button { padding: 0.5rem 1rem; }
  .card { border-radius: 0.5rem; }
}

/* 5. Utilities - helpers de alto nivel */
@layer utilities {
  .text-center { text-align: center; }
  .hidden { display: none; }
}
```

### Framework CSS con Layers

```css
@layer bootstrap, theme, overrides;

/* Framework base */
@layer bootstrap {
  .btn {
    display: inline-block;
    padding: 0.375rem 0.75rem;
    border: 1px solid transparent;
    border-radius: 0.25rem;
  }

  .btn-primary {
    background: #007bff;
    border-color: #007bff;
  }
}

/* Tema personalizado */
@layer theme {
  .btn-primary {
    background: #667eea;
    border-color: #667eea;
  }

  .btn-primary:hover {
    background: #5a67d8;
  }
}

/* Overrides específicos */
@layer overrides {
  .btn-primary {
    background: #764ba2;
    font-weight: bold;
  }
}
```

## Ejemplo Avanzado: Sistema de Diseño

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Sistema de Diseño con Cascade Layers</title>
  <style>
    /* Arquitectura del sistema de diseño */
    @layer tokens, base, patterns, components, utilities, overrides;

    /* 1. Design Tokens */
    @layer tokens {
      :root {
        /* Colores */
        --color-primary: #667eea;
        --color-secondary: #764ba2;
        --color-accent: #f093fb;
        --color-neutral-50: #f8f9fa;
        --color-neutral-100: #e9ecef;
        --color-neutral-900: #212529;

        /* Espaciado */
        --space-xs: 0.25rem;
        --space-sm: 0.5rem;
        --space-md: 1rem;
        --space-lg: 1.5rem;
        --space-xl: 3rem;

        /* Tipografía */
        --font-size-sm: 0.875rem;
        --font-size-base: 1rem;
        --font-size-lg: 1.25rem;
        --font-size-xl: 1.5rem;
        --font-size-2xl: 2rem;

        /* Bordes */
        --border-radius-sm: 0.25rem;
        --border-radius-md: 0.5rem;
        --border-radius-lg: 1rem;

        /* Sombras */
        --shadow-sm: 0 1px 2px rgba(0,0,0,0.05);
        --shadow-md: 0 4px 6px rgba(0,0,0,0.07);
        --shadow-lg: 0 10px 15px rgba(0,0,0,0.1);
      }
    }

    /* 2. Base Layer */
    @layer base {
      * {
        box-sizing: border-box;
      }

      body {
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
        background: var(--color-neutral-50);
        color: var(--color-neutral-900);
        line-height: 1.6;
        margin: 0;
        padding: var(--space-xl);
      }

      h1, h2, h3, h4 {
        line-height: 1.2;
        margin-bottom: var(--space-md);
      }

      h1 { font-size: var(--font-size-2xl); }
      h2 { font-size: var(--font-size-xl); }
      h3 { font-size: var(--font-size-lg); }

      p {
        margin-bottom: var(--space-md);
      }
    }

    /* 3. Patterns Layer */
    @layer patterns {
      .container {
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 var(--space-md);
      }

      .grid {
        display: grid;
        gap: var(--space-md);
      }

      .grid-2 { grid-template-columns: repeat(2, 1fr); }
      .grid-3 { grid-template-columns: repeat(3, 1fr); }
      .grid-4 { grid-template-columns: repeat(4, 1fr); }

      .flex {
        display: flex;
        gap: var(--space-md);
      }

      .flex-column {
        flex-direction: column;
      }

      .justify-center {
        justify-content: center;
      }

      .align-center {
        align-items: center;
      }
    }

    /* 4. Components Layer */
    @layer components {
      .card {
        background: white;
        border-radius: var(--border-radius-lg);
        padding: var(--space-lg);
        box-shadow: var(--shadow-md);
        border: 1px solid var(--color-neutral-100);
      }

      .button {
        display: inline-flex;
        align-items: center;
        gap: var(--space-sm);
        padding: var(--space-sm) var(--space-md);
        background: var(--color-primary);
        color: white;
        border: none;
        border-radius: var(--border-radius-md);
        font-size: var(--font-size-base);
        font-weight: 500;
        cursor: pointer;
        transition: all 0.2s ease;
        text-decoration: none;
      }

      .button:hover {
        background: color-mix(in srgb, var(--color-primary) 90%, black);
        transform: translateY(-1px);
        box-shadow: var(--shadow-lg);
      }

      .button.secondary {
        background: var(--color-secondary);
      }

      .button.secondary:hover {
        background: color-mix(in srgb, var(--color-secondary) 90%, black);
      }

      .button.outline {
        background: transparent;
        color: var(--color-primary);
        border: 2px solid var(--color-primary);
      }

      .button.outline:hover {
        background: var(--color-primary);
        color: white;
      }

      .input {
        width: 100%;
        padding: var(--space-sm) var(--space-md);
        border: 1px solid var(--color-neutral-100);
        border-radius: var(--border-radius-md);
        font-size: var(--font-size-base);
        transition: border-color 0.2s ease;
      }

      .input:focus {
        outline: none;
        border-color: var(--color-primary);
        box-shadow: 0 0 0 3px color-mix(in srgb, var(--color-primary) 20%, transparent);
      }

      .badge {
        display: inline-block;
        padding: var(--space-xs) var(--space-sm);
        background: var(--color-accent);
        color: white;
        border-radius: var(--border-radius-sm);
        font-size: var(--font-size-sm);
        font-weight: 500;
      }
    }

    /* 5. Utilities Layer */
    @layer utilities {
      .text-center { text-align: center; }
      .text-left { text-align: left; }
      .text-right { text-align: right; }

      .font-bold { font-weight: bold; }
      .font-normal { font-weight: normal; }

      .bg-primary { background: var(--color-primary) !important; }
      .bg-secondary { background: var(--color-secondary) !important; }
      .text-primary { color: var(--color-primary) !important; }
      .text-white { color: white !important; }

      .m-0 { margin: 0 !important; }
      .mb-1 { margin-bottom: var(--space-sm) !important; }
      .mb-2 { margin-bottom: var(--space-md) !important; }
      .mb-3 { margin-bottom: var(--space-lg) !important; }

      .p-1 { padding: var(--space-sm) !important; }
      .p-2 { padding: var(--space-md) !important; }
      .p-3 { padding: var(--space-lg) !important; }

      .rounded { border-radius: var(--border-radius-md) !important; }
      .rounded-full { border-radius: 50% !important; }

      .shadow { box-shadow: var(--shadow-md) !important; }
      .shadow-lg { box-shadow: var(--shadow-lg) !important; }

      .hidden { display: none !important; }
      .block { display: block !important; }
      .inline-block { display: inline-block !important; }
      .flex { display: flex !important; }
      .grid { display: grid !important; }
    }

    /* 6. Overrides Layer - para casos específicos */
    @layer overrides {
      .hero-button {
        font-size: var(--font-size-lg);
        padding: var(--space-md) var(--space-xl);
        background: linear-gradient(45deg, var(--color-primary), var(--color-secondary));
      }

      .special-card {
        background: linear-gradient(135deg, var(--color-neutral-50), white);
        border: 2px solid var(--color-primary);
      }
    }

    /* Responsive utilities */
    @media (max-width: 768px) {
      @layer utilities {
        .grid-2, .grid-3, .grid-4 {
          grid-template-columns: 1fr;
        }
      }
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="text-center mb-3">
      <h1>Sistema de Diseño con Cascade Layers</h1>
      <p>Un sistema modular y escalable construido con capas CSS organizadas</p>
    </header>

    <div class="grid grid-2 mb-3">
      <div class="card">
        <h3>🏗️ Arquitectura Modular</h3>
        <p>Las capas permiten organizar el CSS de manera lógica y predecible. Cada capa tiene un propósito específico y un nivel de prioridad definido.</p>
        <span class="badge">Design System</span>
      </div>

      <div class="card">
        <h3>🎨 Design Tokens</h3>
        <p>Variables CSS centralizadas que definen colores, espaciado, tipografía y otros valores base del sistema de diseño.</p>
        <span class="badge">Tokens</span>
      </div>
    </div>

    <div class="card special-card mb-3">
      <h3>🧩 Componentes Reutilizables</h3>
      <p>Componentes construidos sobre patrones consistentes, fáciles de mantener y extender.</p>

      <div class="flex flex-column" style="gap: 1rem; margin-top: 1.5rem;">
        <div class="flex align-center">
          <button class="button">Botón Primario</button>
          <button class="button secondary">Secundario</button>
          <button class="button outline">Outline</button>
        </div>

        <div>
          <input type="text" class="input" placeholder="Campo de entrada" style="max-width: 300px;">
        </div>
      </div>
    </div>

    <div class="grid grid-3 mb-3">
      <div class="card text-center">
        <h3>📱 Responsive</h3>
        <p>El sistema se adapta automáticamente a diferentes tamaños de pantalla.</p>
        <button class="button">Ver más</button>
      </div>

      <div class="card text-center">
        <h3>🎯 Utilities</h3>
        <p>Clases de utilidad de alto nivel para ajustes rápidos y específicos.</p>
        <div class="badge bg-primary text-white">Nuevo</div>
      </div>

      <div class="card text-center">
        <h3>🔧 Mantenible</h3>
        <p>Código organizado que facilita el mantenimiento y la colaboración en equipo.</p>
        <button class="button secondary">Aprender más</button>
      </div>
    </div>

    <div class="card text-center">
      <h3>🚀 Listo para Producción</h3>
      <p>Este sistema de diseño está construido con las mejores prácticas modernas de CSS y es completamente escalable.</p>
      <button class="hero-button">Comenzar Proyecto</button>
    </div>
  </div>
</body>
</html>
```

## Técnicas Avanzadas

### Layers condicionales

```css
@layer base {
  .component { background: white; }
}

/* Solo aplicar en modo oscuro */
@media (prefers-color-scheme: dark) {
  @layer base {
    .component { background: #333; }
  }
}
```

### Layers con CSS-in-JS

```javascript
// styled-components con layers
const Button = styled.button`
  @layer components {
    background: blue;
    padding: 1rem;
  }
`;

// Emotion
const buttonStyles = css`
  @layer components {
    background: green;
    border-radius: 4px;
  }
`;
```

### Debugging de layers

```css
/* Inspeccionar capas en DevTools */
@layer debug {
  * {
    outline: 1px solid red;
  }
}

/* Temporal - remover en producción */
@layer debug;
```

## Mejores Prácticas

### 1. Convenciones de nomenclatura

```css
@layer
  reset,           /* Normalización */
  base,            /* Elementos HTML base */
  layout,          /* Layout y grid */
  patterns,        /* Patrones de diseño */
  components,      /* Componentes específicos */
  utilities,       /* Utilidades de alto nivel */
  overrides;       /* Overrides específicos */
```

### 2. Evitar !important en capas

```css
/* ❌ Mal - !important en capa baja */
@layer base {
  .text { color: black !important; }
}

/* ✅ Bien - usar capas para prioridad */
@layer utilities {
  .text-red { color: red; }
}
```

### 3. Imports organizados

```css
/* importar con capas */
@import url('reset.css') layer(reset);
@import url('base.css') layer(base);
@import url('components.css') layer(components);
```

## Compatibilidad y Fallbacks

### Detectar soporte

```javascript
function supportsCascadeLayers() {
  try {
    CSS.supports('@layer base {}');
    return true;
  } catch {
    return false;
  }
}
```

### Fallback para navegadores antiguos

```css
/* Fallback sin layers */
.button {
  background: blue;
  padding: 1rem;
}

/* Modern browsers con layers */
@supports (@layer) {
  @layer components {
    .button {
      background: green; /* Gana sobre el fallback */
    }
  }
}
```

## Casos de Uso Reales

### 1. Framework CSS

```css
/* Bootstrap con layers */
@layer bootstrap {
  .btn { /* estilos base */ }
}

@layer theme {
  .btn-primary { /* tema personalizado */ }
}

@layer project {
  .btn { /* overrides del proyecto */ }
}
```

### 2. Component Libraries

```css
@layer library {
  @layer base { /* estilos base */ }
  @layer components { /* componentes */ }
  @layer utilities { /* utilities */ }
}
```

### 3. Tema oscuro/claro

```css
@layer theme {
  :root { /* tema base */ }
}

@layer theme.dark {
  :root { /* overrides para tema oscuro */ }
}
```

## Resumen

Cascade Layers revolucionan la organización CSS:

- ✅ **Predecibilidad**: Control total sobre la cascada
- ✅ **Mantenibilidad**: Código organizado por capas lógicas
- ✅ **Escalabilidad**: Fácil de extender y modificar
- ✅ **Colaboración**: Equipos pueden trabajar en capas diferentes
- ✅ **Performance**: Mejor organización del CSS

Cascade Layers permiten construir sistemas de diseño robustos y mantenibles, eliminando muchos de los problemas tradicionales de especificidad CSS.