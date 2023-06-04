package app

import (
	"bookstore_users-api/handlers"
)

func mapURLs() {
	router.GET("/ping", handlers.Ping)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:user_id", handlers.GetUser)
	router.PUT("/users/:user_id", handlers.UpdateUser)
	router.GET("/users/search", handlers.FindUser)

}
