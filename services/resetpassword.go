package services

import (
	"go.messenger/database"
	"go.messenger/models"
)

func GetEmail(email string) bool {
	var user models.User

	result := database.DB.Db.Where("email = ?", email).First(&user)

	return (result.RowsAffected != 0)
}
