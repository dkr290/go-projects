package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
)

func main() {

	if err := run(flag.CommandLine, os.Args[1:], os.Stdout); err != nil {
		os.Exit(1)
	}

}

func run(s *flag.FlagSet, args []string, out io.Writer) error {
	f := &flags{
		n: 100,
		c: runtime.NumCPU(),
	}

	if err := f.parse(s, args); err != nil {
		return err

	}

	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n",
		f.n, f.url, f.c)
	return nil
}
