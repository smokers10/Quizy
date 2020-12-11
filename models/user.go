package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25) NOT NULL"`
	Email    string `gorm:"type:varchar(30) NOT NULL unique"`
	Password string `gorm:"type:text NOT NULL"`
}
