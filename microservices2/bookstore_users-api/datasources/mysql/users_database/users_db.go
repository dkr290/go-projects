package usersdatabase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "pass1234", "mysql", "users")
	var err error
	db, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
