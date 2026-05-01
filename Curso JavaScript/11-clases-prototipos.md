# Módulo 12 - Clases y prototipos

JavaScript permite crear objetos con prototipos. Las clases son una forma más clara de hacerlo.

## Prototipos y La Cadena de Prototipos (Prototype Chain)

En JavaScript todo objeto tiene una propiedad oculta llamada `[[Prototype]]` (a la que podemos acceder mediante `__proto__`). Cuando intentas leer la propiedad de un objeto, si el objeto no la tiene, el motor de JS buscará en su Prototipo. Esto crea una **Cadena de Prototipos** que hereda funcionalidades.

```javascript
// Prototipo clásico con Funciones Constructoras
function Persona(nombre, edad) {
  this.nombre = nombre;
  this.edad = edad;
}

// Agregamos métodos al prototipo para no duplicar memoria en cada objeto nuevo
Persona.prototype.saludar = function() {
  console.log(`Hola, soy ${this.nombre}`);
};

const ana = new Persona('Ana', 28);
ana.saludar();

console.log(ana.__proto__ === Persona.prototype); // true
```

### Herencia directa con `Object.create()`
Si quieres crear objetos con un prototipo específico sin usar clases ni constructores:

```javascript
const animalProto = { 
  vivo: true,
  respirar() { console.log("Respirando..."); }
};

// Creamos un perro que hereda de animalProto
const perro = Object.create(animalProto);
perro.raza = "Labrador";

perro.respirar(); // "Respirando..." (Lo encuentra en su cadena de prototipos)
```

## Clases

```javascript
class Persona {
  constructor(nombre, edad) {
    this.nombre = nombre;
    this.edad = edad;
  }

  saludar() {
    console.log(`Hola, soy ${this.nombre}`);
  }
}

const luis = new Persona('Luis', 32);
luis.saludar();
```

## Los 4 Pilares de la Programación Orientada a Objetos (POO)

JavaScript permite implementar los cuatro pilares fundamentales de la POO:

### 1. Abstracción
Es ocultar los detalles complejos y mostrar solo las funcionalidades esenciales.
En clases, se logra creando métodos que hacen el trabajo pesado internamente.

### 2. Encapsulamiento
Es ocultar el estado (variables) de un objeto para que no pueda ser modificado directamente desde afuera. En JavaScript moderno, usamos `#` para crear propiedades privadas.

```javascript
class CuentaBancaria {
  #saldo; // Propiedad privada

  constructor(titular, saldoInicial) {
    this.titular = titular;
    this.#saldo = saldoInicial;
  }

  depositar(cantidad) {
    if (cantidad > 0) this.#saldo += cantidad;
  }

  obtenerSaldo() {
    return this.#saldo;
  }
}

const miCuenta = new CuentaBancaria('Ana', 100);
miCuenta.depositar(50);
console.log(miCuenta.obtenerSaldo()); // 150
// console.log(miCuenta.#saldo); // Error: Propiedad privada
```

### 3. Herencia
Permite que una clase hija herede métodos y propiedades de una clase padre, usando `extends` y `super()`.

### 4. Polimorfismo
Es la capacidad de clases hijas de sobrescribir (redefinir) métodos heredados de la clase padre para que se comporten diferente.

```javascript
// Ejemplo de Herencia y Polimorfismo
class Animal {
  hacerSonido() {
    console.log("Sonido genérico");
  }
}

class Perro extends Animal {
  // Sobrescribiendo el método (Polimorfismo)
  hacerSonido() {
    console.log("Guau guau!");
  }
}

class Gato extends Animal {
  // Sobrescribiendo el método (Polimorfismo)
  hacerSonido() {
    console.log("Miau!");
  }
}

const animales = [new Animal(), new Perro(), new Gato()];
animales.forEach(a => a.hacerSonido()); 
// Imprime: "Sonido genérico", "Guau guau!", "Miau!"
```

## Herencia Avanzada (Ejemplo extra)
class Empleado extends Persona {
  constructor(nombre, edad, puesto) {
    super(nombre, edad);
    this.puesto = puesto;
  }

  mostrarPuesto() {
    console.log(`Trabajo como ${this.puesto}`);
  }
}

const ana = new Empleado('Ana', 28, 'Desarrolladora');
ana.saludar();
ana.mostrarPuesto();
```

## Métodos estáticos

Un método estático pertenece a la propia CLASE, no a las instancias. Se usa para funciones de utilidad que no dependen de los datos de un objeto.

```javascript
class Matematica {
  static sumar(a, b) {
    return a + b;
  }
}

// No necesitas hacer 'new Matematica()'. Se llama directo en la clase:
console.log(Matematica.sumar(5, 7)); // 12
```

## Alternativa Moderna: Factory Functions (Fábricas)

A muchos desarrolladores avanzados no les gusta la palabra `class` ni tener que pelear con la palabra `this` (que a veces pierde el contexto). Prefieren usar funciones que fabrican y devuelven un objeto utilizando *Closures* para la privacidad.

```javascript
function crearCuenta(titular, saldoInicial) {
  let saldo = saldoInicial; // Totalmente privado, sin necesidad de '#'

  return {
    titular,
    depositar(cantidad) {
      saldo += cantidad;
    },
    obtenerSaldo() {
      return saldo;
    }
  };
}

const miCuenta = crearCuenta('Carlos', 500);
miCuenta.depositar(200);
console.log(miCuenta.obtenerSaldo()); // 700
console.log(miCuenta.saldo); // undefined (No hay forma de tocarlo desde afuera)
```

## Ejemplo práctico

```html
<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Clases</title>
</head>
<body>
  <h1>Clases y prototipos</h1>
  <script>
    class Persona {
      constructor(nombre, edad) {
        this.nombre = nombre;
        this.edad = edad;
      }

      saludar() {
        console.log(`Hola, soy ${this.nombre}`);
      }
    }

    class Estudiante extends Persona {
      constructor(nombre, edad, curso) {
        super(nombre, edad);
        this.curso = curso;
      }

      estudiar() {
        console.log(`${this.nombre} está estudiando ${this.curso}`);
      }
    }

    const estudiante = new Estudiante('Marta', 22, 'JavaScript');
    estudiante.saludar();
    estudiante.estudiar();
  </script>
</body>
</html>
```

## Ejercicio

1. Crea una clase `Animal` con nombre y sonido.
2. Crea una clase `Perro` que herede de `Animal` y añada un método `ladrar`.
3. Crea un método estático en una clase de utilidad.

## Resumen

Las clases son una forma ordenada de crear objetos con comportamiento. Los prototipos son el motor interno que hace que las clases funcionen.
