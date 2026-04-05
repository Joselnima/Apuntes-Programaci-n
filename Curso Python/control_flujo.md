# Control de Flujo: if, elif, else

> **Cómo enseñarle a tu programa a tomar decisiones - El programa hace lo diferente según lo que pase**

---

## ¿QÉ es Control de Flujo?

Es como un **árbol de decisiones**:

```
        ¿Es mayor de 18?
         /            \
       SÍ              NO
       |               |
    Votar          Esperar
```

En la vida real:
- Si llueve → llevo paraguas
- Si NO llueve → llevo gafas de sol
- Si es muy calor → bebo agua

En Python:
```python
if llueve:
    llevar(paraguas)
elif es_calor:
    beber(agua)
else:
    llevar(gafas_sol)
```

---

## ¿PARA QÉ sirve?

### Sin Control de Flujo (Inútil)

```python
# El programa siempre hace LO MISMO, sin importar qué pase
edad = 25
print("Eres mayor")  # Siempre imprime esto
print("Puedes votar")  # Siempre imprime esto

edad = 15
print("Eres mayor")  # ¿Pero aquí también? ¡INCORRECTO!
```

### Con Control de Flujo (Inteligente)

```python
edad = 25
if edad >= 18:
    print("Eres mayor")
    print("Puedes votar")

edad = 15
if edad >= 18:
    print("Eres mayor")
    print("Puedes votar")
else:
    print("Eres menor")
    print("No puedes votar")
```

El programa se adapta 🎯

---

## ¿CÓMO funciona?

### El concepto de "Condición"

Una **condición** es una pregunta que Python responde con SÍ (True) o NO (False):

```python
edad = 25

edad >= 18         # ¿Es 25 mayor o igual a 18? SÍ (True)
edad < 18          # ¿Es 25 menor que 18? NO (False)
edad == 25         # ¿Es 25 igual a 25? SÍ (True)
nombre == "Juan"   # ¿El nombre es "Juan"? SÍ o NO
saldo > 0          # ¿El saldo es positivo? SÍ o NO
```

### Los Operadores de Comparación

| Símbolo | Significado | Ejemplo |
|---------|-------------|---------|
| `==` | ¿Igual a? | `5 == 5` → True |
| `!=` | ¿NO igual a? | `5 != 3` → True |
| `>` | ¿Mayor que? | `5 > 3` → True |
| `<` | ¿Menor que? | `5 < 3` → False |
| `>=` | ¿Mayor o igual? | `5 >= 5` → True |
| `<=` | ¿Menor o igual? | `5 <= 5` → True |

---

## Parte I: if - La Decisión Más Simple

---

## Parte II: if/else (Dos Caminos)

Condición verdadera → un camino; falsa → otro:

```python
edad = 15

if edad >= 18:
    print("Puedes votar")
else:
    print("Eres menor de edad")
```

---

## Parte III: if/elif/else (Múltiples Caminos)

Más de dos opciones:

```python
nota = 85

if nota >= 90:
    print("A - Excelente")
elif nota >= 80:
    print("B - Bueno")
elif nota >= 70:
    print("C - Aceptable")
elif nota >= 60:
    print("D - Pasó")
else:
    print("F - Reprobado")
```

**Python evalúa de arriba hacia abajo y se detiene en la primera que es `True`:**

```python
nota = 85

if nota >= 80:      # ✅ Verdadero
    print("B")      # Se ejecuta
elif nota >= 70:    # No se evalúa (ya encontró un True)
    print("C")
```

---

## Parte IV: Anidamiento

Condicionales dentro de condicionales:

```python
edad = 25
tiene_licencia = True

if edad >= 18:
    if tiene_licencia:
        print("Puede conducir")
    else:
        print("Mayor pero sin licencia")
else:
    print("Menor de edad")

# Mejor con 'and':
if edad >= 18 and tiene_licencia:
    print("Puede conducir")
```

---

## Parte V: Operador Ternario

Forma compacta para if/else simple:

```python
edad = 25

# Forma larga
if edad >= 18:
    estatus = "adulto"
else:
    estatus = "menor"

# Forma compacta (ternario)
estatus = "adulto" if edad >= 18 else "menor"
print(estatus)
```

### Anidado

```python
nota = 75

calificación = "A" if nota >= 90 else \
               "B" if nota >= 80 else \
               "C" if nota >= 70 else \
               "F"

print(calificación)  # C
```

---

## Parte VI: Truthy y Falsy (Decisiones Implícitas)

Casi todo puede ser condición:

### Falsy (se comporta como False)

```python
if 0:               # Falsy
    print("No se ejecuta")

if "":              # String vacío: Falsy
    print("No se ejecuta")

if []:              # Lista vacía: Falsy
    print("No se ejecuta")

if None:            # None: Falsy
    print("No se ejecuta")

if False:           # Obviamente Falsy
    print("No se ejecuta")
```

### Truthy (se comporta como True)

```python
if 1:               # Truthy (cualquier número != 0)
    print("Se ejecuta")

if "texto":         # String no vacío: Truthy
    print("Se ejecuta")

if [1, 2]:          # Lista no vacía: Truthy
    print("Se ejecuta")

if True:            # Obviamente Truthy
    print("Se ejecuta")
```

### Uso Práctico

```python
# Verificar si lista tiene elementos
items = [1, 2, 3]
if items:
    print("Lista tiene elementos")

# Mejor que:
if len(items) > 0:
    print("Lista tiene elementos")

# Verificar si string no está vacío
nombre = input("Nombre: ")
if nombre:
    print(f"Hola {nombre}")
else:
    print("No ingresaste nombre")
```

---

## Parte VII: not, and, or en Condicionales

### not (Negación)

```python
está_lloviendo = False

if not está_lloviendo:
    print("Puedo salir sin sombrilla")

# Equivalente a:
if está_lloviendo == False:
    print("Puedo salir sin sombrilla")

# Con None
resultado = None
if not resultado:
    print("No hay resultado")
```

### and (Ambas deben ser True)

```python
edad = 25
tiene_empleo = True

if edad >= 18 and tiene_empleo:
    print("Puede obtener crédito")

# Si la primera es False, no evalúa la segunda
if False and (1 / 0):  # No da error
    pass
```

### or (Al menos una debe ser True)

```python
es_fin_de_semana = False
tiene_día_libre = True

if es_fin_de_semana or tiene_día_libre:
    print("No trabajo hoy")

# Si la primera es True, no evalúa la segunda
if True or (1 / 0):  # No da error
    pass
```

---

## Parte VIII: Match/Case (Python 3.10+)

Similar a switch en otros lenguajes:

```python
día = 3

match día:
    case 1:
        print("Lunes")
    case 2:
        print("Martes")
    case 3:
        print("Miércoles")
    case 4:
        print("Jueves")
    case 5:
        print("Viernes")
    case 6 | 7:  # 6 O 7
        print("Fin de semana")
    case _:      # Default
        print("Día inválido")
```

### Con Patrones

```python
punto = (0, 0)

match punto:
    case (0, 0):
        print("Origen")
    case (0, y):
        print(f"En eje Y: {y}")
    case (x, 0):
        print(f"En eje X: {x}")
    case (x, y):
        print(f"Punto ({x}, {y})")
```

---

## Parte IX: Comparaciones Compuestas

```python
x = 5

# Forma 1: Múltiples comparaciones
if x > 0 and x < 10:
    print("Entre 1 y 9")

# Forma 2: Encadenada (Pythónica)
if 0 < x < 10:
    print("Entre 1 y 9")

# Más encadenadas
if 0 <= x <= 100:
    print("Entre 0 y 100 inclusive")

# Con diferentes operadores
if 0 < x == 5:
    print("x es 5 y positivo")
```

---

## Parte X: Estrutura Completa Ejemplo

```python
def categorizar_edad(edad):
    \"\"\"Categoriza edad de una persona\"\"\"
    
    if edad < 0:
        print("Edad inválida")
    elif edad < 13:
        print("Niño")
    elif edad < 18:
        print("Adolescente")
    elif edad < 65:
        print("Adulto")
    else:
        print("Jubilato")

# Uso
categorizar_edad(25)  # Adulto
categorizar_edad(12)  # Niño
categorizar_edad(70)  # Jubilado
```

---

## Resumen

| Sentencia | Uso | Ejemplo |
|-----------|-----|---------|
| if | Si verdadero | if x > 5: print("Sí") |
| else | Si falso | else: print("No") |
| elif | Otra condición | elif x == 5: print("Igual") |
| Ternario | Compacto if/else | x if condición else y |
| and | Ambas verdaderas | if a and b: |
| or | Al menos una verdadera | if a or b: |
| not | Negación | if not x: |
| match/case | Múltiples opciones | match x: case 1: ... |

---

**Siguiente**: [bucles_iteraciones.md](bucles_iteraciones.md)
