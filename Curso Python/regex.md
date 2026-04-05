# Expresiones Regulares (Regex)

> **Patrón para encontrar, validar y transformar texto rápidamente**

---

## ¿QUÉ son las Regex?

Una **regex** = descripción de un PATRÓN de texto:

```
Email válido tiene: letras@letras.letras

Regex: \w+@\w+\.\w+

\w+ = una o más letras/números
@   = símbolo @
\w+ = una o más letras/números
\.  = símbolo . (punto)
\w+ = una o más letras/números
```

**Analogía:**
- 🔍 Búsqueda normal = "Encuentra la palabra 'perro'"
- 🎯 Regex = "Encuentra cualquier palabra de 5 letras seguida de número"

---

## ¿PARA QÉ sirven?

### 1. Validar email
```python
email = "juan@example.com"

# ❌ Manual (complicado)
if "@" in email and "." in email:
    print("Probablemente válido")

# ✅ Regex (preciso)
import re
if re.match(r'\w+@\w+\.\w+', email):
    print("✅ Válido")
```

### 2. Buscar patrones
```python
texto = "Mi teléfono es 555-1234 y 555-5678"

# Encontrar TODOS los teléfonos
números = re.findall(r'\d{3}-\d{4}', texto)
# ['555-1234', '555-5678']
```

### 3. Limpiar texto
```python
# Eliminar números
texto = "Hola 123 mundo 456"
limpio = re.sub(r'\d+', '', texto)
# "Hola  mundo "
```

### 4. Extraer información
```python
# Extraer fecha
fecha_texto = "Nacido: 25/12/1990"
fecha = re.search(r'(\d{2})/(\d{2})/(\d{4})', fecha_texto)
print(fecha.group())      # "25/12/1990"
print(fecha.group(1))     # "25" (día)
print(fecha.group(2))     # "12" (mes)
print(fecha.group(3))     # "1990" (año)
```

---

## ¿CÓMO funcionan?

### Componentes Básicos

| Patrón | Significa | Ejemplo |
|--------|-----------|---------|
| `.` | Cualquier carácter | `g.to` → "gato", "gito" |
| `\d` | Dígito (0-9) | `\d{3}` → "123" |
| `\w` | Letra/número/_ | `\w+` → "Hola123_x" |
| `\s` | Espacio/tab/salto | `Hola\s` → "Hola " |
| `^` | Inicio de string | `^Hola` → debe empezar |
| `$` | Final de string | `adiós$` → debe terminar |
| `*` | 0 o más | `ab*c` → "ac", "abc", "abbc" |
| `+` | 1 o más | `ab+c` → "abc", "abbc" (no "ac") |
| `?` | 0 o 1 | `ab?c` → "ac", "abc" (no "abbc") |

**Ejemplo de Lectura:**
```
\d{3}-\d{4}
↓
"3 dígitos - 4 dígitos"
↓
"555-1234" ✅
"55-1234"  ❌ (solo 2 dígitos)
```

---

## Parte I: Básico - Caracteres y Clases

```python
import re

# Búsqueda simple
re.search('gato', 'El gato es lindo')  # Encontrado

# . (punto) = cualquier carácter
re.search('g.to', 'gato')   # Match
re.search('g.to', 'gito')   # Match
re.search('g.to', 'git')    # No match

# \d = dígito
re.search(r'\d', 'Año 2024')        # Match '2'
re.search(r'\d+', 'Año 2024')       # Match '2024'

# \w = letra/número/_ (word character)
re.search(r'\w', 'Hola!')           # Match 'H'
re.search(r'\w+', 'Hola!')          # Match 'Hola'

# \s = espacio/tab/newline
re.search(r'Hola\s', 'Hola ')       # Match

# ^ = inicio
re.search(r'^gato', 'gato lindo')   # Match
re.search(r'^gato', 'el gato')      # No match

# $ = final
re.search(r'lindo$', 'gato lindo')  # Match
re.search(r'lindo$', 'lindo gato')  # No match

# * = 0 o más
re.search(r'go*t', 'gt')            # Match (0 o's)
re.search(r'go*t', 'got')           # Match (1 o)
re.search(r'go*t', 'gooot')         # Match (3 o's)

# + = 1 o más
re.search(r'go+t', 'gt')            # No match
re.search(r'go+t', 'got')           # Match

# ? = 0 o 1
re.search(r'goa?t', 'gt')           # Match
re.search(r'goa?t', 'goat')         # Match
re.search(r'goa?t', 'goaat')        # No match

# {n} = exactamente n
re.search(r'a{3}', 'baaab')         # Match 'aaa'

# {n,m} = entre n y m
re.search(r'a{2,4}', 'baaaab')      # Match 'aaaa'
```

---

## Parte II: Clases de Caracteres

```python
import re

# [ ] = uno de los caracteres
re.search(r'[aeiou]', 'hola')       # Match 'o'
re.search(r'[0-9]', 'año2024')      # Match '2'
re.search(r'[a-z]', 'ABC')          # No match

# [^] = NEGACIÓN (cualquier excepto)
re.search(r'[^0-9]', '12345abc')    # Match 'a' (primer no dígito)
re.search(r'[^aeiou]', 'aaa')       # Match... nada (solo vocales)

# Rangos comunes
r'[A-Z]'    # Mayúscula
r'[a-z]'    # Minúscula
r'[0-9]'    # Dígito
r'[a-zA-Z0-9_]'  # Alfanumérico + underscore (= \w)
r'[^a-zA-Z0-9_ ]'  # Caracteres especiales
```

---

## Parte III: Funciones Principales

### search() - Primera coincidencia

```python
import re

match = re.search(r'\d+', 'El año es 2024')
if match:
    print(match.group())    # '2024'
    print(match.start())    # Posición inicio
    print(match.end())      # Posición fin
```

### findall() - Todas las coincidencias

```python
import re

números = re.findall(r'\d+', 'Tengo 25 años, nací en 1999')
print(números)  # ['25', '1999']

emails = re.findall(r'\w+@\w+\.\w+', 
    'Juan juan@test.com María maría@example.com')
print(emails)  # ['juan@test.com', 'maría@example.com']
```

### match() - Desde el inicio

```python
import re

# search busca en cualquier lugar
re.search(r'gato', 'el gato')   # Match

# match solo desde inicio
re.match(r'gato', 'el gato')    # No match
re.match(r'el', 'el gato')      # Match
```

### sub() - Reemplazar

```python
import re

# Reemplazar simple
nuevo = re.sub(r'\d', 'X', 'Mi número es 12345')
print(nuevo)  # 'Mi número es XXXXX'

# Reemplazar con límite
nuevo = re.sub(r'\d', 'X', 'Es 123 y 456', count=3)
print(nuevo)  # 'Es XXX y 456'

# Reemplazar con función
def mayúscula(match):
    return match.group().upper()

nuevo = re.sub(r'[aeiou]', mayúscula, 'hola mundo')
print(nuevo)  # 'hOlA mUndO'
```

### split() - Dividir por patrón

```python
import re

# Dividir por espacios en blanco
partes = re.split(r'\s+', 'Juan    María   Pedro')
print(partes)  # ['Juan', 'María', 'Pedro']

# Dividir por coma
partes = re.split(r',\s*', 'Juan, María, Pedro')
print(partes)  # ['Juan', 'María', 'Pedro']
```

---

## Parte IV: Grupos

Capturar partes específicas:

```python
import re

# Grupo () captura
email = 'juan@example.com'
match = re.search(r'(\w+)@(\w+\.\w+)', email)
print(match.group(0))   # juan@example.com (completo)
print(match.group(1))   # juan (usuario)
print(match.group(2))   # example.com (dominio)

# Grupos nombrados
match = re.search(r'(?P<usuario>\w+)@(?P<dominio>\w+\.\w+)', email)
print(match.group('usuario'))   # juan
print(match.group('dominio'))   # example.com

# findall con grupos
números = re.findall(r'(\d+)-(\d+)', '123-456 y 789-000')
print(números)  # [('123', '456'), ('789', '000')]
```

---

## Parte V: Validaciones Comunes

```python
import re

def validar_email(email):
    patrón = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
    return bool(re.match(patrón, email))

def validar_teléfono(teléfono):
    patrón = r'^(\+\d{1,2}\s?)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$'
    return bool(re.match(patrón, teléfono))

def validar_url(url):
    patrón = r'^https?://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/\S*)?$'
    return bool(re.match(patrón, url))

def validar_contraseña(contraseña):
    # Mínimo 8 caracteres, mayúscula, minúscula, número
    patrón = r'^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$'
    return bool(re.match(patrón, contraseña))

# Pruebas
print(validar_email("juan@example.com"))      # True
print(validar_email("juan@"))                 # False
print(validar_contraseña("Abc123456"))       # True
print(validar_contraseña("abcdefgh"))        # False
```

---

## Parte VI: Lookahead y Lookbehind

Afirmaciones sin consumir caracteres:

```python
import re

texto = "precio: $100 USD y €50 EUR"

# Lookahead positivo (?=...)
números_con_$ = re.findall(r'\d+(?=\s*USD)', texto)
print(números_con_$)  # ['100']

# Lookahead negativo (?!...)
números_sin_$ = re.findall(r'\d+(?!\s*USD)', texto)
print(números_sin_$)  # ['50']

# Lookbehind positivo (?<=...)
números_después_$ = re.findall(r'(?<=\$)\d+', texto)
print(números_después_$)  # ['100']

# Lookbehind negativo (?<!...)
números_no_$ = re.findall(r'(?<!\$)\d+', texto)
print(números_no_$)  # ['50']
```

---

## Parte VII: Flags

Modificadores de comportamiento:

```python
import re

# re.IGNORECASE - sin diferenciar mayúscula
re.search(r'gato', 'GATO', flags=re.IGNORECASE)  # Match

# re.MULTILINE - ^ y $ por línea
texto = "gato\ngato\ngato"
matches = re.findall(r'^gato$', texto, flags=re.MULTILINE)
print(len(matches))  # 3

# re.DOTALL - . incluye newline
re.search(r'a.b', 'a\nb', flags=re.DOTALL)  # Match

# re.VERBOSE - comentarios
patrón = r'''
    (\w+)      # Usuario
    @          # Arroba
    (\w+)      # Dominio
    \.         # Punto
    (\w+)      # Extensión
'''
re.search(patrón, 'juan@example.com', re.VERBOSE)  # Match
```

---

## Parte VIII: Object Pattern (Compiled)

Compilar regex para reutilizar:

```python
import re

# ✅ BIEN - patrón compilado (más rápido si se usa múltiples veces)
patrón = re.compile(r'\d+')
números = patrón.findall('2024, 2025, 2026')
# Múltiples operaciones sin recompilar
print(patrón.search('Año 2024'))
print(patrón.sub('X', '2024 y 2025'))
```

---

## Parte IX: Ejemplo Real - Extracción de Datos

```python
import re

html = """
<div class="usuario">
    <name>Juan García</name>
    <email>juan@example.com</email>
    <phone>555-123-4567</phone>
</div>
<div class="usuario">
    <name>María López</name>
    <email>maria@test.com</email>
    <phone>555-987-6543</phone>
</div>
"""

# Extraer usuarios
usuarios = re.findall(
    r'<name>([^<]+)</name>\s*'
    r'<email>([^<]+)</email>\s*'
    r'<phone>([^<]+)</phone>',
    html
)

for nombre, email, teléfono in usuarios:
    print(f"{nombre} - {email} - {teléfono}")

# Salida:
# Juan García - juan@example.com - 555-123-4567
# María López - maria@test.com - 555-987-6543
```

---

## Parte X: Tabla de Caracteres Especiales

| Símbolo | Significa |
|---------|-----------|
| `.` | Cualquier carácter |
| `^` | Inicio de línea |
| `$` | Final de línea |
| `*` | 0 o más |
| `+` | 1 o más |
| `?` | 0 o 1 |
| `{n}` | Exactamente n |
| `{n,m}` | Entre n y m |
| `\d` | Dígito [0-9] |
| `\w` | Palabra [a-zA-Z0-9_] |
| `\s` | Espacio |
| `\\` | Escape literal |
| `[ ]` | Clase de caracteres |
| `[^ ]` | Negación |
| `( )` | Grupo/Captura |
| `\|` | O (alternativa) |

---

## Parte XI: Problemas Comunes

```python
import re

# ❌ Olvidar raw string
# re.search('\d', "123")  # Puede fallar

# ✅ Usar raw string
re.search(r'\d', "123")  # OK

# ❌ Regex muy compleja
patrón = r'^(?:[a-zA-Z0-9+._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}|[a-zA-Z0-9.!#$%&\'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")$'

# ✅ Usar bibliotecas
from email_validator import validate_email
validate_email("juan@example.com")
```

---

## Resumen

| Función | Uso |
|---------|-----|
| `search()` | Primera coincidencia |
| `findall()` | Todas las coincidencias |
| `match()` | Desde inicio |
| `sub()` | Reemplazar |
| `split()` | Dividir |
| `compile()` | Compilar para reutilizar |
| Grupos | Capturar partes |
| Lookahead/Lookbehind | Afirmaciones sin consumir |
| Flags | Modificadores (IGNORECASE, etc) |

---

**Siguiente**: [performance.md](performance.md)
