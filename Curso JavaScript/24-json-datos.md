# Módulo 15 - Formato JSON y Manejo de Datos

JSON (JavaScript Object Notation) es el formato de intercambio de datos más utilizado en la web actual. Aunque está inspirado en la sintaxis de los objetos de JavaScript, es texto plano y puede ser leído por casi cualquier lenguaje de programación.

## ¿Qué es JSON?

Es una forma estructurada de representar datos (cadenas, números, booleanos, arreglos y objetos) en formato de texto.

```json
{
  "nombre": "Ana",
  "edad": 28,
  "habilidades": ["JavaScript", "HTML", "CSS"],
  "activo": true,
  "direccion": {
    "ciudad": "Madrid",
    "pais": "España"
  }
}
```

**Reglas importantes del JSON:**
1. Las claves (propiedades) deben ir siempre entre comillas dobles `""`.
2. No permite funciones ni `undefined`.
3. El último elemento de un objeto o arreglo no debe tener una coma `,` al final.

## El objeto global `JSON`

JavaScript proporciona un objeto global llamado `JSON` que tiene dos métodos fundamentales para trabajar con este formato.

### 1. `JSON.stringify()` (Convertir a JSON)

Toma un objeto o valor de JavaScript y lo convierte en una cadena de texto JSON. Esto es necesario cuando quieres enviar datos a un servidor o guardarlos en `localStorage`.

```javascript
const usuario = {
  nombre: "Carlos",
  edad: 30,
  roles: ["admin", "editor"]
};

// Convertir objeto JS a string JSON
const jsonString = JSON.stringify(usuario);

console.log(jsonString);
// Imprime: '{"nombre":"Carlos","edad":30,"roles":["admin","editor"]}'
```

Puedes pasar argumentos extra para formatear el texto (hacerlo bonito o "pretty-print"):
```javascript
// El tercer argumento (2) es el número de espacios de indentación
const jsonBonito = JSON.stringify(usuario, null, 2);
console.log(jsonBonito);
```

### 2. `JSON.parse()` (Convertir de JSON a JavaScript)

Toma una cadena de texto JSON y la convierte en un objeto (o el valor correspondiente) de JavaScript. Es lo que hacemos cuando recibimos datos de una API.

```javascript
const respuestaServidor = '{"producto": "Laptop", "precio": 1200, "stock": 50}';

// Convertir string JSON a objeto JS
const objetoJS = JSON.parse(respuestaServidor);

console.log(objetoJS.producto); // "Laptop"
console.log(objetoJS.precio);   // 1200
```

## Clonación profunda con JSON

Un truco muy común en JavaScript para hacer una copia profunda (deep clone) de un objeto que no tiene funciones ni referencias circulares, es usar ambos métodos juntos:

```javascript
const original = { a: 1, b: { c: 2 } };

// Al convertirlo a texto y volverlo a objeto, se crea una copia completamente nueva
const copia = JSON.parse(JSON.stringify(original));

copia.b.c = 99;
console.log(original.b.c); // 2 (El original no se modificó)
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <title>JSON</title>
</head>
<body>
  <h1>Manejo de JSON</h1>
  <script>
    const configuracion = {
      temaDato: "oscuro",
      notificaciones: true
    };

    // Guardar en LocalStorage (solo acepta strings)
    localStorage.setItem("config", JSON.stringify(configuracion));

    // Recuperar y usar
    const configGuardada = localStorage.getItem("config");
    if (configGuardada) {
      const config = JSON.parse(configGuardada);
      console.log("Tema actual:", config.temaDato);
    }
  </script>
</body>
</html>
```

## Resumen

JSON es vital. Memoriza `JSON.stringify()` para enviar o guardar, y `JSON.parse()` para leer y utilizar datos externos.
