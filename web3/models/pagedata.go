package models

import "github.com/dkr290/go-projects/web3/pkg/forms"

type PageData struct {
	StrMap    map[string]string
	IntMap    map[string]int
	FltMap    map[string]float32
	DataMap   map[string]interface{}
	CSRFToken string
	Warning   string
	Error     string
	Form      *forms.Form
	Data      map[string]interface{}
}
