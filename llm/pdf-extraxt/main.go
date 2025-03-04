package main

import (
	"fmt"
	"log/slog"
	"pdf-extract/pkg/pdf"
)

func main() {
	p := pdf.New("./BOI.pdf")
	if err := p.ReadPdf(); err != nil {
		slog.Error("error reading pdf ", "error", err)
	}
	docs, err := p.SplitText()
	if err != nil {
		slog.Error("error splitting the text", "error", err)
	}

	// for _, doc := range docs {
	// 	fmt.Println("")
	// 	log.Println(doc)
	// }

	mdTextChunks, err := p.AddMetadata(docs, "BOI US FinCEN")
	if err != nil {
		slog.Error("error metadata ", "error", err)
	}
	fmt.Println(mdTextChunks)
}
