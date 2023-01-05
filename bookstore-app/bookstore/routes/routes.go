package routes

import "github.com/gorilla/mux"

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book{bookID}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book{bookID}", controllers.UpdateBookByID).Methods("PUT")
	router.HandleFunc("/book{bookID}", controllers.DeleteBookByID).Methods("DELETE")

}
