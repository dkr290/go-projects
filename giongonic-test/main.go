package main

import (
	"github.com/dkr290/go-projects/giongonic-test/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	hand := handlers.NewHandler()
	router := gin.Default()
	router.GET("/:name", hand.IndexHandler)
	router.GET("/xml", hand.XMLHandler)
	router.Run("127.0.0.1:8080")

}
