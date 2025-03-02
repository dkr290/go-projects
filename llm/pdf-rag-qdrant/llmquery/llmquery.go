package llmquery

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func QuestionResponse(ctx context.Context) error {
	// Load chat model (Ollama 3.2)
	chatModel, err := ollama.New(
		ollama.WithModel("ollama3.2"), // Change to your chat model
		ollama.WithServerURL("http://172.22.0.4/ollama"),
	)
	if err != nil {
		log.Fatal("Error loading chat model:", err)
	}

	// Load embedding model (nomic-embed-text)
	embedModel, err := ollama.New(
		ollama.WithModel("nomic-embed-text"),
		ollama.WithServerURL("http://172.22.0.4/ollama"),
	)
	if err != nil {
		log.Fatal("Error loading embedding model:", err)
	}
	// Format a prompt to direct the model what to do with the content and
	// the question.
	prompt := `You are the AI model that will answer the following question
	Question: %s
	`

	question := `What is About this Book?`

	finalPrompt := fmt.Sprintf(prompt, question)
	ollamaEmbeder, err := embeddings.NewEmbedder(embedModel)
	if err != nil {
		return fmt.Errorf("new embedder error %v", err)
	}
	// Create a new Qdrant vector store.
	url, err := url.Parse("http://qdrant.172.22.0.3.nip.io/")
	if err != nil {
		log.Fatal(err)
	}
	store, err := qdrant.New(
		qdrant.WithURL(*url),
		qdrant.WithCollectionName("vector_store"),
		qdrant.WithEmbedder(ollamaEmbeder),
	)
	if err != nil {
		return err
	}

	OptionVector := []vectorstores.Option{
		vectorstores.WithScoreThreshold(0.80),
	}
	retrreaver := vectorstores.ToRetriever(store, 10, OptionVector...)
	resDocs, err := retrreaver.GetRelevantDocuments(ctx, finalPrompt)
	if err != nil {
		return fmt.Errorf("failed to retreive documents %v", err)
	}

	fmt.Println(resDocs)

	// Search for similar documents using score threshold.
	docs, _ := store.SimilaritySearch(
		ctx,
		"american places",
		10,
		vectorstores.WithScoreThreshold(0.80),
	)
	fmt.Println(docs)
	return nil
}
