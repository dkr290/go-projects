package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

func HandleGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")

}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Danail", City: "Sofia", ZipCode: "100023432"},
		{Name: "Rob", City: "Paris", ZipCode: "32323232"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func HandleGetCustemer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["cust_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Post request received")

}
