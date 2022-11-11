package products

import (
	"database/sql"
	"fmt"
	"log"

	domian "github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/domains"
)

type Repository interface {
	GetByName(name string) (domian.Product, error)
	Store(prod domian.Product) (domian.Product, error)
	GetAll() ([]domian.Product, error)
	Update(id int, name string, productType string, count int, price float64) (domian.Product, error)
	Delete(id int) error
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

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

func (r *repository) GetAll() ([]domian.Product, error) {
	var products []domian.Product
	//Se prepara la consulta
	rows, err := r.db.Query("SELECT id, name, type, count, price FROM  products;")
	if err != nil {
		return []domian.Product{}, fmt.Errorf("error al preparar la consulta - error %v", err)

	}
	//Se recorren las filas de la query

	for rows.Next() {
		var prod domian.Product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Type, &prod.Count, &prod.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		products = append(products, prod)
	}
	return products, nil
}

func (r *repository) Update(id int, name string, productType string, count int, price float64) (domian.Product, error) {

	stmt, err := r.db.Prepare("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?;")
	if err != nil {
		return domian.Product{}, nil
	}

	defer stmt.Close()
	prod := domian.Product{Name: name, Type: productType, Count: count, Price: price, ID: id}
	_, err = stmt.Exec(name, productType, count, price, id)
	if err != nil {
		return domian.Product{}, nil
	}
	return prod, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM products WHERE id = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
