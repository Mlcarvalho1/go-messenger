package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.messenger/database"
	"go.messenger/models"
	"go.messenger/services"
)

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
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

	firebaseId := c.Locals("firebaseId")
	fmt.Println("firebaseId:", firebaseId)

	var user models.User

	result := database.DB.Db.First(&user, "fire_token = ?", firebaseId)
	fmt.Println("User found:", user)

	if result.Error != nil {
		fmt.Println("Error finding user:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}

	user.Name = updates.Name
	user.Avatar = updates.Photo

	if err := database.DB.Db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	contacts, err := services.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contacts not found"})
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}
