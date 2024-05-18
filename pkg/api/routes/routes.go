package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/UserManagementService/pkg/api/controller"
	"github.com/satyamvatstyagi/UserManagementService/pkg/api/middlewares"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/config"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/repository"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/usecase"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/env"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/restclient"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin/v2"
)

func Setup() {
	cerr := env.LoadConfig()
	if cerr != nil {
		log.Fatalf("Loadconfig failed, err=%s", cerr.Error())
	}

	ginMode := env.EnvConfig.GinMode
	gin.SetMode(ginMode)
	router := gin.Default()

	// Set up routes
	setupRoutes(router)

	// Start the server
	port := env.EnvConfig.ServicePort
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(router *gin.Engine) {
	// Initialize the logger
	cfg := config.Config{}
	logger := cfg.InitLogger()
	logger.Info(fmt.Sprintf("Starting the %s service, version: %s", consts.AppName, consts.AppVersion))

	// Use Elastic APM middleware for Gin
	router.Use(apmgin.Middleware(router))
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))
	router.Use(middlewares.LoggingMiddleware(logger))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Initialize the HTTP client
	httpClient := &http.Client{Timeout: consts.MaxTimeout}
	restHTTPClient := restclient.NewHTTPClient(httpClient)

	// Initialize the database
	db := cfg.InitDb()

	// Initialize the repository
	userRepository := repository.NewUserRepository(db)

	// Initialize the usecases
	userUsecase := usecase.NewUserUsecase(userRepository, restHTTPClient)

	// Initialize the controller
	userController := &controller.UserController{UserUsecase: userUsecase}

	username := env.EnvConfig.BasicAuthUser
	password := env.EnvConfig.BasicAuthPassword
	router.GET("/user/health", middlewares.LoggingMiddleware(logger), userController.HealthCheck)
	userService := router.Group("/user", gin.BasicAuth(gin.Accounts{username: password}))
	{
		userService.POST("/register", middlewares.LoggingMiddleware(logger), userController.RegisterUser)
		userService.POST("/login", middlewares.LoggingMiddleware(logger), userController.LoginUser)
		userService.GET("/:username", middlewares.LoggingMiddleware(logger), userController.GetUserByUserName)
		userService.GET("/:username/order", middlewares.LoggingMiddleware(logger), userController.GetOrderByOrderUserName)
	}
}
