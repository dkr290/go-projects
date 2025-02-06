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

func loadBookReades(filepath string) ([]BookReader, error) {

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
	booksOnShelves := booksCount(bookreaders)

	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return commonBooks
}

// bookCount registers all the books and their occurances from the bookreaders shelves.

func booksCount(bookreaders []BookReader) map[Book]uint {

	count := make(map[Book]uint)

	for _, bookreader := range bookreaders {
		for _, book := range bookreader.Books {
			count[book]++
		}
	}

	return count
}
