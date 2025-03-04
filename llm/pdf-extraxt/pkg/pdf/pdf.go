package pdf

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/tmc/langchaingo/textsplitter"
)

// PDF is a struct that represents a PDF document.
type PDF struct {
	Content  string
	Filepath string
}

func New(p string) *PDF {
	return &PDF{
		Filepath: p,
	}
}

func (p *PDF) ReadPdf() error {
	// Open the PDF file
	output, err := exec.Command("pdftotext", p.Filepath, "-").Output()
	if err != nil {
		return fmt.Errorf("failed to extract text: %v", err)
	}

	p.Content = string(output)
	return nil
}

func (p *PDF) SplitText() ([]string, error) {
	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(1200),
		textsplitter.WithChunkOverlap(200),
	)
	// Use the splitter to split the text into chunks
	chunks, err := splitter.SplitText(p.Content)
	if err != nil {
		return nil, fmt.Errorf("error splitting the text %v", err)
	}

	return chunks, nil
}

func (p *PDF) AddMetadata(chunks []string, docTitle string) ([]map[string]any, error) {
	type metadata struct {
		title  string
		author string
		date   string
	}

	var metadataChunks []map[string]interface{}

	// Loop through each chunk
	for _, c := range chunks {
		// Create a metadata struct for the current chunk
		md := metadata{
			title:  docTitle,
			author: "US Business Bureau",
			date:   time.Now().Format("2006-01-02 15:04:05"), // Format the time
		}

		// Create a map for the current chunk and its metadata
		chunkMap := map[string]interface{}{
			"text":     c,
			"metadata": md,
		}

		// Append the map to the slice
		metadataChunks = append(metadataChunks, chunkMap)
	}

	return metadataChunks, nil
}
