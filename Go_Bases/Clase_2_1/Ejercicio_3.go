package main

import "fmt"

/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas

	trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y
la categoría, y que devuelva su salario.
*/
func monthlySalary(minutes int, category string) float64 {
	var salary, hours float64

	hours = float64(minutes) / 60

	switch category {
	case "C":
		salary = hours * 1000
	case "B":
		salary = hours * 1500
		salary = salary * 1.2
	case "A":
		salary = hours * 3000
		salary = salary * 1.5
	}

	return salary
}

func main() {
	fmt.Println("Empleado Categoria C: ", monthlySalary(50, "C"))
	fmt.Println("Empleado Categoria B: ", monthlySalary(50, "B"))
	fmt.Println("Empleado Categoria A: ", monthlySalary(50, "A"))
}
