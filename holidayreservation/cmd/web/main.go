package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dkr290/go-projects/holidayapplication/pkg/config"
	"github.com/dkr290/go-projects/holidayapplication/pkg/render"
)

var (
	portNumber = ":8081"

	UseCache     = true
	InProduction = false
	app          *config.AppConfig
	session      *scs.SessionManager
)

func main() {

	// app is of type appconfig so this way with the NEW fuction we have app as type appConfig and also passing tc so the cache Appconfig.TempLateCache = tc
	// passinf tc to NewConfig mean TempleteCache =  tc
	//Newtemplate gets app (Appconfig) and the global variable app which is AppConfig = app from here
	// we create also repo in the main with newrepo and appconfig and then newhandlers passing the repo back

	///////////////////repository pattern
	// app := config.NewConfig(tc, UseCache) // second argument is to use cache or not
	// repo := handlers.NewRepo(app)
	// handlers.NewHandlers(repo)
	// render.NewTemplate(app)

	/////////////////////interfaces
	//this is going to routes
	// app = config.NewConfig(tc, UseCache) // second argument is to use cache or not
	// h := handlers.NewHandlers(app)
	// render.NewTemplate(app)

	// http.HandleFunc("/", h.HandleHome)
	// http.HandleFunc("/about", h.HandleAbout)
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	session.Cookie.Secure = InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {

		log.Fatalln("error create template cache")

	}

	app = config.NewConfig(tc, UseCache, InProduction, session) // second argument is to use cache or not

	fmt.Printf("Starting the application on port %s\n", portNumber)
	srv := &http.Server{

		Addr:    portNumber,
		Handler: routes(app),
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
