package models

type PageData struct {
	StrMap    map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	DataMap   map[string]any
	CSRFToken string
	Warning   string
	Error     string
}
