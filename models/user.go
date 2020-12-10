package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(30) NOT NULL" json:"name"`
	Email    string `gorm:"type:varchar(40) NOT NULL unique" json:"email"`
	Password string `gorm:"type:varchar(120) NOT NULL" json:"password,omitempty"`
}
