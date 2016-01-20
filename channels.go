package main

import (
	"golang.org/x/net/websocket"
)

/* wsMessage is sent on channels. It contains the message and a pointer to the connection */ 

type wsMessage struct {
	conn *websocket.Conn
	msg []byte
}

/* wsInfo is by readers and writers. 
 * A Reader readS messages from the blocking connection
 * and sends the messages to the messageChan. When a reader sees the remote closing of the
 * connection, it signals true on the closeChan.
 *  
 * A writer waits for the messageChan and closeChan. Messages from the messageChan are sent on the
 * connection. True on the closeChan signals the writer to terminate its goroutine.
 */

type wsInfo struct {
	conn *websocket.Conn
	messageChan chan wsMessage
	closeChan chan bool
}

/* Some channels are global. Others are created at connection accept.
 * The main goroutine selects for the global channels 
 */

var globalMsgChan chan wsMessage
var globalCloseChan chan bool
var globalWriterChan chan wsInfo 


