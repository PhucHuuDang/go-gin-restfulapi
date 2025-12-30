package models

import (
	"time"

	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model
	Title     string `json:"title"`
	User      UserModel
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
}
