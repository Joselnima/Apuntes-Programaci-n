# Módulo 01 - Antes de React

> **Entender cómo funciona la web clásica y qué problema resuelve React**

---

## ¿QÉ es una Página Web?

Toda página web está hecha de **3 capas**:

```
┌─────────────────────────────────┐
│  HTML (estructura)              │  Qué elementos hay
│  • Títulos, párrafos, botones   │
├─────────────────────────────────┤
│  CSS (estilos)                  │  Cómo se ve
│  • Colores, tamaños, posición   │
├─────────────────────────────────┤
│  JavaScript (comportamiento)    │  Qué pasa al hacer clic
│  • Cambios, validaciones, etc   │
└─────────────────────────────────┘
```

**Analogía:**
- 🏗️ HTML = planos de una casa (habitaciones, espacios)
- 🎨 CSS = pintura y decoración
- ⚙️ JavaScript = electricidad (funcionamiento)

---

## ¿PARA QÉ entender esto?

### Sin React (forma antigua - manual)
```html
<!-- Paso 1: Escribir HTML -->
<div id="contador">0</div>
<button id="boton">Sumar</button>

<!-- Paso 2: Escribir JavaScript para cada cosa -->
<script>
  let contador = 0;
  document.getElementById("boton").addEventListener("click", function() {
    contador = contador + 1;
    document.getElementById("contador").textContent = contador;
  });
</script>
```

**Problemas:**
- ❌ Código esparcido por todos lados
- ❌ Difícil reutilizar
- ❌ Fácil confundirse
- ❌ Crece desordenado

### Con React (forma moderna - componentes)
```jsx
function Contador() {
  const [contador, setContador] = useState(0);
  
  return (
    <div>
      <div>{contador}</div>
      <button onClick={() => setContador(contador + 1)}>
        Sumar
      </button>
    </div>
  );
}
```

**Ventajas:**
- ✅ Todo en un solo lugar
- ✅ Reutilizable
- ✅ Claro y organizado
- ✅ Escalable

---

## ¿CÓMO funciona?

### 4 Capas en React

| Capa | Qué es | Ejemplo |
|------|--------|---------|
| **Componente** | Pieza reutilizable | `<Botón />` |
| **JSX** | HTML en JavaScript | `<h1>Hola</h1>` |
| **Estado** | Datos que cambian | contador, nombre |
| **Renderizado** | Mostrar en pantalla | Automático cuando cambia estado |

**Ciclo de React:**
```
1. Usuario hace clic
     ↓
2. Estado cambia (setContador)
     ↓
3. React re-renderiza automáticamente
     ↓
4. Pantalla se actualiza
```

React **automatiza** lo que en JavaScript vanilla es manual y tedioso.

## Ejercicio
Responde:
1. ¿Qué hace HTML?
2. ¿Qué hace CSS?
3. ¿Qué hace JavaScript?
4. ¿Por qué una app grande se vuelve difícil sin componentes?
