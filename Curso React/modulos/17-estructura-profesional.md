# Módulo 16 - Estructura Profesional de Proyectos

> **Organizar código para que crezca y sea mantenible**

---

## ¿QÉ es una Estructura Profesional?

**Estructura** = carpetas organizando código:

```
❌ MAL (todo mezclado)
src/
├── App.jsx
├── Home.jsx
├── About.jsx
├── UserDetail.jsx
├── api.js
├── useAPI.js
└── ...30 archivos más...

✅ BIEN (organizado)
src/
├── components/
├── pages/
├── hooks/
├── services/
├── utils/
├── assets/
├── styles/
└── App.jsx
```

**Analogía:**
- 🏠 Estructura = organización de cajones
- 🗂️ Sin estructura = todo en un cajón
- 📦 Con estructura = cada cosa en su lugar

---

## ¿PARA QÉ sirve?

### 1. Reutilizar Componentes
```
components/
├── Button/
│   ├── Button.jsx
│   ├── Button.css
│   └── Button.test.js
├── Card/
│   ├── Card.jsx
│   ├── Card.css
│   └── Card.test.js
└── Header/
    ├── Header.jsx
    ├── Header.css
    └── Header.test.js

// Uso
import Button from '@/components/Button';
import Card from '@/components/Card';
```

### 2. Lógica Reutilizable (Hooks)
```
hooks/
├── useAPI.js       (para fetch)
├── useAuth.js      (autenticación)
├── usePagination.js (paginación)
└── useForm.js      (formularios)

// Uso
const { datos, cargando } = useAPI('/api/usuarios');
const { usuario, login, logout } = useAuth();
```

### 3. Llamadas a API Centralizadas
```
services/
├── api.js          (configuración)
├── usuarios.js     (endpoints de usuarios)
├── productos.js    (endpoints de productos)
└── auth.js         (endpoints de auth)

// Uso
import * as usuariosService from '@/services/usuarios';
usuariosService.getAll();
usuariosService.create(datos);
```

### 4. Utilidades Compartidas
```
utils/
├── formatters.js   (formatear fechas, moneda)
├── validators.js   (validar email, teléfono)
├── constants.js    (constantes globales)
└── helpers.js      (funciones útiles)

// Uso
import { formatDate, capitalizar } from '@/utils/formatters';
```

---

## ¿CÓMO estructurar?

### Estructura Escalable

```
src/
├── components/                 (UI reutilizable)
│   ├── Button/
│   ├── Card/
│   ├── Header/
│   ├── Nav/
│   └── Footer/
│
├── pages/                      (pantallas completas)
│   ├── Home/
│   ├── About/
│   ├── Products/
│   ├── ProductDetail/
│   └── Dashboard/
│
├── hooks/                      (lógica reutilizable)
│   ├── useAPI.js
│   ├── useAuth.js
│   └── useForm.js
│
├── services/                   (API calls)
│   ├── api.js
│   ├── usuarios.js
│   ├── productos.js
│   └── auth.js
│
├── utils/                      (funciones útiles)
│   ├── formatters.js
│   ├── validators.js
│   ├── constants.js
│   └── helpers.js
│
├── styles/                     (CSS global)
│   ├── globals.css
│   └── variables.css
│
├── assets/                     (imágenes, iconos)
│   ├── images/
│   └── icons/
│
├── context/                    (Context API)
│   ├── AuthContext.jsx
│   └── ThemeContext.jsx
│
├── App.jsx                     (componente raíz)
└── main.jsx                    (entrada)
```

### Dentro de Cada Carpeta

```
components/Button/
├── Button.jsx      (lógica)
├── Button.css      (estilos)
├── Button.test.js  (pruebas)
└── index.js        (exportar)

// index.js permite:
// import Button from '@/components/Button'
// En lugar de:
// import Button from '@/components/Button/Button'
```

### Nomenclatura

```
✅ CORRECTO              ❌ INCORRECTO
components/
  Button/                  btn/
    Button.jsx             b.jsx
    Button.css             style.css
    Button.test.js         test.js

pages/
  ProductDetail/           pd/
    ProductDetail.jsx      ProductDetail.jsx
    useProductDetail.js    hook.js
```

### Importaciones Limpias (Absolute)

```javascript
// ❌ Con rutas relativas (confuso)
import Button from '../../../../components/Button';
import useAPI from '../../hooks/useAPI';

// ✅ Con absolute imports (claro)
import Button from '@/components/Button';
import useAPI from '@/hooks/useAPI';

// Configurar en vite.config.js
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
});
```

### Patrón de Componente

```jsx
// components/Button/Button.jsx
import { useCallback } from 'react';
import './Button.css';

/**
 * Botón reutilizable
 * @param {string} variant - primary, secondary
 * @param {function} onClick - callback
 * @param {string} children - texto
 */
function Button({ variant = 'primary', onClick, children, disabled = false }) {
  const handleClick = useCallback(() => {
    onClick?.();
  }, [onClick]);

  return (
    <button
      className={`btn btn-${variant}`}
      onClick={handleClick}
      disabled={disabled}
    >
      {children}
    </button>
  );
}

export default Button;
```

### Patrón de Página

```jsx
// pages/Products/Products.jsx
import { useState, useEffect } from 'react';
import ProductList from '@/components/ProductList';
import { useAPI } from '@/hooks/useAPI';
import * as productsService from '@/services/productos';

function Products() {
  const { datos: productos, cargando, error } = useAPI(productsService.getAll);
  const [filtro, setFiltro] = useState('');

  if (cargando) return <div>Cargando...</div>;
  if (error) return <div>Error: {error}</div>;

  const filtrados = productos.filter(p => 
    p.nombre.toLowerCase().includes(filtro.toLowerCase())
  );

  return (
    <div>
      <h1>Productos</h1>
      <input
        value={filtro}
        onChange={(e) => setFiltro(e.target.value)}
        placeholder="Buscar..."
      />
      <ProductList productos={filtrados} />
    </div>
  );
}

export default Products;
```

### Patrón de Servicio

```jsx
// services/productos.js
const API_BASE = 'https://api.ejemplo.com';

export const getAll = async () => {
  const res = await fetch(`${API_BASE}/productos`);
  if (!res.ok) throw new Error('Error en getAll');
  return res.json();
};

export const getById = async (id) => {
  const res = await fetch(`${API_BASE}/productos/${id}`);
  if (!res.ok) throw new Error('Error en getById');
  return res.json();
};

export const create = async (datos) => {
  const res = await fetch(`${API_BASE}/productos`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(datos),
  });
  if (!res.ok) throw new Error('Error en create');
  return res.json();
};
```

## Ejercicio
Reorganiza un mini proyecto usando carpetas.
