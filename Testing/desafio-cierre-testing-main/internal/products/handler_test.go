package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockRepository MockRepository) *gin.Engine {
	service := NewService(&mockRepository)
	handler := NewHandler(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("", handler.GetProducts)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	//guarda la response que obtiene el servidor
	return req, httptest.NewRecorder()
}

func TestGetProductsSuccess(t *testing.T) {
	//arrange
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

	repository := MockRepository{
		Data: tp,
	}

	//server creation
	r := createServer(repository)

	//request creation
	request, recorder := createRequestTest(http.MethodGet, "/products?seller_id=FEX112AC", "")

	//act
	r.ServeHTTP(recorder, request)
	//asserts
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetProductsFailNotSellerId(t *testing.T) {
	//arrange
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

	repository := MockRepository{
		Data: tp,
	}

	//server creation
	r := createServer(repository)

	//request creation
	request, recorder := createRequestTest(http.MethodGet, "/products", "")

	//act
	r.ServeHTTP(recorder, request)

	//asserts
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestGetProductsFailBadSellerId(t *testing.T) {
	//arrange
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

	repository := MockRepository{
		Data: tp,
	}

	//server creation
	r := createServer(repository)

	//request creation
	request, recorder := createRequestTest(http.MethodGet, "/products?seller_id=FFFF", "")

	//act
	r.ServeHTTP(recorder, request)

	//asserts
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}
