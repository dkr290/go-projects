package main

import (
	"blob-uploader/handlers"
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//initialize gin
	storageName := os.Getenv("STORAGE_NAME")

	if len(storageName) == 0 {
		log.Println("Missing variable STORAGE_NAME")
		return
	}
	containerName := os.Getenv("CONTAINER_NAME")
	if len(containerName) == 0 {
		log.Println("Missing Variable CONTAINER_NAME")
		return
	}

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*.html")))

	hand := handlers.NewHandlers(r, storageName, containerName)

	r.POST("/upload", hand.UploadHandler)
	r.GET("/", hand.GetHandler)
	r.Run(":8080")

}
