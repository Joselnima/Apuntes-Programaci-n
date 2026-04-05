# Módulo 21 - React Hook Form

> **Formularios sin dolor - Control total con mínimo código**

---

## ¿QUÉ es React Hook Form?

**React Hook Form = formularios inteligentes**

Es como la diferencia entre:
- Vanilla React: Mucho boilerplate, mucho estado, lentos
- React Hook Form: Mínimo código, máxima eficiencia

```javascript
// ❌ Vanilla React (tedioso)
const [nombre, setNombre] = useState("");
const [email, setEmail] = useState("");
const [password, setPassword] = useState("");
// Y así para cada campo... 😫

// ✅ React Hook Form (elegante)
const { register, handleSubmit } = useForm();
```

---

## ¿PARA QUÉ?

### Problema 1: Mucho estado

```javascript
// ❌ Cada campo = línea de código
const [nombre, setNombre] = useState("");
const [email, setEmail] = useState("");
const [telefono, setTelefono] = useState("");
const [edad, setEdad] = useState("");
const [ciudad, setCity] = useState("");
const [pais, setPais] = useState("");
const [activo, setActivo] = useState(false);
// 7 campos = 14 líneas 😱
```

### Problema 2: Validaciones duplicadas

```javascript
// ❌ Validar en onChange Y en submit
const handleChange = (e) => {
  if (e.target.value.length < 3) {
    setErrorNombre("Mínimo 3 caracteres");
  }
};

const handleSubmit = (e) => {
  e.preventDefault();
  if (nombre.length < 3) {
    setErrorNombre("Mínimo 3 caracteres");
  }
};
```

### Problema 3: Re-renders innecesarios

```javascript
// ❌ Cada onChange = re-render de todo
<form>
  <input onChange={handleNombreChange} />
  <input onChange={handleEmailChange} />
  <input onChange={handlePasswordChange} />
  {/* Todo se re-renderiza 😫 */}
</form>
```

**React Hook Form lo resuelve TODO:**

```javascript
// ✅ Natural, simple, rápido
const { register, handleSubmit, formState: { errors } } = useForm();

<form onSubmit={handleSubmit(onSubmit)}>
  <input {...register("nombre", { required: true, minLength: 3 })} />
  {errors.nombre && <p>Error</p>}
</form>
```

---

## ¿CÓMO funciona?

### 1. Setup Básico

```bash
npm install react-hook-form
```

**Componente simple:**

```javascript
import { useForm } from "react-hook-form";

function LoginForm() {
  const { register, handleSubmit } = useForm();

  const onSubmit = (data) => {
    console.log(data);  // { email: "...", password: "..." }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input
        {...register("email")}
        placeholder="Email"
      />
      <input
        {...register("password")}
        type="password"
        placeholder="Contraseña"
      />
      <button type="submit">Enviar</button>
    </form>
  );
}
```

**¿Qué hace `{...register("email")}`?**
- Conecta el input al formulario
- Rastrean cambios sin re-renders extras
- Administra estado internamente

### 2. Validaciones

**Validar antes de enviar:**

```javascript
const { register, handleSubmit, formState: { errors } } = useForm();

const onSubmit = (data) => {
  console.log("Datos válidos:", data);
};

return (
  <form onSubmit={handleSubmit(onSubmit)}>
    {/* Email requerido + formato válido */}
    <input
      {...register("email", {
        required: "Email es requerido",
        pattern: {
          value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
          message: "Email inválido"
        }
      })}
      placeholder="Email"
    />
    {errors.email && <p style={{ color: "red" }}>{errors.email.message}</p>}

    {/* Contraseña: mínimo 8 caracteres */}
    <input
      {...register("password", {
        required: "Contraseña requerida",
        minLength: {
          value: 8,
          message: "Mínimo 8 caracteres"
        }
      })}
      type="password"
      placeholder="Contraseña"
    />
    {errors.password && <p style={{ color: "red" }}>{errors.password.message}</p>}

    <button type="submit">Enviar</button>
  </form>
);
```

### 3. Watch - Escuchar cambios específicos

```javascript
import { useForm } from "react-hook-form";

function FormularioContraseña() {
  const { register, watch, formState: { errors } } = useForm();

  const password = watch("password");
  const confirmPassword = watch("confirmPassword");

  const passwordesIguales = password === confirmPassword;

  return (
    <form>
      <input
        {...register("password", { required: "Requerido" })}
        type="password"
        placeholder="Contraseña"
      />

      <input
        {...register("confirmPassword", { required: "Requerido" })}
        type="password"
        placeholder="Confirmar contraseña"
      />

      {!passwordesIguales && (
        <p style={{ color: "red" }}>Las contraseñas no coinciden</p>
      )}

      <button disabled={!passwordesIguales} type="submit">
        Guardar
      </button>
    </form>
  );
}
```

### 4. FormularioComplejo

```javascript
import { useForm } from "react-hook-form";

interface DatosRegistro {
  nombre: string;
  email: string;
  edad: number;
  ciudad: string;
  pais: string;
  terminos: boolean;
}

function RegistroCompleto() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    watch,
    reset
  } = useForm<DatosRegistro>({
    defaultValues: {
      nombre: "",
      email: "",
      edad: 0,
      ciudad: "",
      pais: "",
      terminos: false
    }
  });

  const onSubmit = async (data: DatosRegistro) => {
    console.log("Enviando:", data);
    // Guardar en servidor
    const response = await fetch("/api/usuarios", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data)
    });

    if (response.ok) {
      alert("¡Registrado!");
      reset();  // Limpiar formulario
    }
  };

  const edad = watch("edad");

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      {/* Nombre */}
      <div>
        <label>Nombre</label>
        <input
          {...register("nombre", {
            required: "Nombre requerido",
            minLength: { value: 3, message: "Mínimo 3 caracteres" }
          })}
          placeholder="Tu nombre"
        />
        {errors.nombre && <span>{errors.nombre.message}</span>}
      </div>

      {/* Email */}
      <div>
        <label>Email</label>
        <input
          {...register("email", {
            required: "Email requerido",
            pattern: {
              value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
              message: "Email inválido"
            }
          })}
          type="email"
          placeholder="tu@email.com"
        />
        {errors.email && <span>{errors.email.message}</span>}
      </div>

      {/* Edad */}
      <div>
        <label>Edad</label>
        <input
          {...register("edad", {
            required: "Edad requerida",
            min: { value: 18, message: "Debes ser mayor de 18" }
          })}
          type="number"
        />
        {errors.edad && <span>{errors.edad.message}</span>}
        {edad > 0 && <p>Tienes {edad} años</p>}
      </div>

      {/* Ciudad */}
      <div>
        <label>Ciudad</label>
        <input
          {...register("ciudad", { required: "Ciudad requerida" })}
          placeholder="Tu ciudad"
        />
        {errors.ciudad && <span>{errors.ciudad.message}</span>}
      </div>

      {/* País */}
      <div>
        <label>País</label>
        <select {...register("pais", { required: "País requerido" })}>
          <option value="">Selecciona país</option>
          <option value="mx">México</option>
          <option value="pe">Perú</option>
          <option value="ar">Argentina</option>
          <option value="co">Colombia</option>
        </select>
        {errors.pais && <span>{errors.pais.message}</span>}
      </div>

      {/* Checkbox - Términos */}
      <div>
        <label>
          <input
            {...register("terminos", {
              required: "Debes aceptar los términos"
            })}
            type="checkbox"
          />
          Acepto términos y condiciones
        </label>
        {errors.terminos && <span>{errors.terminos.message}</span>}
      </div>

      <button type="submit">Registrarse</button>
      <button type="button" onClick={() => reset()}>
        Limpiar
      </button>
    </form>
  );
}
```

### 5. Validaciones Personalizadas

```javascript
const { register } = useForm();

// Validar nombre (solo letras)
const validarNombre = (valor: string) => {
  if (!/^[a-zA-Z\s]+$/.test(valor)) {
    return "Solo se permiten letras";
  }
  return true;
};

<input
  {...register("nombre", {
    required: "Requerido",
    validate: validarNombre
  })}
/>

// Validar que no sea duplicado (async)
const validarEmailUnico = async (email: string) => {
  const response = await fetch(`/api/email-existe?email=${email}`);
  const existe = await response.json();
  return !existe || "Este email ya está registrado";
};

<input
  {...register("email", {
    required: "Requerido",
    validate: validarEmailUnico
  })}
/>
```

---

## Tabla Comparativa

| Aspecto | Vanilla React | React Hook Form |
|--------|---------------|-----------------|
| **Estado por campo** | Múltiples useState | Un único useForm |
| **Re-renders** | Todos los campos | Solo el que cambió |
| **Validaciones** | Manual en cada cambio | Automáticas |
| **Líneas de código** | 100+ | 30-40 |
| **Performance** | Lento con muchos campos | Optimizado |
| **Curva aprendizaje** | Fácil | Media |

---

## Resumen

**React Hook Form es:**
- ✅ Minimalista (poco código)
- ✅ Performante (re-renders inteligentes)
- ✅ Validaciones integradas
- ✅ Fácil de usar
- ✅ Integrable con Zod/Yup para validaciones avanzadas

🚀 **Formularios profesionales en React**