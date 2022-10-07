package main

/*
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo
 en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

import (
	"fmt"
	"os"
)

type myCustomeError struct {
	message string
}

func (err *myCustomeError) Error() string {
	//return fmt.Sprintf("error: el salario ingresado no alcanza el mínimo imponible")
	err.message = "error: el salario ingresado no alcanza el mínimo imponible"
	return err.message
}

func testMyCustomerError(salary float64) string {
	var err = myCustomeError{}
	if salary < 150000 {
		return err.Error()
	}
	return ""
}

func main() {
	var salary float64 = 140000
	err := testMyCustomerError(salary)
	if err != "" {
		fmt.Println(err)
		os.Exit(1)
	}
	println("Debe pagar impuesto")
}
