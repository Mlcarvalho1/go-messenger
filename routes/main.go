package routes

import (
	"firebase.google.com/go/auth"
	"go.messenger/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, fireAuth *auth.Client) {

	api := app.Group("/", logger.New())

	AuthRoutes(api, fireAuth)

	app.Get("/debug", func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		return c.SendString("Authorization Header: " + authHeader)
	})
	// Rotas abaixo deste middleware precisam de autenticação
	api.Use(middlewares.FirebaseAuthMiddleware(fireAuth))

	ChatsRoutes(api)
	UserRoutes(api)
}
