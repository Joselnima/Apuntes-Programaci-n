# Módulo 05 - JSX

> **HTML dentro de JavaScript - la magia de React**

---

## ¿QÉ es JSX?

**JSX** = escribir HTML dentro de JavaScript:

```jsx
const titulo = <h1>Hola React</h1>;
//             ↑ Esto parece HTML
//             Pero es JavaScript
```

**No es HTML puro:**
```jsx
// ❌ Esto NO es válido HTML
const titulo = <h1>Hola {nombre}</h1>;

// ❌ Esto NO es JavaScript normal
const html = <h1>...</h1>;
```

**Es una mágica sintaxis especial de React.**

**Analogía:**
- 🗣️ HTML = español
- 🤖 JavaScript = programación
- 🧙 JSX = hablar español DENTRO de programa

---

## ¿PARA QÉ sirve?

### 1. Escribir UI Más Fácil
```jsx
// ❌ Sin JSX (tedioso)
const elemento = React.createElement(
  'div',
  null,
  React.createElement('h1', null, 'Hola'),
  React.createElement('p', null, 'Mundo')
);

// ✅ Con JSX (natural)
const elemento = (
  <div>
    <h1>Hola</h1>
    <p>Mundo</p>
  </div>
);
```

### 2. Mezclar HTML con Variables
```jsx
const nombre = "Juan";
const edad = 25;

return (
  <div>
    <h1>¡Hola {nombre}!</h1>
    <p>Edad: {edad}</p>
  </div>
);
```

### 3. Lógica Dentro de Componentes
```jsx
const isAdmin = true;

return (
  <div>
    {isAdmin && <button>Eliminar usuario</button>}
  </div>
);
```

---

## ¿CÓMO funciona?

### Reglas Principales del JSX

| Regla | Sí ✅ | No ❌ |
|-------|------|------|
| **Un solo padre** | `<div><h1/><p/></div>` | `<h1/><p/>` (2 padres) |
| **JavaScript en {}** | `<h1>{nombre}</h1>` | `<h1>nombre</h1>` (texto) |
| **className** | `className="mi-clase"` | `class="mi-clase"` |
| **htmlFor** | `htmlFor="input1"` | `for="input1"` |
| **Propiedades camelCase** | `onClick={fn}` | `onclick={fn}` |

### Estructura Completa

```jsx
function Tarjeta({ título, color }) {
  return (
    <div className={`tarjeta ${color}`}>  // ← Un padre
      <h2>{título}</h2>                    // ← {} para JS
      <p className="descripción">          // ← className
        Lorem ipsum
      </p>
    </div>
  );
}
```

### Detrás de Escenas

```jsx
// Lo que escribes (JSX)
return <h1>Hola {nombre}</h1>;

// Lo que React convierte
return React.createElement(
  'h1',
  null,
  'Hola ',
  nombre
);

// Lo que renderiza en HTML
<h1>Hola Juan</h1>
```

### Expresiones Válidas en {}

```jsx
// ✅ Expresiones
{5 + 3}                    // 8
{nombre.toUpperCase()}     // JUAN
{isAdmin ? 'Admin' : 'User'}    // Admin
{items.map(item => item.nombre)}  // Lista

// ❌ Inválidas
{if (x) {...}}   // No IF
{let y = 5}      // No LET
{función()}      // Solo si retorna algo
```

## Ejercicio
Crea un JSX que muestre:
- tu nombre
- tu edad
- una lista de 3 tecnologías
