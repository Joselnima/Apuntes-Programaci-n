# Módulo 17 - Elementos avanzados y atributos

En este módulo exploramos elementos HTML5 avanzados, atributos globales poderosos y características modernas que enriquecen las páginas web.

## Elementos avanzados de HTML5

### `<data>` para datos legibles por máquina

```html
<ul>
  <li><data value="398">Mini Cooper S</data></li>
  <li><data value="399">Mini Cooper</data></li>
  <li><data value="400">Mini Cooper D</data></li>
</ul>
```

Útil para datos que tienen significado tanto para humanos como para máquinas.

### `<time>` para fechas y horas

```html
<p>La reunión es el <time datetime="2026-04-30T14:00">30 de abril a las 2 PM</time>.</p>
<p>Publicado hace <time datetime="2026-04-25">5 días</time>.</p>
```

### `<output>` para resultados de cálculos

```html
<form oninput="resultado.value = cantidad.value * precio.value">
  <input type="number" id="cantidad" value="1">
  <input type="number" id="precio" value="10">
  = <output name="resultado" for="cantidad precio">10</output>
</form>
```

### `<progress>` y `<meter>` para indicadores

```html
<label>Progreso de carga:</label>
<progress value="65" max="100">65%</progress>

<label>Nivel de batería:</label>
<meter value="0.8" min="0" max="1" low="0.2" high="0.8" optimum="1">80%</meter>
```

### `<details>` y `<summary>` para contenido colapsable

```html
<details>
  <summary>Mostrar más información</summary>
  <p>Contenido adicional que se revela al hacer clic.</p>
</details>
```

### `<dialog>` para modales

```html
<dialog id="mi-modal">
  <p>Este es un modal nativo de HTML.</p>
  <button onclick="miModal.close()">Cerrar</button>
</dialog>

<button onclick="miModal.showModal()">Abrir modal</button>
```

### `<mark>` para resaltar texto

```html
<p>El <mark>HTML</mark> es fundamental para la web.</p>
```

### `<ruby>`, `<rt>`, `<rp>` para anotaciones

```html
<ruby>
  漢 <rt>kan</rt>
  字 <rt>ji</rt>
</ruby>
```

Útil para idiomas con anotaciones como el japonés.

### `<bdi>` para aislamiento direccional

```html
<p>Usuario: <bdi>اسم عربي</bdi> - Puntaje: 85</p>
```

Aísla el texto con dirección de escritura diferente.

### `<wbr>` para oportunidades de quiebre de palabra

```html
<p>Esta es una palabra muy larga que puede que necesite <wbr>quebrarse.</p>
```

## Atributos globales avanzados

### `contenteditable` - Contenido editable

```html
<div contenteditable="true">
  Puedes editar este texto directamente en el navegador.
</div>
```

### `spellcheck` - Revisión ortográfica

```html
<textarea spellcheck="true"></textarea>
<input type="text" spellcheck="false">
```

### `translate` - Indicador de traducción

```html
<p translate="yes">This text should be translated.</p>
<p translate="no">HTML</p>
```

### `draggable` - Elementos arrastrables

```html
<div draggable="true" id="elemento-arrastrable">
  Arrástrame
</div>

<div id="zona-soltar">
  Suelta aquí
</div>

<script>
  const arrastrable = document.getElementById('elemento-arrastrable');
  const zonaSoltar = document.getElementById('zona-soltar');

  arrastrable.addEventListener('dragstart', (e) => {
    e.dataTransfer.setData('text/plain', e.target.id);
  });

  zonaSoltar.addEventListener('dragover', (e) => {
    e.preventDefault();
  });

  zonaSoltar.addEventListener('drop', (e) => {
    e.preventDefault();
    const id = e.dataTransfer.getData('text/plain');
    const elemento = document.getElementById(id);
    e.target.appendChild(elemento);
  });
</script>
```

### `hidden` - Ocultar elementos

```html
<div hidden>
  Este contenido está oculto.
</div>
```

### `tabindex` - Orden de tabulación

```html
<button tabindex="1">Primero</button>
<button tabindex="3">Tercero</button>
<button tabindex="2">Segundo</button>
```

### `autocapitalize` - Capitalización automática

```html
<input type="text" autocapitalize="words">
<input type="text" autocapitalize="sentences">
<input type="text" autocapitalize="characters">
```

### `inputmode` - Modo de entrada

```html
<input type="text" inputmode="numeric">
<input type="text" inputmode="email">
<input type="text" inputmode="tel">
```

### `enterkeyhint` - Sugerencia para la tecla Enter

```html
<textarea enterkeyhint="send"></textarea>
<input enterkeyhint="search">
```

## Microdata y datos estructurados

Microdata permite añadir significado semántico a los datos:

```html
<div itemscope itemtype="https://schema.org/Person">
  <span itemprop="name">Juan Pérez</span>
  <span itemprop="jobTitle">Desarrollador Web</span>
  <div itemprop="address" itemscope itemtype="https://schema.org/PostalAddress">
    <span itemprop="streetAddress">Calle Falsa 123</span>
    <span itemprop="addressLocality">Madrid</span>
  </div>
</div>
```

## Elementos obsoletos pero aún válidos

### `<center>` (obsoleto, usa CSS)

```html
<center>Texto centrado</center>
```

### `<font>` (obsoleto, usa CSS)

```html
<font color="red" size="4">Texto con fuente</font>
```

### `<big>` y `<small>` (parcialmente obsoletos)

```html
<big>Texto grande</big>
<small>Texto pequeño</small>
```

## APIs relacionadas con HTML

### Web Storage (localStorage/sessionStorage)

```javascript
// Guardar datos
localStorage.setItem('usuario', 'Juan');

// Recuperar datos
const usuario = localStorage.getItem('usuario');

// Eliminar datos
localStorage.removeItem('usuario');
```

### Geolocation API

```javascript
if (navigator.geolocation) {
  navigator.geolocation.getCurrentPosition((position) => {
    console.log('Latitud:', position.coords.latitude);
    console.log('Longitud:', position.coords.longitude);
  });
}
```

### Intersection Observer (para lazy loading)

```javascript
const observer = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      entry.target.src = entry.target.dataset.src;
      observer.unobserve(entry.target);
    }
  });
});

document.querySelectorAll('img[data-src]').forEach(img => {
  observer.observe(img);
});
```

## Ejemplo completo: página con elementos avanzados

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <title>Elementos Avanzados HTML5</title>
</head>
<body>
  <header>
    <h1>Demostración de HTML5 Avanzado</h1>
  </header>

  <main>
    <section>
      <h2>Datos estructurados</h2>
      <div itemscope itemtype="https://schema.org/Event">
        <h3 itemprop="name">Conferencia Web</h3>
        <time itemprop="startDate" datetime="2026-05-15T10:00">15 de mayo de 2026, 10:00</time>
        <div itemprop="location" itemscope itemtype="https://schema.org/Place">
          <span itemprop="name">Centro de Convenciones</span>
        </div>
      </div>
    </section>

    <section>
      <h2>Contenido editable</h2>
      <div contenteditable="true" spellcheck="true">
        <p>Edita este texto. Tiene revisión ortográfica activada.</p>
      </div>
    </section>

    <section>
      <h2>Elementos interactivos</h2>
      <details>
        <summary>Mostrar detalles</summary>
        <p>Contenido oculto que se expande.</p>
        <progress value="75" max="100">75%</progress>
      </details>

      <dialog id="modal">
        <p>Este es un modal nativo.</p>
        <button onclick="modal.close()">Cerrar</button>
      </dialog>
      <button onclick="modal.showModal()">Abrir modal</button>
    </section>

    <section>
      <h2>Texto con anotaciones</h2>
      <p>El término <data value="42">respuesta</data> es <mark>importante</mark>.</p>
      <p>Palabra en japonés: <ruby>東京 <rt>Tokyo</rt></ruby></p>
    </section>

    <section>
      <h2>Drag and Drop</h2>
      <div draggable="true" id="arrastrable" style="padding: 10px; border: 1px solid black; margin: 10px; display: inline-block;">
        Arrástrame
      </div>
      <div id="zona-soltar" style="width: 200px; height: 100px; border: 2px dashed gray; padding: 10px; margin: 10px;">
        Suelta aquí
      </div>
    </section>
  </main>

  <script>
    // Drag and Drop
    const arrastrable = document.getElementById('arrastrable');
    const zonaSoltar = document.getElementById('zona-soltar');

    arrastrable.addEventListener('dragstart', (e) => {
      e.dataTransfer.setData('text/plain', e.target.id);
    });

    zonaSoltar.addEventListener('dragover', (e) => {
      e.preventDefault();
      zonaSoltar.style.backgroundColor = 'lightblue';
    });

    zonaSoltar.addEventListener('dragleave', () => {
      zonaSoltar.style.backgroundColor = '';
    });

    zonaSoltar.addEventListener('drop', (e) => {
      e.preventDefault();
      zonaSoltar.style.backgroundColor = '';
      const id = e.dataTransfer.getData('text/plain');
      const elemento = document.getElementById(id);
      zonaSoltar.appendChild(elemento);
    });

    // Web Storage demo
    const contenidoEditable = document.querySelector('[contenteditable]');
    const savedContent = localStorage.getItem('editableContent');
    if (savedContent) {
      contenidoEditable.innerHTML = savedContent;
    }

    contenidoEditable.addEventListener('input', () => {
      localStorage.setItem('editableContent', contenidoEditable.innerHTML);
    });
  </script>
</body>
</html>
```

## Resumen

HTML5 incluye muchos elementos y atributos avanzados que permiten crear experiencias web ricas sin depender únicamente de JavaScript o frameworks externos. Estos elementos mejoran la semántica, accesibilidad y funcionalidad nativa de las páginas web.</content>
<parameter name="filePath">e:\Lenguajes de programacion\Curso HTML\17-elementos-avanzados.md