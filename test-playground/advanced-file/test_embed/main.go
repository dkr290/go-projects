package main

import (
	"embed"
	"html/template"
	"os"
)

type Person struct {
	Name string
}

//go:embed templates/*
var f embed.FS

func main() {
	p := Person{
		Name: "Kal",
	}
	tmpl, err := template.ParseFS(f, "templates/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
