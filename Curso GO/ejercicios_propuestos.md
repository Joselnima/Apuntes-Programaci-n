# 01 - Introducción a Go
## Objetivo
Familiarizarse con el entorno, compilación y ejecución de programas en Go.

---

### 1.
Instala Go y verifica la versión desde terminal.

### 2.
Crea un archivo `main.go` que imprima:
`Hola, Go`

### 3.
Compila tu programa con `go build`.

### 4.
Ejecuta tu programa con `go run`.

### 5.
Modifica el programa para imprimir tu nombre.

### 6.
Haz que el programa imprima también tu profesión.

### 7.
Usa `fmt.Println()` para imprimir 3 líneas distintas.

### 8.
Crea un programa que imprima:
- lenguaje favorito
- editor favorito
- sistema operativo

### 9.
Ejecuta `go env` y revisa el valor de `GOROOT`.

### 10.
Ejecuta `go env` y revisa el valor de `GOPATH`.

### 11.
Crea una carpeta de proyecto llamada `go-fundamentos`.

### 12.
Inicializa un módulo con:
`go mod init go-fundamentos`

### 13.
Revisa el contenido del archivo `go.mod`.

### 14.
Ejecuta `go fmt` sobre tu archivo.

### 15.
Agrega comentarios explicando qué hace el programa.

### 16.
Haz un programa que imprima la fecha actual como texto plano.

### 17.
Haz un programa que imprima una frase motivadora.

### 18.
Crea un binario y ejecútalo manualmente.

### 19.
Haz un programa que imprima 10 veces:
`Estoy aprendiendo Go`

### 20.
Explica con tus palabras:
- qué es Go
- para qué sirve
- por qué usarlo

# 02 - Sintaxis Básica de Go

## Objetivo
Aprender variables, constantes, tipos de datos y operadores.

---

## Ejercicios

### 1.
Declara una variable `nombre` de tipo string e imprímela.

### 2.
Declara una variable `edad` de tipo int e imprímela.

### 3.
Declara una constante `PI` con valor `3.1416`.

### 4.
Crea dos variables `a` y `b` y muestra su suma.

### 5.
Muestra la resta de dos números.

### 6.
Muestra la multiplicación de dos números.

### 7.
Muestra la división de dos números.

### 8.
Declara una variable booleana `esActivo`.

### 9.
Usa `:=` para declarar una variable `ciudad`.

### 10.
Declara un `float64` llamado `precio`.

### 11.
Convierte un `int` a `float64`.

### 12.
Convierte un `float64` a `int`.

### 13.
Crea una variable `inicial` de tipo `rune`.

### 14.
Crea una variable `letra` de tipo `byte`.

### 15.
Haz un programa que calcule el área de un rectángulo.

### 16.
Haz un programa que calcule el perímetro de un cuadrado.

### 17.
Evalúa si `10 > 5`.

### 18.
Evalúa si `4 == 4`.

### 19.
Evalúa si `8 != 3`.

### 20.
Haz un programa que imprima:
- nombre
- edad
- altura
- estado civil

### 21. 
Calcular el sueldo mensual de un operario conociendo la cantidad de horas trabajadas y el pago por hora.

# 03 - Control de Flujo

## Objetivo
Usar condicionales y bucles en Go.

---

## Ejercicios

### 1.
Determina si un número es positivo o negativo.

### 2.
Determina si un número es par o impar.

### 3.
Valida si una persona es mayor de edad.

### 4.
Clasifica una nota:
- 0-10 = desaprobado
- 11-20 = aprobado

### 5.
Evalúa si un usuario puede ingresar:
- edad >= 18
- activo = true

### 6.
Usa `switch` para mostrar el día de la semana según un número.

### 7.
Usa `switch` para mostrar el nombre de un mes.

### 8.
Haz un `for` del 1 al 10.

### 9.
Haz un `for` del 10 al 1.

### 10.
Muestra solo números pares del 1 al 20.

### 11.
Muestra solo números impares del 1 al 20.

### 12.
Calcula la suma de los números del 1 al 100.

### 13.
Calcula el factorial de 5.

### 14.
Imprime la tabla de multiplicar del 7.

### 15.
Cuenta cuántas veces aparece un número dentro de un rango.

### 16.
Usa `continue` para saltar múltiplos de 3.

### 17.
Usa `break` para detener un bucle al llegar a 15.

### 18.
Haz un menú repetitivo con `for`.

### 19.
Imprime los primeros 15 números Fibonacci.

### 20.
Haz un programa que cuente del 1 al 50 y muestre:
- `Fizz` si es múltiplo de 3
- `Buzz` si es múltiplo de 5
- `FizzBuzz` si es múltiplo de ambos

### 21.
Ingresar el sueldo de una persona, si supera los 3000 pesos mostrar un mensaje en pantalla
indicando que debe abonar impuestos.

### 22.
Se carga una fecha (día, mes y año) por teclado.
Mostrar un mensaje si corresponde al primer trimestre del año (enero, febrero o marzo)
Cargar por teclado el valor numérico del día, mes y año.
Ejemplo: dia:10 mes:5 año:2017.

### 23.
Escribir un programa que solicite al usuario ingresar dos valores numéricos enteros y luego
determine cuál de los dos valores es mayor

### 24.
Escribir un programa que solicite al usuario ingresar tres notas de un estudiante y luego determine si el estudiante está promocionado (promedio mayor o igual a 7), regular (promedio entre 4 y 6) o reprobado (promedio menor a 4).

# 04 - Funciones en Go

## Objetivo
Aprender a modular código con funciones.

---

## Ejercicios

### 1.
Crea una función `saludar()` que imprima un mensaje.

### 2.
Crea una función `sumar(a, b int) int`.

### 3.
Crea una función `restar(a, b int) int`.

### 4.
Crea una función `multiplicar(a, b int) int`.

### 5.
Crea una función `dividir(a, b float64) float64`.

### 6.
Crea una función que reciba un nombre y lo salude.

### 7.
Crea una función que determine si un número es par.

### 8.
Crea una función que determine si una edad es válida.

### 9.
Crea una función que devuelva el mayor de dos números.

### 10.
Crea una función que devuelva el mayor de tres números.

### 11.
Haz una función que reciba un precio y devuelva el precio con IGV.

### 12.
Haz una función que reciba un porcentaje de descuento.

### 13.
Crea una función con retorno múltiple:
`nombreCompleto() (string, string)`

### 14.
Haz una función que devuelva suma y promedio de dos números.

### 15.
Haz una función variádica que sume varios números.

### 16.
Haz una función anónima que multiplique dos valores.

### 17.
Usa una función como variable.

### 18.
Crea un closure que incremente un contador.

### 19.
Crea una función que valide si un texto está vacío.

### 20.
Haz un mini sistema de operaciones usando funciones separadas.

# 05 - Colecciones en Go

## Objetivo
Trabajar con arrays, slices y maps.

---

## Ejercicios

### 1.
Crea un array de 5 enteros.

### 2.
Imprime todos los elementos de un array.

### 3.
Calcula la suma de un array.

### 4.
Calcula el promedio de un array.

### 5.
Encuentra el número mayor de un array.

### 6.
Encuentra el número menor de un array.

### 7.
Crea un slice de nombres.

### 8.
Agrega un nuevo nombre al slice.

### 9.
Elimina un elemento de un slice.

### 10.
Recorre un slice con `range`.

### 11.
Crea un map de:
`nombre -> edad`

### 12.
Muestra el valor de una clave en un map.

### 13.
Valida si una clave existe en un map.

### 14.
Elimina una clave de un map.

### 15.
Crea un map de productos con precio.

### 16.
Calcula el total de todos los precios del map.

### 17.
Crea una matriz de 3x3.

### 18.
Imprime una matriz fila por fila.

### 19.
Cuenta cuántas veces se repite cada palabra de una lista.

### 20.
Haz un carrito de compras con slice de productos.

# 06 - Strings y Manejo de Texto

## Objetivo
Manipular texto correctamente en Go.

---

## Ejercicios

### 1.
Declara una cadena con tu nombre completo.

### 2.
Convierte una cadena a mayúsculas.

### 3.
Convierte una cadena a minúsculas.

### 4.
Elimina espacios al inicio y final.

### 5.
Divide una cadena usando `Split`.

### 6.
Une una lista de palabras con `Join`.

### 7.
Reemplaza una palabra dentro de un texto.

### 8.
Valida si un texto contiene una palabra.

### 9.
Cuenta cuántos caracteres tiene una cadena.

### 10.
Cuenta cuántas vocales tiene un texto.

### 11.
Invierte una cadena.

### 12.
Convierte un número string a entero.

### 13.
Convierte un entero a string.

### 14.
Convierte un decimal string a `float64`.

### 15.
Valida si un texto empieza con cierto prefijo.

### 16.
Valida si un texto termina con cierto sufijo.

### 17.
Cuenta cuántas veces aparece una palabra.

### 18.
Formatea un texto con `fmt.Sprintf`.

### 19.
Valida si un correo contiene `@`.

### 20.
Limpia una cadena y normalízala para búsqueda.

# 08 - Punteros en Go

## Objetivo
Entender referencias y modificación en memoria.

---

## Ejercicios

### 1.
Declara una variable y obtén su dirección.

### 2.
Imprime el valor de una variable usando puntero.

### 3.
Modifica un valor usando puntero.

### 4.
Crea una función que modifique un entero por referencia.

### 5.
Crea una función que incremente un contador por puntero.

### 6.
Crea una función que actualice un nombre por puntero.

### 7.
Usa punteros en un struct `Usuario`.

### 8.
Crea un método con receiver puntero.

### 9.
Crea un método con receiver por valor y compáralo.

### 10.
Modifica el stock de un producto por puntero.

### 11.
Crea una función que reciba un `*Producto`.

### 12.
Actualiza el precio de un producto por referencia.

### 13.
Haz un método `Activar()` para un usuario.

### 14.
Haz un método `Desactivar()` para un usuario.

### 15.
Explica diferencia entre copiar struct y pasar puntero.

### 16.
Crea una lista de punteros a structs.

### 17.
Recorre una lista de punteros.

### 18.
Haz una función que resetee un contador.

### 19.
Modifica un saldo bancario usando puntero.

### 20.
Haz una pequeña simulación de edición de cliente usando punteros.

# 07 - Structs y Programación Orientada a Objetos

## Objetivo
Organizar datos complejos y usar herencia/composición.

---

## Ejercicios

### 1.
Crea un struct `Persona` con nombre, edad, email.

### 2.
Crea una instancia de `Persona` e imprímela.

### 3.
Accede a los campos de un struct.

### 4.
Crea un struct anidado `Dirección` dentro de `Persona`.

### 5.
Crea un struct `Producto` con nombre, precio, stock.

### 6.
Inicializa un struct usando literales.

### 7.
Haz un método `nombreCompleto()` para `Persona`.

### 8.
Haz un método `aplicarDescuento()` para `Producto`.

### 9.
Haz un método `destock()` para `Producto`.

### 10.
Crea un método receiver con puntero.

### 11.
Usa `reflect` para inspeccionar campos de un struct.

### 12.
Crea un struct `CuentaBancaria` con saldo.

### 13.
Implementa depósito y retiro.

### 14.
Crea un struct `Empleado` que extienda `Persona`.

### 15.
Haz una interfaz `Trabajador` y implementa métodos.

### 16.
Crea múltiples tipos que implementen una interfaz.

### 17.
Haz composición entre structs (no herencia).

### 18.
Crea un struct con campos privados y públicos.

### 19.
Implementa `String()` en un struct.

### 20.
Haz un pequeño sistema de gestión de productos con structs.

# 09 - Concurrencia Básica (Goroutines y Channels)

## Objetivo
Ejecutar código concurrente de forma segura.

---

## Ejercicios

### 1.
Crea una goroutine que imprima un mensaje.

### 2.
Crea dos goroutines que impriman en orden.

### 3.
Crea un channel y envía un valor.

### 4.
Recibe un valor de un channel.

### 5.
Crea un channel con buffer.

### 6.
Haz comunicación entre dos goroutines.

### 7.
Usa `close()` en un channel.

### 8.
Recorre un channel cerrado con `range`.

### 9.
Usa `select` para esperar múltiples channels.

### 10.
Crea un timeout con `time.After()`.

### 11.
Crea un pool de goroutines procesando tareas.

### 12.
Comunica múltiples valores a través de un channel.

### 13.
Usa `sync.WaitGroup` para sincronizar goroutines.

### 14.
Usa `sync.Mutex` para proteger datos compartidos.

### 15.
Haz un productor-consumidor simple.

### 16.
Implementa un fan-out/fan-in pattern.

### 17.
Crea un pipeline de procesamiento.

### 18.
Maneja errores desde goroutines.

### 19.
Cancela goroutines anticipadamente.

### 20.
Haz un simulador de servidor con goroutines manejando clientes.

# 10 - Manejo de Errores en Go

## Objetivo
Aplicar manejo de errores de forma idiomática.

---

## Ejercicios

### 1.
Crea una función que devuelva un `error` si un número es negativo.

### 2.
Crea una función que divida dos números y falle si el divisor es cero.

### 3.
Usa `errors.New()`.

### 4.
Usa `fmt.Errorf()`.

### 5.
Haz una función que valide edad mínima.

### 6.
Haz una función que valide texto vacío.

### 7.
Haz una función que valide email simple.

### 8.
Crea un error personalizado `ErrUsuarioNoEncontrado`.

### 9.
Usa `errors.Is()`.

### 10.
Crea un tipo de error personalizado con struct.

### 11.
Usa `errors.As()`.

### 12.
Lee un archivo inexistente y captura el error.

### 13.
Haz una función que procese una lista y retorne error si está vacía.

### 14.
Usa `defer` para cerrar un recurso.

### 15.
Genera un `panic` intencionalmente.

### 16.
Recupera un `panic` con `recover`.

### 17.
Haz una función que simule un error de negocio.

### 18.
Encadena errores usando wrapping `%w`.

### 19.
Haz una función que valide múltiples reglas.

### 20.
Crea un mini flujo de login con errores controlados.

# 11 - Algoritmos y Complejidad

## Objetivo
Implementar algoritmos comunes y entender complejidad.

---

## Ejercicios

### 1.
Implementa búsqueda lineal en un array.

### 2.
Implementa búsqueda binaria.

### 3.
Implementa bubble sort.

### 4.
Implementa insertion sort.

### 5.
Implementa quick sort.

### 6.
Implementa merge sort.

### 7.
Calcula la complejidad de tus algoritmos.

### 8.
Implementa un árbol binario de búsqueda.

### 9.
Implementa traversal en árbol (pre, in, post order).

### 10.
Implementa búsqueda en profundidad (DFS).

### 11.
Implementa búsqueda en amplitud (BFS).

### 12.
Resuelve el problema de las N reinas.

### 13.
Implementa memoization para Fibonacci.

### 14.
Implementa programación dinámica: cambio de monedas.

### 15.
Resuelve mochila 0/1 con programación dinámica.

### 16.
Implementa longest common subsequence.

### 17.
Calcula distancia de Levenshtein.

### 18.
Implementa selección topológica.

### 19.
Encuentra el camino más corto (Dijkstra).

### 20.
Haz un benchmark comparando tus algoritmos.

# 12 - Archivos y Persistencia

## Objetivo
Leer, escribir y manipular archivos.

---

## Ejercicios

### 1.
Lee un archivo de texto completo.

### 2.
Lee un archivo línea por línea.

### 3.
Escribe contenido en un archivo nuevo.

### 4.
Agrega contenido a un archivo existente.

### 5.
Elimina un archivo.

### 6.
Copia un archivo a otra ubicación.

### 7.
Renombra un archivo.

### 8.
Obtén información de un archivo (tamaño, fecha).

### 9.
Lista archivos de un directorio.

### 10.
Recorre directorios recursivamente.

### 11.
Crea un directorio.

### 12.
Crea una estructura de directorios.

### 13.
Lee un archivo JSON y deserializa.

### 14.
Serializa datos a JSON y guarda en archivo.

### 15.
Lee un archivo CSV.

### 16.
Escribe datos en CSV.

### 17.
Comprime archivos en ZIP.

### 18.
Descomprime archivos ZIP.

### 19.
Copia un directorio completo.

### 20.
Haz un gestor de archivos con operaciones básicas.

# 13 - Bases de Datos

## Objetivo
Conectarse a BD y ejecutar operaciones CRUD.

---

## Ejercicios

### 1.
Conecta a una base de datos SQLite.

### 2.
Crea una tabla usuarios.

### 3.
Inserta un usuario en la BD.

### 4.
Lee todos los usuarios.

### 5.
Lee un usuario por ID.

### 6.
Actualiza un usuario.

### 7.
Elimina un usuario.

### 8.
Usa consultas preparadas para evitar SQL injection.

### 9.
Crea una tabla con relaciones.

### 10.
Implementa un join entre tablas.

### 11.
Usa transacciones para múltiples operaciones.

### 12.
Maneja errores de BD.

### 13.
Realiza una consulta con filtros.

### 14.
Ordena resultados.

### 15.
Usa LIMIT y OFFSET para paginación.

### 16.
Haz agregaciones (SUM, COUNT, AVG).

### 17.
Crea índices en tabla.

### 18.
Usa un ORM como GORM para CRUD básico.

### 19.
Migraciones automáticas con GORM.

### 20.
Haz un mini CRUD de usuarios con BD SQLite.

# 14 - HTTP y Servidores Web

## Objetivo
Crear servidores web y servicios REST.

---

## Ejercicios

### 1.
Crea un servidor HTTP que devuelva "Hola Mundo".

### 2.
Crea un endpoint que reciba parámetros GET.

### 3.
Crea un endpoint POST que reciba JSON.

### 4.
Devuelve JSON desde un endpoint.

### 5.
Implementa routing básico con `http.HandleFunc`.

### 6.
Usa `http.NewServeMux` para rutas.

### 7.
Maneja errores HTTP (404, 500).

### 8.
Implementa middleware simple.

### 9.
Agrega logging a las peticiones.

### 10.
Lee headers de una petición.

### 11.
Devuelve headers personalizados.

### 12.
Implementa autenticación básica.

### 13.
Valida datos de entrada.

### 14.
Limita el tamaño de peticiones.

### 15.
Implementa CORS.

### 16.
Crea un servidor con HTTPS.

### 17.
Implementa un CRUD REST completo.

### 18.
Parsea diferentes tipos de contenido.

### 19.
Usa un framework como Gin (opcional).

### 20.
Haz un mini API de tareas con todas las operaciones CRUD.

# 15 - Paquetes y Módulos

## Objetivo
Organizar código en paquetes y manejar dependencias.

---

## Ejercicios

### 1.
Crea un paquete personalizado.

### 2.
Exporta una función desde un paquete.

### 3.
Usa una función de tu paquete en main.

### 4.
Crea un paquete con múltiples funciones.

### 5.
Crea una constante exportada.

### 6.
Crea un struct exportado.

### 7.
Crea un paquete con submódulos.

### 8.
Usa funciones de múltiples paquetes.

### 9.
Maneja conflicto de nombres entre paquetes.

### 10.
Usa alias de imports.

### 11.
Oculta funciones privadas en un paquete.

### 12.
Inicializa paquetes con `init()`.

### 13.
Crea un módulo Go con `go.mod`.

### 14.
Agrega una dependencia externa.

### 15.
Actualiza una dependencia.

### 16.
Usa vendoring para gestionar dependencias.

### 17.
Crea un paquete con métodos.

### 18.
Usa interfaces en múltiples paquetes.

### 19.
Crea un paquete con tests (básico).

### 20.
Haz un proyecto multi-paquete organizado.

# 16 - Context, Timeouts y Cancelación

## Objetivo
Controlar el ciclo de vida de operaciones.

---

## Ejercicios

### 1.
Crea un contexto con `context.Background()`.

### 2.
Crea un contexto con `context.TODO()`.

### 3.
Crea un contexto con timeout usando `WithTimeout`.

### 4.
Cancela un contexto manualmente con `WithCancel`.

### 5.
Pasa valores en contexto con `WithValue`.

### 6.
Lee valores del contexto.

### 7.
Implementa una función que respete contexto.

### 8.
Crea un timeout en una consulta a BD.

### 9.
Crea un timeout en un cliente HTTP.

### 10.
Maneja `context.DeadlineExceeded`.

### 11.
Maneja `context.Canceled`.

### 12.
Usa función que valida context.Done().

### 13.
Cancela una goroutine dentro de una función.

### 14.
Propaga contexto entre funciones.

### 15.
Crea un deadline absoluto.

### 16.
Implementa cleanup automático con defer.

### 17.
Usa context en servidor HTTP.

### 18.
Crea un contexto anidado.

### 19.
Maneja múltiples timeouts.

### 20.
Haz un simulador de operaciones con timeouts y cancelación.

# 17 - ORM y Query Builders

## Objetivo
Usar herramientas profesionales para acceso a BD.

---

## Ejercicios

### 1.
Instala GORM.

### 2.
Conecta a BD con GORM.

### 3.
Define un modelo básico.

### 4.
Ejecuta auto-migración.

### 5.
Inserta un registro con GORM.

### 6.
Leo todos los registros.

### 7.
Busca por ID.

### 8.
Filtra con WHERE.

### 9.
Actualiza un registro.

### 10.
Elimina un registro.

### 11.
Usa transacciones en GORM.

### 12.
Implementa relaciones One-to-Many.

### 13.
Implementa relaciones Many-to-Many.

### 14.
Agrega indices en modelos.

### 15.
Usa hooks (BeforeSave, AfterCreate, etc).

### 16.
Implementa validaciones en modelo.

### 17.
Usa raw SQL con GORM.

### 18.
Crea un repositorio con GORM.

### 19.
Usa sqlc para queries type-safe.

### 20.
Haz un CRUD profesional con GORM o sqlc.

# 18 - Testing y Benchmarking

## Objetivo
Escribir tests y medir performance.

---

## Ejercicios

### 1.
Escribe un test simple.

### 2.
Corre tests con `go test`.

### 3.
Crea una tabla de casos de prueba.

### 4.
Usa `t.Errorf()` para fallos.

### 5.
Usa `t.Fatal()` para detener.

### 6.
Crea un mock simple.

### 7.
Usa `testing.T` correctamente.

### 8.
Realiza un benchmark.

### 9.
Analiza resultados de benchmark.

### 10.
Optimiza código basado en benchmark.

### 11.
Usa `setup` y `teardown` en tests.

### 12.
Crea subtests.

### 13.
Cubre funciones con tests.

### 14.
Mide cobertura de código.

### 15.
Testa funciones que retornan errores.

### 16.
Testa goroutines.

### 17.
Usa `testify` para assertions.

### 18.
Crea tests de integración.

### 19.
Haz stress testing.

### 20.
Documenta tests con ejemplos.

# 19 - Despliegue y DevOps Básico

## Objetivo
Compilar, empaquetar y desplegar aplicaciones.

---

## Ejercicios

### 1.
Compila un executables para tu SO.

### 2.
Compila cross-platform (Windows, Mac, Linux).

### 3.
Optimiza el tamaño del binario.

### 4.
Crea un Dockerfile para tu app.

### 5.
Construye una imagen Docker.

### 6.
Ejecuta container localmente.

### 7.
Haz multi-stage build en Docker.

### 8.
Crea un docker-compose.yml.

### 9.
Usa variables de entorno en tu app.

### 10.
Lee configuración desde archivo.

### 11.
Haz logging strukturado.

### 12.
Usa flags de línea de comandos.

### 13.
Crea un health check endpoint.

### 14.
Implementa graceful shutdown.

### 15.
Crea un script de deployment.

### 16.
Usa GitHub Actions (CI/CD básico).

### 17.
Automatiza builds y tests.

### 18.
Documenta tu aplicación.

### 19.
Crea un script de desarrollo.

### 20.
Haz deployment a un servidor simple.

# 21 - Autenticación y Seguridad

## Objetivo
Implementar mecanismos seguros de autenticación y proteger aplicaciones.

---

## Ejercicios

### 1.
Haz hash de una contraseña con `bcrypt`.

### 2.
Verifica una contraseña contra su hash.

### 3.
Implementa registro de usuario con contraseña hasheada.

### 4.
Crea un login básico validando contraseña.

### 5.
Genera un JWT (JSON Web Token).

### 6.
Valida un JWT.

### 7.
Extrae claims de un JWT.

### 8.
Implementa refresh tokens.

### 9.
Crea middleware de autenticación JWT.

### 10.
Protege un endpoint con JWT.

### 11.
Implementa autenticación básica HTTP.

### 12.
Valida autenticación básica en middleware.

### 13.
Encripta datos sensibles con `crypto`.

### 14.
Desencripta datos sensibles.

### 15.
Implements OAuth2 flow básico.

### 16.
Valida tokens OAuth2.

### 17.
Implementa CORS correctamente.

### 18.
Usa HTTPS en servidor.

### 19.
Implementa rate limiting por IP.

### 20.
Haz un sistema de autenticación completo (login, JWT, refresh, logout).

# 22 - Proyecto Final Integrador

## Objetivo
Crear una aplicación que integre varios conceptos.

---

## Desafíos

### 1.
Diseña una aplicación pequeña (gestor de tareas, blog, etc).

### 2.
Estructura el proyecto en paquetes organizados.

### 3.
Define los modelos de datos.

### 4.
Implementa la lógica de negocio.

### 5.
Crea endpoints REST.

### 6.
Usa base de datos (SQLite o PostgreSQL).

### 7.
Implementa autenticación con JWT.

### 8.
Valida todas las entradas.

### 9.
Maneja errores correctamente.

### 10.
Haz logging de operaciones.

### 11.
Implementa paginación.

### 12.
Usa context para operaciones largas.

### 13.
Contempla seguridad (SQL injection, XSS, validación).

### 14.
Escribe tests para funciones críticas.

### 15.
Realiza benchmarks en operaciones lentas.

### 16.
Dockeriza tu aplicación.

### 17.
Crea documentación README.

### 18.
Implementa graceful shutdown.

### 19.
Demuestra best practices de Go.

### 20.
Extiende con una característica adicional que aprendiste.