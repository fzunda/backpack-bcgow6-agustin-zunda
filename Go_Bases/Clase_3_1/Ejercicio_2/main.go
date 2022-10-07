package main

import (
	"fmt"
	"os"
	"strings"
)

/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título
(tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total
(Sumando precio por cantidad)
*/

func readFile() {
	data, err := os.ReadFile("./myFile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file := string(data)
	rows := strings.Split(file, "\n")

	var new_rows string = fmt.Sprintf("\nEjemplo: \n\n")
	new_rows += fmt.Sprintf("ID\tPrecio\tCantidad\n")

	for _, value := range rows {
		row := strings.Replace(value, ",", "	", -1)
		new_rows += fmt.Sprintf("%v\n", row)
	}
	fmt.Println(new_rows)

}

func main() {
	readFile()
}
