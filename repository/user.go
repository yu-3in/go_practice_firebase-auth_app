package repository

import (
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func (r *Repository) CreateUser(user *model.User) error {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUser(userID uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateUser(user *model.User) error {
	return r.db.Save(&user).Error
}

func (r *Repository) DeleteUser(userID uint) error {
	return r.db.Delete(&model.User{}, userID).Error
}
