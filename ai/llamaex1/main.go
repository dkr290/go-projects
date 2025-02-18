package main

import (
	"context"
	"fmt"
	"lllamaex1/vector"
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

type data struct {
	Name       string
	Text       string
	Embeddings []float32 // the vector data
}

func (d data) Vector() []float32 {
	return d.Embeddings
}

func main() {
	llm, err := ollama.New(
		ollama.WithModel("mxbai-embed-large"),
		ollama.WithServerURL("http://172.20.0.2/ollama"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Apply the feature vectors to the hand crafted data points. This time you
	// need to use words since we are using a word based model.
	dataPoints := []vector.Data{
		data{Name: "Horse   ", Text: "Animal, Female"},
		data{Name: "Man     ", Text: "Human,  Male,   Pants, Poor, Worker"},
		data{Name: "Woman   ", Text: "Human,  Female, Dress, Poor, Worker"},
		data{Name: "King    ", Text: "Human,  Male,   Pants, Rich, Ruler"},
		data{Name: "Queen   ", Text: "Human,  Female, Dress, Rich, Ruler"},
	}

	// Iterate over each data point and use the LLM to generate the vector
	// embedding related to the model.
	for i, dp := range dataPoints {
		dataPoint := dp.(data)

		vectors, err := llm.CreateEmbedding(context.Background(), []string{dataPoint.Text})
		if err != nil {
			log.Fatal(err)
		}

		dataPoint.Embeddings = vectors[0]
		dataPoints[i] = dataPoint
	}
	// Compare each data point to every other by performing a cosine
	// similarity comparison using the vector embedding from the LLM.
	for _, target := range dataPoints {
		results := vector.Similarity(target, dataPoints...)

		for _, result := range results {
			fmt.Printf("%s -> %s: %.3f%% similar\n",
				result.Target.(data).Name,
				result.DataPoint.(data).Name,
				result.Percentage)
		}
		fmt.Print("\n")
	}

	// -------------------------------------------------------------------------

	// Perform the same vector math as in example2 using the LLM vector embedding.

	// You can perform vector math by adding and subtracting vectors.
	kingSubMan := vector.Sub(dataPoints[3].Vector(), dataPoints[1].Vector())
	kingSubManPlusWoman := vector.Add(kingSubMan, dataPoints[2].Vector())
	queen := dataPoints[4].Vector()

	// Now compare a (king - Man + Woman) to a Queen.
	result := vector.CosineSimilarity(kingSubManPlusWoman, queen)
	fmt.Printf("King - Man + Woman ~= Queen similarity: %.3f%%\n", result*100)
}
