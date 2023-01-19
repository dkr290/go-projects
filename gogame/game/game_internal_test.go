package game

import (
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
			input: "hello",
			want:  []rune("hello"),
		},
		"5 characters in cyrilic": {
			input: "какво",
			want:  []rune("какво"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := NewGame(strings.NewReader(tc.input))
			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("read runes got  = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}
