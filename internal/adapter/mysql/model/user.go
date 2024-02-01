package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
