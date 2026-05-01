# Módulo 09 - Programación asíncrona

La programación asíncrona permite ejecutar acciones que tardan en completarse, como cargar datos desde una API o esperar una respuesta.

## Sincronía vs asíncronía

- Código sincronizado se ejecuta paso a paso.
- Código asíncrono puede esperar tareas sin bloquear el resto.

## `setTimeout` y `setInterval`

```javascript
setTimeout(() => {
  console.log('Esto se muestra después de 2 segundos');
}, 2000);

const intervalo = setInterval(() => {
  console.log('Cada segundo');
}, 1000);

setTimeout(() => {
  clearInterval(intervalo);
}, 5000);
```

## ¿Cómo funciona el Event Loop bajo el capó?

JavaScript es "Single Thread" (tiene un solo hilo de ejecución). El **Event Loop** es el secreto que le permite no bloquearse. Funciona con estos componentes:

1. **Call Stack (Pila de llamadas)**: Donde se ejecuta tu código síncrono.
2. **Web APIs**: Cuando llamas a `setTimeout` o `fetch`, el navegador (C++) toma el control en segundo plano y deja la pila libre.
3. **Colas de Tareas**: Cuando la Web API termina, manda el callback a una cola:
   - **Microtasks** (Pila VIP, muy rápida): Para `Promises` y `async/await`.
   - **Macrotasks** (Pila normal): Para `setTimeout` y `setInterval`.
4. El Event Loop constantemente verifica: si la Pila de llamadas está vacía, mete primero todas las Microtasks pendientes, y luego una Macrotask.

## Promesas

Una promesa es un objeto que representa una operación que puede terminar en éxito o error.

```javascript
const promesa = new Promise((resolve, reject) => {
  const exito = true;

  if (exito) {
    resolve('La operación fue exitosa');
  } else {
    reject('Ocurrió un error');
  }
});

promesa
  .then(resultado => console.log(resultado))
  .catch(error => console.error(error));
```

## `fetch` para obtener datos

```javascript
fetch('https://jsonplaceholder.typicode.com/posts/1')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error(error));
```

## `async` y `await`

```javascript
async function obtenerDatos() {
  try {
    const respuesta = await fetch('https://jsonplaceholder.typicode.com/posts/1');
    const datos = await respuesta.json();
    console.log(datos);
  } catch (error) {
    console.error('Error:', error);
  }
}

obtenerDatos();
```

## Métodos Avanzados de Promesas (Ejecución en paralelo)

Con `async/await` a menudo caemos en el error de esperar una petición, luego la otra, luego otra... de manera secuencial, perdiendo tiempo. Si no dependen entre sí, **ejecútalas en paralelo** usando los métodos de la clase estática `Promise`:

```javascript
const getUsuarios = fetch('https://api.com/usuarios');
const getProductos = fetch('https://api.com/productos');

// 1. Promise.all()
// Útil cuando NECESITAS todas. Si UNA falla, TODO falla.
const [usuarios, productos] = await Promise.all([getUsuarios, getProductos]);

// 2. Promise.allSettled() (ES11)
// Espera a que TODAS terminen. Devuelve el estado individual de cada una aunque alguna haya fallado.
const resultados = await Promise.allSettled([getUsuarios, getProductos]);

// 3. Promise.race()
// Devuelve la PRIMERA que termine (ya sea con éxito o error). Ideal para crear un límite de tiempo (Timeout).
const primera = await Promise.race([getUsuarios, getProductos]);

// 4. Promise.any() (ES12)
// Devuelve la primera que tenga ÉXITO. Ignora las fallidas.
const primerExito = await Promise.any([getUsuarios, getProductos]);
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Programación asíncrona</title>
</head>
<body>
  <h1>Programación asíncrona</h1>
  <button id="cargar">Cargar usuario</button>
  <pre id="resultado"></pre>

  <script>
    const boton = document.getElementById('cargar');
    const resultado = document.getElementById('resultado');

    boton.addEventListener('click', async () => {
      resultado.textContent = 'Cargando...';

      try {
        const respuesta = await fetch('https://jsonplaceholder.typicode.com/users/1');
        const usuario = await respuesta.json();
        resultado.textContent = JSON.stringify(usuario, null, 2);
      } catch (error) {
        resultado.textContent = 'Error al cargar los datos.';
      }
    });
  </script>
</body>
</html>
```

## Ejercicio

1. Usa `setTimeout` para mostrar un mensaje tras 3 segundos.
2. Crea una promesa que resuelva después de 2 segundos.
3. Usa `async/await` para leer datos de una API y mostrar el resultado.

## Resumen

La asíncronía es clave en JavaScript moderno. `Promises`, `fetch` y `async/await` facilitan trabajar con tareas que no terminan de inmediato.
