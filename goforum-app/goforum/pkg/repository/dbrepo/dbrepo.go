package dbrepo

import (
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/repository"
	"github.com/jackc/pgx/v5"
)

type PostgresDBRepo struct {
	App    *config.AppConfig
	DBConn *pgx.Conn
}

func NewPostgresRepo(conn *pgx.Conn, app *config.AppConfig) repository.DatabaseRepo {
	return &PostgresDBRepo{
		App:    app,
		DBConn: conn,
	}
}
