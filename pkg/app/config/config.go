package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/env"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{}

func (c *Config) InitDb() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.EnvConfig.DatabaseHost, env.EnvConfig.DatabaseUser, env.EnvConfig.DatabasePassword, env.EnvConfig.DatabaseName, env.EnvConfig.DatabasePort)
	log.Println("Connecting to database: ", dsn)

	connect := true
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		connect = false
		log.Println("Error connecting to database: ", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		connect = false
		log.Println("Error migrating database: ", err)
	}

	if connect {
		log.Printf("Database connected successfully to host: %s, port: %s, user: %s, dbname: %s", env.EnvConfig.DatabaseHost, env.EnvConfig.DatabasePort, env.EnvConfig.DatabaseUser, env.EnvConfig.DatabaseName)
	} else {
		log.Println("Database connection failed")
	}

	return db
}

// Function to initialize the logger
func (c *Config) InitLogger() logger.Logger {

	fileName := strings.ToLower(consts.AppName) + ".log"

	// Create the log file path
	filePath := env.EnvConfig.LogFilePath + "/" + fileName

	logger, err := logger.NewMtnLogger(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
