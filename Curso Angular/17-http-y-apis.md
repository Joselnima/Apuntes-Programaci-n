# Módulo 17 - HTTP y Consumo de APIs

> **Conecta tu frontend con el backend**

---

## ¿QUÉ es una API HTTP?

**Comunicación entre aplicaciones**

Analogía:
- Restaurante: Haces un pedido (Request)
- Chef prepara (Backend procesa)
- Traen el plato (Response)

En Angular:

```
Angular (Frontend) 
   ↓ GET /api/products
Backend API
   ↓ 200 { products: [...] }
Angular recibe datos
```

---

## ¿PARA QUÉ?

### Problema: Datos Estáticos

**❌ Datos hardcoded:**

```typescript
export class ProductsComponent {
  products = [
    { id: 1, name: 'Laptop', price: 1000 },
    { id: 2, name: 'Mouse', price: 20 }
  ];
}
```

**Limitaciones:**
- No puedes agregar productos
- No persiste en la BD
- Solo lecturas
- No es real

### Solución: API Real

**✅ Datos desde backend:**

```typescript
export class ProductsComponent {
  products$ = this.productService.getProducts();
  
  constructor(private productService: ProductService) {}
}
```

**Con API:**
- Datos dinámicos
- Persistencia en BD
- CRUD operaciones
- App real ✨

---

## ¿CÓMO hacer peticiones HTTP?

### Paso 1: Importar HttpClientModule

**app.module.ts:**

```typescript
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  imports: [
    HttpClientModule  // ← Agregar
  ]
})
export class AppModule { }
```

### Paso 2: Crear Servicio

**product.service.ts:**

```typescript
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class ProductService {
  private apiUrl = 'http://localhost:3000/api/products';
  
  constructor(private http: HttpClient) {}
  
  // GET: Obtener todos
  getProducts() {
    return this.http.get<any[]>(this.apiUrl);
  }
  
  // GET: Obtener uno
  getProductById(id: number) {
    return this.http.get<any>(`${this.apiUrl}/${id}`);
  }
  
  // POST: Crear
  createProduct(product: any) {
    return this.http.post<any>(this.apiUrl, product);
  }
  
  // PUT: Actualizar
  updateProduct(id: number, product: any) {
    return this.http.put<any>(`${this.apiUrl}/${id}`, product);
  }
  
  // DELETE: Eliminar
  deleteProduct(id: number) {
    return this.http.delete<any>(`${this.apiUrl}/${id}`);
  }
}
```

### Paso 3: Usar en Componentes

**products.component.ts:**

```typescript
import { Component, OnInit } from '@angular/core';
import { ProductService } from '../../services/product.service';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html'
})
export class ProductsComponent implements OnInit {
  products: any[] = [];
  loading = true;
  error = '';
  
  constructor(private productService: ProductService) {}
  
  ngOnInit() {
    this.loadProducts();
  }
  
  loadProducts() {
    this.loading = true;
    this.productService.getProducts()
      .subscribe(
        (data) => {
          this.products = data;
          this.loading = false;
        },
        (error) => {
          this.error = 'Error cargando productos';
          this.loading = false;
          console.error(error);
        }
      );
  }
  
  addProduct(name: string, price: number) {
    this.productService.createProduct({ name, price })
      .subscribe(
        (data) => {
          this.products.push(data);
          console.log('Producto añadido');
        },
        (error) => {
          console.error('Error creando producto', error);
        }
      );
  }
  
  deleteProduct(id: number) {
    this.productService.deleteProduct(id)
      .subscribe(
        () => {
          this.products = this.products
            .filter(p => p.id !== id);
        },
        (error) => {
          console.error('Error eliminando', error);
        }
      );
  }
}
```

**products.component.html:**

```html
<div class="products">
  <!-- Loading -->
  <p *ngIf="loading">Cargando...</p>
  
  <!-- Error -->
  <p *ngIf="error" class="error">{{ error }}</p>
  
  <!-- Datos cargados -->
  <div *ngIf="!loading && !error">
    <table>
      <tr>
        <th>Nombre</th>
        <th>Precio</th>
        <th>Acciones</th>
      </tr>
      <tr *ngFor="let product of products">
        <td>{{ product.name }}</td>
        <td>${{ product.price }}</td>
        <td>
          <button (click)="deleteProduct(product.id)">
            Eliminar
          </button>
        </td>
      </tr>
    </table>
  </div>
</div>
```

---

## Los 4 Métodos HTTP

| Método | Qué hace | Ejemplo |
|--------|----------|---------|
| **GET** | Leer datos | `getProducts()` |
| **POST** | Crear | `createProduct(data)` |
| **PUT** | Actualizar | `updateProduct(id, data)` |
| **DELETE** | Eliminar | `deleteProduct(id)` |

---

## Estructura de una Petición

```
GET http://localhost:3000/api/products

↓

RESPONSE:
{
  "status": 200,
  "data": [
    { "id": 1, "name": "Laptop", "price": 1000 },
    { "id": 2, "name": "Mouse", "price": 20 }
  ]
}
```

---

## Manejo de Errores

**Siempre espera errores:**

```typescript
this.productService.getProducts()
  .subscribe(
    (success) => {
      // ✅ Éxito
      this.products = success;
    },
    (error) => {
      // ❌ Error
      console.log('Código:', error.status);
      console.log('Mensaje:', error.error);
      
      if (error.status === 404) {
        this.error = 'No encontrado';
      } else if (error.status === 401) {
        this.error = 'No autenticado';
      } else if (error.status === 500) {
        this.error = 'Error del servidor';
      }
    }
  );
```

---

## Interceptores (Middleware HTTP)

**Intercepta TODAS las peticiones (agregar token, etc)**

**auth.interceptor.ts:**

```typescript
import { Injectable } from '@angular/core';
import {
  HttpInterceptor,
  HttpRequest,
  HttpHandler,
  HttpEvent
} from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  constructor(private authService: AuthService) {}
  
  intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    // Obtener token
    const token = localStorage.getItem('token');
    
    // Agregar header si existe token
    if (token) {
      req = req.clone({
        setHeaders: {
          Authorization: `Bearer ${token}`
        }
      });
    }
    
    // Pasar petición modificada
    return next.handle(req);
  }
}
```

**app.module.ts:**

```typescript
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { AuthInterceptor } from './interceptors/auth.interceptor';

@NgModule({
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    }
  ]
})
export class AppModule { }
```

**Resultado:**

```
Petición automáticamente tiene:
GET /api/products
Headers: {
  Authorization: "Bearer token123..."
}
```

---

## Operadores útiles (RxJS)

### map: Transformar datos

```typescript
// Convertir response a lo que necesitas
product$ = this.http.get('/api/products')
  .pipe(
    map(response => response.data)  // Solo data, ignora metadata
  );
```

### tap: Hacer algo sin cambiar datos

```typescript
// Loguear pero mantener datos
products$ = this.http.get('/api/products')
  .pipe(
    tap(data => console.log('Productos:', data))
  );
```

### catchError: Manejar errores

```typescript
import { catchError } from 'rxjs/operators';
import { of } from 'rxjs';

products$ = this.http.get('/api/products')
  .pipe(
    catchError(error => {
      console.error('Error:', error);
      return of([]);  // Retorna array vacío
    })
  );
```

### switchMap: Peticiones encadenadas

```typescript
// Obtener usuario, luego sus órdenes
user$: Observable<User> = this.getUser();
orders$ = this.user$.pipe(
  switchMap(user => this.getOrdersByUserId(user.id))
);
```

---

## Memory Leaks HTTP

### ❌ Error: No desuscribirse

```typescript
ngOnInit() {
  this.productService.getProducts()
    .subscribe(data => {
      this.products = data;
    });
    // Sin unsubscribe ❌
}
```

### ✅ Solución 1: takeUntil

```typescript
ngOnInit() {
  this.productService.getProducts()
    .pipe(takeUntil(this.destroy$))
    .subscribe(data => {
      this.products = data;
    });
}

ngOnDestroy() {
  this.destroy$.next();
  this.destroy$.complete();
}
```

### ✅ Solución 2: Async Pipe (Recomendado)

```html
<!-- Se desuscribe automáticamente -->
<div *ngFor="let product of (products$ | async)">
  {{ product.name }}
</div>
```

```typescript
products$ = this.productService.getProducts();
```

---

## Testing HTTP

**mock.service.ts:**

```typescript
provideMockError = () =>
  throwError(() => new HttpErrorResponse({ 
    status: 404, 
    error: 'Not found' 
  }));

getMockProducts = () => of([
  { id: 1, name: 'Laptop' },
  { id: 2, name: 'Mouse' }
]);
```

---

## Checklist - HTTP Profesional

- [ ] HttpClientModule importado
- [ ] Peticiones en servicios, NO componentes
- [ ] Todos los métodos HTTP implementados
- [ ] Manejo de errores completo
- [ ] Intercepteur de autenticación
- [ ] Sin memory leaks (async pipe o takeUntil)
- [ ] Tipos TypeScript en responses
- [ ] Loading/error states en UI
- [ ] Validación de datos del backend

---

## Próximo

Módulo 18: RxJS - Streams reactivos

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
