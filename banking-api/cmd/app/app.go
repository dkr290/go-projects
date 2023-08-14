package app

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/banking-api/domain"
	"github.com/dkr290/go-projects/banking-api/service"
	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()

	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepoStub())}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8080", router); err != nil {

		log.Fatal(err)

	}
}
