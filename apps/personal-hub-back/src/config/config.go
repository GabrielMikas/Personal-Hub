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
