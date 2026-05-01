# Módulo 06 - Objetos

Los objetos son colecciones de datos relacionados. En JavaScript, un objeto agrupa propiedades y funciones.

## Crear un objeto

```javascript
const persona = {
  nombre: 'Ana',
  edad: 28,
  ciudad: 'Madrid'
};
```

## Acceder a propiedades

```javascript
console.log(persona.nombre);
console.log(persona['edad']);
```

## Modificar propiedades

```javascript
persona.edad = 29;
persona.pais = 'España';
```

## Métodos

```javascript
const persona = {
  nombre: 'Ana',
  saludar() {
    console.log(`Hola, soy ${this.nombre}`);
  }
};

persona.saludar();
```

## Borrar propiedades

```javascript
delete persona.ciudad;
```

## Recorrer propiedades

```javascript
for (const clave in persona) {
  console.log(clave, persona[clave]);
}
```

## Objetos anidados

```javascript
const producto = {
  nombre: 'Camiseta',
  precio: 25,
  talla: {
    disponible: true,
    colores: ['azul', 'negro']
  }
};

console.log(producto.talla.colores[0]); // 'azul'
// Uso de Optional Chaining (ES11) para evitar errores si la propiedad no existe
console.log(producto.marca?.pais); // undefined (no rompe la app)
```

## Métodos útiles de `Object`

JavaScript proporciona métodos globales muy útiles para trabajar con objetos.

### 1. Claves, Valores y Entradas
Para extraer información rápidamente:
```javascript
const user = { id: 1, rol: 'admin' };

console.log(Object.keys(user)); // ['id', 'rol']
console.log(Object.values(user)); // [1, 'admin']
console.log(Object.entries(user)); // [['id', 1], ['rol', 'admin']]
```

### 2. Copiar y fusionar: `Object.assign()`
Copia las propiedades de uno o más objetos hacia un objeto destino.
```javascript
const base = { a: 1 };
const extra = { b: 2 };
const combinado = Object.assign({}, base, extra); // { a: 1, b: 2 }
```
> *Nota: Hoy en día es más común usar el operador Spread (`{ ...base, ...extra }`).*

### 3. Proteger Objetos: `freeze()` y `seal()`
- `Object.freeze()`: Congela el objeto por completo. No se pueden añadir, borrar ni cambiar sus propiedades.
- `Object.seal()`: Sella el objeto. No se pueden añadir ni borrar propiedades, pero **sí se pueden cambiar** los valores de las existentes.

```javascript
const config = { modo: 'oscuro' };
Object.freeze(config);
config.modo = 'claro'; // No hace nada (o lanza error en modo estricto)

const perfil = { nivel: 5 };
Object.seal(perfil);
perfil.nivel = 6; // Esto SÍ funciona
perfil.nuevaProp = 'hola'; // Esto NO funciona
```

## Getters y Setters

Permiten ejecutar código cada vez que se intenta leer o escribir una propiedad, actuando como propiedades virtuales.

```javascript
const cuenta = {
  _saldo: 100, // Convención: el guión bajo indica que es "privada"
  
  get saldo() {
    return this._saldo;
  },
  
  set saldo(cantidad) {
    if (cantidad < 0) {
      console.log('No puedes tener saldo negativo');
    } else {
      this._saldo = cantidad;
    }
  }
};

console.log(cuenta.saldo); // 100 (Llama al getter)
cuenta.saldo = -50; // "No puedes tener saldo negativo" (Llama al setter)
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Objetos</title>
</head>
<body>
  <h1>Objetos en JavaScript</h1>
  <script>
    const coche = {
      marca: 'Toyota',
      modelo: 'Corolla',
      anio: 2020,
      mostrarInfo() {
        console.log(`${this.marca} ${this.modelo} (${this.anio})`);
      }
    };

    coche.mostrarInfo();
  </script>
</body>
</html>
```

## Ejercicio

1. Crea un objeto `libro` con título, autor y año.
2. Añade un método `describir` que muestre la información.
3. Crea un objeto `pelicula` con una lista de actores.

## Resumen

Los objetos son esenciales en JavaScript porque representan datos del mundo real y permiten agrupar información de forma clara.
