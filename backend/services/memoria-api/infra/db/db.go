package db

import (
	"fmt"

	"memoria-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type NewDTO struct{}

func New() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Tokyo",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)

	logLevel := logger.Silent
	if config.DBLogLevel == "info" {
		logLevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return
	}

	return
}
