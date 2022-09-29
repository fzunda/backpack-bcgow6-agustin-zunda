package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...int) (float64, error) {
	var result float64
	var cantNotas float64
	var acum int

	for _, nota := range notas {
		if nota > 0 {
			acum = acum + nota
			cantNotas++
		} else {
			return 0, errors.New("La nota ingresada debe ser mayor a cero")
		}

	}

	result = float64(acum) / cantNotas
	return result, nil
}

func main() {
	result, err := promedio(6, 4, 1, 10, 8, 9)

	if err == nil {
		fmt.Println("Promedio: ", result)
	} else {
		fmt.Println("Error: ", err)
	}
}
