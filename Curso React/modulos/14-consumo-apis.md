# Módulo 14 - Consumo de APIs

> **Traer datos de internet y mostrarlos en React**

---

## ¿QÉ es una API?

**API** = programa remoto que te dá datos:

```
Tu Aplicación              Servidor (API)
(Browser)                  (internet)
    │
    ├─ "Dame usuarios" ──────────↓
    │                        Busca en BD
    │                        Prepara JSON
    │← "Aquí están:" ◄────────┤
    │  [{id:1, name:"Juan"}, ...]

La API es simplemente una URL que devuelve datos
```

**Analogía:**
- 📞 API = línea telefónica
- 🎤 fetch() = persona que llama
- 📋 JSON = información que recibimos

---

## ¿PARA QÉ sirve?

### 1. Traer Usuarios desde API
```jsx
function Usuarios() {
  const [usuarios, setUsuarios] = useState([]);
  const [cargando, setCargando] = useState(true);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    fetch('https://jsonplaceholder.typicode.com/users')
      .then(res => res.json())
      .then(datos => {
        setUsuarios(datos);
        setCargando(false);
      })
      .catch(err => {
        setError(err.message);
        setCargando(false);
      });
  }, []);
  
  if (cargando) return <p>Cargando usuarios...</p>;
  if (error) return <p>Error: {error}</p>;
  
  return (
    <ul>
      {usuarios.map((u) => (
        <li key={u.id}>{u.name}</li>
      ))}
    </ul>
  );
}
```

### 2. Buscar Productos (GET con Filtro)
```jsx
function Tienda() {
  const [productos, setProductos] = useState([]);
  const [categoria, setCategoria] = useState('todos');
  
  useEffect(() => {
    const url = categoria === 'todos' 
      ? '/api/productos'
      : `/api/productos?categoria=${categoria}`;
    
    fetch(url)
      .then(res => res.json())
      .then(datos => setProductos(datos));
  }, [categoria]);  // Se ejecuta cuando categoría cambia
  
  return (
    <div>
      <button onClick={() => setCategoria('libros')}>Libros</button>
      <button onClick={() => setCategoria('electrónica')}>Electrónica</button>
      
      <ul>
        {productos.map(p => <li key={p.id}>{p.nombre}</li>)}
      </ul>
    </div>
  );
}
```

### 3. Crear Nuevo Item (POST)
```jsx
function CrearUsuario() {
  const [nombre, setNombre] = useState('');
  const [enviando, setEnviando] = useState(false);
  
  const crear = async () => {
    setEnviando(true);
    
    try {
      const res = await fetch('/api/usuarios', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nombre })
      });
      
      const nuevoUsuario = await res.json();
      console.log('Creado:', nuevoUsuario);
      setNombre('');
    } catch (err) {
      console.error('Error:', err);
    } finally {
      setEnviando(false);
    }
  };
  
  return (
    <div>
      <input 
        value={nombre}
        onChange={(e) => setNombre(e.target.value)}
        placeholder="Nombre"
      />
      <button onClick={crear} disabled={enviando}>
        {enviando ? 'Creando...' : 'Crear'}
      </button>
    </div>
  );
}
```

### 4. Filtro + Búsqueda (GET con parámetros)
```jsx
function Buscador() {
  const [resultados, setResultados] = useState([]);
  const [busqueda, setBusqueda] = useState('');
  
  useEffect(() => {
    if (busqueda.length < 2) {
      setResultados([]);
      return;
    }
    
    const params = new URLSearchParams({
      q: busqueda,
      limit: 10
    });
    
    fetch(`/api/buscar?${params}`)
      .then(res => res.json())
      .then(datos => setResultados(datos));
  }, [busqueda]);
  
  return (
    <div>
      <input 
        value={busqueda}
        onChange={(e) => setBusqueda(e.target.value)}
        placeholder="Buscar..."
      />
      <ul>
        {resultados.map(r => (
          <li key={r.id}>{r.título}</li>
        ))}
      </ul>
    </div>
  );
}
```

---

## ¿CÓMO funcionan?

### Métodos HTTP Principales

| Método | Propósito | Ejemplo |
|--------|----------|---------|
| **GET** | Obtener datos | Leer lista de usuarios |
| **POST** | Crear nuevo | Registrar usuario |
| **PUT** | Actualizar todo | Cambiar perfil completo |
| **DELETE** | Borrar | Eliminar usuario |
| **PATCH** | Actualizar parcial | Cambiar solo email |

### Patrón fetch() Básico

```jsx
fetch(url, {
  method: 'GET',           // GET, POST, PUT, DELETE
  headers: {               // Meta información
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({

  })                        // Solo para POST/PUT
})
  .then(res => res.json())  // Convertir a JSON
  .then(datos => {...})     // Usar datos
  .catch(err => {...})      // Manejo de errores
```

### Estados de Carga

```jsx
// 3 Estados necesarios
const [datos, setDatos] = useState(null);
const [cargando, setCargando] = useState(true);
const [error, setError] = useState(null);

// Mostrar según estado
if (cargando) return <Spinner />;
if (error) return <Error msg={error} />;
return <Datos datos={datos} />;
```

### Hook Personalizado para APIs

```jsx
function useAPI(url) {
  const [datos, setDatos] = useState(null);
  const [cargando, setCargando] = useState(true);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    let activo = true;
    
    fetch(url)
      .then(res => res.json())
      .then(d => {
        if (activo) {
          setDatos(d);
          setCargando(false);
        }
      })
      .catch(e => {
        if (activo) {
          setError(e.message);
          setCargando(false);
        }
      });
    
    return () => { activo = false; };
  }, [url]);
  
  return { datos, cargando, error };
}

// USO:
function MiApp() {
  const { datos, cargando, error } = useAPI('/api/usuarios');
  
  if (cargando) return <p>Cargando...</p>;
  if (error) return <p>Error: {error}</p>;
  
  return <ul>{datos.map(d => <li key={d.id}>{d.name}</li>)}</ul>;
}
```

### Validación de Respuesta

```jsx
fetch(url)
  .then(res => {
    // ✅ Status 200-299 = success
    if (!res.ok) {
      throw new Error(`HTTP error! status: ${res.status}`);
    }
    return res.json();
  })
  .then(datos => setDatos(datos))
  .catch(error => {
    console.error('Error:', error);
    setError(error.message);
  });
```

## Ejercicio
Consume:
- usuarios
- posts
- productos
