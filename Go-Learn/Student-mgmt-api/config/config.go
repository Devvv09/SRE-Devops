package config

import (
	"fmt"
	"log"
	"students-mgmt-api/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	c, err := LoadEnviromentVariable()
	if err != nil {
		return fmt.Errorf("error while loading environment variables: %w", err)
	}

	// Corrected connection string format
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHOST, c.DBPORT, viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"))

	// Open the database connection
	DB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	
	// Run AutoMigrate to apply schema changes
	if err := DB.AutoMigrate(&models.Student{}); err != nil {
		log.Printf("Error during migration: %v", err)
		return fmt.Errorf("migration failed: %w", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
	return nil
}
