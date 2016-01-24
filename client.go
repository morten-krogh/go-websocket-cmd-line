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

	readerResultChan := make(chan readerResult)
	go reader(conn, readerResultChan)

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
		case readerResult := <-readerResultChan:
			if readerResult.err == nil {
				output := "Server: type = " + messageTypeString(readerResult.messageType) + ", data = " + string(readerResult.data) + "\n"
				fmt.Printf(output)
			} else {
				fmt.Printf("%s\n", readerResult.err)
				os.Exit(0)
			}
		}
	}
}
