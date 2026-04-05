# Buenas Prácticas en Python

> **Escribir código que otros (y tú futuro) puedan entender**

---

## ¿QUÉ son Buenas Prácticas?

Son **reglas no oficiales** que programadores respetan:

```python
# ❌ Código cabreado (funcionaaa pero... ufff)
def f(x,y,z):
 return x+y*z if x else y/z

# ✅ Código feliz (entiende rápido)
def calcular_costo_total(base, cantidad, descuento):
    """Calcula el costo final aplicando descuento"""
    total = base * cantidad
    return total - descuento
```

**Analogía:**
- 🏚️ Código sin prácticas = casa desordenada
- 🏠 Código con prácticas = casa organizada y limpia

---

## ¿PARA QÉ sirven?

### 1. Entender código después de 6 meses
```python
# Tu futuro: "¿Qué hace esto? ¿Por qué?"
x = [i for i in data if i[2] > 3.5]

# Con buenas prácticas:
estudiantes_sobresalientes = [
    est for est in estudiantes 
    if est['promedio'] > 3.5
]
# Ahora es OBVIO
```

### 2. Trabajo en equipo
```python
# El compañero entiende tu código
# Sin discusiones de "¿qué significa esto?"
```

### 3. Evitar bugs
```python
# Código CLARO = errores fáciles de ver
# Código CONFUSO = bugs ocultos

# "Errors should never pass silently" - Python
```

### 4. Código más eficiente
```python
# Malas prácticas = código lento
# Buenas prácticas = código rápido
```

---

## ¿CÓMO aplicarlas?

### 3 Pilares

| Pilar | Qué es | Beneficio |
|-------|--------|-----------|
| **PEP 8** | Formato de código | Consistencia visual |
| **Naming** | Nombres claros | Legibilidad |
| **Documentation** | Comentarios efectivos | Mantenimiento |

**Resumen rápido:**
- **Rojo**: Código que no sigue PEP 8
- **Amarillo**: Código que funciona pero confuso
- **Verde**: Código limpio y legible

---

## Parte I: PEP 8 - Estilo de Código

### Nombrado

```python
# ✅ BIEN
usuario_activo = True
calcular_promedio()
MÁXIMO_INTENTOS = 3
class Persona:
    pass

# ❌ MAL
usuarioActivo = True        # camelCase en Python no
calcularPromedio()
maxIntents = 3
class persona:              # Clases en PascalCase
    pass
```

### Espacios en Blanco

```python
# ✅ BIEN
x = 5
if x > 3:
    print("Mayor")

diccionario = {"nombre": "Juan", "edad": 25}

# ❌ MAL
x=5
if x>3:
    print("Mayor")

diccionario={"nombre":"Juan","edad":25}

# ✅ Máximo 79 caracteres por línea
# Partir si es necesario
resultado = sumar(variable_larga1, variable_larga2,
                  variable_larga3, variable_larga4)

# ✅ Dos líneas entre funciones de nivel módulo
def función1():
    pass


def función2():
    pass
```

### Imports

```python
# ✅ BIEN
import os
import sys
from pathlib import Path
import numpy as np
from mi_módulo import función

# ❌ MAL
import os, sys  # No múltiples en una línea
from * import  # Nunca esto
```

---

## Parte II: Nombres Significativos

### Variables

```python
# ❌ MAL - confuso
x = 5
y = 10
calc(x, y)

# ✅ BIEN - claro
base = 5
altura = 10
calcular_área(base, altura)
```

### Funciones

```python
# ❌ MAL
def proc():
    pass

def calc():
    pass

# ✅ BIEN
def procesar_datos():
    pass

def calcular_promedio():
    pass

# Verbos (qué hace)
def obtener_usuario():
    pass

def crear_conexión():
    pass

def es_válido():  # Pregunta → bool
    pass

def validar_email():
    pass
```

### Clases

```python
# ❌ MAL
class proc:
    pass

# ✅ BIEN
class ProcesadorDatos:
    pass

class Usuario:
    pass

class ConexiónBD:
    pass
```

---

## Parte III: Código Limpio (DRY)

### DRY - Don't Repeat Yourself

```python
# ❌ MAL - código repetido
def procesar_usuario():
    usuario = obtener_usuario()
    validar(usuario)
    print(usuario.nombre)
    guardar(usuario)

def procesar_producto():
    producto = obtener_producto()
    validar(producto)
    print(producto.nombre)
    guardar(producto)

# ✅ BIEN - abstracto
def procesar(objeto):
    validar(objeto)
    print(objeto.nombre)
    guardar(objeto)

procesar(obtener_usuario())
procesar(obtener_producto())
```

---

## Parte IV: Funciones Pequeñas y Enfocadas

```python
# ❌ MAL - hace demasiado
def procesar_archivo_completo(archivo):
    # Leer archivo (5 líneas)
    with open(archivo) as f:
        datos = f.read()
    
    # Procesar (10 líneas)
    for línea in datos.split("\n"):
        # procesamiento complejo
        pass
    
    # Guardar (5 líneas)
    with open("salida.txt", "w") as f:
        f.write(resultado)

# ✅ BIEN - responsabilidad única
def leer_archivo(archivo):
    with open(archivo) as f:
        return f.read()

def procesar_datos(datos):
    resultado = []
    for línea in datos.split("\n"):
        resultado.append(procesar_línea(línea))
    return resultado

def guardar_resultado(resultado, archivo):
    with open(archivo, "w") as f:
        f.write(resultado)

# Orquestar
datos = leer_archivo("entrada.txt")
procesado = procesar_datos(datos)
guardar_resultado(procesado, "salida.txt")
```

---

## Parte V: Principales SOLID

### S - Single Responsibility

Cada clase/función hace UNA cosa:

```python
# ❌ MAL - múltiples responsabilidades
class Usuario:
    def __init__(self, nombre):
        self.nombre = nombre
    
    def validar_email(self):
        # validación
        pass
    
    def enviar_email(self):
        # enviar
        pass
    
    def guardar_en_bd(self):
        # BD
        pass

# ✅ BIEN - separado
class Usuario:
    def __init__(self, nombre, email):
        self.nombre = nombre
        self.email = email

class ValidadorEmail:
    def validar(self, email):
        pass

class GestorEmail:
    def enviar(self, email):
        pass

class RepositorioUsuario:
    def guardar(self, usuario):
        pass
```

### O - Open/Closed

Abierto para extensión, cerrado para modificación:

```python
# ❌ MAL - debe modificar function
def calcular_descuento(tipo_cliente, monto):
    if tipo_cliente == "regular":
        return monto * 0.05
    elif tipo_cliente == "vip":
        return monto * 0.10
    elif tipo_cliente == "premium":
        return monto * 0.15
    # Cada tipo nuevo = modificar función

# ✅ BIEN - extensible
class EstrategiaDescuento:
    def calcular(self, monto):
        raise NotImplementedError

class DescuentoRegular(EstrategiaDescuento):
    def calcular(self, monto):
        return monto * 0.05

class DescuentoVIP(EstrategiaDescuento):
    def calcular(self, monto):
        return monto * 0.10

# Nuevos tipos = nueva clase (sin modificar)
class DescuentoPremium(EstrategiaDescuento):
    def calcular(self, monto):
        return monto * 0.15
```

---

## Parte VI: Documentación

### Docstrings

```python
def calcular_promedio(calificaciones):
    \"\"\"
    Calcula el promedio de calificaciones.
    
    Args:
        calificaciones (list): Lista de números (0-100)
    
    Returns:
        float: El promedio
    
    Raises:
        ValueError: Si lista está vacía
    
    Ejemplo:
        >>> calcular_promedio([80, 90, 100])
        90.0
    \"\"\"
    if not calificaciones:
        raise ValueError("Lista no puede estar vacía")
    return sum(calificaciones) / len(calificaciones)
```

### Comentarios

```python
# ✅ BIEN - por qué, no qué
x = usuario.edad  # Edad actual en años
descuento = 0.15  # 15% para clientes premium

# ❌ MAL - obvio
x = usuario.edad  # Asignar edad a x
descuento = 0.15  # Descuento del 15%

# ✅ Para código complejo
# Algoritmo KMP para búsqueda de patrón
# Ver: https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
def construir_tabla_fallos(patrón):
    pass
```

---

## Parte VII: Manejo de Errores

```python
# ❌ MAL - captura genérica
try:
    hacer_algo()
except:
    pass

# ✅ BIEN - específico
try:
    hacer_algo()
except FileNotFoundError:
    print("Archivo no encontrado")
except TypeError as e:
    print(f"Error de tipo: {e}")
except Exception as e:
    print(f"Error inesperado: {e}")
    raise  # Re-lanzar si es crítico
```

---

## Parte VIII: Type Hints

Indicar tipos esperados:

```python
# Sin type hints
def procesar(datos):
    return len(datos)

# Con type hints
def procesar(datos: list) -> int:
    return len(datos)

# Complejos
from typing import List, Dict, Optional, Tuple

def buscar_usuario(id: int) -> Optional[Dict[str, str]]:
    \"\"\"Retorna dict o None\"\"\"
    pass

def crear_reporte(usuarios: List[dict]) -> Tuple[str, int]:
    \"\"\"Retorna (contenido, cantidad)\"\"\"
    pass
```

---

## Parte IX: Testing y Debugging

```python
# Sempre test código importante
def calcular_área_círculo(radio):
    import math
    return math.pi * radio ** 2

# Tests
assert calcular_área_círculo(0) == 0
assert 31.4 < calcular_área_círculo(3.18) < 31.5
```

---

## Parte X: Performance

```python
# ❌ MAL - O(n²)
resultado = []
for x in lista1:
    for y in lista2:
        if x == y:
            resultado.append(x)

# ✅ BIEN - O(n)
resultado = [x for x in lista1 if x in set(lista2)]

# O mejor con set
resultado = list(set(lista1) & set(lista2))
```

---

## Parte XI: Versionamiento

```python
# Semántico: MAYOR.MENOR.PARCHE
__version__ = "1.2.3"

# MAYOR = cambios incompatibles
# MENOR = nueva funcionalidad (compatible)
# PARCHE = bugfixes
```

---

## Parte XII: Checklist de Código

Antes de enviar:

- [ ] Código ejecuta sin errores
- [ ] Tests pasan
- [ ] PEP 8 compliance (`pip install flake8`)
- [ ] Docstrings en funciones principales
- [ ] Nombres significativos
- [ ] Sin código repetido (DRY)
- [ ] Funciones pequeñas (< 20 líneas idealmente)
- [ ] Manejo de errores apropiado
- [ ] Type hints (al menos en APIs públicas)
- [ ] Sin variables globales innecesarias

---

## Resumen

| Principio | Aplicar |
|-----------|---------|
| PEP 8 | Formato y estilo |
| Nombres claros | Variables, funciones, clases |
| DRY | No repetir código |
| SOLID | Diseño de clases |
| Funciones pequeñas | Responsabilidad única |
| Docstrings | Documentar qué hace |
| Type hints | Indicar tipos |
| Tests | Validar correctitud |

---

**Siguiente**: [testing.md](testing.md)
