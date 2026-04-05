# Módulo 23 - Custom Hooks - Lógica Reutilizable

> **Hooks personalizados - Comparte lógica entre componentes**

---

## ¿QUÉ es un Custom Hook?

**Custom Hook = Función que reutiliza lógica con hooks**

Es como la diferencia entre:
- Copy-paste: Copias el mismo código en 10 componentes
- Custom Hook: Escribes UNA VEZ, usas en 100 componentes

```javascript
// ❌ Repetir 10 veces en 10 componentes
const [usuarios, setUsuarios] = useState([]);
const [cargando, setCargando] = useState(false);
const [error, setError] = useState(null);

useEffect(() => {
  setCargando(true);
  fetch("/api/usuarios")
    .then(r => r.json())
    .then(data => setUsuarios(data))
    .catch(e => setError(e))
    .finally(() => setCargando(false));
}, []);

// ✅ Custom Hook - UNA VEZ
const useUsuarios = () => {
  const [usuarios, setUsuarios] = useState([]);
  const [cargando, setCargando] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    setCargando(true);
    fetch("/api/usuarios")
      .then(r => r.json())
      .then(data => setUsuarios(data))
      .catch(e => setError(e))
      .finally(() => setCargando(false));
  }, []);

  return { usuarios, cargando, error };
};

// Usar en cualquier componente
function MisComponentes() {
  const { usuarios, cargando, error } = useUsuarios();
  // Una línea, toda la lógica lista
}
```

---

## ¿PARA QUÉ?

### Problema 1: Lógica repetida

```javascript
// ❌ Misma lógica en Componente1, Componente2, Componente3
function Componente1() {
  const [datos, setDatos] = useState([]);
  useEffect(() => {
    fetch("/api/datos").then(r => r.json()).then(setDatos);
  }, []);
  return <div>{datos.length} items</div>;
}

function Componente2() {
  const [datos, setDatos] = useState([]);
  useEffect(() => {
    fetch("/api/datos").then(r => r.json()).then(setDatos);
  }, []);
  return <div>{datos.length} items</div>;
}

function Componente3() {
  const [datos, setDatos] = useState([]);
  useEffect(() => {
    fetch("/api/datos").then(r => r.json()).then(setDatos);
  }, []);
  return <div>{datos.length} items</div>;
}
```

### Problema 2: Mantener cambios en una sola línea

```javascript
// Cambias el custom hook UNA VEZ
// y automáticamente todos los 100+ componentes que lo usan se actualizan
```

### Problema 3: Lógica compleja esparcida

```javascript
// ❌ Componente gigante con lógica mixta
function ComponenteBigote() {
  // Estado
  const [usuarios, setUsuarios] = useState([]);
  const [filtro, setFiltro] = useState("");
  const [ordenar, setOrdenar] = useState("nombre");

  // Efectos
  useEffect(() => { /* cargar */ }, []);
  useEffect(() => { /* filtrar */ }, [filtro]);
  useEffect(() => { /* ordenar */ }, [ordenar]);

  // Funciones
  const agregar = () => { /* ... */ };
  const eliminar = () => { /* ... */ };
  const actualizar = () => { /* ... */ };

  // JSX
  return <div>...</div>;
}

// ✅ Con custom hooks - componente limpio
function ComponenteLimpio() {
  const { usuarios } = useUsuarios();
  const { filtro, setFiltro } = useFiltro();
  const { ordenar, setOrdenar } = useOrdenamiento();

  return <div>...</div>;
}
```

---

## ¿CÓMO funciona?

### 1. Hook Simple - useFetch

```typescript
import { useState, useEffect } from "react";

function useFetch<T>(url: string) {
  const [datos, setDatos] = useState<T | null>(null);
  const [cargando, setCargando] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let isMounted = true;

    const cargar = async () => {
      try {
        const response = await fetch(url);
        if (!response.ok) throw new Error("Error en fetch");
        const data = await response.json();
        if (isMounted) setDatos(data);
      } catch (err) {
        if (isMounted) setError(err instanceof Error ? err.message : "Error");
      } finally {
        if (isMounted) setCargando(false);
      }
    };

    cargar();

    return () => {
      isMounted = false;  // Cleanup
    };
  }, [url]);

  return { datos, cargando, error };
}

// Usar
function MisProductos() {
  const { datos: productos, cargando, error } = useFetch<Producto[]>("/api/productos");

  if (cargando) return <p>Cargando...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <ul>
      {productos?.map(p => (
        <li key={p.id}>{p.nombre}</li>
      ))}
    </ul>
  );
}
```

### 2. Hook con Lógica - useContador

```typescript
import { useState } from "react";

function useContador(inicial: number = 0, paso: number = 1) {
  const [valor, setValor] = useState(inicial);

  const incrementar = () => setValor(v => v + paso);
  const decrementar = () => setValor(v => v - paso);
  const reset = () => setValor(inicial);

  return { valor, incrementar, decrementar, reset };
}

// Usar
function Contador() {
  const { valor, incrementar, decrementar, reset } = useContador(0, 1);

  return (
    <div>
      <p>Valor: {valor}</p>
      <button onClick={incrementar}>+</button>
      <button onClick={decrementar}>-</button>
      <button onClick={reset}>Reset</button>
    </div>
  );
}

// Reutilizar
function PaginaProductos() {
  const pagina = useContador(1, 1);  // Página 1, paso 1

  return (
    <div>
      <p>Página {pagina.valor}</p>
      <button onClick={pagina.incrementar}>Siguiente</button>
      <button onClick={pagina.decrementar}>Anterior</button>
    </div>
  );
}
```

### 3. Hook con Formulario - useFormulario

```typescript
import { useState, useCallback } from "react";

interface Errores {
  [key: string]: string;
}

function useFormulario<T extends Record<string, any>>(
  valoresIniciales: T,
  onSubmit: (datos: T) => Promise<void> | void,
  validar?: (datos: T) => Errores
) {
  const [valores, setValores] = useState(valoresIniciales);
  const [errores, setErrores] = useState<Errores>({});
  const [enviando, setEnviando] = useState(false);

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value, type } = e.target;
    setValores(prev => ({
      ...prev,
      [name]: type === "checkbox" ? (e.target as HTMLInputElement).checked : value
    }));
  }, []);

  const handleSubmit = useCallback(async (e: React.FormEvent) => {
    e.preventDefault();

    if (validar) {
      const erroresNuevos = validar(valores);
      setErrores(erroresNuevos);
      if (Object.keys(erroresNuevos).length > 0) return;
    }

    setEnviando(true);
    try {
      await onSubmit(valores);
    } catch (error) {
      setErrores({ submit: error instanceof Error ? error.message : "Error" });
    } finally {
      setEnviando(false);
    }
  }, [valores, onSubmit, validar]);

  const reset = useCallback(() => {
    setValores(valoresIniciales);
    setErrores({});
  }, [valoresIniciales]);

  return { valores, errores, handleChange, handleSubmit, reset, enviando };
}

// Usar
interface DatosLogin {
  email: string;
  password: string;
}

function Login() {
  const formulario = useFormulario<DatosLogin>(
    { email: "", password: "" },
    async (datos) => {
      const response = await fetch("/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(datos)
      });
      if (!response.ok) throw new Error("Login fallido");
    },
    (datos) => {
      const errores: Errores = {};
      if (!datos.email.includes("@")) errores.email = "Email inválido";
      if (datos.password.length < 6) errores.password = "Mínimo 6 caracteres";
      return errores;
    }
  );

  return (
    <form onSubmit={formulario.handleSubmit}>
      <input
        name="email"
        value={formulario.valores.email}
        onChange={formulario.handleChange}
        placeholder="Email"
      />
      {formulario.errores.email && <p>{formulario.errores.email}</p>}

      <input
        name="password"
        type="password"
        value={formulario.valores.password}
        onChange={formulario.handleChange}
        placeholder="Contraseña"
      />
      {formulario.errores.password && <p>{formulario.errores.password}</p>}

      <button type="submit" disabled={formulario.enviando}>
        {formulario.enviando ? "Enviando..." : "Entrar"}
      </button>
    </form>
  );
}
```

### 4. Hook con LocalStorage - useLocalStorage

```typescript
function useLocalStorage<T>(clave: string, valorInicial: T) {
  const [valor, setValor] = useState<T>(() => {
    try {
      const item = window.localStorage.getItem(clave);
      return item ? JSON.parse(item) : valorInicial;
    } catch {
      return valorInicial;
    }
  });

  const guardar = (nuevoValor: T) => {
    try {
      setValor(nuevoValor);
      window.localStorage.setItem(clave, JSON.stringify(nuevoValor));
    } catch (error) {
      console.error("Error al guardar en localStorage:", error);
    }
  };

  return [valor, guardar] as const;
}

// Usar
function Tema() {
  const [tema, setTema] = useLocalStorage("tema", "light");

  return (
    <div style={{ background: tema === "light" ? "white" : "black" }}>
      <button onClick={() => setTema(tema === "light" ? "dark" : "light")}>
        {tema === "light" ? "🌙" : "☀️"}
      </button>
    </div>
  );
}

// Guarda en localStorage automáticamente
// Persiste entre recargas
```

### 5. Hook con Debounce - useFiltro

```typescript
import { useState, useEffect } from "react";

function useDebounce<T>(valor: T, retraso: number = 500) {
  const [valorDebounced, setValorDebounced] = useState(valor);

  useEffect(() => {
    const timer = setTimeout(() => {
      setValorDebounced(valor);
    }, retraso);

    return () => clearTimeout(timer);
  }, [valor, retraso]);

  return valorDebounced;
}

// Usar
function BuscadorProductos() {
  const [busqueda, setBusqueda] = useState("");
  const busquedaDebounced = useDebounce(busqueda, 300);  // Espera 300ms

  // Solo cambiar cuando termine de escribir
  useEffect(() => {
    console.log("Buscando:", busquedaDebounced);
    // Hacer fetch aquí
  }, [busquedaDebounced]);

  return (
    <input
      placeholder="Buscar..."
      value={busqueda}
      onChange={(e) => setBusqueda(e.target.value)}
    />
  );
}
```

### 6. Hook con Previous Value - usePrevious

```typescript
function usePrevious<T>(valor: T): T | undefined {
  const ref = useRef<T>();

  useEffect(() => {
    ref.current = valor;
  }, [valor]);

  return ref.current;
}

// Usar
function ComponenteConHistorial() {
  const [contador, setContador] = useState(0);
  const contadorAnterior = usePrevious(contador);

  return (
    <div>
      <p>Ahora: {contador}</p>
      <p>Anterior: {contadorAnterior}</p>
      <button onClick={() => setContador(c => c + 1)}>+</button>
    </div>
  );
}
```

---

## Reglas para Custom Hooks

```
✅ DEBE empezar con "use"
   const useAlgo = () => { ... }

✅ PUEDE llamar otros hooks
   useState, useEffect, useContext, etc.

✅ PUEDE tener lógica compleja
   Efectos, cálculos, validaciones

❌ NO llamar condicionalmente
   if (...) useEffect(() => { ... })  // ❌

❌ NO llamar en loops
   for (...) useContext()  // ❌

❌ NO desde funciones regulares
   const fn = () => {
     useEffect()  // ❌
   }
```

---

## Tabla de Custom Hooks Comunes

| Hook | Para | Retorna |
|------|------|---------|
| `useFetch` | Cargar datos de API | datos, cargando, error |
| `useDebounce` | Esperar a que termine de escribir | valorDebounced |
| `useLocalStorage` | Guardar en navegador | [valor, guardar] |
| `useFormulario` | Manejar formulario | valores, errores, handleChange... |
| `useContador` | Incrementar/decrementar | valor, incrementar, decrementar |
| `usePrevious` | Valor anterior | valorAnterior |
| `useAsync` | Operaciones async | data, loading, error |

---

## Resumen

**Custom Hooks son:**
- ✅ Reutilización de lógica
- ✅ Componentes más limpios
- ✅ Fácil de testear
- ✅ Compartir con el equipo
- ✅ DRY (Don't Repeat Yourself)

🚀 **Código limpio y reutilizable**