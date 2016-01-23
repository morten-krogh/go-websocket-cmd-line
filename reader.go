package main

import "fmt"

func reader(info wsInfo) {
	for {
		for {
			messageType, p, err := info.conn.ReadMessage()
			if err != nil {
				fmt.Printf("error = %s\n", err)
				info.closeChan <- info.conn
				return
			}

			messageType = messageType
			
			info.messageChan <- wsMessage{info.conn, p}
		}
	}
}
