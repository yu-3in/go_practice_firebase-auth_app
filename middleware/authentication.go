package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		app, err := util.InitFirebaseApp(c)
		if err != nil {
			return err
		}
		client, err := app.Auth(c.Request().Context())
		if err != nil {
			return err
		}
		token, err := util.VerifyIDToken(c, client)
		if err != nil {
			return err
		}
		fbUser, err := client.GetUser(c.Request().Context(), token.UID)
		if err != nil {
			return err
		}

		c.Set("userID", fbUser.UID)

		return next(c)
	}
}
