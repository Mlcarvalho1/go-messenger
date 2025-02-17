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

	"go.messenger/controllers"
	"go.messenger/database"
	"go.messenger/models"
)

func setupApp() *fiber.App {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

    database.ConnectDb()
    app := fiber.New()
    app.Use(logger.New())
    app.Get("/users", func(c *fiber.Ctx) error {
        return controllers.GetUsers(c)
    })
    return app
}

func TestContacts(t *testing.T) {
    app := setupApp()

    t.Run("app with contacts", func(t *testing.T) {
        // Create a user in the database for testing
        email := "testuser" + strconv.FormatInt(time.Now().UnixNano(), 10) + "@test.com"
        user := models.User{Name: "Testador", Email: email}
        result := database.DB.Db.Create(&user)
        if result.Error != nil {
            t.Fatalf("Failed to create user: %v", result.Error)
        }

        req := httptest.NewRequest(http.MethodGet, "/users", nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var contacts []models.User
        err := json.NewDecoder(resp.Body).Decode(&contacts)
        if err != nil {
            t.Fatalf("Failed to decode response body: %v", err)
        }

        assert.NotEmpty(t, contacts)
        assert.Equal(t, user.ID, contacts[0].ID)
    })

    t.Run("app without contacts", func(t *testing.T) {
        // Ensure the chat table is empty
        result := database.DB.Db.Exec("DELETE FROM chats")
        if result.Error != nil {
            t.Fatalf("Failed to delete chats: %v", result.Error)
        }

		// Ensure the group_members table is empty
		result = database.DB.Db.Exec("DELETE FROM group_members")
		if result.Error != nil {
			t.Fatalf("Failed to delete group_members: %v", result.Error)
		}

        // Ensure the users table is empty
        result = database.DB.Db.Exec("DELETE FROM users")
        if result.Error != nil {
            t.Fatalf("Failed to delete users: %v", result.Error)
        }

        req := httptest.NewRequest(http.MethodGet, "/users", nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var contacts []models.User
        err := json.NewDecoder(resp.Body).Decode(&contacts)
        if err != nil {
            t.Fatalf("Failed to decode response body: %v", err)
        }

        assert.Empty(t, contacts)
    })
}