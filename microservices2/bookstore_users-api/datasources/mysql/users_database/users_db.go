package usersdatabase

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	dbClient       *sql.DB
	mysql_username string
	mysql_password string
	mysql_host     string
	mysql_chema    string
)

func New() *sql.DB {
	return dbClient
}

func init() {
	var err error
	var counts int
	getEnvironment()

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mysql_username, mysql_password, mysql_host, mysql_chema)

	for {
		dbClient, err = sql.Open("mysql", datasourceName)
		if err != nil {
			log.Println("Mysql server is not yet ready")
			counts++
		} else {
			log.Println("Connected to mysql")
			break

		}

		if counts > 30 {
			panic(err)

		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue

	}

	if err = dbClient.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}

func getEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mysql_username = os.Getenv("MYSQL_USERNAME")
	if len(mysql_username) == 0 {
		log.Fatal("You must set your 'MYSQL_USERNAME' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	mysql_password = os.Getenv("MYSQL_PASSWORD")
	if len(mysql_password) == 0 {
		log.Fatal("You must set your 'MYSQL_PASSWORD' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	mysql_host = os.Getenv("MYSQL_HOST")
	if len(mysql_host) == 0 {
		log.Fatal("You must set your 'MYSQL_HOST' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	mysql_chema = os.Getenv("MYSQL_SCHEMA")
	if len(mysql_chema) == 0 {
		log.Fatal("You must set your 'MYSQL_SCHEMA' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
}
