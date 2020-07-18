# WebSocket broadcaster

## A simple browser-based multiuser groupchat application using Go, WebSockets and the Gorilla toolkit

### Overview
This is a simple demo of using [WebSockets](https://en.wikipedia.org/wiki/WebSocket) with [Go](https://golang.org/) to support multiuser chat and broadcast messages across a full-duplex communications channel. This is functionality common in chatrooms, most social messaging applications and WhatsApp. The [websocket package](https://www.gorillatoolkit.org/pkg/websocket) makes this super easy. This web server facilities full-duplex communications with clients, sending/receiving messages across a single TCP connection.

It also leverages Go's concurrency features - asynchronous routines via goroutines and channels - to broadcast messages to all users interacting with it, as a proper chatroom/groupchat application would do.

I've always loved the concept of the realtime web and being able to access information-as-it-happens within a browser. In 2008, there were big strides made with AJAX, HTTP long-polling, Comet and other technologies.

### Servers
I've included are [single-user](/src/server.go) and [multiuser](/src/broadcast_server.go) servers. The latter handles broadcasting any messages to all participants.

- Terminal commands to build this project:
 - `go get -u github.com/gorilla/websocket`
 - Start either server with `go run ./src/server.go`, or `go run ./src/broadcast_server.go`


 The broadcast server sends any new messages to all members, across browser instances or tabs. The dev tools show any received messages from the group.

 What's clever is that Firefox's inspector tools let you filter messages by Sent, Received [AND control frames](https://twitter.com/jasonsalas/status/1283908727821594624). AFAIK, Chrome and Edge don't do this.

 !["Filtering control frames in Firefox's inspector"](https://pbs.twimg.com/media/EdFbAeiU0AAU__c?format=jpg&name=360x360)


 ### UI
 The web interface is a _very_ simplified version of [Ed Zynda's demo](https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets), where he uses a more complex data structure in a proper chatroom, passing JSON. It's a great lesson, so check out his excellent work!

 My demo is more to prove the robustness of Go's concurrency facilities, not exhibit my interface-building prowess (because there isn't much). I'll probably have an update where I represent each user more clearly instead of one long batch of notes.


 ### Playing in traffic
Being able to watch the messages flow back-and-forth in realtime in your favorite browser's network analyzer is the most fun part. Try using it from multiple tabs and/or browsers and see how the messages are logged.

!["Messages being sent and received across the WebSocket in realtime"](https://pbs.twimg.com/media/EdMuB7SUcAQSmcg?format=jpg)
