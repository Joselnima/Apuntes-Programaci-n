# Módulo 13 - Routing: Navegación Entre Páginas

> **Múltiples páginas sin recargar**

---

## ¿QUÉ es Routing?

**Sistema de navegación entre vistas sin reload**

Analogía:
- Librería: Puedes ir del capítulo 1 al 3 sin recargar el libro
- SPA (Single Page App): Cambias vistas sin recargar HTML

```
URL: /productos         → ProductsComponent
URL: /productos/1       → ProductDetailComponent
URL: /carrito           → CartComponent
URL: /checkout          → CheckoutComponent

(Sin recargar página completa ✨)
```

---

## ¿PARA QUÉ?

### Problema: Navegación Rígida

```
❌ Sin router:
- Click → Recarga página (lento)
- Cookie se pierde
- Estado se reinicia

✅ Con router:
- Click → Cambio de vista (rápido  
- Cookie se mantiene
- Estado persiste
```

---

## ¿CÓMO crear rutas?

### Paso 1: Configurar Routes

**app-routing.module.ts:**

```typescript
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { ProductsComponent } from './pages/products/products.component';
import { ProductDetailComponent } from './pages/product-detail/product-detail.component';
import { CartComponent } from './pages/cart/cart.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'productos', component: ProductsComponent },
  { path: 'productos/:id', component: ProductDetailComponent },
  { path: 'carrito', component: CartComponent },
  { path: '404', component: NotFoundComponent },
  { path: '**', redirectTo: '/404' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
```

### Paso 2: Usar <router-outlet>

**app.component.html:**

```html
<app-header></app-header>

<router-outlet></router-outlet>

<app-footer></app-footer>
```

### Paso 3: Navegar

**header.component.html:**

```html
<nav>
  <a routerLink="/">Home</a>
  <a routerLink="/productos">Productos</a>
  <a routerLink="/carrito">Carrito</a>
</nav>
```

**component.ts (Navegación programática):**

```typescript
import { Router } from '@angular/router';

export class ProductsComponent {
  constructor(private router: Router) {}
  
  goToDetail(id: number) {
    this.router.navigate(['/productos', id]);
  }
  
  goHome() {
    this.router.navigate(['/']);
  }
}
```

---

## Rutas con Parámetros

### Parámetro en URL

```typescript
{ path: 'productos/:id', component: ProductDetailComponent }
```

**Capturar en componente:**

```typescript
import { ActivatedRoute } from '@angular/router';

export class ProductDetailComponent implements OnInit {
  productId: number;
  
  constructor(private route: ActivatedRoute) {}
  
  ngOnInit() {
    this.route.params.subscribe(params => {
      this.productId = params['id'];
      this.loadProduct(this.productId);
    });
  }
  
  loadProduct(id: number) {
    // Cargar producto
  }
}
```

### Query Parameters

```typescript
// URL: /productos?category=electronics&sort=price

this.route.queryParams.subscribe(params => {
  const category = params['category'];
  const sort = params['sort'];
});
```

---

## Lazy Loading (Optimización)

**Cargar módulos solo cuando se necesitan:**

```typescript
const routes: Routes = [
  { path: '', component: HomeComponent },
  {
    path: 'admin',
    loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule)
  },
  {
    path: 'dashboard',
    loadChildren: () => import('./dashboard/dashboard.module').then(m => m.DashboardModule)
  }
];
```

---

## Child Routes (Rutas Hijas)

```typescript
const routes: Routes = [
  {
    path: 'productos',
    component: ProductsLayoutComponent,
    children: [
      { path: '', component: ProductsListComponent },
      { path: ':id', component: ProductDetailComponent }
    ]
  }
];
```

---

## Checklist - Routing Profesional

- [ ] AppRoutingModule creado
- [ ] Rutas configuradas correctamente
- [ ] <router-outlet> en template
- [ ] routerLink en navegación
- [ ] router.navigate() funciona
- [ ] Parámetros capturados correctamente
- [ ] Lazy loading implementado
- [ ] 404 page configurada
- [ ] Sin memory leaks en subscriptions

---

## Próximo

Módulo 14: Layouts - Master-Detail

```ts
// Ejemplo simplificado para entender la idea
```

La idea no es memorizar esto, sino entender:
- qué entra
- qué sale
- qué responsabilidad tiene
- cómo se conecta con el resto de la app

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
