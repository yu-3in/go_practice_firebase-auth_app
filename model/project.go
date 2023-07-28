package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	OwnerID     uint
	Owner       User   `gorm:"foreignKey:OwnerID"`
	MemberIDs   []uint `gorm:"-"`
	Members     []User `gorm:"many2many:project_members;"`
	CategoryID  uint
	Category    Category
}
