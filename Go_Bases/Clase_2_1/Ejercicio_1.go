package main

import "fmt"

func impuesto(salario float64) float64 {
	var desc1, desc2, result float64

	if salario >= 50000 {
		desc1 = salario * 0.17
		if salario > 150000 {
			desc2 = salario * 0.1
		}
	}

	result = salario - desc1 - desc2
	return result
}

func main() {
	fmt.Println("Salario total: ", impuesto(160000))
}
