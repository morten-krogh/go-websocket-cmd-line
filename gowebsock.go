package main

import (
	"fmt"
	"os"
)

func printUsage() {
	usageFormatString := "Usage\n  %s client <web-socket-uri>\n  %s server <port>\n  %s server <port> <cert-file> <key-file>\n"
	fmt.Printf(usageFormatString, os.Args[0], os.Args[0], os.Args[0])
}

func main() {

	switch {
	case len(os.Args) == 3 && os.Args[1] == "client":
		wsUri := os.Args[2]
		client(wsUri)
	case len(os.Args) == 3 && os.Args[1] == "server":
		port := os.Args[2]
		server(port, "", "")
	case len(os.Args) == 5 && os.Args[1] == "server":
		port := os.Args[2]
		certFile := os.Args[3]
		keyFile := os.Args[4]
		server(port, certFile, keyFile)
	default:
		printUsage()
	}
}
