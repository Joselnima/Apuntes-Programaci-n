# Generadores y Corrutinas

> **Funciones que PAUSAN y REANUDAN - perfectas para datos enormes**

---

## ¿QUÉ es un Generador?

Un **generador** = función que PAUSA y ESPERA:

```python
def contador():
    yield 1       # Pausa 1
    yield 2       # Pausa 2
    yield 3       # Pausa 3

# Uso
for num in contador():
    print(num)    # Imprime de a uno: 1, 2, 3

# Sin generador: imprime TODO a la vez
def contador_normal():
    return [1, 2, 3]
```

**Analogía:**
- 📺 Función normal = películas descargadas (todo a la vez)
- 🎬 Generador = streaming (cada segundo que necesites)

La palabra clave: **`yield`** = "pausa y devuelve esto, espera llamada siguiente"

---

## ¿PARA QÉ sirven?

### 1. Datos enormes SIN consumir RAM
```python
# ❌ Malo: carga 1 millón en memoria
def leer_números(n):
    return list(range(n))

x = leer_números(1_000_000)  # 💾 GIGABYTES en RAM

# ✅ Bueno: genera bajo demanda
def leer_números(n):
    for i in range(n):
        yield i

for num in leer_números(1_000_000):
    procesar(num)  # Un número por véz, RAM constante
```

### 2. Streams de datos
```python
def leer_archivo_línea(nombre_archivo):
    with open(nombre_archivo) as f:
        for línea in f:
            yield línea.strip()

# Procesa archivo de 10GB línea por línea
for línea in leer_archivo_línea("datos.txt"):
    print(línea)
```

### 3. Secuencias infinitas
```python
def números_infinitos():
    i = 0
    while True:
        yield i
        i += 1

contador = números_infinitos()
print(next(contador))  # 0
print(next(contador))  # 1
print(next(contador))  # 2
# Puedo pedir tantos como quiera
```

### 4. Tuberías de datos
```python
# Transformación paso a paso, sin cargar todo
def elevar_al_cuadrado(números):
    for n in números:
        yield n ** 2

for resultado in elevar_al_cuadrado([1, 2, 3, 4, 5]):
    print(resultado)  # 1, 4, 9, 16, 25
```

---

## ¿CÓMO funcionan?

### Máquina de Estados - Pausa y Reanuda

```python
def demo():
    print("1. Inicio")
    yield "A"           # ← Pausa 1
    print("2. Reanudé")
    yield "B"           # ← Pausa 2
    print("3. Otro más")
    yield "C"           # ← Pausa 3
    print("4. Listo")

gen = demo()
print(next(gen))  # OUTPUT: "1. Inicio" → "A"
#                 Se pausa en yield "A"
print(next(gen))  # OUTPUT: "2. Reanudé" → "B"
#                 Reanuda desde yield "A", continúa hasta "B"
print(next(gen))  # OUTPUT: "3. Otro más" → "C"
print(next(gen))  # OUTPUT: "4. Listo" → StopIteration
```

**Estados del Generador:**
```
┌─────────────┐
│   Creado    │  gen = demo()
└──────┬──────┘
       ↓
┌─────────────┐
│  Suspendido │  Espera next()
└──────┬──────┘
       ↓
┌─────────────┐
│   Ejecutando│  Ejecuta hasta yield
└──────┬──────┘
       ↓
┌─────────────┐
│   Pausado   │  Devuelve valor, espera
└──────┬──────┘
       ↓
┌─────────────┐
│  Finalizado │  No hay más yields (StopIteration)
└─────────────┘
```

---

## Parte I: Básico - yield

```python
def contador(inicio, fin):
    \"\"\"Genera números del inicio al fin\"\"\"
    current = inicio
    while current <= fin:
        yield current  # Pausa aquí
        current += 1

# Uso
for num in contador(1, 5):
    print(num)  # 1 2 3 4 5

# Generadores son iterables
gen = contador(1, 5)
print(next(gen))  # 1
print(next(gen))  # 2
print(next(gen))  # 3
# next(gen) lanzará StopIteration al final
```

---

## Parte II: Estados del Generador

```python
def generador():
    print("Inicio")
    yield 1  # Pausa 1
    print("Después de 1")
    yield 2  # Pausa 2
    print("Después de 2")
    yield 3  # Pausa 3
    print("Fin")

gen = generador()

# Ejecución
print(next(gen))
# Salida: Inicio, output: 1

print(next(gen))
# Salida: Después de 1, output: 2

print(next(gen))
# Salida: Después de 2, output: 3

try:
    next(gen)  # Ejecuta "Fin", luego StopIteration
except StopIteration:
    print("Generador exhausto")
```

---

## Parte III: Expresiones Generadoras

Forma compacta:

```python
# ✅ BIEN - expresión generadora
gen = (x**2 for x in range(10))
print(next(gen))  # 0
print(next(gen))  # 1
print(next(gen))  # 4

# vs lista (genera todo)
lista = [x**2 for x in range(10)]  # Completa en RAM

# Memoria
import sys
print(sys.getsizeof(gen))         # ~128 bytes (vacío)
print(sys.getsizeof(lista))       # ~88 bytes + contenido (mucho más)
```

---

## Parte IV: Generador con Lógica

```python
def filtrar_pares(números):
    \"\"\"Solo genera pares\"\"\"
    for num in números:
        if num % 2 == 0:
            yield num

# Uso
enteros = range(10)
pares = filtrar_pares(enteros)
print(list(pares))  # [0, 2, 4, 6, 8]

# Chain generators
def filtrar_pares(números):
    for num in números:
        if num % 2 == 0:
            yield num

def multiplicar_por_dos(números):
    for num in números:
        yield num * 2

enteros = range(5)
resultado = multiplicar_por_dos(filtrar_pares(enteros))
print(list(resultado))  # [0, 4, 8]
```

---

## Parte V: send() - Comunicación Bidireccional

```python
def calculadora():
    \"\"\"Acumula valores recibidos\"\"\"
    total = 0
    while True:
        valor = yield total  # Envía total, recibe valor
        if valor is None:
            break
        total += valor

# Uso
calc = calculadora()
print(next(calc))           # Inicia, print: 0

print(calc.send(10))        # Suma 10, print: 10
print(calc.send(5))         # Suma 5, print: 15
print(calc.send(3))         # Suma 3, print: 18

try:
    calc.send(None)         # Detiene
except StopIteration:
    pass
```

---

## Parte VI: throw() - Manejo de Errores

```python
def generador_robusto():
    try:
        yield 1
        yield 2
        yield 3
    except ValueError as e:
        print(f"Error capturado: {e}")
        yield 999

gen = generador_robusto()
print(next(gen))          # 1
print(next(gen))          # 2
print(gen.throw(ValueError("Error forzado")))  # Captura, yield 999
```

---

## Parte VII: close() - Detener Generador

```python
def generador_infinito():
    contador = 0
    while True:
        yield contador
        contador += 1

gen = generador_infinito()
print(next(gen))  # 0
print(next(gen))  # 1

gen.close()  # Detiene

try:
    next(gen)  # StopIteration
except StopIteration:
    print("Generador cerrado")
```

---

## Parte VIII: itertools - Utilidades

```python
import itertools

# Encadenamiento
it1 = iter([1, 2, 3])
it2 = iter([4, 5, 6])
encadenado = itertools.chain(it1, it2)
print(list(encadenado))  # [1, 2, 3, 4, 5, 6]

# Repetición
print(list(itertools.repeat("a", 3)))  # ['a', 'a', 'a']

# Combinaciones
print(list(itertools.combinations("ABC", 2)))
# [('A', 'B'), ('A', 'C'), ('B', 'C')]

# Permutaciones
print(list(itertools.permutations("AB")))
# [('A', 'B'), ('B', 'A')]

# Ciclos infinitos
contador = itertools.count(10)
print(next(contador))  # 10
print(next(contador))  # 11
print(next(contador))  # 12

# Grupos
datos = [1, 1, 1, 2, 2, 3, 3, 3, 3]
grupos = itertools.groupby(datos)
for valor, grupo in grupos:
    print(valor, list(grupo))
# 1 [1, 1, 1]
# 2 [2, 2]
# 3 [3, 3, 3, 3]
```

---

## Parte IX: Async/Await (Introducción)

Corrutinas asincrónicas:

```python
import asyncio

async def descargar(url):
    \"\"\"Simula descarga\"\"\"
    print(f"Descargando {url}...")
    await asyncio.sleep(2)  # Pausa 2 segundos
    print(f"Completado {url}")
    return f"Contenido de {url}"

async def main():
    # Ejecutar en paralelo
    resultados = await asyncio.gather(
        descargar("http://ejemplo1.com"),
        descargar("http://ejemplo2.com"),
        descargar("http://ejemplo3.com"),
    )
    return resultados

# Ejecutar
resultados = asyncio.run(main())
print(resultados)

# Sin async/await (sería secuencial):
# Tiempo: 6 segundos
# Con async/await (paralelo):
# Tiempo: 2 segundos (se superponen)
```

---

## Parte X: Ejemplo Real - Lectura Lazy

```python
def leer_archivo_líneas(archivo):
    \"\"\"Lee archivo línea por línea (bajo demanda)\"\"\"
    with open(archivo) as f:
        for línea in f:
            yield línea.strip()

def procesar_archivo_grande(archivo):
    \"\"\"Procesa archivo sin cargarlo completo en RAM\"\"\"
    for línea in leer_archivo_líneas(archivo):
        # Procesar línea individual
        datos = línea.split(",")
        yield {
            "id": datos[0],
            "nombre": datos[1],
            "email": datos[2],
        }

# Uso
for registro in procesar_archivo_grande("millones_de_líneas.csv"):
    enviar_a_bd(registro)
    # Nunca cargamos más de una línea en RAM
```

---

## Parte XI: Comparación - Memoria

```python
import sys
import time

# Versión 1: Lista (carga todo)
def números_lista(n):
    return [i for i in range(n)]

# Versión 2: Generador (bajo demanda)
def números_generador(n):
    for i in range(n):
        yield i

# Comparar
n = 1_000_000

# Lista
lista = números_lista(n)
print(f"Tamaño lista: {sys.getsizeof(lista)} bytes")  # ~8 MB

# Generador
gen = números_generador(n)
print(f"Tamaño generador: {sys.getsizeof(gen)} bytes")  # ~144 bytes

# Velocidad de creación
import timeit

t1 = timeit.timeit(lambda: números_lista(n), number=1)
print(f"Tiempo lista: {t1:.3f}s")

t2 = timeit.timeit(lambda: números_generador(n), number=1)
print(f"Tiempo generador: {t2:.6f}s")
# Generador es casi instantáneo usando generators
```

---

## Parte XII: Mejores Prácticas

```python
# ✅ BIEN - yield eficiente
def líneas_pares(archivo):
    with open(archivo) as f:
        for num, línea in enumerate(f):
            if num % 2 == 0:
                yield línea

# ❌ MAL - construir lista primero
def líneas_pares_mal(archivo):
    with open(archivo) as f:
        resultado = []
        for num, línea in enumerate(f):
            if num % 2 == 0:
                resultado.append(línea)
    return resultado  # Espera a terminar

# Para iterarlo varias veces
def números(n):
    \"\"\"Generador reutilizable\"\"\"
    for i in range(n):
        yield i

# Si necesitas reutilizar, usa itertools.tee
import itertools
it1, it2 = itertools.tee(números(5))
print(list(it1))  # [0, 1, 2, 3, 4]
print(list(it2))  # [0, 1, 2, 3, 4]
```

---

## Resumen

| Concepto | Uso |
|----------|-----|
| `yield` | Pausar y reanudar |
| Generador | Crear bajo demanda |
| `next()` | Obtener siguiente valor |
| `send()` | Enviar valor a generador |
| `throw()` | Lanzar excepción |
| `close()` | Detener generador |
| `itertools` | Combinaciones/permutaciones |
| `async/await` | Ejecución concurrente |
| Lazy eval | Memoria eficiente |

---

**Siguiente**: [stdlib_utiles.md](stdlib_utiles.md)
