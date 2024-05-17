package main

import (
	_ "github.com/satyamvatstyagi/UserManagementService/docs"

	"github.com/joho/godotenv"
	"github.com/satyamvatstyagi/UserManagementService/pkg/api/routes"
)

// @title			User Management Service
// @version			1.0
// @description		API for user management service which allows user to register, login and get user details.
// @schemes			https
// @host			localhost:8080
// @BasePath		/user
func main() {
	godotenv.Load()

	routes.Setup()
}
