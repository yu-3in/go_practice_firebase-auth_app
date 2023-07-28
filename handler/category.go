package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yuuumiravy/go_practice_projects_firebase-auth_app_2023_07/handler/dto"
)

func (h *Handler) CreateCategory(c echo.Context) error {
	req := new(dto.CreateCategoryRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	category := dto.CreateCategoryRequestToCategoryModel(req)

	if err := h.repo.CreateCategory(category); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to create category")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"category": dto.CategoryModelToCategoryResponse(category),
	})
}

func (h *Handler) GetCategories(c echo.Context) error {
	categories, err := h.repo.GetCategories()
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get categories")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"categories": dto.CategoryModelsToCategoryResponses(categories),
	})
}

func (h *Handler) GetCategoryProjects(c echo.Context) error {
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	projects, err := h.repo.GetCategoryProjects(uint(categoryID))
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get category projects")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"projects": dto.ProjectModelsToProjectResponses(projects),
	})
}

func (h *Handler) GetCategory(c echo.Context) error {
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	category, err := h.repo.GetCategory(uint(categoryID))
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get category")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"category": dto.CategoryModelToCategoryResponse(category),
	})
}

func (h *Handler) UpdateCategory(c echo.Context) error {
	req := new(dto.UpdateCategoryRequest)
	if err := c.Bind(&req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	if err := c.Validate(req); err != nil {
		return HTTPError(http.StatusBadRequest, err, "validation error")
	}

	category, err := h.repo.GetCategory(uint(categoryID))
	if err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to get category")
	}
	dto.UpdateCategoryRequestToCategoryModel(req, category)

	if err := h.repo.UpdateCategory(category); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to update category")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"category": dto.CategoryModelToCategoryResponse(category),
	})
}

func (h *Handler) DeleteCategory(c echo.Context) error {
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return HTTPError(http.StatusBadRequest, err, "bad request")
	}

	if err := h.repo.DeleteCategory(uint(categoryID)); err != nil {
		return HTTPError(http.StatusInternalServerError, err, "failed to delete category")
	}

	return c.NoContent(http.StatusNoContent)
}
