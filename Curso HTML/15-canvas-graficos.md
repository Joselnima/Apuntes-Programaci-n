# Módulo 15 - Canvas y gráficos

En este módulo aprenderás a dibujar gráficos y animaciones directamente en HTML usando el elemento `<canvas>`.

## ¿Qué es Canvas?

`<canvas>` es un elemento HTML5 que proporciona un área rectangular para dibujar gráficos usando JavaScript. Es como un lienzo en blanco donde puedes pintar con código.

```html
<canvas id="miCanvas" width="400" height="300">
  Tu navegador no soporta Canvas.
</canvas>
```

## Contexto de dibujo

Para dibujar, necesitas obtener el contexto 2D:

```javascript
const canvas = document.getElementById('miCanvas');
const ctx = canvas.getContext('2d');
```

## Dibujar formas básicas

### Rectángulos

```javascript
// Rectángulo relleno
ctx.fillStyle = 'blue';
ctx.fillRect(10, 10, 100, 50);

// Rectángulo con borde
ctx.strokeStyle = 'red';
ctx.lineWidth = 2;
ctx.strokeRect(120, 10, 100, 50);

// Rectángulo vacío
ctx.clearRect(50, 30, 50, 20);
```

### Círculos y arcos

```javascript
// Círculo relleno
ctx.beginPath();
ctx.arc(200, 100, 50, 0, 2 * Math.PI);
ctx.fillStyle = 'green';
ctx.fill();

// Arco (medialuna)
ctx.beginPath();
ctx.arc(300, 100, 50, 0, Math.PI);
ctx.fillStyle = 'orange';
ctx.fill();
```

### Líneas y caminos

```javascript
ctx.beginPath();
ctx.moveTo(10, 150);
ctx.lineTo(100, 200);
ctx.lineTo(200, 150);
ctx.closePath();
ctx.strokeStyle = 'purple';
ctx.lineWidth = 3;
ctx.stroke();
```

## Texto en Canvas

```javascript
ctx.font = '30px Arial';
ctx.fillStyle = 'black';
ctx.fillText('Hola Canvas!', 10, 250);

ctx.strokeStyle = 'blue';
ctx.strokeText('Texto con borde', 10, 280);
```

## Imágenes en Canvas

```javascript
const img = new Image();
img.onload = function() {
  ctx.drawImage(img, 10, 10, 100, 100);
};
img.src = 'imagen.jpg';
```

## Gradientes y patrones

### Gradiente lineal

```javascript
const gradient = ctx.createLinearGradient(0, 0, 200, 0);
gradient.addColorStop(0, 'red');
gradient.addColorStop(1, 'blue');
ctx.fillStyle = gradient;
ctx.fillRect(10, 10, 200, 100);
```

### Gradiente radial

```javascript
const radialGradient = ctx.createRadialGradient(100, 100, 0, 100, 100, 50);
radialGradient.addColorStop(0, 'white');
radialGradient.addColorStop(1, 'black');
ctx.fillStyle = radialGradient;
ctx.fillRect(50, 50, 100, 100);
```

## Transformaciones

```javascript
// Trasladar
ctx.translate(100, 100);

// Rotar
ctx.rotate(Math.PI / 4);

// Escalar
ctx.scale(2, 1);

// Dibujar después de transformación
ctx.fillRect(0, 0, 50, 50);

// Restaurar
ctx.setTransform(1, 0, 0, 1, 0, 0);
```

## Animaciones simples

```javascript
let x = 0;
function animate() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  ctx.fillRect(x, 100, 50, 50);
  x += 2;
  if (x > canvas.width) x = 0;
  requestAnimationFrame(animate);
}
animate();
```

## SVG vs Canvas

- **Canvas**: Basado en píxeles, mejor para juegos y gráficos complejos
- **SVG**: Basado en vectores, mejor para gráficos escalables e interactivos

### SVG inline básico

```html
<svg width="200" height="200">
  <circle cx="100" cy="100" r="50" fill="red" />
  <rect x="50" y="50" width="100" height="50" fill="blue" />
</svg>
```

## Ejemplo completo: dibujo interactivo

```html
<!DOCTYPE html>
<html>
<body>
  <canvas id="dibujo" width="400" height="300" style="border:1px solid black;"></canvas>
  <br>
  <button onclick="clearCanvas()">Limpiar</button>

  <script>
    const canvas = document.getElementById('dibujo');
    const ctx = canvas.getContext('2d');
    let drawing = false;

    canvas.addEventListener('mousedown', startDrawing);
    canvas.addEventListener('mousemove', draw);
    canvas.addEventListener('mouseup', stopDrawing);

    function startDrawing(e) {
      drawing = true;
      ctx.beginPath();
      ctx.moveTo(e.offsetX, e.offsetY);
    }

    function draw(e) {
      if (!drawing) return;
      ctx.lineTo(e.offsetX, e.offsetY);
      ctx.stroke();
    }

    function stopDrawing() {
      drawing = false;
    }

    function clearCanvas() {
      ctx.clearRect(0, 0, canvas.width, canvas.height);
    }
  </script>
</body>
</html>
```

## Resumen

Canvas permite crear gráficos dinámicos y animaciones en HTML. Es poderoso pero requiere JavaScript. Úsalo cuando necesites gráficos personalizados o juegos en la web.</content>
<parameter name="filePath">e:\Lenguajes de programacion\Curso HTML\15-canvas-graficos.md