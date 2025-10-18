package router

import (
	"echo-template/internal/delivery/rest/handlers"
	"echo-template/internal/infrastructure/logger"
	"echo-template/internal/infrastructure/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	swag "github.com/swaggo/echo-swagger"

	usecase "echo-template/internal/use_case"
)

func RegisterRouter(e *echo.Echo, db *pgxpool.Pool, log *logger.Logger) {
	clientRepo := repository.NewClientRepository(db)
	orgRepo := repository.NewOrgRepository(db)

	clientService := usecase.NewClientService(clientRepo)
	orgService := usecase.NewOrgService(orgRepo)

	authHandler := handlers.NewAuthHandler(clientService)
	orgHandler := handlers.NewOrgHandler(orgService)

	api := e.Group("/api/v1")

	api.GET("/ping", handlers.Ping)

	api.GET("/swagger/*", swag.WrapHandler)

	client := api.Group("/clients")
	authClient := client.Group("/auth")
	{
		authClient.POST("/sign-up", authHandler.SignUpClient)
		authClient.POST("/sign-in", authHandler.SignInClient)
	}

	// clients CRUD
	client.GET("/:id", authHandler.GetClientByID)
	client.PUT("/:id", authHandler.UpdateClient)
	client.DELETE("/:id", authHandler.DeleteClient)

	// organizations CRUD
	org := api.Group("/orgs")
	{
		org.POST("", orgHandler.SignUpOrg)
		org.GET("/:id", orgHandler.GetOrgByID)
		org.PUT("/:id", orgHandler.UpdateOrg)
		org.DELETE("/:id", orgHandler.DeleteOrg)
	}

	log.Info("Routes successfully registered")
}
