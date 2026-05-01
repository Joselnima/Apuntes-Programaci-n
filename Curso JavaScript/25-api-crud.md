# Módulo 09b - Consumo de APIs y CRUD Completo

En el desarrollo web moderno es muy común tener que comunicarse con servidores (APIs) para leer, guardar, actualizar o borrar datos.

A este conjunto de operaciones se le conoce como **CRUD**:
- **C**reate (Crear): Método HTTP `POST`
- **R**ead (Leer): Método HTTP `GET`
- **U**pdate (Actualizar): Método HTTP `PUT` o `PATCH`
- **D**elete (Borrar): Método HTTP `DELETE`

## Configuración Base con Fetch API

La función `fetch` está disponible de forma nativa en el navegador y nos permite hacer estas peticiones. Aquí usamos la sintaxis moderna `async/await`.

### 1. Read (Leer - GET)
Sirve para obtener datos. Es el método por defecto de `fetch`.

```javascript
async function obtenerUsuarios() {
  try {
    const respuesta = await fetch('https://jsonplaceholder.typicode.com/users');
    const usuarios = await respuesta.json(); // Transformar a JSON
    console.log("Usuarios obtenidos:", usuarios);
  } catch (error) {
    console.error("Error al obtener usuarios:", error);
  }
}

obtenerUsuarios();
```

### 2. Create (Crear - POST)
Sirve para enviar datos nuevos al servidor.

```javascript
async function crearUsuario(nuevoUsuario) {
  try {
    const respuesta = await fetch('https://jsonplaceholder.typicode.com/users', {
      method: 'POST', // Especificar el método
      headers: {
        'Content-Type': 'application/json' // Decirle al servidor que enviamos JSON
      },
      body: JSON.stringify(nuevoUsuario) // Convertir el objeto JS a string JSON
    });
    
    const usuarioCreado = await respuesta.json();
    console.log("Usuario creado:", usuarioCreado);
  } catch (error) {
    console.error("Error al crear:", error);
  }
}

// Probando la creación
crearUsuario({ name: 'Juan Perez', email: 'juan@email.com' });
```

### 3. Update (Actualizar - PUT/PATCH)
Sirve para modificar datos existentes. `PUT` suele reemplazar todo el recurso y `PATCH` solo algunos campos.

```javascript
async function actualizarUsuario(id, datosActualizados) {
  try {
    // Nota el ID en la URL
    const respuesta = await fetch(`https://jsonplaceholder.typicode.com/users/${id}`, {
      method: 'PUT', 
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(datosActualizados)
    });
    
    const usuarioActualizado = await respuesta.json();
    console.log("Usuario modificado:", usuarioActualizado);
  } catch (error) {
    console.error("Error al actualizar:", error);
  }
}

// Probando la actualización del usuario con ID 1
actualizarUsuario(1, { name: 'Juan Editado', email: 'nuevo@email.com' });
```

### 4. Delete (Borrar - DELETE)
Sirve para eliminar un recurso.

```javascript
async function borrarUsuario(id) {
  try {
    const respuesta = await fetch(`https://jsonplaceholder.typicode.com/users/${id}`, {
      method: 'DELETE'
    });
    
    if (respuesta.ok) {
      console.log(`Usuario ${id} borrado exitosamente.`);
    }
  } catch (error) {
    console.error("Error al borrar:", error);
  }
}

// Probando el borrado del usuario con ID 1
borrarUsuario(1);
```

## Resumen

Saber cómo interactuar con APIs REST usando **fetch** y métodos **HTTP (GET, POST, PUT, DELETE)** es fundamental para construir aplicaciones dinámicas e interactivas.
