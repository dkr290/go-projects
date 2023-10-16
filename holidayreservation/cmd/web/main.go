package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dkr290/go-projects/holidayapplication/pkg/config"
	"github.com/dkr290/go-projects/holidayapplication/pkg/handlers"
	"github.com/dkr290/go-projects/holidayapplication/pkg/render"
)

var portNumber = ":8080"

func main() {

	var (
		UseCache = true
	)

	tc, err := render.CreateTemplateCache()
	if err != nil {

		log.Fatalln("error create template cache")

	}
	// app is of type appconfig so this way with the NEW fuction we have app as type appConfig and also passing tc so the cache Appconfig.TempLateCache = tc
	// passinf tc to NewConfig mean TempleteCache =  tc
	//Newtemplate gets app (Appconfig) and the global variable app which is AppConfig = app from here
	// we create also repo in the main with newrepo and appconfig and then newhandlers passing the repo back
	app := config.NewConfig(tc, UseCache) // second argument is to use cache or not
	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)
	render.NewTemplate(app)

	http.HandleFunc("/", handlers.Repo.HandleHome)
	http.HandleFunc("/about", handlers.Repo.HandleAbout)

	fmt.Printf("Starting the application on port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		panic(err)
	}

}
