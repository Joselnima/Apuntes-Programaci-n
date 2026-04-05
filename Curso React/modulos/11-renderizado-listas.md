# Módulo 11 - Renderizado de Listas

> **Mostrar múltiples elementos - un patrón fundamental**

---

## ¿QÉ es Renderizado de Listas?

**Renderizar listas** = mostrar múltiples elementos del mismo tipo:

```jsx
// ❌ Manual (repetitivo)
<ul>
  <li>Manzana</li>
  <li>Pera</li>
  <li>Uva</li>
  <li>Naranja</li>
  <li>Plátano</li>
</ul>

// ✅ Con map (automático)
{frutas.map((fruta) => <li>{fruta}</li>)}
```

**Analogía:**
- 📋 Lista = fotocopiadora
- 🖨️ map = botón que copia automáticamente
- ✅ Resultado = múltiples copias

---

## ¿PARA QÉ sirve?

### 1. Mostrar Productos (e-commerce)
```jsx
function Tienda({ productos }) {
  return (
    <div>
      {productos.map((p) => (
        <TarjetaProducto 
          key={p.id}
          nombre={p.nombre}
          precio={p.precio}
        />
      ))}
    </div>
  );
}
```

### 2. Mostrar Usuarios
```jsx
function DirectorioUsuarios({ usuarios }) {
  return (
    <ul>
      {usuarios.map((usuario) => (
        <li key={usuario.id}>
          <h3>{usuario.nombre}</h3>
          <p>{usuario.email}</p>
        </li>
      ))}
    </ul>
  );
}
```

### 3. Mostrar Búsqueda Filtrada
```jsx
function Resultados({ items, filtro }) {
  const filtrados = items.filter(item => 
    item.nombre.includes(filtro)
  );
  
  return (
    <ul>
      {filtrados.map((item) => (
        <li key={item.id}>{item.nombre}</li>
      ))}
    </ul>
  );
}
```

### 4. Tabla de Datos
```jsx
function Tabla({ datos }) {
  return (
    <table>
      <tbody>
        {datos.map((fila) => (
          <tr key={fila.id}>
            <td>{fila.nombre}</td>
            <td>{fila.email}</td>
            <td>{fila.edad}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
```

---

## ¿CÓMO funcionan?

### Sintaxis map()

```jsx
array.map((elemento) => <li>{elemento}</li>)
     ↑    ↑        ↑    ↑
     Método parámetro devuelve JSX
```

**Paso a paso:**
```jsx
const frutas = ["Manzana", "Pera", "Uva"];

frutas.map((fruta) => {
  // Iteración 1: fruta = "Manzana"
  //   → retorna <li>Manzana</li>
  // Iteración 2: fruta = "Pera"
  //   → retorna <li>Pera</li>
  // Iteración 3: fruta = "Uva"
  //   → retorna <li>Uva</li>
  
  return <li>{fruta}</li>;
});
```

### El Atributo `key` (IMPORTANTE)

**¿Qué es key?**
```jsx
{items.map((item) => (
  <li key={item.id}>  // ← key debe ser ÚNICO
    {item.nombre}
  </li>
))}
```

**¿Por qué es importante?**
```jsx
// ❌ MALO (usa index como key)
{items.map((item, index) => (
  <li key={index}>{item}</li>
  // Si eliminas primer item, keys se mezclan
))}

// ✅ CORRECTO (usa ID único)
{items.map((item) => (
  <li key={item.id}>{item.nombre}</li>
  // Id nunca cambia
))}
```

### Casos de Uso de Keys

| Caso | Usa | Ejemplo |
|------|-----|---------|
| **BD/API** | ID único | `key={user.id}` |
| **Items constantes** | index OK | `key={index}` |
| **Sin ID** | UUID | `key={uuid()}` |
| **Listas Dinámicas** | ID único | `key={item.id}` |

### Patrones Comunes

```jsx
// Patrón 1: Objeto simple
{usuarios.map((u) => <p key={u.id}>{u.nombre}</p>)}

// Patrón 2: Componente
{products.map((p) => <Producto key={p.id} {...p} />)}

// Patrón 3: Con filtro
{usuarios
  .filter(u => u.activo)
  .map(u => <Usuario key={u.id} usuario={u} />)
}

// Patrón 4: Con index (solo si es seguro)
{items.map((item, idx) => (
  <span key={idx}>{item}</span>
))}
```

### Ciclo Completo

```
Datos (array)
     ↓
.map() crea JSX
     ↓
React renderiza cada elemento
     ↓
Usa key para identificar
     ↓
Pantalla muestra lista
     ↓
Si datos cambian: React actualiza solo lo necesario
```

## Ejercicio
Renderiza una lista de:
- clientes
- productos
- tareas
