package main

/*
1) Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
	1. Crea un endpoint mediante POST el cual reciba la entidad.
	2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
	deberán ir guardando todas las peticiones que se vayan realizando.
	3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).

2) Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).

3) Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”.
*/

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	Id              int       `json:"id"`
	Code            string    `json:"code" binding:"required"`
	Coin            string    `json:"coin" binding:"required"`
	Amount          float64   `json:"amount" binding:"required"`
	Emitting        string    `json:"emitting" binding:"required"`
	Receptor        string    `json:"Receptor" binding:"required"`
	TransactionDate time.Time `json:"transactionDate"`
}

// array de transacciones
var requestTransactions []transaction

var transaction1 = transaction{Id: 1, Code: "123465A", Coin: "Dolar", Amount: 299, Emitting: "Central Bank", Receptor: "Bank of Córdoba", TransactionDate: time.Now()}
var transaction2 = transaction{Id: 2, Code: "123465B", Coin: "Euro", Amount: 315, Emitting: "Central Bank", Receptor: "Santander Río", TransactionDate: time.Now()}
var transaction3 = transaction{Id: 3, Code: "123465C", Coin: "Peso Chileno", Amount: 152, Emitting: "Central Bank", Receptor: "Galicia", TransactionDate: time.Now()}

func Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req transaction
		token := ctx.GetHeader("token")
		if token != "1234" {
			ctx.JSON(401, gin.H{"error": "token invalido"})
			return
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			lastId := len(requestTransactions)
			lastId++
			req.Id = lastId
			req.TransactionDate = time.Now()
			requestTransactions = append(requestTransactions, req)
			ctx.JSON(http.StatusOK, req)
		}
	}
}

func main() {
	requestTransactions = append(requestTransactions, transaction1)
	requestTransactions = append(requestTransactions, transaction2)
	requestTransactions = append(requestTransactions, transaction3)

	r := gin.Default()
	pr := r.Group("/transactions")
	pr.POST("/add", Guardar())
	r.Run()
}
