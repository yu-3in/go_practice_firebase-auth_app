package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string   `gorm:"not null"`
	Description string   `gorm:"not null"`
	OwnerID     string   `gorm:"not null;size:255"`
	Owner       User     `gorm:"foreignKey:OwnerID"`
	MemberIDs   []string `gorm:"-"`
	Members     []User   `gorm:"many2many:project_members;"`
	CategoryID  uint
	Category    Category
}
