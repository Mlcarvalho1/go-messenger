package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.messenger/services"
	"strconv"
)

func CreateGroup(ctx *fiber.Ctx) error {
	var req services.CreateGroupRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	group, err := services.CreateGroup(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create group",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(group)
}
func GetGroupChatsByUserID(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    chats, err := services.GetGroupChatsByUserID(userId)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Group chats not found",
        })
    }

    return c.Status(fiber.StatusOK).JSON(chats)
}