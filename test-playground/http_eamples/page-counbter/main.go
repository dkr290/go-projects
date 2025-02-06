package main

import (
	"fmt"
	"log"
	"net/http"
)

type pageCounter struct {
	Counter int    `json:"counter"`
	Heading string `json:"heading"`
	Content string `json:"content"`
}

func (p *pageCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Counter++
	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s<p>\n<p>Views: %d</p>", p.Heading, p.Content, p.Counter)
	_, err := w.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	hello := pageCounter{
		Heading: "Home page",
		Content: "This is the main page",
	}
	cha1 := pageCounter{
		Heading: "Chapter 1",
		Content: "This is hte first chapter",
	}
	cha2 := pageCounter{
		Heading: "Chapter 2",
		Content: "This is the second chapter",
	}

	http.Handle("/", &hello)
	http.Handle("/chap1", &cha1)
	http.Handle("/chap2", &cha2)
	fmt.Println("starting the server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
