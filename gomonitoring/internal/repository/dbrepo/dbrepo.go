package dbrepo

import (
	"database/sql"

	"github.com/dkr290/go-projects/gomonitoring/internal/config"
	"github.com/dkr290/go-projects/gomonitoring/internal/repository"
)

var app *config.AppConfig

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo creates the repository
func NewPostgresRepo(Conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	app = a
	return &postgresDBRepo{
		App: a,
		DB:  Conn,
	}
}
