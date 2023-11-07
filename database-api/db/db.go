package db

import (
	"fmt"

	"database-api/config"
	geographicEntities "database-api/domains/geographic/entities"
	userEntities "database-api/domains/users/entities"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// func runMigrations() {
// 	dir := "db/migrations"
// 	command := "up"
// 	if err := goose.Run(command, dbHandle, dir); err != nil {
// 		log.Fatalf("goose run: %v", err)
// 	}
// }

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

	if err != nil {
		return nil, err
	}

	log.Println("Empezando automigración de entidades")

	models := append(userEntities.Models, geographicEntities.Models...)
	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}
	log.Println("Migración finalizada con éxito")

	return db, nil
}

func GetDBInstance() *gorm.DB {
	return db
}
