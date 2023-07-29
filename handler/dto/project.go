package dto

import (
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
	"gorm.io/gorm"
)

type CreateProjectRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	MemberIDs   []string `json:"memberIds"`
	CategoryID  uint     `json:"categoryId" validate:"required"`
}

type CreateProjectResponse struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OwnerID     string   `json:"ownerId"`
	MemberIDs   []string `json:"memberIds"`
	CategoryID  uint     `json:"categoryId"`
}

type UpdateProjectRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	MemberIDs   []string `json:"memberIds"`
	CategoryID  uint     `json:"categoryId" validate:"required"`
}

type ProjectResponse struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Owner       model.User     `json:"owner"`
	Members     []model.User   `json:"members"`
	Category    model.Category `json:"category"`
}

func CreateProjectRequestToProjectModel(r *CreateProjectRequest) *model.Project {
	return &model.Project{
		Name:        r.Name,
		Description: r.Description,
		MemberIDs:   r.MemberIDs,
		CategoryID:  r.CategoryID,
	}
}

func UpdateProjectRequestToProjectModel(r *UpdateProjectRequest, project *model.Project) {
	project.Name = r.Name
	project.Description = r.Description
	project.MemberIDs = r.MemberIDs
	project.CategoryID = r.CategoryID
}

func ProjectModelToProjectResponse(m *model.Project) *ProjectResponse {
	return &ProjectResponse{
		Model:       m.Model,
		Name:        m.Name,
		Description: m.Description,
		Owner:       m.Owner,
		Members:     m.Members,
		Category:    m.Category,
	}
}

func ProjectModelsToProjectResponses(ms []*model.Project) []*ProjectResponse {
	responses := make([]*ProjectResponse, len(ms))
	for i, m := range ms {
		responses[i] = ProjectModelToProjectResponse(m)
	}
	return responses
}
