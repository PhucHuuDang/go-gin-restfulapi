package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name      string      `json:"name"`
	Email     string      `json:"email" gorm:"unique"`
	Age       int         `json:"age"`
	AvatarURL *string     `json:"avatar_url"`
	Posts     []PostModel `json:"posts" gorm:"foreignKey:UserID"`
}
