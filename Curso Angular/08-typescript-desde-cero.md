# Módulo 08 - TypeScript Desde Cero

> **JavaScript con tipos - Type Safety para Angular**

---

## ¿QUÉ es TypeScript?

**TypeScript = JavaScript + Tipos**

```typescript
// ❌ JavaScript (permisivo)
const suma = (a, b) => a + b;
suma(5, "texto");  // ¿Qué pasa?

// ✅ TypeScript (seguro)
const suma = (a: number, b: number): number => a + b;
suma(5, "texto");  // ❌ ERROR INMEDIATO
```

Angular REQUIERE TypeScript. Es obligatorio.

---

## ¿PARA QUÉ?

### Atrapa errores ANTES de producción

```typescript
// ❌ Sin tipos - Bug silencioso
function usuario(datos) {
  return datos.nombre.toUpperCase();
}
usuario(null);  // CRASH 💥

// ✅ Con tipos - Error detectado
function usuario(datos: Usuario): string {
  return datos.nombre.toUpperCase();
}
usuario(null);  // ❌ Error en escritura
```

---

## ¿CÓMO funciona?

### 1. Tipos Básicos

```typescript
// String
const nombre: string = "Juan";

// Number
const edad: number = 25;

// Boolean
const activo: boolean = true;

// Any (evita - es "escape" de tipos)
const cualquier: any = "puede ser cualquier cosa";
```

### 2. Interfaces (MUY IMPORTANTE EN ANGULAR)

```typescript
// Define estructura
interface Usuario {
  id: number;
  nombre: string;
  email: string;
  edad?: number;  // Opcional
}

// Usar la interfaz
const usuario: Usuario = {
  id: 1,
  nombre: "Juan",
  email: "juan@ejemplo.com"
  // edad es opcional, OK
};

// Si falta campo → ERROR
const usuarioMalo: Usuario = {
  id: 1,
  nombre: "Juan"
  // ❌ Falta email
};
```

### 3. Arrays

```typescript
const nombres: string[] = ["Juan", "María"];
const numeros: number[] = [1, 2, 3];

// Array de objetos
interface Producto {
  id: number;
  nombre: string;
  precio: number;
}

const productos: Producto[] = [
  { id: 1, nombre: "Laptop", precio: 1000 }
];
```

### 4. Funciones

```typescript
// Parámetros y retorno tipados
function duplicar(numero: number): number {
  return numero * 2;
}

// Arrow function
const sumar = (a: number, b: number): number => a + b;

// Parámetros opcionales
function saludar(nombre: string, apellido?: string): string {
  return `Hola ${nombre} ${apellido || ""}`;
}

// Parámetros con valores por defecto
function contar(desde: number = 0): void {
  console.log(desde);
}
```

### 5. Clases (Angular usa MUCHO clases)

```typescript
class Persona {
  nombre: string;
  edad: number;
  
  constructor(nombre: string, edad: number) {
    this.nombre = nombre;
    this.edad = edad;
  }
  
  saludar(): string {
    return `Hola, soy ${this.nombre}`;
  }
}

const persona = new Persona("Juan", 25);
console.log(persona.saludar());
```

### 6. Union Types (múltiples opciones)

```typescript
// Puede ser string O number
const id: string | number = 123;
const id2: string | number = "ABC";

// Más complicado
type Resultado = "éxito" | "error" | "pendiente";
const estado: Resultado = "éxito";
```

### 7. Genéricos (reutilizar con diferentes tipos)

```typescript
// Función genérica
function obtenerPrimero<T>(items: T[]): T {
  return items[0];
}

// Usar
const primer = obtenerPrimero<number>([1, 2, 3]);  // type: number
const primerNombre = obtenerPrimero<string>(["Juan", "María"]);  // type: string

// Interfaz genérica
interface Respuesta<T> {
  estado: "éxito" | "error";
  datos: T;
}

const respuestaUsuarios: Respuesta<Usuario[]> = {
  estado: "éxito",
  datos: [{ id: 1, nombre: "Juan", email: "..." }]
};
```

### 8. En Angular (componentes)

```typescript
import { Component } from '@angular/core';

interface Producto {
  id: number;
  nombre: string;
  precio: number;
}

@Component({
  selector: 'app-tienda',
  template: `
    <div *ngFor="let p of productos">
      {{ p.nombre }} - ${{ p.precio }}
    </div>
  `
})
export class TiendaComponent {
  productos: Producto[] = [
    { id: 1, nombre: "Laptop", precio: 1000 },
    { id: 2, nombre: "Mouse", precio: 50 }
  ];
  
  agregarProducto(nombre: string, precio: number): void {
    const id = this.productos.length + 1;
    this.productos.push({ id, nombre, precio });
  }
}
```

---

## Tabla de Tipos

| Tipo | Ejemplo | Descripción |
|------|---------|-------------|
| `string` | `"Juan"` | Texto |
| `number` | `25` | Números |
| `boolean` | `true` | Verdadero/Falso |
| `string[]` | `["a", "b"]` | Array de strings |
| `number \| string` | `25` o `"A"` | Union type |
| `interface` | Define estructura | Objeto tipado |
| `T` genérico | Reutilizable | Para cualquier tipo |
| `?` | Propiedad opcional | Puede no existir |
| `readonly` | No se modifica | Protege datos |

---

## Errores Comunes

### ❌ Usar `any`

```typescript
const datos: any = fetch("/api");  // Pierde todo tipo
```

### ✅ Especificar tipo

```typescript
const datos: Promise<Response> = fetch("/api");
```

---

## Checklist TypeScript

- [ ] Tipos básicos (string, number, boolean)
- [ ] Interfaces (definir estructuras)
- [ ] Arrays de objetos
- [ ] Funciones tipadas
- [ ] Clases en Angular
- [ ] Genéricos básicos
- [ ] Union types

---

## Próximo

Módulo 09: Componentes en Angular

---

## Resumen

Si puedes explicarlo con tus palabras, vas bien.
Si solo lo reconoces cuando lo ves, todavía falta práctica.
