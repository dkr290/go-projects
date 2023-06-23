package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dkr290/go-projects/goforum-app/goforum/models"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/dbdriver"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/handlers"
	"github.com/jackc/pgx/v5"
)

const webPort = "8080"

var sm *scs.SessionManager
var app config.AppConfig

const connString = "postgres://postgres:password@postgres:5432/blog_db"

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routes(&app),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func run() (*pgx.Conn, error) {
	gob.Register(models.Article{})

	// add table models in session

	gob.Register(models.User{})
	gob.Register(models.Post{})

	sm = scs.New()
	sm.Lifetime = 24 * time.Hour
	sm.Cookie.Persist = true
	sm.Cookie.Secure = false
	sm.Cookie.SameSite = http.SameSiteLaxMode
	//save in the config
	app.Session = sm

	dbconn := dbdriver.ConnectDatabase(connString)

	repo := handlers.NewRepo(&app, dbconn)
	handlers.NewHandlers(repo)
	return dbconn, nil
}
