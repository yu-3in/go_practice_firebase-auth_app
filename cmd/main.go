package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/server"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CLIENT_URL")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin},
	}))
	e.Validator = util.NewValidator()

	db, err := util.InitDB()
	if err != nil {
		e.Logger.Fatal("Error connecting to database:", err)
	}
	server.Routing(e, db)

	serverPORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + serverPORT))
}
