package main

import (
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
	conn, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The gowebsock client is connected to %s\n", wsUri)

	readerMessageChan := make(chan wsMessage)
	readerCloseChan := make(chan *websocket.Conn)
	readerInfo := wsInfo{conn, readerMessageChan, readerCloseChan}
	go reader(readerInfo)
	
	writerMessageChan := make(chan wsMessage)
	writerCloseChan := make(chan *websocket.Conn)
	writerInfo := wsInfo{conn, writerMessageChan, writerCloseChan}
	go writer(writerInfo)
	
	stdinReaderChan := make(chan string)
	go stdinReader(stdinReaderChan)

	for {
		select {
		case stdinMessage := <-stdinReaderChan:
			writerMessageChan <- wsMessage{conn, []byte(stdinMessage)}
		case wsMessage := <- readerMessageChan:
			output := "Server: " + string(wsMessage.bytes)
			print(output)
		case <- readerCloseChan:
			output := "The server closed the connection"
			println(output)
			return
		}
	}
}
