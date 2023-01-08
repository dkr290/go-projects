package routes

import (
	"github.com/dkr290/go-projects/bookstore-app/bookstore/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book{bookID}", controllers.UpdateBookByID).Methods("PUT")
	router.HandleFunc("/book{bookID}", controllers.DeleteBookByID).Methods("DELETE")

}
