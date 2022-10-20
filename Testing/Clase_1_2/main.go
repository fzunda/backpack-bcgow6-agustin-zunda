package main

import (
	"Clase_1_2/pkg/operaciones"
	"fmt"
)

func main() {
	num := operaciones.Restar(5, 4)
	fmt.Println("Resta: ", num)
	numericArray := [5]int{10, 5, 14, 2, 100}
	lista := operaciones.ListOrdenadaMayorMenor(numericArray[:], 5)
	fmt.Println(lista)
	div, err := operaciones.Dividir(20, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Divición: ", div)
}
