# Módulo 20 - React + TypeScript

> **Type Safety - Atrapa errores ANTES de producción**

---

## ¿QUÉ es TypeScript?

**TypeScript = JavaScript con "tipos"**

Es como la diferencia entre:
- JavaScript: Papel en blanco (escribes lo que quieras)
- TypeScript: Formulario con casillas específicas (solo ciertas cosas encajan)

```javascript
// ❌ JavaScript - permite CUALQUIER cosa
const suma = (a, b) => a + b;
suma(5, "texto");  // ¿Qué pasa? 🤷‍♂️

// ✅ TypeScript - especifica tipos
const suma = (a: number, b: number): number => a + b;
suma(5, "texto");  // ❌ ERROR en tiempo de escritura
```

---

## ¿PARA QUÉ?

### Problema 1: Errores invisibles

```javascript
// JavaScript
function obtenerNombre(usuario) {
  return usuario.nombre.toUpperCase();
}

obtenerNombre(null);  // ❌ CRASH en producción 💥
```

### Problema 2: Cambios peligrosos

```javascript
// Cambias una función
function procesarDatos(datos) {
  return datos.edad;
}

// Alguien usa así:
const año = procesarDatos({ fecha: "2024" });
// ❌ Ahora es undefined - nadie lo sabe
```

### Problema 3: Falta de documentación

```javascript
// ¿Qué recibe? ¿Qué devuelve?
function calcular(x, y, z) {
  // ???
}
```

**TypeScript resuelve TODO ESTO:**

```typescript
interface Usuario {
  nombre: string;
  edad: number;
}

function obtenerNombre(usuario: Usuario): string {
  return usuario.nombre.toUpperCase();
}

// ✅ Autocompletado
// ✅ Errores inmediatos
// ✅ Documentación automática
```

---

## ¿CÓMO funciona?

### 1. Tipos Básicos

```typescript
// Strings
const nombre: string = "Juan";
// nombre = 123;  ❌ Error

// Numbers
const edad: number = 25;
// edad = "25";  ❌ Error

// Booleans
const activo: boolean = true;
// activo = "true";  ❌ Error

// Any (último recurso - NO RECOMENDADO)
const cualquierCosa: any = "puede ser cualquier cosa";
qualquerCosa = 123;  // OK, pero... ❌
```

### 2. Objetos (Interfaces)

```typescript
// Interfaz define estructura
interface Usuario {
  nombre: string;
  edad: number;
  email?: string;  // Opcional (?)
  activo: boolean;
}

// ✅ Correcto
const usuario: Usuario = {
  nombre: "Juan",
  edad: 25,
  activo: true
  // email es opcional, OK
};

// ❌ Falta "edad"
const usuarioMalo: Usuario = {
  nombre: "Juan"
};

// ❌ "email" debe ser string
const usuarioMalo2: Usuario = {
  nombre: "Juan",
  edad: 25,
  email: 123,
  activo: true
};
```

### 3. Arrays

```typescript
// Array de strings
const nombres: string[] = ["Juan", "María"];
// nombres.push(123);  ❌ Error

// Array de números
const edades: number[] = [25, 30, 35];

// Array de objetos
interface Producto {
  id: number;
  nombre: string;
  precio: number;
}

const productos: Producto[] = [
  { id: 1, nombre: "Laptop", precio: 1000 }
];

// Array mixto (Union)
const mixto: (string | number)[] = [1, "dos", 3];
```

### 4. Funciones con TypeScript

```typescript
// Parámetros tipados + retorno
function duplicar(numero: number): number {
  return numero * 2;
}

duplicar(5);      // ✅ OK
// duplicar("5");  // ❌ Error

// Arrow function
const restar = (a: number, b: number): number => a - b;

// Con parámetros opcionales
function saludar(nombre: string, apellido?: string): string {
  return `Hola ${nombre} ${apellido || ""}`;
}

saludar("Juan");           // ✅ OK
saludar("Juan", "Pérez");  // ✅ OK
```

### 5. EN REACT: Componentes con TypeScript

**Props con interfaz:**

```typescript
// Define qué recibe el componente
interface PerfiltProps {
  nombre: string;
  edad: number;
  email?: string;
  onActualizar: (nuevoNombre: string) => void;
}

// Component
function Perfil({ nombre, edad, onActualizar }: PerfilProps) {
  return (
    <div>
      <h1>{nombre}</h1>
      <p>{edad} años</p>
      <button onClick={() => onActualizar("Nuevo")}>
        Actualizar
      </button>
    </div>
  );
}

// ✅ Autocompletado de props
<Perfil
  nombre="Juan"
  edad={25}
  onActualizar={(n) => console.log(n)}
/>

// ❌ Error
<Perfil
  nombre="Juan"
  edad="25"  // ❌ Debe ser number
/>
```

**useState con TypeScript:**

```typescript
import { useState } from "react";

// Sin especificar (TypeScript adivina)
const [nombre, setNombre] = useState("Juan");  // string

// Con especificación explícita
const [edad, setEdad] = useState<number>(25);

// Con interfaz
interface FormData {
  email: string;
  password: string;
}

const [formulario, setFormulario] = useState<FormData>({
  email: "",
  password: ""
});

// ✅ TypeScript sabe que estas propiedades existen
setFormulario({
  email: "juan@ejemplo.com",
  password: "secreta"
});
```

**useEffect con TypeScript:**

```typescript
useEffect(() => {
  // Código aquí
}, []);  // Dependencies

// Con async/await
useEffect(() => {
  const cargarDatos = async () => {
    try {
      const response = await fetch("/api/usuarios");
      const datos: Usuario[] = await response.json();
      setUsuarios(datos);
    } catch (error: unknown) {
      if (error instanceof Error) {
        console.error(error.message);
      }
    }
  };

  cargarDatos();
}, []);
```

### 6. Generic Types (Genéricos)

**Reutilizar componentes con diferentes tipos:**

```typescript
// Sin genéricos (repetitivo)
interface ListaStrings {
  items: string[];
}

interface ListaNumbers {
  items: number[];
}

// ✅ Con genéricos (flexible)
interface Lista<T> {
  items: T[];
  onSeleccionar: (item: T) => void;
}

// Usar
function ListaProductos() {
  interface Producto {
    id: number;
    nombre: string;
  }

  const productos: Producto[] = [
    { id: 1, nombre: "Laptop" }
  ];

  return (
    <Lista<Producto>
      items={productos}
      onSeleccionar={(prod) => console.log(prod.nombre)}
    />
  );
}

// API genérica
async function fetchDatos<T>(url: string): Promise<T> {
  const response = await fetch(url);
  return response.json();
}

// Usar
const usuarios = await fetchDatos<Usuario[]>("/api/usuarios");
```

---

## Tabla de Referencia

| Tipo | Ejemplo | Descripción |
|------|---------|-------------|
| `string` | `"Juan"` | Texto |
| `number` | `25` | Números |
| `boolean` | `true` | Verdadero/Falso |
| `any` | Cualquier cosa | Evita esto |
| `string[]` | `["a", "b"]` | Array de strings |
| `number \| string` | `25` o `"25"` | Múltiples tipos (Union) |
| `interface` | Define estructura | Objeto tipado |
| `T` (genérico) | Reutilizable | Para cualquier tipo |
| `?` | Propiedad opcional | Puede no existir |
| `readonly` | No se modifica | Protege datos |

---

## Configuración Inicial

### tsconfig.json

```json
{
  "compilerOptions": {
    "target": "ES2020",
    "jsx": "react-jsx",
    "module": "ESNext",
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true
  }
}
```

### En componentes

```typescript
// ✅ Buena estructura
interface ComponentProps {
  titulo: string;
  onClick: () => void;
}

function Componente({ titulo, onClick }: ComponentProps) {
  return <button onClick={onClick}>{titulo}</button>;
}

export default Componente;
```

---

---

## Setup: Crear tu Primer Proyecto

### Opción 1: Vite (RECOMENDADO - más rápido)

**Paso 1: Crear proyecto**

```bash
npm create vite@latest mi-app -- --template react-ts
cd mi-app
```

**Paso 2: Instalar dependencias**

```bash
npm install
```

**Paso 3: Iniciar desarrollo**

```bash
npm run dev
```

**Estructura automática:**

```
mi-app/
├── src/
│   ├── App.tsx          # Componente principal
│   ├── App.css
│   ├── main.tsx         # Entrada
│   └── vite-env.d.ts    # Tipos Vite
├── index.html
├── tsconfig.json        # Configuración TypeScript
├── vite.config.ts
└── package.json
```

### Opción 2: Create React App (más lento)

```bash
npx create-react-app mi-app --template typescript
cd mi-app
npm start
```

### Primer Componente con TypeScript

**App.tsx:**

```typescript
import { useState } from "react";
import "./App.css";

// Interfaz para props
interface GreetingProps {
  initial?: string;
}

// Componente tipado
function App({ initial = "React" }: GreetingProps) {
  const [nombre, setNombre] = useState(initial);
  const [contador, setContador] = useState(0);

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setContador(conteo => conteo + 1);
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setNombre(event.target.value);
  };

  return (
    <div className="app">
      <h1>Hola, {nombre}! 👋</h1>
      
      <input
        type="text"
        value={nombre}
        onChange={handleChange}
        placeholder="Tu nombre"
      />

      <button onClick={handleClick}>
        Clickeaste {contador} veces
      </button>

      <div className="info">
        <p>Estás usando React + TypeScript</p>
      </div>
    </div>
  );
}

export default App;
```

**main.tsx:**

```typescript
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

const root = document.getElementById("root");

if (!root) {
  throw new Error("Root element not found");
}

ReactDOM.createRoot(root).render(
  <React.StrictMode>
    <App initial="TypeScript" />
  </React.StrictMode>
);
```

**App.css:**

```css
.app {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  text-align: center;
  font-family: Arial, sans-serif;
}

input {
  padding: 10px;
  font-size: 16px;
  margin: 10px 0;
  width: 100%;
  max-width: 300px;
}

button {
  padding: 10px 20px;
  font-size: 16px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  margin: 10px 5px;
}

button:hover {
  background: #0056b3;
}

.info {
  margin-top: 20px;
  padding: 15px;
  background: #f0f0f0;
  border-radius: 5px;
}
```

### Ejecutar y Probar

```bash
npm run dev
# Abrirá http://localhost:5173
```

**¿Qué ves?**
- Título: "Hola, TypeScript! 👋"
- Input para cambiar nombre
- Botón que cuenta clicks
- Cambios en tiempo real (hot reload)

---

## Proyecto Pequeño: Contador de Tareas

**tipos.ts** (tipos compartidos):

```typescript
export interface Tarea {
  id: number;
  titulo: string;
  completada: boolean;
}

export interface TareasState {
  tareas: Tarea[];
  agregar: (titulo: string) => void;
  completar: (id: number) => void;
  eliminar: (id: number) => void;
}
```

**TareasContext.tsx:**

```typescript
import { createContext, useState, ReactNode } from "react";
import { Tarea, TareasState } from "./tipos";

export const TareasContext = createContext<TareasState | undefined>(undefined);

export function ProveedorTareas({ children }: { children: ReactNode }) {
  const [tareas, setTareas] = useState<Tarea[]>([]);
  const [proximoId, setProximoId] = useState(1);

  const agregar = (titulo: string) => {
    if (titulo.trim()) {
      setTareas([...tareas, { id: proximoId, titulo, completada: false }]);
      setProximoId(proximoId + 1);
    }
  };

  const completar = (id: number) => {
    setTareas(tareas.map(t => 
      t.id === id ? { ...t, completada: !t.completada } : t
    ));
  };

  const eliminar = (id: number) => {
    setTareas(tareas.filter(t => t.id !== id));
  };

  return (
    <TareasContext.Provider value={{ tareas, agregar, completar, eliminar }}>
      {children}
    </TareasContext.Provider>
  );
}
```

**ListaTareas.tsx:**

```typescript
import { useContext } from "react";
import { TareasContext } from "./TareasContext";

export function ListaTareas() {
  const context = useContext(TareasContext);
  
  if (!context) {
    throw new Error("ListaTareas debe estar dentro de ProveedorTareas");
  }

  const { tareas, completar, eliminar } = context;

  return (
    <ul>
      {tareas.map(tarea => (
        <li key={tarea.id} style={{
          textDecoration: tarea.completada ? "line-through" : "none"
        }}>
          <input
            type="checkbox"
            checked={tarea.completada}
            onChange={() => completar(tarea.id)}
          />
          {tarea.titulo}
          <button onClick={() => eliminar(tarea.id)}>Eliminar</button>
        </li>
      ))}
    </ul>
  );
}
```

**FormularioTarea.tsx:**

```typescript
import { useContext, useState } from "react";
import { TareasContext } from "./TareasContext";

export function FormularioTarea() {
  const [input, setInput] = useState("");
  const context = useContext(TareasContext);

  if (!context) {
    throw new Error("FormularioTarea debe estar dentro de ProveedorTareas");
  }

  const { agregar } = context;

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    agregar(input);
    setInput("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Nueva tarea..."
      />
      <button type="submit">Agregar</button>
    </form>
  );
}
```

**App.tsx:**

```typescript
import { ProveedorTareas } from "./TareasContext";
import { FormularioTarea } from "./FormularioTarea";
import { ListaTareas } from "./ListaTareas";
import "./App.css";

function App() {
  return (
    <ProveedorTareas>
      <div className="app">
        <h1>Mi Lista de Tareas 📝</h1>
        <FormularioTarea />
        <ListaTareas />
      </div>
    </ProveedorTareas>
  );
}

export default App;
```

**Resultado:**
- ✅ Sistema de tareas completo
- ✅ TypeScript tipado 100%
- ✅ Context para estado global
- ✅ Cambios automáticos
- ✅ Manejo de errores

---

## Errores Comunes

### ❌ Usar `any`

```typescript
// ❌ No hagas esto
const datos: any = fetch("/api");
```

### ✅ Especificar tipos

```typescript
// ✅ Hazlo así
const datos: Promise<Response> = fetch("/api");
```

### ❌ Olvidar tipos en Props

```typescript
// ❌ Sin tipos
function Componente(props) {
  return <div>{props.nombre}</div>;
}
```

### ✅ Con interfaz

```typescript
// ✅ Con tipos
interface ComponenteProps {
  nombre: string;
}

function Componente({ nombre }: ComponenteProps) {
  return <div>{nombre}</div>;
}
```

---

## Resumen

TypeScript es tu **guardaespaldas:**
- Atrapa errores ANTES de enviar a producción
- Autocompletado inteligente
- Documentación automática
- Código más mantenible
- Refactorización segura

**Costo:** Escribes un poco más
**Beneficio:** Evitas bugs costosos 💰

🚀 **TypeScript = Código profesional**