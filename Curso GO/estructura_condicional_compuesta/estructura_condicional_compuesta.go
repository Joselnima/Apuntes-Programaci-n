/*
Ejercicio 1: Estructura condicional compuesta
Escribir un programa que solicite al usuario ingresar dos valores numéricos enteros y luego
determine cuál de los dos valores es mayor. Utilizar una estructura condicional compuesta (if-else)
*/
package main

// Para poder utilizar las funciones Print, Println y Scan
import "fmt"

func EstructuraCondicionalCompuesta() {

	var num1, num2 int
	fmt.Print("Ingrese el primer valor:")
	fmt.Scan(&num1)
	fmt.Print("Ingrese el segundo valor:")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Print("El mayor es ", num1)
	} else {
		fmt.Print("El menor es ", num2)
	}
}
