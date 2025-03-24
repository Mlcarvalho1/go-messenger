package routes

import (
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.messenger/controllers"
	"go.messenger/middlewares"
	websockets "go.messenger/webSockets"
)

func ChatsRoutes(api fiber.Router, authClient *auth.Client) {
	chats := api.Group("/chats")

	chats.Get("/user", middlewares.FirebaseAuthMiddleware(authClient), controllers.GetCurrentUserChats)

	websocketServer := websockets.NewWebSocket(authClient)

	chats.Get("/", websocket.New(func(ctx *websocket.Conn) {
		websocketServer.HandleWebSocket(ctx)
	}))

	go websocketServer.HandleMessages()
}
