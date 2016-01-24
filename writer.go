package main

import "github.com/gorilla/websocket"

func writer(conn *websocket.Conn, writerCommandChan chan writerCommand) {
	for {
		select {
		case writerCommand := <-writerCommandChan:
			if writerCommand.close {
				conn.Close()
				return
			} else {
				conn.WriteMessage(writerCommand.messageType, writerCommand.data)
			}
		}
	}
}
