package main

import "github.com/gorilla/websocket"

func reader(conn *websocket.Conn, ch chan readResult) {
	for {
		for {
			msgType, data, err := conn.ReadMessage()
			messageType := messageType (msgType)
			readResult := readResult{conn, messageType, data, err}
			ch <- readResult
			if err != nil {
				return
			}
		}
	}
}
