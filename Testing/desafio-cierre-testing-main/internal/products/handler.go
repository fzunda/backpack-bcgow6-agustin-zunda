package products

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		svc: s,
	}
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	sellerID := ctx.Query("seller_id")
	fmt.Println(sellerID)
	if sellerID == "" {
		ctx.JSON(400, gin.H{"error": "seller_id query param is required"})
		return
	}
	products, err := h.svc.GetAllBySeller(sellerID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, products)

}
