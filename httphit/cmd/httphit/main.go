package main

import (
	"fmt"
	"log"
)

const usageText = `
		  -url
		  		HTTP server URL to make requests (required)
		  -n
		  		Number of requests to make
		  -c
		 		 Concurrency level`

func usage() string { return usageText[1:] }
func main() {

	var f flags

	if err := f.parse(); err != nil {
		fmt.Println(usage())
		log.Fatal(err)

	}
	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n",
		f.n, f.url, f.c)
}
