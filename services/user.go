package services

import (
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"go.messenger/models"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhotoURL string `json:"photo_url"`
}

func GetUser() (models.User, error) {
	user := models.User{
		ID:   1,
		Name: "John Doe",
	}
	return user, nil
}

func CreateUser(c *fiber.Ctx, client *auth.Client) (*auth.UserRecord, error) {
	var request RegisterRequest

	if err := c.BodyParser(&request); err != nil {
		return nil, err
	}

	params := (&auth.UserToCreate{}).
		Email(request.Email).
		EmailVerified(false).
		Password(request.Password).
		DisplayName(request.Name).
		PhotoURL(request.PhotoURL).
		Disabled(false)

	ctx := c.Context()
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return u, nil
}
