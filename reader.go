package main

import "github.com/gorilla/websocket"

/* reader reads input from the web socket connection and sends it to the master.
 * reader runs in its own goroutine.
 * If the connection closes, the reader goroutine terminates.
 */
func reader(conn *websocket.Conn, ch chan readerResult) {
	conn.SetPongHandler(func(appData string) error {
		ch <- readerResult{conn, 10, nil, nil}
		return nil
	})

	conn.SetPingHandler(func(appData string) error {
		ch <- readerResult{conn, 9, nil, nil}
		return nil
	})

	for {
		messageType, data, err := conn.ReadMessage()
		readerResult := readerResult{conn, messageType, data, err}
		ch <- readerResult
		if err != nil {
			return
		}
	}
}
