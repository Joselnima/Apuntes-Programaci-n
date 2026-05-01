# Módulo 24 - Context API - Estado Global

> **Comparte datos entre cualquier componente sin prop drilling**

---

## ¿QUÉ es Context API?

**Context API = Redux light (sin librerías)**

Es como la diferencia entre:
- Props: Pasar datos nivel por nivel como una brigada de bomberos
- Context: Poner datos en un refrigerador central (todos acceden)

```javascript
// ❌ Props (pasar en CADA nivel)
<Abuelo usuario={usuario}>
  <Padre usuario={usuario}>
    <Hijo usuario={usuario}>
      <Nieto usuario={usuario} />  {/* Solo Nieto lo usa! */}
    </Hijo>
  </Padre>
</Abuelo>

// ✅ Context (acceso directo)
<ProveedorUsuario>
  <Abuelo>
    <Padre>
      <Hijo>
        <Nieto />  {/* useContext(UsuarioContext) */}
      </Hijo>
    </Padre>
  </Abuelo>
</ProveedorUsuario>
```

---

## ¿PARA QUÉ?

### Problema 1: Prop Drilling

```javascript
// ❌ Pasar 10 niveles profundo
function App() {
  const usuario = obterUsuario();
  return <Nivel1 usuario={usuario} />;
}

function Nivel1({ usuario }) {
  return <Nivel2 usuario={usuario} />;
}

function Nivel2({ usuario }) {
  return <Nivel3 usuario={usuario} />;
}

function Nivel3({ usuario }) {
  return <Nivel4 usuario={usuario} />;
}

// ... 6 más ...

function Nivel10({ usuario }) {
  return <h1>{usuario.nombre}</h1>;  // Finalmente lo uso
}
```

### Problema 2: Cambios requieren refactorización

```javascript
// Si cambias estructura, refactorizar TODO
```

### Problema 3: Datos globales (tema, usuario, etc)

```javascript
// ❌ Sin Context
const [tema, setTema] = useState("light");
// Pasar tema a TODOS los componentes
// Si 50 componentes lo usan... 😫

// ✅ Con Context
// Todos acceden con useContext(TemaContext)
```

---

## ¿CÓMO funciona?

### 1. Crear Context

```typescript
import { createContext } from "react";

interface Usuario {
  id: number;
  nombre: string;
  email: string;
  rol: "admin" | "usuario";
}

// Crear contexto
export const UsuarioContext = createContext<Usuario | null>(null);
```

### 2. Crear Proveedor (Provider)

```typescript
import { ReactNode, useState } from "react";

export function ProveedorUsuario({ children }: { children: ReactNode }) {
  const [usuario, setUsuario] = useState<Usuario | null>(null);

  const login = async (email: string, password: string) => {
    const response = await fetch("/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password })
    });

    if (response.ok) {
      const datos = await response.json();
      setUsuario(datos);
    }
  };

  const logout = () => {
    setUsuario(null);
  };

  return (
    <UsuarioContext.Provider value={usuario}>
      {children}
    </UsuarioContext.Provider>
  );
}
```

### 3. Usar en App

```typescript
import { ProveedorUsuario } from "./ProveedorUsuario";

function App() {
  return (
    <ProveedorUsuario>
      <div className="app">
        <Header />
        <Sidebar />
        <Main />
      </div>
    </ProveedorUsuario>
  );
}
```

### 4. Acceder desde Componentes

```typescript
import { useContext } from "react";

function Header() {
  const usuario = useContext(UsuarioContext);

  if (!usuario) {
    return <h1>Inicia sesión</h1>;
  }

  return <h1>Hola {usuario.nombre}</h1>;
}

function Sidebar() {
  const usuario = useContext(UsuarioContext);

  if (!usuario?.rol !== "admin") {
    return null;  // No mostrar si no es admin
  }

  return <div>Panel Admin</div>;
}
```

---

## Ejemplo Completo: Tema Global

### 1. Crear Context y Provider

```typescript
import { createContext, useState, ReactNode } from "react";

type Tema = "light" | "dark";

interface TemaContexto {
  tema: Tema;
  cambiarTema: (nuevoTema: Tema) => void;
}

export const TemaContext = createContext<TemaContexto | undefined>(undefined);

export function ProveedorTema({ children }: { children: ReactNode }) {
  const [tema, setTema] = useState<Tema>("light");

  const cambiarTema = (nuevoTema: Tema) => {
    setTema(nuevoTema);
    document.documentElement.setAttribute("data-theme", nuevoTema);
  };

  return (
    <TemaContext.Provider value={{ tema, cambiarTema }}>
      {children}
    </TemaContext.Provider>
  );
}
```

### 2. Custom Hook para usar Context

```typescript
import { useContext } from "react";

export function useTema() {
  const contexto = useContext(TemaContext);
  
  if (!contexto) {
    throw new Error("useTema debe estar dentro de ProveedorTema");
  }
  
  return contexto;
}
```

### 3. Usar en la App

```typescript
import { ProveedorTema } from "./ProveedorTema";

function App() {
  return (
    <ProveedorTema>
      <Header />
      <Main />
      <Footer />
    </ProveedorTema>
  );
}

function Header() {
  const { tema, cambiarTema } = useTema();

  return (
    <header style={{
      background: tema === "light" ? "white" : "#333",
      color: tema === "light" ? "black" : "white"
    }}>
      <button onClick={() => cambiarTema(tema === "light" ? "dark" : "light")}>
        {tema === "light" ? "🌙" : "☀️"}
      </button>
    </header>
  );
}

function Main() {
  const { tema } = useTema();

  return (
    <main style={{
      background: tema === "light" ? "#f5f5f5" : "#1a1a1a",
      color: tema === "light" ? "black" : "white"
    }}>
      {/* Contenido */}
    </main>
  );
}
```

---

## Context Avanzado: Con Actions

```typescript
import { createContext, useReducer, ReactNode } from "react";

type AccionCarrito = 
  | { tipo: "AGREGAR"; producto: Producto }
  | { tipo: "ELIMINAR"; productoId: number }
  | { tipo: "VACIAR" };

interface EstadoCarrito {
  items: Producto[];
  total: number;
}

interface CarritoContexto {
  estado: EstadoCarrito;
  despachar: (accion: AccionCarrito) => void;
}

export const CarritoContext = createContext<CarritoContexto | undefined>(undefined);

function reductorCarrito(estado: EstadoCarrito, accion: AccionCarrito): EstadoCarrito {
  switch (accion.tipo) {
    case "AGREGAR":
      return {
        items: [...estado.items, accion.producto],
        total: estado.total + accion.producto.precio
      };

    case "ELIMINAR":
      const nuevoItems = estado.items.filter(p => p.id !== accion.productoId);
      const nuevoTotal = estado.total - (
        estado.items.find(p => p.id === accion.productoId)?.precio || 0
      );
      return { items: nuevoItems, total: nuevoTotal };

    case "VACIAR":
      return { items: [], total: 0 };

    default:
      return estado;
  }
}

export function ProveedorCarrito({ children }: { children: ReactNode }) {
  const [estado, despachar] = useReducer(reductorCarrito, {
    items: [],
    total: 0
  });

  return (
    <CarritoContext.Provider value={{ estado, despachar }}>
      {children}
    </CarritoContext.Provider>
  );
}

// Usar
export function useCarrito() {
  const contexto = useContext(CarritoContext);
  if (!contexto) throw new Error("useCarrito debe estar en ProveedorCarrito");
  return contexto;
}

// En componente
function ListaProductos() {
  const { despachar } = useCarrito();

  const agregarAlCarrito = (producto: Producto) => {
    despachar({ tipo: "AGREGAR", producto });
  };

  return <button onClick={() => agregarAlCarrito(producto)}>Agregar</button>;
}

function ResumenCarrito() {
  const { estado, despachar } = useCarrito();

  return (
    <div>
      <p>Total: ${estado.total}</p>
      <p>Items: {estado.items.length}</p>
      <button onClick={() => despachar({ tipo: "VACIAR" })}>Vaciar</button>
    </div>
  );
}
```

---

## Múltiples Contextos

```typescript
// Auth Context
export const AuthContext = createContext<AuthContexto | undefined>(undefined);

// Tema Context
export const TemaContext = createContext<TemaContexto | undefined>(undefined);

// Carrito Context
export const CarritoContext = createContext<CarritoContexto | undefined>(undefined);

// Provider raíz
function ProveedorGlobal({ children }: { children: ReactNode }) {
  return (
    <ProveedorAutenticacion>
      <ProveedorTema>
        <ProveedorCarrito>
          {children}
        </ProveedorCarrito>
      </ProveedorTema>
    </ProveedorAutenticacion>
  );
}

// En App
function App() {
  return (
    <ProveedorGlobal>
      {/* Toda app tiene acceso a Auth, Tema, y Carrito */}
    </ProveedorGlobal>
  );
}
```

---

## Context vs Props vs Redux

| Aspecto | Props | Context | Redux |
|--------|-------|---------|-------|
| **Para qué** | Padre → Hijo | Global | Muy global |
| **Complejidad** | Baja | Media | Alta |
| **Estado grande** | ❌ | ⚠️ | ✅ |
| **DevTools** | ❌ | ❌ | ✅ |
| **Performance** | ✅ | ⚠️ | ✅ |
| **Re-renders** | Precisos | Todos suscritos | Optimizado |
| **Curva aprendizaje** | Baja | Media | Alta |

---

## Optimización: useMemo para evitar re-renders

```typescript
import { useMemo } from "react";

export function ProveedorUsuario({ children }: { children: ReactNode }) {
  const [usuario, setUsuario] = useState<Usuario | null>(null);

  // Sin memoizar: cada render crea nuevo objeto
  // Con memoizar: solo nuevo si usuario cambia
  const valor = useMemo(
    () => ({ usuario, setUsuario }),
    [usuario]
  );

  return (
    <UsuarioContext.Provider value={valor}>
      {children}
    </UsuarioContext.Provider>
  );
}
```

---

## Resumen

**Context API es:**
- ✅ Para estado global
- ✅ Sin prop drilling
- ✅ Built-in (sin librerías)
- ✅ Fácil de implementar
- ✅ Perfecto para: tema, autenticación, carrito
- ⚠️ No es para estado muy grande y complejo (usa Redux entonces)

🚀 **Estado global sin complicaciones**