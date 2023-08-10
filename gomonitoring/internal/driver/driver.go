package driver

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn" // need this and next two for pgx
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection information
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 25
const maxIdleDbConn = 25
const maxDbLifetime = 5 * time.Minute

// ConnectPostgres creates database pool for postgres
func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)
	dbConn.SQL = d

	err = testDB(d)

	return dbConn, err
}

// testDB pings database

func testDB(d *sql.DB) error {
	counts := 0
	for {

		err := d.Ping()
		if err != nil {
			log.Println("Postgress server is not yet ready")
			counts++
		} else {
			log.Println("*** Pinged database successfully! ***")
			return nil
		}

		if counts > 10 {
			fmt.Println("Error!", err)
			return err
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue

	}

}
