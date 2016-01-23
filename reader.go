package main

func reader(info wsInfo) {
	for {
		buffer := make([]byte, 1024)
		bytesRead := 0
		allOfFrameRead := false
		for !allOfFrameRead {
			n, err := info.conn.Read(buffer[bytesRead:])

			if err != nil {
				info.closeChan <- info.conn
				return
			}
			bytesRead += n
			if bytesRead < len(buffer) {
				allOfFrameRead = true
			} else {
				newBuffer := make([]byte, 2*len(buffer))
				copy(newBuffer, buffer)
				buffer = newBuffer
			}
		}
		info.messageChan <- wsMessage{info.conn, buffer[:bytesRead]}
	}
}
