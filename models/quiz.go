package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	UserID      uint   `json:"user_id,omitempty"`
	User        User   `json:"user,omitempty"`
	QuizName    string `gorm:"type:varchar(60) NOT NULL" json:"quiz_name,omitempty"`
	IsPublished bool   `gorm:"type:boolean default false" json:"is_published,omitempty"`
	IsPrivate   bool   `gorm:"type:boolean default false" json:"is_private,omitempty"`
	PrivateKey  string `gorm:"type:varchar(256)" json:"PrivateKey,omitempty"`
}

type SelectedQuiz struct {
	ID          uint
	QuizName    string
	IsPublished bool
	IsPrivate   bool
}
