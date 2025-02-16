package controllers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"go.messenger/database"
	"go.messenger/models"
	"go.messenger/routes"
)

func setupApp() *fiber.App {
	database.ConnectDb()
	fireAuth := database.InitFirebaseAuth()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/users", func(c *fiber.Ctx) error {
		return controllers.GetUsers(c) 
	})
	return app
}

func testContacts(t *testing.T){
	app := setupApp()
	t.Run("app with contacts", func(t *testing.T){
		email := "testuser" + strconv.FormatInt(time.Now().UnixNano(), 10) + "@test.com"
        user := models.User{Name: "Testador", Email: email}
        result := database.DB.Db.Create(&user)
		if result.Error != nil{
			t.Fatalf("Failed to create user: %v", result.Error)
		}
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
        resp, _ := app.Test(req)

		var contacts []models.User
        err := json.NewDecoder(resp.Body).Decode(&contacts)
        if err != nil {
            t.Fatalf("Failed to decode response body: %v", err)
        }

		assert.NotEmpty(t, contacts)
        assert.Equal(t, user.ID, contacts[0].ID)
	})

	t.Run("app without contacts", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
        resp, _ := app.Test(req)
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        assert.Empty(t, resp)
	})
}

