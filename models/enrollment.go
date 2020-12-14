package models

import (
	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	UserID     uint
	QuizID     uint
	IsSubmited bool
	Quiz       Quiz
}

type EnrollmentOnCheck struct {
	ID         uint
	UserID     uint
	QuizID     uint
	IsSubmited bool
}
