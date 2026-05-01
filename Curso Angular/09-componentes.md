# Módulo 09 - Componentes: El corazón de Angular

> **Las piezas reutilizables que hacen posible una app real**

---

## ¿Qué es un componente?

Un componente en Angular es una unidad que combina:

- HTML (vista)
- CSS (estilos)
- TypeScript (lógica)

Piensa en él como una caja con nombre:

- tiene entrada (`@Input()`)
- tiene salida (`@Output()`)
- tiene comportamiento propio

### Ejemplo mínimo

```typescript
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  title = 'Mi Tienda';

  handleMenu() {
    console.log('Abrir menú');
  }
}
```

---

## ¿Por qué usar componentes?

### Problema 1: páginas enormes

Si una vista tiene todo el HTML mezclado con lógica, se vuelve difícil de mantener y entender.

### Problema 2: código duplicado

Si repites el mismo bloque en varias páginas, cada cambio se vuelve una pesadilla.

### Problema 3: responsabilidad mezclada

Un componente debe hacer una sola cosa:
- una tarjeta de producto
- un botón de búsqueda
- una lista de mensajes

Si mezclas todo, el código se rompe más rápido.

---

## Módulo 09 explicado con casos reales

### Caso real 1: tienda online

Componentes típicos:

- `HeaderComponent`
- `SearchBarComponent`
- `ProductCardComponent`
- `ProductListComponent`
- `CartComponent`
- `CheckoutComponent`

Flujo real:

1. `ProductListComponent` solicita los productos al servicio.
2. Cada `ProductCardComponent` recibe su producto con `@Input()`.
3. El usuario hace click en "Agregar al carrito".
4. `ProductCardComponent` emite el evento con `@Output()`.
5. `CartComponent` actualiza el carrito desde un servicio compartido.

Eso significa:

- cada componente es independiente
- el diseño es reutilizable
- el mantenimiento es más rápido

### Caso real 2: dashboard de estadísticas

Componentes comunes:

- `StatCardComponent` para cada tarjeta de métrica
- `ChartComponent` para gráficos
- `TableComponent` para tablas
- `FiltersComponent` para filtros globales

Un `StatCardComponent` puede ser así:

```typescript
@Component({
  selector: 'app-stat-card',
  template: `
    <div class="stat-card">
      <h4>{{ title }}</h4>
      <p>{{ value }}</p>
    </div>
  `
})
export class StatCardComponent {
  @Input() title!: string;
  @Input() value!: number;
}
```

Esto permite usar la misma tarjeta con datos diferentes sin repetir código.

### Caso real 3: formulario por secciones

En una app con formularios largos, conviene dividir en componentes pequeños:

- `PersonalInfoComponent`
- `AddressFormComponent`
- `PaymentFormComponent`

Con esto tu `CheckoutComponent` queda más limpio:

```html
<app-personal-info [formGroup]="personalForm"></app-personal-info>
<app-address-form [formGroup]="addressForm"></app-address-form>
<app-payment-form [formGroup]="paymentForm"></app-payment-form>
```

---

## Cómo crear un componente

### Paso 1: generar con Angular CLI

```bash
ng generate component components/product-card
```

Angular crea:

```
src/app/components/product-card/
├── product-card.component.ts
├── product-card.component.html
├── product-card.component.css
└── product-card.component.spec.ts
```

### Paso 2: lógica en TypeScript

```typescript
import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-product-card',
  templateUrl: './product-card.component.html',
  styleUrls: ['./product-card.component.css']
})
export class ProductCardComponent {
  @Input() product: any;
  @Output() addToCart = new EventEmitter<any>();

  quantity = 1;

  handleAddToCart() {
    this.addToCart.emit({
      product: this.product,
      quantity: this.quantity
    });
  }
}
```

### Paso 3: template HTML

```html
<div class="card">
  <img [src]="product.imageUrl" alt="{{ product.name }}">
  <h3>{{ product.name }}</h3>
  <p>{{ product.description }}</p>
  <p class="price">${{ product.price }}</p>

  <div class="quantity">
    <input type="number" [(ngModel)]="quantity" min="1">
  </div>

  <button (click)="handleAddToCart()">Agregar al carrito</button>
</div>
```

### Paso 4: estilos CSS

```css
.card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  max-width: 300px;
}

.card img {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 4px;
}

.price {
  font-size: 20px;
  font-weight: bold;
  color: #28a745;
}

button {
  padding: 10px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background: #0056b3;
}
```

### Paso 5: usar el componente desde un padre

```html
<div class="products-grid">
  <app-product-card
    *ngFor="let product of products"
    [product]="product"
    (addToCart)="handleAddToCart($event)"
  ></app-product-card>
</div>
```

```typescript
import { Component, OnInit } from '@angular/core';
import { ProductService } from '../../services/product.service';

@Component({
  selector: 'app-products-list',
  templateUrl: './products-list.component.html',
  styleUrls: ['./products-list.component.css']
})
export class ProductsListComponent implements OnInit {
  products: any[] = [];

  constructor(private productService: ProductService) {}

  ngOnInit() {
    this.productService.getProducts().subscribe(data => {
      this.products = data;
    });
  }

  handleAddToCart(event: any) {
    console.log('Agregar:', event.product.name, 'Cantidad:', event.quantity);
    // Actualiza el carrito con un servicio real
  }
}
```

---

## Ciclo de vida de un componente Angular

Un componente tiene etapas claras:

- `constructor()` → se crea el componente
- `ngOnChanges()` → cambia una entrada (`@Input`)
- `ngOnInit()` → inicializa datos
- `ngDoCheck()` → chequeos personalizados
- `ngAfterViewInit()` → la vista ya está renderizada
- `ngOnDestroy()` → se destruye y limpia recursos

### Ejemplo real con `ngOnChanges`

```typescript
import { Component, Input, OnChanges, SimpleChanges, OnInit, OnDestroy } from '@angular/core';

@Component({
  selector: 'app-user-card',
  template: `<p>{{ name }}</p>`
})
export class UserCardComponent implements OnChanges, OnInit, OnDestroy {
  @Input() name!: string;

  ngOnChanges(changes: SimpleChanges) {
    if (changes.name) {
      console.log('Nombre cambió:', changes.name.currentValue);
    }
  }

  ngOnInit() {
    console.log('UserCard inicializado');
  }

  ngOnDestroy() {
    console.log('UserCard destruido');
  }
}
```

### Ejemplo real con limpieza en `ngOnDestroy`

```typescript
import { Component, OnInit, OnDestroy } from '@angular/core';
import { interval, Subscription } from 'rxjs';

@Component({
  selector: 'app-clock',
  template: `<p>Hora: {{ time }}</p>`
})
export class ClockComponent implements OnInit, OnDestroy {
  time = new Date().toLocaleTimeString();
  private subscription!: Subscription;

  ngOnInit() {
    this.subscription = interval(1000).subscribe(() => {
      this.time = new Date().toLocaleTimeString();
    });
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }
}
```

---

## Evitar memory leaks en Angular

### Mala práctica: suscripción sin limpiar

```typescript
export class BadComponent implements OnInit {
  ngOnInit() {
    this.service.data$.subscribe(data => {
      this.data = data;
    });
  }
}
```

### Buena práctica real

```typescript
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';

export class GoodComponent implements OnInit, OnDestroy {
  private destroy$ = new Subject<void>();

  ngOnInit() {
    this.service.data$
      .pipe(takeUntil(this.destroy$))
      .subscribe(data => {
        this.data = data;
      });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
```

### Limpiar event listeners

```typescript
scrollHandler = () => console.log('Scrolling');

ngOnInit() {
  window.addEventListener('scroll', this.scrollHandler);
}

ngOnDestroy() {
  window.removeEventListener('scroll', this.scrollHandler);
}
```

### Limpiar timers

```typescript
timerId: any;

ngOnInit() {
  this.timerId = setInterval(() => {
    this.counter++;
  }, 1000);
}

ngOnDestroy() {
  clearInterval(this.timerId);
}
```

---

## Comunicación entre componentes

### Padre → Hijo con `@Input`

```html
<app-child [message]="'Hola'"></app-child>
```

```typescript
@Input() message: string;
```

### Hijo → Padre con `@Output`

```typescript
@Output() onClick = new EventEmitter<string>();

handleClick() {
  this.onClick.emit('Clickeaste');
}
```

```html
<app-child (onClick)="handleChildClick($event)"></app-child>
```

### Componentes no relacionados con servicio

```typescript
@Injectable({ providedIn: 'root' })
export class MessageService {
  private messageSubject = new Subject<string>();
  message$ = this.messageSubject.asObservable();

  sendMessage(msg: string) {
    this.messageSubject.next(msg);
  }
}
```

```typescript
constructor(private messageService: MessageService) {}

sendMessage() {
  this.messageService.sendMessage('Hola!');
}
```

```typescript
constructor(private messageService: MessageService) {}

ngOnInit() {
  this.messageService.message$.subscribe(msg => {
    console.log(msg);
  });
}
```

---

## Optimización real con Change Detection

Angular por defecto revisa todo cuando cambia el estado.

### OnPush para rendimiento

```typescript
import { ChangeDetectionStrategy, Component, Input } from '@angular/core';

@Component({
  selector: 'app-card',
  template: `{{ product.name }}`,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class CardComponent {
  @Input() product: any;
}
```

**Caso real:** en una lista de productos, `CardComponent` solo se verifica cuando cambian las `@Input()`.

### `trackBy` para listas grandes

```html
<div *ngFor="let item of items; trackBy: trackById">
  <app-item [item]="item"></app-item>
</div>
```

```typescript
trackById(index: number, item: any) {
  return item.id;
}
```

Esto evita renderizados innecesarios cuando solo cambian algunos elementos.

---

## Cuándo crear un componente

- cuando el bloque de UI se repite
- cuando la lógica es compleja
- cuando quieres separar responsabilidades
- cuando el elemento puede reutilizarse en otra página

## Buenas prácticas reales

- componentes pequeños y claros
- nombres descriptivos
- usa `@Input()` y `@Output()` para comunicar
- limpia subscripciones en `ngOnDestroy()`
- usa `OnPush` para componentes con datos inmutables

---

## Resumen

Los componentes son el corazón de Angular.

Con buenos componentes:

- la app es más fácil de entender
- el código se mantiene mejor
- los cambios son más seguros
- las pruebas son más sencillas
