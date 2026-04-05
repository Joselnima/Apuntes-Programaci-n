# Módulo 24 - Guards e Interceptors

## ¿Qué aprenderás en este módulo?

En este módulo aprenderás:
- cómo proteger rutas con `CanActivate` y `CanDeactivate`
- cómo controlar acceso por roles con Guardas
- cómo interceptar peticiones HTTP con `HttpInterceptor`
- cómo manejar errores, tokens y carga global desde una sola capa

---

## ¿Para qué sirven los Guards?

Los Guards se usan para decidir si una ruta puede activarse o no.
Sirven para proteger páginas privadas, validar permisos y evitar pérdidas de datos cuando el usuario navega.

---

## ¿Para qué sirven los Interceptors?

Los Interceptors permiten modificar todas las peticiones y respuestas HTTP.
Sirven para agregar tokens, mostrar un loader global, manejar errores y renovar sesiones automáticamente.

---

## Cómo proteger rutas con Guards

### AuthGuard básico

```typescript
@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {
  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  canActivate(): Observable<boolean> {
    return this.authService.isAuthenticated().pipe(
      map(isAuth => {
        if (isAuth) return true;
        this.router.navigate(['/login']);
        return false;
      })
    );
  }
}
```

```typescript
const routes: Routes = [
  { path: 'login', component: LoginComponent },
  {
    path: 'dashboard',
    component: DashboardComponent,
    canActivate: [AuthGuard]
  }
];
```

### RoleGuard para admin

```typescript
@Injectable({ providedIn: 'root' })
export class AdminGuard implements CanActivate {
  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  canActivate(): Observable<boolean> {
    return this.authService.getUser().pipe(
      map(user => {
        if (user?.role === 'admin') return true;
        this.router.navigate(['/forbidden']);
        return false;
      })
    );
  }
}
```

```typescript
{
  path: 'admin',
  component: AdminComponent,
  canActivate: [AuthGuard, AdminGuard]
}
```

### CanDeactivate para formularios

```typescript
@Injectable({ providedIn: 'root' })
export class UnsavedChangesGuard implements CanDeactivate<any> {
  canDeactivate(component: any): boolean {
    if (component.form?.dirty) {
      return confirm('Tienes cambios sin guardar. ¿Salir igual?');
    }
    return true;
  }
}
```

```typescript
{
  path: 'edit',
  component: EditComponent,
  canDeactivate: [UnsavedChangesGuard]
}
```

---

## Cómo usar Interceptors

### AuthInterceptor

```typescript
@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const token = localStorage.getItem('token');
    if (!token) return next.handle(req);

    const authReq = req.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`
      }
    });
    return next.handle(authReq);
  }
}
```

### ErrorInterceptor

```typescript
@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
  constructor(private authService: AuthService, private notifyService: NotifyService) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(
      catchError(error => {
        if (error.status === 401) {
          this.authService.logout();
          this.notifyService.error('Sesión expirada');
        } else if (error.status === 403) {
          this.notifyService.error('No tienes permisos');
        } else if (error.status === 500) {
          this.notifyService.error('Error del servidor');
        }
        return throwError(() => error);
      })
    );
  }
}
```

### LoadingInterceptor

```typescript
@Injectable()
export class LoadingInterceptor implements HttpInterceptor {
  constructor(private loadingService: LoadingService) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    this.loadingService.show();
    return next.handle(req).pipe(
      finalize(() => this.loadingService.hide())
    );
  }
}
```

### RefreshTokenInterceptor

```typescript
@Injectable()
export class RefreshTokenInterceptor implements HttpInterceptor {
  private isRefreshing = false;
  private refreshTokenSubject = new Subject<string>();

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(
      catchError(error => {
        if (error.status === 401 && !this.isRefreshing) {
          this.isRefreshing = true;
          return this.authService.refreshToken().pipe(
            tap(response => {
              localStorage.setItem('token', response.token);
              this.refreshTokenSubject.next(response.token);
              this.isRefreshing = false;
            }),
            switchMap(() => next.handle(req)),
            catchError(() => {
              this.authService.logout();
              return throwError(() => error);
            })
          );
        }
        return throwError(() => error);
      })
    );
  }
}
```

### Registrar los interceptors

```typescript
@NgModule({
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: ErrorInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: LoadingInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: RefreshTokenInterceptor, multi: true }
  ]
})
export class AppModule {}
```

---

## Buenas prácticas con Guards e Interceptors

- Protege rutas en el módulo correcto.
- No pongas lógica de permisos en componentes.
- Ordena interceptors pensando en autenticación primero y errores después.
- Nunca confíes solo en el frontend para seguridad.
- Usa `CanDeactivate` cuando haya formularios con cambios.

---

## Errores comunes

1. No registrar un interceptor con `multi: true`.
2. Usar `RouterModule.forRoot()` en módulos secundarios.
3. No devolver `true` o `false` en los guards.
4. Poner lógica de token en cada componente en lugar de un interceptor.
5. No manejar condiciones de refresco de token.

---

## Mini práctica

1. Crea `AuthGuard` y protege la ruta `/dashboard`.
2. Añade `AdminGuard` y protege la ruta `/admin`.
3. Implementa `AuthInterceptor` que agrega el token a las peticiones.
4. Agrega `ErrorInterceptor` que muestra mensajes según el código HTTP.
5. Crea un `CanDeactivate` para una página de edición con formulario.

---

## Resumen

Guards e interceptors son las herramientas clave para controlar acceso y modificar peticiones HTTP en Angular.
Un buen uso de ellos separa seguridad, manejo de errores y UX del resto de tu app.

---

## Próximo

Módulo 25: Entornos - Dev, Staging, Prod
