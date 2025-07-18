package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	User   User
	Tasks  []Task
}
