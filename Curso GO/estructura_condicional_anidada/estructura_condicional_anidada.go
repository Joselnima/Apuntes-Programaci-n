package main

// Para poder utilizar las funciones Print, Println y Scan
import "fmt"

func EstructuraCondicionalAnidada() {
	var nota1, nota2, nota3, promedio int
	fmt.Print("Ingrese primer nota:")
	fmt.Scan(&nota1)
	fmt.Print("Ingrese segunda nota:")
	fmt.Scan(&nota2)
	fmt.Print("Ingrese tercer nota:")
	fmt.Scan(&nota3)
	promedio = (nota1 + nota2 + nota3) / 3
	if promedio >= 7 {
		fmt.Print("Promocionado")
	} else {
		if promedio >= 4 {
			fmt.Print("Regular")
		} else {
			fmt.Print("Reprobado")
		}
	}
}
