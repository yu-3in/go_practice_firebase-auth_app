package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := util.ParseToken(c)
		if err != nil {
			return echo.ErrUnauthorized
		}
		c.Set("userID", userID)

		return next(c)
	}
}
