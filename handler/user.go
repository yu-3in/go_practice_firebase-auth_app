package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/handler/dto"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func (h *Handler) CreateUser(c echo.Context) error {
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

	req := &dto.CreateUserRequest{
		Email: fbUser.Email,
		Name:  fbUser.DisplayName,
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	user := dto.CreateUserRequestToUserModel(req)
	user.ID = fbUser.UID

	if err := h.repo.CreateUser(user); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to create user")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"user": dto.UserModelToUserResponse(user),
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	userID := c.Get("userID").(string)
	log.Println("GetUser", userID)

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

	userID := c.Get("userID").(string)

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
	userID := c.Get("userID").(string)

	if err := h.repo.DeleteUser(userID); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to delete user")
	}

	return c.NoContent(http.StatusNoContent)
}
