package app

import (
	"bookstore_users-api/handlers/ping"
	"bookstore_users-api/handlers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.FindUser)

}
