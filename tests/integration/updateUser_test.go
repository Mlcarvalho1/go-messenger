package controllers_test

import (
	"bytes"
	"encoding/json"
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

func setupApp3() *fiber.App {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	database.ConnectDb()
	app := fiber.New()
	app.Use(logger.New())
	app.Put("/user", controllers.UpdateUser) // Ajuste a rota para não usar o ID na URL
	return app
}

func TestUpdateUser(t *testing.T) {
	app := setupApp3()

	// Inserir um usuário de teste no banco de dados
	//testUser := models.User{
	//	Name:      "Old Name",
	//	Avatar:    "https://example.com/old-photo.jpg",
	//	FireToken: "test-fire-token",
	//}
	//database.DB.Db.Create(&testUser)

	t.Run("valid user update", func(t *testing.T) {
		updates := controllers.UserUpdates{
			Name:  "New Name",
			Photo: "https://example.com/new-photo.jpg",
		}
		body, _ := json.Marshal(updates)
		req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("firebaseId", "test-fire-token")
		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("invalid fire token", func(t *testing.T) {
		updates := controllers.UserUpdates{
			Name:  "New Name",
			Photo: "https://example.com/new-photo.jpg",
		}
		body, _ := json.Marshal(updates)
		req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("firebaseId", "invalid-fire-token")
		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("user not found", func(t *testing.T) {
		updates := controllers.UserUpdates{
			Name:  "New Name",
			Photo: "https://example.com/new-photo.jpg",
		}
		body, _ := json.Marshal(updates)
		req := httptest.NewRequest(http.MethodPut, "/user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("firebaseId", "non-existent-fire-token")
		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
