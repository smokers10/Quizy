package models

import (
	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	UserID uint
	QuizID uint
}
