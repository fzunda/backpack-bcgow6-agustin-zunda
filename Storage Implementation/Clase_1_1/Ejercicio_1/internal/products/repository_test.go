package products

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domian "github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/domains"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))
	productId := 1
	repo := NewRepository(db)
	prod := domian.Product{
		ID:    productId,
		Name:  "Botines F5",
		Type:  "Deportivo",
		Count: 40,
		Price: 25000.99,
	}

	p, err := repo.Store(prod)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, p.ID, productId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryStoreFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnError(sql.ErrConnDone)
	productId := 1
	repo := NewRepository(db)
	prod := domian.Product{
		ID:    productId,
		Name:  "Botines F5",
		Type:  "Deportivo",
		Count: 40,
		Price: 25000.99,
	}

	p, err := repo.Store(prod)
	assert.Equal(t, sql.ErrConnDone, err)
	assert.Empty(t, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryStoreRowFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare("")
	mock.ExpectExec("INSERT INTO products").WillReturnError(sql.ErrNoRows)
	productId := 1
	repo := NewRepository(db)
	prod := domian.Product{
		ID: productId,
	}

	p, err := repo.Store(prod)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.Empty(t, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryGetByName(t *testing.T) {
	//instanciar el,mock para test
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	//creo las columnas del resultado y el objeto a usar
	rows := sqlmock.NewRows([]string{"id", "name", "type", "count", "price"})
	prod := domian.Product{
		ID:    1,
		Name:  "Botines",
		Type:  "Deportivo",
		Count: 40,
		Price: 25000.99,
	}

	rows.AddRow(prod.ID, prod.Name, prod.Type, prod.Count, prod.Price)
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT id, name, type, count, price FROM  products WHERE name = ?;")).ExpectQuery().WithArgs(prod.Name).WillReturnRows(rows)

	repo := NewRepository(db)
	result, err := repo.GetByName(prod.Name)
	assert.NoError(t, err)
	assert.NotZero(t, result)
	assert.Equal(t, prod, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "type", "count", "price"})
	prod := domian.Product{
		ID:    1,
		Name:  "Botines",
		Type:  "Deportivo",
		Count: 40,
		Price: 25000.99,
	}

	rows.AddRow(prod.ID, prod.Name, prod.Type, prod.Count, prod.Price)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, type, count, price FROM  products;")).WillReturnRows(rows)

	repo := NewRepository(db)
	result, err := repo.GetAll()

	assert.NoError(t, err)
	assert.NotZero(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, type, count, price FROM  products;")).WillReturnError(sql.ErrConnDone)
	repo := NewRepository(db)
	result, err := repo.GetAll()

	assert.Equal(t, fmt.Errorf("error al preparar la consulta - error %v", sql.ErrConnDone), err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

/*func TestGetAllRowFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "count", "price"})
	prod := domian.Product{
		ID:    1,
		Name:  "Botines",
		Count: 40,
		Price: 25000.99,
	}

	rows.AddRow(prod.ID, prod.Name, prod.Type, prod.Count, prod.Price)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, type, count, price FROM  products;")).WillReturnError(sql.ErrTxDone)
	repo := NewRepository(db)
	result, err := repo.GetAll()

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}*/

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := 1

	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id = ?;")).ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewRepository(db)
	err = repo.Delete(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := 1

	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id = ?;")).ExpectExec().WillReturnError(sql.ErrConnDone)
	repo := NewRepository(db)
	err = repo.Delete(id)
	assert.Equal(t, sql.ErrConnDone, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productId := 1
	prod := domian.Product{
		ID:    productId,
		Name:  "Medias",
		Type:  "Deportivo",
		Count: 150,
		Price: 200.99,
	}

	mock.ExpectPrepare("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?;").ExpectExec().WithArgs(prod.Name, prod.Type, prod.Count, prod.Price).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)
	_, err = repo.Update(productId, prod.Name, prod.Type, prod.Count, prod.Price)
	assert.NoError(t, err)
	//assert.NoError(t, mock.ExpectationsWereMet())
}
