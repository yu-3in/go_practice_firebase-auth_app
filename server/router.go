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
	e.POST("/login", h.Login)
	auth.GET("/user", h.GetUser)
	auth.PUT("/user", h.UpdateUser)
	auth.DELETE("/user", h.DeleteUser)
}
