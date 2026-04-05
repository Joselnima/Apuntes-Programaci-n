# Capítulo 18: RxJS Explicado Fácil

> **Domina los flujos de datos reactivos y aprende para qué sirven en Angular**

---

## ¿Qué es RxJS y para qué sirve?

RxJS es una librería para trabajar con **datos que cambian con el tiempo**.

- En Angular, los datos no siempre llegan de una sola vez.
- Llegan desde formularios, eventos del usuario, peticiones HTTP, websockets y temporizadores.
- RxJS te permite escribir código que responde a esos cambios de forma clara, predecible y composable.

### ¿Para qué sirve RxJS?

RxJS sirve para:

- **Manejar flujos de eventos**: clicks, inputs, scroll, actualizaciones en tiempo real.
- **Encadenar peticiones**: ejecutar una consulta cuando termina otra.
- **Combinar datos**: juntar varios streams en uno.
- **Transformar valores**: mapear, filtrar, acumular datos.
- **Manejar errores**: capturar y recuperar sin romper todo el flujo.
- **Cancelar operaciones**: abortar peticiones cuando el componente se destruye o el usuario cambia de pantalla.

### ¿Cuándo usar RxJS?

Usa RxJS cuando trabajas con:

- Formularios reactivos y `valueChanges`
- HTTP en Angular (`HttpClient` devuelve Observables)
- Streams de eventos del DOM
- Websockets y actualizaciones en tiempo real
- Estados compartidos entre componentes
- Operaciones encadenadas o condicionales

> RxJS no es obligatorio para todo, pero en Angular es la herramienta natural para manejar **la mayoría del comportamiento reactivo**.

---

## El gran cambio: de datos estáticos a datos en movimiento

### Analogía: un río frente a un cubo de agua

- Un `Promise` es un **cubo de agua**: lo llenas una vez y recibes un resultado.
- Un `Observable` es un **río**: pueden llegar muchos valores a lo largo del tiempo.
- Un `Subject` es una **puerta de agua controlada**: decides cuándo pasarán los valores.

Esta diferencia es la clave para entender cuándo usar cada herramienta.

---

## Cómo usar RxJS

### Elemento 1: Observable

Un `Observable` es una fuente de datos que puede emitir múltiples valores con el tiempo.

```typescript
import { Observable } from 'rxjs';

const myObservable = new Observable<number>(observer => {
  observer.next(1);
  observer.next(2);
  observer.next(3);
  observer.complete();
});

myObservable.subscribe({
  next: value => console.log('Valor:', value),
  error: err => console.error('Error:', err),
  complete: () => console.log('Terminado')
});
```

### ¿Para qué usar un Observable?

- Para consumir datos de una API con `HttpClient`
- Para escuchar cambios de formularios
- Para reaccionar a eventos del DOM
- Para trabajar con datos en tiempo real

> En Angular, muchas APIs ya devuelven Observables, por lo que RxJS se vuelve parte natural del flujo.

### Elemento 2: Subject

Un `Subject` es un `Observable` al que también puedes empujar valores manualmente.

```typescript
import { Subject } from 'rxjs';

const messageSubject = new Subject<string>();
messageSubject.subscribe(msg => console.log('Listener 1:', msg));
messageSubject.next('Hola');
messageSubject.subscribe(msg => console.log('Listener 2:', msg));
messageSubject.next('Mundo');
```

### ¿Para qué sirve un Subject?

- Para enviar eventos entre componentes
- Para crear streams desde acciones del usuario
- Para construir APIs de eventos personalizadas

### Elemento 3: BehaviorSubject

Un `BehaviorSubject` emite el último valor inmediato a nuevos suscriptores.

```typescript
import { BehaviorSubject } from 'rxjs';

const countSubject = new BehaviorSubject<number>(0);
countSubject.next(1);
countSubject.next(2);
countSubject.subscribe(count => console.log('Subs:', count));
```

### ¿Para qué usar BehaviorSubject?

- Para compartir estado actual con varios componentes
- Para exponer un valor inicial y siempre disponible
- Para implementar servicios de estado simples

**Ejemplo:**

```typescript
@Injectable({ providedIn: 'root' })
export class CartService {
  private cartSubject = new BehaviorSubject<Item[]>([]);
  cart$ = this.cartSubject.asObservable();

  addItem(item: Item) {
    const current = this.cartSubject.value;
    this.cartSubject.next([...current, item]);
  }
}
```

### Elemento 4: Operadores

Los operadores transforman y combinan flujos sin suscribirte manualmente en cada paso.

```typescript
import { of } from 'rxjs';
import { map, filter, take } from 'rxjs/operators';

const numbers = of(1, 2, 3, 4, 5);

numbers
  .pipe(
    filter(n => n > 2),
    map(n => n * 2),
    take(2)
  )
  .subscribe(value => {
    console.log(value); // 6, 8
  });
```

---

## Operadores esenciales y cuándo usarlos

### map: transformar valores

Transfórmalos sin cambiar la estructura del flujo.

```typescript
products$
  .pipe(
    map(products => products.map(product => product.name))
  )
  .subscribe(names => console.log(names));
```

**Uso:** extraer propiedades, convertir datos y preparar la UI.

### filter: seleccionar solo lo relevante

```typescript
products$
  .pipe(
    map(products => products.filter(p => p.price > 500))
  )
  .subscribe(expensive => console.log(expensive));
```

**Uso:** ignorar valores que no interesan.

### switchMap: cancelar la operación anterior

```typescript
searchTerm$
  .pipe(
    debounceTime(300),
    distinctUntilChanged(),
    switchMap(term => this.searchService.search(term))
  )
  .subscribe(results => this.results = results);
```

**Uso:** búsquedas en vivo, peticiones dependientes y cancelación automática.

### mergeMap: ejecutar en paralelo

```typescript
actions$
  .pipe(mergeMap(action => this.api.call(action)))
  .subscribe();
```

**Uso:** múltiples peticiones sin cancelar ninguna.

### concatMap: mantener el orden

```typescript
actions$
  .pipe(concatMap(action => this.api.call(action)))
  .subscribe();
```

**Uso:** cuando cada llamada debe terminar antes de la siguiente.

### exhaustMap: ignorar nuevas emisiones mientras la actual continúa

```typescript
clicks$
  .pipe(exhaustMap(() => this.api.save()))
  .subscribe();
```

**Uso:** prevenir doble envío de formularios.

### tap: efectos secundarios sin cambiar los datos

```typescript
products$
  .pipe(
    tap(data => console.log('Productos:', data)),
    map(products => products.map(p => p.name))
  )
  .subscribe(names => console.log(names));
```

**Uso:** logging, métricas y debugging.

### debounceTime: esperar silencio

```typescript
searchControl.valueChanges
  .pipe(debounceTime(300))
  .subscribe(term => this.search(term));
```

**Uso:** búsquedas, inputs del usuario y eventos muy frecuentes.

### distinctUntilChanged: evitar duplicados

```typescript
searchTerm$
  .pipe(distinctUntilChanged())
  .subscribe(term => this.search(term));
```

**Uso:** no volver a procesar el mismo valor.

### catchError: manejar errores sin colapsar el flujo

```typescript
this.http.get('/api/data')
  .pipe(
    catchError(error => {
      console.error('Error:', error);
      return of([]);
    })
  )
  .subscribe(data => this.data = data);
```

**Uso:** capturar excepciones y retornar un fallback.

### shareReplay: compartir la misma suscripción

```typescript
const sharedData$ = this.http.get('/api/data').pipe(shareReplay(1));
```

**Uso:** evitar peticiones duplicadas y reutilizar resultados.

---

## Casos de uso reales en Angular

### 1. HTTP con cancelación

RxJS permite cancelar la petición anterior cuando llega una nueva.

```typescript
searchTerm$
  .pipe(
    debounceTime(300),
    distinctUntilChanged(),
    switchMap(term => this.http.get<SearchResult[]>(`/api/search?q=${term}`))
  )
  .subscribe(results => this.results = results);
```

**¿Para qué sirve?**
- Evita que el backend trabaje con consultas obsoletas.
- Mejora la experiencia del usuario.

### 2. Formularios reactivos

`valueChanges` devuelve un Observable.

```typescript
this.form.get('email')!.valueChanges
  .pipe(
    debounceTime(300),
    distinctUntilChanged()
  )
  .subscribe(email => this.validateEmail(email));
```

**¿Para qué sirve?**
- Validación instantánea.
- Mostrar mensajes de error dinámicamente.

### 3. Eventos del DOM

```typescript
fromEvent(window, 'scroll')
  .pipe(throttleTime(100))
  .subscribe(() => this.onScroll());
```

**¿Para qué sirve?**
- Optimizar el manejo de scroll y resize.
- Evitar ejecuciones excesivas.

### 4. Estado compartido

Un servicio centralizado expone un stream único.

```typescript
@Injectable({ providedIn: 'root' })
export class UserService {
  private userSubject = new BehaviorSubject<User | null>(null);
  user$ = this.userSubject.asObservable();

  loadUser() {
    this.http.get<User>('/api/user')
      .subscribe(user => this.userSubject.next(user));
  }
}
```

**¿Para qué sirve?**
- Compartir estado entre componentes.
- Mantener el último valor disponible.

### 5. Websockets y datos en tiempo real

```typescript
const socket$ = new Observable<Message>(observer => {
  const socket = new WebSocket('ws://...');

  socket.onmessage = event => observer.next(JSON.parse(event.data));
  socket.onerror = err => observer.error(err);
  socket.onclose = () => observer.complete();

  return () => socket.close();
});
```

**¿Para qué sirve?**
- Mensajes en vivo.
- Actualizaciones instantáneas sin refrescar.

---

## Observables vs Promises vs Callbacks

| Feature | Callbacks | Promises | Observables |
|---------|-----------|----------|-------------|
| Múltiples valores | ❌ | ❌ | ✅ |
| Cancelable | ❌ | ❌ | ✅ |
| Lazy | ❌ | ❌ | ✅ |
| Transformable | ❌ | ✅ | ✅ |
| Error handling | Complejo | Simple | Flexible |
| Integración con Angular | Baja | Media | Alta |

### ¿Cuándo usar cada uno?

- **Callback**: solo en APIs legacy o cuando no hay otra opción.
- **Promise**: para operaciones asíncronas únicas como login, lectura de archivos o `fetch`.
- **Observable**: para flujos continuos, cancelables y combinables.

---

## Buenas prácticas en Angular con RxJS

### 1. Los servicios devuelven Observables

```typescript
getUsers(): Observable<User[]> {
  return this.http.get<User[]>('/api/users');
}
```

- No te suscribas en el servicio.
- Deja que el componente o la plantilla manejen la suscripción.

### 2. Usa `async` pipe cuando sea posible

```html
<div *ngIf="users$ | async as users">
  <p>Cantidad: {{ users.length }}</p>
</div>
```

- Evita suscripciones manuales.
- Angular desuscribe automáticamente.

### 3. Cancela suscripciones en componentes

- Usa `takeUntil`, `take(1)`, `first()` o `async`.
- Si un componente se destruye, la suscripción debe morir.

### 4. No abuses de los Subjects

- `Subject` para eventos y comunicación manual.
- `BehaviorSubject` para estado compartido.
- `ReplaySubject` cuando necesitas reemitir valores anteriores.

### 5. Mantén los side effects en `tap`

- `tap` sirve para logging, métricas o debugging.
- No alteres los datos de la cadena dentro de `tap`.

---

## Errores comunes y cuándo ocurren

### 1. Suscribirse en el servicio

**Error:** subscribes en el servicio.

```typescript
// ❌
getUsers() {
  this.http.get('/api/users').subscribe(users => this.users = users);
}
```

**Mejor:**

```typescript
getUsers(): Observable<User[]> {
  return this.http.get<User[]>('/api/users');
}
```

### 2. No cancelar suscripciones

- Usa `takeUntil`, `take(1)` o `async`.
- Evita memory leaks.

### 3. Usar `switchMap` cuando necesitas todas las respuestas

- `switchMap` cancela la anterior.
- Si necesitas recibir todas, usa `mergeMap` o `concatMap`.

### 4. No manejar errores globalmente

- Usa `catchError` en pipes.
- Usa interceptores HTTP para errores comunes.

### 5. Duplicar streams sin necesidad

- Reutiliza Observables con `shareReplay`.
- No crees el mismo stream varias veces sin necesidad.

---

## Ejemplo avanzado: búsqueda inteligente

```typescript
searchControl = new FormControl('');
results$ = this.searchControl.valueChanges.pipe(
  debounceTime(300),
  distinctUntilChanged(),
  filter(term => term.trim().length >= 2),
  switchMap(term => this.http.get<SearchResult[]>(`/api/search?q=${term}`)),
  catchError(() => of([]))
);
```

```html
<input [formControl]="searchControl" placeholder="Buscar...">
<ul>
  <li *ngFor="let item of results$ | async">{{ item.name }}</li>
</ul>
```

**¿Para qué sirve este patrón?**

- Filtrar consultas demasiado cortas.
- Cancelar búsquedas obsoletas.
- Mostrar resultados actualizados sin ruido.

---

## Checklist - RxJS profesional

- [ ] Entiendes qué problema resuelve RxJS
- [ ] Sabes cuándo usar Observable, Subject y BehaviorSubject
- [ ] Usas operadores adecuados según el caso de uso
- [ ] Manejas cancelación con `switchMap`, `takeUntil` o `exhaustMap`
- [ ] No generas memory leaks
- [ ] Prefieres `async` pipe cuando sea posible
- [ ] Manejas errores con `catchError`
- [ ] Compartes streams con `shareReplay` cuando conviene

---

## Mini práctica profunda

1. Crea un servicio de búsqueda con RxJS.
2. Agrega `debounceTime`, `distinctUntilChanged` y `switchMap`.
3. Muestra resultados en la UI con `async`.
4. Maneja errores con `catchError`.
5. Cambia a `mergeMap` y observa la diferencia.

---

## Resumen

RxJS es la herramienta de Angular para manejar datos que cambian en el tiempo.

- Usa `Observable` para flujos de datos.
- Usa `Subject` cuando necesitas emitir manualmente.
- Usa `BehaviorSubject` para estado con valor actual.
- Usa operadores para transformar y combinar.
- Usa `async` pipe para evitar suscripciones manuales.

Si entiendes **para qué sirve cada pieza** y **en qué caso usarla**, RxJS deja de ser confuso y se convierte en tu mejor aliado.

### forkJoin: Esperar múltiples peticiones

```typescript
// Como Promise.all() pero para Observables
import { forkJoin } from 'rxjs';

forkJoin([
  this.http.get('/api/users'),
  this.http.get('/api/products'),
  this.http.get('/api/categories')
])
  .subscribe(([users, products, categories]) => {
    this.users = users;
    this.products = products;
    this.categories = categories;
  });
```

### takeUntil: Desuscribirse cuando algo ocurra

```typescript
destroy$ = new Subject<void>();

ngOnInit() {
  this.products$
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

---

## Ejemplo Real: Carrito de Compras

```typescript
@Injectable()
export class CartService {
  private itemsSubject = new BehaviorSubject<Item[]>([]);
  items$ = this.itemsSubject.asObservable();
  
  // Total dinámico
  total$ = this.items$.pipe(
    map(items => items.reduce((sum, item) => 
      sum + (item.price * item.quantity), 0
    ))
  );
  
  // Cantidad de items
  itemsCount$ = this.items$.pipe(
    map(items => items.reduce((sum, item) => 
      sum + item.quantity, 0
    ))
  );
  
  addItem(item: Item) {
    const current = this.itemsSubject.value;
    this.itemsSubject.next([...current, item]);
  }
  
  removeItem(id: number) {
    const current = this.itemsSubject.value;
    const filtered = current.filter(item => item.id !== id);
    this.itemsSubject.next(filtered);
  }
  
  clear() {
    this.itemsSubject.next([]);
  }
}
```

**Componente usando CartService:**

```typescript
export class CartComponent {
  items$ = this.cartService.items$;
  total$ = this.cartService.total$;
  itemsCount$ = this.cartService.itemsCount$;
  
  constructor(private cartService: CartService) {}
  
  removeItem(id: number) {
    this.cartService.removeItem(id);
  }
}
```

**Template:**

```html
<div class="cart">
  <p>Items: {{ itemsCount$ | async }}</p>
  <p>Total: ${{ total$ | async }}</p>
  
  <ul>
    <li *ngFor="let item of (items$ | async)">
      {{ item.name }} - ${{ item.price }}
      <button (click)="removeItem(item.id)">
        Eliminar
      </button>
    </li>
  </ul>
</div>
```

---

## Memory Leaks RxJS

### ❌ Error: Suscripción infinita

```typescript
ngOnInit() {
  this.data$.subscribe(data => {
    this.data = data;
  });
  // Sin desuscribirse ❌
}
```

### ✅ Solución: takeUntil

```typescript
destroy$ = new Subject<void>();

ngOnInit() {
  this.data$
    .pipe(takeUntil(this.destroy$))
    .subscribe(data => {
      this.data = data;
    });
}

ngOnDestroy() {
  this.destroy$.next();
  this.destroy$.complete();
}
```

### ✅ Mejor: async pipe

```html
<!-- Automáticamente desuscrito -->
{{ data$ | async }}
```

---

## Comparativa: Callbacks vs Promises vs Observables

| Feature | Callbacks | Promises | Observables |
|---------|-----------|----------|------------|
| Múltiples valores | ❌ | ❌ | ✅ |
| Cancelable | ❌ | ❌ | ✅ |
| Lazy (perezoso) | ❌ | ❌ | ✅ |
| Transformable | ❌ | ✅ (then) | ✅ (operators) |
| Error handling | Complejo | Simple | Simple |
| Async/await | ❌ | ✅ | ❌ |

---

## Checklist - RxJS Profesional

- [ ] Entiendes Observable vs Subject vs BehaviorSubject
- [ ] Usas map, filter, switchMap, tap regularmente
- [ ] Sin memory leaks (takeUntil o async pipe)
- [ ] Combinas múltiples streams (combineLatest, forkJoin)
- [ ] Manejas errores con catchError
- [ ] Usas debounceTime en búsquedas
- [ ] Servicios retornan Observables, NO promesas
- [ ] Documentación de operadores complejos

---

## Próximo

Módulo 19: Estado Local - Gestiona datos complejos

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
