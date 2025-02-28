package main

import (
	"context"
	"fmt"
	"log/slog"
	"pdf-rag/documents"
	"time"

	ops "github.com/opensearch-project/opensearch-go"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/opensearch"
)

func main() {
	ctx := context.Background()
	chunks, err := documents.PdfLoader("./data/how-to-code-in-go.pdf", ctx)
	if err != nil {
		slog.Error("erro chunking the file", "error", err)
	}
	// Initialize OpenSearch client
	client, err := ops.NewClient(ops.Config{
		Addresses: []string{"http://api.172.22.0.3.nip.io"}, // OpenSearch server address
		Username:  "admin",
		Password:  "admin",
	})
	if err != nil {
		slog.Error("Error creating OpenSearch client: %s", "error", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = vector_embed(chunks, ctx, client)
	if err != nil {
		slog.Error("vector_embed error", "error", err)
	}
}

func vector_embed(chunks []schema.Document, ctx context.Context, client *ops.Client) error {
	ollamaLLM, err := ollama.New(ollama.WithModel("nomic-embed-text"))
	if err != nil {
		return fmt.Errorf("cannot load ollama model %v", err)
	}
	ollamaEmbeder, err := embeddings.NewEmbedder(ollamaLLM)
	if err != nil {
		return fmt.Errorf("new embedder error %v", err)
	}

	store, err := opensearch.New(client,
		opensearch.WithEmbedder(ollamaEmbeder),
	)
	if err != nil {
		return err
	}
	_, err = store.AddDocuments(ctx, chunks)
	if err != nil {
		return err
	}
	return nil
}
