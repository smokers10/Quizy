package models

import (
	"quizy/libs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25) NOT NULL"`
	Email    string `gorm:"type:varchar(30) NOT NULL unique"`
	Password string `gorm:"type:text NOT NULL"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = libs.Hashing(u.Password)
	return
}
