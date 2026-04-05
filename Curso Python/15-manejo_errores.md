# Manejo de Excepciones en Python

> **Cómo capturar y manejar errores - Tu código no siempre funcionará perfecto**

---

## ¿QÉ es un Error/Excepción?

Cuando algo sale **mal** en tu código, Python lanza un **error**:

```python
# Error: dividir por cero
10 / 0
# ZeroDivisionError: division by zero

# Error: índice fuera de rango
lista = [1, 2, 3]
print(lista[100])
# IndexError: list index out of range

# Error: convertir texto a número
edad = int("veinticinco")
# ValueError: invalid literal for int()

# El programa se DETIENE aquí → Crash 💥
```

---

## ¿PARA QÉ sirve manejar errores?

### Caso 1: Usuario escribe algo inválido
```python
# ❌ Sin manejo (se cae)
edad_texto = input("¿Cuántos años? ")
edad = int(edad_texto)  # Si escribe "abc", ¡ERROR!
print(f"Tienes {edad} años")

# ✅ Con manejo (controlas)
try:
    edad = int(input("¿Cuántos años? "))
    print(f"Tienes {edad} años")
except ValueError:
    print("ERROR: Debes escribir un número")
    # Programa sigue funcionando
```

### Caso 2: Archivo que no existe
```python
# ❌ Sin manejo
archivo = open("datos.txt")  # Si no existe...
# FileNotFoundError

# ✅ Con manejo
try:
    archivo = open("datos.txt")
    contenido = archivo.read()
except FileNotFoundError:
    print("Archivo no encontrado, usando datos vacíos")
    contenido = ""
```

### Caso 3: Operación matemática
```python
# ❌ Sin manejo
resultado = 100 / numero_usuario  # Si es 0...
# ZeroDivisionError

# ✅ Con manejo
try:
    resultado = 100 / numero_usuario
except ZeroDivisionError:
    resultado = 0
    print("No se puede dividir entre cero")
```

---

## ¿CÓMO funcionan?

### La Estructura: try/except

```python
try:
    # Código que PODRÍA generar error
    resultado = 10 / 0
except ZeroDivisionError:
    # Qué HACER si ocurre ese error
    print("No puedo dividir entre cero")
```

**Flujo:**
1. Python ejecuta el código en `try`
2. Si NO hay error → sigue normalmente
3. Si HAY error → salta al `except` correspondiente
4. Ejecuta el `except` → programa sigue

---

## Parte I: Try/Except (Captura)  Básico

```python
try:
    # Código que PODRÍA fallar
    x = 10 / 0
except ZeroDivisionError:
    # Qué hacer si falla
    print("No puedes dividir entre cero")
```

### Capturar La Excepción

```python
try:
    resultado = int("abc")
except ValueError as e:
    print(f"Error: {e}")
    print(type(e))
```

---

## Parte II: Múltiples Excepciones

### Opción 1: Excepciones Separadas

```python
try:
    edad = int(input("Edad: "))
    año = 2026 - edad
    print(f"Naciste en {año}")
except ValueError:
    print("Debes ingresar un número")
except ZeroDivisionError:
    print("No puedes dividir entre cero")
```

### Opción 2: Una Excepción General

```python
try:
    edad = int(input("Edad: "))
except (ValueError, TypeError):
    print("Entrada inválida")
```

### Opción 3: Cualquier Excepción

```python
try:
    operación_riesgosa()
except Exception as e:  # Capta CUALQUIER excepción
    print(f"Error inesperado: {e}")
```

**⚠️ EVITA capturar `Exception` genéricamente - especifica qué errores esperas.**

---

## Parte III: else en Try/Except

Se ejecuta si **NO** hay excepción:

```python
try:
    edad = int(input("Edad: "))
except ValueError:
    print("Debes ingresar un número")
else:
    # Se ejecuta si todo está bien
    if edad >= 18:
        print("Mayor de edad")
    else:
        print("Menor de edad")
```

---

## Parte IV: finally - Siempre se Ejecuta

Se ejecuta independientemente de si hay error o no:

```python
try:
    archivo = open("datos.txt", "r")
    contenido = archivo.read()
except FileNotFoundError:
    print("Archivo no encontrado")
finally:
    # SIEMPRE se ejecuta
    archivo.close()  # Limpiar recursos
    print("Operación completada")
```

### Ejemplo: Conexión a Base de Datos

```python
try:
    conectar_bd()
    hacer_consulta()
except ConnectionError:
    print("Error de conexión")
finally:
    desconectar_bd()  # Siempre desconectar
```

---

## Parte V: Lanzar Excepciones (raise)

Crear errores personalizados:

```python
def validar_edad(edad):
    if edad < 0:
        raise ValueError("La edad no puede ser negativa")
    if edad > 150:
        raise ValueError("Edad poco realista")
    return True

try:
    validar_edad(-5)
except ValueError as e:
    print(f"Error: {e}")  # Error: La edad no puede ser negativa
```

---

## Parte VI: Excepciones Personalizadas

Clases propias de excepciones:

```python
class EdadInválida(Exception):
    pass

class SaldoInsuficiente(Exception):
    pass

def retirar(saldo, cantidad):
    if cantidad > saldo:
        raise SaldoInsuficiente(f"Solo tienes {saldo}")
    return saldo - cantidad

try:
    nuevo_saldo = retirar(100, 200)
except SaldoInsuficiente as e:
    print(f"No puedes: {e}")
```

### Con Mensajes Personalizados

```python
class CuentaBloqueada(Exception):
    def __init__(self, mensaje, tiempo_espera):
        super().__init__(mensaje)
        self.tiempo_espera = tiempo_espera

try:
    raise CuentaBloqueada("Demasiados intentos", 3600)
except CuentaBloqueada as e:
    print(f"{e} - Espera {e.tiempo_espera}s")
```

---

## Parte VII: Tipos Comunes de Excepciones

| Excepción | Cuándo Ocurre | Ejemplo |
|-----------|--------------|---------|
| ValueError | Valor inválido | int("abc") |
| TypeError | Tipo incorrecto | "5" + 5 |
| IndexError | Índice fuera de rango | lista[10] cuando tiene 5 |
| KeyError | Clave no existe | dict["clave_inexistente"] |
| FileNotFoundError | Archivo no existe | open("no_existe.txt") |
| ZeroDivisionError | División entre 0 | 10 / 0 |
| AttributeError | Atributo no existe | objeto.atributo_inexistente |
| ImportError | No puede importar | import módulo_inexistente |
| RuntimeError | Error en tiempo ejecución | Error general |
|NameError | Variable/nombre no existe | usar variable no definida |

---

## Parte VIII: Debugging con Traceback

Cuando ocurre error, Python muestra el traceback (camino del error):

```
Traceback (most recent call last):
  File "programa.py", line 10, in <module>
    resultado = 10 / 0
ZeroDivisionError: division by zero
```

Lee de **abajo hacia arriba**:
1. **Qué error**: ZeroDivisionError
2. **Dónde**: Línea 10 de programa.py
3. **Qué código**: `resultado = 10 / 0`

---

## Parte IX: Contexto de Excepciones (Context)

Acceder a información del error:

```python
try:
    numero = int("abc")
except ValueError as error:
    print(f"Tipo: {type(error).__name__}")  # ValueError
    print(f"Mensaje: {error}")               # invalid literal...
    print(f"Args: {error.args}")             # ('invalid literal...',)
```

---

## Parte X: with para Manejo Automático

Cierra recursos automáticamente:

```python
# ❌ Manual
archivo = open("datos.txt")
contenido = archivo.read()
archivo.close()

# ✅ Automático con with
with open("datos.txt") as archivo:
    contenido = archivo.read()
# archivo.close() se ejecuta automáticamente

# Múltiples recursos
with open("entrada.txt") as entrada, open("salida.txt", "w") as salida:
    for línea in entrada:
        salida.write(línea.upper())
```

---

## Parte XI: Ejemplo Integral

```python
class CuentaBancaria:
    def __init__(self, saldo=0):
        self._saldo = saldo
    
    def depositar(self, cantidad):
        if cantidad <= 0:
            raise ValueError("Depósito debe ser positivo")
        self._saldo += cantidad
        return self._saldo
    
    def retirar(self, cantidad):
        if cantidad <= 0:
            raise ValueError("Retiro debe ser positivo")
        if cantidad > self._saldo:
            raise ValueError(f"Saldo insuficiente. Tienes {self._saldo}")
        self._saldo -= cantidad
        return self._saldo
    
    def obtener_saldo(self):
        return self._saldo

def operar_cuenta():
    cuenta = CuentaBancaria(1000)
    
    operaciones = [
        ("depositar", 500),
        ("retirar", 300),
        ("retirar", 2000),  # ❌ Falla
        ("depositar", -50)   # ❌ Nunca se ejecuta
    ]
    
    for operación, cantidad in operaciones:
        try:
            if operación == "depositar":
                nuevo_saldo = cuenta.depositar(cantidad)
                print(f"Depositado {cantidad}. Saldo: {nuevo_saldo}")
            elif operación == "retirar":
                nuevo_saldo = cuenta.retirar(cantidad)
                print(f"Retirado {cantidad}. Saldo: {nuevo_saldo}")
        except ValueError as e:
            print(f"❌ Error: {e}")
        except Exception as e:
            print(f"❌ Error inesperado: {e}")
        finally:
            print(f"   Saldo actual: {cuenta.obtener_saldo()}\n")

operar_cuenta()
```

---

## Parte XII: Best Practices

### ✅ BIEN
```python
try:
    resultado = int(entrada_usuario)
except ValueError:
    print("Ingresa un número válido")
```

### ❌ MAL
```python
try:
    resultado = int(entrada_usuario)
except:  # Demasiado genérico
    pass  # El silencio oculta errores
```

### Reglas de Oro

1. **Sé específico**: Captura excepciones específicas, no genéricas
2. **Log de errores**: Registra qué falló
3. **Limpiar recursos**: Usa `finally` o `with`
4. **No silencies errores**: `except: pass` es malo
5. **Proporciona contexto**: Lanza excepciones con mensajes útiles

---

## Resumen

| Concepto | Uso | Ejemplo |
|----------|-----|---------|
| try | Código que puede fallar | try: operación() |
| except | Manejar error | except ValueError: |
| else | Si NO falla | else: print("OK") |
| finally | Siempre ejecutar | finally: limpiar() |
| raise | Lanzar error | raise ValueError("...") |
| with | Auto-cerrar recursos | with open(...) as f: |
| Exception | Exención genérica | except Exception: |
| as | Alias | except Error as e: |

---

**Siguiente**: [archivos_io.md](archivos_io.md)
