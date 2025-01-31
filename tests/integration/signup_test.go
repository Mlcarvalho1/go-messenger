package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"go.messenger/controllers"
)

type RegisterResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhotoURL string `json:"photo_url"`
}

var app *fiber.App

func setup() {
    app = fiber.New()
    app.Use(logger.New())
    app.Post("/auth/sign-up", controllers.Signup)
}

func TestUserRegistration(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	// Create a new Fiber app
	setup()

	t.Run("successful registration", func(t *testing.T) {
		registrationPayload := map[string]string{
			"name":      "Rafael Alves",
			"email":     "raaffa2@email.com",
			"Password":  "senhaSegura123.",
			"photo_url": "base64encodedimage==",
		}
		body, _ := json.Marshal(registrationPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response RegisterResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

	})

	t.Run("invalid email registration", func(t *testing.T) {
		registrationPayload := map[string]string{
			"name":      "Rafael Alves",
			"email":     "rrafaemail.com",
			"Password":  "senhaSegura123.",
			"photo_url": "base64encodedimage==",
		}
		body, _ := json.Marshal(registrationPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response ErrorResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, "Invalid email format", response.Error)
	})

	t.Run("email already registered", func(t *testing.T) {
		registrationPayload := map[string]string{
			"name":      "Rafael Alves",
			"email":     "example@email.com",
			"Password":  "senhaSegura123.",
			"photo_url": "base64encodedimage==",
		}
		body, _ := json.Marshal(registrationPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response ErrorResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, "Email already in use", response.Error)
	})

    t.Run("weak password registration", func(t *testing.T) {
		registrationPayload := map[string]string{
			"name":      "Rafael Alves",
			"email":     "example@email.com",
			"Password":  "1234",
			"photo_url": "base64encodedimage==",
		}
		body, _ := json.Marshal(registrationPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.Equal(t, "WEAK_PASSWORD : Password should be at least 6 characters", response.Error)
	})

	t.Run("registration without profile picture", func(t *testing.T) {
		registrationPayload := map[string]string{
			"name":      "Rafael Alves",
			"email":     "example5@email.com",
			"Password":  "senhaSegura123.",
			"photo_url": "base64encodedimage==",
		}
		body, _ := json.Marshal(registrationPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response RegisterResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

	})
}
