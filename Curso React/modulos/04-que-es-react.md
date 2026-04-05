# Módulo 04 - ¿Qué es React?

> **La librería para construir interfaces modernas y reactivas que responden al usuario**

---

## ¿QÉ es React?

**React** = librería que **reacciona** a cambios:

```
Usuario hace algo
     ↓
React detecta cambio
     ↓
React actualiza pantalla (automático)
```

**No tienes que:**
```javascript
// ❌ Buscar elementos
document.getElementById('nombre').textContent = 'Juan';
// Boring!
```

**Con React:**
```javascript
// ✅ Describe qué quieres
setNombre('Juan');
// React lo hace automáticamente
```

**Analogía:**
- 🤖 React = criada inteligente
- Tú = jefe que describe qué quieres
- Criada = lo hace automáticamente

---

## ¿PARA QÉ sirve?

### 1. Reutilizar componentes
```jsx
// Define UNA VEZ
function Tarjeta({ título, descripción }) {
  return (
    <div className="tarjeta">
      <h2>{título}</h2>
      <p>{descripción}</p>
    </div>
  );
}

// Usa 100 VECES
<Tarjeta título="Producto 1" descripción="..." />
<Tarjeta título="Producto 2" descripción="..." />
<Tarjeta título="Producto 3" descripción="..." />
```

### 2. Mantener Sincronizado
```jsx
// Sin React: debe cambiar lista Y contador
user.push(nuevoUser);  // Actualiza lista
contador.textContent = user.length;  // Actualiza contador

// Con React: todo sincronizado automáticamente
setUsers([...users, nuevoUser]);  // React actualiza TODO
```

### 3. Manejar Interacción
```jsx
// Click, cambios en input, validaciones
// Todo declarativo, claro, sin magia
<button onClick={comprar}>Comprar</button>
```

### 4. Aplicaciones Modernas
```jsx
// Gmail, Netflix, Spotify, Discord...
// Todas usan React para ser rápidas y reactivas
```

---

## ¿CÓMO funciona?

### Estructura Principal de React

```
┌─────────────────────────────────┐
│         App (raíz)              │
│  • Coordina todo                │
├─────────────────────────────────┤
│      ┌──────┐  ┌──────┐         │
│      │Header│  │Navbar│         │
│      └──────┘  └──────┘         │
│                                 │
│  ┌──────────────────────────┐   │
│  │      Main Content        │   │
│  │  • Listas                │   │
│  │  • Formularios           │   │
│  │  • Cards                 │   │
│  └──────────────────────────┘   │
│     ┌─────┐  ┌──────────────┐   │
│     │Card │  │  Comentarios │   │
│     └─────┘  └──────────────┘   │
└─────────────────────────────────┘
```

### 3 Máquinas Mentales

| Concepto | Mental | Realidad |
|----------|--------|----------|
| **Componente** | Pieza de LEGO | Función que devuelve HTML |
| **Estado** | Memoria de pieza | Variable que React observa |
| **Props** | Instrucciones | Parámetros a componente |

**Ciclo Life de React:**
```
1. Componente nace (se crea)
2. Recibe props (instrucciones)
3. Renderiza (muestra en pantalla)
4. Usuario interactúa
5. Estado cambia
6. Componente se actualiza
7. Vuelva a renderizar
```

## Ejercicio
Piensa en una página web y sepárala en componentes.
