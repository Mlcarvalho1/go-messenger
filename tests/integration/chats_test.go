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


func setupAppChats() *fiber.App {
    if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

    // Initialize the database (assuming you have a function to do this)
    database.ConnectDb()

    // Create a new Fiber app
    app := fiber.New()
    app.Use(logger.New())

    // Setup routes
    api := app.Group("/api")
    routes.ChatsRoutes(api)

    return app
}
func TestGetChatsByUserID(t *testing.T) {

    app := setupAppChats()
    t.Run("valid user ID with chats", func(t *testing.T) {
        // Create a user and chats in the database for testing
        email := "testuser" + strconv.FormatInt(time.Now().UnixNano(), 10) + "@test.com"
        user := models.User{Name: "Testador", Email: email}
        result := database.DB.Db.Create(&user)
        if result.Error != nil {
            t.Fatalf("Failed to create user: %v", result.Error)
        }

        chat := models.Chat{UserID: user.ID, ReceiverID: user.ID}
        result = database.DB.Db.Create(&chat)
        if result.Error != nil {
            t.Fatalf("Failed to create chat: %v", result.Error)
        }

        req := httptest.NewRequest(http.MethodGet, "/api/chats/user/"+strconv.Itoa(int(user.ID)), nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var chats []models.Chat
        err := json.NewDecoder(resp.Body).Decode(&chats)
        if err != nil {
            t.Fatalf("Failed to decode response body: %v", err)
        }

        assert.NotEmpty(t, chats)
        assert.Equal(t, user.ID, chats[0].UserID)
    })

    t.Run("valid user ID without chats", func(t *testing.T) {
        // Create a user without chats in the database for testing
        email := "testuser" + strconv.FormatInt(time.Now().UnixNano(), 10) + "@test.com"
        user := models.User{Name: "Testador", Email: email}
        result := database.DB.Db.Create(&user)
        if result.Error != nil {
            t.Fatalf("Failed to create user: %v", result.Error)
        }

        req := httptest.NewRequest(http.MethodGet, "/api/chats/user/"+strconv.Itoa(int(user.ID)), nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var chats []models.Chat
        err := json.NewDecoder(resp.Body).Decode(&chats)
        if err != nil {
            t.Fatalf("Failed to decode response body: %v", err)
        }

        assert.Empty(t, chats)
    })
   
    t.Run("invalid user ID", func(t *testing.T) {
        req := httptest.NewRequest(http.MethodGet, "/api/chats/user/invalid", nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
    })

    t.Run("non-existent user ID", func(t *testing.T) {
        req := httptest.NewRequest(http.MethodGet, "/api/chats/user/999999", nil)
        resp, _ := app.Test(req)

        assert.Equal(t, http.StatusOK, resp.StatusCode)

        var chats []models.Chat
        err := json.NewDecoder(resp.Body).Decode(&chats)
        if err != nil {
            t.Fatalf("Failed to decode response body: %v", err)
        }

        assert.Empty(t, chats)
    })
}