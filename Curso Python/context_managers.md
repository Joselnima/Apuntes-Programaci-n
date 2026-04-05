# Context Managers - Gestión de Recursos

> **Asegurar que abres y cierras recursos automáticamente**

---

## ¿QUÉ es un Context Manager?

Un **context manager** = garantiza limpieza automática:

```python
# ❌ Manual (peligroso)
archivo = open("datos.txt")
contenido = archivo.read()
archivo.close()  # Confío que NO olvides

# ✅ Automático (seguro)
with open("datos.txt") as archivo:
    contenido = archivo.read()
# Se cierra automáticamente aquí
```

**Analogía:**
- 🏠 `with` = "entra a la casa, usa lo que quieras, cierta la puerta al salir"
- 🔑 Automático = la puerta se cierra SOLA

---

## ¿PARA QÉ sirven?

### 1. Archivos (cierre garantizado)
```python
# ❌ Riesgo: Se queda abierto
try:
    archivo = open("datos.txt")
    contenido = archivo.read()
except Exception as e:
    print(f"Error: {e}")
    # ⚠️ El archivo SIGUE abierto

# ✅ Seguro: Cierra siempre
with open("datos.txt") as archivo:
    contenido = archivo.read()
# Cierra incluso si hay error
```

### 2. Conexiones a BD
```python
with conexión.transaction():
    insertar_usuario("Juan")
    insertar_usuario("María")
# Si todo va bien: COMMIT automático
# Si hay error: ROLLBACK automático
```

### 3. Locks en multithreading
```python
lock = threading.Lock()

with lock:
    # Solo un thread aquí
    variable_compartida += 1
# Lock se libera automáticamente
```

### 4. Liberar memoria
```python
with recurso.usar() as r:
    procesar(r)
# Memoria liberada automáticamente
```

---

## ¿CÓMO funcionan?

### Estructura del `with`

```
with RECURSO as VARIABLE:
    ↑       ↑  ↑
    Abre  Variable disponible
    
    # Usa VARIABLE aquí
    # Ejecuta __enter__() del recurso
    
# Sale del bloque
# Ejecuta __exit__() automáticamente (LIMPIEZA)
```

**Proceso Interno:**
```python
1. archivo = open(...)      # __enter__() ejecutada
2. contenido = archivo...   # Tu código aquí
3. archivo.close()          # __exit__() ejecutada (automática)
```

---

## Parte I: `with` Statement

```python
# Estructura
with expresión as variable:
    # variable disponible aquí
    hacer_algo(variable)
# variable no disponible aquí
```

Ejemplos comunes:

```python
# Archivos
with open("archivo.txt") as f:
    contenido = f.read()

# Locks en multithreading
import threading
lock = threading.Lock()
with lock:
    sección_crítica()

# Transacciones en BD
with conexión.transaction():
    conexión.ejecutar("INSERT ...")
    conexión.ejecutar("UPDATE ...")
# Commit automático si no hay error
```

---

## Parte II: Crear Context Managers - Clase

```python
class ConexiónBD:
    def __init__(self, url):
        self.url = url
        self.conexión = None
    
    def __enter__(self):
        \"\"\"Se ejecuta cuando entra en with\"\"\"
        print(f"Conectando a {self.url}...")
        self.conexión = conectar(self.url)
        return self.conexión  # Qué obtiene la variable
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        \"\"\"Se ejecuta cuando sale de with\"\"\"
        print("Cerrando conexión...")
        if self.conexión:
            self.conexión.cerrar()
        
        # Return False para propagar excepciones
        # Return True para suprimir excepciones
        return False

# Uso
with ConexiónBD("postgresql://localhost/db") as conn:
    conn.ejecutar("SELECT * FROM usuarios")
    # Cierra automáticamente
```

---

## Parte III: Decorador `@contextmanager`

Forma más simple con generador:

```python
from contextlib import contextmanager

@contextmanager
def abrir_recurso(nombre):
    print(f"Abriendo {nombre}...")
    recurso = crear_recurso(nombre)
    try:
        yield recurso  # Lo que obtiene la variable
    finally:
        print(f"Cerrando {nombre}...")
        cerrar_recurso(recurso)

# Uso
with abrir_recurso("archivo.txt") as r:
    usar(r)
# Cierra automáticamente
```

### Manejo de Excepciones

```python
from contextlib import contextmanager

@contextmanager
def manejo_errores():
    try:
        yield
    except ValueError as e:
        print(f"Error de valor: {e}")
    except Exception as e:
        print(f"Error inesperado: {e}")
        raise
    finally:
        print("Limpieza...")

with manejo_errores():
    x = int(input())  # If bad: captura ValueError
```

---

## Parte IV: Archivo con Manejo Manual

```python
@contextmanager
def abrir_texto(archivo, modo="r", encoding="utf-8"):
    \"\"\"Context manager para archivos\"\"\"
    f = open(archivo, modo, encoding=encoding)
    try:
        yield f
    finally:
        f.close()

# Uso
with abrir_texto("datos.txt") as f:
    for línea in f:
        print(línea.strip())
```

---

## Parte V: Timer Context Manager

```python
from contextlib import contextmanager
import time

@contextmanager
def medir_tiempo(nombre="Operación"):
    inicio = time.time()
    print(f"[{nombre}] Iniciando...")
    try:
        yield
    finally:
        tiempo = time.time() - inicio
        print(f"[{nombre}] Completó en {tiempo:.2f}s")

# Uso
with medir_tiempo("Procesamiento"):
    results = [x**2 for x in range(1000000)]
    
# Salida:
# [Procesamiento] Iniciando...
# [Procesamiento] Completó en 0.05s
```

---

## Parte VI: Múltiples Context Managers

```python
# Python 3.10+
with open("entrada.txt") as entrada, open("salida.txt", "w") as salida:
    for línea in entrada:
        salida.write(línea.upper())

# Python 3.9 y anterior
from contextlib import ExitStack
with ExitStack() as stack:
    entrada = stack.enter_context(open("entrada.txt"))
    salida = stack.enter_context(open("salida.txt", "w"))
    for línea in entrada:
        salida.write(línea.upper())
```

---

## Parte VII: ExitStack - Apilamiento Dinámico

```python
from contextlib import ExitStack

def procesar_archivos(nombres):
    with ExitStack() as stack:
        # Abrir archivos dinámicamente
        archivos = [
            stack.enter_context(open(nombre))
            for nombre in nombres
        ]
        
        for archivo in archivos:
            procesar(archivo)
        # Todos se cierran automáticamente en orden inverso
```

---

## Parte VIII: BD - Transacciones

```python
@contextmanager
def transacción(conexión):
    try:
        yield conexión
        conexión.commit()
        print("✓ Transacción confirmada")
    except Exception as e:
        conexión.rollback()
        print(f"✗ Transacción revertida: {e}")
        raise

# Uso
with transacción(conn) as c:
    c.ejecutar("INSERT INTO usuarios VALUES (...)")
    c.ejecutar("UPDATE cuentas SET saldo = saldo - 100")
    # Si todo OK: commit automático
    # Si error: rollback automático
```

---

## Parte IX: Lock y Threading

```python
from threading import Lock
from contextlib import contextmanager

@contextmanager
def sección_crítica(lock):
    lock.acquire()
    try:
        yield
    finally:
        lock.release()

# O mejor - Lock es ya context manager:
import threading

lock = threading.Lock()

with lock:
    variable_compartida += 1
    # Thread-safe aquí
# Libera automáticamente
```

---

## Parte X: Suprimir Excepciones

```python
from contextlib import suppress

# Sin suppress
try:
    operación_que_puede_fallar()
except ValueError:
    pass

# Con suppress
with suppress(ValueError):
    operación_que_puede_fallar()

# Múltiples excepciones
with suppress(ValueError, FileNotFoundError):
    algo_que_puede_fallar()
```

---

## Parte XI: Redirect STDOUT/STDERR

```python
import io
from contextlib import redirect_stdout, redirect_stderr

# Capturar salida
buffer = io.StringIO()
with redirect_stdout(buffer):
    print("Esto se captura")
    print("No se ve en terminal")

salida = buffer.getvalue()
print(f"Fue capturado: {salida}")

# Redirigir stderr
log_file = open("errores.log", "w")
with redirect_stderr(log_file):
    print("error", file=__import__("sys").stderr)
log_file.close()
```

---

## Parte XII: Ejemplo Completo

```python
from contextlib import contextmanager
import time
import sqlite3

@contextmanager
def conexión_bd(base_datos):
    \"\"\"Context manager para BD\"\"\"
    print(f"Conectando a {base_datos}...")
    conn = sqlite3.connect(base_datos)
    try:
        yield conn
        conn.commit()
        print("✓ Cambios guardados")
    except Exception as e:
        conn.rollback()
        print(f"✗ Error: {e}")
        raise
    finally:
        conn.close()
        print("Conexión cerrada")

@contextmanager
def timing():
    \"\"\"Medir tiempo de ejecución\"\"\"
    inicio = time.time()
    try:
        yield
    finally:
        print(f"Tiempo: {time.time() - inicio:.3f}s")

# Uso
with conexión_bd("datos.db") as conn, timing():
    cursor = conn.cursor()
    cursor.execute("CREATE TABLE usuarios (id INT, nombre TEXT)")
    cursor.execute("INSERT INTO usuarios VALUES (1, 'Juan')")
    
# Salida:
# Conectando a datos.db...
# ✓ Cambios guardados
# Conexión cerrada
# Tiempo: 0.002s
```

---

## Resumen

| Concepto | Uso |
|----------|-----|
| `with` | Entra/sale automático |
| `__enter__/__exit__` | Protocolo context manager |
| `@contextmanager` | Usar con generador |
| `ExitStack` | Múltiples dinámicos |
| `finally` | Siempre ejecutar limpieza |
| `suppress()` | Ignorar excepciones |
| Transacciones | Commit/rollback automático |

---

**Siguiente**: [generators_corrutinas.md](generators_corrutinas.md)
