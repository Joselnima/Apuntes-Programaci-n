# Módulos y Paquetes en Python

> **Dividir código en partes para reutilizar y mantener organizado**

---

## ¿QUÉ es un Módulo?

Un **módulo** = archivo `.py` que contiene código:

```
proyecto/
├── utilidades.py     ← Módulo 1
├── matemática.py     ← Módulo 2
└── main.py           ← Usa módulos 1 y 2
```

**Analogía:**
- 📄 Archivo = página de libro
- 📚 Módulo = capítulo de libro
- 🏢 Paquete = colección de libros

---

## ¿PARA QÉ sirven?

### 1. No repetir código
```python
# ❌ Sin módulos (repetir en cada lugar)
def saludar(nombre):
    return f"Hola {nombre}"

# archivo2.py
def saludar(nombre):      # REPETIRRR
    return f"Hola {nombre}"

# ✅ Con módulos (una sola vez)
# utilidades.py - define UNA VEZ
def saludar(nombre):
    return f"Hola {nombre}"

# main.py - usa
from utilidades import saludar
```

### 2. Mantener organizado
```python
proyecto/
├── usuarios/      (todo sobre usuarios)
├── pagos/         (todo sobre pagos)
├── reportes/      (todo sobre reportes)
```

### 3. Compartir código
```python
# Tu amigo usa tu módulo
from tu_proyecto.utilidades import saludar
```

### 4. Reutilizar sin copiar
```python
# 5 proyectos usan tu módulo de autenticación
# Cambias 1 vez, todos se actualizan
```

---

## ¿CÓMO funcionan?

### Ciclo de Importación

```
"from utilidades import saludar"
         ↓
Python busca: utilidades.py
         ↓
Ejecuta: TODO el archivo (funciones, variables, clases)
         ↓
Trae SOLO: saludar
         ↓
Disponible en main.py
```

**3 Formas de Importar:**

| Forma | Ejemplo | Cuándo usar |
|-------|---------|-----------|
| **Módulo completo** | `import math` | Muchas funciones |
| **Función específica** | `from math import sin` | 1-2 funciones |
| **Con nombre corto** | `import numpy as np` | Nombre muy largo |

---

## Parte I: Módulos

Un **módulo** es un archivo `.py`:

### Crear un Módulo

```python
# archivo: utilidades.py
def saludar(nombre):
    return f"Hola {nombre}"

def sumar(a, b):
    return a + b

VERSIÓN = "1.0"
```

### Importar un Módulo

```python
# archivo: main.py
import utilidades

print(utilidades.saludar("Juan"))  # Hola Juan
print(utilidades.sumar(5, 3))      # 8
print(utilidades.VERSIÓN)          # 1.0
```

### Importar Funciones Específicas

```python
from utilidades import saludar, sumar

print(saludar("Juan"))  # Sin prefijo
print(sumar(5, 3))
```

### Importar con Alias

```python
from utilidades import saludar as hola
import utilidades as util

print(hola("Juan"))     # Hola Juan
print(util.sumar(5, 3)) # 8
```

### Importar Todo

```python
# ❌ No recomendado (dudoso qué se importa)
from utilidades import *

# ✅ Mejor: sé específico
from utilidades import saludar, sumar
```

---

## Parte II: Paquetes

Un **paquete** es una carpeta con un `__init__.py`:

```
mi_proyecto/
├── main.py
└── milib/                # Paquete
    ├── __init__.py
    ├── utils.py
    ├── helpers.py
    └── subpaquete/       # Sub-paquete
        ├── __init__.py
        └── core.py
```

### Usar Paquete

```python
# Importar módulo de paquete
from milib import utils
from milib.utils import función

# Importar de sub-paquete
from milib.subpaquete.core import otra_función

# O dentro de milib/__init__.py:
from .utils import función
from .helpers import helper_func
```

---

## Parte III: __init__.py

Marca carpeta como paquete. Puede estar vacío o tener imports:

```python
# milib/__init__.py

# Importar funciones para acceso directo
from .utils import saludar
from .helpers import procesar

# Variables de paquete
__version__ = "1.0"
__author__ = "Tu Nombre"

# Publicidad: qué exporta cuando hacen from milib import *
__all__ = ["saludar", "procesar"]
```

Ahora:

```python
# Acceso directo
from milib import saludar, procesar

# En lugar de
from milib.utils import saludar
from milib.helpers import procesar
```

---

## Parte IV: Librerías Estándar Útiles

### os - Sistema Operativo

```python
import os

print(os.getcwd())              # Directorio actual
os.chdir("/nueva/ruta")
print(os.listdir("."))          # Archivos en carpeta
print(os.path.exists("archivo.txt"))
print(os.path.isdir("carpeta"))
os.makedirs("nueva/carpeta/anidada")
```

### sys - Sistema Python

```python
import sys

print(sys.version)              # Versión Python
print(sys.platform)             # Windows, Linux, etc.
print(sys.argv)                 # Argumentos línea comando
sys.exit(0)                     # Salir programa
print(sys.path)                 # Dónde busca módulos
```

### datetime - Fechas y Horas

```python
from datetime import datetime, timedelta, date

ahora = datetime.now()
print(ahora)                    # 2026-...

hoy = date.today()
print(hoy)

fecha_especifica = datetime(2026, 1, 15)

# Diferencia
diferencia = ahora - fecha_especifica
print(diferencia.days)

# Sumar días
mañana = ahora + timedelta(days=1)
```

### collections - Estructuras Útiles

```python
from collections import defaultdict, Counter, namedtuple

# defaultdict
contador = defaultdict(int)
contador["a"] += 1
contador["a"] += 1
print(contador)  # defaultdict(<class 'int'>, {'a': 2})

# Counter
contar = Counter("aabbcc")
print(contar)   # Counter({'a': 2, 'b': 2, 'c': 2})

# namedtuple
Punto = namedtuple("Punto", ["x", "y"])
p = Punto(10, 20)
print(p.x, p.y)
```

### itertools - Iteraciones Avanzadas

```python
from itertools import combinations, permutations, cycle, repeat

# Combinaciones
for combo in combinations("ABC", 2):
    print(combo)  # ('A', 'B'), ('A', 'C'), ('B', 'C')

# Permutaciones
for perm in permutations("ABC"):
    print(perm)  # ('A', 'B', 'C'), ('A', 'C', 'B'), ...

# Repetir infinito
contador = cycle([1, 2, 3])
print(next(contador))  # 1
print(next(contador))  # 2
print(next(contador))  # 3
print(next(contador))  # 1 (repite)

# Repetir valor
for x in repeat("A", 3):
    print(x)  # A, A, A
```

### functools - Utilidades Funcionales

```python
from functools import reduce, lru_cache

# reduce
números = [1, 2, 3, 4, 5]
producto = reduce(lambda x, y: x * y, números)
print(producto)  # 120

# lru_cache (caché)
@lru_cache(maxsize=32)
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print(fibonacci(100))  # Rápido (cachéado)
```

### math - Funciones Matemáticas

```python
import math

print(math.pi)           # 3.14159...
print(math.e)            # 2.71828...
print(math.sqrt(16))     # 4.0
print(math.sin(math.pi))  # 0.0
print(math.floor(3.7))   # 3
print(math.ceil(3.2))    # 4
print(math.factorial(5)) # 120
```

### random - Números Aleatorios

```python
import random

print(random.random())              # Float 0-1
print(random.randint(1, 10))        # Int 1-10
print(random.choice([1, 2, 3]))     # Elemento aleatorio
print(random.sample([1,2,3,4,5], 3))  # 3 aleatorios sin repetir
aleatorio = [1, 2, 3]
random.shuffle(aleatorio)           # Barajaro in-place
```

### pathlib - Rutas Modernas

```python
from pathlib import Path

# Crear rutas
ruta = Path("Documentos") / "archivo.txt"

# Operaciones
print(ruta.exists())       # Existe
print(ruta.is_file())      # Es archivo
print(ruta.parent)         # Carpeta padre
print(ruta.name)           # Nombre archivo
print(ruta.suffix)         # Extensión

# Crear carpeta
ruta.mkdir(parents=True, exist_ok=True)

# Listar archivos
for archivo in ruta.parent.glob("*.txt"):
    print(archivo)
```

### re - Expresiones Regulares

```python
import re

texto = "email@ejemplo.com"

# Buscar
if re.search(r"\w+@\w+\.\w+", texto):
    print("Email válido")

# Reemplazar
nuevo = re.sub(r"\d+", "X", "Tengo 25 años")  # Tengo X años

# Buscar todos
palabras = re.findall(r"\w+", "Hola Mundo Python")  # ['Hola', 'Mundo', 'Python']
```

---

## Parte V: sys.path - Dónde Busca Python

```python
import sys

# Dónde Python busca módulos
print(sys.path)

# Agregar carpeta
sys.path.append("/ruta/a/módulos")

# Ahora puede importar de esa ruta
import mi_módulo_custom
```

---

## Parte VI: Instalar Paquetes con pip

```bash
# Instalar
pip install requests

# Versión específica
pip install requests==2.28.1

# Múltiples
pip install requests flask numpy

# Desinstalar
pip install uninstall requests

# Listar instalados
pip list

# Guardar requirements
pip freeze > requirements.txt

# Instalar desde requirements
pip install -r requirements.txt
```

---

## Parte VII: Crear tu Propio Paquete

### Estructura

```
mi_paquete/
├── setup.py
├── README.md
├── mi_paquete/
│   ├── __init__.py
│   ├── core.py
│   └── utils.py
```

### setup.py

```python
from setuptools import setup, find_packages

setup(
    name="mi_paquete",
    version="1.0",
    author="Tu Nombre",
    description="Descripción breve",
    packages=find_packages(),
    install_requires=[
        "requests>=2.25.0",
        "numpy>=1.20.0"
    ],
)
```

### Instalar en Desarrollo

```bash
cd mi_paquete
pip install -e .  # -e = editable
```

---

## Resumen de Imports

| Tipo | Código |
|------|--------|
| Módulo entero | `import módulo` |
| Función específica | `from módulo import función` |
| Múltiples | `from módulo import f1, f2` |
| Con alias | `import módulo as m` |
| Todo (evitar) | `from módulo import *` |
| Relativo (paquete) | `from .módulo import función` |

---

**Siguiente**: [bases_datos.md](bases_datos.md)
