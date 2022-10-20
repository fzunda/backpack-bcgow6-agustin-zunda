package transactions

import (
	"fmt"
	"time"

	"Clase_2_1/pkg/store"
)

type Transaction struct {
	Id              int       `json:"id"`
	Code            string    `json:"code"`
	Coin            string    `json:"coin"`
	Amount          float64   `json:"amount"`
	Emitting        string    `json:"emitting"`
	Receptor        string    `json:"Receptor"`
	TransactionDate time.Time `json:"transactionDate"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error)
	LastId() (int, error)
	Update(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error)
	UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error)
	Delete(id int) error
}

// struct que implementa todos los metodos de la interface
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() (result []Transaction, err error) {

	err = r.db.Read(&result)
	return
}

func (r *repository) LastId() (int, error) {

	var ts []Transaction
	err := r.db.Read(ts)
	if err != nil {
		return 0, err
	}
	if len(ts) == 0 {
		return 0, nil
	}

	return ts[len(ts)-1].Id, nil
}

func (r *repository) Store(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error) {
	var ts []Transaction

	date = time.Now()

	err := r.db.Read(&ts)
	if err != nil {
		return Transaction{}, err
	}

	t := Transaction{Id: id, Code: code, Coin: coin, Amount: amount, Emitting: emitting, Receptor: receptor, TransactionDate: date}
	ts = append(ts, t)

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}

	return t, nil
}

func (r *repository) Update(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error) {
	t := Transaction{Code: code, Coin: coin, Amount: amount, Emitting: emitting, Receptor: receptor, TransactionDate: time.Now()}
	updated := false

	var ts []Transaction
	err := r.db.Read(&ts)
	if err != nil {
		return Transaction{}, err
	}

	for i := range ts {
		if ts[i].Id == id {
			t.Id = id
			ts[i] = t
			updated = true
		}
	}

	if !updated {
		return Transaction{}, fmt.Errorf("Transacción %d no encontrada", id)
	}

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func (r *repository) UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error) {
	var update bool
	t := Transaction{}

	date := time.Now()
	var ts []Transaction
	err := r.db.Read(&ts)
	if err != nil {
		return Transaction{}, err
	}

	for i := range ts {
		if ts[i].Id == id {
			ts[i].Code = code
			ts[i].Amount = amount
			ts[i].TransactionDate = date
			t = ts[i]
			update = true
		}
	}
	if !update {
		return Transaction{}, fmt.Errorf("Transacción %d no encontrada", id)
	}

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}

	return t, nil
}

func (r *repository) Delete(id int) error {
	var deleted bool
	var pos int

	var ts []Transaction
	err := r.db.Read(&ts)
	if err != nil {
		return err
	}

	for i := range ts {
		if ts[i].Id == id {
			fmt.Println("Id: ", ts[i].Id)
			pos = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Transacción %d no encontrada", pos)
	}
	ts = append(ts[:pos], ts[pos+1:]...)

	if err := r.db.Write(ts); err != nil {
		return err
	}

	return nil
}
