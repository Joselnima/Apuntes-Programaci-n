# Módulo 07 - CSS Desde Cero

> **Estilos profesionales - Hacer que se vea bonito**

---

## ¿QUÉ es CSS?

**CSS = Lenguaje para diseñar**

Flexbox es lo más importante. Grid para layouts complejos.

### Selectores

```css
p { color: black; }         /* Elemento */
.card { border: 1px solid; } /* Clase */
#header { background: blue; } /* ID */
```

### Propiedades Clave

```css
/* Layout */
display: flex;              /* Lado a lado */
justify-content: center;    /* Centrar horizontal */
align-items: center;        /* Centrar vertical */
gap: 20px;                  /* Espacio entre */

/* Espaciado */
margin: 10px;               /* Fuera */
padding: 20px;              /* Dentro */
border: 1px solid;          /* Borde */

/* Tipografía */
color: blue;
font-size: 16px;
font-weight: bold;

/* Responsive */
@media (max-width: 768px) { /* Móvil */ }
```

### Flexbox (LO MÁS IMPORTANTE)

```css
.container {
  display: flex;           /* Activa flexbox */
  justify-content: center; /* Centra horizontal */
  align-items: center;     /* Centra vertical */
  gap: 20px;              /* Espacio entre items */
  flex-wrap: wrap;        /* Saltar a nueva línea */
}

.item {
  flex: 1;                /* Crecer igual */
}
```

### Media Queries

```css
/* Desktop (defecto) */
.container { width: 1200px; }

/* Tablet */
@media (max-width: 768px) {
  .container { width: 100%; }
}

/* Móvil */
@media (max-width: 480px) {
  h1 { font-size: 20px; }
}
```

---

## En Angular

```typescript
@Component({
  template: `<h1>Título</h1>`,
  styles: [`
    h1 { color: blue; font-size: 32px; }
  `]
})
export class AppComponent {}
```

---

## Checklist

- [ ] Flexbox básico
- [ ] Media queries
- [ ] Responsive design
- [ ] Box model

---

## Próximo: TypeScript desde cero
