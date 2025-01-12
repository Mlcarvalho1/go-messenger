package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FirebaseSignInResponse struct {
	IDToken string `json:"idToken"`
}

func Login(c *fiber.Ctx) error {
	var payload LoginRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if payload.Email == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email and password are required"})
	}

	firebaseApiKey, fileExi := os.LookupEnv("FIREBASE_API_KEY")

	if !fileExi {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Firebase API key not found"})
	}

	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", firebaseApiKey)
	requestBody, _ := json.Marshal(map[string]string{
		"email":             payload.Email,
		"password":          payload.Password,
		"returnSecureToken": "true",
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to authenticate with Firebase"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	var firebaseResponse FirebaseSignInResponse
	if err := json.NewDecoder(resp.Body).Decode(&firebaseResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse Firebase response"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": firebaseResponse.IDToken,
	})
}
