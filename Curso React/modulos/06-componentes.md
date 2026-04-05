# Módulo 06 - Componentes

> **Las piezas de LEGO de React - bloques reutilizables**

---

## ¿QÉ es un Componente?

Un **componente** = función que devuelve HTML:

```jsx
function Saludo() {              // ← Función
  return <h1>¡Hola!</h1>;       // ← Devuelve HTML (JSX)
}
```

**Analogía:**
- 🧑 Persona = función
- 🗣️ Lo que dice = lo que devuelve
- 👕 Reusable = puedes llamarla 100 veces

---

## ¿PARA QÉ sirven?

### 1. No Repetir Código
```jsx
// ❌ SIN componentes (copiar-pegar)
function HomePage() {
  return (
    <div>
      <div className="tarjeta">
        <h2>Producto 1</h2>
        <button>Comprar</button>
      </div>
      <div className="tarjeta">
        <h2>Producto 2</h2>
        <button>Comprar</button>
      </div>
      <div className="tarjeta">
        <h2>Producto 3</h2>
        <button>Comprar</button>
      </div>
    </div>
  );
}

// ✅ CON componentes (reutilizable)
function TarjetaProducto({ nombre }) {
  return (
    <div className="tarjeta">
      <h2>{nombre}</h2>
      <button>Comprar</button>
    </div>
  );
}

function HomePage() {
  return (
    <div>
      <TarjetaProducto nombre="Producto 1" />
      <TarjetaProducto nombre="Producto 2" />
      <TarjetaProducto nombre="Producto 3" />
    </div>
  );
}
```

### 2. Mantener Organizado
```
App/
├── Header/
│   └── Header.jsx      ← Componente
├── Navbar/
│   └── Navbar.jsx      ← Componente
├── TarjetaProducto/
│   └── TarjetaProducto.jsx  ← Componente
└── Footer/
    └── Footer.jsx      ← Componente
```

### 3. Reutilizar en Proyectos
```jsx
// Tu librería de componentes
import Button from './componentes/Button';
import Card from './componentes/Card';
// Usas en 10 proyectos diferentes
```

### 4. Cambiar de un Lugar
```jsx
// Cambias TarjetaProducto UNA VEZ
// 50 páginas se actualizan automáticamente
```

---

## ¿CÓMO funcionan?

### Partes de un Componente

```jsx
function Tarjeta(props) {          // ← Nombre en PascalCase
  const { título } = props;        // ← Recibe información
  
  return (                         // ← Siempre retorna HTML
    <div className="tarjeta">
      <h2>{título}</h2>
      <p>Descripción</p>
    </div>
  );
}
```

### Cómo se Usa

```jsx
// Esto...
<Tarjeta título="Mi Tarjeta" />

// Es como llamar...
Tarjeta({ título: "Mi Tarjeta" })
```

### Tipos de Componentes

| Tipo | Qué es | Cuándo usar |
|------|--------|-------------|
| **Funcional** | Función normal | Siempre (2024+) |
| **Clase** | extends React.Component | Raro (antiguo) |
| **Controlado** | Tiene estado | Interactivos |
| **No controlado** | Sin estado | Solo muestra |

### Flujo Mental: Componente en Acción

```
1. DEFINIR
   function Botón({ texto }) {
     return <button>{texto}</button>;
   }

2. USAR
   <Botón texto="Haz clic" />

3. RENDERIZADO
   <button>Haz clic</button>

4. PANTALLA
   ┌────────────────┐
   │  [Haz clic]    │
   └────────────────┘
```

## Ejercicio
Crea estos componentes:
- `Header`
- `Footer`
- `TarjetaUsuario`
