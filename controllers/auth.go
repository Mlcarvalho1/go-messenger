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


func Signup(c *fiber.Ctx) error {
	var payload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		PhotoURL string `json:"photo_url"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validação básica dos campos
	if payload.Email == "" || payload.Password == "" || payload.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name, email and password are required"})
	}

	firebaseApiKey, exists := os.LookupEnv("FIREBASE_API_KEY")
	if !exists {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Firebase API key not found"})
	}

	// URL para criar usuário no Firebase
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s", firebaseApiKey)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"email":             payload.Email,
		"password":          payload.Password,
		"displayName":       payload.Name,
		"photoUrl":          payload.PhotoURL,
		"returnSecureToken": true,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user in Firebase"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Lendo o corpo da resposta de erro do Firebase
		var firebaseError struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&firebaseError); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Mapeando erros comuns do Firebase
		switch firebaseError.Error.Message {
		case "EMAIL_EXISTS":
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already in use"})
		case "INVALID_EMAIL":
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email format"})
		case "WEAK_PASSWORD":
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password is too weak"})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": firebaseError.Error.Message})
		}
	}

	var firebaseResponse struct {
		IDToken      string `json:"idToken"`
		Email        string `json:"email"`
		DisplayName  string `json:"displayName"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		LocalID      string `json:"localId"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&firebaseResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse Firebase response"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": firebaseResponse.IDToken,
		"user": fiber.Map{
			"id":    firebaseResponse.LocalID,
			"name":  firebaseResponse.DisplayName,
			"email": firebaseResponse.Email,
		},
	})
}