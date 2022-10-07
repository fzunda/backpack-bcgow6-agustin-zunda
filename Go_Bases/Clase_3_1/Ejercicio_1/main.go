package main

import (
	"fmt"
	"os"
)

/*
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable

*/

type product struct {
	id_product int
	price      float64
	amount     int
}

var p1 = product{id_product: 1, price: 500.0, amount: 4}
var p2 = product{id_product: 2, price: 1500.0, amount: 10}
var p3 = product{id_product: 3, price: 5000.0, amount: 40}

func createFile() {
	message := fmt.Sprintf("%v,%v,%v\n", p1.id_product, p1.price, p1.amount)
	message += fmt.Sprintf("%v,%v,%v\n", p2.id_product, p2.price, p2.amount)
	message += fmt.Sprintf("%v,%v,%v\n", p3.id_product, p3.price, p3.amount)
	fmt.Println(message)
	d1 := []byte(message)
	err := os.WriteFile("./myFile.txt", d1, 777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	createFile()
}
