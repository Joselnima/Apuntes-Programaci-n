# Web Básico - HTTP, Requests, y Frameworks

> **Comunicación entre computadoras a través de internet**

---

## ¿QUÉ es la Web?

La **web** = computadoras hablando entre sí:

```
Tu Computadora          Internet           Servidor Google
(Cliente)                 📡                 (Servidor)
    │
    └─ "Dame la página google.com" ──────────→
                                           Servidor procesa
                                           Servidor prepara
    ← "Aquí está: <html>...</html>" ───────┘
    │
    └─ Navegador muestra página
```

**Conceptos:**
- 💻 **Cliente** = tu computadora (pide cosas)
- 🖥️ **Servidor** = computadora remota (responde)
- 🌐 **HTTP** = idioma que hablan

---

## ¿PARA QÉ sirve?

### 1. Obtener datos de internet
```python
# "Dame la lista de usuarios del servidor"
respuesta = requests.get('https://api.ejemplo.com/usuarios')
usuarios = respuesta.json()
```

### 2. Enviar datos al servidor
```python
# "Guarda este nuevo usuario"
nuevo_usuario = {"nombre": "Juan", "edad": 25}
respuesta = requests.post('https://api.ejemplo.com/usuarios', json=nuevo_usuario)
```

### 3. Crear tu propio servidor
```python
# Tu código recibe peticiones de usuarios
@app.route('/saludar')
def saludar():
    return "¡Hola!"

# Usuario accede: http://misitiio.com/saludar
```

### 4. APIs - Integrar servicios
```python
# Usar API de clima, maps, pagos, etc
# Sin crear los servicios tú mismo
```

---

## ¿CÓMO funcionan?

### HTTP: Pregunta y Respuesta

```
Cliente                           Servidor
  │                                 │
  ├─ PREGUNTA ─────────────────────→
  │  GET /api/usuarios
  │  Host: ejemplo.com
  │
  ←──────────────── RESPUESTA ──────┤
  │  Status: 200 OK
  │  Content: [...]
  │
```

### Ciclo Completo: GET

```python
1. ENVÍO      → requests.get("url", params={...})
2. SERVIDOR   → Recibe, procesa
3. RESPUESTA  → 200 OK + datos JSON
4. CLIENTE    → respuesta.json() obtiene datos
```

### 4 Operaciones Básicas

| Operación | Qué hace | Ejemplo |
|-----------|----------|---------|
| **GET** | Obtener datos | Leer lista de usuarios |
| **POST** | Crear nuevo | Agregar usuario |
| **PUT** | Cambiar todo | Actualizar usuario |
| **DELETE** | Borrar | Eliminar usuario |

---

## Parte I: HTTP Básico

HTTP = protocolo de transferencia de hipertexto

### Métodos HTTP

```
GET     - Obtener recurso
POST    - Crear recurso
PUT     - Actualizar completo
PATCH   - Actualizar parcial
DELETE  - Eliminar recurso
HEAD    - GET sin cuerpo
OPTIONS - Métodos permitidos
```

### Códigos de Respuesta

```
200 OK              - Éxito
201 Created         - Creado
400 Bad Request     - Solicitud inválida
401 Unauthorized    - No autenticado
403 Forbidden       - Acceso denegado
404 Not Found       - No existe
500 Server Error    - Error del servidor
```

---

## Parte II: requests - Cliente HTTP

```bash
pip install requests
```

### GET - Obtener datos

```python
import requests

# Solicitud simple
respuesta = requests.get('https://api.ejemplo.com/usuarios/1')

# Verificar éxito
if respuesta.status_code == 200:
    datos = respuesta.json()
    print(datos)

# Con parámetros
params = {'página': 1, 'límite': 10}
respuesta = requests.get('https://api.ejemplo.com/usuarios', params=params)

# Con headers
headers = {'Authorization': 'Bearer mi_token'}
respuesta = requests.get('https://api.ejemplo.com/usuarios', headers=headers)

# Timeout
respuesta = requests.get('https://api.ejemplo.com/usuarios', timeout=5)
```

### POST - Crear recurso

```python
import requests
import json

# Con diccionario
datos = {
    'nombre': 'Juan',
    'email': 'juan@example.com',
    'edad': 25
}

respuesta = requests.post(
    'https://api.ejemplo.com/usuarios',
    json=datos,  # Automáticamente serializado a JSON
    headers={'Content-Type': 'application/json'}
)

if respuesta.status_code == 201:
    usuario_creado = respuesta.json()
    print(usuario_creado)
```

### PUT/PATCH - Actualizar

```python
import requests

# Actualizar completo (PUT)
respuesta = requests.put(
    'https://api.ejemplo.com/usuarios/1',
    json={'nombre': 'Juan García', 'edad': 26}
)

# Actualizar parcial (PATCH)
respuesta = requests.patch(
    'https://api.ejemplo.com/usuarios/1',
    json={'nombre': 'Juan García'}
)
```

### DELETE - Eliminar

```python
import requests

respuesta = requests.delete('https://api.ejemplo.com/usuarios/1')

if respuesta.status_code == 204:
    print("Eliminado correctamente")
```

### Manejo de Errores

```python
import requests
from requests.exceptions import RequestException, Timeout, ConnectionError

try:
    respuesta = requests.get('https://api.ejemplo.com/usuarios', timeout=5)
    respuesta.raise_for_status()  # Lanza error si status >= 400
    datos = respuesta.json()
except Timeout:
    print("Solicitud se agotó (timeout)")
except ConnectionError:
    print("Error de conexión")
except requests.HTTPError as e:
    print(f"Error HTTP: {e}")
except Exception as e:
    print(f"Error inesperado: {e}")
```

---

## Parte III: JSONPlaceholder - API de Prueba

```python
import requests

# API pública gratuita para pruebas
# https://jsonplaceholder.typicode.com

# Obtener posts
respuesta = requests.get('https://jsonplaceholder.typicode.com/posts/1')
post = respuesta.json()
print(post)
# {
#   "userId": 1,
#   "id": 1,
#   "title": "...",
#   "body": "..."
# }

# Crear post
nuevo_post = {
    'userId': 1,
    'title': 'Mi primer post',
    'body': 'Contenido del post'
}
respuesta = requests.post(
    'https://jsonplaceholder.typicode.com/posts',
    json=nuevo_post
)
print(respuesta.status_code)  # 201
print(respuesta.json())
```

---

## Parte IV: Beautiful Soup - Web Scraping

```bash
pip install beautifulsoup4
```

```python
import requests
from bs4 import BeautifulSoup

# Descargar página
respuesta = requests.get('https://ejemplo.com')
html = respuesta.content

# Parsear HTML
soup = BeautifulSoup(html, 'html.parser')

# Buscar elementos
títulos = soup.find_all('h1')
for título in títulos:
    print(título.text)

# Buscar por clase
artículos = soup.find_all('div', class_='artículo')

# Buscar por ID
header = soup.find(id='header')

# Navegar árbol
enlace = soup.select_one('a.importante')
print(enlace.get('href'))
```

---

## Parte V: Flask - Framework Web Minimalista

```bash
pip install flask
```

### Servidor Básico

```python
from flask import Flask, request, jsonify

app = Flask(__name__)

# Ruta GET
@app.route('/usuarios', methods=['GET'])
def obtener_usuarios():
    usuarios = [
        {'id': 1, 'nombre': 'Juan'},
        {'id': 2, 'nombre': 'María'},
    ]
    return jsonify(usuarios)

# Ruta GET con parámetro
@app.route('/usuarios/<int:id>', methods=['GET'])
def obtener_usuario(id):
    usuario = {'id': id, 'nombre': 'Juan'}
    return jsonify(usuario)

# Ruta POST
@app.route('/usuarios', methods=['POST'])
def crear_usuario():
    datos = request.json
    nuevo_usuario = {
        'id': 3,
        'nombre': datos['nombre']
    }
    return jsonify(nuevo_usuario), 201

# Ruta DELETE
@app.route('/usuarios/<int:id>', methods=['DELETE'])
def eliminar_usuario(id):
    return '', 204

if __name__ == '__main__':
    app.run(debug=True, port=5000)
```

Ejecutar:
```bash
python app.py
# Acceder: http://localhost:5000/usuarios
```

---

## Parte VI: FastAPI - Framework Moderno

```bash
pip install fastapi uvicorn
```

```python
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class Usuario(BaseModel):
    nombre: str
    email: str
    edad: int

usuarios_db = []

@app.get("/usuarios")
def obtener_usuarios():
    return usuarios_db

@app.get("/usuarios/{id}")
def obtener_usuario(id: int):
    return {"id": id, "nombre": "Juan"}

@app.post("/usuarios", status_code=201)
def crear_usuario(usuario: Usuario):
    usuarios_db.append(usuario.dict())
    return usuario

@app.put("/usuarios/{id}")
def actualizar_usuario(id: int, usuario: Usuario):
    return usuario

@app.delete("/usuarios/{id}", status_code=204)
def eliminar_usuario(id: int):
    pass

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
```

Ejecutar:
```bash
uvicorn app:app --reload
# Documentación automática: http://localhost:8000/docs
```

---

## Parte VII: Sesiones y Cookies

```python
import requests

# Sesión mantiene cookies automáticamente
session = requests.Session()

# Login
respuesta = session.post('https://ejemplo.com/login', 
    json={'usuario': 'juan', 'contraseña': '123'})

# Las cookies se guardan automáticamente
respuesta = session.get('https://ejemplo.com/perfil')
print(respuesta.json())  # Autenticado

# Ver cookies
print(session.cookies)
```

---

## Parte VIII: API REST Ejemplo Completo

```python
from flask import Flask, request, jsonify

app = Flask(__name__)

# BD en memoria
usuarios = {
    1: {'id': 1, 'nombre': 'Juan', 'email': 'juan@example.com'},
    2: {'id': 2, 'nombre': 'María', 'email': 'maria@example.com'},
}

@app.route('/usuarios', methods=['GET'])
def list_usuarios():
    return jsonify(list(usuarios.values()))

@app.route('/usuarios/<int:id>', methods=['GET'])
def get_usuario(id):
    if id not in usuarios:
        return {'error': 'No encontrado'}, 404
    return jsonify(usuarios[id])

@app.route('/usuarios', methods=['POST'])
def create_usuario():
    datos = request.json
    id_nuevo = max(usuarios.keys()) + 1
    usuario = {
        'id': id_nuevo,
        'nombre': datos['nombre'],
        'email': datos['email']
    }
    usuarios[id_nuevo] = usuario
    return jsonify(usuario), 201

@app.route('/usuarios/<int:id>', methods=['PUT'])
def update_usuario(id):
    if id not in usuarios:
        return {'error': 'No encontrado'}, 404
    datos = request.json
    usuarios[id].update(datos)
    return jsonify(usuarios[id])

@app.route('/usuarios/<int:id>', methods=['DELETE'])
def delete_usuario(id):
    if id not in usuarios:
        return {'error': 'No encontrado'}, 404
    del usuarios[id]
    return '', 204

if __name__ == '__main__':
    app.run(debug=True)
```

---

## Parte IX: Autenticación con JWT

```bash
pip install pyjwt
```

```python
from flask import Flask, request, jsonify
import jwt
import datetime

app = Flask(__name__)
app.config['SECRET_KEY'] = 'tu-clave-secreta'

@app.route('/login', methods=['POST'])
def login():
    datos = request.json
    usuario = datos.get('usuario')
    
    # Validar credenciales (simplificado)
    if usuario == 'juan':
        token = jwt.encode({
            'usuario': usuario,
            'exp': datetime.datetime.utcnow() + datetime.timedelta(hours=1)
        }, app.config['SECRET_KEY'], algorithm='HS256')
        
        return jsonify({'token': token})
    
    return {'error': 'Credenciales inválidas'}, 401

def verificar_token(request):
    token = request.headers.get('Authorization', '').replace('Bearer ', '')
    try:
        datos = jwt.decode(token, app.config['SECRET_KEY'], algorithms=['HS256'])
        return datos
    except:
        return None

@app.route('/perfil', methods=['GET'])
def perfil():
    datos = verificar_token(request)
    if not datos:
        return {'error': 'No autorizado'}, 401
    
    return jsonify({'usuario': datos['usuario']})

if __name__ == '__main__':
    app.run(debug=True)
```

---

## Parte X: CORS - Compartir Entre Dominios

```bash
pip install flask-cors
```

```python
from flask import Flask
from flask_cors import CORS

app = Flask(__name__)
CORS(app)  # Permitir todas las solicitudes CORS

# O específico
CORS(app, resources={r"/api/*": {"origins": ["https://ejemplo.com"]}})

@app.route('/api/datos')
def datos():
    return {'mensaje': 'Accesible desde cualquier dominio'}

if __name__ == '__main__':
    app.run()
```

---

## Parte XI: WebSockets - Comunicación Bidireccional

```bash
pip install flask-socketio python-socketio
```

```python
from flask import Flask
from flask_socketio import SocketIO, emit

app = Flask(__name__)
socketio = SocketIO(app)

@socketio.on('conectar')
def handle_conectar():
    print('Cliente conectado')
    emit('respuesta', {'datos': 'Bienvenido'})

@socketio.on('mensaje')
def handle_mensaje(datos):
    print(f'Mensaje: {datos}')
    emit('respuesta', {'datos': f'Eco: {datos}'}, broadcast=True)

if __name__ == '__main__':
    socketio.run(app, debug=True)
```

---

## Resumen

| Herramienta | Usa |
|-------------|-----|
| requests | Cliente HTTP |
| BeautifulSoup | Web scraping |
| Flask | Framework web simple |
| FastAPI | Framework web moderno |
| JWT | Autenticación basada en tokens |
| CORS | Compartir entre dominios |
| WebSockets | Comunicación en vivo |

---

**Siguiente**: [data_science.md](data_science.md)
