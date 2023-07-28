package util

import (
	"fmt"
	"os"

	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	MigrateDB(db)
	return db, nil
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Project{},
	)
}
