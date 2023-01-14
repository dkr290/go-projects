package main

import "fmt"

type language string

var phrazebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",      // English
	"fr": "Bonjour le monde", // French
	"he": "שלום עולם",
	"bg": "Здравей Свят",
	"ru": "Всем привет",
}

func main() {

	greeting := greet("en")
	fmt.Println(greeting)
}

func greet(l language) string {
	if greeting, ok := phrazebook[l]; ok {

		return greeting

	} else {
		return fmt.Sprintf("unsupported language: %q", l)
	}

}
