# Lección: Estructura Condicional

Duración estimada: 10-20 minutos

Objetivos:

- Comprender la instrucción `if` en Go.
- Tomar decisiones simples basadas en entradas del usuario.

Temas cubiertos:

- Sintaxis de `if` en Go.
- Comparaciones y operadores relacionales.

Descripción:

La lección solicita el sueldo de una persona y muestra un mensaje si supera 3000 pesos.

Archivos:

- `estructura_condicional.go` — contiene `EstructuraCondicional()`.
- `main.go` — envoltorio que invoca la función para ejecutar la lección.

Cómo ejecutar (PowerShell):

```powershell
cd "e:\GO courses\estructura_condicionales"
go run main.go estructura_condicional.go
```

Actividades:

1. Ejecuta la lección e ingresa distintos sueldos para ver el comportamiento.
2. Agrega `else` para mostrar un mensaje cuando no corresponde pagar impuestos.

Ejercicios sugeridos:

- Agregar `else` con mensaje `No paga impuestos`.
- Permitir ingreso repetido hasta que el usuario escriba `0` para salir.

Pista de solución:

- Usa `if sueldo > 3000 { ... } else { ... }` y un bucle `for` para repetir.
