package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	UserID      uint
	User        User
	Title       string `gorm:"type:varchar(40)"`
	IsPublished bool   `gorm:"type:boolean"`
	IsPrivate   bool   `gorm:"type:boolean"`
	PrivateKey  string `gorm:"type:text"`
}

type SelectedQuiz struct {
	ID          uint
	Title       string
	IsPublished bool
	IsPrivate   bool
}
