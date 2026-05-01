# Módulo 08 - Arrays

Los arrays (o arreglos) son estructuras de datos que almacenan listas de valores. Pueden contener números, texto, objetos u otros arrays. En JavaScript, el objeto `Array` tiene decenas de métodos muy poderosos.

## 1. Crear un Array

```javascript
// Literal (Recomendado)
const numeros = [1, 2, 3, 4, 5];
const frutas = ['manzana', 'pera', 'plátano'];

// Constructor global
const vacio = new Array(); // []
const conTamaño = new Array(3); // [empty x 3]
```

## 2. Acceder y Modificar (Índices Básicos)

```javascript
const letras = ['a', 'b', 'c'];

// Acceso clásico
console.log(letras[0]); // 'a'

// at() (ES13) - Soporta negativos para leer desde el final
console.log(letras.at(-1)); // 'c'

// Sobrescribir
letras[1] = 'x'; 
console.log(letras); // ['a', 'x', 'c']
```

## 3. Agregar y Eliminar elementos

Estos métodos mutan el array original:

```javascript
const animales = ['perro'];

// Al final del array
animales.push('gato'); // Retorna nueva longitud (2) -> ['perro', 'gato']
const ultimo = animales.pop(); // Retorna el elemento quitado ('gato') -> ['perro']

// Al principio del array
animales.unshift('conejo'); // Desplaza el resto -> ['conejo', 'perro']
const primero = animales.shift(); // Quita el primero -> ['perro']
```

## 4. Manipulación de la Estructura (Mutables)

```javascript
const arr = [1, 2, 3, 4, 5];

// splice(indice, borrarCantidad, itemsAAgregar...)
arr.splice(2, 1, 'x', 'y'); // Borra 1 item en el índice 2 y añade 'x', 'y'
console.log(arr); // [1, 2, 'x', 'y', 4, 5]

// sort() - Ordena. Por defecto lo hace alfabéticamente (convirtiendo a texto)
const desordenados = [10, 2, 5];
desordenados.sort((a, b) => a - b); // Ordenamiento numérico [2, 5, 10]

// reverse() - Invierte el orden in-situ
desordenados.reverse(); // [10, 5, 2]

// fill(valor, inicio, fin) - Rellena con un valor fijo
const ceros = [1, 2, 3, 4].fill(0, 1, 3); // [1, 0, 0, 4]

// copyWithin(destino, inicio, fin) - Copia secuencias internamente
const letrasCop = ['a', 'b', 'c', 'd', 'e'];
letrasCop.copyWithin(0, 3, 5); // ['d', 'e', 'c', 'd', 'e']
```

## 5. Manipulación de la Estructura (Inmutables - Devuelven copia)

```javascript
// slice(inicio, fin) - Extrae una porción
const porcion = [1, 2, 3, 4].slice(1, 3); // [2, 3]

// concat() - Fusiona arrays
const mezcla = [1, 2].concat([3, 4]); // [1, 2, 3, 4]

// flat(profundidad) - Aplana arrays anidados
const anidado = [1, [2, 3], [[4]]];
console.log(anidado.flat(2)); // [1, 2, 3, 4]

// Métodos inmutables ES14 (2023)
const original = [3, 1, 2];
console.log(original.toSorted());    // [1, 2, 3]
console.log(original.toReversed());  // [2, 1, 3]
console.log(original.toSpliced(1,1));// [3, 2]
console.log(original.with(1, 9));    // [3, 9, 2]
console.log(original); // Sigue siendo [3, 1, 2]
```

## 6. Métodos de Búsqueda y Verificación

```javascript
const num = [10, 20, 30, 40, 20];

// includes() - ¿Existe el valor?
console.log(num.includes(20)); // true

// indexOf() / lastIndexOf() - Posición exacta
console.log(num.indexOf(20)); // 1
console.log(num.lastIndexOf(20)); // 4

// find() / findIndex() - Busca usando una función de condición
console.log(num.find(n => n > 25)); // 30 (el primero que encuentra)
console.log(num.findIndex(n => n > 25)); // 2 (su índice)

// findLast() / findLastIndex() (ES14) - Busca desde el final
console.log(num.findLast(n => n > 25)); // 40
console.log(num.findLastIndex(n => n > 25)); // 3

// some() - ¿Al menos UNO cumple la condición?
console.log(num.some(n => n > 35)); // true

// every() - ¿TODOS cumplen la condición?
console.log(num.every(n => n > 0)); // true
```

## 7. Métodos Funcionales (Iteración Avanzada)

Estos métodos reciben un "Callback" (una función) que se ejecuta por cada elemento.

```javascript
const base = [1, 2, 3];

// forEach() - Ejecuta código para cada elemento (No retorna nada)
base.forEach((valor, indice) => console.log(`${indice}: ${valor}`));

// map() - Transforma cada elemento y devuelve un NUEVO array
const dobles = base.map(n => n * 2); // [2, 4, 6]

// filter() - Devuelve NUEVO array con los elementos que pasen el test
const pares = [1, 2, 3, 4].filter(n => n % 2 === 0); // [2, 4]

// reduce() - Acumula los valores en un único resultado
// reduce(callback(acumulador, valorActual), valorInicial)
const suma = base.reduce((acc, val) => acc + val, 0); // 6

// reduceRight() - Igual, pero de derecha a izquierda
const restaInversa = [10, 5, 2].reduceRight((acc, val) => acc - val); // 2 - 5 - 10 = -13

// flatMap() - Mapea y luego aplana profundidad 1
const frases = ["Hola Mundo"];
console.log(frases.flatMap(f => f.split(" "))); // ["Hola", "Mundo"]
```

## 8. Conversión e Iteradores de Objeto

```javascript
const chars = ['A', 'B', 'C'];

// join() - Une a string usando un separador
console.log(chars.join('-')); // "A-B-C"

// toString() / toLocaleString()
console.log(chars.toString()); // "A,B,C"

// Iteradores explícitos
const iteradorKeys = chars.keys(); // Devuelve iterador de índices [0, 1, 2]
const iteradorValues = chars.values(); // Devuelve iterador de valores ['A', 'B', 'C']
const iteradorEntries = chars.entries(); // Devuelve iterador de pares [0, 'A'], [1, 'B']

for (const [indice, valor] of chars.entries()) {
  console.log(indice, valor);
}
```

## 9. Métodos Estáticos de la Clase Array

Métodos que se llaman directamente desde `Array`, no desde una instancia.

```javascript
// Array.from() - Convierte estructuras similares a array (NodeLists, Strings, Sets) a Arrays reales.
// También acepta un mapFn opcional.
const arrayDeString = Array.from("HOLA"); // ['H', 'O', 'L', 'A']
const mapeadoInmediato = Array.from([1, 2, 3], x => x * 10); // [10, 20, 30]

// Array.isArray() - Comprobación de seguridad
console.log(Array.isArray([1, 2])); // true
console.log(Array.isArray({})); // false

// Array.of() - Crea un array a partir de argumentos sueltos (Evita el bug de new Array(numero))
console.log(Array.of(7)); // [7] (new Array(7) crearía un array de 7 vacíos)
```
