package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dkr290/go-projects/banking-api/domain"
	"github.com/dkr290/go-projects/banking-api/pkg/logger"
	"github.com/dkr290/go-projects/banking-api/service"
	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	var (
		sAddress string
		sPort    string
	)

	v := envsCheck()

	if v {

		sAddress = os.Getenv("SERVER_ADDRESS")
		sPort = os.Getenv("SERVER_PORT")
		if sAddress == "" {
			sAddress = ""
		}
		if sPort == "" {
			sPort = "4000"
		}

	} else {
		sAddress = os.Getenv("SERVER_ADDRESS")
		sPort = os.Getenv("SERVER_PORT")
	}

	router := mux.NewRouter()

	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepoStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepoDb())}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	logger.Info("listen on ..." + sAddress + sPort)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", sAddress, sPort), router); err != nil {

		log.Fatal(err)

	}
}

func envsCheck() bool {

	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		return true
	}

	if len(os.Getenv("DB_USER")) == 0 {
		log.Fatal("You must set your 'DB_USER' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(os.Getenv("DB_PASS")) == 0 {
		log.Fatal("You must set your 'DB_PASS' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(os.Getenv("DB_ADDR")) == 0 {
		log.Fatal("You must set your 'DB_ADDR' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(os.Getenv("DB_PORT")) == 0 {
		log.Fatal("You must set your 'DB_PORT' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(os.Getenv("DB_NAME")) == 0 {
		log.Fatal("You must set your 'DB_NAME' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	return false
}
