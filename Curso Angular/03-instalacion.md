# Módulo 03 - Instalación de Angular

> **Setup completo - Desde cero hasta tu primer componente funcionando**

---

## ¿QUÉ necesitas instalar?

**Angular necesita una cadena de herramientas moderna:**

```
Tu código (TypeScript) 
    ↓
Angular CLI (herramienta)
    ↓
Node.js + npm (gestor de paquetes)
    ↓
VS Code (editor)
    ↓
Navegador (Chrome/Firefox)
```

**Analogía:**
- Carpintero = Angular
- Sierras/martillos = CLI, Node.js
- Mesa de trabajo = VS Code
- Resultado = Tu app web

---

## ¿PARA QUÉ cada herramienta?

| Herramienta | Para qué | Análogo |
|------------|----------|---------|
| **Node.js** | Ejecutar herramientas en tu PC | Motor del carro |
| **npm** | Descargar librerías | Gasolinera |
| **Angular CLI** | Crear proyectos, compilar | Caja de herramientas |
| **VS Code** | Escribir código | Cuaderno y pluma |
| **TypeScript** | JavaScript tipado | Idioma mejorado |
| **Webpack** | Empaquetar código (automático) | Empacar regalo |

---

## ¿CÓMO instalar? (Paso a paso)

### Paso 1: Instalar Node.js

**Ir a:** https://nodejs.org/ (versión LTS)

Descarga e instala (click siguiente, siguiente, siguiente)

**Verificar instalación:**

```bash
node --version
npm --version
```

Deberías ver versiones (ej: v18.19.0)

### Paso 2: Instalar Angular CLI

```bash
npm install -g @angular/cli
```

`-g` = global (en toda tu computadora)

**Verificar:**

```bash
ng version
```

### Paso 3: Crear tu primer proyecto

```bash
ng new mi-app
cd mi-app
```

**Preguntará:**
```
Routing? → No (por ahora)
Stylesheet → CSS
```

**¿Qué hace?**
- Crea carpeta `mi-app`
- Descarga 500+ dependencias
- Genera estructura del proyecto

⏳ Espera 5-10 minutos (primera vez es lenta)

### Paso 4: Iniciar servidor de desarrollo

```bash
ng serve
```

o

```bash
ng serve --open
```

**Resultado:**
```
✔ Compiled successfully.
Local:            http://localhost:4200
```

Abrirá automáticamente, verás logo de Angular ✅

### Paso 5: Editar en VS Code

```bash
code .
```

Abrirá VS Code en tu proyecto

---

## Estructura de Carpeta

```
mi-app/
├── src/
│   ├── app/
│   │   ├── app.component.ts      ← Tu principal componente
│   │   ├── app.component.html    ← Template
│   │   ├── app.component.css     ← Estilos
│   │   ├── app.component.spec.ts ← Tests
│   │   └── app.module.ts         ← Módulo
│   ├── main.ts                   ← Entrada
│   ├── index.html                ← HTML raíz
│   └── styles.css                ← CSS global
├── angular.json                  ← Configuración
├── tsconfig.json                 ← TypeScript config
└── package.json                  ← Dependencias
```

---

## Tu Primer Cambio

**Abrir:** `src/app/app.component.ts`

```typescript
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Mi Primera App en Angular';
}
```

**Abrir:** `src/app/app.component.html`

```html
<h1>¡Hola {{ title }}! 👋</h1>
<p>Bienvenido a Angular</p>
```

**Guardar (Ctrl+S)**

La app se actualiza automáticamente en el navegador ✨

---

## Comandos Útiles

| Comando | Para qué |
|---------|----------|
| `ng serve` | Iniciar desarrollo (http://localhost:4200) |
| `ng generate component nombre` | Crear componente |
| `ng build` | Compilar para producción |
| `ng test` | Ejecutar tests |
| `ng lint` | Verificar errores de código |

---

## Troubleshooting

### "ng command not found"
```bash
npm install -g @angular/cli
```

### "Puerto 4200 ya en uso"
```bash
ng serve --port 5000
```

### "npm ERR!"
```bash
npm cache clean --force
npm install
```

### "Cambios no se ven"
```bash
Ctrl+C (detener servidor)
ng serve
```

---

## ✅ Checklist

- [ ] Node.js instalado
- [ ] Angular CLI instalado
- [ ] Proyecto creado (`ng new mi-app`)
- [ ] Servidor ejecutando (`ng serve`)
- [ ] Página visible en http://localhost:4200
- [ ] Cambio en HTML reflejado en navegador
- [ ] VS Code abierto en la carpeta

Si todo OK, ¡estás listo para aprender Angular! 🚀

---

## Próximo Paso

Módulo 04: Tu primer componente profesional
