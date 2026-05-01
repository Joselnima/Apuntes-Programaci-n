# Módulo 27 - Manejo de Archivos (File, Blob, FormData)

Manipular archivos y enviarlos al servidor es una parte esencial del desarrollo web. El navegador provee herramientas seguras para leer datos seleccionados por el usuario.

## 1. El objeto `File` y el input file

Cuando el usuario selecciona un archivo usando `<input type="file">`, JavaScript puede acceder a sus metadatos (nombre, tamaño, tipo) mediante la propiedad `.files` del input.

```html
<input type="file" id="subidaArchivo" accept="image/png, image/jpeg" />

<script>
  const input = document.getElementById('subidaArchivo');
  
  input.addEventListener('change', (evento) => {
    // evento.target.files es una lista de archivos (FileList)
    const archivo = evento.target.files[0];
    
    if (archivo) {
      console.log(`Nombre: ${archivo.name}`);
      console.log(`Tamaño: ${archivo.size} bytes`);
      console.log(`Tipo: ${archivo.type}`);
    }
  });
</script>
```

## 2. Leer archivos con `FileReader`

Si deseas mostrar una vista previa de una imagen o leer el texto de un archivo `.txt`, necesitas usar la clase `FileReader`.

```javascript
function leerTextoDeArchivo(archivo) {
  const lector = new FileReader();

  // El evento 'load' se dispara cuando la lectura termina
  lector.addEventListener('load', (e) => {
    console.log("Contenido del archivo:", e.target.result);
  });

  // Ordenamos leer el archivo como texto
  lector.readAsText(archivo); 
  
  // Para imágenes usarías: lector.readAsDataURL(archivo);
}
```

## 3. `Blob` (Binary Large Object)

Un Blob representa datos inmutables en bruto (raw data). El objeto `File` en realidad hereda de `Blob`. Puedes crear Blobs desde cero en JavaScript para, por ejemplo, generar un archivo y forzar su descarga.

```javascript
// Creamos un Blob con texto plano
const contenido = ["Hola, esto es un archivo de texto creado desde JS!"];
const miBlob = new Blob(contenido, { type: 'text/plain' });

// Crear un enlace temporal para descargar el Blob
const urlBlob = URL.createObjectURL(miBlob);

const enlace = document.createElement('a');
enlace.href = urlBlob;
enlace.download = 'archivo_generado.txt';
enlace.click(); // Forzar descarga

// Liberar memoria
URL.revokeObjectURL(urlBlob);
```

## 4. `FormData` (Para enviar archivos)

La forma estándar de enviar archivos por una API es usando `multipart/form-data`. El objeto `FormData` construye este formato automáticamente por nosotros.

```javascript
async function subirArchivo(archivo) {
  const formData = new FormData();
  
  // Añadimos datos al formulario
  formData.append("nombreUsuario", "Carlos"); // Texto normal
  formData.append("fotoPerfil", archivo);     // Archivo (File o Blob)

  try {
    const respuesta = await fetch("https://api.tu-servidor.com/upload", {
      method: "POST",
      // ¡NO pongas Content-Type en los headers manualmente con FormData!
      // El navegador lo calcula y añade los límites automáticamente
      body: formData 
    });
    
    const resultado = await respuesta.json();
    console.log(resultado);
  } catch (error) {
    console.error("Error al subir archivo", error);
  }
}
```
