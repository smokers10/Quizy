package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	QuizID      uint   `json:"quiz_id,omitempty"`
	Quiz        Quiz   `json:"quiz,omitmepty"`
	Question    string `gorm:"type:text"`
	OptionA     string `gorm:"type:text"`
	OptionB     string `gorm:"type:text"`
	OptionC     string `gorm:"type:text"`
	OptionD     string `gorm:"type:text"`
	RightAnswer string `gorm:"type:varchar(1)"`
}
