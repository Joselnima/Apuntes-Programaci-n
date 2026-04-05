# Módulo 17 - Buenas Prácticas en React

> **Escribir código profesional que escala y es fácil mantener**

---

## ¿QÉ son Buenas Prácticas?

**Buenas prácticas** = patrones probados que evitan problemas:

```jsx
// ❌ MALO (funciona pero desordenado)
function Page() {
  const [user, setUser] = useState();
  useEffect(() => { fetch(...) }, []);
  const handleX = () => { ... };
  const handleY = () => { ... };
  return <div>... 200 líneas ...</div>;
}

// ✅ CORRECTO (claro, mantenible)
function UserPage() {
  const { user, loading } = useAPI('/api/user');
  return <UserDetail user={user} />;
}
```

**Analogía:**
- 🏗️ Construcción = buenas prácticas = casa duradera
- 🏚️ Sin prácticas = casa que se cae

---

## ¿PARA QÉ sirven?

### 1. Código Legible (tu futuro te agradece)
```jsx
// ❌ CONFUSO
const u = (d) => d.map(x => <li>{x.a}</li>);

// ✅ CLARO
const UserList = (usuarios) => 
  usuarios.map(usuario => <li>{usuario.nombre}</li>);
```

### 2. Reutilizable (no repetir)
```jsx
// ❌ Repite en 3 lugares
<button onClick={handleDelete}>Eliminar</button>
<button onClick={handleDelete}>Eliminar</button>
<button onClick={handleDelete}>Eliminar</button>

// ✅ Componente único
<DeleteButton onClick={handleDelete} />
```

### 3. Testeable (validar sin bugs)
```jsx
// ✅ TESTEABLE (función pura)
export const calcularTotal = (items) => {
  return items.reduce((sum, item) => sum + item.precio, 0);
};

// ❌ NO TESTEABLE (lógica en componente)
function Carrito() {
  const total = ... // Lógica mezclada
}
```

### 4. Performance (rápido)
```jsx
// ❌ Re-renderiza TODO
{items.map((item, index) => (
  <Item key={index} />  // Key inestable
))}

// ✅ Renderiza solo lo necesario
{items.map((item) => (
  <Item key={item.id} />  // Key estable
))}
```

---

## ¿CÓMO aplicarlas?

### 10 Reglas de Oro

| Regla | Cómo | Mal ❌ | Bien ✅ |
|-------|------|--------|---------|
| **Componentes pequeños** | Max 200 líneas | Componente gigante | Dividir 3 componentes |
| **Un trabajo** | Una sola responsabilidad | Hace todo | Un componente por tarea |
| **Props claros** | Nombres descriptivos | `data`, `fn` | `usuario`, `onDelete` |
| **State mínimo** | Solo lo que cambia | `state = {}` enorme | Solo necesario |
| **Keys únicos** | ID real, no index | `key={i}` | `key={item.id}` |
| **useEffect limpio** | Evitar loops | Sin dependencies | `[id]` correcto |
| **Nombres claros** | `getUserById()` | `fn()`, `getData` | `fetchUserProfile` |
| **Sin lógica en JSX** | Extraer a función | Condicionales complejos | Función separada |
| **Validar props** | PropTypes/TypeScript | Sin validación | PropTypes.string |
| **Reutilizar hooks** | Si repite 2x | copiar código | Custom hook |

### Patrón: División de Responsabilidades

```jsx
// ❌ TODO EN UNOCOMPONENTE
function ProductPage() {
  const [products, setProducts] = useState([]);
  const [search, setSearch] = useState('');
  const [cargando, setCargando] = useState(true);
  
  useEffect(() => {
    fetch('/api/products')
      .then(r => r.json())
      .then(d => setProducts(d));
  }, []);
  
  const filtered = products.filter(p => 
    p.name.includes(search)
  );
  
  return (
    <div>
      <input onChange={e => setSearch(e.target.value)} />
      <ul>
        {filtered.map(p => (
          <li key={p.id}>
            <h3>{p.name}</h3>
            <button onClick={() => { /* delete */ }}>Delete</button>
            <button onClick={() => { /* edit */ }}>Edit</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

// ✅ BIEN DIVIDIDO
// pages/ProductPage.jsx
function ProductPage() {
  const { products, loading } = useAPI('/api/products');
  const [search, setSearch] = useState('');
  
  if (loading) return <Spinner />;
  
  return (
    <div>
      <SearchBar value={search} onChange={setSearch} />
      <ProductList products={products} search={search} />
    </div>
  );
}

// components/SearchBar.jsx
function SearchBar({ value, onChange }) {
  return (
    <input
      value={value}
      onChange={(e) => onChange(e.target.value)}
      placeholder="Buscar..."
    />
  );
}

// components/ProductList.jsx
function ProductList({ products, search }) {
  const filtered = products.filter(p => 
    p.name.toLowerCase().includes(search.toLowerCase())
  );
  
  return (
    <ul>
      {filtered.map(p => (
        <ProductCard key={p.id} product={p} />
      ))}
    </ul>
  );
}

// components/ProductCard.jsx
function ProductCard({ product }) {
  return (
    <li>
      <h3>{product.name}</h3>
      <ProductActions product={product} />
    </li>
  );
}

// components/ProductActions.jsx
function ProductActions({ product }) {
  const handleDelete = () => { /* delete */ };
  const handleEdit = () => { /* edit */ };
  
  return (
    <div>
      <button onClick={handleDelete}>Delete</button>
      <button onClick={handleEdit}>Edit</button>
    </div>
  );
}
```

### Patrón: Funciones Puras

```jsx
// ❌ IMPURA (depende de afuera)
let total = 0;
function addPrice(price) {
  total += price;  // Modifica variable global
  return total;
}

// ✅ PURA (solo usa parámetros)
function calculateTotal(prices) {
  return prices.reduce((sum, price) => sum + price, 0);
}

// En React
<button onClick={() => {
  const total = calculateTotal(cartItems);
  setTotal(total);
}}>
  Calculate
</button>
```

### Patrón: useEffect Seguro

```jsx
// ❌ MALO (fuga de memoria)
useEffect(() => {
  fetch('/api/user')
    .then(r => r.json())
    .then(d => setState(d));  // Puede setState si desmontado
}, []);

// ✅ CORRECTO (flag)
useEffect(() => {
  let mounted = true;
  
  fetch('/api/user')
    .then(r => r.json())
    .then(d => {
      if (mounted) setState(d);  // Solo si montado
    });
  
  return () => { mounted = false; };
}, []);
```

### Patrón: Nombres Descriptivos

```jsx
// ❌ MALO
const handleClick = () => {};
const data = [];
const fn = (x) => x * 2;

// ✅ CORRECTO
const handleDeleteUser = () => {};
const usersList = [];
const doubleValue = (number) => number * 2;
```

## Ejercicio
Toma un componente grande y sepáralo en 3.

---

## Memory Leaks - El Enemigo Invisible

### ¿QUÉ es un Memory Leak?

**Memory Leak = Memoria que se queda en la RAM aunque no la uses**

Es como:
- 🚗 Dejar el auto en marcha mientras trabajas (gasta gasolina sin usar)
- 💧 Grifo que gotea (pierde agua todo el tiempo)

```javascript
// ❌ MEMORY LEAK (suscripción activa pero sin usar)
useEffect(() => {
  const subscription = socket.on('message', handleMessage);
  // Si el componente se desmonta...
  // La suscripción sigue activa 👻
}, []);

// ✅ CORRECTO (cleanup)
useEffect(() => {
  const subscription = socket.on('message', handleMessage);
  
  return () => {
    subscription.unsubscribe();  // Limpieza
  };
}, []);
```

### Common Memory Leaks

#### 1. Timers sin limpiar

```javascript
// ❌ MEMORY LEAK
useEffect(() => {
  const interval = setInterval(() => {
    console.log("Cada segundo");
  }, 1000);
  
  // Si el componente se desmonta, interval sigue corriendo
  // Memoria aumenta infinitamente 📈
}, []);

// ✅ CORRECTO
useEffect(() => {
  const interval = setInterval(() => {
    console.log("Cada segundo");
  }, 1000);
  
  return () => {
    clearInterval(interval);  // Limpieza
  };
}, []);
```

#### 2. Event Listeners sin remover

```javascript
// ❌ MEMORY LEAK
useEffect(() => {
  window.addEventListener('resize', handleResize);
  // Si desmonta, listener sigue escuchando
}, []);

// ✅ CORRECTO
useEffect(() => {
  window.addEventListener('resize', handleResize);
  
  return () => {
    window.removeEventListener('resize', handleResize);
  };
}, []);
```

#### 3. Suscripciones a observables

```javascript
// ❌ MEMORY LEAK (RxJS)
useEffect(() => {
  this.data$ = dataService.getData().subscribe(data => {
    setState(data);
  });
  // Subscription sigue activa después de desmontar
}, []);

// ✅ CORRECTO
useEffect(() => {
  const subscription = dataService.getData().subscribe(data => {
    setState(data);
  });
  
  return () => {
    subscription.unsubscribe();
  };
}, []);
```

#### 4. Referencias que no se limpian

```javascript
// ❌ MEMORY LEAK
let globalRef = null;

function Component() {
  useEffect(() => {
    globalRef = document.querySelector('#big-element');
    // La referencia persiste aunque el elemento se borre
  }, []);
}

// ✅ CORRECTO
function Component() {
  const ref = useRef<HTMLDivElement>(null);
  
  useEffect(() => {
    // Ref se limpia automáticamente
    return () => {
      ref.current = null;
    };
  }, []);
  
  return <div ref={ref} />;
}
```

#### 5. setState después de desmontar

```javascript
// ❌ MEMORY LEAK (advertencia en consola)
useEffect(() => {
  fetch('/api/data')
    .then(r => r.json())
    .then(data => {
      setState(data);  // ¿Y si desmontó antes?
    });
}, []);

// ✅ CORRECTO (con flag)
useEffect(() => {
  let isMounted = true;
  
  fetch('/api/data')
    .then(r => r.json())
    .then(data => {
      if (isMounted) {  // Verificar si montado
        setState(data);
      }
    });
  
  return () => {
    isMounted = false;  // Desmontar = false
  };
}, []);

// ✅ MÁS MODERNO (AbortController)
useEffect(() => {
  const controller = new AbortController();
  
  fetch('/api/data', { signal: controller.signal })
    .then(r => r.json())
    .then(data => setState(data))
    .catch(err => {
      if (err.name !== 'AbortError') console.error(err);
    });
  
  return () => {
    controller.abort();  // Abortar fetch
  };
}, []);
```

---

## Errores Comunes que Rompen la App

### 1. Infinite Loops (Loop Infinito)

```javascript
// ❌ LOOP INFINITO
function Component() {
  const [count, setCount] = useState(0);
  
  useEffect(() => {
    setCount(count + 1);
  }, [count]);  // count cambia → efecto → setCount → count cambia → ...
  
  // Re-render infinito 💥
}

// ✅ CORRECTO
useEffect(() => {
  setCount(c => c + 1);
}, []);  // Una sola vez

// ✅ O SI NECESITAS DEPENDENCY
useEffect(() => {
  // Hacer algo cuando count cambia
  console.log("Count es:", count);
}, [count]);  // Sin cambiar count aquí
```

### 2. Closure Stale (Variable vieja)

```javascript
// ❌ STALE CLOSURE
function Counter() {
  const [count, setCount] = useState(0);
  
  const handleClick = () => {
    setTimeout(() => {
      console.log(count);  // Siempre 0 (old value)
    }, 1000);
  };
  
  return <button onClick={handleClick}>Click</button>;
}

// Si clickeas rápido y esperas 1 segundo, muestra 0
// No importa cuántas veces clickees

// ✅ CORRECTO (setter funcional)
const handleClick = () => {
  setTimeout(() => {
    setCount(c => {
      console.log(c);  // c es el valor ACTUAL
      return c;
    });
  }, 1000);
};
```

### 3. Key inestable

```javascript
// ❌ MALO (key = index)
{items.map((item, index) => (
  <Item key={index} item={item} />
))}

// Problema:
// Si borras del medio, los índices cambian
// Item 0 → Item 1 (por el borrado) → React lo re-renderiza todo
// State se pierde

// ✅ CORRECTO
{items.map(item => (
  <Item key={item.id} item={item} />
))}

// item.id es único y estable
```

### 4. Modificar estado directamente

```javascript
// ❌ NUNCA HAGAS ESTO
const usuario = usuarios[0];
usuario.nombre = "Nuevo";  // ❌ Mutación
setUsuarios(usuarios);     // React no detecta cambio

// ✅ CORRECTO (copia)
const usuarioActualizado = { ...usuarios[0], nombre: "Nuevo" };
setUsuarios([usuarioActualizado, ...usuarios.slice(1)]);

// ✅ O MÁS SIMPLE
setUsuarios(usuarios.map(u => 
  u.id === usuarios[0].id ? { ...u, nombre: "Nuevo" } : u
));
```

### 5. Dependencias faltantes en useEffect

```javascript
// ❌ PELIGROSO (dependencies vacías pero usa variables)
function Component() {
  const [id, setId] = useState(1);
  
  useEffect(() => {
    fetch(`/api/user/${id}`)  // Usa id
      .then(r => r.json())
      .then(data => setUser(data));
  }, []);  // ❌ Falta id
}

// Si id cambia, el efecto no se ejecuta
// Cargan datos del user 1 siempre

// ✅ CORRECTO
useEffect(() => {
  fetch(`/api/user/${id}`)
    .then(r => r.json())
    .then(data => setUser(data));
}, [id]);  // ✅ Incluir id
```

---

## Performance - Haciendo la App Rápida

### ¿QUÉ afecta el performance?

1. Re-renders innecesarios
2. Componentes grandes
3. Computaciones pesadas
4. Efectos sin limpieza
5. Imágenes sin optimizar

### Re-renders Innecesarios

```javascript
// ❌ TODO re-renderiza
function App() {
  const [user, setUser] = useState(null);
  const [tema, setTema] = useState("light");
  
  return (
    <UserProvider value={user}>  {/* Si tema cambia, User re-renderiza */}
      <Theme value={tema} />
    </UserProvider>
  );
}

// ✅ Separar contextos
function App() {
  return (
    <UserProvider>       {/* Cambios de user */}
      <ThemeProvider>    {/* Cambios de tema */}
        <Content />
      </ThemeProvider>
    </UserProvider>
  );
}
```

### React.memo - Evitar re-renders

```javascript
// ❌ RE-RENDERIZA SIEMPRE
function UserCard({ user }) {
  console.log("Renderizando...");  // Se ejecuta cada vez
  return <h1>{user.nombre}</h1>;
}

// ✅ MEMORIZAR (solo si props cambian)
const UserCard = React.memo(function UserCard({ user }) {
  console.log("Renderizando...");  // Solo si user cambió
  return <h1>{user.nombre}</h1>;
});

// O con función comparativa
const UserCard = React.memo(
  ({ user }) => <h1>{user.nombre}</h1>,
  (prevProps, nextProps) => prevProps.user.id === nextProps.user.id
);
```

### useMemo - Cálculos caros

```javascript
// ❌ Se recalcula cada render
function ListaFiltrada({ items, filtro }) {
  const filtrados = items.filter(i => 
    i.nombre.includes(filtro)
  );  // Con 10,000 items = LENTO
  
  return <div>{filtrados.map(...)}</div>;
}

// ✅ Memorizar resultado
function ListaFiltrada({ items, filtro }) {
  const filtrados = useMemo(
    () => items.filter(i => i.nombre.includes(filtro)),
    [items, filtro]  // Solo recalcular si items o filtro cambian
  );
  
  return <div>{filtrados.map(...)}</div>;
}
```

### useCallback - Evitar propagar re-renders

```javascript
// ❌ Nueva función cada render
function Parent() {
  const handleClick = () => { /* ... */ };  // Nueva cada vez
  
  return <Child onClick={handleClick} />;  // Child re-renderiza siempre
}

// ✅ Misma función en memoria
function Parent() {
  const handleClick = useCallback(() => {
    /* ... */
  }, []);  // Misma función siempre
  
  return <Child onClick={handleClick} />;  // Child solo si necesario
}
```

### Code Splitting - Cargar componentes cuando sea necesario

```javascript
// ❌ TODO en el bundle
import Administrador from './pages/Administrador';
import Dashboard from './pages/Dashboard';

function App() {
  return admin ? <Administrador /> : <Dashboard />;
}
// Bundle gigante, carga lenta

// ✅ Lazy loading (solo cuando necesitas)
const Administrador = React.lazy(() => 
  import('./pages/Administrador')
);
const Dashboard = React.lazy(() => 
  import('./pages/Dashboard')
);

function App() {
  return (
    <Suspense fallback={<div>Cargando...</div>}>
      {admin ? <Administrador /> : <Dashboard />}
    </Suspense>
  );
}
```

### Imágenes Optimizadas

```javascript
// ❌ Carga toda la imagen siempre
<img src="imagen-gigante.jpg" alt="..." />

// ✅ Múltiples tamaños (responsive)
<img 
  srcSet="small.jpg 400w, medium.jpg 800w, large.jpg 1200w"
  src="medium.jpg"
  alt="..."
/>

// ✅ Lazy loading
<img 
  src="imagen.jpg" 
  alt="..."
  loading="lazy"  {/* Cargar cuando sea visible */}
/>
```

---

## Tabla: Errores vs Solución

| Problema | Síntoma | Solución |
|----------|---------|----------|
| **Memory Leak** | RAM sube con el tiempo 📈 | Limpiar en return `useEffect` |
| **Loop Infinito** | App se congela 🔄 | Verificar dependencies en `useEffect` |
| **Stale Closure** | Datos anticuados 👻 | Usar setter funcional o ref |
| **Key inestable** | Errores al borrar/añadir | Usar ID único, no index |
| **Mutación** | React no detecta cambios | Siempre hacer copia: `{...obj}` |
| **Dependencies faltantes** | Datos no se actualizan | Incluir todas las variables usadas |
| **Re-renders extra** | App lenta 🐌 | Usar `React.memo`, `useMemo` |
| **Cálculos pesados** | Frame drops ⚠️ | Memorizar con `useMemo` |
| **Imagen 4MB** | Página lenta a cargar | Optimizar y responsive images |

---

## Checklist de Performance

```
❌ Antes de producción, verificar:

□ No memory leaks (DevTools)
□ No infinite loops
□ No mutaciones de estado
□ Dependencies correctas en useEffect
□ Componentes grandes divididos
□ React.memo donde necesario
□ useMemo para cálculos caros
□ useCallback para funciones en props
□ Imágenes optimizadas
□ Code splitting implementado
□ No console.log en producción
□ Cleanup en todos los efectos
```

---

---

## Resumen Final

### Buenas Prácticas = 3 Pilares

**1. Código Limpio**
- Componentes pequeños (max 200 líneas)
- Nombres descriptivos
- Una responsabilidad por componente
- Reutilización

**2. Sin Memory Leaks**
- Cleanup en useEffect
- Limpiar timers, listeners, suscripciones
- Verificar si montado antes de setState
- Usar AbortController para fetch

**3. Performance Rápido**
- Evitar re-renders innecesarios
- React.memo para componentes puros
- useMemo para cálculos caros
- useCallback para funciones en props
- Code splitting con lazy loading
- Optimizar imágenes

### Beneficios Reales

```
Código profesional = 
  👷 Mantenible
  🚀 Rápido
  🐛 Menos bugs
  😊 Felicidad del equipo
  💰 Ahorro de tiempo
```

---

## Ejercicios Prácticos

**1. Identifica Memory Leaks**
```javascript
// ¿Dónde está el memory leak?
useEffect(() => {
  const listener = window.addEventListener('scroll', handle);
}, []);
```

**2. Optimiza un componente grande**
Toma 500 líneas y córtalo en 3 componentes con responsabilidades claras

**3. Implementa cleanup**
Crea un componente que:
- Escuche window.resize
- Haga fetch a API
- Setee un timeout
- Limpie TODO

**Soluciones en próximos módulos** ✅

---

🏆 **Ahora escribes código que escala y es fácil mantener**
