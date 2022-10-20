package operaciones

/*
Ejercicio 1 - Test Unitario Restar

Para el método Restar() visto en la clase, realizar el test unitario correspondiente. Para esto:
Dentro de la carpeta go-testing crear un archivo calculadora.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo calculadora_test.go con el test diseñado.
*/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 5
	num2 := 4
	resultadoEsperado := 1

	resultado := Restar(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("Fuct restar arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
	}
}

func TestDividir(t *testing.T) {
	//arrange
	num1 := 20
	num2 := 4
	ExpectedResult := 5

	//act
	Result, err := Dividir(num1, num2)

	//assert
	assert.Equal(t, Result, ExpectedResult)
	assert.Nil(t, err, nil)
}
