# 🚀 CURSO DE PROGRAMACIÓN (ENFOQUE PROFESIONAL)

## 🧩 MÓDULO 0: Fundamentos (Base sólida)

> *Aquí evitas el típico error de “copiar código sin entender”.*

### ¿Qué es programar?
Programar es el proceso de darle instrucciones a una computadora para que realice una tarea específica o resuelva un problema. Es como escribir una receta detallada, donde cada paso debe ser lógico y ordenado. El objetivo principal es resolver problemas mediante la automatización de procesos. A nivel más profundo, programar trata sobre la **gestión de la complejidad** y la estructuración del pensamiento para que una máquina, que carece de raciocinio o intuición humana, pueda ejecutar tareas de forma predecible y segura.

### ¿Cómo piensa una computadora?
A nivel fundamental, una computadora "piensa" a través de impulsos eléctricos (encendido/apagado), representados por ceros y unos (**código binario**). La computadora opera basándose puramente en la **Lógica Booleana** (Verdadero y Falso). A diferencia de los humanos, no tiene intuición ni sentido común; ejecuta exactamente lo que se le ordena de forma literal, secuencial y sin cuestionar. Por ello, nuestras instrucciones deben ser 100% precisas y sin ambigüedades. Cuando un programa falla (lo que llamamos un *bug* o error), casi siempre se debe a que la lógica humana proporcionó instrucciones erróneas o incompletas.

### ¿Qué es un algoritmo?
Un algoritmo es un conjunto finito de pasos lógicos, secuenciales y bien definidos que resuelven un problema o realizan una tarea. Un algoritmo es una idea matemática, existe independientemente de si se programa en una computadora o no. 

Las características fundamentales de un buen algoritmo son:
1. **Precisión:** Cada paso debe estar definido claramente, sin dejar lugar a dobles interpretaciones.
2. **Finitud:** Por más largo que sea, el algoritmo debe tener un fin en algún momento.
3. **Definición / Determinismo:** Si se ejecuta el algoritmo dos veces bajo las mismas condiciones y datos de entrada, el resultado debe ser exactamente el mismo.

En computación, existe el concepto de **Pensamiento Algorítmico**, que implica descomponer un problema masivo y abrumador en partes pequeñas, entendibles y manejables (*divide y vencerás*), y buscar patrones repetitivos para resolverlos de la manera más eficiente.

### ¿Qué es un lenguaje de programación?
Es un idioma artificial, estructurado matemáticamente, diseñado para expresar algoritmos de manera que una computadora los pueda entender y ejecutar. Actúa como un puente o interfaz entre la lógica del pensamiento humano y el código máquina.

#### Niveles de Abstracción
La abstracción es un concepto crítico en informática: consiste en ocultar detalles complejos para que el programador pueda centrarse en el problema de negocio y no en la electrónica de la máquina.
- **Lenguajes de Bajo Nivel:** Están muy cerca del hardware (Ej: Lenguaje Ensamblador o Assembly). Son difíciles de leer para los humanos porque no ocultan los detalles, pero dan un control absoluto y extremadamente veloz sobre la memoria RAM y los registros del procesador.
- **Lenguajes de Alto Nivel:** Se asemejan más al lenguaje humano (inglés estructurado) y a las matemáticas (Ej: JavaScript, Python, Java). Poseen un alto nivel de abstracción; ocultan automáticamente la administración de memoria y los ciclos del procesador.

#### Compiladores e Intérpretes (¿Cómo se traduce el lenguaje?)
Dado que el procesador físico de tu computadora solo entiende ceros y unos (Código Máquina), todo lenguaje de alto nivel debe ser traducido. Hay dos formas principales:
- **Compilador:** Es un programa que toma todo tu código completo y lo traduce de un solo golpe creando un archivo "ejecutable" nuevo (como un `.exe`). Si hay un solo error en tu lógica, el archivo no se crea. (Ej: C++, Rust, Go).
- **Intérprete:** Es un programa que va leyendo tu código línea por línea y lo traduce e ejecuta en tiempo real sobre la marcha. Si hay un error en la línea 10, el programa ejecutará perfectamente hasta la línea 9 y luego estallará en un error. (Ej: JavaScript, Python).

### Paradigmas de Programación
Un paradigma es un "estilo", un enfoque filosófico o una escuela de pensamiento para estructurar y organizar el código. No se trata de sintaxis exacta, sino de la forma en que le planteamos el problema a la computadora:
- **Paradigma Imperativo:** Nos enfocamos en el *CÓMO*. Le damos a la computadora instrucciones explícitas de estado y control paso a paso.
- **Paradigma Declarativo:** Nos enfocamos en el *QUÉ*. Le pedimos a la computadora lo que queremos obtener como resultado final, y delegamos a la computadora el trabajo de decidir cómo llegar a él internamente.
- **Paradigma Orientado a Objetos (POO):** Modelamos el problema del mundo real dividiéndolo en "objetos" conceptuales. Cada objeto tiene características (estado) y acciones (comportamiento) propias, e interactúan entre sí.
