package main

import (
	//	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

func client(wsUri string) {

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

	typeMsg := "Type a message to send on the websocket and press return\n"
	
	println(typeMsg)
	
	wsReaderChan := make(chan []byte)

	go wsReader(ws, wsReaderChan)

	stdinReaderChan := make(chan string)

	go stdinReader(stdinReaderChan)

	for {
		select {
		case stdinMsg := <-stdinReaderChan:
			_, err = ws.Write([]byte(stdinMsg))
			if err != nil {
				println("\nError sending message to the websocket server")
			} else {
				println("\nMessage sent to the websocket server")
			}
			println(typeMsg)
		case wsMsg := <- wsReaderChan:
			println("The server replied:\n ")
			println(string(wsMsg))
			println(typeMsg)
		}
	}
}
