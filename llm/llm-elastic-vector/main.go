package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type VectorIndexSettings struct {
	NumDimensions int
	Path          string
	Similarity    string
}

func createVectorIndex(
	client *opensearch.Client,
	indexName string,
	settings VectorIndexSettings,
) error {
	// Define the index mapping for vector search
	mapping := map[string]interface{}{
		"settings": map[string]interface{}{
			"index": map[string]interface{}{
				"knn": true, // Enable k-NN search
			},
		},
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				settings.Path: map[string]interface{}{
					"type":      "knn_vector", // OpenSearch uses "knn_vector" for vector fields
					"dimension": settings.NumDimensions,
					"method": map[string]interface{}{
						"name":       "hnsw",
						"space_type": settings.Similarity,
						"engine":     "nmslib",
					},
				},
			},
		},
	}

	// Convert the mapping to JSON
	mappingJSON, err := json.Marshal(mapping)
	if err != nil {
		return fmt.Errorf("failed to marshal mapping: %w", err)
	}

	// Create the index
	req := opensearchapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(string(mappingJSON)),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index: %s", res.String())
	}

	log.Printf("Index '%s' created successfully", indexName)
	return nil
}

func main() {
	// Initialize OpenSearch client
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{"http://api.172.22.0.3.nip.io"}, // OpenSearch server address
		Username:  "admin",
		Password:  "admin",
	})
	if err != nil {
		log.Fatalf("Error creating OpenSearch client: %s", err)
	}

	// Define vector index settings
	settings := VectorIndexSettings{
		NumDimensions: 4,
		Path:          "embedding",
		Similarity:    "cosinesimil", // OpenSearch uses "cosinesimil" for cosine similarity
	}

	// Create the vector index
	indexName := "vector_index"
	err = createVectorIndex(client, indexName, settings)
	if err != nil {
		log.Fatalf("Failed to create vector index: %s", err)
	}

	log.Println("Vector index created successfully")
}
