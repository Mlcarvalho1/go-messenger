package routes

import (
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(api fiber.Router) {
	groups := api.Group("/groups")

	// returns all group chats where a user is involved
	groups.Get("/user/:userId", controllers.GetGroupChatsByUserID)

	groups.Post("/", controllers.CreateGroup)
}
