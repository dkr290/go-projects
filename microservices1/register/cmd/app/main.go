package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"register/data"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Printf("Starting register service at port %s", webPort)

	//connect to Database

	conn := connectToDb()
	if conn == nil {
		log.Panic("Can't connect to the postgres")
	}

	app := Config{

		DB:     conn,
		Models: data.New(conn),
	}

	// going to the chi router in routes and then to each handler
	srv := &http.Server{

		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func connectToDb() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress server is not yet ready")
			counts++
		} else {
			log.Println("Connected to postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue

	}
}
