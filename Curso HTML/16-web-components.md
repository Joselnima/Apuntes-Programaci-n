# Módulo 16 - Web Components

En este módulo aprenderás sobre Web Components, una tecnología que permite crear elementos HTML reutilizables y encapsulados.

## ¿Qué son Web Components?

Web Components son un conjunto de tecnologías que permiten crear componentes reutilizables para la web. Incluyen:

- Custom Elements: Elementos HTML personalizados
- Shadow DOM: DOM encapsulado
- HTML Templates: Plantillas reutilizables
- ES Modules: Importación de módulos

## HTML Templates

`<template>` permite definir fragmentos de HTML que no se renderizan hasta que se usan.

```html
<template id="mi-template">
  <style>
    .caja { border: 1px solid black; padding: 10px; }
  </style>
  <div class="caja">
    <h3>Contenido del template</h3>
    <slot name="contenido">Contenido por defecto</slot>
  </div>
</template>
```

## Usando templates

```javascript
const template = document.getElementById('mi-template');
const clone = template.content.cloneNode(true);
document.body.appendChild(clone);
```

## Custom Elements

Crear elementos HTML personalizados:

```javascript
class MiElemento extends HTMLElement {
  constructor() {
    super();
    // Crear shadow DOM
    const shadow = this.attachShadow({ mode: 'open' });
    
    // Añadir contenido
    shadow.innerHTML = `
      <style>
        .mi-estilo { color: blue; }
      </style>
      <div class="mi-estilo">
        <slot>Hola desde Web Component!</slot>
      </div>
    `;
  }
}

// Registrar el elemento
customElements.define('mi-elemento', MiElemento);
```

## Usando Custom Elements

```html
<mi-elemento>Hola mundo</mi-elemento>
<mi-elemento>
  <span slot="contenido">Contenido personalizado</span>
</mi-elemento>
```

## Shadow DOM

El Shadow DOM encapsula el estilo y markup del componente:

```javascript
class TarjetaProducto extends HTMLElement {
  constructor() {
    super();
    const shadow = this.attachShadow({ mode: 'open' });
    
    shadow.innerHTML = `
      <style>
        .tarjeta {
          border: 1px solid #ccc;
          border-radius: 8px;
          padding: 16px;
          max-width: 300px;
        }
        .titulo { font-size: 1.2em; margin-bottom: 8px; }
        .precio { color: green; font-weight: bold; }
      </style>
      
      <div class="tarjeta">
        <div class="titulo"><slot name="titulo">Producto</slot></div>
        <div class="precio"><slot name="precio">$0.00</slot></div>
        <slot name="descripcion">Sin descripción</slot>
      </div>
    `;
  }
}

customElements.define('tarjeta-producto', TarjetaProducto);
```

## Slots nombrados

```html
<tarjeta-producto>
  <span slot="titulo">iPhone 15</span>
  <span slot="precio">$999</span>
  <p slot="descripcion">El último smartphone de Apple con cámara avanzada.</p>
</tarjeta-producto>
```

## Ciclo de vida de Custom Elements

```javascript
class MiComponente extends HTMLElement {
  constructor() {
    super();
    console.log('Constructor llamado');
  }
  
  connectedCallback() {
    console.log('Elemento conectado al DOM');
  }
  
  disconnectedCallback() {
    console.log('Elemento desconectado del DOM');
  }
  
  adoptedCallback() {
    console.log('Elemento movido a nuevo documento');
  }
  
  attributeChangedCallback(name, oldValue, newValue) {
    console.log(`Atributo ${name} cambió de ${oldValue} a ${newValue}`);
  }
  
  static get observedAttributes() {
    return ['titulo', 'activo'];
  }
}

customElements.define('mi-componente', MiComponente);
```

## Propiedades y métodos

```javascript
class ContadorClicks extends HTMLElement {
  constructor() {
    super();
    this.clicks = 0;
    const shadow = this.attachShadow({ mode: 'open' });
    
    shadow.innerHTML = `
      <button id="btn">Haz clic: 0</button>
    `;
    
    this.btn = shadow.getElementById('btn');
    this.btn.addEventListener('click', () => this.incrementar());
  }
  
  incrementar() {
    this.clicks++;
    this.btn.textContent = `Haz clic: ${this.clicks}`;
  }
  
  // Propiedad getter/setter
  get valor() {
    return this.clicks;
  }
  
  set valor(nuevoValor) {
    this.clicks = nuevoValor;
    this.btn.textContent = `Haz clic: ${this.clicks}`;
  }
}

customElements.define('contador-clicks', ContadorClicks);
```

## Usando el componente

```html
<contador-clicks id="contador"></contador-clicks>
<script>
  const contador = document.getElementById('contador');
  console.log(contador.valor); // 0
  contador.valor = 5; // Cambia el contador
</script>
```

## Estilos y temas

```javascript
class BotonPersonalizado extends HTMLElement {
  constructor() {
    super();
    const shadow = this.attachShadow({ mode: 'open' });
    
    shadow.innerHTML = `
      <style>
        :host {
          display: inline-block;
        }
        button {
          background: var(--color-primario, blue);
          color: white;
          border: none;
          padding: 8px 16px;
          border-radius: 4px;
          cursor: pointer;
        }
        button:hover {
          opacity: 0.8;
        }
      </style>
      <button><slot>Botón</slot></button>
    `;
  }
}

customElements.define('boton-personalizado', BotonPersonalizado);
```

## CSS desde fuera

```css
boton-personalizado {
  --color-primario: red;
}
```

## Ejemplo completo: componente de tarjeta

```html
<!DOCTYPE html>
<html>
<head>
  <title>Web Components Ejemplo</title>
</head>
<body>
  <template id="tarjeta-template">
    <style>
      .tarjeta {
        border: 1px solid #ddd;
        border-radius: 8px;
        padding: 16px;
        margin: 8px;
        box-shadow: 0 2px 4px rgba(0,0,0,0.1);
      }
      .titulo { font-size: 1.2em; font-weight: bold; }
      .contenido { margin: 8px 0; }
    </style>
    <div class="tarjeta">
      <div class="titulo"><slot name="titulo">Título</slot></div>
      <div class="contenido"><slot name="contenido">Contenido</slot></div>
      <slot name="acciones"></slot>
    </div>
  </template>

  <tarjeta-componente>
    <span slot="titulo">Mi Tarjeta</span>
    <p slot="contenido">Este es el contenido de la tarjeta.</p>
    <button slot="acciones">Acción</button>
  </tarjeta-componente>

  <script>
    class TarjetaComponente extends HTMLElement {
      constructor() {
        super();
        const template = document.getElementById('tarjeta-template');
        const shadow = this.attachShadow({ mode: 'open' });
        shadow.appendChild(template.content.cloneNode(true));
      }
    }
    
    customElements.define('tarjeta-componente', TarjetaComponente);
  </script>
</body>
</html>
```

## Compatibilidad y polyfills

Web Components funcionan en navegadores modernos. Para soporte antiguo, usa polyfills como webcomponents.js.

## Resumen

Web Components permiten crear componentes reutilizables y encapsulados. Son el futuro de los componentes web nativos, sin necesidad de frameworks externos.</content>
<parameter name="filePath">e:\Lenguajes de programacion\Curso HTML\16-web-components.md