package app

import "github.com/dkr290/go-projects/secure-notes/handlers"

func mapURLs() {
	router.GET("/ping", handlers.Ping)
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlers.GetNotes)
	//router.POST("/", handlers.PostNotes)

}
