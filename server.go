package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handles users")
	})

	port := 5000

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server running on port: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}

	// HTTP1 server without tls
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP version: ", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersion(r.TLS.Version)
		fmt.Println("Received request with TLS version: ", tlsVersion)
	}
}

func getTLSVersion(version uint16) string {
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
		return "TLS version not found!"
	}
}
