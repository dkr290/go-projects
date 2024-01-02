package url_test

import (
	"fmt"
	"log"

	"github.com/dkr290/go-projects/url"
)

func ExampleURL() {

	u, err := url.Parse("https://foo.com/python")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)

	//Output:
	//https://foo.com/python
}

func ExampleURL_fields() {

	u, err := url.Parse("https://foo.com/python")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u)

	//Output:
	//https
	//foo.com
	//python
	//https://foo.com/python
}
