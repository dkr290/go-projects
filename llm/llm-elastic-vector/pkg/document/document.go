package document

import (
	"context"

	"github.com/opensearch-project/opensearch-go"
)

type Document struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Embedding []float64 `json:"embedding"`
}

type SearchResult struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Embedding []float64 `json:"embedding"`
	Score     float64   `json:"score"`
}

func StorDocument(cts context.Context, client *opensearch.Client, indexName string) error {
	return nil
}
