# Módulo 19 - Estado Local: Gestión de Datos

> **Controla el estado sin que se salga de control**

---

## ¿QÉ es estado local?

**Información que vive en servicios/componentes**

```
Estado Local:
- Carrito con items
- Usuario logueado
- Tema seleccionado
- Filtros aplicados

Vs

Estado Global (Redux, NGRX):
- Datos compartidos TODA la app
```

---

## Estrategia 1: BehaviorSubject + Observable

**Centro de verdad único**

```typescript
@Injectable()
export class CartStateService {
  private cartSubject = new BehaviorSubject<CartItem[]>([]);
  cart$ = this.cartSubject.asObservable();
  
  private totalSubject = new BehaviorSubject<number>(0);
  total$ = this.totalSubject.asObservable();
  
  addItem(item: CartItem) {
    const current = this.cartSubject.value;
    const updated = [...current, item];
    this.cartSubject.next(updated);
    this.updateTotal(updated);
  }
  
  private updateTotal(items: CartItem[]) {
    const total = items.reduce((sum, item) =>
      sum + (item.price * item.quantity), 0
    );
    this.totalSubject.next(total);
  }
}
```

**Componente usa async pipe:**

```html
<div *ngFor="let item of (cart$ | async)">
  {{ item.name }}
</div>
<p>Total: ${{ total$ | async }}</p>
```

---

## Estrategia 2: Service con Métodos

```typescript
@Injectable()
export class AppStateService {
  private state = {
    user: null,
    theme: 'light',
    sidebarOpen: true
  };
  
  getUser() {
    return this.state.user;
  }
  
  setUser(user: any) {
    this.state.user = user;
  }
  
  toggleTheme() {
    this.state.theme = this.state.theme === 'light' ? 'dark' : 'light';
  }
  
  getState() {
    return { ...this.state };  // Retorna copia
  }
}
```

---

## Estrategia 3: Usando Signals (Angular 14+)

```typescript
@Injectable()
export class TodoStateService {
  todos = signal<Todo[]>([]);
  selectedTodo = signal<Todo | null>(null);
  
  addTodo(todo: Todo) {
    this.todos.update(current => [...current, todo]);
  }
  
  selectTodo(id: number) {
    const todo = this.todos().find(t => t.id === id);
    this.selectedTodo.set(todo || null);
  }
}
```

**Componente:**

```typescript
export class TodoComponent {
  todos = this.stateService.todos;
  selectedTodo = this.stateService.selectedTodo;
  
  constructor(private stateService: TodoStateService) {}
}
```

```html
<div *ngFor="let todo of todos()">
  {{ todo.title }}
</div>
<p *ngIf="selectedTodo()">
  Selected: {{ selectedTodo()?.title }}
</p>
```

---

## Performance: OnPush + Signals

```typescript
@Component({
  selector: 'app-todo',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <div *ngFor="let todo of todos()">
      {{ todo.title }}
    </div>
  `
})
export class TodoComponent {
  todos = inject(TodoStateService).todos;
}
```

---

## Memory Leaks en Estado

### ❌ Error: Suscripciones olvidadas

```typescript
ngOnInit() {
  this.stateService.items$.subscribe(items => {
    this.items = items;
  });
  // Sin unsubscribe en ngOnDestroy ❌
}
```

### ✅ Solución: takeUntil

```typescript
destroy$ = new Subject<void>();

ngOnInit() {
  this.stateService.items$
    .pipe(takeUntil(this.destroy$))
    .subscribe(items => this.items = items);
}

ngOnDestroy() {
  this.destroy$.next();
}
```

### ✅ Mejor: async pipe

```html
<div *ngFor="let item of (items$ | async)">
  {{ item }}
</div>
```

---

## Ejemplo Completo: Todo App

```typescript
@Injectable()
export class TodoStateService {
  private todosSubject = new BehaviorSubject<Todo[]>([]);
  todos$ = this.todosSubject.asObservable();
  
  private filterSubject = new BehaviorSubject<'all'|'done'|'pending'>('all');
  filter$ = this.filterSubject.asObservable();
  
  filtered$ = combineLatest([
    this.todos$,
    this.filter$
  ]).pipe(
    map(([todos, filter]) =>
      filter === 'all' ? todos :
      filter === 'done' ? todos.filter(t => t.done) :
      todos.filter(t => !t.done)
    )
  );
  
  addTodo(title: string) {
    const todo: Todo = {
      id: Date.now(),
      title,
      done: false
    };
    const current = this.todosSubject.value;
    this.todosSubject.next([...current, todo]);
  }
  
  toggleTodo(id: number) {
    const current = this.todosSubject.value.map(t =>
      t.id === id ? { ...t, done: !t.done } : t
    );
    this.todosSubject.next(current);
  }
  
  setFilter(filter: 'all'|'done'|'pending') {
    this.filterSubject.next(filter);
  }
}
```

---

## Checklist - Estado Local

- [ ] Un servicio por dominio (CartService, UserService, etc)
- [ ] BehaviorSubject para estado compartido
- [ ] Observable pattern sin memory leaks
- [ ] async pipe o takeUntil
- [ ] OnPush change detection
- [ ] Immutabilidad (spread operator)
- [ ] Sin lógica duplicada

---

## Próximo

Módulo 20: Formularios Template-Driven (COMPLETADO)

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
