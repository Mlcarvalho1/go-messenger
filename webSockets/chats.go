package websockets

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gofiber/websocket/v2"
	"go.messenger/database"
	"go.messenger/models"
	"gorm.io/gorm"
)

type Message struct {
	Text       string       `json:"text"`
	SenderId   uint         `json:"senderId"`
	ReceiverId uint         `json:"receiverId"`
	Chat       *models.Chat `json:"-"`
	Time       time.Time    `json:"time"`
}

type WebSocketServer struct {
	clients    map[uint]*websocket.Conn
	mu         sync.Mutex
	broadcast  chan *Message
	AuthClient *auth.Client
}

func NewWebSocket(authClient *auth.Client) *WebSocketServer {
	return &WebSocketServer{
		clients:    make(map[uint]*websocket.Conn),
		broadcast:  make(chan *Message),
		AuthClient: authClient,
	}
}

func (s *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {
	// Recupera o token do query param
	tokenString := ctx.Query("token")
	if tokenString == "" {
		log.Println("Token not provided")
		ctx.Close()
		return
	}

	// Valida o token com Firebase
	token, err := s.AuthClient.VerifyIDToken(context.Background(), tokenString)
	if err != nil {
		log.Println("Invalid token:", err)
		ctx.Close()
		return
	}

	firebaseId := token.UID

	var user models.User
	result := database.DB.Db.First(&user, "fire_token = ?", firebaseId)
	if result.Error != nil {
		log.Println("Error fetching user from DB:", result.Error)
		ctx.Close()
		return
	}

	// Agora sim: registrando a conexão do usuário autenticado
	s.mu.Lock()
	s.clients[user.ID] = ctx
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, user.ID)
		s.mu.Unlock()
		ctx.Close()
	}()

	log.Printf("User %s connected via WebSocket", user.Name)

	// Loop de leitura de mensagens
	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		// Parse the received message
		var message Message
		var chat models.Chat

		message.SenderId = user.ID

		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error Unmarshalling:", err)
			continue
		}

		result := database.DB.Db.Where(
			"(user_id = ? AND receiver_id = ?) OR (user_id = ? AND receiver_id = ?)",
			user.ID, message.ReceiverId, message.ReceiverId, user.ID,
		).First(&chat)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Create a new chat if it doesn't exist
				chat = models.Chat{
					UserID:     user.ID,
					ReceiverID: message.ReceiverId,
					CreatedAt:  time.Now(),
				}

				if createErr := database.DB.Db.Create(&chat).Error; createErr != nil {
					log.Println("Error creating new chat:", createErr)
					continue
				}
			} else {
				log.Println("Database Error:", result.Error)
				continue
			}
		}

		message.Chat = &chat
		// Send the message to the intended recipient
		s.broadcast <- &message
	}
}

func (s *WebSocketServer) HandleMessages() {
	for {
		msg := <-s.broadcast

		msg.Time = time.Now()
		// Convert message to JSON
		messageJSON, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error marshalling message: %v", err)
			continue
		}

		var messages []interface{}

		if len(msg.Chat.Messages) > 0 {
			if err := json.Unmarshal(msg.Chat.Messages, &messages); err != nil {
				log.Printf("Error unmarshalling chat messages: %v", err)
				continue
			}
		}

		messages = append(messages, msg)

		updatedMessages, err := json.Marshal(messages)
		if err != nil {
			log.Printf("Error marshalling updated messages: %v", err)
			continue
		}

		if err := database.DB.Db.Model(&models.Chat{}).
			Where("id = ?", msg.Chat.ID).
			Update("messages", updatedMessages).Error; err != nil {
			log.Printf("Error updating chat messages: %v", err)
			continue
		}

		// Find the recipient and send the message
		s.mu.Lock()
		receiverConn, exists := s.clients[msg.ReceiverId]
		s.mu.Unlock()

		if exists {
			err := receiverConn.WriteMessage(websocket.TextMessage, messageJSON)
			if err != nil {
				log.Printf("Write Error: %v ", err)
				s.mu.Lock()
				receiverConn.Close()
				delete(s.clients, msg.ReceiverId)
				s.mu.Unlock()
			}

		} else {
			log.Printf("User %d is not connected", msg.ReceiverId)
		}
	}
}
