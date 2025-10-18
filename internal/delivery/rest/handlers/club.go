package handlers

import (
	"echo-template/internal/models"
	"echo-template/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	usecase "echo-template/internal/use_case"
)

type ClubHandler struct {
	clubService *usecase.ClubService
	validate    *validator.Validate
}

func NewClubHandler(clubService *usecase.ClubService) *ClubHandler {
	return &ClubHandler{
		clubService: clubService,
		validate:    validator.New(),
	}
}

// CreateClub godoc
//
//	@Summary		Create club
//	@Description	create club with given data
//	@Tags			Clubs
//	@Accept			json
//	@Produce		json
//	@Param			club	body		models.ClubCreate	true	"Club data"
//	@Success		201		{object}	db.Club
//	@Failure		400 {object} utils.Err
//	@Failure		500 {object} utils.Err
//	@Router			/clubs [post]
func (h *ClubHandler) CreateClub(c echo.Context) error {
	var club models.ClubCreate
	if err := c.Bind(&club); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	if err := h.validate.Struct(club); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	createdClub, err := h.clubService.CreateClub(c.Request().Context(), &club)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, createdClub)
}

// UpdateClub godoc
//
//	@Summary		Update club
//	@Description	Update club fields by id
//	@Tags			Clubs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Club ID"
//	@Param			club	body		models.ClubUpdate	true	"Fields to update"
//	@Success		200		{object}	db.Club
//	@Failure		400		{object}	utils.Err
//	@Failure		404		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/clubs/{id} [put]
func (h *ClubHandler) UpdateClub(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "id is required"})
	}
	uid, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	var req models.ClubUpdate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	updated, err := h.clubService.UpdateClub(c.Request().Context(), &req, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

// DeleteClub godoc
//
//	@Summary		Delete club
//	@Description	Delete club by id
//	@Tags			Clubs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Club ID"
//	@Success		204
//	@Failure		400		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/clubs/{id} [delete]
func (h *ClubHandler) DeleteClub(c echo.Context) error {
	var req models.ClubDelete
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	if err := h.clubService.DeleteClub(c.Request().Context(), req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// GetClubByID godoc
//
//	@Summary		Get club by id
//	@Description	Returns a single club by id
//	@Tags			Clubs
//	@Produce		json
//	@Param			id	path	string	true	"Club ID"
//	@Success		200	{object}	db.GetClubByIDRow
//	@Failure		400	{object}	utils.Err
//	@Failure		404	{object}	utils.Err
//	@Router			/clubs/{id} [get]
func (h *ClubHandler) GetClubByID(c echo.Context) error {
	var req models.ClubGet
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	club, err := h.clubService.GetClubByID(c.Request().Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, club)
}
