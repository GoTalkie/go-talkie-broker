package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// TODO: make room and there add connections
var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{}

func (app *Config) Chat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error during upgradin connection to WebSocket: ", err)
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("Client connected")
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(clients, conn)
			break
		}

		// Broadcast message to all connected clients
		for client := range clients {
			err = client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error broadcasting message to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}

}
