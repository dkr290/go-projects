package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()

	router := mux.NewRouter()

	router.HandleFunc("/greet", HandleGreet).Methods(http.MethodGet)
	router.HandleFunc("/customers", GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{cust_id:[0-9]+}", HandleGetCustemer)

	if err := http.ListenAndServe(":8080", router); err != nil {

		log.Fatal(err)

	}
}
