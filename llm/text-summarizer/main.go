package main

import (
	"context"
	"log"
	"net/http"
	"text-summarizer/pkg/handlers"
	"text-summarizer/pkg/helpers"
	"text-summarizer/pkg/summarizer"
	"time"

	"github.com/ollama/ollama/api"
)

func main() {
	run()
}

func run() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Second)
	defer cancel() // Always call cancel to release resources
	temp, err := template.ParseGlob("templates/*.html")

	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}
	prompt := `Summmarize the following text:
  Artificial intelligence is transforming the industried acress the world. AI models like DeppSeek R1 enable
  businesses to automate tasks,  analize large datasets, and enahance productiviti. With advancements of AI, applications
  range from virtusal assistants to predictive analytics and personalized reccomendations.
  `
	summarizer := summarizer.New(client, "deepseek-r1", prompt, ctx)
	h := handlers.New(s, client, temp)

	http.HandleFunc("/", helpers.MakeHandler(next func(http.ResponseWriter, *http.Request) error))
	http.HandleFunc("/summarize", summarizeHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
