package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	f := &flags{
		n: 100,
		c: runtime.NumCPU(),
	}

	if err := f.parse(); err != nil {
		os.Exit(1)

	}
	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n",
		f.n, f.url, f.c)
}
