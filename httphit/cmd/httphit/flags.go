package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"strings"
)

type flags struct {
	url  string
	n, c int
}

const usageText = `
Usage:
  httphit [options] url
Options:`

func (f *flags) parse(s *flag.FlagSet, args []string) error {

	flag.Usage = func() {
		fmt.Fprintf(s.Output(), usageText[1:])
		fmt.Println()
		flag.PrintDefaults()
	}

	//remofing it since customizing message
	//flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")
	s.Var(toNumber(&f.n), "n", "Number of requests to make")
	s.Var(toNumber(&f.c), "c", "Concurrency level")

	if err := s.Parse(args); err != nil {
		return err
	}
	//define url argument to the url field before validation
	f.url = s.Arg(0)

	if err := f.urlValidate(); err != nil {
		fmt.Fprintln(s.Output(), err)
		fmt.Println()
		if strings.Contains(err.Error(), "url: required") {
			s.Usage()
		}
		return err

	}
	if err := f.validateCN(); err != nil {
		fmt.Fprintln(s.Output(), err)
		fmt.Println()
		return err
	}

	return nil
}

func (f *flags) urlValidate() error {

	u, err := url.Parse(f.url)

	switch {
	case strings.TrimSpace(f.url) == "":
		err = errors.New("url: required")
	case err != nil:
		err = errors.New("parse error")

	case chechUscheme(u.Scheme):

		err = errors.New("only supported scheme is http or https")
	case u.Host == "":
		err = errors.New("missing host")

	}

	return err
}

func chechUscheme(s string) bool {

	if s != "https" && s != "http" {

		return true
	}
	return false
}

func (f *flags) validateCN() error {
	if f.c > f.n {

		return fmt.Errorf("-c=%d: should be less then or equal to -n=%d", f.c, f.n)
	}
	return nil
}
