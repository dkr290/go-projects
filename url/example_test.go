package url_test

import (
	"fmt"
	"log"

	"github.com/dkr290/go-projects/url"
)

func ExampleUrl() {

	u, err := url.Parse("http://foo.com/python")

	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "https"
	fmt.Println(u)

	//Output:
	//&{https foo.com python}
}
