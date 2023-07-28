package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func HTTPError(code int, err error, message ...string) *echo.HTTPError {
	var errMsg string
	if len(message) > 0 {
		errMsg = fmt.Errorf("%s: %w", message[0], err).Error()
	} else {
		errMsg = fmt.Errorf("%w", err).Error()
	}

	return echo.NewHTTPError(code, ErrorResponse{Message: errMsg})
}
