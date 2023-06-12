package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/handlers"
)

const webPort = "8080"

var sm *scs.SessionManager
var app config.AppConfig

func main() {

	sm = scs.New()
	sm.Lifetime = 24 * time.Hour
	sm.Cookie.Persist = true
	sm.Cookie.Secure = false
	sm.Cookie.SameSite = http.SameSiteLaxMode
	//save in the config
	app.Session = sm

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routes(&app),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
