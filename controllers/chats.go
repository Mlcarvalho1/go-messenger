package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.messenger/database"
	"go.messenger/models"
	"go.messenger/services"
)

func GetCurrentUserChats(c *fiber.Ctx) error {
	firebaseId := c.Locals("firebaseId")

	var user models.User

	result := database.DB.Db.First(&user, "fire_token = ?", firebaseId)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	chats, err := services.GetChatsByUserID(int(user.ID))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chats not found"})
	}

	return c.Status(fiber.StatusOK).JSON(chats)
}
