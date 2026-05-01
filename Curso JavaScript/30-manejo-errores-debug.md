# Módulo 29 - Manejo de Errores y Debugging

Ningún código es perfecto. Saber manejar los errores en ejecución y utilizar las herramientas de debugging es lo que distingue a un profesional de un principiante.

## 1. Bloque `try...catch...finally`

El mecanismo principal para manejar errores que detendrían la aplicación. 

```javascript
try {
  // Código que podría fallar
  const datos = JSON.parse('{"formato_invalido: 123');
  console.log(datos);

} catch (error) {
  // Solo se ejecuta si algo en el 'try' falló
  console.error("Hubo un fallo al analizar el JSON:");
  
  // El objeto Error tiene propiedades útiles:
  console.error("Nombre:", error.name);       // "SyntaxError"
  console.error("Mensaje:", error.message);   // "Unexpected token f in JSON..."
  // console.error("Rastro:", error.stack);   // Rastro completo de llamadas

} finally {
  // Se ejecuta SIEMPRE, haya habido error o no.
  // Ideal para cerrar conexiones o limpiar interfaces (ej. quitar spinners de carga).
  console.log("Terminó el intento de análisis.");
}
```

## 2. Lanzar Errores Manualmente (`throw`)

A veces, el código técnicamente es correcto pero lógicamente es inválido (ej. transferir dinero negativo). Puedes forzar un error usando `throw`.

```javascript
function procesarPago(monto) {
  if (monto <= 0) {
    // throw detiene la función inmediatamente y salta al catch más cercano
    throw new Error("El monto debe ser mayor a cero."); 
  }
  console.log(`Procesando $${monto}...`);
}

try {
  procesarPago(-50);
} catch (error) {
  console.error(error.message); // "El monto debe ser mayor a cero."
}
```

## 3. Clases de Error Customizadas

Extender la clase base `Error` te permite crear errores personalizados para poder diferenciarlos en bloques `catch` grandes.

```javascript
class ValidationError extends Error {
  constructor(mensaje) {
    super(mensaje);
    this.name = "ValidationError";
  }
}

try {
  throw new ValidationError("El email no es válido");
} catch (err) {
  if (err instanceof ValidationError) {
    console.log("Error de usuario:", err.message);
  } else {
    console.log("Error desconocido de sistema:", err);
  }
}
```

## 4. Herramientas de Debugging en el Navegador

### El objeto `console` es más que `console.log()`

```javascript
const usuarios = [{ id: 1, nombre: "Ana" }, { id: 2, nombre: "Luis" }];

console.error("Texto en rojo (para errores serios)");
console.warn("Texto en amarillo (advertencias)");

// Imprime arrays u objetos como una tabla interactiva:
console.table(usuarios);

// Medir rendimiento:
console.time("Bucle");
for (let i = 0; i < 100000; i++) {} // trabajo...
console.timeEnd("Bucle"); // "Bucle: 2.5ms"

// Agrupar mensajes
console.group("Detalles del usuario");
console.log("Nombre: Ana");
console.log("Rol: Admin");
console.groupEnd();
```

### La palabra clave `debugger`

Si tienes la consola de desarrollo (DevTools, F12) abierta en tu navegador, colocar `debugger;` en tu código hará que la ejecución se pause exactamente en esa línea. Podrás inspeccionar el valor de todas las variables en ese momento preciso y avanzar línea por línea.

```javascript
function calcularImpuestos(precio) {
  let tasa = 0.21;
  debugger; // La ejecución se congelará aquí para inspección
  return precio + (precio * tasa);
}
```
