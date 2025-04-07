package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

func checkAndSaveBody(url string, c chan string, errCh chan string) {
	// Background context - root of all contexts
	ctx := context.Background()
	// Context with timeout
	ctxWithTimeout, timeoutCancel := context.WithTimeout(ctx, 5*time.Second)

	defer timeoutCancel()

	req, err := http.NewRequestWithContext(ctxWithTimeout, "GET", url, nil)
	if err != nil {
		s := fmt.Sprintf("error %v:\n", err)
		errCh <- s
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("error %v:\n", err)
		errCh <- s
	} else {
		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				errCh <- fmt.Sprintf("error %v:", err)
			}
			file := strings.Split(url, "//")[1]
			file += ".txt"
			s += fmt.Sprintf("Writing response Body to %s\n", file)
			err = os.WriteFile(file, bodyBytes, 0664)
			if err != nil {

				s += "Error writing to file!\n"

				// sending s over the channel
				errCh <- s
			}
		}
		s += fmt.Sprintf("%s is UP\n", url)
		c <- s
		errCh <- ""
	}
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com"}
	c := make(chan string)
	errCh := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, c, errCh)
	}
	fmt.Println("Number of goroutines is:", runtime.NumGoroutine())
	select {
	case err := <-errCh:
		if err != "" {
			fmt.Println(err)
			return
		}
		// Channel had a value, but it was nil
	default:
		// Channel is empty - no value available
	}

	fmt.Println("Printing results")
	fmt.Println(strings.Repeat("-", 10))
	for range urls {
		fmt.Println(<-c)
	}
}
