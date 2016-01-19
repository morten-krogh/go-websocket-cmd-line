package main

import (
	"bufio"
//	"fmt"
	"log"
	"os"
)

func stdinReader(c chan string) {

	inputReader := bufio.NewReader(os.Stdin)

	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		c <- line
	}
}
