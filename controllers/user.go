package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.messenger/services"
)

func GetUser(c *fiber.Ctx) error {
	// id, err := strconv.Atoi(c.Params("id"))
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	// }

	user, err := services.GetUser()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
