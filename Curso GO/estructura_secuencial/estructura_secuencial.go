/*
Calcular el sueldo mensual de un operario conociendo la cantidad de horas trabajadas y el pago por hora.
*/
package main

// Para poder utilizar las funciones Print, Println y Scan
import "fmt"

func EstructuraSecuencial() {
	var horasTrabajadas int
	var costoHora float32
	var sueldo float32

	fmt.Print("Ingrese las horas trabajadas por el empleado:")
	fmt.Scan(&horasTrabajadas)
	fmt.Print("Ingrese el pago por hora:")
	fmt.Scan(&costoHora)
	sueldo = float32(horasTrabajadas) * costoHora
	fmt.Print("El sueldo total del operario es ", sueldo)
}
