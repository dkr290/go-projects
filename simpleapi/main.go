package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	port := 8080

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incomming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling users ")
	})

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	if err := http2.ConfigureServer(&server, &http2.Server{}); err != nil {
		log.Fatalln("Error configuing the http2 server", err)
	}

	fmt.Println("server is running on port:", port)

	// if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
	// 	log.Fatalln("Could not start the server", err)
	// }

	if err := server.ListenAndServeTLS(cert, key); err != nil {
		log.Fatalln("Could not start the server", err)
	}
}
