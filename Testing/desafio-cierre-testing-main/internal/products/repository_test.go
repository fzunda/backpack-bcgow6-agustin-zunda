package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {
	id := "FEX112AC"
	prodList := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}

	repo := NewRepository()
	result, err := repo.GetAllBySeller(id)

	assert.Empty(t, err)
	assert.Equal(t, prodList, result)
}
