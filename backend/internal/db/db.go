package db

import (
	"fmt"

	"backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection to DB

var DB *gorm.DB

func ConnectDB() error {
	host := config.GetDBHost()
	user := config.GetDBUser()
	password := config.GetDBPassword()
	dbName := config.GetDBName()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		host, user, password, dbName,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
