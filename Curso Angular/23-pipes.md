# Capítulo 23: Pipes - Los Transformadores de Datos

Imagina que tienes una caja de herramientas en tu taller. Cada herramienta tiene un propósito específico: el martillo clava clavos, el destornillador gira tornillos, las tenazas cortan alambres. En Angular, los pipes son esas herramientas especializadas: transforman datos crudos en información presentable y útil.

Este capítulo te enseñará a usar las herramientas integradas de Angular y a crear las tuyas propias para dar formato a tus datos.

## ¿Qué Aprenderás en Este Capítulo?

En este capítulo, exploraremos:
- La filosofía detrás de los pipes y su importancia
- Cómo usar los pipes integrados de Angular
- La diferencia entre pipes puros e impuros
- Cómo crear tus propios pipes personalizados
- El poderoso pipe `async` para trabajar con observables
- Mejores prácticas y errores comunes

## La Filosofía de los Pipes

Piensa en un pipe como un chef especializado. Tú le das los ingredientes crudos (datos), y él los transforma en algo delicioso y presentable (información formateada). Los pipes mantienen tu template limpio y legible, separando la lógica de presentación de la lógica de negocio.

### Analogía de la Fábrica

- **Sin pipes**: Tendrías que procesar cada dato manualmente en el componente
- **Con pipes**: La transformación ocurre automáticamente en la vista

```typescript
// ❌ Sin pipes - lógica en el componente
export class ProductComponent {
  getFormattedPrice() {
    return '$' + this.product.price.toFixed(2);
  }

  getFormattedDate() {
    return this.product.releaseDate.toLocaleDateString();
  }
}
```

```html
<!-- Template cluttered -->
<p>Price: {{ getFormattedPrice() }}</p>
<p>Release: {{ getFormattedDate() }}</p>
```

```typescript
// ✅ Con pipes - transformación en la vista
export class ProductComponent {
  // Solo datos, sin lógica de formateo
}
```

```html
<!-- Template clean -->
<p>Price: {{ product.price | currency:'USD' }}</p>
<p>Release: {{ product.releaseDate | date:'shortDate' }}</p>
```

## Pipes Integrados de Angular

Angular viene con una caja de herramientas completa. Vamos a explorar las más útiles:

### Pipes de Texto

```html
<!-- Mayúsculas y minúsculas -->
<p>{{ "hola mundo" | uppercase }}</p>        <!-- "HOLA MUNDO" -->
<p>{{ "HOLA MUNDO" | lowercase }}</p>        <!-- "hola mundo" -->
<p>{{ "hola mundo" | titlecase }}</p>        <!-- "Hola Mundo" -->

<!-- Recorte de texto -->
<p>{{ "Texto muy largo que necesito cortar" | slice:0:20 }}</p>  <!-- "Texto muy largo que" -->
<p>{{ "Texto muy largo que necesito cortar" | slice:-10 }}</p>   <!-- "cortar" -->
```

### Pipes de Números y Moneda

```html
<!-- Números -->
<p>{{ 3.14159 | number:'1.2-2' }}</p>        <!-- "3.14" -->
<p>{{ 1234.56 | number:'1.0-0' }}</p>        <!-- "1,235" -->
<p>{{ 0.5 | percent }}</p>                   <!-- "50%" -->

<!-- Moneda -->
<p>{{ 1234.56 | currency:'USD' }}</p>        <!-- "$1,234.56" -->
<p>{{ 1234.56 | currency:'EUR':'symbol':'1.0-0' }}</p>  <!-- "€1,235" -->
<p>{{ 1234.56 | currency:'COP':'symbol-narrow' }}</p>   <!-- "$1,234.56" -->
```

### Pipes de Fecha

```html
<!-- Formatos de fecha -->
<p>{{ today | date:'short' }}</p>             <!-- "1/15/23, 2:30 PM" -->
<p>{{ today | date:'mediumDate' }}</p>        <!-- "Jan 15, 2023" -->
<p>{{ today | date:'longDate' }}</p>          <!-- "January 15, 2023" -->
<p>{{ today | date:'shortTime' }}</p>         <!-- "2:30 PM" -->

<!-- Formato personalizado -->
<p>{{ today | date:'dd/MM/yyyy' }}</p>        <!-- "15/01/2023" -->
<p>{{ today | date:'EEEE, MMMM d' }}</p>      <!-- "Sunday, January 15" -->
```

### Pipes de Arrays y Objetos

```html
<!-- Arrays -->
<p>{{ [1,2,3,4,5] | slice:1:4 }}</p>          <!-- "2,3,4" -->
<p>{{ ['a','b','c'] | join:' - ' }}</p>       <!-- "a - b - c" -->

<!-- JSON para debugging -->
<pre>{{ user | json }}</pre>

<!-- KeyValue para objetos -->
<div *ngFor="let item of user | keyvalue">
  {{ item.key }}: {{ item.value }}
</div>
```

## Creando Tus Propios Pipes

Cuando los pipes integrados no alcanzan, crea los tuyos. Es como fabricar una herramienta personalizada para tu taller.

### Un Pipe Simple: Truncar Texto

```typescript
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'truncate'
})
export class TruncatePipe implements PipeTransform {
  transform(value: string, limit: number = 20, suffix: string = '...'): string {
    if (!value) return '';

    if (value.length <= limit) return value;

    return value.slice(0, limit) + suffix;
  }
}
```

```html
<!-- Uso básico -->
<p>{{ "Este es un texto muy largo que necesito cortar" | truncate }}</p>
<!-- "Este es un texto mu..." -->

<!-- Con parámetros personalizados -->
<p>{{ longText | truncate:50:' [leer más]' }}</p>
<!-- "Este es un texto muy largo que necesito cortar [leer más]" -->
```

### Un Pipe Más Complejo: Calcular Edad

```typescript
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'age'
})
export class AgePipe implements PipeTransform {
  transform(birthDate: Date | string): number {
    if (!birthDate) return 0;

    const birth = new Date(birthDate);
    const today = new Date();

    let age = today.getFullYear() - birth.getFullYear();
    const monthDiff = today.getMonth() - birth.getMonth();

    // Si no ha cumplido años este año
    if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
      age--;
    }

    return age;
  }
}
```

```html
<p>{{ user.birthDate | age }} años</p>
```

### Un Pipe con Múltiples Parámetros: Formato Condicional

```typescript
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'conditionalFormat'
})
export class ConditionalFormatPipe implements PipeTransform {
  transform(
    value: number,
    threshold: number = 0,
    goodClass: string = 'text-success',
    badClass: string = 'text-danger'
  ): string {
    const className = value >= threshold ? goodClass : badClass;
    return `<span class="${className}">${value}</span>`;
  }
}
```

```html
<!-- Para temperaturas -->
<p innerHTML="{{ temperature | conditionalFormat:20:'text-hot':'text-cold' }}"></p>

<!-- Para calificaciones -->
<p innerHTML="{{ score | conditionalFormat:70 }}"></p>
```

## La Pureza de los Pipes: Un Concepto Crucial

Los pipes pueden ser **puros** o **impuros**. Esta diferencia afecta el rendimiento y el comportamiento.

### Pipes Puros (Pure Pipes)

```typescript
@Pipe({
  name: 'square',
  pure: true  // Este es el valor por defecto
})
export class SquarePipe implements PipeTransform {
  transform(value: number): number {
    console.log('Calculando cuadrado...');
    return value * value;
  }
}
```

**Comportamiento**: Solo se ejecuta cuando cambia la **referencia** del valor de entrada.

```typescript
// En el componente
export class AppComponent {
  numbers = [1, 2, 3, 4, 5];

  addNumber() {
    this.numbers.push(6);  // ❌ El pipe NO se ejecuta (misma referencia)
  }

  replaceArray() {
    this.numbers = [...this.numbers, 6];  // ✅ El pipe SÍ se ejecuta (nueva referencia)
  }
}
```

### Pipes Impuros (Impure Pipes)

```typescript
@Pipe({
  name: 'filter',
  pure: false  // ¡Atención! Afecta el rendimiento
})
export class FilterPipe implements PipeTransform {
  transform(items: any[], searchTerm: string): any[] {
    console.log('Filtrando...');
    if (!searchTerm) return items;

    return items.filter(item =>
      item.name.toLowerCase().includes(searchTerm.toLowerCase())
    );
  }
}
```

**Comportamiento**: Se ejecuta en **cada ciclo de detección de cambios**.

```typescript
// Se ejecuta cada vez que Angular detecta cambios
// Útil para arrays que se modifican internamente
this.filteredItems = this.allItems.filter(/*...*/);  // Se ejecuta constantemente
```

### ¿Cuándo Usar Cada Uno?

| Situación | Pipe Tipo | Razón |
|-----------|-----------|-------|
| Formatear números/fechas | Puro | Los valores primitivos cambian por referencia |
| Filtrar arrays mutables | Impuro | El array cambia internamente |
| Transformar objetos | Puro | Los objetos cambian por referencia |
| Buscar en tiempo real | Impuro | Necesitas recalcular en cada tecla |

## El Pipe Async: Tu Aliado con Observables

El pipe `async` es como un mayordomo que se encarga de todas las suscripciones por ti.

### El Problema que Resuelve

```typescript
// ❌ Sin async pipe - gestión manual de suscripciones
export class UserComponent implements OnInit, OnDestroy {
  user: User | null = null;
  private subscription?: Subscription;

  ngOnInit() {
    this.subscription = this.userService.getUser(1).subscribe(user => {
      this.user = user;
    });
  }

  ngOnDestroy() {
    this.subscription?.unsubscribe();  // ¡Fácil de olvidar!
  }
}
```

```typescript
// ✅ Con async pipe - Angular maneja las suscripciones
export class UserComponent {
  user$ = this.userService.getUser(1);  // Observable
}
```

```html
<!-- El pipe async se suscribe automáticamente -->
<div *ngIf="user$ | async as user">
  <h2>{{ user.name }}</h2>
  <p>{{ user.email }}</p>
</div>
```

### Usos Avanzados del Async Pipe

```typescript
export class DashboardComponent {
  // Múltiples observables
  users$ = this.userService.getUsers();
  stats$ = this.statsService.getStats();
  notifications$ = this.notificationService.getNotifications();

  // Combinar observables
  dashboardData$ = combineLatest({
    users: this.users$,
    stats: this.stats$,
    notifications: this.notifications$
  });
}
```

```html
<!-- Usar datos combinados -->
<div *ngIf="dashboardData$ | async as data">
  <h2>Dashboard</h2>
  <p>Usuarios: {{ data.users.length }}</p>
  <p>Estadísticas: {{ data.stats.total }}</p>
  <div *ngFor="let notification of data.notifications">
    {{ notification.message }}
  </div>
</div>

<!-- O usar por separado -->
<ul>
  <li *ngFor="let user of users$ | async">
    {{ user.name }}
  </li>
</ul>
```

## Un Ejemplo Completo: Sistema de Comentarios

Vamos a crear un sistema de comentarios que use múltiples pipes:

```typescript
// Pipes personalizados
@Pipe({ name: 'relativeTime' })
export class RelativeTimePipe implements PipeTransform {
  transform(date: Date): string {
    const now = new Date();
    const diffMs = now.getTime() - date.getTime();
    const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));

    if (diffDays === 0) return 'Hoy';
    if (diffDays === 1) return 'Ayer';
    if (diffDays < 7) return `Hace ${diffDays} días`;
    if (diffDays < 30) return `Hace ${Math.floor(diffDays / 7)} semanas`;

    return date.toLocaleDateString();
  }
}

@Pipe({ name: 'highlightMentions' })
export class HighlightMentionsPipe implements PipeTransform {
  transform(text: string): string {
    return text.replace(/@(\w+)/g, '<span class="mention">@$1</span>');
  }
}

@Pipe({ name: 'filterComments', pure: false })
export class FilterCommentsPipe implements PipeTransform {
  transform(comments: Comment[], filter: string): Comment[] {
    if (!filter) return comments;

    return comments.filter(comment =>
      comment.text.toLowerCase().includes(filter.toLowerCase()) ||
      comment.author.toLowerCase().includes(filter.toLowerCase())
    );
  }
}
```

```typescript
// Componente
export class CommentsComponent implements OnInit {
  comments$: Observable<Comment[]>;
  filterText = '';

  constructor(private commentService: CommentService) {}

  ngOnInit() {
    this.comments$ = this.commentService.getComments();
  }
}
```

```html
<div class="comments-section">
  <h3>Comentarios</h3>

  <!-- Filtro -->
  <input
    [(ngModel)]="filterText"
    placeholder="Buscar comentarios..."
    class="filter-input"
  >

  <!-- Lista de comentarios -->
  <div *ngIf="comments$ | async as comments">
    <div
      *ngFor="let comment of comments | filterComments:filterText"
      class="comment"
    >
      <div class="comment-header">
        <strong>{{ comment.author }}</strong>
        <span class="comment-date">
          {{ comment.createdAt | relativeTime }}
        </span>
      </div>

      <div
        class="comment-text"
        [innerHTML]="comment.text | highlightMentions"
      ></div>

      <div class="comment-actions">
        <span class="likes">{{ comment.likes | number:'1.0-0' }} likes</span>
        <button>Responder</button>
      </div>
    </div>
  </div>
</div>
```

## Mejores Prácticas con Pipes

### ✅ Hazlo:
- Pipes pequeños y enfocados en una sola responsabilidad
- Pipes puros siempre que sea posible
- Pipes declarados en módulos compartidos para reutilización
- Pipes con nombres descriptivos
- Pipes testeables unitariamente

### ❌ Evítalo:
- Pipes con lógica de negocio compleja
- Pipes impuros sin necesidad
- Pipes que modifican los datos originales
- Pipes con efectos secundarios
- Pipes excesivamente anidados

## Errores Comunes y Cómo Evitarlos

1. **Usar pipes impuros innecesariamente**: Reduce el rendimiento. Usa puros cuando puedas.

2. **Olvidar declarar el pipe**: Asegúrate de incluirlo en `declarations` del módulo.

3. **Pipes que mutan datos**: Los pipes deben ser funciones puras, sin efectos secundarios.

4. **No manejar valores null/undefined**: Siempre verifica la entrada.

5. **Pipes complejos en templates**: Si el pipe necesita muchos parámetros, considera usar un método en el componente.

6. **No usar OnPush con pipes impuros**: Puede causar problemas de detección de cambios.

## Mini Práctica: Tu Sistema de Pipes

1. **Crea un pipe `fileSize`** que convierta bytes a KB, MB, GB con formato legible
2. **Crea un pipe `creditCardMask`** que oculte dígitos de tarjetas de crédito
3. **Crea un pipe `weatherIcon`** que retorne íconos basados en condiciones climáticas
4. **Implementa un pipe `sort`** impuro para ordenar arrays dinámicamente
5. **Usa el pipe `async`** para mostrar datos de un servicio meteorológico

## Resumen del Capítulo

Los pipes son las herramientas especializadas de Angular para transformar datos en la vista. Mantienen tus templates limpios, tu código reutilizable, y tu aplicación performante cuando se usan correctamente.

Recuerda: pipes puros para la mayoría de casos, pipes impuros solo cuando necesites recalcular por mutaciones. El pipe `async` es tu mejor amigo para trabajar con observables. Y siempre, crea pipes pequeños y enfocados.

En el próximo capítulo, exploraremos los guards e interceptors, los guardianes de seguridad de tu aplicación Angular.
