# Stdlib Útiles - Módulos Importantes

> **Herramientas poderosas que Python ya trae - no necesitas instalar**

---

## ¿QUÉ es la Stdlib (Standard Library)?

La **stdlib** = caja de herramientas de Python:

```python
import sys
print(len(sys.modules))  # ~200+ módulos listos

# Sin instalar NADA (no necesitas pip)
import json
import random
import datetime
# Todas disponibles
```

**Analogía:**
- 📦 Npm/pip = comprar herramientas nuevas (internet)
- 🧰 Stdlib = herramientas que ya tienes (sin conexión)

---

## ¿PARA QÉ sirve?

### 1. Contar ocurrencias (Counter)
```python
votos = "Juan Juan María Pedro Juan".split()

# ❌ Manual (tedioso)
frecuencia = {}
for voto in votos:
    if voto not in frecuencia:
        frecuencia[voto] = 0
    frecuencia[voto] += 1

# ✅ Con Counter (una línea)
from collections import Counter
frecuencia = Counter(votos)
# {'Juan': 3, 'María': 1, 'Pedro': 1}
```

### 2. Diccionarios inteligentes (defaultdict)
```python
# ❌ Manual
grupos = {}
for país, persona in [("Chile", "Juan"), ("Chile", "María"), ("Perú", "Pedro")]:
    if país not in grupos:
        grupos[país] = []
    grupos[país].append(persona)

# ✅ Con defaultdict (automático)
from collections import defaultdict
grupos = defaultdict(list)
for país, persona in [("Chile", "Juan"), ("Chile", "María"), ("Perú", "Pedro")]:
    grupos[país].append(persona)
# {'Chile': ['Juan', 'María'], 'Perú': ['Pedro']}
```

### 3. Tuplas con etiquetas (namedtuple)
```python
Persona = namedtuple('Persona', ['nombre', 'edad', 'ciudad'])
juan = Persona("Juan", 25, "Santiago")

print(juan.nombre)  # "Juan" (más claro que juan[0])
```

### 4. Colas eficientes (deque)
```python
from collections import deque

# Lista normal: lento en inicio
lista = [1, 2, 3]
lista.pop(0)  # O(n) - caro

# deque: rápido en ambos lados
cola = deque([1, 2, 3])
cola.popleft()  # O(1) - barato
```

---

## ¿CÓMO funcionan?

### 5 Módulos Principales

| Módulo | Para qué | Ejemplo |
|--------|----------|---------|
| **collections** | Estructuras especiales | Counter, defaultdict |
| **itertools** | Permutaciones/combinaciones | permutations, combinations |
| **functools** | Funcional programming | reduce, lru_cache |
| **datetime** | Fechas y horas | datetime.now() |
| **random** | Números aleatorios | random.choice([1,2,3]) |

---

## Parte I: collections - Estructuras Especializadas

```python
from collections import Counter, defaultdict, namedtuple, deque

# Counter - contar ocurrencias
palabras = "la la la gata gata gato".split()
contador = Counter(palabras)
print(contador)
# Counter({'la': 3, 'gata': 2, 'gato': 1})

print(contador.most_common(2))
# [('la', 3), ('gata', 2)]

# defaultdict - valor por defecto
d = defaultdict(list)
d['usuarios'].append('Juan')
d['usuarios'].append('María')
print(d)
# defaultdict(<class 'list'>, {'usuarios': ['Juan', 'María']})

# namedtuple - tupla con nombres
Punto = namedtuple('Punto', ['x', 'y'])
p = Punto(3, 4)
print(p.x, p.y)  # 3 4

# deque - cola doble (O(1) en ambos extremos)
cola = deque([1, 2, 3])
cola.appendleft(0)
cola.append(4)
print(cola)  # deque([0, 1, 2, 3, 4])
x = cola.popleft()  # 0 (eficiente)
```

---

## Parte II: itertools - Combinaciones y Permutaciones

```python
import itertools

# permutaciones (orden importa)
print(list(itertools.permutations('AB')))
# [('A', 'B'), ('B', 'A')]

# combinaciones (orden no importa)
print(list(itertools.combinations('ABC', 2)))
# [('A', 'B'), ('A', 'C'), ('B', 'C')]

# producto cartesiano
print(list(itertools.product('AB', '12')))
# [('A', '1'), ('A', '2'), ('B', '1'), ('B', '2')]

# Encadenar iterables
it = itertools.chain([1, 2], [3, 4], [5, 6])
print(list(it))  # [1, 2, 3, 4, 5, 6]

# Repetir
print(list(itertools.repeat('x', 3)))  # ['x', 'x', 'x']

# Ciclos infinitos (cuidado)
contador = itertools.count(10, 2)
print(next(contador))  # 10
print(next(contador))  # 12
print(next(contador))  # 14

# Tomar primeros N
primeros_5 = itertools.islice(range(100), 5)
print(list(primeros_5))  # [0, 1, 2, 3, 4]

# Agrupar
datos = [1, 1, 1, 2, 2, 3]
for valor, grupo in itertools.groupby(datos):
    print(valor, list(grupo))
# 1 [1, 1, 1]
# 2 [2, 2]
# 3 [3]
```

---

## Parte III: functools - Utilidades Funcionales

```python
from functools import reduce, lru_cache, wraps

# reduce - acelumular
numeros = [1, 2, 3, 4, 5]
producto = reduce(lambda x, y: x * y, numeros)
print(producto)  # 120

# lru_cache - caché automático
@lru_cache(maxsize=128)
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print(fibonacci(30))  # Muy rápido (caché)

# partial - fijar argumentos
from functools import partial

def potencia(base, exponente):
    return base ** exponente

cuadrado = partial(potencia, exponente=2)
print(cuadrado(5))  # 25
print(cuadrado(3))  # 9

# wraps - preservar metadatos
def mi_decorador(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)
    return wrapper
```

---

## Parte IV: operator - Operaciones Como Funciones

```python
from operator import add, mul, itemgetter, attrgetter

# Operadores como funciones
print(add(2, 3))        # 5
print(mul(4, 5))        # 20

# Extractores - data structures
personas = [
    {'nombre': 'Juan', 'edad': 25},
    {'nombre': 'María', 'edad': 30},
]

nombres = list(map(itemgetter('nombre'), personas))
print(nombres)  # ['Juan', 'María']

# Ordenar por atributo
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad

personas = [
    Persona('Juan', 25),
    Persona('María', 30),
    Persona('Pedro', 20),
]

ordenado = sorted(personas, key=attrgetter('edad'))
for p in ordenado:
    print(f"{p.nombre}: {p.edad}")
```

---

## Parte V: datetime - Fechas y Horas

```python
from datetime import datetime, timedelta, date

# Ahora
ahora = datetime.now()
print(ahora)  # 2024-01-15 14:30:45.123456

# Crear fecha y hora específica
fecha = datetime(2024, 1, 15, 14, 30, 0)
print(fecha)

# Solo fecha
hoy = date.today()
print(hoy)  # 2024-01-15

# Operaciones
mañana = date.today() + timedelta(days=1)
semana_pasada = date.today() - timedelta(weeks=1)

# Diferencia
ahora_inicio = datetime(2024, 1, 1)
ahora_final = datetime.now()
diferencia = ahora_final - ahora_inicio
print(diferencia.days)  # Días transcurridos

# Formato
fecha = datetime.now()
print(fecha.strftime("%d/%m/%Y %H:%M:%S"))
# 15/01/2024 14:30:45

# Parsear
fecha_str = "15/01/2024"
fecha = datetime.strptime(fecha_str, "%d/%m/%Y")
print(fecha)  # 2024-01-15 00:00:00
```

---

## Parte VI: math - Operaciones Matemáticas

```python
import math

# Constantes
print(math.pi)      # 3.14159...
print(math.e)       # 2.71828...

# Funciones
print(math.sqrt(16))        # 4.0
print(math.pow(2, 8))       # 256.0
print(math.ceil(3.2))       # 4 (techo)
print(math.floor(3.9))      # 3 (piso)
print(math.fabs(-5))        # 5.0

# Trigonometría
print(math.sin(math.pi/2))  # 1.0
print(math.cos(0))          # 1.0

# Logaritmos
print(math.log(math.e))     # 1.0
print(math.log10(100))      # 2.0

# GCD y LCM
print(math.gcd(12, 8))      # 4
print(math.gcd(30, 45))     # 15

# Factorial
print(math.factorial(5))    # 120
```

---

## Parte VII: random - Números Aleatorios

```python
import random

# Entero aleatorio
print(random.randint(1, 10))    # Entre 1 y 10
print(random.random())          # Float entre 0 y 1

# Uniforme
print(random.uniform(1, 5))     # Float entre 1 y 5

# Elección
colores = ['rojo', 'verde', 'azul']
print(random.choice(colores))   # Un color aleatorio

# Múltiples sin repetición
cartas = list(range(1, 53))
mano = random.sample(cartas, 5)  # 5 cartas distintas
print(mano)

# Barajar
lista = [1, 2, 3, 4, 5]
random.shuffle(lista)
print(lista)  # Orden aleatorio

# Distribuciones
print(random.gauss(0, 1))       # Distribución normal
print(random.expovariate(1))    # Distribución exponencial

# Reproducible (seed)
random.seed(42)
print(random.randint(1, 100))
random.seed(42)
print(random.randint(1, 100))   # Mismo número
```

---

## Parte VIII: re - Expresiones Regulares

```python
import re

# Búsqueda simple
texto = "El email es juan@example.com"
si_está = "juan" in texto

# Búsqueda con regex
patrón = r'\b\w+@\w+\.\w+\b'
email = re.search(patrón, texto)
if email:
    print(email.group())  # juan@example.com

# Encontrar todos
texto = "Emails: juan@test.com y maría@test.com"
emails = re.findall(r'\b\w+@\w+\.\w+\b', texto)
print(emails)  # ['juan@test.com', 'maría@test.com']

# Reemplazo
nuevo_texto = re.sub(r'\d{3}-\d{4}', 'XXXX-XXXX', "Teléfono: 555-1234")
print(nuevo_texto)  # Teléfono: XXXX-XXXX

# Dividir
partes = re.split(r',\s*', "Juan, María, Pedro")
print(partes)  # ['Juan', 'María', 'Pedro']

# Grupos
patrón = r'(\w+)@(\w+)'
match = re.search(patrón, "juan@example.com")
print(match.group(1))  # juan
print(match.group(2))  # example
```

---

## Parte IX: pathlib - Manejo de Rutas

```python
from pathlib import Path

# Ruta actual
p = Path('.')
print(p.resolve())  # Ruta absoluta

# Crear ruta
archivo = Path('datos') / 'usuarios.csv'
print(archivo)  # datos/usuarios.csv

# Propiedades
print(archivo.name)          # usuarios.csv
print(archivo.stem)          # usuarios
print(archivo.suffix)        # .csv
print(archivo.parent)        # datos

# Operaciones
print(archivo.exists())      # ¿Existe?
print(archivo.is_file())     # ¿Es archivo?
print(archivo.is_dir())      # ¿Es directorio?

# Listar archivos
for archivo in Path('.').glob('*.py'):
    print(archivo)

# Leer/escribir
p = Path('archivo.txt')
p.write_text('Hola mundo')
contenido = p.read_text()
print(contenido)

# Permisos
p.chmod(0o755)
```

---

## Parte X: shutil - Operaciones Archivo

```python
import shutil
from pathlib import Path

# Copiar archivo
shutil.copy('origen.txt', 'copia.txt')

# Copiar árbol
shutil.copytree('carpeta_origen', 'carpeta_copia')

# Mover/renombrar
shutil.move('archivo.txt', 'nueva_ubicación.txt')

# Eliminar árbol
shutil.rmtree('carpeta')

# Espacio disco
espacio = shutil.disk_usage('.')
print(f"Total: {espacio.total}")
print(f"Libre: {espacio.free}")
```

---

## Parte XI: json - JSON Serialización

```python
import json

# Codificar
datos = {
    'nombre': 'Juan',
    'edad': 25,
    'hobbies': ['leer', 'programar'],
}

json_str = json.dumps(datos, indent=2)
print(json_str)

# Guardar a archivo
with open('datos.json', 'w') as f:
    json.dump(datos, f, indent=2)

# Decodificar
json_str = '{"nombre": "Juan", "edad": 25}'
datos = json.loads(json_str)
print(datos)

# Cargar de archivo
with open('datos.json') as f:
    datos = json.load(f)
```

---

## Parte XII: logging - Registrar Eventos

```python
import logging

# Configurar
logging.basicConfig(
    level=logging.DEBUG,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    filename='app.log'
)

logger = logging.getLogger(__name__)

# Verbosidad (de menor a mayor)
logger.debug("Mensaje de depuración")
logger.info("Información")
logger.warning("Advertencia")
logger.error("Error")
logger.critical("Error crítico")

# Ejemplo
try:
    resultado = 10 / 0
except ZeroDivisionError as e:
    logger.error(f"Error aritmético: {e}")
```

---

## Resumen

| Módulo | Uso |
|--------|-----|
| collections | Counter, defaultdict, namedtuple |
| itertools | Permutaciones, combinaciones |
| functools | reduce, lru_cache, partial |
| operator | Operadores como funciones |
| datetime | Fechas y horas |
| math | Operaciones matemáticas |
| random | Números aleatorios |
| re | Expresiones regulares |
| pathlib | Manejo de rutas |
| json | Serialización JSON |

---

**Siguiente**: [regex.md](regex.md)
