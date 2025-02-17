package models

import (
	"time"

	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model

	ID        uint     	`json:"id" gorm:"primaryKey"`
	GroupID   uint      `json:"group_id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	Group     Group     `json:"-" gorm:"foreignKey:GroupID"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`

	CreatedAt time.Time `json:"created_at,omitempty" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"-"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"-"`
}