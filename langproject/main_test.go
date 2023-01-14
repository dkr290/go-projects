package main

import "testing"

func TestGreet(t *testing.T) {

	type testCase struct {
		lang             language
		expectedGreeting string
	}

	var tests = map[string]testCase{
		"English": {
			lang:             "en",
			expectedGreeting: "Hello world",
		},
		"French": {
			lang:             "fr",
			expectedGreeting: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang:             "akk",
			expectedGreeting: `unsupported language: "akk"`,
		},
		"Hebrew": {
			lang:             "he",
			expectedGreeting: "שלום עולם",
		},
		"Bulgarian": {
			lang:             "bg",
			expectedGreeting: "Здравей Свят",
		},
		"Russian": {
			lang:             "ru",
			expectedGreeting: "Всем привет",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			greeting := greet(tc.lang)

			if greeting != tc.expectedGreeting {
				t.Errorf("expected %s, got %s", greeting, tc.expectedGreeting)
			}
		})
	}

}
