# Módulo 13 - Formularios y Eventos en JS

Los formularios son la principal vía de interacción y captura de datos del usuario. JavaScript nos permite interceptar esos datos, validarlos y procesarlos sin tener que recargar la página.

## 1. El Evento `submit` y `.preventDefault()`

Por defecto, cuando envías un formulario HTML (haciendo clic en el botón de submit), el navegador recarga la página. En aplicaciones modernas (SPAs), cancelamos esto con `event.preventDefault()`.

```html
<form id="miFormulario">
  <input type="text" id="nombre" placeholder="Tu nombre" />
  <button type="submit">Enviar</button>
</form>

<script>
  const form = document.getElementById("miFormulario");

  form.addEventListener("submit", function(evento) {
    // 1. Detener la recarga de la página
    evento.preventDefault(); 
    
    // 2. Capturar el valor
    const inputNombre = document.getElementById("nombre");
    console.log("Formulario enviado. Nombre:", inputNombre.value);
  });
</script>
```

## 2. Eventos de Inputs (`input`, `change`, `focus`, `blur`)

Para reaccionar en tiempo real a lo que el usuario hace, tenemos múltiples eventos:

- `input`: Se dispara **cada vez que una tecla es presionada** o el valor cambia. Perfecto para barras de búsqueda en tiempo real.
- `change`: Se dispara **cuando el usuario sale del input** después de cambiar el valor (o al seleccionar un Checkbox / Select).
- `focus`: Cuando el input recibe el cursor.
- `blur`: Cuando el input pierde el cursor.

```javascript
const nombreInput = document.getElementById("nombre");

nombreInput.addEventListener("input", (e) => {
  console.log("Escribiendo:", e.target.value);
});

nombreInput.addEventListener("blur", () => {
  console.log("El usuario salió de la caja de texto");
});
```

## 3. Capturando datos fácilmente con `FormData`

Si tu formulario tiene docenas de campos, usar `document.getElementById` para cada uno es horrible. El objeto `FormData` lee automáticamente el atributo `name` de todos tus inputs y recopila sus valores en un solo paso.

```html
<form id="registro">
  <input type="text" name="usuario" value="Carlos" />
  <input type="password" name="clave" value="1234" />
  <button type="submit">Registrar</button>
</form>

<script>
  document.getElementById("registro").addEventListener("submit", (e) => {
    e.preventDefault();
    
    // Pasamos el formulario (e.target) a FormData
    const datos = new FormData(e.target);
    
    // Obtenemos los valores individuales por su 'name'
    console.log(datos.get("usuario")); // "Carlos"
    console.log(datos.get("clave"));   // "1234"
    
    // O convertimos todo a un objeto literal (ES10)
    const objeto = Object.fromEntries(datos.entries());
    console.log(objeto); // { usuario: "Carlos", clave: "1234" }
  });
</script>
```

## 4. Validación Nativa (HTML5) en JS

Puedes utilizar las propias reglas de validación de HTML5 (`required`, `minlength`, `type="email"`) pero controladas desde JavaScript usando el método `checkValidity()`.

```javascript
const formulario = document.getElementById("registro");

formulario.addEventListener("submit", (e) => {
  e.preventDefault();

  // Comprueba si el formulario cumple todas las reglas HTML5
  if (!formulario.checkValidity()) {
    console.log("El formulario tiene errores. Por favor, revísalo.");
    // Muestra los globos de error nativos del navegador
    formulario.reportValidity(); 
    return;
  }

  console.log("Formulario perfecto, enviando al servidor...");
});
```
