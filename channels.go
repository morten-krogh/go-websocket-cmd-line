package main

import "github.com/gorilla/websocket"

/* readerResult is sent from the reader to the master
 * when the reader has read from the connection
 */
type readerResult struct {
	conn        *websocket.Conn
	messageType int
	data        []byte
	err         error
}

/* writerCommand is sent by the master to the writer to send it on the writer's connection. */
type writerCommand struct {
	close       bool
	messageType int
	data        []byte
}

/* writerInit is sent by the http handler in a new goroutine to the master */
type writerInit struct {
	conn              *websocket.Conn
	writerCommandChan chan writerCommand
}

/* The gobal channel to send the writerInit on. The master listens to this channel. */
var writerInitChan chan writerInit
