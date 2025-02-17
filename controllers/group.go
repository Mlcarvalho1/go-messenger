package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.messenger/services"
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