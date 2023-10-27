package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null;unique_index" validate:"required"`
	Description string `gorm:"not null;unique_index"                   validate:"required"`
	Done        bool   `gorm:"default:false"                                               json:"omitempty"`
	UserID      uint   `gorm:"not null"                                validate:"required" json:"user"`
}
