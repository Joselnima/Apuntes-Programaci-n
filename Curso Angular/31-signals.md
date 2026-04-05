# Capítulo 31: Signals en Angular - La Revolución Reactiva

## La Analogía del Sistema Nervioso Humano

Imagina que tu aplicación Angular es un organismo vivo. Los **Signals** son el sistema nervioso central: transmiten información instantáneamente, coordinan respuestas automáticas, y mantienen todo sincronizado sin esfuerzo consciente.

En el organismo digital:
- **Señales (Signals)**: Las neuronas que transmiten información
- **Computados (Computed)**: Los reflejos automáticos del cuerpo
- **Efectos (Effects)**: Las respuestas musculares a estímulos
- **Linked Signals**: Las sinapsis que conectan diferentes partes del cerebro

En este capítulo, te convertirás en el neurocirujano de tus aplicaciones. Aprenderemos a diseñar sistemas nerviosos digitales que respondan instantáneamente, se curen automáticamente, y evolucionen con el tiempo.

---

## ¿Por Qué Signals Representan una Revolución?

### El Problema del Estado en Angular Clásico

```typescript
// ❌ ANTIGUO: Estado disperso y difícil de rastrear
@Component({...})
export class UserProfileComponent implements OnInit, OnDestroy {
  user: User | null = null;
  loading = false;
  error: string | null = null;

  private subscription?: Subscription;

  ngOnInit() {
    this.loading = true;
    this.userService.getCurrentUser().subscribe({
      next: (user) => {
        this.user = user;
        this.loading = false;
      },
      error: (err) => {
        this.error = err.message;
        this.loading = false;
      }
    });
  }

  ngOnDestroy() {
    this.subscription?.unsubscribe();
  }

  // ¿Cómo saber si todos los estados están sincronizados?
  // ¿Dónde está el estado realmente?
  // ¿Qué pasa si múltiples componentes necesitan este estado?
}
```

### La Solución: Signals como Sistema Nervioso Central

```typescript
// ✅ NUEVO: Estado centralizado y automáticamente reactivo
@Injectable({
  providedIn: 'root'
})
export class UserState {
  private _user = signal<User | null>(null);
  private _loading = signal(false);
  private _error = signal<string | null>(null);

  // Estado público de solo lectura
  readonly user = this._user.asReadonly();
  readonly loading = this._loading.asReadonly();
  readonly error = this._error.asReadonly();

  // Estado computado
  readonly isAuthenticated = computed(() => !!this._user());
  readonly displayName = computed(() =>
    this._user() ? `${this._user()!.firstName} ${this._user()!.lastName}` : 'Invitado'
  );

  // Acciones
  loadUser() {
    this._loading.set(true);
    this._error.set(null);

    this.http.get<User>('/api/user').subscribe({
      next: (user) => {
        this._user.set(user);
        this._loading.set(false);
      },
      error: (err) => {
        this._error.set(err.message);
        this._loading.set(false);
      }
    });
  }

  logout() {
    this._user.set(null);
    this._error.set(null);
  }
}

// Componente ultra-simple
@Component({...})
export class UserProfileComponent {
  readonly user = inject(UserState).user;
  readonly loading = inject(UserState).loading;
  readonly error = inject(UserState).error;
  readonly displayName = inject(UserState).displayName;
  readonly isAuthenticated = inject(UserState).isAuthenticated;

  private userState = inject(UserState);

  ngOnInit() {
    this.userState.loadUser();
  }

  logout() {
    this.userState.logout();
  }
}
```

---

## Arquitectura Fundamental de Signals

### 1. Signal: La Unidad Básica de Estado

```typescript
// Tipos de signals
const count = signal(0);                    // WritableSignal<number>
const user = signal<User | null>(null);     // WritableSignal<User | null>
const items = signal<CartItem[]>([]);       // WritableSignal<CartItem[]>

// Solo lectura (para APIs públicas)
const readonlyCount = count.asReadonly();   // ReadonlySignal<number>

// Tipos genéricos avanzados
const config = signal<AppConfig>({
  theme: 'light',
  language: 'es',
  notifications: true
});
```

### 2. Computed: Cálculos Automáticos y Memoizados

```typescript
// ✅ Computados simples
const doubled = computed(() => count() * 2);
const isEven = computed(() => count() % 2 === 0);

// ✅ Computados complejos con objetos
const userSummary = computed(() => {
  const user = user();
  if (!user) return null;

  return {
    fullName: `${user.firstName} ${user.lastName}`,
    initials: `${user.firstName[0]}${user.lastName[0]}`,
    isAdult: user.age >= 18,
    membershipLevel: user.points > 1000 ? 'Premium' : 'Basic'
  };
});

// ✅ Computados anidados
const premiumBenefits = computed(() => {
  const summary = userSummary();
  if (!summary || summary.membershipLevel !== 'Premium') return [];

  return ['Priority Support', 'Free Shipping', 'Exclusive Discounts'];
});
```

### 3. Effect: Reacciones Automáticas a Cambios

```typescript
// ✅ Efectos básicos
effect(() => {
  console.log('Count changed:', count());
});

// ✅ Efectos con cleanup automático
effect((onCleanup) => {
  const subscription = someObservable.subscribe();

  onCleanup(() => {
    subscription.unsubscribe();
  });
});

// ✅ Efectos condicionales
effect(() => {
  const currentUser = user();
  const isLoading = loading();

  if (currentUser && !isLoading) {
    // Solo ejecutar cuando hay usuario y no está cargando
    analytics.track('user_loaded', { userId: currentUser.id });
  }
});
```

---

## Linked Signals: Conexiones Inteligentes Entre Estados

### ¿Qué son los Linked Signals?

Los **Linked Signals** son conexiones automáticas entre diferentes partes del estado. Cuando un signal cambia, automáticamente actualiza otros signals relacionados, creando una red de dependencias reactivas.

```typescript
// ❌ SIN LINKS: Estado desconectado
const firstName = signal('John');
const lastName = signal('Doe');
const fullName = signal(''); // Necesita actualización manual

// Actualización manual tediosa
firstName.set('Jane');
fullName.set(`${firstName()} ${lastName()}`); // ¡Fácil olvidar!

// ❌ SIN LINKS: Componentes duplicados
@Component({...})
export class UserFormComponent {
  firstName = signal('John');
  lastName = signal('Doe');

  // En otro componente...
}

@Component({...})
export class UserDisplayComponent {
  firstName = signal('John'); // ¡Duplicado!
  lastName = signal('Doe');   // ¡Duplicado!
}
```

```typescript
// ✅ CON LINKED SIGNALS: Estado conectado automáticamente
@Injectable({
  providedIn: 'root'
})
export class UserState {
  // Señales base
  readonly firstName = signal('John');
  readonly lastName = signal('Doe');

  // Señales computadas (linked automáticamente)
  readonly fullName = computed(() =>
    `${this.firstName()} ${this.lastName()}`
  );

  readonly initials = computed(() =>
    `${this.firstName()[0]}${this.lastName()[0]}`.toUpperCase()
  );

  readonly displayName = computed(() => {
    const full = this.fullName();
    return full.length > 20 ? this.initials() : full;
  });

  // Acciones que actualizan múltiples signals
  updateName(first: string, last: string) {
    this.firstName.set(first);
    this.lastName.set(last);
    // fullName, initials, y displayName se actualizan automáticamente
  }
}

// ✅ USO: Un solo estado compartido
@Component({...})
export class UserFormComponent {
  private userState = inject(UserState);

  // Todos los componentes usan el mismo estado
  readonly firstName = this.userState.firstName;
  readonly lastName = this.userState.lastName;
  readonly fullName = this.userState.fullName; // Se actualiza solo

  updateName() {
    this.userState.updateName('Jane', 'Smith');
    // ¡Todo se actualiza automáticamente!
  }
}
```

### Patrón: State Machines con Linked Signals

```typescript
// Estado de autenticación como máquina de estados
@Injectable({
  providedIn: 'root'
})
export class AuthState {
  // Estados base
  private _user = signal<User | null>(null);
  private _loading = signal(false);
  private _error = signal<string | null>(null);

  // Estados computados (linked)
  readonly user = this._user.asReadonly();
  readonly loading = this._loading.asReadonly();
  readonly error = this._error.asReadonly();

  readonly isAuthenticated = computed(() => !!this._user());
  readonly isLoading = computed(() => this._loading());
  readonly hasError = computed(() => !!this._error());

  readonly authStatus = computed(() => {
    if (this._loading()) return 'loading';
    if (this._error()) return 'error';
    if (this._user()) return 'authenticated';
    return 'unauthenticated';
  });

  readonly canAccessAdmin = computed(() => {
    const user = this._user();
    return user?.role === 'admin' && !this._error();
  });

  // Acciones que cambian el estado
  async login(credentials: LoginCredentials) {
    this._loading.set(true);
    this._error.set(null);

    try {
      const user = await this.authService.login(credentials);
      this._user.set(user);
    } catch (error) {
      this._error.set(error.message);
    } finally {
      this._loading.set(false);
    }
  }

  logout() {
    this._user.set(null);
    this._error.set(null);
    this.authService.logout();
  }

  clearError() {
    this._error.set(null);
  }
}
```

### Patrón: Form State con Linked Signals

```typescript
// Estado de formulario inteligente
@Injectable()
export class ContactFormState {
  // Campos del formulario
  readonly name = signal('');
  readonly email = signal('');
  readonly message = signal('');

  // Estados computados
  readonly isNameValid = computed(() => this.name().trim().length >= 2);
  readonly isEmailValid = computed(() =>
    /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(this.email())
  );
  readonly isMessageValid = computed(() => this.message().trim().length >= 10);

  readonly isFormValid = computed(() =>
    this.isNameValid() && this.isEmailValid() && this.isMessageValid()
  );

  readonly fieldErrors = computed(() => ({
    name: this.name() && !this.isNameValid() ? 'Nombre muy corto' : null,
    email: this.email() && !this.isEmailValid() ? 'Email inválido' : null,
    message: this.message() && !this.isMessageValid() ? 'Mensaje muy corto' : null
  }));

  readonly touchedFields = signal<Set<string>>(new Set());

  readonly showErrors = computed(() => {
    const touched = this.touchedFields();
    const errors = this.fieldErrors();
    return Object.keys(errors).reduce((acc, field) => ({
      ...acc,
      [field]: touched.has(field) && errors[field]
    }), {});
  });

  // Acciones
  updateField(field: string, value: string) {
    switch (field) {
      case 'name': this.name.set(value); break;
      case 'email': this.email.set(value); break;
      case 'message': this.message.set(value); break;
    }
  }

  markFieldAsTouched(field: string) {
    this.touchedFields.update(touched => new Set([...touched, field]));
  }

  async submit() {
    if (!this.isFormValid()) {
      // Marcar todos los campos como touched para mostrar errores
      this.touchedFields.set(new Set(['name', 'email', 'message']));
      return;
    }

    const formData = {
      name: this.name(),
      email: this.email(),
      message: this.message()
    };

    await this.contactService.submit(formData);
    this.reset();
  }

  reset() {
    this.name.set('');
    this.email.set('');
    this.message.set('');
    this.touchedFields.set(new Set());
  }
}
```

---

## Signals en Formularios Reactivos

### Patrón: Formularios con Signals

```typescript
// ✅ FORMULARIO REACTIVO CON SIGNALS
@Component({
  selector: 'app-reactive-form',
  template: `
    <form [formGroup]="form" (ngSubmit)="onSubmit()">
      <mat-form-field>
        <input matInput
               formControlName="name"
               placeholder="Nombre"
               [class.error]="showErrors().name">
        <mat-error *ngIf="showErrors().name">
          {{ showErrors().name }}
        </mat-error>
      </mat-form-field>

      <mat-form-field>
        <input matInput
               type="email"
               formControlName="email"
               placeholder="Email"
               [class.error]="showErrors().email">
        <mat-error *ngIf="showErrors().email">
          {{ showErrors().email }}
        </mat-error>
      </mat-form-field>

      <button mat-raised-button
              color="primary"
              type="submit"
              [disabled]="!isFormValid()">
        Enviar
      </button>
    </form>
  `
})
export class ReactiveFormComponent {
  private fb = inject(FormBuilder);
  private formState = inject(ContactFormState);

  // Formulario reactivo conectado a signals
  readonly form = this.fb.group({
    name: ['', [Validators.required, Validators.minLength(2)]],
    email: ['', [Validators.required, Validators.email]],
    message: ['', [Validators.required, Validators.minLength(10)]]
  });

  // Estados computados desde signals
  readonly isFormValid = this.formState.isFormValid;
  readonly showErrors = this.formState.showErrors;

  ngOnInit() {
    // Sincronizar FormGroup con Signals
    this.syncFormToSignals();
    this.syncSignalsToForm();
  }

  private syncFormToSignals() {
    // Cuando el formulario cambia, actualizar signals
    this.form.valueChanges.subscribe(value => {
      if (value.name !== undefined) this.formState.name.set(value.name);
      if (value.email !== undefined) this.formState.email.set(value.email);
      if (value.message !== undefined) this.formState.message.set(value.message);
    });
  }

  private syncSignalsToForm() {
    // Cuando los signals cambian, actualizar formulario
    effect(() => {
      const name = this.formState.name();
      const email = this.formState.email();
      const message = this.formState.message();

      this.form.patchValue({ name, email, message }, { emitEvent: false });
    });
  }

  onSubmit() {
    if (this.form.valid) {
      this.formState.submit();
    } else {
      // Marcar campos como touched
      Object.keys(this.form.controls).forEach(key => {
        this.form.get(key)?.markAsTouched();
        this.formState.markFieldAsTouched(key);
      });
    }
  }
}
```

### Patrón: Formularios Template-Driven con Signals

```typescript
// ✅ TEMPLATE-DRIVEN CON SIGNALS
@Component({
  selector: 'app-template-form',
  template: `
    <form #form="ngForm" (ngSubmit)="onSubmit(form)">
      <div class="form-field">
        <input type="text"
               name="name"
               [(ngModel)]="name"
               required
               minlength="2"
               #nameField="ngModel"
               [class.error]="nameField.invalid && nameField.touched">

        <div class="error-message" *ngIf="nameField.invalid && nameField.touched">
          <span *ngIf="nameField.errors?.['required']">Nombre requerido</span>
          <span *ngIf="nameField.errors?.['minlength']">Mínimo 2 caracteres</span>
        </div>
      </div>

      <div class="form-field">
        <input type="email"
               name="email"
               [(ngModel)]="email"
               required
               email
               #emailField="ngModel"
               [class.error]="emailField.invalid && emailField.touched">

        <div class="error-message" *ngIf="emailField.invalid && emailField.touched">
          <span *ngIf="emailField.errors?.['required']">Email requerido</span>
          <span *ngIf="emailField.errors?.['email']">Email inválido</span>
        </div>
      </div>

      <button type="submit"
              [disabled]="!isFormValid()"
              [class.submitting]="isSubmitting()">
        {{ isSubmitting() ? 'Enviando...' : 'Enviar' }}
      </button>
    </form>
  `
})
export class TemplateFormComponent {
  private formState = inject(ContactFormState);

  // Signals conectados al formulario
  readonly name = this.formState.name;
  readonly email = this.formState.email;
  readonly isFormValid = this.formState.isFormValid;
  readonly isSubmitting = signal(false);

  onSubmit(form: NgForm) {
    if (form.valid) {
      this.isSubmitting.set(true);
      this.formState.submit()
        .finally(() => this.isSubmitting.set(false));
    } else {
      // Marcar todos los campos como touched
      Object.values(form.controls).forEach(control => {
        control.markAsTouched();
      });
    }
  }
}
```

### Patrón: Validación Asíncrona con Signals

```typescript
// Servicio de validación con signals
@Injectable({
  providedIn: 'root'
})
export class ValidationService {
  // Cache de validaciones para performance
  private emailCache = new Map<string, boolean>();

  isEmailAvailable(email: string): Observable<boolean> {
    if (this.emailCache.has(email)) {
      return of(this.emailCache.get(email)!);
    }

    return this.http.get<boolean>(`/api/validate-email?email=${email}`).pipe(
      tap(result => this.emailCache.set(email, result))
    );
  }

  isUsernameAvailable(username: string): Observable<boolean> {
    return this.http.get<boolean>(`/api/validate-username?username=${username}`);
  }
}

// Estado de formulario con validación asíncrona
@Injectable()
export class RegistrationFormState {
  readonly email = signal('');
  readonly username = signal('');
  readonly password = signal('');

  // Validaciones síncronas
  readonly emailValid = computed(() =>
    /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(this.email())
  );

  readonly usernameValid = computed(() =>
    this.username().length >= 3 && /^[a-zA-Z0-9_]+$/.test(this.username())
  );

  readonly passwordValid = computed(() => {
    const pwd = this.password();
    return pwd.length >= 8 &&
           /[A-Z]/.test(pwd) &&
           /[a-z]/.test(pwd) &&
           /[0-9]/.test(pwd);
  });

  // Validaciones asíncronas con signals
  readonly emailAvailable = signal<boolean | null>(null);
  readonly usernameAvailable = signal<boolean | null>(null);

  readonly emailChecking = signal(false);
  readonly usernameChecking = signal(false);

  // Estado general
  readonly isFormValid = computed(() =>
    this.emailValid() &&
    this.usernameValid() &&
    this.passwordValid() &&
    this.emailAvailable() !== false &&
    this.usernameAvailable() !== false
  );

  readonly canSubmit = computed(() =>
    this.isFormValid() &&
    !this.emailChecking() &&
    !this.usernameChecking()
  );

  constructor(private validationService: ValidationService) {
    // Validación asíncrona automática
    effect(() => {
      const email = this.email();
      if (email && this.emailValid()) {
        this.emailChecking.set(true);
        this.validationService.isEmailAvailable(email).subscribe({
          next: (available) => {
            this.emailAvailable.set(available);
            this.emailChecking.set(false);
          },
          error: () => {
            this.emailAvailable.set(null);
            this.emailChecking.set(false);
          }
        });
      } else {
        this.emailAvailable.set(null);
      }
    });

    effect(() => {
      const username = this.username();
      if (username && this.usernameValid()) {
        this.usernameChecking.set(true);
        this.validationService.isUsernameAvailable(username).subscribe({
          next: (available) => {
            this.usernameAvailable.set(available);
            this.usernameChecking.set(false);
          },
          error: () => {
            this.usernameAvailable.set(null);
            this.usernameChecking.set(false);
          }
        });
      } else {
        this.usernameAvailable.set(null);
      }
    });
  }

  async submit() {
    if (!this.canSubmit()) return;

    const registrationData = {
      email: this.email(),
      username: this.username(),
      password: this.password()
    };

    await this.registrationService.register(registrationData);
    this.reset();
  }

  reset() {
    this.email.set('');
    this.username.set('');
    this.password.set('');
    this.emailAvailable.set(null);
    this.usernameAvailable.set(null);
  }
}
```

---

## Debugging y Troubleshooting de Signals

### Herramientas de Debugging para Signals

```typescript
// 1. LOGGING AVANZADO
@Injectable({
  providedIn: 'root'
})
export class SignalDebugger {
  private logs: SignalLog[] = [];

  watch<T>(signal: Signal<T>, name: string): Signal<T> {
    effect(() => {
      const value = signal();
      const log: SignalLog = {
        name,
        value,
        timestamp: Date.now(),
        stack: new Error().stack
      };
      this.logs.push(log);
      console.log(`🔍 ${name}:`, value);
    });
    return signal;
  }

  getLogs(name?: string): SignalLog[] {
    return name
      ? this.logs.filter(log => log.name === name)
      : [...this.logs];
  }

  clearLogs() {
    this.logs = [];
  }
}

// Uso
export class DebugComponent {
  private debugger = inject(SignalDebugger);

  count = this.debugger.watch(signal(0), 'count');
  doubled = this.debugger.watch(
    computed(() => this.count() * 2),
    'doubled'
  );
}
```

### Troubleshooting: Errores Comunes con Signals

#### Error 1: "Cannot read property of undefined"

```typescript
// ❌ ERROR: Acceso a propiedad sin verificar
const userName = computed(() => user().name); // Crashea si user es null

// ✅ SOLUCIÓN: Verificación segura
const userName = computed(() => user()?.name ?? 'Unknown');

// ✅ SOLUCIÓN: Computed condicional
const userName = computed(() => {
  const currentUser = user();
  return currentUser ? currentUser.name : 'Unknown';
});
```

#### Error 2: "Maximum call stack size exceeded"

```typescript
// ❌ ERROR: Dependencias circulares
const a = signal(1);
const b = computed(() => a() + 1);
const c = computed(() => b() + 1);
const d = computed(() => c() + a()); // ¡Dependencia circular!

// ✅ SOLUCIÓN: Reestructurar dependencias
const a = signal(1);
const b = computed(() => a() + 1);
const c = computed(() => b() + 1);
const d = computed(() => c() + 1); // Sin circularidad
```

#### Error 3: Effects que no se ejecutan

```typescript
// ❌ ERROR: Effect con dependencias implícitas
effect(() => {
  console.log('User changed');
  // El effect no sabe que depende de 'user'
});

// ✅ SOLUCIÓN: Acceder explícitamente a las dependencias
effect(() => {
  const currentUser = user(); // Ahora Angular sabe la dependencia
  console.log('User changed:', currentUser);
});
```

#### Error 4: Memory Leaks en Effects

```typescript
// ❌ ERROR: Effect sin cleanup
effect(() => {
  const subscription = interval(1000).subscribe(console.log);
  // ¡La suscripción nunca se cancela!
});

// ✅ SOLUCIÓN: Cleanup function
effect((onCleanup) => {
  const subscription = interval(1000).subscribe(console.log);

  onCleanup(() => {
    subscription.unsubscribe();
  });
});
```

### Debugging de Dependency Injection con Signals

```typescript
// Servicio con signals para debugging
@Injectable({
  providedIn: 'root'
})
export class DependencyDebugger {
  private injectionHistory = signal<Map<string, InjectionInfo[]>>(new Map());

  trackInjection<T>(token: Type<T> | InjectionToken<T>, instance: T, context: string) {
    const key = this.getTokenKey(token);
    const info: InjectionInfo = {
      token: key,
      instance,
      context,
      timestamp: Date.now(),
      stack: new Error().stack
    };

    this.injectionHistory.update(history => {
      const existing = history.get(key) || [];
      existing.push(info);
      return new Map(history.set(key, existing));
    });

    console.log(`🔧 Injected ${key} in ${context}`);
  }

  getInjectionHistory(token?: string): InjectionInfo[] {
    const history = this.injectionHistory();
    if (token) {
      return history.get(token) || [];
    }
    return Array.from(history.values()).flat();
  }

  detectCircularDependencies(): CircularDependency[] {
    const history = this.injectionHistory();
    const circular: CircularDependency[] = [];

    // Algoritmo para detectar dependencias circulares
    // (Implementación simplificada)
    return circular;
  }

  private getTokenKey(token: Type<any> | InjectionToken<any>): string {
    return typeof token === 'function' ? token.name : token.toString();
  }
}

// Provider con debugging
export function debugProvider<T>(
  token: Type<T> | InjectionToken<T>,
  factory: () => T,
  context: string
) {
  return {
    provide: token,
    useFactory: () => {
      const instance = factory();
      inject(DependencyDebugger).trackInjection(token, instance, context);
      return instance;
    }
  };
}

// Uso
@NgModule({
  providers: [
    debugProvider(UserService, () => new UserService(), 'AppModule'),
    debugProvider(AuthService, () => new AuthService(), 'AppModule')
  ]
})
export class AppModule {}
```

---

## Proyecto Práctico: Sistema de Gestión de Tareas con Signals

### Arquitectura del Sistema

```typescript
// types/task.types.ts
export interface Task {
  id: string;
  title: string;
  description: string;
  completed: boolean;
  priority: 'low' | 'medium' | 'high';
  dueDate?: Date;
  tags: string[];
  createdAt: Date;
  updatedAt: Date;
}

export interface TaskFilter {
  search: string;
  completed?: boolean;
  priority?: Task['priority'];
  tags: string[];
}

export interface TaskStats {
  total: number;
  completed: number;
  pending: number;
  overdue: number;
  byPriority: Record<Task['priority'], number>;
}
```

### Task State Management

```typescript
// services/task.state.ts
@Injectable({
  providedIn: 'root'
})
export class TaskState {
  // Estado base
  private _tasks = signal<Task[]>([]);
  private _loading = signal(false);
  private _error = signal<string | null>(null);
  private _filter = signal<TaskFilter>({
    search: '',
    tags: []
  });

  // Estado público
  readonly tasks = this._tasks.asReadonly();
  readonly loading = this._loading.asReadonly();
  readonly error = this._error.asReadonly();
  readonly filter = this._filter.asReadonly();

  // Estado computado
  readonly filteredTasks = computed(() => {
    const tasks = this._tasks();
    const filter = this._filter();

    return tasks.filter(task => {
      // Filtro de búsqueda
      if (filter.search) {
        const search = filter.search.toLowerCase();
        const matchesTitle = task.title.toLowerCase().includes(search);
        const matchesDesc = task.description.toLowerCase().includes(search);
        if (!matchesTitle && !matchesDesc) return false;
      }

      // Filtro de completado
      if (filter.completed !== undefined && task.completed !== filter.completed) {
        return false;
      }

      // Filtro de prioridad
      if (filter.priority && task.priority !== filter.priority) {
        return false;
      }

      // Filtro de tags
      if (filter.tags.length > 0) {
        const hasMatchingTag = filter.tags.some(tag =>
          task.tags.includes(tag)
        );
        if (!hasMatchingTag) return false;
      }

      return true;
    });
  });

  readonly stats = computed((): TaskStats => {
    const tasks = this._tasks();
    const now = new Date();

    const completed = tasks.filter(t => t.completed).length;
    const pending = tasks.length - completed;
    const overdue = tasks.filter(t =>
      !t.completed && t.dueDate && t.dueDate < now
    ).length;

    const byPriority = tasks.reduce((acc, task) => {
      acc[task.priority] = (acc[task.priority] || 0) + 1;
      return acc;
    }, {} as Record<Task['priority'], number>);

    return {
      total: tasks.length,
      completed,
      pending,
      overdue,
      byPriority
    };
  });

  readonly hasOverdueTasks = computed(() =>
    this.stats().overdue > 0
  );

  readonly completionRate = computed(() => {
    const stats = this.stats();
    return stats.total > 0 ? (stats.completed / stats.total) * 100 : 0;
  });

  // Acciones
  async loadTasks() {
    this._loading.set(true);
    this._error.set(null);

    try {
      const tasks = await this.taskService.getTasks();
      this._tasks.set(tasks);
    } catch (error) {
      this._error.set(error.message);
    } finally {
      this._loading.set(false);
    }
  }

  async addTask(taskData: Omit<Task, 'id' | 'createdAt' | 'updatedAt'>) {
    try {
      const newTask = await this.taskService.createTask(taskData);
      this._tasks.update(tasks => [...tasks, newTask]);
    } catch (error) {
      this._error.set(error.message);
    }
  }

  async updateTask(id: string, updates: Partial<Task>) {
    try {
      const updatedTask = await this.taskService.updateTask(id, updates);
      this._tasks.update(tasks =>
        tasks.map(task =>
          task.id === id ? { ...task, ...updatedTask } : task
        )
      );
    } catch (error) {
      this._error.set(error.message);
    }
  }

  async deleteTask(id: string) {
    try {
      await this.taskService.deleteTask(id);
      this._tasks.update(tasks => tasks.filter(task => task.id !== id));
    } catch (error) {
      this._error.set(error.message);
    }
  }

  updateFilter(newFilter: Partial<TaskFilter>) {
    this._filter.update(current => ({ ...current, ...newFilter }));
  }

  clearFilter() {
    this._filter.set({ search: '', tags: [] });
  }

  // Linked signals para tareas relacionadas
  getTaskById(id: string) {
    return computed(() => this._tasks().find(task => task.id === id));
  }

  getTasksByTag(tag: string) {
    return computed(() =>
      this._tasks().filter(task => task.tags.includes(tag))
    );
  }

  getTasksDueToday() {
    return computed(() => {
      const today = new Date();
      today.setHours(0, 0, 0, 0);
      const tomorrow = new Date(today);
      tomorrow.setDate(tomorrow.getDate() + 1);

      return this._tasks().filter(task =>
        task.dueDate && task.dueDate >= today && task.dueDate < tomorrow
      );
    });
  }
}
```

### Componentes con Signals

```typescript
// components/task-list.component.ts
@Component({
  selector: 'app-task-list',
  template: `
    <div class="task-list">
      <!-- Filtros -->
      <app-task-filters
        [filter]="taskState.filter()"
        (filterChange)="onFilterChange($event)">
      </app-task-filters>

      <!-- Estadísticas -->
      <app-task-stats [stats]="taskState.stats()"></app-task-stats>

      <!-- Lista de tareas -->
      <div class="tasks" *ngIf="filteredTasks().length > 0; else noTasks">
        <app-task-item
          *ngFor="let task of filteredTasks(); trackBy: trackByTaskId"
          [task]="task"
          (toggle)="onToggleTask($event)"
          (delete)="onDeleteTask($event)">
        </app-task-item>
      </div>

      <ng-template #noTasks>
        <div class="no-tasks">
          <p>{{ hasActiveFilter() ? 'No hay tareas que coincidan con el filtro' : 'No hay tareas' }}</p>
          <button mat-button (click)="clearFilter()">Limpiar filtros</button>
        </div>
      </ng-template>

      <!-- Loading y error -->
      <div *ngIf="taskState.loading()" class="loading">
        <mat-spinner diameter="40"></mat-spinner>
        <p>Cargando tareas...</p>
      </div>

      <div *ngIf="taskState.error()" class="error">
        <mat-icon color="warn">error</mat-icon>
        <p>{{ taskState.error() }}</p>
        <button mat-button (click)="retryLoad()">Reintentar</button>
      </div>
    </div>
  `
})
export class TaskListComponent {
  private taskState = inject(TaskState);

  readonly filteredTasks = this.taskState.filteredTasks;
  readonly hasActiveFilter = computed(() => {
    const filter = this.taskState.filter();
    return !!(filter.search || filter.completed !== undefined ||
             filter.priority || filter.tags.length > 0);
  });

  ngOnInit() {
    this.taskState.loadTasks();
  }

  onFilterChange(newFilter: Partial<TaskFilter>) {
    this.taskState.updateFilter(newFilter);
  }

  onToggleTask(task: Task) {
    this.taskState.updateTask(task.id, { completed: !task.completed });
  }

  onDeleteTask(task: Task) {
    if (confirm('¿Eliminar tarea?')) {
      this.taskState.deleteTask(task.id);
    }
  }

  clearFilter() {
    this.taskState.clearFilter();
  }

  retryLoad() {
    this.taskState.loadTasks();
  }

  trackByTaskId(index: number, task: Task): string {
    return task.id;
  }
}
```

### Formulario de Tareas con Signals

```typescript
// components/task-form.component.ts
@Component({
  selector: 'app-task-form',
  template: `
    <form [formGroup]="taskForm" (ngSubmit)="onSubmit()" class="task-form">
      <h2>{{ isEditing ? 'Editar Tarea' : 'Nueva Tarea' }}</h2>

      <mat-form-field>
        <input matInput
               formControlName="title"
               placeholder="Título de la tarea"
               [class.error]="titleField?.invalid && titleField?.touched">
        <mat-error *ngIf="titleField?.invalid && titleField?.touched">
          <span *ngIf="titleField?.errors?.['required']">Título requerido</span>
          <span *ngIf="titleField?.errors?.['minlength']">Mínimo 3 caracteres</span>
        </mat-error>
      </mat-form-field>

      <mat-form-field>
        <textarea matInput
                  formControlName="description"
                  placeholder="Descripción"
                  rows="3"></textarea>
      </mat-form-field>

      <mat-form-field>
        <mat-select formControlName="priority">
          <mat-option value="low">Baja</mat-option>
          <mat-option value="medium">Media</mat-option>
          <mat-option value="high">Alta</mat-option>
        </mat-select>
      </mat-form-field>

      <mat-form-field>
        <input matInput
               [matDatepicker]="dueDatePicker"
               formControlName="dueDate"
               placeholder="Fecha límite">
        <mat-datepicker-toggle matSuffix [for]="dueDatePicker"></mat-datepicker-toggle>
        <mat-datepicker #dueDatePicker></mat-datepicker>
      </mat-form-field>

      <mat-form-field>
        <mat-chip-grid #chipGrid>
          <mat-chip-row
            *ngFor="let tag of selectedTags()"
            (removed)="removeTag(tag)">
            {{ tag }}
            <button matChipRemove>
              <mat-icon>cancel</mat-icon>
            </button>
          </mat-chip-row>
          <input placeholder="Etiquetas..."
                 [matChipInputFor]="chipGrid"
                 (matChipInputTokenEnd)="addTag($event)">
        </mat-chip-grid>
      </mat-form-field>

      <div class="form-actions">
        <button mat-button type="button" (click)="onCancel()">
          Cancelar
        </button>
        <button mat-raised-button
                color="primary"
                type="submit"
                [disabled]="!taskForm.valid || isSubmitting()">
          {{ isSubmitting() ? 'Guardando...' : (isEditing ? 'Actualizar' : 'Crear') }}
        </button>
      </div>
    </form>
  `,
  styles: [`
    .task-form {
      display: flex;
      flex-direction: column;
      gap: 1rem;
      max-width: 500px;
    }

    .form-actions {
      display: flex;
      justify-content: flex-end;
      gap: 0.5rem;
    }
  `]
})
export class TaskFormComponent implements OnInit {
  private fb = inject(FormBuilder);
  private taskState = inject(TaskState);

  @Input() taskToEdit: Task | null = null;
  @Output() saved = new EventEmitter<void>();
  @Output() cancelled = new EventEmitter<void>();

  readonly taskForm = this.fb.group({
    title: ['', [Validators.required, Validators.minLength(3)]],
    description: [''],
    priority: ['medium'],
    dueDate: [null],
    tags: [[]]
  });

  readonly selectedTags = signal<string[]>([]);
  readonly isSubmitting = signal(false);
  readonly isEditing = computed(() => !!this.taskToEdit);

  get titleField() {
    return this.taskForm.get('title');
  }

  ngOnInit() {
    if (this.taskToEdit) {
      this.taskForm.patchValue({
        title: this.taskToEdit.title,
        description: this.taskToEdit.description,
        priority: this.taskToEdit.priority,
        dueDate: this.taskToEdit.dueDate,
        tags: this.taskToEdit.tags
      });
      this.selectedTags.set([...this.taskToEdit.tags]);
    }

    // Sincronizar tags del formulario con signal
    this.taskForm.get('tags')?.valueChanges.subscribe(tags => {
      this.selectedTags.set(tags || []);
    });
  }

  addTag(event: MatChipInputEvent) {
    const value = (event.value || '').trim();
    if (value) {
      const currentTags = this.selectedTags();
      if (!currentTags.includes(value)) {
        const newTags = [...currentTags, value];
        this.selectedTags.set(newTags);
        this.taskForm.patchValue({ tags: newTags });
      }
    }
    event.chipInput!.clear();
  }

  removeTag(tag: string) {
    const newTags = this.selectedTags().filter(t => t !== tag);
    this.selectedTags.set(newTags);
    this.taskForm.patchValue({ tags: newTags });
  }

  async onSubmit() {
    if (!this.taskForm.valid) return;

    this.isSubmitting.set(true);

    try {
      const formValue = this.taskForm.value;
      const taskData = {
        title: formValue.title!,
        description: formValue.description || '',
        priority: formValue.priority as Task['priority'],
        dueDate: formValue.dueDate,
        tags: this.selectedTags(),
        completed: false
      };

      if (this.taskToEdit) {
        await this.taskState.updateTask(this.taskToEdit.id, taskData);
      } else {
        await this.taskState.addTask(taskData);
      }

      this.saved.emit();
      if (!this.taskToEdit) {
        this.taskForm.reset();
        this.selectedTags.set([]);
      }
    } catch (error) {
      console.error('Error saving task:', error);
    } finally {
      this.isSubmitting.set(false);
    }
  }

  onCancel() {
    this.cancelled.emit();
  }
}
```

---

## Checklist de Signals Mastery

### Fundamentos de Signals
- [ ] **Signal()**: Crear y actualizar señales básicas
- [ ] **Computed()**: Cálculos automáticos y memoizados
- [ ] **Effect()**: Reacciones automáticas a cambios
- [ ] **asReadonly()**: Exponer señales de solo lectura

### Linked Signals
- [ ] **Estado conectado**: Múltiples signals relacionados automáticamente
- [ ] **State machines**: Estados complejos con transiciones automáticas
- [ ] **Form state**: Estados de formulario inteligentes
- [ ] **Dependency chains**: Cadenas de dependencias complejas

### Signals en Formularios
- [ ] **Reactive Forms**: Formularios reactivos con signals
- [ ] **Template Forms**: Formularios template-driven con signals
- [ ] **Validación asíncrona**: Validaciones en tiempo real
- [ ] **Form state management**: Estados complejos de formularios

### Debugging y Troubleshooting
- [ ] **Signal debugger**: Herramientas para rastrear cambios
- [ ] **Error patterns**: Identificar y solucionar errores comunes
- [ ] **Memory leaks**: Prevenir fugas de memoria en effects
- [ ] **Performance issues**: Optimizar signals lentos

### Arquitectura Avanzada
- [ ] **State management**: Gestión de estado a gran escala
- [ ] **Service layer**: Servicios con signals
- [ ] **Component communication**: Comunicación entre componentes
- [ ] **Testing signals**: Estrategias de testing

---

## Reflexiones Finales

Los Signals representan una evolución fundamental en la programación reactiva. No son solo una nueva API, sino una nueva forma de pensar sobre el estado y las dependencias en tus aplicaciones.

**Recuerda:**
- **Signals son síncronos**: A diferencia de Observables, siempre tienen un valor actual
- **La reactividad es automática**: Los computed y effects se actualizan automáticamente
- **El performance es clave**: Signals están optimizados para cambios frecuentes
- **La simplicidad importa**: Menos código boilerplate, más lógica de negocio

En el próximo módulo, exploraremos **Signals en el mundo real**: integración con NgRx, testing avanzado, y patrones enterprise para aplicaciones a gran escala.

¿Estás listo para dominar la reactividad del futuro?

---

## 📚 Recursos Adicionales

- [Angular Signals Documentation](https://angular.io/guide/signals) - Documentación oficial
- [Signals RFC](https://github.com/angular/angular/discussions/49685) - Propuesta original
- [RxJS vs Signals](https://blog.angular.io/rxjs-when-to-use-signals/) - Cuándo usar cada uno
- [Signals Best Practices](https://dev.to/angular/signals-best-practices-4b0l) - Patrones recomendados

## 🎯 Proyecto de Práctica

Construye una aplicación completa de gestión de tareas usando signals:

### Nivel Básico:
1. **Signals básicos**: Crear, leer, actualizar signals
2. **Computed simple**: Calcular estadísticas básicas
3. **Effects básicos**: Logging de cambios

### Nivel Intermedio:
1. **Linked signals**: Estado conectado entre componentes
2. **Form integration**: Formularios reactivos con signals
3. **Filtering & sorting**: Filtros dinámicos con computed

### Nivel Avanzado:
1. **State management**: Arquitectura completa con servicios
2. **Real-time updates**: Sincronización con backend
3. **Offline support**: Persistencia local con signals
4. **Performance optimization**: Memoización y lazy loading

### Debugging Challenge:
1. **Implementa logging**: Sistema de debugging para signals
2. **Detect memory leaks**: Herramientas para encontrar fugas
3. **Performance profiling**: Optimizar signals lentos
4. **Error boundaries**: Manejo robusto de errores

¡Comparte tu aplicación en la comunidad y recibe feedback sobre tu implementación de signals!

## 🚀 Próximos Pasos

Con signals dominados, estás preparado para:
- **Aplicaciones enterprise**: Estado complejo y escalable
- **Performance crítica**: Aplicaciones que manejan miles de actualizaciones
- **Arquitecturas modernas**: Signals + RxJS + NgRx
- **Testing avanzado**: Estrategias específicas para signals
- **Open source**: Contribuir a librerías con signals

¡Felicitaciones! Has aprendido la tecnología del futuro de Angular. Los signals cambiarán cómo escribes aplicaciones reactive para siempre. 🎉
