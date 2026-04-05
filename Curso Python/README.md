# 📚 Curso Completo: Fundamentos de Python

> **Una guía pedagógica completa y exhaustiva para dominar Python desde cero hasta nivel intermedio/avanzado** 
> 
> *Con explicaciones detalladas, ejemplos prácticos, mejores prácticas y ejercicios integrados*

---

## 🎯 ¿Qué es Python?

**Python** es un lenguaje de programación versátil y accesible creado por Guido van Rossum en 1991. Fue diseñado para ser:

- **Legible**: Sintaxis clara y cercana al lenguaje natural
- **Versátil**: Sirve para casi cualquier cosa (web, ciencia, automatización, IA)
- **Productivo**: Escribes menos código, más rápidamente
- **Dinámico**: Tipado en tiempo de ejecución, flexible
- **Poderoso**: La librería estándar es inmensa y hay millones de paquetes

### ¿Por qué aprender Python?

| Aspecto | Beneficio |
|--------|----------|
| **Aplicabilidad** | Web, Data Science, ML, Automatización, Scripting, Games, IoT |
| **Comunidad** | Millones de programadores, miles de librerías (numpy, pandas, django, flask) |
| **Curva aprensión** | Syntax simple, ideal para principiantes pero poderoso para expertos |
| **Salario** | Muy demandado, especialmente en Data Science e IA |
| **Industria** | Usado en: Google, Netflix, Spotify, Instagram, NASA, Tesla |
| **Velocidad desarrollo** | Prototipas rápido, iteración ágil |

### Características Únicas de Python

✅ **Interpretado**: No necesitas compilar, ejecutas directamente  
✅ **Tipado dinámico**: Las variables descubren su tipo automáticamente  
✅ **Multiparadigma**: OOP, funcional, imperativo, todo junto  
✅ **Indentación obligatoria**: Fuerza código legible  
✅ **Duck typing**: "Si camina como pato y grazna como pato..."  
✅ **Batteries included**: Librería estándar gigante  
✅ **Ecosistema masivo**: pip con millones de paquetes  
✅ **Comunidad activa**: Stack Overflow, Github, foros siempre ayudando

---

## 📖 Índice Completo del Curso

### **PARTE I: Fundamentos Básicos**

#### 1️⃣ **Instalación y Configuración** (`install.md`)
   - Descarga e instalación de Python
   - Configuración de entornos virtuales (venv)
   - Instalación de paquetes con pip
   - IDEs recomendados (VS Code, PyCharm)
   - Verificación de instalación

#### 2️⃣ **Estructura Básica de Python** (`estructura_basica_python.md`)
   - Tu primer programa `print("Hola")`
   - Estructura de archivos `.py`
   - Módulos y paquetes
   - Importaciones (from/import)
   - Comentarios y documentación (docstrings)
   - `if __name__ == "__main__":`
   - Argumentos de línea de comandos (sys.argv)

#### 3️⃣ **Variables y Asignación** (`variables.md`)
   - Mutable vs Immutable
   - Declaración: simple vs múltiple
   - Naming conventions (snake_case)
   - Scope: local, global, nonlocal
   - Unpacking y swap de variables
   - Constantes (convención con MAYÚSCULAS)
   - F-strings y formateo

#### 4️⃣ **Tipos de Datos Básicos** (`tipos_datos.md`)
   - Números: `int`, `float`, `complex`
   - Booleanos: `bool` (True/False)
   - Cadenas: `str` (inmutables)
   - Conversión de tipos (casting)
   - type() e isinstance()
   - Operaciones específicas de cada tipo

#### 5️⃣ **Operadores** (`operadores.md`)
   - Aritméticos: +, -, *, /, //, %, **
   - Comparación: ==, !=, <, >, <=, >=
   - Lógicos: and, or, not
   - Asignación: =, +=, -=, *=, /=, //=, %=, **=
   - Identidad: is, is not
   - Membresía: in, not in
   - Precedencia de operadores

#### 6️⃣ **Control de Flujo** (`control_flujo.md`)
   - Sentencias if/elif/else
   - Operador ternario: `x if condición else y`
   - Bloques anidados
   - Truthy y Falsy
   - Match/case (Python 3.10+)

#### 7️⃣ **Bucles e Iteraciones** (`bucles_iteraciones.md`)
   - for: iteración sobre secuencias
   - while: iteración condicional
   - break, continue, pass
   - range() y enumerate()
   - zip() y reversed()
   - else en bucles
   - Anidación de bucles

#### 8️⃣ **Cadenas de Texto (Strings)** (`strings.md`)
   - Creación: comillas simples/dobles/triples
   - Acceso e indexing
   - Slicing (rebanadas)
   - Métodos: split, join, strip, replace, find, etc.
   - F-strings (formatted string literals)
   - Raw strings r""
   - Escapado de caracteres
   - Unicode y encoding

---

### **PARTE II: Estructuras de Datos**

#### 9️⃣ **Listas, Tuplas y Conjuntos** (`listas_tuplas_conjuntos.md`)
   - **Listas** (mutables): creación, acceso, modificación
   - Métodos de listas: append, extend, insert, remove, pop, sort, reverse
   - Slicing de listas
   - **Tuplas** (inmutables): creación, unpacking
   - Named tuples
   - **Conjuntos (sets)**: operaciones de teoría de conjuntos
   - Diferencias y casos de uso

#### 🔟 **Diccionarios y Mapeos** (`diccionarios.md`)
   - Creación: literal {} y dict()
   - Acceso: claves, .get(), .keys(), .values(), .items()
   - Modificación: actualización, pop, del
   - Métodos útiles: setdefault, update, clear
   - Diccionarios anidados
   - Iteración eficiente
   - defaultdict y Counter

#### 1️⃣1️⃣ **Comprehensiones** (`comprehensiones.md`) ⭐ TEMA AVANZADO
   - List comprehensions: [x*2 for x in range(10)]
   - Diccionario comprehensions: {k: v for k, v in items}
   - Set comprehensions
   - Generator expressions (generadores)
   - Nested comprehensions
   - Condiciones en comprehensions

---

### **PARTE III: Funciones y Programación Funcional**

#### 1️⃣2️⃣ **Funciones en Python** (`funciones.md`)
   - Definición: def nombre(parámetros):
   - Parámetros: por posición, por nombre, por defecto
   - *args y **kwargs (argumentos variables)
   - Retorno: simple, múltiple, tupla
   - Docstrings y type hints
   - Scope (LEGB)
   - Funciones anónimas (lambda)
   - Funciones como objetos (first-class functions)

#### 1️⃣3️⃣ **Decoradores** (`decoradores.md`) ⭐ TEMA INTERMEDIO
   - ¿Qué son los decoradores?
   - Funciones que retornan funciones
   - El símbolo @ y su significado
   - Decoradores simples
   - Decoradores con parámetros
   - Preservando metadata con @wraps
   - Decoradores de clase
   - Estacked decorators (pilas)

#### 1️⃣4️⃣ **Programación Orientada a Objetos** (`poo.md`)
   - Clases y objetos: definición e instanciación
   - Atributos: de instancia y de clase
   - Métodos: instancia, clase (@classmethod), estática (@staticmethod)
   - `__init__`, `__str__`, `__repr__`
   - Herencia simple y múltiple
   - Polimorfismo
   - Métodos mágicos (__add__, __len__, etc.)
   - Propiedades (@property)
   - Encapsulación

#### 1️⃣5️⃣ **Manejo de Excepciones** (`manejo_errores.md`)
   - Try/except/finally
   - Múltiples excepciones
   - Excepciones personalizadas
   - Contexto de excepción (e)
   - As (alias de excepciones)
   - Raising exceptions (lanzar)
   - Best practices en error handling

---

### **PARTE IV: Entrada/Salida y Datos**

#### 1️⃣6️⃣ **Entrada y Salida de Archivos** (`archivos_io.md`)
   - Lectura: read(), readline(), readlines()
   - Escritura: write(), writelines()
   - El contexto manager: with open()
   - Modos: r, w, a, r+, w+, a+
   - Rutas: pathlib vs os.path
   - Encodings: UTF-8, ASCII, latin-1
   - Lectura línea por línea (memory efficient)

#### 1️⃣7️⃣ **JSON y Serialización** (`json_serializacion.md`)
   - JSON: formato universal
   - json.dumps() y json.loads()
   - json.dump() y json.load()
   - Objetos personalizados con JSONEncoder/Decoder
   - Pickle (serialización Python)
   - CSV: csv module
   - YAML introducción

#### 1️⃣8️⃣ **Módulos y Paquetes** (`modulos_paquetes.md`)
   - Diferencia: módulos vs paquetes
   - Creación de módulos reutilizables
   - Import: absoluto vs relativo
   - Namespace packages
   - __init__.py
   - sys.path
   - Librerías estándar importantes

#### 1️⃣9️⃣ **Bases de Datos** (`bases_datos.md`)
   - SQLite (en la stdlib)
   - Conexiones y cursores
   - SQL básico con Python
   - ORM: SQLAlchemy introducción
   - Consultas seguras (prevención de SQL injection)
   - Context managers para conexiones
   - Migraciones básicas

---

### **PARTE V: Tests y Buenas Prácticas**

#### 2️⃣0️⃣ **Testing** (`testing.md`) ⭐ TEMA IMPORTANTE
   - unittest módulo
   - pytest framework
   - Test discovery y fixtures
   - Assertions
   - Mocking y patches
   - Cobertura de código
   - TDD (Test-Driven Development)

#### 2️⃣1️⃣ **Buenas Prácticas** (`buenas_practicas.md`)
   - PEP 8: Estilo de código
   - Naming: variables, funciones, clases
   - DRY (Don't Repeat Yourself)
   - SOLID: Single Responsibility, etc.
   - Code reviews
   - Documentación
   - Versionamiento semántico

---

### **PARTE VI: Temas Avanzados Únicos de Python**

#### 2️⃣2️⃣ **Generators y Corrutinas** (`generators_corrutinas.md`)
   - yield vs return
   - Generator functions
   - Expresiones generadoras
   - send() y throw()
   - Async/await (Python 3.5+)
   - asyncio basics

#### 2️⃣3️⃣ **Context Managers** (`context_managers.md`)
   - with statement
   - __enter__ y __exit__
   - contextlib.contextmanager
   - ExitStack
   - Casos de uso prácticos

#### 2️⃣4️⃣ **Librerías Estándar Útiles** (`stdlib_utiles.md`)
   - collections: defaultdict, Counter, namedtuple, deque
   - itertools: combinations, permutations, cycle
   - functools: reduce, partial, lru_cache
   - operator: itemgetter, attrgetter
   - datetime: date, time, datetime, timedelta
   - pathlib: Paths modernos
   - os: Sistema operativo
   - sys: Sistema Python
   - re: Expresiones regulares
   - math: Constantes y funciones matemáticas

#### 2️⃣5️⃣ **Expresiones Regulares** (`regex.md`)
   - Patrones básicos
   - re.match, re.search, re.findall, re.sub
   - Groupos de captura
   - Lookahead y lookbehind
   - Compilación de patrones
   - Raw strings r""

#### 2️⃣6️⃣ **Performance y Optimización** (`performance.md`)
   - Timing: timeit
   - Profiling: cProfile, line_profiler
   - Memory profiling
   - Optimizaciones comunes
   - Numba y Cython para cálculos intensivos

#### 2️⃣7️⃣ **Web con Python** (`web_basico.md`)
   - HTTP: requests library
   - Parsing HTML: BeautifulSoup
   - APIs REST: json
   - Flask: microframework web
   - FastAPI: alternativa moderna

#### 2️⃣8️⃣ **Data Science Intro** (`data_science.md`)
   - NumPy: arrays multidimensionales
   - Pandas: DataFrames
   - Matplotlib: visualización
   - Scikit-learn: ML básico

---

## 📊 Estructura del Curso

```
Curso Python/
├── README.md (este archivo)
├── 
├─ PARTE I: FUNDAMENTOS
│  ├── install.md
│  ├── estructura_basica_python.md
│  ├── variables.md
│  ├── tipos_datos.md
│  ├── operadores.md
│  ├── control_flujo.md
│  └── bucles_iteraciones.md
├─ PARTE II: ESTRUCTURAS DE DATOS
│  ├── strings.md
│  ├── listas_tuplas_conjuntos.md
│  ├── diccionarios.md
│  └── comprehensiones.md
├─ PARTE III: FUNCIONES Y OOP
│  ├── funciones.md
│  ├── decoradores.md
│  ├── poo.md
│  └── manejo_errores.md
├─ PARTE IV: IO Y DATOS
│  ├── archivos_io.md
│  ├── json_serializacion.md
│  ├── modulos_paquetes.md
│  └── bases_datos.md
├─ PARTE V: TESTS Y BUENAS PRÁCTICAS
│  ├── testing.md
│  └── buenas_practicas.md
├─ PARTE VI: AVANZADO
│  ├── generators_corrutinas.md
│  ├── context_managers.md
│  ├── stdlib_utiles.md
│  ├── regex.md
│  ├── performance.md
│  ├── web_basico.md
│  └── data_science.md
└── [Ejercicios y recursos]
```

---

## 🚀 Cómo Usar Este Curso

### Para Beginners Absolutos
1. Lee de forma secuencial: Instalación → Estructura → Variables → ...
2. Ejecuta TODOS los ejemplos de código
3. Intenta modificarlos: ¿Qué pasa si cambio esto?
4. No avances hasta que entiendas completamente

### Para Users con experiencia en otros lenguajes
1. Lee "Estructura Básica" y "Tipos de Datos" rápidamente
2. Salta a "Características Únicas de Python": Comprehensiones, Decoradores
3. Profundiza en OOP y librerías útiles
4. Experimenta con ejemplos de Data Science

### Aprendizaje Práctico
- ✏️ Cada sección tiene ejemplos completos y ejecutables
- 🎯 Intenta resolver los ejercicios propuestos
- 🔄 Repite: No entiende algo → Copia el código → Modifícalo
- 📖 Usa este curso como referencia mientras te programas

---

## ⚡ Requisitos Previos

- **Computadora** con Windows, macOS o Linux
- **0 experiencia previa** necesaria (si aprendiste con otro lenguaje, mejor)
- **Curiosidad**: Lo más importante
- **Paciencia**: La programación requiere práctica

---

## 🎓 ¿Qué aprenderás?

Al completar este curso, serás capaz de:

✅ Ejecutar programas Python desde la terminal  
✅ Manipular datos con tipos y estructuras  
✅ Crear funciones reutilizables y bien organizadas  
✅ Usar OOP para código escalable  
✅ Leer/escribir archivos y JSON  
✅ Manejar errores elegantemente  
✅ Escribir código legible y mantenible  
✅ Entender la filosofía Pythónica  
✅ Trabajar con librerías populares  
✅ Resolver problemas reales  

---

## 💡 Filosofía de Python

Python sigue una filosofía simple pero poderosa: **"La legibilidad es importante"**. Esto se resume en el "Zen de Python":

```
The Zen of Python, by Tim Peters

Beautiful is better than ugly.
Explicit is better than implicit.
Simple is better than complex.
Complex is better than complicated.
Readability counts.
```

En todo este curso, seguimos estos principios. El código que escribas será hermoso, claro y efectivo.

---

## 📞 Tips mientras aprendes

1. **Experimenta en REPL**: `python3` en terminal para código interactivo
2. **Lee los errores**: Los errores de Python en español son MUY específicos
3. **Usa `help()`**: `help(str.split)` te muestra cómo usar algo
4. **Google es tu amigo**: "Python how to X" siempre tiene respuestas
5. **Comunidad**: r/learnprogramming, Python.org, StackOverflow

---

**¡Bienvenido a Python! Que disfrutes el viaje.** 🚀

---

*Última actualización: 2026*  
*Versión: 1.0 - Curso Completo*
