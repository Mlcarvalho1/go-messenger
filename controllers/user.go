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

func GetUsers(c *fiber.Ctx) error{
	contacts, err := services.GetUsers()
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contacts not found"})
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}