package main

import (
	"context"
	"log/slog"
	"pdf-extract/pkg/pdf"

	"github.com/qdrant/go-client/qdrant"
)

func main() {
	// new PDF reading
	p := pdf.New("./BOI.pdf")
	if err := p.ReadPdf(); err != nil {
		slog.Error("error reading pdf ", "error", err)
		return
	}
	// split text in chunks
	docs, err := p.SplitText()
	if err != nil {
		slog.Error("error splitting the text", "error", err)
		return
	}

	// for _, doc := range docs {
	// 	fmt.Println("")
	// 	log.Println(doc)
	// }
	// add metadata to the chunks
	mdTextChunks, err := p.AddMetadata(docs, "BOI US FinCEN")
	if err != nil {
		slog.Error("error metadata ", "error", err)
		return
	}

	// err = createCollection("vector_store1", "qdrant.172.22.0.5.nip.io/grpc")
	// if err != nil {
	// 	slog.Error("error creating collection", "error", err)
	// 	return
	// }

	err = p.GenEmbeddings(
		mdTextChunks,
		"nomic-embed-text",
		"http://172.22.0.3/ollama",
		"http://qdrant.172.22.0.5.nip.io",
		"vector_store1",
	)
	if err != nil {
		slog.Error("error adding the embeddings", "error", err)
	}
}

func uint32Ptr(i uint32) *uint32 { return &i }

func createCollection(coll string, host string) error {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   host,
		Port:   80,
		UseTLS: false,
	})
	if err != nil {
		return err
	}
	err = client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName:    coll,
		ShardNumber:       uint32Ptr(6),
		ReplicationFactor: uint32Ptr(3),
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     768,
			Distance: qdrant.Distance_Cosine,
		}),
	})
	if err != nil {
		return err
	}
	return nil
}

// curl -X PUT "http://localhost:6333/collections/my_collection" \
//      -H "Content-Type: application/json" \
//      -d '{
//            "vectors": {
//                "size": 128,
//                "distance": "Cosine"
//            },
//            "shard_number": 6,
//            "replication_factor": 3
//          }'
