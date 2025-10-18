package handlers

import (
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/utils"
	"errors"
	"net/http"
	"strconv"

	usecase "echo-template/internal/use_case"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// FeedHandler handles feed HTTP requests
type FeedHandler struct {
	feedService *usecase.FeedService
	validator   *utils.Validator
}

// NewFeedHandler creates a new feed handler
func NewFeedHandler(feedService *usecase.FeedService, validator *utils.Validator) *FeedHandler {
	return &FeedHandler{
		feedService: feedService,
		validator:   validator,
	}
}

// GetSubscriptionFeed retrieves posts from subscribed clubs for a user
// @Summary Get subscription feed
// @Description Retrieve posts from clubs that the user is subscribed to
// @Tags feed
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param limit query int false "Number of posts to return (max 100)" default(20)
// @Param offset query int false "Number of posts to skip" default(0)
// @Param tag query string false "Filter posts by tag"
// @Success 200 {object} models.FeedResponse
// @Failure 400 {object} utils.Err
// @Failure 404 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /feed/subscription/{user_id} [get]
func (h *FeedHandler) GetSubscriptionFeed(c echo.Context) error {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid user ID"})
	}

	limitStr := c.QueryParam("limit")
	limit := int32(20) // default
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid limit parameter"})
		}
		limit = int32(limitInt)
	}

	offsetStr := c.QueryParam("offset")
	offset := int32(0) // default
	if offsetStr != "" {
		offsetInt, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid offset parameter"})
		}
		offset = int32(offsetInt)
	}

	tagStr := c.QueryParam("tag")
	var tag *string = nil
	if tagStr != "" {
		tag = &tagStr
	}

	feed, err := h.feedService.GetSubscriptionFeed(c.Request().Context(), userID, tag, limit, offset)
	if err != nil {
		if errors.Is(err, repository.ErrFeedNotFound) {
			return c.JSON(http.StatusNotFound, utils.Err{Message: "Feed not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to retrieve subscription feed"})
	}

	return c.JSON(http.StatusOK, feed)
}

// GetEveryoneFeed retrieves all posts ordered by creation date
// @Summary Get everyone feed
// @Description Retrieve all posts ordered by creation date
// @Tags feed
// @Accept json
// @Produce json
// @Param limit query int false "Number of posts to return (max 100)" default(20)
// @Param offset query int false "Number of posts to skip" default(0)
// @Param tag query string false "Filter posts by tag"
// @Success 200 {object} models.FeedResponse
// @Failure 400 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /feed/everyone [get]
func (h *FeedHandler) GetEveryoneFeed(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit := int32(20) // default
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid limit parameter"})
		}
		limit = int32(limitInt)
	}

	offsetStr := c.QueryParam("offset")
	offset := int32(0) // default
	if offsetStr != "" {
		offsetInt, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid offset parameter"})
		}
		offset = int32(offsetInt)
	}

	tagStr := c.QueryParam("tag")
	var tag *string = nil
	if tagStr != "" {
		tag = &tagStr
	}

	// Validate request parameters
	if limit <= 0 || limit > 100 {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Limit must be between 1 and 100"})
	}
	if offset < 0 {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Offset must be non-negative"})
	}

	feed, err := h.feedService.GetEveryoneFeed(c.Request().Context(), tag, limit, offset)
	if err != nil {
		// Handle specific service errors
		if errors.Is(err, errors.New("invalid pagination parameters")) {
			return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid pagination parameters"})
		}
		if errors.Is(err, errors.New("failed to retrieve posts from database")) {
			return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to retrieve posts from database"})
		}
		if errors.Is(err, errors.New("failed to retrieve total post count")) {
			return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to retrieve total post count"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to retrieve everyone feed"})
	}

	return c.JSON(http.StatusOK, feed)
}

// GetFeedByClubID retrieves posts from a specific club
// @Summary Get club feed
// @Description Retrieve posts from a specific club
// @Tags feed
// @Accept json
// @Produce json
// @Param club_id path string true "Club ID"
// @Param limit query int false "Number of posts to return (max 100)" default(20)
// @Param offset query int false "Number of posts to skip" default(0)
// @Param tag query string false "Filter posts by tag"
// @Success 200 {object} models.FeedResponse
// @Failure 400 {object} utils.Err
// @Failure 404 {object} utils.Err
// @Failure 500 {object} utils.Err
// @Router /feed/club/{club_id} [get]
func (h *FeedHandler) GetFeedByClubID(c echo.Context) error {
	clubIDStr := c.Param("club_id")
	clubID, err := uuid.Parse(clubIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid club ID"})
	}

	limitStr := c.QueryParam("limit")
	limit := int32(20) // default
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid limit parameter"})
		}
		limit = int32(limitInt)
	}

	offsetStr := c.QueryParam("offset")
	offset := int32(0) // default
	if offsetStr != "" {
		offsetInt, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid offset parameter"})
		}
		offset = int32(offsetInt)
	}

	tagStr := c.QueryParam("tag")
	var tag *string = nil
	if tagStr != "" {
		tag = &tagStr
	}

	feed, err := h.feedService.GetFeedByClubID(c.Request().Context(), clubID, tag, limit, offset)
	if err != nil {
		if errors.Is(err, repository.ErrFeedNotFound) {
			return c.JSON(http.StatusNotFound, utils.Err{Message: "Feed not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to retrieve club feed"})
	}

	return c.JSON(http.StatusOK, feed)
}
