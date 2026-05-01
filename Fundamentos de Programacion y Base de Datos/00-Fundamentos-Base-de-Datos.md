# 🗄️ CURSO DE BASE DE DATOS (DE CERO A AVANZADO)

## 🧩 MÓDULO 0: Fundamentos (lo que casi todos subestiman)

### ¿Qué es una base de datos (BD)?
Es un conjunto de datos lógicamente coherente, estructurado y organizado que se almacena sistemáticamente para su uso posterior. El propósito fundamental de una base de datos no es solo "guardar información", sino garantizar que esos datos mantengan su **integridad** (no se corrompan), su **disponibilidad** y su **seguridad** a lo largo del tiempo, incluso ante caídas de sistema o fallos eléctricos.

### El Sistema Gestor de Bases de Datos (SGBD / DBMS)
Es el software complejo y robusto que actúa como intermediario entre la base de datos física, las aplicaciones y los usuarios. El SGBD abstrae la complejidad de cómo los bytes se escriben y leen del disco duro físico. Gracias a este software, el desarrollador solo necesita enviar instrucciones lógicas, y el gestor se encarga de traducirlas a operaciones seguras de hardware.

---

### Arquitecturas de Bases de Datos y sus Filosofías

#### 1. Relacionales (SQL)
Organizan los datos en **tablas** con un esquema estricto de filas y columnas. Su arquitectura está basada matemáticamente en la **Teoría de Conjuntos** y en el **Álgebra Relacional** propuesta por Edgar Codd en 1970. Están diseñadas priorizando absolutamente la estructura, la falta de anomalías y la consistencia total de la información (no contradecirse jamás).
*Ejemplos:* MySQL, PostgreSQL, Oracle Database.

#### 2. No relacionales (NoSQL)
Nacen de la necesidad de manejar "Big Data" y aplicaciones web masivas distribuidas en múltiples servidores. Al prescindir del rigor matemático de las tablas y relaciones formales, son más flexibles y escalan de manera espectacular (crecimiento horizontal). Los datos se guardan en su forma más natural, como documentos tipo JSON, redes de grafos o pares clave-valor.
*Ejemplos:* MongoDB, Redis, Cassandra, Neo4j.

---

### Conceptos Teóricos Críticos para Perfiles Avanzados

#### Las Propiedades ACID (El corazón de las BD Relacionales)
Para garantizar la fiabilidad del sistema en el mundo financiero, empresarial e industrial, las BD Relacionales cumplen 4 propiedades teóricas fundamentales para cada transacción de datos:
1. **Atomicidad (Atomicity):** El principio de "todo o nada". Una operación que involucra múltiples pasos se ejecuta por completo o no se realiza en absoluto. Si se interrumpe a la mitad (por ejemplo, en una transferencia de dinero de la cuenta A a la cuenta B), el sistema revierte todos los cambios a su estado inicial.
2. **Consistencia (Consistency):** Cualquier transacción llevará a la base de datos de un estado válido inicial a otro estado válido final, respetando estrictamente todas las reglas matemáticas y constraints (restricciones) definidas en el esquema.
3. **Aislamiento (Isolation):** Si varios usuarios ejecutan transacciones simultáneamente al mismo tiempo, el resultado final debe ser idéntico al que se obtendría si dichas transacciones se hubieran realizado en fila, una tras otra. Evita choques de datos.
4. **Durabilidad (Durability):** Una vez que una transacción ha sido exitosa (confirmada / *commit*), los cambios son permanentes y sobreviven de inmediato a un fallo abrupto del servidor o pérdida de energía.

#### El Teorema CAP (El corazón de los Sistemas Distribuidos / NoSQL)
Este teorema es obligatorio en arquitectura de software moderna. Dicta que es imposible que un sistema de datos distribuido en múltiples servidores garantice estas tres cosas al 100% al mismo tiempo:
- **Consistencia (Consistency):** Todos los usuarios, sin importar a qué servidor se conecten, leen exactamente la misma información actualizada al mismo milisegundo.
- **Disponibilidad (Availability):** Todo usuario que haga una petición siempre obtendrá una respuesta de éxito o fracaso, garantizando que el sistema jamás parece estar caído.
- **Tolerancia a Particiones (Partition tolerance):** El sistema seguirá operando incluso si la red se corta y los servidores no pueden hablar entre sí momentáneamente.
En la realidad, dado que las fallas de red existen obligatoriamente (P), un arquitecto siempre debe elegir qué sacrificar: si sacrifica disponibilidad a favor de consistencia (CP - común en bancos) o si prefiere que siempre responda aunque dé datos viejos (AP - común en redes sociales).

#### Normalización
Es el proceso formal y metodológico de organizar los datos de la base de datos aplicando una serie de "Reglas Normales". Su propósito es eliminar drásticamente la redundancia (no repetir la misma información en varios lugares) y evitar "anomalías" cuando se inserta, actualiza o borra información.

#### Índices
Conceptualmente, funcionan de la misma manera que el índice analítico de palabras al final de un libro muy grueso. Es una estructura de datos abstracta y adicional que la base de datos genera y mantiene internamente en el disco. Su propósito es **acelerar dramáticamente la búsqueda y lectura** de información. La base de datos, en lugar de escanear millones de registros uno por uno, va directo al índice y de ahí a la ubicación física del dato. 

---

### Componentes Lógicos del Modelo Relacional:

- **Entidad / Tabla:** La abstracción de un objeto del mundo real que se quiere modelar y sobre el cual queremos guardar información (ej. la entidad "Paciente").
- **Tupla / Registro (Fila):** Una ocurrencia única y materializada de esa entidad.
- **Atributo / Columna:** La característica atómica indivisible de esa entidad (ej. el atributo "Fecha de Nacimiento").
- **Clave Primaria (Primary Key - PK):** Un concepto de unicidad. Un atributo (o grupo de atributos) que se elige arquitectónicamente para distinguir e identificar de manera irrepetible cada tupla dentro de una tabla.
- **Clave Foránea (Foreign Key - FK):** El concepto de referencia de negocio. Es un atributo dentro de una tabla cuyo valor coincide obligatoriamente con la Clave Primaria de otra tabla, estableciendo así el "vínculo" conceptual entre ambas entidades.
- **Modelo Entidad-Relación (ER):** La representación abstracta, formal y gráfica (el diagrama maestro) que detalla el universo del negocio: cuáles son sus Entidades, qué Atributos tienen y, especialmente, qué Cardinalidad existe entre ellas (cómo se relacionan de 1 a 1, de 1 a Muchos o de Muchos a Muchos).
