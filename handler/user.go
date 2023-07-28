package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/handler/dto"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func (h *Handler) CreateUser(c echo.Context) error {
	req := new(dto.UserRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	user := dto.UserRequestToUserModel(req)

	if err := h.repo.CreateUser(user); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to create user")
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to generate token")
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"token": token,
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	userID, err := util.GetUserID(c)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user_id")
	}
	user, err := h.repo.GetUser(userID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": dto.UserModelToUserResponse(user),
	})
}
