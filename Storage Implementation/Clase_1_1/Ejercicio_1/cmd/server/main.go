package main

import (
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/cmd/server/handler"
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/db"
	"github.com/fzunda/backpack-bcgow6-agustin-zunda/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := db.Connection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/:name", p.GetByName())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
