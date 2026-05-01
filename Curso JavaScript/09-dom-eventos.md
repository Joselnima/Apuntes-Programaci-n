# Módulo 09 - DOM y Eventos Avanzados

El **DOM (Document Object Model)** es la representación en árbol que hace el navegador del documento HTML. Mediante JavaScript, tenemos acceso a una inmensa API para leer, crear, modificar y eliminar cualquier parte de la página, así como reaccionar a las interacciones del usuario.

A continuación, la guía exhaustiva de manipulación del DOM y Eventos.

## 1. Seleccionar Elementos del DOM

Existen formas antiguas y modernas para capturar los nodos HTML.

```javascript
// --- MÉTODOS TRADICIONALES ---
const porId = document.getElementById("miBoton"); // Devuelve 1 elemento
const porClase = document.getElementsByClassName("caja"); // Devuelve una HTMLCollection (Viva)
const porEtiqueta = document.getElementsByTagName("div"); // Devuelve una HTMLCollection

// --- MÉTODOS MODERNOS (Recomendados) ---
// querySelector devuelve el PRIMER elemento que coincida con el selector CSS
const primerBoton = document.querySelector(".btn-primario"); 
const header = document.querySelector("header > h1");

// querySelectorAll devuelve TODOS los elementos en un NodeList (Estático)
// Nota: Un NodeList permite usar .forEach(), una HTMLCollection antigua no.
const todosLosBotones = document.querySelectorAll("button");
todosLosBotones.forEach(btn => console.log(btn));
```

## 2. Crear, Clonar y Eliminar Elementos

```javascript
// --- CREAR ---
const nuevoParrafo = document.createElement("p"); // Crea el nodo 
nuevoParrafo.textContent = "Este es un párrafo creado desde JS";

// --- CLONAR ---
// clonar(true) clona el elemento y TODOS sus hijos. clonar(false) solo la etiqueta vacía.
const copia = nuevoParrafo.cloneNode(true); 

// --- ELIMINAR ---
// Método moderno
copia.remove(); 

// Método clásico (desde el padre)
// document.body.removeChild(nuevoParrafo);
```

## 3. Insertar Elementos en la Página

Una vez creado un elemento, está en la memoria RAM, no en la pantalla. Hay que insertarlo.

```javascript
const contenedor = document.querySelector("#contenedor");
const div = document.createElement("div");

// --- MÉTODOS MODERNOS ---
contenedor.append(div, "Texto suelto"); // Añade al final por dentro (Acepta múltiples nodos y texto)
contenedor.prepend(div); // Añade al principio por dentro
contenedor.before(div);  // Lo inserta JUSTO ANTES del contenedor (como hermano)
contenedor.after(div);   // Lo inserta JUSTO DESPUÉS del contenedor (como hermano)
contenedor.replaceWith(div); // Reemplaza todo el contenedor con el nuevo div

// --- MÉTODOS CLÁSICOS ---
contenedor.appendChild(div); // Añade al final (Solo acepta 1 nodo, no texto)
// contenedor.insertBefore(nodoNuevo, nodoReferencia);

// --- INSERTAR HTML DIRECTAMENTE (Poderoso pero cuidado con XSS) ---
// insertAdjacentHTML(posición, stringHTML)
// Posiciones: "beforebegin", "afterbegin", "beforeend", "afterend"
contenedor.insertAdjacentHTML("beforeend", "<strong>Hola Mundo</strong>");
```

## 4. Modificar Contenido

```javascript
const titulo = document.querySelector("h1");

// .textContent -> Lee/Escribe solo texto puro (Ignora etiquetas HTML, es más seguro y rápido)
titulo.textContent = "Nuevo Título"; 

// .innerText -> Parecido a textContent, pero respeta el CSS (No lee texto oculto con display:none)
console.log(titulo.innerText);

// .innerHTML -> Lee/Escribe renderizando las etiquetas HTML.
titulo.innerHTML = 'Nuevo <span style="color:red">Título</span>';

// .outerHTML -> Reemplaza/Lee TODO el elemento, incluyendo la propia etiqueta h1
console.log(titulo.outerHTML); // "<h1>...</h1>"
```

## 5. Modificar Atributos y Clases

La forma correcta de dar estilo no es inyectando CSS directo, sino cambiando clases.

```javascript
const caja = document.querySelector(".caja");

// --- CLASES (classList) ---
caja.classList.add("activa", "visible"); // Añadir clases
caja.classList.remove("visible");        // Eliminar clases
caja.classList.toggle("activa");         // Alternar: si la tiene la quita, si no, la pone.
caja.classList.replace("caja", "tarjeta"); // Cambia una por otra
console.log(caja.classList.contains("activa")); // true o false

// --- ATRIBUTOS GLOBALES ---
caja.setAttribute("id", "nuevaId");
caja.setAttribute("disabled", "");
console.log(caja.getAttribute("id")); // "nuevaId"
caja.removeAttribute("disabled");
console.log(caja.hasAttribute("disabled")); // false

// --- DATA-ATTRIBUTES (Atributos de datos personalizados) ---
// HTML: <div id="jugador" data-vidas="3" data-nombre-poder="fuego"></div>
const jugador = document.getElementById("jugador");
console.log(jugador.dataset.vidas);       // "3"
console.log(jugador.dataset.nombrePoder); // "fuego" (JS convierte kebab-case a camelCase)
jugador.dataset.vidas = "2";              // Escribir en el HTML
```

## 6. Estilos en Línea y Computados

```javascript
const btn = document.querySelector("button");

// Escribir estilos en línea (Prioridad alta en CSS)
btn.style.backgroundColor = "blue"; // JS usa camelCase para propiedades CSS (background-color -> backgroundColor)
btn.style.marginTop = "20px";

// Leer el estilo REAL aplicado (Combinando CSS, inline, navegador...)
const estilosReales = getComputedStyle(btn);
console.log(estilosReales.padding); // ej. "10px 15px"
```

## 7. Moverse por el DOM (Traversing)

Cómo saltar de un nodo a sus familiares sin usar `querySelector`.

```javascript
const elemento = document.querySelector(".hijo-medio");

// Hacia arriba (Padres)
console.log(elemento.parentElement); // Padre directo
// closest() busca el ancestro más cercano que coincida con el selector (muy útil)
console.log(elemento.closest(".contenedor-principal")); 

// Hacia los lados (Hermanos)
console.log(elemento.nextElementSibling);     // Siguiente hermano HTML
console.log(elemento.previousElementSibling); // Hermano anterior HTML

// Hacia abajo (Hijos)
const padre = document.querySelector("ul");
console.log(padre.children); // HTMLCollection de los <li>
console.log(padre.firstElementChild); // Primer <li>
console.log(padre.lastElementChild);  // Último <li>
```

## 8. Escuchar Eventos Avanzados (`addEventListener`)

La forma moderna y segura de escuchar acciones del usuario.

```javascript
const btnAccion = document.querySelector("#accion");

// addEventListener(evento, callback, opciones)
const saludar = (e) => {
  console.log("Hiciste clic!");
  
  // El Objeto Evento (e) contiene TODO sobre la acción
  console.log(e.type); // "click"
  console.log(e.target); // El nodo HTML exacto que originó el click
  console.log(e.clientX, e.clientY); // Coordenadas del mouse en pantalla
};

// Asignar evento
btnAccion.addEventListener("click", saludar);

// Quitar evento (Ojo: requiere que la función tenga nombre, no sirve con flechas anónimas)
btnAccion.removeEventListener("click", saludar);

// Opciones Avanzadas en eventos
btnAccion.addEventListener("click", () => console.log("Solo una vez"), { 
  once: true, // Se auto-destruye después del primer clic
  capture: false, 
  passive: true // Optimiza el rendimiento en eventos de scroll
});
});
```

## 9. Catálogo de Eventos Principales

JavaScript puede reaccionar a cientos de eventos diferentes. Aquí están los más utilizados y cómo manejarlos.

### Eventos de Ratón (Mouse)
```javascript
const zona = document.querySelector(".zona-interactiva");

// Clics
zona.addEventListener("click", () => console.log("Clic izquierdo"));
zona.addEventListener("dblclick", () => console.log("Doble clic"));
zona.addEventListener("contextmenu", (e) => {
  e.preventDefault(); // Evita que salga el menú por defecto del navegador
  console.log("Clic derecho");
});

// Movimiento
zona.addEventListener("mouseenter", () => console.log("El cursor ENTRO a la zona"));
zona.addEventListener("mouseleave", () => console.log("El cursor SALIÓ de la zona"));
zona.addEventListener("mousemove", (e) => {
  // e.offsetX y e.offsetY te dan las coordenadas relativas dentro del elemento
  console.log(`Moviendo en X: ${e.offsetX}, Y: ${e.offsetY}`);
});

// Botones del ratón físicos y rueda
zona.addEventListener("mousedown", () => console.log("Botón presionado"));
zona.addEventListener("mouseup", () => console.log("Botón soltado"));
zona.addEventListener("wheel", (e) => console.log(`Rueda del ratón girada: ${e.deltaY}`)); // deltaY positivo baja, negativo sube

// Diferencia entre mouseover y mouseenter:
// mouseover "burbujea" y se dispara constantemente al pasar sobre hijos internos.
// mouseenter NO burbujea, es más seguro y eficiente para menús desplegables.
```

### Eventos de Teclado
```javascript
const inputTexto = document.querySelector("input[type='text']");

// keydown: Cuando la tecla se presiona (se dispara repetidamente si se mantiene pulsada)
// keyup: Cuando se suelta la tecla
inputTexto.addEventListener("keydown", (e) => {
  console.log("Tecla presionada:", e.key); // Ej: "Enter", "A", "Escape"
  
  // Detectar combinaciones
  if (e.ctrlKey && e.key === "c") {
    console.log("El usuario intentó copiar!");
  }
});
```

### Eventos de Ventana y Documento
```javascript
// DOMContentLoaded: Se dispara cuando el HTML ha sido completamente cargado
// y analizado, sin esperar a CSS, imágenes ni iframes. ¡Muy útil!
document.addEventListener("DOMContentLoaded", () => {
  console.log("El DOM está listo para ser manipulado");
});

// load: Espera a que todo (imágenes, CSS externo) esté descargado.
window.addEventListener("load", () => {
  console.log("Toda la página, incluyendo imágenes, ha cargado.");
});

// resize: Cuando el usuario cambia el tamaño de la ventana
window.addEventListener("resize", () => {
  console.log(`Nuevo tamaño: ${window.innerWidth}x${window.innerHeight}`);
});

// scroll: Cuando el usuario hace scroll (Ojo: dispara cientos de veces, usar 'passive')
window.addEventListener("scroll", () => {
  console.log(`Haciendo scroll. Posición en Y: ${window.scrollY}`);
}, { passive: true });

// beforeunload: Avisar al usuario antes de que cierre la pestaña accidentalmente
window.addEventListener("beforeunload", (e) => {
  e.preventDefault();
  e.returnValue = ""; // Muestra el cartel nativo: "¿Seguro que quieres salir?"
});

// hashchange: Fundamental para Single Page Applications (SPAs)
// Se dispara cuando la URL cambia la parte del hashtag (ej: misitio.com/#contacto)
window.addEventListener("hashchange", () => {
  console.log("Navegando a la sección:", window.location.hash);
});
```

### Eventos de Formularios y Portapapeles (Clipboard)
> **Nota:** Formularios en profundidad en el **Módulo 13**.

```javascript
const input = document.querySelector("input");

// Foco
input.addEventListener("focus", () => console.log("El input está activo"));
input.addEventListener("blur", () => console.log("El input perdió el foco"));

// Foco con Burbujeo: focusin y focusout sí burbujean, focus y blur NO.
document.querySelector("form").addEventListener("focusin", () => console.log("Un hijo recibió foco"));

// Portapapeles (Copiar, Cortar y Pegar)
input.addEventListener("copy", (e) => {
  console.log("El usuario copió el texto!");
  // e.preventDefault(); // Puedes prohibir que copien si lo deseas
});
input.addEventListener("paste", (e) => {
  const textoPegado = e.clipboardData.getData("text");
  console.log("Pegaron esto:", textoPegado);
});
```

### Eventos de Arrastrar y Soltar (Drag & Drop)
Para hacer tableros tipo Trello o subida de archivos. El elemento a arrastrar debe tener `draggable="true"`.

```javascript
const cajaArrastrable = document.querySelector(".draggable");
const zonaCaida = document.querySelector(".dropzone");

// Eventos del elemento que viaja
cajaArrastrable.addEventListener("dragstart", () => console.log("Empezó el arrastre"));
cajaArrastrable.addEventListener("dragend", () => console.log("Terminó el arrastre"));

// Eventos de la zona de destino
zonaCaida.addEventListener("dragover", (e) => {
  e.preventDefault(); // OBLIGATORIO para permitir que algo caiga aquí
  console.log("Elemento sobrevolando la zona");
});
zonaCaida.addEventListener("drop", (e) => {
  e.preventDefault();
  console.log("Elemento soltado con éxito!");
});
```

### Eventos Multimedia (Audio / Video) y Animaciones
```javascript
const video = document.querySelector("video");

video.addEventListener("play", () => console.log("El video comenzó"));
video.addEventListener("pause", () => console.log("El video fue pausado"));
video.addEventListener("ended", () => console.log("El video terminó"));
video.addEventListener("timeupdate", () => console.log("Segundo actual:", video.currentTime));

// Animaciones CSS
const elementoAnimado = document.querySelector(".animacion");
elementoAnimado.addEventListener("animationstart", () => console.log("Inicia animación CSS"));
elementoAnimado.addEventListener("animationend", () => console.log("Termina animación CSS"));
elementoAnimado.addEventListener("transitionend", () => console.log("Termina transición (ej: hover)"));
```

### Eventos Táctiles (Móviles)
```javascript
const zonaTouch = document.querySelector(".touch");

zonaTouch.addEventListener("touchstart", (e) => {
  console.log("Dedos apoyados:", e.touches.length); // Útil para detectar gestos de zoom (2 dedos)
}, { passive: true });

zonaTouch.addEventListener("touchmove", () => console.log("Arrastrando dedo por la pantalla"));
zonaTouch.addEventListener("touchend", () => console.log("Dedo levantado"));
```

## 10. Propagación de Eventos (Bubbling y Delegation)

### Event Bubbling (Burbujeo)
Cuando haces clic en un elemento interno (ej. un `<span>` dentro de un `<button>`), el evento se dispara para el `<span>`, luego "burbujea" hacia arriba disparando el evento del `<button>`, luego el `<div>`, hasta llegar a `document`.

```javascript
const hijo = document.querySelector(".hijo");
hijo.addEventListener("click", (e) => {
  e.stopPropagation(); // Detiene el burbujeo inmediatamente.
  console.log("Clic solo procesado en el hijo");
});
```

### Event Delegation (Delegación)
El patrón de oro en JS: en vez de asignar 1000 listeners a 1000 botones, le pones 1 listener al padre y averiguas quién fue clickeado con `e.target`. Ahorra muchísima memoria RAM.

```javascript
const tablaDinamica = document.querySelector("#tabla");

tablaDinamica.addEventListener("click", (e) => {
  // e.target es donde el usuario hizo el click físico
  if (e.target.matches(".btn-eliminar")) {
    const fila = e.target.closest("tr"); // Busca la fila padre
    fila.remove();
    console.log("Fila eliminada de forma eficiente!");
  }
});
```

## 11. APIs del DOM Esenciales

```javascript
const caja = document.querySelector(".caja");

// 1. getBoundingClientRect() - Obtiene tamaño y posición exacta en la pantalla
const medidas = caja.getBoundingClientRect();
console.log(`Ancho: ${medidas.width}, Alto: ${medidas.height}, Top: ${medidas.top}`);

// 2. IntersectionObserver (Se usa para Lazy Loading o animaciones al scrollear)
const observador = new IntersectionObserver((entradas) => {
  entradas.forEach(entrada => {
    if (entrada.isIntersecting) {
      console.log("¡La caja ya es visible en la pantalla!");
      entrada.target.classList.add("animar");
    }
  });
});
observador.observe(caja);
```
