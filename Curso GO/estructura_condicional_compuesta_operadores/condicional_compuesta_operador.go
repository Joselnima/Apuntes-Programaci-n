/*
Se carga una fecha (día, mes y año) por teclado.
Mostrar un mensaje si corresponde al primer trimestre del año (enero, febrero o marzo)
Cargar por teclado el valor numérico del día, mes y año.
Ejemplo: dia:10 mes:5 año:2017.
*/
package main

import "fmt"

func CondicionalCompuestaOperador() {
	var dia, mes, año int
	fmt.Print("Ingrese número de día:")
	fmt.Scan(&dia)
	fmt.Print("Ingrese número de mes:")
	fmt.Scan(&mes)
	fmt.Print("Ingrese número de año:")
	fmt.Scan(&año)
	if mes == 1 || mes == 2 || mes == 3 {
		fmt.Print("Corresponde al primer trimestre")
	}
}
