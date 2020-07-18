// Package: a simple WebSocket server based on the Gorilla toolkit for Go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client.html")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade the client's connection from HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// infinite loop listening for inbound messages sent to the WebSocket
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// print the inbound message to the server console
		fmt.Printf("%s said | %s\n", conn.RemoteAddr(), string(msg))

		// echo the message back to the client
		if err = conn.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func configRoutes() {
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/ws", wsHandler)
}

func main() {
	configRoutes()
	fmt.Println("Server running at http://localhost:5309")
	log.Fatal(http.ListenAndServe(":5309", nil))
}
