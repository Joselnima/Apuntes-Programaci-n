# Capítulo 26: Autenticación y Seguridad - El Guardián de la Fortaleza

## La Analogía del Club Exclusivo

Imagina que estás organizando una fiesta en un club nocturno exclusivo. Para entrar, los invitados necesitan mostrar una invitación especial en la puerta. Tú eres el portero que verifica cada invitación y decide quién puede pasar. Una vez dentro, algunos invitados tienen acceso a la zona VIP, mientras que otros solo pueden estar en la pista de baile.

En el mundo digital, la **autenticación** es exactamente eso: un sistema que verifica la identidad de los usuarios y controla qué partes de tu aplicación pueden acceder. Es como ser el portero de un club exclusivo donde:

- La **invitación** es el token de autenticación
- El **portero** es el `AuthGuard`
- La **zona VIP** son las rutas protegidas
- El **sistema de sonido** que "anuncia" tu presencia son los interceptores HTTP

En este capítulo, construiremos juntos el sistema de seguridad completo de nuestra aplicación Angular, desde el login hasta la protección de rutas sensibles.

---

## El Flujo de Autenticación: Un Viaje Paso a Paso

### Paso 1: El Punto de Entrada - El Formulario de Login

Antes de construir el sistema de autenticación, necesitamos un lugar donde los usuarios puedan identificarse. Vamos a crear un componente de login que sea intuitivo y seguro.

```typescript
// login.component.ts
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginForm: FormGroup;
  isLoading = false;
  errorMessage = '';

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router
  ) {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  onSubmit() {
    if (this.loginForm.valid) {
      this.isLoading = true;
      this.errorMessage = '';

      this.authService.login(this.loginForm.value).subscribe({
        next: (success) => {
          if (success) {
            this.router.navigate(['/dashboard']);
          } else {
            this.errorMessage = 'Credenciales incorrectas';
          }
        },
        error: (error) => {
          this.errorMessage = 'Error al iniciar sesión. Inténtalo de nuevo.';
          console.error('Error de login:', error);
        },
        complete: () => {
          this.isLoading = false;
        }
      });
    } else {
      this.markFormGroupTouched();
    }
  }

  private markFormGroupTouched() {
    Object.keys(this.loginForm.controls).forEach(key => {
      this.loginForm.get(key)?.markAsTouched();
    });
  }

  get email() { return this.loginForm.get('email'); }
  get password() { return this.loginForm.get('password'); }
}
```

```html
<!-- login.component.html -->
<div class="login-container">
  <div class="login-card">
    <h2 class="login-title">Bienvenido de vuelta</h2>
    <p class="login-subtitle">Ingresa tus credenciales para continuar</p>

    <form [formGroup]="loginForm" (ngSubmit)="onSubmit()" class="login-form">
      <!-- Campo de email -->
      <div class="form-group">
        <label for="email" class="form-label">Correo electrónico</label>
        <input
          type="email"
          id="email"
          formControlName="email"
          class="form-input"
          placeholder="tu@email.com"
          [class.error]="email?.invalid && email?.touched"
        >
        <div class="error-message" *ngIf="email?.invalid && email?.touched">
          <span *ngIf="email?.errors?.['required']">El email es requerido</span>
          <span *ngIf="email?.errors?.['email']">Ingresa un email válido</span>
        </div>
      </div>

      <!-- Campo de contraseña -->
      <div class="form-group">
        <label for="password" class="form-label">Contraseña</label>
        <input
          type="password"
          id="password"
          formControlName="password"
          class="form-input"
          placeholder="Tu contraseña"
          [class.error]="password?.invalid && password?.touched"
        >
        <div class="error-message" *ngIf="password?.invalid && password?.touched">
          <span *ngIf="password?.errors?.['required']">La contraseña es requerida</span>
          <span *ngIf="password?.errors?.['minlength']">Mínimo 6 caracteres</span>
        </div>
      </div>

      <!-- Mensaje de error general -->
      <div class="alert alert-error" *ngIf="errorMessage">
        {{ errorMessage }}
      </div>

      <!-- Botón de submit -->
      <button
        type="submit"
        class="btn btn-primary btn-full"
        [disabled]="isLoading"
      >
        <span *ngIf="isLoading" class="spinner"></span>
        {{ isLoading ? 'Iniciando sesión...' : 'Iniciar sesión' }}
      </button>
    </form>
  </div>
</div>
```

```css
/* login.component.css */
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.login-card {
  background: white;
  border-radius: 12px;
  padding: 40px;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

.login-title {
  color: #333;
  text-align: center;
  margin-bottom: 8px;
  font-size: 28px;
  font-weight: 600;
}

.login-subtitle {
  color: #666;
  text-align: center;
  margin-bottom: 32px;
  font-size: 16px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.form-label {
  font-weight: 500;
  color: #333;
  font-size: 14px;
}

.form-input {
  padding: 12px 16px;
  border: 2px solid #e1e5e9;
  border-radius: 8px;
  font-size: 16px;
  transition: border-color 0.3s ease;
}

.form-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.form-input.error {
  border-color: #e74c3c;
}

.error-message {
  color: #e74c3c;
  font-size: 14px;
  margin-top: 4px;
}

.btn {
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.btn-primary {
  background: #667eea;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #5a6fd8;
  transform: translateY(-1px);
}

.btn-full {
  width: 100%;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.spinner {
  width: 16px;
  height: 16px;
  border: 2px solid #ffffff;
  border-top: 2px solid transparent;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.alert {
  padding: 12px 16px;
  border-radius: 8px;
  font-size: 14px;
}

.alert-error {
  background-color: #fee;
  color: #c33;
  border: 1px solid #fcc;
}
```

### Paso 2: El Servicio de Autenticación - El Corazón del Sistema

El servicio de autenticación es el cerebro de nuestro sistema de seguridad. Maneja el login, logout, y mantiene el estado de autenticación.

```typescript
// auth.service.ts
import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { BehaviorSubject, Observable, throwError } from 'rxjs';
import { tap, catchError, map } from 'rxjs/operators';
import { Router } from '@angular/router';

export interface User {
  id: number;
  email: string;
  name: string;
  role: string;
}

export interface AuthResponse {
  token: string;
  user: User;
  expiresIn: number;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly TOKEN_KEY = 'auth_token';
  private readonly USER_KEY = 'auth_user';
  private readonly TOKEN_EXPIRY_KEY = 'token_expiry';

  // BehaviorSubject para mantener el estado de autenticación
  private isAuthenticatedSubject = new BehaviorSubject<boolean>(this.hasValidToken());
  private currentUserSubject = new BehaviorSubject<User | null>(this.getSavedUser());

  public isAuthenticated$ = this.isAuthenticatedSubject.asObservable();
  public currentUser$ = this.currentUserSubject.asObservable();

  constructor(
    private http: HttpClient,
    private router: Router
  ) {
    // Verificar token periódicamente
    this.startTokenValidationTimer();
  }

  /**
   * Inicia sesión con credenciales
   */
  login(credentials: LoginCredentials): Observable<boolean> {
    return this.http.post<AuthResponse>('/api/auth/login', credentials)
      .pipe(
        tap(response => this.handleAuthentication(response)),
        map(() => true),
        catchError(error => {
          console.error('Error en login:', error);
          return throwError(() => new Error('Credenciales incorrectas'));
        })
      );
  }

  /**
   * Registra un nuevo usuario
   */
  register(userData: { name: string; email: string; password: string }): Observable<boolean> {
    return this.http.post<AuthResponse>('/api/auth/register', userData)
      .pipe(
        tap(response => this.handleAuthentication(response)),
        map(() => true),
        catchError(error => {
          console.error('Error en registro:', error);
          return throwError(() => new Error('Error al registrar usuario'));
        })
      );
  }

  /**
   * Cierra la sesión del usuario
   */
  logout(): void {
    // Limpiar almacenamiento local
    localStorage.removeItem(this.TOKEN_KEY);
    localStorage.removeItem(this.USER_KEY);
    localStorage.removeItem(this.TOKEN_EXPIRY_KEY);

    // Actualizar estado
    this.isAuthenticatedSubject.next(false);
    this.currentUserSubject.next(null);

    // Redirigir al login
    this.router.navigate(['/login']);
  }

  /**
   * Verifica si el usuario está autenticado
   */
  isAuthenticated(): boolean {
    return this.isAuthenticatedSubject.value;
  }

  /**
   * Obtiene el usuario actual
   */
  getCurrentUser(): User | null {
    return this.currentUserSubject.value;
  }

  /**
   * Obtiene el token actual
   */
  getToken(): string | null {
    const token = localStorage.getItem(this.TOKEN_KEY);
    const expiry = localStorage.getItem(this.TOKEN_EXPIRY_KEY);

    if (!token || !expiry) {
      return null;
    }

    // Verificar si el token ha expirado
    if (Date.now() > parseInt(expiry, 10)) {
      this.logout(); // Token expirado, cerrar sesión
      return null;
    }

    return token;
  }

  /**
   * Verifica si el usuario tiene un rol específico
   */
  hasRole(role: string): boolean {
    const user = this.getCurrentUser();
    return user ? user.role === role : false;
  }

  /**
   * Refresca el token si está próximo a expirar
   */
  refreshToken(): Observable<boolean> {
    const token = this.getToken();
    if (!token) {
      return throwError(() => new Error('No hay token para refrescar'));
    }

    return this.http.post<AuthResponse>('/api/auth/refresh', { token })
      .pipe(
        tap(response => this.handleAuthentication(response)),
        map(() => true),
        catchError(error => {
          this.logout();
          return throwError(() => new Error('Error al refrescar token'));
        })
      );
  }

  /**
   * Maneja la respuesta de autenticación
   */
  private handleAuthentication(response: AuthResponse): void {
    const expiryTime = Date.now() + (response.expiresIn * 1000);

    // Guardar en localStorage
    localStorage.setItem(this.TOKEN_KEY, response.token);
    localStorage.setItem(this.USER_KEY, JSON.stringify(response.user));
    localStorage.setItem(this.TOKEN_EXPIRY_KEY, expiryTime.toString());

    // Actualizar estado
    this.isAuthenticatedSubject.next(true);
    this.currentUserSubject.next(response.user);

    // Reiniciar timer de validación
    this.startTokenValidationTimer();
  }

  /**
   * Verifica si hay un token válido guardado
   */
  private hasValidToken(): boolean {
    const token = localStorage.getItem(this.TOKEN_KEY);
    const expiry = localStorage.getItem(this.TOKEN_EXPIRY_KEY);

    if (!token || !expiry) {
      return false;
    }

    return Date.now() < parseInt(expiry, 10);
  }

  /**
   * Obtiene el usuario guardado
   */
  private getSavedUser(): User | null {
    const userStr = localStorage.getItem(this.USER_KEY);
    if (!userStr) return null;

    try {
      return JSON.parse(userStr);
    } catch {
      return null;
    }
  }

  /**
   * Inicia el timer para validar el token periódicamente
   */
  private startTokenValidationTimer(): void {
    // Verificar cada 5 minutos si el token está próximo a expirar
    setTimeout(() => {
      const expiry = localStorage.getItem(this.TOKEN_EXPIRY_KEY);
      if (expiry) {
        const timeLeft = parseInt(expiry, 10) - Date.now();
        // Si quedan menos de 10 minutos, refrescar
        if (timeLeft < 10 * 60 * 1000) {
          this.refreshToken().subscribe();
        }
      }
      this.startTokenValidationTimer();
    }, 5 * 60 * 1000); // 5 minutos
  }
}
```

### Paso 3: Los Guards - Los Porteros de las Rutas

Los guards son como los porteros de un club exclusivo. Deciden quién puede entrar a cada zona de tu aplicación.

```typescript
// auth.guard.ts
import { Injectable } from '@angular/core';
import { CanActivate, Router, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Observable, of } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { AuthService } from '../services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<boolean> | Promise<boolean> | boolean {

    if (this.authService.isAuthenticated()) {
      // Verificar roles si es necesario
      const requiredRole = route.data['role'];
      if (requiredRole && !this.authService.hasRole(requiredRole)) {
        this.router.navigate(['/unauthorized']);
        return false;
      }
      return true;
    }

    // Usuario no autenticado, redirigir al login
    this.router.navigate(['/login'], {
      queryParams: { returnUrl: state.url }
    });
    return false;
  }
}

// role.guard.ts - Guard para roles específicos
@Injectable({
  providedIn: 'root'
})
export class RoleGuard implements CanActivate {

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  canActivate(route: ActivatedRouteSnapshot): boolean {
    const requiredRole = route.data['role'];

    if (!requiredRole) {
      return true; // No se requiere rol específico
    }

    if (this.authService.hasRole(requiredRole)) {
      return true;
    }

    // Usuario no tiene el rol requerido
    this.router.navigate(['/unauthorized']);
    return false;
  }
}

// login.guard.ts - Evita que usuarios logueados vayan al login
@Injectable({
  providedIn: 'root'
})
export class LoginGuard implements CanActivate {

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  canActivate(): boolean {
    if (this.authService.isAuthenticated()) {
      this.router.navigate(['/dashboard']);
      return false;
    }
    return true;
  }
}
```

### Paso 4: Los Interceptores - Los Mensajeros Silenciosos

Los interceptores son como mensajeros que automáticamente incluyen tu "invitación" (token) en cada petición HTTP que haces.

```typescript
// auth.interceptor.ts
import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, switchMap } from 'rxjs/operators';
import { AuthService } from '../services/auth.service';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  private isRefreshing = false;

  constructor(private authService: AuthService) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    // Agregar token a la petición si existe
    const token = this.authService.getToken();
    if (token) {
      request = this.addTokenToRequest(request, token);
    }

    return next.handle(request).pipe(
      catchError((error: HttpErrorResponse) => {
        // Si es un error 401 (no autorizado), intentar refrescar el token
        if (error.status === 401 && !this.isRefreshing) {
          return this.handle401Error(request, next);
        }

        return throwError(() => error);
      })
    );
  }

  /**
   * Agrega el token de autorización a la petición
   */
  private addTokenToRequest(request: HttpRequest<any>, token: string): HttpRequest<any> {
    return request.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`
      }
    });
  }

  /**
   * Maneja errores 401 intentando refrescar el token
   */
  private handle401Error(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (!this.isRefreshing) {
      this.isRefreshing = true;

      return this.authService.refreshToken().pipe(
        switchMap(() => {
          this.isRefreshing = false;
          const newToken = this.authService.getToken();
          if (newToken) {
            return next.handle(this.addTokenToRequest(request, newToken));
          }
          return next.handle(request);
        }),
        catchError(() => {
          this.isRefreshing = false;
          this.authService.logout();
          return throwError(() => new Error('Sesión expirada'));
        })
      );
    }

    return next.handle(request);
  }
}

// error.interceptor.ts - Interceptor para manejar errores globalmente
@Injectable()
export class ErrorInterceptor implements HttpInterceptor {

  constructor(private authService: AuthService) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(request).pipe(
      catchError((error: HttpErrorResponse) => {
        let errorMessage = 'Ha ocurrido un error desconocido';

        if (error.error instanceof ErrorEvent) {
          // Error del lado del cliente
          errorMessage = error.error.message;
        } else {
          // Error del lado del servidor
          switch (error.status) {
            case 400:
              errorMessage = 'Datos inválidos';
              break;
            case 401:
              errorMessage = 'No autorizado';
              this.authService.logout();
              break;
            case 403:
              errorMessage = 'Acceso denegado';
              break;
            case 404:
              errorMessage = 'Recurso no encontrado';
              break;
            case 500:
              errorMessage = 'Error interno del servidor';
              break;
            default:
              errorMessage = `Error ${error.status}: ${error.message}`;
          }
        }

        console.error('Error HTTP:', errorMessage);
        return throwError(() => new Error(errorMessage));
      })
    );
  }
}
```

### Paso 5: Configuración de Rutas Protegidas

Ahora vamos a configurar las rutas de nuestra aplicación con los guards apropiados.

```typescript
// app-routing.module.ts
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { RoleGuard } from './guards/role.guard';
import { LoginGuard } from './guards/login.guard';

const routes: Routes = [
  // Rutas públicas
  {
    path: 'login',
    loadComponent: () => import('./components/login/login.component').then(m => m.LoginComponent),
    canActivate: [LoginGuard] // Evita que usuarios logueados vayan al login
  },
  {
    path: 'register',
    loadComponent: () => import('./components/register/register.component').then(m => m.RegisterComponent),
    canActivate: [LoginGuard]
  },

  // Rutas protegidas
  {
    path: 'dashboard',
    loadComponent: () => import('./components/dashboard/dashboard.component').then(m => m.DashboardComponent),
    canActivate: [AuthGuard]
  },
  {
    path: 'profile',
    loadComponent: () => import('./components/profile/profile.component').then(m => m.ProfileComponent),
    canActivate: [AuthGuard]
  },
  {
    path: 'admin',
    loadComponent: () => import('./components/admin/admin.component').then(m => m.AdminComponent),
    canActivate: [AuthGuard, RoleGuard],
    data: { role: 'admin' } // Solo para administradores
  },
  {
    path: 'users',
    loadComponent: () => import('./components/users/users.component').then(m => m.UsersComponent),
    canActivate: [AuthGuard, RoleGuard],
    data: { role: 'manager' } // Para managers o superiores
  },

  // Ruta por defecto
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },

  // Ruta 404
  {
    path: '**',
    loadComponent: () => import('./components/not-found/not-found.component').then(m => m.NotFoundComponent)
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
```

### Paso 6: Configuración Global de Interceptores

```typescript
// app.module.ts
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

// Interceptores
import { AuthInterceptor } from './interceptors/auth.interceptor';
import { ErrorInterceptor } from './interceptors/error.interceptor';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: ErrorInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
```

---

## Mejores Prácticas de Seguridad

### 1. **Nunca Confíes Solo en el Frontend**

```typescript
// ❌ MAL: Verificación solo en el frontend
canActivate(): boolean {
  const user = this.authService.getCurrentUser();
  if (user && user.role === 'admin') {
    return true;
  }
  return false;
}

// ✅ BIEN: Verificación en backend + frontend
canActivate(): Observable<boolean> {
  return this.http.get('/api/user/permissions').pipe(
    map(permissions => permissions.includes('admin')),
    catchError(() => of(false))
  );
}
```

### 2. **Uso Seguro de localStorage**

```typescript
// auth.service.ts - Método seguro para manejar tokens
private secureStorage = {
  setItem(key: string, value: string) {
    try {
      // En producción, considera usar HttpOnly cookies
      localStorage.setItem(key, value);
    } catch (error) {
      console.error('Error guardando en localStorage:', error);
    }
  },

  getItem(key: string): string | null {
    try {
      return localStorage.getItem(key);
    } catch (error) {
      console.error('Error leyendo de localStorage:', error);
      return null;
    }
  },

  removeItem(key: string) {
    try {
      localStorage.removeItem(key);
    } catch (error) {
      console.error('Error eliminando de localStorage:', error);
    }
  }
};
```

### 3. **Validación de Tokens en el Backend**

Asegúrate de que tu API valide los tokens en cada petición:

```typescript
// backend (Node.js/Express ejemplo)
const jwt = require('jsonwebtoken');

function authenticateToken(req, res, next) {
  const authHeader = req.headers['authorization'];
  const token = authHeader && authHeader.split(' ')[1];

  if (!token) {
    return res.status(401).json({ message: 'Token requerido' });
  }

  jwt.verify(token, process.env.JWT_SECRET, (err, user) => {
    if (err) {
      return res.status(403).json({ message: 'Token inválido' });
    }
    req.user = user;
    next();
  });
}

// Uso en rutas protegidas
app.get('/api/protected', authenticateToken, (req, res) => {
  res.json({ message: 'Acceso concedido', user: req.user });
});
```

### 4. **Manejo de Sesiones Expiradas**

```typescript
// auth.service.ts
handleExpiredSession() {
  // Mostrar notificación
  this.showNotification('Tu sesión ha expirado. Por favor, inicia sesión nuevamente.');

  // Limpiar estado
  this.logout();

  // Opcional: Intentar refresh automático
  // this.refreshToken().subscribe(success => {
  //   if (!success) this.logout();
  // });
}

private showNotification(message: string) {
  // Implementa tu sistema de notificaciones
  alert(message); // En producción usa un servicio de notificaciones
}
```

---

## Errores Comunes y Cómo Evitarlos

### ❌ **Error 1: Guardar Contraseñas en localStorage**
```typescript
// MAL
localStorage.setItem('password', userPassword);

// Por qué es malo:
// - Cualquiera puede ver localStorage
// - Persiste hasta que se borre manualmente
// - No es seguro

// ✅ BIEN: Solo guardar tokens
localStorage.setItem('auth_token', token);
```

### ❌ **Error 2: No Proteger Rutas Sensibles**
```typescript
// MAL
const routes: Routes = [
  { path: 'admin', component: AdminComponent } // Sin guard
];

// ✅ BIEN
const routes: Routes = [
  {
    path: 'admin',
    component: AdminComponent,
    canActivate: [AuthGuard, RoleGuard],
    data: { role: 'admin' }
  }
];
```

### ❌ **Error 3: Exposición de Datos Sensibles**
```typescript
// MAL: Enviar datos sensibles al frontend
{
  "user": {
    "id": 1,
    "email": "user@example.com",
    "password": "hashed_password", // ❌ Nunca
    "creditCard": "1234-5678-9012-3456" // ❌ Nunca
  }
}

// ✅ BIEN: Solo enviar datos necesarios
{
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe",
    "role": "user"
  }
}
```

### ❌ **Error 4: Tokens sin Expiración**
```typescript
// MAL
const token = jwt.sign({ userId: 1 }, 'secret'); // Sin expiración

// ✅ BIEN
const token = jwt.sign(
  { userId: 1 },
  'secret',
  { expiresIn: '24h' }
);
```

---

## Proyecto Práctico: Sistema de Autenticación Completo

Vamos a crear una aplicación completa con autenticación. Este proyecto te ayudará a entender cómo todos los conceptos encajan juntos.

### Paso 1: Configuración del Proyecto

```bash
ng new auth-app --routing --style=css
cd auth-app
ng generate service services/auth
ng generate guard guards/auth
ng generate guard guards/role
ng generate guard guards/login
ng generate interceptor interceptors/auth
ng generate interceptor interceptors/error
```

### Paso 2: Implementar el Servicio de Autenticación

Crea el `AuthService` completo como se mostró anteriormente.

### Paso 3: Crear Componentes de Autenticación

```bash
ng generate component components/login
ng generate component components/register
ng generate component components/dashboard
ng generate component components/profile
```

### Paso 4: Configurar Rutas Protegidas

Actualiza `app-routing.module.ts` con las rutas protegidas.

### Paso 5: Crear un Layout con Navegación

```typescript
// app.component.ts
import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService, User } from './services/auth.service';

@Component({
  selector: 'app-root',
  template: `
    <div class="app-container">
      <nav *ngIf="isAuthenticated$ | async" class="navbar">
        <div class="nav-brand">
          <h3>Mi App</h3>
        </div>
        <div class="nav-links">
          <a routerLink="/dashboard" routerLinkActive="active">Dashboard</a>
          <a routerLink="/profile" routerLinkActive="active">Perfil</a>
          <a *ngIf="currentUser$ | async as user"
             routerLink="/admin" routerLinkActive="active"
             *ngIf="user.role === 'admin'">Admin</a>
          <button (click)="logout()" class="btn-logout">Cerrar Sesión</button>
        </div>
      </nav>

      <main class="main-content">
        <router-outlet></router-outlet>
      </main>
    </div>
  `,
  styles: [`
    .app-container {
      min-height: 100vh;
      display: flex;
      flex-direction: column;
    }

    .navbar {
      background: #2c3e50;
      color: white;
      padding: 1rem 2rem;
      display: flex;
      justify-content: space-between;
      align-items: center;
      box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }

    .nav-brand h3 {
      margin: 0;
      color: #ecf0f1;
    }

    .nav-links {
      display: flex;
      align-items: center;
      gap: 1rem;
    }

    .nav-links a {
      color: #bdc3c7;
      text-decoration: none;
      padding: 0.5rem 1rem;
      border-radius: 4px;
      transition: all 0.3s ease;
    }

    .nav-links a:hover,
    .nav-links a.active {
      background: #34495e;
      color: white;
    }

    .btn-logout {
      background: #e74c3c;
      color: white;
      border: none;
      padding: 0.5rem 1rem;
      border-radius: 4px;
      cursor: pointer;
      transition: background 0.3s ease;
    }

    .btn-logout:hover {
      background: #c0392b;
    }

    .main-content {
      flex: 1;
      padding: 2rem;
    }
  `]
})
export class AppComponent implements OnInit {
  isAuthenticated$: Observable<boolean>;
  currentUser$: Observable<User | null>;

  constructor(private authService: AuthService) {
    this.isAuthenticated$ = this.authService.isAuthenticated$;
    this.currentUser$ = this.authService.currentUser$;
  }

  ngOnInit() {
    // Verificar autenticación al iniciar la app
    if (this.authService.isAuthenticated()) {
      // Opcional: Validar token con el servidor
    }
  }

  logout() {
    this.authService.logout();
  }
}
```

### Paso 6: Probar el Sistema

1. **Regístrate** como nuevo usuario
2. **Inicia sesión** con tus credenciales
3. **Navega** por las rutas protegidas
4. **Verifica** que no puedas acceder sin autenticación
5. **Cierra sesión** y confirma que te redirige al login

---

## Reflexiones Finales

La autenticación es como construir una fortaleza digital. Has aprendido a:

- 🏰 **Construir murallas**: Proteger rutas con guards
- 🔑 **Gestionar llaves**: Manejar tokens JWT de forma segura
- 📨 **Enviar mensajeros**: Usar interceptores para comunicación automática
- 👮 **Poner porteros**: Controlar acceso basado en roles
- 🛡️ **Defender contra ataques**: Implementar mejores prácticas de seguridad

Recuerda que la seguridad es un proceso continuo. Mantén tus dependencias actualizadas, revisa regularmente tu código en busca de vulnerabilidades, y siempre valida tanto en frontend como en backend.

En el próximo capítulo, construiremos una aplicación CRUD completa que utilizará todo lo que hemos aprendido sobre autenticación. ¿Estás listo para el desafío?

---

## 📚 Recursos Adicionales

- [JWT.io](https://jwt.io/) - Herramientas para JWT
- [OWASP Authentication Cheat Sheet](https://owasp.org/www-project-cheat-sheets/cheatsheets/Authentication_Cheat_Sheet.html)
- [Angular Authentication Guide](https://angular.io/guide/security)
- [Web Security Best Practices](https://owasp.org/www-project-top-ten/)

## 🎯 Checklist de Implementación

- [ ] Servicio de autenticación con manejo de tokens
- [ ] Guards para protección de rutas
- [ ] Interceptores para inyección automática de tokens
- [ ] Componentes de login y registro
- [ ] Manejo de errores y sesiones expiradas
- [ ] Validación de roles y permisos
- [ ] Pruebas de seguridad
- [ ] Documentación de la API de autenticación
