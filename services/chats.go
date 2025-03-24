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

	for i := range chats {
		chat := &chats[i]

		// If receiverID matches the userId, swap User and Receiver
		if chat.ReceiverID == uint(userId) {
			newReceiver := chat.User

			chat.UserID = uint(userId)
			chat.User = chat.Receiver

			chat.ReceiverID = uint(newReceiver.ID)
			chat.Receiver = newReceiver
		} else {
			chat.UserID = uint(userId)
		}
	}

	return chats, nil
}
