package routes

import (
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	user := api.Group("/users")

	user.Get("/", controllers.GetUser)
	user.Patch("/", controllers.UpdateUser)

	user.Get("/", controllers.GetUsers)
}
