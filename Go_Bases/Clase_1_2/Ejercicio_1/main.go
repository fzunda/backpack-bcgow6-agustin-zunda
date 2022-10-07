package main

import (
	"fmt"
	"strings"
)

var nombre string = "Agustin"

func main() {
	texto := strings.Split(nombre, "")
	fmt.Println("Cantidad Letras: ", len(texto))

	for i := 0; i < len(texto); i++ {
		fmt.Println(texto[i])
	}
}
