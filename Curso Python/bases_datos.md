# Bases de Datos en Python

> **Guardar grandes cantidades de datos organizados y buscarlos rápidamente**

---

## ¿QUÉ es una Base de Datos?

Una **BD** = carpeta super organizada:

```
Archivo de texto:
Juan 25 juan@email.com
María 30 maria@email.com
Pedro 22 pedro@email.com

Base de Datos:
┌──────┬──────┬────────────┐
│ ID   │ Edad │ Email      │
├──────┼──────┼────────────┤
│ 1    │ 25   │ juan@...   │
│ 2    │ 30   │ maria@...  │
│ 3    │ 22   │ pedro@...  │
└──────┴──────┴────────────┘
```

**3 Conceptos:**
- **Tabla** = hoja de cálculo (filas y columnas)
- **Fila** = un registro (una persona)
- **Columna** = un campo (nombre, edad, etc)

**Analogía:**
- 📋 Lista en Python = notas en papel pegadas
- 🗄️ Base de Datos = archivador profesional con índices

---

## ¿PARA QÉ sirven?

### 1. Guardar MUCHOS datos
```python
# ❌ Con listas (lento con millones)
usuarios = [{"id": 1, "nombre": "Juan"}, ...]

# ✅ Con BD (rápido)
# Busca "Juan" en 1 millón de registros en microsegundos
```

### 2. Buscar datos específicos
```python
# "Dame todos los usuarios mayores de 25 años"
SELECT * FROM usuarios WHERE edad > 25;
```

### 3. Evitar duplicados
```python
# BD asegura: no hay dos usuarios con mismo email
email TEXT UNIQUE
```

### 4. Actualizar datos
```python
# El usuario cambió su edad sin borrar el resto
UPDATE usuarios SET edad = 26 WHERE id = 1;
```

---

## ¿CÓMO funcionan?

### 4 Operaciones Básicas (CRUD)

| Operación | Qué hace | Ejemplo |
|-----------|----------|---------|
| **C**reate | Insertar nuevo registro | `INSERT INTO usuarios...` |
| **R**ead | Buscar/leer datos | `SELECT * FROM usuarios...` |
| **U**pdate | Cambiar datos existentes | `UPDATE usuarios SET...` |
| **D**elete | Borrar registros | `DELETE FROM usuarios...` |

### Pasos para usar una BD

```
1. Conectar          → "Abro el archivo de BD"
2. Crear tabla       → "Hago las columnas"
3. Hacer operaciones → Insert, Select, Update, Delete
4. Guardar cambios   → commit()
5. Cerrar           → Desconectar
```

---

## Parte I: SQLite3 (En la Stdlib)

SQLite es una BD embebida, perfecta para aprender:

### Conectar y Crear Tabla

```python
import sqlite3

# Conectar (crea archivo si no existe)
conexión = sqlite3.connect("mibase.db")

# Obtener cursor (ejecutor de comandos)
cursor = conexión.cursor()

# Crear tabla
cursor.execute("""
    CREATE TABLE IF NOT EXISTS usuarios (
        id INTEGER PRIMARY KEY,
        nombre TEXT NOT NULL,
        edad INTEGER,
        email TEXT UNIQUE
    )
""")

# Guardar cambios
conexión.commit()

# Cerrar
conexión.close()
```

### Insertar Datos

```python
import sqlite3

conexión = sqlite3.connect("mibase.db")
cursor = conexión.cursor()

# Insertar uno
cursor.execute("""
    INSERT INTO usuarios (nombre, edad, email)
    VALUES (?, ?, ?)
""", ("Juan", 25, "juan@ejemplo.com"))

# Insertar múltiples
datos = [
    ("María", 30, "maria@ejemplo.com"),
    ("Pedro", 22, "pedro@ejemplo.com")
]
cursor.executemany("""
    INSERT INTO usuarios (nombre, edad, email)
    VALUES (?, ?, ?)
""", datos)

conexión.commit()
conexión.close()
```

**Importante**: Usa `?` para parámetros, NUNCA f-strings (SQL injection).

### Leer Datos

```python
import sqlite3

conexión = sqlite3.connect("mibase.db")
cursor = conexión.cursor()

# Leer todos
cursor.execute("SELECT * FROM usuarios")
usuarios = cursor.fetchall()
for usuario in usuarios:
    print(usuario)  # (1, 'Juan', 25, 'juan@ejemplo.com')

# Leer uno
cursor.execute("SELECT * FROM usuarios WHERE edad > ?", (25,))
usuario = cursor.fetchone()
print(usuario)

# Leer varios
cursor.execute("SELECT nombre FROM usuarios LIMIT 2")
nombres = cursor.fetchmany(2)
print(nombres)

conexión.close()
```

### Actualizar y Eliminar

```python
import sqlite3

conexión = sqlite3.connect("mibase.db")
cursor = conexión.cursor()

# Actualizar
cursor.execute("""
    UPDATE usuarios SET edad = ? WHERE nombre = ?
""", (31, "Juan"))

# Eliminar
cursor.execute("DELETE FROM usuarios WHERE edad < ?", (20,))

conexión.commit()
conexión.close()
```

### Context Manager (Recomendado)

Auto-cierra y maneja errores:

```python
import sqlite3

with sqlite3.connect("mibase.db") as conexión:
    cursor = conexión.cursor()
    
    cursor.execute("SELECT * FROM usuarios")
    for usuario in cursor:
        print(usuario)
    
    # commit() automático si no haya error
```

---

## Parte II: Clase para Encapsular DB

Patrón común:

```python
import sqlite3
from contextlib import contextmanager

class BD:
    def __init__(self, archivo="mibase.db"):
        self.archivo = archivo
    
    @contextmanager
    def conectar(self):
        conexión = sqlite3.connect(self.archivo)
        try:
            yield conexión
        finally:
            conexión.commit()
            conexión.close()
    
    def crear_tabla(self):
        with self.conectar() as con:
            cursor = con.cursor()
            cursor.execute("""
                CREATE TABLE IF NOT EXISTS usuarios (
                    id INTEGER PRIMARY KEY,
                    nombre TEXT,
                    edad INTEGER
                )
            """)
    
    def insertar(self, nombre, edad):
        with self.conectar() as con:
            cursor = con.cursor()
            cursor.execute(
                "INSERT INTO usuarios (nombre, edad) VALUES (?, ?)",
                (nombre, edad)
            )
    
    def obtener_todos(self):
        with self.conectar() as con:
            cursor = con.cursor()
            cursor.execute("SELECT * FROM usuarios")
            return cursor.fetchall()
    
    def obtener_por_nombre(self, nombre):
        with self.conectar() as con:
            cursor = con.cursor()
            cursor.execute("SELECT * FROM usuarios WHERE nombre = ?", (nombre,))
            return cursor.fetchone()
    
    def actualizar(self, id, edad):
        with self.conectar() as con:
            cursor = con.cursor()
            cursor.execute("UPDATE usuarios SET edad = ? WHERE id = ?", (edad, id))
    
    def eliminar(self, id):
        with self.conectar() as con:
            cursor = con.cursor()
            cursor.execute("DELETE FROM usuarios WHERE id = ?", (id,))

# Uso
bd = BD()
bd.crear_tabla()
bd.insertar("Juan", 25)
bd.insertar("María", 30)

print(bd.obtener_todos())
print(bd.obtener_por_nombre("Juan"))
```

---

## Parte III: SQLAlchemy ORM (Alternativa Moderna)

ORM: mapear tablas a clases de Python

### Instalación

```bash
pip install sqlalchemy
```

### Definir Modelos

```python
from sqlalchemy import create_engine, Column, String, Integer
from sqlalchemy.orm import declarative_base, Session

# Crear "base"
Base = declarative_base()

# Definir modelo
class Usuario(Base):
    __tablename__ = "usuarios"
    
    id = Column(Integer, primary_key=True)
    nombre = Column(String, nullable=False)
    edad = Column(Integer)
    email = Column(String, unique=True)
    
    def __repr__(self):
        return f"<Usuario(nombre={self.nombre}, edad={self.edad})>"

# Motor de BD
engine = create_engine("sqlite:///mibase.db")

# Crear tablas
Base.metadata.create_all(engine)
```

### CRUD con ORM

```python
from sqlalchemy.orm import Session

# Crear sesión
session = Session(engine)

# Create
usuario = Usuario(nombre="Juan", edad=25, email="juan@ejemplo.com")
session.add(usuario)
session.commit()

# Read
juan = session.query(Usuario).filter_by(nombre="Juan").first()
print(juan)

# Update
juan.edad = 26
session.commit()

# Delete
session.delete(juan)
session.commit()

session.close()
```

### Consultas Complejas

```python
# Todos
todos = session.query(Usuario).all()

# Filtro
mayores_25 = session.query(Usuario).filter(Usuario.edad > 25).all()

# Ordenar
ordenado = session.query(Usuario).order_by(Usuario.edad).all()

# Contar
cantidad = session.query(Usuario).count()

# SQL crudo (si necesitas)
resultado = session.execute("SELECT * FROM usuarios WHERE edad > ?", (25,))
```

---

## Parte IV: Relaciones (Foreign Keys)

```python
from sqlalchemy import ForeignKey
from sqlalchemy.orm import relationship

class Departamento(Base):
    __tablename__ = "departamentos"
    id = Column(Integer, primary_key=True)
    nombre = Column(String, nullable=False)

class Empleado(Base):
    __tablename__ = "empleados"
    id = Column(Integer, primary_key=True)
    nombre = Column(String, nullable=False)
    departamento_id = Column(Integer, ForeignKey("departamentos.id"))
    
    # Relación
    departamento = relationship("Departamento")

# Uso
dept = Departamento(nombre="IT")
emp = Empleado(nombre="Juan", departamento=dept)
session.add(emp)
session.commit()

# Acceder
print(emp.departamento.nombre)  # IT
```

---

## Parte V: Mejores Prácticas

### 1. Siempre Usa Parámetros

```python
# ❌ MAL - SQL Injection
nombre = "'; DROP TABLE usuarios; --"
cursor.execute(f"SELECT * FROM usuarios WHERE nombre = '{nombre}'")

# ✅ BIEN
cursor.execute("SELECT * FROM usuarios WHERE nombre = ?", (nombre,))
```

### 2. Usa Context Managers

```python
# ❌ MAL
conexión = sqlite3.connect("bd.db")
try:
    # hacer cosas
finally:
    conexión.close()

# ✅ BIEN
with sqlite3.connect("bd.db") as conexión:
    # hacer cosas
```

### 3. Transacciones

```python
with sqlite3.connect("bd.db") as conexión:
    cursor = conexión.cursor()
    try:
        cursor.execute("INSERT INTO usuarios VALUES (...)")
        cursor.execute("INSERT INTO logs VALUES (...)")
        # Si todo OK, auto-commit
    except Exception as e:
        conexión.rollback()  # Deshacer cambios
        raise
```

### 4. Índices para Búsquedas Rápidas

```python
cursor.execute("""
    CREATE INDEX idx_nombre ON usuarios(nombre)
""")
```

---

## Parte VI: Migrar datos

Ejemplo: importar CSV a BD

```python
import csv
import sqlite3

def importar_csv(archivo_csv, conexión):
    cursor = conexión.cursor()
    
    with open(archivo_csv, "r") as f:
        lector = csv.DictReader(f)
        for fila in lector:
            cursor.execute("""
                INSERT INTO usuarios (nombre, edad, email)
                VALUES (?, ?, ?)
            """, (fila["nombre"], fila["edad"], fila["email"]))
    
    conexión.commit()

with sqlite3.connect("mibase.db") as con:
    importar_csv("usuarios.csv", con)
```

---

## Parte VII: Exportar a JSON

```python
import json
import sqlite3

def exportar_json(tabla, archivo_salida):
    with sqlite3.connect("mibase.db") as con:
        con.row_factory = sqlite3.Row  # Convertir a diccionarios
        cursor = con.cursor()
        
        cursor.execute(f"SELECT * FROM {tabla}")
        filas = [dict(fila) for fila in cursor.fetchall()]
        
        with open(archivo_salida, "w") as f:
            json.dump(filas, f, indent=2)

exportar_json("usuarios", "usuarios.json")
```

---

## Comparativa: SQL puro vs ORM

| Aspecto | SQL puro | ORM |
|--------|----------|-----|
| Aprendizaje | Fácil, SQL explícito | Más abstracto |
| Flexibilidad | Total | Limitada a ORM |
| Performance | Muy bueno | Bueno, más overhead |
| Cambiar BD | Debe reescribirse | Relativamente fácil |
| Seguridad | Manual (parámetros) | Automática |
| Código | Verboso | Conciso |

---

## Resumen

| Operación | sqlite3 | SQLAlchemy |
|-----------|---------|-----------|
| Conectar | `sqlite3.connect()` | `create_engine()` |
| Crear tabla | cursor.execute(CREATE) | declarative_base() |
| Insertar | cursor.execute(INSERT) | session.add() |
| Leer | cursor.execute(SELECT) | session.query() |
| Actualizar | cursor.execute(UPDATE) | modificar + commit |
| Eliminar | cursor.execute(DELETE) | session.delete() |

---

**Siguiente**: [buenas_practicas.md](buenas_practicas.md)
