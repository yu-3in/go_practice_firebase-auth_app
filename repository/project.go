package repository

import (
	"log"

	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *Repository) CreateProject(project *model.Project) error {
	if err := r.SetProjectMembers(project); err != nil {
		return err
	}
	log.Println(project)
	return r.db.Create(project).Error
}

func (r *Repository) GetProject(projectID uint) (*model.Project, error) {
	var project model.Project
	if err := r.db.Preload(clause.Associations).First(&project, projectID).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *Repository) GetProjects() ([]*model.Project, error) {
	var projects []*model.Project
	if err := r.db.Preload(clause.Associations).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) GetOwnerProjects(ownerID uint) ([]*model.Project, error) {
	var projects []*model.Project
	if err := r.db.Preload(clause.Associations).Where("owner_id = ?", ownerID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) GetMemberProjects(userID uint) ([]*model.Project, error) {
	var projects []*model.Project
	if err := r.db.Preload(clause.Associations).Joins("JOIN project_members ON projects.id = project_members.project_id").Where("project_members.user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) UpdateProject(project *model.Project) error {
	if err := r.SetProjectMembers(project); err != nil {
		return err
	}
	return r.db.Save(project).Error
}

func (r *Repository) DeleteProject(projectID uint) error {
	return r.db.Delete(&model.Project{}, projectID).Error
}

func (r *Repository) SetProjectMembers(project *model.Project) error {
	members := make([]model.User, len(project.MemberIDs))
	for i, memberID := range project.MemberIDs {
		members[i] = model.User{Model: gorm.Model{ID: memberID}}
	}
	project.Members = members
	return nil
}
