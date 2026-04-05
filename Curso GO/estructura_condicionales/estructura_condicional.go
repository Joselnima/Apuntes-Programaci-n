/*
Ingresar el sueldo de una persona, si supera los 3000 pesos mostrar un mensaje en pantalla
indicando que debe abonar impuestos.
*/
package main

// Para poder utilizar las funciones Print, Println y Scan
import "fmt"

func EstructuraCondicional() {
	var sueldo float32
	fmt.Print("Ingrese el sueldo:")
	fmt.Scan(&sueldo)
	if sueldo > 3000 {
		fmt.Print("Esta persona debe abonar impuestos")
	}
}
