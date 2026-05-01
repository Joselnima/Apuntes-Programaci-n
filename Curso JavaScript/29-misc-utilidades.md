# Módulo 29 - Utilidades Misceláneas (Date, Intl, URL, Timers)

JavaScript provee múltiples APIs globales en el navegador para lidiar con el tiempo, la internacionalización y la manipulación de URLs de manera profesional.

## 1. Fechas con el objeto `Date`

El objeto `Date` permite manipular fechas y tiempos. Trabaja con milisegundos desde el 1 de enero de 1970 (Epoch de UNIX).

```javascript
// Obtener la fecha y hora actual
const ahora = new Date();

// Crear fecha específica (Año, Mes (0-indexado, enero es 0!), Día, Horas, Min, Sec)
const miCumple = new Date(1995, 11, 25); // 25 de diciembre de 1995

// Leer propiedades
console.log(ahora.getFullYear()); // 2026
console.log(ahora.getMonth());    // Mes del 0 al 11
console.log(ahora.getDate());     // Día del mes (1 - 31)
console.log(ahora.getDay());      // Día de la semana (0 domingo, 6 sábado)

// Modificar
ahora.setFullYear(2030);

// Cálculo de tiempo
const inicio = Date.now(); // Milisegundos actuales
// ... ejecutar algo pesado ...
const fin = Date.now();
console.log(`Tardó ${fin - inicio} ms`);
```

## 2. Formato de Localización (`Intl`)

El objeto `Intl` (Internacionalización) es fantástico para mostrar fechas o monedas en el formato que espera el usuario según su país, sin usar librerías de terceros (como moment.js).

```javascript
const monto = 1250000.5;

// Formatear como Moneda (ej. Euros en España)
const formatoEuro = new Intl.NumberFormat('es-ES', { 
  style: 'currency', 
  currency: 'EUR' 
}).format(monto);
console.log(formatoEuro); // "1.250.000,50 €"

// Formatear Fechas (ej. Español de México)
const fecha = new Date();
const formatoFecha = new Intl.DateTimeFormat('es-MX', { 
  dateStyle: 'full' 
}).format(fecha);
console.log(formatoFecha); // "martes, 1 de mayo de 2026"
```

## 3. URLs y Query Strings (`URL`, `URLSearchParams`)

Construir URLs a mano concatenando strings (`"?user="+id`) es propenso a errores (por caracteres especiales). La API `URL` lo soluciona.

```javascript
const miUrl = new URL("https://api.sitio.com/buscar");

// Añadir parámetros de forma segura
miUrl.searchParams.append("q", "camisa azul");
miUrl.searchParams.append("limit", 10);

console.log(miUrl.href); 
// "https://api.sitio.com/buscar?q=camisa+azul&limit=10"

// Leer parámetros de la URL actual del navegador
const params = new URLSearchParams(window.location.search);
if (params.has("id")) {
  console.log("El ID solicitado es:", params.get("id"));
}
```

### Funciones globales de Codificación de URIs
Cuando envías datos por la URL, ciertos caracteres (como espacios, `&`, `=`, `?`) pueden romper la estructura de la petición. Por ello, JavaScript ofrece métodos para "codificar" esos caracteres especiales.

- **`encodeURI()`**: Codifica toda una URL. Ignora caracteres especiales que son legales en las URLs (como `:`, `/`, `?`, `&`).
- **`encodeURIComponent()`**: Codifica de forma agresiva. Convierte TODOS los caracteres especiales. Es ideal para codificar un *parámetro* específico antes de inyectarlo en la URL.
- **`decodeURI()` / `decodeURIComponent()`**: Hacen el proceso inverso.

```javascript
// 1. encodeURI (Para URLs completas)
const urlMala = "https://misitio.com/mi archivo con espacios.pdf";
const urlArreglada = encodeURI(urlMala); 
console.log(urlArreglada); // "https://misitio.com/mi%20archivo%20con%20espacios.pdf"

// 2. encodeURIComponent (Para variables o valores de búsqueda)
const parametroPeligroso = "a+b=c & d";
// Si pusieras esto directo en la URL se rompería porque '&' y '=' significan otra cosa.
const paramCodificado = encodeURIComponent(parametroPeligroso);
console.log(paramCodificado); // "a%2Bb%3Dc%20%26%20d"

// Construcción manual infalible:
const busquedaSegura = `https://api.com/buscar?q=${paramCodificado}`;

// 3. Decodificar (Cuando recibes los datos del servidor)
console.log(decodeURIComponent("a%2Bb%3Dc%20%26%20d")); // "a+b=c & d"
```

## 4. Timers (`setTimeout` y `setInterval`)

Métodos para retrasar o repetir la ejecución de código.

```javascript
// setTimeout: Ejecuta una sola vez después de X milisegundos
const idTimeout = setTimeout(() => {
  console.log("Han pasado 2 segundos!");
}, 2000);

// clearTimeout cancela el timer antes de que ocurra
clearTimeout(idTimeout);

// setInterval: Ejecuta repetidamente cada X milisegundos
let contador = 0;
const idIntervalo = setInterval(() => {
  contador++;
  console.log(`Tick ${contador}`);
  
  if (contador === 5) {
    clearInterval(idIntervalo); // Detener el intervalo
  }
}, 1000);
```
