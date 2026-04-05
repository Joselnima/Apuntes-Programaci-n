# Lección: Estructura Condicional Compuesta

# Lección: Estructura Condicional Compuesta

Duración estimada: 15-25 minutos

Objetivos:

- Aprender el uso de `if-else` para seleccionar entre dos opciones.
- Manejar la comparación entre valores y casos de igualdad.

Temas cubiertos:

- `if-else` y flujo alternativo.
- Comparaciones y mensajes de salida.

Descripción:

Se solicitan dos enteros y se determina cuál es mayor; incluye el uso de `else` para manejar el caso contrario.

Archivos:

- `estructura_condicional_compuesta.go` — contiene `EstructuraCondicionalCompuesta()`.
- `main.go` — envoltorio para ejecutar la lección.

Cómo ejecutar (PowerShell):

```powershell
cd "e:\GO courses\estructura_condicional_compuesta"
go run main.go estructura_condicional_compuesta.go
```

Actividades:

1. Ejecuta la lección con distintos pares de números.
2. Modifica el programa para detectar y mostrar `Son iguales` cuando corresponde.

Ejercicios sugeridos:

- Manejar igualdad con un tercer mensaje.
- Extender a una lista de números y devolver el mayor.

Pista de solución:

- Añade una condición `if num1 == num2 { fmt.Print("Son iguales") } else if num1 > num2 { ... } else { ... }`.
