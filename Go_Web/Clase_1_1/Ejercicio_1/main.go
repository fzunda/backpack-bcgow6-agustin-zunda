package main

/*
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
	- Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
	- Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
	- Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.
Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.

*/

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type transaction struct {
	Id              int
	Code            string
	Coin            string
	Amount          float64
	Emitting        string
	Receptor        string
	TransactionDate time.Time
}

var transaction1 = transaction{Id: 1, Code: "123465A", Coin: "Dolar", Amount: 299, Emitting: "Central Bank", Receptor: "Bank of Córdoba", TransactionDate: time.Now()}
var transaction2 = transaction{Id: 2, Code: "123465B", Coin: "Euro", Amount: 315, Emitting: "Central Bank", Receptor: "Santander Río", TransactionDate: time.Now()}
var transaction3 = transaction{Id: 3, Code: "123465C", Coin: "Peso Chileno", Amount: 152, Emitting: "Central Bank", Receptor: "Galicia", TransactionDate: time.Now()}

func main() {
	var arrayTransactions = make([]transaction, 0)
	arrayTransactions = append(arrayTransactions, transaction1)
	arrayTransactions = append(arrayTransactions, transaction2)
	arrayTransactions = append(arrayTransactions, transaction3)
	jsonData, err := json.Marshal(arrayTransactions)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Printf("Transacción: \n %s \n", string(jsonData))
}
