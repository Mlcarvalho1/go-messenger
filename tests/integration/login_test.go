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

type AuthResponse struct {
	Token string `json:"token"`
}

func TestLogin(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Create a new Fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/auth", controllers.Login)

	t.Run("valid account", func(t *testing.T) {
		loginPayload := map[string]string{
			"email":    "teste@teste.com",
			"password": "12345678",
		}
		body, _ := json.Marshal(loginPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response AuthResponse
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		assert.NotEmpty(t, response.Token)
		assert.IsType(t, "", response.Token)
	})

	t.Run("invalid account", func(t *testing.T) {
		loginPayload := map[string]string{
			"email":    "teste2@teste.com",
			"password": "12345678",
		}
		body, _ := json.Marshal(loginPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("missing email or password", func(t *testing.T) {
		loginPayload := map[string]string{
			"password": "",
		}
		body, _ := json.Marshal(loginPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
