package main

import (
	"log/slog"
	"pdf-extract/pkg/pdf"
)

func main() {
	// new PDF reading
	p := pdf.New("./BOI.pdf")
	if err := p.ReadPdf(); err != nil {
		slog.Error("error reading pdf ", "error", err)
	}
	// split text in chunks
	docs, err := p.SplitText()
	if err != nil {
		slog.Error("error splitting the text", "error", err)
	}

	// for _, doc := range docs {
	// 	fmt.Println("")
	// 	log.Println(doc)
	// }
	// add metadata to the chunks
	mdTextChunks, err := p.AddMetadata(docs, "BOI US FinCEN")
	if err != nil {
		slog.Error("error metadata ", "error", err)
	}

	err = p.GenEmbeddings(
		mdTextChunks,
		"nomic-embed-text",
		"http://172.22.0.3/ollama",
		"http://qdrant.172.22.0.4.nip.io",
		"vector_store1",
	)
	if err != nil {
		slog.Error("error adding the embeddings", "error", err)
	}
}
