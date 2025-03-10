package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"text-summarizer/pkg/summarizer"

	"github.com/ollama/ollama/api"
)

type Handlers struct {
	template     template.Template
	ollamaClient *api.Client
	summarizer   *summarizer.Config
}

func New(s *summarizer.Config, olClient *api.Client, temp template.Template) *Handlers {
	return &Handlers{
		template:     temp,
		ollamaClient: olClient,
		summarizer:   s,
	}
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) error {
	err := h.template.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) SummarizeHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("method not allowed")
	}

	text := r.FormValue("text")
	if text == "" {
		renderResult(w, "Error: Please input text to summarize", true)
		return fmt.Errorf("error no text")
	}

	err := h.summarizer.SummarizeText()
	if err != nil {
		renderResult(w, "Error: Failed to generate summary: "+err.Error(), true)
		return err
	}

	renderResult(w, summary, false)
}

func renderResult(w http.ResponseWriter, message string, isError bool) {
	data := struct {
		Message string
		IsError bool
	}{
		Message: message,
		IsError: isError,
	}

	err := templates.ExecuteTemplate(w, "result.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
