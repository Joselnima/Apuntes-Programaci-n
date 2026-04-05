# Estructura Básica de Python

> **Cómo organizar un programa Python: orden correcto y buenas prácticas**

---

## ¿QUÉ es la Estructura?

**Estructura** = el ORDEN en que escribes las cosas:

```python
# ❌ MAL (confuso)
print(x)
x = 5
import math

# ✅ BIEN (claro)
import math
x = 5
print(x)
```

El ORDEN importa para:
- 📖 Legibilidad (otros entienden rápido)
- 🔧 Funcionalidad (menos errores)
- 📚 Mantenibilidad (futuro tú lo agradece)

**Analogía:**
- 📖 Libro desordenado = confuso
- 📚 Libro with índice = fácil de leer

---

## ¿PARA QÉ?

### 1. Evitar errores
```python
# ❌ Error: x no existe aún
print(x)
x = 5

# ✅ Correcto: x ya existe
x = 5
print(x)
```

### 2. Código limpio y profesional
```python
# Orden estándar:
# 1. Importaciones (librerías)
# 2. Constantes (valores fijos)
# 3. Funciones (tareas reutilizables)
# 4. Código principal (que se ejecuta)
```

### 3. Trabajar en equipo
```python
# Todos siguen MISMO orden
# Fácil revisar código de otros
```

---

## ¿CÓMO?

### Orden Estándar Python

```python
"""
Módulo: descripción de qué hace este archivo
"""

# 1. IMPORTACIONES (de librerías)
import os
import math
from datetime import datetime

# 2. CONSTANTES (valores que no cambian)
MÁXIMO_INTENTOS = 3
RUTA_BASE = "/datos"

# 3. FUNCIONES (código reutilizable)
def saludo(nombre):
    return f"Hola {nombre}"

# 4. CLASES (si necesitas POO)
class Persona:
    def __init__(self, nombre):
        self.nombre = nombre

# 5. CÓDIGO PRINCIPAL (se ejecuta)
if __name__ == "__main__":
    print(saludo("Juan"))
    persona = Persona("María")
```

**Por qué `if __name__ == "__main__"`:**
```python
# Permite reutilizar: Si otro archivo importa este,
# el código main NO se ejecuta automáticamente
```

---

## Parte I: El Script Más Simple

### Ejemplo Mínimo

```python
# archivo: hola.py
print("Hola Mundo")
```

Ejecutas:
```bash
python hola.py
# Output: Hola Mundo
```

Eso es todo. Pero en la práctica, los scripts son más organizados.

---

## Parte II: Input y Output

### Input del usuario

```python
# Leer del usuario y guardar en variable
nombre = input("¿Cuál es tu nombre? ")
print(f"Hola {nombre}")
```

**Importante**: `input()` SIEMPRE retorna un **string**, nunca número:

```python
edad = input("¿Cuántos años tienes? ")
print(type(edad))  # <class 'str'>

# Para convertir:
edad = int(input("¿Cuántos años tienes? "))
print(type(edad))  # <class 'int'>
```

### Output con print()

```python
# Básico
print("Hola")

# Múltiples valores
print("Hola", "Mundo", 123)  # Hola Mundo 123

# Con separador personalizado
print("A", "B", "C", sep="-")  # A-B-C

# Sin salto de línea
print("Hola", end="")
print(" Mundo")  # HolaMundo (en la misma línea)

# F-strings (Python 3.6+) - Recomendado
nombre = "Juan"
edad = 25
print(f"Me llamo {nombre} y tengo {edad} años")

# Formato antiguo (aún funciona)
print("Me llamo {} y tengo {} años".format(nombre, edad))
```

---

## Parte III: Estructura de un Script Real

Este es el patrón que seguirás en casi TODO:

```python
# -*- coding: utf-8 -*-
"""
módulo_ejemplo.py

Descripción breve de qué hace este módulo.
Puede ser de varias líneas si es necesario.

Autor: Tu Nombre
Fecha: 2026
"""

# ============================================================================
# IMPORTACIONES
# ============================================================================
import os
import sys
from pathlib import Path
from datetime import datetime

# Librerías de terceros
import requests


# ============================================================================
# CONSTANTES
# ============================================================================
VERSIÓN = "1.0"
DIRECTORIO_BASE = Path(__file__).parent
RUTA_DATOS = DIRECTORIO_BASE / "datos"


# ============================================================================
# FUNCIONES
# ============================================================================
def saludar(nombre):
    """
    Saluda a una persona por su nombre.
    
    Args:
        nombre (str): Nombre de la persona
    
    Returns:
        str: Saludo personalizado
    
    Ejemplo:
        >>> saludar("Juan")
        'Hola Juan'
    """
    return f"Hola {nombre}"


def main():
    """Función principal del programa."""
    print("Iniciando programa...")
    
    # Código del programa aquí
    nombre = input("¿Tu nombre? ")
    resultado = saludar(nombre)
    print(resultado)


# ============================================================================
# CLASES (si las necesitas)
# ============================================================================
class Persona:
    """Representa una persona."""
    
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad
    
    def presentarse(self):
        """Devuelve una presentación."""
        return f"Soy {self.nombre} y tengo {self.edad} años"


# ============================================================================
# PUNTO DE ENTRADA
# ============================================================================
if __name__ == "__main__":
    main()
```

---

## Parte IV: El Mágico `if __name__ == "__main__":`

Este es probablemente el patrón más importante en Python.

### ¿Qué significa?

```python
# archivo: ejemplo.py

def saludar():
    print("Hola desde saludar()")

print("Línea 1: Importando módulo")
saludar()
```

**Si importas este módulo desde otro archivo**:

```python
# archivo: otro.py
import ejemplo

# Output:
# Línea 1: Importando módulo
# Hola desde saludar()
```

¡El código se ejecutó automáticamente! Esto es un problema si solo querías importar:

```python
def saludar():
    print("Hola desde saludar()")
    # Esto SÍ se ejecuta cuando importas

# ¿Pero esto solo ejecuta si corro el archivo directamente?
print("Hola Mundo")
```

### La Solución: `if __name__ == "__main__":`

```python
# archivo: ejemplo.py

def saludar():
    print("Hola desde saludar()")

if __name__ == "__main__":
    print("Este código SOLO se ejecuta si corro este archivo directamente")
    print("Si alguien lo importa, isto NO se ejecuta")
    saludar()
```

**Ahora**:
- `python ejemplo.py` → Imprime todo
- `import ejemplo` desde otro archivo → No imprime nada, solo puede usar saludar()

### ¿Por qué funciona?

Python asigna el nombre `__main__` SOLO al archivo que ejecutas directamente. Otros archivos tienen nombres diferentes (`__nombre_modulo__`).

---

## Parte V: Importaciones

### Importar un módulo completo

```python
import math

print(math.pi)      # 3.14159...
print(math.sqrt(9)) # 3
```

### Importar funciones específicas

```python
from math import sqrt, pi

print(pi)    # Acceso directo
print(sqrt(9))
```

### Importar con alias

```python
import numpy as np
from pandas import DataFrame as DF

datos = np.array([1, 2, 3])
df = DF(datos)
```

### Orden recomendado de importaciones

1. Módulos estándar
2. Módulos de terceros
3. Módulos locales

```python
# 1. Estándar
import os
import sys
from pathlib import Path

# 2. Terceros
import numpy as np
import requests

# 3. Locales
from mi_modulo import mi_funcion
```

---

## Parte VI: Comentarios y Docstrings

### Comentarios (con #)

```python
# Esto es un comentario de una línea

# Esto es un comentario
# de múltiples líneas
# que explica algo complejo
x = 5  # Comentario al final de una línea

# ¡Evita comentarios obvios!
# ❌ MAL
x = 5  # Asignar 5 a x

# ✅ BIEN
x = 5  # Cantidad de intentos reintentos permitidos
```

### Docstrings (Triple comilla)

Los docstrings documentan funciones/clases para usuarios:

```python
def multiplicar(a, b):
    """
    Multiplica dos números.
    
    Args:
        a: Primer número
        b: Segundo número
    
    Returns:
        El producto a * b
    
    Ejemplo:
        >>> multiplicar(3, 4)
        12
    """
    return a * b

# Acceso a docstring:
print(multiplicar.__doc__)
help(multiplicar)
```

---

## Parte VII: Argumentos de Línea de Comandos

### Acceso con sys.argv

```python
# archivo: script.py
import sys

print(sys.argv)
```

```bash
python script.py hola mundo 123
# Output: ['script.py', 'hola', 'mundo', '123']
```

**Importante**: El primer elemento (index 0) siempre es el nombre del script.

```python
import sys

if len(sys.argv) > 1:
    nombre = sys.argv[1]
    print(f"Hola {nombre}")
else:
    print("Uso: python script.py <nombre>")
```

Utilizar:
```bash
python script.py Juan
# Output: Hola Juan
```

### Librería argparse (Más profesional)

```python
import argparse

parser = argparse.ArgumentParser(description="Mi programa")
parser.add_argument("nombre", help="Tu nombre")
parser.add_argument("--edad", type=int, help="Tu edad (opcional)")
parser.add_argument("--verbose", action="store_true", help="Modo verbose")

args = parser.parse_args()

print(f"Nombre: {args.nombre}")
if args.edad:
    print(f"Edad: {args.edad}")
if args.verbose:
    print("Modo verbose activado")
```

Utilizar:
```bash
python script.py Juan --edad 25 --verbose
```

---

## Parte VIII: Estructura de Directorios para Proyectos

### Proyecto simple

```
mi_proyecto/
├── venv/                # Entorno virtual
├── main.py             # Archivo principal
├── utils.py            # Funciones útiles
└── README.md
```

### Proyecto mediano

```
mi_proyecto/
├── venv/
├── src/
│   ├── __init__.py      # Marca como paquete
│   ├── main.py
│   ├── utils.py
│   └── modelos.py
├── tests/
│   ├── __init__.py
│   ├── test_utils.py
│   └── test_modelos.py
├── data/
│   └── muestra.csv
├── requirements.txt     # Paquetes necesarios
└── README.md
```

---

## Parte IX: Paquetes y Módulos

### Crear un paquete

Un **paquete** es una carpeta con un `__init__.py`:

```
mi_paquete/
├── __init__.py          # Marca como paquete
├── modulo1.py
├── modulo2.py
└── subpaquete/
    ├── __init__.py
    └── modulo3.py
```

### Usar el paquete

```python
# Importar módulo
from mi_paquete import modulo1
modulo1.mi_funcion()

# Importar función específica
from mi_paquete.modulo1 import mi_funcion
mi_funcion()

# Importar subpaquete
from mi_paquete.subpaquete import modulo3
```

---

## Parte X: El archivo `__init__.py`

El `__init__.py` puede estar vacío o contener código de inicialización:

```python
# mi_paquete/__init__.py

# Importa funciones para acceso directo
from .modulo1 import mi_funcion
from .modulo2 import otra_funcion

# Versión del paquete
__version__ = "1.0.0"

# Variables públicas
__all__ = ["mi_funcion", "otra_funcion"]
```

Ahora otros pueden hacer:

```python
from mi_paquete import mi_funcion
```

---

## Parte XI: Ejecución de Módulos como Scripts

Puedes ejecutar un módulo dentro de un paquete:

```bash
python -m mi_paquete.modulo1
```

Esto busca `if __name__ == "__main__":` dentro de ese módulo.

---

## Resumen

| Concepto | Uso |
|----------|-----|
| `print()` | Salida a pantalla |
| `input()` | Entrada del usuario |
| `if __name__ == "__main__":` | Código que solo corre si ejecutas este archivo |
| Importaciones | Reutilizar código |
| Docstrings | Documentación de funciones |
| `sys.argv` | Argumentos de línea de comandos |
| Paquetes | Organización de código en carpetas |

---

**Siguiente**: [variables.md](variables.md)
