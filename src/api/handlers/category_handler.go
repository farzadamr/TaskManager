package handlers

import (
	"net/http"
	"strconv"

	"github.com/farzadamr/TaskManager/services"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

type createCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	var req createCategoryRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	category, err := h.categoryService.CreateCategory(c.Request().Context(), userID, req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetUserCategories(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	categories, err := h.categoryService.GetUserCategories(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := h.categoryService.DeleteCategory(c.Request().Context(), userID, uint(categoryID)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
