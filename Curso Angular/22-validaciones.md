# Módulo 22 - Validaciones Profesionales

## ¿Qué aprenderás en este módulo?

En este módulo aprenderás:
- cómo asegurar que los datos del formulario sean válidos
- cómo crear validadores personalizados y grupales
- cómo usar validadores asíncronos con APIs
- cómo mostrar mensajes claros al usuario
- por qué las validaciones mejoran la UX y la seguridad

---

## ¿Para qué sirven las validaciones?

Las validaciones evitan que el usuario envíe datos incorrectos.
También guían al usuario para que escriba la información esperada, reduciendo errores de backend y mejorando la experiencia.

---

## Cómo validar en Angular

### 1. Validadores built-in

Angular incluye validadores listos para usar:
- `Validators.required`
- `Validators.email`
- `Validators.minLength(n)`
- `Validators.maxLength(n)`
- `Validators.min(n)`
- `Validators.max(n)`
- `Validators.pattern(regex)`

```typescript
const form = this.fb.group({
  email: ['', [Validators.required, Validators.email]],
  password: ['', [Validators.required, Validators.minLength(8)]],
  age: ['', [Validators.min(18), Validators.max(100)]]
});
```

### 2. Validadores custom

```typescript
import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

export function strongPasswordValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;
    if (!value) return null;

    const hasNumber = /[0-9]/.test(value);
    const hasUpper = /[A-Z]/.test(value);
    const hasLower = /[a-z]/.test(value);
    const isLongEnough = value.length >= 8;

    return hasNumber && hasUpper && hasLower && isLongEnough
      ? null
      : { weakPassword: true };
  };
}
```

```typescript
const form = this.fb.group({
  password: ['', [Validators.required, strongPasswordValidator()]]
});
```

### 3. Validadores cross-field

```typescript
export function passwordsMatchValidator(formGroup: FormGroup): ValidationErrors | null {
  const password = formGroup.get('password')?.value;
  const confirm = formGroup.get('confirmPassword')?.value;

  return password === confirm ? null : { passwordMismatch: true };
}

const form = this.fb.group({
  password: ['', Validators.required],
  confirmPassword: ['', Validators.required]
}, { validators: passwordsMatchValidator });
```

### 4. Validadores asincrónicos

```typescript
export function usernameAvailableValidator(
  userService: UserService
): AsyncValidatorFn {
  return (control: AbstractControl) => {
    if (!control.value) return of(null);

    return userService.checkUsername(control.value).pipe(
      debounceTime(300),
      map(isAvailable => isAvailable ? null : { usernameTaken: true }),
      catchError(() => of(null))
    );
  };
}
```

```typescript
const form = this.fb.group({
  username: ['', Validators.required, [usernameAvailableValidator(this.userService)]]
});
```

### 5. Mensajes claros para el usuario

```html
<div *ngIf="emailControl.invalid && emailControl.touched">
  <p *ngIf="emailControl.errors?.['required']">Email requerido</p>
  <p *ngIf="emailControl.errors?.['email']">Email inválido</p>
</div>
```

### 6. Validación en tiempo real

```typescript
this.passwordStrength$ = this.form.get('password')!.valueChanges.pipe(
  debounceTime(300),
  map(pwd => this.calculateStrength(pwd))
);
```

```html
<div class="strength-bar" [style.width.%]="(passwordStrength$ | async)"></div>
```

---

## Errores comunes

1. Mostrar errores antes de que el usuario toque el campo.
2. Usar validadores en el template cuando se necesita lógica dinámica.
3. No limpiar mensajes cuando el valor se corrige.
4. No validar campos relacionados (por ejemplo, contraseña y confirmación).
5. No manejar errores de validadores asíncronos.

---

## Mini práctica

1. Crea un formulario con email, contraseña y confirmación.
2. Agrega validadores custom para fuerza de contraseña.
3. Crea un validador cross-field que compare contraseña y confirmación.
4. Agrega un validador async que consulte una API simulada para username.

---

## Resumen

Las validaciones profesionales combinan validadores built-in, custom y asíncronos.
Muestra siempre mensajes claros y evita que el formulario se envíe si hay errores.
Una buena validación mejora tanto la seguridad como la experiencia del usuario.
