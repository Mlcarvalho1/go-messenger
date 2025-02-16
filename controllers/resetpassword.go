package controllers

import (
	"log"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

type PasswordResetRequest struct {
	Email string `json:"email"`
}

var FireAuth *auth.Client

func PasswordReset(c *fiber.Ctx) error {
	var request PasswordResetRequest

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	authClient := FireAuth // Reuse the initialized Firebase Auth client

	link, err := authClient.PasswordResetLink(c.Context(), request.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate password reset link")
	}

	log.Printf("Password reset link: %s", link)

	return c.JSON(fiber.Map{
		"message": "Password reset link sent",
	})
}
