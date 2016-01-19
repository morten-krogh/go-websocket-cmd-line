package main

import (
//	"bufio"
//	"fmt"
	"golang.org/x/net/websocket"
	"log"
//	"os"
)

func wsReader(ws *websocket.Conn, c chan []byte) {

	for {
		msg := make([]byte, 4096)

		_, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}

		c <- msg
	}
}
