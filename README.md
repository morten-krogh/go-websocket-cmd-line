# Gowebsock

## Command line program

Gowebsock is a command line program for websocket client and server. Gowebsock is written in the go language.

## Usage

```
Usage
  gowebsock client <web-socket-uri>
  gowebsock server <port>
  
```

Websocket messages are typed at the terminal prompt. At the press of the return key, the message is sent on the websocket. The incoming messages are written to the terminal. Messages typed into the server terminal are sent to all connected clients.

### Example server and two clients

##### Server
```
~/go> bin/gowebsock server 9000
The gowebsock server is listening on port 9000
New connection: 127.0.0.1:65120
New connection: 127.0.0.1:65125
127.0.0.1:65120: Hi server!
127.0.0.1:65125: Hello server!
Good day to both of you!
Connection closed: 127.0.0.1:65120
Connection closed: 127.0.0.1:65125
``` 

##### Client 1
```
~/go> bin/gowebsock client ws://127.0.0.1:9000
The gowebsock client is connected to ws://127.0.0.1:9000
Hi server!
Server: Good day to both of you!
```

##### Client 2
```
~/go> bin/gowebsock client ws://127.0.0.1:9000
The gowebsock client is connected to ws://127.0.0.1:9000
Hello server!
Server: Good day to both of you!
```

## Installation

Gowebsock is a standard Go program and can be installed with

```
go install github.com/morten-krogh/gowebsock/
```

or just 

```
go install
```

from within the gowebsock directory. The executable is located in the bin directory and can be copied to anywhere.


## Purpose of Gowebsock 


The Gowebsock client can of course be used to communicate with public websocket servers. In that regard, the client acts like a simple browser.

However, the real utility is to test websocket implementations. A server implementation can be tested by use of several clients. The clients can send hand crafted messages. In this way, the server can be interactively tested, and various corner cases can be resolved quickly.

Likewise, but less common, a client implementation can be tested by having the client connect to a Gowebsock server. Server replies can be manually typed.

## Internal architecture of Gowebsock

Gowebsock is built using the standard Go websocket package "golang.org/x/net/websocket". New websocket connections are created in their own goroutine. Reading from a websocket is a blocking operation. When a user types a message, the message must be sent on the sockets. It is not possible for one goroutine to be blocked on reading and at the same time write to the socket. The simplest solution, and the one employed by Gowebsock, is to have one goroutine for reading and one for writing. 

A Gowebsock program has one master goroutine, one goroutine for reading from stdin, one reader goroutine for each open websocket, and one writer goroutine for each open websocket. The master goroutine coordinates everything and all communication goes through the master. For the server, there is also a goroutine that listens for new connections.

##### Reader goroutine
A reader goroutine blocks on reading from its websocket. When a message has arrived, the message and the address of the socket is sent to the master goroutine through a global channel. If the websocket cnnection is closed by the remote end, the reader notices and sends the close signal to the master thorugh another global channel.

##### Stdin goroutine
The stdin goroutine reads from the terminal and sends the message to the master thorugh a dedicated channel.

##### Listener goroutine
The listener, in case of the server, is a special goroutine that accets new websocket connections and creates a new goroutine for each connection. The new goroutine becomes a writer.

##### Writer goroutine
A writer goroutine is created at connection acceptance. The writer tells the master that it has been created thorugh a global channel. The writer also creates a new channel that it sends to the master. The master will send messages to thw writer through this new channel in the future. The master tells the writer to write on its websocket and to terminate itself when the websocket has been closed.

##### Master goroutine
The master coordinates everything. It keeps a map of all open web sockets and channels to the writers. The master blocks on a channel select where it waits for new connections, reader messages, stdin messages, and closing of connections. The server sends stdin messages to all connected websockets.  
