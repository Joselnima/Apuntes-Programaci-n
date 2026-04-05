# Módulo 03 - JavaScript Necesario para React

> **JavaScript essentials - solo lo que necesitas saber para React**

---

## ¿QÉ JavaScript necesito?

**BUENA NOTICIA:** No necesitas TODO JavaScript, solo ~10% para empezar React

```
JavaScript completo = 🎪 Circo completo
React necesita = 🎯 Solo los elementos esenciales
```

---

## ¿PARA QÉ?

Estos conceptos aparecen en **CADA aplicación React**:
- Funciones modernas
- Objetos y arrays
- map(), filter(), find()
- Spread operator
- Desestructuración

---

## ¿CÓMO?

### 1. Variables: const y let

```javascript
// ✅ Siempre usa const (no cambia)
const nombre = "Juan";
// nombre = "Pedro";  ❌ Error

// ✅ USA let solo si va a cambiar
let edad = 25;
age = 26;  // OK

// ❌ NUNCA var (antiguo)
var color = "rojo";
```

**Regla:** Const por defecto, let si necesitas cambiar

### 2. Funciones Modernas

**Forma Antigua:**
```javascript
function saludar(nombre) {
  return "Hola " + nombre;
}
```

**Arrow Functions (modern):**
```javascript
const saludar = (nombre) => "Hola " + nombre;

// Si es complicado
const saludar = (nombre) => {
  const mensaje = `Hola ${nombre}`;
  return mensaje;
};
```

**En React:**
```jsx
// Arrow functions es lo que usas
const App = () => {
  return <h1>Hola</h1>;
};
```

### 3. Objetos

```javascript
// Crear
const usuario = {
  nombre: "Juan",
  edad: 25,
  email: "juan@ejemplo.com"
};

// Acceder
console.log(usuario.nombre);  // "Juan"
console.log(usuario["edad"]);  // 25

// Modificar
usuario.edad = 26;
usuario.ciudad = "Santiago";

// En React
function Perfil({ usuario }) {
  return <h1>{usuario.nombre}</h1>;
}
```

### 4. Arrays (Listas)

```javascript
const frutas = ["manzana", "pera", "uva"];

// Acceder
frutas[0]  // "manzana"

// Agregar
frutas.push("naranja");

// Largo
frutas.length  // 4

// En React (MÁS COMÚN)
function Lista({ items }) {
  return (
    <ul>
      {items.map((item) => (
        <li key={item.id}>{item.nombre}</li>
      ))}
    </ul>
  );
}
```

### 5. map() - LO MÁS IMPORTANTE

**map()** transforma cada elemento:

```javascript
const numeros = [1, 2, 3, 4, 5];

// Duplicar cada número
const duplicado = numeros.map((n) => n * 2);
// [2, 4, 6, 8, 10]

// Crear strings
const mensajes = numeros.map((n) => `Número: ${n}`);
// ["Número: 1", "Número: 2", ...]

// EN REACT (¡MUY COMÚN!)
const usuarios = [
  { id: 1, nombre: "Juan" },
  { id: 2, nombre: "María" }
];

function ListaUsuarios() {
  return (
    <ul>
      {usuarios.map((usuario) => (
        <li key={usuario.id}>{usuario.nombre}</li>
      ))}
    </ul>
  );
}
```

### 6. filter() y find()

**filter()** - mantiene elementos que cumplen condición

```javascript
const usuarios = [
  { id: 1, nombre: "Juan", activo: true },
  { id: 2, nombre: "María", activo: false },
  { id: 3, nombre: "Pedro", activo: true }
];

// Solo activos
const activos = usuarios.filter((u) => u.activo);
// [{id: 1, ...}, {id: 3, ...}]

// EN REACT
function UsuariosActivos() {
  const activos = usuarios.filter((u) => u.activo);
  return (
    <ul>
      {activos.map((u) => <li key={u.id}>{u.nombre}</li>)}
    </ul>
  );
}
```

**find()** - obtiene PRIMER elemento que cumple

```javascript
const usuario = usuarios.find((u) => u.id === 2);
// { id: 2, nombre: "María", activo: false }

// EN REACT
function UsuarioDetalle({ id }) {
  const usuario = usuarios.find((u) => u.id === id);
  return <h1>{usuario.nombre}</h1>;
}
```

### 7. Desestructuración

**Extraer valores sin repetir:**

```javascript
// ❌ Sin desestructuración (repetitivo)
const usuario = { nombre: "Juan", edad: 25 };
const nombre = usuario.nombre;
const edad = usuario.edad;

// ✅ Con desestructuración (limpio)
const { nombre, edad } = usuario;

// EN REACT (MÁS COMÚN)
function Perfil({ usuario }) {
  const { nombre, edad, email } = usuario;
  return (
    <div>
      <h1>{nombre}</h1>
      <p>{edad} años</p>
      <p>{email}</p>
    </div>
  );
}

// Arrays también
const frutas = ["manzana", "pera", "uva"];
const [primera, segunda] = frutas;
// primera = "manzana", segunda = "pera"
```

### 8. Spread Operator (...)

**Copiar y extender:**

```javascript
// Copiar objeto
const usuario = { nombre: "Juan", edad: 25 };
const usuarioNuevo = { ...usuario, ciudad: "Santiago" };
// { nombre: "Juan", edad: 25, ciudad: "Santiago" }

// Copiar array
const frutas = ["manzana", "pera"];
const nuevaLista = [...frutas, "uva"];
// ["manzana", "pera", "uva"]

// EN REACT (IMPORTANTE)
function ActualizarDatos() {
  const [usuario, setUsuario] = useState({ nombre: "Juan" });
  
  const actualizarNombre = (nuevoNombre) => {
    setUsuario({ ...usuario, nombre: nuevoNombre });
  };
}
```

### 9. Template Strings (backticks)

```javascript
// ❌ Viejo (concatenación)
const nombre = "Juan";
const mensaje = "Hola " + nombre + ", bienvenido";

// ✅ Moderno (template strings)
const mensaje = `Hola ${nombre}, bienvenido`;

// EN REACT
function Saludo({ nombre }) {
  return <h1>{`Hola ${nombre}`}</h1>;
}
```

### 10. Ternario (if compacto)

```javascript
// ❌ if-else normal
if (edad >= 18) {
  console.log("Mayor");
} else {
  console.log("Menor");
}

// ✅ Ternario (compact)
const estado = edad >= 18 ? "Mayor" : "Menor";

// EN REACT (COMÚN)
function Acceso({ edad }) {
  return (
    <h1>
      {edad >= 18 ? "Puedes entrar" : "Muy joven"}
    </h1>
  );
}

// Incluso más compacto con &&
{edad >= 18 && <button>Votar</button>}
```

---

## Tabla de Referencia

| Concepto | Uso | Ejemplo |
|----------|-----|---------|
| **const** | Variables | `const nombre = "Juan"` |
| **Arrow Fn** | Funciones | `const fn = () => {}` |
| **map()** | Transformar arrays | `arr.map(x => x * 2)` |
| **filter()** | Filtrar arrays | `arr.filter(x => x > 5)` |
| **find()** | Buscar uno | `arr.find(x => x.id === 1)` |
| **{}** | Desestructuración | `const { nombre } = obj` |
| **...** | Spread | `{ ...obj, city: "NY" }` |
| **\${...}** | Template strings | `"Hola ${nombre}"` |
| **? :** | Ternario | `x ? "sí" : "no"` |
| **&&** | AND compacto | `x && <Comp />` |

---

## ¡Listo!

Con estos 10 conceptos estás listo para React 🚀

Ahora ir a Módulo 04 para aprender React real.

## Ejercicio
Crea:
- un array de productos
- un objeto usuario
- una función que retorne un saludo
