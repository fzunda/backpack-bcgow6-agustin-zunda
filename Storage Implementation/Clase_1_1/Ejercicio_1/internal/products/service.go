package products

import domian "github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/domains"

type Service interface {
	GetByName(name string) (domian.Product, error)
	Store(prod domian.Product) (domian.Product, error)
	GetAll() ([]domian.Product, error)
	Update(id int, name string, productType string, count int, price float64) (domian.Product, error)
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

func (s *service) GetByName(name string) (domian.Product, error) {
	return s.repository.GetByName(name)
}

func (s *service) Store(prod domian.Product) (domian.Product, error) {
	return s.repository.Store(prod)
}

func (s *service) GetAll() ([]domian.Product, error) {
	return s.repository.GetAll()
}

func (s *service) Update(id int, name string, productType string, count int, price float64) (domian.Product, error) {
	return s.repository.Update(id, name, productType, count, price)
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
