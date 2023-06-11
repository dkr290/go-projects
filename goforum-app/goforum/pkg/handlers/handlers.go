package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.html", &models.PageData{})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About us"
	strMap["intro"] = "This page is where we talk about ourselves this is like that"

	render.RenderTemplate(w, "about.html", &models.PageData{
		StrMap: strMap,
	})
}
