package main

func writer(info wsInfo) {

	for {
		select {
		case wsMsg := <-info.messageChan:
			info.conn.Write([]byte("OK"))
			info.conn.Write(wsMsg.bytes)
		case <-info.closeChan:
			return
		}
	}
}
