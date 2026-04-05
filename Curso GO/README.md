# 📚 Curso Completo: Fundamentos de Go

> **Una guía pedagógica completa para dominar Go desde cero hasta nivel intermedio/avanzado**

---

## 🎯 ¿Qué es Go?

**Go** (también conocido como **Golang**) es un lenguaje de programación moderno creado por Google en 2009. Fue diseñado para:

- **Ser simple**: Sintaxis clara y directa, sin complejidad innecesaria
- **Ser rápido**: Compilación rápida, ejecución eficiente
- **Ser concurrente**: Goroutines y channels para paralelismo fácil
- **Ser escalable**: Ideal para sistemas grandes y distribuidos
- **Ser productivo**: Menos líneas de código, más claridad

### ¿Por qué aprender Go?

| Aspecto | Beneficio |
|--------|----------|
| **Performance** | Compilado a binario nativo, tan rápido como C |
| **Concurrencia** | Goroutines (ligeras), no threads pesados |
| **Simplicidad** | Menos ceremonias que Java, más seguridad que Python |
| **Escalabilidad** | Diseñado para sistemas distribuidos |
| **Industria** | Usado en: Google, Netflix, Docker, Kubernetes, CloudFlare |
| **Salario** | Uno de los lenguajes mejor pagados |

### Características Únicas de Go

✅ **Compilación a binario**: Un solo archivo ejecutable  
✅ **No tiene clases**: Usa structs + métodos (OOP simplificado)  
✅ **Manejo explícito de errores**: `if err != nil` cada paso  
✅ **Concurrencia nativa**: Goroutines son baratas (~2KB vs ~2MB threads)  
✅ **Interfaces implícitas**: Duck typing type-safe  
✅ **Un estilo de código**: gofmt obliga convenciones  
✅ **Cero dependencias**: Stdlib muy completa  

---

## 📖 Índice Completo del Curso

### **PARTE I: Fundamentos Básicos**

#### 1️⃣ **Estructura Básica de Go** (`02-estructura_basica_go.md`)
   - Package declaration y función main()
   - Importaciones (orden y best practices)
   - Comentarios y documentación (Godoc)
   - Orden típico de un archivo .go
   - Paquetes locales
   - Argumentos de línea de comandos
   - init() y ciclo de ejecución

#### 2️⃣ **Variables, Constantes y Tipos** (`03-variables.md`)
   - Declaración: `var`, `:=`, constantes
   - Tipos de datos básicos: int, float, string, bool
   - Zero values (valores por defecto)
   - Conversión de tipos
   - Operadores de asignación
   - Convenciones de nombres (camelCase, PascalCase)
   - Ejemplo integrado: Calculadora de presupuesto
   - Anti-patrones y mejores prácticas

#### 3️⃣ **Tipos de Datos** (`04-tipos_datos.md`)
   - Tipos básicos: bool, int, float, string
   - Tipos especiales: byte, rune, time.Time
   - Tipos compuestos: array, slice, map, struct
   - Pointers, functions, interfaces
   - Ejercicio integrado: Sistema de biblioteca

#### 4️⃣ **Operadores** (`05-operadores.md`)
   - Aritméticos: +, -, *, /, %
   - Lógicos: &&, ||, !
   - Comparación: ==, !=, <, >
   - Asignación: =, +=, -=, etc.

#### 5️⃣ **Estructura Secuencial** (`06-estructura_secuencial.md`)
   - Variables y asignaciones
   - Operaciones secuenciales
   - Actualizaciones de variables
   - Entrada y salida (lectura de usuarios)
   - Flujo de datos (pipelines)
   - Ejemplo integrado: Carrito de compras
   - Mejores prácticas

#### 6️⃣ **Strings** (`07-string.md`)
   - Creación y manipulación
   - Funciones de strings (split, join, contains, etc.)
   - Conversiones y codificación UTF-8
   - Ejemplos prácticos

#### 7️⃣ **Arrays** (`08-array.md`)
   - Arrays vs slice
   - Iteración y acceso
   - Operaciones comunes
   - Casos de uso

#### 8️⃣ **Slices** (`09-slices.md`)
   - Dinámicos y flexibles
   - Append, slicing, capacity
   - Diferencias con arrays
   - Patrones comunes

---

### **PARTE II: Control de Flujo y Bucles**

#### 9️⃣ **Control de Flujo** (`10-control_flujo.md`)
   - `if/else/else if`: Decisiones simples
   - `switch`: Múltiples opciones
   - `goto`: Evitar (generalmente)
   - 3 ejemplos integrados: Calculadora, Validador, Procesador

#### 🔟 **Bucles e Iteraciones** (`11-bucles_iteraciones.md`)
   - `for` clásico: El único loop de Go
   - `for range`: Iterar sobre colecciones
   - `break` y `continue`
   - Patrones: buscar, acumular, filtrar, mapear
   - 5 ejemplos integrados complejos

---

### **PARTE III: Funciones y Métodos**

#### 1️⃣1️⃣ **Funciones** (`12-funciones.md`)
   - Funciones básicas: parámetros y retornos
   - Múltiples retornos y error handling
   - Funciones anónimas y closures
   - Funciones como argumentos (callbacks, map/filter)
   - Defer, panic, recover
   - Recursión y memoización
   - Orden superior y composición

#### 1️⃣2️⃣ **Programación Orientada a Objetos** (`13-poo.md`)
   - Structs como objetos (sin clases)
   - Métodos: receivers por valor/referencia
   - Interfaces: contratos y polimorfismo
   - Composición vs herencia
   - Encapsulación (público/privado)
   - Type assertion y type switch
   - Ejemplo completo: Sistema bancario

---

### **PARTE IV: Estándar Library (Standard Packages)**

#### 1️⃣3️⃣ **Time y Date** (`14-time.md`, `15-date.md`)
   - `time.Now()`: Hora actual
   - Parsing y formatting
   - Durations
   - Operaciones con fechas
   - Ejemplos de uso

#### 1️⃣4️⃣ **Math y Random** (`16-math.md`, `17-math_rand.md`)
   - Funciones matemáticas: Sqrt, Pow, Sin, Cos, Log
   - Números aleatorios
   - Seeding for reproducibility

---

### **PARTE V: Estructuras de Datos Avanzadas**

#### 1️⃣5️⃣ **Estructuras de Datos Complejas** (`18-data_structures.md`)
   - Arrays y slices: casos de uso
   - Maps: tablas hash eficientes
   - Structs: agrupación de datos
   - Interfaces: abstracciones
   - Pointers: referencias y memoria
   - Linked Lists: nodo a nodo
   - Stacks: LIFO
   - Queues: FIFO
   - Sets: elementos únicos
   - Binary Trees: búsqueda eficiente
   - Graphs: estructuras complejas
   - **19 secciones con ejemplos completos**

---

### **PARTE VI: I/O y Persistencia**

#### 1️⃣6️⃣ **Archivos y Persistencia** (`19-archivos_persistencia_serializacion.md`)
   - Leer/escribir archivos
   - Directorios
   - Serialización: JSON, CSV, XML, Binary
   - Base de datos simple con JSON
   - Persistencia con gob (binary)
   - Copiar archivos
   - Manejo de paths multiplataforma
   - Tabla de formatos
   - Mejores prácticas

---

### **PARTE VII: Algoritmos y Complejidad**

#### 1️⃣7️⃣ **Algoritmos y Big O** (`20-algoritmos_complejidad.md`)
   - Notación Big O: O(1), O(n), O(n²), O(2ⁿ)
   - Búsqueda lineal vs binaria
   - Ordenamiento: Bubble, Quick, Merge
   - Estructuras de datos y complejidad
   - Memoización y programación dinámica
   - Divide y conquista
   - Problemas clásicos: Fibonacci, torres de Hanoi
   - Performance real vs teoría

---

### **PARTE VIII: Manejo de Errores y Debugging**

#### 1️⃣8️⃣ **Errores y Debugging** (`21-manejo_errores_depuracion.md`)
   - Modelo de errores de Go (valores, no excepciones)
   - `errors.New()` y `fmt.Errorf()`
   - Errores personalizados
   - Wrapping errors con contexto
   - `defer`, `panic`, `recover`
   - Print statements y logging
   - Structured logging
   - Stack traces
   - Testing y benchmarks
   - Delve debugger y profiling

---

### **PARTE IX: Bases de Datos**

#### 1️⃣9️⃣ **Bases de Datos Completo** (`22-bases_datos.md`)
   - Conceptos: SQL vs NoSQL, estructura, relaciones
   - **PostgreSQL**: Driver pq, CRUD, conexiones, queries
   - **MySQL**: Driver mysql, diferencias con PostgreSQL
   - **Oracle**: Driver godror, sintaxis específica
   - **SQL Server**: Driver mssqldb, integración Microsoft
   - **SQLite**: Local, sin servidor, perfecto para testing
   - **MongoDB**: Documentos JSON, colecciones, BSON
   - **Redis**: Clave-valor, cache, sesiones, colas
   - **Elasticsearch**: Full-text search, indexación
   - Patrón general (funciona para toda BD)
   - Mejores prácticas y anti-patrones
   - Tabla comparativa: cuál usar según caso

---

### **PARTE X: Web y APIs**

#### 2️⃣0️⃣ **HTTP y Servidores Web** (`23-http_servidores_web.md`)
   - Servidor HTTP mínimo
   - Objeto Request (métodos, parámetros, headers, body)
   - Objeto ResponseWriter (JSON, status codes, headers)
   - Enrutamiento (routing simples y avanzado)
   - Ejemplo práctico: API REST completa
   - Middleware (logging, CORS)
   - Manejo de errores
   - Anti-patrones

#### 2️⃣1️⃣ **Paquetes, Módulos y Organización** (`24-paquetes_modulos_organizacion.md`)
   - Qué son los paquetes (packages)
   - Módulos y go.mod
   - Importaciones (locales y terceros)
   - Estructura típica de proyecto (cmd/, pkg/, internal/)
   - Visibilidad (exportado vs privado)
   - Circular imports y soluciones
   - Ejemplo práctico: API estructurada
   - Best practices de organización
   - Anti-patrones

---

### **PARTE XI: Patrones Avanzados y Control de Ejecución**

#### 2️⃣2️⃣ **Context, Timeouts y Cancelación** (`25-context_timeouts_cancelacion.md`)
   - Qué es Context y por qué existe
   - Deadline y Done channels
   - context.Background() y context.TODO()
   - Timeouts con context.WithTimeout()
   - Cancelación manual con context.WithCancel()
   - Pasar valores con context.WithValue()
   - Patrón profesional HTTP con context
   - Timeouts en BD, HTTP client, goroutines
   - Anti-patrones comunes
   - Mejores prácticas en producción
   - Cleanup automático de recursos

#### 2️⃣3️⃣ **ORM y Query Builders** (`26-orm_query_builders.md`)
   - SQL puro vs ORM (seguridad, type-safety)
   - **GORM**: Instalación, modelos, CRUD
   - Queries avanzadas y relaciones
   - Transacciones en GORM
   - Patrón Repository profesional
   - **sqlc**: Type-safe SQL generator
   - Ejemplo completo con sqlc
   - **ent**: Entity framework para Go
   - Cuándo usar: GORM, sqlc, ent, SQL puro
   - Anti-patrones (SQL injection, ignorar errores)
   - Mejores prácticas en producción
   - Jerarquía de recomendación profesional

---

### **PARTE XII: Autenticación y Seguridad**

#### 2️⃣4️⃣ **Autenticación y Seguridad** (`27-autenticacion_seguridad.md`)
   - ¿Por qué existe la seguridad?
   - Hashing de contraseñas con bcrypt
   - Registro y login seguros
   - JWT (JSON Web Tokens)
   - Generación y validación de tokens
   - Middleware de autenticación
   - Validación de entradas (prevenir SQL injection)
   - Rate limiting contra ataques
   - CORS - Control de acceso entre dominios
   - HTTPS/TLS - Encriptación de conexión
   - Encriptación de datos sensibles (AES)
   - Anti-patrones comunes
   - Checklist de seguridad en producción
   - Best practices profesionales

---

### **PARTE XIII: Logs, Configuración y Observabilidad**

#### 2️⃣5️⃣ **Logs, Configuración y Observabilidad** (`28-logs_configuracion_observabilidad.md`)
   - Logging estructurado vs desordenado
   - Slog (standard de Go 1.21+)
   - Logging en JSON para producción
   - Contextualización de logs (requestID)
   - Configuración por variables de entorno
   - Struct de configuración
   - Viper para archivos .env
   - Observabilidad: Logs, Métricas, Trazas
   - Recolectar métricas simples
   - Prometheus para métricas
   - Health checks
   - Trazas distribuidas con OpenTelemetry
   - Anti-patrones en logging
   - Best practices en observabilidad
   - Ejemplo completo integrado

---

### **PARTE XIV: Mejores Prácticas**

#### 2️⃣6️⃣ **Buenas Prácticas** (`29-buenas_practicas.md`)
   - Convenciones de nombres: públicos/privados
   - Organización de código: paquetes y archivos
   - Manejo de errores idiomático
   - Testing: tabla de casos, benchmarks, mocks
   - Documentación y comentarios
   - Performance: pre-aloquing, strings.Builder
   - Seguridad: validación, cerrar recursos
   - Patrones: Constructor, Builder, Dependency Injection
   - Anti-patrones a evitar
   - Herramientas: gofmt, golint, go vet, godoc

---

## 🚀 Cómo Usar Este Repositorio

### Opción 1: Aprendizaje Secuencial

Sigue el orden del índice arriba. Cada documento es auto-contenido pero construye sobre anteriores.

**Camino recomendado**:
```
1. estructura_basica_go.md               (1-2 horas)
2. variables.md                          (1-2 horas)
3. tipos_datos.md                        (2-3 horas)
4. operadores.md                         (1 hora)
5. estructura_secuencial.md              (2 horas)
6. control_flujo.md                      (2 horas)
7. bucles_iteraciones.md                 (2 horas)
8. funciones.md                          (3 horas)
9. poo.md                                (3 horas)
10. paquetes_modulos_organizacion.md     (2 horas)
11. http_servidores_web.md               (3 horas)
12. context_timeouts_cancelacion.md      (2 horas)
13. orm_query_builders.md                (3 horas)
14. autenticacion_seguridad.md           (3 horas)
15. logs_configuracion_observabilidad.md (3 horas) ← NUEVO
16. archivos_persistencia_serializacion.md  (2 horas)
17. bases_datos.md                       (3 horas)
18. algoritmos_complejidad.md            (2 horas)
19. manejo_errores_depuracion.md         (2 horas)
20. buenas_practicas.md                  (3 horas)
19. buenas_practicas.md                  (3 horas)
... (el resto según interés)
```

### Opción 2: Por Tema de Interés

Si quieres aprender algo específico:
- **"Cómo empezar con Go"** → `estructura_basica_go.md` + `variables.md`
- **"Cómo escribir funciones"** → `funciones.md`
- **"Cómo organizar mi proyecto"** → `paquetes_modulos_organizacion.md`
- **"Cómo crear un servidor web"** → `http_servidores_web.md`
- **"Cómo manejar timeouts y cancelación"** → `context_timeouts_cancelacion.md`
- **"Cómo trabajar con bases de datos"** → `bases_datos.md` + `orm_query_builders.md`
- **"Cómo elegir ORM"** → `orm_query_builders.md`
- **"Cómo autenticar usuarios"** → `autenticacion_seguridad.md`
- **"Cómo asegurar mi aplicación"** → `autenticacion_seguridad.md`
- **"Cómo trabajar con tokens JWT"** → `autenticacion_seguridad.md`
- **"Cómo loguear eventos"** → `logs_configuracion_observabilidad.md`
- **"Cómo configurar mi aplicación"** → `logs_configuracion_observabilidad.md`
- **"Cómo monitorear mi aplicación"** → `logs_configuracion_observabilidad.md`
- **"Cómo trabajar con archivos"** → `archivos_persistencia_serializacion.md`
- **"Cómo manejar errores"** → `manejo_errores_depuracion.md`
- **"Cómo optimizar código"** → `algoritmos_complejidad.md`

### Opción 3: Ejecutar Ejemplos

Estructura del repositorio:

```
e:\GO courses\
├── README.md                              (este archivo)
├── tipos_datos.md
├── control_flujo.md
├── bucles_iteraciones.md
├── ... (14 documentos más)
│
├── ejercicio01/                           (Ejemplos ejecutables)
│   ├── ejercicio01.go
│   └── README.md
├── estructura_secuencial/
│   ├── estructura_secuencial.go
│   ├── main.go
│   └── README.md
├── estructura_condicionales/
│   ├── estructura_condicional.go
│   ├── main.go
│   └── README.md
└── estructura_condicional_compuesta/
    ├── estructura_condicional_compuesta.go
    ├── main.go
    └── README.md
```

**Para ejecutar un ejemplo**:
```bash
cd e:\GO courses\ejercicio01
go run ejercicio01.go
```

O con main.go:
```bash
cd e:\GO courses\estructura_secuencial
go run main.go
```

---

## 📊 Estadísticas del Curso

| Métrica | Valor |
|---------|-------|
| **Documentos** | 26 archivos markdown |
| **Líneas de documentación** | ~38,000+ líneas |
| **Ejemplos de código** | 480+ ejemplos completos |
| **Secciones** | 190+ subsecciones |
| **Partes principales** | 14 capítulos |
| **Horas de estudio** | 55-75 horas (ritmo moderado) |
| **Nivel final** | Intermedio/Avanzado/Profesional |

---

## ✨ Características Destacadas

### 📖 Estilo Pedagógico
- Cada documento **explicado como un libro**, no como referencia cruda
- **Introducción y conclusión narrativa** en cada tema
- **Metáforas y analogías** para conceptos complejos
- **Viaje del programador**: "Cómo aprenderás esto"

### 💡 Ejemplos Prácticos
- **300+ ejemplos de código** ejecutables
- Desde básico hasta avanzado
- Casos reales de uso
- Comparaciones: **Bien ✅ vs Mal ❌**

### 🎯 Tablas de Referencia
- Big O notation
- Comparación de tipos de datos
- Estructuras de datos vs complejidad
- Formatos de serialización
- Patrones de diseño

### 🔍 Integración Total
- Conceptos se construyen uno sobre otro
- Cross-referencias entre documentos
- Un único "sistema educativo" coherente

---

## 🎓 Progresión de Aprendizaje

```
Semana 1: Fundamentos Básicos
├─ Estructura básica (.go, package, main)
├─ Variables y constantes
├─ Tipos de datos
├─ Operadores
└─ Estructura secuencial

Semana 2: Control y Funcionalidad
├─ Control de flujo (if/switch)
├─ Bucles e iteraciones
├─ Funciones básicas
└─ Strings, arrays, slices

Semana 3: Programación Avanzada
├─ Funciones avanzadas (closures, orden superior)
├─ POO sin clases (structs, métodos, interfaces)
└─ Paquetes y módulos

Semana 4: Desarrollo Web
├─ HTTP y servidores
├─ APIs REST
└─ Middleware y routing

Semana 5: Datos y Persistencia
├─ Archivos y serialización
├─ Bases de datos (SQL y NoSQL)
└─ Patrones de acceso a datos

Semana 6: Patrones Avanzados
├─ Context, timeouts y cancelación
├─ ORM y Query Builders (GORM, sqlc)
└─ Gestión de ciclo de vida

Semana 7: Optimización y Calidad
├─ Algoritmos y complejidad
├─ Manejo de errores
└─ Buenas prácticas
```

---

## 🛠️ Cómo Ejecutar Ejemplos

### Requisitos
- Go 1.18+ instalado
- PowerShell o terminal
- Editor de texto (VS Code recomendado)

### Paso a Paso

1. **Abre PowerShell** en la carpeta de la lección:
```powershell
cd "e:\GO courses\estructura_secuencial"
```

2. **Ejecuta el código**:
```powershell
go run main.go
```

3. **Modifica y experimenta**: Abre `estructura_secuencial.go` y cambia valores.

4. **Compila si quieres**:
```powershell
go build -o programa.exe
.\programa.exe
```

---

## 📝 Notas Importantes

- **Cada carpeta con código** tiene su propio `README.md` con instrucciones específicas
- **Los documentos son independientes** pero se construyen uno sobre otro
- **Las buenas prácticas** se enfatizan desde el inicio
- **No es un referencial seco**: Es pedagogía, es para aprender pensando

---

## 🎯 Objetivos de Aprendizaje

Al completar este curso, podrás:

✅ **Escribir programas básicos**: Variables, tipos, bucles, condicionales  
✅ **Estructurar código**: Funciones, structs, métodos  
✅ **Manejar datos**: Arrays, slices, maps, estructuras de datos  
✅ **Trabajar con archivos**: Lectura, escritura, serialización  
✅ **Escribir código limpio**: Siguiendo buenas prácticas de Go  
✅ **Debuggear problems**: Manejo de errores, testing, profiling  
✅ **Entender complejidad**: Big O, optimización, algoritmos  
✅ **Diseñar arquitectura**: Interfaces, composición, patrones  

---

## 🚀 Próximos Pasos (Después de Este Curso)

Una vez completes esto, estará listo para:
- **Goroutines y Channels**: Concurrencia avanzada
- **HTTP y Web**: APIs, servidores, REST
- **Bases de datos**: SQL, queries, ORM
- **Testing avanzado**: Coverage, CI/CD
- **Microservicios**: Proyectos en producción

---

## 💬 Filosofía de Este Curso

> "Make it easy to read, not easy to write" — Go Proverbs

Go prioriza **claridad y simplicidad**. Este curso refleja eso:
- Código claro sobre código ingenioso
- Explícito sobre implícito
- Una forma de hacer algo, no diez
- Menos características, mejor comprensión

**Tu meta**: Escribir Go como lo hacen los profesionales de Google.

---

## 📞 Estructura de Carpetas

```
e:\GO courses\
│
├── 📄 Documentación (18 archivos .md)
│   ├── tipos_datos.md                    (Todos los tipos)
│   ├── operadores.md                     (Operaciones)
│   ├── estructura_secuencial.md          (Paso a paso)
│   ├── control_flujo.md                  (if/switch)
│   ├── bucles_iteraciones.md             (for y range)
│   ├── funciones.md                      (Función completa)
│   ├── poo.md                            (Structs y métodos)
│   ├── string.md                         (Manipulación texto)
│   ├── array.md                          (Arrays fijos)
│   ├── slices.md                         (Slices dinámicos)
│   ├── time.md                           (Fecha/hora)
│   ├── math.md                           (Matemáticas)
│   ├── math_rand.md                      (Números aleatorios)
│   ├── date.md                           (Operaciones de fecha)
│   ├── data_structures.md                (Estructuras complejas)
│   ├── archivos_persistencia_serializacion.md
│   ├── algoritmos_complejidad.md         (Big O y optimización)
│   ├── manejo_errores_depuracion.md      (Errores y debug)
│   ├── bases_datos.md                    (PostgreSQL, MySQL, Oracle, SQL Server, MongoDB, Redis, etc.)
│   └── buenas_practicas.md               (Profesionalismo)
│
└── 📁 Ejemplos ejecutables (4 carpetas)
    ├── ejercicio01/                      (Hola Mundo)
    ├── estructura_secuencial/            (Secuencia)
    ├── estructura_condicionales/         (If/else)
    └── estructura_condicional_compuesta/ (Lógica compleja)
```

---

## ¡Listo para empezar!

**Recomendación**: Comienza con `tipos_datos.md`, lee despacio, ejecuta los ejemplos, modifica el código y experimenta.

Go es simple una vez entiendes sus principios. Vamos. 🚀
