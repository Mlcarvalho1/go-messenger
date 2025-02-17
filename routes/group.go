package routes

import (
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(api fiber.Router) {
	groups := api.Group("/groups")
	
	groups.Post("/", controllers.CreateGroup)
}
