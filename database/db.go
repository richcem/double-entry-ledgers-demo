package database

import (
	"log"
	"os"

	"github.com/richcem/double-entry-ledgers-demo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection and performs migrations
func InitDB() {
	var err error

	// For development, we'll use SQLite
	DB, err = gorm.Open(sqlite.Open("ledger.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	// Auto-migrate models
	err = DB.AutoMigrate(
		&models.Account{},
		&models.Transaction{},
		&models.Entry{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		os.Exit(1)
	}

	log.Println("Database connected and migrated successfully")
}
