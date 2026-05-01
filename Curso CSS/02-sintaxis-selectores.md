# Módulo 02 - Sintaxis y selectores básicos

En este módulo aprenderás la sintaxis completa de CSS y todos los tipos de selectores básicos para aplicar estilos a elementos HTML específicos.

## Sintaxis de CSS

### Estructura básica de una regla CSS

```css
selector {
  propiedad: valor;
  propiedad2: valor2;
}
```

**Ejemplo completo:**
```css
h1 {
  color: blue;
  font-size: 24px;
  margin-bottom: 20px;
}
```

### Componentes de una regla CSS

1. **Selector**: Identifica qué elementos HTML se van a estilizar
2. **Declaración**: Propiedad + valor entre llaves
3. **Propiedad**: Lo que quieres cambiar (color, tamaño, etc.)
4. **Valor**: El valor específico de la propiedad

## Comentarios en CSS

```css
/* Este es un comentario de una línea */

/*
Este es un comentario
de múltiples líneas
útil para explicar
secciones complejas
*/
```

## Selectores básicos

### 1. Selector universal
Selecciona todos los elementos:
```css
* {
  margin: 0;
  padding: 0;
}
```

### 2. Selectores de tipo (etiqueta)
Selecciona elementos por su nombre de etiqueta:
```css
h1 { color: red; }
p { font-size: 16px; }
div { background-color: lightblue; }
```

### 3. Selectores de clase
Selecciona elementos con un atributo `class` específico:
```html
<div class="caja">Contenido</div>
<p class="destacado">Texto importante</p>
```

```css
.caja {
  border: 1px solid black;
  padding: 10px;
}

.destacado {
  font-weight: bold;
  color: red;
}
```

### 4. Selectores de ID
Selecciona un elemento con un `id` específico (único por página):
```html
<header id="cabecera">Título principal</header>
```

```css
#cabecera {
  background-color: navy;
  color: white;
  padding: 20px;
}
```

## Selectores de atributo

### Atributo presente
```css
input[type] {
  border: 1px solid gray;
}
```

### Atributo con valor específico
```css
input[type="text"] {
  width: 200px;
}

input[type="email"] {
  background-color: lightyellow;
}
```

### Atributo que contiene una palabra
```css
[class~="boton"] {
  cursor: pointer;
}
```

### Atributo que comienza con
```css
[href^="https"] {
  color: green;
}
```

### Atributo que termina con
```css
[src$=".png"] {
  border: 2px solid blue;
}
```

### Atributo que contiene
```css
[class*="card"] {
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}
```

## Selectores combinadores

### 1. Selector descendente (espacio)
Selecciona elementos dentro de otros:
```css
article p {
  line-height: 1.6;
}

nav ul li {
  display: inline-block;
}
```

### 2. Selector de hijo directo (>)
Selecciona hijos directos únicamente:
```css
ul > li {
  margin-bottom: 5px;
}

div > p {
  font-weight: bold;
}
```

### 3. Selector adyacente (+)
Selecciona el elemento inmediatamente siguiente:
```css
h1 + p {
  margin-top: 0;
  color: gray;
}
```

### 4. Selector general de hermanos (~)
Selecciona todos los hermanos siguientes:
```css
h1 ~ p {
  font-style: italic;
}
```

## Agrupación de selectores

Puedes aplicar las mismas reglas a múltiples selectores separándolos con comas:

```css
h1, h2, h3 {
  font-family: 'Arial', sans-serif;
  color: #333;
}

.caja, .tarjeta, .panel {
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 20px;
}
```

## Especificidad de selectores

La especificidad determina qué regla CSS se aplica cuando hay conflictos:

### Cálculo de especificidad:
- **ID**: 100 puntos cada uno
- **Clase/Atributo/Pseudo-clase**: 10 puntos cada uno
- **Elemento/Pseudo-elemento**: 1 punto cada uno

**Ejemplos:**
```css
/* Especificidad: 001 (1 elemento) */
p { color: blue; }

/* Especificidad: 010 (1 clase) */
.destacado { color: red; }

/* Especificidad: 100 (1 ID) */
#titulo { color: green; }

/* Especificidad: 111 (1 ID + 1 clase + 1 elemento) */
#titulo.destacado span { color: purple; }
```

## Herencia en CSS

Muchas propiedades CSS se heredan de elementos padre a hijos:

### Propiedades que se heredan:
- `color`
- `font-family`, `font-size`, `font-weight`
- `text-align`, `line-height`
- `list-style`

### Propiedades que NO se heredan:
- `margin`, `padding`
- `border`
- `width`, `height`
- `background`

### Forzar herencia:
```css
h1 {
  color: inherit; /* Hereda del padre */
  border: inherit; /* Fuerza herencia */
}
```

## Valores y unidades

### Unidades absolutas:
- `px` - píxeles
- `pt` - puntos (72 por pulgada)
- `pc` - picas (12 puntos)
- `in` - pulgadas
- `cm` - centímetros
- `mm` - milímetros

### Unidades relativas:
- `em` - relativo al tamaño de fuente del elemento
- `rem` - relativo al tamaño de fuente del root
- `%` - porcentaje relativo al elemento padre
- `vw` - 1% del ancho del viewport
- `vh` - 1% del alto del viewport
- `vmin` - menor entre vw o vh
- `vmax` - mayor entre vw o vh

## Ejemplo práctico completo

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <title>Selectores CSS</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <header id="cabecera">
    <h1 class="titulo-principal">Mi Sitio Web</h1>
    <nav>
      <ul>
        <li><a href="#inicio">Inicio</a></li>
        <li><a href="#servicios">Servicios</a></li>
        <li><a href="#contacto">Contacto</a></li>
      </ul>
    </nav>
  </header>

  <main>
    <section id="inicio">
      <h2>Bienvenido</h2>
      <p class="introduccion">Esta es una página de ejemplo para demostrar selectores CSS.</p>
      <div class="caja">
        <p>Contenido dentro de una caja.</p>
      </div>
    </section>

    <section id="servicios">
      <h3>Nuestros Servicios</h3>
      <article class="servicio destacado">
        <h4>Servicio Premium</h4>
        <p>El mejor servicio disponible.</p>
      </article>
      <article class="servicio">
        <h4>Servicio Básico</h4>
        <p>Servicio confiable y económico.</p>
      </article>
    </section>
  </main>
</body>
</html>
```

```css
/* Selector universal */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* Selectores de tipo */
body {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  color: #333;
}

h1, h2, h3, h4 {
  margin-bottom: 10px;
}

/* Selector de ID */
#cabecera {
  background-color: #2c3e50;
  color: white;
  padding: 20px;
  text-align: center;
}

/* Selector de clase */
.titulo-principal {
  font-size: 2.5em;
  margin-bottom: 0;
}

/* Selectores descendentes */
#cabecera nav ul {
  list-style: none;
}

#cabecera nav ul li {
  display: inline-block;
  margin: 0 15px;
}

#cabecera nav a {
  color: white;
  text-decoration: none;
  padding: 5px 10px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

#cabecera nav a:hover {
  background-color: #34495e;
}

/* Selector de hijo directo */
main > section {
  margin-bottom: 40px;
  padding: 20px;
}

/* Selectores de atributo */
a[href^="#"] {
  font-weight: bold;
}

/* Selector de clase múltiple */
.servicio.destacado {
  border: 2px solid #e74c3c;
  background-color: #ffeaea;
  padding: 15px;
  border-radius: 8px;
}

.servicio {
  margin-bottom: 20px;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 8px;
}

/* Selector adyacente */
h3 + article {
  margin-top: 10px;
}

/* Selectores de atributo avanzados */
input[type="text"] {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

input[type="email"]:focus {
  border-color: #3498db;
  outline: none;
}
```

## Resumen

En este módulo aprendiste:

- ✅ La sintaxis completa de CSS
- ✅ Todos los selectores básicos: universal, tipo, clase, ID
- ✅ Selectores de atributo con diferentes operadores
- ✅ Combinadores: descendente, hijo directo, adyacente, hermanos
- ✅ Agrupación de selectores
- ✅ Especificidad y cascada
- ✅ Herencia de propiedades
- ✅ Unidades absolutas y relativas

Los selectores son la base para aplicar estilos específicos a elementos HTML. Practica combinándolos para crear reglas CSS eficientes y mantenibles.