package main

import "github.com/gorilla/websocket"

func writer(info wsInfo) {
	for {
		select {
		case wsMsg := <-info.messageChan:
			info.conn.WriteMessage(websocket.TextMessage, wsMsg.bytes)
		case <-info.closeChan:
			return
		}
	}
}
