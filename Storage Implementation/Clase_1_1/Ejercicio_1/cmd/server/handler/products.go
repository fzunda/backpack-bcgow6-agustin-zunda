package handler

import (
	"net/http"
	"strconv"

	"github.com/fzunda/backpack-bcgow6-agustin-zunda/cmd/server/handler/request"
	domian "github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/domains"
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/products"
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (s *Product) GetByName() gin.HandlerFunc {

	return func(c *gin.Context) {
		name := (c.Param("name"))
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		}
		prod, err := s.service.GetByName(name)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(prod, "", http.StatusOK))
	}
}

func (s *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request.ProductRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, "Name is required", http.StatusBadRequest))
			return
		}
		if req.Type == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, "Type is required", http.StatusBadRequest))
			return
		}
		if req.Count == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, "Count is required", http.StatusBadRequest))
			return
		}
		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, "Price is required", http.StatusBadRequest))
			return
		}

		result, err := s.service.Store(domian.Product(req))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(result, "", http.StatusOK))
	}
}

func (s *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}
		if len(res) == 0 {
			ctx.JSON(404, web.NewResponse(nil, "no products stored", http.StatusNotFound))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(res, "", http.StatusOK))
	}
}

func (s *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domian.Product
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, "id is required", http.StatusBadRequest))
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(nil, err.Error(), http.StatusNotFound))
		}
		p, err := s.service.Update(id, req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(nil, err.Error(), http.StatusNotFound))
		}
		ctx.JSON(http.StatusOK, web.NewResponse(p, "", http.StatusOK))
	}
}

func (s *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, "id is required", http.StatusBadRequest))
			return
		}

		err = s.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(nil, "", http.StatusNoContent))
	}
}
