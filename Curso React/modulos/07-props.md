# Módulo 07 - Props

> **Pasar instrucciones de padre a hijo - la comunicación en React**

---

## ¿QÉ son las Props?

**Props** = parámetros que un componente padre envía a un hijo:

```jsx
// Componente hijo
function Saludo(props) {
  return <h1>Hola {props.nombre}</h1>;
}

// Componente padre
function App() {
  return <Saludo nombre="Jose" />;
  //              ↑ Props (instrucción)
}
```

**Analogía:**
- 🎬 Componente = actor
- 🎞️ Props = guión (instrucciones)
- 🎪 Padre = director (da instrucciones)
- 🎭 Hijo = actor (sigue instrucciones)

---

## ¿PARA QÉ sirven?

### 1. Personalizar Componentes
```jsx
// Define UNA VEZ
function TarjetaProducto({ nombre, precio }) {
  return (
    <div className="tarjeta">
      <h3>{nombre}</h3>
      <p>${precio}</p>
    </div>
  );
}

// Usa MUCHAS VECES con instrucciones diferentes
<TarjetaProducto nombre="Laptop" precio={999} />
<TarjetaProducto nombre="Mouse" precio={25} />
<TarjetaProducto nombre="Teclado" precio={75} />
```

### 2. Pasar Datos Complejos
```jsx
function Perfil({ usuario }) {
  return (
    <div>
      <h2>{usuario.nombre}</h2>
      <p>Email: {usuario.email}</p>
      <p>Ciudad: {usuario.ciudad}</p>
    </div>
  );
}

const datos = {
  nombre: "Juan",
  email: "juan@ejemplo.com",
  ciudad: "Santiago"
};

<Perfil usuario={datos} />
```

### 3. Pasar Funciones (Callbacks)
```jsx
function Botón({ onClick, texto }) {
  return <button onClick={onClick}>{texto}</button>;
}

function App() {
  const manejarClic = () => alert('Clic!');
  
  return (
    <Botón 
      texto="Haz clic" 
      onClick={manejarClic}  // ← Función como prop
    />
  );
}
```

### 4. Pasar Listas (children)
```jsx
function Tarjeta({ children, título }) {
  return (
    <div className="tarjeta">
      <h2>{título}</h2>
      {children}  // ← Lo que va adentro
    </div>
  );
}

<Tarjeta título="Mi Tarjeta">
  <p>Contenido 1</p>
  <p>Contenido 2</p>
</Tarjeta>
```

---

## ¿CÓMO funcionan?

### 2 Formas de Acceder a Props

```jsx
// Forma 1: Sin desestructuración
function Saludo(props) {
  return <h1>Hola {props.nombre}</h1>;
}

// Forma 2: Desestructuración (moderna)
function Saludo({ nombre }) {
  return <h1>Hola {nombre}</h1>;
}

// Ambas hacen lo MISMO
```

### Flujo: Padre → Hijo

```
Padre (App)
├─ State: nombre = "Juan"
├─ Envía: <Hijo nombre={nombre} />
│
Hijo (Saludo)
├─ Recibe: nombre = "Juan"
├─ Renderiza: <h1>Hola Juan</h1>
│
Pantalla
└─ ¡Hola Juan!
```

### Props vs State

| Aspecto | Props | State |
|--------|-------|-------|
| **Quién controla** | Padre | Componente mismo |
| **¿Se puede cambiar?** | ❌ No (read-only) | ✅ Sí (setState) |
| **Cuándo cambia** | Padre lo decide | Componente lo decide |
| **Analogía** | Parámetro función | Variable local |

### Validación de Props (Bonus)

```jsx
import PropTypes from 'prop-types';

function Producto({ nombre, precio }) {
  return (
    <div>
      <h3>{nombre}</h3>
      <p>${precio}</p>
    </div>
  );
}

Producto.propTypes = {
  nombre: PropTypes.string.isRequired,
  precio: PropTypes.number.isRequired
};
```

### Props Comunes

```jsx
// Props normales
<Componente 
  nombre="Juan"
  edad={25}
  activo={true}
  items={[1,2,3]}
  callback={() => console.log('click')}
/>

// Children (especial)
<Componente>
  Contenido adentro
</Componente>
```

## Ejercicio
Crea un componente `Producto` que reciba:
- nombre
- precio
- stock
