package products

import (
	"database/sql"
	"fmt"

	domian "github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/domains"
)

type Repository interface {
	GetByName(name string) (domian.Product, error)
	Store(prod domian.Product) (domian.Product, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

/*
	Ejercicio 1 - Implementar GetByName()

Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:
Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorna una estructura del tipo Product.
Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.
*/
func (r *repository) GetByName(name string) (domian.Product, error) {
	stmt, err := r.db.Prepare("SELECT id, name, type, count, price FROM  products WHERE name = ?;")
	if err != nil {
		return domian.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}

	defer stmt.Close()

	var prod domian.Product
	err = stmt.QueryRow(name).Scan(&prod.ID, &prod.Name, &prod.Type, &prod.Count, &prod.Price)
	if err != nil {
		return domian.Product{}, fmt.Errorf("no registros para %s - error %v", name, err)
	}
	return prod, nil
}

/*
	Ejercicio 2 - Replicar Store()

Tomar el ejemplo visto en la clase y diseñar el método Store():
Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
Implementar el método Store.
*/
func (r *repository) Store(prod domian.Product) (domian.Product, error) {
	stmt, err := r.db.Prepare("INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?);")
	if err != nil {
		return domian.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(&prod.Name, &prod.Type, &prod.Count, &prod.Price)
	if err != nil {
		return domian.Product{}, err
	}
	insertedId, _ := result.LastInsertId()
	prod.ID = int(insertedId)

	return prod, nil
}
