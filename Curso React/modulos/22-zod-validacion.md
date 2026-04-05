# Módulo 22 - Zod - Validación TypeScript

> **Validación de datos con TypeScript - seguridad total**

---

## ¿QUÉ es Zod?

**Zod = TypeScript + Validación en Runtime**

Es como tener un inspector de calidad:
- TypeScript verifica tipos en escritura
- Zod verifica datos cuando llegan (API, formularios, etc)

```typescript
// ❌ Sin Zod (confío ciegamente)
const usuario = await fetch("/api/usuario").then(r => r.json());
console.log(usuario.nombre.toUpperCase());  // ¿Existe .nombre? 🤔

// ✅ Con Zod (verifico todo)
const usuarioSchema = z.object({
  nombre: z.string(),
  email: z.string().email()
});

const usuario = usuarioSchema.parse(datos);
console.log(usuario.nombre.toUpperCase());  // Garantizado 100%
```

---

## ¿PARA QUÉ?

### Problema 1: Datos del servidor sin validar

```typescript
// ❌ API devuelve cualquier cosa
const response = await fetch("/api/usuarios");
const usuarios = await response.json();

// ¿Tiene "nombre"? ¿Es string? ¿Tiene "edad"? ¿Es number?
usuarios[0].nombre.toUpperCase();  // ¿CRASH?
```

### Problema 2: Datos del usuario sin verificar

```typescript
// ❌ Formulario - datos inseguros
const usuario = {
  nombre: datos.nombre,  // ¿Vacio? ¿Muy corto?
  email: datos.email,    // ¿Email válido?
  edad: datos.edad       // ¿Number? ¿Negativo?
};
```

### Problema 3: TypeScript solo valida en escritura

```typescript
interface Usuario {
  nombre: string;
  email: string;
}

// TypeScript dice OK en escritura
const usuario: Usuario = JSON.parse(miJSON);
// Pero en runtime, puede ser: { nombre: 123, email: null }
// ❌ TypeError cuando accesas usuario.nombre.toUpperCase()
```

**Zod lo resuelve:**

```typescript
// ✅ Define esquema
const usuarioSchema = z.object({
  nombre: z.string().min(3),
  email: z.string().email(),
  edad: z.number().positive().optional()
});

// Extraer tipo automáticamente
type Usuario = z.infer<typeof usuarioSchema>;

// Validar datos
try {
  const usuario = usuarioSchema.parse(datosDesconocidos);
  // Aquí usuario es 100% seguro
} catch (error) {
  console.error("Datos inválidos:", error.errors);
}
```

---

## ¿CÓMO funciona?

### 1. Setup

```bash
npm install zod
```

### 2. Tipos Básicos

```typescript
import { z } from "zod";

// String
const nombre = z.string();
nombre.parse("Juan");      // ✅ OK
// nombre.parse(123);      // ❌ Error

// Number
const edad = z.number();
edad.parse(25);            // ✅ OK
// edad.parse("25");       // ❌ Error

// Boolean
const activo = z.boolean();
activo.parse(true);        // ✅ OK

// Literal (valores específicos)
const rol = z.literal("admin");
rol.parse("admin");        // ✅ OK
// rol.parse("user");      // ❌ Error

// Enum
const estado = z.enum(["activo", "inactivo", "pendiente"]);
estado.parse("activo");           // ✅ OK
// estado.parse("desconocido");    // ❌ Error
```

### 3. Strings con Validaciones

```typescript
const email = z.string().email();
email.parse("juan@ejemplo.com");     // ✅ OK
// email.parse("no-es-email");       // ❌ Error

// Longitud
const contrasena = z.string().min(8).max(20);
contrasena.parse("1234567890");      // ✅ OK
// contrasena.parse("123");           // ❌ Error (muy corta)

// URL
const web = z.string().url();
web.parse("https://ejemplo.com");    // ✅ OK

// Regex personalizado
const telefono = z.string().regex(/^\d{10}$/);
telefono.parse("1234567890");        // ✅ OK

// Transformar
const nombreMayuscula = z.string().transform(s => s.toUpperCase());
nombreMayuscula.parse("juan");       // "JUAN"
```

### 4. Objects (Objetos)

```typescript
// Objeto simple
const usuarioSchema = z.object({
  nombre: z.string(),
  email: z.string().email(),
  edad: z.number().positive()
});

usuarioSchema.parse({
  nombre: "Juan",
  email: "juan@ejemplo.com",
  edad: 25
});  // ✅ OK

// Propiedades opcionales
const perfilSchema = z.object({
  nombre: z.string(),
  descripcion: z.string().optional(),  // Puede no existir
  website: z.string().url().nullable()  // Puede ser null
});

// Campos con valores por defecto
const configSchema = z.object({
  tema: z.enum(["light", "dark"]).default("light"),
  idioma: z.string().default("es")
});
```

### 5. Arrays

```typescript
// Array de strings
const etiquetas = z.array(z.string());
etiquetas.parse(["react", "typescript"]);  // ✅ OK

// Array de objetos
const productosSchema = z.array(z.object({
  id: z.number(),
  nombre: z.string(),
  precio: z.number().positive()
}));

productosSchema.parse([
  { id: 1, nombre: "Laptop", precio: 1000 },
  { id: 2, nombre: "Mouse", precio: 50 }
]);  // ✅ OK

// Con longitud
const palabrasSchema = z.array(z.string()).min(1).max(5);
```

### 6. Con React Hook Form

**La combinación PERFECTA:**

```typescript
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";

// 1. Define esquema Zod
const registroSchema = z.object({
  nombre: z.string().min(3, "Mínimo 3 caracteres"),
  email: z.string().email("Email inválido"),
  edad: z.number().min(18, "Debes ser mayor de 18"),
  contrasena: z.string().min(8, "Mínimo 8 caracteres"),
  confirmar: z.string()
}).refine((data) => data.contrasena === data.confirmar, {
  message: "Las contraseñas no coinciden",
  path: ["confirmar"]
});

// 2. Extraer tipo
type DatosRegistro = z.infer<typeof registroSchema>;

// 3. Usar en formulario
function Registro() {
  const { register, handleSubmit, formState: { errors } } = useForm<DatosRegistro>({
    resolver: zodResolver(registroSchema)  // Zod maneja validación
  });

  const onSubmit = (data: DatosRegistro) => {
    // Aquí data está garantizado válido
    console.log("Datos seguros:", data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input
        {...register("nombre")}
        placeholder="Tu nombre"
      />
      {errors.nombre && <p>{errors.nombre.message}</p>}

      <input
        {...register("email")}
        type="email"
        placeholder="tu@email.com"
      />
      {errors.email && <p>{errors.email.message}</p>}

      <input
        {...register("edad", { valueAsNumber: true })}
        type="number"
      />
      {errors.edad && <p>{errors.edad.message}</p>}

      <input
        {...register("contrasena")}
        type="password"
      />
      {errors.contrasena && <p>{errors.contrasena.message}</p>}

      <input
        {...register("confirmar")}
        type="password"
      />
      {errors.confirmar && <p>{errors.confirmar.message}</p>}

      <button type="submit">Registrarse</button>
    </form>
  );
}
```

### 7. Validación de APIs

```typescript
// Esquema para respuesta de API
const usuariosAPISchema = z.array(z.object({
  id: z.number(),
  nombre: z.string(),
  email: z.string().email(),
  activo: z.boolean()
}));

async function cargarUsuarios() {
  const response = await fetch("/api/usuarios");
  const datos = await response.json();

  try {
    // Validar estructura
    const usuarios = usuariosAPISchema.parse(datos);
    console.log("Usuarios válidos:", usuarios);
    return usuarios;
  } catch (error) {
    console.error("Datos inválidos de API:", error.errors);
    // Mostrar error al usuario
    return [];
  }
}

// Manejo de errores detallado
try {
  const resultado = usuariosAPISchema.parse(datos);
} catch (error) {
  if (error instanceof z.ZodError) {
    error.errors.forEach(err => {
      console.log(`${err.path.join(".")}: ${err.message}`);
      // Output: "0.email: Invalid email"
    });
  }
}
```

### 8. Transformaciones Complejas

```typescript
// Transformar datos mientras validas
const usuarioTransformSchema = z.object({
  nombre: z.string().transform(s => s.trim().toUpperCase()),
  email: z.string().email().transform(s => s.toLowerCase()),
  fecha_nacimiento: z.string().transform(s => new Date(s))
});

const usuario = usuarioTransformSchema.parse({
  nombre: "  juan  ",
  email: "JUAN@EJEMPLO.COM",
  fecha_nacimiento: "1990-01-15"
});

console.log(usuario);
// {
//   nombre: "JUAN",
//   email: "juan@ejemplo.com",
//   fecha_nacimiento: Date object
// }
```

---

## Tabla de Métodos

| Método | Para | Ejemplo |
|--------|------|---------|
| `.parse()` | Validar (ERROR si falla) | `schema.parse(data)` |
| `.safeParse()` | Validar (devuelve success/error) | `result.success ? ... : ...` |
| `.parseAsync()` | Async validation | Validar contra base de datos |
| `.refine()` | Validación personalizada | Contraseñas que coincidan |
| `.transform()` | Transformar datos | Convertir a mayúsculas |
| `.default()` | Valor por defecto | Si no existe |
| `.optional()` | Campo opcional | Puede no existir |
| `.array()` | Array de tipo | `z.array(z.string())` |

---

## Comparación

| Aspecto | TypeScript | Zod |
|--------|-----------|-----|
| **Cuándo valida** | Escritura | Runtime |
| **Detecta errores API** | ❌ No | ✅ Sí |
| **Transformar datos** | ❌ No | ✅ Sí |
| **Con React Hook Form** | ❌ No | ✅ Sí |
| **Mensajes de error** | Genéricos | Personalizables |

---

## Resumen

**Zod es:**
- ✅ Validación en runtime
- ✅ TypeScript-first
- ✅ Integración perfecta con React Hook Form
- ✅ Mensajes de error claros
- ✅ Transformación de datos
- ✅ Para APIs, formularios, configuraciones

🚀 **Datos 100% seguros con Zod**