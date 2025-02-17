package models

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model

	ID           uint          `json:"id" gorm:"primaryKey"`
	Name         string        `json:"name" gorm:"not null"`
	Description  string        `json:"description"`
	GroupMembers []GroupMember `json:"group_members" gorm:"foreignKey:GroupID"`

	CreatedAt time.Time `json:"created_at,omitempty" gorm:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"-"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"-"`
}
