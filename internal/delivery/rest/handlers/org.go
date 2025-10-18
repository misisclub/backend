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

type OrgHandler struct {
	orgService *usecase.OrgService
	validate   *validator.Validate
}

func NewOrgHandler(orgService *usecase.OrgService) *OrgHandler {
	return &OrgHandler{
		orgService: orgService,
		validate:   validator.New(),
	}
}

// CreateOrg godoc
//
//	@Summary		Create organization
//	@Description	create organization with given data
//	@Tags			Orgs
//	@Accept			json
//	@Produce		json
//	@Param			org	body		models.OrgCreate	true	"Organization data"
//	@Success		201		{object}	db.Org
//	@Failure		400 {object} utils.Err
//	@Failure		500 {object} utils.Err
//	@Router			/orgs [post]
func (h *OrgHandler) CreateOrg(c echo.Context) error {
	var client models.OrgCreate
	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	if err := h.validate.Struct(client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	org, err := h.orgService.CreateOrg(c.Request().Context(), &client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, org)
}

// UpdateOrg godoc
//
//	@Summary		Update organization
//	@Description	Update organization fields by id
//	@Tags			Orgs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Org ID"
//	@Param			org	body		models.OrgUpdate	true	"Fields to update"
//	@Success		200		{object}	db.Org
//	@Failure		400		{object}	utils.Err
//	@Failure		404		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/orgs/{id} [put]
func (h *OrgHandler) UpdateOrg(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "id is required"})
	}
	uid, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	var req models.OrgUpdate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	updated, err := h.orgService.UpdateOrg(c.Request().Context(), &req, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

// DeleteOrg godoc
//
//	@Summary		Delete organization
//	@Description	Delete organization by id
//	@Tags			Orgs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Org ID"
//	@Success		204
//	@Failure		400		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/orgs/{id} [delete]
func (h *OrgHandler) DeleteOrg(c echo.Context) error {
	var req models.OrgDelete
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	if err := h.orgService.DeleteOrg(c.Request().Context(), req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// GetOrgByID godoc
//
//	@Summary		Get organization by id
//	@Description	Returns a single organization by id
//	@Tags			Orgs
//	@Produce		json
//	@Param			id	path	string	true	"Org ID"
//	@Success		200	{object}	db.Org
//	@Failure		400	{object}	utils.Err
//	@Failure		404	{object}	utils.Err
//	@Router			/orgs/{id} [get]
func (h *OrgHandler) GetOrgByID(c echo.Context) error {
	var req models.OrgGet
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	org, err := h.orgService.GetOrgByID(c.Request().Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, org)
}
