package models

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	UserID     uint
	User       User
	QuestionID uint
	Question   Question
	Answer     string
}
