# ¿Qué es la Web?

La **Web** (o **World Wide Web**) es un sistema de información que permite acceder a documentos, aplicaciones, imágenes, videos y servicios a través de **Internet** usando un navegador como Chrome, Edge o Firefox.

> **Importante:**
> **Internet** y **Web** no son exactamente lo mismo.
>
> - **Internet** = la infraestructura de red mundial.
> - **Web** = uno de los servicios que funciona sobre Internet.

### Ejemplo simple

Piensa en esto:

- **Internet** = carreteras y autopistas.
- **Web** = autos, buses, delivery y personas moviéndose por esas carreteras.

La Web usa esa infraestructura para que tú puedas abrir una página como:

```txt
https://www.google.com
```

y ver contenido desde cualquier parte del mundo.

---

# Historia breve de la Web

La Web fue propuesta por **Tim Berners-Lee** en 1989 mientras trabajaba en CERN.

Su objetivo era resolver un problema muy concreto:

> **¿Cómo compartir información entre investigadores de forma simple y conectada?**

Para resolverlo, se crearon tres piezas fundamentales:

1. **HTML** → para estructurar documentos.
2. **HTTP** → para comunicar cliente y servidor.
3. **URL** → para identificar recursos.

Estas tres ideas siguen siendo el corazón de la Web moderna.

---

# Cómo funciona la Web paso a paso

Cuando tú escribes una dirección en el navegador, por ejemplo:

```txt
https://www.youtube.com
```

pasan muchas cosas por detrás.

## Flujo general

```txt
Usuario → Navegador → DNS → Servidor → Respuesta → Navegador renderiza
```

Ahora lo explicamos bien.

---

## 1) El usuario solicita una URL

Tú escribes una URL en el navegador:

```txt
https://midominio.com/productos
```

Una URL tiene partes:

```txt
https://midominio.com/productos
│      │             │
│      │             └── Ruta o recurso
│      └──────────────── Dominio
└────────────────────── Protocolo
```

### Partes de una URL

- **Protocolo**: cómo se va a comunicar (`http`, `https`)
- **Dominio**: nombre del sitio (`midominio.com`)
- **Ruta**: qué recurso quieres (`/productos`)
- **Parámetros** (opcional): filtros o datos (`?id=10`)

Ejemplo:

```txt
https://miapp.com/productos?id=10&categoria=ropa
```

---

## 2) El navegador pregunta al DNS

Las computadoras no “entienden” nombres como:

```txt
google.com
```

Ellas trabajan mejor con direcciones IP como:

```txt
142.250.190.14
```

Entonces el navegador consulta a un **DNS (Domain Name System)**.

### ¿Qué hace el DNS?

Traduce esto:

```txt
google.com
```

a esto:

```txt
142.250.190.14
```

### Analogía

El DNS es como la **agenda de contactos de Internet**.

- Tú buscas: “Google”
- El DNS te dice: “su número es esta IP”

---

## 3) El navegador se conecta al servidor

Una vez que tiene la IP, el navegador intenta conectarse al **servidor** donde está alojada la web.

Ese servidor es una computadora (o conjunto de computadoras) que está esperando solicitudes.

### El servidor puede hacer varias cosas

- Enviar una página HTML
- Enviar imágenes
- Enviar archivos CSS
- Enviar JavaScript
- Consultar una base de datos
- Ejecutar lógica de negocio
- Devolver datos JSON

---

## 4) Se envía una petición HTTP

El navegador envía una **petición HTTP** al servidor.

Ejemplo conceptual:

```http
GET /productos HTTP/1.1
Host: miapp.com
```

### ¿Qué significa eso?

Le está diciendo:

> “Hola servidor, quiero el recurso `/productos` del dominio `miapp.com`”

---

## 5) El servidor procesa la solicitud

El servidor recibe la petición y decide qué hacer.

### Puede ser algo simple

Por ejemplo:

- El usuario pide `/nosotros`
- El servidor devuelve un HTML estático

### O algo más complejo

Por ejemplo:

- El usuario pide `/productos?id=10`
- El servidor consulta una base de datos
- Procesa reglas de negocio
- Devuelve resultados

Ejemplo de backend:

```txt
Petición → API → Servicio → Base de datos → Respuesta
```

---

## 6) El servidor responde

El servidor devuelve una **respuesta HTTP**.

Ejemplo conceptual:

```http
HTTP/1.1 200 OK
Content-Type: text/html
```

Y además envía el contenido.

### Posibles respuestas HTTP

- **200 OK** → todo salió bien
- **201 Created** → recurso creado
- **400 Bad Request** → solicitud incorrecta
- **401 Unauthorized** → no autenticado
- **403 Forbidden** → no autorizado
- **404 Not Found** → recurso no existe
- **500 Internal Server Error** → error del servidor

---

## 7) El navegador renderiza el contenido

Cuando recibe el contenido, el navegador lo interpreta y lo muestra visualmente.

Si el servidor devuelve HTML, el navegador:

1. Lee el HTML
2. Descarga CSS
3. Descarga JavaScript
4. Construye la interfaz
5. Muestra la página

---

# Protocolos principales

## HTTP

**HTTP (HyperText Transfer Protocol)** es el protocolo base de la Web.

Sirve para intercambiar información entre:

- **Cliente** (navegador)
- **Servidor**

## HTTPS

**HTTPS** es HTTP pero **seguro**.

Usa cifrado para proteger la información.

### ¿Por qué es importante?

Porque evita que otros puedan leer fácilmente:

- contraseñas
- tokens
- formularios
- información sensible

---

# Frontend y Backend

La mayoría de aplicaciones web se entienden mejor separándolas en dos mundos.

## Frontend

Es la parte que **ve y usa el usuario**.

Incluye:

- botones
- formularios
- tablas
- menús
- inputs
- modales
- dashboards

### Tecnologías típicas

- HTML
- CSS
- JavaScript
- TypeScript
- Angular
- React
- Vue

## Backend

Es la parte que **procesa la lógica y maneja datos**.

Incluye:

- autenticación
- reglas de negocio
- consultas a base de datos
- generación de reportes
- APIs

### Tecnologías típicas

- Python / FastAPI / Django
- Node.js / Express / NestJS
- Java / Spring
- C# / .NET
- PHP / Laravel

---

# Cómo se construye una página web

Una web moderna normalmente usa tres tecnologías base.

## 1) HTML

Define la **estructura**.

Ejemplo:

```html
<h1>Hola mundo</h1>
<p>Esta es una página web</p>
```

## 2) CSS

Define la **apariencia**.

Ejemplo:

```css
h1 {
  color: blue;
}
```

## 3) JavaScript

Define el **comportamiento**.

Ejemplo:

```javascript
document.querySelector('button').addEventListener('click', () => {
  alert('Hola');
});
```

# Resumen final

## La Web

La Web es un sistema que funciona sobre Internet y permite acceder a recursos mediante navegadores.

Su funcionamiento básico es:

```txt
Usuario → Navegador → DNS → Servidor → Respuesta → Renderizado
```

Está basada en conceptos como:

- URL
- HTTP / HTTPS
- HTML
- CSS
- JavaScript
- cliente / servidor

---