package main

import (
	"github.com/gorilla/websocket"
)

/* wsMessage is sent on channels. It contains the message and a pointer to the connection */

type wsMessage struct {
	conn  *websocket.Conn
	bytes []byte
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
	conn        *websocket.Conn
	messageChan chan wsMessage
	closeChan   chan *websocket.Conn
}

/* The globalConnChan is used to send the connection from the
 * original goroutine that is created at accept to the main goroutine.
 * The main goroutine will distribute the connection to a reader.
 * The original goroutine itself becomes the writer.
 */
type newConnInfo struct {
	conn       *websocket.Conn
	wsInfoChan chan wsInfo
}

var globalConnChan chan newConnInfo
