package transactions

import (
	"time"
)

type Service interface {
	GetAll() ([]Transaction, error)
	Store(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error)
	Update(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error)
	UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	ts, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) Store(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return Transaction{}, err
	}

	lastId++
	transaccion, err := s.repository.Store(id, code, coin, amount, emitting, receptor, date)
	if err != nil {
		return Transaction{}, err
	}
	return transaccion, nil
}

func (s *service) Update(id int, code string, coin string, amount float64, emitting string, receptor string, date time.Time) (Transaction, error) {
	return s.repository.Update(id, code, coin, amount, emitting, receptor, date)
}

func (s *service) UpdateCodeAndAmount(id int, code string, amount float64) (Transaction, error) {
	return s.repository.UpdateCodeAndAmount(id, code, amount)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
