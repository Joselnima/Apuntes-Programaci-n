# Módulo 12 - Formularios

> **Capturar datos del usuario - inputs controlados**

---

## ¿QÉ es un Formulario Controlado?

**Formulario controlado** = React controla el valor del input:

```jsx
// Valor vive en estado de React
const [nombre, setNombre] = useState('');

// Input sincronizado con estado
<input 
  value={nombre}                           // ← Muestra estado
  onChange={(e) => setNombre(e.target.value)}  // ← Actualiza estado
/>

// Todo conectado: User escribe → onChange → setState → Re-render
```

**Analogía:**
- 🎮 Joystick = input
- 🧠 Brain = estado (React)
- 📺 Pantalla = lo que ves

---

## ¿PARA QÉ sirven?

### 1. Validación en Tiempo Real
```jsx
function Email() {
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  
  const handleChange = (e) => {
    const valor = e.target.value;
    setEmail(valor);
    
    if (valor && !valor.includes('@')) {
      setError('Email inválido');
    } else {
      setError('');
    }
  };
  
  return (
    <div>
      <input value={email} onChange={handleChange} />
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
}
```

### 2. Múltiples Campos
```jsx
function Registro() {
  const [datos, setDatos] = useState({
    nombre: '',
    email: '',
    edad: ''
  });
  
  const handleChange = (e) => {
    const { name, value } = e.target;
    setDatos({
      ...datos,
      [name]: value
    });
  };
  
  const enviar = () => {
    console.log('Enviando:', datos);
  };
  
  return (
    <form onSubmit={(e) => { e.preventDefault(); enviar(); }}>
      <input name="nombre" value={datos.nombre} onChange={handleChange} />
      <input name="email" value={datos.email} onChange={handleChange} />
      <input name="edad" value={datos.edad} onChange={handleChange} />
      <button type="submit">Enviar</button>
    </form>
  );
}
```

### 3. Desabilitar Botón Mientras Se Envía
```jsx
function Contacto() {
  const [mensaje, setMensaje] = useState('');
  const [enviando, setEnviando] = useState(false);
  
  const enviar = async () => {
    setEnviando(true);
    await fetch('/api/contacto', {
      method: 'POST',
      body: JSON.stringify({ mensaje })
    });
    setEnviando(false);
  };
  
  return (
    <div>
      <textarea 
        value={mensaje} 
        onChange={(e) => setMensaje(e.target.value)}
      />
      <button disabled={enviando} onClick={enviar}>
        {enviando ? 'Enviando...' : 'Enviar'}
      </button>
    </div>
  );
}
```

### 4. Select y Checkbox
```jsx
function Preferencias() {
  const [pais, setPais] = useState('Chile');
  const [newsletter, setNewsletter] = useState(false);
  
  return (
    <div>
      <select value={pais} onChange={(e) => setPais(e.target.value)}>
        <option>Chile</option>
        <option>Perú</option>
        <option>Argentina</option>
      </select>
      
      <label>
        <input 
          type="checkbox"
          checked={newsletter}
          onChange={(e) => setNewsletter(e.target.checked)}
        />
        Recibir newsletter
      </label>
    </div>
  );
}
```

---

## ¿CÓMO funcionan?

### Componentes de Formulario

| Tipo | Property | Event | Valor |
|------|----------|-------|--------|
| **Input** | `value` | `onChange` | string |
| **Textarea** | `value` | `onChange` | string |
| **Select** | `value` | `onChange` | string |
| **Checkbox** | `checked` | `onChange` | boolean |
| **Radio** | `checked` | `onChange` | boolean |

### Ciclo de Controlado

```
1. User escribe en input
     ↓
2. onChange se ejecuta
     ↓
3. Extrae nuevo valor: e.target.value
     ↓
4. setState actualiza estado
     ↓
5. Componente re-renderiza
     ↓
6. Input muestra NUEVO valor
```

### Patrón Principal

```jsx
// Estado
const [valor, setValor] = useState('');

// Input
<input
  value={valor}                    // ← Sincronizado
  onChange={(e) => setValor(e.target.value)}  // ← Actualiza
/>

// Acceso
console.log(valor);  // Siempre disponible
```

### Formulario Completo

```jsx
function Formulario() {
  const [datos, setDatos] = useState({
    nombre: '',
    email: '',
    mensaje: ''
  });
  
  const handleChange = (e) => {
    const { name, value } = e.target;
    setDatos(prev => ({
      ...prev,
      [name]: value
    }));
  };
  
  const handleSubmit = (e) => {
    e.preventDefault();  // No recarga página
    console.log('Enviando:', datos);
    // Enviar a API
  };
  
  return (
    <form onSubmit={handleSubmit}>
      <input 
        name="nombre"
        value={datos.nombre}
        onChange={handleChange}
        placeholder="Tu nombre"
      />
      <input 
        name="email"
        type="email"
        value={datos.email}
        onChange={handleChange}
        placeholder="Tu email"
      />
      <textarea
        name="mensaje"
        value={datos.mensaje}
        onChange={handleChange}
        placeholder="Tu mensaje"
      />
      <button type="submit">Enviar</button>
    </form>
  );
}
```

### Validación Común

```jsx
// Email válido
const esEmailValido = (email) => /^[^@]+@[^@]+\.[^@]+$/.test(email);

// Teléfono válido
const esTelefonoValido = (tel) => /^\d{10}$/.test(tel);

// Campo no vacío
const esNoVacio = (valor) => valor.trim() !== '';
```

## Ejercicio
Haz un formulario con:
- nombre
- correo
- edad
