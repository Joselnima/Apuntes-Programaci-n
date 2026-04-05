# Módulo 04 - Tu Primer Proyecto Angular

> **De cero a app funcional en 10 minutos**

---

## ¿QUÉ vas a hacer?

**Crearás una aplicación Angular completa y funcionando, entendiendo cada carpeta y archivo**

Analogía:
- Constructor: `ng new` (prepara todo)
- Casa lista para habitar: Proyecto generado
- Ocupar la casa: Tu primer cambio visible

---

## ¿PARA QUÉ?

### Necesitas saber cómo:

```
1. Crear proyecto
2. Levantar servidor
3. Ver cambios en tiempo real
4. Entender estructura generada
5. Primeras líneas de código
```

---

## ¿CÓMO funciona?

### Paso 1: Crear el Proyecto

```bash
ng new mi-tienda
```

**Angular preguntará:**

```
? Would you like to add Angular routing? (y/N) → N (por ahora)
? Which stylesheet format would you like to use?
  CSS
  SCSS
  SASS
  LESS
→ CSS (selecciona)
```

**¿Qué hace `ng new`?**
- Descarga dependencias (500 MB 😅)
- Crea estructura de carpetas
- Configura TypeScript
- Prepara todo para desarrollo

⏳ Espera 5-10 minutos la primera vez

### Paso 2: Entrar a la Carpeta

```bash
cd mi-tienda
```

### Paso 3: Iniciar Servidor

```bash
ng serve
```

**Verás:**

```
✔ Compiled successfully.
✔ Built successfully.

Local:            http://localhost:4200
```

Abrirá automáticamente en tu navegador 🎉

### Paso 4: Tu Primer Cambio

**Abre:** `src/app/app.component.ts`

```typescript
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Mi primera app Angular!';  // ← Cambia esto
}
```

**Abre:** `src/app/app.component.html`

```html
<h1>{{ title }}</h1>  <!-- Muestra el título -->
<p>Bienvenido a Angular</p>
```

**Guarda (Ctrl+S) y mira el navegador → ¡Cambio instantáneo!** ✨

---

## Estructura Generada

```
mi-tienda/
├── src/
│   ├── app/
│   │   ├── app.component.ts      ← Tu componente principal
│   │   ├── app.component.html    ← Template (HTML)
│   │   ├── app.component.css     ← Estilos (CSS)
│   │   ├── app.component.spec.ts ← Tests (ignorar)
│   │   └── app.module.ts         ← Módulo (después)
│   │
│   ├── main.ts                   ← Punto de entrada
│   ├── index.html                ← HTML raíz (raro editar)
│   ├── styles.css                ← CSS global
│   ├── polyfills.ts              ← Compatibilidad
│   └── environments/             ← Dev vs Prod
│
├── angular.json                  ← Configuración build
├── tsconfig.json                 ← TypeScript config
├── package.json                  ← Dependencias
└── README.md                      ← Documentación
```

**¿Qué es cada cosa?**

| Archivo | Para qué |
|---------|----------|
| `app.component.ts` | Lógica del componente |
| `app.component.html` | Estructura (HTML) |
| `app.component.css` | Estilos (CSS) |
| `app.module.ts` | Configura componentes, servicios |
| `main.ts` | Inicia la app |
| `angular.json` | Build, puertos, configuración |

---

## Comandos Esenciales

```bash
ng serve                    # Iniciar dev (http://localhost:4200)
ng serve --port 3000       # Puerto diferente
ng serve --open            # Abrir navegador automáti
camente
ng generate component nombre  # Crear componente
ng build                    # Compilar para producción
ng test                     # Ejecutar tests
ng lint                     # Verificar código
```

---

## Hot Module Replacement (HMR) - Lo Más Cool

**Angular recompila y recarga SIN perder estado:**

```
1. Escribes código
2. Guardas (Ctrl+S)
3. Angular recompila
4. Navegador se actualiza
5. ¡Ves cambios al instante!
```

**Todo sin recargar manualmente.** 🚀

---

## Tu Primera Interactividad

**Ahora haremos un botón que haga algo:**

**app.component.ts:**

```typescript
export class AppComponent {
  title = 'Mi tienda Angular';
  contador = 0;
  
  incrementar() {
    this.contador++;
  }
}
```

**app.component.html:**

```html
<h1>{{ title }}</h1>

<button (click)="incrementar()">
  Clickeaste {{ contador }} veces
</button>
```

**app.component.css:**

```css
button {
  padding: 10px 20px;
  font-size: 16px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background: #0056b3;
}
```

**Guarda y prueba en el navegador** ✅

---

## Estructura Con Carpetas Personales

**Para un proyecto real, mejora la estructura:**

```
src/app/
├── components/           ← Componentes reutilizables
│   ├── header/
│   ├── footer/
│   └── sidebar/
├── pages/                ← Páginas principales
│   ├── home/
│   ├── products/
│   └── cart/
├── services/             ← Lógica compartida
│   ├── product.service.ts
│   └── cart.service.ts
├── models/               ← Interfaces TypeScript
│   └── product.ts
└── app.component.ts      ← Raíz
```

Haremos esto en próximos módulos ✅

---

## Debugging: Consola del Navegador

**Abre:** F12 en el navegador → Console

```javascript
// Aquí aparecen los console.log() que escribas en TypeScript
// Y errores si tu código se rompe
```

**Agrega logs en tu código:**

```typescript
export class AppComponent {
  contador = 0;
  
  incrementar() {
    this.contador++;
    console.log('Contador ahora es:', this.contador);
  }
}
```

Abre Console y haz clic en el botón.

---

## Checklist Primer Proyecto

- [ ] Angular CLI instalado (`ng version` funciona)
- [ ] Proyecto creado (`ng new mi-tienda`)
- [ ] Servidor ejecutando (`ng serve`)
- [ ] Página visible en http://localhost:4200
- [ ] Logo Angular aparece
- [ ] Cambios en HTML se ven al guardar
- [ ] Botón funciona
- [ ] Console sin errores rojo

Si todo ✅, estás listo para aprender Angular.

---

## Próximo

Módulo 05: Estructura de proyecto profesional

---

## Resumen

Si puedes explicarlo con tus palabras, vas bien.
Si solo lo reconoces cuando lo ves, todavía falta práctica.
