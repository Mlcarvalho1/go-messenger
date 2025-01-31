package routes

import (
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	user := api.Group("/auth")

	user.Post("/", controllers.Login)
	user.Post("/sign-up", controllers.Signup)
}
