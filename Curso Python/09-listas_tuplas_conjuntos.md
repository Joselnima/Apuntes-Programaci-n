# Listas, Tuplas y Conjuntos

> **Cómo guardar múltiples datos juntos - Colecciones de información**

---

## ¿QÉ son Listas, Tuplas y Conjuntos?

Son **contenedores** que guardan múltiples valores en uno solo:

### Lista: Tu lista de compras
```python
compra = ["manzanas", "pan", "leche"]
# La puedes editar: agregar, quitar, cambiar
```

### Tupla: Coordenadas de ubicación
```python
ubicacion = (40.4168, -3.7038)  # Latitud, Longitud
# NO puedes cambiar (fija)
```

### Set: Amigos únicos (sin repetidos)
```python
amigos = {"Ana", "Bruno", "Carlos"}
# Sin duplicados, operaciones matemáticas
```

---

## ¿PARA QÉ sirven?

### Listas: Datos que CAMBIAN
```python
tareas = ["hacer cama", "lavar platos"]
tareas.append("comprar leche")    # Agregar
tareas.remove("hacer cama")       # Quitar
print(len(tareas))                # Saber cuántas
```

### Tuplas: Datos que NO cambian (Seguros)
```python
coordenadas = (40.4, -3.7)  # No puedo modificar
# Garantizado que siempre vale lo mismo
# Más seguro y rápido
```

### Conjuntos: Valores ÚNICOS
```python
etiquetas = {"rojo", "azul", "verde"}
etiquetas.add("amarillo")     # Agregar
etiquetas.remove("rojo")      # Quitar

# Operaciones matemáticas
conjunto1 | conjunto2         # Unión
conjunto1 & conjunto2         # Intersección
```

---

## Parte I: Listas (List) - La Estructura Más Usada

### Creación

```python
# Vacía
lista = []

# Con elementos
números = [1, 2, 3, 4, 5]

# Tipos mixtos (legal pero no recomendado)
mixta = [1, "texto", 3.14, True, None]

# Anidada
matriz = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]

# Con range
números = list(range(10))  # [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
```

### Acceso

```python
lista = [10, 20, 30, 40, 50]

print(lista[0])       # 10 (primer elemento)
print(lista[-1])      # 50 (último)
print(lista[1:3])     # [20, 30] (slice)
print(len(lista))     # 5
```

### Modificación

```python
lista = [1, 2, 3, 4, 5]

# Cambiar elemento
lista[0] = 10         # [10, 2, 3, 4, 5]

# Cambiar rango
lista[1:3] = [20, 30]  # [10, 20, 30, 4, 5]
```

### Métodos Importantes

```python
lista = [3, 1, 4, 1, 5]

# .append() - agregar al final
lista.append(9)           # [3, 1, 4, 1, 5, 9]

# .extend() - agregar varios
lista.extend([2, 6])      # [3, 1, 4, 1, 5, 9, 2, 6]

# .insert() - insertar en posición
lista.insert(0, 0)        # [0, 3, 1, 4, 1, 5, 9, 2, 6]

# .remove() - elimina primer valor
lista.remove(1)           # Quita el primer 1

# .pop() - elimina y devuelve
último = lista.pop()      # Elimina último
primero = lista.pop(0)    # Elimina primero

# .clear() - vacía
# lista.clear()           # []

# .sort() - ordena
lista = [3, 1, 4, 1, 5]
lista.sort()              # [1, 1, 3, 4, 5]
lista.sort(reverse=True)  # [5, 4, 3, 1, 1]

# .reverse() - invierte
lista = [1, 2, 3]
lista.reverse()           # [3, 2, 1]

# .index() - posición de elemento
lista = [1, 2, 3, 2]
print(lista.index(2))     # 1 (primer 2)

# .count() - cuántas veces aparece
print(lista.count(2))     # 2
```

### Búsqueda

```python
lista = [1, 2, 3, 4, 5]

# in
if 3 in lista:
    print("Encontrado")

# index
posición = lista.index(3)  # 2

# Condicional
pares = [x for x in lista if x % 2 == 0]  # [2, 4]
```

---

## Parte II: Tuplas (Tuple)

Secuencias **inmutables** (no se pueden cambiar):

### Creación

```python
# Con paréntesis (recomendado)
tupla = (1, 2, 3)

# Sin paréntesis (válido)
tupla = 1, 2, 3

# Vacía
tupla_vacía = ()

# Un elemento (NOTA la coma)
singleton = (1,)

# Tipos mixtos
tupla = (1, "texto", 3.14, True)

# Anidada
tupla = (1, (2, 3), (4, 5))
```

### Acceso

```python
tupla = (10, 20, 30, 40, 50)

print(tupla[0])       # 10
print(tupla[-1])      # 50
print(tupla[1:3])     # (20, 30)
print(len(tupla))     # 5
```

### ❌ No Mutable

```python
tupla = (1, 2, 3)

# Esto FALLA
# tupla[0] = 10        # TypeError

# Aber puedes convertir a lista
lista = list(tupla)
lista[0] = 10
tupla = tuple(lista)   # (10, 2, 3)
```

### Métodos

```python
tupla = (1, 2, 3, 2)

# .index() - posición
print(tupla.index(2))   # 1

# .count() - cuántas
print(tupla.count(2))   # 2
```

### Unpacking

Desempacar en variables:

```python
# Básico
dimensiones = (1920, 1080)
ancho, alto = dimensiones
print(ancho, alto)  # 1920 1080

# Múltiple
a, b, c = (1, 2, 3)
print(a, b, c)      # 1 2 3

# Con *
primero, *resto = (1, 2, 3, 4, 5)
print(primero)      # 1
print(resto)        # [2, 3, 4, 5]

# Intercambiar
a, b = 1, 2
a, b = b, a         # Swap perfecto
print(a, b)         # 2 1
```

### Named Tuples

Tuplas con nombres de campos:

```python
from collections import namedtuple

Persona = namedtuple("Persona", ["nombre", "edad", "ciudad"])

p = Persona("Juan", 25, "Madrid")

print(p.nombre)     # Juan
print(p.edad)       # 25
print(p.ciudad)     # Madrid

# También indexado
print(p[0])         # Juan
```

---

## Parte III: Conjuntos (Set)

Colecciones **sin duplicados**, **mutables**:

### Creación

```python
# Con {}
conjunto = {1, 2, 3, 4, 5}

# Con set()
conjunto = set([1, 1, 2, 2, 3])  # {1, 2, 3}

# Vacío (debe ser set(), no {})
vacío = set()

# Strings
letras = set("aabbcc")  # {'a', 'b', 'c'}
```

### No Duplicados

```python
conjunto = {1, 1, 1, 2, 2, 3}
print(conjunto)         # {1, 2, 3}
print(len(conjunto))    # 3
```

### Métodos

```python
s = {1, 2, 3}

# .add() - agregar
s.add(4)                # {1, 2, 3, 4}

# .remove() - eliminar (error si no existe)
s.remove(1)             # {2, 3, 4}

# .discard() - eliminar (sin error)
s.discard(5)            # No error si no existe

# .pop() - aleatorio
s = {1, 2, 3}
elemento = s.pop()      # Elimina y devuelve aleatorio

# .clear() - vaciar
# s.clear()              # set()
```

### Operaciones de Conjuntos

```python
a = {1, 2, 3, 4}
b = {3, 4, 5, 6}

# Unión (todos elementos)
print(a | b)            # {1, 2, 3, 4, 5, 6}
print(a.union(b))       # equivalente

# Intersección (elementos comunes)
print(a & b)            # {3, 4}
print(a.intersection(b)) # equivalente

# Diferencia (en a pero no en b)
print(a - b)            # {1, 2}
print(a.difference(b))  # equivalente

# Diferencia simétrica (en uno o otro, no en ambos)
print(a ^ b)            # {1, 2, 5, 6}
print(a.symmetric_difference(b)) # equivalente
```

### Subconjuntos

```python
a = {1, 2}
b = {1, 2, 3, 4}

print(a <= b)           # True (a subconjunto de b)
print(a.issubset(b))    # equivalente

print(b >= a)           # True (b superconjunto de a)
print(b.issuperset(a))  # equivalente

print(a.isdisjoint({5, 6}))  # True (sin elementos comunes)
```

### Búsqueda

```python
s = {1, 2, 3, 4, 5}

# O(1) búsqueda rápida
if 3 in s:
    print("Encontrado")

# Eliminar duplicados
lista = [1, 1, 2, 2, 3, 3]
unico = list(set(lista))  # [1, 2, 3]
```

---

## Comparación de Estructuras

| Característica | Lista | Tupla | Conjunto |
|---|---|---|---|
| **Mutable** | ✅ Sí | ❌ No | ✅ Sí |
| **Ordenada** | ✅ Sí | ✅ Sí | ❌ No |
| **Duplicados** | ✅ Permite | ✅ Permite | ❌ No |
| **Indexable** | ✅ Sí [i] | ✅ Sí [i] | ❌ No |
| **Hashable** | ❌ No | ✅ Sí | ❌ No |
| **Búsqueda O(n)** | ✅ | ✅ | ❌ O(1) |
| **Uso** | General | Retornos, claves | Únicos, intersección |

---

## Conversiones

```python
# Lista a Tupla
lista = [1, 2, 3]
tupla = tuple(lista)  # (1, 2, 3)

# Tupla a Lista
tupla = (1, 2, 3)
lista = list(tupla)   # [1, 2, 3]

# Lista a Conjunto
lista = [1, 1, 2, 3]
conjunto = set(lista)  # {1, 2, 3}

# String a Lista
palabra = "Hola"
lista = list(palabra)  # ['H', 'o', 'l', 'a']

# String a Conjunto
conjunto = set("aabbcc")  # {'a', 'b', 'c'}
```

---

## List Comprehension

Crear listas compactamente:

```python
# Tradicional
cuadrados = []
for x in range(5):
    cuadrados.append(x ** 2)

# Comprehension
cuadrados = [x ** 2 for x in range(5)]  # [0, 1, 4, 9, 16]

# Con condición
pares = [x for x in range(10) if x % 2 == 0]  # [0, 2, 4, 6, 8]

# Anidada
matriz = [[i + j for j in range(3)] for i in range(3)]
# [[0, 1, 2], [1, 2, 3], [2, 3, 4]]
```

---

## Resumen

| Operación | Lista | Tupla | Conjunto |
|-----------|-------|-------|----------|
| Crear | `[]` | `()` | `{}` |
| Agregar | `.append()` | N/A | `.add()` |
| Acceso | `[i]` | `[i]` | N/A |
| Búsqueda | `in` O(n) | `in` O(n) | `in` O(1) |
| Ordenar | `.sort()` | N/A | N/A |
| Copiar | `.copy()` | N/A | `.copy()` |

---

**Siguiente**: [diccionarios.md](diccionarios.md)
