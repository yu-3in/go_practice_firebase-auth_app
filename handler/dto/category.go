package dto

import "github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateCategoryRequestToCategoryModel(r *CreateCategoryRequest) *model.Category {
	return &model.Category{
		Name: r.Name,
	}
}

func UpdateCategoryRequestToCategoryModel(r *UpdateCategoryRequest, category *model.Category) {
	category.Name = r.Name
}

func CategoryModelToCategoryResponse(m *model.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:   m.ID,
		Name: m.Name,
	}
}

func CategoryModelsToCategoryResponses(ms []*model.Category) []*CategoryResponse {
	responses := make([]*CategoryResponse, len(ms))
	for i, m := range ms {
		responses[i] = CategoryModelToCategoryResponse(m)
	}
	return responses
}
