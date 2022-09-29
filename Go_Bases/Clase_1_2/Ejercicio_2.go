package main

import (
	"fmt"
)

var nombres = [3]string{"Agustin", "Eduardo", "Matias"}
var edades = [3]int{19, 30, 26}
var empleos = [3]int{2, 1, 3}
var sueldos = [3]int{80000, 150000, 300000}

func main() {
	for i := 0; i < len(nombres); i++ {
		if edades[i] > 22 && empleos[i] >= 1 {
			fmt.Printf("Se le otorga prestamo al empleado %v \n", nombres[i])
			if sueldos[i] > 100000 {
				fmt.Println("Sin intereses")
			} else {
				fmt.Println("Con intereses")
			}
		}
	}

}
