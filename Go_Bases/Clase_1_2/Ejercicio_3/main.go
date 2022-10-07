package main

import (
	"fmt"
)

var dia int = 1

var meses = [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

func main() {

	if dia > 0 && dia < 13 {
		dia--
		fmt.Println(meses[dia])
	} else {
		fmt.Println("El mes ingresado no existe")
	}

}
