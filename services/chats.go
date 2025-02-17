package services

import (
	"go.messenger/models"
    "go.messenger/database"
)

func GetChatsByUserID(userId int) ([]models.Chat, error) {
	var chats []models.Chat
    result := database.DB.Db.Where("user_id = ? OR receiver_id = ?", userId, userId).Find(&chats)
    if result.Error != nil {
        return nil, result.Error
    }
	return chats, nil
}