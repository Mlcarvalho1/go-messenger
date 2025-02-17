package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"` // Hidden from JSON
	Avatar    string `json:"avatar,omitempty"`
	FireToken string `json:"fire_token,omitempty"`

	// Override gorm.Model fields to exclude them from JSON response
	CreatedAt string `json:"created_at,omitempty" gorm:"-"`
	UpdatedAt string `json:"updated_at,omitempty" gorm:"-"`
	DeletedAt string `json:"deleted_at,omitempty" gorm:"-"`
}