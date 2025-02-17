package controllers 

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"go.messenger/services"
)

func GetChatsByUserID(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
    userId, err := strconv.Atoi(userIdStr)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

	chats, err := services.GetChatsByUserID(userId)
	
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chats not found"})
	}

	return c.Status(fiber.StatusOK).JSON(chats)
}