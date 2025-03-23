package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.messenger/controllers"
	websockets "go.messenger/webSockets"
)

func ChatsRoutes(api fiber.Router) {
	chats := api.Group("/chats")

	// returns all chats where a user is involved
	chats.Get("/user", controllers.GetCurrentUserChats)

	websocketServer := websockets.NewWebSocket()

	chats.Get("/", websocket.New(func(ctx *websocket.Conn) {
		websocketServer.HandleWebSocket(ctx)
	}))

	go websocketServer.HandleMessages()

}
