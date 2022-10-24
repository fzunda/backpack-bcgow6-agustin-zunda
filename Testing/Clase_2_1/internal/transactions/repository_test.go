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

/*
Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre
de un producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa el método “Read” del
Storage para buscar el producto. Para esto:
	1. Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico cuyo
	 nombre puede ser “Before Update”.
	2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede
	ser a través de un boolean como se observó en la clase.
	3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del
	producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización.
	También debe validarse que el método Read haya sido ejecutado durante el test.
*/

type MockRepositoryTransaction struct {
	UpdateCodeAndAmountCalled bool
	Data                      []Transaction
}

func (ts *MockRepositoryTransaction) Read(data interface{}) error {
	ts.UpdateCodeAndAmountCalled = true
	dataRead := data.(*[]Transaction)
	*dataRead = ts.Data
	return nil
}

func (ts *MockRepositoryTransaction) Write(data interface{}) error {
	return nil
}

func TestUpdateCodeAndAmount(t *testing.T) {
	//arrange
	id, code := 1, "0015"
	var amount float64 = 250000
	ts := []Transaction{
		{Id: 1, Code: "001", Coin: "dolar", Amount: 15000, Emitting: "Vendedor", Receptor: "Comprador"},
	}
	myMock := MockRepositoryTransaction{Data: ts}
	repositoryts := NewRepository(&myMock)

	//Act
	result, err := repositoryts.UpdateCodeAndAmount(id, code, amount)

	//asserts
	assert.Nil(t, err)
	assert.Equal(t, id, result.Id)
	assert.Equal(t, code, result.Code)
	assert.Equal(t, amount, result.Amount)
}

func TestUpdate(t *testing.T) {
	//arrange
	id, code, coin, emitting, receptor := 1, "0015", "Euro", "Vendedor23", "Comprador23"
	var amount float64 = 300000

	ts := []Transaction{
		{Id: 1, Code: "001", Coin: "dolar", Amount: 15000, Emitting: "Vendedor", Receptor: "Comprador"},
	}
	myMock := MockRepositoryTransaction{Data: ts}
	repositoryts := NewRepository(&myMock)

	//Act
	result, err := repositoryts.Update(id, code, coin, amount, emitting, receptor)

	//asserts
	assert.Nil(t, err)
	assert.Equal(t, id, result.Id)
	assert.Equal(t, code, result.Code)
	assert.Equal(t, amount, result.Amount)
	assert.Equal(t, emitting, result.Emitting)
	assert.Equal(t, receptor, result.Receptor)
}

func TestStore(t *testing.T) {
	//arrange
	ts := []Transaction{}
	id, code, coin := 1, "001", "dolar"
	var amount float64 = 15000
	emitting, receptor := "Vendedor", "Comprador"

	myMock := MockRepositoryTransaction{Data: ts}
	repositoryts := NewRepository(&myMock)
	//Act
	result, err := repositoryts.Store(id, code, coin, amount, emitting, receptor)

	//asserts
	assert.Nil(t, err)
	assert.Equal(t, id, result.Id)
	assert.Equal(t, code, result.Code)
	assert.Equal(t, coin, result.Coin)
	assert.Equal(t, amount, result.Amount)
	assert.Equal(t, emitting, result.Emitting)
	assert.Equal(t, receptor, result.Receptor)
}

func TestDelete(t *testing.T) {
	//arrange
	id := 1
	id2 := 15
	ts := []Transaction{
		{Id: 1, Code: "001", Coin: "dolar", Amount: 15000, Emitting: "Vendedor", Receptor: "Comprador"},
	}

	myMock := MockRepositoryTransaction{Data: ts}
	repositoryts := NewRepository(&myMock)
	//Act
	err := repositoryts.Delete(id)
	err2 := repositoryts.Delete(id2)

	//asserts
	assert.Nil(t, err)
	assert.NotNil(t, err2)
}
