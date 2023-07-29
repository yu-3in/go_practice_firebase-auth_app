package util

import (
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func InitFirebaseApp(c echo.Context) (*firebase.App, error) {
	app, err := firebase.NewApp(c.Request().Context(), nil, option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON"))))
	if err != nil {
		return nil, err
	}
	return app, nil
}

func VerifyIDToken(c echo.Context, client *auth.Client) (*auth.Token, error) {
	tokenStr := c.Request().Header.Get("Authorization")
	if tokenStr == "" {
		return nil, echo.ErrUnauthorized
	}
	tokenStr = tokenStr[len("Bearer "):]
	token, err := client.VerifyIDToken(c.Request().Context(), tokenStr)
	if err != nil {
		return nil, err
	}
	return token, nil
}
