package services

import (
	"go.messenger/database"
	"go.messenger/models"
)

func GetUser(id int) (models.User, error) {
	var user models.User

	database.DB.Db.First(&user, id)

	return user, nil
}
