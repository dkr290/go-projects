package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dkr290/go-projects/bookstore-app/bookstore/models"
	"github.com/dkr290/go-projects/bookstore-app/bookstore/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {

	nb := models.GetAllBooks()
	res, err := json.Marshal(nb)
	if err != nil {
		log.Fatal("Error marshal in GetBook ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatalln("error while parsing", err)
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		log.Fatalln("Error marshaling the Users in CreateBook", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {

}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {

}
