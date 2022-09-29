package main

import "fmt"

/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
 de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus
 calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o
 promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se
le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func funcMinimum(notes ...int) float64 {
	var min int = notes[0]
	var result float64

	for _, value := range notes {
		if min > value {
			min = value
		}
	}
	result = float64(min)
	return result
}

func funcMaximum(notes ...int) float64 {
	var max int = notes[0]
	var result float64

	for _, value := range notes {
		if max < value {
			max = value
		}
	}
	result = float64(max)
	return result
}

func funcAverage(notes ...int) float64 {
	var result, count float64

	for _, value := range notes {
		result += float64(value)
		count++
	}
	result = result / count
	return result
}

func funcOperator(x string) func(notes ...int) float64 {
	switch x {
	case "minimum":
		return funcMinimum
	case "maximum":
		return funcMaximum
	case "average":
		return funcAverage
	default:
		fmt.Println("La operación solicitada no existe")
	}
	return nil
}

func main() {
	test1 := funcOperator("minimum")
	if test1 != nil {
		fmt.Println(test1(5, 5, 3, 2, 8, 1))
	}

	test2 := funcOperator("maximum")
	if test2 != nil {
		fmt.Println(test2(5, 5, 3, 2, 8, 1))
	}

	test3 := funcOperator("average")
	if test3 != nil {
		fmt.Println(test3(5, 5, 3, 2, 8, 1))
	}

	test4 := funcOperator("addition")
	if test4 != nil {
		fmt.Println(test4(5, 5, 3, 2, 8, 1))
	}
}
