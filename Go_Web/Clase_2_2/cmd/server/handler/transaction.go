package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/transactions"
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/pkg/web"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id              int       `json:"id"`
	Code            string    `json:"code"`
	Coin            string    `json:"coin"`
	Amount          float64   `json:"amount"`
	Emitting        string    `json:"emitting"`
	Receptor        string    `json:"Receptor"`
	TransactionDate time.Time `json:"transactionDate"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

// ListTransactions godoc
// @Summary  Show list transactions
// @Tags     Transactions
// @Produce  json
// @Param    token  header    string        true  "token"
// @Success  200    {object}  web.Response  "List transactions"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  500    {object}  web.Response  "Internal server error "
// @Failure  404    {object}  web.Response  "Not found products"
// @Router   /transactions [GET]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		res, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		if len(res) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "No hay transacciones almacenadas"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, res, ""))
	}
}

// Store Transaction godoc
// @Summary  Store transaction
// @Tags     Transactions
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "token requerido"
// @Param    transaction  body      request  true  "Transaction to Store"
// @Success  200      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  404      {object}  web.Response  "Not found products"
// @Failure  409      {object}  web.Response
// @Router   /transactions [POST]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}

		if req.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo es requerido"))
			return
		}

		if req.Coin == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "la moneda es requerido"))
			return
		}

		if req.Amount == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El monto es requerido"))
			return
		}

		if req.Emitting == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El emisor es requerido"))
			return
		}

		if req.Receptor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El receptor es requerido"))
			return
		}

		c, err := t.service.Store(req.Id, req.Code, req.Coin, req.Amount, req.Emitting, req.Receptor, req.TransactionDate)
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, c, ""))
	}
}

// UpdateTransactions godoc
// @Summary  Update transaction
// @Tags     Transactions
// @Accept   json
// @Produce  json
// @Param    Id       path      int             true   "Id transactions"
// @Param    token    header    string          false  "token"
// @Param    transactions  body      request  true   "transaction to update"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  404      {object}  web.Response
// @Router   /transactions/{Id} [PUT]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("Id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El id es requerido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if req.Code == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El codigo es requerido"))
			return
		}
		if req.Coin == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La moneda es requerido"))
			return
		}
		if req.Amount <= 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El monto es requerido"))
			return
		}
		if req.Emitting == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El emisor es requerido"))
			return
		}
		if req.Receptor == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El receptor es requerido"))
			return
		}
		t, err := t.service.Update(id, req.Code, req.Coin, req.Amount, req.Emitting, req.Receptor, req.TransactionDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))
	}
}

// Update Code and Amount Transactions godoc
// @Summary      Update code and amount transactions
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Description  This endpoint update field code and amount from an transaction
// @Param        token  header    string               true  "Token header"
// @Param        Id     path      int                  true  "Transaction Id"
// @Param        transactions  body   request          true   "transaction to update"
// @Success      200    {object}  web.Response
// @Failure      401    {object}  web.Response
// @Failure      400    {object}  web.Response
// @Failure      404    {object}  web.Response
// @Router       /transactions/{id} [PATCH]
func (t *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("Id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El id es requerido"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if req.Code == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El codigo es requerido"))
			return
		}
		if req.Amount <= 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El monto es requerido"))
			return
		}
		t, err := t.service.UpdateCodeAndAmount(id, req.Code, req.Amount)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))
	}
}

// Delete Transaction
// @Summary  Delete transaction
// @Tags     Transactions
// @Param    id     path      int     true  "Transaction id"
// @Param    token  header    string  true  "Token"
// @Success  200    {object}  web.Response
// @Failure  401    {object}  web.Response
// @Failure  400    {object}  web.Response
// @Failure  404    {object}  web.Response
// @Router   /transactions/{id} [DELETE]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		id, err := strconv.Atoi(ctx.Param("Id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El id es requerido"))
			return
		}
		if err := t.service.Delete(id); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, "deleted transaction", ""))
	}
}
