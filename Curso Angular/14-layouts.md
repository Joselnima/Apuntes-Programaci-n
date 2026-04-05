# Módulo 14 - Layouts: Master-Detail y Navegación

> **Patrones profesionales de estructura**

---

## ¿QÉ es un layout?

**Estructura de componentes para navegación**

Patrón Master-Detail:
- Izquierda: Lista (Master)
- Derecha: Detalle (Detail)

```
┌─────────────────┐
│    Header       │
├────────┬────────┤
│        │        │
│Master  │Detail  │
│        │        │
├────────┴────────┤
│    Footer       │
└─────────────────┘
```

---

## ¿PARA QÉ?

```
❌ Sin layout: Componentes desorganizados
✅ Con layout: Estructura clara + reutilizable
```

---

## Master-Detail Completo

**app.component.html:**

```html
<app-header></app-header>

<div class="container">
  <app-master 
    [items]="products"
    (onSelect)="selectProduct($event)"
    [selectedId]="selectedProductId"
  ></app-master>
  
  <app-detail 
    *ngIf="selectedProduct"
    [item]="selectedProduct"
    (onUpdate)="updateProduct($event)"
  ></app-detail>
</div>

<app-footer></app-footer>
```

**container.component.ts:**

```typescript
export class ContainerComponent implements OnInit {
  products: any[] = [];
  selectedProductId: number | null = null;
  selectedProduct: any = null;
  
  constructor(private productService: ProductService) {}
  
  ngOnInit() {
    this.productService.getProducts()
      .subscribe(data => this.products = data);
  }
  
  selectProduct(id: number) {
    this.selectedProductId = id;
    this.selectedProduct = this.products.find(p => p.id === id);
  }
  
  updateProduct(updated: any) {
    const index = this.products.findIndex(p => p.id === updated.id);
    this.products[index] = updated;
  }
}
```

---

## Sidebar Layout

```html
<div class="app-layout">
  <app-sidebar 
    [menuItems]="menuItems"
    (onSelect)="navigateTo($event)"
  ></app-sidebar>
  
  <main>
    <router-outlet></router-outlet>
  </main>
</div>
```

```css
.app-layout {
  display: flex;
}

app-sidebar {
  width: 250px;
  background: #f8f9fa;
}

main {
  flex: 1;
  padding: 20px;
}
```

---

## Nested Routes + Layout

```typescript
const routes: Routes = [
  {
    path: 'dashboard',
    component: DashboardLayoutComponent,
    children: [
      { path: '', component: DashboardComponent },
      { path: 'profile', component: ProfileComponent },
      { path: 'settings', component: SettingsComponent }
    ]
  }
];
```

---

## Checklist

- [ ] Layout component creado
- [ ] Master-detail funciona
- [ ] Sidebar en navegación
- [ ] Child outlets configurados
- [ ] Responsive design
- [ ] Sin memory leaks

---

## Próximo

Módulo 15: Organización Profesional

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
