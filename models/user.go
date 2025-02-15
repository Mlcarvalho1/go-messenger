package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	Avatar    string `json:"avatar"`
	FireToken string `json:"fire_token"`
}