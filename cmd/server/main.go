package main

import (
	"github.com/joho/godotenv"
	"github.com/satyamvatstyagi/UserManagementService/pkg/api/routes"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
)

func main() {
	godotenv.Load()
	logger.NewLogger().SetLogger()
	routes.Setup()
}
