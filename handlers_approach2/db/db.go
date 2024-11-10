package db

import (
	"database/sql"
	"log"
)

type Database interface {
	GetDB() error
}

type Mysql struct {
	DB *sql.DB
}

func (m *Mysql) GetDB() error {
	log.Println("The database simulated SELECT * FROM")
	return nil
}
