package main

/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que
represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
	1. Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
	2. Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)

La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
si es cuadrática y cuál es el valor máximo.

*/

import "fmt"

type Matriz struct {
	values []float64
	alto   int
	ancho  int
}

func (m Matriz) Set() int {
	values := m.alto * m.ancho
	return values
}

func (m Matriz) Print() {

	if len(m.values) == 0 {
		fmt.Println("La Matriz esta vacia")
	} else {
		for fila := 0; fila < m.alto; fila++ {
			fmt.Printf("\t%.0f\n", m.values[fila*m.ancho:fila*m.ancho+m.ancho])
		}
	}
}

func (m Matriz) Max() float64 {
	var value float64

	for _, val := range m.values {
		if val > value {
			value = val
		}
	}
	return value
}

func (m Matriz) Cuadratica() bool {
	if m.ancho == m.alto {
		return true
	}
	return false
}

func main() {

	m := Matriz{
		ancho:  3,
		alto:   3,
		values: []float64{5, 6, 12, 15, 99, 14, 76, 8, 90},
	}
	fmt.Printf("Matriz iniciada con %v valores.\n", m.Set())
	m.Print()
	fmt.Println("Matriz Cuadratica: ", m.Cuadratica())
	fmt.Println("Maximo valor de la Matriz: ", m.Max())
}
