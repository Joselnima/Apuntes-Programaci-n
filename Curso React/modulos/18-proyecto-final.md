# Módulo 18 - Proyecto Final Guiado

> **Construir una aplicación React completa - Tienda de Productos**

---

## ¿QUÉ haremos?

**Proyecto: Sistema de Productos (Tienda)**

Una tienda simple donde puedes:
- ✅ Ver lista de productos
- ✅ Buscar productos
- ✅ Agregar producto
- ✅ Eliminar producto
- ✅ Ver detalle completo
- ✅ Guardar en API

---

## ¿PARA QÉ?

Practicar TODO lo aprendido:

| Concepto | Uso |
|----------|-----|
| **useState** | Guardar productos, búsqueda |
| **useEffect** | Cargar API al iniciar |
| **Props** | Pasar datos entre componentes |
| **Eventos** | Click, onChange |
| **Listas** | Renderizar con .map() |
| **Condicionales** | Mostrar/ocultar según estado |
| **Formularios** | Crear nuevo producto |
| **Estructura** | Carpetas organizadas |

---

## ¿CÓMO?

### Paso 1: Estructura de Carpetas

```
proyecto-tienda/
├── src/
│   ├── components/
│   │   ├── Header/
│   │   │   ├── Header.jsx
│   │   │   └── Header.css
│   │   ├── SearchBar/
│   │   │   ├── SearchBar.jsx
│   │   │   └── SearchBar.css
│   │   ├── ProductForm/
│   │   │   ├── ProductForm.jsx
│   │   │   └── ProductForm.css
│   │   ├── ProductList/
│   │   │   ├── ProductList.jsx
│   │   │   └── ProductList.css
│   │   ├── ProductCard/
│   │   │   ├── ProductCard.jsx
│   │   │   └── ProductCard.css
│   │   └── ProductDetail/
│   │       ├── ProductDetail.jsx
│   │       └── ProductDetail.css
│   │
│   ├── pages/
│   │   ├── Home/
│   │   │   └── Home.jsx
│   │   └── ProductDetailPage/
│   │       └── ProductDetailPage.jsx
│   │
│   ├── services/
│   │   ├── api.js
│   │   └── productos.js
│   │
│   ├── hooks/
│   │   ├── useAPI.js
│   │   └── useProducts.js
│   │
│   ├── App.jsx
│   └── main.jsx
```

### Paso 2: Hook Personalizado (useProducts)

```jsx
// hooks/useProducts.js
import { useState, useEffect } from 'react';

export function useProducts() {
  const [productos, setProductos] = useState([]);
  const [cargando, setCargando] = useState(true);
  const [error, setError] = useState(null);
  
  // Cargar al inicio
  useEffect(() => {
    cargarProductos();
  }, []);
  
  const cargarProductos = async () => {
    try {
      // Simulando API
      const res = await fetch('/api/productos');
      const datos = await res.json();
      setProductos(datos);
    } catch (err) {
      setError(err.message);
    } finally {
      setCargando(false);
    }
  };
  
  const agregar = (producto) => {
    setProductos([...productos, { ...producto, id: Date.now() }]);
  };
  
  const eliminar = (id) => {
    setProductos(productos.filter(p => p.id !== id));
  };
  
  const obtenerPorId = (id) => {
    return productos.find(p => p.id === Number(id));
  };
  
  return {
    productos,
    cargando,
    error,
    agregar,
    eliminar,
    obtenerPorId,
    recargar: cargarProductos
  };
}
```

### Paso 3: Página Principal (Home)

```jsx
// pages/Home/Home.jsx
import { useState } from 'react';
import SearchBar from '@/components/SearchBar';
import ProductList from '@/components/ProductList';
import ProductForm from '@/components/ProductForm';
import Header from '@/components/Header';
import { useProducts } from '@/hooks/useProducts';

function Home() {
  const { productos, cargando, error, agregar, eliminar } = useProducts();
  const [busqueda, setBusqueda] = useState('');
  const [mostrarForm, setMostrarForm] = useState(false);
  
  // Filtrar productos
  const filtrados = productos.filter(p =>
    p.nombre.toLowerCase().includes(busqueda.toLowerCase())
  );
  
  const handleAgregar = (nuevoProducto) => {
    agregar(nuevoProducto);
    setMostrarForm(false);
  };
  
  if (cargando) return <div className="spinner">Cargando...</div>;
  if (error) return <div className="error">Error: {error}</div>;
  
  return (
    <div className="home">
      <Header />
      
      <div className="container">
        <SearchBar 
          valor={busqueda}
          onChange={setBusqueda}
        />
        
        <button 
          className="btn-agregar"
          onClick={() => setMostrarForm(!mostrarForm)}
        >
          {mostrarForm ? 'Cancelar' : '+ Agregar Producto'}
        </button>
        
        {mostrarForm && (
          <ProductForm onAgregar={handleAgregar} />
        )}
        
        <ProductList 
          productos={filtrados}
          onEliminar={eliminar}
        />
        
        {filtrados.length === 0 && (
          <p>No hay productos que coincidan la búsqueda</p>
        )}
      </div>
    </div>
  );
}

export default Home;
```

### Paso 4: Componentes

```jsx
// components/SearchBar/SearchBar.jsx
function SearchBar({ valor, onChange }) {
  return (
    <div className="searchbar">
      <input
        type="text"
        value={valor}
        onChange={(e) => onChange(e.target.value)}
        placeholder="🔍 Buscar productos..."
      />
    </div>
  );
}

export default SearchBar;

// components/ProductForm/ProductForm.jsx
import { useState } from 'react';

function ProductForm({ onAgregar }) {
  const [datos, setDatos] = useState({
    nombre: '',
    precio: '',
    descripción: '',
    imagen: ''
  });
  
  const handleChange = (e) => {
    const { name, value } = e.target;
    setDatos(prev => ({
      ...prev,
      [name]: value
    }));
  };
  
  const handleSubmit = (e) => {
    e.preventDefault();
    
    if (!datos.nombre || !datos.precio) {
      alert('Nombre y precio son requeridos');
      return;
    }
    
    onAgregar({
      ...datos,
      precio: parseFloat(datos.precio)
    });
    
    setDatos({
      nombre: '',
      precio: '',
      descripción: '',
      imagen: ''
    });
  };
  
  return (
    <form onSubmit={handleSubmit} className="form-producto">
      <input
        name="nombre"
        value={datos.nombre}
        onChange={handleChange}
        placeholder="Nombre del producto"
        required
      />
      <input
        name="precio"
        type="number"
        value={datos.precio}
        onChange={handleChange}
        placeholder="Precio"
        step="0.01"
        required
      />
      <textarea
        name="descripción"
        value={datos.descripción}
        onChange={handleChange}
        placeholder="Descripción"
      />
      <input
        name="imagen"
        value={datos.imagen}
        onChange={handleChange}
        placeholder="URL de imagen"
      />
      <button type="submit">Crear Producto</button>
    </form>
  );
}

export default ProductForm;

// components/ProductList/ProductList.jsx
import ProductCard from '@/components/ProductCard';

function ProductList({ productos, onEliminar }) {
  return (
    <div className="producto-grid">
      {productos.map(producto => (
        <ProductCard
          key={producto.id}
          producto={producto}
          onEliminar={onEliminar}
        />
      ))}
    </div>
  );
}

export default ProductList;

// components/ProductCard/ProductCard.jsx
function ProductCard({ producto, onEliminar }) {
  return (
    <div className="producto-card">
      {producto.imagen && (
        <img src={producto.imagen} alt={producto.nombre} />
      )}
      <h3>{producto.nombre}</h3>
      <p className="descripción">{producto.descripción}</p>
      <p className="precio">${producto.precio.toFixed(2)}</p>
      
      <div className="acciones">
        <button className="btn-ver">Ver detalle</button>
        <button
          className="btn-eliminar"
          onClick={() => {
            if (window.confirm('¿Eliminar este producto?')) {
              onEliminar(producto.id);
            }
          }}
        >
          Eliminar
        </button>
      </div>
    </div>
  );
}

export default ProductCard;
```

### Paso 5: Estilo Base

```css
/* App.css */
:root {
  --primary: #007bff;
  --danger: #dc3545;
  --success: #28a745;
  --gray: #f8f9fa;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  background-color: var(--gray);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.searchbar {
  margin: 20px 0;
}

.searchbar input {
  width: 100%;
  padding: 10px 15px;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 16px;
}

.producto-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  margin: 20px 0;
}

.producto-card {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  transition: transform 0.2s;
}

.producto-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
}

.producto-card img {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.producto-card h3 {
  padding: 15px;
  text-align: center;
}

.precio {
  font-size: 24px;
  font-weight: bold;
  color: var(--primary);
  text-align: center;
  padding: 0 15px;
}

.acciones {
  display: flex;
  gap: 10px;
  padding: 15px;
}

button {
  flex: 1;
  padding: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-weight: bold;
  transition: 0.2s;
}

.btn-ver {
  background-color: var(--primary);
  color: white;
}

.btn-eliminar {
  background-color: var(--danger);
  color: white;
}

button:hover {
  opacity: 0.85;
}
```

---

## Retos Adicionales

### Reto 1: Editar Productos
Agregar funcionalidad de editar

### Reto 2: Persistencia (LocalStorage)
Guardar productos en localStorage

### Reto 3: API Real
Conectar con API como JSONPlaceholder

### Reto 4: React Router
Agregar página de detalle (ruta `/producto/:id`)

### Reto 5: Validaciones
Validar formulario con reglas

### Reto 6: Filtros
Filtrar por precio mínimo/máximo

### Reto 7: Carrito
Agregar sistema de carrito

---

## Checklist de Finalización

- ✅ Mostrar lista de productos
- ✅ Buscar productos
- ✅ Agregar producto
- ✅ Eliminar producto
- ✅ Validar formulario
- ✅ Mostrar estados (cargando, error)
- ✅ Componentes organizados
- ✅ Hook personalizado
- ✅ Estilos responsive
- ✅ Sin console.log innecesarios
