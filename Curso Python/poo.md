# Programación Orientada a Objetos (POO) en Python

> **Cómo organizar código en estructuras que modelan cosas del mundo real - La forma profesional de programar**

---

## ¿QÉ es POO (Programación Orientada a Objetos)?

### Una Analogía del Mundo Real

Imagina que quieres modelar **autos** en tu programa.

**Sin POO** (Caótico):
```python
auto1_marca = "Toyota"
auto1_color = "azul"
auto1_año = 2020
auto1_velocidad = 0

auto2_marca = "Honda"
auto2_color = "rojo"
auto2_año = 2019
auto2_velocidad = 0

def acelerar(auto, velocidad):
    # Qué auto? ¿cuáles variables pertenecen a cuál?
    # Esto se vuelve un desastre...
    pass
```

**Con POO** (Organizado):
```python
class Auto:
    def __init__(self, marca, color, año):
        self.marca = marca
        self.color = color
        self.año = año
        self.velocidad = 0
    
    def acelerar(self, cantidad):
        self.velocidad += cantidad
        print(f"Acelerando... Velocidad: {self.velocidad}")

# Ahora es claro: cada auto es su propia cosa
toyota = Auto("Toyota", "azul", 2020)
honda = Auto("Honda", "rojo", 2019)

toyota.acelerar(50)  # El TOYOTA acelera
honda.acelerar(30)   # El HONDA acelera
```

---

## ¿PARA QUÉ sirve POO?

### 1. **Organiza datos relacionados**
```python
# Sin POO - confuso
persona_nombre = "Juan"
persona_edad = 25
persona_ciudad = "Madrid"
persona_otro = "?"
persona_otro_mas = "?"
# ¿Cuáles variables van juntas?

# Con POO - claro
class Persona:
    def __init__(self, nombre, edad, ciudad):
        self.nombre = nombre
        self.edad = edad
        self.ciudad = ciudad

juan = Persona("Juan", 25, "Madrid")  # Claro: todo aquí
```

### 2. **Agrupa funciones relacionadas**
```python
# Sin POO - desorganizado
def calcular_edad(persona_edad):
    pass

def validar_edad(persona_edad):
    pass

def es_mayor_de_edad(persona_edad):
    pass

# Con POO - todo junto
class Persona:
    def calcular_edad(self):
        pass
    
    def validar_edad(self):
        pass
    
    def es_mayor_de_edad(self):
        pass
```

### 3. **Reutiliza código fácilmente**
```python
# Creas la clase una vez
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad

# La usas múltiples veces
juan = Persona("Juan", 25)
maria = Persona("María", 30)
pedro = Persona("Pedro", 28)
```

---

## ¿CÓMO funciona POO?

### Los 4 Conceptos Clave

1. **Clase** = El plano/molde
2. **Objeto** = Lo que creas del molde
3. **Atributo** = Las características
4. **Método** = Las acciones

### Analogía: Una Fábrica de Galletas

```
📋 CLASE (El plano)
   - Cómo hacer una galleta
   - Qué ingredientes
   - Qué pasos

🍪 OBJETO (La galleta hecha)
   - Una galleta real de chocolate
   - Otra galleta real de vainilla

📍 ATRIBUTO (Las características)
   - Sabor: chocolate
   - Tamaño: grande
   - Fresquedad: sí

⚙️ MÉTODO (Las acciones)
   - Morder()
   - Compartir()
   - Romper()
```

---

## Parte I: Tu Primera Clase - Lo Más Básico

### Definición

```python
class Persona:
    \"\"\"Representa una persona\"\"\"
    pass

# Crear instancia (objeto)
p = Persona()
print(type(p))  # <class '__main__.Persona'>
```

---

## Parte II: __init__ (Constructor)

Se ejecuta automáticamente al crear una instancia:

```python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre  # Atributo de instancia
        self.edad = edad

# Crear instancia
p = Persona("Juan", 25)
print(p.nombre)  # Juan
print(p.edad)    # 25
```

### self - La Palabra Clave

`self` representa la instancia misma:

```python
class Auto:
    def __init__(self, marca, modelo):
        self.marca = marca
        self.modelo = modelo
        # self accede a la instancia

auto = Auto("Toyota", "Corolla")
print(auto.marca)  # Toyota
```

---

## Parte III: Atributos de Instancia

Propiedades de cada objeto:

```python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre   # Atributo
        self.edad = edad       # Atributo
        self.amigos = []       # Atributo (lista)

juan = Persona("Juan", 25)
print(juan.nombre)  # Juan

# Modificar atributo
juan.edad = 26
print(juan.edad)    # 26

# Agregar atributo (Python permite)
juan.teléfono = "555-1234"
print(juan.teléfono)  # 555-1234
```

---

## Parte IV: Métodos

Funciones dentro de una clase:

```python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad
    
    def saludar(self):
        return f"Hola, soy {self.nombre}"
    
    def es_mayor(self):
        return self.edad >= 18
    
    def cumpleaños(self):
        self.edad += 1

# Usar métodos
juan = Persona("Juan", 25)
print(juan.saludar())     # Hola, soy Juan
print(juan.es_mayor())    # True
juan.cumpleaños()
print(juan.edad)          # 26
```

---

## Parte V: __str__ y __repr__ (Representación)

Controla cómo se ve el objeto como string:

```python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad
    
    def __str__(self):
        \"\"\"Para print() y str()\"\"\"
        return f"{self.nombre} ({self.edad})"
    
    def __repr__(self):
        \"\"\"Para representación técnica\"\"\"
        return f"Persona('{self.nombre}', {self.edad})"

juan = Persona("Juan", 25)
print(juan)           # Juan (25)
print(str(juan))      # Juan (25)
print(repr(juan))     # Persona('Juan', 25)
```

---

## Parte VI: Atributos de Clase vs Instancia

### De Instancia (pertenecen a cada objeto)

```python
class Auto:
    def __init__(self, marca, modelo):
        self.marca = marca    # Cada instancia tiene su marca
        self.modelo = modelo

auto1 = Auto("Toyota", "Corolla")
auto2 = Auto("Honda", "Civic")

print(auto1.marca)  # Toyota
print(auto2.marca)  # Honda (diferente)
```

### De Clase (compartidos por todos)

```python
class Auto:
    ruedas = 4  # Atributo de clase
    
    def __init__(self, marca):
        self.marca = marca

print(Auto.ruedas)  # 4 (acceso desde clase)

auto = Auto("Toyota")
print(auto.ruedas)  # 4 (acceso desde instancia)

Auto.ruedas = 6     # Cambiar para toda la clase
print(auto.ruedas)  # 6
```

---

## Parte VII: Métodos Especiales (@classmethod, @staticmethod)

### @classmethod

Recibe la clase, no la instancia:

```python
class Persona:
    contador = 0
    
    def __init__(self, nombre):
        self.nombre = nombre
        Persona.contador += 1
    
    @classmethod
    def total_personas(cls):
        return cls.contador

p1 = Persona("Juan")
p2 = Persona("María")

print(Persona.total_personas())  # 2
print(p1.total_personas())       # 2
```

### @staticmethod

Sin acceso a instancia ni clase:

```python
class Utilidad:
    @staticmethod
    def sumar(a, b):
        return a + b

print(Utilidad.sumar(5, 3))  # 8

u = Utilidad()
print(u.sumar(10, 5))        # 15
```

---

## Parte VIII: Encapsulación (Privacidad)

Controlar acceso a atributos:

### Convención (single underscore)

```python
class CuentaBancaria:
    def __init__(self, saldo):
        self._saldo = saldo  # Convención: privado
    
    def depositar(self, cantidad):
        self._saldo += cantidad
    
    def retirar(self, cantidad):
        if cantidad <= self._saldo:
            self._saldo -= cantidad
        else:
            print("Saldo insuficiente")
    
    def obtener_saldo(self):
        return self._saldo

cuenta = CuentaBancaria(1000)
cuenta.depositar(500)
print(cuenta.obtener_saldo())  # 1500

# Técnicamente accesible (convención)
print(cuenta._saldo)  # 1500
```

### Privado (double underscore)

```python
class CuentaBancaria:
    def __init__(self, saldo):
        self.__saldo = saldo  # Realmente privado
    
    def obtener_saldo(self):
        return self.__saldo

cuenta = CuentaBancaria(1000)
print(cuenta.obtener_saldo())  # 1000

# ❌ Esto da error
# print(cuenta.__saldo)  # AttributeError

# Pero Python permite evadir (name mangling)
print(cuenta._CuentaBancaria__saldo)  # 1000 (evitable)
```

---

## Parte IX: Properties (@property)

Atributos con getters y setters:

```python
class Persona:
    def __init__(self, nombre, año_nacimiento):
        self.nombre = nombre
        self._año_nacimiento = año_nacimiento
    
    @property
    def edad(self):
        \"\"\"Calcular edad automáticamente\"\"\"
        from datetime import datetime
        return datetime.now().year - self._año_nacimiento
    
    @edad.setter
    def edad(self, valor):
        \"\"\"Establecer año de nacimiento desde edad\"\"\"
        self._año_nacimiento = datetime.now().year - valor

juan = Persona("Juan", 1998)
print(juan.edad)  # 26 (o la edad actual)

juan.edad = 30
print(juan._año_nacimiento)  # 1993
```

---

## Parte X: Herencia

Reutilizar código de una clase base:

```python
class Animal:
    def __init__(self, nombre):
        self.nombre = nombre
    
    def hacer_sonido(self):
        return "Sonido genérico"

class Perro(Animal):
    def hacer_sonido(self):
        return "Guau guau"

class Gato(Animal):
    def hacer_sonido(self):
        return "Miau"

perro = Perro("Rex")
print(perro.nombre)           # Rex
print(perro.hacer_sonido())   # Guau guau

gato = Gato("Whiskers")
print(gato.hacer_sonido())    # Miau
```

### super() - Llamar a la Clase Base

```python
class Animal:
    def __init__(self, nombre):
        self.nombre = nombre

class Perro(Animal):
    def __init__(self, nombre, raza):
        super().__init__(nombre)  # Llamar __init__ del padre
        self.raza = raza

perro = Perro("Rex", "Labrador")
print(perro.nombre)  # Rex
print(perro.raza)    # Labrador
```

---

## Parte XI: Polimorfismo

Diferentes clases, mismo método:

```python
class Figura:
    def área(self):
        pass

class Círculo(Figura):
    def __init__(self, radio):
        self.radio = radio
    
    def área(self):
        return 3.14 * self.radio ** 2

class Cuadrado(Figura):
    def __init__(self, lado):
        self.lado = lado
    
    def área(self):
        return self.lado ** 2

# Código polimórfico
figuras = [Círculo(5), Cuadrado(4)]

for figura in figuras:
    print(figura.área())  # Cada uno calcula su área
```

---

## Parte XII: Métodos Mágicos (__add__, __len__, etc.)

Operador atormentan:

```python
class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def __add__(self, otro):
        \"\"\"Suma de vectores\"\"\"
        return Vector(self.x + otro.x, self.y + otro.y)
    
    def __len__(self):
        \"\"\"Magnitud del vector\"\"\"
        return int((self.x ** 2 + self.y ** 2) ** 0.5)
    
    def __repr__(self):
        return f"Vector({self.x}, {self.y})"

v1 = Vector(3, 4)
v2 = Vector(1, 2)

v3 = v1 + v2  # __add__
print(v3)     # Vector(4, 6)

print(len(v1))  # __len__ → 5
```

---

## Parte XIII: Interfaces y ABC (Abstract Base Classes)

Forzar que subclases implementen métodos:

```python
from abc import ABC, abstractmethod

class Animal(ABC):
    @abstractmethod
    def hacer_sonido(self):
        pass

class Perro(Animal):
    def hacer_sonido(self):
        return "Guau guau"

# ❌ Esto falla
# animal = Animal()  # TypeError

# ✅ Esto funciona
perro = Perro()
print(perro.hacer_sonido())  # Guau guau
```

---

## Parte XIV: Ejemplo Completo: Sistema de Biblioteca

```python
class Libro:
    def __init__(self, título, autor, isbn):
        self.título = título
        self.autor = autor
        self.isbn = isbn
        self.disponible = True
    
    def __str__(self):
        return f"{self.título} por {self.autor}"

class Usuario:
    def __init__(self, nombre, id):
        self.nombre = nombre
        self.id = id
        self.libros_prestados = []
    
    def prestar_libro(self, libro):
        if libro.disponible:
            self.libros_prestados.append(libro)
            libro.disponible = False
            return True
        return False

class Biblioteca:
    def __init__(self):
        self.libros = []
        self.usuarios = []
    
    def agregar_libro(self, libro):
        self.libros.append(libro)
    
    def registrar_usuario(self, usuario):
        self.usuarios.append(usuario)
    
    def prestar(self, usuario_id, isbn):
        usuario = next((u for u in self.usuarios if u.id == usuario_id), None)
        libro = next((l for l in self.libros if l.isbn == isbn), None)
        
        if usuario and libro:
            return usuario.prestar_libro(libro)
        return False

# Uso
bib = Biblioteca()

libro1 = Libro("Python 101", "Guido", "123")
libro2 = Libro("Web Dev", "Bob", "456")
bib.agregar_libro(libro1)
bib.agregar_libro(libro2)

juan = Usuario("Juan", 1)
bib.registrar_usuario(juan)

bib.prestar(1, "123")
print(juan.libros_prestados)  # [Python 101 por Guido]
```

---

## Resumen

| Concepto | Uso | Ejemplo |
|----------|-----|---------|
| class | Definir clase | class Persona: |
| __init__ | Constructor | def __init__(self, x): |
| self | Instancia | self.x = x |
| Atributo | Propiedad | self.nombre |
| Método | Función | def saludar(self): |
| @property | Atributo computed | @property def edad(): |
| Herencia | Reutilizar | class Perro(Animal): |
| super() | Clase base | super().__init__() |
| @classmethod | Método de clase | @classmethod def crear(): |
| @staticmethod | Método estático | @staticmethod def sumar(): |

---

**Siguiente**: [manejo_errores.md](manejo_errores.md)
