package main

import (
	"html/template"
	"test-sp-monitor/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	hand := handlers.NewHandlers(r)
	// Load HTML template
	r.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")))
	r.GET("/", hand.GetHandler)
	r.POST("/add", hand.AddHandler)
	r.POST("/delete", hand.DeleteHandler)

	// Run the server
	r.Run(":8080")

}
