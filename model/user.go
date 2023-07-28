package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique;email"`
	Password string `gorm:"not null;min=8"`
	Point    int    `gorm:"not null;default:0"`
}
