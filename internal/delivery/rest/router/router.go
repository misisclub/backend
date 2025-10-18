package router

import (
	"echo-template/internal/delivery/rest/handlers"
	"echo-template/internal/infrastructure/logger"
	"echo-template/internal/infrastructure/repository"
	"echo-template/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	swag "github.com/swaggo/echo-swagger"

	usecase "echo-template/internal/use_case"
)

func RegisterRouter(e *echo.Echo, db *pgxpool.Pool, log *logger.Logger) {
	clientRepo := repository.NewClientRepository(db)
	orgRepo := repository.NewOrgRepository(db)
	clubRepo := repository.NewClubRepository(db)
	postRepo := repository.NewPostRepository(db)
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	feedRepo := repository.NewFeedRepository(db)

	clientService := usecase.NewClientService(clientRepo)
	orgService := usecase.NewOrgService(orgRepo)
	clubService := usecase.NewClubService(clubRepo)
	postService := usecase.NewPostService(postRepo)
	subscriptionService := usecase.NewSubscriptionService(subscriptionRepo)
	feedService := usecase.NewFeedService(feedRepo)

	validator := utils.NewValidator()

	authHandler := handlers.NewAuthHandler(clientService)
	orgHandler := handlers.NewOrgHandler(orgService)
	clubHandler := handlers.NewClubHandler(clubService)
	postHandler := handlers.NewPostHandler(postService, validator)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)
	feedHandler := handlers.NewFeedHandler(feedService, validator)

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
		org.POST("", orgHandler.CreateOrg)
		org.GET("/:id", orgHandler.GetOrgByID)
		org.PUT("/:id", orgHandler.UpdateOrg)
		org.DELETE("/:id", orgHandler.DeleteOrg)
	}

	// clubs CRUD
	club := api.Group("/clubs")
	{
		club.POST("", clubHandler.CreateClub)
		club.GET("/:id", clubHandler.GetClubByID)
		club.PUT("/:id", clubHandler.UpdateClub)
		club.DELETE("/:id", clubHandler.DeleteClub)
	}

	// posts CRUD
	post := api.Group("/posts")
	{
		post.POST("", postHandler.CreatePost)
		post.GET("/:id", postHandler.GetPost)
		post.PUT("/:id", postHandler.UpdatePost)
		post.DELETE("/:id", postHandler.DeletePost)
	}

	// subscriptions CRUD
	subscription := api.Group("/subscriptions")
	{
		subscription.POST("", subscriptionHandler.CreateSubscription)
		subscription.GET("/:id", subscriptionHandler.GetSubscriptionByID)
		subscription.GET("/user/:user_id", subscriptionHandler.GetSubscriptionsByUserID)
		subscription.GET("/club/:club_id", subscriptionHandler.GetSubscriptionsByClubID)
		subscription.GET("/user/:user_id/club/:club_id", subscriptionHandler.GetSubscriptionByUserAndClub)
		subscription.DELETE("/:id", subscriptionHandler.DeleteSubscription)
		subscription.DELETE("/user/:user_id/club/:club_id", subscriptionHandler.DeleteSubscriptionByUserAndClub)
	}

	// feed endpoints
	feed := api.Group("/feed")
	{
		feed.GET("/subscription/:user_id", feedHandler.GetSubscriptionFeed)
		feed.GET("/everyone", feedHandler.GetEveryoneFeed)
		feed.GET("/club/:club_id", feedHandler.GetFeedByClubID)
	}

	log.Info("Routes successfully registered")
}
