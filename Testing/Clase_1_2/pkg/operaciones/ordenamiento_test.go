package operaciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Ejercicio 2 - Test Unitario Método Ordenar
Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente, posteriormente diseñar
un test unitario que valide el funcionamiento del mismo.
	1. Dentro de la carpeta go-testing crear un archivo ordenamiento.go con la función a probar.
	2. Dentro de la carpeta go-testing crear un archivo ordenamiento_test.go con el test diseñado.

*/

func TestListOrdenadaMayorMenor(t *testing.T) {
	//arrange
	ListValue := []int{10, 5, 14, 2, 100}
	ExpectedResult := []int{2, 5, 10, 14, 100}
	//act
	Result := ListOrdenadaMayorMenor(ListValue, 5)
	//assert
	assert.Equal(t, Result, ExpectedResult)
}
