package main

import (
	"errors"
	"fmt"
	"net/http"
)

var portNumber = ":8080"

func main() {

	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/about", HandleAbout)
	http.HandleFunc("/devide", HandleDevide)

	fmt.Printf("Starting the application on port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		panic(err)
	}

}

// this is the about functions
func HandleAbout(w http.ResponseWriter, r *http.Request) {

	owner, saying := getData()

	sum := addValues(2, 2)

	_, _ = fmt.Fprintf(w, "This is the about page of %s\nI like to say %s\nand as a side note, 2 + 2 is %d", owner, saying, sum)
}

// this is the about page
func HandleHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is a home page")

}

func HandleDevide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 0.0

	f, err := devideValues(x, y)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error: Division by 0 is not valid operation. Error %s", err)
		return
	}

	_, _ = fmt.Fprintf(w, "%f devided by %f is %f", x, y, f)

}

func getData() (string, string) {
	o := "Rick Sanchez"
	s := "Wubba Luba Dup Dup"

	return o, s
}

func addValues(x, y int) int {

	return x + y

}

func devideValues(x, y float64) (float64, error) {

	if y == 0 {
		err := errors.New("devisor is 0")
		return 0, err
	}
	result := x / y
	return result, nil
}
