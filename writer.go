package main

import (
	//	"bufio"
//	"fmt"
	//	"net"
//	"golang.org/x/net/websocket"
//	"log"
//	"net/http"
	//	"os"
)

func writer(info wsInfo) {

	for {
		select {
		case wsMsg := <- info.messageChan:
			info.conn.Write(wsMsg.bytes)
		case <- info.closeChan:
			return
		}
	}
}
