# Gowebsock

## Command line program

Gowebsock is a command line program for websocket clients and servers. Gowebsock is written in the go language.

## Usage

```
Usage
  gowebsock client <web-socket-uri>
  gowebsock server <port>
  gowebsock server <port> <cert-file> <key-file>
  
```

Websocket messages are typed at the terminal prompt. At the press of the return key, the message is sent on the websocket. The incoming messages are written to the terminal. Messages typed into the server terminal are sent to all connected clients.

There are three special messages: ping, pong and close. They can be invoked by typing the corresponding word.  

## Example with a server and two clients

##### Server
```
~/go> gowebsock server 9000
The gowebsock server is listening on port 9000
New connection: 127.0.0.1:50039
New connection: [::1]:50042
127.0.0.1:50039: type = TextMessage, data = Hi server!
[::1]:50042: type = TextMessage, data = Hello server!
Good day to both of you!
127.0.0.1:50039: type = PingMessage, data = 
pong
close
127.0.0.1:50039: websocket: close 1005 
[::1]:50042: websocket: close 1005
```

##### Client 1
```
~/go> gowebsock client ws://127.0.0.1:9000
The gowebsock client is connected to ws://127.0.0.1:9000
Hi server!
Server: type = TextMessage, data = Good day to both of you!
ping
Server: type = PongMessage, data = 
websocket: close 1005 
```

##### Client 2
```
~/go> gowebsock client ws://localhost:9000
The gowebsock client is connected to ws://localhost:9000
Hello server!
Server: type = TextMessage, data = Good day to both of you!
Server: type = PongMessage, data = 
websocket: close 1005 
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

#### Dependencies

The only dependency besides the Go standard library is the Gorilla websocket library:

```
go get github.com/gorilla/websocket
```


## Purpose of Gowebsock 


The Gowebsock client can, of course, be used to communicate with public websocket servers. In that regard, the client acts like a simple browser.

However, the real utility is to test websocket implementations. A server implementation can be tested by use of several clients. The clients can send hand crafted messages. In this way, the server can be interactively tested, and various corner cases can be resolved quickly.

Likewise, but less common, a client implementation can be tested by having the client connect to a Gowebsock server. Server replies can then be manually typed.


## Internal architecture of Gowebsock

Gowebsock is built using the standard Go websocket package "github.com/gorilla/websocket". New websocket connections are created in their own goroutine. Reading from a websocket is a blocking operation. When a user types a message, the message must be sent on the sockets. It is not possible for one goroutine to be blocked on reading and at the same time write to the socket. The simplest solution, and the one employed by Gowebsock, is to have one goroutine for reading and one for writing. 

A Gowebsock program has one master goroutine, one goroutine for reading from stdin, one reader goroutine for each open websocket, and one writer goroutine for each open websocket. The master goroutine coordinates everything and all communication goes through the master. For the server, there is also a goroutine that listens for new connections.

##### Reader goroutine
A reader goroutine blocks on reading from its websocket. When a message has arrived, the message and the address of the socket is sent to the master goroutine through a global channel. If the websocket connection is closed by the remote end, the reader notices and sends the close signal to the master through the same channel.

##### Stdin goroutine
The stdin goroutine reads from the terminal and sends the message to the master through a dedicated channel.

##### Listener goroutine
The listener, in case of the server, is a special goroutine that accepts new websocket connections and creates a new goroutine for each connection. The new goroutine becomes a writer.

##### Writer goroutine
A writer goroutine is created at connection acceptance. The writer tells the master that it has been created through a global channel. The writer also creates a new channel that it sends to the master. The master will send messages to the writer through this new channel in the future. The master tells the writer to write on its websocket and to terminate itself when the websocket has been closed.

##### Master goroutine
The master coordinates everything. It keeps a map of all open web sockets and channels to the writers. The master blocks on a channel select where it waits for new connections, reader messages, stdin messages, and closing of connections. The server sends stdin messages to all connected websockets.  

## Transport layer security

Gowebsock servers and clients can use encrypted connections using TLS. The client specifies a uri starting with "wss://" instead of "ws://". The server must be started with a certificate and a private key. The gowebsock repository includes an example certificate and key in the directory cert.


## Example with a TLS server and a client

##### server
```
~/go> gowebsock server 9000 src/github.com/morten-krogh/gowebsock/cert/localhost.crt src/github.com/morten-krogh/gowebsock/cert/localhost.key 
The gowebsock server is listening on port 9000 using TLS
New connection: 127.0.0.1:49817
127.0.0.1:49817: Hello TLS server!
Hi TLS client
```

##### client
```
~/go> gowebsock client wss://127.0.0.1:9000
The gowebsock client is connected to wss://127.0.0.1:9000
Hello TLS server!
Server: Hi TLS client
The server closed the connection
```
