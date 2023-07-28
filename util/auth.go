package util

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

var (
	jwtSecretKey = os.Getenv("JWT_SECRET")
	jwtExpiresIn = os.Getenv("JWT_EXPIRATION")
)

func GenerateToken(userID uint) (string, error) {
	jwtExpiresIn, err := time.ParseDuration(jwtExpiresIn)
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(jwtExpiresIn).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

func ParseToken(c echo.Context) (uint, error) {
	tokenStr := c.Request().Header.Get("Authorization")
	if tokenStr == "" {
		return 0, echo.ErrUnauthorized
	}
	// "Bearer <token>"の形式で送られてくるので、"Bearer "を除去する
	tokenStr = tokenStr[len("Bearer "):]

	// トークンを検証
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 署名の検証に使用する鍵を指定
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	// トークンが有効であることを確認
	if !token.Valid {
		return 0, echo.ErrUnauthorized
	}

	// トークンからユーザーIDを取得
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("failed to get claims")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("failed to get user_id")
	}
	return uint(userID), nil
}

func GetUserID(c echo.Context) (uint, error) {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return 0, echo.ErrUnauthorized
	}
	return userID, nil
}
