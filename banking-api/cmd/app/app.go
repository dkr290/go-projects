package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dkr290/go-projects/banking-api/domain"
	"github.com/dkr290/go-projects/banking-api/pkg/logger"
	"github.com/dkr290/go-projects/banking-api/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepoDb(dbClient)
	//accountRepositoryDb := domain.NewAccountRepoDb(dbClient)

	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepoStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}

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

func getDbClient() *sqlx.DB {
	db_User := os.Getenv("DB_USER")
	db_Pass := os.Getenv("DB_PASS")
	db_Addr := os.Getenv("DB_ADDR")
	db_Port := os.Getenv("DB_PORT")
	db_Name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_User, db_Pass, db_Addr, db_Port, db_Name)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	err = testDb(client)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
func testDb(client *sqlx.DB) error {
	counts := 0

	for {
		err := client.Ping()
		if err != nil {
			logger.Error("Mysql server is not yet ready")
			counts++
		} else {
			logger.Info("*** Pinged database successfully! ***")
			return nil
		}
		if counts > 10 {
			logger.Error("Error connection to the database" + err.Error())
			return err
		}

		logger.Info("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
