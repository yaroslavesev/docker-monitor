package db

import (
	"backend/internal/models"
)

// Migrations

func RunMigrations() error {
	err := DB.AutoMigrate(&models.Container{})
	return err
}
