# Módulo 16 - Proyecto práctico

En este proyecto aplicarás lo aprendido para construir una pequeña aplicación web con JavaScript.

## Objetivo

Crear una lista de tareas dinámica donde el usuario pueda:

- añadir tareas
- marcar tareas como completas
- eliminar tareas
- ver el estado actual

## Estructura del proyecto

```
Curso JavaScript/
  index.html
  js/
    app.js
  css/
    styles.css
```

## HTML básico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Lista de tareas</title>
  <link rel="stylesheet" href="css/styles.css">
</head>
<body>
  <main>
    <h1>Lista de tareas</h1>
    <form id="formulario">
      <input type="text" id="tareaInput" placeholder="Escribe una tarea" required>
      <button type="submit">Añadir</button>
    </form>
    <ul id="listaTareas"></ul>
  </main>
  <script type="module" src="js/app.js"></script>
</body>
</html>
```

## CSS sencillo

```css
body {
  font-family: Arial, sans-serif;
  background: #f0f4f8;
  margin: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
}

main {
  background: white;
  padding: 20px;
  border-radius: 12px;
  width: 100%;
  max-width: 500px;
  box-shadow: 0 10px 20px rgba(0,0,0,0.1);
}

form {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 8px;
}

button {
  padding: 10px 16px;
  border: none;
  background: #007bff;
  color: white;
  border-radius: 8px;
  cursor: pointer;
}

li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #eee;
}

li.completa {
  text-decoration: line-through;
  color: #7a7a7a;
}

button.eliminar {
  background: #e55353;
}
```

## JavaScript

```javascript
const formulario = document.getElementById('formulario');
const tareaInput = document.getElementById('tareaInput');
const listaTareas = document.getElementById('listaTareas');

let tareas = [];

formulario.addEventListener('submit', event => {
  event.preventDefault();
  const texto = tareaInput.value.trim();

  if (texto === '') return;

  const tarea = {
    id: Date.now(),
    texto,
    completada: false
  };

  tareas.push(tarea);
  tareaInput.value = '';
  renderizarTareas();
});

function renderizarTareas() {
  listaTareas.innerHTML = '';

  tareas.forEach(tarea => {
    const li = document.createElement('li');
    li.textContent = tarea.texto;

    if (tarea.completada) {
      li.classList.add('completa');
    }

    li.addEventListener('click', () => {
      tarea.completada = !tarea.completada;
      renderizarTareas();
    });

    const botonEliminar = document.createElement('button');
    botonEliminar.textContent = 'Eliminar';
    botonEliminar.className = 'eliminar';
    botonEliminar.addEventListener('click', event => {
      event.stopPropagation();
      tareas = tareas.filter(item => item.id !== tarea.id);
      renderizarTareas();
    });

    li.appendChild(botonEliminar);
    listaTareas.appendChild(li);
  });
}

renderizarTareas();
```

## Extensiones sugeridas

- Guardar tareas en `localStorage`.
- Añadir prioridad a cada tarea.
- Mostrar contadores de tareas completadas y pendientes.
- Agregar edición de tareas.

## Ejercicio final

1. Implementa `localStorage` para conservar las tareas al recargar.
2. Añade un filtro para ver solo tareas completas o pendientes.
3. Mejora el estilo visual con colores y efectos.

## Resumen

Este proyecto integra variables, funciones, DOM, eventos y arrays. Es una excelente práctica para comenzar a crear aplicaciones reales con JavaScript.
