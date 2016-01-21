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


