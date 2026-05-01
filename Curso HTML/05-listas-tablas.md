# Módulo 05 - Listas y tablas

En este módulo veremos listas y tablas con todos sus elementos importantes.

## Listas

### Lista desordenada

```html
<ul>
  <li>Manzanas</li>
  <li>Peras</li>
  <li>Plátanos</li>
</ul>
```

### Lista ordenada

```html
<ol>
  <li>Preparar ingredientes</li>
  <li>Mezclar</li>
  <li>Hornear</li>
</ol>
```

### Lista de definición

```html
<dl>
  <dt>HTML</dt>
  <dd>Lenguaje de marcado que estructura páginas web.</dd>
  <dt>CSS</dt>
  <dd>Lenguaje para dar estilo a páginas web.</dd>
</dl>
```

### Listas anidadas

```html
<ul>
  <li>Frutas
    <ul>
      <li>Manzanas</li>
      <li>Naranjas</li>
    </ul>
  </li>
  <li>Verduras</li>
</ul>
```

## Tablas

```html
<table>
  <caption>Horario de clases</caption>
  <thead>
    <tr>
      <th scope="col">Día</th>
      <th scope="col">Hora</th>
      <th scope="col">Asignatura</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Lunes</td>
      <td>09:00</td>
      <td>HTML</td>
    </tr>
    <tr>
      <td>Martes</td>
      <td>11:00</td>
      <td>JavaScript</td>
    </tr>
  </tbody>
</table>
```

### Componentes de una tabla

- `<caption>` describe la tabla.
- `<thead>` contiene los encabezados.
- `<tbody>` contiene los datos.
- `<tfoot>` puede usarse para totales.
- `<th scope="col">` o `scope="row"` mejora accesibilidad.

## Columnas con `<colgroup>`

```html
<table>
  <colgroup>
    <col span="1">
    <col span="2">
  </colgroup>
  <thead>...</thead>
  <tbody>...</tbody>
</table>
```

## Ejemplo real: comparación de características

```html
<table>
  <caption>Comparativa de planes</caption>
  <thead>
    <tr>
      <th scope="col">Plan</th>
      <th scope="col">Precio</th>
      <th scope="col">Espacio</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row">Básico</th>
      <td>$5</td>
      <td>10 GB</td>
    </tr>
    <tr>
      <th scope="row">Pro</th>
      <td>$12</td>
      <td>50 GB</td>
    </tr>
  </tbody>
</table>
```

## Resumen

Usa listas para elementos relacionados y tablas para datos estructurados. Asegúrate de que cada tabla tenga encabezados claros y una descripción con `<caption>`.
