package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Print("usage: %s web-socket-uri\n", os.Args[0])
		os.Exit(1)
	}

	wsUri := os.Args[1]

	client(wsUri)
}
