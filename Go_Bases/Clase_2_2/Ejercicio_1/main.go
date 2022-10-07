package main

import "fmt"

type students struct {
	name    string
	surname string
	dni     int
	date    string
}

func (x students) studentsDetail() {
	fmt.Println("-----------------------")
	fmt.Println("Nombre: ", x.name)
	fmt.Println("Apellido: ", x.surname)
	fmt.Println("Dni: ", x.dni)
	fmt.Println("Fecha: ", x.date)
}

var a = students{name: "Martin", surname: "Día", dni: 12345678, date: "25/01/87"}
var b = students{name: "Jose", surname: "Perez", dni: 12345661, date: "22/01/90"}
var c = students{name: "Sofia", surname: "Gomez", dni: 1334577, date: "01/04/91"}
var d = students{name: "Maria", surname: "Aguirre", dni: 13345566, date: "01/01/87"}
var e = students{name: "Martin", surname: "Ludueña", dni: 12344478, date: "15/01/95"}

func main() {
	a.studentsDetail()
	b.studentsDetail()
	c.studentsDetail()
	d.studentsDetail()
	e.studentsDetail()
}
