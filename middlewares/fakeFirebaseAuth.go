package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func FakeFirebaseAuth(fakeFirebaseId string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}

		token := authHeader[len("Bearer "):]
		if token != fakeFirebaseId {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("firebaseId", fakeFirebaseId)
		return c.Next()
	}
}
