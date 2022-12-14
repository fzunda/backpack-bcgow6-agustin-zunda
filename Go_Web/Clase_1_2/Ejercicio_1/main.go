package main

/*
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder
filtrar por todos los campos.
	1. Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
	2. Luego genera la lógica de filtrado de nuestro array.
	3. Devolver por el endpoint el array filtrado.

*/

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Go_Web/Clase_2_2/pkg/web"
	"github.com/gin-gonic/gin"
)

type transaction struct {
	Id              int
	Code            string
	Coin            string
	Amount          float64
	Emitting        string
	Receptor        string
	TransactionDate time.Time
}

func GetAll(ctx *gin.Context) {
	jsonData, err := os.ReadFile("./transactions.json")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
	} else {
		var arrayTransactions []transaction
		err := json.Unmarshal([]byte(jsonData), &arrayTransactions)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": arrayTransactions})

		}
	}
}

func searchTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jsonData, err := os.ReadFile("./transactions.json")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		} else {
			var arrayTransactions []transaction
			err := json.Unmarshal([]byte(jsonData), &arrayTransactions)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
			} else {
				var req transaction
				if err := ctx.ShouldBindJSON(&req); err != nil {
					ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
					return
				}

				var result transaction

				for _, val := range arrayTransactions {
					if req.Id != 0 && val.Id == req.Id {
						result = val
					}
					if req.Id != 0 && val.Id == req.Id && req.Code != "" && val.Code == req.Code {
						result = val
					}
					if req.Id != 0 && val.Id == req.Id && req.Code != "" && val.Code == req.Code {
						result = val
					}
					if req.Id != 0 && val.Id == req.Id && req.Code != "" && val.Code == req.Code && req.Coin != "" && val.Coin == req.Coin {
						result = val
					}
					if req.Id != 0 && val.Id == req.Id && req.Code != "" && val.Code == req.Code && req.Coin != "" && val.Coin == req.Coin && req.Amount != 0 && val.Amount == req.Amount {
						result = val
					}
					if req.Id != 0 && val.Id == req.Id && req.Code != "" && val.Code == req.Code && req.Coin != "" && val.Coin == req.Coin && req.Amount != 0 && val.Amount == req.Amount && req.Emitting != "" && val.Emitting == req.Emitting {
						result = val
					}
					if req.Id != 0 && val.Id == req.Id && req.Code != "" && val.Code == req.Code && req.Coin != "" && val.Coin == req.Coin && req.Amount != 0 && val.Amount == req.Amount && req.Emitting != "" && val.Emitting == req.Emitting && req.Receptor != "" && val.Receptor == req.Receptor {
						result = val
					}
				}
				ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, result, ""))
				return
			}
		}

	}
}

func main() {
	router := gin.Default()
	router.GET("/transactions", func(c *gin.Context) {
		GetAll(c)
	})
	router.GET("/transactions/search", searchTransaction())
	router.Run()
}
