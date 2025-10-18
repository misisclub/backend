package handlers

import (
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/models"
	"echo-template/internal/utils"
	"errors"
	"net/http"

	usecase "echo-template/internal/use_case"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postService *usecase.PostService
	validator   *utils.Validator
}

func NewPostHandler(postService *usecase.PostService, validator *utils.Validator) *PostHandler {
	return &PostHandler{
		postService: postService,
		validator:   validator,
	}
}

// CreatePost creates a new post
// @Summary Create a new post
// @Description Create a new post with the provided information
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "Post creation data"
// @Success 201 {object} models.PostResponse
// @Failure 400 {object} utils.Err
// @Failure 409 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /posts [post]
func (h *PostHandler) CreatePost(c echo.Context) error {
	var req models.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request body"})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	post, err := h.postService.CreatePost(c.Request().Context(), &req)
	if err != nil {
		if errors.Is(err, repository.ErrPostExists) {
			return c.JSON(http.StatusConflict, utils.Err{Message: "Post already exists for this owner"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to create post"})
	}

	return c.JSON(http.StatusCreated, post)
}

// GetPost retrieves a post by ID
// @Summary Get a post by ID
// @Description Retrieve a post by its unique identifier
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.PostResponse
// @Failure 400 {object} utils.Err
// @Failure 404 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /posts/{id} [get]
func (h *PostHandler) GetPost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid post ID"})
	}

	post, err := h.postService.GetPost(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return c.JSON(http.StatusNotFound, utils.Err{Message: "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to retrieve post"})
	}

	return c.JSON(http.StatusOK, post)
}

// UpdatePost updates an existing post
// @Summary Update a post
// @Description Update an existing post with new information
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body models.UpdatePostRequest true "Post update data"
// @Success 200 {object} models.PostResponse
// @Failure 400 {object} utils.Err
// @Failure 404 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /posts/{id} [put]
func (h *PostHandler) UpdatePost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid post ID"})
	}

	var req models.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request body"})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	post, err := h.postService.UpdatePost(c.Request().Context(), id, &req)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return c.JSON(http.StatusNotFound, utils.Err{Message: "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to update post"})
	}

	return c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post by ID
// @Summary Delete a post
// @Description Delete a post by its unique identifier
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.Err
// @Failure 404 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /posts/{id} [delete]
func (h *PostHandler) DeletePost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid post ID"})
	}

	err = h.postService.DeletePost(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return c.JSON(http.StatusNotFound, utils.Err{Message: "Post not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to delete post"})
	}

	return c.NoContent(http.StatusNoContent)
}
