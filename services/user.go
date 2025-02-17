package services

import (
	"go.messenger/database"
	"go.messenger/models"
	"gorm.io/gorm"
)

func GetUser(id int) (models.User, error) {
	var user models.User
	err := database.DB.Db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, err
		}
		return models.User{}, err
	}
	return user, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User

	result := database.DB.Db.Select("id", "name", "email", "avatar").Where("id > 1").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
