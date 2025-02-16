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
	"go.messenger/database"
)

func TestResetPassword(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	database.ConnectDb()
	fireAuth := database.InitFirebaseAuth()
	if fireAuth == nil {
		t.Fatalf("Failed to initialize Firebase Auth client")
	}
	// Create a new Fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/resetpassword", controllers.PasswordReset)

	t.Run("valid submission", func(t *testing.T) {
		resetPayload := map[string]string{
			"email": "vma3@cin.ufpe.br",
		}
		body, _ := json.Marshal(resetPayload)

		req := httptest.NewRequest(http.MethodPost, "/resetpassword", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotEmpty(t, bodyBytes) // Verifica se o corpo da resposta não está vazio

	})

	t.Run("missing email", func(t *testing.T) {
		resetPayload := map[string]string{
			"email": "",
		}
		body, _ := json.Marshal(resetPayload)

		req := httptest.NewRequest(http.MethodPost, "/resetpassword", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.NotEmpty(t, bodyBytes) // Verifica se o corpo da resposta não está vazio
	})

	t.Run("incorrect email", func(t *testing.T) {
		resetPayload := map[string]string{
			"email": "joaovictor",
		}
		body, _ := json.Marshal(resetPayload)

		req := httptest.NewRequest(http.MethodPost, "/resetpassword", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.NotEmpty(t, bodyBytes) // Verifica se o corpo da resposta não está vazio
	})
}
