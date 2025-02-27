package documents

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/textsplitter"
)

func PdfLoader(f *os.File, size int64) error {
	// Get file size

	loader := documentloaders.NewPDF(f, size)
	// Load the document pages
	docs, err := loader.LoadAndSplit(context.Background(), textsplitter.RecursiveCharacter{
		ChunkSize:    1200,
		ChunkOverlap: 300,
	})
	if err != nil {
		return fmt.Errorf("error pdf load  %v", err)
	}
	fmt.Println(len(docs))
	return nil
}
