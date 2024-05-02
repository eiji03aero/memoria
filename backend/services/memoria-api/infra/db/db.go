package db

import (
	"fmt"

	"memoria-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Tokyo",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	return
}
