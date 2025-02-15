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
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/auth/sign-up", controllers.Signup)
	return app
}

func TestSignupSuccessWithPhoto(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	app := setupApp()

	loginPayload := map[string]string{
		"email":    "teste@teste.com",
		"password": "12345678",
	}
	
	body, _ := json.Marshal(loginPayload)

	req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// var body map[string]interface{}

	// t.Logf("Response body: %v", body) // Log para depuração
	
	// // Verifica se "user" existe antes de acessar seus campos
	// userData, ok := body["user"].(map[string]interface{})
	// if !ok {
	// 	t.Fatalf("Expected 'user' key in response, but got: %v", body)
	// }
	// json.NewDecoder(resp.Body).Decode(&body)

	// assert.Equal(t, "User created successfully", body["message"])
	// assert.Equal(t, "John Doe", body["user"].(map[string]interface{})["name"])
	// assert.Equal(t, "https://example.com/photo.jpg", body["user"].(map[string]interface{})["photo"])

	// Cleanup: Remover usuário do Firebase e BD
	// cleanupTestUser(userData["email"].(string))
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
