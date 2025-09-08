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
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incomming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
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

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto

	fmt.Println("Received request with HTTP Version:", httpVersion)
	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with tls version:", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"

	default:
		return "Unknown TLS version"
	}
}
