package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title" gorm:"notnull"`
	Description string    `json:"description"`
	UserID      uint      `gorm:"not null"`
	CategoryID  uint      `gorm:"not null"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	CreateAt    time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt    time.Time `json:"update_at" gorm:"autoUpdateTime"`
	User        User
	Category    *Category
}
