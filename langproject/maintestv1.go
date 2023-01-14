package main

import "testing"

func TestGreet_English(t *testing.T) {

	lang := "en"
	expectedGreeting := "Hello World"
	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got:%q", expectedGreeting, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	lang := "fr"
	expectedGreeting := "Bonjour le monde"
	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got:%q", expectedGreeting, greeting)
	}
}

func TestGreet_Empty(t *testing.T) {

	lang := "akk"
	expectedGreeting := ""
	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got:%q", expectedGreeting, greeting)
	}
}
