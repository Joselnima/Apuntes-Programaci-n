# Módulo 09 - Componentes: El Corazón de Angular

> **Las piezas reutilizables de toda aplicación**

---

## ¿QUÉ es un componente?

**Una caja que contiene: HTML + CSS + Lógica**

Analogía:
- LEGO: Cada pieza es independiente
- Puedes combinarlas de infinitas formas
- Una pieza rota no afecta otras

En Angular:
- HeaderComponent
- ProductCardComponent
- CartComponent
- FooterComponent

**Cada uno vive su vida.**

```typescript
// Estructura mínima
@Component({
  selector: 'app-header',           // Nombre HTML
  templateUrl: './header.html',     // HTML
  styleUrls: ['./header.css']       // CSS
})
export class HeaderComponent {
  title = 'Mi Tienda';              // Datos
  
  handleClick() {
    // Lógica
  }
}
```

---

## ¿PARA QUÉ?

### Problema 1: Una Página, 500 Líneas

```html
<!-- app.component.html - SIN componentes -->
<header>
  <h1>Mi Tienda</h1>
  <nav>...</nav>
</header>

<main>
  <div class="product-list">
    <div class="product-card">...</div>
    <div class="product-card">...</div>
    <!-- 50 líneas de HTML -->
  </div>
</main>

<footer>
  <!-- 20 líneas más -->
</footer>
```

**❌ Imposible de mantener. ¿Dónde termina header?**

### Solución: Descomponer en Componentes

```html
<!-- app.component.html - CON componentes -->
<app-header></app-header>

<main>
  <app-product-list></app-product-list>
</main>

<app-footer></app-footer>
```

**✅ Claro. Cada componente en su archivo.**

### Problema 2: Código Duplicado

```typescript
// Sin componentes: Header aparece en 3 páginas
// Copias ese HTML/CSS 3 veces
// Cambio: actualizar en 3 lugares ❌

// Con componentes: HeaderComponent existe una vez
// Las 3 páginas lo usan
// Cambio: actualizar en 1 lugar ✅
```

### Problema 3: Propiedades Locales Compartidas

```typescript
// Sin servicios, cómo pasan datos?

// ❌ Aquí: cartCount en AppComponent
// ProductCard necesita cartCount
// Frontend necesita cartCount
// ¿Quién es la fuente de verdad?

// ✅ CartService es la fuente de verdad
// Todos los componentes leen de CartService
```

---

## ¿CÓMO crear componentes?

### Paso 1: Generar el Componente

```bash
ng generate component components/product-card
```

**Angular crea:**

```
src/app/components/product-card/
├── product-card.component.ts       ← Lógica
├── product-card.component.html     ← Template
├── product-card.component.css      ← Estilos
└── product-card.component.spec.ts  ← Tests (ignorar)
```

### Paso 2: Definir la Lógica

**product-card.component.ts:**

```typescript
import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-product-card',
  templateUrl: './product-card.component.html',
  styleUrls: ['./product-card.component.css']
})
export class ProductCardComponent {
  // @Input: Datos QUE ENTRA (props)
  @Input() product: any;
  
  // @Output: Eventos QUE SALE
  @Output() onAddToCart = new EventEmitter<any>();
  
  // Propiedades locales
  quantity = 1;
  
  // Métodos
  addToCart() {
    this.onAddToCart.emit({
      product: this.product,
      quantity: this.quantity
    });
  }
}
```

**¿Qué hace?**

| Concepto | Qué es |
|----------|--------|
| `@Input()` | Propiedad que recibe datos del padre |
| `@Output()` | Evento que envía datos al padre |
| `EventEmitter` | Objeto que dispara eventos |
| `emit()` | Envía un evento |

### Paso 3: Template HTML

**product-card.component.html:**

```html
<div class="card">
  <img [src]="product.imageUrl" alt="{{ product.name }}">
  
  <h3>{{ product.name }}</h3>
  <p>{{ product.description }}</p>
  
  <p class="price">
    ${{ product.price }}
  </p>
  
  <div class="quantity">
    <input 
      type="number" 
      [(ngModel)]="quantity" 
      min="1"
    >
  </div>
  
  <button (click)="addToCart()">
    Agregar al carrito
  </button>
</div>
```

### Paso 4: CSS

**product-card.component.css:**

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

### Paso 5: Usar el Componente

**products-list.component.html:**

```html
<div class="products-grid">
  <app-product-card
    *ngFor="let product of products"
    [product]="product"
    (onAddToCart)="handleAddToCart($event)"
  ></app-product-card>
</div>
```

**products-list.component.ts:**

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
    // Aquí actualizas el carrito
  }
}
```

---

## Ciclo de Vida de un Componente

**Los componentes nacen, viven, mueren.**

```
constructor()          ← Nace (NO accedes a propiedades)
  ↓
ngOnInit()            ← Inicializa (AQUÍ cargas datos)
  ↓
ngDoCheck()           ← Angular verifica cambios
  ↓
ngAfterViewInit()     ← Template ya renderizado
  ↓
ngOnDestroy()         ← Muere (LIMPIA aquí)
```

**Uso práctico:**

```typescript
import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-example',
  templateUrl: './example.html'
})
export class ExampleComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  
  constructor(private productService: ProductService) {}
  
  ngOnInit() {
    // ✅ CARGAR DATOS AQUÍ
    this.subscription = this.productService
      .getProducts()
      .subscribe(data => {
        console.log('Productos:', data);
      });
  }
  
  ngOnDestroy() {
    // ✅ LIMPIAR RECURSOS AQUÍ
    this.subscription.unsubscribe();
  }
}
```

---

## Memory Leaks (Gotchas Importantes)

### ❌ Memory Leak 1: Subscripciones Sin Limpiar

```typescript
// ❌ INCORRECTO
export class BadComponent implements OnInit {
  ngOnInit() {
    // Cada vez que entra al componente, nueva suscripción
    this.service.data$.subscribe(data => {
      this.data = data;
    });
    // Sin unsubscribe en ngOnDestroy
    // La memoria crece infinitamente
  }
}
```

**✅ CORRECCIÓN:**

```typescript
export class GoodComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  
  ngOnInit() {
    this.subscription = this.service.data$.subscribe(data => {
      this.data = data;
    });
  }
  
  ngOnDestroy() {
    this.subscription.unsubscribe(); // ✅ Limpiar
  }
}
```

### ❌ Memory Leak 2: Event Listeners

```typescript
// ❌ INCORRECTO
ngOnInit() {
  window.addEventListener('scroll', () => {
    console.log('Scrolling');
  });
  // ¿Quién lo remueve?
}
```

**✅ CORRECCIÓN:**

```typescript
scrollHandler = () => console.log('Scrolling');

ngOnInit() {
  window.addEventListener('scroll', this.scrollHandler);
}

ngOnDestroy() {
  window.removeEventListener('scroll', this.scrollHandler);
}
```

### ❌ Memory Leak 3: Timers

```typescript
// ❌ INCORRECTO
ngOnInit() {
  setInterval(() => {
    this.counter++;
  }, 1000);
  // El interval corre eternamente
}
```

**✅ CORRECCIÓN:**

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

## Comunicación Entre Componentes

### Patrón 1: Padre → Hijo (@Input)

```typescript
// Padre
<app-child [message]="'Hola'"></app-child>

// Hijo recibe
@Input() message: string;
template: {{ message }}  // "Hola"
```

### Patrón 2: Hijo → Padre (@Output)

```typescript
// Hijo emite
@Output() onClick = new EventEmitter<string>();

handleClick() {
  this.onClick.emit('Clickeaste');
}

// Padre escucha
<app-child (onClick)="handleChildClick($event)"></app-child>

handleChildClick(message: string) {
  console.log(message); // "Clickeaste"
}
```

### Patrón 3: Components NO Relacionados (Servicio)

```typescript
// servicio.ts
@Injectable({ providedIn: 'root' })
export class MessageService {
  private messageSubject = new Subject<string>();
  message$ = this.messageSubject.asObservable();
  
  sendMessage(msg: string) {
    this.messageSubject.next(msg);
  }
}

// Componente A emite
constructor(private messageService: MessageService) {}
sendMessage() {
  this.messageService.sendMessage('Hola!');
}

// Componente B recibe
constructor(private messageService: MessageService) {}
ngOnInit() {
  this.messageService.message$.subscribe(msg => {
    console.log(msg); // "Hola!"
  });
}
```

---

## Change Detection (Optimización)

**Por defecto, Angular revisa TODO cada cambio:**

```typescript
// ❌ LENTO
@Component({
  selector: 'app-card',
  template: `{{ product.name }}`,
  changeDetection: ChangeDetectionStrategy.Default
})
export class CardComponent {
  @Input() product: any;
}
```

**Con OnPush, solo revisa si @Input cambia:**

```typescript
// ✅ RÁPIDO
import { ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'app-card',
  template: `{{ product.name }}`,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class CardComponent {
  @Input() product: any;
}
```

**Usa OnPush siempre que puedas.** (Lo enseñaremos mejor en Optimización)

---

## Checklist - Componentes Profesionales

- [ ] Componente generado con `ng generate`
- [ ] Tiene `@Input()` para datos del padre
- [ ] Tiene `@Output()` para eventos
- [ ] Implementa `OnInit` para datos
- [ ] Implementa `OnDestroy` para limpieza
- [ ] Sin memory leaks (sin suscripciones olvidadas)
- [ ] CSS aislado (solo aplica al componente)
- [ ] Una responsabilidad clara
- [ ] Nombres descriptivos
- [ ] Menos de 300 líneas de código

---

## Próximo

Módulo 10: Templates y Data Binding - Conecta datos con HTML

---

## Resumen

Si puedes explicarlo con tus palabras, vas bien.
Si solo lo reconoces cuando lo ves, todavía falta práctica.
