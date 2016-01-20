package main

import (
	//	"bufio"
	"fmt"
	//	"net"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	//	"os"
)

var readerChan chan string
var closeChan chan string

type writerInfo struct {
	address  string
	msgChan  chan string
	doneChan chan bool
}

var writerInfoChan chan writerInfo 

func handShake(config *websocket.Config, request *http.Request) error {
	return nil
}

func wsHandler(conn *websocket.Conn) {
	remoteAddr := conn.Request().RemoteAddr

	go wsReaderServer(conn, readerChan)

	msgChan := make(chan string)
	doneChan := make(chan bool)

	writerInfoChan <- writerInfo{remoteAddr, msgChan, doneChan}
	
	for {
		select {
		case msg := <-msgChan:
			conn.Write([]byte(msg))
		case done := <-doneChan:
			if done {
				return
			}
		}
	}
}

func server(port string) {
	readerChan = make(chan string)
	closeChan = make(chan string)
	writerInfoChan = make(chan writerInfo)
	
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
		fmt.Printf("Web socket is listening on port %s\n", port)
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	writerInfos := make(map[string]writerInfo)

	for {
		select {
		case stdinMsg := <-stdinReaderChan:
			for _, writerInfo := range writerInfos {
				writerInfo.msgChan <- stdinMsg
			}
		case writerInfo := <-writerInfoChan:
			fmt.Printf("New connection: %s\n", writerInfo.address)
			writerInfos[writerInfo.address] = writerInfo
		case msg := <-readerChan:
			fmt.Print(msg)
		case msg := <-closeChan:
			writerInfo, ok := writerInfos[msg]
			if ok {
				fmt.Printf("Connection closed: %s\n", writerInfo.address)
				writerInfo.doneChan <- true
				delete(writerInfos, msg)
			}
		}
	}
}
