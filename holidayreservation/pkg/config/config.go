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
	Data
}

type Data struct {
	Title       string
	Description string
}

// the new appconfig fuction just to create new appconfig and initialize the template
func NewConfig(a map[string]*pongo2.Template, c bool) *AppConfig {
	return &AppConfig{
		TempleteCache: a,
		UseCache:      c,
	}
}
