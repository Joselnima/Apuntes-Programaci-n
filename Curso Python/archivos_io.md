# Entrada y Salida de Archivos (I/O)

> **Leer y escribir archivos - Cómo guardar datos permanentemente**

---

## ¿QÉ es I/O de Archivos?

**I/O** = Input/Output (lectura y escritura)

Es guardar **información permanente** en tu computadora:

```
Programa Python
    ↓ (escribe)
  archivo.txt  ← Datos guardados en disco
    ↓ (lee después)
Programa Python (próxima vez)
```

**Analogía:**
- 📝 Tu programa = notas en un papel
- 💾 Archivo = guardar el papel en una carpeta
- 📂 Récuperación = abrir esa carpeta después

---

## ¿PARA QÉ sirven?

### 1. Guardar datos del usuario
```python
usuario = "Juan"
crear_archivo("usuarios.txt")
guardar(usuario)  # Ahora está PERMANENTE
```

### 2. Guardar configuraciones
```python
configuración = {
    "tema": "oscuro",
    "idioma": "español"
}
guardar_config(configuración)
```

### 3. Procesar archivos grandes
```python
# Leer millón de líneas SIN cargar todo en RAM
for línea in leer_línea_por_línea("datos.csv"):
    procesar(línea)
```

### 4. Logs de eventos
```python
registrar("Error en línea 42")
registrar("Usuario inició sesión")
```

---

## ¿CÓMO funcionan?

### 3 Pasos: Abrir, Trabajar, Cerrar

```python
# Paso 1: ABRIR
archivo = open("letra.txt", "r")  # r = read (leer)

# Paso 2: USAR
contenido = archivo.read()

# Paso 3: CERRAR
archivo.close()

# PERO: Si ocurre error, NUNCA se cierra...
```
`
**Mejor forma - Cierra automáticamente:**
```python
with open("datos.txt", "r") as archivo:
    contenido = archivo.read()
# Se cierra automáticamente, incluso con error
```

---

## Parte I: Modos de Apertura - Diferentes Formas de Abrir

| Modo | Descripción | Si existe | Si no existe |
|------|-------------|-----------|--------------|
| "r" | Lectura | Abre | Error |
| "w" | Escritura | Crea nuevo | Crea |
| "a" | Agregar (append) | Añade al final | Crea |
| "x" | Crear exclusivo | Error | Crea |
| "r+" | Leer + Escribir | Abre | Error |
| "w+" | Escribir + Leer | Crea nuevo | Crea |
| "a+" | Agregar + Leer | Añade al final | Crea |

### Binario

Agrega "b" para binario: "rb", "wb", "xb"

---

## Parte II: Lectura Básica

### read() - Todo el Contenido

```python
# Leer todo completo
with open("archivo.txt", "r") as f:
    contenido = f.read()
    print(contenido)
    print(type(contenido))  # <class 'str'>
```

### readline() - Una Línea

```python
with open("archivo.txt", "r") as f:
    línea1 = f.readline()  # Primera línea
    línea2 = f.readline()  # Segunda línea
    print(línea1)
    print(línea2)
```

### readlines() - Todas como Lista

```python
with open("archivo.txt", "r") as f:
    líneas = f.readlines()  # ['Línea 1\n', 'Línea 2\n', ...]
    for línea in líneas:
        print(línea.strip())  # .strip() quita \n
```

### Iterar Directamente

```python
# ✅ MEJOR - memory efficient
with open("archivo.txt", "r") as f:
    for línea in f:
        print(línea.strip())
```

---

## Parte III: Escritura Básica

### write() - Escribir String

```python
with open("salida.txt", "w") as f:
    f.write("Primera línea\n")
    f.write("Segunda línea\n")
    f.write("Tercera línea")
```

### writelines() - Escribir Lista

```python
líneas = ["Línea 1\n", "Línea 2\n", "Línea 3\n"]

with open("salida.txt", "w") as f:
    f.writelines(líneas)
```

### Append (Agregar)

```python
with open("archivo.txt", "a") as f:
    f.write("Línea nueva al final\n")
```

---

## Parte IV: print() a Archivo

Redirigir print directamente:

```python
with open("log.txt", "w") as f:
    print("Mensaje 1", file=f)
    print("Mensaje 2", file=f)
    print("Mensaje 3")  # Sin file= va a terminal
```

---

## Parte V: Operaciones de Archivos

### Verificar si Existe

```python
import os

if os.path.exists("archivo.txt"):
    print("Existe")
else:
    print("No existe")

# Con pathlib (moderno)
from pathlib import Path

if Path("archivo.txt").exists():
    print("Existe")
```

### Eliminar Archivo

```python
import os

os.remove("archivo.txt")

# Con pathlib
from pathlib import Path
Path("archivo.txt").unlink()
```

### Renombrar

```python
import os

os.rename("viejo_nombre.txt", "nuevo_nombre.txt")

# Con pathlib
from pathlib import Path
Path("viejo.txt").rename("nuevo.txt")
```

### Directorio Actual

```python
import os

print(os.getcwd())  # Directorio actual
os.chdir("/Nueva/Ruta")  # Cambiar directorio

# Con pathlib
from pathlib import Path
print(Path.cwd())
```

---

## Parte VI: Rutas (Paths)

### Strings (Viejo)

```python
# Problema: \ es escape en strings
ruta = "C:\Usuarios\Juan"  # ❌ Problemas

# Soluciones
ruta = "C:\\Usuarios\\Juan"  # Escapar
ruta = r"C:\Usuarios\Juan"    # Raw string
ruta = "C:/Usuarios/Juan"     # Barras
```

### pathlib (Moderno)

```python
from pathlib import Path

# Multiplataforma (\ en Windows, / en Linux)
ruta = Path("Usuarios") / "Juan" / "Documentos"
print(ruta)  # Usuarios/Juan/Documentos (o con \ en Windows)

# Componentes
print(ruta.parent)   # Usuarios/Juan
print(ruta.name)     # Documentos
print(ruta.suffix)   # .txt (extensión)

# Relativo a absoluto
ruta_abs = ruta.absolute()

# Crear carpeta
ruta.mkdir(parents=True, exist_ok=True)

# Listar archivos
for archivo in ruta.glob("*.txt"):
    print(archivo)
```

---

## Parte VII: Leer CSV

### Con Strings Simples

```python
def leer_csv(archivo):
    usuarios = []
    with open(archivo, "r") as f:
        for línea in f:
            datos = línea.strip().split(",")
            usuarios.append(datos)
    return usuarios

# archivo.txt: Juan,25,Madrid
#              María,30,Barcelona
usuarios = leer_csv("usuarios.txt")
```

### Con csv Module

```python
import csv

with open("usuarios.csv", "r") as f:
    lector = csv.DictReader(f)  # Con encabezados
    for fila in lector:
        print(fila["nombre"], fila["edad"])

# O sin encabezados
with open("usuarios.csv", "r") as f:
    lector = csv.reader(f)
    for fila in lector:
        print(fila)  # ['Juan', '25', 'Madrid']
```

### Escribir CSV

```python
import csv

datos = [
    ["nombre", "edad", "ciudad"],
    ["Juan", 25, "Madrid"],
    ["María", 30, "Barcelona"]
]

with open("usuarios.csv", "w", newline="") as f:
    escritor = csv.writer(f)
    for fila in datos:
        escritor.writerow(fila)
```

---

## Parte VIII: Codificación (Encoding)

Especificar cómo se interpreta el texto:

```python
# UTF-8 (default, recomendado)
with open("archivo.txt", "r", encoding="utf-8") as f:
    contenido = f.read()

# Latin-1 (para archivos antiguos)
with open("archivo.txt", "r", encoding="latin-1") as f:
    contenido = f.read()

# ASCII
with open("archivo.txt", "r", encoding="ascii") as f:
    contenido = f.read()
```

**Regla**: Siempre especifica `encoding="utf-8"` para compatibilidad.

---

## Parte IX: Tamaño de Lectura Control

Para archivos gigantes, lee en bloques:

```python
# Leer en chunks (bloques)
with open("archivo_gigante.txt", "r") as f:
    while True:
        chunk = f.read(4096)  # 4KB por vez
        if not chunk:
            break
        procesar(chunk)

# O línea por línea (mejor)
with open("archivo_gigante.txt", "r") as f:
    for línea in f:
        procesar(línea)
```

---

## Parte X: Busca y Reemplaza en Archivo

```python
# Leer todo
with open("archivo.txt", "r") as f:
    contenido = f.read()

# Reemplazar
nuevo_contenido = contenido.replace("viejo", "nuevo")

# Escribir
with open("archivo.txt", "w") as f:
    f.write(nuevo_contenido)

# O directamente con re (regex)
import re

with open("archivo.txt", "r") as f:
    contenido = f.read()

nuevo = re.sub(r"\d+", "X", contenido)  # Reemplazar números

with open("archivo.txt", "w") as f:
    f.write(nuevo)
```

---

## Parte XI: Manejo de Errores

```python
try:
    with open("archivo.txt", "r") as f:
        contenido = f.read()
except FileNotFoundError:
    print("El archivo no existe")
except IOError:
    print("Error al leer el archivo")
except Exception as e:
    print(f"Error inesperado: {e}")
```

---

## Parte XII: Ejemplo Completo - Sistema de Log

```python
from datetime import datetime
from pathlib import Path

class Logger:
    def __init__(self, nombre_archivo="log.txt"):
        self.archivo = Path(nombre_archivo)
    
    def escribir(self, nivel, mensaje):
        """Escribir mensaje timestampeado en log"""
        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        línea = f"[{timestamp}] {nivel}: {mensaje}\n"
        
        with open(self.archivo, "a", encoding="utf-8") as f:
            f.write(línea)
    
    def info(self, mensaje):
        self.escribir("INFO", mensaje)
    
    def error(self, mensaje):
        self.escribir("ERROR", mensaje)
    
    def warning(self, mensaje):
        self.escribir("WARNING", mensaje)
    
    def leer_últimas(self, cantidad=10):
        """Leer últimas líneas del log"""
        with open(self.archivo, "r", encoding="utf-8") as f:
            líneas = f.readlines()
        return líneas[-cantidad:]

# Uso
logger = Logger("app.log")
logger.info("App iniciada")
logger.error("Error de conexión")
logger.warning("Usuario no encontrado")

print(" Últimas líneas:")
for línea in logger.leer_últimas(2):
    print(línea.strip())
```

---

## Resumen

| Operación | Código |
|-----------|--------|
| Leer todo | `f.read()` |
| Leer línea | `f.readline()` |
| Leer todas | `f.readlines()` |
| Iterar | `for línea in f:` |
| Escribir | `f.write(texto)` |
| Agregar | `open(file, "a")` |
| Existe | `Path(file).exists()` |
| Eliminar | `Path(file).unlink()` |
| Renombrar | `Path(file).rename(new)` |
| Pasta actual | `Path.cwd()` |

---

**Siguiente**: [json_serializacion.md](json_serializacion.md)
