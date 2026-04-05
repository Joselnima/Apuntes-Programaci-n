# Comprehensiones en Python

> **Una forma mágica y rápida de crear listas, diccionarios y sets**

---

## ¿QUÉ es una Comprehension?

Una **comprehension** = escribir un **for loop en una línea**

```python
# Forma larga (4 líneas)
cuadrados = []
for x in range(5):
    cuadrados.append(x ** 2)
print(cuadrados)  # [0, 1, 4, 9, 16]

# Forma comprehension (1 línea) - MISMA función
cuadrados = [x ** 2 for x in range(5)]
print(cuadrados)  # [0, 1, 4, 9, 16]
```

**Analogía:**
- 📋 Forma larga = escribir receta paso a paso
- 🧙 Comprehension = hechizo que hace TODO de una vez

---

## ¿PARA QÉ sirven?

### 1. Transformar listas (cambiar cada item)
```python
precios = [10, 20, 30]
precios_con_impuesto = [p * 1.21 for p in precios]  # Aplicar 21% impuesto
print(precios_con_impuesto)  # [12.1, 24.2, 36.3]
```

### 2. Filtrar (elegir solo algunos items)
```python
números = [1, 2, 3, 4, 5, 6, 7, 8]
solo_pares = [x for x in números if x % 2 == 0]
print(solo_pares)  # [2, 4, 6, 8]
```

### 3. Crear tablas de multiplicar
```python
tabla_3 = [3 * i for i in range(1, 11)]
print(tabla_3)  # [3, 6, 9, 12, 15, 18, 21, 24, 27, 30]
```

### 4. Limpiar texto
```python
texto = "Python 3.11"
solo_letras = [c for c in texto if c.isalpha()]
print(solo_letras)  # ['P', 'y', 't', 'h', 'o', 'n']
```

---

## ¿CÓMO funcionan?

### Estructura General

```python
[  TRANSFORMACIÓN  for  VARIABLE  in  ITERABLE  if  CONDICIÓN  ]
   └─ opcional            └─ cada item      └─ opcional
```

**Ejemplo paso a paso:**
```python
[x * 2 for x in [1, 2, 3] if x > 1]
 ↑       ↑ ↑ ↑  ↑ ↑   ↑  ↑ ↑  ↑
 resultado transformación cada item   condición
```

**Cómo se ejecuta:**
1. Toma CADA item de `[1, 2, 3]`
2. Verifica SI `x > 1`
3. Si pasa, aplica TRANSFORMACIÓN `x * 2`
4. Guarda resultado en lista

Resultado: `[4, 6]` (solo 2 y 3 pasan la condición)

---

## Parte I: List Comprehensions

### Básico

```python
# Transformar cada elemento
números = [1, 2, 3, 4, 5]
cuadrados = [x ** 2 for x in números]
print(cuadrados)  # [1, 4, 9, 16, 25]

# Con función
def doble(x):
    return x * 2

dobrados = [doble(x) for x in números]
print(dobrados)   # [2, 4, 6, 8, 10]

# Strings
palabra = "Python"
mayúsculas = [c.upper() for c in palabra]
print(mayúsculas)  # ['P', 'Y', 'T', 'H', 'O', 'N']
```

### Con Condición

```python
números = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

# Solo pares
pares = [x for x in números if x % 2 == 0]
print(pares)  # [2, 4, 6, 8, 10]

# Filtro y transformación
pares_dobles = [x * 2 for x in números if x % 2 == 0]
print(pares_dobles)  # [4, 8, 12, 16, 20]

# Múltiples condiciones
resultado = [x for x in números if x > 3 if x < 8]
print(resultado)  # [4, 5, 6, 7]
```

### If/Else en Comprehension

```python
números = [1, 2, 3, 4, 5]

# Transformar cada elemento
pares_impares = ["par" if x % 2 == 0 else "impar" for x in números]
print(pares_impares)  # ['impar', 'par', 'impar', 'par', 'impar']

# Con función
def clasificar(x):
    if x < 3:
        return "pequeño"
    elif x < 5:
        return "medio"
    else:
        return "grande"

clasificados = [clasificar(x) for x in números]
print(clasificados)  # ['pequeño', 'pequeño', 'medio', 'medio', 'grande']
```

### Anidadas

```python
# Matriz 3x3
matriz = [[i + j for j in range(3)] for i in range(3)]
print(matriz)
# [[0, 1, 2], [1, 2, 3], [2, 3, 4]]

# Aplanar matriz
matriz = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
plana = [x for fila in matriz for x in fila]
print(plana)  # [1, 2, 3, 4, 5, 6, 7, 8, 9]

# Aplanar con condición
números_mayores = [x for fila in matriz for x in fila if x > 5]
print(números_mayores)  # [6, 7, 8, 9]
```

---

## Parte II: Dict Comprehensions

### Básico

```python
números = [1, 2, 3, 4, 5]

# Crear diccionario
cuadrados = {x: x ** 2 for x in números}
print(cuadrados)  # {1: 1, 2: 4, 3: 9, 4: 16, 5: 25}

# Desde strings
código_países = {país: código for país, código in [("España", "ES"), ("Francia", "FR")]}
print(código_países)  # {'España': 'ES', 'Francia': 'FR'}
```

### Con Condición

```python
números = range(10)

# Solo pares
pares_cuadrados = {x: x ** 2 for x in números if x % 2 == 0}
print(pares_cuadrados)  # {0: 0, 2: 4, 4: 16, 6: 36, 8: 64}

# Invertir diccionario
diccionario = {"a": 1, "b": 2, "c": 3}
invertido = {v: k for k, v in diccionario.items()}
print(invertido)  # {1: 'a', 2: 'b', 3: 'c'}
```

### Transformar Diccionarios

```python
diccionario = {"Juan": 25, "María": 30, "Pedro": 22}

# Sumar años
años_después = {nombre: edad + 5 for nombre, edad in diccionario.items()}
print(años_después)  # {'Juan': 30, 'María': 35, 'Pedro': 27}

# Filtrar mayores de 25
mayores = {nombre: edad for nombre, edad in diccionario.items() if edad > 25}
print(mayores)  # {'María': 30}
```

---

## Parte III: Set Comprehensions

```python
# Crear conjunto sin duplicados
números = [1, 1, 2, 2, 3, 3, 4, 4, 5, 5]
únicos = {x for x in números}
print(únicos)  # {1, 2, 3, 4, 5}

# Transformar
letras = "aabbcc"
únicas_mayúsculas = {c.upper() for c in letras}
print(únicas_mayúsculas)  # {'A', 'B', 'C'}

# Con condición
pares = {x for x in range(10) if x % 2 == 0}
print(pares)  # {0, 2, 4, 6, 8}

# Intersection (elementos comunes)
a = {x for x in range(10) if x % 2 == 0}
b = {x for x in range(10) if x % 3 == 0}
comunes = a & b
print(comunes)  # {0, 6}
```

---

## Parte IV: Generator Expressions

Comprensiones que generan datos bajo demanda (memory efficient):

```python
# List comprehension (todo en memoria)
lista = [x ** 2 for x in range(1000000)]

# Generator expression (bajo demanda)
generador = (x ** 2 for x in range(1000000))

# Usar
for valor in generador:
    print(valor)  # Imprime uno a uno, sin usar mucha memoria

# Convertir a lista si necesitas
lista = list(generador)
```

### Generator Functions

```python
# Sin generator
def números_hasta(n):
    resultado = []
    for i in range(n):
        resultado.append(i)
    return resultado

# Con generator (memory efficient)
def números_generador(n):
    for i in range(n):
        yield i  # yield = devuelve valor y pausa

# Usar
for num in números_generador(5):
    print(num)  # 0, 1, 2, 3, 4 (uno a uno)

# next() para obtener lo próximo
gen = números_generador(3)
print(next(gen))  # 0
print(next(gen))  # 1
print(next(gen))  # 2
# print(next(gen))  # StopIteration
```

---

## Parte V: Comparación: Performance

```python
import sys

# List comprehension
lista = [x ** 2 for x in range(10000)]
print(f"Lista: {sys.getsizeof(lista)} bytes")  # Mucho

# Generator expression
gen = (x ** 2 for x in range(10000))
print(f"Generator: {sys.getsizeof(gen)} bytes")  # Poco

# Usar generador no consume memoria hasta iterar
valores = sum(x ** 2 for x in range(10000))
print(valores)
```

---

## Parte VI: Casos de Uso Prácticos

### Extraer Datos

```python
usuarios = [
    {"nombre": "Juan", "edad": 25, "activo": True},
    {"nombre": "María", "edad": 30, "activo": False},
    {"nombre": "Pedro", "edad": 22, "activo": True}
]

# Nombres de usuarios activos
activos = [u["nombre"] for u in usuarios if u["activo"]]
print(activos)  # ['Juan', 'Pedro']

# Edades menores de 26
jóvenes_edades = [u["edad"] for u in usuarios if u["edad"] < 26]
print(jóvenes_edades)  # [25, 22]
```

### Procesararchivos

```python
# Leer líneas no vacías
with open("archivo.txt", "r") as f:
    líneas = [línea.strip() for línea in f if línea.strip()]

# Palabras únicas
with open("texto.txt", "r") as f:
    palabras = {palabra for línea in f for palabra in línea.split()}
```

### Crear Combinaciones

```python
# Combinaciones de dos listas
colores = ["rojo", "azul"]
tamaños = ["pequeño", "mediano", "grande"]

combinaciones = [(color, tamaño) for color in colores for tamaño in tamaños]
print(combinaciones)
# [('rojo', 'pequeño'), ('rojo', 'mediano'), ('rojo', 'grande'),
#  ('azul', 'pequeño'), ('azul', 'mediano'), ('azul', 'grande')]
```

---

## Parte VII: Cuándo NO Usar Comprehensions

```python
# ✅ Uso apropiado (simple, legible)
pares = [x for x in range(100) if x % 2 == 0]

# ❌ Muy compleja (dificulta lectura)
resultado = [
    {f"clave_{y}": x**y for y in range(3)}
    for x in range(5)
    if x % 2 == 0
    for y in range(3)
    if y != 1
]

# Mejor como función explícita
def crear_datos():
    resultado = []
    for x in range(5):
        if x % 2 != 0:
            continue
        item = {}
        for y in range(3):
            if y == 1:
                continue
            item[f"clave_{y}"] = x ** y
        resultado.append(item)
    return resultado
```

---

## Resumen

| Tipo | Sintaxis | Ejemplo |
|------|----------|---------|
| List | `[expr for item in seq]` | `[x*2 for x in l]` |
| List+filter | `[expr for item in seq if cond]` | `[x for x in l if x>5]` |
| Dict | `{key: val for item in seq}` | `{x: x**2 for x in l}` |
| Set | `{expr for item in seq}` | `{x for x in l}` |
| Generator | `(expr for item in seq)` | `(x*2 for x in l)` |

---

**Siguiente**: [decoradores.md](decoradores.md)
