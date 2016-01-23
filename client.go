package main

import (
	"net/http"
	"crypto/tls"
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

func client(wsUri string) {

	tlsConfig := tls.Config{}
	tlsConfig.InsecureSkipVerify = true
	dialer := websocket.Dialer{TLSClientConfig: &tlsConfig}
	requestHeader := http.Header{}
	requestHeader.Set("origin", "http://localhost/")
	conn, _, err := dialer.Dial(wsUri, requestHeader)
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
		case wsMessage := <-readerMessageChan:
			output := "Server: " + string(wsMessage.bytes) + "\n"
			print(output)
		case <-readerCloseChan:
			output := "The server closed the connection"
			println(output)
			return
		}
	}
}
