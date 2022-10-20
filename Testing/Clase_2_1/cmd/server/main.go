package main

import (
	"fmt"
	"os"

	"Clase_2_1/cmd/server/handler"
	"Clase_2_1/docs"
	"Clase_2_1/internal/transactions"
	"Clase_2_1/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI BootCamp API
// @version 1.0
// @description this API handler MELI transactions
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	_ = godotenv.Load()
	db := store.New(store.FileType, "./transactions.json")
	repo := transactions.NewRepository(db)
	ser := transactions.NewService(repo)

	t := handler.NewTransaction(ser)

	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	tr.PUT("/:Id", t.Update())
	tr.PATCH("/:Id", t.UpdateCodeAndAmount())
	tr.DELETE("/:Id", t.Delete())

	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
