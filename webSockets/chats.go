package websockets

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
	"go.messenger/database"
	"go.messenger/models"
)

type Message struct {
	Text       string `json:"text"`
	SenderId   uint   `json:"senderId"`
	ReceiverId uint   `json:"receiverId"`
}

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *Message
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
	}
}

func (s *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {
	// Register a new Client
	s.clients[ctx] = true
	defer func() {
		delete(s.clients, ctx)
		ctx.Close()
	}()

	firebaseId := ctx.Locals("firebaseId")

	log.Println("Firebase ID:", firebaseId)

	var user models.User

	result := database.DB.Db.First(&user, "fire_token = ?", firebaseId)

	if result.Error != nil {
		log.Fatalf("Error fetching user")
	}

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		// send the message to the broadcast channel
		var message Message
		message.SenderId = user.ID

		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error Unmarshalling:", err)
		}

		s.broadcast <- &message
	}
}

func (s *WebSocketServer) HandleMessages() {
	for {
		msg := <-s.broadcast

		// Send the message to all Clients
		messageJSON, err := json.Marshal(msg)

		log.Println("Message Received:", string(messageJSON))

		if err != nil {
			log.Printf("Error marshalling message: %v", err)
			continue
		}

		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, messageJSON)
			if err != nil {
				log.Printf("Write  Error: %v ", err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}
