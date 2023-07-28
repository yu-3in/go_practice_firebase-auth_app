package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/handler/dto"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func (h *Handler) CreateUser(c echo.Context) error {
	req := new(dto.CreateUserRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	user := dto.CreateUserRequestToUserModel(req)

	if err := h.repo.CreateUser(user); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to create user")
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to generate token")
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"token": token,
		"user":  dto.UserModelToUserResponse(user),
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	userID, err := util.GetUserID(c)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get userID")
	}
	user, err := h.repo.GetUser(userID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": dto.UserModelToUserResponse(user),
	})
}

func (h *Handler) UpdateUser(c echo.Context) error {
	req := new(dto.UpdateUserRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	userID, err := util.GetUserID(c)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get userID")
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	user, err := h.repo.GetUser(userID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user")
	}
	dto.UpdateUserRequestToUserModel(req, user)

	if err := h.repo.UpdateUser(user); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to update user")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": dto.UserModelToUserResponse(user),
	})
}

func (h *Handler) DeleteUser(c echo.Context) error {
	userID, err := util.GetUserID(c)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get userID")
	}

	if err := h.repo.DeleteUser(userID); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to delete user")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) Login(c echo.Context) error {
	req := new(dto.LoginRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	user, err := h.repo.GetUserByEmail(req.Email)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user")
	}

	if err := util.VerifyPassword(user.Password, req.Password); err != nil {
		return HTTPError(http.StatusBadRequest, err, "invalid email or password")
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to generate token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  dto.UserModelToUserResponse(user),
	})
}

func (h *Handler) Logout(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
