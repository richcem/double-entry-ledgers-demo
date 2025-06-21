package database

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/richcem/double-entry-ledgers-demo/config"
	"github.com/richcem/double-entry-ledgers-demo/migrations"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection and performs migrations
func InitDB() {
	var err error

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		os.Exit(1)
	}

	// For development, we'll use SQLite with pure Go driver
	// Using glebarez/sqlite which is a pure Go driver
	DB, err = gorm.Open(sqlite.Open(cfg.DBName+"?cache=shared&_fk=1"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	// Run migrations using gormigrate
	if err := migrations.Migrate(DB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		os.Exit(1)
	}

	log.Println("Database connected and migrated successfully")
}
