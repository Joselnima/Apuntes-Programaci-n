# Módulo 10 - Formularios avanzados

Este módulo completa el trabajo de formularios con elementos avanzados y validación detallada.

## `fieldset` y `legend`

```html
<fieldset>
  <legend>Información personal</legend>
  <label for="nombre">Nombre</label>
  <input id="nombre" name="nombre" type="text">
</fieldset>
```

## `datalist`

```html
<label for="ciudad">Ciudad</label>
<input list="ciudades" id="ciudad" name="ciudad">
<datalist id="ciudades">
  <option value="Madrid">
  <option value="Barcelona">
  <option value="Valencia">
</datalist>
```

## `output`

```html
<form oninput="resultado.value=parseInt(cantidad.value)*precio.value">
  <label for="cantidad">Cantidad</label>
  <input id="cantidad" type="number" value="1">

  <label for="precio">Precio</label>
  <input id="precio" type="number" value="10">

  <output name="resultado" for="cantidad precio">10</output>
</form>
```

## `progress` y `meter`

```html
<progress value="40" max="100">40%</progress>
<meter value="0.6">60%</meter>
```

## Tipos de input avanzados

```html
<input type="url" name="sitio">
<input type="tel" name="telefono">
<input type="range" name="volumen" min="0" max="100">
<input type="color" name="color">
<input type="datetime-local" name="fecha-hora">
<input type="month" name="mes">
<input type="week" name="semana">
```

## `placeholder`, `autofocus` y `readonly`

```html
<input type="text" placeholder="Tu nombre" autofocus>
<input type="text" readonly value="Solo lectura">
```

## Validación con atributos HTML

- `required`
- `pattern`
- `min` / `max`
- `step`
- `maxlength`

```html
<input type="email" name="correo" required>
<input type="text" name="usuario" pattern="[A-Za-z0-9]{4,10}">
```

## `form` independiente

```html
<form id="mi-form">
  <input type="text" name="nombre" form="mi-form">
</form>
```

## Ejemplo real: formulario de registro avanzado

```html
<form action="/registro" method="post">
  <fieldset>
    <legend>Cuenta</legend>
    <label for="usuario">Usuario</label>
    <input id="usuario" name="usuario" type="text" required>

    <label for="correo">Correo</label>
    <input id="correo" name="correo" type="email" required>
  </fieldset>

  <fieldset>
    <legend>Preferencias</legend>
    <label for="color">Color favorito</label>
    <input id="color" name="color" type="color">

    <label for="volumen">Volumen</label>
    <input id="volumen" name="volumen" type="range" min="0" max="100">
  </fieldset>

  <button type="submit">Crear cuenta</button>
</form>
```

## Resumen

Los formularios avanzados aprovechan los elementos nativos de HTML para enriquecer la experiencia y validar datos sin JavaScript.
