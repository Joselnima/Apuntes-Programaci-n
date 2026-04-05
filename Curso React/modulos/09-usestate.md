# Módulo 09 - useState

> **La memoria de React - cómo guardar y cambiar datos en pantalla**

---

## ¿QÉ es useState?

**useState** = Hook (herramienta) para guardar información:

```jsx
const [contador, setContador] = useState(0);
       ↑              ↑              ↑
    variable      función         valor inicial
```

**Analogía:**
- 🧠 Brain = tu computadora
- 📝 Memoria = useState (recuerda cosas)
- 🔄 Cambio = setContador (actualiza memoria)

---

## ¿PARA QÉ sirve?

### 1. Contar clics
```jsx
function Contador() {
  const [contador, setContador] = useState(0);
  
  return (
    <div>
      <p>Clics: {contador}</p>
      <button onClick={() => setContador(contador + 1)}>
        Sumar
      </button>
    </div>
  );
}

// Cada clic: contador aumenta
// React re-renderiza automáticamente
```

### 2. Guardar Texto
```jsx
function Formulario() {
  const [nombre, setNombre] = useState('');
  
  return (
    <input
      value={nombre}
      onChange={(e) => setNombre(e.target.value)}
    />
  );
}

// Mientras escribes: nombre se actualiza
// Pantalla refleja cambios al instante
```

### 3. Toggle (encendido/apagado)
```jsx
function Menu() {
  const [cerrado, setCerrado] = useState(true);
  
  return (
    <div>
      <button onClick={() => setCerrado(!cerrado)}>
        Menú
      </button>
      {!cerrado && <nav>Opciones...</nav>}
    </div>
  );
}

// Cada clic: toggle true/false
```

### 4. Mantener Sincronizado
```jsx
function Carrito() {
  const [items, setItems] = useState([]);
  
  const agregar = (item) => {
    setItems([...items, item]);  // Actualiza estado
    // React automáticamente:
    // - Re-renderiza
    // - Muestra new items
    // - Actualiza contador
  };
}
```

---

## ¿CÓMO funciona?

### Fórmula useState

```jsx
const [valor, setValor] = useState(inicial);
```

**Ejemplo paso a paso:**
```jsx
// 1. Inicializo
const [edad, setEdad] = useState(25);
//              ↓ edad = 25 al inicio

// 2. Usado en pantalla
<p>{edad}</p>  // Muestra 25

// 3. Cambio
<button onClick={() => setEdad(26)}>Cumpleaños</button>

// 4. Automático:
//    - edad ahora es 26
//    - React re-renderiza
//    - Pantalla muestra 26
```

### Ciclo de useState

```
1. Usuario hace clic
     ↓
2. onClick ejecuta setContador(valor_nuevo)
     ↓
3. React detecta cambio de estado
     ↓
4. Componente se re-renderiza
     ↓
5. Pantalla se actualiza
```

### Reglas Importantes

| Regla | Sí ✅ | No ❌ |
|-------|------|------|
| **En la raíz** | `function Comp() { useState... }` | `if (true) { useState... }` |
| **Nunca modificar** | `setValor(nuevo)` | `valor = nuevo` |
| **Usar setter** | `setArray([...arr, item])` | `arr.push(item)` |
| **Previo state** | `setCount(c => c + 1)` | `setCount(count + 1)` |

### Comparación: Sin useState vs Con useState

```jsx
// ❌ MALO (no funciona en React)
function Contador() {
  let contador = 0;  // Variable normal
  return (
    <button onClick={() => contador++}>
      {contador}  // Nunca cambia en pantalla
    </button>
  );
}

// ✅ CORRECTO (React lo ve)
function Contador() {
  const [contador, setContador] = useState(0);
  return (
    <button onClick={() => setContador(contador + 1)}>
      {contador}  // Cambia automáticamente
    </button>
  );
}
```

## Ejercicio
Haz:
- contador +
- contador -
- reset
