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
	// e.POST("/login", h.login)
	auth.GET("/user", h.GetUser)
	// e.PUT("/users", h.updateUser)
	// e.DELETE("/users/:id", h.deleteUser)
	// e.POST("/login", h.login)
	// e.GET("/me", h.me)
}
