package main

import (
//	"bufio"
//	"fmt"
	"golang.org/x/net/websocket"
//	"log"
//	"os"
)

func wsReaderServer(conn *websocket.Conn, c chan string) {
	remoteAddr := conn.Request().RemoteAddr
	remoteAddr = remoteAddr
	
	buffer := make([]byte, 65536)
	
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			closeChan <- remoteAddr
			return
		}

		msg := remoteAddr + ": " + string(buffer[:n])
		c <- msg
	}
}
