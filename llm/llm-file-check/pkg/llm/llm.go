package llm

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

func LLmreplace(valuesFile, chartFile []byte, internalRegistry string, ollamaHost string) {
	// Prompt for Ollama AI
	prompt := fmt.Sprintf(`
Here is a Helm values.yaml file:

%s

And here is the Chart.yaml file:

%s

Please update the values.yaml file:
- Replace any Docker images from "docker.io" with "%s".
- If an image has no tag or tag is null, replace it with the "appVersion" from Chart.yaml.
- If the docker.io/repository is missing and you see only repository: image/name this is by default in docker.io and also needs to be replaced

Return only the corrected YAML content.
`, valuesFile, chartFile, internalRegistry)

	os.Setenv("OLLAMA_HOST", ollamaHost) // Default value
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Second)
	defer cancel()
	genReq := &api.GenerateRequest{
		Model:  "llama3.2",
		Prompt: prompt,
	}
	// Capture the full response
	var fullResponse string
	genResp := func(r api.GenerateResponse) error {
		fullResponse += r.Response
		return nil
	}
	err = client.Generate(ctx, genReq, genResp)
	if err != nil {
		log.Fatal("Error generating response from Ollama:", err)
	}
	// Write modified YAML to a new file
	err = os.WriteFile("modified_values.yaml", []byte(fullResponse), 0644)
	if err != nil {
		log.Fatal("Error writing modified_values.yaml:", err)
	}

	fmt.Println("✅ Updated values.yaml saved as modified_values.yaml!")
}
