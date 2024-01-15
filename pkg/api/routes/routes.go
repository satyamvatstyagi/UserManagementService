package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/ELKTestService1/pkg/app/config"
)

func Setup() {

	cfg := config.Config{}

	// Initialize the database
	cfg.InitDb()

	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)
	router := gin.Default()

	// Setup the routes

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
