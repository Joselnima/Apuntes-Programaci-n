# Módulo 06 - Formularios

Este módulo explica todos los elementos de un formulario en HTML, desde básicos hasta avanzados, con ejemplos detallados y atributos importantes.

## La etiqueta `<form>`

```html
<form action="/submit" method="post">
  <!-- Campos del formulario -->
</form>
```

Atributos importantes:

- `action`: URL que recibirá los datos (puede ser relativa o absoluta).
- `method`: `get` (datos en URL) o `post` (datos en cuerpo de la petición).
- `enctype`: tipo de codificación (`application/x-www-form-urlencoded` por defecto, `multipart/form-data` para archivos).
- `autocomplete`: `on` o `off` para activar/desactivar autocompletado.
- `novalidate`: evita la validación del navegador.
- `target`: donde abrir el resultado (`_self`, `_blank`, etc.).

## Etiquetas básicas de formulario

```html
<label for="nombre">Nombre:</label>
<input type="text" id="nombre" name="nombre">
```

### Asociación `label` / `input`

- Usa `for` en el `label` apuntando al `id` del `input`.
- Mejora accesibilidad y aumenta el área de clic.
- También puedes anidar el input dentro del label.

## Tipos de input comunes

```html
<input type="text" name="nombre" required placeholder="Tu nombre">
<input type="email" name="correo" required placeholder="nombre@ejemplo.com">
<input type="password" name="password" minlength="8">
<input type="tel" name="telefono" pattern="[0-9]{9}">
<input type="url" name="web" placeholder="https://ejemplo.com">
<input type="number" name="cantidad" min="1" max="10" step="1">
<input type="date" name="fecha">
<input type="time" name="hora">
<input type="datetime-local" name="fecha-hora">
<input type="month" name="mes">
<input type="week" name="semana">
<input type="color" name="color">
<input type="range" name="volumen" min="0" max="100" value="50">
<input type="search" name="busqueda" placeholder="Buscar...">
<input type="file" name="archivo" accept="image/*">
<input type="hidden" name="token" value="abc123">
```

### Detalles de tipos de input

- `text`: Texto general.
- `email`: Valida formato de email.
- `password`: Oculta el texto.
- `tel`: Para números de teléfono (sin validación automática).
- `url`: Valida formato de URL.
- `number`: Solo números, con `min`, `max`, `step`.
- `date`: Selector de fecha.
- `time`: Selector de hora.
- `datetime-local`: Fecha y hora local.
- `month`: Selector de mes y año.
- `week`: Selector de semana.
- `color`: Selector de color.
- `range`: Slider numérico.
- `search`: Campo de búsqueda (puede tener estilo especial).
- `file`: Para subir archivos, `accept` filtra tipos.
- `hidden`: Campo oculto, útil para datos internos.

## Checkbox y radio

```html
<label><input type="checkbox" name="acepto" required value="si"> Acepto los términos</label>
<label><input type="checkbox" name="notificaciones" checked> Recibir notificaciones</label>

<label><input type="radio" name="plan" value="basic" checked> Básico</label>
<label><input type="radio" name="plan" value="pro"> Pro</label>
<label><input type="radio" name="plan" value="premium"> Premium</label>
```

- Checkbox: Múltiples selecciones, `checked` para preseleccionar.
- Radio: Una selección por grupo (mismo `name`), `checked` para preseleccionar.

## `textarea` y `select`

```html
<label for="mensaje">Mensaje:</label>
<textarea id="mensaje" name="mensaje" rows="4" cols="50" maxlength="500" placeholder="Escribe tu mensaje..."></textarea>

<label for="pais">País:</label>
<select id="pais" name="pais" required>
  <option value="">Selecciona un país</option>
  <option value="es">España</option>
  <option value="mx">México</option>
  <option value="ar" selected>Argentina</option>
</select>
```

### `select` con `optgroup`

```html
<select name="color">
  <optgroup label="Colores fríos">
    <option value="azul">Azul</option>
    <option value="verde">Verde</option>
    <option value="violeta">Violeta</option>
  </optgroup>
  <optgroup label="Colores cálidos">
    <option value="rojo">Rojo</option>
    <option value="amarillo">Amarillo</option>
    <option value="naranja">Naranja</option>
  </optgroup>
</select>
```

- `optgroup`: Agrupa opciones para mejor organización.
- `selected`: Preselecciona una opción.

## `datalist` para sugerencias

```html
<label for="ciudad">Ciudad:</label>
<input list="ciudades" id="ciudad" name="ciudad">
<datalist id="ciudades">
  <option value="Madrid">
  <option value="Barcelona">
  <option value="Valencia">
  <option value="Sevilla">
</datalist>
```

- `datalist`: Proporciona sugerencias mientras el usuario escribe.

## Botones de formulario

```html
<button type="submit">Enviar</button>
<button type="reset">Limpiar formulario</button>
<button type="button" onclick="alert('Hola')">Botón normal</button>
<input type="submit" value="Enviar">
<input type="reset" value="Limpiar">
<input type="button" value="Acción" onclick="miFuncion()">
```

- `submit`: Envía el formulario.
- `reset`: Restaura valores iniciales.
- `button`: Botón genérico, necesita JavaScript para acción.

## Atributos de validación y control

### Atributos básicos

- `required`: Campo obligatorio.
- `placeholder`: Texto de ayuda que desaparece al escribir.
- `readonly`: Solo lectura, no editable.
- `disabled`: Deshabilitado, no se envía.
- `maxlength`: Máximo caracteres.
- `minlength`: Mínimo caracteres.

### Atributos numéricos y de rango

- `min`, `max`: Valores mínimo y máximo.
- `step`: Incremento para números/range.

```html
<input type="number" name="edad" min="18" max="120" step="1">
<input type="range" name="satisfaccion" min="1" max="5" step="1" value="3">
```

### Atributos de patrón

- `pattern`: Expresión regular para validación.

```html
<input type="text" name="codigo" pattern="[A-Z]{3}[0-9]{3}" title="Tres letras mayúsculas seguidas de tres números">
<input type="tel" name="telefono" pattern="^\+?[0-9\s\-\(\)]+$" title="Número de teléfono válido">
```

### Atributos de autocompletado

- `autocomplete`: Sugerencias del navegador.

```html
<input type="email" name="correo" autocomplete="email">
<input type="text" name="direccion" autocomplete="street-address">
<input type="text" name="ciudad" autocomplete="address-level2">
```

Valores comunes: `name`, `email`, `tel`, `address-line1`, `postal-code`, etc.

## `fieldset` y `legend` para agrupar

```html
<fieldset>
  <legend>Información personal</legend>
  <label for="nombre">Nombre:</label>
  <input id="nombre" name="nombre" type="text" required>

  <label for="apellido">Apellido:</label>
  <input id="apellido" name="apellido" type="text" required>
</fieldset>

<fieldset>
  <legend>Preferencias</legend>
  <label><input type="checkbox" name="newsletter"> Recibir newsletter</label>
</fieldset>
```

- `fieldset`: Agrupa campos relacionados.
- `legend`: Título del grupo.

## `output` para resultados calculados

```html
<form oninput="total.value = (parseInt(cantidad.value) || 0) * (parseFloat(precio.value) || 0)">
  <label for="cantidad">Cantidad:</label>
  <input id="cantidad" type="number" value="1" min="1">

  <label for="precio">Precio unitario:</label>
  <input id="precio" type="number" value="10.50" step="0.01">

  <label for="total">Total:</label>
  <output name="total" for="cantidad precio">10.50</output>
</form>
```

- `output`: Muestra resultados de cálculos, actualizado automáticamente.

## `progress` y `meter` para indicadores

```html
<label>Progreso de carga:</label>
<progress value="65" max="100">65%</progress>

<label>Nivel de batería:</label>
<meter value="0.8" min="0" max="1" low="0.2" high="0.8" optimum="1">80%</meter>
```

- `progress`: Barra de progreso (valor actual hacia máximo).
- `meter`: Indicador de medida (con rangos bajo/alto/óptimo).

## Validación avanzada con JavaScript (opcional)

Aunque HTML5 valida automáticamente, puedes usar JavaScript para validaciones personalizadas:

```html
<input type="text" name="usuario" pattern="[a-zA-Z0-9]{4,}" title="Al menos 4 caracteres alfanuméricos" required>
```

## Ejemplo real: formulario de registro completo

```html
<form action="/registro" method="post" enctype="multipart/form-data">
  <fieldset>
    <legend>Datos personales</legend>
    
    <label for="nombre">Nombre:</label>
    <input id="nombre" name="nombre" type="text" required placeholder="Tu nombre completo">
    
    <label for="email">Correo electrónico:</label>
    <input id="email" name="email" type="email" required autocomplete="email">
    
    <label for="telefono">Teléfono:</label>
    <input id="telefono" name="telefono" type="tel" pattern="[0-9]{9}" title="9 dígitos">
    
    <label for="fecha-nacimiento">Fecha de nacimiento:</label>
    <input id="fecha-nacimiento" name="fecha-nacimiento" type="date" required>
    
    <label for="pais">País:</label>
    <select id="pais" name="pais" required>
      <option value="">Selecciona tu país</option>
      <option value="es">España</option>
      <option value="mx">México</option>
      <option value="ar">Argentina</option>
    </select>
  </fieldset>

  <fieldset>
    <legend>Cuenta</legend>
    
    <label for="usuario">Nombre de usuario:</label>
    <input id="usuario" name="usuario" type="text" required minlength="4" maxlength="20" pattern="[a-zA-Z0-9_]+" title="Solo letras, números y guiones bajos">
    
    <label for="password">Contraseña:</label>
    <input id="password" name="password" type="password" required minlength="8" placeholder="Mínimo 8 caracteres">
    
    <label for="avatar">Foto de perfil:</label>
    <input id="avatar" name="avatar" type="file" accept="image/*">
  </fieldset>

  <fieldset>
    <legend>Preferencias</legend>
    
    <label for="color-favorito">Color favorito:</label>
    <input id="color-favorito" name="color-favorito" type="color" value="#ff0000">
    
    <label for="volumen-notificaciones">Volumen de notificaciones:</label>
    <input id="volumen-notificaciones" name="volumen-notificaciones" type="range" min="0" max="100" value="50">
    
    <label><input type="checkbox" name="terminos" required> Acepto los términos y condiciones</label>
    <label><input type="checkbox" name="newsletter" checked> Suscribirme al newsletter</label>
    
    <label>Plan preferido:</label><br>
    <label><input type="radio" name="plan" value="basico" checked> Básico</label>
    <label><input type="radio" name="plan" value="premium"> Premium</label>
  </fieldset>

  <button type="submit">Crear cuenta</button>
  <button type="reset">Limpiar formulario</button>
</form>
```

## Resumen

Los formularios HTML ofrecen una gran variedad de inputs y atributos para crear experiencias de usuario ricas y validaciones robustas. Combina tipos de input apropiados con atributos de validación para formularios accesibles y funcionales. Recuerda que la validación HTML5 es básica; para validaciones complejas, considera JavaScript.

