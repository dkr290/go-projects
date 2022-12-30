package repository

import "database/sql"

type Repository struct {
	DB      *sql.DB
	DbUser  string
	DbPass  string
	DbPort  string
	DbName  string
	DbHost  string
	APPPort string
}

// creates new repository
func NewRepo() *Repository {
	r := Repository{}
	return &r
}
