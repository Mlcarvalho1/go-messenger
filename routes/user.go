package routes

import (
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	user := api.Group("/user")

	user.Get("/:id", controllers.GetUser)
}
