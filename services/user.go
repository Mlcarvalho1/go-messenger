package services

import (
	"go.messenger/models"
)

func GetUser() (models.User, error) {
	user := models.User{
		ID:   1,
		Name: "John Doe",
	}
	return user, nil
}
