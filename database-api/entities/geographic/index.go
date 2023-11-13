package entities

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrateEntities(db *gorm.DB) error {
	log.Println("Starting automigration of entities of geographic domain")

	err := db.AutoMigrate(
		&Country{},
		&Region{},
		&City{},
	)
	if err != nil {
		return err
	}
	log.Println("Succesful automigration of geographic domain")
	return nil
}
