# Módulo 25 - Manejo de Estados Avanzado

> **De setState a Redux - Elegir la herramienta correcta**

---

## ¿QUÉ es manejo de estados?

**Manejo de estados = Decidir CÓMO guardar y actualizar datos**

Es como decidir dónde guardar dinero:
- Bolsillo: useState (simple, cerca)
- Cuenta de ahorro: Context (mediano, acceso medio)
- Bóveda de banco: Redux (grande, seguro, complejo)

```
Complejidad → ────────────────────────────────────
             useState → Context → Zustand → Redux
             
Tamaño estado → ────────────────────────────────
             Pequeño ← ─────────────────→ Grande
```

---

## ¿PARA QUÉ?

### 1. useState - Muy Simple

```javascript
// ✅ USO: Un componente, datos locales
function Contador() {
  const [valor, setValor] = useState(0);
  return <button onClick={() => setValor(v => v + 1)}>{valor}</button>;
}

// ❌ NO USAR para:
// - Múltiples componentes
// - Lógica compleja
// - Estado compartido
```

### 2. Context - Global Simple

```javascript
// ✅ USO: Tema, idioma, autenticación
function App() {
  return (
    <ProveedorTema>
      <ProveedorAutenticacion>
        <Aplicacion />
      </ProveedorAutenticacion>
    </ProveedorTema>
  );
}

// ❌ NO USAR para:
// - Mucho estado
// - Cambios frecuentes (performance)
// - Lógica muy compleja
```

### 3. Zustand - Balance

```javascript
// ✅ USO: Balance entre simple y poderoso
// - Carrito de compras
// - Notificaciones
// - UI state

// ❌ NO USAR para:
// - Estado muy grande
// - Equipos que ya usan Redux
```

### 4. Redux - Máquina de Producción

```javascript
// ✅ USO: Aplicaciones GRANDES
// - E-commerce gigante
// - SaaS complejo
// - Múltiples equipos

// ❌ NO USAR para:
// - Proyecto pequeño
// - Cambiar tema app
// - Estado local de componente
```

---

## Comparación: Estados Simples

### useState - Problema

```javascript
// ❌ Mucho código para algo grande
const [usuarios, setUsuarios] = useState([]);
const [cargando, setCargando] = useState(false);
const [error, setError] = useState(null);
const [currentPage, setCurrentPage] = useState(1);
const [filtro, setFiltro] = useState("");
const [ordenar, setOrdenar] = useState("nombre");
// ... 50 líneas más

useEffect(() => { /* 20 líneas de lógica */ }, []);
useEffect(() => { /* validar filtro */ }, [filtro]);
// ... más efectos
```

### Context - Mejor pero...

```javascript
// ✅ Más organizado
// ❌ Re-renders innecesarios si cambias pequeña parte
// Si cambias usuario, re-renderiza TODO el árbol
```

### Zustand - Solución Elegante

```javascript
import { create } from "zustand";

interface UsuariosStore {
  usuarios: any[];
  cargando: boolean;
  error: string | null;
  currentPage: number;
  filtro: string;
  
  setUsuarios: (usuarios: any[]) => void;
  setCargando: (cargando: boolean) => void;
  setError: (error: string | null) => void;
  setFiltro: (filtro: string) => void;
  
  cargar: () => Promise<void>;
  siguiente: () => void;
  anterior: () => void;
}

export const useUsuariosStore = create<UsuariosStore>((set) => ({
  usuarios: [],
  cargando: false,
  error: null,
  currentPage: 1,
  filtro: "",
  
  setUsuarios: (usuarios) => set({ usuarios }),
  setCargando: (cargando) => set({ cargando }),
  setError: (error) => set({ error }),
  setFiltro: (filtro) => set({ filtro, currentPage: 1 }),
  
  cargar: async () => {
    set({ cargando: true });
    try {
      const response = await fetch("/api/usuarios");
      const data = await response.json();
      set({ usuarios: data, cargando: false });
    } catch (err) {
      set({ error: "Error al cargar", cargando: false });
    }
  },
  
  siguiente: () => set((state) => ({ currentPage: state.currentPage + 1 })),
  anterior: () => set((state) => ({ currentPage: state.currentPage - 1 }))
}));

// Usar
function ListaUsuarios() {
  const { usuarios, cargando, error, cargar } = useUsuariosStore();
  
  useEffect(() => {
    cargar();
  }, [cargar]);
  
  if (cargando) return <p>Cargando...</p>;
  if (error) return <p>Error: {error}</p>;
  
  return (
    <ul>
      {usuarios.map(u => <li key={u.id}>{u.nombre}</li>)}
    </ul>
  );
}
```

---

## Redux - Para Casos Grandes

### 1. Setup

```bash
npm install @reduxjs/toolkit react-redux
```

### 2. Crear Slice (acción + reducer)

```typescript
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface Usuario {
  id: number;
  nombre: string;
}

interface UsuariosState {
  items: Usuario[];
  cargando: boolean;
  error: string | null;
}

const estadoInicial: UsuariosState = {
  items: [],
  cargando: false,
  error: null
};

export const usuariosSlice = createSlice({
  name: "usuarios",
  initialState: estadoInicial,
  reducers: {
    setCargando: (state, action: PayloadAction<boolean>) => {
      state.cargando = action.payload;
    },
    setUsuarios: (state, action: PayloadAction<Usuario[]>) => {
      state.items = action.payload;
    },
    setError: (state, action: PayloadAction<string | null>) => {
      state.error = action.payload;
    }
  }
});

export const { setCargando, setUsuarios, setError } = usuariosSlice.actions;
export default usuariosSlice.reducer;
```

### 3. Crear Store

```typescript
import { configureStore } from "@reduxjs/toolkit";
import usuariosReducer from "./usuariosSlice";

export const store = configureStore({
  reducer: {
    usuarios: usuariosReducer
  }
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
```

### 4. Usar en App

```typescript
import { Provider } from "react-redux";

function App() {
  return (
    <Provider store={store}>
      <div className="app">
        <ListaUsuarios />
      </div>
    </Provider>
  );
}
```

### 5. En Componentes

```typescript
import { useDispatch, useSelector } from "react-redux";
import { RootState, AppDispatch } from "./store";
import { setCargando, setUsuarios, setError } from "./usuariosSlice";

function ListaUsuarios() {
  const dispatch = useDispatch<AppDispatch>();
  const { items, cargando, error } = useSelector((state: RootState) => state.usuarios);

  useEffect(() => {
    const cargar = async () => {
      dispatch(setCargando(true));
      try {
        const response = await fetch("/api/usuarios");
        const datos = await response.json();
        dispatch(setUsuarios(datos));
      } catch (err) {
        dispatch(setError(err instanceof Error ? err.message : "Error"));
      }
    };
    
    cargar();
  }, [dispatch]);

  if (cargando) return <p>Cargando...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <ul>
      {items.map(u => <li key={u.id}>{u.nombre}</li>)}
    </ul>
  );
}
```

---

## Tabla de Decisión

| Situación | Herramienta | Por qué |
|-----------|------------|--------|
| **Login de usuario** | useState | Temporal, un componente |
| **Tema app** | Context | Global, cambia poco |
| **Carrito compras** | Zustand | Compartido, caótico |
| **E-commerce gigante** | Redux | Mucho estado, muchos equipos |
| **Contador simple** | useState | ¿Por qué complica? |
| **Notificaciones** | Context o Zustand | Global, simple |
| **Dashboard datos** | Redux | Muchas acciones, historial |
| **Form local** | useState | Solo ese componente |

---

## Patrones Pro

### 1. Selectors en Redux

```typescript
// ❌ Acceso directo
const usuarios = useSelector(state => state.usuarios.items);

// ✅ Selector memorable y reutilizable
export const selectUsuarios = (state: RootState) => state.usuarios.items;
export const selectCargando = (state: RootState) => state.usuarios.cargando;

// Uso
const usuarios = useSelector(selectUsuarios);
const cargando = useSelector(selectCargando);
```

### 2. Async Thunks (operaciones async)

```typescript
import { createAsyncThunk } from "@reduxjs/toolkit";

export const cargarUsuarios = createAsyncThunk(
  "usuarios/cargarUsuarios",
  async (_, { rejectWithValue }) => {
    try {
      const response = await fetch("/api/usuarios");
      if (!response.ok) throw new Error("Error");
      return await response.json();
    } catch (error) {
      return rejectWithValue(error instanceof Error ? error.message : "Error");
    }
  }
);

// En slice
export const usuariosSlice = createSlice({
  name: "usuarios",
  initialState,
  reducers: { /* ... */ },
  extraReducers: (builder) => {
    builder
      .addCase(cargarUsuarios.pending, (state) => {
        state.cargando = true;
      })
      .addCase(cargarUsuarios.fulfilled, (state, action) => {
        state.items = action.payload;
        state.cargando = false;
      })
      .addCase(cargarUsuarios.rejected, (state, action) => {
        state.error = action.payload as string;
        state.cargando = false;
      });
  }
});

// Usar
function ListaUsuarios() {
  const dispatch = useDispatch<AppDispatch>();
  
  useEffect(() => {
    dispatch(cargarUsuarios());
  }, [dispatch]);
  
  // ...
}
```

---

## DevTools - Tiempo de Desarrollo

### Redux DevTools

```typescript
// Automático con Redux Toolkit y la extensión
// Ve cada acción, estado a lo largo del tiempo
// Viaja en el tiempo entre estados ⏰
```

### Zustand DevTools (simple)

```typescript
import { devtools } from "zustand/middleware";

const useStore = create<Store>(
  devtools((set) => ({
    // ...
  }), { name: "MiStore" })
);
```

---

## Resumen Final

### Elegir según:

```
¿Solo un componente?
└── useState

¿Varios componentes, datos simples?
└── Context

¿Varios componentes, pero caótico?
└── Zustand (simple y poderoso)

¿Aplicación GRANDE con muchos datos?
└── Redux (máquina de producción)
```

### Pro Tips

1. **Empieza simple** - useState es suficiente 80% de casos
2. **Refactoriza cuando sientas dolor** - No anticipar problemas
3. **Context es gratis** - Úsalo para global simple
4. **Zustand es felicidad** - Si necesitas algo entre Context y Redux
5. **Redux es profesional** - Equipos grandes lo aman (DevTools, testing)

---

## Próximas Mejoras

Ahora que conoces todos los patrones:

✅ React + TypeScript (módulo 20)
✅ React Hook Form + Zod (módulos 21-22)
✅ Custom Hooks (módulo 23)
✅ Context API (módulo 24)
✅ Manejo de Estados Avanzado (módulo 25)

**Temas opcionales para profundizar:**
- Rendimiento: React.memo, useMemo, useCallback
- Testing: Vitest, RTL (React Testing Library)
- Arquitectura: Atomic Design, Clean Code
- Next.js: Framework base en React
- SSR/SSG: Renderizado en servidor

🚀 **¡Ya eres developer React profesional!**