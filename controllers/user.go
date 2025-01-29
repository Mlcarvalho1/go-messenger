package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.messenger/services"
	"go.messenger/database"
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


func CreateUser(c *fiber.Ctx) error {
	client := database.InitFirebaseAuth()

	user, err := services.CreateUser(c, client)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"uid":     user.UID,
		"name":    user.DisplayName,
		"email":   user.Email,
		"photo":   user.PhotoURL,
	})
}