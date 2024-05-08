package db

import (
	"fmt"

	"memoria-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() (db *gorm.DB, err error) {
	dbName := config.DBName
	if config.Env == "test" {
		dbName = "memoria-test"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Tokyo",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, dbName,
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
