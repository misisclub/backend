package handlers

import (
	"echo-template/internal/models"
	usecase "echo-template/internal/use_case"
	"echo-template/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type SubscriptionHandler struct {
	service   *usecase.SubscriptionService
	validator *validator.Validate
}

func NewSubscriptionHandler(service *usecase.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		service:   service,
		validator: validator.New(),
	}
}

// CreateSubscription godoc
//
//	@Summary		Create subscription
//	@Description	create subscription with given data
//	@Tags			Subscriptions
//	@Accept			json
//	@Produce		json
//	@Param			subscription	body		models.SubscriptionCreate	true	"Subscription data"
//	@Success		201		{object}	models.SubscriptionModel
//	@Failure		400 {object} utils.Err
//	@Failure		500 {object} utils.Err
//	@Router			/subscriptions [post]
func (h *SubscriptionHandler) CreateSubscription(c echo.Context) error {
	var req models.SubscriptionCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request body"})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Validation failed"})
	}

	subscription, err := h.service.CreateSubscription(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to create subscription"})
	}

	return c.JSON(http.StatusCreated, subscription)
}

// GetSubscriptionByID godoc
//
//	@Summary		Get subscription by id
//	@Description	Returns a single subscription by id
//	@Tags			Subscriptions
//	@Produce		json
//	@Param			id	path	string	true	"Subscription ID"
//	@Success		200	{object}	models.SubscriptionModel
//	@Failure		400	{object}	utils.Err
//	@Failure		404	{object}	utils.Err
//	@Failure		500	{object}	utils.Err
//	@Router			/subscriptions/{id} [get]
func (h *SubscriptionHandler) GetSubscriptionByID(c echo.Context) error {
	var req models.SubscriptionGet
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request"})
	}

	subscription, err := h.service.GetSubscriptionByID(c.Request().Context(), req.ID)
	if err != nil {
		if err.Error() == "subscription not found" {
			return c.JSON(http.StatusNotFound, utils.Err{Message: "Subscription not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to get subscription"})
	}

	return c.JSON(http.StatusOK, subscription)
}

// GetSubscriptionsByUserID godoc
//
//	@Summary		Get subscriptions by user id
//	@Description	Returns all subscriptions for a specific user
//	@Tags			Subscriptions
//	@Produce		json
//	@Param			user_id	path	string	true	"User ID"
//	@Success		200	{object}	[]models.SubscriptionModel
//	@Failure		400	{object}	utils.Err
//	@Failure		500	{object}	utils.Err
//	@Router			/subscriptions/user/{user_id} [get]
func (h *SubscriptionHandler) GetSubscriptionsByUserID(c echo.Context) error {
	var req models.SubscriptionGetByUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request"})
	}

	subscriptions, err := h.service.GetSubscriptionsByUserID(c.Request().Context(), req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to get subscriptions"})
	}

	return c.JSON(http.StatusOK, subscriptions)
}

// GetSubscriptionsByClubID godoc
//
//	@Summary		Get subscriptions by club id
//	@Description	Returns all subscriptions for a specific club
//	@Tags			Subscriptions
//	@Produce		json
//	@Param			club_id	path	string	true	"Club ID"
//	@Success		200	{object}	[]models.SubscriptionModel
//	@Failure		400	{object}	utils.Err
//	@Failure		500	{object}	utils.Err
//	@Router			/subscriptions/club/{club_id} [get]
func (h *SubscriptionHandler) GetSubscriptionsByClubID(c echo.Context) error {
	var req models.SubscriptionGetByClub
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request"})
	}

	subscriptions, err := h.service.GetSubscriptionsByClubID(c.Request().Context(), req.ClubID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to get subscriptions"})
	}

	return c.JSON(http.StatusOK, subscriptions)
}

// GetSubscriptionByUserAndClub godoc
//
//	@Summary		Get subscription by user and club
//	@Description	Returns a single subscription for specific user and club
//	@Tags			Subscriptions
//	@Produce		json
//	@Param			user_id	path	string	true	"User ID"
//	@Param			club_id	path	string	true	"Club ID"
//	@Success		200	{object}	models.SubscriptionModel
//	@Failure		400	{object}	utils.Err
//	@Failure		404	{object}	utils.Err
//	@Failure		500	{object}	utils.Err
//	@Router			/subscriptions/user/{user_id}/club/{club_id} [get]
func (h *SubscriptionHandler) GetSubscriptionByUserAndClub(c echo.Context) error {
	var req models.SubscriptionGetByUserAndClub
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request"})
	}

	subscription, err := h.service.GetSubscriptionByUserAndClub(c.Request().Context(), req.UserID, req.ClubID)
	if err != nil {
		if err.Error() == "subscription not found" {
			return c.JSON(http.StatusNotFound, utils.Err{Message: "Subscription not found"})
		}
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to get subscription"})
	}

	return c.JSON(http.StatusOK, subscription)
}

// DeleteSubscription godoc
//
//	@Summary		Delete subscription
//	@Description	Delete subscription by id
//	@Tags			Subscriptions
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Subscription ID"
//	@Success		204
//	@Failure		400		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/subscriptions/{id} [delete]
func (h *SubscriptionHandler) DeleteSubscription(c echo.Context) error {
	var req models.SubscriptionDelete
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request"})
	}

	err := h.service.DeleteSubscription(c.Request().Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to delete subscription"})
	}

	return c.NoContent(http.StatusNoContent)
}

// DeleteSubscriptionByUserAndClub godoc
//
//	@Summary		Delete subscription by user and club
//	@Description	Delete subscription by user id and club id
//	@Tags			Subscriptions
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	string	true	"User ID"
//	@Param			club_id	path	string	true	"Club ID"
//	@Success		204
//	@Failure		400		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/subscriptions/user/{user_id}/club/{club_id} [delete]
func (h *SubscriptionHandler) DeleteSubscriptionByUserAndClub(c echo.Context) error {
	var req models.SubscriptionDeleteByUserAndClub
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "Invalid request"})
	}

	err := h.service.DeleteSubscriptionByUserAndClub(c.Request().Context(), req.UserID, req.ClubID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: "Failed to delete subscription"})
	}

	return c.NoContent(http.StatusNoContent)
}
