# Variables y Asignación en Python

> **Cómo guardar y recordar información en tu programa - El primer paso para aprender a programar**

---

## ¿QUÉ es una Variable?

Imagina que estás en una biblioteca. Variable es como **una caja etiquetada** donde guardas libros:

- 📦 La **caja** = espacio en memoria de la computadora
- 📌 La **etiqueta** = nombre de la variable
- 📚 El **contenido** = valor que guardas

```python
# Creamos 3 cajas con etiquetas
nombre = "Juan"        # Una caja llamada "nombre" que contiene el texto "Juan"
edad = 25              # Una caja llamada "edad" que contiene el número 25
altura = 1.75          # Una caja llamada "altura" que contiene 1.75
```

**¿Cómo lo ve Python internamente?**
```
En la memoria de tu computadora:
┌─────────────────────────────────┐
│ Etiqueta: nombre                │
│ Contenido: "Juan"               │
└─────────────────────────────────┘

┌─────────────────────────────────┐
│ Etiqueta: edad                  │
│ Contenido: 25                   │
└─────────────────────────────────┘
```

---

## ¿PARA QUÉ sirven las Variables?

### Casos reales:
- **Guardar datos del usuario**: Cuando alguien entra a tu programa y se registra
- **Recordar valores**: La edad, el nombre, el saldo bancario
- **Hacer cálculos**: Sumar gastos, calcular promedios, etc.
- **Mostrar información**: Luego imprimes o usas esos datos

**Ejemplo práctico - Un banco:**
```python
# Sin variables (imposible)
# "¿Cuál es el saldo del cliente?"
# No hay forma de recordarlo...

# Con variables (posible)
saldo_juan = 5000
saldo_maria = 3500

print(f"Saldo de Juan: {saldo_juan}")  # Saldo de Juan: 5000
```

---

## ¿CÓMO funcionan las Variables?

### Paso 1: Declaración
```python
edad = 25
```
Le dices a Python: "Crea una variable llamada 'edad' y pon ahí el valor 25"

### Paso 2: Almacenamiento
Python automáticamente:
1. Reserva espacio en la memoria
2. Le pone la etiqueta "edad"
3. Guarda el valor 25 ahí

### Paso 3: Usar la Variable
```python
print(edad)  # Imprime: 25
nuevo_edad = edad + 1  # 25 + 1 = 26
```

**Lo importante**: Python automaticamente entiende QUÉ TIPO de dato es (esto se llama "tipado dinámico"):

---

## Parte I: Tu Primera Variable - Lo Básico Básico

### El concepto más simple

```python
nombre = "Juan"
```

**¿Qué pasa línea por línea?**

1. `nombre` ← Este es el nombre de la variable (la etiqueta)
2. `=` ← Esto significa "almacena esto en"
3. `"Juan"` ← Este es el valor (lo que guardas en la caja)

**En español**: "Guarda el texto 'Juan' en una caja llamada nombre"

### Otros ejemplos

```python
edad = 25                # Guarda el número 25 en una caja llamada "edad"
precio = 99.99           # Guarda 99.99 en una caja llamada "precio"
activo = True            # Guarda verdadero (True) en "activo"
ciudad = "Madrid"        # Guarda el texto "Madrid" en "ciudad"
```

### Python Adivina el Tipo

Lo especial de Python es que **automáticamente entiende** qué tipo de dato es:

```python
tipo_nombre = type(nombre)
print(tipo_nombre)       # <class 'str'> (significa TEXTO)

tipo_edad = type(edad)
print(tipo_edad)         # <class 'int'> (significa NÚMERO ENTERO)

tipo_precio = type(precio)
print(tipo_precio)       # <class 'float'> (significa NÚMERO CON DECIMALES)

tipo_activo = type(activo)
print(tipo_activo)       # <class 'bool'> (significa VERDADERO/FALSO)
```

**¿Por qué lo adivina?** Porque:
- Si ves `"texto"` entre comillas → es texto (str)
- Si ves `25` sin comillas → es número entero (int)
- Si ves `99.99` con punto → es número decimal (float)
- Si ves `True` o `False` → es booleano (bool)

---

## Parte II: Convenciones de Nombres - Cómo Nombrar Variables Correctamente

### ✅ BIEN - Nombres que hablan claro

```python
# Estos nombres te dicen exactamente qué contienen
edad_usuario = 25
nombre_completo = "Juan García"
es_administrador = True
saldo_cuenta_bancaria = 5000.50
cantidad_productos = 10

# Constantes (valores que NO cambian)
MAX_INTENTOS_LOGIN = 3
PRECIO_FIJO = 99.99
PI = 3.14159
```

**¿Por qué son buenos?**
- Al leer `edad_usuario` entiendes inmediatamente qué es
- Si otro programador lee tu código, lo entiende
- Cuando vuelves a tu código en 6 meses, lo entiendes

### ❌ MAL - Nombres confusos (EVITA)

```python
# ❌ Demasiado cortos - ¿qué es "x"?
x = 25

# ❌ Nombres en camelCase - es de Java, no Python
miEdad = 25
nombreCompleto = "Juan"

# ❌ Nombres muy genéricos
dato = 25
valor = "Juan"
cosa = True

# ❌ Nombres con espacios - da error
mi nombre = 25  # SyntaxError

# ❌ Empezar con número
1nombre = "Juan"  # SyntaxError

# ❌ Usar caracteres especiales
mi-edad = 25  # Python lo interpreta como resta
mi.edad = 25  # Python lo interpreta como propiedad
```

### Las Reglas del Juego (Obligatorias)

```python
# ✅ PERMITIDO
nombre = "Juan"           # letra simple
primer_nombre = "Juan"    # contiene guion bajo (_)
nombre1 = "Juan"          # contiene números (pero NO al principio)
CONSTANTE = 100           # MAYÚSCULAS para constantes

# ❌ NO PERMITIDO
1nombre = "Juan"          # No puede empezar con número
mi nombre = "Juan"        # No puede tener espacios
mi-nombre = "Juan"        # No puede tener guiones normales
mi.nombre = "Juan"        # Puntos tienen otro significado
mi@nombre = "Juan"        # Caracteres especiales no allowed

# ⚠️ ESPECIAL - Palabras reservadas de Python
# Estos nombres NO puedes usar (son palabras del lenguaje)
class = 5        # Error - "class" es palabra reservada
def = "Juan"     # Error - "def" es palabra reservada
if = True        # Error - "if" es palabra reservada
```

### Resumen de Convenciones

| Tipo | Formato | Ejemplo |
|------|---------|---------|
| Variable normal | snake_case | `edad_usuario`, `precio_producto` |
| Constante | MAYÚSCULAS | `MAX_INTENTOS`, `TAX_RATE` |
| NO usar | camelCase | ~~`miEdad`~~ |
| NO usar | PascalCase | ~~`MiEdad`~~ (solo para clases) |
| Regla oro | **Nombres descriptivos** | `es_mayor` ✅ vs `x` ❌ |

---

## Parte III: Asignación Múltiple

### Asignación simultánea

```python
a, b, c = 1, 2, 3
print(a, b, c)  # 1 2 3

# Con diferentes tipos
x, nombre, activo = 5, "Juan", True
```

### Intercambiar (swap)

```python
a = 5
b = 10

# En otros lenguajes: necesitarías una variable temporal
# En Python:
a, b = b, a
print(a, b)  # 10 5
```

### Asignación desde listas (unpacking)

```python
valores = [1, 2, 3]
x, y, z = valores
print(x, y, z)  # 1 2 3

# Con ignorar valores
primero, _, tercero = valores
print(primero, tercero)  # 1 3
```

---

## Parte IV: Tipos de Asignación

### Forma 1: Asignación Directa

```python
edad = 25
```

### Forma 2: Operadores de Asignación Compuesta

```python
x = 10
x += 5    # x = x + 5  →   x = 15
x -= 3    # x = x - 3  →   x = 12
x *= 2    # x = x * 2  →   x = 24
x /= 4    # x = x / 4  →   x = 6.0
x //= 2   # x = x // 2 →   x = 3.0
x %= 2    # x = x % 2  →   x = 1.0
x **= 2   # x = x ** 2 →   x = 1.0
```

### Forma 3: No modificar (strings/tuplas)

Las strings y tuplas son **inmutables**:

```python
texto = "Hola"
# ❌ Esto no funciona:
# texto[0] = "J"  # Error: 'str' object does not support item assignment

# En su lugar, crea una nueva:
texto = "J" + texto[1:]  # "Jola"
```

---

## Parte V: Scope (Alcance) - IMPORTANTE

El **scope** determina dónde puedes acceder a una variable.

### Scope Local

Variables dentro de funciones solo existen en esa función:

```python
def saludar():
    nombre = "Juan"  # Local
    print(nombre)

saludar()            # Juan
print(nombre)        # NameError: nombre no está definida
```

### Scope Global

Variables fuera de funciones son globales:

```python
nombre = "Juan"  # Global

def saludar():
    print(nombre)  # Acceso a global

saludar()  # Juan
print(nombre)  # Juan
```

### Scope Local Modifica Global

```python
contador = 0

def incrementar():
    global contador  # Declara que usarás la global
    contador += 1

incrementar()
print(contador)  # 1

# Sin global:
def otra():
    contador += 1   # X Error: no existe local, no escribes global
```

### Nonlocal (Para Anidación)

```python
def externa():
    x = 10  # Scope de externa
    
    def interna():
        nonlocal x
        x += 1
    
    interna()
    print(x)  # 11

externa()
```

### El Modelo LEGB

Python busca variables en este orden:

1. **L**ocal: dentro de función
2. **E**nclosing: en función que encierra (anidadas)
3. **G**lobal: en el módulo
4. **B**uilt-in: funciones incorporadas (print, len, etc.)

```python
x = "global"

def externa():
    x = "enclosing"
    
    def interna():
        x = "local"
        print(x)  # ¿Cuál?
    
    interna()

externa()  # "local" (la más cercana gana)
```

---

## Parte VI: F-Strings (Python 3.6+)

La forma moderna de formatear strings:

### Básico

```python
nombre = "Juan"
edad = 25

# F-strings
print(f"Me llamo {nombre} y tengo {edad} años")

# Alternativa vieja (aún funciona):
print("Me llamo {} y tengo {} años".format(nombre, edad))
```

### Con Expresiones

```python
a = 5
b = 3

print(f"{a} + {b} = {a + b}")  # 5 + 3 = 8
print(f"¿Es {a} > {b}? {a > b}")  # ¿Es 5 > 3? True

def doble(x):
    return x * 2

print(f"El doble de 5 es {doble(5)}")  # El doble de 5 es 10
```

### Formato de Números

```python
pi = 3.14159

print(f"{pi:.2f}")      # 3.14 (2 decimales)
print(f"{pi:.4f}")      # 3.1416 (4 decimales)

num = 1234567
print(f"{num:,}")       # 1,234,567 (con comas)
print(f"{num:_}")       # 1_234_567 (con guiones)

# Porcentajes
porcentaje = 0.85
print(f"{porcentaje:.1%}")  # 85.0%
```

---

## Parte VII: Conversión de Tipos (Casting)

### String a Número

```python
texto = "42"
numero = int(texto)
print(numero + 8)  # 50

precio = "19.99"
flotante = float(precio)
print(flotante * 2)  # 39.98

# Si el formato es inválido:
int("abc")  # ValueError: invalid literal for int() with base 10: 'abc'
```

### Número a String

```python
edad = 25
texto = str(edad)
print(tipo(texto))  # <class 'str'>

print(f"Tienes {texto} años")  # Tienes 25 años
```

### A Boolean

```python
print(bool(1))          # True
print(bool(0))          # False
print(bool("texto"))    # True
print(bool(""))         # False
print(bool([1, 2]))     # True
print(bool([]))         # False (lista vacía es falsy)
```

### Entre Números

```python
# Int a Float
x = 5
y = float(x)
print(y)  # 5.0

# Float a Int (trunca)
pi = 3.14159
entero = int(pi)
print(entero)  # 3
```

---

## Parte VIII: Constantes

Python no tiene constantes de verdad, pero por convención usas MAYÚSCULAS:

```python
# Constantes del programa
MAX_INTENTOS = 3
MIN_EDAD = 18
DIRECTORIO_DATOS = "/datos"
PI = 3.14159

# Variables
intentos_restantes = MAX_INTENTOS
edad_usuario = 25

if edad_usuario < MIN_EDAD:
    print("Muy joven")
```

---

## Parte IX: Truthy y Falsy

En Python, casi todo puede evaluarse como `True` o `False`:

### Falsy (se consideran False)

```python
print(bool(0))              # False
print(bool(0.0))            # False
print(bool(""))             # False (string vacío)
print(bool([]))             # False (lista vacía)
print(bool({}))             # False (dict vacío)
print(bool(None))           # False (ninguno/nada)
print(bool(False))          # False
```

### Truthy (se consideran True)

```python
print(bool(1))              # True (cualquier número distinto de 0)
print(bool(-5))             # True
print(bool("texto"))        # True (cualquier string no vacío)
print(bool([1]))            # True (lista no vacía)
print(bool({"a": 1}))       # True (dict no vacío)
print(bool(True))           # True
```

### En Condicionales

```python
# Esto funciona:
if []:
    print("No se imprime")  # Lista vacía = falsy

if "":
    print("No se imprime")  # String vacío = falsy

if 0:
    print("No se imprime")  # 0 = falsy

# Esto sí:
if [1, 2]:
    print("Se imprime")     # Lista no vacía = truthy

if "hola":
    print("Se imprime")     # String no vacío = truthy

if 42:
    print("Se imprime")     # Número != 0 = truthy
```

---

## Parte X: Mutable vs Immutable

### Immutable (no se pueden cambiar después de crear)

```python
# String
s = "Hola"
# s[0] = "J"  # Error
s = "Jola"  # OK: crear nueva

# Tupla
t = (1, 2, 3)
# t[0] = 10  # Error
t = (10, 2, 3)  # OK

# Número
x = 5
# No puedes "modificar" el 5
x = 10  # OK: asignar nueva variable
```

### Mutable (SÍ se pueden cambiar)

```python
# Lista
lista = [1, 2, 3]
lista[0] = 10   # OK: modificar elemento
lista.append(4) # OK: agregar elemento

# Diccionario
d = {"nombre": "Juan"}
d["nombre"] = "Pedro"  # OK: cambiar
d["edad"] = 25         # OK: agregar

# Conjunto
s = {1, 2, 3}
s.add(4)       # OK
s.remove(1)    # OK
```

**Importante**: Esto afecta cómo funcionan los parámetros de funciones:

```python
def modificar(x):
    x[0] = 999  # Modifica la lista original

lista = [1, 2, 3]
modificar(lista)
print(lista)  # [999, 2, 3]
```

---

## Parte XI: Variables Especiales de Python

### None

Representa "nada":

```python
resultado = None
print(resultado)        # None
print(type(resultado))  # <class 'NoneType'>

# Verificar:
if resultado is None:
    print("No hay resultado")

# Falsy:
if not resultado:
    print("Falsy también")
```

### Reserved Keywords

Palabras que no puedes usar como nombres:

```
False, None, True, and, as, assert, async, await, break, 
class, continue, def, del, elif, else, except, finally, 
for, from, global, if, import, in, is, lambda, nonlocal, 
not, or, pass, raise, return, try, while, with, yield
```

---

## Resumen

| Concepto | Uso | Ejemplo |
|----------|-----|---------|
| Variable | Almacenar valor | `nombre = "Juan"` |
| snake_case | Convención Python | `mi_variable` |
| Múltiple | Varios a la vez | `a, b = 1, 2` |
| Operadores | +=, -=, etc. | `x += 5` |
| Scope | Dónde acceder | local, global |
| F-strings | Formateo | `f"Hola {x}"` |
| Casting | Convertir tipos | `int("5")` |
| Truthy/Falsy | Valores booleanos | `bool([])` = False |
| Mutable | Se pueden cambiar | listas, dicts |

---

**Siguiente**: [tipos_datos.md](tipos_datos.md)
