package application

import (
	"echo-template/internal/delivery/rest/router"
	"echo-template/internal/infrastructure"
	"echo-template/internal/infrastructure/database"
	"echo-template/internal/infrastructure/logger"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/karagenc/zap4echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Application struct {
	e         *echo.Echo
	Address   string
	DB        *pgxpool.Pool
	Logger    *logger.Logger
	SecretKey string
}

func NewApplication(config *infrastructure.Config) *Application {
	e := echo.New()
	l := logger.NewLogger()
	l.Infof("im born!")

	db, err := database.NewPostgresDB(config, l)
	if err != nil {
		l.Errorf("failed to connect to database: %s", err.Error())
		return nil
	}

	return &Application{
		e:         e,
		Address:   config.Server.Address,
		DB:        db,
		Logger:    l,
		SecretKey: config.Other.JWTKey,
	}
}

func (a *Application) RunServer() error {
	e := initServer(a)

	c := jaegertracing.New(e, nil)
	defer c.Close()

	a.Logger.Info("Starting server on " + a.Address)

	if err := e.Start(a.Address); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			a.Logger.Errorf("Failed to start server: %s", err.Error())
		}
	}
	return nil
}

func initServer(a *Application) *echo.Echo {
	e := a.e
	e.Use(zap4echo.Logger(a.Logger.Desugar()))
	e.Use(middleware.Recover())

	router.RegisterRouter(e, a.DB, a.Logger)

	return e
}
