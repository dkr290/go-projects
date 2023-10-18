package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/flosch/pongo2"
)

// it is for the application configuration
type AppConfig struct {
	TempleteCache map[string]*pongo2.Template
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

// the new appconfig fuction just to create new appconfig and initialize the template
func NewConfig(a map[string]*pongo2.Template, c bool, inProduction bool, s *scs.SessionManager) *AppConfig {
	return &AppConfig{
		TempleteCache: a,
		UseCache:      c,
		InProduction:  inProduction,
		Session:       s,
	}
}
