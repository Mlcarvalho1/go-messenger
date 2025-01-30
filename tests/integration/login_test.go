package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stretchr/testify/assert"

	"go.messenger/controllers"
)

type MockFirebaseResponse struct {
	IDToken string `json:"idToken"`
}

func TestLogin(t *testing.T) {
	// Set up environment variable for Firebase API key
	os.Setenv("FIREBASE_API_KEY", "mock-api-key")

	// Create a new Fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/auth", controllers.Login)

	// Mock Firebase response
	mockResponse := MockFirebaseResponse{IDToken: "mocked_token"}
	mockResponseJSON, _ := json.Marshal(mockResponse)

	// Create a test HTTP server to mock Firebase API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponseJSON)
	}))
	defer ts.Close()

	// Replace Firebase API URL with the test server URL
	os.Setenv("FIREBASE_API_KEY", ts.URL)

	// Test cases
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
