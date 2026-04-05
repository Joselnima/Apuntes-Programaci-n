# Diccionarios en Python

> **Cómo guardar información etiquetada - Datos con nombres, no números**

---

## ¿QÉ es un Diccionario?

Un **diccionario** es como una **guía telefónica**:

```
"Juan" → 555-1234
"María" → 555-5678
```

En Python:
```python
edades = {
    "Juan": 25,
    "María": 30,
    "Pedro": 28
}

print(edades["Juan"])  # 25
```

**La diferencia con Listas:**
```python
# Lista - acceso por NÚMERO
nombres = ["Juan", "María", "Pedro"]
print(nombres[0])          # Juan (pero ¿por qué posición?)

# Diccionario - acceso por CLAVE (más claro)
edades = {"Juan": 25, "María": 30}
print(edades["Juan"])      # 25 (claro: edad de Juan)
```

---

## ¿PARA QÉ sirven?

### 1. Agrupar datos de una cosa
```python
persona = {
    "nombre": "Juan",
    "edad": 25,
    "ciudad": "Madrid"
}
# Mucho más claro que: persona = ["Juan", 25, "Madrid"]
```

### 2. Contar frecuencias
```python
votos = {"A": 15, "B": 12, "C": 8}
```

### 3. Traducir palabras
```python
traductor = {"perro": "dog", "gato": "cat"}
```

### 4. Guardar datos del usuario
```python
usuario = {
    "email": "juan@gmail.com",
    "contraseña_hash": "abc123...",
    "activo": True,
    "último_login": "2024-01-15"
}
```

---

## ¿CÓMO funcionan?

### Concepto: Clave → Valor

```
Clave       Valor
────────────────────
"nombre" → "Juan"
"edad"   → 25
"ciudad" → "Madrid"
```

En Python: `{"nombre": "Juan", "edad": 25}`

---

## Parte I: Creación de Diccionarios

### Forma 1: Literal

```python
# Vacío
diccionario = {}

# Con pares clave: valor
persona = {
    "nombre": "Juan",
    "edad": 25,
    "ciudad": "Madrid"
}

# Tipos mixtos como valores (permitido)
datos = {
    "nombre": "Juan",
    "edad": 25,
    "emails": ["juan@ejemplo.com", "j.juan@otro.com"],
    "activo": True
}
```

### Forma 2: Función dict()

```python
# Vacío
diccionario = dict()

# Desde lista de tuplas
diccionario = dict([("nombre", "Juan"), ("edad", 25)])

# Desde argumentos nombrados
diccionario = dict(nombre="Juan", edad=25, ciudad="Madrid")
```

### Anidados

```python
empleados = {
    "emp1": {"nombre": "Juan", "edad": 25},
    "emp2": {"nombre": "María", "edad": 30}
}

print(empleados["emp1"]["nombre"])  # Juan
```

---

## Parte II: Acceso a Valores

### Acceso Directo

```python
persona = {"nombre": "Juan", "edad": 25}

print(persona["nombre"])  # Juan
# print(persona["teléfono"])  # KeyError (no existe)
```

### .get() - Seguro

```python
persona = {"nombre": "Juan", "edad": 25}

print(persona.get("nombre"))           # Juan
print(persona.get("teléfono"))         # None (no error)
print(persona.get("teléfono", "N/A"))  # N/A (valor por defecto)
```

**El `.get()` es mejor porque no da error si la clave no existe.**

---

## Parte III: Modificación

### Agregar/Actualizar

```python
persona = {"nombre": "Juan"}

# Añadir
persona["edad"] = 25        # {"nombre": "Juan", "edad": 25}

# Actualizar
persona["nombre"] = "Pedro" # {"nombre": "Pedro", "edad": 25}

# Con .update()
persona.update({"ciudad": "Madrid", "edad": 26})
# {"nombre": "Pedro", "edad": 26, "ciudad": "Madrid"}
```

### Eliminar

```python
persona = {"nombre": "Juan", "edad": 25, "ciudad": "Madrid"}

# .pop() - elimina y devuelve
edad = persona.pop("edad")  # 25
print(persona)              # {"nombre": "Juan", "ciudad": "Madrid"}

# .pop() con defecto
teléfono = persona.pop("teléfono", "N/A")  # N/A

# del - elimina sin devolver
del persona["ciudad"]

# .clear() - vacía todo
# persona.clear()      # {}
```

---

## Parte IV: Iteración

### Iterar Claves

```python
persona = {"nombre": "Juan", "edad": 25, "ciudad": "Madrid"}

for clave in persona:
    print(clave)

# Output: nombre, edad, ciudad
```

### Iterar Valores

```python
for valor in persona.values():
    print(valor)

# Output: Juan, 25, Madrid
```

### Iterar Pares (Clave-Valor)

```python
for clave, valor in persona.items():
    print(f"{clave}: {valor}")

# Output:
# nombre: Juan
# edad: 25
# ciudad: Madrid
```

---

## Parte V: Métodos Útiles

```python
diccionario = {"a": 1, "b": 2, "c": 3}

# .keys() - todas las claves
claves = diccionario.keys()  # dict_keys(['a', 'b', 'c'])
print("a" in claves)          # True

# .values() - todos los valores
valores = diccionario.values()  # dict_values([1, 2, 3])
print(2 in valores)            # True

# .items() - pares clave-valor
pares = diccionario.items()    # dict_items([('a', 1), ('b', 2), ('c', 3)])

# .setdefault() - obtén valor o establece por defecto
diccionario.setdefault("d", 4)  # {"a": 1, "b": 2, "c": 3, "d": 4}

# .copy() - copia superficial
copia = diccionario.copy()

# .fromkeys() - crear dict con claves y valor por defecto
nuevo = dict.fromkeys(["x", "y", "z"], 0)  # {'x': 0, 'y': 0, 'z': 0}
```

---

## Parte VI: Búsqueda y Verificación

```python
persona = {"nombre": "Juan", "edad": 25}

# Clave existe
if "nombre" in persona:
    print("Encontrada")

# Clave no existe
if "teléfono" not in persona:
    print("No existe")

# Valor existe
if 25 in persona.values():
    print("Edad encontrada")

# Longitud
print(len(persona))  # 2
```

---

## Parte VII: Dict Comprehension

Crear diccionarios compactamente:

```python
# Tradicional
cuadrados = {}
for x in range(5):
    cuadrados[x] = x ** 2

# Comprehension
cuadrados = {x: x ** 2 for x in range(5)}
# {0: 0, 1: 1, 2: 4, 3: 9, 4: 16}

# Con condición
pares = {x: x ** 2 for x in range(10) if x % 2 == 0}
# {0: 0, 2: 4, 4: 16, 6: 36, 8: 64}

# Desde dos listas
claves = ["a", "b", "c"]
valores = [1, 2, 3]
diccionario = {k: v for k, v in zip(claves, valores)}
# {'a': 1, 'b': 2, 'c': 3}
```

---

## Parte VIII: defaultdict

Para evitar KeyError:

```python
from collections import defaultdict

# Sin defaultdict
contador = {}
for letra in "aaabbc":
    if letra not in contador:
        contador[letra] = 0
    contador[letra] += 1

# Con defaultdict
contador = defaultdict(int)
for letra in "aaabbc":
    contador[letra] += 1

print(contador)  # defaultdict(<class 'int'>, {'a': 3, 'b': 2, 'c': 1})

# Convertir a dict normal
print(dict(contador))  # {'a': 3, 'b': 2, 'c': 1}
```

---

## Parte IX: Counter

Contar elementos:

```python
from collections import Counter

# Crear contador
contador = Counter("aabbcc")
print(contador)  # Counter({'a': 2, 'b': 2, 'c': 2})

# From lista
números = [1, 1, 1, 2, 2, 3]
contador = Counter(números)
print(contador)  # Counter({1: 3, 2: 2, 3: 1})

# Métodos útiles
print(contador.most_common(2))    # [(1, 3), (2, 2)] - top 2
print(contador.most_common(1)[0]) # (1, 3) - el más común
print(contador.total())           # 6 - total elementos
```

---

## Parte X: Ordenamiento

```python
diccionario = {"c": 3, "a": 1, "b": 2}

# Ordenar por claves
ordenado = dict(sorted(diccionario.items()))
print(ordenado)  # {'a': 1, 'b': 2, 'c': 3}

# Ordenar por valores
ordenado = dict(sorted(diccionario.items(), key=lambda x: x[1]))
print(ordenado)  # {'a': 1, 'b': 2, 'c': 3}

# Descendiente
ordenado = dict(sorted(diccionario.items(), reverse=True))
print(ordenado)  # {'c': 3, 'b': 2, 'a': 1}
```

---

## Parte XI: Ejemplo Práctico: Contador de Palabras

```python
texto = "python python programación python código programación"

# Con Counter
from collections import Counter

palabras = texto.split()
contador = Counter(palabras)

print(contador)  # Counter({'python': 3, 'programación': 2, 'código': 1})

# Más comunes
for palabra, count in contador.most_common(2):
    print(f"{palabra}: {count}")

# Output:
# python: 3
# programación: 2
```

---

## Parte XII: Fusionando Diccionarios

```python
# Python 3.9+
d1 = {"a": 1, "b": 2}
d2 = {"c": 3, "d": 4}
d3 = d1 | d2  # {'a': 1, 'b': 2, 'c': 3, 'd': 4}

# Método universal
d3 = {**d1, **d2}  # {'a': 1, 'b': 2, 'c': 3, 'd': 4}

# .update()
d1.update(d2)  # Modifica d1
```

---

## Resumen de Métodos

| Método | Función | Ejemplo |
|--------|---------|---------|
| `[clave]` | Acceso | `dic["nombre"]` |
| `.get(clave, defecto)` | Acceso seguro | `dic.get("edad", 0)` |
| `.keys()` | Todas las claves | `dic.keys()` |
| `.values()` | Todos los valores | `dic.values()` |
| `.items()` | Pares clave-valor | `dic.items()` |
| `.pop(clave)` | Elimina y devuelve | `dic.pop("edad")` |
| `.update()` | Actualiza | `dic.update({})` |
| `.setdefault()` | Obtén o establece | `dic.setdefault("edad", 25)` |
| `.clear()` | Vacía | `dic.clear()` |

---

**Siguiente**: [poo.md](poo.md)
