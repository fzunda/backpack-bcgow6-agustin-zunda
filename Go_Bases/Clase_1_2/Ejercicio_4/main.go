package main

import (
	"fmt"
)

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
var mayores int = 0

func main() {
	for i, v := range employees {
		if v > 21 {
			mayores++
		}
		fmt.Println(i, v)
	}

	fmt.Println("Empleados Adultos: ", mayores)

	delete(employees, "Pedro")
	fmt.Println(employees)
	employees["Federico"] = 25
	fmt.Println(employees)
}
