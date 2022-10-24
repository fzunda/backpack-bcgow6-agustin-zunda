package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySellerSuccess(t *testing.T) {
	//arrange
	var (
		id string = "FEX112AC"
	)

	tp := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		}, {
			ID:          "mock2",
			SellerID:    "FEX112AC",
			Description: "generic product 1",
			Price:       500.50,
		}, {
			ID:          "mock3",
			SellerID:    "FEX115263A",
			Description: "generic product2",
			Price:       254.50,
		},
	}

	resultExpected := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		}, {
			ID:          "mock2",
			SellerID:    "FEX112AC",
			Description: "generic product 1",
			Price:       500.50,
		},
	}
	//Act
	MockRepo := MockRepository{
		Data: tp,
	}
	serv := NewService(&MockRepo)
	result, err := serv.GetAllBySeller(id)
	//asserts
	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result)
}

func TestGetAllBySellerError(t *testing.T) {
	//arrange
	var incorrectId string = "FA45A1"
	errExpected := errors.New("no products was found")
	//Act
	MockRepo := MockRepository{
		Data: nil,
	}
	serv := NewService(&MockRepo)
	//asserts
	result, err := serv.GetAllBySeller(incorrectId)
	assert.Nil(t, result)
	assert.EqualError(t, errExpected, err.Error())
}
