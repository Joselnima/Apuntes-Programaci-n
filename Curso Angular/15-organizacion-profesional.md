# 15 - Organización profesional de carpetas

## ¿Qué aprenderás en este módulo?

En este módulo aprenderás:
- cómo organizar un proyecto Angular cuando crece
- qué va en `core/`, `shared/`, `features/`, `pages/` y `app/`
- cuándo crear un módulo de funcionalidad y cuándo usar un componente compartido
- buenas prácticas para mantener el código legible y escalable

---

## ¿Para qué sirve esta organización?

Una organización profesional ayuda a mantener el código comprensible, reutilizable y fácil de escalar.
Con una estructura clara, el equipo trabaja más rápido y los cambios se hacen con menos riesgo de romper otras partes.

---

## Cómo organizar carpetas en Angular

Una estructura común y útil es:

- `src/app/core/`
  - servicios singleton como `auth.service.ts`, `http.interceptor.ts`
  - guardas, interceptores y clases de utilidad globales
- `src/app/shared/`
  - componentes pequeños reutilizables como botones, tarjetas y directivas
  - pipes y modelos comunes
- `src/app/features/`
  - módulos por dominio, por ejemplo `users/`, `products/`, `orders/`
  - cada feature con su propio módulo, rutas y componentes
- `src/app/pages/`
  - páginas de alto nivel que combinan features y layout
- `src/app/app-routing.module.ts`
  - solo rutas principales y lazy loading

### Ejemplo de estructura

```bash
src/app/
  core/
    services/
      auth.service.ts
      session.service.ts
    guards/
      auth.guard.ts
    interceptors/
      auth.interceptor.ts
  shared/
    components/
      card/
        card.component.ts
    pipes/
      date-format.pipe.ts
    directives/
      highlight.directive.ts
  features/
    users/
      users.module.ts
      users-routing.module.ts
      list/
        user-list.component.ts
      edit/
        user-edit.component.ts
  pages/
    home/
      home.component.ts
    settings/
      settings.component.ts
  app.component.ts
  app.module.ts
  app-routing.module.ts
```

---

## Ejemplo explicado

Supón que tu app necesita un formulario de usuario y un listado.

- `core/` guarda el servicio `UserService` y la lógica de autenticación.
- `shared/` guarda un componente `UserCard` que se puede usar en varias páginas.
- `features/users/` guarda todo lo relacionado con usuarios: listado, edición, creación.
- `app-routing.module.ts` carga `UsersModule` de forma perezosa si el usuario entra a `/users`.

Esto separa responsabilidades y evita mezclar lógica de negocio con UI.

---

## Errores comunes

1. Guardar todo en `app/` y no separar por dominio.
2. Usar `shared/` como basura para imports sin criterio.
3. Poner servicios que dependen de features dentro de `core/`.
4. Incluir rutas internas en `app-routing.module.ts` en lugar de usar lazy loading.

---

## Mini práctica

1. Crea la estructura de carpetas en un proyecto Angular nuevo.
2. Mueve un servicio global a `core/` y un componente reutilizable a `shared/`.
3. Crea un feature `products/` con su propio módulo y ruta.
4. Explica por qué cada archivo está en su carpeta.

---

## Resumen

La organización profesional no es solo estética: es una forma de reducir la complejidad.
Si puedes decir qué debe vivir en `core/`, `shared/`, `features/` y `pages/`, ya dominaste la idea.
