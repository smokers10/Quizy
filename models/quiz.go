package models

import (
	"quizy/libs"

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

type SelectedQuizEnrollment struct {
	ID          uint
	Title       string
	IsPublished bool
	IsPrivate   bool
	PrivateKey  string
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	if q.PrivateKey != "" {
		q.PrivateKey = libs.Hashing(q.PrivateKey)
	}
	return
}
