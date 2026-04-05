# Cadenas de Texto (Strings)

> **Cómo trabajar con texto - Quizá hagas 50% de tu programación con strings**

---

## ¿QÉ es un String?

Un **string** (cadena de texto) es simplemente **texto**:

```python
nombre = "Juan"
mensaje = "Hola, ¿cómo estás?"
linea = "Python es genial"
```

Es una **secuencia de caracteres**:
```
"Hola"
 ↓ ↓ ↓ ↓
 H o l a
```

Cada carácter tiene una posición:
```
"Hola"
 0 1 2 3
 H o l a
```

---

## ¿PARA QÉ sirven?

### 1. Mostrar mensajes
```python
print("Bienvenido al programa")
```

### 2. Guardar nombres y datos
```python
usuario = input("¿Cuál es tu nombre? ")
idioma = "Python"
ciudad = "Madrid"
```

### 3. Procesar información
```python
email = "juan@gmail.com"
if "@" in email:
    print("Email válido")
```

### 4. Manipular texto
```python
frase = "python"
frase_bonita = frase.upper()  # "PYTHON"
frase_limpia = frase.strip()  # Quita espacios
```

---

## ¿CÓMO funcionan?

### Creación: Lo Básico

```python
# Python entiende automáticamente que es texto (cuando está entre comillas)
nombre = "Juan"      # String
edad = "25"          # TEXTO (no número)
numero = 25          # Número (sin comillas)

# Son DIFERENTES
print(type("25"))    # <class 'str'> (texto)
print(type(25))      # <class 'int'> (número)
```

### Acceso: Tomar caracteres

```python
palabra = "Python"
#          0123456  (posiciones)

print(palabra[0])    # P (primer carácter)
print(palabra[1])    # y (segundo)
print(palabra[-1])   # n (último - contas desde atrás)
print(palabra[-2])   # o (penúltimo)
```

### Rebanada (Slice): Tomar pedazos

```python
texto = "Programacion"
#        012345678901  (posiciones)

print(texto[0:7])    # "Program" (del 0 al 7, excluye el 7)
print(texto[4:])     # "amacion" (del 4 hasta el final)
print(texto[:6])     # "Progra" (del inicio hasta el 6)
print(texto[-4:])    # "cion" (últimos 4)
```

---

## Parte I: Creación de Strings - Las Formas de Crear Texto

```python
nombre = "Juan"
nombre = 'Juan'

# Diferencia si necesitas comilla adentro
frase = "Dijo: 'Hola'"        # Comillas adentro: OK
frase = 'Dijo: "Hola"'        # Comillas adentro: OK
frase = "Dijo: \"Hola\""      # Escapar: OK
```

### Strings Multi-línea

```python
descripción = """
Primera línea
Segunda línea
Tercera línea
"""

# O con saltos explícitos
descripción = "Primera línea\nSegunda línea\nTercera línea"
```

### Raw Strings

No interpreta caracteres especiales (perfecto para rutas):

```python
# Sin r: interpreta \n como salto
ruta1 = "C:\Usuarios\Juan"  # Problema con \U y \J

# Con r: literal
ruta2 = r"C:\Usuarios\Juan"  # Funciona correctamente

# Perfecto para regex
patrón = r"\d+"  # Busca dígitos
```

---

## Parte II: Acceso e Indexing

### Acceso Individual

```python
texto = "Python"

print(texto[0])      # P
print(texto[1])      # y
print(texto[-1])     # n (último)
print(texto[-2])     # o (penúltimo)

# Out of range
# print(texto[10])   # IndexError
```

### Longitud

```python
texto = "Python"
print(len(texto))    # 6
```

### Iteración

```python
for letra in "Hola":
    print(letra)

# Output: H, o, l, a (cada una en línea)
```

---

## Parte III: Slicing (Rebanadas)

Obtén subsecciones:

```python
texto = "Python 3.12"

print(texto[0:6])        # Python (índices 0-5)
print(texto[7:])         # 3.12 (desde 7 hasta el final)
print(texto[:6])         # Python (desde inicio hasta 5)
print(texto[::2])        # Pto.1 (cada 2 caracteres)
print(texto[::-1])       # 21.3 nohtyP (invertido)
print(texto[-3:])        # .12 (últimos 3)
```

**Sintaxis**: `string[inicio:fin:paso]`
- `inicio`: Índice inicial (inclusivo), default 0
- `fin`: Índice final (EXCLUSIVO), default fin del string
- `paso`: Incremento, default 1

---

## Parte IV: Strings son Inmutables

No puedes cambiar un string una vez creado:

```python
texto = "Hola"

# ❌ Error
# texto[0] = "J"  # TypeError: 'str' object does not support item assignment

# ✅ Correcto: crear uno nuevo
texto = "Jola"
```

---

## Parte V: Operadores con Strings

### Concatenación

```python
saludo = "Hola" + " " + "Mundo"
print(saludo)  # Hola Mundo

# Múltiples
cadena = "A" + "B" + "C" + "D"
print(cadena)  # ABCD

# Con variables
nombre = "Juan"
mensaje = "Hola " + nombre  # Hola Juan
```

### Repetición

```python
print("Hola " * 3)      # Hola Hola Hola
print("-" * 20)         # --------------------

# Útil para separadores
print("Mesa de precios")
print("-" * 50)
```

### Pertenencia

```python
print("a" in "hola")        # True
print("x" in "hola")        # False
print("ola" in "hola")      # True

if "Juan" in "Mi nombre es Juan":
    print("Encontrado")
```

---

## Parte VI: Métodos Útiles

### .upper() y .lower()

```python
texto = "PyThOn"

print(texto.upper())    # PYTHON
print(texto.lower())    # python
print(texto.title())    # Python (primer letra de cada palabra)
print(texto.capitalize())  # Python (solo primer letra)
```

### .strip(), .lstrip(), .rstrip()

Elimina espacios:

```python
texto = "  Hola Mundo  "

print(texto.strip())    # Hola Mundo
print(texto.lstrip())   # Hola Mundo   (izquierda)
print(texto.rstrip())   #   Hola Mundo (derecha)

# También con caracteres específicos
código = "***Hola***"
print(código.strip("*"))  # Hola
```

### .replace()

```python
texto = "Hola Mundo"

print(texto.replace("Mundo", "Python"))  # Hola Python

# Reemplazar múltiple
texto = "aaa"
print(texto.replace("a", "b", 2))  # bba (solo 2 primeros)
```

### .split() y .join()

**split**: Divide string en lista

```python
frase = "Hola Mundo Python"
palabras = frase.split()         # ['Hola', 'Mundo', 'Python']

csv = "Juan,25,Madrid"
datos = csv.split(",")           # ['Juan', '25', 'Madrid']

# Con límite
partes = "a:b:c:d".split(":", 2)  # ['a', 'b', 'c:d']

# Líneas
documento = "Línea 1\nLínea 2\nLínea 3"
líneas = documento.split("\n")   # ['Línea 1', 'Línea 2', 'Línea 3']
```

**join**: Combina lista en string

```python
palabras = ["Hola", "Mundo", "Python"]
frase = " ".join(palabras)       # Hola Mundo Python

datos = ["Juan", "25", "Madrid"]
csv = ",".join(datos)            # Juan,25,Madrid

# Newlines
líneas = ["Primera", "Segunda", "Tercera"]
documento = "\n".join(líneas)    # Con saltos
```

### .find() y .index()

```python
texto = "Hola Mundo"

print(texto.find("Mundo"))      # 5 (índice)
print(texto.find("xyz"))        # -1 (no encontrado)

print(texto.index("Mundo"))     # 5
# print(texto.index("xyz"))     # ValueError (no encontrado)

# Busque desde posición específica
print(texto.find("o", 5))       # 7 (desde índice 5)
```

### .count()

```python
texto = "Hola Hola Hola"

print(texto.count("Hola"))      # 3
print(texto.count("a"))         # 3
```

### Verificación

```python
print("123".isdigit())          # True
print("abc".isalpha())          # True
print("abc123".isalnum())       # True
print("   ".isspace())          # True
print("".isempty())             # AttributeError (no existe)
print(len("") == 0)             # True (así se verifica)
```

### .startswith() y .endswith()

```python
archivo = "documento.pdf"

print(archivo.startswith("documento"))   # True
print(archivo.endswith(".pdf"))          # True

url = "https://ejemplo.com"
print(url.startswith("https"))           # True
```

---

## Parte VII: F-Strings (Python 3.6+)

La forma moderna e recomendada:

### Básico

```python
nombre = "Juan"
edad = 25

print(f"Me llamo {nombre} y tengo {edad} años")
```

### Con Expresiones

```python
a = 5
b = 3

print(f"{a} + {b} = {a + b}")
print(f"Es {a} > {b}? {a > b}")

lista = [1, 2, 3]
print(f"La lista tiene {len(lista)} elementos")
```

### Formato de Números

```python
pi = 3.14159265
print(f"{pi:.2f}")          # 3.14 (2 decimales)
print(f"{pi:.4f}")          # 3.1416 (4 decimales)

número = 1234567
print(f"{número:,}")        # 1,234,567
print(f"{número:_}")        # 1_234_567

porcentaje = 0.85
print(f"{porcentaje:.1%}")  # 85.0%
print(f"{porcentaje:.0%}")  # 85%
```

### Con Diccionarios

```python
persona = {"nombre": "Juan", "edad": 25}

print(f"Nombre: {persona['nombre']}")
print(f"Edad: {persona['edad']}")
```

---

## Parte VIII: Method Chaining

Encadenar métodos:

```python
texto = "  HOLA MUNDO  "

resultado = (texto
    .lower()           # hola mundo
    .strip()           # hola mundo
    .replace("mundo", "python")  # hola python
    .title()           # Hola Python
)

print(resultado)  # Hola Python
```

---

## Parte IX: Encoding y Unicode

```python
# UTF-8 es el estándar
texto = "Hola Mundo"
bytes_utf8 = texto.encode("utf-8")
print(bytes_utf8)  # b'Hola Mundo'

# Decodificar
restaurado = bytes_utf8.decode("utf-8")
print(restaurado)  # Hola Mundo

# Con caracteres especiales
español = "Niño Español"
print(español)
print(len(español))  # 13 caracteres
```

---

## Parte X: Comparación y Ordenamiento

```python
# Comparación alfabética
print("apple" < "banana")  # True
print("apple" == "apple")  # True

# Case-sensitive
print("A" < "a")           # True (mayúsculas primero)
print("ABC" > "ABC")       # False

# Ordenar strings
palabras = ["zebra", "apple", "banana"]
palabras.sort()
print(palabras)  # ['apple', 'banana', 'zebra']
```

---

## Parte XI: Expresiones Regulares (Regex)

Búsquedas avanzadas:

```python
import re

texto = "Mi email es test@ejemplo.com"

# Buscar patrón
match = re.search(r"\w+@\w+\.\w+", texto)
if match:
    print(f"Email encontrado: {match.group()}")
    # Output: Email encontrado: test@ejemplo.com

# Reemplazar
nuevo = re.sub(r"\d+", "X", "Tengo 25 años")
print(nuevo)  # Tengo X años
```

(Más en `regex.md`)

---

## Resumen de Métodos Útiles

| Método | Qué hace | Ejemplo |
|--------|----------|---------|
| .upper() | Mayúscula | "hola".upper() = "HOLA" |
| .lower() | Minúscula | "HOLA".lower() = "hola" |
| .strip() | Quita espacios | "  hola  ".strip() = "hola" |
| .replace() | Reemplaza | "hola".replace("a","o") = "holo" |
| .split() | Divide en lista | "a,b,c".split(",") = ['a','b','c'] |
| .join() | Combina lista | ",".join(['a','b']) = "a,b" |
| .find() | Encuentra índice | "hola".find("l") = 2 |
| .startswith() | ¿Empieza con? | "hola".startswith("ho") = True |
| .endswith() | ¿Termina con? | "hola".endswith("la") = True |
| .count() | Cuenta ocurrencias | "hola".count("l") = 1 |

---

**Siguiente**: [listas_tuplas_conjuntos.md](listas_tuplas_conjuntos.md)
