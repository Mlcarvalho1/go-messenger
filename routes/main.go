package routes

import (
	"context"

	"go.messenger/middlewares"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, fireApp *firebase.App) {
	authClient, err := fireApp.Auth(context.Background())

	if err != nil {
		panic(err)
	}

	api := app.Group("/", logger.New())

	// Rotas abaixo deste middleware utilizam o firebase.App
	api.Use(middlewares.FirebaseMiddleware(fireApp))

	AuthRoutes(api)

	app.Get("/debug", func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		return c.SendString("Authorization Header: " + authHeader)
	})
	// Rotas abaixo deste middleware precisam de autenticação
	api.Use(middlewares.FirebaseAuthMiddleware(authClient))

	UserRoutes(api)
}
