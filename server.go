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

/*
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
*/

func wsHandler(conn *websocket.Conn) {
	globalConnChan <- conn
/*
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
*/
}

func server(port string) {

	var wsServer websocket.Server
	//wsServer.Handshake = handShake
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

	stdinReaderChan := make(chan string)
	go stdinReader(stdinReaderChan)

	writerInfoMap := make(map[string]wsInfo)

	readerMessageChan := make(chan wsMessage)
	readerCloseChan := make(chan *websocket.Conn)
	
	for {
		select {
		case conn := <- globalConnChan:
			readerInfo := wsInfo{conn, readerMessageChan, readerCloseChan}
			go reader(readerInfo)

			writerMessageChan := make(chan wsMessage)
			writerCloseChan := make(chan *websocket.Conn)
			writerInfo := wsInfo{conn, writerMessageChan, writerCloseChan}
			go writer(writerInfo)

			address := conn.RemoteAddr().String()
			writerInfoMap[address] = writerInfo
			fmt.Printf("New connection: %s\n", address)
		case stdinMsg := <-stdinReaderChan:
			for _, wsInfo := range writerInfoMap {
				wsInfo.messageChan <- wsMessage{wsInfo.conn, []byte(stdinMsg)}
			}
		case wsMessage := <-readerMessageChan:
			address := wsMessage.conn.RemoteAddr().String()
			output := address + ": " + string(wsMessage.bytes)
			fmt.Print(output)
		case conn := <-readerCloseChan:
			address := conn.RemoteAddr().String()
			writerInfo := writerInfoMap[address]
			writerInfo.closeChan <- conn
			delete(writerInfoMap, address)
			fmt.Printf("Connection closed: %s\n", address)
		}
	}
}
