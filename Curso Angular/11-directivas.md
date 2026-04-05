# Módulo 11 - Directivas

> **Controla el HTML dinámicamente**

---

## ¿QÉ es una directiva?

**Una orden que cambia el comportamiento del HTML**

Analogía:
- Director de cine: "Aquí la cámara se mueve así"
- Directiva: "Aquí el HTML se muestra así"

```html
<!-- Sin directivas: JSON estático -->
<div>
  <p>Producto 1</p>
  <p>Producto 2</p>
  <p>Producto 3</p>
  <!-- Copiar-pegar 100 veces = No escala -->
</div>

<!-- Con directivas: Dinámico -->
<div>
  <p *ngFor="let product of products">
    {{ product.name }}
  </p>
  <!-- Una línea, 100 productos. ✨ -->
</div>
```

---

## ¿PARA QUÉ?

### Problema 1: HTML Repetitivo

**❌ Sin directivas:**

```html
<div>
  <h1 *ngIf="user">Hola {{ user.name }}</h1>
  <h1 *ngIf="!user">Inicia sesión</h1>
</div>

<!-- Mejor: -->

<h1 *ngIf="user; else noUser">
  Hola {{ user.name }}
</h1>

<ng-template #noUser>
  <h1>Inicia sesión</h1>
</ng-template>
```

### Problema 2: Listas sin escala

**❌ Sin ngFor:**

```html
<div>
  <p>{{ products[0].name }}</p>
  <p>{{ products[1].name }}</p>
  <p>{{ products[2].name }}</p>
  <!-- 500 líneas de copiar-pegar -->
</div>
```

**✅ Con ngFor:**

```html
<div>
  <p *ngFor="let product of products">
    {{ product.name }}
  </p>
  <!-- 2 líneas, escala a infinito -->
</div>
```

### Problema 3: HTML Rígido

**❌ Sin ngClass/ngStyle:**

```html
<!-- HTML fijo -->
<div class="card active">
  {{ product.name }}
</div>

<!-- Si status cambia, ¿qué pasa? Nada. -->
```

**✅ Con ngClass:**

```html
<!-- HTML dinámico -->
<div [ngClass]="{ 'active': isActive, 'disabled': isDisabled }">
  {{ product.name }}
</div>

<!-- Cambia automáticamente -->
```

---

## ¿CÓMO?

### Directiva 1: *ngIf (Mostrar/Ocultar)

**Controla si el elemento existe en el DOM**

```typescript
export class HelloComponent {
  isLoggedIn = false;
  user = { name: 'Juan', role: 'admin' };
}
```

```html
<!-- Simple: mostrar si true -->
<button *ngIf="isLoggedIn">Logout</button>

<!-- O mostrar solo si false -->
<button *ngIf="!isLoggedIn">Login</button>

<!-- If-else con ng-template -->
<div *ngIf="user; else notLoggedIn">
  <p>Welcome, {{ user.name }}</p>
</div>

<ng-template #notLoggedIn>
  <p>Por favor inicia sesión</p>
</ng-template>

<!-- If-else-if (then-else-then) -->
<div *ngIf="user.role === 'admin'; then adminTemplate else userTemplate">
</div>

<ng-template #adminTemplate>
  <button>Panel de administración</button>
</ng-template>

<ng-template #userTemplate>
  <button>Mis datos</button>
</ng-template>
```

**Diferencia:** *ngIf **elimina** del DOM, no solo oculta

```
*ngIf vs CSS display:none

*ngIf (no existe)           | CSS display:none (existe)
- No existe en DOM          | - Existe en DOM
- Componente no inicia      | - Componente inicia
- Menos memoria             | - Más memoria
- Para ocultar (lógica)     | - Para estilos
```

---

### Directiva 2: *ngFor (Listas)

**Itera sobre un array**

```typescript
export class ProductListComponent {
  products = [
    { id: 1, name: 'Laptop', price: 1000 },
    { id: 2, name: 'Mouse', price: 20 },
    { id: 3, name: 'Teclado', price: 50 }
  ];
}
```

```html
<!-- Básico: itera products -->
<ul>
  <li *ngFor="let product of products">
    {{ product.name }} - ${{ product.price }}
  </li>
</ul>

<!-- Con índice -->
<ul>
  <li *ngFor="let product of products; let i = index">
    #{{ i + 1 }}: {{ product.name }}
  </li>
</ul>

<!-- Con primero/último -->
<ul>
  <li *ngFor="let product of products; let first = first; let last = last"
      [class.first-item]="first"
      [class.last-item]="last">
    {{ product.name }}
  </li>
</ul>

<!-- Con even/odd -->
<ul>
  <li *ngFor="let product of products; let even = even; let odd = odd"
      [class.row-even]="even"
      [class.row-odd]="odd">
    {{ product.name }}
  </li>
</ul>

<!-- Contar desde index -->
<ul>
  <li *ngFor="let product of products; let count = count">
    Item {{ count }} of {{ products.length }}
  </li>
</ul>
```

**Variables disponibles:**

| Variable | Qué es |
|----------|--------|
| `let item` | Elemento actual |
| `let i = index` | Posición (0, 1, 2...) |
| `let first = first` | ¿Es el primero? |
| `let last = last` | ¿Es el último? |
| `let even = even` | ¿Índice par? |
| `let odd = odd` | ¿Índice impar? |
| `let count = count` | Contar desde 1 |

### Optimizar ngFor con trackBy

**❌ Lento (recalcula todo):**

```html
<ul>
  <li *ngFor="let product of products">
    {{ product.name }}
  </li>
</ul>
```

**✅ Rápido (solo nuevos):**

```typescript
trackByProductId(index: number, item: any) {
  return item.id;
}
```

```html
<ul>
  <li *ngFor="let product of products; trackBy: trackByProductId">
    {{ product.name }}
  </li>
</ul>
```

**Por qué:**
- Sin trackBy: Angular recrea todos los elementos cuando cambia la lista
- Con trackBy: Angular solo actualiza los que cambiaron

---

### Directiva 3: [ngClass] y [ngStyle]

**Asigna clases/estilos dinámicamente**

```typescript
export class StatusComponent {
  status = 'success';  // 'success', 'warning', 'error'
  isHighlight = true;
  fontSize = 16;
}
```

#### ngClass

```html
<!-- Objeto: clase si condition -->
<div [ngClass]="{ 'active': isActive, 'disabled': isDisabled }">
  Contenido
</div>

<!-- String condicional -->
<div [ngClass]="status === 'success' ? 'text-green' : 'text-red'">
  {{ status }}
</div>

<!-- Array de clases -->
<div [ngClass]="['base-class', status, isHighlight ? 'highlight' : '']">
  Contenido
</div>
```

**CSS:**

```css
.active {
  background: green;
}

.disabled {
  opacity: 0.5;
}

.text-green {
  color: green;
}

.text-red {
  color: red;
}

.highlight {
  border: 2px solid yellow;
}
```

#### ngStyle

```html
<!-- Objeto de estilos -->
<div [ngStyle]="{
  'color': isError ? 'red' : 'black',
  'font-size.px': fontSize,
  'background': isHighlight ? 'yellow' : 'white'
}">
  Contenido con estilos dinámicos
</div>

<!-- Condicional completo -->
<div [ngStyle]="isActive ? { 'background': 'green' } : { 'background': 'gray' }">
  Estado
</div>
```

---

### Directiva 4: [ngSwitch]

**Como switch-case en HTML**

```typescript
export class UserRoleComponent {
  userRole = 'user';  // 'admin', 'moderator', 'user'
}
```

```html
<div [ngSwitch]="userRole">
  <div *ngSwitchCase="'admin'">
    <button>Panel de Administración</button>
  </div>
  
  <div *ngSwitchCase="'moderator'">
    <button>Moderar Contenido</button>
  </div>
  
  <div *ngSwitchCase="'user'">
    <button>Ver Perfil</button>
  </div>
  
  <div *ngSwitchDefault>
    <p>Usuario no reconocido</p>
  </div>
</div>
```

---

### Directiva 5: @for y @if (Angular 17+)

**Nueva sintaxis más limpia:**

```html
<!-- Viejo: *ngFor -->
<div *ngFor="let item of items">
  {{ item.name }}
</div>

<!-- Nuevo: @for -->
<div @for="let item of items">
  {{ item.name }}
</div>

<!-- Viejo: *ngIf -->
<div *ngIf="isVisible">
  Contenido
</div>

<!-- Nuevo: @if -->
@if (isVisible) {
  <div>Contenido</div>
}

<!-- Viejo: *ngIf else -->
<div *ngIf="user; else noUser">
  {{ user.name }}
</div>
<ng-template #noUser>
  No hay usuario
</ng-template>

<!-- Nuevo: @if else -->
@if (user) {
  <div>{{ user.name }}</div>
} @else {
  <div>No hay usuario</div>
}
```

---

## Ejemplo Real: Tabla Dinámica

```typescript
export class TableComponent {
  data = [
    { id: 1, name: 'Laptop', category: 'Electronics', active: true },
    { id: 2, name: 'Teclado', category: 'Accesorios', active: false },
    { id: 3, name: 'Mouse', category: 'Accesorios', active: true }
  ];
  
  sortBy = 'name';
  
  trackByItemId(index: number, item: any) {
    return item.id;
  }
}
```

```html
<table>
  <thead>
    <tr>
      <th>#</th>
      <th>Nombre</th>
      <th>Categoría</th>
      <th>Estado</th>
    </tr>
  </thead>
  <tbody>
    <tr *ngFor="let item of data; let i = index; trackBy: trackByItemId"
        [ngClass]="{ 'row-inactive': !item.active }">
      <td>{{ i + 1 }}</td>
      <td>{{ item.name }}</td>
      <td>{{ item.category }}</td>
      <td>
        <span *ngIf="item.active; else inactive" class="badge-active">
          Activo
        </span>
        <ng-template #inactive>
          <span class="badge-inactive">Inactivo</span>
        </ng-template>
      </td>
    </tr>
  </tbody>
</table>

<!-- Mensaje si no hay datos -->
<div *ngIf="!data || data.length === 0" class="empty-message">
  No hay datos para mostrar
</div>
```

---

## Directivas Custom (Avanzado)

```typescript
import { Directive, ElementRef, HostListener, Input } from '@angular/core';

@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective {
  @Input() appHighlight = 'yellow';
  
  constructor(private el: ElementRef) {}
  
  @HostListener('mouseenter')
  onMouseEnter() {
    this.el.nativeElement.style.backgroundColor = this.appHighlight;
  }
  
  @HostListener('mouseleave')
  onMouseLeave() {
    this.el.nativeElement.style.backgroundColor = '';
  }
}
```

**Usar:**

```html
<p [appHighlight]="'lightblue'">
  Pasa el ratón
</p>
```

---

## Performance: *ngFor best practices

```html
<!-- ❌ LENTO: Sin trackBy -->
<div *ngFor="let item of items">
  {{ item.name }}
</div>

<!-- ✅ RÁPIDO: Con trackBy -->
<div *ngFor="let item of items; trackBy: trackById">
  {{ item.name }}
</div>

<!-- ✅ RÁPIDO: OnPush en componentes -->
<div *ngFor="let item of items; trackBy: trackById">
  <app-card [item]="item"></app-card>
</div>

<!-- ✅ RÁPIDO: async pipe -->
<div *ngFor="let item of items$ | async; trackBy: trackById">
  {{ item.name }}
</div>
```

---

## Checklist - Directivas Profesionales

- [ ] Usas *ngIf para condicionales
- [ ] Usas *ngFor con trackBy
- [ ] Usas [ngClass] para clases dinámicas
- [ ] Usas [ngStyle] solo cuando necesita
- [ ] Sin HTML duplicado
- [ ] *ngFor tiene `trackBy` si lista es grande
- [ ] `ng-template` solo para casos especiales
- [ ] Performance: OnPush en componentes iterados
- [ ] Sin lógica compleja en template

---

## Próximo

Módulo 12: Comunicación Entre Componentes - Padre, hijo y hermanos

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
