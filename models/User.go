package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `validate:"required"       gorm:"not null"`
	Lastname  string `validate:"required"       gorm:"not null"`
	Email     string `validate:"required,email" gorm:"not null;unique_index"`
	Tasks     []Task `                                                       json:"tasks"`
}
