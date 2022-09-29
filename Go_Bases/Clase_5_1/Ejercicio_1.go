package main

/*Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones. Para ello, cuentan con todo el detalle
necesario en un archivo .txt.
	1. Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo, no han pasado el archivo a leer por nuestro
	programa.
	2. Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt” (recuerda lo visto sobre el pkg “os”).
	Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al intentar leer un archivo que no
	existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.

*/
import (
	"fmt"
	"os"
	"strings"
)

type Employee struct {
	id          int
	name        string
	surname     string
	dni         int
	grossSalary float64
}

var emp1 = Employee{id: 1, name: "Estela", surname: "Perez", dni: 12346888, grossSalary: 200000}
var emp2 = Employee{id: 2, name: "Gerardo", surname: "Gomez", dni: 12346777, grossSalary: 250000}
var emp3 = Employee{id: 1, name: "Carola", surname: "Rodriguez", dni: 12345711, grossSalary: 150000}

func createFile() (string, error) {
	message := fmt.Sprintf("%v,%v,%v,%v,%v\n", emp1.id, emp1.name, emp1.surname, emp1.dni, emp1.grossSalary)
	message += fmt.Sprintf("%v,%v,%v,%v,%v\n", emp2.id, emp2.name, emp2.surname, emp2.dni, emp2.grossSalary)
	message += fmt.Sprintf("%v,%v,%v,%v,%v\n", emp3.id, emp3.name, emp3.surname, emp3.dni, emp3.grossSalary)
	fmt.Println(message)
	d1 := []byte(message)
	err := os.WriteFile("./myFile.txt", d1, 777)
	if err != nil {
		return "", err
	}
	return "", nil
}

func readFile() (string, error) {

	data, err := os.ReadFile("./myFile.txt")
	if err != nil {
		return "", err
	}
	file := string(data)
	rows := strings.Split(file, "\n")

	var new_rows string = fmt.Sprintf("\nEjemplo: \n\n")
	new_rows += fmt.Sprintf("ID\tNombre\tApellido\tDni\tSalario\n")

	for _, value := range rows {
		row := strings.Replace(value, ",", "\t", -1)
		new_rows += fmt.Sprintf("%v\n", row)
	}
	fmt.Println(new_rows)
	return "", nil
}

func main() {

	defer func() {
		fmt.Println("\nFinalizado ...")
	}()

	fmt.Println("\nIniciando ...")
	_, err := createFile()
	if err != nil {
		panic(err)
	}
	_, err2 := readFile()
	if err2 != nil {
		panic(err2)
	}
}
