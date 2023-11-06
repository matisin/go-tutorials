package db

import (
	"fmt"

	"database-api/config"
	userEntities "database-api/domains/users/entities"

	"log"

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

	log.Println("Empezando migración de entidades")
	err = db.AutoMigrate(
		&userEntities.Module{},
		&userEntities.Action{},
		&userEntities.Profile{},
		&userEntities.User{},
		&userEntities.Session{},
		&userEntities.Audit{},
		&userEntities.Group{},
		&userEntities.Permission{},
		&userEntities.Resource{},
	)
	if err != nil {
		return nil, err
	}
	log.Println("Migración finalizada con éxito")

	return db, nil
}

func GetDBInstance() *gorm.DB {
	return db
}
