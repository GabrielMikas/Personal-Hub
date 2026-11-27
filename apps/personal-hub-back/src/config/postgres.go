package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		// Not fatal: variables may come from the real environment.
		fmt.Println("Error loading environment variables:", err)
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, errors.New("DB_DSN not set")
	}

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if err := Migrate(conn); err != nil {
		return nil, err
	}
	return conn, nil
}

// Migrate creates or updates every table the application uses.
func Migrate(conn *gorm.DB) error {
	if err := conn.AutoMigrate(
		&hobbies_schemas.Order{},
		&hobbies_schemas.Card{},
		&hobbies_schemas.Reimbursement{},
	); err != nil {
		return fmt.Errorf("running migrations: %w", err)
	}
	return nil
}
