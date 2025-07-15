package models

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"notnull"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	CreateAt    time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt    time.Time `json:"update_at" gorm:"autoUpdateTime"`
}
