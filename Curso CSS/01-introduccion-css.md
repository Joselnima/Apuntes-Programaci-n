# Módulo 01 - Introducción a CSS

En este módulo aprenderás qué es CSS, para qué sirve y cómo se relaciona con HTML para crear páginas web visualmente atractivas.

## ¿Qué es CSS?

**CSS (Cascading Style Sheets)** es un lenguaje de hojas de estilo en cascada que se utiliza para describir la presentación visual de un documento HTML. Mientras HTML estructura el contenido, CSS se encarga de cómo se ve ese contenido.

### Analogía simple
- **HTML** es como el esqueleto y órganos de una persona (estructura)
- **CSS** es como la piel, ropa y apariencia (estilo)
- **JavaScript** es como los músculos y movimientos (interactividad)

## ¿Para qué sirve CSS?

CSS permite controlar:
- **Colores**: Texto, fondos, bordes
- **Tipografía**: Fuentes, tamaños, estilos
- **Layout**: Posicionamiento de elementos
- **Espaciado**: Márgenes, padding, dimensiones
- **Efectos visuales**: Sombras, gradientes, animaciones
- **Responsive design**: Adaptación a diferentes pantallas

## Cómo funciona CSS

### 1. Sintaxis básica

```css
selector {
  propiedad: valor;
}
```

**Ejemplo:**
```css
h1 {
  color: blue;
  font-size: 24px;
}
```

### 2. Formas de aplicar CSS

#### a) CSS inline (en línea)
```html
<h1 style="color: blue; font-size: 24px;">Título</h1>
```

#### b) CSS interno (en el documento)
```html
<head>
  <style>
    h1 {
      color: blue;
      font-size: 24px;
    }
  </style>
</head>
```

#### c) CSS externo (archivo separado) - RECOMENDADO
```html
<head>
  <link rel="stylesheet" href="styles.css">
</head>
```

**Archivo styles.css:**
```css
h1 {
  color: blue;
  font-size: 24px;
}
```

## Ventajas de usar CSS externo

- ✅ **Reutilizable**: Un archivo CSS para múltiples páginas
- ✅ **Mantenible**: Cambios en un solo lugar
- ✅ **Rendimiento**: Mejor caching del navegador
- ✅ **Separación de responsabilidades**: HTML para estructura, CSS para estilo

## El concepto de "Cascada"

La "C" en CSS significa **Cascading** (en cascada). Esto significa que las reglas CSS se aplican en un orden específico de prioridad:

### Orden de prioridad (de menor a mayor):
1. **Reglas por defecto del navegador**
2. **CSS externo**
3. **CSS interno**
4. **CSS inline**
5. **!important** (último recurso)

### Especificidad de selectores
Cuando dos reglas afectan al mismo elemento, gana la más específica:
- Selectores de etiqueta: `h1` (especificidad baja)
- Selectores de clase: `.clase` (especificidad media)
- Selectores de ID: `#id` (especificidad alta)

## Herramientas para trabajar con CSS

### Editores de código
- **VS Code** con extensiones CSS
- **Sublime Text**
- **Atom**

### Navegadores con herramientas de desarrollo
- **Chrome DevTools** (F12)
- **Firefox Developer Tools**
- **Edge DevTools**

### Recursos de aprendizaje
- **MDN Web Docs**: Documentación oficial
- **CSS-Tricks**: Tutoriales y guías
- **Can I Use**: Compatibilidad de navegadores

## Primer ejemplo práctico

Vamos a crear nuestra primera página con CSS:

**index.html:**
```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Mi primera página con CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <header>
    <h1>¡Hola CSS!</h1>
    <p>Esta página tiene estilo gracias a CSS.</p>
  </header>

  <main>
    <section>
      <h2>¿Qué aprenderemos?</h2>
      <ul>
        <li>Colores y fuentes</li>
        <li>Layout y posicionamiento</li>
        <li>Animaciones y efectos</li>
      </ul>
    </section>
  </main>
</body>
</html>
```

**styles.css:**
```css
/* Estilos para el body */
body {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  margin: 0;
  padding: 20px;
  background-color: #f4f4f4;
}

/* Estilos para el header */
header {
  background-color: #333;
  color: white;
  padding: 20px;
  text-align: center;
  border-radius: 8px;
}

/* Estilos para h1 */
h1 {
  margin: 0;
  font-size: 2.5em;
}

/* Estilos para párrafos */
p {
  color: #666;
  font-size: 1.1em;
}

/* Estilos para listas */
ul {
  background-color: white;
  padding: 15px;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

li {
  margin-bottom: 10px;
}
```

## Conceptos clave aprendidos

- ✅ CSS separa el estilo del contenido HTML
- ✅ Las reglas CSS tienen selector, propiedad y valor
- ✅ CSS externo es la mejor práctica
- ✅ La cascada determina qué reglas se aplican
- ✅ Los selectores tienen diferentes niveles de especificidad

## Resumen

CSS es el lenguaje que transforma documentos HTML simples en páginas web visualmente atractivas. En este módulo aprendiste los fundamentos: qué es CSS, cómo funciona, cómo aplicarlo y por qué es importante usar buenas prácticas desde el principio.

En el próximo módulo profundizaremos en la sintaxis y selectores de CSS.