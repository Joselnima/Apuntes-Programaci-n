# Módulo 14 - Shapes

En este módulo aprenderás sobre CSS Shapes, una característica avanzada que permite crear diseños de texto que fluyen alrededor de formas no rectangulares, creando layouts visualmente atractivos y modernos.

## Introducción a CSS Shapes

### ¿Qué son las CSS Shapes?

CSS Shapes permiten que el contenido de texto fluya alrededor de formas geométricas complejas, en lugar de solo rectángulos. Esto crea diseños más dinámicos y visualmente interesantes.

```css
/* Texto fluyendo alrededor de un círculo */
.shape-circle {
  shape-outside: circle(50%);
}

/* Texto fluyendo alrededor de una elipse */
.shape-ellipse {
  shape-outside: ellipse(50% 30%);
}
```

## Shape Functions Básicas

### Circle

```css
.element {
  shape-outside: circle(50%);
  width: 200px;
  height: 200px;
  float: left;
}
```

### Ellipse

```css
.element {
  shape-outside: ellipse(50% 30%);
  width: 300px;
  height: 200px;
  float: left;
}
```

### Inset Rectangle

```css
.element {
  shape-outside: inset(20px 10px 30px 15px);
  width: 200px;
  height: 150px;
  float: left;
}
```

### Polygon

```css
.element {
  shape-outside: polygon(0 0, 100% 0, 100% 100%, 0 100%);
  width: 200px;
  height: 200px;
  float: left;
}
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Shapes - Básico</title>
  <style>
    .container {
      max-width: 800px;
      margin: 0 auto;
      padding: 20px;
    }

    .shape-circle {
      width: 200px;
      height: 200px;
      background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
      shape-outside: circle(50%);
      float: left;
      margin-right: 20px;
      margin-bottom: 10px;
      border-radius: 50%;
    }

    .content {
      font-family: 'Arial', sans-serif;
      line-height: 1.6;
    }

    .shape-ellipse {
      width: 250px;
      height: 150px;
      background: linear-gradient(45deg, #a8e6cf, #ffd3a5);
      shape-outside: ellipse(50% 40%);
      float: right;
      margin-left: 20px;
      margin-bottom: 10px;
      border-radius: 50%;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>CSS Shapes - Texto Fluyendo</h1>

    <div class="shape-circle"></div>

    <div class="content">
      <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

      <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
    </div>

    <div class="shape-ellipse"></div>

    <div class="content">
      <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>

      <p>Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.</p>
    </div>
  </div>
</body>
</html>
```

## Shape Margin

### Añadiendo margen a las shapes

```css
.element {
  shape-outside: circle(50%);
  shape-margin: 20px;
  width: 200px;
  height: 200px;
  float: left;
}
```

## Shapes con Imágenes

### Usando imágenes como shapes

```css
.image-shape {
  width: 300px;
  height: 200px;
  shape-outside: url('shape-image.png');
  float: left;
  margin-right: 20px;
}
```

### Shape from image con threshold

```css
.image-shape {
  shape-outside: url('shape-image.png');
  shape-image-threshold: 0.5;
  float: left;
}
```

## Polygon Shapes Avanzadas

### Triángulo

```css
.triangle-shape {
  width: 0;
  height: 0;
  shape-outside: polygon(50% 0%, 0% 100%, 100% 100%);
  float: left;
  margin-right: 20px;
}
```

### Forma compleja

```css
.complex-shape {
  width: 250px;
  height: 200px;
  background: linear-gradient(45deg, #667eea, #764ba2);
  shape-outside: polygon(
    0% 0%,
    70% 0%,
    100% 30%,
    100% 70%,
    70% 100%,
    0% 100%,
    0% 70%,
    30% 30%
  );
  float: left;
  margin-right: 20px;
}
```

## Clip Path vs Shape Outside

### Diferencias clave

```css
/* Clip-path: Recorta el elemento visualmente */
.element {
  clip-path: circle(50%);
}

/* Shape-outside: Afecta el flujo del contenido alrededor */
.element {
  shape-outside: circle(50%);
}
```

## Ejemplo Complejo con Múltiples Shapes

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Shapes - Complejo</title>
  <style>
    .container {
      max-width: 1000px;
      margin: 0 auto;
      padding: 20px;
    }

    .shape-grid {
      display: grid;
      grid-template-columns: 1fr 2fr;
      gap: 30px;
      margin-bottom: 40px;
    }

    .shape-element {
      width: 200px;
      height: 200px;
      shape-outside: polygon(
        50% 0%,
        100% 25%,
        100% 75%,
        50% 100%,
        0% 75%,
        0% 25%
      );
      shape-margin: 15px;
      float: left;
      margin-right: 20px;
      margin-bottom: 15px;
      background: linear-gradient(45deg, #ff9a9e, #fecfef);
      border-radius: 10px;
    }

    .shape-circle-large {
      width: 150px;
      height: 150px;
      shape-outside: circle(50%);
      shape-margin: 20px;
      float: right;
      margin-left: 20px;
      margin-bottom: 15px;
      background: linear-gradient(45deg, #a8e6cf, #ffd3a5);
      border-radius: 50%;
    }

    .content {
      font-family: 'Georgia', serif;
      line-height: 1.7;
      font-size: 16px;
    }

    .content p {
      margin-bottom: 20px;
    }

    .shape-rectangle {
      width: 100px;
      height: 300px;
      shape-outside: inset(20px);
      shape-margin: 10px;
      float: left;
      margin-right: 20px;
      background: linear-gradient(45deg, #667eea, #764ba2);
      border-radius: 5px;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>CSS Shapes - Diseño Complejo</h1>

    <div class="shape-grid">
      <div>
        <div class="shape-element"></div>
        <div class="shape-circle-large"></div>
      </div>

      <div class="content">
        <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>

        <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      </div>
    </div>

    <div class="shape-rectangle"></div>

    <div class="content">
      <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>

      <p>Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.</p>

      <p>Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.</p>
    </div>
  </div>
</body>
</html>
```

## Shape Inside

### Controlando el flujo interno

```css
.element {
  shape-inside: circle(50%);
  width: 300px;
  height: 300px;
}
```

## Compatibilidad y Fallbacks

### Detectando soporte

```css
@supports (shape-outside: circle()) {
  .shape-supported {
    shape-outside: circle(50%);
  }
}

.shape-fallback {
  /* Fallback para navegadores sin soporte */
  border-radius: 50%;
}
```

## Casos de Uso Prácticos

### 1. Diseño de revistas
```css
.magazine-layout {
  shape-outside: polygon(0 0, 100% 0, 80% 100%, 0 100%);
}
```

### 2. Avatares circulares
```css
.avatar {
  shape-outside: circle(50%);
  clip-path: circle(50%);
}
```

### 3. Elementos decorativos
```css
.decorative-shape {
  shape-outside: ellipse(60% 40%);
  background: linear-gradient(45deg, #ff6b6b, #4ecdc4);
}
```

## Performance Considerations

### Optimizaciones

```css
/* Evitar shapes complejas en elementos que cambian frecuentemente */
.simple-shape {
  shape-outside: circle(50%);
  /* Mejor performance que polygons complejos */
}

/* Usar will-change para animaciones */
.shape-animated {
  will-change: shape-outside;
  transition: shape-outside 0.3s ease;
}
```

## Resumen

CSS Shapes revolucionan el diseño web al permitir layouts no rectangulares:

- ✅ **Shape-outside**: Controla el flujo de texto alrededor de formas
- ✅ **Funciones geométricas**: circle(), ellipse(), polygon(), inset()
- ✅ **Shape-margin**: Añade espacio alrededor de las shapes
- ✅ **Compatibilidad**: Funciona en navegadores modernos
- ✅ **Performance**: Optimizado para uso en producción

Las shapes son ideales para diseños creativos, layouts de revistas, y elementos visualmente atractivos que van más allá de los tradicionales diseños basados en cajas rectangulares.