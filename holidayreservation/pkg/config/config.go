package config

import (
	"log"

	"github.com/flosch/pongo2"
)

// it is for the application configuration
type AppConfig struct {
	TempleteCache map[string]*pongo2.Template
	UseCache      bool
	InfoLog       *log.Logger
	TemplateData
}

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]string
	FlotMap   map[string]string
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

// the new appconfig fuction just to create new appconfig and initialize the template
func NewConfig(a map[string]*pongo2.Template, c bool) *AppConfig {
	return &AppConfig{
		TempleteCache: a,
		UseCache:      c,
	}
}
