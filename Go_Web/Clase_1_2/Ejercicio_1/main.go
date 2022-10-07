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

func searchTransaction(ctx *gin.Context) {
	transaction, err 
}

func main() {
	router := gin.Default()
	router.GET("/transactions", func(c *gin.Context) {
		GetAll(c)
	})
	router.Run()
}
