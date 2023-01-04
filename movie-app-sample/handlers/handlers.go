package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dkr290/go-projects/movie-app-sample/models"
	"github.com/gorilla/mux"
)

var movies = models.CreateMovies()

func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovie() {

}

func CreateMovie() {

}

func UpdateMovie() {

}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append()
		}

	}

}
