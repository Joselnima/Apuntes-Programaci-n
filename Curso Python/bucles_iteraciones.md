# Bucles e Iteraciones en Python

> **Cómo repetir acciones automáticamente - Haz que la computadora trabaje por ti**

---

## ¿QÉ es un Bucle?

Un **bucle** te permite **repetir el mismo código múltiples veces** sin escribirlo una y otra vez:

**Analogía del mundo real:**
- 📚 **Sin bucle**: Tienes que fotocopiar una hoja manualmente 100 veces (tedioso)
- 🏭 **Con bucle**: Pones la hoja y le dices a la máquina "copia 100 veces" (automatizado)

---

## ¿PARA QÉ sirven los Bucles?

### Caso 1: Procesar múltiples datos
```python
# ❌ Sin bucle (ridículo)
email_a = "juan@gmail.com"
enviar_email(email_a)

email_b = "maria@gmail.com"
enviar_email(email_b)

email_c = "pedro@gmail.com"
enviar_email(email_c)
# ¿Y si tengo 1000 emails?

# ✅ Con bucle (inteligente)
emails = ["juan@gmail.com", "maria@gmail.com", "pedro@gmail.com"]
for email in emails:
    enviar_email(email)  # Se repite automáticamente
```

### Caso 2: Generar secuencias
```python
# Tabla de multiplicar del 5
for i in range(1, 11):
    print(f"5 x {i} = {5 * i}")

# Output:
# 5 x 1 = 5
# 5 x 2 = 10
# ...
# 5 x 10 = 50
```

### Caso 3: Buscar/filtrar datos
```python
# Encontrar números pares
for numero in [1, 2, 3, 4, 5, 6]:
    if numero % 2 == 0:
        print(numero)  # Imprime: 2, 4, 6
```

---

## ¿CÓMO funcionan los Bucles?

Hay **dos tipos**:
1. **for** - cuando repites sobre una colección
2. **while** - cuando repites mientras se cumpla una condición

---

## Parte I: for - El Bucle Más Común

### Concepto: "Para cada... en..."

```python
for fruta in frutas:
    # Aquí adentro, "fruta" vale una diferente cada iteración
    print(fruta)
```

### Iterar sobre String

```python
palabra = "Python"

for letra in palabra:
    print(letra)

# Output: P, y, t, h, o, n (cada una en nueva línea)
```

### Iterar sobre Range

`range(inicio, fin, paso)`:

```python
# Del 0 al 4 (fin NO inclusivo)
for i in range(5):
    print(i)  # 0, 1, 2, 3, 4

# Del 5 al 9
for i in range(5, 10):
    print(i)  # 5, 6, 7, 8, 9

# Con paso de 2
for i in range(0, 10, 2):
    print(i)  # 0, 2, 4, 6, 8

# Hacia atrás
for i in range(10, 0, -1):
    print(i)  # 10, 9, 8, ..., 1
```

---

## Parte II: Bucle while

Repite **mientras una condición sea verdadera**:

### While Básico

```python
contador = 0

while contador < 5:
    print(contador)
    contador += 1

# Output: 0, 1, 2, 3, 4
```

### Cuidado: Bucles Infinitos

```python
# ❌ PELIGRO - se ejecuta para siempre
x = 0
while True:
    print(x)
    # Falta incrementar x

# ✅ CORRECTO - tiene condición de salida
x = 0
while x < 5:
    print(x)
    x += 1
```

### Entrada Validada con While

```python
edad = -1

while edad < 0 or edad > 120:
    edad = int(input("Ingresa edad válida (0-120): "))

print(f"Tu edad: {edad}")
```

---

## Parte III: break - Salir del Bucle

Sale inmediatamente del bucle:

```python
for i in range(10):
    if i == 5:
        break  # Sale del bucle
    print(i)

# Output: 0, 1, 2, 3, 4
```

### Caso Práctico: Buscar en Lista

```python
números = [1, 3, 5, 7, 9, 11]
buscado = 7

for num in números:
    if num == buscado:
        print(f"Encontrado: {buscado}")
        break
else:
    print("No encontrado")
```

---

## Parte IV: continue - Saltar Iteración

Salta al siguiente ciclo del bucle:

```python
for i in range(5):
    if i == 2:
        continue  # Salta esta iteración
    print(i)

# Output: 0, 1, 3, 4 (falta el 2)
```

### Caso Práctico: Ignorar Valores

```python
# Imprimir solo números pares
for i in range(1, 11):
    if i % 2 != 0:
        continue
    print(i)

# Output: 2, 4, 6, 8, 10
```

---

## Parte V: pass - No Hacer Nada

Placeholder cuando necesitas un bloque pero no tienes código aún:

```python
for i in range(5):
    pass  # Hacer nada, de momento

# Común en estructuras incompletas
if edad < 18:
    pass  # Implementaré después
else:
    print("Mayor de edad")
```

---

## Parte VI: enumerate() - Índice + Valor

Obtén el índice y el valor simultáneamente:

```python
frutas = ["manzana", "plátano", "cereza"]

for índice, fruta in enumerate(frutas):
    print(f"{índice}: {fruta}")

# Output:
# 0: manzana
# 1: plátano
# 2: cereza

# Comenzar en 1
for índice, fruta in enumerate(frutas, start=1):
    print(f"{índice}: {fruta}")

# Output:
# 1: manzana
# 2: plátano
# 3: cereza
```

---

## Parte VII: zip() - Combinar Listas

Empareja elementos de múltiples listas:

```python
nombres = ["Juan", "María", "Pedro"]
edades = [25, 30, 28]

for nombre, edad in zip(nombres, edades):
    print(f"{nombre} tiene {edad} años")

# Output:
# Juan tiene 25 años
# María tiene 30 años
# Pedro tiene 28 años

# Con 3+ listas
ciudades = ["Madrid", "Barcelona", "Valencia"]

for nombre, edad, ciudad in zip(nombres, edades, ciudades):
    print(f"{nombre} ({edad}) vive en {ciudad}")
```

**Nota**: Si las listas tienen diferente largura, zip() se detiene en la más corta.

---

## Parte VIII: reversed() - Invertido

Itera en orden inverso:

```python
números = [1, 2, 3, 4, 5]

for num in reversed(números):
    print(num)

# Output: 5, 4, 3, 2, 1

# Con strings
palabra = "Python"
for letra in reversed(palabra):
    print(letra, end="")

# Output: nohtyP
```

---

## Parte IX: else en Bucles

Se ejecuta si el bucle completa **sin break**:

```python
# Buscar número
for i in range(1, 6):
    if i == 10:
        print("Encontrado")
        break
else:
    print("No encontrado")

# Output: No encontrado


# Si hay break:
for i in range(1, 6):
    if i == 3:
        print("Encontrado")
        break
else:
    print("No encontrado")

# Output: Encontrado
# El else NO se ejecuta cuando hay break
```

---

## Parte X: Bucles Anidados

Bucles dentro de bucles:

```python
# Tabla de multiplicación
for i in range(1, 4):
    for j in range(1, 4):
        print(f"{i} x {j} = {i*j}")

# Output:
# 1 x 1 = 1
# 1 x 2 = 2
# 1 x 3 = 3
# 2 x 1 = 2
# ... etc
```

### Salir de Bucle Anidado

```python
# break sale solo del bucle más interno
encontrado = False

for i in range(10):
    for j in range(10):
        if i * j == 35:
            print(f"Encontrado: {i} x {j} = 35")
            encontrado = True
            break
    if encontrado:
        break
```

---

## Parte XI: List Comprehension (Avanzado)

Forma compacta de crear listas con bucles:

```python
# Tradicional
cuadrados = []
for i in range(5):
    cuadrados.append(i ** 2)
print(cuadrados)  # [0, 1, 4, 9, 16]

# Comprehension (equivalente)
cuadrados = [i ** 2 for i in range(5)]
print(cuadrados)  # [0, 1, 4, 9, 16]

# Con condición
pares = [i for i in range(10) if i % 2 == 0]
print(pares)  # [0, 2, 4, 6, 8]
```

---

## Parte XII: Comparación for vs while

| Situación | Usar | Ejemplo |
|-----------|------|----------|
| Iterar sobre secuencia | for | for x in lista: |
| Repetir mientras condición | while | while x < 10: |
| Número conocido de iteraciones | for | for i in range(100): |
| Condición desconocida | while | while usuario input válido |
| Alternar entre valores | while | while True/False changes |

---

## Parte XIII: Performance

```python
# Lento (muy recomendado para entendimiento)
for i in range(1000000):
    print(i)

# Rápido (lista completa en memoria)
números = list(range(1000000))

# Memory-efficient (genera sobre la marcha)
for i in range(1000000):
    pass
```

---

## Ejemplos Prácticos Completos

### Ejemplo 1: Tabla de Multiplicar

```python
numero = int(input("Tabla del: "))

for i in range(1, 11):
    print(f"{numero} x {i} = {numero * i}")
```

### Ejemplo 2: Validar Entrada

```python
opción = ""
while opción not in ["1", "2", "3"]:
    opción = input("Elige 1, 2 o 3: ")
print(f"Elegiste: {opción}")
```

### Ejemplo 3: Fibonacci

```python
n = int(input("¿Cuántos Fibonacci? "))
a, b = 0, 1

for _ in range(n):
    print(a, end=" ")
    a, b = b, a + b

# Output: 0 1 1 2 3 5 8 13 21 34
```

---

## Resumen

| Concepto | Uso | Ejemplo |
|----------|-----|---------|
| for | Iterar secuencia | for x in lista: |
| while | Repetir condición | while x < 10: |
| break | Salir bucle | break |
| continue | Saltar iteración | continue |
| pass | No hacer nada | pass |
| range() | Secuencia números | range(5) |
| enumerate() | Índice + valor | enumerate(lista) |
| zip() | Emparejar listas | zip(a, b) |
| reversed() | Invertir | reversed(lista) |
| List comp. | Bucle compacto | [x*2 for x in l] |

---

**Siguiente**: [funciones.md](funciones.md)
