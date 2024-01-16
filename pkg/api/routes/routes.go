package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/UserManagementService/pkg/api/controller"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/config"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/repository"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/usecase"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/restclient"
)

func Setup() {

	cfg := config.Config{}

	// Initialize the logger
	logger := cfg.InitLogger()
	logger.Info(map[string]interface{}{"message": "Logger Initialized"})

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

	// Setup the routes
	setupUserRoutes(userController, router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupUserRoutes(c *controller.UserController, router *gin.Engine) {
	username := os.Getenv("BASIC_AUTH_USER")
	password := os.Getenv("BASIC_AUTH_PASSWORD")
	userService := router.Group("/user", gin.BasicAuth(gin.Accounts{username: password}))
	{
		userService.POST("/register", c.RegisterUser)
		userService.POST("/login", c.LoginUser)
		userService.GET("/:username", c.GetUserByUserName)
		userService.GET("/:username/order", c.GetOrderByOrderUserName)
	}
}
