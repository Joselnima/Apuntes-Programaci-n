# Operadores en Python

> **Símbolos para comparar, calcular y tomar decisiones**

---

## ¿QUÉ es un Operador?

Un **operador** = símbolo que **transforma** datos:

```
3  +  5
↑  ↑  ↑
dato operador dato = RESULTADO
```

**Analogía:**
- 🧮 Operador aritmetico = calculadora (suma, resta, etc)
- 🔄 Operador comparación = balanza (¿cuál es más grande?)
- ⚡ Operador lógico = interruptor (sí/no, encendido/apagado)

---

## ¿PARA QÉ sirven?

### 1. Cálculos matemáticos
```python
total_productos = 5 + 3 + 2  # Contar items
precio_final = precio * 1.21  # Aplicar impuesto
```

### 2. Tomar decisiones (comparar)
```python
edad = 25
if edad >= 18:  # ¿Es adulto? (comparar)
    print("Puedes votar")
```

### 3. Validar condiciones lógicas
```python
tiene_dinero = True
tiene_entrada = True
if tiene_dinero and tiene_entrada:  # ¿Cumple TODO?
    print("¡Puedes entrar al cine!")
```

---

## ¿CÓMO funcionan?

### 4 Tipos de Operadores

| Tipo | Propósito | Ejemplo | Resultado |
|------|-----------|---------|-----------|
| **Aritméticos** | Calcular números | `10 + 3` | `13` |
| **Comparación** | Verdadero/Falso | `10 > 3` | `True` |
| **Lógicos** | AND, OR, NOT | `True and False` | `False` |
| **Asignación** | Guardar valores | `x = 5` | `x` vale `5` |

---

## Parte I: Operadores Aritméticos

Los básicos que conoces desde matemáticas:

```python
# Suma
print(5 + 3)        # 8

# Resta
print(10 - 4)       # 6

# Multiplicación
print(4 * 3)        # 12

# División (retorna float)
print(10 / 3)       # 3.333...

# División entera (trunca decimales)
print(10 // 3)      # 3

# Módulo (resto)
print(10 % 3)       # 1

# Potencia
print(2 ** 8)       # 256

# Negación
print(-5)           # -5
```

### Orden de operaciones

```python
# Misma que matemáticas: * / % antes que + -
resultado = 2 + 3 * 4  # 14, no 20

# Paréntesis funcionan igual
resultado = (2 + 3) * 4  # 20
```

---

## Parte II: Operadores de Comparación

Devuelven `True` o `False`:

```python
# Igual
print(5 == 5)       # True
print(5 == 4)       # False

# No igual
print(5 != 4)       # True
print(5 != 5)       # False

# Mayor que
print(5 > 3)        # True
print(5 > 5)        # False

# Mayor o igual
print(5 >= 5)       # True
print(5 >= 6)       # False

# Menor que
print(3 < 5)        # True
print(5 < 5)        # False

# Menor o igual
print(3 <= 5)       # True
print(5 <= 5)       # True
```

### Con Strings

```python
print("a" == "a")       # True
print("apple" < "banana")  # True (orden alfabético)
print("A" < "a")        # True (mayúsculas vienen primero)
```

---

## Parte III: Operadores Lógicos

Combinan condiciones booleanas:

### and (Y - AMBAS deben ser True)

```python
print(True and True)    # True
print(True and False)   # False
print(False and False)  # False

# Uso práctico
edad = 25
tiene_licencia = True

if edad >= 18 and tiene_licencia:
    print("Puede conducir")
```

### or (O - AL MENOS UNA debe ser True)

```python
print(True or False)    # True
print(False or False)   # False
print(True or True)     # True

# Uso práctico
es_fin_de_semana = False
tiene_día_libre = True

if es_fin_de_semana or tiene_día_libre:
    print("No hay que trabajar")
```

### not (NO - Invierte el booleano)

```python
print(not True)         # False
print(not False)        # True

# Uso práctico
activo = True
if not activo:
    print("Usuario inactivo")
else:
    print("Usuario activo")
```

### Cortocircuito

```python
# AND: Si primero es False, no evalúa segundo
print(False and (1 / 0))  # False (no da error)

# OR: Si primero es True, no evalúa segundo
print(True or (1 / 0))    # True (no da error)
```

---

## Parte IV: Operadores de Asignación

Combinan asignación con otra operación:

```python
x = 10

# +=  (suma y asigna)
x += 5      # x = x + 5  →  x = 15

# -=
x -= 3      # x = x - 3  →  x = 12

# *=
x *= 2      # x = x * 2  →  x = 24

# /=
x /= 4      # x = x / 4  →  x = 6.0

# //=
x //= 2     # x = x // 2 →  x = 3.0

# %=
x %= 2      # x = x % 2  →  x = 1.0

# **=
x **= 2     # x = x ** 2 →  x = 1.0
```

---

## Parte V: Operadores de Identidad

Comparan si dos variables apuntan al **mismo objeto en memoria**:

### is (ES - misma identidad)

```python
a = [1, 2, 3]
b = [1, 2, 3]
c = a

print(a == b)    # True (mismo contenido)
print(a is b)    # False (objetos diferentes)
print(a is c)    # True (mismo objeto)

# Con None (siempre usa 'is')
x = None
if x is None:    # ✅ Correcto
    print("x es None")
    
if x == None:    # Funciona, pero no es Pythónico
    print("x es None")
```

### is not

```python
if x is not None:
    print("x tiene valor")
```

---

## Parte VI: Operadores de Membresía

Verifican si un elemento está **en** una colección:

### in

```python
# Strings
print("a" in "hola")        # True
print("x" in "hola")        # False

# Listas
print(2 in [1, 2, 3])       # True
print(5 in [1, 2, 3])       # False

# Tuplas
print("Juan" in ("Juan", "Pedro", "Luis"))  # True

# Diccionarios (busca en CLAVES)
persona = {"nombre": "Juan", "edad": 25}
print("nombre" in persona)  # True
print("Juan" in persona)    # False

# Sets
print(3 in {1, 2, 3, 4})    # True
```

### not in

```python
if "admin" not in usuarios:
    print("Usuario no es admin")
```

---

## Parte VII: Operadores Bitwise (Avanzado)

Operan a nivel de bits (0s y 1s):

```python
a = 5      # 0101 en binario
b = 3      # 0011 en binario

print(a & b)    # 1 (AND)    → 0001
print(a | b)    # 7 (OR)     → 0111
print(a ^ b)    # 6 (XOR)    → 0110
print(~a)       # -6 (NOT)
print(a << 1)   # 10 (desplazar izquierda)
print(a >> 1)   # 2 (desplazar derecha)
```

---

## Parte VIII: Precedencia de Operadores

El orden en que Python evalúa múltiples operadores:

```
1. ()              Paréntesis
2. **              Potencia
3. +x, -x, ~x      Unarios
4. *, /, //, %     Multiplicación, división
5. +, -            Suma, resta
6. <<, >>          Bitwise shift
7. &               Bitwise AND
8. ^               Bitwise XOR
9. |               Bitwise OR
10. ==, !=, <, >, <=, >=, is, in  Comparación
11. not            Lógico NOT
12. and            Lógico AND
13. or             Lógico OR
```

### Ejemplo

```python
resultado = 2 + 3 * 4 ** 2 - 5
#          = 2 + 3 * 16 - 5
#          = 2 + 48 - 5
#          = 45
```

---

## Parte IX: Operador Condicional Ternario

Forma compacta de if/else:

```python
# Sintaxis: valor_si_verdadero if condición else valor_si_falso

edad = 25
estatus = "adulto" if edad >= 18 else "menor"
print(estatus)  # adulto

# Anidado
nota = 85
calificación = "A" if nota >= 90 else "B" if nota >= 80 else "C"
print(calificación)  # B
```

---

## Parte X: Comparaciones Encadenadas

```python
x = 5

# De una forma
if x > 0 and x < 10:
    print("Entre 1 y 9")

# De forma elegante (Python)
if 0 < x < 10:
    print("Entre 1 y 9")

# Más comparaciones
if 0 <= x <= 100:
    print("Entre 0 y 100 inclusive")
```

---

## Parte XI: Operadores sobre Strings y Listas

```python
# Concatenación
print("Hola " + "Mundo")    # Hola Mundo

# Repetición
print("Hola " * 3)          # Hola Hola Hola

# Concatenación de listas
print([1, 2] + [3, 4])      # [1, 2, 3, 4]

# Repetición de listas
print([0] * 5)              # [0, 0, 0, 0, 0]
```

---

## Resumen - Tabla de Operadores

| Tipo | Operadores | Ejemplo |
|------|-----------|---------|
| Aritméticos | +, -, *, /, //, %, ** | 5 + 3 = 8 |
| Comparación | ==, !=, <, >, <=, >= | 5 > 3 = True |
| Lógicos | and, or, not | True and False = False |
| Asignación | =, +=, -=, *=, /=, //=, %=, **= | x += 5 |
| Identidad | is, is not | a is b |
| Membresía | in, not in | 2 in [1,2,3] = True |
| Bitwise | &, \|, ^, ~, <<, >> | 5 & 3 = 1 |

---

**Siguiente**: [control_flujo.md](control_flujo.md)
