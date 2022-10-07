package main

/*
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
Genera un handler para el endpoint llamado “GetAll”.
Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
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

func main() {
	router := gin.Default()
	router.GET("/transactions", func(c *gin.Context) {
		GetAll(c)
	})
	router.Run()
}
