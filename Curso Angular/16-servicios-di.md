# Módulo 16 - Servicios e Inyección de Dependencias

> **Lógica compartida y reutilizable**

---

## ¿QUÉ es un servicio?

**Una caja de herramientas compartida**

Analogía:
- Casa: Necesitas acceso a agua, luz, gas
- No lo creas cada vez, lo compartes
- Un servicio atiende a muchas casas

En Angular:
- `CartService`: Maneja carrito (muchos componentes lo usan)
- `ProductService`: Obtiene productos
- `AuthService`: Maneja autenticación

**Estructura mínima:**

```typescript
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'  // Disponible en toda la app
})
export class CartService {
  private items: any[] = [];
  
  addItem(item: any) {
    this.items.push(item);
  }
  
  getItems() {
    return this.items;
  }
}
```

---

## ¿PARA QUÉ?

### Problema 1: Duplicación de Código

**❌ Sin servicios:**

```typescript
// ProductsComponent
export class ProductsComponent {
  ngOnInit() {
    // Aquí genero la petición HTTP
    fetch('http://api.com/products')
      .then(r => r.json())
      .then(data => this.products = data);
  }
}

// DetailComponent
export class DetailComponent {
  ngOnInit() {
    // Aquí repito lo mismo
    fetch('http://api.com/products')
      .then(r => r.json())
      .then(data => this.products = data);
  }
}

// HeaderComponent
export class HeaderComponent {
  ngOnInit() {
    // Y aquí repito OTRA VEZ
    fetch('http://api.com/products')
      .then(r => r.json())
      .then(data => this.products = data);
  }
}
```

**❌ Cambio en API?**
- Necesitas actualizar en 3 lugares (¡error!)
- Alguien olvida actualizar (¡error!)

### Solución: Centralizar en Servicio

**✅ Con servicio:**

```typescript
// product.service.ts
@Injectable({ providedIn: 'root' })
export class ProductService {
  constructor(private http: HttpClient) {}
  
  getProducts() {
    return this.http.get('http://api.com/products');
  }
}

// ProductsComponent
export class ProductsComponent {
  constructor(private productService: ProductService) {}
  
  ngOnInit() {
    this.productService.getProducts()
      .subscribe(data => this.products = data);
  }
}

// DetailComponent - Misma lógica, mismo servicio
export class DetailComponent {
  constructor(private productService: ProductService) {}
  
  ngOnInit() {
    this.productService.getProducts()
      .subscribe(data => this.products = data);
  }
}

// HeaderComponent - También usa el mismo servicio
export class HeaderComponent {
  constructor(private productService: ProductService) {}
  
  ngOnInit() {
    this.productService.getProducts()
      .subscribe(data => this.products = data);
  }
}
```

**✅ Cambio en API?**
- Actualizas en UN lugar: `productService.ts`
- Todos los componentes se actualizan automáticamente ✨

### Problema 2: Estado Compartido

**❌ Sin servicio:**

```typescript
// App tiene estado
export class AppComponent {
  cartCount = 0;
  
  addToCart() {
    this.cartCount++;
  }
}

// Header necesita cartCount
// Product necesita cartCount
// ¿Cómo se lo pasas? ¿Props? ¿Global? ¿Caos?
```

**✅ Con servicio:**

```typescript
// cart.service.ts
@Injectable({ providedIn: 'root' })
export class CartService {
  private countSubject = new BehaviorSubject(0);
  count$ = this.countSubject.asObservable();
  
  addItem() {
    this.countSubject.next(this.countSubject.value + 1);
  }
}

// Cualquier componente accede al carrito
export class HeaderComponent {
  constructor(private cartService: CartService) {}
  
  cartCount$ = this.cartService.count$;
}
```

---

## ¿CÓMO crear servicios?

### Paso 1: Generar el Servicio

```bash
ng generate service services/cart
```

**Resultado:**

```
services/
├── cart.service.ts
└── cart.service.spec.ts  (ignorar)
```

### Paso 2: Escribir la Lógica

**cart.service.ts:**

```typescript
import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'  // ✅ Disponible globalmente
})
export class CartService {
  private cart: any[] = [];
  private cartSubject = new BehaviorSubject<any[]>([]);
  
  cart$ = this.cartSubject.asObservable();
  
  // Constructor
  constructor() {
    console.log('CartService inicializado');
  }
  
  // Métodos
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
  
  getTotalPrice() {
    return this.cart.reduce((sum, item) => 
      sum + (item.price * item.quantity), 0
    );
  }
  
  clearCart() {
    this.cart = [];
    this.cartSubject.next([]);
  }
}
```

### Paso 3: Usar en Componentes

**products.component.ts:**

```typescript
import { Component, OnInit } from '@angular/core';
import { CartService } from '../../services/cart.service';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html'
})
export class ProductsComponent implements OnInit {
  products = [
    { id: 1, name: 'Laptop', price: 1000 },
    { id: 2, name: 'Mouse', price: 20 },
    { id: 3, name: 'Teclado', price: 50 }
  ];
  
  // ✅ Inyectar servicio en constructor
  constructor(private cartService: CartService) {}
  
  ngOnInit() {
    console.log('Componente inicializado');
  }
  
  addToCart(product: any) {
    this.cartService.addItem({
      ...product,
      quantity: 1
    });
  }
}
```

**products.component.html:**

```html
<div class="products">
  <div *ngFor="let product of products" class="product-card">
    <h3>{{ product.name }}</h3>
    <p>${{ product.price }}</p>
    <button (click)="addToCart(product)">
      Agregar al carrito
    </button>
  </div>
</div>
```

---

## Inyección de Dependencias (DI)

**Angular entrega automáticamente las dependencias**

```typescript
// Constructor
constructor(private cartService: CartService) {
  // Angular automáticamente inyecta CartService
  // Disponible como this.cartService
}
```

**Flujo:**

```
1. Declara dependencia en constructor
   constructor(private cartService: CartService)
   
2. Angular busca CartService
   ¿Lo conoce? ¿Está registrado?
   
3. Crea instancia (Si no existe)
   O reutiliza si ya existe
   
4. Inyecta en el componente
   this.cartService está listo
```

---

## Niveles de Inyección

### Nivel 1: Singleton (@Injectable root)

**UN servicio para toda la app**

```typescript
@Injectable({
  providedIn: 'root'  // Singleton global
})
export class CartService { }
```

**Todos los componentes comparten la misma instancia:**

```
HeaderComponent  →
                  CartService (UNA SOLA INSTANCIA)
ProductComponent →
```

### Nivel 2: Módulo Específico

```typescript
@Injectable({
  providedIn: ProductsModule  // Solo en ProductModule
})
export class ProductService { }
```

### Nivel 3: Componente Específico

```typescript
@Component({
  selector: 'app-product',
  providers: [ProductService]  // Nueva instancia por componente
})
export class ProductComponent { }
```

**En este caso cada componente tiene su propia instancia:**

```
ProductComponent A → ProductService instancia #1
ProductComponent B → ProductService instancia #2
```

---

## Servicios con Dependencias

**Un servicio puede usar otros servicios:**

```typescript
@Injectable({ providedIn: 'root' })
export class OrderService {
  constructor(
    private cartService: CartService,
    private http: HttpClient,
    private authService: AuthService
  ) {}
  
  placeOrder() {
    const cartItems = this.cartService.getCart();
    const user = this.authService.getUser();
    
    return this.http.post('/api/orders', {
      items: cartItems,
      user: user
    });
  }
}
```

**Angular resuelve la cadena de dependencias automáticamente:**

```
OrderService necesita:
  → CartService
  → HttpClient
  → AuthService
```

---

## Memory Leaks en Servicios

### ❌ Error: Suscripciones Infinitas

```typescript
// ❌ INCORRECTO
export class BadComponent implements OnInit {
  ngOnInit() {
    this.dataService.data$.subscribe(data => {
      this.data = data;
    });
    // Sin unsubscribe ❌
  }
}
```

**✅ CORRECCIÓN:**

```typescript
export class GoodComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  
  constructor(private dataService: DataService) {}
  
  ngOnInit() {
    this.subscription = this.dataService.data$
      .subscribe(data => {
        this.data = data;
      });
  }
  
  ngOnDestroy() {
    this.subscription.unsubscribe();
  }
}
```

**O mejor aún, con async pipe:**

```html
<!-- Async pipe se desuscribe automáticamente -->
<div>{{ dataService.data$ | async }}</div>
```

---

## Ejemplos Útiles

### Servicio de Autenticación

```typescript
@Injectable({ providedIn: 'root' })
export class AuthService {
  private userSubject = new BehaviorSubject<any>(null);
  user$ = this.userSubject.asObservable();
  
  login(email: string, password: string) {
    return this.http.post('/api/login', { email, password })
      .pipe(
        tap(response => {
          this.userSubject.next(response.user);
          localStorage.setItem('token', response.token);
        })
      );
  }
  
  logout() {
    this.userSubject.next(null);
    localStorage.removeItem('token');
  }
  
  isLoggedIn() {
    return this.user$.pipe(
      map(user => !!user)
    );
  }
}
```

### Servicio de Notificaciones

```typescript
@Injectable({ providedIn: 'root' })
export class NotificationService {
  private notificationSubject = new Subject<{
    message: string,
    type: 'success' | 'error' | 'warning'
  }>();
  
  notification$ = this.notificationSubject.asObservable();
  
  success(message: string) {
    this.notificationSubject.next({ message, type: 'success' });
  }
  
  error(message: string) {
    this.notificationSubject.next({ message, type: 'error' });
  }
  
  warning(message: string) {
    this.notificationSubject.next({ message, type: 'warning' });
  }
}
```

**Uso en componentes:**

```typescript
export class MyComponent {
  notifications$ = this.notifiService.notification$;
  
  constructor(private notifiService: NotificationService) {}
  
  doSomething() {
    this.notifiService.success('¡Guardado!');
  }
}
```

---

## Checklist - Servicios Profacionales

- [ ] Servicio generado con `ng generate service`
- [ ] @Injectable({ providedIn: 'root' })
- [ ] Métodos con responsabilidad clara
- [ ] Usa Observables para datos reactivos
- [ ] Lógica NO en componentes, EN servicios
- [ ] Sin memory leaks (unsubscribe en ngOnDestroy)
- [ ] Servicios reutilizables (no duplicar código)
- [ ] Tipos TypeScript en métodos
- [ ] Documentación clara

---

## Próximo

Módulo 17: HTTP y APIs - Conecta tu Backend

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
