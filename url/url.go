package url

import (
	"errors"
	"fmt"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(url string) (*URL, error) {

	i := strings.Index(url, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}
	scheme, rest := url[:i], url[i+3:]
	host, path := rest, ""
	if i := strings.Index(rest, "/"); i > 0 {
		host, path = rest[:i], rest[i+1:]
	}

	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil

}

func (u *URL) Port() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return ""
	}
	return u.Host[i+1:]
}

// Hostname returns u.Host, stripping any port number if present.

func (u *URL) HostName() string {
	i := strings.Index(u.Host, ":")

	if i < 0 {
		return u.Host
	}

	return u.Host[:i]
}

func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
