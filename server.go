package main

import (
	//	"bufio"
	"fmt"
	//	"net"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	//	"os"
	"time"
)

var wsReaderChan chan string = make(chan string)
var wsWriterChan chan string = make(chan string)

func handShake(config *websocket.Config, request *http.Request) error {
	//fmt.Print("handshake\n")
	return nil
}

func wsHandler(conn *websocket.Conn) {
	remoteAddr := conn.Request().RemoteAddr
	fmt.Printf("New connection: %s\n", remoteAddr)

	go wsReaderServer(conn, wsReaderChan)

	for {
		//fmt.Println(conn.IsServerConn())
		select {
		case msg := <- wsWriterChan:
			_, err := conn.Write([]byte(msg))
			if err != nil {
				return
			}
		case <-time.After(5 * time.Second):
		}
	}
}

func server(port string) {

	stdinReaderChan := make(chan string)
	go stdinReader(stdinReaderChan)
	
	var wsServer websocket.Server
	wsServer.Handshake = handShake
	//wsServer.Config
	wsServer.Handler = websocket.Handler(wsHandler)

	var httpServer http.Server
	httpServer.Addr = ":" + port
	httpServer.Handler = wsServer
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Web socket is listening on port %s\n", port)
	}()

	
	
	for {
		select {
		case stdinMsg := <- stdinReaderChan:
			fmt.Print("something\n")
			stdinMsg = stdinMsg
		case wsMsg := <- wsReaderChan:
			fmt.Print(wsMsg)
		}
	}
	
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
