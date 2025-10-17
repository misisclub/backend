package handlers

import (
	"echo-template/internal/models"
	"echo-template/internal/utils"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"

	usecase "echo-template/internal/use_case"
)

type ClientHandler struct {
	clientService *usecase.ClientService
	validate      *validator.Validate
}

func NewAuthHandler(clientService *usecase.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
		validate:      validator.New(),
	}
}

// ClientAuth godoc
//
//	@Summary		Create client
//	@Description	create client with given data
//	@Tags			Clients
//	@Accept			json
//	@Produce		json
//	@Param			client	body		models.ClientSignUp	true	"Credentials to use"
//	@Success		201		{object}	models.SignSuccess
//	@Failure		400 {object} utils.Err
//	@Failure		500 {object} utils.Err
//	@Router			/clients/auth/sign-up [post]
func (h *ClientHandler) SignUpClient(c echo.Context) error {
	var client models.ClientSignUp
	if err := c.Bind(&client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	if err := h.validate.Struct(client); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}

	sign, err := h.clientService.SignUpClient(c.Request().Context(), &client)
	if errors.Is(err, &pgconn.PgError{Code: "23505"}) {
		return c.JSON(http.StatusConflict, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, sign)
}

// PartnerAuth godoc
//
//	@Summary		Sign-in for partners
//	@Description	sign-in in partners with given data
//	@Tags			Clients
//	@Accept			json
//	@Produce		json
//	@Param			partner	body		models.ClientSignIn	true	"Credentials to use"
//	@Success		200		{object}	models.SignSuccess
//	@Failure		400		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/clients/auth/sign-in [post]
func (h *ClientHandler) SignInClient(c echo.Context) error {
	var client models.ClientSignIn
	if err := c.Bind(&client); err != nil {
		return c.JSON(utils.BadRequestError())
	}
	if err := h.validate.Struct(client); err != nil {
		return c.JSON(utils.BadRequestError())
	}
	sign, err := h.clientService.SignInClient(c.Request().Context(), &client)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, sign)
}

// UpdateClient godoc
//
//	@Summary		Update client
//	@Description	Update client fields by id
//	@Tags			Clients
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Client ID"
//	@Param			client	body		models.ClientUpdate	true	"Fields to update"
//	@Success		200		{object}	models.ClientModel
//	@Failure		400		{object}	utils.Err
//	@Failure		404		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/clients/{id} [put]
func (h *ClientHandler) UpdateClient(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "id is required"})
	}
	uid, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	var req models.ClientUpdate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	updated, err := h.clientService.UpdateClient(c.Request().Context(), &req, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

// DeleteClient godoc
//
//	@Summary		Delete client
//	@Description	Delete client by id
//	@Tags			Clients
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Client ID"
//	@Success		204
//	@Failure		400		{object}	utils.Err
//	@Failure		500		{object}	utils.Err
//	@Router			/clients/{id} [delete]
func (h *ClientHandler) DeleteClient(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: "id is required"})
	}
	uid, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	if err := h.clientService.DeleteClient(c.Request().Context(), uid); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Err{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// GetClientByID godoc
//
//	@Summary		Get client by id
//	@Description	Returns a single client by id
//	@Tags			Clients
//	@Produce		json
//	@Param			id	path	string	true	"Client ID"
//	@Success		200	{object}	models.ClientModel
//	@Failure		400	{object}	utils.Err
//	@Failure		404	{object}	utils.Err
//	@Router			/clients/{id} [get]
func (h *ClientHandler) GetClientByID(c echo.Context) error {
	var req models.ClientGet
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Err{Message: err.Error()})
	}
	c.Logger().Debug(req.ID.String())
	client, err := h.clientService.GetClientByID(c.Request().Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, client)
}
