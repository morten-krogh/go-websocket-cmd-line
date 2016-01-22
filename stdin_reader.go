package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func stdinReader(inputChan chan string) {

	inputReader := bufio.NewReader(os.Stdin)

	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSuffix(line, "\n")
		inputChan <- line
	}
}
