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
	QuizId     uint
	Answer     string `gorm:"type:varchar(1)"`
	IsRight    bool
}
