package main

import (
	"bufio"
	"log"
	"os"
)

func stdinReader(inputChan chan string) {

	inputReader := bufio.NewReader(os.Stdin)

	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		inputChan <- line
	}
}
