package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/holidayapplication/pkg/render"
)

type Data struct {
	Title string
}

// this is the about functions
func HandleAbout(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Title: "About Page",
	}

	render.RenderTemplate(w, "about-page.html", data)

}

// this is the about page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Title: "Home Page",
	}
	render.RenderTemplate(w, "home-page.html", data)
}
