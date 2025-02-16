package routes

import (
	"github.com/gofiber/websocket/v2"
	"go.messenger/webSockets"

	"github.com/gofiber/fiber/v2"
)

func ChatsRoutes(api fiber.Router) {
	chats := api.Group("/chats")

	// returns all chats where a user is involved
	chats.GetChatsByUserID("/user/:userId", controllers.GetChatByUserId)
	
	websocketServer := websockets.NewWebSocket()

	chats.Get("/", websocket.New(func(ctx *websocket.Conn) {
		websocketServer.HandleWebSocket(ctx)
	}))

	go websocketServer.HandleMessages()
}