package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	port   = "8080"
)

func StartApplication() {

	mapURLs()
	fmt.Printf("Starting web server, listening on %s\n", port)
	router.Run(fmt.Sprintf(":%s", port))

}
