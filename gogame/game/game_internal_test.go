package game

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameAsk(t *testing.T) {

	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in cyrilic": {
			input: "КАКВО",
			want:  []rune("КАКВО"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := NewGame(strings.NewReader(tc.input), string(tc.want), 5)
			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("read runes got  = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

// func TestValidateGuess(t *testing.T) {

// 	tt := map[string]struct { // define required inputand the expected output
// 		word     []rune
// 		expected error
// 	}{
// 		"nominal": { // write our scenarious, good one and not good
// 			word:     []rune("GUESS"),
// 			expected: nil,
// 		},
// 		"too long": {
// 			word:     []rune("verylongstring"),
// 			expected: errInvlidWorldLength,
// 		},
// 		"too short": {
// 			word:     []rune("ggg"),
// 			expected: errInvlidWorldLength,
// 		},
// 	}
// 	for name, tc := range tt {
// 		t.Run(name, func(t *testing.T) {
// 			g := NewGame(nil)                 //it does not require reader and we can give it nil
// 			err := g.validateGuess(tc.word)   // call the method using game object and give it the word
// 			if !errors.Is(err, tc.expected) { // compare the two errors
// 				t.Errorf("error %c, expected %q, got %q", tc.word, tc.expected, err) //%c can be used to print content of the slice
// 			}
// 		})
// 	}
// }

// same tests but with struct
func TestValidateGuess1(t *testing.T) {

	type tt struct { // define required inputand the expected output
		description string
		input       string
		maxAttempts int
		word        []rune
		expected    error
	}

	for _, scenario := range []tt{
		{
			description: "nominal",
			input:       "GUESS",
			maxAttempts: 5,
			word:        []rune("GUESS"),
			expected:    nil,
		},
		{
			description: "too long",
			input:       "too long",
			maxAttempts: 5,
			word:        []rune("verylongstring"),
			expected:    errInvlidWorldLength,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			g := NewGame(nil, scenario.input, scenario.maxAttempts) //it does not require reader and we can give it nil
			err := g.validateGuess(scenario.word)                   // call the method using game object and give it the word
			if !errors.Is(err, scenario.expected) {                 // compare the two errors
				t.Errorf("error %c, expected %q, got %q", scenario.word, scenario.expected, err) //%c can be used to print content of the slice
			}
		})
	}

}
