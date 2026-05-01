# Módulo 15 - React Router

> **Navegación entre "páginas" en una SPA (Single Page Application)**

---

## ¿QÉ es React Router?

**React Router** = navegar en una aplicación sin recargar:

```
URL normal:
Inicio → /       (recarga total)
Sobre → /about   (recarga total)

Con React Router:
Inicio → /       (cambio instantáneo, sin recarga)
Sobre → /about   (cambio instantáneo, sin recarga)
```

**Analogía:**
- 🏫 Escuela normal = ir entre edificios (recarga)
- 🚀 SPA = cambiar cuartos del mismo edificio (sin salir)

---

## ¿PARA QÉ sirve?

### 1. Navegación Múltiple
```jsx
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';

function App() {
  return (
    <BrowserRouter>
      <nav>
        <Link to="/">Inicio</Link>
        <Link to="/about">Sobre</Link>
        <Link to="/contact">Contacto</Link>
      </nav>
      
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/contact" element={<Contact />} />
      </Routes>
    </BrowserRouter>
  );
}
```

### 2. Rutas Dinámicas (parámetros)
```jsx
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/producto/:id" element={<Producto />} />
        <Route path="/usuario/:username" element={<Perfil />} />
      </Routes>
    </BrowserRouter>
  );
}

// Obtener parámetro
function Producto() {
  const { id } = useParams();
  
  return <h1>Producto {id}</h1>;
}

// URL: /producto/42
// id = "42"
```

### 3. Rutas Anidadas (sub-rutas)
```jsx
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/admin" element={<AdminLayout />}>
          <Route path="usuarios" element={<Users />} />
          <Route path="productos" element={<Products />} />
          <Route path="configuración" element={<Settings />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

// /admin/usuarios → AdminLayout + Users
// /admin/productos → AdminLayout + Products
```

### 4. Redirecciones
```jsx
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/admin" element={
          <PrivateRoute>
            <Admin />
          </PrivateRoute>
        } />
        
        {/* Ruta no encontrada */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
}

// Componente protegido
function PrivateRoute({ children }) {
  const { logueado } = useContext(AuthContext);
  
  return logueado ? children : <Navigate to="/login" />;
}
```

---

## ¿CÓMO funcionan?

### Componentes Principales

| Componente | Propósito |
|-----------|----------|
| `<BrowserRouter>` | Wrapper principal (una vez en App) |
| `<Routes>` | Contenedor de rutas |
| `<Route>` | Una ruta específica |
| `<Link>` | Navegar sin recarga |
| `<Navigate>` | Redirigir |
| `useParams()` | Obtener parámetros dinámicos |
| `useNavigate()` | Navegar programáticamente |
| `useLocation()` | Obtener ubicación actual |

### Estructura de Ruta

```jsx
<Routes>
  <Route 
    path="/usuarios/:id"          // ← URL pattern
    element={<UsuarioDetalle />}  // ← Component
  />
</Routes>

// Match:
/usuarios/123  ✅ (id="123")
/usuarios      ❌ (no coincide)
/usuarios/     ✅ (id="")
```

### Tipos de Parámetros

```jsx
// Ruta
<Route path="/post/:id" element={<Post />} />

// Acceso
const { id } = useParams();

// Query strings (búsqueda)
<Route path="/buscar" element={<Buscar />} />
// URL: /buscar?q=react&sort=desc
const searchParams = new URLSearchParams(useLocation().search);
const q = searchParams.get('q');
```

### Navegación Programática

```jsx
function LoginForm() {
  const navigate = useNavigate();
  
  const handleLogin = async (credenciales) => {
    const res = await fetch('/api/login', {
      method: 'POST',
      body: JSON.stringify(credenciales)
    });
    
    if (res.ok) {
      navigate('/dashboard');  // Ir a dashboard
    }
  };
  
  return <form onSubmit={handleLogin}>...</form>;
}
```

### Patrón de Rutas Común

```jsx
function App() {
  return (
    <BrowserRouter>
      <Header />
      <Routes>
        {/* Públicas */}
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/contact" element={<Contact />} />
        
        {/* Autenticación */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        
        {/* Protegidas */}
        <Route 
          path="/dashboard" 
          element={<PrivateRoute><Dashboard /></PrivateRoute>}
        />
        
        {/* Error 404 */}
        <Route path="*" element={<NotFound />} />
      </Routes>
      <Footer />
    </BrowserRouter>
  );
}
```

### Active Link (resaltar actual)

```jsx
import { NavLink } from 'react-router-dom';

function Nav() {
  return (
    <nav>
      <NavLink to="/" className={({ isActive }) => isActive ? 'active' : ''}>
        Home
      </NavLink>
      <NavLink to="/about">
        About
      </NavLink>
    </nav>
  );
}

/* CSS */
a.active {
  font-weight: bold;
  color: blue;
}
```

## Ejercicio
Crea rutas:
- Inicio
- Nosotros
- Contacto
