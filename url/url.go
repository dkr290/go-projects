package url

import "strings"

type URL struct {
	Scheme string
}

func Parse(url string) (*URL, error) {
	i := strings.Index(url, "://")
	scheme := url[:i]

	return &URL{Scheme: scheme}, nil

}
