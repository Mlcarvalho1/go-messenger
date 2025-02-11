package middlewares

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

type contextKey string

const fireAppKey contextKey = "fireApp"

func FirebaseMiddleware(fireApp *firebase.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.Context(), fireAppKey, fireApp)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

func GetFirebaseApp(c *fiber.Ctx) (*firebase.App, error) {
	ctx := c.UserContext()

	fireApp, ok := ctx.Value(fireAppKey).(*firebase.App)
	if !ok {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Firebase app not found in context")
	}

	return fireApp, nil
}
