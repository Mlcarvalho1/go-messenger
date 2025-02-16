package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model

	ID         uint           `json:"id" gorm:"primaryKey"`
	UserID     uint           `json:"user_id" gorm:"not null"`
	User       User           `json:"user" gorm:"foreignKey:UserID"`
	ReceiverID uint           `json:"receiver_id" gorm:"not null"`
	Receiver   User           `json:"receiver" gorm:"foreignKey:ReceiverID"`
	Messages   datatypes.JSON `json:"messages" gorm:"type:jsonb"`
	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"not null"`
}
