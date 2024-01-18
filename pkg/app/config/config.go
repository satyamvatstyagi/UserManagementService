package config

import (
	"log"
	"os"

	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{}

func (c *Config) InitDb() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	DropUnusedColumns(db, &models.User{})

	return db
}

// Removes any unsed column that exist when a mmodel struct field changes
func DropUnusedColumns(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		stmt := &gorm.Statement{DB: db}
		stmt.Parse(model)
		fields := stmt.Schema.Fields
		columns, _ := db.Migrator().ColumnTypes(model)

		for i := range columns {
			found := false
			for j := range fields {
				if columns[i].Name() == fields[j].DBName {
					found = true
					break
				}
			}
			if !found {
				db.Migrator().DropColumn(model, columns[i].Name())
			}
		}
	}
}

// Function to initialize the logger
func (c *Config) InitLogger() *logger.MtnLogger {
	fileName := os.Getenv("LOG_FILE_NAME")

	// If the LOG_FILE_NAME environment variable is not set, set it to "app.log"
	if fileName == "" {
		fileName = "app.log"
	}

	// Check if "log" directory exists in the current directory
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		// Create the directory
		err := os.Mkdir("log", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Check if the file exists in the "log" directory
	filePath := "log/" + fileName
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	logger, err := logger.NewMtnLogger(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
