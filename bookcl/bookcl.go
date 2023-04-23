package main

import (
	"encoding/json"
	"os"
)

// BookCl is a collection of books on a shelf
type BookReader struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book descrfibes the book in a shelf
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookCl(filepath string) ([]BookReader, error) {

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookreader []BookReader
	//decode the json and store it to the variable bookscls
	err = json.NewDecoder(f).Decode(&bookreader)
	if err != nil {
		return nil, err
	}

	return bookreader, nil
}

func findCommonBooks(bookreaders []BookReader) []Book {

}
