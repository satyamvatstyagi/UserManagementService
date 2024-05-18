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
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("Error connecting to database: ", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Error migrating database: ", err)
	}
	//DropUnusedColumns(db, &models.User{})

	return db
}

// Removes any unsed column that exist when a mmodel struct field changes
// func DropUnusedColumns(db *gorm.DB, models ...interface{}) {
// 	for _, model := range models {
// 		stmt := &gorm.Statement{DB: db}
// 		stmt.Parse(model)
// 		fields := stmt.Schema.Fields
// 		columns, _ := db.Migrator().ColumnTypes(model)

// 		for i := range columns {
// 			found := false
// 			for j := range fields {
// 				if columns[i].Name() == fields[j].DBName {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				db.Migrator().DropColumn(model, columns[i].Name())
// 			}
// 		}
// 	}
// }

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
