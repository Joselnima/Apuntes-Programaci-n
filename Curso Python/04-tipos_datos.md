# Tipos de Datos en Python

> **Aprende sobre los diferentes tipos de información que Python puede almacenar - El lenguaje entiende automáticamente qué tipo es cada cosa**

---

## ¿QÉ es un "Tipo de Dato"?

Un **tipo de dato** es la categoría de una información:

- 🔢 **25** es un NÚMERO (entero)
- 📝 **"Hola"** es TEXTO (cadena)
- ✅ **True** es VERDADERO/FALSO (booleano)
- 📦 **[1,2,3]** es una LISTA (colección)

Es como en la vida real:
- Un **teléfono** es un número (tipo: int)
- Un **nombre** es texto (tipo: str)
- **¿Está activo?** es sí/no (tipo: bool)

---

## ¿PARA QUÉ importa el tipo?

### Cada tipo tiene operaciones diferentes

```python
# Con NÚMEROS puedes hacer operaciones matemáticas
5 + 3          # Suma = 8 ✅
5 - 3          # Resta = 2 ✅
5 * 3          # Multiplicación = 15 ✅

# Con TEXTO puedes combinarlos
"Hola" + "Mundo"      # Concatenación = "HolaMundo" ✅
"Hola" * 3            # Repetición = "HolaHolaHola" ✅

# Pero NO puedes mezclar (normalmente)
"Hola" + 5            # ERROR ❌
# Python dice: "¿Cómo sumo TEXTO con NÚMERO?"
```

### Cada tipo tiene métodos propios

```python
# Los NÚMEROS tienen métodos matemáticos
numero = 5
numero.bit_length()   # Qué número de bits necesita

# Los TEXTOS tienen métodos de texto
texto = "hola"
texto.upper()         # "HOLA" (mayúsculas)
texto.replace("h", "H")  # "Hola"

# Las LISTAS tienen métodos de lista
lista = [1, 2, 3]
lista.append(4)       # Añade al final: [1, 2, 3, 4]
lista.sort()          # Ordena: [1, 2, 3, 4]
```

---

## ¿CÓMO Python sabe el tipo?

Lo adivina automáticamente (tipado dinámico):

```python
# Por el valor que le das
x = 5              # Python: "Esto es un int"
y = 5.0            # Python: "Esto es un float"
z = "5"            # Python: "Esto es un str"
w = True           # Python: "Esto es un bool"

# Para verificar
print(type(x))     # <class 'int'>
print(type(y))     # <class 'float'>
print(type(z))     # <class 'str'>
print(type(w))     # <class 'bool'>
```

---

## Parte I: Tipos Numéricos - Los Números

```python
edad = 25
temperatura = -10
año = 2026

# Grandes números (sin límite en Python)
numero_grande = 999999999999999999999

# Binario, octal, hexadecimal
binario = 0b1010        # 10 en decimal
octal = 0o12            # 10 en decimal
hexadecimal = 0xa       # 10 en decimal

# Operaciones
suma = 5 + 3        # 8
resta = 10 - 3      # 7
multiplicación = 4 * 3    # 12
división_entera = 10 // 3  # 3 (trunca)
módulo = 10 % 3     # 1
potencia = 2 ** 8   # 256
```

### 2. Floats (float)

Números con decimales:

```python
precio = 19.99
pi = 3.14159
gravedad = 9.81
notación_científica = 1.5e-3  # 0.0015

# Cuidado con precisión
print(0.1 + 0.2)    # 0.30000000000000004 (no exacto)
print(0.1 + 0.2 == 0.3)  # False (¡!)

# Solución: usar Decimal para finanzas
from decimal import Decimal
dinero = Decimal("0.1") + Decimal("0.2")  # Exacto
```

### 3. Complex (números complejos)

Para matemáticas avanzadas:

```python
z = 3 + 4j
print(z.real)       # 3.0
print(z.imag)       # 4.0
print(abs(z))       # 5.0 (magnitud)
```

### Conversión entre Números

```python
# Int a Float
x = int(5.99)       # 5 (trunca)
y = float(5)        # 5.0

# String a número
num = int("123")    # 123
precio = float("19.99")  # 19.99

# Con base diferente
x = int("1010", 2)  # 10 (binario)
x = int("1C", 16)   # 28 (hexadecimal)
```

---

## Parte II: Boolean (bool)

Solo dos valores: `True` o `False`:

```python
es_mayor = True
tiene_licencia = False

# Se crea comparando
x = 5
es_positivo = x > 0  # True
es_par = x % 2 == 0  # False

# Operaciones lógicas
print(True and False)    # False
print(True or False)     # True
print(not True)          # False

# Conversión
print(bool(1))           # True
print(bool(0))           # False
print(bool(""))          # False (string vacío)
print(bool("texto"))     # True
print(bool([]))          # False (lista vacía)
print(bool([1]))         # True
```

---

## Parte III: Strings (str)

Secuencias de caracteres. **Inmutables** (no se pueden cambiar):

### Creación

```python
# Comillas simples o dobles (igual)
nombre = "Juan"
nombre = 'Juan'

# Triples para multi-línea
descripción = """
Esto es una descripción
que ocupa varias líneas
sin errores
"""

# Raw strings (para \, rutas, regex)
ruta = r"C:\Users\Juan\Documentos"  # No interpreta \
patron = r"\d+"  # Expresión regular
```

### Acceso e Indexing

```python
texto = "Python"

# Acceso por índice (comienza en 0)
print(texto[0])     # P
print(texto[1])     # y
print(texto[-1])    # n (último)
print(texto[-2])    # o (penúltimo)

# Longitud
print(len(texto))   # 6

# ❌ No puedes modificar
# texto[0] = "J"  # Error: strings son inmutables
```

### Slicing (Rebanadas)

```python
texto = "Python"

print(texto[0:3])      # Pyt (índices 0, 1, 2)
print(texto[3:])       # hon (desde 3 hasta el final)
print(texto[:3])       # Pyt (desde inicio hasta 3)
print(texto[::2])      # Pto (cada 2 caracteres)
print(texto[::-1])     # nohtyP (invertido)
```

### Métodos útiles

```python
texto = "  Hola Mundo  "

# Manipulación
print(texto.strip())           # Hola Mundo (sin espacios)
print(texto.lower())           # hola mundo (minúscula)
print(texto.upper())           # HOLA MUNDO (mayúscula)
print(texto.replace("Hola", "Chao"))  # Chao Mundo

# Búsqueda
print("Hola" in "Hola Mundo")  # True
print(texto.find("Mundo"))     # 7 (índice donde empieza)
print(texto.count("o"))        # 2 (cuántas veces aparece)

# División y unión
palabras = "Hola Mundo Python".split()  # ['Hola', 'Mundo', 'Python']
unido = " - ".join(palabras)   # Hola - Mundo - Python

# Verificación
print("123".isdigit())         # True (solo dígitos)
print("abc".isalpha())         # True (solo letras)
print("abc123".isalnum())      # True (letras y números)
print("  ".isspace())          # True (solo espacios)
```

### F-strings (Formateo Moderno)

```python
nombre = "Juan"
edad = 25
pi = 3.14159

print(f"Me llamo {nombre}")
print(f"{nombre} tiene {edad} años")
print(f"Pi es aproximadamente {pi:.2f}")
```

---

## Parte IV: None (NoneType)

Representa la ausencia de valor:

```python
resultado = None
print(resultado)  # None

# Usado como "sin valor"
def buscar(lista, valor):
    if valor in lista:
        return lista.index(valor)
    return None  # No encontrado

# Verificar
if resultado is None:
    print("No hay resultado")
```

---

## Parte V: Tipos de Colecciones

### Lists (Listas)

Secuencias **mutables** de elementos:

```python
# Creación
números = [1, 2, 3, 4, 5]
mixta = [1, "texto", 3.14, True]
vacía = []

# Acceso
print(números[0])       # 1
print(números[-1])      # 5

# Modificación
números[0] = 10
números.append(6)
números.extend([7, 8])

# Métodos
print(len(números))
print(max(números))
print(min(números))
print(5 in números)
```

### Tuples (Tuplas)

Secuencias **inmutables**:

```python
# Creación
punto = (10, 20)
colores = ("rojo", "verde", "azul")
singleton = (1,)  # Nota la coma

# Acceso
print(punto[0])     # 10

# ❌ No mutable
# punto[0] = 5  # Error

# Unpacking
x, y = punto  # x=10, y=20

# Tuplas nombradas
from collections import namedtuple
Persona = namedtuple("Persona", ["nombre", "edad"])
p = Persona("Juan", 25)
print(p.nombre)  # Juan
```

### Dictionaries (Diccionarios)

Colecciones **mutables** de pares clave-valor:

```python
# Creación
persona = {"nombre": "Juan", "edad": 25, "ciudad": "Madrid"}
vacío = {}

# Acceso
print(persona["nombre"])    # Juan
print(persona.get("edad"))  # 25
print(persona.get("teléfono", "No disponible"))  # No disponible

# Modificación
persona["edad"] = 26
persona["teléfono"] = "555-1234"

# Métodos
print(persona.keys())       # dict_keys(['nombre', 'edad', ...])
print(persona.values())     # dict_values(['Juan', 26, ...])
print(persona.items())      # dict_items([('nombre', 'Juan'), ...])

# Eliminación
del persona["teléfono"]
```

### Sets (Conjuntos)

Colecciones **mutables** sin duplicados:

```python
# Creación
numeros = {1, 2, 3, 4, 5}
vacío = set()  # No usas {}

# No hay duplicados
conjunto = {1, 1, 1, 2, 2, 3}
print(conjunto)  # {1, 2, 3}

# Operaciones matemáticas
a = {1, 2, 3}
b = {2, 3, 4}

print(a | b)      # {1, 2, 3, 4} - unión
print(a & b)      # {2, 3} - intersección
print(a - b)      # {1} - diferencia
print(a ^ b)      # {1, 4} - diferencia simétrica

# Métodos
a.add(4)
a.remove(1)
print(3 in a)
```

---

## Parte VI: Tipos Dinámicos vs Verificación de Tipos

### Type Checking en Runtime

```python
x = 5
print(type(x))          # <class 'int'>
print(isinstance(x, int))  # True

lista = [1, 2, 3]
print(type(lista))      # <class 'list'>
print(isinstance(lista, list))  # True
```

### Type Hints (Python 3.5+)

Documentan qué tipos esperas, pero **no se fuerzan**:

```python
def sumar(a: int, b: int) -> int:
    """Suma dos enteros."""
    return a + b

resultado = sumar(5, 3)  # OK

resultado = sumar("5", "3")  # Funciona, pero no es lo esperado
# → "53" (concatenación, no suma)
```

### Type Checking con mypy

Para verificar tipos antes de ejecutar:

```bash
pip install mypy
mypy archivo.py
```

---

## Parte VII: Conversión de Tipos

| De/A | int | float | str | bool |
|------|-----|-------|-----|------|
| int(x) | - | int(3.14)=3 | str(5)="5" | bool(1)=T |
| float(x) | float(5)=5.0 | - | str(3.14)="3.14" | bool(0.1)=T |
| str(x) | int("5")=5 | float("3.14")=3.14 | - | bool("x")=T |
| bool(x) | bool(1)=T | bool(0.0)=F | bool("")=F | - |

---

## Parte VIII: Type Coercion Implícita

Python a veces convierte automáticamente:

```python
# Suma int + float → float
resultado = 5 + 2.5  # 7.5

# String multiplica int
texto = "Hola " * 3  # Hola Hola Hola

# Comparación
print(5 == 5.0)      # True (igual valor)
print(5 is 5.0)      # False (tipos diferentes)
```

---

## Parte IX: Memoria y Referencia

```python
# Immutable
a = 5
b = a      # Copia el valor
a = 10
print(b)   # 5 (sin cambios)

# Mutable
lista1 = [1, 2, 3]
lista2 = lista1      # Referencia al mismo objeto
lista1.append(4)
print(lista2)        # [1, 2, 3, 4] (afectado)

# Para copiar:
lista3 = lista1.copy()  # Nueva lista
lista1.append(5)
print(lista3)        # [1, 2, 3, 4] (sin cambios)
```

---

## Resumen de Tipos

| Tipo | Ejemplo | Mutable | Ordenado | Duplicados |
|------|---------|---------|----------|-----------|
| int | 42 | N/A | - | - |
| float | 3.14 | N/A | - | - |
| bool | True, False | N/A | - | - |
| str | "Hola" | ❌ | ✅ | ✅ |
| list | [1, 2, 3] | ✅ | ✅ | ✅ |
| tuple | (1, 2, 3) | ❌ | ✅ | ✅ |
| dict | {"a": 1} | ✅ | ✅* | ✅ (clave) |
| set | {1, 2, 3} | ✅ | ❌ | ❌ |

---

**Siguiente**: [operadores.md](operadores.md)
