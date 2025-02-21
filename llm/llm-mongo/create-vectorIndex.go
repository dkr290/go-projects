package main

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const indexName = "vector_index"

// VectorIndexSettings represents setting to create a vector index.
type VectorIndexSettings struct {
	NumDimensions int
	Path          string
	Similarity    string
}

func CreateVectorIndex(client *mongo.Client) error {
	settings := VectorIndexSettings{
		NumDimensions: 4,
		Path:          "embedding",
		Similarity:    "cosine",
	}

	return nil
}
