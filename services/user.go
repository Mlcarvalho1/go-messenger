package services

import (
	"go.messenger/database"
	"go.messenger/models"
)

func GetUser(id int) (models.User, error) {
	var user models.User

	database.DB.Db.Select("id", "name", "email", "avatar").Where("id = ?", id).First(&user)

	return user, nil
}
