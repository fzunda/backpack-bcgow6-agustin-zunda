package products

import domian "github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/domains"

type Service interface {
	GetByName(name string) (domian.Product, error)
	Store(prod domian.Product) (domian.Product, error)
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
