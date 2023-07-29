package dto

import (
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
)

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Point int    `json:"point"`
}

type UserResponse struct {
	model.StringIDModel
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Point int    `json:"point"`
}

func CreateUserRequestToUserModel(r *CreateUserRequest) *model.User {
	return &model.User{
		Name:  r.Name,
		Email: r.Email,
	}
}

func UpdateUserRequestToUserModel(r *UpdateUserRequest, user *model.User) {
	user.Name = r.Name
	user.Email = r.Email
	user.Point = r.Point
}

func UserModelToUserResponse(m *model.User) *UserResponse {
	return &UserResponse{
		StringIDModel: m.StringIDModel,
		Name:          m.Name,
		Email:         m.Email,
		Point:         m.Point,
	}
}

func UserModelToUserResponses(m []*model.User) []*UserResponse {
	responses := make([]*UserResponse, len(m))
	for i, user := range m {
		responses[i] = UserModelToUserResponse(user)
	}
	return responses
}
