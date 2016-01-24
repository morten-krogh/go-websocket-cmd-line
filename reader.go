package main

import "github.com/gorilla/websocket"

func pongHandler(appData string) error {
	
	return nil
}

func reader(conn *websocket.Conn, ch chan readerResult) {
	conn.SetPongHandler(func (appData string) error {
		ch <- readerResult{conn, 10, nil, nil}
		return nil
	})

	conn.SetPingHandler(func (appData string) error {
		ch <- readerResult{conn, 9, nil, nil}
		return nil
	})
	
	for {
		for {
			messageType, data, err := conn.ReadMessage()
			readerResult := readerResult{conn, messageType, data, err}
			ch <- readerResult
			if err != nil {
				return
			}
		}
	}
}
