package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/about", HandleAbout)

	http.ListenAndServe(":8080", nil)

}

// this is the about functions
func HandleAbout(w http.ResponseWriter, r *http.Request) {

	owner, saying := getData()

	sum := addValues(2, 2)

	fmt.Fprintf(w, "This is the about page of %s\nI like to say %s\nand as a side note, 2 + 2 is %d", owner, saying, sum)
}

// this is the about page
func HandleHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is a home page")

}

func getData() (string, string) {
	o := "Rick Sanchez"
	s := "Wubba Luba Dup Dup"

	return o, s
}

func addValues(x, y int) int {

	return x + y

}
