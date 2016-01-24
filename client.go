package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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

	readResultChan := make(chan readResult)
	go reader(conn, readResultChan)

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
		case readResult := <-readResultChan:
			if readResult.err == nil {
				output := "Server: type = " + string(readResult.messageType.string()) + ", data = " + string(readResult.data) + "\n"
				fmt.Printf(output)
			} else {
				fmt.Printf("%s\n", readResult.err)
			}
		}
	}
}
