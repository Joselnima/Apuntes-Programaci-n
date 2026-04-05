# Módulo 13 - useEffect

> **Ejecutar código cuando el componente se monta o cambia algo - el segundo Hook más importante**

---

## ¿QÉ es useEffect?

**useEffect** = Hook para ejecutar código FUERA del renderizado:

```jsx
useEffect(() => {
  console.log('Se ejecutó');
  // Llamar API
  // Escuchar eventos
  // Configurar temporizadores
}, [])  // ← Dependencies
```

**Analogía:**
- 🎬 Render = actuación principal
- 🎥 useEffect = escenas detrás de cámaras

---

## ¿PARA QÉ sirve?

### 1. Llamar API al Montar (data fetching)
```jsx
function Usuarios() {
  const [usuarios, setUsuarios] = useState([]);
  const [cargando, setCargando] = useState(true);
  
  useEffect(() => {
    // Se ejecuta UNA SOLA VEZ al montar
    fetch('/api/usuarios')
      .then(res => res.json())
      .then(datos => {
        setUsuarios(datos);
        setCargando(false);
      });
  }, []);  // ← [] = solo al montar
  
  if (cargando) return <p>Cargando...</p>;
  
  return (
    <ul>
      {usuarios.map(u => <li key={u.id}>{u.nombre}</li>)}
    </ul>
  );
}
```

### 2. Reaccionar a Cambios
```jsx
function Buscador() {
  const [busqueda, setBusqueda] = useState('');
  const [resultados, setResultados] = useState([]);
  
  useEffect(() => {
    // Se ejecuta cada vez que "busqueda" cambia
    if (busqueda.length > 2) {
      fetch(`/api/buscar?q=${busqueda}`)
        .then(res => res.json())
        .then(datos => setResultados(datos));
    } else {
      setResultados([]);
    }
  }, [busqueda]);  // ← Se ejecuta cuando busqueda cambia
  
  return (
    <div>
      <input 
        value={busqueda}
        onChange={(e) => setBusqueda(e.target.value)}
      />
      <ul>
        {resultados.map(r => <li key={r.id}>{r.nombre}</li>)}
      </ul>
    </div>
  );
}
```

### 3. Limpiar Recursos (cleanup)
```jsx
function Reloj() {
  const [tiempo, setTiempo] = useState(new Date());
  
  useEffect(() => {
    const intervalo = setInterval(() => {
      setTiempo(new Date());
    }, 1000);
    
    // CLEANUP: Se ejecuta cuando componente se desmonta
    return () => {
      clearInterval(intervalo);  // Detener intervalo
    };
  }, []);
  
  return <p>{tiempo.toLocaleTimeString()}</p>;
}
```

### 4. Escuchar Eventos del Navegador
```jsx
function DetectorRedimensión() {
  const [ancho, setAncho] = useState(window.innerWidth);
  
  useEffect(() => {
    const handleResize = () => setAncho(window.innerWidth);
    
    window.addEventListener('resize', handleResize);
    
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);
  
  return <p>Ancho: {ancho}px</p>;
}
```

---

## ¿CÓMO funcionan?

### Los 3 Casos de Dependencies

**Caso 1: Sin Dependencies []**
```jsx
useEffect(() => {
  console.log('Se ejecuta SOLO al montar');
  // Perfecto para:
  // - Llamar API
  // - Inicializar
}, [])  // ← Array vacío
```

**Caso 2: Dependencies específicas**
```jsx
useEffect(() => {
  console.log('Se ejecuta al montar Y cuando id cambia');
  // Se ejecuta si id cambia
}, [id])  // ← Solo cuando id cambia
```

**Caso 3: Sin dependencies**
```jsx
useEffect(() => {
  console.log('Se ejecuta SIEMPRE (mal idea)');
  // ¡EVITAR! - Causa loops infinitos
});  // ← SIN array
```

### Ciclo de Vida

```
MONTAJE:
  1. Componente se crea
  2. Renderiza
  3. useEffect se ejecuta
  
ACTUALIZACIÓN:
  4. User interactúa o props cambian
  5. Componente re-renderiza
  6. useEffect se ejecuta (si dependencies cambiaron)
  
DESMONTAJE:
  7. Componente se borra
  8. Cleanup (return) se ejecuta
```

### Tabla: Cuándo se ejecuta

| Dependencies | Cuándo se ejecuta |
|--------------|------------------|
| `[]` | Al montar (1 vez) |
| `[id]` | Al montar + cuando id cambia |
| `[id, nombre]` | Al montar + cuando id o nombre cambian |
| Sin array | CADA render (¡evitar!) |

### Patrón Completo

```jsx
function Componente({ id }) {
  const [datos, setDatos] = useState(null);
  const [error, setError] = useState(null);
  const [cargando, setCargando] = useState(true);
  
  useEffect(() => {
    // Flag para evitar memory leaks
    let activo = true;
    
    const cargar = async () => {
      try {
        setCargando(true);
        const res = await fetch(`/api/datos/${id}`);
        const resultado = await res.json();
        
        // Solo actualizar si componente sigue montado
        if (activo) {
          setDatos(resultado);
          setError(null);
        }
      } catch (err) {
        if (activo) {
          setError(err.message);
          setDatos(null);
        }
      } finally {
        if (activo) {
          setCargando(false);
        }
      }
    };
    
    cargar();
    
    // Cleanup
    return () => {
      activo = false;
    };
  }, [id]);
  
  if (cargando) return <p>Cargando...</p>;
  if (error) return <p>Error: {error}</p>;
  
  return <div>Datos: {JSON.stringify(datos)}</div>;
}
```

### Errores Comunes

```jsx
// ❌ MALO: Llama API en cada render
useEffect(() => {
  fetch('/api/datos').then(...);
});  // Sin dependencies

// ✅ CORRECTO: Llama API al montar
useEffect(() => {
  fetch('/api/datos').then(...);
}, []);

// ❌ MALO: Crea loop infinito
useEffect(() => {
  setDatos([...datos]); // Actualiza datos
}, [datos]);  // Que causa que datos cambie

// ✅ CORRECTO: Depende de lo correcto
useEffect(() => {
  fetchData(userId);
}, [userId]);
```

## Ejercicio
Muestra un `console.log` cada vez que cambie un input.
