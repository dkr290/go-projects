package summarizer

import (
	"context"
	"fmt"

	"github.com/ollama/ollama/api"
)

type Config struct {
	client *api.Client
	model  string
	prompt string
	ctx    context.Context
}

func New(client *api.Client, model, prompt string, ctx context.Context) *Config {
	return &Config{
		model:  model,
		client: client,
		prompt: prompt,
		ctx:    ctx,
	}
}

func toPtr[T any](t T) *T {
	return &t
}

func (c *Config) SummarizeText() (string, error) {
	req := &api.GenerateRequest{
		Model:  c.model,
		Prompt: c.prompt,
		Stream: toPtr(false),
	}
	respFunc := func(resp api.GenerateResponse) error {
		fmt.Println(resp.Response)
		return nil
	}
	err := c.client.Generate(c.ctx, req, respFunc)
	if err != nil {
		return fmt.Errorf("cannot query ollama %v", err)
	}

	return nil
}
