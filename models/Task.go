package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null;unique_index"`
	Description string `gorm:"not null;unique_index"`
	Done        bool   `gorm:"default:false"`
	UserID      uint
}
