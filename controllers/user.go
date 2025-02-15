package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.messenger/services"
)

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid user ID: %v", err.Error()),
		})
	}

	user, err := services.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

type UserUpdates struct {
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

func UpdateUser(c *fiber.Ctx) error {

	var updates UserUpdates

	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if updates.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name cannot be empty"})
	}

	//id, err := strconv.Atoi(c.Params("id"))
	//if err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	//}

	user, err := services.UpdateUser(updates.Name, updates.Photo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
