package main

import (
	//	"bufio"
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"os"
)

func handShake(config *websocket.Config, request *http.Request) error {

	fmt.Print("handshaek\n")
	return nil
}

func server(port string) {

	fmt.Printf("server : %s\n", port)
	os.Exit(0)


	var wsServer websocket.Server
	wsServer.Handshake = handShake
	

	var httpServer http.Server 
	httpServer.Addr = ":" + port
	httpServer.Handler = wsServer
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Web socket is listening on port %s\n", port)	
	
	/*
	origin := "http://localhost/"

	config, err := websocket.NewConfig(wsUri, origin)

	if err != nil {
		log.Fatal(err)
	}

	ws, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}



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

*/
}
