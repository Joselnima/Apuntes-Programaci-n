# Decoradores en Python

> **Una forma profesional de mejorar funciones sin cambiar su código**

---

## ¿QUÉ es un Decorador?

Un **decorador** = función que **modifica otra función**

```python
@mi_decorador
def mi_función():
    return "Hola"

# Es lo MISMO que:
def mi_función():
    return "Hola"
mi_función = mi_decorador(mi_función)
```

**Analogía:**
- 🎁 Tu función = caja de regalo
- 🎀 Decorador = papel y cinta decorativo
- 📦 Función decorada = caja hermosa, pero sigue siendo lo mismo adentro

**Símbolo `@`:**
- `@algo` = "envuelve la función que viene después"

---

## ¿PARA QÉ sirven?

### 1. Agregar funcionalidad sin cambiar código
```python
def cronómetro(func):
    """Mide cuánto tarda la función"""
    def envuelta():
        import time
        inicio = time.time()
        resultado = func()
        print(f"Tardó {time.time() - inicio} segundos")
        return resultado
    return envuelta

@cronómetro
def buscar_clientes():
    return "Clientes encontrados"

buscar_clientes()  # Automáticamente mide tiempo
```

### 2. Validar entrada de datos
```python
def requiere_admin(func):
    def envuelta(usuario):
        if usuario["nivel"] != "admin":
            return "❌ No tienes permisos"
        return func(usuario)
    return envuelta

@requiere_admin
def eliminar_usuario(usuario):
    return "✅ Usuario eliminado"
```

### 3. Repetir lógica en muchas funciones
```python
# En lugar de:
def login():
    print("Registrando login...")
    # lógica
    print("Login registrado")

def logout():
    print("Registrando logout...")
    # lógica
    print("Logout registrado")

# Usa decorador:
@registrar_evento
def login():
    pass

@registrar_evento
def logout():
    pass
```

---

## ¿CÓMO funcionan?

### Concepto paso a paso

```python
# Paso 1: Decorador recibe función
def decorador(func):
    
    # Paso 2: Crea función envuelta
    def envuelta():
        print("ANTES")
        func()  # Ejecuta función original
        print("DESPUÉS")
    
    # Paso 3: Retorna la envuelta
    return envuelta

# Paso 4: Usa @
@decorador
def saludar():
    print("¡Hola!")

# Resultado:
saludar()
# OUTPUT:
# ANTES
# ¡Hola!
# DESPUÉS
```

---

## Parte I: Funciones que Retornan Funciones

Requisito para entender decoradores:

```python
def crear_saludador(saludo):
    def saludar(nombre):
        return f"{saludo}, {nombre}"
    return saludar

buenos_días = crear_saludador("Buenos días")
buenas_noches = crear_saludador("Buenas noches")

print(buenos_días("Juan"))      # Buenos días, Juan
print(buenas_noches("María"))   # Buenas noches, María
```

---

## Parte II: Decorador Simple

### Estructura Básica

```python
def mi_decorador(func):
    \"\"\"Decorador simple\"\"\"
    def envuelta():
        print("Antes de ejecutar función")
        func()
        print("Después de ejecutar función")
    return envuelta

@mi_decorador
def saludar():
    print("Hola")

saludar()
# Output:
# Antes de ejecutar función
# Hola
# Después de ejecutar función
```

**Equivalente a**:
```python
def saludar():
    print("Hola")

saludar = mi_decorador(saludar)
saludar()
```

---

## Parte III: Decorador con Parámetros en Función

```python
def registrar(func):
    def envuelta(*args, **kwargs):
        print(f"Llamando: {func.__name__}({args}, {kwargs})")
        resultado = func(*args, **kwargs)
        print(f"Resultado: {resultado}")
        return resultado
    return envuelta

@registrar
def sumar(a, b):
    return a + b

sumar(5, 3)
# Output:
# Llamando: sumar((5, 3), {})
# Resultado: 8
```

---

## Parte IV: Decorador con Parámetros

Decorador que recibe argumentos:

```python
def repetir(veces):
    \"\"\"Decorador que repetir la función N veces\"\"\"
    def decorador(func):
        def envuelta(*args, **kwargs):
            resultados = []
            for _ in range(veces):
                resultados.append(func(*args, **kwargs))
            return resultados
        return envuelta
    return decorador

@repetir(3)
def saludar(nombre):
    return f"Hola {nombre}"

print(saludar("Juan"))
# Output: ['Hola Juan', 'Hola Juan', 'Hola Juan']
```

---

## Parte V: @wraps - Preservar Metadata

Sin `@wraps`, pierdes información de la función:

```python
def decorador1(func):
    def envuelta():
        \"\"\"Envuelta\"\"\"
        return func()
    return envuelta

@decorador1
def mi_función():
    \"\"\"Mi función\"\"\"
    pass

print(mi_función.__name__)  # envuelta (❌ perdimos el nombre)
print(mi_función.__doc__)   # Envuelta (❌ perdimos docstring)
```

**Con @wraps**:

```python
from functools import wraps

def decorador2(func):
    @wraps(func)  # Copia metadata
    def envuelta():
        return func()
    return envuelta

@decorador2
def mi_función():
    \"\"\"Mi función\"\"\"
    pass

print(mi_función.__name__)  # mi_función (✅)
print(mi_función.__doc__)   # Mi función (✅)
```

---

## Parte VI: Decoradores Prácticos

### Timing (Medir Tiempo)

```python
import time
from functools import wraps

def timing(func):
    @wraps(func)
    def envuelta(*args, **kwargs):
        inicio = time.time()
        resultado = func(*args, **kwargs)
        tiempo = time.time() - inicio
        print(f"{func.__name__} tomó {tiempo:.4f} segundos")
        return resultado
    return envuelta

@timing
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

fibonacci(10)
# Output: fibonacci tomó 0.0015 segundos
```

### Autenticación

```python
def requiere_login(func):
    @wraps(func)
    def envuelta(usuario, *args, **kwargs):
        if not usuario.get("logueado"):
            raise ValueError("Debes estar logueado")
        return func(usuario, *args, **kwargs)
    return envuelta

@requiere_login
def acceso_admin(usuario):
    print(f"Acceso para {usuario['nombre']}")

# Falla
try:
    acceso_admin({"nombre": "Juan", "logueado": False})
except ValueError as e:
    print(f"Error: {e}")

# Funciona
acceso_admin({"nombre": "Juan", "logueado": True})
```

### Cachéo (Memoización)

```python
def cachear(func):
    cache = {}
    
    @wraps(func)
    def envuelta(n):
        if n not in cache:
            cache[n] = func(n)
        return cache[n]
    return envuelta

@cachear
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print(fibonacci(100))  # Muy rápido después de la primera llamada
```

**Alternativa**: usar `@lru_cache` de functools

```python
from functools import lru_cache

@lru_cache(maxsize=32)
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print(fibonacci(100))
```

### Validación

```python
def validar_tipos(**tipos_esperados):
    def decorador(func):
        @wraps(func)
        def envuelta(*args, **kwargs):
            # Validar argumentos
            for nombre, tipo in tipos_esperados.items():
                if nombre in kwargs:
                    if not isinstance(kwargs[nombre], tipo):
                        raise TypeError(f"{nombre} debe ser {tipo.__name__}")
            return func(*args, **kwargs)
        return envuelta
    return decorador

@validar_tipos(edad=int, nombre=str)
def crear_usuario(nombre, edad):
    return f"{nombre} ({edad})"

print(crear_usuario(nombre="Juan", edad=25))  # OK
# print(crear_usuario(nombre="Juan", edad="25"))  # TypeError
```

---

## Parte VII: Decoradores de Clase

Decoradores que actúan sobre clases:

```python
def singleton(cls):
    \"\"\"Asegurar que solo existe una instancia\"\"\"
    instancia = {}
    
    def obtener_instancia(*args, **kwargs):
        if cls not in instancia:
            instancia[cls] = cls(*args, **kwargs)
        return instancia[cls]
    
    return obtener_instancia

@singleton
class BaseDatos:
    def __init__(self):
        self.conexión = None

db1 = BaseDatos()
db2 = BaseDatos()

print(db1 is db2)  # True (misma instancia)
```

---

## Parte VIII: Stacked Decorators (Apilar)

```python
def decorador1(func):
    def envuelta():
        print("1. Antes")
        resultado = func()
        print("1. Después")
        return resultado
    return envuelta

def decorador2(func):
    def envuelta():
        print("2. Antes")
        resultado = func()
        print("2. Después")
        return resultado
    return envuelta

@decorador1
@decorador2
def mi_función():
    print("Función")

mi_función()
# Output:
# 1. Antes
# 2. Antes
# Función
# 2. Después
# 1. Después
```

Se ejecutan de abajo hacia arriba en la definición.

---

## Parte IX: Decoradores de Métodos

### @property

```python
class Persona:
    def __init__(self, nombre, año_nacimiento):
        self._nombre = nombre
        self._año_nacimiento = año_nacimiento
    
    @property
    def edad(self):
        from datetime import datetime
        return datetime.now().year - self._año_nacimiento
    
    @edad.setter
    def edad(self, valor):
        self._año_nacimiento = datetime.now().year - valor

p = Persona("Juan", 1998)
print(p.edad)      # Llama al getter
p.edad = 30        # Llama al setter
```

### @classmethod y @staticmethod

```python
class Utilidades:
    contador = 0
    
    @classmethod
    def incrementar(cls):
        cls.contador += 1
    
    @staticmethod
    def sumar(a, b):
        return a + b

print(Utilidades.sumar(5, 3))  # 8
Utilidades.incrementar()
print(Utilidades.contador)      # 1
```

---

## Resumen

| Tipo | Sintaxis | Uso |
|------|----------|-----|
| Simple | `@decorador` | Envolver función |
| Con args | `@decorador(args)` | Pasar parámetros |
| @wraps | `@wraps(func)` | Preservar metadata |
| @property | `@property` | Atributo computed |
| @classmethod | `@classmethod` | Método de clase |
| @staticmethod | `@staticmethod` | Método sin self |
| Stacked | `@dec1, @dec2` | Múltiples decoradores |

---

**Siguiente**: [bases_datos.md](bases_datos.md)
