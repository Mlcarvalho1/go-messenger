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

func UpdateUser(name string, photo string) (models.User, error) {
	var user models.User
	user = models.User{
		Name:   name,
		Avatar: photo,
	}

	return user, nil
}
