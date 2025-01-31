package services

import (
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

