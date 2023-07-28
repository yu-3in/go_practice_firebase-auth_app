package dto

import (
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Point int    `json:"point"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	gorm.Model
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Point    int               `json:"point"`
	Projects []ProjectResponse `json:"projects"`
}

func CreateUserRequestToUserModel(r *CreateUserRequest) *model.User {
	return &model.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

func UpdateUserRequestToUserModel(r *UpdateUserRequest, user *model.User) {
	user.Name = r.Name
	user.Email = r.Email
	user.Point = r.Point
}

func LoginRequestToUserModel(r *LoginRequest) *model.User {
	return &model.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

func UserModelToUserResponse(m *model.User) *UserResponse {
	return &UserResponse{
		Model: m.Model,
		Name:  m.Name,
		Email: m.Email,
		Point: m.Point,
	}
}

func UserModelToUserResponses(m []*model.User) []*UserResponse {
	responses := make([]*UserResponse, len(m))
	for i, user := range m {
		responses[i] = UserModelToUserResponse(user)
	}
	return responses
}
