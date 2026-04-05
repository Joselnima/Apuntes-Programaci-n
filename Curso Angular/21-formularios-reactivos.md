# Capítulo 21: Formularios Reactivos - El Control Maestro

Imagina que estás dirigiendo una orquesta. En los formularios template-driven, los músicos (HTML) tocan según su intuición. Pero en los formularios reactivos, tú eres el director: controlas cada nota, cada pausa, cada dinámica. Tienes el poder absoluto, pero también la responsabilidad total.

Este capítulo te llevará del control básico al dominio completo de los formularios reactivos en Angular.

## ¿Qué Aprenderás en Este Capítulo?

En este capítulo, exploraremos:
- La filosofía del control total en formularios
- Cómo construir formularios desde TypeScript
- Validaciones dinámicas que responden al contexto
- FormArrays para listas complejas
- Observables y reacciones en tiempo real
- Cuándo elegir este enfoque sobre los template-driven

## La Filosofía del Control Reactivo

Piensa en un formulario como una máquina compleja. En los template-driven, la máquina funciona automáticamente: pones los engranajes (HTML) y ella hace su trabajo. En los reactivos, tú diseñas cada engranaje, controlas cada movimiento, y puedes cambiar la máquina mientras funciona.

Es como la diferencia entre conducir un auto automático y uno manual: el automático es más fácil para el día a día, pero el manual te da control absoluto cuando lo necesitas.

### Analogía del Restaurante Evolucionada

- **Template-Driven**: El menú está fijo. El cliente elige, el camarero valida.
- **Reactive**: Tú eres el chef. Creas el menú dinámicamente, cambias los platos según los ingredientes disponibles, y controlas cada aspecto de la experiencia.

## ¿Para Qué Sirven los Formularios Reactivos?

Los formularios reactivos brillan cuando necesitas:

- **Validación condicional**: "Este campo es obligatorio solo si..."
- **Formularios dinámicos**: Agregar/eliminar campos en tiempo real
- **Testing avanzado**: Probar lógica de validación aisladamente
- **Integración con RxJS**: Reacciones en tiempo real a cambios
- **Formularios complejos**: Múltiples secciones con lógica interdependiente

### El Problema que Resuelven

Imagina un formulario de envío internacional:

```typescript
// ❌ Template-Driven no puede hacer esto fácilmente
// ¿Cómo hacer que "Estado" sea obligatorio solo para direcciones US?
// ¿Cómo agregar campos de aduana dinámicamente?
```

Con reactivos, esto es trivial.

## Cómo Construir un Formulario Reactivo

### Paso 1: Preparar el Terreno

Primero, importa ReactiveFormsModule:

```typescript
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  imports: [ReactiveFormsModule],
  // ...
})
export class AppModule { }
```

### Paso 2: Crear el Formulario en TypeScript

El corazón de los formularios reactivos está en TypeScript:

```typescript
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-contact-form',
  templateUrl: './contact-form.component.html'
})
export class ContactFormComponent implements OnInit {
  contactForm: FormGroup;

  constructor(private fb: FormBuilder) {
    // FormBuilder: tu asistente para crear formularios
    this.contactForm = this.fb.group({
      // Cada propiedad es un FormControl
      name: ['', [Validators.required, Validators.minLength(2)]],
      email: ['', [Validators.required, Validators.email]],
      message: ['', [Validators.required, Validators.maxLength(500)]],
      urgent: [false] // Sin validadores para checkboxes opcionales
    });
  }

  ngOnInit() {
    // Aquí puedes configurar observadores, validaciones dinámicas, etc.
  }

  onSubmit() {
    if (this.contactForm.valid) {
      console.log('Enviando mensaje:', this.contactForm.value);
      // Llamar al servicio de envío
    }
  }
}
```

### Paso 3: Conectar en el Template

El template es sorprendentemente simple:

```html
<form [formGroup]="contactForm" (ngSubmit)="onSubmit()">
  <!-- Nombre -->
  <div class="form-group">
    <label for="name">Nombre:</label>
    <input
      id="name"
      type="text"
      formControlName="name"
    >

    <div *ngIf="contactForm.get('name')?.invalid && contactForm.get('name')?.touched" class="error">
      <span *ngIf="contactForm.get('name')?.errors?.['required']">El nombre es obligatorio</span>
      <span *ngIf="contactForm.get('name')?.errors?.['minlength']">Mínimo 2 caracteres</span>
    </div>
  </div>

  <!-- Email -->
  <div class="form-group">
    <label for="email">Email:</label>
    <input
      id="email"
      type="email"
      formControlName="email"
    >

    <div *ngIf="contactForm.get('email')?.invalid && contactForm.get('email')?.touched" class="error">
      <span *ngIf="contactForm.get('email')?.errors?.['required']">El email es obligatorio</span>
      <span *ngIf="contactForm.get('email')?.errors?.['email']">Ingresa un email válido</span>
    </div>
  </div>

  <!-- Mensaje -->
  <div class="form-group">
    <label for="message">Mensaje:</label>
    <textarea
      id="message"
      formControlName="message"
      rows="4"
      maxlength="500"
    ></textarea>

    <div class="char-count">
      {{ contactForm.get('message')?.value?.length || 0 }}/500
    </div>

    <div *ngIf="contactForm.get('message')?.invalid && contactForm.get('message')?.touched" class="error">
      <span *ngIf="contactForm.get('message')?.errors?.['required']">El mensaje es obligatorio</span>
      <span *ngIf="contactForm.get('message')?.errors?.['maxlength']">Máximo 500 caracteres</span>
    </div>
  </div>

  <!-- Urgente -->
  <div class="form-group">
    <input
      id="urgent"
      type="checkbox"
      formControlName="urgent"
    >
    <label for="urgent">Marcar como urgente</label>
  </div>

  <!-- Botón inteligente -->
  <button
    type="submit"
    [disabled]="contactForm.invalid"
    class="btn-submit"
  >
    Enviar Mensaje
  </button>

  <!-- Información del estado -->
  <div class="form-status">
    <p>Estado: {{ contactForm.valid ? 'Válido' : 'Inválido' }}</p>
    <p>Modificado: {{ contactForm.dirty ? 'Sí' : 'No' }}</p>
  </div>
</form>
```

## Entendiendo las Estructuras Básicas

### FormControl: El Átomo del Formulario

Un FormControl representa un campo individual. Es como una caja que contiene un valor y sabe si ese valor es válido.

```typescript
// Crear un control básico
const nameControl = new FormControl('');

// Crear con valor inicial y validadores
const emailControl = new FormControl('', [
  Validators.required,
  Validators.email
]);

// Propiedades útiles
console.log(emailControl.value);      // El valor actual
console.log(emailControl.valid);      // true/false
console.log(emailControl.errors);     // { required: true, email: true }
console.log(emailControl.touched);    // Usuario interactuó?
console.log(emailControl.dirty);      // Valor cambió?
```

### FormGroup: El Organizador

Un FormGroup agrupa múltiples FormControls. Es como un director de orquesta que coordina varios instrumentos.

```typescript
// Crear un grupo manualmente
const addressGroup = new FormGroup({
  street: new FormControl('', Validators.required),
  city: new FormControl('', Validators.required),
  zipCode: new FormControl('', [Validators.required, Validators.pattern('[0-9]{5}')])
});

// O usar FormBuilder (más limpio)
const addressGroup = this.fb.group({
  street: ['', Validators.required],
  city: ['', Validators.required],
  zipCode: ['', [Validators.required, Validators.pattern('[0-9]{5}')]]
});

// Acceder a controles
addressGroup.get('street')?.value;
addressGroup.get('city')?.setValue('Nueva York');

// Estado del grupo
addressGroup.valid;    // Todos los controles válidos?
addressGroup.value;    // { street: '...', city: '...', zipCode: '...' }
```

### FormArray: Para Listas Dinámicas

FormArray maneja listas de controles que pueden crecer o encogerse dinámicamente.

```typescript
export class SkillsComponent {
  skillsForm: FormGroup;

  constructor(private fb: FormBuilder) {
    this.skillsForm = this.fb.group({
      name: ['', Validators.required],
      skills: this.fb.array([]) // Array vacío inicialmente
    });
  }

  // Getter para acceder fácilmente al array
  get skills() {
    return this.skillsForm.get('skills') as FormArray;
  }

  // Agregar una nueva skill
  addSkill() {
    const skillGroup = this.fb.group({
      name: ['', Validators.required],
      level: ['beginner', Validators.required]
    });

    this.skills.push(skillGroup);
  }

  // Eliminar una skill
  removeSkill(index: number) {
    this.skills.removeAt(index);
  }

  onSubmit() {
    console.log('Skills:', this.skillsForm.value);
  }
}
```

Template correspondiente:

```html
<form [formGroup]="skillsForm" (ngSubmit)="onSubmit()">
  <input formControlName="name" placeholder="Tu nombre">

  <div formArrayName="skills">
    <div *ngFor="let skill of skills.controls; let i = index" [formGroupName]="i">
      <input formControlName="name" placeholder="Skill name">
      <select formControlName="level">
        <option value="beginner">Principiante</option>
        <option value="intermediate">Intermedio</option>
        <option value="expert">Experto</option>
      </select>
      <button type="button" (click)="removeSkill(i)">Eliminar</button>
    </div>
  </div>

  <button type="button" (click)="addSkill()">Agregar Skill</button>
  <button type="submit" [disabled]="skillsForm.invalid">Guardar</button>
</form>
```

## Validadores: Los Guardianes de la Calidad

### Validadores Built-in

Angular incluye validadores comunes:

```typescript
const userForm = this.fb.group({
  // Obligatorio
  name: ['', Validators.required],

  // Email
  email: ['', Validators.email],

  // Longitud de texto
  username: ['', [Validators.minLength(3), Validators.maxLength(20)]],

  // Números
  age: ['', [Validators.min(18), Validators.max(100)]],

  // Patrones (regex)
  phone: ['', Validators.pattern('[0-9]{10}')],

  // Múltiples validadores
  password: ['', [
    Validators.required,
    Validators.minLength(8),
    Validators.pattern(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/) // Mayúscula, minúscula, número
  ]]
});
```

### Validadores Personalizados

Cuando los built-in no alcanzan, crea los tuyos:

```typescript
// Validador simple (función)
function noWhitespace(control: AbstractControl): ValidationErrors | null {
  const value = control.value;
  if (value && value.trim().length === 0) {
    return { 'whitespace': true };
  }
  return null;
}

// Validador con parámetros
function ageRange(min: number, max: number) {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;
    if (value < min || value > max) {
      return { 'ageRange': { min, max, actual: value } };
    }
    return null;
  };
}

// Usarlos
const form = this.fb.group({
  nickname: ['', [Validators.required, noWhitespace]],
  age: ['', [Validators.required, ageRange(18, 65)]]
});
```

### Validadores de Grupo

Validan relaciones entre múltiples campos:

```typescript
// Validar que contraseñas coincidan
function passwordsMatch(group: AbstractControl): ValidationErrors | null {
  const password = group.get('password')?.value;
  const confirm = group.get('confirmPassword')?.value;

  if (password !== confirm) {
    return { 'passwordMismatch': true };
  }
  return null;
}

// Validar fecha lógica
function dateRange(group: AbstractControl): ValidationErrors | null {
  const start = new Date(group.get('startDate')?.value);
  const end = new Date(group.get('endDate')?.value);

  if (start >= end) {
    return { 'invalidDateRange': true };
  }
  return null;
}

// Aplicar al formulario
const form = this.fb.group({
  password: ['', Validators.required],
  confirmPassword: ['', Validators.required]
}, { validators: passwordsMatch });
```

## Reacciones en Tiempo Real con Observables

Los formularios reactivos brillan cuando se combinan con RxJS:

```typescript
export class SearchComponent implements OnInit {
  searchForm: FormGroup;
  results$ = new BehaviorSubject<any[]>([]);

  constructor(
    private fb: FormBuilder,
    private searchService: SearchService
  ) {
    this.searchForm = this.fb.group({
      query: [''],
      category: ['all']
    });
  }

  ngOnInit() {
    // Reaccionar a cambios en la búsqueda
    this.searchForm.get('query')!.valueChanges.pipe(
      debounceTime(300),              // Esperar 300ms
      distinctUntilChanged(),         // Solo si cambió
      filter(query => query.length > 2), // Mínimo 3 caracteres
      switchMap(query =>              // Buscar
        this.searchService.search(query, this.searchForm.get('category')!.value)
      )
    ).subscribe(results => {
      this.results$.next(results);
    });

    // Reaccionar a cambios de categoría
    this.searchForm.get('category')!.valueChanges.subscribe(category => {
      const query = this.searchForm.get('query')!.value;
      if (query.length > 2) {
        // Rebúsqueda inmediata al cambiar categoría
        this.searchService.search(query, category)
          .subscribe(results => this.results$.next(results));
      }
    });
  }
}
```

Template:

```html
<form [formGroup]="searchForm">
  <input formControlName="query" placeholder="Buscar...">

  <select formControlName="category">
    <option value="all">Todo</option>
    <option value="users">Usuarios</option>
    <option value="products">Productos</option>
  </select>
</form>

<div *ngIf="results$.value.length > 0">
  <h3>Resultados ({{ results$.value.length }})</h3>
  <ul>
    <li *ngFor="let result of results$ | async">
      {{ result.name }}
    </li>
  </ul>
</div>
```

## Validación Dinámica: Cambiando las Reglas en Tiempo Real

Uno de los superpoderes de los formularios reactivos:

```typescript
export class ShippingComponent {
  shippingForm: FormGroup;
  isInternational = false;

  constructor(private fb: FormBuilder) {
    this.shippingForm = this.fb.group({
      country: ['US'],
      state: [''],      // Será obligatorio solo para US
      province: [''],   // Será obligatorio solo para internacional
      postalCode: ['']
    });
  }

  ngOnInit() {
    // Escuchar cambios en el país
    this.shippingForm.get('country')!.valueChanges.subscribe(country => {
      this.updateValidation(country);
    });

    // Validación inicial
    this.updateValidation('US');
  }

  private updateValidation(country: string) {
    this.isInternational = country !== 'US';

    const stateControl = this.shippingForm.get('state')!;
    const provinceControl = this.shippingForm.get('province')!;
    const postalControl = this.shippingForm.get('postalCode')!;

    if (this.isInternational) {
      // Para internacional: province obligatorio, state opcional
      stateControl.clearValidators();
      provinceControl.setValidators([Validators.required]);
      postalControl.setValidators([Validators.required]); // Código postal siempre requerido
    } else {
      // Para US: state obligatorio, province opcional
      stateControl.setValidators([Validators.required]);
      provinceControl.clearValidators();
      postalControl.setValidators([Validators.required, Validators.pattern('[0-9]{5}')]);
    }

    // Actualizar validación de todos los controles
    stateControl.updateValueAndValidity();
    provinceControl.updateValueAndValidity();
    postalControl.updateValueAndValidity();
  }
}
```

## Un Ejemplo Completo: Formulario de Reserva de Hotel

Vamos a crear un formulario complejo que demuestre todo lo aprendido:

```typescript
export class HotelBookingComponent implements OnInit {
  bookingForm: FormGroup;
  totalPrice = 0;

  roomTypes = [
    { id: 'single', name: 'Individual', price: 80 },
    { id: 'double', name: 'Doble', price: 120 },
    { id: 'suite', name: 'Suite', price: 200 }
  ];

  constructor(private fb: FormBuilder) {
    this.bookingForm = this.fb.group({
      guestInfo: this.fb.group({
        name: ['', [Validators.required, Validators.minLength(2)]],
        email: ['', [Validators.required, Validators.email]],
        phone: ['', [Validators.required, Validators.pattern('[0-9]{10}')]]
      }),
      bookingDetails: this.fb.group({
        roomType: ['single', Validators.required],
        checkIn: ['', Validators.required],
        checkOut: ['', Validators.required],
        guests: [1, [Validators.required, Validators.min(1), Validators.max(4)]]
      }),
      extras: this.fb.array([]),
      termsAccepted: [false, Validators.requiredTrue]
    });
  }

  ngOnInit() {
    // Calcular precio cuando cambien las fechas o tipo de habitación
    this.bookingForm.get('bookingDetails')!.valueChanges.subscribe(details => {
      this.calculatePrice();
    });

    // Agregar algunos extras por defecto
    this.addExtra('wifi', 'WiFi gratuito', 0);
    this.addExtra('parking', 'Estacionamiento', 15);
  }

  get extras() {
    return this.bookingForm.get('extras') as FormArray;
  }

  addExtra(id: string, name: string, price: number) {
    const extraGroup = this.fb.group({
      id: [id],
      name: [name],
      price: [price],
      selected: [false]
    });

    this.extras.push(extraGroup);
  }

  calculatePrice() {
    const details = this.bookingForm.get('bookingDetails')!.value;
    const roomType = this.roomTypes.find(r => r.id === details.roomType);

    if (!roomType || !details.checkIn || !details.checkOut) {
      this.totalPrice = 0;
      return;
    }

    const nights = Math.ceil(
      (new Date(details.checkOut).getTime() - new Date(details.checkIn).getTime())
      / (1000 * 60 * 60 * 24)
    );

    let price = roomType.price * nights;

    // Agregar extras seleccionados
    this.extras.controls.forEach(extra => {
      if (extra.get('selected')!.value) {
        price += extra.get('price')!.value * nights;
      }
    });

    this.totalPrice = price;
  }

  onSubmit() {
    if (this.bookingForm.valid) {
      const bookingData = {
        ...this.bookingForm.value,
        totalPrice: this.totalPrice,
        bookingDate: new Date()
      };

      console.log('Reserva confirmada:', bookingData);
      // Enviar al servicio de reservas
    }
  }
}
```

Template correspondiente:

```html
<div class="booking-container">
  <h2>Reserva de Hotel</h2>

  <form [formGroup]="bookingForm" (ngSubmit)="onSubmit()">
    <!-- Información del Huésped -->
    <fieldset formGroupName="guestInfo">
      <legend>Información del Huésped</legend>

      <div class="form-row">
        <div class="form-group">
          <label>Nombre completo:</label>
          <input type="text" formControlName="name">
          <div *ngIf="bookingForm.get('guestInfo.name')?.invalid && bookingForm.get('guestInfo.name')?.touched" class="error">
            <span *ngIf="bookingForm.get('guestInfo.name')?.errors?.['required']">Nombre obligatorio</span>
            <span *ngIf="bookingForm.get('guestInfo.name')?.errors?.['minlength']">Mínimo 2 caracteres</span>
          </div>
        </div>

        <div class="form-group">
          <label>Email:</label>
          <input type="email" formControlName="email">
          <div *ngIf="bookingForm.get('guestInfo.email')?.invalid && bookingForm.get('guestInfo.email')?.touched" class="error">
            <span *ngIf="bookingForm.get('guestInfo.email')?.errors?.['required']">Email obligatorio</span>
            <span *ngIf="bookingForm.get('guestInfo.email')?.errors?.['email']">Email inválido</span>
          </div>
        </div>
      </div>

      <div class="form-group">
        <label>Teléfono:</label>
        <input type="tel" formControlName="phone">
        <div *ngIf="bookingForm.get('guestInfo.phone')?.invalid && bookingForm.get('guestInfo.phone')?.touched" class="error">
          <span *ngIf="bookingForm.get('guestInfo.phone')?.errors?.['required']">Teléfono obligatorio</span>
          <span *ngIf="bookingForm.get('guestInfo.phone')?.errors?.['pattern']">Formato inválido (10 dígitos)</span>
        </div>
      </div>
    </fieldset>

    <!-- Detalles de la Reserva -->
    <fieldset formGroupName="bookingDetails">
      <legend>Detalles de la Reserva</legend>

      <div class="form-row">
        <div class="form-group">
          <label>Tipo de habitación:</label>
          <select formControlName="roomType">
            <option *ngFor="let room of roomTypes" [value]="room.id">
              {{ room.name }} - ${{ room.price }}/noche
            </option>
          </select>
        </div>

        <div class="form-group">
          <label>Número de huéspedes:</label>
          <input type="number" formControlName="guests" min="1" max="4">
          <div *ngIf="bookingForm.get('bookingDetails.guests')?.invalid && bookingForm.get('bookingDetails.guests')?.touched" class="error">
            <span *ngIf="bookingForm.get('bookingDetails.guests')?.errors?.['min']">Mínimo 1 huésped</span>
            <span *ngIf="bookingForm.get('bookingDetails.guests')?.errors?.['max']">Máximo 4 huéspedes</span>
          </div>
        </div>
      </div>

      <div class="form-row">
        <div class="form-group">
          <label>Check-in:</label>
          <input type="date" formControlName="checkIn">
        </div>

        <div class="form-group">
          <label>Check-out:</label>
          <input type="date" formControlName="checkOut">
        </div>
      </div>
    </fieldset>

    <!-- Extras -->
    <fieldset>
      <legend>Servicios Adicionales</legend>

      <div formArrayName="extras">
        <div *ngFor="let extra of extras.controls; let i = index" [formGroupName]="i" class="extra-item">
          <input type="checkbox" formControlName="selected">
          <label>{{ extra.get('name')!.value }}</label>
          <span class="price">{{ extra.get('price')!.value === 0 ? 'Gratis' : '$' + extra.get('price')!.value + '/noche' }}</span>
        </div>
      </div>
    </fieldset>

    <!-- Términos -->
    <div class="form-group">
      <input type="checkbox" formControlName="termsAccepted">
      <label>Acepto los términos y condiciones</label>
      <div *ngIf="bookingForm.get('termsAccepted')?.invalid && bookingForm.get('termsAccepted')?.touched" class="error">
        Debes aceptar los términos para continuar
      </div>
    </div>

    <!-- Precio Total -->
    <div class="total-price" *ngIf="totalPrice > 0">
      <h3>Total: ${{ totalPrice }}</h3>
    </div>

    <!-- Botón -->
    <button type="submit" [disabled]="bookingForm.invalid" class="btn-submit">
      Confirmar Reserva
    </button>
  </form>
</div>
```

## Cuándo Elegir Formularios Reactivos

### ✅ Úsalos cuando:
- Necesites validación condicional o dinámica
- El formulario sea complejo (muchos campos interrelacionados)
- Quieras unit testing avanzado
- Necesites reaccionar a cambios en tiempo real
- El formulario deba escalar o cambiar dinámicamente

### ❌ Evítalos si:
- El formulario es muy simple (1-3 campos)
- No necesitas lógica compleja
- Quieres desarrollo rápido
- Los miembros del equipo no están familiarizados con RxJS

## Errores Comunes y Cómo Evitarlos

1. **Olvidar importar ReactiveFormsModule**: Sin esto, `formControlName` no funciona.
2. **Usar `[(ngModel)]` en formularios reactivos**: Mezcla paradigmas, causa conflictos.
3. **No llamar `updateValueAndValidity()`**: Los cambios de validadores no se aplican.
4. **Acceder a controles sin `?.`**: Puede causar errores si el control no existe.
5. **No manejar errores de async validators**: Los validadores asíncronos necesitan manejo especial.
6. **Crear formularios demasiado anidados**: Puede hacer el código difícil de mantener.

## Mini Práctica: Tu Propio Formulario Reactivo

1. Crea un componente `UserProfileComponent`
2. Incluye campos para nombre, email, fecha de nacimiento, y una lista de hobbies
3. Implementa validación de edad (mínimo 18 años)
4. Permite agregar/eliminar hobbies dinámicamente
5. Muestra un precio calculado basado en la cantidad de hobbies
6. Incluye términos y condiciones con validación

## Resumen del Capítulo

Los formularios reactivos son la herramienta definitiva para formularios complejos en Angular. Te dan control total sobre la validación, el estado, y el comportamiento, permitiendo crear experiencias de usuario sofisticadas que responden inteligentemente a las acciones del usuario.

Son más complejos de aprender que los template-driven, pero esa complejidad se traduce en poder y flexibilidad. Cuando domines los reactivos, podrás crear formularios que se adaptan dinámicamente, validan inteligentemente, y proporcionan feedback en tiempo real.

En el próximo capítulo, exploraremos validaciones avanzadas que harán tus formularios aún más robustos y seguros.
