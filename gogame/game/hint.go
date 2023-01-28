package game

import "strings"

type hint byte

type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		return "💔"
	}
}

// StringConcat is a naive implementation to build feedback as a string.
// It is used only to benchmark it against the strings.Builder version.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}
