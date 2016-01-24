package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	writerCommandChan := make(chan writerCommand)
	writerInitChan <- writerInit{conn, writerCommandChan}
	writer(conn, writerCommandChan)

}

func server(port string, certFile string, keyFile string) {
	writerInitChan = make(chan writerInit)

	var httpServer http.Server
	httpServer.Addr = ":" + port
	httpServer.Handler = http.HandlerFunc(serverHandler)

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

	writerMap := make(map[string]writerInit)

	readerResultChan := make(chan readerResult)

	for {
		select {
		case writerInit := <- writerInitChan:
			conn := writerInit.conn
			go reader(conn, readerResultChan)
			
			address := conn.RemoteAddr().String()
			writerMap[address] = writerInit
			fmt.Printf("New connection: %s\n", address)

/*
			wsInfoChan := newConnInfo.wsInfoChan
			writerMessageChan := make(chan wsMessage)
			writerCloseChan := make(chan *websocket.Conn)
			writerInfo := wsInfo{conn, writerMessageChan, writerCloseChan}
			wsInfoChan <- writerInfo
*/
			
		case stdinMsg := <-stdinReaderChan:
			for _, writerInit := range writerMap {
				writerInit.writerCommandChan <- writerCommand{false, 1, []byte(stdinMsg)}
			}
		case readerResult := <-readerResultChan:
			address := readerResult.conn.RemoteAddr().String()
			output := address + ": " + string(readerResult.data) + "\n"
			fmt.Print(output)
/*
		case conn := <-readerCloseChan:
			address := conn.Request().RemoteAddr
			writerInfo := writerInfoMap[address]
			writerInfo.closeChan <- conn
			delete(writerInfoMap, address)
			fmt.Printf("Connection closed: %s\n", address)
*/
		}
	}
}
