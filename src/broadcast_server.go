// Multiuser message broadcasting via WebSockets
// using Go and the Gorilla toolkit
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client.html")
	})

	http.HandleFunc("/ws", handleConnections)

	// deliver inbound messages to groupchat participants
	go broadcastMessages()

	log.Println("server running on port 5309")
	err := http.ListenAndServe(":5309", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// add connection as a groupchat participant
	clients[conn] = true

	// infinite loop listening for new messages sent to the WebSocket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, conn)
			break
		}

		// send the message to the broadcast channel
		broadcast <- msg
	}
}

func broadcastMessages() {
	for {
		// extract latest message from the channel
		msg := <-broadcast

		// deliver message to all connected clients
		for client := range clients {
			if err := client.WriteMessage(1, msg); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
