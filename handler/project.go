package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/handler/dto"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/util"
)

func (h *Handler) CreateProject(c echo.Context) error {
	req := new(dto.CreateProjectRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	userID, err := util.GetUserID(c)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user_id")
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	project := dto.CreateProjectRequestToProjectModel(req)
	project.OwnerID = userID

	if err := h.repo.CreateProject(project); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to create project")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"project": dto.ProjectModelToProjectResponse(project),
	})
}

func (h *Handler) GetProjects(c echo.Context) error {
	projects, err := h.repo.GetProjects()
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get projects")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"projects": dto.ProjectModelsToProjectResponses(projects),
	})
}

// 自分がオーナーであるプロジェクトを取得
func (h *Handler) GetOwnerProjects(c echo.Context) error {
	userID, err := util.GetUserID(c)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user_id")
	}

	projects, err := h.repo.GetOwnerProjects(userID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get projects")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"projects": dto.ProjectModelsToProjectResponses(projects),
	})
}

// 自分がメンバーであるプロジェクトを取得
func (h *Handler) GetMemberProjects(c echo.Context) error {
	userID, err := util.GetUserID(c)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get user_id")
	}

	projects, err := h.repo.GetMemberProjects(userID)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get projects")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"projects": dto.ProjectModelsToProjectResponses(projects),
	})
}

func (h *Handler) GetProject(c echo.Context) error {
	projectIDStr := c.Param("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}
	project, err := h.repo.GetProject(uint(projectID))
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get project")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"project": dto.ProjectModelToProjectResponse(project),
	})
}

func (h *Handler) UpdateProject(c echo.Context) error {
	req := new(dto.UpdateProjectRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	projectIDStr := c.Param("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}
	project, err := h.repo.GetProject(uint(projectID))
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get project")
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	dto.UpdateProjectRequestToProjectModel(req, project)

	if err := h.repo.UpdateProject(project); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to update project")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"project": dto.ProjectModelToProjectResponse(project),
	})
}

func (h *Handler) DeleteProject(c echo.Context) error {
	projectIDStr := c.Param("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}
	if err := h.repo.DeleteProject(uint(projectID)); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to delete project")
	}

	return c.NoContent(http.StatusNoContent)
}
