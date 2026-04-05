# Módulo 10 - Renderizado Condicional

> **Mostrar u ocultar componentes según condiciones**

---

## ¿QÉ es Renderizado Condicional?

**Renderizado condicional** = mostrar componentes según condición:

```jsx
// Sin condicional (siempre muestra)
<h1>Bienvenido</h1>

// Con condicional (solo si logueado)
{logueado && <h1>Bienvenido</h1>}
{logueado ? <h1>Hola admin</h1> : <h1>Inicia sesión</h1>}
```

**Analogía:**
- 🚦 Semáforo = condicional
- 🟢 Verde = muestra
- 🔴 Rojo = oculta

---

## ¿PARA QÉ sirve?

### 1. Login / Logout
```jsx
function Header() {
  const [logueado, setLogueado] = useState(false);
  
  return (
    <header>
      {logueado ? (
        <div>
          <p>Bienvenido Juan</p>
          <button onClick={() => setLogueado(false)}>Logout</button>
        </div>
      ) : (
        <button onClick={() => setLogueado(true)}>Login</button>
      )}
    </header>
  );
}
```

### 2. Mostrar Errores
```jsx
function Formulario() {
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  
  const validar = (valor) => {
    if (!valor.includes('@')) {
      setError('Email inválido');
    } else {
      setError('');
    }
    setEmail(valor);
  };
  
  return (
    <div>
      <input onChange={(e) => validar(e.target.value)} />
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
}
```

### 3. Mostrar Diferentes Contenidos
```jsx
function Dashboard({ rol }) {
  return (
    <div>
      {rol === 'admin' && <AdminPanel />}
      {rol === 'usuario' && <UsuarioPanel />}
      {rol === 'guest' && <GuestPanel />}
    </div>
  );
}
```

### 4. Listas Vacías
```jsx
function Carrito({ items }) {
  return (
    <div>
      {items.length > 0 ? (
        <ul>
          {items.map(item => <li key={item.id}>{item.nombre}</li>)}
        </ul>
      ) : (
        <p>Tu carrito está vacío</p>
      )}
    </div>
  );
}
```

---

## ¿CÓMO funcionan?

### 3 Patrones Principales

**Patrón 1: Ternario (if-else)**
```jsx
{condición ? <ComponenteA /> : <ComponenteB />}

// Ejemplo
{esAdmin ? <DeleteButton /> : <ViewButton />}
```

**Patrón 2: AND (&&)**
```jsx
{condición && <Componente />}

// Ejemplo: solo muestra si esLogueado es true
{esLogueado && <h1>Hola usuario</h1>}
```

**Patrón 3: Elemento Temporal**
```jsx
let contenido;

if (estado === 'cargando') {
  contenido = <Spinner />;
} else if (estado === 'error') {
  contenido = <Error />;
} else {
  contenido = <Datos />;
}

return contenido;
```

### Comparación de Patrones

| Situación | Usa | Código |
|-----------|-----|--------|
| If-else simple | Ternario | `x ? A : B` |
| Si o No | AND | `x && A` |
| Múltiples casos | Switch o if/else | Variable |
| Complejo | Variable | `let content; if...` |

### Flujo Mental

```
Componente renderiza
     ↓
Evalúa condición
     ↓
¿Verdadero?
  ├─ Sí → Renderiza Componente A
  └─ No → Renderiza Componente B o nada
     ↓
Pantalla se actualiza
```

### Validaciones Comunes

```jsx
// Toggle
{showMenu && <Menu />}

// Múltiples niveles
{user && user.admin && <AdminPanel />}

// Negación
{!isLoading && <Data />}

// Comparación
{count > 0 && <ShowCount count={count} />}

// Existencia
{items.length > 0 && <ItemsList items={items} />}
```

### Evitar Errores Comunes

```jsx
// ❌ MALO (renderiza "false" en pantalla)
{false}
{0}
{undefined}

// ✅ CORRECTO (no renderiza nada)
{condición && <Componente />}

// ❌ MALO (null + "text" = problema)
{error || ""}

// ✅ CORRECTO
{error && <p>{error}</p>}
```

## Ejercicio
Muestra:
- “Stock disponible” si stock > 0
- “Agotado” si stock = 0
