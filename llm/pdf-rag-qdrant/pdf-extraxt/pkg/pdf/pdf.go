package pdf

import (
	"fmt"
	"os/exec"

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
