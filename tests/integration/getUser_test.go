package controllers_test

import (
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

func setupApp2() *fiber.App {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	database.ConnectDb()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/user/:id", controllers.GetUser)
	return app
}

func TestGetUser(t *testing.T) {
	app := setupApp2()

	t.Run("valid user ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/user/125", nil)
		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("invalid user ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/user/invalid", nil)
		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("user not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/user/9999", nil)
		resp, err := app.Test(req, -1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
