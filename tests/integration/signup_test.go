package controllers_test

import (
	"bytes"
	"encoding/json"

	"log"
	"net/http"
	"net/http/httptest"

	"context"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.messenger/controllers"
	"go.messenger/database"
)

// Configura o ambiente de teste
func setupApp() *fiber.App {
	database.ConnectDb()
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/auth/sign-up", controllers.Signup)
	return app
}

func TestSignupSuccess(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	app := setupApp()
	t.Run("valid account", func(t *testing.T) {
		SignupPayload := map[string]string{
			"email":     "example@teste.com",
			"password":  "12345678fgfd",
			"photo_url": "https://example.com/photo.jpg",
			"name":      "John Doe",
		}
		
		cleanupTestUser(SignupPayload["email"]) // Limpa o usuário antes de criar um novo

		body, _ := json.Marshal(SignupPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("email already exists", func(t *testing.T) {
		SignupPayload := map[string]string{
			"email":     "example@teste.com",
			"password":  "12345678fgfd",
			"photo_url": "https://example.com/photo.jpg",
			"name":      "John Doe",
		}

		// Create user first
		body, _ := json.Marshal(SignupPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// app.Test(req, -1)

		// Try to create the same user again
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("weak password", func(t *testing.T) {
		SignupPayload := map[string]string{
			"email":     "example@teste.com",
			"password":  "123",
			"photo_url": "https://example.com/photo.jpg",
			"name":      "John Doe",
		}

		// cleanupTestUser(SignupPayload["email"]) // Limpa o usuário antes de criar um novo

		body, _ := json.Marshal(SignupPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("invalid email", func(t *testing.T) {
		SignupPayload := map[string]string{
			"email":     "invalid-email",
			"password":  "12345678fgfd",
			"photo_url": "https://example.com/photo.jpg",
			"name":      "John Doe",
		}

		body, _ := json.Marshal(SignupPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("success without photo", func(t *testing.T) {
		SignupPayload := map[string]string{
			"email":    "example@teste.com",
			"password": "12345678fgfd",
			"name":     "John Doe",
			"photo_url": "",
		}

		cleanupTestUser(SignupPayload["email"]) // Limpa o usuário antes de criar um novo

		body, _ := json.Marshal(SignupPayload)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

}

// Função para remover o usuário do Firebase e banco de dados após o teste
func cleanupTestUser(email string) {
	authClient := database.InitFirebaseAuth() // Implemente essa função para pegar o cliente do Firebase
	if authClient != nil {
		user, err := authClient.GetUserByEmail(context.Background(), email)
		if err == nil {
			authClient.DeleteUser(context.Background(), user.UID)
		}
	}

	db := database.DB.Db // Implemente essa função para acessar o banco
	if db != nil {
		db.Exec("DELETE FROM users WHERE email = ?", email)
	}
}
