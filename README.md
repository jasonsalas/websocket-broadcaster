# WebSocket broadcaster

## A simple browser-based multiuser chat application using Go, WebSockets and the Gorilla toolkit

This is a simple demo of using (https://en.wikipedia.org/wiki/WebSocket)[WebSockets] with Go to support multiuser chat and broadcast messages across full-duplex communications channel. The [https://www.gorillatoolkit.org/pkg/websocket](websocket package) makes this super easy. This web server facilities full-duplex communications with clients, sending/receiving messages across a single TCP connection.

It also leverages Go's concurrency features - asynchronous routines and channels - to broadcast messages to all users interacting with it, as a proper chatroom would do.

I've always loved the concept of the realtime web and being able to access information-as-it-happens within a browser. In 2008, there were big strides made with AJAX, HTTP long-polling, Comet and other technologies.

- Terminal commands to build this project on your own:
 - `mkdir src && touch ./src/server.go client.html README.md`
 - `go mod init github.com/jasonsalas/ws-broadcaster`
 - `go get -u github.com/gorilla/websocket`
 - Start the server with `go run ./src/server.go`, or run `go build ./src/server.go && ./server`

 What's clever is that Firefox's inspector tools let you filter messages by Sent, Received [AND control frames](https://twitter.com/jasonsalas/status/1283908727821594624). AFAIK, Chrome and Edge don't do this.

 !["Filtering control frames in Firefox's inspector"](https://pbs.twimg.com/media/EdFbAeiU0AAU__c?format=jpg&name=360x360)
