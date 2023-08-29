package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dkr290/go-projects/banking-api/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(params)

	if err != nil {
		writeResponse(w, err.Code, err.ReturnMessage())

	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {

	cust := mux.Vars(r)
	id := cust["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {

		writeResponse(w, err.Code, err.ReturnMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)

	}
}

func HandleGetCustemer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["cust_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Post request received")

}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal("Error encoding to json", err.Error())
	}
}
