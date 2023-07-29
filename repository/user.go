package repository

import (
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
)

func (r *Repository) CreateUser(user *model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUser(userID string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateUser(user *model.User) error {
	return r.db.Save(&user).Error
}

func (r *Repository) DeleteUser(userID string) error {
	return r.db.Delete(&model.User{}, userID).Error
}
