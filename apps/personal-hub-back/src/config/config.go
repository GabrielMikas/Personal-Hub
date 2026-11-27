package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	conn, err := initializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing postgres: %w", err)
	}
	db = conn
	return nil
}

func GetDB() *gorm.DB {
	return db
}

// SetDB replaces the global database connection. It exists so tests can
// inject an isolated database instead of the development/production one.
func SetDB(conn *gorm.DB) {
	db = conn
}
