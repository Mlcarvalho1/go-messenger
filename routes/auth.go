package routes

import (
	"firebase.google.com/go/auth"
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router, fireAuth *auth.Client) {
	user := api.Group("/auth")

	user.Post("/", controllers.Login)
	user.Post("/sign-up", func(c *fiber.Ctx) error {
		return controllers.Signup(c, fireAuth) // Pass fireAuth to the controller
	})
}
