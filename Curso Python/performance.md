# Performance y Optimización

> **Hacer código rápido - DESPUÉS de que funcione correctamente**

---

## ¿QUÉ es Performance?

**Performance** = qué tan rápido se ejecuta tu código:

```python
# ❌ Lento (tarda 10 segundos)
def buscar_usuario(usuario_buscado):
    for usuario in [millón de usuarios]:
        if usuario == usuario_buscado:
            return usuario

# ✅ Rápido (tarda 0.001 segundos)
diccionario = {usuario.id: usuario}
return diccionario.get(usuario_buscado)
```

**Analogía:**
- 🚗 Coche lento = tarda 1 hora en ir a trabajo
- 🚄 Tren rápido = tarda 20 minutos

---

## ¿PARA QÉ sirve?

### 1. Aplicaciones web no cuelgan
```python
# Lento: usuario espera 10 segundos
# Rápido: usuario ve respuesta en 0.1 segundos
```

### 2. Procesar datos masivos
```python
# Analizar 1 millón de transacciones:
# Lento: tarda 1 hora
# Rápido: tarda 5 minutos
```

### 3. Ahorrar dinero
```python
# Servidor lento necesita 100 máquinas
# Servidor rápido necesita 1 máquina
# ($ optimizado = $ ahorrado)
```

### 4. Mejor experiencia usuario
```python
# Aplicación móvil lenta = batería agotada
# Aplicación rápida = batería dura
```

---

## ¿CÓMO optimizar?

### 4 Pasos (Importante: en ESTE orden)

| Paso | Qué | Por qué |
|------|-----|--------|
| 1. **Correcto** | Código que funciona | No optimices código roto |
| 2. **Medir** | ¿Cuánto tarda? | Mide ANTES de optimizar |
| 3. **Identificar** | ¿Cuál parte es lenta? | No optimices al azar |
| 4. **Optimizar** | Cambia eso específico | Ahora sí, mejora |

**Lema del Ingeniero:**
> "No optimices las suposiciones - mide primero"

**Proceso Visual:**
```
Código ✅ → Medir ⏱️ → Lento? ❌
                Sí ↓
            Identificar el cuello 🔍
                ↓
            Optimizar ese parte 💨
                ↓
            Medir de nuevo ⏱️
```

---

## Parte I: Medir Tiempo - timeit

```python
import timeit

# Medir una línea
tiempo = timeit.timeit('sum(range(100))', number=100000)
print(f"Tiempo: {tiempo:.6f} segundos")

# Medir función
def función_lenta():
    return sum(i**2 for i in range(100))

tiempo = timeit.timeit(función_lenta, number=10000)
print(f"10000 iteraciones: {tiempo:.4f}s")

# Comparar dos métodos
def método_1():
    return [x**2 for x in range(100)]  # Lista comprehension

def método_2():
    resultado = []
    for x in range(100):
        resultado.append(x**2)
    return resultado  # Append manual

t1 = timeit.timeit(método_1, number=10000)
t2 = timeit.timeit(método_2, number=10000)

print(f"Método 1: {t1:.4f}s")
print(f"Método 2: {t2:.4f}s")
print(f"Método 1 es {t2/t1:.1f}x más rápido")
```

---

## Parte II: Profiling - cProfile

Ver qué funciones usan más tiempo:

```python
import cProfile
import pstats

def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

# Perfilar
profiler = cProfile.Profile()
profiler.enable()

resultado = fibonacci(30)

profiler.disable()
stats = pstats.Stats(profiler)
stats.sort_stats('cumulative')
stats.print_stats(10)  # Top 10
```

Línea de comando:
```bash
python -m cProfile -s cumulative mi_script.py
```

---

## Parte III: memory_profiler - Uso de RAM

```bash
pip install memory-profiler
```

```python
from memory_profiler import profile

@profile
def función_pesada():
    lista = [i**2 for i in range(1000000)]
    diccionario = {i: i**2 for i in range(100000)}
    return len(lista) + len(diccionario)

función_pesada()
```

Ejecutar:
```bash
python -m memory_profiler mi_script.py
```

Salida muestra RAM usada línea por línea.

---

## Parte IV: time.perf_counter() - Precisión Alta

```python
import time

inicio = time.perf_counter()

# Operación a medir
resultado = sum(i**2 for i in range(10000000))

fin = time.perf_counter()

print(f"Tiempo: {fin - inicio:.4f} segundos")
```

---

## Parte V: Optimizaciones Comunes

### 1. Buscar en Set vs Lista

```python
import timeit

lista = list(range(10000))
conjunto = set(range(10000))

# En lista: O(n)
t_lista = timeit.timeit(lambda: 9999 in lista, number=100000)

# En set: O(1)
t_conjunto = timeit.timeit(lambda: 9999 in conjunto, number=100000)

print(f"Lista: {t_lista:.4f}s")
print(f"Set: {t_conjunto:.4f}s")
print(f"Set es {t_lista/t_conjunto:.0f}x más rápido")
```

### 2. List Comprehension vs Append

```python
import timeit

# Comprehension
def comp():
    return [x**2 for x in range(1000)]

# Append
def append_manual():
    resultado = []
    for x in range(1000):
        resultado.append(x**2)
    return resultado

t1 = timeit.timeit(comp, number=10000)
t2 = timeit.timeit(append_manual, number=10000)

print(f"Comprehension: {t1:.4f}s (más rápido)")
print(f"Append: {t2:.4f}s")
```

### 3. String Concatenation

```python
import timeit

# ❌ Lento - crea nuevas strings
def concat_lento():
    resultado = ""
    for i in range(1000):
        resultado += str(i)  # Crea nueva string cada vez
    return resultado

# ✅ Rápido - join
def concat_rápido():
    return "".join(str(i) for i in range(1000))

t1 = timeit.timeit(concat_lento, number=100)
t2 = timeit.timeit(concat_rápido, number=100)

print(f"Concatenacion: {t1:.4f}s")
print(f"Join: {t2:.4f}s (mucho más rápido)")
```

### 4. Generator vs Lista

```python
import timeit

# Lista (carga toda en memoria)
def números_lista():
    return [x**2 for x in range(1000000)]

# Generador (bajo demanda)
def números_gen():
    return (x**2 for x in range(1000000))

# Crear
t1 = timeit.timeit(números_lista, number=10)
t2 = timeit.timeit(números_gen, number=10)

print(f"Lista: {t1:.4f}s")
print(f"Generador: {t2:.6f}s (casi gratis)")
```

---

## Parte VI: Local Variables - Acceso Más Rápido

```python
import timeit

# ❌ Global (más lento, lookup en dict global)
contador_global = 0
def función_global():
    global contador_global
    for i in range(1000):
        contador_global += 1

# ✅ Local (más rápido, variable en stack)
def función_local():
    contador = 0
    for i in range(1000):
        contador += 1
    return contador

t1 = timeit.timeit(función_global, number=10000)
t2 = timeit.timeit(función_local, number=10000)

print(f"Global: {t1:.4f}s")
print(f"Local: {t2:.4f}s (más rápido)")
```

---

## Parte VII: numpy para Operaciones Numéricas

```python
import numpy as np
import timeit

# Python puro
def python_suma():
    lista = list(range(1000000))
    return sum(lista)

# NumPy
def numpy_suma():
    arr = np.arange(1000000)
    return np.sum(arr)

t1 = timeit.timeit(python_suma, number=100)
t2 = timeit.timeit(numpy_suma, number=100)

print(f"Python: {t1:.4f}s")
print(f"NumPy: {t2:.4f}s (10-100x más rápido)")
```

---

## Parte VIII: Caché Local

```python
from functools import lru_cache
import timeit

@lru_cache(maxsize=None)
def fibonacci_cache(n):
    if n < 2:
        return n
    return fibonacci_cache(n-1) + fibonacci_cache(n-2)

def fibonacci_sin_cache(n):
    if n < 2:
        return n
    return fibonacci_sin_cache(n-1) + fibonacci_sin_cache(n-2)

# Con caché
t1 = timeit.timeit(lambda: fibonacci_cache(30), number=10)

# Sin caché (muy lento)
t2 = timeit.timeit(lambda: fibonacci_sin_cache(25), number=1)

print(f"Con caché (fib(30)): {t1:.4f}s")
print(f"Sin caché (fib(25)): {t2:.4f}s")
print("Caché es MUCHO más rápido")
```

---

## Parte IX: PyPy - Interpretador Más Rápido

```bash
# Instalar PyPy
# Windows/Mac/Linux desde https://www.pypy.org/

pypy3 -c "print(sum(i**2 for i in range(100000000)))"
```

PyPy vs CPython:
- PyPy: ~0.5 segundos (JIT compilation)
- CPython: ~5 segundos (interpretación)

**10x más rápido** en loops intensivos.

---

## Parte X: Algoritmos - Lo Más Importante

```python
import timeit

# ❌ O(n²) - búsqueda ingenua
def búsqueda_ingenua(lista, objetivo):
    for x in lista:
        for y in lista:
            if x + y == objetivo:
                return (x, y)

# ✅ O(n) - usando set
def búsqueda_eficiente(lista, objetivo):
    visto = set()
    for x in lista:
        complemento = objetivo - x
        if complemento in visto:
            return (complemento, x)
        visto.add(x)

lista = list(range(1000))
objetivo = 1500

t1 = timeit.timeit(lambda: búsqueda_ingenua(lista, objetivo), number=10)
t2 = timeit.timeit(lambda: búsqueda_eficiente(lista, objetivo), number=10000)

print(f"Ingenua (n²): {t1:.4f}s")
print(f"Eficiente (n): {t2:.4f}s")
print("El algoritmo es MÁS importante que PyPy o compilación")
```

---

## Parte XI: Profiling Real

```python
import cProfile
import pstats
from io import StringIO

def procesar_datos():
    datos = [i for i in range(100000)]
    resultados = []
    
    for dato in datos:
        resultado = dato ** 2  # O(1)
        if resultado > 50000000:
            resultados.append(resultado)
    
    return resultados

# Perfilar
profiler = cProfile.Profile()
profiler.enable()

for _ in range(100):
    procesar_datos()

profiler.disable()

# Mostrar resultados
s = StringIO()
ps = pstats.Stats(profiler, stream=s).sort_stats('cumulative')
ps.print_stats()
print(s.getvalue())
```

---

## Parte XII: Checklist de Optimización

1. ✅ ¿El código es correcto? (tests pasando)
2. ✅ ¿Es realmente lento? (timeit/profiling)
3. ✅ ¿Dónde está el cuello? (cProfile)
4. ✅ ¿Es problema de algoritmo? (Big-O complexity)
5. ✅ ¿Es problema de memoria? (memory-profiler)
6. ✅ ¿Vale la pena optimizar? (costo/beneficio)

---

## Resumen

| Herramienta | Usa |
|-------------|-----|
| timeit | Medir tiempo de pequeños fragmentos |
| cProfile | Profiling de funciones |
| memory_profiler | Uso de RAM línea por línea |
| perf_counter | Precisión alta en tiempo |
| Big-O | Análisis de algoritmo |
| numpy | Operaciones numéricas rápidas |
| lru_cache | Caché automática |
| PyPy | Interpretador JIT más rápido |

---

**Siguiente**: [web_basico.md](web_basico.md)
