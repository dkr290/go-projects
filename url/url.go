package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func Parse(url string) (*URL, error) {

	scheme, rest, ok := parseScheme(url)
	if !ok {
		return nil, errors.New("missing scheme")
	}

	host, path := parseHost(rest)

	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil

}
func parseScheme(rawurl string) (scheme, rest string, ok bool) {

	i := strings.Index(rawurl, "://")
	if i < 1 {
		return "", "", false
	}

	return rawurl[:i], rawurl[i+3:], true

}

func parseHost(rest string) (host, path string) {

	if i := strings.Index(rest, "/"); i > 0 {
		host, path = rest[:i], rest[i+1:]
	}

	return host, path

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

	if u == nil {
		return ""
	}

	var s string
	if sc := u.Scheme; sc != "" {
		s += sc
		s += "://"
	}
	if h := u.Host; h != "" {

		s += h
	}
	if p := u.Path; p != "" {

		s += "/"
		s += p
	}
	return s

}
