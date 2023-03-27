package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/web3/models"
	"github.com/dkr290/go-projects/web3/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.PageData{})

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var strMap = map[string]string{
		"title": "About",
		"intro": "This page is where we talk about ourselves.",
	}

	render.RenderTemplate(w, "about.page.html", &models.PageData{
		StrMap: strMap,
	})

}
