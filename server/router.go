package server

import (
	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/handler"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/middleware"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/repository"
	"gorm.io/gorm"
)

func Routing(e *echo.Echo, db *gorm.DB) {
	repo := repository.NewRepository(db)
	h := handler.NewHandler(repo)

	// ミドルウェア
	auth := e.Group("", middleware.AuthenticationMiddleware)

	e.POST("/signup", h.CreateUser)

	auth.GET("/user", h.GetUser)
	auth.PUT("/user", h.UpdateUser)
	auth.DELETE("/user", h.DeleteUser)

	auth.POST("/projects", h.CreateProject)
	auth.GET("/projects", h.GetProjects)
	auth.GET("/projects/member", h.GetMemberProjects)
	auth.GET("/projects/owner", h.GetOwnerProjects)
	auth.GET("/projects/:id", h.GetProject)
	auth.PUT("/projects/:id", h.UpdateProject)
	auth.DELETE("/projects/:id", h.DeleteProject)

	auth.POST("/categories", h.CreateCategory)
	auth.GET("/categories", h.GetCategories)
	auth.GET("/categories/:id/projects", h.GetCategoryProjects)
	auth.GET("/categories/:id", h.GetCategory)
	auth.PUT("/categories/:id", h.UpdateCategory)
	auth.DELETE("/categories/:id", h.DeleteCategory)

}
