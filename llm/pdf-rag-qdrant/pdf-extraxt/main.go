package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Open the PDF file
	filePath := "./BOI.pdf"
	output, err := exec.Command("pdftotext", filePath, "-").Output()
	if err != nil {
		log.Fatalf("Failed to extract text: %v", err)
	}

	fmt.Printf("Extracted text:\n%s\n", output)
}
