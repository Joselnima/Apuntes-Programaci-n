# Lección: Estructura Secuencial

Duración estimada: 15-25 minutos

Objetivos:

- Entender el flujo secuencial en un programa Go.
- Leer entrada desde consola y realizar cálculos básicos.

Temas cubiertos:

- `package main`, `func main()` y llamadas a funciones.
- Lectura con `fmt.Scan` y conversiones entre tipos numéricos.

Descripción:

Esta lección muestra cómo calcular el sueldo de un operario a partir de las horas trabajadas y el pago por hora.

Archivos:

- `estructura_secuencial.go` — lógica de la lección (`EstructuraSecuencial()`).
- `main.go` — envoltorio que llama a `EstructuraSecuencial()` para ejecutar la lección.

Cómo ejecutar (PowerShell):

```powershell
cd "e:\GO courses\estructura_secuencial"
go run main.go estructura_secuencial.go
```

Actividad (pasos):

1. Ejecuta la lección y proporciona horas y pago por hora.
2. Observa el resultado y modifica los datos de entrada.

Ejercicios sugeridos:

- Añadir cálculo de horas extras (1.5x) si las horas > 40.
- Validar la entrada: rechazar horas negativas.

Pista de solución:

- Implementa una condición `if horas > 40 { ... }` que calcule extras y los sume al sueldo base.
