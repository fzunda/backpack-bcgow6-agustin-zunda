package main

import (
	"errors"
	"fmt"
)

const (
	OperadorSuma           = "+"
	OperadorResta          = "-"
	OperadorMultiplicacion = "*"
	OperadorDivision       = "/"
)

func operacionAritmetica(num1, num2 float64, operador string) float64 {
	//var result float64

	switch operador {
	case OperadorSuma:
		return num1 + num2
	case OperadorResta:
		return num1 - num2
	case OperadorMultiplicacion:
		return num1 * num2
	case OperadorDivision:
		if num2 > 0 {
			return num1 / num2
		}
	}
	return 0
}

func sum(numeros ...float64) float64 {
	var result float64

	for _, value := range numeros {
		result += value
	}
	return result
}

func calcular(num1, num2 float64) (float64, float64, float64, float64) {
	suma := num1 + num2
	resta := num1 - num2
	division := num1 / num2
	multiplication := num1 * num2

	return suma, resta, division, multiplication
}

func dividir(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("el divisor no puede ser cero, tiene que ser mayor o igual a uno")
	}
	result := dividendo / divisor
	return result, nil
}

func main() {
	fmt.Println(operacionAritmetica(10, 5.5, OperadorSuma))
	fmt.Println(operacionAritmetica(10, 5.5, OperadorResta))
	fmt.Println(operacionAritmetica(10, 5.5, OperadorMultiplicacion))
	fmt.Println(operacionAritmetica(10, 0, OperadorDivision))
	fmt.Println(sum(1, 2, 3, 4, 5, 31, 52))

	suma, resta, mult, div := calcular(10, 15)

	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println(mult)
	fmt.Println(div)

	result, err := dividir(10, 2)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
