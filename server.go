package main

import (
//	"fmt"
//	"log"
//	"net/http"
	"github.com/gorilla/websocket"
)

func wsHandler(conn *websocket.Conn) {
	wsInfoChan := make(chan wsInfo)
	globalConnChan <- newConnInfo{conn, wsInfoChan}

	wsInfo := <-wsInfoChan

	writer(wsInfo)
}

func server(port string, certFile string, keyFile string) {
	globalConnChan = make(chan newConnInfo)

	/*
	var wsServer websocket.Server
	wsServer.Handler = websocket.Handler(wsHandler)

	var httpServer http.Server
	httpServer.Addr = ":" + port
	httpServer.Handler = wsServer

	go func() {
		var err error
		if certFile == "" {
			fmt.Printf("The gowebsock server is listening on port %s\n", port)
			err = httpServer.ListenAndServe()
		} else {
			fmt.Printf("The gowebsock server is listening on port %s using TLS\n", port)
			err = httpServer.ListenAndServeTLS(certFile, keyFile)
		}
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
		case newConnInfo := <-globalConnChan:
			conn := newConnInfo.conn
			wsInfoChan := newConnInfo.wsInfoChan

			readerInfo := wsInfo{conn, readerMessageChan, readerCloseChan}
			go reader(readerInfo)

			writerMessageChan := make(chan wsMessage)
			writerCloseChan := make(chan *websocket.Conn)
			writerInfo := wsInfo{conn, writerMessageChan, writerCloseChan}
			wsInfoChan <- writerInfo

			address := conn.Request().RemoteAddr
			writerInfoMap[address] = writerInfo
			fmt.Printf("New connection: %s\n", address)
		case stdinMsg := <-stdinReaderChan:
			for _, wsInfo := range writerInfoMap {
				wsInfo.messageChan <- wsMessage{wsInfo.conn, []byte(stdinMsg)}
			}
		case wsMessage := <-readerMessageChan:
			address := wsMessage.conn.Request().RemoteAddr
			output := address + ": " + string(wsMessage.bytes) + "\n"
			fmt.Print(output)
		case conn := <-readerCloseChan:
			address := conn.Request().RemoteAddr
			writerInfo := writerInfoMap[address]
			writerInfo.closeChan <- conn
			delete(writerInfoMap, address)
			fmt.Printf("Connection closed: %s\n", address)
		}
	}
*/
}
