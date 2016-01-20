package main

import (
//	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
//	"log"
	"os"
)

func wsReaderClient(conn *websocket.Conn, c chan []byte) {
	buffer := make([]byte, 65536)
	
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Connection closed\n")
			os.Exit(0)
		}

		c <- buffer[:n]
	}
}
