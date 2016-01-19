package main

import (
	//	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	//	"os"
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

	wsReaderChan := make(chan []byte)

	go wsReader(ws, wsReaderChan)

	stdinReaderChan := make(chan string)

	go stdinReader(stdinReaderChan)

	for {
		select {
		case stdinMsg := <-stdinReaderChan:
			print("stdin: ", stdinMsg)
		}
	}

	/*
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

	*/
}
