package model

type User struct {
	StringIDModel
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique;email"`
	Point int    `gorm:"not null;default:0"`
}
