# Módulo 25 - Estructuras de Datos Avanzadas

JavaScript proporciona estructuras de datos especializadas para casos donde un Array o un Object básico no son suficientes. Las más importantes son `Set` y `Map`.

## 1. `Set` (Conjuntos)

Un `Set` es una colección de valores **únicos**. A diferencia de un Array, en un Set no puede haber elementos duplicados.

```javascript
const miSet = new Set();
miSet.add(1);
miSet.add(2);
miSet.add(2); // Ignorado, ya existe
miSet.add("Hola");

console.log(miSet); // Set(3) { 1, 2, "Hola" }
console.log(miSet.size); // 3 (Se usa 'size', no 'length')
console.log(miSet.has(2)); // true
```

**Truco: Eliminar duplicados de un array**
```javascript
const numerosConDuplicados = [1, 2, 2, 3, 4, 4, 5];
const arrayLimpio = [...new Set(numerosConDuplicados)];
console.log(arrayLimpio); // [1, 2, 3, 4, 5]
```

## 2. `Map` (Mapas / Diccionarios)

Un `Map` es una colección de pares clave-valor. Aunque se parece a un Objeto regular, tiene ventajas cruciales:
1. **Cualquier tipo de clave**: En un objeto, las claves solo pueden ser Strings o Symbols. En un Map, la clave puede ser un número, otro objeto, una función, etc.
2. **Mantiene el orden** de inserción.
3. **Iteración directa** (tiene métodos como `.forEach`, lo cual un objeto normal no tiene).

```javascript
const miMapa = new Map();

// set(clave, valor)
miMapa.set("nombre", "Carlos");
miMapa.set(100, "Cien"); // Clave numérica
miMapa.set({ id: 1 }, "Usuario"); // Clave objeto

console.log(miMapa.get("nombre")); // "Carlos"
console.log(miMapa.size); // 3

// Iteración
for (const [clave, valor] of miMapa) {
  console.log(`Clave: ${clave}, Valor: ${valor}`);
}
```

## 3. `WeakSet` y `WeakMap`

Son variantes "débiles" de Set y Map.
- Sus claves o valores **deben ser objetos**.
- No previenen que el *Garbage Collector* de JavaScript elimine los objetos si ya no hay otras referencias a ellos. Esto los hace ideales para gestionar cachés temporales de objetos del DOM sin causar "fugas de memoria" (Memory Leaks).
- No son iterables (no puedes usar `for...of` ni `.size`).

```javascript
let nodoDOM = document.getElementById("boton");
const clicks = new WeakMap();

// Almacenar info asociada a este elemento DOM
clicks.set(nodoDOM, 5);

// Si en otra parte eliminamos nodoDOM (ej. nodoDOM.remove()), 
// el registro en 'clicks' se borrará automáticamente de la memoria RAM.
```
