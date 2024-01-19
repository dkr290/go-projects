package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type flags struct {
	url  string
	n, c int
}

func (f *flags) parse() error {

	flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")
	flag.IntVar(&f.n, "n", f.n, "Number of requests to make")
	flag.IntVar(&f.c, "c", f.c, "Concurrency level")

	flag.Parse()

	if err := f.urlValidate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println()
		flag.Usage()
		return err
	}

	return nil
}

func (f *flags) urlValidate() error {

	if strings.TrimSpace(f.url) == "" {
		return errors.New("-url: required")
	}

	return nil
}
