# Módulo 02 - Preparando tu Entorno

> **Instalar herramientas para empezar a programar en React**

---

## ¿QÉ necesito?

Para programar en React necesitas **3 cosas**:

```
┌─────────────────────────────────────┐
│ 1. Node.js (engine de JavaScript)   │
├─────────────────────────────────────┤
│ 2. Editor (VS Code)                 │
├─────────────────────────────────────┤
│ 3. Navegador (Chrome recomendado)   │
└─────────────────────────────────────┘
```

**Analogía:**
- 🔨 Node.js = herramientas para construir
- 📝 VS Code = lugar para escribir
- 🌐 Navegador = lugar donde se ve

---

## ¿PARA QÉ?

### 1. Node.js
Te permite ejecutar JavaScript fuera del navegador:
```bash
node -v      # Ver versión
node app.js  # Ejecutar archivo
npm install  # Instalar librerías
```

### 2. VS Code
Editor para escribir código:
- Resaltado de sintaxis (colores)
- Autocompletado
- Terminal integrada
- Extensiones

### 3. Navegador
Ver la aplicación funcionando

---

## ¿CÓMO?

### Paso 1: Instalar Node.js

**En Windows:**
1. Ve a [nodejs.org](https://nodejs.org)
2. Descarga LTS (recomendado)
3. Ejecuta instalador
4. Próximo, próximo, siguiente...

**En Mac:**
```bash
brew install node
```

**En Linux:**
```bash
sudo apt update
sudo apt install nodejs npm
```

**Verificar:**
```bash
node -v    # v18.x.x
npm -v     # 9.x.x
```

### Paso 2: Instalar VS Code

1. Ve a [code.visualstudio.com](https://code.visualstudio.com)
2. Descarga
3. Instala

**Extensiones recomendadas:**
- ES7+ React/Redux/React-Native snippets
- Prettier
- ES Lint

### Paso 3: Crear Proyecto React

```bash
# Crear proyecto
npm create vite@latest mi-app -- --template react

# Entrar a carpeta
cd mi-app

# Instalar dependencias
npm install

# Ejecutar localmente
npm run dev
```

**Salida:**
```
VITE v4.x.x ready in xxx ms

➜  Local:   http://localhost:5173/
```

Abre en navegador: `http://localhost:5173`

### Paso 4: Estructura del Proyecto

```
mi-app/
├── src/
│   ├── App.jsx           ← Componente principal
│   ├── App.css
│   ├── main.jsx          ← Entrada
│   └── index.css
├── public/
│   └── logo.png
├── package.json          ← Dependencias
├── vite.config.js        ← Configuración
└── index.html
```

**Archivos importantes:**
- `src/App.jsx` = tu componente principal
- `package.json` = librerías instaladas
- `index.html` = punto de entrada HTML

### Paso 5: Editar y Ver Cambios

```jsx
// src/App.jsx
function App() {
  return <h1>¡Hola React!</h1>;
}

export default App;
```

**Guarda y recarga navegador → ¡Cambios automáticos!**

---

## Comandos Útiles

| Comando | Qué hace |
|---------|----------|
| `npm run dev` | Inicia servidor local |
| `npm install` | Instala dependencias |
| `npm run build` | Crear versión para producción |
| `npm install nombre` | Agregar nueva librería |
| `npm uninstall nombre` | Remover librería |

---

## ¡Primer Componente!

```jsx
// src/App.jsx
import './App.css'

function App() {
  const nombre = "Juan";
  
  return (
    <div>
      <h1>¡Hola {nombre}!</h1>
      <p>Bienvenido a React</p>
    </div>
  );
}

export default App;
```

Guarda → ¡Debería verse cambios en navegador! 🎉

## Ejercicio
1. Instala Node.js
2. Crea tu primer proyecto
3. Ejecuta `npm run dev`
4. Abre la app en el navegador
