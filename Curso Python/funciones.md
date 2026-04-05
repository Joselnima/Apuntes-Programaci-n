# Funciones en Python

> **Cómo empaquetarcódigo en cajas reutilizables - Evita repetirte y escribe código más limpio**

---

## ¿QÉ es una Función?

Una **función** es como un **robot que haces**:

- 🤖 Algunos robots preparan café
- 🤖 Otros calculan precios
- 🤖 Otros envían emails

Cada robot:
1. **Recibe órdenes** (parámetros)
2. **Hace su trabajo**
3. **Te devuelve un resultado** (return)

**Ejemplo en la vida real:**
```
Tú:    "¡Oye camarero, dame un café!"
Camarero: (recibe la orden)
Camarero: (prepara el café)
Camarero: (te lo entrega)
Tú:    "Gracias!"
```

En Python:
```python
def preparar_cafe(tipo):
    # Recibe orden (parámetro)
    print(f"Preparando {tipo}...")
    cafe = f"{tipo} listo"
    return cafe  # Devuelve resultado

mi_cafe = preparar_cafe("expreso")
print(mi_cafe)  # expreso listo
```

---

## ¿PARA QUÉ sirven las Funciones?

### Problema: Código Repetido (Difícil)

```python
# ❌ Imagina que necesitas validar 3 emails

# Validación 1
email1 = "juan@gmail.com"
if "@" in email1 and "." in email1:
    print(f"{email1} es válido")
else:
    print(f"{email1} es inválido")

# Validación 2 (copiar-pegar...)
email2 = "maria@hotmail.com"
if "@" in email2 and "." in email2:
    print(f"{email2} es válido")
else:
    print(f"{email2} es inválido")

# Validación 3 (más copiar-pegar...)
email3 = "pedro@yahoo.com"
if "@" in email3 and "." in email3:
    print(f"{email3} es válido")
else:
    print(f"{email3} es inválido")

# ¿Problemas?
# - 20 líneas para hacer LO MISMO 3 veces
# - Si descubres un error, debes cambiarlo en 3 lugares
# - Es aburrido escribir lo mismo una y otra vez
```

### Solución: Función (Fácil)

```python
# ✅ Defines la validación UNA VEZ
def validar_email(email):
    if "@" in email and "." in email:
        return True
    else:
        return False

# Luego la usas TANTAS VECES como quieras
print(validar_email("juan@gmail.com"))    # True
print(validar_email("maria@hotmail.com"))  # True
print(validar_email("pedro123"))           # False

# ¿Ventajas?
# - Solo 9 líneas (3 en la función, 3 en uso)
# - Si cambias la lógica, solo cambias la función
# - Código limpio y profesional
# - Otros programadores lo entienden
```

---

## ¿CÓMO funcionan las Funciones?

### Paso 1: Definición (Crear el Robot)

```python
def saludar(nombre):
    """Función que saluda a alguien"""
    print(f"Hola {nombre}")
```

Esto **no ejecuta** nada. Solo le dices a Python "aquí hay una función llamada 'saludar'"

### Paso 2: Invocación (Activar el Robot)

```python
saludar("Juan")  # Aquí SÍ se ejecuta
# Output: Hola Juan
```

### Paso 3: Resultado (El Robot Entrega)

```python
def calcular_precio_con_iva(precio):
    """Calcula el precio con 21% IVA"""
    iva = precio * 0.21
    precio_final = precio + iva
    return precio_final  # Devuelve el resultado

precio_sin_iva = 100
precio_con_iva = calcular_precio_con_iva(precio_sin_iva)
print(precio_con_iva)  # 121.0
```

---

## Parte I: Tu Primera Función - Lo Más Básico

### Definición

```python
def saludar():
    print("Hola Mundo")

# Llamar/ejecutar
saludar()  # Output: Hola Mundo
```

### Anatomía

```python
def nombre_función(parámetro1, parámetro2):
    \"\"\"Docstring: qué hace la función\"\"\"
    # Cuerpo de la función
    resultado = parámetro1 + parámetro2
    return resultado

# Llamar
resultado_final = nombre_función(5, 3)
```

**Reglas**:
- Nombre en `snake_case`
- Parámetros entre paréntesis
- Docstring en triple comilla (opcional pero recomendado)
- `return` para devolver valor (opcional)

---

## Parte II: Parámetros y Return

### Parámetro Simple

```python
def cuadrado(x):
    return x ** 2

print(cuadrado(5))  # 25
```

### Múltiples Parámetros

```python
def sumar(a, b, c):
    return a + b + c

print(sumar(1, 2, 3))  # 6
```

### Parámetros por Defecto

```python
def saludar(nombre, saludo="Hola"):
    return f"{saludo} {nombre}"

print(saludar("Juan"))                      # Hola Juan
print(saludar("María", "Buenos días"))      # Buenos días María
```

### Return Múltiple (Tupla)

```python
def obtener_datos():
    nombre = "Juan"
    edad = 25
    ciudad = "Madrid"
    return nombre, edad, ciudad

# Desempaquetar
n, e, c = obtener_datos()
print(n, e, c)  # Juan 25 Madrid

# O como tupla
datos = obtener_datos()
print(datos)    # ('Juan', 25, 'Madrid')
```

---

## Parte III: *args (Argumentos Variables)

Pasar **número variable de argumentos**:

```python
def sumar_todos(*números):
    total = 0
    for num in números:
        total += num
    return total

print(sumar_todos(1, 2, 3))           # 6
print(sumar_todos(1, 2, 3, 4, 5))     # 15
print(sumar_todos(10))                # 10
```

**Importante**: `*args` es una **tupla** dentro de la función.

### Con Otros Parámetros

```python
def imprimir_info(nombre, *edades):
    print(f"Nombre: {nombre}")
    print(f"Edades: {edades}")

imprimir_info("Juan", 25, 30, 35)
# Output:
# Nombre: Juan
# Edades: (25, 30, 35)
```

---

## Parte IV: **kwargs (Argumentos con Nombre)

Pasar argumentos con **nombres clave**:

```python
def crear_usuario(**propiedades):
    for clave, valor in propiedades.items():
        print(f"{clave}: {valor}")

crear_usuario(nombre="Juan", edad=25, ciudad="Madrid")
# Output:
# nombre: Juan
# edad: 25
# ciudad: Madrid
```

**Importante**: `**kwargs` es un **diccionario** dentro de la función.

### Combinado

```python
def completo(pos1, pos2, *args, **kwargs):
    print(f"pos1: {pos1}")
    print(f"pos2: {pos2}")
    print(f"args: {args}")
    print(f"kwargs: {kwargs}")

completo(1, 2, 3, 4, nombre="Juan", edad=25)
# Output:
# pos1: 1
# pos2: 2
# args: (3, 4)
# kwargs: {'nombre': 'Juan', 'edad': 25}
```

---

## Parte V: Docstrings

Documenta qué hace tu función:

```python
def calcular_promedio(calificaciones):
    \"\"\"
    Calcula el promedio de un conjunto de calificaciones.
    
    Args:
        calificaciones (list): Lista de números
    
    Returns:
        float: El promedio
    
    Raises:
        ValueError: Si la lista está vacía
    
    Ejemplo:
        >>> calcular_promedio([80, 90, 100])
        90.0
    \"\"\"
    if not calificaciones:
        raise ValueError("La lista no puede estar vacía")
    return sum(calificaciones) / len(calificaciones)

# Ver docstring
print(calcular_promedio.__doc__)
help(calcular_promedio)
```

---

## Parte VI: Type Hints (Python 3.5+)

Indican qué tipos espera:

```python
def sumar(a: int, b: int) -> int:
    \"\"\"Suma dos integers\"\"\"
    return a + b

sumar(5, 3)      # OK
sumar("5", "3")  # Funciona, pero no es lo esperado

# Con tipos complejos
from typing import List, Dict

def procesar(nombres: List[str]) -> Dict[str, int]:
    return {nombre: len(nombre) for nombre in nombres}

resultado = procesar(["Juan", "María"])
# {'Juan': 4, 'María': 5}
```

---

## Parte VII: Funciones Anónimas (Lambda)

Funciones sin nombre, pequeñas:

```python
# Forma tradicional
def cuadrado(x):
    return x ** 2

# Lambda (equivalente)
cuadrado = lambda x: x ** 2
print(cuadrado(5))  # 25

# Parámetros múltiples
suma = lambda x, y: x + y
print(suma(3, 4))  # 7

# Usado típicamente en filter, map, sorted
números = [1, 2, 3, 4, 5]
pares = list(filter(lambda x: x % 2 == 0, números))
print(pares)  # [2, 4]
```

---

## Parte VIII: map() y filter()

Funciones de programación funcional:

### map()

Aplica función a cada elemento:

```python
números = [1, 2, 3, 4, 5]

# Tradicional
cuadrados = []
for num in números:
    cuadrados.append(num ** 2)

# Con map
cuadrados = list(map(lambda x: x ** 2, números))
print(cuadrados)  # [1, 4, 9, 16, 25]

# Mejor: list comprehension
cuadrados = [x ** 2 for x in números]
```

### filter()

Filtra elementos según condición:

```python
números = [1, 2, 3, 4, 5, 6]

# Tradicional
pares = []
for num in números:
    if num % 2 == 0:
        pares.append(num)

# Con filter
pares = list(filter(lambda x: x % 2 == 0, números))
print(pares)  # [2, 4, 6]

# Mejor: list comprehension
pares = [x for x in números if x % 2 == 0]
```

---

## Parte IX: Scope (Alcance de Variables)

### Local

```python
def mi_función():
    x = 10  # Local
    print(x)

mi_función()  # 10
print(x)      # NameError: x no está definida globalmente
```

### Global

```python
x = 10  # Global

def mi_función():
    print(x)  # Accede a global

mi_función()  # 10
print(x)      # 10
```

### Modificar Global

```python
contador = 0

def incrementar():
    global contador  # Declara que modificaremos la global
    contador += 1

incrementar()
print(contador)  # 1
```

### Nonlocal (Anidadas)

```python
def externa():
    x = 10
    
    def interna():
        nonlocal x  # Modifica la de la función que encierra
        x += 1
        print(x)
    
    interna()
    print(x)

externa()
# Output: 11, 11
```

---

## Parte X: Closures (Funciones que Retornan Funciones)

```python
def crear_multiplicador(factor):
    def multiplicar(x):
        return x * factor
    return multiplicar

duplicar = crear_multiplicador(2)
triplicar = crear_multiplicador(3)

print(duplicar(5))      # 10
print(triplicar(5))     # 15
```

---

## Parte XI: Decoradores (Introducción)

Funciones que envuelven otras funciones:

```python
def saludar(func):
    def envuelta():
        print("--- Decorador ---")
        func()
        print("--- Fin ---")
    return envuelta

@saludar
def mi_función():
    print("Contenido de función")

mi_función()
# Output:
# --- Decorador ---
# Contenido de función
# --- Fin ---
```

(Más detalles en `decoradores.md`)

---

## Parte XII: Funciones Predefinidas Útiles

```python
# len() - longitud
print(len([1, 2, 3]))       # 3
print(len("Hola"))          # 4

# sum() - suma
print(sum([1, 2, 3, 4]))    # 10

# max() y min()
print(max([1, 5, 3]))       # 5
print(min([1, 5, 3]))       # 1

# abs() - valor absoluto
print(abs(-5))              # 5

# round() - redondear
print(round(3.14159, 2))    # 3.14

# sorted() - ordenar
print(sorted([3, 1, 4, 1, 5]))  # [1, 1, 3, 4, 5]

# reversed() - invertir
print(list(reversed([1, 2, 3])))  # [3, 2, 1]

# enumerate() - índice + valor
for i, v in enumerate(['a', 'b', 'c']):
    print(f"{i}: {v}")
```

---

## Resumen

| Concepto | Uso | Ejemplo |
|----------|-----|---------|
| def | Definir función | def sumar(a, b): |
| return | Devolver valor | return a + b |
| Parámetros | Entradas | def f(x, y): |
| Por defecto | Valor predeterminado | def f(x=5): |
| *args | Argumentos variables | def f(*args): |
| **kwargs | Argumentos nombrados | def f(**kwargs): |
| Docstring | Documentación | \"\"\"Descripción\"\"\" |
| Type hints | Tipos esperados | def f(x: int) -> str: |
| lambda | Función anónima | lambda x: x*2 |
| map() | Aplicar función | map(func, lista) |
| filter() | Filtrar | filter(condición, lista) |
| Closure | Función dentro función | def f(): def g(): ... |

---

**Siguiente**: [strings.md](strings.md)
