package config

import (
	"log"
	"os"

	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"gorm.io/driver/postgres"
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
