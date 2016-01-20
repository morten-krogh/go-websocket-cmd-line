package main

import (
//	"bufio"
//	"fmt"
//	"golang.org/x/net/websocket"
//	"log"
//	"os"
)

func reader(info wsInfo) {
	buffer := make([]byte, 65536)
	
	for {
		n, err := info.conn.Read(buffer)
		if err != nil {
			info.closeChan <- info.conn
			return
		}
		info.messageChan <- wsMessage{info.conn, buffer[:n]}
	}
}
