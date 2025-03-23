package services

import (
	"go.messenger/database"
	"go.messenger/models"
)

func GetChatsByUserID(userId int) ([]models.Chat, error) {
	var chats []models.Chat
	result := database.DB.Db.
		Preload("User").
		Preload("Receiver").
		Where("user_id = ? OR receiver_id = ?", userId, userId).
		Find(&chats)

	if result.Error != nil {
		return nil, result.Error
	}

	return chats, nil
}
