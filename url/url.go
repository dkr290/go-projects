package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
}

func Parse(url string) (*URL, error) {
	i := strings.Index(url, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}
	scheme := url[:i]

	return &URL{Scheme: scheme}, nil

}
