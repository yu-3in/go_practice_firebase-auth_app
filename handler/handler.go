package handler

import "github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/repository"

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}
