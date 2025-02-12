package services

import (
	"context"
	"errors"

	"firebase.google.com/go/auth"
	"go.messenger/database"
	"go.messenger/models"
)

type SignupPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhotoURL string `json:"photo_url"`
}

// CreateUser envia os dados para o Firebase e salva no banco de dados
func CreateUser(payload SignupPayload) (*models.User, error) {
	if payload.Email == "" || payload.Password == "" || payload.Name == "" {
		return nil, errors.New("name, email, and password are required")
	}

	authClient := database.InitFirebaseAuth()
	if authClient == nil {
		return nil, errors.New("failed to initialize Firebase auth client")
	}

	params := (&auth.UserToCreate{}).
		Email(payload.Email).
		EmailVerified(false).
		Password(payload.Password).
		DisplayName(payload.Name).
		PhotoURL(payload.PhotoURL).
		Disabled(false)

	userRecord, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		if err.Error() == "auth/email-already-exists" {
			return nil, errors.New("email already exists")
		}
		return nil, errors.New(err.Error())
	}

	user := &models.User{
		Name:      payload.Name,
		Email:     payload.Email,
		Avatar:    payload.PhotoURL,
		FireToken: userRecord.UID,
	}

	// Salva o usuário no banco de dados
	db := database.DB.Db
	if db == nil {
		return nil, errors.New("failed to connect to database")
	}

	if err := db.Create(user).Error; err != nil {
		return nil, errors.New("failed to create user in database")
	}

	return user, nil
}
