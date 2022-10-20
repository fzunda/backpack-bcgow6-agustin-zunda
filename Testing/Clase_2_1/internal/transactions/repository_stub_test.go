package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne
 dos productos con las especificaciones que deseen.
  Comprobar que GetAll() retorne la información
exactamente igual a la esperada. Para esto:
Dentro de la carpeta
/internal/(producto/usuario/transacción),
crear un archivo repository_test.go con el test
diseñado.
*/

type StubDB struct {
	Data []Transaction
}

func (fs *StubDB) Read(data interface{}) error {
	dataStub := data.(*[]Transaction)
	*dataStub = fs.Data
	return nil
}

func (fs *StubDB) Write(data interface{}) error {
	Casterdata := data.(Transaction)
	fs.Data = append(fs.Data, Casterdata)
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange
	ts := []Transaction{
		{Id: 1, Code: "001", Coin: "dolar", Amount: 15000, Emitting: "Vendedor", Receptor: "Comprador"},
	}
	myStubDB := &StubDB{Data: ts}
	repositoryts := NewRepository(myStubDB)

	//Act
	result, err := repositoryts.GetAll()

	//asserts
	assert.Nil(t, err)
	assert.Equal(t, ts, result)

}
