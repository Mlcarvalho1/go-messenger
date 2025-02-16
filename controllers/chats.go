package controllers 

import (
	"github.com/gofiber/fiber/v2"
	"go.messenger/services"
)

func GetChatsByUserID(c *fiber.Ctx) error {
	userId := c.Params("userId")

	chats, err := services.GetChatsByUserId(userId)
	
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chats not found"})
	}

	return c.Status(fiber.StatusOK).JSON(chats)
}