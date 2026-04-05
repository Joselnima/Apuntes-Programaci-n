# JSON y Serialización

> **Convertir datos a formato universal que internet entiende**

---

## ¿QUÉ es JSON?

**JSON** = formato para guardar datos que TODOS entienden:

```
Python:
{"nombre": "Juan", "edad": 25}

JavaScript:
{"nombre": "Juan", "edad": 25}

PHP:
{"nombre": "Juan", "edad": 25}

Todos leen IGUAL ✅
```

**Analogía:**
- 🗣️ Español = tu amigo entiende
- 🌍 JSON = todos en el mundo entienden

**JSON = Texto ESTRUCTURADO**

---

## ¿PARA QÉ sirve?

### 1. APIs web (cliente-servidor)
```python
# Python → envía JSON → Servidor
cliente_manda = {"usuario": "juan", "contraseña": "123"}

# Servidor → responde JSON → Python
servidor_responde = {"token": "abc123", "estado": "autenticado"}
```

### 2. Guardar configuraciones
```python
# Archivo config.json
{
  "base_datos": "localhost",
  "puerto": 5432,
  "debug": true
}
```

### 3. Cachear datos
```python
# En lugar de consultar BD cada vez:
# 1. Obtén datos
# 2. Guarda como JSON (rápido)
# 3. Lee JSON en vez de BD
```

### 4. Intercambiar datos entre lenguajes
```python
# Python envía JSON
# Node.js lo recibe
# PHP lo procesa
# Todos entienden
```

---

## ¿CÓMO funcionan?

### 2 Operaciones Principales

```
Python Objeto      ↔ JSON String
  (diccionario)      (texto)
                    
  dict            dumps() → JSON
  list                        string
  str             loads() → Python
  int                       valor
  bool
```

**Ejemplo Visual:**
```python
Tengo:     {"nombre": "Juan", "edad": 25}  (diccionario Python)

dumps()↓   "{\\"nombre\\": \\"Juan\\", \\"edad\\": 25}"  (texto JSON)

loads()↓   {"nombre": "Juan", "edad": 25}  (diccionario Python)
```

---

## Parte I: Conceptos de JSON

### Tipos JSON

```
Strings:    "texto"
Números:    123, 3.14
Booleanos:  true, false
Null:       null
Arrays:     [1, 2, 3]
Objetos:    {"clave": "valor"}
```

### Ejemplo JSON

```json
{
  "nombre": "Juan",
  "edad": 25,
  "emails": ["j1@ejemplo.com", "j2@otro.com"],
  "activo": true,
  "datos_adicionales": null
}
```

---

## Parte II: Convertir a JSON (dumps, dump)

### json.dumps() - String a JSON

Convierte Python objeto a string JSON:

```python
import json

persona = {
    "nombre": "Juan",
    "edad": 25,
    "emails": ["j1@ejemplo.com", "j2@otro.com"]
}

# Convertir a JSON string
json_string = json.dumps(persona)
print(json_string)
# {"nombre": "Juan", "edad": 25, "emails": ["j1@ejemplo.com", "j2@otro.com"]}

print(type(json_string))  # <class 'str'>
```

### json.dump() - Guardar Directo a Archivo

```python
import json

persona = {
    "nombre": "Juan",
    "edad": 25
}

# Escribir JSON a archivo
with open("persona.json", "w") as f:
    json.dump(persona, f)  # Escribe automáticamente

# Leer archivo
with open("persona.json", "r") as f:
    contenido = f.read()
    print(contenido)
```

### Formateo Bonito (Pretty Print)

```python
import json

datos = {"nombre": "Juan", "edad": 25, "ciudad": "Madrid"}

# Sin formato
print(json.dumps(datos))
# {"nombre": "Juan", "edad": 25, "ciudad": "Madrid"}

# Con formato
print(json.dumps(datos, indent=2))
# {
#   "nombre": "Juan",
#   "edad": 25,
#   "ciudad": "Madrid"
# }

# Guardar bonito a archivo
with open("datos.json", "w") as f:
    json.dump(datos, f, indent=2)
```

---

## Parte III: Parsear JSON (loads, load)

### json.loads() - String a Python

```python
import json

json_string = '{"nombre": "Juan", "edad": 25}'

# Convertir a Python objeto (dict)
persona = json.loads(json_string)
print(persona)           # {'nombre': 'Juan', 'edad': 25}
print(type(persona))     # <class 'dict'>
print(persona["nombre"]) # Juan
```

### json.load() - Leer de Archivo

```python
import json

# Leer JSON de archivo
with open("persona.json", "r") as f:
    persona = json.load(f)  # Lee y parsea automáticamente
    print(persona["nombre"])
```

---

## Parte IV: Convertir Tipos Python-JSON

| Python | JSON | Inverso |
|--------|------|---------|
| dict | object | objeto |
| list | array | array |
| str | string | string |
| int, float | number | número |
| bool (True) | true | bool |
| bool (False) | false | bool |
| None | null | None |

```python
python_obj = {
    "texto": "valor",
    "número": 42,
    "flotante": 3.14,
    "booleano": True,
    "nada": None,
    "lista": [1, 2, 3]
}

json_string = json.dumps(python_obj)
print(json_string)
# {"texto": "valor", "número": 42, "flotante": 3.14, ...}

recuperado = json.loads(json_string)
print(recuperado == python_obj)  # True
```

---

## Parte V: Objetos Personalizados a JSON

### El Problema

```python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad

juan = Persona("Juan", 25)

# ❌ Esto falla
json.dumps(juan)  # TypeError: Object of type Persona is not JSON serializable
```

### Solución 1: JSONEncoder Personalizado

```python
import json

class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad

class PersonaEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, Persona):
            return {"nombre": obj.nombre, "edad": obj.edad}
        return super().default(obj)

juan = Persona("Juan", 25)

# Usar encodificador personalizado
json_string = json.dumps(juan, cls=PersonaEncoder)
print(json_string)  # {"nombre": "Juan", "edad": 25}
```

### Solución 2: __dict__

```python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad

juan = Persona("Juan", 25)

# Convertir atributos a dict
json_string = json.dumps(juan.__dict__)
print(json_string)  # {"nombre": "Juan", "edad": 25}
```

### Solución 3: Método to_dict()

```python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad
    
    def to_dict(self):
        return {"nombre": self.nombre, "edad": self.edad}

juan = Persona("Juan", 25)

json_string = json.dumps(juan.to_dict())
print(json_string)  # {"nombre": "Juan", "edad": 25}
```

---

## Parte VI: Parsear JSON Personalizado

### Con default_object_hook

```python
import json

def persona_decoder(d):
    if "nombre" in d and "edad" in d:
        return Persona(d["nombre"], d["edad"])
    return d

class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad
    
    def __repr__(self):
        return f"Persona({self.nombre}, {self.edad})"

json_string = '{"nombre": "Juan", "edad": 25}'

# Parsear con función personalizada
persona = json.loads(json_string, object_hook=persona_decoder)
print(persona)  # Persona(Juan, 25)
```

---

## Parte VII: Validación de JSON

```python
import json

json_string = "JSON inválido"

try:
    datos = json.loads(json_string)
except json.JSONDecodeError as e:
    print(f"JSON inválido: {e}")
    print(f"Línea: {e.lineno}, Columna: {e.colno}")
```

---

## Parte VIII: API REST Simple Ejemplo

Trabajar con APIs web que retornan JSON:

```python
import json
import requests  # pip install requests

# Obtener datos de API
respuesta = requests.get("https://api.example.com/usuarios/1")

# Parsear JSON
datos = respuesta.json()  # Equivale a json.loads(respuesta.text)

print(datos["nombre"])
print(datos["email"])
```

---

## Parte IX: Guardar y Cargar Configuración

```python
import json
from pathlib import Path

class Configuración:
    def __init__(self, archivo="config.json"):
        self.archivo = Path(archivo)
        self.datos = {}
        self.cargar()
    
    def cargar(self):
        """Cargar configuración desde archivo"""
        if self.archivo.exists():
            with open(self.archivo, "r") as f:
                self.datos = json.load(f)
        else:
            self.datos = {"debug": False, "puerto": 8000}
    
    def guardar(self):
        """Guardar configuración a archivo"""
        with open(self.archivo, "w") as f:
            json.dump(self.datos, f, indent=2)
    
    def obtener(self, clave, defecto=None):
        return self.datos.get(clave, defecto)
    
    def establecer(self, clave, valor):
        self.datos[clave] = valor
        self.guardar()

# Uso
config = Configuración()
print(config.obtener("debug"))  # False
config.establecer("debug", True)
print(config.obtener("debug"))  # True (guardado automáticamente)
```

---

## Parte X: Pickle (Serialización Python)

Para serializar objetos Python complejos (no es universal como JSON):

```python
import pickle

class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad

juan = Persona("Juan", 25)

# Guardar
with open("juan.pkl", "wb") as f:
    pickle.dump(juan, f)

# Cargar
with open("juan.pkl", "rb") as f:
    juan_recuperado = pickle.load(f)
    print(juan_recuperado.nombre)
```

**⚠️ Pickle SOLO funciona en Python. JSON es universal. Usa JSON cuando sea posible.**

---

## Resumen

| Operación | Función | Entrada | Salida |
|-----------|---------|---------|--------|
| A string JSON | json.dumps() | Object/dict | str |
| A archivo | json.dump() | Object/dict | File |
| De string JSON | json.loads() | str | dict/list |
| De archivo | json.load() | File | dict/list |
| Validar JSON | JSONDecodeError | str | Error |
| Personalizado | JSONEncoder | Object | str |

---

**Siguiente**: [modulos_paquetes.md](modulos_paquetes.md)
