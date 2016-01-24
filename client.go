package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
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

	readResultChan := make(chan readerResult)
	go reader(conn, readResultChan)

	writerCommandChan := make(chan writerCommand)
	go writer(conn, writerCommandChan)

	stdinReaderChan := make(chan string)
	go stdinReader(stdinReaderChan)

	for {
		select {
		case stdinMessage := <-stdinReaderChan:
			var messageType int
			data := ""
			switch stdinMessage {
			case "close":
				messageType = 8
			case "ping":
				messageType = 9
			case "pong":
				messageType = 10
			default:
				messageType = 1
				data = stdinMessage
			}
			writerCommandChan <- writerCommand{false, messageType, []byte(data)}
		case readResult := <-readResultChan:
			if readResult.err == nil {
				output := "Server: type = " + messageTypeString(readResult.messageType) + ", data = " + string(readResult.data) + "\n"
				fmt.Printf(output)
			} else {
				fmt.Printf("%s\n", readResult.err)
				os.Exit(0)
			}
		}
	}
}
