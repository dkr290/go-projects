package models

import "github.com/dkr290/go-projects/goforum-app/goforum/pkg/forms"

type PageData struct {
	StrMap    map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	DataMap   map[string]any
	CSRFToken string
	Warning   string
	Error     string
	Form      *forms.Form
	Data      map[string]any
}
