package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/UserManagementService/pkg/api/controller"
	"github.com/satyamvatstyagi/UserManagementService/pkg/api/middlewares"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/config"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/repository"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/usecase"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/restclient"
	"go.elastic.co/apm/module/apmgin/v2"
)

func Setup() {

	cfg := config.Config{}

	// Initialize the logger
	logger := cfg.InitLogger()

	// Initialize the database
	db := cfg.InitDb()

	// Initialize the repositories
	userRepository := repository.NewUserRepository(db)

	httpClient := &http.Client{Timeout: consts.MaxTimeout}
	restHTTPClient := restclient.NewHTTPClient(httpClient)

	// Initialize the usecases
	userUsecase := usecase.NewUserUsecase(userRepository, restHTTPClient)

	// Initialize the controller
	userController := &controller.UserController{UserUsecase: userUsecase}

	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)
	router := gin.Default()

	// Use Elastic APM middleware for Gin
	router.Use(apmgin.Middleware(router))

	// Setup the routes
	setupUserRoutes(userController, router, logger)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupUserRoutes(c *controller.UserController, router *gin.Engine, l *logger.MtnLogger) {
	username := os.Getenv("BASIC_AUTH_USER")
	password := os.Getenv("BASIC_AUTH_PASSWORD")
	userService := router.Group("/user", gin.BasicAuth(gin.Accounts{username: password}))
	{
		userService.POST("/register", middlewares.LoggingMiddleware(l), c.RegisterUser)
		userService.POST("/login", middlewares.LoggingMiddleware(l), c.LoginUser)
		userService.GET("/:username", middlewares.LoggingMiddleware(l), c.GetUserByUserName)
		userService.GET("/:username/order", middlewares.LoggingMiddleware(l), c.GetOrderByOrderUserName)
	}
}
