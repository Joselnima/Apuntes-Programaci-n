# Módulo 05 - Estructura Profesional del Proyecto

> **Organiza tu código como los expertos lo hacen**

---

## ¿QUÉ es la estructura de carpetas?

**Plan de tu casa ANTES de construir**

Analogía:
- Caótico: todo en el piso (¡confusión!)
- Organizado: cocina, dormitorio, baño (¡funciona!)

En Angular:
- ❌ Todos los archivos en `src/app/`
- ✅ Carpetas: `components/`, `services/`, `pages/`, `models/`

**Resultado:**
```
Escalable → 10 archivos o 1000 archivos, se entiende igual
Mantenible → Dos desarrolladores no pisarse
Profesional → Se ve como código de verdad
```

---

## ¿PARA QUÉ necesitas estructura?

### Problema 1: El Proyecto Crece

```
Semana 1: 5 archivos (todo está bien)
Mes 1:   30 archivos (¿dónde está todo?)
Mes 6:   300 archivos (¡CAOS!)
```

**Sin estructura:**
- Tardan 10 minutos en encontrar un archivo
- Componentes duplicados
- Servicios perdidos
- Nadie entiende nada

**Con estructura clara:**
- Todo en su lugar
- Reutilizable
- Escalable

### Problema 2: Lógica Compartida

```
❌ Lógica de carrito en app.component.ts
✅ Lógica en cart.service.ts
    ↓
    app.component usa el servicio
    ↓
    product.component también usa el servicio
```

### Problema 3: Cambios de Diseño

```
❌ Una página con 500 líneas
✅ Descompuesta en: header, product-list, cart, footer
   
Cambio en header:
❌ Toca 500 líneas (riesgo)
✅ Toca solo header.component.ts (seguro)
```

---

## ¿CÓMO organizar profesionalmente?

### Estructura Base (Para Proyectos Medianos)

```
src/
├── app/
│   ├── components/
│   │   ├── header/
│   │   │   ├── header.component.ts
│   │   │   ├── header.component.html
│   │   │   ├── header.component.css
│   │   │   └── header.component.spec.ts
│   │   ├── footer/
│   │   ├── sidebar/
│   │   └── navigation/
│   │
│   ├── pages/
│   │   ├── home/
│   │   │   └── home.component.ts
│   │   ├── products/
│   │   │   └── products.component.ts
│   │   ├── product-detail/
│   │   ├── cart/
│   │   └── checkout/
│   │
│   ├── services/
│   │   ├── product.service.ts
│   │   ├── cart.service.ts
│   │   ├── auth.service.ts
│   │   └── api.service.ts
│   │
│   ├── models/
│   │   ├── product.ts
│   │   ├── user.ts
│   │   ├── cart-item.ts
│   │   └── order.ts
│   │
│   ├── guards/
│   │   ├── auth.guard.ts
│   │   └── admin.guard.ts
│   │
│   ├── interceptors/
│   │   └── auth.interceptor.ts
│   │
│   ├── pipes/
│   │   ├── currency.pipe.ts
│   │   └── date.pipe.ts
│   │
│   ├── app.component.ts      ← Raíz
│   ├── app.component.html
│   ├── app.component.css
│   └── app.module.ts
│
├── assets/
│   ├── images/
│   ├── icons/
│   └── data/
│
├── environments/
│   ├── environment.ts        ← Dev
│   └── environment.prod.ts   ← Producción
│
├── main.ts
├── index.html
├── styles.css                ← CSS global
└── polyfills.ts
```

---

### Capas Explicadas

| Carpeta | Qué Contiene | Ejemplo |
|---------|-------------|---------|
| `components/` | Piezas reutilizables | Header, botón, tarjeta |
| `pages/` | Vistas completas | Home, productos, carrito |
| `services/` | Lógica compartida | API calls, auth, cart |
| `models/` | Interfaces TypeScript | Product, User, Order |
| `guards/` | Protección de rutas | Solo admin puede ver |
| `interceptors/` | Middleware HTTP | Agregar token automático |
| `pipes/` | Transformar datos | Formatear precio |
| `assets/` | Recursos estáticos | Imágenes, fonts, data |
| `environments/` | Configuración | URLs dev vs prod |

---

## Patrón de Organización: Feature Folders (AVANZADO)

**Para proyectos grandes (100+ componentes):**

```
src/app/
├── shared/
│   ├── components/
│   │   ├── header/
│   │   └── footer/
│   ├── services/
│   ├── models/
│   └── pipes/
│
├── products/                ← Feature 1
│   ├── components/
│   │   └── product-list/
│   ├── services/
│   │   └── product.service.ts
│   ├── models/
│   │   └── product.ts
│   └── products.module.ts
│
├── cart/                    ← Feature 2
│   ├── components/
│   ├── services/
│   │   └── cart.service.ts
│   ├── models/
│   │   └── cart.ts
│   └── cart.module.ts
│
└── checkout/               ← Feature 3
    ├── components/
    ├── services/
    └── checkout.module.ts
```

**Ventaja:** Puedes trabajar en `products/` sin tocar `cart/`

---

## Ejemplos: Cómo Se Ve en Código

### Componente: `header/header.component.ts`

```typescript
import { Component, OnInit } from '@angular/core';
import { CartService } from '../../services/cart.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  cartCount = 0;
  
  constructor(private cartService: CartService) {}
  
  ngOnInit() {
    this.cartService.getCartCount().subscribe(count => {
      this.cartCount = count;
    });
  }
}
```

### Template: `header/header.component.html`

```html
<header>
  <h1>Mi Tienda</h1>
  <nav>
    <a href="/">Inicio</a>
    <a href="/productos">Productos</a>
    <a href="/carrito">Carrito ({{ cartCount }})</a>
  </nav>
</header>
```

### Service: `services/cart.service.ts`

```typescript
import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CartService {
  private cartCount = new BehaviorSubject(0);
  
  constructor() { }
  
  addToCart() {
    this.cartCount.next(this.cartCount.value + 1);
  }
  
  getCartCount() {
    return this.cartCount.asObservable();
  }
}
```

### Model: `models/product.ts`

```typescript
export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  inStock: boolean;
}
```

---

## Convenciones de Nombres

Cosas están siempre donde ESPERAS encontrarlas:

| Cosa | Nombre | Archivo |
|------|--------|---------|
| Componente | `HeaderComponent` | `header.component.ts` |
| Servicio | `CartService` | `cart.service.ts` |
| Model | `Product` | `product.ts` |
| Guard | `AuthGuard` | `auth.guard.ts` |
| Pipe | `CurrencyPipe` | `currency.pipe.ts` |

**Regla:** `función.tipo-de-archivo.ts`

---

## Cómo Generar Automáticamente

Angular CLI crea estructura correcta:

```bash
# Componente
ng generate component components/header

# Service
ng generate service services/cart

# Model (archivo)
ng generate interface models/product

# Guard
ng generate guard guards/auth

# Pipe
ng generate pipe pipes/currency
```

**Resultado:** Carpetas creadas automáticamente en el lugar correcto ✅

---

## Errores Comunes

### ❌ Error 1: Todo en app.component

```typescript
// Aquí hay COMPONENTES, SERVICIOS, LÓGICA
// 2000 líneas en UN archivo
// ¡Imposible de mantener!
export class AppComponent {
  // ... 2000 líneas de caos
}
```

**✅ Corrección:** Separa en componentes, servicios, modelos

### ❌ Error 2: No Reutilizar Servicios

```typescript
// En ProductsComponent
const url = 'http://api.com/products';
// Generas GET request acá

// En CartComponent
const url = 'http://api.com/products';
// Igual, código duplicado
```

**✅ Corrección:** Crea `product.service.ts` que usano ambos

### ❌ Error 3: Carpetas Anidadas Demasiado Profundas

```
src/app/features/products/components/
  product-list/views/table/container/
    product-table.component.ts
```

**✅ Regla:** Máximo 3-4 niveles de profundidad

---

## Referencia Rápida

```
CREAR COMPONENTE CORRECTO:

1. Archivo: src/app/components/my-component/
2. Ejecuta: ng generate component components/my-component
3. Resultado:
   my-component.component.ts      ← Lógica
   my-component.component.html    ← Template
   my-component.component.css     ← Estilos
   my-component.component.spec.ts ← Tests

CONECTAR SERVICIO:

import { MyService } from '../../services/my.service';

constructor(private myService: MyService) {}
```

---

## Checklist - Estructura Profesional

- [ ] Carpetas creadas: components/, pages/, services/, models/
- [ ] Ningún archivo TypeScript en raíz de `src/app/`
- [ ] Nombres siguen patrón: `nombre.tipo.ts`
- [ ] Servicios en `services/`
- [ ] Interfaces en `models/`
- [ ] Componentes en `components/` o `pages/`
- [ ] Sin archivos más profundos que 4 niveles
- [ ] CLI puede generar archivos correctamente

---

## Próximo

Módulo 06: Componentes - La Base de Angular

---

## Resumen

Si puedes explicarlo con tus palabras, vas bien.
Si solo lo reconoces cuando lo ves, todavía falta práctica.
