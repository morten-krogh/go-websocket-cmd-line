package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"golang.org/x/net/websocket"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Print("usage: %s web-socket-uri\n", os.Args[0])
		os.Exit(1)
	}
	
	wsUri := os.Args[1]

	// wsUri := "ws://echo.websocket.org/"
	origin := "http://localhost/"

	config, err := websocket.NewConfig(wsUri, origin)

	if err != nil {
		log.Fatal(err)
	}

	ws, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The web socket is connected to %s\n", wsUri)

	inputReader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("\nMessage: ")

		line, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		_, err = ws.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
		}

		msg := make([]byte, 512)
	
		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
	
		fmt.Printf("Reply: %s", msg[:n])
	}
}
