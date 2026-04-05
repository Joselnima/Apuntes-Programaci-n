# Módulo 10 - Templates y Data Binding

> **Conecta datos con HTML en tiempo real**

---

## ¿QUÉ es data binding?

**Sincronización automática: Datos ↔ Pantalla**

Analogía:
- Pizarra: Escribes un número
- Espejo: Se refleja automáticamente
- En Angular: cambias `counter`
- HTML se actualiza **sin recargar** ✨

---

## ¿PARA QUÉ?

### Problema 1: Actualizar HTML Manualmente (Pasado)

```html
<!-- jQuery era así -->
<h1 id="title">Hola</h1>

<script>
  // Cambias datos
  const name = 'Juan';
  
  // Actualizas HTML manualmente
  document.getElementById('title').textContent = name;
  // ¡Tedioso!
</script>
```

### Solución: Angular lo Hace Automático

```html
<!-- Angular -->
{{ name }}  <!-- HTML se actualiza automáticamente -->
```

**Tipos de binding:**

| Tipo | Símbolo | Dirección | Ejemplo |
|------|---------|-----------|---------|
| Interpolation | `{{ }}` | → (datos a HTML) | `{{ userName }}` |
| Property | `[ ]` | → (datos a HTML) | `[disabled]="isLoading"` |
| Event | `( )` | ← (HTML a datos) | `(click)="onClick()"` |
| Two-Way | `[( )]` | ↔ (bidireccional) | `[(ngModel)]="text"` |
| String | `[attr.]` | → (atributos) | `[attr.aria-label]="label"` |

---

## ¿CÓMO? Los 4 Tipos de Binding

### Tipo 1: Interpolation `{{ }}`

**Mostrar datos en HTML**

```typescript
// Componente
export class MyComponent {
  name = 'Juan';
  age = 25;
  price = 99.99;
}
```

```html
<!-- Template -->
<p>Hola, {{ name }}</p>
<p>Tienes {{ age }} años</p>
<p>Precio: ${{ price }}</p>

<!-- También puedes hacer operaciones -->
<p>Edad en 5 años: {{ age + 5 }}</p>
<p>Mayúsculas: {{ name.toUpperCase() }}</p>
<p>¿Mayor de edad?: {{ age >= 18 ? 'Si' : 'No' }}</p>
```

**Resultado:**

```
Hola, Juan
Tienes 25 años
Precio: $99.99

Edad en 5 años: 30
Mayúsculas: JUAN
¿Mayor de edad?: Si
```

### Tipo 2: Property Binding `[ ]`

**Cambiar propiedades HTML / Angular**

```typescript
export class MyComponent {
  isDisabled = false;
  imageUrl = 'https://example.com/photo.jpg';
  cssClass = 'active';
  isVisible = true;
}
```

```html
<!-- Cambiar atributos HTML -->
<input [disabled]="isDisabled">
<img [src]="imageUrl" alt="Foto">

<!-- Cambiar propiedades ng -->
<div [ngClass]="cssClass"></div>
<div [ngStyle]="{ color: 'red', fontSize: '20px' }"></div>

<!-- Cambiar propiedades customizadas (componentes propios) -->
<app-card [title]="'Mi Producto'" [price]="99.99"></app-card>
```

**Equivalentes:**

```html
<!-- ✅ Angular (recomendado)  -->
<input [disabled]="isDisabled">

<!-- ❌ HTML clásico NO funciona en Angular -->
<input disabled="isDisabled">

<!-- ✅ Se ve en el navegador -->
<input disabled>  <!-- Si isDisabled es true -->
```

### Tipo 3: Event Binding `( )`

**Escuchar eventos del usuario**

```typescript
export class MyComponent {
  counter = 0;
  
  increment() {
    this.counter++;
  }
  
  handleClick(event: MouseEvent) {
    console.log('Clickeaste en:', event);
  }
  
  handleKeyPress(event: KeyboardEvent) {
    if (event.key === 'Enter') {
      console.log('Presionaste Enter');
    }
  }
}
```

```html
<!-- Click -->
<button (click)="increment()">Incrementar</button>
<p>Contador: {{ counter }}</p>

<!-- Eventos con parámetros -->
<button (click)="handleClick($event)">
  Detectar evento
</button>

<!-- Teclado -->
<input (keypress)="handleKeyPress($event)">

<!-- Submit (formularios) -->
<form (submit)="onSubmit()">
  <input type="text">
</form>

<!-- Otros eventos -->
<input (blur)="onBlur()"     type="text">
<input (focus)="onFocus()"    type="text">
<input (change)="onChange()"  type="text">
<div (mouseover)="onHover()"></div>
```

**Eventos comunes:**

| Evento | Cuándo |
|--------|--------|
| `(click)` | Click del ratón |
| `(keypress)` | Presiona una tecla |
| `(keyup)` | Suelta una tecla |
| `(keydown)` | Baja una tecla |
| `(submit)` | Envía formulario |
| `(blur)` | Sale del input |
| `(focus)` | Entra al input |
| `(change)` | Valor cambió |
| `(scroll)` | Scroll |
| `(mouseover)` | Ratón sobre elemento |
| `(mouseout)` | Ratón sale |

### Tipo 4: Two-Way Binding `[( )]`

**Bidireccional: Usuario cambia HTML → datos se actualizan**

```typescript
export class MyComponent {
  fullName = '';
  email = '';
  agreed = false;
}
```

```html
<!-- Input de texto -->
<input [(ngModel)]="fullName" placeholder="Nombre completo">
<p>Escribiste: {{ fullName }}</p>

<!-- Textarea -->
<textarea [(ngModel)]="bio"></textarea>

<!-- Checkbox -->
<input [(ngModel)]="agreed" type="checkbox">
<p>Aceptaste: {{ agreed }}</p>

<!-- Select -->
<select [(ngModel)]="selectedCity">
  <option>Madrid</option>
  <option>Barcelona</option>
  <option>Valencia</option>
</select>
```

**Cómo funciona:**

```
1. Usuario escribe en <input>
2. Angular detecta cambio
3. Actualiza fullName en TypeScript
4. Actualiza {{ fullName }} en HTML
5. Todo sincronizado ✨

fullName: ''
↓
Usuario escribe "Juan"
↓
fullName: 'Juan'
↓
{{ fullName }} muestra "Juan"
```

**⚠️ Necesitas importar FormsModule:**

```typescript
import { FormsModule } from '@angular/forms';

@NgModule({
  imports: [FormsModule]
})
export class AppModule { }
```

---

## Combinando los 4 Tipos

**Ejemplo real: Carrito de compras**

```typescript
export class CartComponent {
  products = [
    { name: 'Laptop', price: 1000, quantity: 0 },
    { name: 'Mouse', price: 20, quantity: 0 },
    { name: 'Teclado', price: 50, quantity: 0 }
  ];
  
  getTotalPrice() {
    return this.products.reduce((sum, p) => 
      sum + (p.price * p.quantity), 0
    );
  }
  
  checkout() {
    console.log('Total:', this.getTotalPrice());
  }
}
```

```html
<div class="cart">
  <h1>Carrito</h1>
  
  <!-- Interpolación: mostrar datos -->
  <p>Total: ${{ getTotalPrice() }}</p>
  
  <!-- Event: botón -->
  <button (click)="checkout()">Finalizar compra</button>
  
  <!-- Lista con property + two-way -->
  <table>
    <tr>
      <th>Producto</th>
      <th>Precio</th>
      <th>Cantidad</th>
      <th>Subtotal</th>
    </tr>
    
    <tr *ngFor="let product of products">
      <!-- Interpolación -->
      <td>{{ product.name }}</td>
      <td>${{ product.price }}</td>
      
      <!-- Two-way: cambiar cantidad -->
      <td>
        <input [(ngModel)]="product.quantity" type="number">
      </td>
      
      <!-- Interpolación con operaciones -->
      <td>${{ product.price * product.quantity }}</td>
    </tr>
  </table>
</div>
```

---

## Property Binding vs Atributos

**¿Cuándo usar `[property]` vs `[attr.]`?**

```html
<!-- Property Binding (Lo normal) -->
<img [src]="imageUrl">
<input [disabled]="isDisabled">
<div [id]="elementId"></div>

<!-- Attribute Binding (Para atributos HTML especiales) -->
<div [attr.data-id]="123"></div>
<div [attr.aria-label]="'Cerrar'"></div>
<table [attr.border]="1"></table>

<!-- La mayoría de casos: property binding -->
<!-- Uses: src, href, id, class, style, disabled, checked, etc -->
```

---

## ngClass y ngStyle (Dinámicos)

### ngClass: Clases dinámicas

```typescript
export class StatusComponent {
  status = 'active';  // 'active' o 'inactive'
  warnings = 5;
}
```

```html
<!-- Forma 1: String -->
<div [ngClass]="status">
  Estado: {{ status }}
</div>

<!-- Forma 2: Objeto (más común) -->
<div [ngClass]="{
  'active': status === 'active',
  'warning': warnings > 0,
  'error': warnings > 10
}">
  Estado indicador
</div>

<!-- Forma 3: Array -->
<div [ngClass]="['base-class', status]"></div>
```

**CSS:**

```css
.active {
  color: green;
  border: 2px solid green;
}

.warning {
  color: orange;
  background: lightyellow;
}

.error {
  color: red;
  border: 2px solid red;
}
```

### ngStyle: Estilos dinámicos

```typescript
export class StyleComponent {
  textColor = 'blue';
  fontSize = 16;
  isHighlight = true;
}
```

```html
<!-- Forma 1: Objeto (recomendado) -->
<p [ngStyle]="{
  'color': textColor,
  'font-size.px': fontSize,
  'background': isHighlight ? 'yellow' : 'white'
}">
  Texto con estilos dinámicos
</p>

<!-- Forma 2: Variable con estilos -->
<p [ngStyle]="myStyles"></p>
```

```typescript
setTextColor() {
  this.myStyles = {
    'color': 'red',
    'font-weight': 'bold'
  };
}
```

---

## Safe Navigation y Elvis Operator

**Cuando no estás seguro si existe la propiedad:**

```typescript
export class UserComponent {
  user: any = null;  // ¿Qué pasa al inicio?
}
```

```html
<!-- ❌ Error: user es null -->
<p>{{ user.name }}</p>

<!-- ✅ Safe Navigation (sin error) -->
<p>{{ user?.name }}</p>

<!-- ✅ Alternativa con ||  -->
<p>{{ user?.name || 'Usuario no cargado' }}</p>
```

---

## Performance: Evitar Cálculos en Templates

**❌ Lento:**

```html
<!-- Esto se recalcula CADA cambio -->
<p>{{ products.filter(p => p.price > 100).length }}</p>

<!-- Se ejecuta muchísimas veces -->
<p>{{ expensiveCalculation() }}</p>
```

**✅ Rápido:**

```typescript
export class ListComponent {
  // Calcula UNA VEZ en ngOnInit
  expensiveProducts: Product[];
  
  ngOnInit() {
    this.expensiveProducts = this.products
      .filter(p => p.price > 100);
  }
}
```

```html
<!-- Muestra resultado calculado -->
<p>{{ expensiveProducts.length }}</p>
```

---

## Checklist - Templates Profundionales

- [ ] Usas `{{ }}` para mostrar datos
- [ ] Usas `[ ]` para cambiar propiedades HTML
- [ ] Usas `( )` para escuchar eventos
- [ ] Usas `[( )]` para dos direcciones
- [ ] Tienes FormsModule importado
- [ ] Sin lógica compleja en templates
- [ ] Usas safe navigation `?.`
- [ ] Sin memory leaks (event listeners)
- [ ] Performance: cálculos en TypeScript, no HTML

---

## Próximo

Módulo 11: Directivas - Control el flujo del template

---

## Errores comunes

1. Aprender la sintaxis sin entender el propósito
2. Mezclar lógica de UI con lógica de negocio
3. No practicar con mini ejemplos propios

---

## Mini práctica

Haz esto por tu cuenta:

1. Repite el ejemplo sin copiar
2. Cámbiale nombres y estructura
3. Agrega una pequeña mejora
4. Explica en voz alta qué hace cada parte

---

## Resumen

Si puedes explicarlo con tus palabras, vas bien.
Si solo lo reconoces cuando lo ves, todavía falta práctica.
