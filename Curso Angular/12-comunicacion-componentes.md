# Módulo 12 - Comunicación Entre Componentes

> **Componentes hablan asignatura entre sí**

---

## ¿QÉ es la comunicación?

**Intercambiar información entre componentes**

Analogía:
- Casa 1 necesita saber si llueve
- Casa 2 sabe si llueve
- Mensajero comunica: "Llueve!"
- Casa 1 se prepara

En Angular:
- Padre → Hijo: @Input (datos)
- Hijo → Padre: @Output (eventos)
- Sin relación: Servicio (Subject/BehaviorSubject)

---

## ¿PARA QUÉ?

### Problema: Componentes Aislados

**❌ Componentes que no se comunican:**

```typescript
// ProductsComponent
export class ProductsComponent {
  products = [
    { id: 1, name: 'Laptop' },
    { id: 2, name: 'Mouse' }
  ];
}

// CartComponent (No sabe qué productos agregaron!)
export class CartComponent {
  cartItems = [];
  // ¿Cómo se entera de nuevos productos?
}
```

### Solución: Comunicación Estructurada

**✅ Componentes conectados:**

```
ProductsComponent              
    ↓ (líneas de datos)        
ProductCardComponent           
    ↓ (evento: addToCart)      
CartComponent
    ↓ (actualiza carrito)
```

---

## ¿CÓMO?

### Patrón 1: Padre → Hijo (@Input)

**Pasar datos del padre al hijo**

**Padre: products-list.component.ts**

```typescript
import { Component } from '@angular/core';

@Component({
  selector: 'app-products-list',
  template: `
    <app-product-card
      *ngFor="let product of products"
      [product]="product"
      [onSale]="true"
      [discount]="10"
    ></app-product-card>
  `
})
export class ProductsListComponent {
  products = [
    { id: 1, name: 'Laptop', price: 1000 },
    { id: 2, name: 'Mouse', price: 20 }
  ];
}
```

**Hijo: product-card.component.ts**

```typescript
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-product-card',
  template: `
    <div class="card">
      <h3>{{ product.name }}</h3>
      <p>${{ product.price }}</p>
      
      <p *ngIf="onSale">
        Descuento: {{ discount }}%
      </p>
    </div>
  `
})
export class ProductCardComponent {
  @Input() product: any;      // Recibe producto
  @Input() onSale: boolean;   // Recibe estado
  @Input() discount: number;  // Recibe descuento
}
```

**Cómo funciona:**

```
Padre:
products = [...]
[product]="products[0]"    ← Pasa datos

       ↓↓↓ FLUJO ↓↓↓

Hijo:
@Input() product           ← Recibe datos
{{ product.name }}         ← Los usa en template
```

#### @Input Avanzado

```typescript
// Nombre diferente
@Input('productData') product: any;

// Requerido (error si no pasas)
@Input() required product: any;

// Valor por defecto
@Input() discount: number = 0;

// Getter/Setter (reacciona a cambios)
private _quantity = 0;

@Input()
set quantity(value: number) {
  this._quantity = value;
  this.calculateTotal();
}

get quantity(): number {
  return this._quantity;
}
```

---

### Patrón 2: Hijo → Padre (@Output)

**Enviar evento del hijo al padre**

**Hijo: product-card.component.ts**

```typescript
import { Component, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-product-card',
  template: `
    <div class="card">
      <h3>{{ product.name }}</h3>
      <p>${{ product.price }}</p>
      
      <input type="number" [(ngModel)]="quantity" min="1">
      
      <button (click)="onAddClick()">
        Agregar al carrito
      </button>
    </div>
  `
})
export class ProductCardComponent {
  @Input() product: any;
  
  // Crea evento
  @Output() onAddToCart = new EventEmitter<any>();
  
  quantity = 1;
  
  // Emite evento
  onAddClick() {
    this.onAddToCart.emit({
      product: this.product,
      quantity: this.quantity
    });
  }
}
```

**Padre: products-list.component.ts**

```typescript
export class ProductsListComponent {
  products = [...];
  cartItems = [];
  
  // Escucha evento del hijo
  onProductAddToCart(event: any) {
    console.log('Producto agregado:', event);
    this.cartItems.push(event);
  }
}
```

**Padre: Template**

```html
<app-product-card
  *ngFor="let product of products"
  [product]="product"
  (onAddToCart)="onProductAddToCart($event)"
></app-product-card>

<p>Carrito: {{ cartItems.length }} items</p>
```

**Cómo funciona:**

```
Hijo:
@Output() onAddToCart = new EventEmitter()
onAddToCart.emit(data)     ← Emite evento

       ↓↓↓ FLUJO ↓↓↓

Padre:
(onAddToCart)="handler($event)"  ← Escucha
handler() recibe el data          ← Procesa
```

#### @Output Avanzado

```typescript
// Nombre diferente
@Output('addToCart') onAdd = new EventEmitter();

// Emitir múltiples veces
updateQuantity(qty: number) {
  this.onAdd.emit(qty);
}

// Emitir con typing
@Output() priceChange = new EventEmitter<{
  productId: number,
  newPrice: number
}>();

// Emitir
this.priceChange.emit({
  productId: 1,
  newPrice: 99.99
});
```

---

### Patrón 3: Componentes No Relacionados (Servicio + Subject)

**Dos componentes sin relación padre-hijo**

**Escenario:**

```
CartComponent              ProductsComponent
    ↓                      ↓
    └─── CartService ───┘
         (Subject)
```

**cart.service.ts:**

```typescript
import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class CartService {
  private cartSubject = new Subject<any[]>();
  cart$ = this.cartSubject.asObservable();
  
  private cart: any[] = [];
  
  addItem(item: any) {
    this.cart.push(item);
    this.cartSubject.next([...this.cart]);
  }
  
  removeItem(id: number) {
    this.cart = this.cart.filter(item => item.id !== id);
    this.cartSubject.next([...this.cart]);
  }
  
  getCart() {
    return this.cart;
  }
}
```

**products.component.ts:**

```typescript
export class ProductsComponent implements OnInit {
  products = [...];
  
  constructor(private cartService: CartService) {}
  
  ngOnInit() {
    // No necesita suscribirse (CartComponent lo hace)
  }
  
  addToCart(product: any) {
    this.cartService.addItem(product);
  }
}
```

**cart.component.ts:**

```typescript
export class CartComponent implements OnInit, OnDestroy {
  cartItems$ = this.cartService.cart$;
  destroy$ = new Subject<void>();
  
  constructor(private cartService: CartService) {}
  
  ngOnInit() {
    // No necesita suscribirse (async pipe)
  }
  
  removeItem(id: number) {
    this.cartService.removeItem(id);
  }
  
  ngOnDestroy() {
    this.destroy$.next();
  }
}
```

**cart.component.html:**

```html
<div>
  <h2>Carrito</h2>
  <ul>
    <li *ngFor="let item of (cartItems$ | async)">
      {{ item.name }} - ${{ item.price }}
      <button (click)="removeItem(item.id)">Eliminar</button>
    </li>
  </ul>
</div>
```

---

## 4 Patrones de Comunicación

| Caso | Patrón | Tool |
|------|--------|------|
| Padre → Hijo | Direct | @Input |
| Hijo → Padre | Event | @Output |
| Entre hermanos | Mediator | Servicio |
| Global (app) | Global State | Service + Subject |

---

## Ejemplo Completo: Carrito de Compras

**cart-item.component.ts (Hijo):**

```typescript
@Component({
  selector: 'app-cart-item',
  template: `
    <div class="item">
      <span>{{ item.name }}</span>
      <input 
        [(ngModel)]="quantity"
        (change)="onQuantityChange()"
        type="number"
      >
      <button (click)="onRemove()">Borrar</button>
    </div>
  `
})
export class CartItemComponent {
  @Input() item: any;
  @Output() onQuantityChanged = new EventEmitter<{id, qty}>();
  @Output() onRemoved = new EventEmitter<number>();
  
  quantity = 1;
  
  onQuantityChange() {
    this.onQuantityChanged.emit({
      id: this.item.id,
      qty: this.quantity
    });
  }
  
  onRemove() {
    this.onRemoved.emit(this.item.id);
  }
}
```

**cart.component.ts (Padre):**

```typescript
@Component({
  selector: 'app-cart',
  template: `
    <ul>
      <app-cart-item
        *ngFor="let item of cartItems"
        [item]="item"
        (onQuantityChanged)="updateQuantity($event)"
        (onRemoved)="removeItem($event)"
      ></app-cart-item>
    </ul>
  `
})
export class CartComponent {
  cartItems = [
    { id: 1, name: 'Laptop' },
    { id: 2, name: 'Mouse' }
  ];
  
  updateQuantity(event: any) {
    console.log('Cantidad actualizada:', event);
  }
  
  removeItem(id: number) {
    this.cartItems = this.cartItems.filter(item => item.id !== id);
  }
}
```

---

## Memory Leaks en Comunicación

### ❌ Error: Suscripciones sin limpiar

```typescript
ngOnInit() {
  this.cartService.cart$.subscribe(items => {
    this.items = items;
  });
  // Sin unsubscribe ❌
}
```

### ✅ Solución: takeUntil

```typescript
destroy$ = new Subject<void>();

ngOnInit() {
  this.cartService.cart$
    .pipe(takeUntil(this.destroy$))
    .subscribe(items => {
      this.items = items;
    });
}

ngOnDestroy() {
  this.destroy$.next();
}
```

### ✅ Mejor: async pipe

```html
<!-- Desuscrita automáticamente -->
<div *ngFor="let item of (cartItems$ | async)">
  {{ item.name }}
</div>
```

---

## Checklist - Comunicación Profesional

- [ ] @Input para datos padre → hijo
- [ ] @Output para eventos hijo → padre  
- [ ] Servicios para componentes no relacionados
- [ ] BehaviorSubject para estado compartido
- [ ] Sin memory leaks (unsubscribe o async)
- [ ] Tipado correcto (TypeScript)
- [ ] Nombres descriptivos (@Input/Output)
- [ ] Documentación clara
- [ ] Sin acoplamientos innecesarios

---

## Próximo

Módulo 13: Routing - Navegación entre páginas

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
