package main

/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var salary float64 = 160000

	if salary < 150000 {
		fmt.Println(errors.New("error: El salario ingresado no alcanza el mínimo imponible"))
		os.Exit(1)
	}
	println("Debe pagar impuesto")
}
