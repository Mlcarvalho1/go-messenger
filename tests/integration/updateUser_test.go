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
	"go.messenger/middlewares"
	"go.messenger/models"
)

func setupApp3() *fiber.App {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	database.ConnectDb()
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/auth", controllers.Login)
	app.Patch("/user", middlewares.FakeFirebaseAuth("fake-firebase-id-for-test20"), controllers.UpdateUser)
	return app
}

func TestUpdateUser(t *testing.T) {
	app := setupApp3()

	t.Run("valid user update", func(t *testing.T) {
		// Cria o usuário no banco de dados real (sem transação, para ser visível)
		user := models.User{
			FireToken: "fake-firebase-id-for-test20",
			Email:     "teste20@gmail.com",
			Name:      "Old Name",
			Avatar:    "https://example.com/old-photo.jpg",
		}
		database.DB.Db.Create(&user)
		defer database.DB.Db.Delete(&user) // Limpeza

		updates := controllers.UserUpdates{
			Name:  "New Name",
			Photo: "https://example.com/new-photo1.jpg",
		}
		body, _ := json.Marshal(updates)

		req := httptest.NewRequest(http.MethodPatch, "/user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer fake-firebase-id-for-test20")

		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var updatedUser models.User
		database.DB.Db.First(&updatedUser, "fire_token = ?", "fake-firebase-id-for-test20")

		assert.Equal(t, updates.Name, updatedUser.Name)
		assert.Equal(t, updates.Photo, updatedUser.Avatar)
	})

	t.Run("invalid token", func(t *testing.T) {
		updates := controllers.UserUpdates{
			Name:  "Hacker Name50",
			Photo: "https://example.com/hacker-photo.jpg",
		}
		body, _ := json.Marshal(updates)

		req := httptest.NewRequest(http.MethodPatch, "/user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer invalid-token")

		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("update with no name", func(t *testing.T) {
		user := models.User{
			FireToken: "fake-firebase-id-for-test20",
			Email:     "teste21@gmail.com",
			Name:      "Old Name",
			Avatar:    "https://example.com/old-photo.jpg",
		}
		database.DB.Db.Create(&user)
		defer database.DB.Db.Delete(&user)

		updates := controllers.UserUpdates{
			Name:  "", // Nome vazio
			Photo: "https://example.com/new-photo.jpg",
		}
		body, _ := json.Marshal(updates)

		req := httptest.NewRequest(http.MethodPatch, "/user", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer fake-firebase-id-for-test20")

		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
