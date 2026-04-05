# Capítulo 20: Formularios Template-Driven - La Simplicidad Elegante

Imagina que estás en una cafetería, pidiendo un café. El camarero te pregunta: "¿Cómo lo quieres? ¿Con azúcar, leche, o negro?" Tú respondes, y él prepara exactamente lo que pediste. En Angular, los formularios template-driven funcionan de manera similar: el HTML (el camarero) toma la iniciativa, y TypeScript (tú) solo interviene cuando es necesario.

Este capítulo te guiará a través de los formularios más simples y directos de Angular, perfectos para cuando quieres algo rápido sin complicaciones.

## ¿Qué Aprenderás en Este Capítulo?

En este capítulo, exploraremos:
- La filosofía detrás de los formularios template-driven
- Cómo conectar HTML con TypeScript de forma intuitiva
- Validaciones básicas que guían al usuario
- Estados del formulario y cómo responder a ellos
- Cuándo elegir este enfoque sobre los reactivos

## La Filosofía de los Formularios Template-Driven

Piensa en un formulario como una conversación entre el usuario y tu aplicación. En los template-driven, el HTML lleva la voz cantante: define qué campos existen, qué validaciones aplicar, y cómo mostrar errores. TypeScript solo escucha y actúa cuando el formulario se envía.

Es como escribir una carta: el papel (HTML) contiene las palabras, pero tú (TypeScript) decides qué hacer con ella una vez escrita.

### Analogía del Restaurante

- **Template-Driven**: El menú está en la mesa (HTML). Tú eliges qué pedir, y el camarero valida si tu pedido es correcto.
- **Reactive**: Tú dictas exactamente qué quieres, y el camarero solo ejecuta tus instrucciones.

## ¿Para Qué Sirven Estos Formularios?

Los formularios template-driven brillan en situaciones cotidianas:
- Formularios de contacto simples
- Registros de usuario básicos
- Configuraciones de perfil
- Cualquier formulario donde la lógica sea directa

Son ideales cuando quieres enfocarte en la experiencia del usuario sin perder tiempo en código complejo.

### El Problema que Resuelven

Antes de Angular, crear formularios era un dolor de cabeza. Imagina tener que:

```javascript
// JavaScript puro - ¡qué tedioso!
function validateForm() {
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;

  if (!email.includes('@')) {
    document.getElementById('emailError').textContent = 'Email inválido';
    return false;
  }

  if (password.length < 8) {
    document.getElementById('passwordError').textContent = 'Contraseña muy corta';
    return false;
  }

  // Y así para cada campo...
}
```

Con Angular template-driven, esto se simplifica enormemente.

## Cómo Construir un Formulario Template-Driven

### Paso 1: Preparar el Terreno

Primero, importa `FormsModule` en tu módulo:

```typescript
import { FormsModule } from '@angular/forms';

@NgModule({
  imports: [FormsModule],
  // ...
})
export class AppModule { }
```

### Paso 2: Crear el Template

El HTML es donde ocurre la magia:

```html
<form #loginForm="ngForm" (ngSubmit)="onSubmit(loginForm.value)">
  <!-- Campo de email -->
  <div class="form-group">
    <label for="email">Correo electrónico:</label>
    <input
      id="email"
      type="email"
      name="email"
      [(ngModel)]="user.email"
      required
      email
      #emailField="ngModel"
    >

    <!-- Mostrar errores solo cuando sea necesario -->
    <div *ngIf="emailField.invalid && emailField.touched" class="error">
      <span *ngIf="emailField.errors?.['required']">El email es obligatorio</span>
      <span *ngIf="emailField.errors?.['email']">Ingresa un email válido</span>
    </div>
  </div>

  <!-- Campo de contraseña -->
  <div class="form-group">
    <label for="password">Contraseña:</label>
    <input
      id="password"
      type="password"
      name="password"
      [(ngModel)]="user.password"
      required
      minlength="8"
      #passwordField="ngModel"
    >

    <div *ngIf="passwordField.invalid && passwordField.touched" class="error">
      <span *ngIf="passwordField.errors?.['required']">La contraseña es obligatoria</span>
      <span *ngIf="passwordField.errors?.['minlength']">Mínimo 8 caracteres</span>
    </div>
  </div>

  <!-- Botón inteligente -->
  <button
    type="submit"
    [disabled]="!loginForm.valid"
    class="btn-submit"
  >
    Iniciar Sesión
  </button>

  <!-- Información del estado -->
  <p>¿Válido? {{ loginForm.valid }}</p>
  <p>¿Modificado? {{ loginForm.dirty }}</p>
</form>
```

### Paso 3: La Lógica en TypeScript

El componente es sorprendentemente simple:

```typescript
import { Component } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html'
})
export class LoginComponent {
  user = {
    email: '',
    password: ''
  };

  onSubmit(formData: any) {
    console.log('Datos del formulario:', formData);
    // Aquí enviarías los datos al servidor
  }
}
```

## Entendiendo los Estados del Formulario

Un formulario template-driven tiene varios estados que puedes usar para mejorar la UX:

- **`valid`**: Todos los campos cumplen las reglas
- **`invalid`**: Al menos un campo tiene errores
- **`pristine`**: Ningún campo ha sido tocado
- **`dirty`**: Al menos un campo ha sido modificado
- **`touched`**: El usuario ha interactuado con al menos un campo
- **`untouched`**: Ningún campo ha sido tocado

### Estrategia de Validación Inteligente

```html
<!-- Mostrar errores solo después de interactuar -->
<div *ngIf="field.invalid && field.touched" class="error">
  <!-- Mensajes específicos -->
</div>

<!-- Deshabilitar envío si hay errores -->
<button [disabled]="!form.valid">Enviar</button>

<!-- Cambiar apariencia cuando se modifica -->
<input [class.modified]="form.dirty">
```

## Validadores Integrados

Angular incluye validadores listos para usar:

| Validador | Uso | Ejemplo |
|-----------|-----|---------|
| `required` | Campo obligatorio | `<input required>` |
| `email` | Formato de email | `<input email>` |
| `minlength` | Longitud mínima | `<input minlength="8">` |
| `maxlength` | Longitud máxima | `<input maxlength="100">` |
| `pattern` | Expresión regular | `<input pattern="[0-9]+">` |

## Tipos de Campos Comunes

### Select (Desplegable)

```html
<select [(ngModel)]="user.country" name="country" required>
  <option value="">Selecciona un país</option>
  <option value="es">España</option>
  <option value="mx">México</option>
  <option value="ar">Argentina</option>
</select>
```

### Checkbox

```html
<input
  type="checkbox"
  [(ngModel)]="user.agree"
  name="agree"
  required
>
Acepto los términos y condiciones
```

### Radio Buttons

```html
<div>
  <input type="radio" [(ngModel)]="user.gender" name="gender" value="male"> Masculino
  <input type="radio" [(ngModel)]="user.gender" name="gender" value="female"> Femenino
</div>
```

### Textarea

```html
<textarea
  [(ngModel)]="user.bio"
  name="bio"
  rows="4"
  maxlength="500"
  placeholder="Cuéntanos sobre ti..."
></textarea>
<span>{{ user.bio?.length || 0 }}/500</span>
```

## Un Ejemplo Completo: Formulario de Registro

Vamos a crear un formulario de registro completo paso a paso.

### El Componente TypeScript

```typescript
import { Component } from '@angular/core';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  user = {
    name: '',
    email: '',
    password: '',
    confirmPassword: '',
    country: '',
    agree: false
  };

  submitted = false;

  onSubmit(formData: any) {
    this.submitted = true;

    if (this.isValidForm(formData)) {
      console.log('Registrando usuario:', formData);
      // Llamar al servicio de autenticación
    }
  }

  private isValidForm(data: any): boolean {
    return data.name && data.email && data.password && data.agree;
  }
}
```

### El Template HTML

```html
<div class="register-container">
  <h2>Crear Cuenta</h2>

  <form #registerForm="ngForm" (ngSubmit)="onSubmit(registerForm.value)">
    <!-- Nombre -->
    <div class="form-group">
      <label for="name">Nombre completo:</label>
      <input
        id="name"
        type="text"
        name="name"
        [(ngModel)]="user.name"
        required
        minlength="2"
        #nameField="ngModel"
      >
      <div *ngIf="nameField.invalid && nameField.touched" class="error">
        <span *ngIf="nameField.errors?.['required']">El nombre es obligatorio</span>
        <span *ngIf="nameField.errors?.['minlength']">Mínimo 2 caracteres</span>
      </div>
    </div>

    <!-- Email -->
    <div class="form-group">
      <label for="email">Correo electrónico:</label>
      <input
        id="email"
        type="email"
        name="email"
        [(ngModel)]="user.email"
        required
        email
        #emailField="ngModel"
      >
      <div *ngIf="emailField.invalid && emailField.touched" class="error">
        <span *ngIf="emailField.errors?.['required']">El email es obligatorio</span>
        <span *ngIf="emailField.errors?.['email']">Ingresa un email válido</span>
      </div>
    </div>

    <!-- Contraseña -->
    <div class="form-group">
      <label for="password">Contraseña:</label>
      <input
        id="password"
        type="password"
        name="password"
        [(ngModel)]="user.password"
        required
        minlength="8"
        #passwordField="ngModel"
      >
      <div *ngIf="passwordField.invalid && passwordField.touched" class="error">
        <span *ngIf="passwordField.errors?.['required']">La contraseña es obligatoria</span>
        <span *ngIf="passwordField.errors?.['minlength']">Mínimo 8 caracteres</span>
      </div>
    </div>

    <!-- País -->
    <div class="form-group">
      <label for="country">País:</label>
      <select
        id="country"
        name="country"
        [(ngModel)]="user.country"
        required
        #countryField="ngModel"
      >
        <option value="">Selecciona tu país</option>
        <option value="es">España</option>
        <option value="mx">México</option>
        <option value="ar">Argentina</option>
        <option value="co">Colombia</option>
      </select>
      <div *ngIf="countryField.invalid && countryField.touched" class="error">
        <span>Selecciona un país</span>
      </div>
    </div>

    <!-- Términos -->
    <div class="form-group">
      <input
        id="agree"
        type="checkbox"
        name="agree"
        [(ngModel)]="user.agree"
        required
        #agreeField="ngModel"
      >
      <label for="agree">Acepto los términos y condiciones</label>
      <div *ngIf="agreeField.invalid && agreeField.touched" class="error">
        <span>Debes aceptar los términos</span>
      </div>
    </div>

    <!-- Botón -->
    <button
      type="submit"
      [disabled]="!registerForm.valid"
      class="btn-submit"
    >
      Crear Cuenta
    </button>

    <!-- Feedback -->
    <div *ngIf="submitted && registerForm.valid" class="success">
      ¡Cuenta creada exitosamente!
    </div>
  </form>
</div>
```

### Estilos CSS

```css
.register-container {
  max-width: 500px;
  margin: 0 auto;
  padding: 20px;
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
}

input, select {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

input:focus, select:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.error {
  color: #dc3545;
  font-size: 14px;
  margin-top: 5px;
}

.success {
  color: #28a745;
  background: #d4edda;
  padding: 10px;
  border-radius: 4px;
  margin-top: 20px;
}

.btn-submit {
  width: 100%;
  padding: 12px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
}

.btn-submit:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.btn-submit:hover:not(:disabled) {
  background: #0056b3;
}
```

## Cuándo Elegir Template-Driven

### ✅ Úsalos cuando:
- El formulario es simple (menos de 5 campos)
- Las validaciones son básicas
- Quieres código rápido y legible
- No necesitas lógica compleja de validación

### ❌ Evítalos si:
- Necesitas validaciones dinámicas complejas
- El formulario tiene muchos campos interdependientes
- Quieres control total programático
- Necesitas testing unitario avanzado

## Errores Comunes y Cómo Evitarlos

1. **Olvidar importar FormsModule**: Sin esto, `ngModel` no funciona.
2. **No usar `name` en los campos**: Angular necesita este atributo para trackear los campos.
3. **Mostrar errores demasiado pronto**: Usa `field.touched` para no molestar al usuario.
4. **No deshabilitar el botón**: Permite envíos inválidos.
5. **Ignorar el estado del formulario**: Usa `form.valid`, `form.dirty`, etc.

## Mini Práctica: Tu Propio Formulario

1. Crea un componente `ContactFormComponent`
2. Incluye campos para nombre, email, mensaje y un checkbox de "contactarme"
3. Agrega validaciones apropiadas
4. Muestra errores solo cuando el usuario interactúe
5. Deshabilita el envío hasta que todo sea válido
6. En `onSubmit`, muestra los datos en consola

## Resumen del Capítulo

Los formularios template-driven son la forma más intuitiva de manejar formularios en Angular. El HTML lleva la iniciativa, definiendo campos y validaciones, mientras TypeScript se mantiene simple y enfocado en la lógica de negocio.

Son perfectos para formularios cotidianos donde la simplicidad es más importante que el control absoluto. Recuerda: cuando dudes entre template-driven y reactive, elige el que haga tu código más legible y mantenible.

En el próximo capítulo, exploraremos los formularios reactivos, donde TypeScript toma el control completo.


