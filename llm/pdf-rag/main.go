package main

import (
	"log/slog"
	"os"
	"pdf-rag/documents"
)

func main() {
	f, err := os.Open("./data/sample.pdf")
	if err != nil {
		slog.Error("Cannot open a file", "error", err)
	}
	defer f.Close()
	// Read the entire file into memory
	s, _ := f.Stat()
	size := s.Size()
	err = documents.PdfLoader(f, size)
	if err != nil {
		slog.Error("Erro pdf loader", "error", err)
	}
	// documents.ReadPDF("./data/sample.pdf")
}
