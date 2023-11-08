package db

import (
	"fmt"

	"database-api/config"
	geographicEntities "database-api/domains/geographic/entities"
	userEntities "database-api/domains/users/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DatabaseHost,
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort,
		config.DatabaseSslMode,
		config.TimeZone)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = userEntities.AutoMigrateEntities(db)
	if err != nil {
		return nil, err
	}
	err = geographicEntities.AutoMigrateEntities(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDBInstance() *gorm.DB {
	return db
}
