package main

/*
	1. Crea dentro de la carpeta go-web un archivo llamado main.go
	2. Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
	3. Pegale al endpoint para corroborar que la respuesta sea la correcta.
*/

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Agus"})
	})
	router.Run()
}

//http://localhost:8080/hello
