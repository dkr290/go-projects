package main

import (
	"fmt"
	"os"
)

func main() {

	bookscl, err := loadBookReades("testdata/bookcl.json")

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load bookcl %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookscl)
}
