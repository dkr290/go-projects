package documents

import (
	"fmt"
	"log"

	"rsc.io/pdf"
)

func ReadPDF(filePath string) {
	// Open the PDF file
	f, err := pdf.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the total number of pages
	numPages := f.NumPage()
	fmt.Println("Number of pages:", numPages)

	// Iterate over each page and print text
	for i := 1; i <= numPages; i++ {
		page := f.Page(i)
		// Extract text from the page
		fmt.Println(page.Content())
	}
}
