package repository

import (
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/model"
	"gorm.io/gorm/clause"
)

func (r *Repository) CreateCategory(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *Repository) GetCategory(categoryID uint) (*model.Category, error) {
	var category model.Category
	if err := r.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *Repository) GetCategories() ([]*model.Category, error) {
	var categories []*model.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *Repository) GetCategoryProjects(categoryID uint) ([]*model.Project, error) {
	var projects []*model.Project
	if err := r.db.Preload(clause.Associations).Where("category_id = ?", categoryID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Repository) UpdateCategory(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *Repository) DeleteCategory(categoryID uint) error {
	return r.db.Delete(&model.Category{}, categoryID).Error
}
