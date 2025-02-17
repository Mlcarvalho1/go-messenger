package services

import (
	"go.messenger/database"
	"go.messenger/models"
	//"go.messenger/database"
)

func GetUser(id int) (models.User, error) {
	var user models.User

	database.DB.Db.Select("id", "name", "email", "avatar").Where("id = ?", id).First(&user)

	return user, nil
}

func GetUsers() ([]models.User, error){
	var users []models.User

    result := database.DB.Db.Select("id", "name", "email", "avatar").Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }

    return users, nil
}