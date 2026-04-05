# Capítulo 25: Variables de Entorno - Los Ajustes de tu Aplicación

Imagina que tienes un coche deportivo. En la ciudad usas el modo económico para ahorrar gasolina, pero en la autopista cambias a modo deportivo para máxima performance. En Angular, las variables de entorno son esos "modos" que ajustan tu aplicación según dónde se ejecute: desarrollo, pruebas, o producción.

Este capítulo te enseñará a configurar tu aplicación para que se comporte correctamente en cada ambiente.

## ¿Qué Aprenderás en Este Capítulo?

En este capítulo, exploraremos:
- La importancia de los entornos en el desarrollo web
- Cómo configurar variables de entorno en Angular
- La diferencia entre archivos de desarrollo y producción
- Mejores prácticas de seguridad
- Configuración avanzada con múltiples entornos
- Errores comunes y cómo evitarlos

## La Filosofía de los Entornos

Piensa en tu aplicación como un actor que debe adaptarse a diferentes escenarios. En el teatro (desarrollo), puede ensayar con libertad. En los ensayos generales (staging), se prepara para el público. En la función principal (producción), debe dar lo mejor de sí.

Los entornos permiten que tu aplicación use diferentes configuraciones sin cambiar el código fuente.

### Analogía del Restaurante

- **Desarrollo**: La cocina de pruebas donde experimentas con nuevas recetas
- **Staging**: El restaurante de práctica donde los chefs se entrenan
- **Producción**: El restaurante real donde atiendes a los clientes

## ¿Por Qué Necesitas Variables de Entorno?

### El Problema Sin Entornos

```typescript
// ❌ Código hardcoded - ¡Peligroso!
export class ApiService {
  private apiUrl = 'https://api.miproyecto.com'; // ¿Qué pasa en desarrollo?

  constructor(private http: HttpClient) {}

  getUsers() {
    return this.http.get(`${this.apiUrl}/users`);
  }
}
```

**Problemas:**
- No puedes desarrollar localmente
- Un cambio afecta todos los ambientes
- Información sensible expuesta
- Dificulta el trabajo en equipo

### La Solución: Variables de Entorno

```typescript
// ✅ Configuración por ambiente
import { environment } from '../environments/environment';

export class ApiService {
  private apiUrl = environment.apiUrl; // Automáticamente correcto

  constructor(private http: HttpClient) {}

  getUsers() {
    return this.http.get(`${this.apiUrl}/users`);
  }
}
```

## Configuración Básica de Entornos

### Estructura de Archivos

Angular crea automáticamente la carpeta `src/environments/` con dos archivos:

```
src/
  environments/
    environment.ts          // Desarrollo
    environment.prod.ts     // Producción
```

### Archivo de Desarrollo

```typescript
// src/environments/environment.ts
export const environment = {
  production: false,
  name: 'desarrollo',

  // APIs
  apiUrl: 'http://localhost:3000/api',
  websocketUrl: 'ws://localhost:3000',

  // Configuración de desarrollo
  enableDebugTools: true,
  logLevel: 'debug',

  // Servicios externos (llaves públicas)
  firebaseConfig: {
    apiKey: "dev-api-key-publica",
    authDomain: "dev-app.firebaseapp.com",
    projectId: "dev-project"
  },

  // Features flags
  features: {
    newDashboard: true,
    betaFeatures: true,
    analytics: false
  }
};
```

### Archivo de Producción

```typescript
// src/environments/environment.prod.ts
export const environment = {
  production: true,
  name: 'producción',

  // APIs de producción
  apiUrl: 'https://api.miproyecto.com',
  websocketUrl: 'wss://api.miproyecto.com',

  // Configuración de producción
  enableDebugTools: false,
  logLevel: 'error',

  // Servicios externos de producción
  firebaseConfig: {
    apiKey: "prod-api-key-publica",
    authDomain: "app.firebaseapp.com",
    projectId: "prod-project"
  },

  // Features flags
  features: {
    newDashboard: true,
    betaFeatures: false,
    analytics: true
  }
};
```

## Cómo Angular Selecciona el Archivo Correcto

### Configuración en angular.json

El archivo `angular.json` contiene la "receta" para construir tu aplicación:

```json
{
  "projects": {
    "mi-app": {
      "architect": {
        "build": {
          "configurations": {
            "production": {
              "fileReplacements": [
                {
                  "replace": "src/environments/environment.ts",
                  "with": "src/environments/environment.prod.ts"
                }
              ]
            }
          }
        }
      }
    }
  }
}
```

**¿Qué significa esto?**
- En desarrollo: usa `environment.ts`
- En producción (`ng build --configuration=production`): reemplaza `environment.ts` con `environment.prod.ts`

### Comandos de Build

```bash
# Desarrollo (usa environment.ts)
ng build
ng serve

# Producción (usa environment.prod.ts)
ng build --configuration=production

# Personalizado (puedes crear más)
ng build --configuration=staging
```

## Uso Avanzado: Múltiples Entornos

### Creando un Entorno de Staging

1. **Crear el archivo de staging:**

```typescript
// src/environments/environment.staging.ts
export const environment = {
  production: false,
  name: 'staging',

  apiUrl: 'https://staging-api.miproyecto.com',
  websocketUrl: 'wss://staging-api.miproyecto.com',

  enableDebugTools: true,
  logLevel: 'warn',

  firebaseConfig: {
    apiKey: "staging-api-key",
    authDomain: "staging-app.firebaseapp.com",
    projectId: "staging-project"
  },

  features: {
    newDashboard: true,
    betaFeatures: true,
    analytics: true
  }
};
```

2. **Configurar angular.json:**

```json
{
  "projects": {
    "mi-app": {
      "architect": {
        "build": {
          "configurations": {
            "production": {
              "fileReplacements": [
                {
                  "replace": "src/environments/environment.ts",
                  "with": "src/environments/environment.prod.ts"
                }
              ]
            },
            "staging": {
              "fileReplacements": [
                {
                  "replace": "src/environments/environment.ts",
                  "with": "src/environments/environment.staging.ts"
                }
              ]
            }
          }
        }
      }
    }
  }
}
```

3. **Usar el nuevo entorno:**

```bash
# Construir para staging
ng build --configuration=staging

# Servir con configuración de staging
ng serve --configuration=staging
```

## Seguridad: Lo Que NUNCA Debes Hacer

### ❌ Errores Comunes de Seguridad

```typescript
// ¡NUNCA hagas esto!
export const environment = {
  production: false,

  // ❌ Llaves privadas en el frontend
  databasePassword: 'supersecreto123',
  jwtSecret: 'mi-clave-super-secreta',

  // ❌ Tokens de API privados
  stripeSecretKey: 'sk_live_...',
  awsSecretAccessKey: 'AKIA...',

  // ❌ Información sensible
  adminPassword: 'admin123'
};
```

### ✅ Lo Que SÍ Debes Hacer

```typescript
export const environment = {
  production: false,

  // ✅ Llaves públicas (visibles en el navegador)
  firebaseConfig: {
    apiKey: "public-key-visible",
    authDomain: "app.firebaseapp.com"
  },

  // ✅ URLs públicas
  apiUrl: 'https://api.miproyecto.com',

  // ✅ Configuración de features
  features: {
    premiumFeatures: true
  }
};
```

### ¿Por Qué es Seguro?

- El código JavaScript es visible en el navegador
- Cualquier usuario puede ver las variables de entorno
- Las llaves privadas deben estar en el **servidor backend**
- Usa variables de entorno del servidor para información sensible

## Patrones Avanzados de Configuración

### Configuración Dinámica

```typescript
// services/config.service.ts
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {
  private config = environment;

  // Getters tipados
  get apiUrl(): string {
    return this.config.apiUrl;
  }

  get isProduction(): boolean {
    return this.config.production;
  }

  get logLevel(): string {
    return this.config.logLevel;
  }

  // Métodos de configuración
  isFeatureEnabled(feature: keyof typeof this.config.features): boolean {
    return this.config.features[feature];
  }
}
```

### Logging Condicional

```typescript
// services/logger.service.ts
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LoggerService {
  private levels = ['error', 'warn', 'info', 'debug'];

  log(level: string, message: string, data?: any) {
    if (this.shouldLog(level)) {
      console.log(`[${level.toUpperCase()}] ${message}`, data);
    }
  }

  private shouldLog(level: string): boolean {
    const currentLevelIndex = this.levels.indexOf(environment.logLevel);
    const messageLevelIndex = this.levels.indexOf(level);
    return messageLevelIndex <= currentLevelIndex;
  }
}
```

### Guards Basados en Entorno

```typescript
// guards/feature.guard.ts
import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class FeatureGuard implements CanActivate {
  canActivate(): boolean {
    // Solo permite acceso a features beta en desarrollo/staging
    return environment.features.betaFeatures;
  }
}
```

## Un Ejemplo Completo: Aplicación Multi-entorno

Vamos a crear una aplicación que demuestre el uso completo de entornos:

```typescript
// services/api.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private http: HttpClient) {}

  getUsers() {
    return this.http.get(`${environment.apiUrl}/users`);
  }

  getProducts() {
    return this.http.get(`${environment.apiUrl}/products`);
  }
}

// services/analytics.service.ts
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AnalyticsService {
  private enabled = environment.features.analytics;

  trackEvent(event: string, data?: any) {
    if (!this.enabled) return;

    // Enviar a servicio de analytics
    console.log('Tracking:', event, data);
  }
}

// components/environment-info.component.ts
import { Component } from '@angular/core';
import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-environment-info',
  template: `
    <div class="env-info" *ngIf="!environment.production">
      <strong>Entorno: {{ environment.name }}</strong>
      <p>API: {{ environment.apiUrl }}</p>
      <p>Debug: {{ environment.enableDebugTools }}</p>
    </div>
  `
})
export class EnvironmentInfoComponent {
  environment = environment;
}
```

```html
<!-- app.component.html -->
<app-environment-info></app-environment-info>

<div class="app-content">
  <h1>Mi Aplicación</h1>

  <!-- Features condicionales -->
  <div *ngIf="configService.isFeatureEnabled('newDashboard')">
    <app-new-dashboard></app-new-dashboard>
  </div>

  <div *ngIf="configService.isFeatureEnabled('betaFeatures')">
    <app-beta-features></app-beta-features>
  </div>
</div>
```

## Configuración de Build Avanzada

### Scripts en package.json

```json
{
  "scripts": {
    "build:dev": "ng build",
    "build:staging": "ng build --configuration=staging",
    "build:prod": "ng build --configuration=production",
    "serve:staging": "ng serve --configuration=staging",
    "deploy:staging": "npm run build:staging && firebase deploy --project staging",
    "deploy:prod": "npm run build:prod && firebase deploy --project production"
  }
}
```

### Variables de Entorno del Sistema

Para información realmente sensible, usa variables del sistema operativo:

```typescript
// Solo para desarrollo local
export const environment = {
  production: false,

  // Variables del sistema (solo locales)
  localApiPort: process.env['API_PORT'] || '3000',
  databaseUrl: process.env['DATABASE_URL'],

  // Configuración normal
  apiUrl: `http://localhost:${process.env['API_PORT'] || '3000'}/api`
};
```

**Nota:** `process.env` solo funciona en Node.js, no en el navegador. Para producción, estas variables deben venir del servidor.

## Mejores Prácticas

### ✅ Hazlo:
- Mantén entornos separados y bien documentados
- Usa solo información pública en los archivos de entorno
- Crea servicios de configuración centralizados
- Documenta qué variables necesita cada entorno
- Usa feature flags para controlar funcionalidades

### ❌ Evítalo:
- Guardar secretos en el frontend
- Hardcodear URLs o configuraciones
- Mezclar configuraciones de diferentes entornos
- Cambiar archivos de entorno manualmente
- Usar `process.env` para información que va al bundle

## Errores Comunes y Cómo Evitarlos

1. **Archivo incorrecto**: Siempre importa desde `../environments/environment`, nunca directamente de `environment.prod.ts`

2. **Variables undefined**: Asegúrate de que todas las variables estén definidas en todos los archivos de entorno

3. **Build incorrecto**: Verifica que `fileReplacements` esté configurado correctamente en `angular.json`

4. **Secrets expuestos**: Recuerda que todo en `environment.ts` es visible en el navegador

5. **Configuración runtime**: Los archivos de entorno son estáticos; no cambian en tiempo de ejecución

6. **Merging manual**: No intentes combinar entornos manualmente; deja que Angular lo haga

## Mini Práctica: Sistema de Entornos Completo

1. **Crea tres entornos**: desarrollo, staging, y producción con diferentes configuraciones
2. **Implementa un servicio de logging** que respete el `logLevel` del entorno
3. **Crea feature flags** para mostrar/ocultar funcionalidades según el entorno
4. **Configura un servicio de configuración** centralizado con getters tipados
5. **Agrega un indicador visual** que muestre en qué entorno estás (solo en no-producción)
6. **Configura scripts de build** para cada entorno en `package.json`

## Resumen del Capítulo

Las variables de entorno son el sistema nervioso de tu aplicación Angular. Te permiten adaptar el comportamiento de tu app según dónde se ejecute, manteniendo el código limpio y la configuración organizada.

Recuerda: desarrollo para experimentar, staging para probar, producción para servir. Mantén la información sensible en el servidor, y usa los entornos para todo lo demás.

En el próximo capítulo, exploraremos la autenticación, el guardián que protege el acceso a tu aplicación.
