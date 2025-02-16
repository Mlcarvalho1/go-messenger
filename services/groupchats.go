package services

import (
	"go.messenger/models"
    "go.messenger/database"
)

func GetGroupChatsByUserID(userId int) ([]models.GroupChat, error) {
    var chats []models.GroupChat
    result := database.DB.Db.Where("JSON_CONTAINS(member_lists, '[?]', '$')", userId).Find(&chats)
    if result.Error != nil {
        return nil, result.Error
    }
    return chats, nil
}