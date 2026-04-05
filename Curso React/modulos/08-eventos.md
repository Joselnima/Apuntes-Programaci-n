# Módulo 08 - Eventos

> **Responder a lo que hace el usuario - clics, escritura, etc**

---

## ¿QÉ son los Eventos?

**Eventos** = acciones del usuario que tu código puede responder:

```jsx
<button onClick={miFunción}>Click me</button>
        ↑                ↑
      Evento         Qué hacer
```

**Analogía:**
- 🔔 Evento = campana
- 👂 Escuchar = addEventListener
- 🎯 Responder = ejecutar función

---

## ¿PARA QÉ sirven?

### 1. Botones (onClick)
```jsx
function Boton() {
  const comprar = () => {
    console.log('¡Comprado!');
    // Lógica de compra
  };
  
  return <button onClick={comprar}>Comprar</button>;
}

// Usuario hace clic → Se ejecuta comprar()
```

### 2. Inputs (onChange)
```jsx
function Buscador() {
  const [busqueda, setBusqueda] = useState('');
  
  const handleChange = (e) => {
    setBusqueda(e.target.value);
    // Mientras escribe: busca automáticamente
  };
  
  return <input onChange={handleChange} />;
}
```

### 3. Formularios (onSubmit)
```jsx
function Formulario() {
  const [email, setEmail] = useState('');
  
  const enviar = (e) => {
    e.preventDefault();  // Evitar recarga
    console.log(`Enviando: ${email}`);
  };
  
  return (
    <form onSubmit={enviar}>
      <input 
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <button type="submit">Enviar</button>
    </form>
  );
}
```

### 4. Mouse (onMouseEnter, onMouseLeave)
```jsx
function Hover() {
  const [sobre, setSobre] = useState(false);
  
  return (
    <div
      onMouseEnter={() => setSobre(true)}
      onMouseLeave={() => setSobre(false)}
      style={{ background: sobre ? 'red' : 'blue' }}
    >
      Pasa el mouse
    </div>
  );
}
```

### 5. Teclado (onKeyPress, onKeyDown)
```jsx
function Buscar() {
  const buscar = (e) => {
    if (e.key === 'Enter') {
      console.log('Buscando...');
    }
  };
  
  return <input onKeyPress={buscar} />;
}
```

---

## ¿CÓMO funcionan?

### Eventos Comunes en React

| Evento | Cuándo se ejecuta | Ejemplo |
|--------|-------------------|---------|
| `onClick` | Clic del mouse | Botón |
| `onChange` | Cambio de valor | Input |
| `onSubmit` | Envío de formulario | Form |
| `onFocus` | Elemento recibe atención | Input |
| `onBlur` | Elemento pierde atención | Input |
| `onHover` / `onMouseEnter` | Hover | Cualquier elemento |
| `onKeyPress` | Se presiona tecla | Input |

### Sintaxis Correcta

```jsx
// ✅ Función nombrada
<button onClick={handleClick}>Click</button>

// ✅ Arrow function
<button onClick={() => console.log('click')}>Click</button>

// ✅ Función con parámetro
<button onClick={() => agregar(5)}>Agregar</button>

// ❌ INCORRECTO (ejecuta automáticamente)
<button onClick={handleClick()}>Click</button>

// ❌ INCORRECTO (string)
<button onClick="handleClick()">Click</button>
```

### Acceso a Eventos (event object)

```jsx
function Input() {
  const handleChange = (e) => {
    // e = event object
    console.log(e.target.value);    // Valor actual
    console.log(e.target.name);     // Nombre del input
  };
  
  return (
    <input 
      onChange={handleChange}
      name="email"
    />
  );
}
```

### Prevenir Comportamiento Defecto

```jsx
function Link() {
  const handleClick = (e) => {
    e.preventDefault();  // ← Evita ir al link
    console.log('Link bloqueado');
  };
  
  return (
    <a href="https://google.com" onClick={handleClick}>
      Google
    </a>
  );
}
```

### Ciclo Completo: Evento

```
1. Usuario hace acción (e.g., clic)
     ↓
2. React detecta el evento
     ↓
3. Ejecuta el handler (función)
     ↓
4. Handler actualiza estado (setState)
     ↓
5. React re-renderiza
     ↓
6. Pantalla se actualiza
```

## Ejercicio
Crea:
- un botón que muestre un mensaje
- un input que imprima lo que escribes
