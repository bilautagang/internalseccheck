package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <IP>")
		return
	}

	ip := os.Args[1]

	// Ensure the URL has a scheme
	if ip[:4] != "http" {
		ip = "https://" + ip
	}

	// Create a custom HTTP client with a transport that skips certificate verification
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: customTransport}

	// Get the response and capture the TLS connection state
	resp, err := client.Get(ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Check for HSTS by looking for the Strict-Transport-Security header
	hsts := resp.Header.Get("Strict-Transport-Security")
	if hsts != "" {
		fmt.Println("HSTS: Enabled")
	} else {
		fmt.Println("HSTS: Not Enabled")
	}

	// Check for X-Frame-Options header
	xFrameOptions := resp.Header.Get("X-Frame-Options")
	if xFrameOptions != "" {
		fmt.Println("X-Frame-Options:", xFrameOptions)
	} else {
		fmt.Println("X-Frame-Options: Not Set")
	}

	// Check the TLS version
	tlsVersion := resp.TLS.Version
	switch tlsVersion {
	case tls.VersionTLS13:
		fmt.Println("TLS Version: 1.3")
	case tls.VersionTLS12:
		fmt.Println("TLS Version: 1.2")
	case tls.VersionTLS11:
		fmt.Println("TLS Version: 1.1")
	case tls.VersionTLS10:
		fmt.Println("TLS Version: 1.0")
	default:
		fmt.Println("TLS Version: Unknown")
	}
}
