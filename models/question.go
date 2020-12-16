package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	QuizID      uint
	Quiz        Quiz
	Question    string `gorm:"type:text"`
	OptionA     string `gorm:"type:text"`
	OptionB     string `gorm:"type:text"`
	OptionC     string `gorm:"type:text"`
	OptionD     string `gorm:"type:text"`
	RightAnswer string `gorm:"type:varchar(1)"`
}

type SelectedQuestion struct {
	ID          uint
	Question    string
	OptionA     string
	OptionB     string
	OptionC     string
	OptionD     string
	RightAnswer string
}

type SelectedQuestionOnTake struct {
	ID       uint
	Question string
	OptionA  string
	OptionB  string
	OptionC  string
	OptionD  string
}

type SelectedQuestionOnCheck struct {
	ID          uint
	RightAnswer string
}
