package products

import "errors"

type MockRepository struct {
	Data []Product
}

func (m *MockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	var productsList []Product

	for _, prod := range m.Data {
		if prod.SellerID == sellerID {
			productsList = append(productsList, prod)
		}
	}

	if len(productsList) == 0 {
		return nil, errors.New("no products was found")
	}

	return productsList, nil
}
